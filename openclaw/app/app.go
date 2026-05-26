//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package app provides an OpenClaw-like runnable wiring for
// `trpc-agent-go`.
//
// It is designed to be imported by downstream distributions that want to
// inject internal-only plugins (channels, backends, tools) via anonymous
// imports.
package app

import (
	"context"
	"net/http"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/claudecode"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/model/openai"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/admin"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/channel"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/cron"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/debugrecorder"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/deps"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/gateway"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/memoryfile"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/octool"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/outbound"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/persona"
	ocskills "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/skills"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/subagentrun"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/uploads"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/registry"
	openclawsubagent "trpc.group/trpc-go/trpc-agent-go/openclaw/subagent"
)

const (
	appName = "openclaw"

	defaultHTTPAddr = ":8080"

	modeMock   = "mock"
	modeOpenAI = "openai"

	defaultOpenAIModel = "gpt-5"

	defaultSkillsDir        = "skills"
	defaultAgentsDir        = ".agents"
	defaultBundledSkillsDir = "bundled-skills"

	csvDelimiter = ","

	defaultDebugRecorderDir = "debug"

	defaultAgentName = "assistant"

	preambleOnlyFinalAnswerRule = "A preamble-only response such as " +
		"`I will ...`, `I'll ...`, `我先...`, or " +
		"`接下来...` is not a valid final answer. If " +
		"tool work is needed, the same assistant message " +
		"must include the tool call; if no tool is needed, " +
		"return the requested content or completed result " +
		"instead of an announcement. "
	substantiveSameTurnRule = "For writing, summarization, " +
		"recommendation, explanation, or analysis tasks, " +
		"produce the requested content in the same turn. "
	artifactCompletionRule = "For requests to create, write, send, " +
		"publish, upload, schedule, or update an artifact " +
		"or external resource, the turn is not complete " +
		"until you have performed the action and returned " +
		"the resulting link, id, file marker, or exact " +
		"blocker after recovery."
	openClawPostToolPrompt = "[OpenClaw Tool Result Prompt] " +
		"Treat tool results as mid-task state, not as " +
		"permission to stop. Compare the latest tool result " +
		"with the user's original/current request and keep " +
		"working autonomously until the requested content, " +
		"artifact, or external action is complete or blocked. " +
		"If the request is complete, return the concrete " +
		"user-facing result: link, id, file marker, sent " +
		"status, created document title, scheduled job id, " +
		"or the exact blocker. If work is still needed and " +
		"an available tool can advance, verify, or recover " +
		"the task, call that tool in the same assistant turn. " +
		"Do not answer only with what you will do next. Do " +
		"not stop at a plan, progress note, or tool-result " +
		"summary when a next tool call is available. Do not " +
		"ask for confirmation for an in-scope next step unless " +
		"it is destructive, expensive, or genuinely ambiguous. " +
		"Do not claim that a document, message, upload, " +
		"schedule, file, wiki page, or external resource was " +
		"created, sent, written, published, or updated unless " +
		"the latest tool result proves it. If a tool failed " +
		"or returned a partial result, recover with available " +
		"tools when there is a clear path: corrected " +
		"parameters, canonical ids, alternate lookup, retry, " +
		"or verification. Only return a blocker when no safe " +
		"next step remains. For files and media, return " +
		"`MEDIA:` or `MEDIA_DIR:` lines only after the " +
		"file or directory exists and is intended to be " +
		"sent. For docs and iWiki, return the link, id, " +
		"or title when available. " +
		"Keep final answers concise and user-facing; avoid " +
		"exposing tool/source/process details unless they are " +
		"the exact result or blocker."
	skillPreambleOnlyRule = "A preamble-only skill response is " +
		"invalid. "
	skillSameTurnToolRule = "If you say you will read, load, " +
		"use, write, create, send, or publish through a " +
		"skill, the same assistant message must include " +
		"the required tool call. Do not stop after " +
		"announcing the skill-backed next step. "

	defaultAgentInstruction = "You are a helpful assistant. " +
		"Keep replies concise, but make each reply " +
		"substantive enough to complete the user's request. " +
		"Act without asking for confirmation when the next " +
		"step is clearly in scope, cheap, and reversible. " +
		preambleOnlyFinalAnswerRule +
		substantiveSameTurnRule +
		artifactCompletionRule
	openClawSkillsGuidance = "Treat the skill overview below as the " +
		"skills available in this session. Each entry " +
		"includes a path to that skill's SKILL.md on disk. " +
		"This is a blocking requirement for matching " +
		"skills. " +
		"If the user names a skill, names a slash command, " +
		"or the task clearly matches a skill description, " +
		"you must use that skill in the same turn. Start " +
		"with one brief user-visible preamble about the " +
		"immediate next step, then call `skill_load` for " +
		"that skill right away in the same turn. That " +
		"brief preamble is part of acting immediately, " +
		"not a pause to ask what to do next. That preamble may " +
		"not be the whole reply. " +
		skillPreambleOnlyRule +
		skillSameTurnToolRule +
		"That preamble may " +
		"announce the immediate task, but do not use it " +
		"for substantive guidance, capability " +
		"disclaimers, or explanations about which " +
		"subsystem loads versus runs the skill. Never mention " +
		"reading, loading, or using a matching skill " +
		"unless you already called `skill_load` for it in " +
		"this turn. Never say that you could read or load " +
		"a matching skill later without actually doing it " +
		"first. Do not answer a matching skill task from " +
		"the short summary, prior knowledge, or partial " +
		"memory. Even if you think you already know the " +
		"answer, load `SKILL.md` first. Load `SKILL.md` " +
		"before giving substantive guidance or acting on " +
		"the workflow. When " +
		"`SKILL.md` references " +
		"relative paths, resolve them from the skill " +
		"directory first. Read only the supporting docs, " +
		"scripts, assets, examples, or templates you still " +
		"need. Do not respond with capability disclaimers " +
		"such as `I can read the skill` when you can load " +
		"it now. Announce the next step briefly and do it. " +
		"When the user asks you to add, teach, configure, " +
		"preserve, or reuse a durable capability, workflow, " +
		"integration, domain rule, team process, API, CLI, " +
		"MCP endpoint, document convention, or tool usage " +
		"pattern, or to remember an executable workflow or " +
		"integration, prefer creating or updating a local " +
		"skill over treating it as a one-off answer. For " +
		"lightweight facts, preferences, or simple standing " +
		"rules, use memory instead. " +
		"Use platform code and tools for stable safety " +
		"boundaries, secrets, permissions, file paths, " +
		"validation, and execution guarantees; use skill " +
		"context for evolving behavior, triggers, " +
		"constraints, examples, recovery paths, and domain " +
		"knowledge. If you create or update a skill, do not " +
		"stop after describing the idea: choose a writable " +
		"user-managed skill root, not bundled skills unless " +
		"explicitly asked to edit them, write the skill files, " +
		"avoid storing raw secrets, validate or inspect the skill, " +
		"refresh or reload skills when the runtime provides " +
		"that path, and then use the skill to complete the " +
		"current task. " +
		"Reuse bundled scripts, templates, and assets " +
		"when they already fit. If multiple skills match, " +
		"use the smallest set that covers the task. Keep " +
		"context small and avoid bulk-loading docs. Do not " +
		"invent commands, flags, auth steps, file layouts, " +
		"or workflows from a short summary or partial " +
		"memory. Keep exploring nearby runtime facts, " +
		"retries, and recovery paths yourself before asking " +
		"for more input. If a matching skill is missing, " +
		"unreadable, or still lacks a required external " +
		"input after reasonable local exploration, state " +
		"the issue briefly and continue with the best " +
		"fallback."
	openClawSkillLoadToolDescription = "Load a skill body and optional " +
		"docs. This is a blocking requirement when the user " +
		"names a listed skill, names a slash command, or the " +
		"task clearly matches a listed skill description. " +
		"Before the first matching load, start with one " +
		"brief user-visible preamble that says which skill " +
		"or skill docs you are reading next, then call this " +
		"tool right away in the same turn. Do not pause " +
		"after that preamble to ask what to do next, and do " +
		"not send the preamble as the whole reply. " +
		skillPreambleOnlyRule +
		skillSameTurnToolRule +
		"Do not " +
		"use that preamble for " +
		"capability disclaimers, implementation-split " +
		"explanations, or substantive guidance about the " +
		"task. Do not answer from a short skill summary, " +
		"prior knowledge, or partial memory when a matching " +
		"skill exists. Load `SKILL.md` first, then load " +
		"only the extra docs you still need."
	openClawToolingGuidance = "For common PDF, DOCX, text, CSV, " +
		"and spreadsheet uploads already in the chat, prefer " +
		"read_document or read_spreadsheet before falling back " +
		"to exec_command. " +
		"For questions about the active chat history, recent " +
		"turns, or who said something in the current session, use " +
		"conversation_history before searching long-term memory. " +
		"Only use long-term memory tools for facts that are not " +
		"available in the current session. " +
		"Do not call exec_command just to print OPENCLAW_* upload " +
		"vars or inspect recent upload metadata when a matching " +
		"chat file is already available. For other general local " +
		"shell work, use exec_command. For interactive follow-up " +
		"input, use " +
		"write_stdin and kill_session when needed. Use message " +
		"to send to the current chat or an explicit target. " +
		artifactCompletionRule + " " +
		"Use the available tool path to complete the request " +
		"in this turn. " +
		"Chat uploads are saved to stable host paths. For host " +
		"commands, prefer OPENCLAW_LAST_UPLOAD_PATH or " +
		"OPENCLAW_SESSION_UPLOADS_DIR, OPENCLAW_LAST_UPLOAD_HOST_REF, " +
		"OPENCLAW_LAST_UPLOAD_NAME, " +
		"OPENCLAW_LAST_UPLOAD_MIME, and " +
		"OPENCLAW_MEMORY_FILE, " +
		"OPENCLAW_RECENT_UPLOADS_JSON instead of guessing " +
		"attachment paths. For long-running work, independent " +
		"verification, or background work that can continue after " +
		"this turn, use subagents_spawn with mode=async. When a " +
		"subagent result is required before continuing, use " +
		"mode=sync. When the user must review the subagent result " +
		"before you continue, use mode=review, show the result, " +
		"and wait for the next user reply. Do not use subagents " +
		"for small, tightly-coupled steps, and do not spawn " +
		"nested subagents. " +
		"When a user follows up about a " +
		"recent upload in the current chat, assume they mean " +
		"that existing upload unless the reference is " +
		"genuinely ambiguous. Match by media kind first: " +
		"prefer OPENCLAW_LAST_PDF_PATH, " +
		"OPENCLAW_LAST_AUDIO_PATH, OPENCLAW_LAST_VIDEO_PATH, or " +
		"OPENCLAW_LAST_IMAGE_PATH when the request clearly targets " +
		"one of those kinds. Telegram voice notes count as audio, " +
		"video notes count as video, and documents with image/audio/" +
		"video MIME types still count as that media kind. If the " +
		"user replies to an earlier media message, treat that " +
		"replied media as " +
		"the default target unless they clearly ask for something " +
		"else. Do not ask the user to re-upload a file or provide " +
		"a local path when the recent upload context already lists " +
		"a matching upload for this chat. If the user wants a " +
		"derived file sent back in the current chat, send it with " +
		"message instead of asking which channel or delivery " +
		"method to use. For exec_command, do " +
		"not assume skill workspace paths like work/inputs. Do not " +
		"expose local host paths to the user; when acknowledging a " +
		"new upload, refer to it only by filename and media kind, " +
		"not by OPENCLAW_* vars or a machine path. If the channel " +
		"gives you an opaque placeholder filename, avoid surfacing " +
		"that raw placeholder to the user unless they explicitly " +
		"ask for the exact filename. Refer to uploads " +
		"and generated files by user-facing filenames, and use " +
		"OPENCLAW_LAST_*_NAME instead of basename(" +
		"OPENCLAW_LAST_*_PATH) when deriving output filenames, " +
		"because stored host paths may include internal dedupe " +
		"prefixes. Use message " +
		"with host refs when possible, or with local file " +
		"paths/artifact refs when needed, to send " +
		"PDFs, images, audio, or video back to the current chat " +
		"when needed instead of asking for chat_id or another " +
		"upload. Merely mentioning a filename in text does not " +
		"send it; call message with files when the user should " +
		"actually receive media or documents. If a command " +
		"returns media_files or media_dirs, call message with " +
		"those paths unless your final reply already includes " +
		"`MEDIA:` or `MEDIA_DIR:` lines for OpenClaw to " +
		"auto-attach and hide from the user. When exec_command " +
		"or write_stdin generates images that you need to inspect, " +
		"prefer printing `MEDIA:` / `MEDIA_DIR:` lines or the " +
		"absolute image paths on their own lines. OpenClaw can " +
		"reattach those generated images to the model for direct " +
		"visual inspection, so inspect the image before assuming " +
		"OCR failed. If you intentionally " +
		"use that directive path, keep the visible prose separate " +
		"from the `MEDIA:` lines. If a compatible audio reply " +
		"should arrive as a Telegram voice bubble instead of a " +
		"generic audio file, call message with as_voice=true or " +
		"include `[[audio_as_voice]]` in the final reply along " +
		"with the `MEDIA:` lines. If a command " +
		"produces multiple files in one " +
		"directory, send that directory or the matching files " +
		"directly with message instead of only describing their " +
		"paths. When you mention generated files in the final " +
		"reply, use concise filenames rather than local machine " +
		"paths, and ensure those filenames actually exist under " +
		"the current working directory or " +
		"OPENCLAW_SESSION_UPLOADS_DIR. Prefer writing derived " +
		"files under " +
		"OPENCLAW_SESSION_UPLOADS_DIR when you will send them " +
		"back to the user. OPENCLAW_MEMORY_FILE is a visible " +
		"MEMORY.md file for the current scope, not hidden " +
		"internal state. If the user asks what you remember or " +
		"asks to inspect that file, read it and quote or " +
		"summarize the relevant " +
		"lines. If the user explicitly says 'remember this' " +
		"or asks you to remember a durable fact, preference, " +
		"or workflow rule, update OPENCLAW_MEMORY_FILE with a " +
		"short bullet in the same turn. Use " +
		"OPENCLAW_MEMORY_FILE only for stable cross-session " +
		"facts, preferences, or working style. Do not store " +
		"secrets or large transcripts in that file. " +
		"If a memory file does not exist yet, you may create it " +
		"at that exact path. Prefer already installed local tools " +
		"for OCR, PDF, audio, image, and video work before " +
		"trying package installs or long downloads. " +
		"When creating a cron job from chat, omit channel and " +
		"target to send results back to the current chat by " +
		"default. When adding cron jobs, write the stored task " +
		"as a one-time execution instruction, not as another " +
		"scheduling request. Prefer concise, outcome-oriented " +
		"tasks over brittle shell transcripts unless exact " +
		"commands are truly required. Use cron for future or " +
		"recurring work."

	browserToolingGuidance = "For real browser automation, use " +
		"browser. Prefer browser snapshot plus act for page " +
		"interaction, use browser screenshot when visual " +
		"verification matters, and keep using the same targetId " +
		"after tabs or snapshot calls. When the user mentions " +
		"their current browser tab, relay, or extension attach " +
		"flow, use profile=\"chrome\" when that profile exists."

	agentTypeLLM        = "llm"
	agentTypeClaudeCode = "claude-code"

	openAIVariantAuto = "auto"

	defaultOpenAIVariant = openAIVariantAuto

	deepSeekAPIHost = "api.deepseek.com"
	qwenAPIHost     = "dashscope.aliyuncs.com"
	hunyuanAPIHost  = "api.hunyuan.cloud.tencent.com"

	openAIBaseURLEnvName = "OPENAI_BASE_URL"
	openAIModelEnvName   = "OPENAI_MODEL"

	errClaudeCodeAgentNoPrompts = "claude-code agent does not support " +
		"agent prompts"
)

// Main runs the OpenClaw-like CLI and returns an exit code.
//
// args should not include the program name.
func Main(args []string) int { _ = "STUB: not implemented"; return 0 }

// MainWithOptions runs the OpenClaw-like CLI with runtime options and returns
// an exit code.
//
// args should not include the program name.
func MainWithOptions(args []string, options ...RuntimeOption) int {
	_ = "STUB: not implemented"
	return 0
}

// RunWithOptions runs OpenClaw until ctx is canceled or the runtime exits.
func RunWithOptions(
	ctx context.Context,
	args []string,
	options ...RuntimeOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

type exitError struct {
	Code int
	Err  error
}

type startupLogLine struct {
	warn bool
	text string
}

func applyOpenClawToolDefaults(
	agentType string,
	opts *runOptions,
) {
	_ = "STUB: not implemented"
	return
}

func (e *exitError) Error() string { _ = "STUB: not implemented"; return "" }

// ExitCode returns the suggested process exit code for this error.
func (e *exitError) ExitCode() int { _ = "STUB: not implemented"; return 0 }

func shouldLogExitError(err error) bool { _ = "STUB: not implemented"; return false }

func logStartupLines(lines []startupLogLine) { _ = "STUB: not implemented"; return }

func runtimeStartupLines(
	opts runOptions,
	stateDir string,
	channels []channel.Channel,
	needsModel bool,
) []startupLogLine {
	_ = "STUB: not implemented"
	return nil
}

func configStartupSummary(configPath string) string { _ = "STUB: not implemented"; return "" }

func startupPathSummary(path string) string { _ = "STUB: not implemented"; return "" }

func channelStartupSummary(channels []channel.Channel) string { _ = "STUB: not implemented"; return "" }

func modelStartupSummary(
	opts runOptions,
	needsModel bool,
) string {
	_ = "STUB: not implemented"
	return ""
}

func gatewayStartupLines(
	httpAddr string,
	gwSrv *gateway.Server,
) []startupLogLine {
	_ = "STUB: not implemented"
	return nil
}

func adminStartupLines(
	preferredAddr string,
	binding *adminBinding,
) []startupLogLine {
	_ = "STUB: not implemented"
	return nil
}

// Runtime wires OpenClaw components without owning the HTTP listener.
//
// Downstream distributions can mount Gateway.Handler into any HTTP server
// implementation (including framework-managed servers) while reusing the
// default OpenClaw runner + channel wiring.
type Runtime struct {
	Gateway  Gateway
	A2A      A2ASurface
	Admin    AdminSurface
	Channels []channel.Channel
	prompts  *RuntimePromptController
	adminCfg *admin.Config
	appName  string
	session  session.Service
	subagent SubagentService

	runner            runner.Runner
	cronRunner        closeFunc
	sessionSvc        closeFunc
	memorySvc         closeFunc
	cronSvc           closeFunc
	subagentSvc       closeFunc
	skillsWatch       closeFunc
	toolSets          []tool.ToolSet
	telemetryShutdown func(context.Context) error
}

// Gateway provides the HTTP handler and routes served by OpenClaw.
type Gateway struct {
	Handler      http.Handler
	HealthPath   string
	MessagesPath string
	StatusPath   string
	CancelPath   string
}

type AdminSurface struct {
	Handler http.Handler
	Addr    string
	URL     string
}

type PromptSnapshot struct {
	Instruction  string
	SystemPrompt string
}

type RuntimePromptController struct {
	agent agent.Agent

	mu       sync.RWMutex
	snapshot PromptSnapshot
}

func newRuntimePromptController(
	agt agent.Agent,
	instruction string,
	systemPrompt string,
) *RuntimePromptController {
	_ = "STUB: not implemented"
	return nil
}

// PromptController exposes runtime prompt updates without changing
// Runtime's exported struct layout.
func (r *Runtime) PromptController() *RuntimePromptController {
	_ = "STUB: not implemented"
	return nil
}

func (r *Runtime) AppName() string { _ = "STUB: not implemented"; return "" }

func (r *Runtime) SessionService() session.Service {
	_ = "STUB: not implemented"
	return *new(session.Service)
}

// SubagentService is the OpenClaw subagent control-plane service exposed by
// Runtime.
type SubagentService interface {
	ListForUser(
		userID string,
		filter openclawsubagent.ListFilter,
	) []openclawsubagent.Run
	GetForUser(userID string, runID string) (*openclawsubagent.Run, error)
	CancelForUser(
		userID string,
		runID string,
	) (*openclawsubagent.Run, bool, error)
}

func (r *Runtime) SubagentService() SubagentService {
	_ = "STUB: not implemented"
	return *new(SubagentService)
}

func (r *Runtime) ConfigureAdmin(
	configure func(*admin.Config),
) {
	_ = "STUB: not implemented"
	return
}

// AddAdminOptions appends runtime-scoped admin options without changing
// Runtime's exported struct layout.
func (r *Runtime) AddAdminOptions(opts ...admin.Option) { _ = "STUB: not implemented"; return }

func (r *Runtime) applyAdminConfig(cfg admin.Config) { _ = "STUB: not implemented"; return }

func (c *RuntimePromptController) Snapshot() PromptSnapshot {
	_ = "STUB: not implemented"
	return *new(PromptSnapshot)
}

func (c *RuntimePromptController) SetInstruction(
	instruction string,
) {
	_ = "STUB: not implemented"
	return
}

func (c *RuntimePromptController) SetSystemPrompt(
	systemPrompt string,
) {
	_ = "STUB: not implemented"
	return
}

func (c *RuntimePromptController) SetPrompts(
	instruction string,
	systemPrompt string,
) {
	_ = "STUB: not implemented"
	return
}

// NewRuntime constructs an OpenClaw runtime based on CLI args / config file,
// but does not start an HTTP server.
func NewRuntime(
	ctx context.Context,
	args []string,
) (*Runtime, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewRuntimeWithOptions constructs an embedded OpenClaw runtime with options.
func NewRuntimeWithOptions(
	ctx context.Context,
	args []string,
	options ...RuntimeOption,
) (rt *Runtime, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close releases owned resources (session/memory services, toolsets, runner).
func (r *Runtime) Close() error { _ = "STUB: not implemented"; return nil }

func run(
	ctx context.Context,
	args []string,
	options ...RuntimeOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:gosec

//nolint:gosec

type closeFunc interface {
	Close() error
}

func closeSessionService(svc closeFunc) { _ = "STUB: not implemented"; return }

func closeMemoryService(svc closeFunc) { _ = "STUB: not implemented"; return }

func memoryServiceTools(svc memory.Service) []tool.Tool { _ = "STUB: not implemented"; return nil }

func fileMemoryStoreForBackend(
	backend string,
	store *memoryfile.Store,
) *memoryfile.Store {
	_ = "STUB: not implemented"
	return nil
}

func appendMemoryServiceRunnerOption(
	opts []runner.Option,
	svc memory.Service,
) []runner.Option {
	_ = "STUB: not implemented"
	return nil
}

func closeToolSets(sets []tool.ToolSet) { _ = "STUB: not implemented"; return }

func shutdownTelemetry(
	shutdown func(context.Context) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func shutdownTelemetryWithContext(
	ctx context.Context,
	shutdown func(context.Context) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func newCronRunner(
	appName string,
	ag agent.Agent,
	memSvc memory.Service,
	rlCfg *runner.RalphLoopConfig,
) runner.Runner {
	_ = "STUB: not implemented"
	return *new(runner.Runner)
}

func runtimeInstanceID(
	agentType string,
	opts runOptions,
	needsModel bool,
	stateDir string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func channelIDs(channels []channel.Channel) []string { _ = "STUB: not implemented"; return nil }

func listenURL(addr string) string { _ = "STUB: not implemented"; return "" }

func defaultOpenAIModelName() string { _ = "STUB: not implemented"; return "" }

func makeGatewayOptions(
	users []string,
	requireMention bool,
	mentionPatterns []string,
) []gateway.Option {
	_ = "STUB: not implemented"
	return nil
}

func normalizeAgentType(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func validateAgentRunOptions(agentType string, opts runOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func ralphLoopConfigFromRunOptions(
	opts runOptions,
) (*runner.RalphLoopConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseKVOverrides(items []string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseClaudeOutputFormat(
	raw string,
) (claudecode.OutputFormat, error) {
	_ = "STUB: not implemented"
	return *new(claudecode.OutputFormat), nil
}

func newClaudeCodeAgent(opts runOptions) (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func newAgent(
	mdl model.Model,
	cfg agentConfig,
	extraTools []tool.Tool,
	toolSets []tool.ToolSet,
) (agent.Agent, *ocskills.Repository, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil, nil
}

func buildOpenClawToolingGuidance(cfg agentConfig) string { _ = "STUB: not implemented"; return "" }

func buildOpenClawSkillsGuidance(cfg agentConfig) string { _ = "STUB: not implemented"; return "" }

func hasToolNamed(tools []tool.Tool, name string) bool { _ = "STUB: not implemented"; return false }

func toolsFromProviders(
	mdl model.Model,
	appName string,
	stateDir string,
	specs []pluginSpec,
) ([]tool.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toolSetsFromProviders(
	mdl model.Model,
	appName string,
	stateDir string,
	specs []pluginSpec,
) ([]tool.ToolSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func channelsFromRegistry(
	ctx context.Context,
	gw registry.GatewayClient,
	appName string,
	stateDir string,
	allowUsers []string,
	specs []pluginSpec,
) ([]channel.Channel, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type agentConfig struct {
	AppName string

	AddSessionSummary                             bool
	EnableContextCompaction                       bool
	ContextCompactionOversizedToolResultMaxTokens int
	MaxHistoryRuns                                int
	PreloadMemory                                 int
	GenerationConfig                              *model.GenerationConfig
	Instruction                                   string
	SystemPrompt                                  string

	SkillsRoot          string
	SkillsExtraDirs     []string
	SkillsDebug         bool
	SkillsAllowBundled  []string
	SkillConfigs        map[string]ocskills.SkillConfig
	SkillConfigKeys     []string
	SkillsWatch         bool
	SkillsWatchBundled  bool
	SkillsWatchDebounce time.Duration
	SkillsToolProfile   string
	SkillsLoadMode      string
	SkillsMaxLoaded     int
	SkillsToolResults   bool
	SkillsSkipFallback  bool
	SkillsToolingGuide  *string
	KnowledgesConfig    []knowledgeEntry

	StateDir string

	MemoryFileStore *memoryfile.Store

	EnableLocalExec bool

	EnableOpenClawTools  bool
	OpenClawToolingGuide *string
	EnableParallelTools  bool

	ToolProviders []pluginSpec

	ToolSets []pluginSpec

	RefreshToolSetsOnRun bool
}

type openClawToolsBundle struct {
	tools         []tool.Tool
	execMgr       *octool.Manager
	router        *outbound.Router
	cronTool      *cron.Tool
	subagentTools subagentrun.Tools
	deps          *deps.Report
}

type runtimeStores struct {
	uploads     *uploads.Store
	personas    *persona.Store
	memoryFiles *memoryfile.Store
}

func newRuntimeStores(stateDir string) (runtimeStores, error) {
	_ = "STUB: not implemented"
	return *new(runtimeStores), nil
}

func buildOpenClawTools(
	enabled bool,
	stateDir string,
	uploadStore *uploads.Store,
	memoryFileStore *memoryfile.Store,
) openClawToolsBundle {
	_ = "STUB: not implemented"
	return *new(openClawToolsBundle)
}

func resolveSkillRoots(cwd string, cfg agentConfig) []string { _ = "STUB: not implemented"; return nil }

func newSkillsWatchService(
	cwd string,
	cfg agentConfig,
	repo *ocskills.Repository,
) *ocskills.WatchService {
	_ = "STUB: not implemented"
	return nil
}

func resolveWorkspaceSkillsRoot(cwd, raw string) string { _ = "STUB: not implemented"; return "" }

func dirExists(path string) bool { _ = "STUB: not implemented"; return false }

func resolveBundledSkillsRoot(cwd, stateDir string) string { _ = "STUB: not implemented"; return "" }

func resolveStateDir(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func maybeEnableDebugRecorder(
	ctx context.Context,
	opts runOptions,
) (context.Context, *debugrecorder.Recorder, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func configFingerprint(parts ...string) string { _ = "STUB: not implemented"; return "" }

func newMockModel(_ registry.ModelSpec) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

func recordDebugOpenAIChatRequestJSON(
	ctx context.Context,
	raw []byte,
	marshalErr error,
) {
	_ = "STUB: not implemented"
	return
}

func newOpenAIModel(spec registry.ModelSpec) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

func modelFromOptions(opts runOptions) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

func parseOpenAIVariant(
	raw string,
	baseURL string,
) (openai.Variant, error) {
	_ = "STUB: not implemented"
	return *new(openai.Variant), nil
}

func inferOpenAIVariant(baseURL string) openai.Variant {
	_ = "STUB: not implemented"
	return *new(openai.Variant)
}

func openAIBaseURLHost(raw string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func splitCSV(input string) []string { _ = "STUB: not implemented"; return nil }

type echoModel struct {
	name string
}

func (m *echoModel) Info() model.Info { _ = "STUB: not implemented"; return *new(model.Info) }

func (m *echoModel) GenerateContent(
	ctx context.Context,
	req *model.Request,
) (<-chan *model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func lastUserText(req *model.Request) string { _ = "STUB: not implemented"; return "" }
