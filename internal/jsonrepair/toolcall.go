//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package jsonrepair

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// IsToolCallArgumentsJSONRepairEnabled reports whether tool call arguments JSON repair is enabled.
func IsToolCallArgumentsJSONRepairEnabled(invocation *agent.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}

// RepairToolCallArguments returns repaired tool call arguments when the input is not valid JSON.
func RepairToolCallArguments(ctx context.Context, toolName string, arguments []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// chooseToolCallArguments prefers repaired when it is a non-empty JSON payload.
func chooseToolCallArguments(arguments []byte, repaired []byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// RepairToolCallArgumentsInPlace repairs the tool call arguments in place when needed.
func RepairToolCallArgumentsInPlace(ctx context.Context, toolCall *model.ToolCall) {
	_ = "STUB: not implemented"
	return
}

// RepairToolCallsArgumentsInPlace repairs tool call arguments in place when needed.
func RepairToolCallsArgumentsInPlace(ctx context.Context, toolCalls []model.ToolCall) {
	_ = "STUB: not implemented"
	return
}

// RepairResponseToolCallArgumentsInPlace repairs tool call arguments inside the response in place when needed.
func RepairResponseToolCallArgumentsInPlace(ctx context.Context, response *model.Response) {
	_ = "STUB: not implemented"
	return
}
