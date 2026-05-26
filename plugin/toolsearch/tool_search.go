//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package toolsearch provides a Tool Search plugin.
package toolsearch

import (
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/plugin"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// ToolSearch uses an LLM to select relevant tools before the main
// model call by mutating `args.Request.Tools` in a BeforeModel callback.
type ToolSearch struct {
	name          string
	searcher      searcher
	maxTools      int
	alwaysInclude []string
	failOpen      bool
}

const defaultToolSearchPluginName = "tool_search"

// New creates a new ToolSearch.
func New(m model.Model, opts ...Option) (*ToolSearch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Name implements plugin.Plugin.
func (s *ToolSearch) Name() string { _ = "STUB: not implemented"; return "" }

// Register implements plugin.Plugin.
func (s *ToolSearch) Register(r *plugin.Registry) { _ = "STUB: not implemented"; return }

// Callback returns a BeforeModel callback that performs tool selection.
func (s *ToolSearch) Callback() model.BeforeModelCallbackStructured {
	_ = "STUB: not implemented"
	return *new(model.BeforeModelCallbackStructured)
}

// Fallback to original full tool set (do not mutate req.Tools).

// If no tools are available for selection, nothing to do.

// Rebuild request tools map.

func requestFromBeforeModelArgs(args *model.BeforeModelArgs) *model.Request {
	_ = "STUB: not implemented"
	return nil
}

func (s *ToolSearch) validateAlwaysIncludeToolsExist(baseTools map[string]tool.Tool) error {
	_ = "STUB: not implemented"
	return nil
}

func sortedToolNames(tools map[string]tool.Tool) []string { _ = "STUB: not implemented"; return nil }

func (s *ToolSearch) buildCandidateTools(baseTools map[string]tool.Tool) map[string]tool.Tool {
	_ = "STUB: not implemented"
	// Prepare candidate tools for selection (exclude always-include).
	return nil
}

func lastUserMessage(messages []model.Message) (model.Message, error) {
	_ = "STUB: not implemented"
	return *new(model.Message), nil
}

func buildSelectedTools(
	baseTools map[string]tool.Tool,
	selected []string,
	alwaysInclude []string,
) map[string]tool.Tool {
	_ = "STUB: not implemented"
	return nil
}
