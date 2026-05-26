//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/plugin"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	demoPluginName = "demo_plugin"
	demoTag        = "plugin_demo"
	denyKeyword    = "/deny"

	assistantPrefix = "[plugin] "
	denyText        = "Blocked by Runner plugin (BeforeModel short-circuit)."
)

type demoPlugin struct {
	debug bool
}

func newDemoPlugin(debug bool) plugin.Plugin { _ = "STUB: not implemented"; return *new(plugin.Plugin) }

func (p *demoPlugin) Name() string { _ = "STUB: not implemented"; return "" }

func (p *demoPlugin) Register(r *plugin.Registry) { _ = "STUB: not implemented"; return }

func (p *demoPlugin) beforeAgent(
	ctx context.Context,
	args *agent.BeforeAgentArgs,
) (*agent.BeforeAgentResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *demoPlugin) afterAgent(
	ctx context.Context,
	args *agent.AfterAgentArgs,
) (*agent.AfterAgentResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *demoPlugin) beforeModel(
	ctx context.Context,
	args *model.BeforeModelArgs,
) (*model.BeforeModelResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func requestHasUserKeyword(req *model.Request, keyword string) bool {
	_ = "STUB: not implemented"
	return false
}

func denyResponse() *model.Response { _ = "STUB: not implemented"; return nil }

func (p *demoPlugin) beforeTool(
	ctx context.Context,
	args *tool.BeforeToolArgs,
) (*tool.BeforeToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *demoPlugin) afterTool(
	ctx context.Context,
	args *tool.AfterToolArgs,
) (*tool.AfterToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *demoPlugin) onEvent(
	_ context.Context,
	_ *agent.Invocation,
	e *event.Event,
) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addTag(e *event.Event, tag string) { _ = "STUB: not implemented"; return }

func addAssistantPrefix(e *event.Event, prefix string) { _ = "STUB: not implemented"; return }
