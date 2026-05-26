//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package e2b

import (
	"context"
	"os"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
)

// Compile-time check that workspaceRuntime satisfies the expected interfaces.
var (
	_ codeexecutor.WorkspaceManager = (*workspaceRuntime)(nil)
	_ codeexecutor.ProgramRunner    = (*workspaceRuntime)(nil)
)

const (
	// Base directory inside the E2B sandbox where per-execution workspaces
	// are created. /tmp is writable in the default template.
	defaultSandboxRunBase = "/tmp/run"

	defaultCreateTimeout  = 15 * time.Second
	defaultRmTimeout      = 15 * time.Second
	defaultStageTimeout   = 60 * time.Second
	defaultCollectTimeout = 30 * time.Second
	defaultRunTimeout     = 30 * time.Second

	// Maximum bytes read back from the sandbox for a single file when
	// collecting outputs.
	maxReadSizeBytes = 4 * 1024 * 1024 // 4 MiB

	// Sentinels used to frame RunProgram stdout / stderr / exit code so
	// wrapper-script noise is stripped before returning to callers.
	sentinelStdoutBegin = "__E2B_STDOUT_BEGIN__"
	sentinelStdoutEnd   = "__E2B_STDOUT_END__"
	sentinelStderrBegin = "__E2B_STDERR_BEGIN__"
	sentinelStderrEnd   = "__E2B_STDERR_END__"
	sentinelExitPrefix  = "__E2B_EXITCODE__="

	metadataFileMode = 0o600
)

// workspaceRuntime implements WorkspaceManager / WorkspaceFS / ProgramRunner
// for the E2B sandbox.
type workspaceRuntime struct {
	ce  *CodeExecutor
	cfg runtimeConfig
}

type runtimeConfig struct {
	runBase string
}

func newWorkspaceRuntime(c *CodeExecutor) *workspaceRuntime { _ = "STUB: not implemented"; return nil }

// CreateWorkspace creates a per-execution directory inside the sandbox.
func (r *workspaceRuntime) CreateWorkspace(
	ctx context.Context,
	execID string,
	pol codeexecutor.WorkspacePolicy,
) (codeexecutor.Workspace, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.Workspace), nil
}

// Cleanup removes the workspace directory from the sandbox.
func (r *workspaceRuntime) Cleanup(
	ctx context.Context,
	ws codeexecutor.Workspace,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PutFiles writes files into the sandbox workspace. Content is shipped via
// base64 to stay binary-safe.
func (r *workspaceRuntime) PutFiles(
	ctx context.Context,
	ws codeexecutor.Workspace,
	files []codeexecutor.PutFile,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PutDirectory packs a host directory into tar.gz then extracts it in the
// sandbox under ws.Path/to.
func (r *workspaceRuntime) PutDirectory(
	ctx context.Context,
	ws codeexecutor.Workspace,
	hostPath string,
	to string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// StageDirectory stages a directory with options (ReadOnly + AllowMount).
func (r *workspaceRuntime) StageDirectory(
	ctx context.Context,
	ws codeexecutor.Workspace,
	src string,
	to string,
	opt codeexecutor.StageOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Collect returns files in the workspace that match the supplied globs.
func (r *workspaceRuntime) Collect(
	ctx context.Context,
	ws codeexecutor.Workspace,
	patterns []string,
) ([]codeexecutor.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StageInputs maps external inputs into the sandbox workspace.
func (r *workspaceRuntime) StageInputs(
	ctx context.Context,
	ws codeexecutor.Workspace,
	specs []codeexecutor.InputSpec,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *workspaceRuntime) stageInputsLocked(
	ctx context.Context,
	ws codeexecutor.Workspace,
	specs []codeexecutor.InputSpec,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *workspaceRuntime) stageInput(
	ctx context.Context,
	ws codeexecutor.Workspace,
	md codeexecutor.WorkspaceMetadata,
	sp codeexecutor.InputSpec,
	mode string,
	to string,
) (string, *int, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (r *workspaceRuntime) stageArtifactInput(
	ctx context.Context,
	md codeexecutor.WorkspaceMetadata,
	sp codeexecutor.InputSpec,
	to string,
	dest string,
) (string, *int, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// stageCopyInsideSandbox resolves a workspace:// or skill:// reference and
// copies/links the target into dest, using only commands inside the sandbox.
func (r *workspaceRuntime) stageCopyInsideSandbox(
	ctx context.Context,
	ws codeexecutor.Workspace,
	sp codeexecutor.InputSpec,
	scheme string,
	mode string,
	dest string,
	rootSub string,
) (string, *int, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func pinnedArtifactVersion(
	md codeexecutor.WorkspaceMetadata,
	name string,
	to string,
) *int {
	_ = "STUB: not implemented"
	return nil
}

// CollectOutputs applies sandbox-side globs and optionally persists artifacts.
func (r *workspaceRuntime) CollectOutputs(
	ctx context.Context,
	ws codeexecutor.Workspace,
	spec codeexecutor.OutputSpec,
) (codeexecutor.OutputManifest, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.OutputManifest), nil
}

// RunProgram runs an arbitrary command inside the sandbox workspace.
func (r *workspaceRuntime) RunProgram(
	ctx context.Context,
	ws codeexecutor.Workspace,
	spec codeexecutor.RunProgramSpec,
) (codeexecutor.RunResult, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.RunResult), nil
}

// buildRunWrapper produces a bash script that executes `inner` while framing
// stdout/stderr/exit-code with sentinels, so the driver can parse them out.
func buildRunWrapper(inner string) string { _ = "STUB: not implemented"; return "" }

// parseFramedOutput extracts the user's stdout/stderr and exit code from the
// framed streaming output produced by buildRunWrapper.
func parseFramedOutput(rawStdout, rawStderr string) (string, string, int) {
	_ = "STUB: not implemented"
	return "", "", 0
}

// extractBetween returns the text between begin and end sentinels. Surrounding
// newlines added by `echo` are trimmed.
func extractBetween(s, begin, end string) string { _ = "STUB: not implemented"; return "" }

// ExecuteInline writes each code block into the sandbox workspace and runs
// it, aggregating stdout/stderr from all blocks.
func (r *workspaceRuntime) ExecuteInline(
	ctx context.Context,
	execID string,
	blocks []codeexecutor.CodeBlock,
	timeout time.Duration,
) (codeexecutor.RunResult, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.RunResult), nil
}

// runBash runs a bash snippet in the sandbox and returns the captured
// stdout, stderr and exit code.
func (r *workspaceRuntime) runBash(
	ctx context.Context, script string, timeout time.Duration,
) (string, string, int, error) {
	_ = "STUB: not implemented"
	return "", "", 0, nil
}

// runBashStreaming is the low-level primitive: it invokes Sandbox.RunCode
// with LanguageBash and collects the combined stream outputs.
func (r *workspaceRuntime) runBashStreaming(
	ctx context.Context, script string, timeout time.Duration,
) (string, string, int, error) {
	_ = "STUB: not implemented"
	return "", "", 0, nil
}

func isTimeoutErr(err error) bool { _ = "STUB: not implemented"; return false }

func tarGzFromFiles(files []codeexecutor.PutFile) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tarGzFromDir(root string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Skip symlinks/devices/etc for simplicity.

func (r *workspaceRuntime) uploadTarGzAndExtract(
	ctx context.Context, dest string, data []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// writeBytesToSandbox writes arbitrary bytes to dest inside the sandbox.
func (r *workspaceRuntime) writeBytesToSandbox(
	ctx context.Context, dest string, data []byte, mode os.FileMode,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *workspaceRuntime) readFile(
	ctx context.Context, full string, limit int64,
) ([]byte, int64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// Echo size then base64 payload on stdout, framed so we can parse it
// unambiguously even when file content contains ASCII sentinels.

// Strip newlines inserted by `base64`.

// listFilesByGlob resolves the provided patterns inside the sandbox using
// bash's globstar semantics and returns absolute file paths.
func (r *workspaceRuntime) listFilesByGlob(
	ctx context.Context, wsPath string, patterns []string,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Resolve the workspace root so we can reject matches that escape it
// (e.g. when patterns contain "..", symlinks, or absolute paths). The
// resolved base is also emitted on stdout so the Go side can use the
// identical form for prefix comparison even if wsPath is itself a
// symlink.

// The first line carries the resolved workspace root used by the
// sandbox filter; fall back to wsPath if it is missing.

// Defence-in-depth: also filter on the Go side against both the
// caller-supplied wsPath and the sandbox-resolved base, in case
// the shell filter was ever bypassed. path.Clean collapses any
// ".." segments so a literal prefix like "/ws/../escape" can't
// slip through.

// pathUnder reports whether p is equal to base or nested below it. Both
// arguments are expected to be absolute POSIX paths.
func pathUnder(p, base string) bool { _ = "STUB: not implemented"; return false }

func (r *workspaceRuntime) loadWorkspaceMetadata(
	ctx context.Context, ws codeexecutor.Workspace,
) (codeexecutor.WorkspaceMetadata, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.WorkspaceMetadata), nil
}

// fall back to fresh metadata

func (r *workspaceRuntime) saveWorkspaceMetadata(
	ctx context.Context, ws codeexecutor.Workspace,
	md codeexecutor.WorkspaceMetadata,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *workspaceRuntime) cleanupMetadataTemp(
	ctx context.Context,
	tmpPath string,
	committed *bool,
) {
	_ = "STUB: not implemented"
	return
}

// Input scheme prefixes mirrored from the container runtime.
const (
	inputSchemeArtifact  = "artifact://"
	inputSchemeHost      = "host://"
	inputSchemeWorkspace = "workspace://"
	inputSchemeSkill     = "skill://"
)

func inputBase(from string) string { _ = "STUB: not implemented"; return "" }

func sanitize(s string) string { _ = "STUB: not implemented"; return "" }

func shellQuote(s string) string { _ = "STUB: not implemented"; return "" }
