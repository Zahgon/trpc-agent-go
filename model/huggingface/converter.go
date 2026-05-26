//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package huggingface

import (
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// convertRequest converts a model.Request to a HuggingFace ChatCompletionRequest.
func (m *Model) convertRequest(req *model.Request) (*ChatCompletionRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert messages.

// Convert tools.

// Convert structured output to response format.

// convertMessage converts a model.Message to a HuggingFace ChatMessage.
func convertMessage(msg model.Message) (ChatMessage, error) {
	_ = "STUB: not implemented"
	return *new(ChatMessage), nil
}

// Handle tool message fields.

// Convert content - prioritize Content string field.

// Check if all content parts are text.

// Single text content - use string format.

// Multiple parts or non-text content - use array format.

// Convert tool calls.

// convertContentPart converts a model.ContentPart to a HuggingFace ContentPart.
func convertContentPart(part model.ContentPart) (ContentPart, error) {
	_ = "STUB: not implemented"
	return *new(ContentPart), nil
}

// Convert image to image_url format.

// If data is provided, create a data URL.

// convertTool converts a tool.Tool to a HuggingFace Tool.
func convertTool(t tool.Tool) (Tool, error) {
	_ = "STUB: not implemented"
	// Get tool declaration.
	return *new(Tool), nil
}

// Convert parameters to map.

// convertToolCall converts a model.ToolCall to a HuggingFace ToolCall.
func convertToolCall(tc model.ToolCall) (ToolCall, error) {
	_ = "STUB: not implemented"
	return *new(ToolCall), nil
}

// convertResponse converts a HuggingFace ChatCompletionResponse to a model.Response.
func (m *Model) convertResponse(hfResp *ChatCompletionResponse) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

// Convert choices.

// Convert usage.

// convertChoice converts a HuggingFace ChatCompletionChoice to a model.Choice.
func convertChoice(choice ChatCompletionChoice) model.Choice {
	_ = "STUB: not implemented"
	return *new(model.Choice)
}

// convertChunk converts a HuggingFace ChatCompletionChunk to a model.Response.
func (m *Model) convertChunk(chunk *ChatCompletionChunk) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

// Convert choices.

// Convert usage.

// convertMessageToModel converts a HuggingFace ChatMessage to a model.Message.
func convertMessageToModel(hfMsg ChatMessage) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}

// Handle tool message fields.

// Convert content.

// String content.

// Array content.

// Typed array content.

// Convert tool calls.

// convertContentPartToModel converts a content part map to a model.ContentPart.
func convertContentPartToModel(partMap map[string]any) model.ContentPart {
	_ = "STUB: not implemented"
	return *new(model.ContentPart)
}

// convertContentPartStructToModel converts a ContentPart struct to a model.ContentPart.
func convertContentPartStructToModel(part ContentPart) model.ContentPart {
	_ = "STUB: not implemented"
	return *new(model.ContentPart)
}
