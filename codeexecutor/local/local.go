//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package local provides a CodeExecutor that executes code blocks in the
// local environment. It supports Python and Bash scripts by writing them
// to files and invoking the appropriate interpreter.
package local

import (
	"context"
	"os"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
)

// CodeExecutor that executes code on the local host (unsafe).
type CodeExecutor struct {
	WorkDir            string
	Timeout            time.Duration
	CleanTempFiles     bool
	codeBlockDelimiter codeexecutor.CodeBlockDelimiter
	ws                 *Runtime
	inputsHostBase     string
	autoInputs         bool
	workspaceMode      WorkspaceMode
}

// CodeExecutorOption configures a local CodeExecutor.
type CodeExecutorOption func(*CodeExecutor)

// WithWorkDir sets the working directory used for execution.
func WithWorkDir(workDir string) CodeExecutorOption {
	_ = "STUB: not implemented"
	return *new(CodeExecutorOption)
}

// WithTimeout sets the per-command timeout.
func WithTimeout(timeout time.Duration) CodeExecutorOption {
	_ = "STUB: not implemented"
	return *new(CodeExecutorOption)
}

// WithCleanTempFiles toggles cleanup of temporary helper files.
func WithCleanTempFiles(clean bool) CodeExecutorOption {
	_ = "STUB: not implemented"
	return *new(CodeExecutorOption)
}

// WithWorkspaceInputsHostBase sets the host inputs directory that
// will be exposed under work/inputs when auto inputs are enabled.
func WithWorkspaceInputsHostBase(host string) CodeExecutorOption {
	_ = "STUB: not implemented"
	return *new(CodeExecutorOption)
}

// WithWorkspaceAutoInputs enables or disables automatic mapping of
// the host inputs directory (when configured) into work/inputs for
// each workspace.
func WithWorkspaceAutoInputs(enable bool) CodeExecutorOption {
	_ = "STUB: not implemented"
	return *new(CodeExecutorOption)
}

// WithWorkspaceMode configures how local workspaces are created.
//
// The default is WorkspaceModeIsolated, which creates a unique workspace per
// run. WorkspaceModeTrustedLocal reuses WorkDir as the workspace root.
func WithWorkspaceMode(mode WorkspaceMode) CodeExecutorOption {
	_ = "STUB: not implemented"
	return *new(CodeExecutorOption)
}

// WithCodeBlockDelimiter sets the code block delimiter.
func WithCodeBlockDelimiter(
	delimiter codeexecutor.CodeBlockDelimiter,
) CodeExecutorOption {
	_ = "STUB: not implemented"
	return *new(CodeExecutorOption)
}

var defaultCodeBlockDelimiter = codeexecutor.CodeBlockDelimiter{
	Start: "```",
	End:   "```",
}

const (
	codeFilePatternBase = "code_*"
	pythonFileExt       = ".py"
	shellFileExt        = ".sh"
	prepareFileErrFmt   = "failed to prepare %s file: %w"
)

// New creates a local CodeExecutor.
func New(options ...CodeExecutorOption) *CodeExecutor { _ = "STUB: not implemented"; return nil }

// ExecuteCode executes code blocks and returns combined output.
func (e *CodeExecutor) ExecuteCode(
	ctx context.Context, input codeexecutor.CodeExecutionInput,
) (codeexecutor.CodeExecutionResult, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.CodeExecutionResult), nil
}

// Determine working directory for the command CWD and a separate
// script directory for writing intermediate script files.
// When WorkDir is set, we create a unique temp subdirectory inside it
// for script files to avoid collisions from concurrent ExecuteCode calls
// (e.g. multiple calls all writing to code_0.sh).
// CWD for the executed command
// directory where script files are written

// Create a unique temp subdirectory for script files to prevent
// concurrent calls from overwriting each other's code_0.sh.

// Fall back to writing scripts directly into WorkDir.
// Per-block errors will surface via result.Output.

// Clean up the temp script directory after execution.

// When CleanTempFiles is false, write scripts directly into
// WorkDir so they can be inspected after execution.

func (e *CodeExecutor) executeCodeBlock(
	ctx context.Context, cmdDir, scriptDir string,
	block codeexecutor.CodeBlock,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// prepareCodeFile writes code to a temporary helper file.
func (e *CodeExecutor) prepareCodeFile(
	workDir string, block codeexecutor.CodeBlock,
) (filePath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func writeHelperFile(
	filePath, content string,
	fileMode os.FileMode,
) error {
	_ = "STUB: not implemented"
	return nil
}

func helperFileExtension(language string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (e *CodeExecutor) getFileMode(language string) os.FileMode {
	_ = "STUB: not implemented"
	return *new(os.FileMode)
}

func (e *CodeExecutor) buildCommandArgs(
	language, filePath string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *CodeExecutor) executeCommand(
	ctx context.Context, workDir string, cmdArgs []string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// #nosec G204 — interpreter and path are controlled by us

func removeHelperFile(path string) { _ = "STUB: not implemented"; return }

// CodeBlockDelimiter returns the code block delimiter used by the local
// executor.
func (e *CodeExecutor) CodeBlockDelimiter() codeexecutor.CodeBlockDelimiter {
	_ = "STUB: not implemented"
	return *new(codeexecutor.CodeBlockDelimiter)
}

// Workspace methods

// CreateWorkspace creates a new workspace directory.
func (e *CodeExecutor) ensureWS() *Runtime { _ = "STUB: not implemented"; return nil }

// CreateWorkspace delegates to the local workspace runtime.
func (e *CodeExecutor) CreateWorkspace(
	ctx context.Context, execID string,
	pol codeexecutor.WorkspacePolicy,
) (codeexecutor.Workspace, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.Workspace), nil
}

// Cleanup delegates to the local workspace runtime.
func (e *CodeExecutor) Cleanup(
	ctx context.Context, ws codeexecutor.Workspace,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PutFiles delegates to the local workspace runtime.
func (e *CodeExecutor) PutFiles(
	ctx context.Context, ws codeexecutor.Workspace,
	files []codeexecutor.PutFile,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PutDirectory delegates to the local workspace runtime.
func (e *CodeExecutor) PutDirectory(
	ctx context.Context, ws codeexecutor.Workspace,
	hostPath, to string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// RunProgram delegates to the local workspace runtime.
func (e *CodeExecutor) RunProgram(
	ctx context.Context, ws codeexecutor.Workspace,
	spec codeexecutor.RunProgramSpec,
) (codeexecutor.RunResult, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.RunResult), nil
}

// Collect delegates to the local workspace runtime.
func (e *CodeExecutor) Collect(
	ctx context.Context, ws codeexecutor.Workspace,
	patterns []string,
) ([]codeexecutor.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExecuteInline delegates to the local workspace runtime.
func (e *CodeExecutor) ExecuteInline(
	ctx context.Context, execID string,
	blocks []codeexecutor.CodeBlock,
	timeout time.Duration,
) (codeexecutor.RunResult, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.RunResult), nil
}

// Engine exposes the local runtime as an Engine for skills.
func (e *CodeExecutor) Engine() codeexecutor.Engine {
	_ = "STUB: not implemented"
	return *new(codeexecutor.Engine)
}
