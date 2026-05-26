//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package tool

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// NamedToolSet wraps a ToolSet to automatically prefix tool names with the toolset name.
// This prevents tool name conflicts when multiple toolsets provide tools with the same name.
type NamedToolSet struct {
	toolSet tool.ToolSet
}

// NewNamedToolSet creates a new named toolset wrapper.
// If the toolSet is already a NamedToolSet, it returns itself to avoid double-wrapping.
func NewNamedToolSet(toolSet tool.ToolSet) *NamedToolSet { _ = "STUB: not implemented"; return nil }

// Tools returns tools with names prefixed by the toolset name to avoid conflicts.
func (s *NamedToolSet) Tools(ctx context.Context) []tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

// Create tools with prefixed names to avoid conflicts

// Close implements the ToolSet interface.
func (s *NamedToolSet) Close() error { _ = "STUB: not implemented"; return nil }

// Name implements the ToolSet interface.
func (s *NamedToolSet) Name() string { _ = "STUB: not implemented"; return "" }

// NamedTool wraps an original tool with a prefixed name to avoid conflicts.
type NamedTool struct {
	original tool.Tool
	name     string
}

// Declaration returns the tool declaration with a prefixed name.
func (t *NamedTool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

// Original returns the underlying Tool instance wrapped by the NamedTool.
func (t *NamedTool) Original() tool.Tool {
	_ = "STUB: not implemented"

	// ToolSetName returns the source ToolSet name for runtime policy checks.
	return *new(tool.Tool)
}

func (t *NamedTool) ToolSetName() string {
	_ = "STUB: not implemented"

	// Call delegates to the original tool's Call method.
	return ""
}

func (t *NamedTool) Call(ctx context.Context, jsonArgs []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// StreamableCall delegates to the original tool's StreamableCall method.
func (t *NamedTool) StreamableCall(ctx context.Context, jsonArgs []byte) (*tool.StreamReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SkipSummarization delegates to the original tool when it implements
// a SkipSummarization() bool preference; otherwise returns false.
func (t *NamedTool) SkipSummarization() bool { _ = "STUB: not implemented"; return false }
