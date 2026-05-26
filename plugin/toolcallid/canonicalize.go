//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package toolcallid

import (
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const canonicalToolCallIDPrefix = "trpc-agent-go-toolcall"

func canonicalizeResponse(inv *agent.Invocation, rsp *model.Response) (*model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func canonicalizeToolCalls(
	toolCalls []model.ToolCall,
	invocationID string,
	responseID string,
	choiceIndex int,
) ([]model.ToolCall, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func canonicalToolCallID(
	invocationID string,
	responseID string,
	rawToolCallID string,
	choiceIndex int,
	slotIndex int,
) string {
	_ = "STUB: not implemented"
	return ""
}
