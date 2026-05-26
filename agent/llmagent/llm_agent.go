//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package llmagent provides an LLM agent implementation.
package llmagent

import (
	"context"
	"reflect"
	"sync"

	sdktrace "go.opentelemetry.io/otel/trace"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/internal/flow"
	"trpc.group/trpc-go/trpc-agent-go/internal/flow/llmflow"
	"trpc.group/trpc-go/trpc-agent-go/internal/skillprofile"
	itelemetry "trpc.group/trpc-go/trpc-agent-go/internal/telemetry"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/planner"
	"trpc.group/trpc-go/trpc-agent-go/prompt"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
	toolskill "trpc.group/trpc-go/trpc-agent-go/tool/skill"
	toolworkspaceexec "trpc.group/trpc-go/trpc-agent-go/tool/workspaceexec"
)

// localruntimeFallback returns a simple local workspace executor used when
// no explicit executor is provided.
func defaultCodeExecutor() codeexecutor.CodeExecutor {
	_ = "STUB: not implemented"
	return *

	// LLMAgent is an agent that uses an LLM to generate responses.
	new(codeexecutor.CodeExecutor)
}

type LLMAgent struct {
	name                    string
	mu                      sync.RWMutex
	model                   model.Model
	models                  map[string]model.Model // Registered models for switching
	modelSelector           agent.ModelSelector
	description             string
	instruction             prompt.Text
	systemPrompt            prompt.Text
	modelInstructions       map[string]prompt.Text
	modelGlobalInstructions map[string]prompt.Text
	genConfig               model.GenerationConfig
	flow                    flow.Flow
	tools                   []tool.Tool     // All tools (user tools + framework tools)
	userToolNames           map[string]bool // Names of tools explicitly registered
	// via WithTools and WithToolSets.
	codeExecutor         codeexecutor.CodeExecutor
	workspaceRegistry    *codeexecutor.WorkspaceRegistry
	workspaceRegistries  map[workspaceRegistryKey]*codeexecutor.WorkspaceRegistry
	planner              planner.Planner
	subAgents            []agent.Agent // Sub-agents that can be delegated to
	agentCallbacks       *agent.Callbacks
	outputKey            string         // Key to store output in session state
	outputSchema         map[string]any // JSON schema for output validation
	inputSchema          map[string]any // JSON schema for input validation
	structuredOutput     *model.StructuredOutput
	structuredOutputType reflect.Type
	option               Options
}

const invalidOutputSchemaAwaitUserReply = "" +
	"Invalid LLMAgent configuration: if output_schema is set, " +
	"await_user_reply must be disabled"

// New creates a new LLMAgent with the given options.
func New(name string, opts ...Option) *LLMAgent { _ = "STUB: not implemented"; return nil }

// Apply function options.

// Wire agent-scoped extensions (WithExtensions) before any
// consumer of the callback fields runs. Order matters here:
//
//   - applyExtensionContributions must rewrite options.AgentCallbacks
//     / ModelCallbacks / ToolCallbacks BEFORE we copy them into
//     a.agentCallbacks, flowOpts.ModelCallbacks and the
//     FunctionCallResponseProcessor below — otherwise extension
//     hooks would silently no-op for this agent.
//   - extensionContributedTools is cached on Options so the
//     tool-surface builders can re-apply the same set after user
//     and framework tools (including per-invocation framework
//     tools such as transfer_to_agent). Extension callbacks are
//     static once merged; only tools need to flow through the
//     surface builders more than once.

// Validate output_schema configuration before registering tools.

// Extension-contributed tools (WithExtensions →
// extension.Registry.Tools) must be covered by the same
// guard. They are appended to the outbound tool surface
// alongside WithTools/WithToolSets, so allowing them
// through would let an extension silently break the
// "structured output ⇒ no callable tools" contract.
// Hook-only extensions (no Tools(...) call inside Register)
// leave this slice empty and remain compatible with
// WithOutputSchema.

// Register tools from both tools and toolsets, including knowledge search tool if provided.
// Also track which tools are user-registered (via WithTools) for filtering purposes.

// Initialize models map and determine the initial model.

// Construct the agent first so request processors can access dynamic getters.

// Prepare request processors in the correct order, wiring dynamic getters.

// Prepare response processors.

// Add planning response processor if planner is configured.

// Add output response processor if output_key or output_schema is configured or structured output is requested.

// Configure default transfer message for direct sub-agent calls.
// Default behavior (when not configured): enabled with built-in default message.

// Explicitly configured via WithDefaultTransferMessage.

// Always install the transfer processor so dynamic sub-agent updates
// (for example, via SubAgentSetter) can enable transfer_to_agent later.

// Create flow with the provided processors and options.

// buildRequestProcessors constructs the request processors in the required order.
func buildRequestProcessorsWithAgent(a *LLMAgent, options *Options) []flow.RequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

// 1. Basic processor - handles generation config.

// 2. Planning processor - handles planning instructions if planner is configured.

// 3. Instruction processor - adds instruction content and system prompt.

// Fallback injection for structured output when the provider doesn't enforce JSON Schema natively.

// static value unused when resolver is present
// static value unused when resolver is present

// 4. Identity processor - sets agent identity.

// 5. Skills processor - injects skill overview and loaded contents
// when a skills repository is configured statically or resolved at
// invocation time. This ensures the model sees available skills
// (names/descriptions) and any loaded SKILL.md/doc texts before
// deciding on tool calls.

// 6. Workspace exec processor - injects executor/workspace guidance
// when the current invocation exposes workspace_exec capability.

// 7. Content processor - appends conversation/context history.

// 8. On-demand session processor - injects a small overview for
// session_search / session_load when enabled.

// 9. Post-tool processor - injects dynamic prompt after tool results.

// 10. Skills tool result processor - materializes loaded skill content
// into tool result messages.

// 11. Time processor - adds current time information if enabled.
// Moved after content processor to avoid invalidating system message cache.
// Time information changes frequently, so placing it last allows previous
// stable content (instructions, identity, skills, history) to be cached.

func hasStaticOutputResponseProcessor(options *Options) bool {
	_ = "STUB: not implemented"
	return false
}

func hasInvocationStructuredOutput(ctx context.Context, invocation *agent.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}

func appendPostToolProcessor(options *Options, requestProcessors []flow.RequestProcessor) []flow.RequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

func appendOnDemandSessionProcessor(options *Options, requestProcessors []flow.RequestProcessor) []flow.RequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

func appendSkillsToolResultProcessor(a *LLMAgent, options *Options, requestProcessors []flow.RequestProcessor) []flow.RequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

func appendTimeProcessor(options *Options, requestProcessors []flow.RequestProcessor) []flow.RequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

// buildRequestProcessors preserves the original helper signature for tests and
// legacy callers. It constructs a temporary agent instance and forwards to
// buildRequestProcessorsWithAgent. Dynamic updates are not supported when using
// this legacy function; use New() which wires the real agent for runtime getters.
func buildRequestProcessors(name string, options *Options) []flow.RequestProcessor {
	_ = "STUB: not implemented" // nolint:deadcode
	return nil
}

func newTextPrompt(template string) prompt.Text {
	_ = "STUB: not implemented"
	return *new(prompt.Text)
}

func cloneTextPromptMap(src map[string]string) map[string]prompt.Text {
	_ = "STUB: not implemented"
	return nil
}

func prepareSkillsRepository(options *Options) { _ = "STUB: not implemented"; return }

// applySkillsExecutorFallback auto-wires a local code executor when the
// caller enabled skills via WithSkills but did not provide an executor.
// This preserves the zero-config upgrade path (WithSkills alone should
// keep working) while still letting callers fully opt out.
//
// The fallback is intentionally skipped when:
//   - an executor was already configured via WithCodeExecutor,
//   - the caller used WithAllowedSkillTools to drive fine-grained tool
//     selection (they are being explicit about what they want), or
//   - the caller explicitly selected SkillToolProfileKnowledgeOnly,
//     which is the opt-out signal for "no convenience execution wiring
//     from the framework".
//
// Note: the distinction between the unconfigured default and an
// explicit KnowledgeOnly profile is intentional; both normalize to the
// same built-in skill tool set, but only an explicit opt-in disables
// the fallback.
//
// Scope of the fallback: the auto-injected CodeExecutor exists to power
// execution tools such as workspace_exec. It must not silently expand
// the agent's execution surface to also auto-execute fenced code from
// assistant replies. Therefore, when this path injects an executor and
// the caller has NOT explicitly configured
// WithEnableCodeExecutionResponseProcessor, the function also disables
// EnableCodeExecutionResponseProcessor. Callers who explicitly set
// that option (true or false) keep their configured value.
func applySkillsExecutorFallback(options *Options) { _ = "STUB: not implemented"; return }

// initializeModels initializes the models map and determines the initial
// model based on WithModel and WithModels options.
func initializeModels(options *Options) (model.Model, map[string]model.Model) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

// Case 1: No models configured at all.

// Case 2: Only WithModel is set, no WithModels.

// Case 3: WithModels is set (with or without WithModel).

// If WithModel is also set, use it as the initial model.

// Check if the model is already in the models map.

// If not found, add it with the default name.

// WithModels is set but WithModel is not, use the first model from map.
// Note: map iteration order is not guaranteed.

// Should not reach here, but return nil for safety.

// registerTools assembles the full tool list for an LLMAgent from user-
// registered tools, knowledge tools, skill tools, and workspace tools.
//
// existingReg, when non-nil, is a WorkspaceRegistry carried over from a
// previous call (e.g. agent construction or tool refresh). Reusing it
// ensures that successive invocations share the same workspace cache
// instead of creating duplicate directories. The returned registry
// should be stored on the agent for future reuse.
func registerTools(
	options *Options,
	existingReg *codeexecutor.WorkspaceRegistry,
) ([]tool.Tool, map[string]bool, *codeexecutor.WorkspaceRegistry) {
	_ = "STUB: not implemented"
	// Step 1: collect user-registered tools (WithTools + WithToolSets)
	// and knowledge search tools.
	return nil, nil, nil
}

// Step 2: determine workspace registry and skill_run tool based on
// which capabilities the caller configured.
//
// Three scenarios:
//   a) skills + executor → shared registry for both skill_run and
//      workspace_exec; create one only if not inherited.
//   b) skills only (no executor) → skill_run gets nil registry and
//      builds its own internally; workspace_exec is not wired.
//   c) executor only (no skills) → registry for workspace_exec;
//      no skill_run needed.

// Step 3: wire workspace_exec (and session companions) when an
// executor with EngineProvider is available.

// Step 4: wire skill tools (skill_load, skill_run, etc.) when a
// skill repository is configured.

func collectUserToolNames(tools []tool.Tool) map[string]bool { _ = "STUB: not implemented"; return nil }

func appendStaticToolSetTools(
	allTools []tool.Tool,
	userToolNames map[string]bool,
	options *Options,
) ([]tool.Tool, map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendKnowledgeTools(
	allTools []tool.Tool,
	options *Options,
) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

func appendSkillTools(
	allTools []tool.Tool,
	options *Options,
	runTool *toolskill.RunTool,
) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

func appendSkillToolsWithRepo(
	allTools []tool.Tool,
	options *Options,
	repo skill.Repository,
	reg *codeexecutor.WorkspaceRegistry,
	runTool *toolskill.RunTool,
) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

func appendSkillToolsWithRepoAndFlags(
	allTools []tool.Tool,
	options *Options,
	repo skill.Repository,
	reg *codeexecutor.WorkspaceRegistry,
	runTool *toolskill.RunTool,
	exec codeexecutor.CodeExecutor,
	skillFlags skillprofile.Flags,
) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

func mustResolveSkillToolFlags(options *Options) skillprofile.Flags {
	_ = "STUB: not implemented"
	return *new(skillprofile.Flags)
}

func mustResolveSkillToolFlagsWithExecutor(
	options *Options,
	exec codeexecutor.CodeExecutor,
) skillprofile.Flags {
	_ = "STUB: not implemented"
	return *new(skillprofile.Flags)
}

func appendWorkspaceExecTool(
	allTools []tool.Tool,
	options *Options,
	reg *codeexecutor.WorkspaceRegistry,
	inv *agent.Invocation,
) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

// appendWorkspaceExecToolWithExecutor wires workspace_exec and its
// companion tools into allTools.
//
// loadedSkillsRepo is the effective skill repository for the current
// invocation. Callers on the invocation-scoped path pass the result
// of skillRepositoryForInvocation so that surface-patch repo
// overrides propagate into workspace_exec's loaded-skills reconcile;
// static callers (agent construction time) pass
// options.skillsRepository because no invocation context exists yet.
// Passing nil disables loaded-skills reconcile for this ExecTool.
func appendWorkspaceExecToolWithExecutor(
	allTools []tool.Tool,
	exec codeexecutor.CodeExecutor,
	enabled bool,
	sessional bool,
	reg *codeexecutor.WorkspaceRegistry,
	inv *agent.Invocation,
	options *Options,
	loadedSkillsRepo skill.Repository,
) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

func appendOnDemandSessionTools(
	allTools []tool.Tool,
	options *Options,
	inv *agent.Invocation,
) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

func buildWorkspaceRegistry() *codeexecutor.WorkspaceRegistry {
	_ = "STUB: not implemented"
	return nil
}

type workspaceRegistryKey struct {
	exec codeexecutor.CodeExecutor
}

func workspaceRegistryKeyForExecutor(
	exec codeexecutor.CodeExecutor,
) (workspaceRegistryKey, bool) {
	_ = "STUB: not implemented"
	return *new(workspaceRegistryKey), false
}

func workspaceRegistryMap(
	exec codeexecutor.CodeExecutor,
	reg *codeexecutor.WorkspaceRegistry,
) map[workspaceRegistryKey]*codeexecutor.WorkspaceRegistry {
	_ = "STUB: not implemented"
	return nil
}

func (a *LLMAgent) ensureWorkspaceRegistryForExecutor(
	exec codeexecutor.CodeExecutor,
) (*codeexecutor.WorkspaceRegistry, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (a *LLMAgent) workspaceRegistryForKeyLocked(
	key workspaceRegistryKey,
) *codeexecutor.WorkspaceRegistry {
	_ = "STUB: not implemented"
	return nil
}

func (a *LLMAgent) workspaceRegistryForInvocation(
	inv *agent.Invocation,
	exec codeexecutor.CodeExecutor,
) *codeexecutor.WorkspaceRegistry {
	_ = "STUB: not implemented"
	return nil
}

// workspacePrepOptions translates llmagent-level workspace options
// (WithWorkspaceBootstrap, invocation-scoped loaded-skills wiring,
// explicit disable switch) into the public workspaceexec options.
// The workspaceexec package owns reconciler construction and
// conversation-files wiring, so no internal workspaceprep type ever
// crosses this boundary.
//
// loadedSkillsRepo is the repository the caller has already resolved
// for the current invocation; see appendWorkspaceExecToolWithExecutor
// for how callers pick between invocation-scoped and agent-default
// repos. Passing a nil repo skips the loaded-skills wiring entirely,
// which is what we want when the agent has no skill support
// configured at all.
func workspacePrepOptions(
	opts *Options,
	loadedSkillsRepo skill.Repository,
) []func(*toolworkspaceexec.ExecTool) {
	_ = "STUB: not implemented"
	return nil
}

func buildSkillRunTool(
	options *Options,
	reg *codeexecutor.WorkspaceRegistry,
) *toolskill.RunTool {
	_ = "STUB: not implemented"
	return nil
}

func buildSkillRunToolWithRepo(
	options *Options,
	repo skill.Repository,
	reg *codeexecutor.WorkspaceRegistry,
	exec codeexecutor.CodeExecutor,
) *toolskill.RunTool {
	_ = "STUB: not implemented"
	return nil
}

// executorSupportsInteractive reports whether the effective engine
// behind the configured code executor exposes an
// InteractiveProgramRunner.  The check mirrors the fallback logic in
// RunTool.ensureEngine: when the executor does not implement
// EngineProvider (or returns a nil engine), the runtime falls back to
// a local engine which does support interactive sessions, so we
// return true in those cases.
func executorSupportsInteractive(options *Options) bool { _ = "STUB: not implemented"; return false }

func codeExecutorSupportsInteractive(exec codeexecutor.CodeExecutor) bool {
	_ = "STUB: not implemented"
	return false
}

// ensureEngine falls back to localexec which supports interactive.

// executorSupportsWorkspaceExec reports whether the agent has an
// explicit executor that exposes a live workspace engine suitable for
// generic workspace-side command execution. Unlike skill_run, this does
// not fall back to the local engine because that would silently move
// commands onto the agent host instead of the configured executor.
func executorSupportsWorkspaceExec(options *Options) bool { _ = "STUB: not implemented"; return false }

func codeExecutorSupportsWorkspaceExec(exec codeexecutor.CodeExecutor) bool {
	_ = "STUB: not implemented"
	return false
}

// codeExecutorHasLiveWorkspace reports whether exec exposes an Engine
// with Manager + FS — enough to drive Workspace.Collect / PutFiles /
// StageInputs / SaveArtifact. Looser than
// codeExecutorSupportsWorkspaceExec, which additionally requires
// Runner for the workspace_exec LLM tool. Workspace.RunProgram still
// surfaces a targeted "no program runner" error when Runner is nil,
// so executors that intentionally omit ProgramRunner can keep the
// other facade methods usable from callbacks.
func codeExecutorHasLiveWorkspace(exec codeexecutor.CodeExecutor) bool {
	_ = "STUB: not implemented"
	return false
}

// executorSupportsWorkspaceExecSessions reports whether workspace_exec can
// expose interactive session helpers such as workspace_write_stdin.
func executorSupportsWorkspaceExecSessions(options *Options) bool {
	_ = "STUB: not implemented"
	return false
}

func workspaceExecSurfaceEnabled(options *Options) bool { _ = "STUB: not implemented"; return false }

func codeExecutorSupportsWorkspaceExecSessions(
	exec codeexecutor.CodeExecutor,
) bool {
	_ = "STUB: not implemented"
	return false
}

// Run implements the agent.Agent interface.
// It executes the LLM agent flow and returns a channel of events.
func (a *LLMAgent) Run(ctx context.Context, invocation *agent.Invocation) (e <-chan *event.Event, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if this is a custom response error (early return)

// Handle actual errors

// executeAgentFlow executes the agent flow with before agent callbacks.
// Returns the updated context, event channel, and any error that occurred.
func (a *LLMAgent) executeAgentFlow(ctx context.Context, invocation *agent.Invocation) (context.Context, <-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

// Use the context from result if provided.

// Create a channel that returns the custom response and then closes.

// Create an event from the custom response.

// Use the underlying flow to execute the agent logic.

// haveCustomResponseError represents an early return due to a custom response from before agent callbacks.
// This is not an actual error but a signal to return early with the custom response.
type haveCustomResponseError struct {
	EventChan <-chan *event.Event
}

func (e *haveCustomResponseError) Error() string { _ = "STUB: not implemented"; return "" }

type baseModelResolution struct {
	Model              model.Model
	AllowAgentSelector bool
}

func (a *LLMAgent) resolveFlowBaseModel(inv *agent.Invocation) llmflow.ModelBaseResolution {
	_ = "STUB: not implemented"
	return *new(llmflow.ModelBaseResolution)
}

func (a *LLMAgent) resolveBaseModel(inv *agent.Invocation) baseModelResolution {
	_ = "STUB: not implemented"
	return *new(baseModelResolution)
}

// setupInvocation sets up the invocation.
func (a *LLMAgent) setupInvocation(invocation *agent.Invocation) {
	_ = "STUB: not implemented"
	// Set agent identity before resolving node-scoped surfaces.
	return
}

// Set the base model once for compatibility with existing callbacks.

// Lift run-scoped structured output into the current invocation once.

// Keep run-scoped values on RunOptions so clone-based handoffs can reuse the same output contract.

// Propagate per-agent safety limits into the invocation. These limits are
// evaluated by the Invocation helpers (IncLLMCallCount / IncToolIteration)
// and enforced at the flow layer. When the values are <= 0, the helpers
// treat them as "no limit", preserving existing behavior.

// withWorkspace installs a workspaceio.Workspace into ctx so that
// AgentCallbacks (BeforeAgent/AfterAgent/BeforeTool/AfterTool) observe
// the same invocation workspace as workspace_exec.
//
// The helper resolves the effective executor and registry the same
// way InvocationToolSurface does — honoring an invocation-scoped
// RunOptions.CodeExecutor override and only installing a facade when
// the executor actually supports workspace execution. This keeps the
// callback view and the workspace_exec tool view in lockstep across
// run-options overrides; otherwise callbacks could read or write a
// stale executor while the LLM tool uses the override.
func (a *LLMAgent) withWorkspace(
	ctx context.Context,
	inv *agent.Invocation,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Gate on "live workspace" (Manager + FS) rather than
// "workspace_exec" (Manager + FS + Runner). Executors that
// intentionally omit ProgramRunner can still serve Collect /
// PutFiles / StageInputs / SaveArtifact through the facade;
// RunProgram itself returns a targeted error when Runner is nil.

// wrapEventChannel wraps the event channel to apply after agent callbacks.
func (a *LLMAgent) wrapEventChannelWithTelemetry(
	ctx context.Context,
	invocation *agent.Invocation,
	originalChan <-chan *event.Event,
	span sdktrace.Span,
	tracker *itelemetry.InvokeAgentTracker,
	startedSpan bool,
) <-chan *event.Event {
	_ = "STUB: not implemented"
	// Create a new channel with the same capacity as the original channel
	return nil
}

// Forward all events from the original channel

func finalizeWrappedTelemetry(
	span sdktrace.Span,
	tracker *itelemetry.InvokeAgentTracker,
	fullRespEvent *event.Event,
	responseErrorType string,
	tokenUsage *itelemetry.TokenUsage,
	startedSpan bool,
	wrappedChan chan *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func resolveWrappedResponseErrorType(fullRespEvent *event.Event, responseErrorType string) string {
	_ = "STUB: not implemented"
	return ""
}

func recordWrappedEventTelemetry(
	evt *event.Event,
	tracker *itelemetry.InvokeAgentTracker,
	tokenUsage *itelemetry.TokenUsage,
	responseErrorType *string,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func addWrappedTokenUsage(tokenUsage *itelemetry.TokenUsage, usage *model.Usage) {
	_ = "STUB: not implemented"
	return
}

func (a *LLMAgent) runAfterAgentCallback(
	ctx context.Context,
	invocation *agent.Invocation,
	fullRespEvent *event.Event,
) (context.Context, *event.Event) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func wrappedAgentError(fullRespEvent *event.Event) error { _ = "STUB: not implemented"; return nil }

func wrappedAfterAgentEvent(
	invocation *agent.Invocation,
	result *agent.AfterAgentResult,
	callbackErr error,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// Info implements the agent.Agent interface.
// It returns the basic information about this agent.
func (a *LLMAgent) Info() agent.Info { _ = "STUB: not implemented"; return *new(agent.Info) }

// getAllToolsLocked builds the full tool list.
// It combines user tools with framework tools like transfer_to_agent
// under the caller's read lock. It always returns a fresh slice so
// callers can safely use it after releasing the lock without data
// races.
//
// This variant is used by methods that don't accept a context (for
// example, Tools()). It uses context.Background() when refreshing tools
// from ToolSets.
func (a *LLMAgent) getAllToolsLocked() []tool.Tool { _ = "STUB: not implemented"; return nil }

func (a *LLMAgent) getAllToolsLockedWithContext(
	ctx context.Context,
) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

// When RefreshToolSetsOnRun is enabled, rebuild tools from ToolSets
// on each call to keep ToolSet-provided tools in sync with their
// underlying dynamic source (for example, MCP ListTools).

// Tools implements the agent.Agent interface.
// It returns the list of tools available to the agent, including
// transfer tools.
func (a *LLMAgent) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

// SubAgents returns the list of sub-agents for this agent.
func (a *LLMAgent) SubAgents() []agent.Agent { _ = "STUB: not implemented"; return nil }

// FindSubAgent finds a sub-agent by name.
// Returns nil if no sub-agent with the given name is found.
func (a *LLMAgent) FindSubAgent(name string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

// UserTools returns the list of tools that were explicitly registered
// by the user via WithTools and WithToolSets options.
//
// User tools (can be filtered):
//   - Tools registered via WithTools
//   - Tools registered via WithToolSets
//
// Framework tools (never filtered, not included in this list):
//   - knowledge_search / agentic_knowledge_search (auto-added when
//     WithKnowledge is set)
//   - transfer_to_agent (auto-added when WithSubAgents is set)
//   - await_user_reply (auto-added when WithAwaitUserReplyTool(true))
//
// This method is used by the tool filtering logic to distinguish user
// tools from framework tools.
func (a *LLMAgent) UserTools() []tool.Tool { _ = "STUB: not implemented"; return nil }

// When ToolSets are static, user tool tracking is based on the
// snapshot captured at construction time.

// When ToolSets are refreshed on each run, user tools include:
//   - Tools from WithTools (tracked in userToolNames)
//   - Tools coming from ToolSets (wrapped as NamedTool).
// Framework tools (knowledge_search, transfer_to_agent,
// await_user_reply, etc.)
// remain excluded.

// FilterTools filters the list of tools based on the provided filter
// function.
func (a *LLMAgent) FilterTools(ctx context.Context) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

// CodeExecutor returns the code executor used by this agent.
// implements the agent.CodeExecutor interface.
// This allows the agent to execute code blocks in different environments.
func (a *LLMAgent) CodeExecutor() codeexecutor.CodeExecutor {
	_ = "STUB: not implemented"
	return *

	// SetSubAgents replaces the sub-agents for this agent in a
	// concurrency-safe way. This enables dynamic sub-agent discovery from
	// registries without recreating the agent instance.
	new(codeexecutor.CodeExecutor)
}

func (a *LLMAgent) SetSubAgents(subAgents []agent.Agent) { _ = "STUB: not implemented"; return }

// refreshToolsLocked recomputes the aggregated tool list and user tool
// tracking map from the current options. Caller must hold a.mu.Lock.
func (a *LLMAgent) refreshToolsLocked() { _ = "STUB: not implemented"; return }

// AddToolSet adds or replaces a tool set at runtime in a
// concurrency-safe way. If another ToolSet with the same Name()
// already exists, it will be replaced. Subsequent invocations see the
// updated tool list without recreating the agent.
func (a *LLMAgent) AddToolSet(toolSet tool.ToolSet) { _ = "STUB: not implemented"; return }

// RemoveToolSet removes all tool sets whose Name() matches the given
// name. It returns true if at least one ToolSet was removed. Tools
// from the removed tool sets will no longer be exposed on future
// invocations.
func (a *LLMAgent) RemoveToolSet(name string) bool { _ = "STUB: not implemented"; return false }

// SetToolSets replaces the agent ToolSets with the provided slice in a
// concurrency-safe way. Subsequent invocations will see tools from
// exactly these ToolSets plus framework tools (knowledge, skills).
func (a *LLMAgent) SetToolSets(toolSets []tool.ToolSet) { _ = "STUB: not implemented"; return }

// SetModel sets the model for this agent in a concurrency-safe way.
// This allows callers to manage multiple models externally and switch
// dynamically during runtime.
func (a *LLMAgent) SetModel(m model.Model) { _ = "STUB: not implemented"; return }

// SetModelByName switches the model by name in a concurrency-safe way.
// The model must be registered via WithModels option when creating the agent.
// Returns an error if the specified model name is not found.
func (a *LLMAgent) SetModelByName(modelName string) error { _ = "STUB: not implemented"; return nil }

// SetInstruction updates the agent's instruction at runtime in a concurrency-safe way.
// Subsequent requests will use the new instruction without recreating the agent.
func (a *LLMAgent) SetInstruction(instruction string) { _ = "STUB: not implemented"; return }

// SetGlobalInstruction updates the agent's global system prompt at runtime.
// This affects the system-level prompt prepended to requests.
func (a *LLMAgent) SetGlobalInstruction(systemPrompt string) { _ = "STUB: not implemented"; return }

// SetPrompts updates the instruction and global system prompt together.
// Subsequent requests observe the pair from a single agent lock.
func (a *LLMAgent) SetPrompts(
	instruction string,
	systemPrompt string,
) {
	_ = "STUB: not implemented"
	return
}

// SetModelInstructions updates the model-specific instruction overrides.
// Key: model.Info().Name, Value: instruction text.
func (a *LLMAgent) SetModelInstructions(instructions map[string]string) {
	_ = "STUB: not implemented"
	return
}

// SetModelGlobalInstructions updates the model-specific system prompt
// overrides.
// Key: model.Info().Name, Value: system prompt text.
func (a *LLMAgent) SetModelGlobalInstructions(prompts map[string]string) {
	_ = "STUB: not implemented"
	return
}

// getInstruction returns the current instruction with read lock.
func (a *LLMAgent) getInstruction() string { _ = "STUB: not implemented"; return "" }

// getSystemPrompt returns the current system prompt with read lock.
func (a *LLMAgent) getSystemPrompt() string { _ = "STUB: not implemented"; return "" }

func (a *LLMAgent) instructionForInvocation(inv *agent.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}

func (a *LLMAgent) instructionPromptForInvocation(inv *agent.Invocation) prompt.Text {
	_ = "STUB: not implemented"
	return *new(prompt.Text)
}

func (a *LLMAgent) systemPromptForInvocation(inv *agent.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}

func (a *LLMAgent) systemPromptTextForInvocation(inv *agent.Invocation) prompt.Text {
	_ = "STUB: not implemented"
	return *new(prompt.Text)
}
