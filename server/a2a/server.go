//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package a2a provides utilities for creating a2a servers.
package a2a

import (
	"context"
	"net/http"

	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	a2a "trpc.group/trpc-go/trpc-a2a-go/server"
	"trpc.group/trpc-go/trpc-a2a-go/taskmanager"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// New creates a new a2a server.
func New(opts ...Option) (*a2a.A2AServer, error) { _ = "STUB: not implemented"; return nil, nil }

// Enable ADK compatibility by default.

// Default to ADK-style streaming: artifacts for content.

// Host is only required if we need to build an agent card
// If user provides a custom agent card, host is optional

func buildAgentCard(options *options) (a2a.AgentCard, error) {
	_ = "STUB: not implemented"
	return *new(a2a.AgentCard), nil
}

// buildRuntimeState makes a shallow copy of message metadata for RuntimeState.
func buildRuntimeState(metadata map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func cloneMetadata(metadata map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

func buildProcessor(
	agent agent.Agent,
	sessionService session.Service,
	serverIdentity string,
	options *options,
) (*messageProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use custom converters if provided, otherwise use defaults

func buildA2AServer(options *options) (*a2a.A2AServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create a task manager that wraps the session service

// Set default UserID header if not configured

// Extract base path from agent card URL for request routing.
// If the URL contains a path component (e.g., "http://example.com/api/v1"),
// it will be extracted and used as the base path for routing incoming requests.

// traceContextMiddleware extracts W3C Trace Context from HTTP headers and injects
// it into the request context. This enables distributed tracing across A2A
// agent boundaries.
type traceContextMiddleware struct{}

// Wrap implements the a2a.Middleware interface.
func (m *traceContextMiddleware) Wrap(next http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Extract trace context from HTTP headers using the global propagator

// Continue with the enriched context

// extractBasePath extracts the path component from a URL for request routing.
// It parses the URL and returns the path if the URL has a valid scheme.
//
// Examples:
//   - "http://example.com/api/v1" → "/api/v1"
//   - "https://example.com/docs" → "/docs"
//   - "grpc://service:9090/rpc" → "/rpc"
//   - "http://example.com" → "" (no path)
//   - "invalid-url" → "" (no scheme)
//
// The extracted path is used as the base path for routing incoming A2A requests.
func extractBasePath(urlStr string) string { _ = "STUB: not implemented"; return "" }

// Extract path if URL has a valid scheme

// No valid scheme, return empty string

// messageProcessor is the message processor for the a2a server.
type messageProcessor struct {
	runner               runner.Runner
	a2aToAgentConverter  A2AMessageToAgentMessage
	eventToA2AConverter  EventToA2AMessage
	errorHandler         ErrorHandler
	debugLogging         bool
	adkCompatibility     bool
	responseRewriter     ResponseRewriter
	streamingEventType   StreamingEventType
	structuredTaskErrors bool
	agentName            string
	runOptions           []agent.RunOption
}

func isFinalStreamingEvent(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

// The only truly final event is runner.completion
// This ensures we don't miss postprocessing events (code execution, etc.)

func buildFinalStreamingMetadata(evt *event.Event) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func finalStreamingResponseID(evt *event.Event) string { _ = "STUB: not implemented"; return "" }

func cloneStreamingMetadata(metadata map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func isStructuredTaskErrorEvent(
	evt *event.Event,
) bool {
	_ = "STUB: not implemented"
	return false
}

func taskErrorState(
	respErr *model.ResponseError,
) protocol.TaskState {
	_ = "STUB: not implemented"
	return *new(protocol.TaskState)
}

func buildTaskErrorMetadata(
	agentEvent *event.Event,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func buildTaskErrorMessage(
	taskID string,
	ctxID string,
	agentEvent *event.Event,
	metadata map[string]any,
) *protocol.Message {
	_ = "STUB: not implemented"
	return nil
}

func buildStructuredFailureTask(
	taskID string,
	ctxID string,
	history []protocol.Message,
	agentEvent *event.Event,
) *protocol.Task {
	_ = "STUB: not implemented"
	return nil
}

// handleDefaultError provides a fallback error handling mechanism
func (m *messageProcessor) handleError(
	ctx context.Context,
	msg *protocol.Message,
	streaming bool,
	err error,
) (*taskmanager.MessageProcessingResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *messageProcessor) handleStreamingProcessingError(
	ctx context.Context,
	msg *protocol.Message,
	subscriber taskmanager.TaskSubscriber,
	err error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ProcessMessage is the main entry point for processing messages.
func (m *messageProcessor) ProcessMessage(
	ctx context.Context,
	message protocol.Message,
	options taskmanager.ProcessOptions,
	handler taskmanager.TaskHandler,
) (*taskmanager.MessageProcessingResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// It should not reach here, if client transfers an empty ctx id, trpc-a2a-go will generate one

// Get user ID from auth context, or generate from context ID if not available
// This follows ADK pattern: use auth user if available, otherwise use A2A_USER_{context_id}

// Convert A2A message to agent message

// Apply user-defined runOptions first, then merge A2A metadata into RuntimeState.
// This avoids conflicts when user also sets WithRuntimeState in runOptions,
// since WithRuntimeState uses overwrite semantics.

// Overlay structured graph resume state (e.g. ResumeCommand) so that
// it takes precedence over the raw flattened metadata keys.

// Copy existing state to avoid mutating the shared map from WithRunOptions.

func (m *messageProcessor) processStreamingMessage(
	ctx context.Context,
	userID string,
	ctxID string,
	a2aMsg *protocol.Message,
	agentMsg *model.Message,
	handler taskmanager.TaskHandler,
	runnerOpts []agent.RunOption,
) (*taskmanager.MessageProcessingResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run the agent and get streaming events

// Start processing in a goroutine with error recovery

// Send error to subscriber before closing

// processAgentStreamingEvents handles streaming events from the agent runner using tunnel for batch processing.
func (m *messageProcessor) processAgentStreamingEvents(
	ctx context.Context,
	taskID string,
	userID string,
	sessionID string,
	a2aMsg *protocol.Message,
	agentMsgChan <-chan *event.Event,
	subscriber taskmanager.TaskSubscriber,
	handler taskmanager.TaskHandler,
) {
	_ = "STUB: not implemented"
	return
}

// define consume function

// run event tunnel

func (m *messageProcessor) abortStreamingOnError(
	ctx context.Context,
	a2aMsg *protocol.Message,
	subscriber taskmanager.TaskSubscriber,
	err error,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *messageProcessor) sendTaskSubmittedEvent(
	ctx context.Context,
	taskID string,
	userID string,
	sessionID string,
	a2aMsg *protocol.Message,
	subscriber taskmanager.TaskSubscriber,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Add ADK-compatible metadata if enabled.

func (m *messageProcessor) sendFinalArtifactEvent(
	ctx context.Context,
	taskID string,
	ctxID string,
	finalStreamingMetadata map[string]any,
	subscriber taskmanager.TaskSubscriber,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *messageProcessor) sendTaskCompletedEvent(
	ctx context.Context,
	taskID string,
	ctxID string,
	finalStreamingMetadata map[string]any,
	subscriber taskmanager.TaskSubscriber,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *messageProcessor) sendStreamingResult(
	ctx context.Context,
	subscriber taskmanager.TaskSubscriber,
	result protocol.StreamingMessageResult,
	errorMessage string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// processBatchStreamingEvents processes a batch of streaming events and sends them through msgChan.
func (m *messageProcessor) processBatchStreamingEvents(
	ctx context.Context,
	taskID string,
	a2aMsg *protocol.Message,
	batch []*event.Event,
	subscriber taskmanager.TaskSubscriber,
	terminalTaskError *bool,
	finalStreamingMetadata *map[string]any,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// continue processing

// Check context cancellation

// Convert event to A2A message for streaming

// Send message if conversion successful

// Continue processing - need more data

func (m *messageProcessor) processMessage(
	ctx context.Context,
	userID string,
	ctxID string,
	a2aMsg *protocol.Message,
	agentMsg *model.Message,
	runnerOpts []agent.RunOption,
) (*taskmanager.MessageProcessingResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Collect converted A2A messages

func (m *messageProcessor) processUnaryEvent(
	ctx context.Context,
	a2aMsg *protocol.Message,
	agentEvent *event.Event,
	eventCount int,
	messages *[]protocol.Message,
) error {
	_ = "STUB: not implemented"
	return nil
}

func shouldSkipRunnerCompletionEvent(
	agentEvent *event.Event,
	messages []protocol.Message,
) bool {
	_ = "STUB: not implemented"
	return false
}

// Keep runner-completion state delta, but merge it into the latest
// converted message to avoid turning metadata-only completion into
// the final artifact message.

func appendConvertedUnaryResult(
	messages *[]protocol.Message,
	convertedResult any,
) {
	_ = "STUB: not implemented"
	return
}

// Extract messages from task artifacts.

func buildMessageProcessingResult(
	a2aMsg *protocol.Message,
	ctxID string,
	messages []protocol.Message,
) *taskmanager.MessageProcessingResult {
	_ = "STUB: not implemented"
	// If only one message, return it directly.
	return nil
}

// Multiple messages: return a Task with history containing
// intermediate messages and the final message in artifacts.

func mergeRunnerCompletionStateDeltaIntoLastMessage(
	messages []protocol.Message,
	stateDelta map[string][]byte,
) bool {
	_ = "STUB: not implemented"
	return false
}

func cloneStateDeltaBytes(raw []byte) []byte { _ = "STUB: not implemented"; return nil }

// addTaskMetadata adds ADK-compatible metadata to task status update events.
// Only writes metadata when ADK compatibility is enabled, using ADK-prefixed keys
// (adk_app_name, adk_user_id, adk_session_id) for interoperability with ADK clients.
func (m *messageProcessor) addTaskMetadata(event *protocol.TaskStatusUpdateEvent, userID, sessionID string) {
	_ = "STUB: not implemented"
	return
}

func (m *messageProcessor) rewriteStreamingResult(
	ctx context.Context,
	result protocol.StreamingMessageResult,
) protocol.StreamingMessageResult {
	_ = "STUB: not implemented"
	return *new(protocol.StreamingMessageResult)
}

func (m *messageProcessor) rewriteUnaryResult(
	ctx context.Context,
	result protocol.UnaryMessageResult,
) protocol.UnaryMessageResult {
	_ = "STUB: not implemented"
	return *new(protocol.UnaryMessageResult)
}

func normalizeStreamingResult(
	result protocol.StreamingMessageResult,
) protocol.StreamingMessageResult {
	_ = "STUB: not implemented"
	return *new(protocol.StreamingMessageResult)
}

func normalizeUnaryResult(
	result protocol.UnaryMessageResult,
) protocol.UnaryMessageResult {
	_ = "STUB: not implemented"
	return *new(protocol.UnaryMessageResult)
}

func normalizeProtocolMessage(msg *protocol.Message) *protocol.Message {
	_ = "STUB: not implemented"
	return nil
}

func normalizeTask(task *protocol.Task) *protocol.Task { _ = "STUB: not implemented"; return nil }

func normalizeTaskArtifactUpdateEvent(
	event *protocol.TaskArtifactUpdateEvent,
) *protocol.TaskArtifactUpdateEvent {
	_ = "STUB: not implemented"
	return nil
}

func normalizeTaskStatusUpdateEvent(
	event *protocol.TaskStatusUpdateEvent,
) *protocol.TaskStatusUpdateEvent {
	_ = "STUB: not implemented"
	return nil
}

func normalizeArtifact(artifact *protocol.Artifact) *protocol.Artifact {
	_ = "STUB: not implemented"
	return nil
}

func normalizeMetadataMap(metadata map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func droppedUnaryMessageProcessingResult() *taskmanager.MessageProcessingResult {
	_ = "STUB: not implemented"
	return nil
}

func droppedStreamingMessageProcessingResult() *taskmanager.MessageProcessingResult {
	_ = "STUB: not implemented"
	return nil
}
