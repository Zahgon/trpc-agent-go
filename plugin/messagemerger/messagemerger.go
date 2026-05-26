//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package messagemerger provides a runner-scoped plugin that merges
// consecutive messages with the same role.
package messagemerger

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/plugin"
)

// messageMergerPlugin merges consecutive system, user, and assistant messages before a
// model request is sent.
//
// Tool messages are intentionally preserved one-by-one because their ToolID and
// ToolName fields carry per-call semantics that must not be collapsed.
type messageMergerPlugin struct {
	name      string
	separator string
}

// New creates a new message merger plugin.
func New(options ...Option) plugin.Plugin { _ = "STUB: not implemented"; return *new(plugin.Plugin) }

// Name implements plugin.Plugin.
func (p *messageMergerPlugin) Name() string {
	_ = "STUB: not implemented"

	// Register implements plugin.Plugin.
	return ""
}

func (p *messageMergerPlugin) Register(r *plugin.Registry) { _ = "STUB: not implemented"; return }

func (p *messageMergerPlugin) beforeModel(
	_ context.Context,
	args *model.BeforeModelArgs,
) (*model.BeforeModelResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
