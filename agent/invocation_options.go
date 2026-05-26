//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package agent

import (
	"reflect"

	"trpc.group/trpc-go/trpc-agent-go/artifact"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// InvocationOptions is the options for the Invocation.
type InvocationOptions func(*Invocation)

// WithInvocationID set invocation id for the Invocation.
func WithInvocationID(id string) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationAgent set agent for the Invocation.
func WithInvocationAgent(agent Agent) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationBranch set branch for the Invocation.
func WithInvocationBranch(branch string) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationEndInvocation set endInvocation for the Invocation.
func WithInvocationEndInvocation(endInvocation bool) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationSession set session for the Invocation.
func WithInvocationSession(session *session.Session) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationSessionService set session service for the Invocation.
func WithInvocationSessionService(sessionService session.Service) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationModel set model for the Invocation.
func WithInvocationModel(model model.Model) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationMessage set message for the Invocation.
func WithInvocationMessage(message model.Message) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationRunOptions set runOptions for the Invocation.
func WithInvocationRunOptions(runOptions RunOptions) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationTransferInfo set transferInfo for the Invocation.
func WithInvocationTransferInfo(transferInfo *TransferInfo) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationStructuredOutput set structuredOutput for the Invocation.
func WithInvocationStructuredOutput(structuredOutput *model.StructuredOutput) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationStructuredOutputType set structuredOutputType for the Invocation.
func WithInvocationStructuredOutputType(structuredOutputType reflect.Type) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationMemoryService set memoryService for the Invocation.
func WithInvocationMemoryService(memoryService memory.Service) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationArtifactService set artifactService for the Invocation.
func WithInvocationArtifactService(artifactService artifact.Service) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationPlugins sets the PluginManager for this invocation.
func WithInvocationPlugins(pm PluginManager) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}

// WithInvocationEventFilterKey set eventFilterKey for the Invocation.
func WithInvocationEventFilterKey(key string) InvocationOptions {
	_ = "STUB: not implemented"
	return *new(InvocationOptions)
}
