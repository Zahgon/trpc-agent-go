//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package toolcallid provides a plugin that canonicalizes final tool call IDs.
package toolcallid

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/model"
	pluginbase "trpc.group/trpc-go/trpc-agent-go/plugin"
)

const defaultPluginName = "tool_call_id"

type plugin struct {
	name string
}

// New creates a ToolCall ID plugin.
func New() pluginbase.Plugin { _ = "STUB: not implemented"; return *new(pluginbase.Plugin) }

func newPlugin() *plugin { _ = "STUB: not implemented"; return nil }

// Name implements plugin.Plugin.
func (p *plugin) Name() string { _ = "STUB: not implemented"; return "" }

// Register implements plugin.Plugin.
func (p *plugin) Register(r *pluginbase.Registry) { _ = "STUB: not implemented"; return }

func (p *plugin) afterModel(
	ctx context.Context,
	args *model.AfterModelArgs,
) (*model.AfterModelResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
