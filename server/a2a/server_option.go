//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package a2a

import (
	"context"
	"net/http"

	"trpc.group/trpc-go/trpc-a2a-go/auth"
	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	a2a "trpc.group/trpc-go/trpc-a2a-go/server"
	"trpc.group/trpc-go/trpc-a2a-go/taskmanager"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// serverUserIDHeader is the default header that a2a server get UserID of invocation.
var serverUserIDHeader = "X-User-ID"

// UserIDFromContext returns the user ID from the context.
func UserIDFromContext(ctx context.Context) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// NewContextWithUserID returns a new context with the user ID.
func NewContextWithUserID(ctx context.Context, userID string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// ProcessorBuilder returns a message processor for the given agent.
type ProcessorBuilder func(agent agent.Agent, sessionService session.Service) taskmanager.MessageProcessor

// ProcessMessageHook is a function that wraps the message processor with additional functionality.
type ProcessMessageHook func(next taskmanager.MessageProcessor) taskmanager.MessageProcessor

// TaskManagerBuilder returns a task manager for the given agent.
type TaskManagerBuilder func(processor taskmanager.MessageProcessor) taskmanager.TaskManager

// ResponseRewriter rewrites outbound A2A responses before they are returned or
// sent to the remote peer.
//
// RewriteUnary receives the request context and final unary result after
// server-side aggregation of converted events. It rewrites what the caller will
// actually receive, rather than every intermediate converter output.
//
// Returning nil drops the outbound result.
type ResponseRewriter interface {
	RewriteUnary(ctx context.Context, result protocol.UnaryMessageResult) protocol.UnaryMessageResult
	RewriteStreaming(ctx context.Context, result protocol.StreamingMessageResult) protocol.StreamingMessageResult
}

// ResponseRewriterFuncs adapts plain functions into a ResponseRewriter.
type ResponseRewriterFuncs struct {
	Unary     func(ctx context.Context, result protocol.UnaryMessageResult) protocol.UnaryMessageResult
	Streaming func(ctx context.Context, result protocol.StreamingMessageResult) protocol.StreamingMessageResult
}

// RewriteUnary implements ResponseRewriter.
func (f ResponseRewriterFuncs) RewriteUnary(
	ctx context.Context,
	result protocol.UnaryMessageResult,
) protocol.UnaryMessageResult {
	_ = "STUB: not implemented"
	return *new(protocol.UnaryMessageResult)
}

// RewriteStreaming implements ResponseRewriter.
func (f ResponseRewriterFuncs) RewriteStreaming(
	ctx context.Context,
	result protocol.StreamingMessageResult,
) protocol.StreamingMessageResult {
	_ = "STUB: not implemented"
	return *new(protocol.StreamingMessageResult)
}

// EventToA2APartMapper converts an agent event into additional A2A parts.
//
// Returning nil or empty parts means this mapper contributes nothing and the
// default converter continues with its normal behavior.
type EventToA2APartMapper func(ctx context.Context, event *event.Event) ([]protocol.Part, error)

type defaultAuthProvider struct {
	userIDHeader string
}

func (d *defaultAuthProvider) Authenticate(r *http.Request) (*auth.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type options struct {
	sessionService            session.Service
	agent                     agent.Agent
	runner                    runner.Runner
	enableStreaming           bool
	graphEventObjectAllowlist []string
	responseRewriter          ResponseRewriter
	streamingEventType        StreamingEventType
	agentCard                 *a2a.AgentCard
	processorBuilder          ProcessorBuilder
	processorHook             ProcessMessageHook
	taskManagerBuilder        TaskManagerBuilder
	runOptions                []agent.RunOption
	a2aToAgentConverter       A2AMessageToAgentMessage
	eventToA2AConverter       EventToA2AMessage
	eventPartMappers          []EventToA2APartMapper
	host                      string
	extraOptions              []a2a.Option
	errorHandler              ErrorHandler
	debugLogging              bool
	userIDHeader              string
	adkCompatibility          bool
	structuredTaskErrors      bool
}

// Option is a function that configures a Server.
type Option func(*options)

// StreamingEventType controls how the A2A server emits agent output events in
// streaming mode.
//
// By default, streaming output is emitted as TaskArtifactUpdateEvent.
// This follows the ADK pattern: artifacts for content, status for state
// changes.
type StreamingEventType int

const (
	// StreamingEventTypeTaskArtifactUpdate emits agent output as
	// TaskArtifactUpdateEvent (default).
	StreamingEventTypeTaskArtifactUpdate StreamingEventType = iota

	// StreamingEventTypeMessage emits agent output as Message.
	StreamingEventTypeMessage
)

// WithSessionService sets the session service to use.
func WithSessionService(service session.Service) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithAgent sets the agent to use.
// It is mutually exclusive with WithRunner.
func WithAgent(agent agent.Agent, enableStreaming bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRunner sets the runner to use.
// It is mutually exclusive with WithAgent and requires WithAgentCard.
func WithRunner(r runner.Runner) Option { _ = "STUB: not implemented"; return *new(Option) }

func normalizeMetadataKeys(keys []string) []string { _ = "STUB: not implemented"; return nil }

// WithAgentCard sets the agent card to use.
// Use BuildBasicAgentCard to derive a basic card from an agent when needed.
func WithAgentCard(agentCard a2a.AgentCard) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithProcessorBuilder sets the processor builder to use.
func WithProcessorBuilder(builder ProcessorBuilder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithProcessMessageHook sets the process message hook to use.
// The hook can be used to wrap the message processor with additional functionality.
func WithProcessMessageHook(hook ProcessMessageHook) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHost sets the host address for the A2A server's agent card URL.
// The host will be normalized to a complete URL and used by other agents to discover and communicate with this agent.
//
// Supported formats:
//   - "localhost:8080" → "http://localhost:8080"
//   - "example.com" → "http://example.com"
//   - "http://example.com/api/v1" → "http://example.com/api/v1" (used as-is)
//   - "https://example.com" → "https://example.com" (used as-is)
//   - "grpc://service:9090" → "grpc://service:9090" (custom schemes supported)
//
// If the URL contains a path (e.g., "http://example.com/api/v1"), the path will be
// automatically extracted and set as the base path for routing requests.
//
// Example:
//
//	server, _ := a2a.New(
//	    a2a.WithAgent(myAgent),
//	    a2a.WithHost("localhost:8080"),  // URL: "http://localhost:8080", basePath: ""
//	    // or
//	    a2a.WithHost("http://example.com/api/v1"),  // URL: "http://example.com/api/v1", basePath: "/api/v1"
//	)
func WithHost(host string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUserIDHeader sets the HTTP header name to extract UserID from requests.
// If not set, defaults to "X-User-ID".
func WithUserIDHeader(header string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExtraA2AOptions passes extra options to the underlying A2A server.
// For example, it can be combined with a2a.WithAgentCardHandler and
// NewAgentCardHandler(...) to serve a dynamically updated AgentCard.
func WithExtraA2AOptions(opts ...a2a.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTaskManagerBuilder sets the task manager builder to use.
func WithTaskManagerBuilder(builder TaskManagerBuilder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRunOptions appends additional run options for every agent invocation.
// These options are applied before the A2A message metadata is merged into RuntimeState.
// If both WithRunOptions and A2A message metadata set the same RuntimeState key,
// the A2A metadata value takes precedence (last-write-wins).
func WithRunOptions(runOpts ...agent.RunOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithA2AToAgentConverter sets the A2A message to agent message converter to use.
func WithA2AToAgentConverter(converter A2AMessageToAgentMessage) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Converter-related options.
//
// The options in this section control how A2A requests/responses are converted.
// Unless otherwise noted, options marked as "default event converter only"
// affect only the built-in EventToA2AMessage implementation created by
// buildProcessor.

// WithEventToA2AConverter sets the event to A2A message converter to use.
//
// Providing a custom converter bypasses the built-in event conversion behavior.
// The default-event-converter options below do not rewrite custom converter
// output unless their comments explicitly say they also affect server-generated
// metadata.
func WithEventToA2AConverter(converter EventToA2AMessage) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGraphEventObjectAllowlist configures which graph object types
// (`evt.Response.Object`) are forwarded through A2A.
//
// Default event converter only.
// Matching applies only when object type starts with `graph.`.
//   - default (option not set): only graph.execution is forwarded.
//   - exact rule: "graph.node.start"
//   - prefix rule: "graph.node.*" or "graph.node*" (trailing '*' means prefix match)
//   - suffix rule: "*step" or "*.step" (leading '*' means suffix match)
//   - wildcard rule: "*" (allow all graph.* object types)
func WithGraphEventObjectAllowlist(objectTypes ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithResponseRewriter rewrites outbound A2A results before they are returned
// or sent to the remote peer.
//
// This option affects:
//   - unary Message / Task results returned to the caller
//   - streaming Message / TaskArtifactUpdateEvent / TaskStatusUpdateEvent results
//   - server-generated final streaming completion events
//   - server-generated structured task error results
//   - messages returned by ErrorHandler
//
// For unary responses, the rewriter sees the final result returned by the A2A
// server after it aggregates converted events. For streaming responses, it sees
// each outbound streaming event immediately before send. The request context is
// passed through so rewriters can use request-scoped values for logging.
//
// Returning nil drops the outbound result.
func WithResponseRewriter(rewriter ResponseRewriter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithADKCompatibility enables ADK compatibility mode.
//
// This option affects the default event converter and server-generated task
// metadata/status updates. It does not rewrite metadata produced by a custom
// EventToA2AConverter.
//
// When enabled, metadata keys in A2A messages will use the "adk_" prefix
// (e.g., "adk_app_name", "adk_user_id", "adk_session_id") to be compatible
// with ADK (Agent Development Kit) Python implementation.
func WithADKCompatibility(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStreamingEventType configures which A2A protocol type is used to emit
// agent output in streaming mode.
//
// Default event converter only.
// This option affects streaming output events converted from agent events
// (assistant text/tool calls/code execution). Task status updates
// (submitted/completed) are still emitted as TaskStatusUpdateEvent.
func WithStreamingEventType(eventType StreamingEventType) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEventToA2APartMapper registers a lightweight event-to-part mapper on the
// default event converter.
//
// Built-in tool-call and code-execution handling still takes precedence. For
// regular text events, mapper-generated parts are appended after reasoning and
// content TextParts so natural-language output is preserved.
//
// The mapper is ignored when WithEventToA2AConverter is used to replace the
// converter entirely.
func WithEventToA2APartMapper(mapper EventToA2APartMapper) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDebugLogging sets the debug logging to use.
func WithDebugLogging(debug bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithErrorHandler sets a custom error handler.
func WithErrorHandler(handler ErrorHandler) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStructuredTaskErrors enables structured propagation of agent
// Response.Error values through A2A task status metadata.
func WithStructuredTaskErrors(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// ErrorHandler converts errors to user-friendly messages
type ErrorHandler func(ctx context.Context, msg *protocol.Message, err error) (*protocol.Message, error)

// DefaultErrorHandler provides intelligent error handling based on error type
func defaultErrorHandler(ctx context.Context, msg *protocol.Message, err error) (*protocol.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type singleMsgSubscriber struct {
	ch chan protocol.StreamingMessageEvent
}

func newSingleResultSubscriber(result protocol.StreamingMessageResult) *singleMsgSubscriber {
	_ = "STUB: not implemented"
	return nil
}

func (e *singleMsgSubscriber) Send(event protocol.StreamingMessageEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// Channel returns the channel of the task subscriber
func (e *singleMsgSubscriber) Channel() <-chan protocol.StreamingMessageEvent {
	_ = "STUB: not implemented"

	// Closed returns true if the task subscriber is closed
	return nil
}

func (e *singleMsgSubscriber) Closed() bool {
	_ = "STUB: not implemented"

	// Close close the task subscriber
	return false
}

func (e *singleMsgSubscriber) Close() { _ = "STUB: not implemented"; return }
