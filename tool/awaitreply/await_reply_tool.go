//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package awaitreply provides the await_user_reply framework tool.
package awaitreply

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	// ToolName is the await_user_reply framework tool name.
	ToolName = "await_user_reply"
)

// Request is the request payload for await_user_reply.
type Request struct{}

// Response is the tool response payload for await_user_reply.
type Response struct {
	// Success indicates whether the route was recorded successfully.
	Success bool `json:"success"`
	// Message describes the result.
	Message string `json:"message"`
	// AgentName is the agent that will receive the next user turn.
	AgentName string `json:"agent_name,omitempty"`
}

// Tool marks the current agent as waiting for the next user reply.
type Tool struct{}

// New creates a new await_user_reply tool.
func New() *Tool {
	_ = "STUB: not implemented"

	// Declaration implements tool.Tool.
	return nil
}

func (t *Tool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// Call implements tool.CallableTool.
func (t *Tool) Call(ctx context.Context, jsonArgs []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
