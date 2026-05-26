//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package llmagent

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
	"trpc.group/trpc-go/trpc-agent-go/internal/skillprofile"
	"trpc.group/trpc-go/trpc-agent-go/internal/surfacepatch"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

func (a *LLMAgent) rootSurfacePatch(
	inv *agent.Invocation,
) (surfacepatch.Patch, bool) {
	_ = "STUB: not implemented"
	return *new(surfacepatch.Patch), false
}

func (a *LLMAgent) fewShotForInvocation(
	inv *agent.Invocation,
) [][]model.Message {
	_ = "STUB: not implemented"
	return nil
}

func (a *LLMAgent) skillRepositoryForInvocation(
	inv *agent.Invocation,
) skill.Repository {
	_ = "STUB: not implemented"
	return *new(skill.Repository)
}

func (a *LLMAgent) modelSurfaceForInvocation(
	inv *agent.Invocation,
) (model.Model, bool) {
	_ = "STUB: not implemented"
	return *new(model.Model), false
}

func (a *LLMAgent) codeExecutorForInvocation(
	inv *agent.Invocation,
) codeexecutor.CodeExecutor {
	_ = "STUB: not implemented"
	return *new(codeexecutor.CodeExecutor)
}

func (a *LLMAgent) supportsWorkspaceExecForInvocation(
	inv *agent.Invocation,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *LLMAgent) supportsWorkspaceExecSessionsForInvocation(
	inv *agent.Invocation,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *LLMAgent) skillToolFlagsForInvocation(
	inv *agent.Invocation,
) skillprofile.Flags {
	_ = "STUB: not implemented"
	return *new(skillprofile.Flags)
}

// ExecutionTraceAppliedSurfaceIDs reports the effective surfaces that affected one invocation step.
func (a *LLMAgent) ExecutionTraceAppliedSurfaceIDs(inv *agent.Invocation) []string {
	_ = "STUB: not implemented"
	return nil
}

// InvocationToolSurface returns the invocation-scoped tool surface and user tool names.
func (a *LLMAgent) InvocationToolSurface(
	ctx context.Context,
	inv *agent.Invocation,
) ([]tool.Tool, map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Pass effectiveSkills so workspace_exec's loaded-skills
// reconcile reads the same repository that skill tools and the
// skills request processor use on this invocation. Without this
// alignment, a surface-patch repo override would be honored by
// the skill tools but silently ignored by the reconciler path
// added in this change set, causing the model context and the
// materialized skill working copy to drift apart.

// Extension-contributed tools (WithExtensions →
// extension.Registry.Tools) sit at the same logical layer as
// other framework-managed auto-injected tools: not folded into
// userToolNames, yet present on the outbound tool surface.
//
// Append them after every framework tool (knowledge, workspace,
// skills, session recall, await_user_reply and transfer) so
// earlier-wins dedup also protects later framework declarations
// from extension name collisions.

func (a *LLMAgent) userToolsForInvocation(
	ctx context.Context,
	patch surfacepatch.Patch,
) ([]tool.Tool, map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyUserToolPatch(
	userTools []tool.Tool,
	userToolNames map[string]bool,
	patch surfacepatch.Patch,
) ([]tool.Tool, map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterInvocationUserTools(
	ctx context.Context,
	userTools []tool.Tool,
	userToolNames map[string]bool,
	filter tool.FilterFunc,
) ([]tool.Tool, map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil
}
