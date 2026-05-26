//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package codeexecutor

import "context"

// RunEnvProvider returns per-run environment variables derived from the
// execution context. Implementations typically read caller-supplied
// state (e.g. user tokens) from the context and return them as
// key-value pairs to be merged into every RunProgramSpec executed by
// the wrapped engine.
//
// Returned entries never override values already present in the spec's
// Env map, preserving explicit per-call overrides from tools.
type RunEnvProvider func(ctx context.Context) map[string]string

// NewEnvInjectingEngine wraps eng so that every RunProgram and
// StartProgram call merges environment variables from provider into
// the spec before delegating to the underlying runner.
//
// If eng or provider is nil the original engine is returned unchanged.
func NewEnvInjectingEngine(eng Engine, provider RunEnvProvider) Engine {
	_ = "STUB: not implemented"
	return *new(Engine)
}

type envEngine struct {
	inner    Engine
	provider RunEnvProvider
}

func (e *envEngine) Manager() WorkspaceManager {
	_ = "STUB: not implemented"
	return *new(WorkspaceManager)
}
func (e *envEngine) FS() WorkspaceFS        { _ = "STUB: not implemented"; return *new(WorkspaceFS) }
func (e *envEngine) Describe() Capabilities { _ = "STUB: not implemented"; return *new(Capabilities) }

func (e *envEngine) Runner() ProgramRunner { _ = "STUB: not implemented"; return *new(ProgramRunner) }

// envRunner wraps a ProgramRunner to inject per-run env vars.
type envRunner struct {
	inner    ProgramRunner
	provider RunEnvProvider
}

func (r *envRunner) RunProgram(
	ctx context.Context,
	ws Workspace,
	spec RunProgramSpec,
) (RunResult, error) {
	_ = "STUB: not implemented"
	return *new(RunResult), nil
}

// envInteractiveRunner extends envRunner with InteractiveProgramRunner
// support, so that type assertions from workspace_exec etc. continue
// to work through the wrapper.
type envInteractiveRunner struct {
	envRunner
	interactive InteractiveProgramRunner
}

func (r *envInteractiveRunner) StartProgram(
	ctx context.Context,
	ws Workspace,
	spec InteractiveProgramSpec,
) (ProgramSession, error) {
	_ = "STUB: not implemented"
	return *new(ProgramSession), nil
}

// NewEnvInjectingCodeExecutor wraps exec so that Engine() returns an
// env-injecting engine. exec must implement EngineProvider; if it does
// not (or is nil), the original executor is returned unchanged.
//
// This is the recommended top-level entry point: pass the wrapped
// executor to llmagent.WithCodeExecutor and all tool paths (skill_run,
// workspace_exec, interactive sessions) will automatically receive the
// injected environment variables.
func NewEnvInjectingCodeExecutor(
	exec CodeExecutor,
	provider RunEnvProvider,
) CodeExecutor {
	_ = "STUB: not implemented"
	return *new(CodeExecutor)
}

type envCodeExecutor struct {
	CodeExecutor
	ep       EngineProvider
	provider RunEnvProvider
}

func (e *envCodeExecutor) Engine() Engine { _ = "STUB: not implemented"; return *new(Engine) }

// mergeProviderEnv builds a fresh Env map that contains all entries
// from spec.Env plus any provider-supplied entries whose keys are not
// already present. The original spec.Env map is never mutated.
func mergeProviderEnv(
	ctx context.Context,
	provider RunEnvProvider,
	spec *RunProgramSpec,
) {
	_ = "STUB: not implemented"
	return
}
