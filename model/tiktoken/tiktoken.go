//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package tiktoken provides a tiktoken-go based token counter implementation
// that is compatible with the root model.TokenCounter interface.
package tiktoken

import (
	"context"

	"github.com/tiktoken-go/tokenizer"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// Counter implements a tiktoken-based token counter compatible with model.TokenCounter.
// It uses a tokenizer.Codec to encode message text and counts tokens as the
// length of the returned token slice.
type Counter struct {
	encoding tokenizer.Codec
}

// New creates a tiktoken-based counter.
//
// Parameters:
//   - modelName: OpenAI model name (e.g., "gpt-4o"). The tokenizer is chosen with tokenizer.ForModel.
//     If the model is not supported, falls back to cl100k_base.
//
// Returns:
// - *Counter on success; error if codec initialization fails.
func New(modelName string) (*Counter, error) { _ = "STUB: not implemented"; return nil, nil }

// Fallback to cl100k_base for broad compatibility.

// CountTokens returns the token count for a single message using tiktoken-go.
// It encodes Message.Content, Message.ReasoningContent, text ContentParts, and ToolCalls.
func (c *Counter) CountTokens(_ context.Context, message model.Message) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Count tokens for tool calls.

// countToolCallTokens calculates the token count for a single tool call.
// It encodes the tool call's type, ID, function name, description, and arguments.
func (c *Counter) countToolCallTokens(toolCall model.ToolCall) (int, error) {
	_ = "STUB: not implemented"

	// Count tokens for tool call type (e.g., "function").
	return 0, nil
}

// Count tokens for tool call ID.

// Count tokens for function name.

// Count tokens for function description.

// Count tokens for function arguments (JSON string).

// CountTokensRange returns the token count for a range of messages using tiktoken-go.
// This is more efficient than calling CountTokens multiple times.
func (c *Counter) CountTokensRange(ctx context.Context, messages []model.Message, start, end int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
