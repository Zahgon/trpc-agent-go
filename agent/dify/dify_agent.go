//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package dify provides an agent that can communicate with dify workflow or chatflow.
package dify

import (
	"context"

	"github.com/cloudernative/dify-sdk-go"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultStreamingChannelSize    = 1024
	defaultNonStreamingChannelSize = 10
)

// DifyMode represents the Dify service mode
type DifyMode string

const (
	// ModeChatflow represents Dify chatflow mode (default)
	ModeChatflow DifyMode = "chatflow"
	// ModeWorkflow represents Dify workflow mode
	ModeWorkflow DifyMode = "workflow"
)

// DifyAgent is an agent that communicates with a remote Dify service.
type DifyAgent struct {
	// options
	baseUrl           string // dify base url
	apiSecret         string // dify api secret
	name              string
	description       string
	mode              DifyMode                     // Dify service mode: chatflow or workflow (default: chatflow)
	eventConverter    DifyEventConverter           // Custom event converters
	requestConverter  DifyRequestConverter         // Custom Dify chatflow request converter
	workflowConverter DifyWorkflowRequestConverter // Custom Dify workflow request converter

	streamingBufSize        int                  // Buffer size for streaming responses
	streamingRespHandler    StreamingRespHandler // Handler for streaming responses
	transferStateKey        []string             // Keys in session state to transfer to the A2A agent message by metadata
	enableStreaming         *bool                // Explicitly set streaming mode; nil means use agent card capability
	autoGenConversationName *bool                // Whether to auto generate conversation name

	difyClient        *dify.Client
	getDifyClientFunc func(*agent.Invocation) (*dify.Client, error)
}

// New creates a new DifyAgent.
func New(opts ...Option) (*DifyAgent, error) { _ = "STUB: not implemented"; return nil, nil }

// Default to chatflow mode

// Validate that required fields are set

// Validate mode

// sendErrorEvent sends an error event to the event channel
func (r *DifyAgent) sendErrorEvent(ctx context.Context, eventChan chan<- *event.Event,
	invocation *agent.Invocation, errorMessage string) {
	_ = "STUB: not implemented"
	return
}

// Run implements the Agent interface
func (r *DifyAgent) Run(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// shouldUseStreaming determines whether to use streaming protocol
func (r *DifyAgent) shouldUseStreaming(invocation *agent.Invocation) bool {
	_ = "STUB: not implemented"
	// Per-run override.
	return false
}

// If explicitly set via option, use that value

// Default to non-streaming if capabilities are not specified

// buildDifyRequest constructs Dify request from invocation
func (r *DifyAgent) buildDifyRequest(
	ctx context.Context,
	invocation *agent.Invocation,
	isStream bool,
) (*dify.ChatMessageRequest,
	error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Transfer additional state keys

// processStreamEvent processes a single stream event and returns the content to aggregate
func (r *DifyAgent) processStreamEvent(
	ctx context.Context,
	streamEvent dify.ChatMessageStreamChannelResponse,
	invocation *agent.Invocation,
) (*event.Event, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// Handle nil event (e.g., when Answer is empty)

// Aggregate content from delta

// buildStreamingRequest builds and sends streaming request to Dify chatflow
func (r *DifyAgent) buildStreamingRequest(
	ctx context.Context,
	invocation *agent.Invocation,
) (<-chan dify.ChatMessageStreamChannelResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildWorkflowStreamingRequest builds and sends streaming request to Dify workflow
func (r *DifyAgent) buildWorkflowStreamingRequest(
	ctx context.Context,
	invocation *agent.Invocation,
	eventChan chan<- *event.Event,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Transfer additional state keys

// Track workflow_run_id from streaming response for tracing correlation with Dify logs

// Convert workflow streaming response to event
// Extract text from outputs

// Try to get answer from common output fields

// Send final aggregated event with Dify-assigned workflow_run_id for proper tracing

// sendFinalStreamingEvent sends the final aggregated event for streaming
func (r *DifyAgent) sendFinalStreamingEvent(
	ctx context.Context,
	eventChan chan<- *event.Event,
	invocation *agent.Invocation,
	aggregatedContent string,
	messageID string,
) {
	_ = "STUB: not implemented"
	return
}

// runStreaming handles streaming communication
func (r *DifyAgent) runStreaming(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle workflow and chatflow differently due to different SDK APIs

// Workflow uses callback-based streaming

// Chatflow uses channel-based streaming

// Dify SDK emits Err with io.EOF on normal stream close; ignore it.
// Only surface real streaming failures (e.g., token limit exceeded, server errors).

// Skip nil events (empty responses)

// Record the streaming event's MessageID to keep the final event consistent

// executeNonStreamingRequest executes a non-streaming Dify request
func (r *DifyAgent) executeNonStreamingRequest(
	ctx context.Context,
	invocation *agent.Invocation,
) (*dify.ChatMessageResponse, error) {
	_ = "STUB: not implemented"
	// Handle workflow mode
	return nil, nil
}

// Transfer additional state keys

// Convert WorkflowResponse to ChatMessageResponse
// Extract answer from workflow outputs

// Try to get answer from common output fields

// Handle chatflow mode

// convertAndEmitNonStreamingEvent converts result to event and emits it
func (r *DifyAgent) convertAndEmitNonStreamingEvent(
	ctx context.Context,
	eventChan chan<- *event.Event,
	invocation *agent.Invocation,
	result *dify.ChatMessageResponse,
) {
	_ = "STUB: not implemented"
	return
}

// runNonStreaming handles non-streaming A2A communication
func (r *DifyAgent) runNonStreaming(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tools implements the Agent interface
func (r *DifyAgent) Tools() []tool.Tool {
	_ = "STUB: not implemented"
	// Remote A2A agents don't expose tools directly
	// Tools are handled by the remote agent
	return nil
}

// Info implements the Agent interface
func (r *DifyAgent) Info() agent.Info { _ = "STUB: not implemented"; return *new(agent.Info) }

// SubAgents implements the Agent interface
func (r *DifyAgent) SubAgents() []agent.Agent {
	_ = "STUB: not implemented"
	// Remote A2A agents don't have sub-agents in the local context
	return nil
}

// FindSubAgent implements the Agent interface
func (r *DifyAgent) FindSubAgent(name string) agent.Agent {
	_ = "STUB: not implemented"
	// Remote A2A agents don't have sub-agents in the local context
	return *new(agent.Agent)
}

// getDifyClient returns a Dify client instance, preferring the custom getDifyClientFunc if set,
// otherwise creating a client with the default configuration.
func (r *DifyAgent) getDifyClient(
	invocation *agent.Invocation,
) (*dify.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
