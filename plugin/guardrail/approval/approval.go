//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package approval provides a runner-scoped tool approval plugin.
package approval

import (
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/plugin"
	"trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/approval/review"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// Plugin is the approval plugin implementation.
type Plugin struct {
	name              string
	reviewer          review.Reviewer
	defaultToolPolicy ToolPolicy
	toolPolicies      map[string]ToolPolicy
	tokenCounter      model.TokenCounter
}

// New creates a new approval plugin.
func New(options ...Option) (*Plugin, error) { _ = "STUB: not implemented"; return nil, nil }

// Name implements plugin.Plugin.
func (p *Plugin) Name() string {
	_ = "STUB: not implemented"

	// Register implements plugin.Plugin.
	return ""
}

func (p *Plugin) Register(r *plugin.Registry) { _ = "STUB: not implemented"; return }

func (p *Plugin) beforeTool() tool.BeforeToolCallbackStructured {
	_ = "STUB: not implemented"
	return *new(tool.BeforeToolCallbackStructured)
}

func (p *Plugin) resolveToolPolicy(toolName string) ToolPolicy {
	_ = "STUB: not implemented"
	return *new(ToolPolicy)
}

func requiresReviewer(opts *options) bool { _ = "STUB: not implemented"; return false }
