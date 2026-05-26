//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package skill provides skill-related tools (function calls)
// for executing skill scripts without inlining code into prompts.
package skill

import (
	"context"
	"strings"
	"time"
	"unicode/utf8"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
	"trpc.group/trpc-go/trpc-agent-go/internal/skillstage"
	"trpc.group/trpc-go/trpc-agent-go/internal/workspaceinput"
	"trpc.group/trpc-go/trpc-agent-go/internal/workspacesession"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// RunTool lets the LLM execute commands inside a skill workspace.
// It stages the entire skill directory and runs a single command.
type RunTool struct {
	repo skill.Repository
	exec codeexecutor.CodeExecutor
	reg  *codeexecutor.WorkspaceRegistry
	wsr  *workspacesession.Resolver
	sst  *skillstage.Stager

	skillStager SkillStager

	allowedCmds map[string]struct{}
	deniedCmds  map[string]struct{}

	forceSaveArtifacts bool
	requireSkillLoaded bool
	outputLimits       RunOutputLimits
}

// RunOutputLimits controls how much inline text skill_run returns.
//
// These limits apply to stdout/stderr and primary_output selection only.
// Collected files still follow the workspace collector limits.
type RunOutputLimits struct {
	// StdoutStderrBytes is the per-stream inline limit for stdout/stderr.
	StdoutStderrBytes int
	// PrimaryOutputBytes is the largest text file eligible for
	// primary_output.
	PrimaryOutputBytes int
}

// SkillRunEnvProvider is an optional interface for skill repositories that
// want to inject environment variables into skill_run executions.
//
// Returned variables are merged into runInput.Env with least privilege:
// - Never overrides explicit tool-call env (runInput.Env)
// - Never overrides existing host env (os.LookupEnv)
// - Blocks known-dangerous env keys
type SkillRunEnvProvider interface {
	SkillRunEnv(
		ctx context.Context,
		skillName string,
	) (map[string]string, error)
}

// NewRunTool creates a new RunTool.
func NewRunTool(
	repo skill.Repository,
	exec codeexecutor.CodeExecutor,
	opts ...func(*RunTool),
) *RunTool {
	_ = "STUB: not implemented"
	return nil
}

const envAllowedCommands = "TRPC_AGENT_SKILL_RUN_ALLOWED_COMMANDS"
const envDeniedCommands = "TRPC_AGENT_SKILL_RUN_DENIED_COMMANDS"

const defaultSkillRunTimeout = 5 * time.Minute

const (
	defaultAutoExportPattern = codeexecutor.DirOut + "/**"
	defaultAutoExportMax     = 20
)

const (
	skillDirInputs = "inputs"
	skillDirVenv   = ".venv"
	skillDirBin    = "bin"
)

const (
	envPath            = "PATH"
	envVirtualEnv      = "VIRTUAL_ENV"
	envEditor          = "EDITOR"
	envVisual          = "VISUAL"
	posixPathListSep   = ":"
	windowsPathListSep = ";"
)

const (
	envLDPreload           = "LD_PRELOAD"
	envLDLibraryPath       = "LD_LIBRARY_PATH"
	envDYLDInsertLibraries = "DYLD_INSERT_LIBRARIES"
	envDYLDLibraryPath     = "DYLD_LIBRARY_PATH"
	envDYLDForceFlatNS     = "DYLD_FORCE_FLAT_NAMESPACE"
	envOpenSSLConf         = "OPENSSL_CONF"
)

const workspaceMetadataFileMode uint32 = 0o600

const (
	editorHelperDir     = ".trpc_agent"
	editorContentFile   = "editor_input.txt"
	editorScriptFile    = "editor_write.sh"
	editorScriptMissing = "editor wrapper: missing target file"
)

// WithAllowedCommands restricts skill_run to a single program execution
// whose command name is in the allowlist.
//
// When enabled, shell features (pipes, redirects, separators) are
// rejected and the command is executed without a shell.
func WithAllowedCommands(cmds ...string) func(*RunTool) { _ = "STUB: not implemented"; return nil }

// WithDeniedCommands rejects a single program execution whose command name
// matches the denylist.
//
// When enabled, shell features (pipes, redirects, separators) are rejected
// and the command is executed without a shell.
func WithDeniedCommands(cmds ...string) func(*RunTool) { _ = "STUB: not implemented"; return nil }

// WithForceSaveArtifacts forces skill_run to persist collected outputs
// via the artifact service when possible.
//
// It applies to both:
//   - legacy output_files + save_as_artifacts
//   - declarative outputs.save
func WithForceSaveArtifacts(enable bool) func(*RunTool) { _ = "STUB: not implemented"; return nil }

// WithRunOutputLimits customizes the inline stdout/stderr limit and the
// maximum file size eligible for primary_output.
//
// These limits do not change output_files collection limits. For large text
// payloads, prefer writing files under out/ and collecting them with
// output_files or outputs.
func WithRunOutputLimits(limits RunOutputLimits) func(*RunTool) {
	_ = "STUB: not implemented"
	return nil
}

// WithRequireSkillLoaded rejects skill_run calls unless the skill has been
// loaded via skill_load in the current session state.
//
// When enabled, models must call skill_load first to bring SKILL.md (and any
// selected docs) into context, reducing hallucinated commands/scripts.
func WithRequireSkillLoaded(enable bool) func(*RunTool) { _ = "STUB: not implemented"; return nil }

// WithWorkspaceRegistry reuses a caller-provided workspace registry so
// skill_run can share the same invocation workspace with other tools.
func WithWorkspaceRegistry(
	reg *codeexecutor.WorkspaceRegistry,
) func(*RunTool) {
	_ = "STUB: not implemented"
	return nil
}

func (t *RunTool) loadAllowedCommandsFromEnv() { _ = "STUB: not implemented"; return }

func (t *RunTool) setAllowedCommands(cmds []string) { _ = "STUB: not implemented"; return }

func (t *RunTool) loadDeniedCommandsFromEnv() { _ = "STUB: not implemented"; return }

func (t *RunTool) setDeniedCommands(cmds []string) { _ = "STUB: not implemented"; return }

func splitCommandList(raw string) []string { _ = "STUB: not implemented"; return nil }

// runInput is the JSON schema for skill_run.
type runInput struct {
	Skill          string            `json:"skill"`
	Command        string            `json:"command"`
	Cwd            string            `json:"cwd,omitempty"`
	Env            map[string]string `json:"env,omitempty"`
	Stdin          string            `json:"stdin,omitempty"`
	EditorText     string            `json:"editor_text,omitempty"`
	OutputFiles    []string          `json:"output_files,omitempty"`
	Timeout        int               `json:"timeout,omitempty"`
	SaveArtifacts  bool              `json:"save_as_artifacts,omitempty"`
	OmitInline     bool              `json:"omit_inline_content,omitempty"`
	ArtifactPrefix string            `json:"artifact_prefix,omitempty"`

	Inputs  []codeexecutor.InputSpec `json:"inputs,omitempty"`
	Outputs *codeexecutor.OutputSpec `json:"outputs,omitempty"`
}

// runOutput is the structured result returned by skill_run.
type runOutput struct {
	StagedInputs  []stagedInput `json:"staged_inputs,omitempty"`
	OutputFiles   []runFile     `json:"output_files"`
	PrimaryOutput *runFile      `json:"primary_output,omitempty"`
	Stdout        string        `json:"stdout"`
	Stderr        string        `json:"stderr"`
	ExitCode      int           `json:"exit_code"`
	TimedOut      bool          `json:"timed_out"`
	Duration      int64         `json:"duration_ms"`
	ArtifactFiles []artifactRef `json:"artifact_files,omitempty"`
	Warnings      []string      `json:"warnings,omitempty"`
}

type stagedInput = workspaceinput.StagedInput

type runFile struct {
	codeexecutor.File
	Ref string `json:"ref,omitempty"`
}

type artifactRef struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

type artifactStateRef struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
	Ref     string `json:"ref"`
}

type skillRunArtifactsDelta struct {
	ToolCallID string             `json:"tool_call_id"`
	Artifacts  []artifactStateRef `json:"artifacts"`
}

// Declaration implements tool.Tool.
func (t *RunTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

func formatCommandPreview(cmds map[string]struct{}, max int) string {
	_ = "STUB: not implemented"
	return ""
}

// Call executes the run request.
func (t *RunTool) Call(
	ctx context.Context, args []byte,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

var _ tool.Tool = (*RunTool)(nil)
var _ tool.CallableTool = (*RunTool)(nil)

func isSkillLoadedInContext(ctx context.Context, name string) bool {
	_ = "STUB: not implemented"
	return false
}

// StateDelta returns a stable, replayable artifact ref list when skill_run
// persisted outputs via Artifact service.
//
// It is consumed by the flow to attach StateDelta onto tool.response events.
func (t *RunTool) StateDelta(
	toolCallID string,
	_ []byte,
	resultJSON []byte,
) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

var _ stateDeltaProvider = (*RunTool)(nil)

func (t *RunTool) applyArtifactSaveOverrides(
	ctx context.Context,
	in runInput,
) (runInput, bool, string) {
	_ = "STUB: not implemented"
	return *new(runInput), false, ""
}

func (t *RunTool) prepareWorkspaceForRun(
	ctx context.Context,
	in runInput,
) (
	codeexecutor.Engine,
	codeexecutor.Workspace,
	string,
	context.Context,
	[]stagedInput,
	[]string,
	error,
) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.Engine), *new(codeexecutor.Workspace), "", *new(context.Context), nil, nil, nil
}

func (t *RunTool) stageSkillForRun(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	name string,
) (SkillStageResult, error) {
	_ = "STUB: not implemented"
	return *new(SkillStageResult), nil
}

func (t *RunTool) buildRunOutput(
	ctx context.Context,
	rr codeexecutor.RunResult,
	autoFiles []codeexecutor.File,
	files []codeexecutor.File,
	manifest *codeexecutor.OutputManifest,
	in runInput,
	saveRequested bool,
	outputsSaveSkipReason string,
) (runOutput, error) {
	_ = "STUB: not implemented"
	return *new(runOutput), nil
}

const (
	userFileInputFromPrefix  = "user_message://"
	userFileInputModePut     = "put"
	userFileInputNameFmt     = "upload_%d"
	userFileInputDefaultName = "upload"

	userFileInputKeyFileIDPrefix = "file_id/"
	userFileInputKeySHA256Prefix = "sha256/"
	userFileInputHostPrefix      = "host://"

	userFileInputWarnPrefix     = "user file input:"
	userFileInputWarnMissingRef = userFileInputWarnPrefix +
		" missing bytes and file_id"
	userFileInputWarnNoDownloader = userFileInputWarnPrefix +
		" model does not support file download"
	userFileInputWarnArtifactNoService = userFileInputWarnPrefix +
		" artifact service is not configured"
)

func (t *RunTool) stageUserFileInputs(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
) ([]stagedInput, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func fileNameFromArtifactRef(fileID string) string { _ = "STUB: not implemented"; return "" }

func sanitizeUserFileName(name string) string { _ = "STUB: not implemented"; return "" }

func uniqueUserFileName(
	used map[string]struct{},
	existingTo map[string]struct{},
	name string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func userFileInputBytes(
	ctx context.Context,
	mdl model.Model,
	f model.File,
) ([]byte, string, string) {
	_ = "STUB: not implemented"
	return nil, "", ""
}

func userFileInputsFromMessage(msg model.Message) []model.File {
	_ = "STUB: not implemented"
	return nil
}

func userFileInputsFromSession(sess *session.Session) []model.File {
	_ = "STUB: not implemented"
	return nil
}

func (t *RunTool) autoExportWorkspaceOut(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	in runInput,
) []codeexecutor.File {
	_ = "STUB: not implemented"
	return nil
}

// parseRunArgs validates and decodes input args.
func (t *RunTool) parseRunArgs(args []byte) (runInput, error) {
	_ = "STUB: not implemented"
	return *new(runInput), nil
}

func normalizeRunInput(in *runInput) { _ = "STUB: not implemented"; return }

func normalizeInputTo(to string) string { _ = "STUB: not implemented"; return "" }

// ensureEngine gets engine from executor or builds a local one.
func (t *RunTool) ensureEngine() codeexecutor.Engine {
	_ = "STUB: not implemented"
	return *new(codeexecutor.Engine)
}

func (t *RunTool) createWorkspace(
	ctx context.Context, eng codeexecutor.Engine, name string,
) (codeexecutor.Workspace, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.Workspace), nil
}

// stageSkill materializes the skill source under skills/<name>.
// As of the workspaceprep migration, skillstage.StageSkill defaults
// to a writable working copy. skill_run inherits the same writable
// semantics during its deprecation period; this matches the
// workspace_exec / reconciler path and avoids carrying two skill
// staging modes in parallel. If a future caller specifically needs
// the legacy read-only tree, it should switch to the lower-level
// StageSkillWithOptions API.
func (t *RunTool) stageSkill(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	root string,
	name string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *RunTool) loadWorkspaceMetadata(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
) (codeexecutor.WorkspaceMetadata, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.WorkspaceMetadata), nil
}

func (t *RunTool) saveWorkspaceMetadata(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	md codeexecutor.WorkspaceMetadata,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *RunTool) skillLinksPresent(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	name string,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t *RunTool) removeWorkspacePath(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	rel string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// shellQuote wraps a string for safe single-quoted usage in a
// POSIX shell. It escapes embedded single quotes by closing,
// inserting an escaped quote, and reopening.
func shellQuote(s string) string { _ = "STUB: not implemented"; return "" }

const (
	envVarPrefix = "$"
	envVarLBrace = "${"
	envVarRBrace = "}"
)

func hasEnvPrefix(s string, name string) bool { _ = "STUB: not implemented"; return false }

func isWorkspaceEnvPath(s string) bool { _ = "STUB: not implemented"; return false }

func isAllowedWorkspacePath(rel string) bool { _ = "STUB: not implemented"; return false }

func sanitizeWorkspaceRelPath(rel string, fallback string) string {
	_ = "STUB: not implemented"
	return ""
}

func resolveCWD(cwd string, name string) string {
	_ = "STUB: not implemented"
	// Default: run at the skill root. Relative cwd resolves under the
	// skill root. "$WORK_DIR" style paths resolve to workspace-relative
	// roots. Absolute paths are treated as workspace-absolute and must
	// start with known workspace dirs like "/skills" or "/work".
	return ""
}

// filepathBase returns the last element of a path, trimming trailing
// separators. It avoids importing path/filepath at top-level.
func filepathBase(p string) string { _ = "STUB: not implemented"; return "" }

func (t *RunTool) runProgram(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	skillRoot string,
	cwd string,
	in runInput,
) (codeexecutor.RunResult, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.RunResult), nil
}

func (t *RunTool) buildRunProgramSpec(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	skillRoot string,
	cwd string,
	in runInput,
) (codeexecutor.RunProgramSpec, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.RunProgramSpec), nil
}

func venvRelPaths(cwd string, skillRoot string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func skillLocalRelPaths(
	cwd string,
	skillRoot string,
) (string, string, string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

func (t *RunTool) maybeInjectSkillEnv(
	ctx context.Context,
	skillName string,
	env map[string]string,
) {
	_ = "STUB: not implemented"
	return
}

func isBlockedSkillEnvKey(key string) bool { _ = "STUB: not implemented"; return false }

func isValidEnvVarName(key string) bool { _ = "STUB: not implemented"; return false }

func slashRel(base string, target string) string { _ = "STUB: not implemented"; return "" }

func injectVenvEnv(env map[string]string, venv string, venvBin string) {
	_ = "STUB: not implemented"
	return
}

func injectSkillLocalEnv(
	env map[string]string,
	venv string,
	venvBin string,
	skillBin string,
) {
	_ = "STUB: not implemented"
	return
}

func injectSkillLocalEnvWithSep(
	env map[string]string,
	venv string,
	venvBin string,
	skillBin string,
	sep string,
) {
	_ = "STUB: not implemented"
	return
}

func pathEnvKey(env map[string]string, sep string) string { _ = "STUB: not implemented"; return "" }

type pathListSeparatorProvider interface {
	PathListSeparator() string
}

func pathListSeparatorForEngine(eng codeexecutor.Engine) string {
	_ = "STUB: not implemented"
	return ""
}

func cloneStringMap(src map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

func (t *RunTool) prepareEditorEnv(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	env map[string]string,
	editorText string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func buildEditorWrapperScript(contentPath string) string { _ = "STUB: not implemented"; return "" }

func wrapWithVenvPrefix(cmd string, venv string, venvBin string) string {
	_ = "STUB: not implemented"
	return ""
}

func wrapWithSkillLocalPrefix(
	cmd string,
	venv string,
	venvBin string,
	skillBin string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func writeShellSkillPATH(
	sb *strings.Builder,
	venvBin string,
	skillBin string,
) {
	_ = "STUB: not implemented"
	return
}

func isBareCommandName(cmd string) bool { _ = "STUB: not implemented"; return false }

func matchSkillLocalCommand(
	list map[string]struct{},
	cmd string,
	venvBinRel string,
	skillBinRel string,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func skillLocalCommandPath(dirRel string, cmd string) string { _ = "STUB: not implemented"; return "" }

func forceExplicitRelativeCommand(rel string) string { _ = "STUB: not implemented"; return "" }

const (
	disallowedShellMeta = "\n\r;&|<>"

	errShellMetaFmt = "skill_run: shell metacharacter %q is not allowed " +
		"when command restrictions are enabled " +
		"(allowed_commands/denied_commands set). " +
		"Use a single executable with args only " +
		"(no redirects/pipes/chaining). " +
		"To allow shell syntax, clear " +
		"allowed_commands/denied_commands in the tool config"
)

func cmdInList(list map[string]struct{}, cmd string) bool { _ = "STUB: not implemented"; return false }

func normalizeCommandForList(cmd string) string { _ = "STUB: not implemented"; return "" }

func splitCommandLine(s string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// withArtifactContext returns a context augmented with artifact
// service and session info when available from the invocation.
func withArtifactContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// prepareOutputs collects files either through OutputSpec or legacy
// output_files patterns. It returns collected files and optional
// manifest.
func (t *RunTool) prepareOutputs(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	in runInput,
) ([]codeexecutor.File, *codeexecutor.OutputManifest, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func outputFilesFromManifest(
	manifest codeexecutor.OutputManifest,
) []codeexecutor.File {
	_ = "STUB: not implemented"
	return nil
}

// buildRunOutput converts a RunResult and files into runOutput.
func buildRunOutput(
	rr codeexecutor.RunResult, files []codeexecutor.File,
) runOutput {
	_ = "STUB: not implemented"
	return *new(runOutput)
}

func buildRunOutputWithLimits(
	rr codeexecutor.RunResult,
	files []codeexecutor.File,
	limits RunOutputLimits,
) runOutput {
	_ = "STUB: not implemented"
	return *new(runOutput)
}

type failedOutputFilterResult struct {
	files        []codeexecutor.File
	manifest     *codeexecutor.OutputManifest
	omittedNames []string
	warnings     []string
}

func filterFailedEmptyOutputs(
	rr codeexecutor.RunResult,
	files []codeexecutor.File,
	manifest *codeexecutor.OutputManifest,
) failedOutputFilterResult {
	_ = "STUB: not implemented"
	return *new(failedOutputFilterResult)
}

func failedRunResult(rr codeexecutor.RunResult) bool { _ = "STUB: not implemented"; return false }

func emptyCollectedFile(f codeexecutor.File) bool { _ = "STUB: not implemented"; return false }

func emptyCollectedFileRef(f codeexecutor.FileRef) bool { _ = "STUB: not implemented"; return false }

func appendFilteredFileName(
	names []string,
	seen map[string]struct{},
	name string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func filterFailedEmptyManifestFiles(
	manifest *codeexecutor.OutputManifest,
	seen map[string]struct{},
	omittedNames *[]string,
) *codeexecutor.OutputManifest {
	_ = "STUB: not implemented"
	return nil
}

const (
	defaultStdoutStderrBytes = 16 * 1024
	defaultPrimaryOutputSize = 32 * 1024
)

const (
	warnStdoutTruncated     = "stdout truncated"
	warnStderrTruncated     = "stderr truncated"
	warnPartialOutputCommit = "output files were collected, but " +
		"workspace metadata could not be updated"
	warnFailedRunEmptyOutputFiles = "empty output_files omitted " +
		"because command failed; shell redirections can create " +
		"empty files before execution fails"
)

func truncateOutput(s string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func truncateOutputWithLimit(s string, limit int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func toRunFiles(files []codeexecutor.File) []runFile { _ = "STUB: not implemented"; return nil }

func shouldInlineFileContent(f codeexecutor.File) bool { _ = "STUB: not implemented"; return false }

const maxTrimUTF8SuffixBytes = utf8.UTFMax - 1

func trimTruncatedUTF8TextFiles(files []codeexecutor.File) { _ = "STUB: not implemented"; return }

func validUTF8PrefixLen(s string) int { _ = "STUB: not implemented"; return 0 }

func selectPrimaryOutput(files []runFile) *runFile { _ = "STUB: not implemented"; return nil }

func selectPrimaryOutputWithLimit(
	files []runFile,
	limit int,
) *runFile {
	_ = "STUB: not implemented"
	return nil
}

func mergeAutoPrimaryOutput(files []codeexecutor.File, out *runOutput) {
	_ = "STUB: not implemented"
	return
}

func mergeAutoPrimaryOutputWithLimit(
	files []codeexecutor.File,
	out *runOutput,
	limit int,
) {
	_ = "STUB: not implemented"
	return
}

func defaultRunOutputLimits() RunOutputLimits {
	_ = "STUB: not implemented"
	return *new(RunOutputLimits)
}

func normalizeRunOutputLimits(limits RunOutputLimits) RunOutputLimits {
	_ = "STUB: not implemented"
	return *new(RunOutputLimits)
}

// attachArtifactsIfRequested saves files as artifacts when requested
// and optionally omits inline content in the returned files.
func (t *RunTool) attachArtifactsIfRequested(
	ctx context.Context,
	out *runOutput,
	files []codeexecutor.File,
	prefix string,
	save bool,
	omitInline bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Only act when caller requests artifact persistence.

// Best-effort behavior: keep output_files content since
// artifacts were not persisted anywhere.

const warnSaveArtifactsSkippedTmpl = "save_as_artifacts requested but " +
	"%s; outputs are not persisted"

const warnOutputsSaveSkippedTmpl = "outputs.save requested but " +
	"%s; outputs are not persisted"

const warnOutputFilesWorkspaceOnly = "output_files are workspace-" +
	"relative; prefer output_files content; when you need a " +
	"stable reference, use output_files[*].ref (workspace://...)"

const (
	reasonNoInvocation = "invocation is missing from context"
	reasonNoService    = "artifact service is not configured"
	reasonNoSession    = "session is missing from invocation"
	reasonNoSessionIDs = "session app/user/session IDs are missing"
)

func artifactSaveSkipReason(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func appendWarning(out *runOutput, reason string) { _ = "STUB: not implemented"; return }

func appendOutputsSaveWarning(out *runOutput, reason string) { _ = "STUB: not implemented"; return }

const warnOmitInlineNoFallback = "omit_inline_content requested but " +
	"invocation is missing; returning inline output_files"

func applyOmitInlineContent(
	ctx context.Context,
	out *runOutput,
	omit bool,
) {
	_ = "STUB: not implemented"
	return
}

func hasOmitInlineFallback(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func skillRunOutputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func stagedInputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func runFileSchema(desc string) *tool.Schema { _ = "STUB: not implemented"; return nil }

func artifactRefSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func inputSpecsSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func outputSpecSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

// mergeManifestArtifactRefs appends artifact refs derived from a
// manifest when inline files were not already saved.
func mergeManifestArtifactRefs(
	manifest *codeexecutor.OutputManifest, out *runOutput,
) {
	_ = "STUB: not implemented"
	return
}

func (t *RunTool) collectFiles(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	patterns []string,
) ([]codeexecutor.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *RunTool) saveArtifacts(
	ctx context.Context,
	files []codeexecutor.File,
	prefix string,
) ([]artifactRef, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
