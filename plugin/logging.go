//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package plugin

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultLoggingPluginName = "logging"
)

type (
	agentStartTimeKey struct{}
	modelStartTimeKey struct{}
	toolStartTimeKey  struct{}
)

// Logging logs high-level lifecycle events for agents, tools, and models.
type Logging struct {
	name string
}

// NewLogging creates a Logging plugin with a default name.
func NewLogging() *Logging { _ = "STUB: not implemented"; return nil }

// NewNamedLogging creates a Logging plugin with a custom name.
func NewNamedLogging(name string) *Logging { _ = "STUB: not implemented"; return nil }

// Name implements Plugin.
func (p *Logging) Name() string {
	_ = "STUB: not implemented"

	// Register implements Plugin.
	return ""
}

func (p *Logging) Register(r *Registry) { _ = "STUB: not implemented"; return }

func (p *Logging) beforeAgent(
	ctx context.Context,
	args *agent.BeforeAgentArgs,
) (*agent.BeforeAgentResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Logging) afterAgent(
	ctx context.Context,
	args *agent.AfterAgentArgs,
) (*agent.AfterAgentResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Logging) beforeModel(
	ctx context.Context,
	_ *model.BeforeModelArgs,
) (*model.BeforeModelResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Logging) afterModel(
	ctx context.Context,
	args *model.AfterModelArgs,
) (*model.AfterModelResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Logging) beforeTool(
	ctx context.Context,
	args *tool.BeforeToolArgs,
) (*tool.BeforeToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Logging) afterTool(
	ctx context.Context,
	args *tool.AfterToolArgs,
) (*tool.AfterToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
