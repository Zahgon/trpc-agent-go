//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package graph

import (
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/internal/surfacepatch"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

func graphInvocationFromState(state State) *agent.Invocation { _ = "STUB: not implemented"; return nil }

func graphSurfacePatch(
	invocation *agent.Invocation,
	localNodeID string,
) (surfacepatch.Patch, bool) {
	_ = "STUB: not implemented"
	return *new(surfacepatch.Patch), false
}

func graphPatchedModel(
	invocation *agent.Invocation,
	localNodeID string,
	fallback model.Model,
) model.Model {
	_ = "STUB: not implemented"
	return *new(model.Model)
}

func toolSliceToMap(tools []tool.Tool) map[string]tool.Tool { _ = "STUB: not implemented"; return nil }

func toolMapToSlice(tools map[string]tool.Tool) []tool.Tool { _ = "STUB: not implemented"; return nil }

func applyToolMapPatch(
	base map[string]tool.Tool,
	patch surfacepatch.Patch,
) (map[string]tool.Tool, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *llmRunner) currentNodeID(state State) string { _ = "STUB: not implemented"; return "" }

func (r *llmRunner) insertFewShot(
	state State,
	messages []model.Message,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}
