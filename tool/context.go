//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package tool

import "context"

// ContextKeyToolCallID is the context key type for tool call ID.
// It's exported so that the framework can inject tool call ID into context.
type ContextKeyToolCallID struct{}

type contextKeyStructuredStreamErrors struct{}
type contextKeyFinalResultChunks struct{}

// ToolCallIDFromContext retrieves tool call ID from context.
// Returns the tool call ID and true if found, empty string and false
// otherwise.
func ToolCallIDFromContext(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// WithStructuredStreamErrors marks a streamable-tool invocation as expecting
// structured error chunks instead of plain text fallback content.
func WithStructuredStreamErrors(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// StructuredStreamErrorsFromContext reports whether structured stream error
// chunks are enabled for the current streamable-tool invocation.
func StructuredStreamErrorsFromContext(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// WithFinalResultChunks marks a streamable-tool invocation as expecting final
// result chunks for framework-managed completion handling.
func WithFinalResultChunks(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// FinalResultChunksFromContext reports whether final result chunks are enabled
// for the current streamable-tool invocation.
func FinalResultChunksFromContext(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}
