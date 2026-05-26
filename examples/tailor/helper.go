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
	anthropicsdk "github.com/anthropics/anthropic-sdk-go"
	"github.com/ollama/ollama/api"
	openaisdk "github.com/openai/openai-go"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

func long(s string) string { _ = "STUB: not implemented"; return "" }

func repeat(s string, n int) string { _ = "STUB: not implemented"; return "" }

// summarizeMessages returns a concise multi-line preview of messages with truncation.
func summarizeMessages(msgs []model.Message, maxItems int) string {
	_ = "STUB: not implemented"
	return ""
}

// replace newlines to keep one-line per message

// summarizeMessagesHeadTail shows head and tail messages with omitted middle count.
func summarizeMessagesHeadTail(msgs []model.Message, headCount, tailCount int) string {
	_ = "STUB: not implemented"
	return ""
}

// All messages fit, show them all.

// Show head messages.

// Show omitted count.

// Show tail messages.

func firstTextPart(m model.Message) string { _ = "STUB: not implemented"; return "" }

func firstNonEmpty(values ...string) string { _ = "STUB: not implemented"; return "" }

func truncate(s string, max int) string { _ = "STUB: not implemented"; return "" }

// convertFromOpenAIMessages converts OpenAI SDK messages back to model.Message format.
// This is a simplified conversion that extracts the basic content for token counting.
func convertFromOpenAIMessages(openaiMsgs []openaisdk.ChatCompletionMessageParamUnion) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func extractSystemContent(content openaisdk.ChatCompletionSystemMessageParamContentUnion) string {
	_ = "STUB: not implemented"
	return ""
}

func extractUserContent(content openaisdk.ChatCompletionUserMessageParamContentUnion) string {
	_ = "STUB: not implemented"
	return ""
}

func extractAssistantContent(content openaisdk.ChatCompletionAssistantMessageParamContentUnion) string {
	_ = "STUB: not implemented"
	return ""
}

func extractToolContent(content openaisdk.ChatCompletionToolMessageParamContentUnion) string {
	_ = "STUB: not implemented"
	return ""
}

// convertFromAnthropicMessages converts Anthropic SDK messages back to model.Message format.
// This is a simplified conversion that extracts the basic content for token counting.
func convertFromAnthropicMessages(
	anthropicMsgs []anthropicsdk.MessageParam,
	systemPrompts []anthropicsdk.TextBlockParam,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

// Add system messages first if present.

// Convert conversation messages.

// Extract content from message.

// Tool use blocks are not included in token count for simplification.

func convertFromOllamaMessages(ollamaMsgs []api.Message) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

// loadMessagesFromJSON loads messages from a JSON file in the format of input.json.
func loadMessagesFromJSON(filename string) ([]model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For assistant messages, preserve tool_calls if present.

// If there are tool_calls but no content, keep the message.
// The API will handle it correctly.

// Preserve the message with tool_calls.

// Skip unknown roles.
