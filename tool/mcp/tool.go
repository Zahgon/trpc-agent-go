//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package mcp

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/tool"
	mcp "trpc.group/trpc-go/trpc-mcp-go"
)

// mcpToolResult wraps MCP tool result for backward compatibility.
// It marshals as the Content slice, but provides Meta access via GetMeta().
type mcpToolResult struct {
	Content []mcp.Content
	Meta    map[string]any
	IsError bool
}

// MarshalJSON implements json.Marshaler.
// It marshals only the Content slice for backward compatibility.
func (r *mcpToolResult) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GetMeta returns the metadata from the tool result.
func (r *mcpToolResult) GetMeta() map[string]any {
	_ = "STUB: not implemented"

	// GetCallbackResult returns the callback-facing result payload.
	// AfterTool callbacks should continue to receive the raw content slice.
	return nil
}

func (r *mcpToolResult) GetCallbackResult() any {
	_ = "STUB: not implemented"

	// RetryResultError reports whether the MCP result should be treated as a result-level failure.
	return *new(any)
}

func (r *mcpToolResult) RetryResultError() bool {
	_ = "STUB: not implemented"

	// mcpTool implements the Tool interface for MCP tools.
	return false
}

type mcpTool struct {
	mcpToolRef     *mcp.Tool
	inputSchema    *tool.Schema
	outputSchema   *tool.Schema
	sessionManager *mcpSessionManager
}

// newMCPTool creates a new MCP tool wrapper.
func newMCPTool(mcpToolData mcp.Tool, sessionManager *mcpSessionManager) *mcpTool {
	_ = "STUB: not implemented"
	return nil
}

// Convert MCP input schema to inner Schema.

// Convert MCP output schema to inner Schema.

// Call implements the Tool interface.
func (t *mcpTool) Call(ctx context.Context, jsonArgs []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Parse raw arguments.

// callOnce performs a single call to the MCP tool.
// Returns a wrapped result that marshals as Content for backward compatibility.
func (t *mcpTool) callOnce(ctx context.Context, arguments map[string]any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Wrap for backward compatibility: marshals as Content array, but Meta is accessible

// Declaration implements the Tool interface.
func (t *mcpTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }
