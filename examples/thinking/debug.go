//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"

	openaisdk "github.com/openai/openai-go"
)

// printChatRequestMessages prints the messages being sent to the model API.
// This helps verify whether reasoning_content is included in the request.
func printChatRequestMessages(_ context.Context, req *openaisdk.ChatCompletionNewParams) {
	_ = "STUB: not implemented"
	return
}

// Check for reasoning_content by marshaling to JSON.

// Also check tool calls.

// Fallback: marshal the whole message.

func getSystemContent(msg *openaisdk.ChatCompletionSystemMessageParam) string {
	_ = "STUB: not implemented"
	return ""
}

func getUserContent(msg *openaisdk.ChatCompletionUserMessageParam) string {
	_ = "STUB: not implemented"
	return ""
}

func getAssistantContent(msg *openaisdk.ChatCompletionAssistantMessageParam) string {
	_ = "STUB: not implemented"
	return ""
}

func truncateString(s string, maxLen int) string { _ = "STUB: not implemented"; return "" }

// checkReasoningContentInAssistantMessage checks if reasoning_content is present
// in the assistant message by marshaling to JSON and inspecting the result.
func checkReasoningContentInAssistantMessage(msg *openaisdk.ChatCompletionAssistantMessageParam) {
	_ = "STUB: not implemented"
	return
}

// Check if reasoning_content key exists in the JSON.
