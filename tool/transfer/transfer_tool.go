//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package transfer provides transfer_to_agent tool implementation.
package transfer

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	// TransferToolName is the name of the transfer_to_agent tool.
	TransferToolName = "transfer_to_agent"
	// FieldAgentName is the name of the agent_name field.
	FieldAgentName = "agent_name"
	// FieldMessage is the name of the message field.
	FieldMessage = "message"
)

// Request represents the request structure for transfer_to_agent tool.
type Request struct {
	// AgentName is the name of the target agent to transfer to.
	AgentName string `json:"agent_name" jsonschema:"description=Name of the agent to transfer control to"`
	// Message is the message to send to the target agent (optional).
	Message string `json:"message,omitempty" jsonschema:"description=Optional message to pass to the target agent"`
}

// Response represents the response from transfer_to_agent tool.
type Response struct {
	// Success indicates if the transfer was successful.
	Success bool `json:"success"`
	// Message provides details about the transfer.
	Message string `json:"message"`
	// TargetAgent is the name of the agent control was transferred to.
	TargetAgent string `json:"target_agent,omitempty"`
	// TransferType indicates the type of transfer performed.
	TransferType string `json:"transfer_type"`
}

// Tool implements the transfer_to_agent functionality.
type Tool struct {
	availableAgents []agent.Info
}

// New creates a new transfer_to_agent tool with the provided agent information.
func New(agents []agent.Info) *Tool { _ = "STUB: not implemented"; return nil }

// findAgentInfo finds agent information by name.
// Returns nil if no agent with the given name is found.
func (t *Tool) findAgentInfo(name string) *agent.Info { _ = "STUB: not implemented"; return nil }

// Declaration implements the tool.Tool interface.
func (t *Tool) Declaration() *tool.Declaration {
	_ = "STUB: not implemented"
	// Build detailed agent descriptions.
	return nil
}

// Call implements the tool.CallableTool interface.
func (t *Tool) Call(ctx context.Context, jsonArgs []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Find the target agent information.

// Get invocation from context.

// Set transfer information in the invocation with just the agent name.
