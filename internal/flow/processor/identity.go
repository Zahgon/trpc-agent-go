//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package processor

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// IdentityRequestProcessor implements identity processing logic.
type IdentityRequestProcessor struct {
	// AgentName is the name of the agent.
	AgentName string
	// Description is the description of the agent.
	Description string

	addNameToInstruction bool
}

// Option is a function that can be used to configure the identity request processor.
type Option func(*IdentityRequestProcessor)

// WithAddNameToInstruction adds the agent name to the instruction if true.
func WithAddNameToInstruction(addNameToInstruction bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// NewIdentityRequestProcessor creates a new identity request processor.
func NewIdentityRequestProcessor(agentName, description string, opts ...Option) *IdentityRequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

// ProcessRequest implements the flow.RequestProcessor interface.
// It adds agent identity information to the request if provided.
func (p *IdentityRequestProcessor) ProcessRequest(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// Get agent name.

// Initialize messages slice if nil.

// Create identity message if we have name or description.

// Find existing system message or create new one

// There's already a system message, check if it contains identity

// Prepend identity to existing system message

// No existing system message, create new one

// containsIdentity checks if the given content already contains the identity.
func containsIdentity(content, identity string) bool { _ = "STUB: not implemented"; return false }
