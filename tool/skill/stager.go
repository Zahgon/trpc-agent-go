//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package skill

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
	rootskill "trpc.group/trpc-go/trpc-agent-go/skill"
)

const (
	errSkillStagerNotConfigured = "skill stager is not configured"
	errSkillRepoNotConfigured   = "skill repository is not configured"
)

// SkillStager materializes a skill into the current workspace.
//
// Implementations may copy, mount, preload, or no-op. The returned
// WorkspaceSkillDir must be the workspace-relative directory of the
// specific staged skill and must remain within the known workspace
// roots.
type SkillStager interface {
	StageSkill(
		ctx context.Context,
		req SkillStageRequest,
	) (SkillStageResult, error)
}

// SkillStageRequest describes the skill staging context for one run.
type SkillStageRequest struct {
	SkillName  string
	Repository rootskill.Repository
	Engine     codeexecutor.Engine
	Workspace  codeexecutor.Workspace
}

// SkillStageResult reports where the staged skill lives inside the
// workspace.
type SkillStageResult struct {
	// WorkspaceSkillDir is the workspace-relative directory of the
	// staged skill, such as "skills/weather" or
	// "work/custom/weather". It must point to the specific skill
	// directory, not just the shared "skills" root, and it must not
	// be a sandbox absolute path like "/sandbox/workspace/skills".
	WorkspaceSkillDir string
}

// WithSkillStager overrides the strategy used to materialize skills
// into the workspace.
//
// When unset, RunTool uses the default copy-based stager, which stages
// the skill under "skills/<skill-name>".
func WithSkillStager(stager SkillStager) func(*RunTool) { _ = "STUB: not implemented"; return nil }

type copySkillStager struct {
	tool *RunTool
}

func newCopySkillStager(tool *RunTool) SkillStager {
	_ = "STUB: not implemented"
	return *new(SkillStager)
}

func (s *copySkillStager) StageSkill(
	ctx context.Context,
	req SkillStageRequest,
) (SkillStageResult, error) {
	_ = "STUB: not implemented"
	return *new(SkillStageResult), nil
}

func defaultWorkspaceSkillDir(name string) string { _ = "STUB: not implemented"; return "" }

func normalizeSkillStageResult(
	res SkillStageResult,
) (SkillStageResult, error) {
	_ = "STUB: not implemented"
	return *new(SkillStageResult), nil
}

func normalizeWorkspaceSkillDir(dir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
