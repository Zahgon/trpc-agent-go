//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package a2aagent

import (
	"encoding/json"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-a2a-go/protocol"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// A2AEventConverter defines an interface for converting A2A protocol types to Event.
type A2AEventConverter interface {
	// ConvertToEvents converts an A2A protocol type to multiple Events.
	// In non-streaming mode, A2A server returns a Task with history containing
	// intermediate messages (tool calls, tool responses, etc.) and artifacts for final response.
	ConvertToEvents(result protocol.MessageResult, agentName string, invocation *agent.Invocation) ([]*event.Event, error)

	// ConvertStreamingToEvents converts a streaming A2A protocol type to Events.
	ConvertStreamingToEvents(result protocol.StreamingMessageEvent, agentName string, invocation *agent.Invocation) ([]*event.Event, error)
}

// InvocationA2AConverter defines an interface for converting invocations to A2A protocol messages.
type InvocationA2AConverter interface {
	// ConvertToA2AMessage converts an invocation to an A2A protocol Message.
	ConvertToA2AMessage(isStream bool, agentName string, invocation *agent.Invocation) (*protocol.Message, error)
}

type defaultA2AEventConverter struct {
	dataPartMappers []A2ADataPartMapper
}

func (d *defaultA2AEventConverter) ConvertToEvents(
	result protocol.MessageResult,
	agentName string,
	invocation *agent.Invocation,
) ([]*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Single message: build event from its parts

// Task with history: convert history messages first, then artifacts
// History contains intermediate messages (tool calls, tool responses, etc.)

// Artifacts contain the final response

// Handle unknown response types

// Mark the last event as done

func (d *defaultA2AEventConverter) ConvertStreamingToEvents(
	result protocol.StreamingMessageEvent,
	agentName string,
	invocation *agent.Invocation,
) ([]*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// submitted/completed updates without structured errors are control signals.

// Final artifact chunk is either an aggregated result or a termination signal,
// not incremental content for the user.

type defaultEventA2AConverter struct {
}

// ConvertToA2AMessage converts an event to an A2A protocol Message.
func (d *defaultEventA2AConverter) ConvertToA2AMessage(
	isStream bool,
	agentName string,
	invocation *agent.Invocation,
) (*protocol.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildA2AParts converts invocation message content and content parts to A2A protocol parts.
func (d *defaultEventA2AConverter) buildA2AParts(invocation *agent.Invocation) []protocol.Part {
	_ = "STUB: not implemented"
	return nil
}

// appendContentPart converts a single model.ContentPart and appends it to parts.
func appendContentPart(parts []protocol.Part, cp model.ContentPart) []protocol.Part {
	_ = "STUB: not implemented"
	return nil
}

func appendTextPart(parts []protocol.Part, cp model.ContentPart) []protocol.Part {
	_ = "STUB: not implemented"
	return nil
}

func appendImagePart(parts []protocol.Part, cp model.ContentPart) []protocol.Part {
	_ = "STUB: not implemented"
	return nil
}

func appendAudioPart(parts []protocol.Part, cp model.ContentPart) []protocol.Part {
	_ = "STUB: not implemented"
	return nil
}

func appendFilePart(parts []protocol.Part, cp model.ContentPart) []protocol.Part {
	_ = "STUB: not implemented"
	return nil
}

// buildRespEvent converts A2A response to tRPC event (used for both streaming and non-streaming mode)
func (d *defaultA2AEventConverter) buildRespEvent(
	isStreaming bool,
	msg *protocol.Message,
	agentName string,
	invocation *agent.Invocation) *event.Event {
	_ = "STUB: not implemented"

	// Parse A2A message parts to extract content and tool information
	return nil
}

// Create event with appropriate response structure

// parseResult holds the parsed information from A2A message parts
type parseResult struct {
	// textContent holds plain text content from TextParts
	textContent string

	// reasoningContent holds thought/reasoning content from TextParts with thought metadata
	reasoningContent string

	// toolCalls holds function call requests (assistant -> tool)
	toolCalls []model.ToolCall

	// toolResponses holds function response data (tool -> assistant)
	// Multiple tool responses can exist in a single message
	toolResponses []toolResponseData

	// codeExecution holds executable code content
	codeExecution string

	// codeExecutionResult holds code execution result content
	codeExecutionResult string

	// objectType holds the type of the object
	objectType string

	// tag holds the event tag from A2A message metadata
	tag string

	// responseID holds the original LLM Response.ID from A2A message metadata
	responseID string

	// taskState holds the remote task lifecycle state when present.
	taskState protocol.TaskState

	// responseError holds structured error fields reconstructed from metadata.
	responseError *model.ResponseError

	// stateDelta holds structured state updates reconstructed from A2A metadata.
	stateDelta map[string][]byte

	// extensions holds custom event payloads reconstructed by DataPart mappers.
	extensions map[string]json.RawMessage
}

// toolResponseData holds tool response information
type toolResponseData struct {
	id      string
	name    string
	content string
}

func newDataPartMappingResult(result *parseResult) *A2ADataPartMappingResult {
	_ = "STUB: not implemented"
	return nil
}

func applyDataPartMappingResult(dst *parseResult, mapped *A2ADataPartMappingResult) {
	_ = "STUB: not implemented"
	return
}

// parseA2AMessageParts processes all parts in the A2A message and extracts content and tool information
func parseA2AMessageParts(msg *protocol.Message) *parseResult {
	_ = "STUB: not implemented"
	return nil
}

func parseA2AMessagePartsWithMappers(
	msg *protocol.Message,
	mappers []A2ADataPartMapper,
) *parseResult {
	_ = "STUB: not implemented"
	return nil
}

func flushParseResultText(
	result *parseResult,
	textBuilder *strings.Builder,
	reasoningBuilder *strings.Builder,
) {
	_ = "STUB: not implemented"
	return
}

// processTextPart processes a TextPart and returns its content and whether it's a thought
func processTextPart(part protocol.Part) (text string, isThought bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Check if this is a thought/reasoning content by looking at metadata
// Support both "thought" and "adk_thought" keys for ADK compatibility

// processDataPart processes a DataPart and updates the parseResult accordingly
func processDataPart(part protocol.Part, result *parseResult) { _ = "STUB: not implemented"; return }

func processDataPartWithMappers(
	part protocol.Part,
	result *parseResult,
	mappers []A2ADataPartMapper,
) {
	_ = "STUB: not implemented"
	return
}

// Use GetDataPartType to get the type with correct precedence (adk_type first, then type)
// GetDataPartType handles nil metadata internally

// processFunctionCall processes a function call DataPart and returns the ToolCall
func processFunctionCall(d *protocol.DataPart) *model.ToolCall {
	_ = "STUB: not implemented"
	return nil
}

// Validate that we have at least a name

// processFunctionResponse processes a function response DataPart and returns the response content and metadata
func processFunctionResponse(d *protocol.DataPart) (content string, id string, name string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// Extract tool response metadata

// Extract response content. Keep strings as-is, otherwise prefer JSON for
// structured values.

// extractStringField extracts a string value from data map, trying primary key first, then fallback key
func extractStringField(data map[string]any, primary, fallback string) string {
	_ = "STUB: not implemented"
	return ""
}

// processExecutableCode processes an executable code DataPart and returns the code content
func processExecutableCode(d *protocol.DataPart) string { _ = "STUB: not implemented"; return "" }

// processCodeExecutionResult processes a code execution result DataPart and returns the result content
func processCodeExecutionResult(d *protocol.DataPart) string { _ = "STUB: not implemented"; return "" }

// convertA2ARoleToModelRole converts A2A protocol role to internal model role
func convertA2ARoleToModelRole(role protocol.MessageRole) model.Role {
	_ = "STUB: not implemented"
	return *new(model.Role)
}

// Default to assistant for unknown roles

// buildEventResponse creates an event with the appropriate response structure
func buildEventResponse(
	isStreaming bool,
	messageID string,
	result *parseResult,
	invocation *agent.Invocation,
	agentName string,
	role protocol.MessageRole,
) *event.Event {
	_ = "STUB: not implemented"
	return nil

	// Restore tag from A2A message metadata if present
}

// Use llm_response_id from metadata when available (preserves original LLM Response.ID),
// fall back to messageID (which is ArtifactID in streaming, or Message.MessageID in unary).

func markGraphCompletionEvent(evt *event.Event, result *parseResult) {
	_ = "STUB: not implemented"
	return
}

// graph.execution is a terminal subgraph event. Preserve completion
// semantics so parent graph agent nodes can reconstruct final state.

// buildStreamingResponse creates a response for streaming mode.
// In streaming mode:
// - Tool calls and tool responses use Message (not Delta) since they are complete units
// - Text content uses Delta for incremental updates
func buildStreamingResponse(messageID string, result *parseResult, role protocol.MessageRole) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

// Tool call: use Message (tool calls are complete units, not streamed incrementally)

// Tool response: use Message (tool responses are complete units)

// Text content: use Delta for streaming incremental updates

// Convert A2A protocol role to internal model role

func terminalStreamingResponseError(result *parseResult) *model.ResponseError {
	_ = "STUB: not implemented"
	return nil
}

func streamingResponseContent(result *parseResult) string { _ = "STUB: not implemented"; return "" }

// Some A2A servers surface invocation errors as a regular message decorated
// with structured error metadata. Preserve that message as normal stream
// content so callers can drain the stream to EOF instead of short-circuiting.

func streamingResponseObjectType(result *parseResult) string { _ = "STUB: not implemented"; return "" }

// extractObjectType determines the response object type from parseResult.
// Priority: 1) objectType from message metadata (for third-party framework compatibility,
// as some frameworks like ADK include object type in metadata)
// 2) Infer from content type (toolCalls, codeExecution, codeExecutionResult)
// 3) Return empty string to let caller use default value
func extractObjectType(result *parseResult) string { _ = "STUB: not implemented"; return "" }

// buildNonStreamingResponse creates a response for non-streaming mode.
// In non-streaming mode, all content uses Message (not Delta).
func buildNonStreamingResponse(messageID string, result *parseResult, role protocol.MessageRole) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

// Tool call: assistant requesting tool execution

// Tool response: tool returning results

// Text content: final assistant response
// Only add if no tool calls (tool calls already include text content)

// If no content at all, add empty assistant message

func buildErrorResponse(
	messageID string,
	respErr *model.ResponseError,
	now time.Time,
) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

func taskResponseError(
	result *parseResult,
) *model.ResponseError {
	_ = "STUB: not implemented"
	return nil
}

func buildRecoverableErrorResponse(messageID string, result *parseResult, role protocol.MessageRole, now time.Time) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

func markTerminalStructuredErrorEvent(evt *event.Event, result protocol.StreamingMessageResult) {
	_ = "STUB: not implemented"
	return
}

func hasStructuredErrorMetadata(metadata map[string]any) bool {
	_ = "STUB: not implemented"
	return false
}

func nonStreamingResponseContent(result *parseResult) string { _ = "STUB: not implemented"; return "" }

func nonStreamingResponseObjectType(result *parseResult) string {
	_ = "STUB: not implemented"
	return ""
}

func taskFailureMessage(
	state protocol.TaskState,
) string {
	_ = "STUB: not implemented"
	return ""
}

func isTaskFailureState(
	state protocol.TaskState,
) bool {
	_ = "STUB: not implemented"
	return false
}

func taskStateFromMetadata(
	metadata map[string]any,
) protocol.TaskState {
	_ = "STUB: not implemented"
	return *new(protocol.TaskState)
}

func cloneTaskMetadata(
	metadata map[string]any,
	taskState protocol.TaskState,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// convertTaskToMessage converts a Task to a Message
func convertTaskToMessage(task *protocol.Task) *protocol.Message {
	_ = "STUB: not implemented"
	return nil
}

// Add artifacts if any

// convertTaskStatusToMessage converts a TaskStatusUpdateEvent to a Message
func convertTaskStatusToMessage(event *protocol.TaskStatusUpdateEvent) *protocol.Message {
	_ = "STUB: not implemented"
	return nil
}

// convertTaskArtifactToMessage converts a TaskArtifactUpdateEvent to a Message.
func convertTaskArtifactToMessage(event *protocol.TaskArtifactUpdateEvent) *protocol.Message {
	_ = "STUB: not implemented"
	return nil
}
