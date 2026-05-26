//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package weknora provides an agent that can communicate with WeKnora service.
package weknora

import (
	"context"
	"time"

	"github.com/Tencent/WeKnora/client"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultStreamingChannelSize = 1024
)

// WeKnoraAgent is an agent that communicates with a remote WeKnora service.
type WeKnoraAgent struct {
	baseUrl          string // weknora base url
	token            string // weknora token
	name             string
	description      string
	agentID          string
	timeout          time.Duration
	knowledgeBaseIDs []string
	webSearchEnabled bool

	weknoraClient        *client.Client
	getWeKnoraClientFunc func(*agent.Invocation) (*client.Client, error)
}

// New creates a new WeKnoraAgent.
func New(opts ...Option) (*WeKnoraAgent, error) { _ = "STUB: not implemented"; return nil, nil }

// sendErrorEvent sends an error event to the event channel
func (r *WeKnoraAgent) sendErrorEvent(ctx context.Context, eventChan chan<- *event.Event,
	invocation *agent.Invocation, errorMessage string) {
	_ = "STUB: not implemented"
	return
}

// Run implements the Agent interface
func (r *WeKnoraAgent) Run(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildWeKnoraRequest constructs WeKnora request from invocation
func (r *WeKnoraAgent) buildWeKnoraRequest(
	ctx context.Context,
	invocation *agent.Invocation,
) (*client.AgentQARequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// runStreaming handles streaming communication
func (r *WeKnoraAgent) runStreaming(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Session might not exist, try to create it

// Send final aggregated event

// sendFinalStreamingEvent sends the final aggregated event for streaming
func (r *WeKnoraAgent) sendFinalStreamingEvent(
	ctx context.Context,
	eventChan chan<- *event.Event,
	invocation *agent.Invocation,
	aggregatedContent string,
	aggregatedReasoning string,
) {
	_ = "STUB: not implemented"
	return
}

// Tools implements the Agent interface
func (r *WeKnoraAgent) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

// Info implements the Agent interface
func (r *WeKnoraAgent) Info() agent.Info { _ = "STUB: not implemented"; return *new(agent.Info) }

// SubAgents implements the Agent interface
func (r *WeKnoraAgent) SubAgents() []agent.Agent { _ = "STUB: not implemented"; return nil }

// FindSubAgent implements the Agent interface
func (r *WeKnoraAgent) FindSubAgent(name string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func (r *WeKnoraAgent) getWeKnoraClient(
	invocation *agent.Invocation,
) (*client.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func genSessionKey(sessionID string) string { _ = "STUB: not implemented"; return "" }
