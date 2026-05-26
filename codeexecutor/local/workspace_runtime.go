//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package local

// Workspace runtime provides workspace-based execution on local host.

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
)

const (
	defaultTimeoutSec = 10
	defaultFileMode   = 0o644
	maxReadSizeBytes  = 4 * 1024 * 1024 // 4 MiB per output file
)

// Runtime implements the workspace-based executor using local processes.
type Runtime struct {
	WorkRoot            string
	ReadOnlyStagedSkill bool
	InputsHostBase      string
	AutoInputs          bool
	Mode                WorkspaceMode
}

// WorkspaceMode controls how the local runtime chooses workspace roots.
type WorkspaceMode int

const (
	// WorkspaceModeIsolated creates a unique workspace directory for each
	// CreateWorkspace call. This is the default and is safer.
	WorkspaceModeIsolated WorkspaceMode = iota
	// WorkspaceModeTrustedLocal reuses WorkRoot as the workspace root.
	// Cleanup becomes a no-op to avoid deleting user directories.
	WorkspaceModeTrustedLocal
)

// NewRuntime creates a new local Runtime. When workRoot is empty, a
// temporary directory will be used per workspace.
func NewRuntime(workRoot string) *Runtime { _ = "STUB: not implemented"; return nil }

// PathListSeparator returns the local host separator used for PATH-like
// environment variables.
func (*Runtime) PathListSeparator() string { _ = "STUB: not implemented"; return "" }

// RuntimeOption customizes the local Runtime behavior.
type RuntimeOption func(*Runtime)

// WithRuntimeWorkspaceMode sets the workspace mode for the runtime.
func WithRuntimeWorkspaceMode(mode WorkspaceMode) RuntimeOption {
	_ = "STUB: not implemented"
	return *new(RuntimeOption)
}

// WithReadOnlyStagedSkill toggles making staged skill trees read-only.
func WithReadOnlyStagedSkill(readOnly bool) RuntimeOption {
	_ = "STUB: not implemented"
	return *new(RuntimeOption)
}

// WithInputsHostBase sets the host directory that will be exposed
// under work/inputs inside each workspace when auto inputs are
// enabled.
func WithInputsHostBase(host string) RuntimeOption {
	_ = "STUB: not implemented"
	return *new(RuntimeOption)
}

// WithAutoInputs enables or disables automatic mapping of the host
// inputs directory (when configured) into work/inputs for each
// workspace.
func WithAutoInputs(enable bool) RuntimeOption {
	_ = "STUB: not implemented"
	return *new(RuntimeOption)
}

// NewRuntimeWithOptions creates a Runtime with optional settings.
func NewRuntimeWithOptions(
	workRoot string, opts ...RuntimeOption,
) *Runtime {
	_ = "STUB: not implemented"
	return nil
}

// CreateWorkspace creates an execution workspace directory.
func (r *Runtime) CreateWorkspace(
	ctx context.Context,
	execID string,
	pol codeexecutor.WorkspacePolicy,
) (codeexecutor.Workspace, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.Workspace), nil
}

// Sanitize execID to be filesystem friendly.

// replace others

// Make workspace path unique to avoid collisions between runs.

// Persist is respected by callers deciding whether to call Cleanup.

// Ensure standard layout and metadata.json.

func (r *Runtime) createTrustedWorkspace(
	ctx context.Context,
	execID string,
) (codeexecutor.Workspace, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.Workspace), nil
}

// Cleanup removes workspace directory if it exists.
func (r *Runtime) Cleanup(
	ctx context.Context,
	ws codeexecutor.Workspace,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PutFiles writes file blobs under the workspace root.
func (r *Runtime) PutFiles(
	ctx context.Context,
	ws codeexecutor.Workspace,
	files []codeexecutor.PutFile,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PutDirectory copies an entire directory from host into workspace.
func (r *Runtime) PutDirectory(
	ctx context.Context,
	ws codeexecutor.Workspace,
	hostPath string,
	to string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// StageDirectory stages a host directory into the workspace.
// Behavior depends on options, e.g., making the tree read-only.
func (r *Runtime) StageDirectory(
	ctx context.Context,
	ws codeexecutor.Workspace,
	src string,
	to string,
	opt codeexecutor.StageOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

// RunProgram runs a command inside the workspace.
func (r *Runtime) RunProgram(
	ctx context.Context,
	ws codeexecutor.Workspace,
	spec codeexecutor.RunProgramSpec,
) (codeexecutor.RunResult, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.RunResult), nil
}

// Resolve cwd under workspace.

// Ensure cwd exists.

// Keep exitCode as 0 when killed by timeout.

// Map other errors to exitCode -1 for visibility.

// Collect finds output files by glob patterns relative to workspace root.
func (r *Runtime) Collect(
	ctx context.Context,
	ws codeexecutor.Workspace,
	patterns []string,
) ([]codeexecutor.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Canonicalize root to make prefix checks robust on platforms
// where different paths may refer to the same location.

// Use doublestar to support ** patterns.

// Doublestar on os.DirFS("/") expects patterns relative to "/".

// Convert match back to absolute path.

// Ensure it is within root.

// Collapse symlinks to canonical path and dedupe.

// Re-check containment against canonical root.

const (
	inputSchemeArtifact  = "artifact://"
	inputSchemeHost      = "host://"
	inputSchemeWorkspace = "workspace://"
	inputSchemeSkill     = "skill://"
)

func pinnedArtifactVersion(
	md codeexecutor.WorkspaceMetadata,
	name string,
	to string,
) *int {
	_ = "STUB: not implemented"
	return nil
}

// StageInputs maps external inputs into the workspace.
func (r *Runtime) StageInputs(
	ctx context.Context,
	ws codeexecutor.Workspace,
	specs []codeexecutor.InputSpec,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Runtime) stageInputsLocked(
	ctx context.Context,
	ws codeexecutor.Workspace,
	specs []codeexecutor.InputSpec,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Runtime) stageInput(
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

// CollectOutputs implements the declarative collector with limits.
// It returns codeexecutor.ErrPartialOutputCommit when collected files exist
// but the metadata record could not be committed.
func (r *Runtime) CollectOutputs(
	ctx context.Context,
	ws codeexecutor.Workspace,
	spec codeexecutor.OutputSpec,
) (codeexecutor.OutputManifest, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.OutputManifest), nil
}

func collectOutputMatch(
	ctx context.Context,
	wsPath string,
	absPath string,
	spec codeexecutor.OutputSpec,
	maxFileBytes int64,
	leftTotal int64,
) (codeexecutor.FileRef, int64, bool, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.FileRef), 0, false, nil
}

func withinWorkspacePath(wsPath string, absPath string) bool {
	_ = "STUB: not implemented"
	return false
}

// ExecuteInline writes temp files for code blocks and runs them.
func (r *Runtime) ExecuteInline(
	ctx context.Context,
	execID string,
	blocks []codeexecutor.CodeBlock,
	timeout time.Duration,
) (codeexecutor.RunResult, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.RunResult), nil
}

// Helpers

func defaultTimeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (r *Runtime) writeFileSafe(root string, f codeexecutor.PutFile) error {
	_ = "STUB: not implemented"
	return nil
}

// Resolve and ensure inside root.

func copyDir(src, dst string) error {
	_ = "STUB: not implemented"
	// Ensure destination exists.
	return nil
}

func readLimited(path string) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func readLimitedWithCap(path string, capBytes int) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func makeSymlink(root, toRel, target string) error { _ = "STUB: not implemented"; return nil }

// Remove existing path if present.

func copyPath(src, dst string) error { _ = "STUB: not implemented"; return nil }

func inputDefaultName(from string) string { _ = "STUB: not implemented"; return "" }

// Strip scheme and keep tail element as default name.

// makeTreeReadOnly removes write bits from the entire tree.
func makeTreeReadOnly(root string) error { _ = "STUB: not implemented"; return nil }

// Clear write bits: owner/group/other.
