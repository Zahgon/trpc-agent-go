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
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// createAgentCallbacks creates agent callbacks for user context injection.
func (e *userContextExample) createAgentCallbacks() *agent.Callbacks {
	_ = "STUB: not implemented"
	return nil
}

// createToolCallbacks creates tool callbacks for authorization and audit.
func (e *userContextExample) createToolCallbacks() *tool.Callbacks {
	_ = "STUB: not implemented"
	return nil
}

// createBeforeAgentCallback creates the before agent callback.
func (e *userContextExample) createBeforeAgentCallback() agent.BeforeAgentCallbackStructured {
	_ = "STUB: not implemented"
	return *new(agent.BeforeAgentCallbackStructured)
}

// Inject user context into invocation state.
// In a real application, you would get this from request metadata,
// JWT token, session, etc.

// createAfterAgentCallback creates the after agent callback.
func (e *userContextExample) createAfterAgentCallback() agent.AfterAgentCallbackStructured {
	_ = "STUB: not implemented"
	return *new(agent.AfterAgentCallbackStructured)
}

// Print audit summary.

// Clean up audit log.

// Clean up user context.

// createBeforeToolCallback creates the before tool callback for authorization.
func (e *userContextExample) createBeforeToolCallback() tool.BeforeToolCallbackStructured {
	_ = "STUB: not implemented"
	return *new(tool.BeforeToolCallbackStructured)
}

// Get invocation from context.

// Get user context from invocation state.

// Check if user has permission to use this tool.

// Log failed authorization attempt.

// createAfterToolCallback creates the after tool callback for audit logging.
func (e *userContextExample) createAfterToolCallback() tool.AfterToolCallbackStructured {
	_ = "STUB: not implemented"
	return *new(tool.AfterToolCallbackStructured)
}

// Get invocation from context.

// Get user context from invocation state.

// Create audit entry.

// Append to audit log.

// appendAuditLog appends an audit entry to the invocation state.
func (e *userContextExample) appendAuditLog(inv *agent.Invocation, entry AuditEntry) {
	_ = "STUB: not implemented"
	return
}
