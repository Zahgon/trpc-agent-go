//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package dify

import (
	"context"

	"github.com/cloudernative/dify-sdk-go"
	"trpc.group/trpc-go/trpc-a2a-go/protocol"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// DifyEventConverter defines an interface for converting Dify response to Event.
type DifyEventConverter interface {
	// ConvertToEvent converts an A2A protocol type to an Event.
	ConvertToEvent(
		resp *dify.ChatMessageResponse,
		agentName string,
		invocation *agent.Invocation,
	) *event.Event

	// ConvertStreamingToEvent converts a streaming A2A protocol type to an Event.
	ConvertStreamingToEvent(
		resp dify.ChatMessageStreamChannelResponse,
		agentName string,
		invocation *agent.Invocation,
	) *event.Event
}

// DifyRequestConverter defines an interface for converting invocations to Dify request messages.
type DifyRequestConverter interface {
	// ConvertToDifyRequest converts agent invocation to Dify request
	ConvertToDifyRequest(
		ctx context.Context,
		invocation *agent.Invocation,
		isStream bool,
	) (*dify.ChatMessageRequest, error)
}

// DifyWorkflowRequestConverter defines an interface for converting invocations to Dify workflow requests.
type DifyWorkflowRequestConverter interface {
	// ConvertToWorkflowRequest converts agent invocation to Dify workflow request
	ConvertToWorkflowRequest(
		ctx context.Context,
		invocation *agent.Invocation,
	) (dify.WorkflowRequest, error)
}

// defaultDifyEventConverter is the default implementation of DifyEventConverter.
// It converts Dify chatflow/workflow responses to internal event format.
type defaultDifyEventConverter struct {
}

// ConvertToEvent converts a Dify ChatMessageResponse to an internal Event.
// If resp is nil, it returns a default empty assistant message event.
func (d *defaultDifyEventConverter) ConvertToEvent(
	resp *dify.ChatMessageResponse,
	agentName string,
	invocation *agent.Invocation,
) (evt *event.Event) {
	_ = "STUB: not implemented"
	return nil
}

// set Dify response ID, ensure AG-UI translator can correctly identify the message

// ConvertStreamingToEvent converts a Dify streaming response to an internal Event.
// Returns nil if the response Answer is empty.
func (d *defaultDifyEventConverter) ConvertStreamingToEvent(
	resp dify.ChatMessageStreamChannelResponse,
	agentName string,
	invocation *agent.Invocation,
) (evt *event.Event) {
	_ = "STUB: not implemented"
	return nil
}

// set Dify return MessageID as Response.ID, ensure AG-UI translator
// can correctly trigger TextMessageStartEvent. If MessageID is empty, fall back in order.

// defaultEventDifyConverter is the default implementation of DifyRequestConverter.
// It converts agent invocations to Dify ChatMessageRequest format.
type defaultEventDifyConverter struct {
}

// ConvertToDifyRequest converts an agent invocation to a Dify ChatMessageRequest.
// It handles text, image, and file content parts from the invocation message.
func (d *defaultEventDifyConverter) ConvertToDifyRequest(
	ctx context.Context,
	invocation *agent.Invocation,
	isStream bool,
) (*dify.ChatMessageRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Enable streaming response mode

// Handle content parts if available

// Append text content to query

// For now, we can't directly handle images in Dify requests
// This would need to be handled based on specific Dify capabilities

// Similar to images, file handling depends on Dify capabilities

// Handle other content types as needed

// defaultWorkflowRequestConverter is the default implementation of DifyWorkflowRequestConverter.
// It converts agent invocations to Dify WorkflowRequest format.
type defaultWorkflowRequestConverter struct{}

// ConvertToWorkflowRequest converts an agent invocation to a Dify WorkflowRequest.
// It extracts query, image, and file inputs from the invocation message content parts.
func (d *defaultWorkflowRequestConverter) ConvertToWorkflowRequest(
	ctx context.Context,
	invocation *agent.Invocation,
) (dify.WorkflowRequest, error) {
	_ = "STUB: not implemented"
	return *new(dify.WorkflowRequest), nil
}

// Handle content parts if available

// workflow default mode

// extractTextFromParts extracts text content from protocol message parts
func extractTextFromParts(parts []protocol.Part) string { _ = "STUB: not implemented"; return "" }

// buildResponseForEvent creates a Response object based on streaming mode
func buildResponseForEvent(isStreaming bool, content string) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

// buildRespEvent converts A2A response to tRPC event
func (d *defaultDifyEventConverter) buildRespEvent(
	isStreaming bool,
	msg *protocol.Message,
	agentName string,
	invocation *agent.Invocation) *event.Event {
	_ = "STUB: not implemented"

	// Extract text content from parts
	return nil
}

// Create event with appropriate response

// convertTaskToMessage converts a Task to a Message
func convertTaskToMessage(task *protocol.Task) *protocol.Message {
	_ = "STUB: not implemented"
	return nil

	// Add artifacts if any
}

// convertTaskStatusToMessage converts a TaskStatusUpdateEvent to a Message
func convertTaskStatusToMessage(event *protocol.TaskStatusUpdateEvent) *protocol.Message {
	_ = "STUB: not implemented"
	return nil
}

// convertTaskArtifactToMessage converts a TaskArtifactUpdateEvent to a Message
func convertTaskArtifactToMessage(event *protocol.TaskArtifactUpdateEvent) *protocol.Message {
	_ = "STUB: not implemented"
	return nil
}
