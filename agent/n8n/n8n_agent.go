//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package n8n provides an agent that communicates with n8n workflows via webhook.
package n8n

import (
	"context"
	"net/http"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultStreamingChannelSize    = 1024
	defaultNonStreamingChannelSize = 10
	// n8n workflows can involve multiple external API calls and AI model invocations,
	// so a generous default timeout is used to avoid premature cancellation.
	defaultHTTPTimeout  = time.Hour
	maxResponseBodySize = 1 << 20 // 1MB limit for response body reads
)

// N8nAgent is an agent that communicates with a remote n8n webhook.
type N8nAgent struct {
	webhookURL           string
	name                 string
	description          string
	authType             AuthType
	authConfig           *AuthConfig
	requestConverter     RequestConverter
	responseConverter    ResponseConverter
	streamingBufSize     int
	streamingRespHandler StreamingRespHandler
	enableStreaming      *bool
	httpClient           *http.Client
	getHTTPClientFunc    func(*agent.Invocation) (*http.Client, error)
	transferStateKey     []string
}

// New creates a new N8nAgent with the given options.
func New(opts ...Option) (*N8nAgent, error) { _ = "STUB: not implemented"; return nil, nil }

// Initialize default HTTP client if none provided, so the connection pool is reused.

// Run implements the Agent interface.
func (a *N8nAgent) Run(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tools implements the Agent interface.
func (a *N8nAgent) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

// Info implements the Agent interface.
func (a *N8nAgent) Info() agent.Info { _ = "STUB: not implemented"; return *new(agent.Info) }

// SubAgents implements the Agent interface.
func (a *N8nAgent) SubAgents() []agent.Agent { _ = "STUB: not implemented"; return nil }

// FindSubAgent implements the Agent interface.
func (a *N8nAgent) FindSubAgent(name string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func (a *N8nAgent) shouldUseStreaming(invocation *agent.Invocation) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *N8nAgent) getHTTPClient(invocation *agent.Invocation) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildHTTPRequest constructs the HTTP request for the n8n webhook.
func (a *N8nAgent) buildHTTPRequest(
	ctx context.Context,
	invocation *agent.Invocation,
	isStreaming bool,
) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *N8nAgent) applyAuth(req *http.Request) { _ = "STUB: not implemented"; return }

func (a *N8nAgent) sendErrorEvent(
	ctx context.Context,
	eventChan chan<- *event.Event,
	invocation *agent.Invocation,
	errorMessage string,
) {
	_ = "STUB: not implemented"
	return
}

func (a *N8nAgent) sendFinalStreamingEvent(
	ctx context.Context,
	eventChan chan<- *event.Event,
	invocation *agent.Invocation,
	aggregatedContent string,
) {
	_ = "STUB: not implemented"
	return
}

func (a *N8nAgent) runNonStreaming(
	ctx context.Context,
	invocation *agent.Invocation,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *N8nAgent) runStreaming(
	ctx context.Context,
	invocation *agent.Invocation,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use bufio.Reader instead of Scanner to avoid line size limits on large SSE chunks.

// Per SSE spec, if the value starts with a space, remove it (exactly one).

// Handler error is treated as fatal: partial aggregated content is
// intentionally discarded because the handler may have encountered
// malformed data, making the accumulated content unreliable.
