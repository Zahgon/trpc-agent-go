//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package container

import (
	"context"
	"io"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
)

const (
	maxReadSizeBytes        = 4 * 1024 * 1024 // 4 MiB
	defaultCreateTimeoutSec = 5
	defaultRmTimeoutSec     = 10
	defaultStageTimeoutSec  = 10
	// Use a neutral, writable base inside the container for workspaces.
	// Avoid /mnt and /workspace which may be read-only in some setups.
	// /tmp is typically a writable tmpfs inside containers.
	defaultRunContainerBase = "/tmp/run"
	// Bind-mounted skills default inside container.
	defaultSkillsContainer = "/opt/trpc-agent/skills"
	// Bind-mounted inputs default inside container.
	defaultInputsContainer = "/opt/trpc-agent/inputs"

	inputSchemeArtifact  = "artifact://"
	inputSchemeHost      = "host://"
	inputSchemeWorkspace = "workspace://"
	inputSchemeSkill     = "skill://"

	metadataFileMode = 0o600
)

// workspaceRuntime provides workspace execution on Docker.
type workspaceRuntime struct {
	ce  *CodeExecutor
	cfg runtimeConfig
}

type runtimeConfig struct {
	skillsHostBase      string
	skillsContainerBase string
	runHostBase         string
	runContainerBase    string
	inputsHostBase      string
	inputsContainerBase string
	autoMapInputs       bool
}

// newWorkspaceRuntime builds a runtime bound to the provided executor.
func newWorkspaceRuntime(c *CodeExecutor) (*workspaceRuntime, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default inputs mount location inside container.

// Infer host bases that are bind-mounted at the default skills
// and inputs locations when present.

// findBindSource returns the host path whose bind dest equals dest.
// Bind spec is source:dest[:mode]. We parse from right to handle ':'
// that may appear in the source path (Windows not considered here).
func findBindSource(binds []string, dest string) string { _ = "STUB: not implemented"; return "" }

// Last part may be mode; second last is dest.

// Join all but the last two parts as source.

// CreateWorkspace ensures a per‑execution directory inside container.
func (r *workspaceRuntime) CreateWorkspace(
	ctx context.Context,
	execID string,
	pol codeexecutor.WorkspacePolicy,
) (codeexecutor.Workspace, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.Workspace), nil
}

// Make workspace path unique to avoid collisions.

// Create standard layout and metadata.json inside container.

// Cleanup removes the workspace directory.
func (r *workspaceRuntime) Cleanup(
	ctx context.Context,
	ws codeexecutor.Workspace,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PutFiles writes files via CopyToContainer.
func (r *workspaceRuntime) PutFiles(
	ctx context.Context,
	ws codeexecutor.Workspace,
	files []codeexecutor.PutFile,
) error {
	_ = "STUB: not implemented"
	return nil
}

// PutDirectory copies a host directory into the workspace.
func (r *workspaceRuntime) PutDirectory(
	ctx context.Context,
	ws codeexecutor.Workspace,
	hostPath string,
	to string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Fast path: within skills mount; copy inside container.

// fall through to tar copy on error

// Pack dir into tar stream.

// Ensure destination exists in container.

// StageDirectory stages a directory with options.
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

// RunProgram runs a command in the workspace with timeout.
func (r *workspaceRuntime) RunProgram(
	ctx context.Context,
	ws codeexecutor.Workspace,
	spec codeexecutor.RunProgramSpec,
) (codeexecutor.RunResult, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.RunResult), nil
}

// Prepare standard dirs and env injection.

// Build env parts with defaults first, then overlay user env.

// Ensure run/output dirs exist before cd/exec.

// Collect copies out files by glob patterns (simple exact path here).
func (r *workspaceRuntime) Collect(
	ctx context.Context,
	ws codeexecutor.Workspace,
	patterns []string,
) ([]codeexecutor.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use bash globstar to approximate doublestar semantics.

// Canonicalize path via readlink/realpath to collapse symlinks.

// Convert to workspace-relative canonical path and dedupe.

// StageInputs maps external inputs using container semantics.
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

func (r *workspaceRuntime) stageHostInput(
	ctx context.Context,
	ws codeexecutor.Workspace,
	sp codeexecutor.InputSpec,
	mode string,
	to string,
	dest string,
) (string, *int, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// If under inputsHostBase, prefer symlink/cp (zero-copy).

// Fallback: tar copy host path to dest dir.

func (r *workspaceRuntime) stageWorkspaceInput(
	ctx context.Context,
	ws codeexecutor.Workspace,
	sp codeexecutor.InputSpec,
	mode string,
	dest string,
) (string, *int, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (r *workspaceRuntime) stageSkillInput(
	ctx context.Context,
	ws codeexecutor.Workspace,
	sp codeexecutor.InputSpec,
	mode string,
	dest string,
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

func (r *workspaceRuntime) saveWorkspaceMetadata(
	ctx context.Context,
	ws codeexecutor.Workspace,
	md codeexecutor.WorkspaceMetadata,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *workspaceRuntime) loadWorkspaceMetadata(
	ctx context.Context,
	ws codeexecutor.Workspace,
) (codeexecutor.WorkspaceMetadata, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.WorkspaceMetadata), nil
}

func (r *workspaceRuntime) cleanupMetadataTemp(
	ctx context.Context,
	tmpPath string,
	committed *bool,
) {
	_ = "STUB: not implemented"
	return
}

// CollectOutputs applies container-side glob and optional save.
func (r *workspaceRuntime) CollectOutputs(
	ctx context.Context,
	ws codeexecutor.Workspace,
	spec codeexecutor.OutputSpec,
) (codeexecutor.OutputManifest, error) {
	_ = "STUB: not implemented"
	// Build bash to expand globstar and echo absolute file paths.
	return *new(codeexecutor.OutputManifest), nil
}

func (r *workspaceRuntime) copyBytesTo(
	ctx context.Context, dest string, data []byte, mode uint32,
) error {
	_ = "STUB: not implemented"
	// Create a tar with single file named as dest's base.
	return nil
}

// Ensure parent exists.

// Copy to parent dir.

func inputBase(from string) string { _ = "STUB: not implemented"; return "" }

// ExecuteInline writes code blocks and runs them.
func (r *workspaceRuntime) ExecuteInline(
	ctx context.Context,
	execID string,
	blocks []codeexecutor.CodeBlock,
	timeout time.Duration,
) (codeexecutor.RunResult, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.RunResult), nil
}

// Internal helpers

func (r *workspaceRuntime) execCmd(
	ctx context.Context,
	argv []string,
	timeout time.Duration,
) (string, string, int, bool, error) {
	_ = "STUB: not implemented"
	return "", "", 0, false, nil
}

func (r *workspaceRuntime) execCmdWithStdin(
	ctx context.Context,
	argv []string,
	timeout time.Duration,
	stdin string,
) (string, string, int, bool, error) {
	_ = "STUB: not implemented"
	return "", "", 0, false, nil
}

func sanitize(s string) string { _ = "STUB: not implemented"; return "" }

func shellQuote(s string) string { _ = "STUB: not implemented"; return "" }

func tarFromFiles(files []codeexecutor.PutFile) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (r *workspaceRuntime) copyFileOut(
	ctx context.Context,
	fullPath string,
) ([]byte, int64, string, string, error) {
	_ = "STUB: not implemented"
	return nil, 0, "", "", nil
}
