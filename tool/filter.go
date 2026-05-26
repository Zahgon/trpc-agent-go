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

// FilterFunc is a function that filters tools based on a context and a tool.
type FilterFunc func(ctx context.Context, tool Tool) bool

// FilterTools filters tools from a list of tools based on a filter function.
func FilterTools(ctx context.Context, tools []Tool, filter FilterFunc) []Tool {
	_ = "STUB: not implemented"
	return nil
}

// FilterToolSet creates a new ToolSet that filters tools from the original ToolSet.
func FilterToolSet(toolset ToolSet, filter FilterFunc) ToolSet {
	_ = "STUB: not implemented"
	return *new(ToolSet)
}

// filteredToolSet wraps a ToolSet to filter its tools based on their names.
type filteredToolSet struct {
	original ToolSet
	filter   FilterFunc
}

// Tools returns filtered tools from the original ToolSet.
func (f *filteredToolSet) Tools(ctx context.Context) []Tool { _ = "STUB: not implemented"; return nil }

// Create new slice for filtered tools

// Close implements the ToolSet interface.
func (f *filteredToolSet) Close() error { _ = "STUB: not implemented"; return nil }

// Name implements the ToolSet interface.
func (f *filteredToolSet) Name() string { _ = "STUB: not implemented"; return "" }

// NewIncludeToolNamesFilter creates a FilterFunc that includes only the specified tool names.
func NewIncludeToolNamesFilter(names ...string) FilterFunc {
	_ = "STUB: not implemented"
	return *new(FilterFunc)
}

// NewExcludeToolNamesFilter creates a FilterFunc that excludes the specified tool names.
func NewExcludeToolNamesFilter(names ...string) FilterFunc {
	_ = "STUB: not implemented"
	return *new(FilterFunc)
}
