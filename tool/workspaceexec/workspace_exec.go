//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package workspaceexec exposes shared executor-workspace tools such as
// workspace_exec and workspace_save_artifact.
package workspaceexec

import (
	"context"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
	"trpc.group/trpc-go/trpc-agent-go/internal/workspaceprep"
	"trpc.group/trpc-go/trpc-agent-go/internal/workspacesession"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultWorkspaceExecTimeout = 5 * time.Minute
	defaultWorkspaceWriteYield  = 200
)

// ExecTool executes shell commands in the shared executor workspace.
type ExecTool struct {
	exec      codeexecutor.CodeExecutor
	reg       *codeexecutor.WorkspaceRegistry
	resolver  *workspacesession.Resolver
	sessional bool

	// providers contribute workspace requirements (loaded skills,
	// bootstrap files, conversation files, ...). When non-empty,
	// reconciler is always non-nil. These fields are populated only
	// through the WithWorkspaceBootstrap / WithLoadedSkills options
	// so internal workspaceprep types never leak into this tool's
	// public surface.
	providers              []workspaceprep.Provider
	reconciler             workspaceprep.Reconciler
	conversationFilesWired bool

	mu       sync.Mutex
	sessions map[string]*execSession
	ttl      time.Duration
	clock    func() time.Time
}

// WriteStdinTool sends additional stdin to a running workspace_exec session.
type WriteStdinTool struct {
	exec *ExecTool
}

// KillSessionTool terminates a running workspace_exec session.
type KillSessionTool struct {
	exec *ExecTool
}

type execSession struct {
	mu sync.Mutex

	proc        codeexecutor.ProgramSession
	exitedAt    time.Time
	finalized   bool
	finalizedAt time.Time
}

type execInput struct {
	Command       string            `json:"command"`
	Cwd           string            `json:"cwd,omitempty"`
	Env           map[string]string `json:"env,omitempty"`
	Stdin         string            `json:"stdin,omitempty"`
	YieldTimeMS   *int              `json:"yield_time_ms,omitempty"`
	YieldMs       *int              `json:"yieldMs,omitempty"`
	Background    bool              `json:"background,omitempty"`
	Timeout       int               `json:"timeout,omitempty"`
	TimeoutSec    *int              `json:"timeout_sec,omitempty"`
	TimeoutSecOld *int              `json:"timeoutSec,omitempty"`
	TTY           *bool             `json:"tty,omitempty"`
	PTY           *bool             `json:"pty,omitempty"`
}

type writeInput struct {
	SessionID     string `json:"session_id,omitempty"`
	SessionIDOld  string `json:"sessionId,omitempty"`
	Chars         string `json:"chars,omitempty"`
	YieldTimeMS   *int   `json:"yield_time_ms,omitempty"`
	YieldMs       *int   `json:"yieldMs,omitempty"`
	AppendNewline *bool  `json:"append_newline,omitempty"`
	Submit        *bool  `json:"submit,omitempty"`
}

type killInput struct {
	SessionID    string `json:"session_id,omitempty"`
	SessionIDOld string `json:"sessionId,omitempty"`
}

type execOutput struct {
	Status     string `json:"status"`
	Output     string `json:"output,omitempty"`
	ExitCode   *int   `json:"exit_code,omitempty"`
	SessionID  string `json:"session_id,omitempty"`
	Offset     int    `json:"offset"`
	NextOffset int    `json:"next_offset"`
}

type execRequest struct {
	background bool
	tty        bool
	yield      *int
	eng        codeexecutor.Engine
	ws         codeexecutor.Workspace
	spec       codeexecutor.RunProgramSpec
}

type killOutput struct {
	OK        bool   `json:"ok"`
	SessionID string `json:"session_id"`
	Status    string `json:"status"`
}

// NewExecTool creates a workspace_exec tool for the provided executor.
func NewExecTool(
	exec codeexecutor.CodeExecutor,
	opts ...func(*ExecTool),
) *ExecTool {
	_ = "STUB: not implemented"
	return nil
}

// NewWriteStdinTool creates a tool for continuing or polling a running session.
func NewWriteStdinTool(exec *ExecTool) *WriteStdinTool { _ = "STUB: not implemented"; return nil }

// NewKillSessionTool creates a tool for terminating a running session.
func NewKillSessionTool(exec *ExecTool) *KillSessionTool { _ = "STUB: not implemented"; return nil }

// WithWorkspaceRegistry reuses a caller-provided workspace registry so
// workspace_exec can share the same invocation workspace with other tools.
func WithWorkspaceRegistry(
	reg *codeexecutor.WorkspaceRegistry,
) func(*ExecTool) {
	_ = "STUB: not implemented"
	return nil
}

// WithWorkspaceBootstrap declares static files and one-shot commands
// that must exist/run in the workspace before workspace_exec runs
// user commands. The spec is converted into reconciler Requirements
// internally; idempotency and skip-on-fingerprint-match are handled
// by the framework. Files are staged first, then commands run, both
// in declaration order.
//
// A malformed spec panics during option application: silently
// dropping a partially-configured bootstrap would leave the agent
// running in a state the caller did not ask for.
//
// Passing an empty spec (no files, no commands) is a no-op.
func WithWorkspaceBootstrap(
	spec codeexecutor.WorkspaceBootstrapSpec,
) func(*ExecTool) {
	_ = "STUB: not implemented"
	return nil
}

// WithLoadedSkills wires the reconciler to materialize skills that
// have been recorded in session state via skill_load, using the
// supplied repository to resolve skill sources. Skills are staged
// into skills/<name> as writable working copies.
//
// Passing a nil repository is a no-op.
func WithLoadedSkills(repo skill.Repository) func(*ExecTool) { _ = "STUB: not implemented"; return nil }

// addPreparer records a provider and ensures the companion pieces
// (default reconciler, conversation-files provider) are installed.
// It is only used by the options above so that callers never see
// an internal workspaceprep type.
func (t *ExecTool) addPreparer(p workspaceprep.Provider) { _ = "STUB: not implemented"; return }

// toInternalBootstrapSpec bridges codeexecutor.WorkspaceBootstrapSpec
// (the stable public type) to workspaceprep.BootstrapSpec (the
// internal representation). Keeping the two struct families
// nominally distinct lets the public surface evolve independently of
// the reconciler internals.
func toInternalBootstrapSpec(
	in codeexecutor.WorkspaceBootstrapSpec,
) workspaceprep.BootstrapSpec {
	_ = "STUB: not implemented"
	return *new(workspaceprep.BootstrapSpec)
}

// Declaration returns the schema for workspace_exec.
func (t *ExecTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// Declaration returns the schema for workspace_write_stdin.
func (t *WriteStdinTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// Declaration returns the schema for workspace_kill_session.
func (t *KillSessionTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// Call executes workspace_exec once or starts a resumable session.
func (t *ExecTool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func parseExecInput(args []byte) (execInput, error) {
	_ = "STUB: not implemented"
	return *new(execInput), nil
}

func (t *ExecTool) prepareExec(
	ctx context.Context,
	in execInput,
) (execRequest, error) {
	_ = "STUB: not implemented"
	return *new(execRequest), nil
}

func (t *ExecTool) callNonSessional(
	ctx context.Context,
	req execRequest,
) (execOutput, error) {
	_ = "STUB: not implemented"
	return *new(execOutput), nil
}

func (t *ExecTool) callSessional(
	ctx context.Context,
	req execRequest,
) (execOutput, error) {
	_ = "STUB: not implemented"
	return *new(execOutput), nil
}

func runOneShot(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	spec codeexecutor.RunProgramSpec,
) (execOutput, error) {
	_ = "STUB: not implemented"
	return *new(execOutput), nil
}

func (t *ExecTool) startInteractive(
	ctx context.Context,
	req execRequest,
) (execOutput, error) {
	_ = "STUB: not implemented"
	return *new(execOutput), nil
}

func initialPoll(
	proc codeexecutor.ProgramSession,
	background bool,
	yield *int,
) codeexecutor.ProgramPoll {
	_ = "STUB: not implemented"
	return *new(codeexecutor.ProgramPoll)
}

// Call writes stdin to an interactive workspace_exec session or polls it.
func (t *WriteStdinTool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Call terminates a running workspace_exec session.
func (t *KillSessionTool) Call(_ context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// reconcileWorkspace converges the workspace to the desired state
// before executing a user command. When no provider is configured
// the function preserves the legacy behavior of staging conversation
// files inline; otherwise it delegates to the reconciler which
// collects Requirements from every provider and applies them in
// phase order (file -> skill -> command).
func (t *ExecTool) reconcileWorkspace(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *ExecTool) liveEngine() (codeexecutor.Engine, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.Engine), nil
}

func (t *ExecTool) putSession(id string, sess *execSession) { _ = "STUB: not implemented"; return }

func (t *ExecTool) getSession(id string) (*execSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *ExecTool) removeSession(id string) (*execSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *ExecTool) finalizeAndRemoveSession(id string) error { _ = "STUB: not implemented"; return nil }

func (t *ExecTool) cleanupExpiredLocked() { _ = "STUB: not implemented"; return }

func (t *ExecTool) markSessionFinalized(sess *execSession) { _ = "STUB: not implemented"; return }

// normalizeCWD forwards to workspacefacade.NormalizeWorkspaceCWD so the
// LLM workspace_exec tool and Workspace.RunProgram share one canonical
// containment policy. Keep the wrapper for the existing call sites.
func normalizeCWD(raw string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func supportsInteractiveSessions(exec codeexecutor.CodeExecutor) bool {
	_ = "STUB: not implemented"
	return false
}

func execOutputSchema(desc string) *tool.Schema { _ = "STUB: not implemented"; return nil }

func pollOutput(sessionID string, poll codeexecutor.ProgramPoll) execOutput {
	_ = "STUB: not implemented"
	return *new(execOutput)
}

func execTimeout(raw int) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func execYield(background bool, raw *int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func writeYield(raw *int) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func combineOutput(stdout, stderr string) string { _ = "STUB: not implemented"; return "" }

func firstIntPtr(vs ...*int) *int { _ = "STUB: not implemented"; return nil }

func firstIntValue(vs ...*int) int { _ = "STUB: not implemented"; return 0 }

func firstBoolValue(vs ...*bool) bool { _ = "STUB: not implemented"; return false }

func firstNonEmpty(values ...string) string { _ = "STUB: not implemented"; return "" }

func intPtrValue(v int) *int {
	_ = "STUB: not implemented"

	// isAllowedWorkspacePath forwards to
	// workspacefacade.IsAllowedWorkspaceRoot. Kept as a package-private
	// alias for the few remaining call sites; new code should call the
	// facade helper directly.
	return nil
}

func isAllowedWorkspacePath(rel string) bool { _ = "STUB: not implemented"; return false }

var _ tool.Tool = (*ExecTool)(nil)
var _ tool.CallableTool = (*ExecTool)(nil)
var _ tool.Tool = (*WriteStdinTool)(nil)
var _ tool.CallableTool = (*WriteStdinTool)(nil)
var _ tool.Tool = (*KillSessionTool)(nil)
var _ tool.CallableTool = (*KillSessionTool)(nil)
