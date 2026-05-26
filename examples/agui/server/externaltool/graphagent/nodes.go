//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"

	demotool "trpc.group/trpc-go/trpc-agent-go/examples/agui/server/externaltool/graphagent/tool"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

var executeInternalToolsNode = graph.NewToolsNodeFunc(
	demotool.NewInternalTools(),
	graph.WithEnableParallelTools(true),
)

func internalToolNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func externalInterruptNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func externalToolResultMessage(
	ctx context.Context,
	state graph.State,
	pendingToolCall model.ToolCall,
	pendingToolCalls []model.ToolCall,
) (model.Message, error) {
	_ = "STUB: not implemented"
	return *new(model.Message), nil
}

func externalInterruptPrompt(toolCalls []model.ToolCall) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func latestAssistantToolCalls(messages []model.Message) []model.ToolCall {
	_ = "STUB: not implemented"
	return nil
}

func hasToolResult(messages []model.Message, toolID string) bool {
	_ = "STUB: not implemented"
	return false
}
