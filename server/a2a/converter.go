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

	"trpc.group/trpc-go/trpc-a2a-go/protocol"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// A2AMessageToAgentMessage defines an interface for converting A2A protocol messages to Agent messages.
type A2AMessageToAgentMessage interface {
	// ConvertToAgentMessage converts an A2A protocol message to an Agent message.
	ConvertToAgentMessage(ctx context.Context, message protocol.Message) (*model.Message, error)
}

// EventToA2AUnaryOptions is the options for the EventToA2AMessage.
type EventToA2AUnaryOptions struct {
	CtxID string
}

// EventToA2AStreamingOptions is the options for the EventToA2AMessage.
type EventToA2AStreamingOptions struct {
	CtxID  string
	TaskID string
}

// EventToA2AMessage defines an interface for converting Agent events to A2A protocol messages.
type EventToA2AMessage interface {
	// ConvertToA2AMessage converts an Agent event to an A2A protocol message.
	ConvertToA2AMessage(
		ctx context.Context,
		event *event.Event,
		options EventToA2AUnaryOptions,
	) (protocol.UnaryMessageResult, error)

	// ConvertStreaming converts an Agent event to an A2A protocol message for streaming.
	ConvertStreamingToA2AMessage(
		ctx context.Context,
		event *event.Event,
		options EventToA2AStreamingOptions,
	) (protocol.StreamingMessageResult, error)
}

// defaultA2AMessageToAgentMessage is the default implementation of A2AMessageToAgentMessageConverter.
type defaultA2AMessageToAgentMessage struct{}

// ConvertToAgentMessage converts an A2A protocol message to an Agent message.
func (c *defaultA2AMessageToAgentMessage) ConvertToAgentMessage(
	ctx context.Context,
	message protocol.Message,
) (*model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Process all parts in the A2A message

// Only add to content string, not to contentParts
// to avoid duplication when converting back to A2A message

// Convert FilePart to model.ContentPart.
// The original content type is primarily preserved by metadata["content_type"].
// MimeType and Name are only fallback signals for compatibility.

// Create message with both content and content parts

// defaultEventToA2AMessage is the default implementation of EventToA2AMessageConverter.
type defaultEventToA2AMessage struct {
	// Enable ADK-compatible metadata keys (for example, "adk_type" instead
	// of "type").
	adkCompatibility          bool
	graphEventObjectAllowlist []string
	streamingEventType        StreamingEventType
	eventPartMappers          []EventToA2APartMapper
}

const graphObjectPrefix = "graph."

var defaultAllowedGraphObjectTypes = []string{
	graph.ObjectTypeGraphExecution,
}

// setMetadata writes value under the standard key, and additionally under the
// ADK-prefixed key when ADK compatibility is enabled.
func (c *defaultEventToA2AMessage) setMetadata(m map[string]any, key string, value any) {
	_ = "STUB: not implemented"
	return
}

// setPartTypeMetadata sets the DataPart type metadata.
func (c *defaultEventToA2AMessage) setPartTypeMetadata(dataPart *protocol.DataPart, typeValue string) {
	_ = "STUB: not implemented"
	return
}

// setThoughtMetadata sets the thought metadata on a TextPart.
func (c *defaultEventToA2AMessage) setThoughtMetadata(textPart *protocol.TextPart) {
	_ = "STUB: not implemented"
	return
}

func (c *defaultEventToA2AMessage) buildMessageMetadata(evt *event.Event) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func hasStructuredMetadata(metadata map[string]any) bool { _ = "STUB: not implemented"; return false }

// hasContentfulMetadata reports whether metadata contains fields that are
// meaningful enough to warrant emitting an otherwise-empty A2A message.
// A message that carries only llm_response_id (and nothing else) is not
// useful to downstream consumers, so we exclude that key from the check.
func hasContentfulMetadata(metadata map[string]any) bool { _ = "STUB: not implemented"; return false }

func matchesAllowedGraphObjectType(objectType string, allowedObjectTypes []string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *defaultEventToA2AMessage) shouldEmitEvent(evt *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// ConvertToA2AMessage converts an Agent event to an A2A protocol message.
// For non-streaming responses, it returns the full content including
// tool calls.
func (c *defaultEventToA2AMessage) ConvertToA2AMessage(
	ctx context.Context,
	event *event.Event,
	options EventToA2AUnaryOptions,
) (protocol.UnaryMessageResult, error) {
	_ = "STUB: not implemented"
	return *new(protocol.UnaryMessageResult), nil
}

// Additional safety check for choices array bounds.

// Check if this is a tool call event.

// Check if this is a code execution event.

// Fallback to plain content conversion.

// convertCodeExecutionToA2AMessage converts code execution events to A2A DataPart messages.
// This handles both code execution and code execution result events.
func (c *defaultEventToA2AMessage) convertCodeExecutionToA2AMessage(
	evt *event.Event,
) (protocol.UnaryMessageResult, error) {
	_ = "STUB: not implemented"
	return *new(protocol.UnaryMessageResult), nil
}

// convertContentToA2AMessage converts message content to A2A message.
// It creates a message with text parts containing the content.
func (c *defaultEventToA2AMessage) buildTextParts(msg model.Message) []protocol.Part {
	_ = "STUB: not implemented"
	return nil

	// Add reasoning content as a separate TextPart with thought metadata
	// Following ADK pattern: thought content is stored in TextPart metadata
}

// Add main content

// convertContentToA2AMessage converts message content to A2A message.
// It creates a message with text parts containing the content.
func (c *defaultEventToA2AMessage) convertContentToA2AMessage(
	ctx context.Context,
	event *event.Event,
) (protocol.UnaryMessageResult, error) {
	_ = "STUB: not implemented"
	return *new(protocol.UnaryMessageResult), nil
}

// ConvertStreamingToA2AMessage converts an Agent event to an A2A protocol
// message for streaming.
//
// For streaming responses, it converts delta content, tool calls, and code
// execution events into A2A streaming results. The concrete A2A type can be
// configured via WithStreamingEventType.
func (c *defaultEventToA2AMessage) ConvertStreamingToA2AMessage(
	ctx context.Context,
	evt *event.Event,
	options EventToA2AStreamingOptions,
) (protocol.StreamingMessageResult, error) {
	_ = "STUB: not implemented"
	return *new(protocol.StreamingMessageResult), nil
}

// Additional safety check for choices array bounds

// Check if this is a tool call event

func (c *defaultEventToA2AMessage) runEventPartMappers(
	ctx context.Context,
	evt *event.Event,
) ([]protocol.Part, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *defaultEventToA2AMessage) convertMetadataOnlyToA2AMessageResult(
	evt *event.Event,
) protocol.UnaryMessageResult {
	_ = "STUB: not implemented"
	return *new(protocol.UnaryMessageResult)
}

func (c *defaultEventToA2AMessage) convertMetadataOnlyToA2AStreamingMessage(
	evt *event.Event,
	options EventToA2AStreamingOptions,
) (protocol.StreamingMessageResult, bool) {
	_ = "STUB: not implemented"
	return *new(protocol.StreamingMessageResult), false
}

func (c *defaultEventToA2AMessage) convertPartsToA2AStreamingResult(
	evt *event.Event,
	options EventToA2AStreamingOptions,
	parts []protocol.Part,
) protocol.StreamingMessageResult {
	_ = "STUB: not implemented"
	return *new(protocol.StreamingMessageResult)
}

func (c *defaultEventToA2AMessage) convertPartsToA2AStreamingResultWithMetadata(
	evt *event.Event,
	options EventToA2AStreamingOptions,
	parts []protocol.Part,
	metadata map[string]any,
) protocol.StreamingMessageResult {
	_ = "STUB: not implemented"
	return *new(protocol.StreamingMessageResult)
}

// convertDeltaContentToA2AStreamingMessage converts delta content to an A2A
// streaming result.
func (c *defaultEventToA2AMessage) convertDeltaContentToA2AStreamingMessage(
	ctx context.Context,
	event *event.Event,
	options EventToA2AStreamingOptions,
) (protocol.StreamingMessageResult, error) {
	_ = "STUB: not implemented"
	return *new(protocol.StreamingMessageResult), nil
}

// isToolCallEvent checks if an event is related to tool calls.
// It filters out both tool call requests and tool call responses.
func isToolCallEvent(event *event.Event) bool { _ = "STUB: not implemented"; return false }

// Check if this event contains tool calls in the response choices

// Check for tool call requests (assistant making tool calls)

// Check for tool call responses (tool returning results)

// Check for tool ID in the message (indicates tool response)

func isCodeExecutionEvent(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

// Check if the event object type is code execution related

// convertToolCallToA2AMessage converts tool call events to A2A DataPart messages.
// This handles both tool call requests and tool call responses.
func (c *defaultEventToA2AMessage) convertToolCallToA2AMessage(
	event *event.Event,
) (protocol.UnaryMessageResult, error) {
	_ = "STUB: not implemented"
	return *new(protocol.UnaryMessageResult), nil
}

// Handle tool call requests (assistant making function calls)
// OpenAI returns tool calls in a single choice with multiple ToolCalls

// Convert ToolCall to map for DataPart

// Handle tool call responses (tool returning results)
// OpenAI returns each tool response in a separate choice

// Convert tool response to DataPart

// Pass content as-is without parsing
// Client will receive the raw response string and display it directly

// convertToolCallToA2AStreamingMessage converts tool call events to A2A streaming messages.
func (c *defaultEventToA2AMessage) convertToolCallToA2AStreamingMessage(
	event *event.Event,
	options EventToA2AStreamingOptions,
) (protocol.StreamingMessageResult, error) {
	_ = "STUB: not implemented"
	// First get the message parts using the unary converter
	return *new(protocol.StreamingMessageResult), nil
}

// convertCodeExecutionToA2AStreamingMessage converts code execution events to A2A streaming messages.
func (c *defaultEventToA2AMessage) convertCodeExecutionToA2AStreamingMessage(
	evt *event.Event,
	options EventToA2AStreamingOptions,
) (protocol.StreamingMessageResult, error) {
	_ = "STUB: not implemented"
	// First get the message parts using the unary converter
	return *new(protocol.StreamingMessageResult), nil
}

// convertFilePart converts a protocol.FilePart to one or more model.ContentPart values.
//
// Content type resolution order (highest to lowest priority):
//  1. FilePart.Metadata["content_type"] — set explicitly by trpc-agent-go clients
//  2. MimeType or common format value — "image/*"/"png" → ContentTypeImage,
//     "audio/*"/"mp3" → ContentTypeAudio
//  3. FilePart.Name — legacy fallback for older clients that used name="image"/"audio"
//
// FileWithBytes.Bytes is a base64-encoded string per the A2A spec; it is decoded here.
func convertFilePart(filePart *protocol.FilePart) []model.ContentPart {
	_ = "STUB: not implemented"
	// Resolve content type using metadata > mimeType > name (legacy).
	return nil
}

// Decode base64-encoded bytes per the A2A spec.

// Non-base64 content (e.g. plain text in tests): use raw bytes.

// Audio with URI and other file types all use ContentTypeFile with FileID.

// resolveFilePartContentType determines the logical content type of a FilePart.
//
// Priority:
//  1. Metadata["content_type"] (set by trpc-agent-go clients, unambiguous)
//  2. MimeType or common format value ("image/*", "audio/*", "png", "mp3")
//  3. Name field (legacy: older clients used name="image"/"audio" as a type hint)
func resolveFilePartContentType(filePart *protocol.FilePart) string {
	_ = "STUB: not implemented"
	// 1. Explicit metadata (highest priority, set by current client)
	return ""
}

// 2. Infer from MimeType or common format value

// 3. Legacy name-based fallback (older clients that used name="image"/"audio")

func inferContentTypeFromMimeType(mimeType string) string { _ = "STUB: not implemented"; return "" }
