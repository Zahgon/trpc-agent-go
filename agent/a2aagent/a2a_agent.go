//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package a2aagent provides an agent that can communicate with remote A2A agents.
package a2aagent

import (
	"context"
	"strings"
	"time"

	sdktrace "go.opentelemetry.io/otel/trace"

	"trpc.group/trpc-go/trpc-a2a-go/client"
	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-a2a-go/server"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	itelemetry "trpc.group/trpc-go/trpc-agent-go/internal/telemetry"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultStreamingChannelSize    = 1024
	defaultNonStreamingChannelSize = 10
	defaultUserIDHeader            = "X-User-ID"
)

// A2AAgent is an agent that communicates with a remote A2A agent via A2A protocol.
type A2AAgent struct {
	// options
	name                 string
	description          string
	agentCard            *server.AgentCard      // Agent card and resolution state
	agentURL             string                 // URL of the remote A2A agent
	eventConverter       A2AEventConverter      // Custom A2A event converters
	dataPartMappers      []A2ADataPartMapper    // Lightweight inbound DataPart mappers for default converter
	a2aMessageConverter  InvocationA2AConverter // Custom A2A message converters for requests
	extraA2AOptions      []client.Option        // Additional A2A client options
	streamingBufSize     int                    // Buffer size for streaming responses
	streamingRespHandler StreamingRespHandler   // Handler for streaming responses
	transferStateKey     []string               // Keys in session state to transfer to the A2A agent message by metadata
	buildMessageHook     BuildMessageHook       // Hook called after A2A message is built but before it is sent
	userIDHeader         string                 // HTTP header name to send UserID to A2A server
	enableStreaming      *bool                  // Explicitly set streaming mode; nil means use agent card capability

	a2aClient *client.A2AClient
}

// New creates a new A2AAgent.
func New(opts ...Option) (*A2AAgent, error) { _ = "STUB: not implemented"; return nil, nil }

// Normalize the URL to ensure it has a proper scheme

// Create A2A client first

// If agent card is not set, fetch it using A2A client's GetAgentCard method

// Set name and description from agent card if not already set

// Normalize the agent card URL to ensure it has a proper scheme

// Rebuild a2a client if URL changed

// sendErrorEvent sends an error event to the event channel.
func (r *A2AAgent) sendErrorEvent(
	ctx context.Context,
	eventChan chan<- *event.Event,
	invocation *agent.Invocation,
	err error,
) *model.ResponseError {
	_ = "STUB: not implemented"
	return nil
}

// validateA2ARequestOptions validates that all A2A request options are of the correct type
func (r *A2AAgent) validateA2ARequestOptions(invocation *agent.Invocation) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *A2AAgent) setupInvocation(invocation *agent.Invocation) { _ = "STUB: not implemented"; return }

// Run implements the Agent interface
func (r *A2AAgent) Run(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate A2A request options early

// shouldUseStreaming determines whether to use streaming protocol.
//
// Priority:
//  1. Per-run override (agent.WithStream / invocation.RunOptions.Stream)
//  2. Agent option (WithEnableStreaming)
//  3. Agent card capability
//  4. Default false
func (r *A2AAgent) shouldUseStreaming(invocation *agent.Invocation) bool {
	_ = "STUB: not implemented"
	// Per-run override.
	return false
}

// If explicitly set via option, use that value

// Otherwise check if agent card supports streaming

// Default to non-streaming if capabilities are not specified

// buildA2AMessage constructs A2A message from session events.
// It assembles a middleware chain around the base converter:
//
//	transferStateKey → user hook → base converter
//
// transferStateKey is the outermost layer so it always runs even if
// the user hook short-circuits (skips calling next).
func (r *A2AAgent) buildA2AMessage(invocation *agent.Invocation, isStream bool) (*protocol.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Base converter function.

// User hook layer wraps the base converter.

// Built-in layer (outermost): transfer state keys into message metadata.
// Placed after hook so it always runs regardless of hook behavior.

// wrapWithTransferState returns a middleware that injects transferStateKey values
// from RuntimeState into the message metadata after calling next.
//
// Supported patterns:
//   - "*"        — transfer all keys
//   - "prefix*"  — transfer keys with the given prefix (e.g. "user.*" or "user*")
//   - "*suffix"  — transfer keys with the given suffix (e.g. "*.id" or "*id")
//   - "exact"    — transfer only the exact key
func (r *A2AAgent) wrapWithTransferState(next ConvertToA2AMessageFunc) ConvertToA2AMessageFunc {
	_ = "STUB: not implemented"
	return *new(ConvertToA2AMessageFunc)
}

// matchStateKeys copies keys from src to dst that match the given pattern.
func matchStateKeys(pattern string, src map[string]any, dst map[string]any) {
	_ = "STUB: not implemented"
	return
}

// runStreaming handles streaming A2A communication
func (r *A2AAgent) runStreaming(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// executeStreaming executes the streaming A2A communication workflow.
func (r *A2AAgent) executeStreaming(ctx context.Context, invocation *agent.Invocation, eventChan chan<- *event.Event) {
	_ = "STUB: not implemented"
	return
}

// buildRequestOptions constructs A2A request options from invocation.
func (r *A2AAgent) buildRequestOptions(ctx context.Context, invocation *agent.Invocation) []client.RequestOption {
	_ = "STUB: not implemented"
	return nil
}

// Add UserID header if session has UserID

// Propagate trace context via HTTP headers (W3C Trace Context).

type streamingEventResult struct {
	responseID        string
	aggregatedContent string
	terminalError     *model.ResponseError
}

// processStreamingEvents processes streaming events and aggregates content.
// Returns the response ID, aggregated content, and terminal error state.
func (r *A2AAgent) processStreamingEvents(
	ctx context.Context,
	invocation *agent.Invocation,
	eventChan chan<- *event.Event,
	streamChan <-chan protocol.StreamingMessageEvent,
) streamingEventResult {
	_ = "STUB: not implemented"
	return *new(streamingEventResult)
}

// flushBufferedContent emits buffered streaming text as a complete assistant
// message before forwarding a non-partial event such as a tool call or tool
// response. This preserves the original turn order in session history.
func (r *A2AAgent) flushBufferedContent(
	ctx context.Context,
	invocation *agent.Invocation,
	eventChan chan<- *event.Event,
	responseID string,
	anchorTimestamp time.Time,
	contentBuilder *strings.Builder,
) {
	_ = "STUB: not implemented"
	return
}

// aggregateEventContent aggregates content from event delta.
// Returns updated responseID and any terminal error that occurred.
func (r *A2AAgent) aggregateEventContent(
	ctx context.Context,
	invocation *agent.Invocation,
	eventChan chan<- *event.Event,
	evt *event.Event,
	responseID string,
	contentBuilder *strings.Builder,
) (string, *model.ResponseError) {
	_ = "STUB: not implemented"
	return "", nil
}

// emitFinalEvent emits the final completion event.
func (r *A2AAgent) emitFinalEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	eventChan chan<- *event.Event,
	responseID string,
	aggregatedContent string,
) {
	_ = "STUB: not implemented"
	return
}

// runNonStreaming handles non-streaming A2A communication
func (r *A2AAgent) runNonStreaming(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Construct A2A message from session

// Convert A2A response to multiple events

// Emit all events

func (r *A2AAgent) wrapEventChannelWithTelemetry(
	ctx context.Context,
	invocation *agent.Invocation,
	originalChan <-chan *event.Event,
	span sdktrace.Span,
	tracker *itelemetry.InvokeAgentTracker,
	startedSpan bool,
) <-chan *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// Tools implements the Agent interface
func (r *A2AAgent) Tools() []tool.Tool {
	_ = "STUB: not implemented"
	// Remote A2A agents don't expose tools directly
	// Tools are handled by the remote agent
	return nil
}

// Info implements the Agent interface
func (r *A2AAgent) Info() agent.Info { _ = "STUB: not implemented"; return *new(agent.Info) }

// SubAgents implements the Agent interface
func (r *A2AAgent) SubAgents() []agent.Agent {
	_ = "STUB: not implemented"
	// Remote A2A agents don't have sub-agents in the local context
	return nil
}

// FindSubAgent implements the Agent interface
func (r *A2AAgent) FindSubAgent(name string) agent.Agent {
	_ = "STUB: not implemented"
	// Remote A2A agents don't have sub-agents in the local context
	return *new(agent.Agent)
}

// GetAgentCard returns the resolved agent card
func (r *A2AAgent) GetAgentCard() *server.AgentCard {
	_ = "STUB: not implemented"

	// extractTraceHeaders extracts W3C Trace Context headers from ctx using the
	// globally registered OpenTelemetry propagator. Returns a map of header
	// key-value pairs (e.g. "traceparent" -> "00-..."). Returns nil when ctx
	// carries no valid span context.
	return nil
}

func extractTraceHeaders(ctx context.Context) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
