//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package skillstage

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
)

const (
	skillDirInputs = "inputs"
	skillDirVenv   = ".venv"
)

const workspaceMetadataFileMode uint32 = 0o600

// Stager materializes skill package contents into a workspace and maintains
// the corresponding workspace metadata and links.
type Stager struct{}

// New creates a skill stager.
func New() *Stager {
	_ = "STUB: not implemented"

	// StageOptions tunes how StageSkillWithOptions materializes a skill
	// working copy. The zero value matches the default behavior expected
	// by the workspaceprep reconciler: a writable session-level working
	// copy that scripts may freely modify.
	return nil
}

type StageOptions struct {
	// ReadOnly flips the staged tree to read-only after copy. This
	// is the legacy behavior used by the now-deprecated skill_run
	// tool and should not be used by new callers; treating skills/
	// as a writable working copy is the default contract.
	ReadOnly bool
}

// StageSkill copies a skill into the shared workspace and links the shared
// work/out roots under skills/<name>. The staged tree is writable by
// default; callers that need the legacy read-only semantics can use
// StageSkillWithOptions.
func (s *Stager) StageSkill(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	root string,
	name string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// StageSkillWithOptions is StageSkill with explicit knobs. It exists
// so legacy entry points can request the old read-only behavior while
// new workspace-preparation code keeps the writable-by-default
// contract.
func (s *Stager) StageSkillWithOptions(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	root string,
	name string,
	opts StageOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stager) stageSkillWithOptionsLocked(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	root string,
	name string,
	dg string,
	opts StageOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

// LoadWorkspaceMetadata reads workspace metadata from the shared metadata file
// and returns a normalized in-memory view with defaults applied.
func (s *Stager) LoadWorkspaceMetadata(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
) (codeexecutor.WorkspaceMetadata, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.WorkspaceMetadata), nil
}

// SaveWorkspaceMetadata persists workspace metadata into the shared metadata
// file within the current workspace.
func (s *Stager) SaveWorkspaceMetadata(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	md codeexecutor.WorkspaceMetadata,
) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanupMetadataTemp(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	tmpFile string,
	committed *bool,
) {
	_ = "STUB: not implemented"
	return
}

// SkillLinksPresent reports whether the staged skill directory still exposes
// the expected shared-directory symlinks back into the workspace roots.
func (s *Stager) SkillLinksPresent(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	name string,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *Stager) linkWorkspaceDirs(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	name string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveWorkspacePath removes a workspace-relative path after first making
// non-symlink files writable so cleanup can succeed on read-only staged trees.
func (s *Stager) RemoveWorkspacePath(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	rel string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stager) readOnlyExceptSymlinks(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	dest string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func runProgramExitError(
	op string,
	res codeexecutor.RunResult,
	err error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func shellQuote(s string) string { _ = "STUB: not implemented"; return "" }
