//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package codeexecutor

import (
	"context"
	"errors"
	"time"
)

// ErrWorkspaceInitNeedsEngineProvider is returned when hooks are requested but
// the executor does not implement [EngineProvider].
var ErrWorkspaceInitNeedsEngineProvider = errors.New(
	"codeexecutor: workspace init hooks require CodeExecutor to implement EngineProvider",
)

// ErrWorkspaceInitIncompleteEngine is returned when Engine() or Manager() is nil.
var ErrWorkspaceInitIncompleteEngine = errors.New(
	"codeexecutor: Engine() or Engine.Manager() is nil",
)

const (
	workspaceInitErrorOutputMax = 1024
	workspaceInitCleanupTimeout = 30 * time.Second
)

// WorkspaceInitHook is a callback run after [WorkspaceManager.CreateWorkspace]
// succeeds and before that workspace is returned to callers. Use it for
// deterministic setup (stage inputs, install dependencies). Hooks are scoped to
// workspace creation: they do not watch files on disk or re-run later solely
// because workspace contents changed.
//
// Multiple hooks run in declaration order; failure labels use hook index
// (0, 1, ...). Use [NewWorkspaceInitHook] for declarative staging and commands with
// per-command diagnostic names in errors.
type WorkspaceInitHook func(context.Context, WorkspaceInitEnv) error

// WorkspaceInitEnv is the capability bundle passed to each [WorkspaceInitHook].
// The workspace directory already exists when the hook runs.
type WorkspaceInitEnv struct {
	Workspace    Workspace
	ExecID       string
	Policy       WorkspacePolicy
	FS           WorkspaceFS
	Runner       ProgramRunner
	Capabilities Capabilities
}

// WorkspaceInitSpec is a declarative hook: stage [InputSpec] inputs, then run
// init commands in order. It does not express per-tool-call idempotency; callers
// that need convergence on every tool invocation handle that at a higher layer.
type WorkspaceInitSpec struct {
	// Inputs are staged via [WorkspaceFS.StageInputs] (artifact://, host://, etc.).
	Inputs []InputSpec
	// Commands run sequentially after inputs; non-zero exit code fails the hook.
	Commands []WorkspaceInitCommand
}

// WorkspaceInitCommand describes one init-time program run. Fields align with
// [RunProgramSpec] minus [ResourceLimits], which init hooks omit in v1.
// Name is optional; when set it appears in errors for that command.
type WorkspaceInitCommand struct {
	Name    string
	Cmd     string
	Args    []string
	Env     map[string]string
	Cwd     string
	Stdin   string
	Timeout time.Duration
}

// NewWorkspaceInitHook wraps [WorkspaceInitSpec] as a [WorkspaceInitHook] function.
func NewWorkspaceInitHook(spec WorkspaceInitSpec) WorkspaceInitHook {
	_ = "STUB: not implemented"
	return *new(WorkspaceInitHook)
}

// NewWorkspaceInitExecutor wraps exec so every [WorkspaceManager.CreateWorkspace]
// runs the given hooks before returning the workspace.
//
// When hooks is non-empty, exec must implement [EngineProvider] with a non-nil
// [Engine] and non-nil [WorkspaceManager]; otherwise this function returns an
// error satisfying [ErrWorkspaceInitNeedsEngineProvider] or
// [ErrWorkspaceInitIncompleteEngine].
//
// For [InputSpec] values that use artifact://, the context passed to
// CreateWorkspace must carry the artifact service and (when applicable) session
// information—the same requirements as [WorkspaceFS.StageInputs]. Standard agent
// workspace-session tooling injects that context before [WorkspaceRegistry.Acquire],
// so init hooks can load artifacts without extra setup.
func NewWorkspaceInitExecutor(
	exec CodeExecutor,
	hooks ...WorkspaceInitHook,
) (CodeExecutor, error) {
	_ = "STUB: not implemented"
	return *new(CodeExecutor), nil
}

func cloneStringStringMap(m map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func truncateInitOut(s string) string { _ = "STUB: not implemented"; return "" }

type workspaceInitExecutor struct {
	CodeExecutor
	ep    EngineProvider
	hooks []WorkspaceInitHook
}

func (e *workspaceInitExecutor) Engine() Engine { _ = "STUB: not implemented"; return *new(Engine) }

type workspaceInitEngine struct {
	inner Engine
	hooks []WorkspaceInitHook
}

func newWorkspaceInitEngine(inner Engine, hooks []WorkspaceInitHook) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

func (e *workspaceInitEngine) Manager() WorkspaceManager {
	_ = "STUB: not implemented"
	return *new(WorkspaceManager)
}

func (e *workspaceInitEngine) FS() WorkspaceFS { _ = "STUB: not implemented"; return *new(WorkspaceFS) }

func (e *workspaceInitEngine) Runner() ProgramRunner {
	_ = "STUB: not implemented"
	return *new(ProgramRunner)
}

func (e *workspaceInitEngine) Describe() Capabilities {
	_ = "STUB: not implemented"
	return *new(Capabilities)
}

type workspaceInitManager struct {
	inner WorkspaceManager
	eng   Engine
	hooks []WorkspaceInitHook
}

func (m *workspaceInitManager) CreateWorkspace(
	ctx context.Context,
	execID string,
	pol WorkspacePolicy,
) (Workspace, error) {
	_ = "STUB: not implemented"
	return *new(Workspace), nil
}

// Best-effort cleanup without inheriting deadline/cancel from ctx,
// which may already be expired when the hook failed for timeout.

func (m *workspaceInitManager) Cleanup(
	ctx context.Context,
	ws Workspace,
) error {
	_ = "STUB: not implemented"
	return nil
}
