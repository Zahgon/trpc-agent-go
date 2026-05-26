//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package huggingface provides HuggingFace-compatible model implementations.
package huggingface

import (
	"context"
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	// defaultChannelBufferSize is the default channel buffer size.
	defaultChannelBufferSize = 256
	// defaultBaseURL is the default HuggingFace API base URL.
	defaultBaseURL = "https://router.huggingface.co"
	// defaultAPIKeyEnvVar is the default environment variable for HuggingFace API key.
	defaultAPIKeyEnvVar = "HUGGINGFACE_API_KEY"
)

// ChatRequestCallbackFunc is the function type for the chat request callback.
type ChatRequestCallbackFunc func(
	ctx context.Context,
	chatRequest *ChatCompletionRequest,
)

// ChatResponseCallbackFunc is the function type for the chat response callback.
type ChatResponseCallbackFunc func(
	ctx context.Context,
	chatRequest *ChatCompletionRequest,
	chatResponse *ChatCompletionResponse,
)

// ChatChunkCallbackFunc is the function type for the chat chunk callback.
type ChatChunkCallbackFunc func(
	ctx context.Context,
	chatRequest *ChatCompletionRequest,
	chatChunk *ChatCompletionChunk,
)

// ChatStreamCompleteCallbackFunc is the function type for the chat stream completion callback.
// This callback is invoked when streaming is completely finished (success or error).
type ChatStreamCompleteCallbackFunc func(
	ctx context.Context,
	chatRequest *ChatCompletionRequest,
	streamErr error, // nil if streaming completed successfully
)

// options contains configuration options for creating a Model.
type options struct {
	// API key for the HuggingFace client.
	APIKey string
	// Base URL for the HuggingFace API. Default is https://router.huggingface.co.
	BaseURL string
	// Buffer size for response channels (default: 256).
	ChannelBufferSize int
	// HTTP client for making requests.
	HTTPClient *http.Client
	// Callback for the chat request.
	ChatRequestCallback ChatRequestCallbackFunc
	// Callback for the chat response.
	ChatResponseCallback ChatResponseCallbackFunc
	// Callback for the chat chunk.
	ChatChunkCallback ChatChunkCallbackFunc
	// Callback for the chat stream completion.
	ChatStreamCompleteCallback ChatStreamCompleteCallbackFunc
	// Extra headers to be added to HTTP requests.
	ExtraHeaders map[string]string
	// Extra fields to be added to the HTTP request body.
	ExtraFields map[string]any
	// EnableTokenTailoring enables automatic token tailoring based on model context window.
	EnableTokenTailoring bool
	// TokenCounter count tokens for token tailoring.
	TokenCounter model.TokenCounter
	// TailoringStrategy defines the strategy for token tailoring.
	TailoringStrategy model.TailoringStrategy
	// MaxInputTokens is the max input tokens for token tailoring.
	MaxInputTokens int
	// ContextWindow is the model context window size in tokens.
	ContextWindow int
	// TokenTailoringConfig allows customization of token tailoring parameters.
	TokenTailoringConfig *model.TokenTailoringConfig
}

var (
	defaultOptions = options{
		BaseURL:           defaultBaseURL,
		ChannelBufferSize: defaultChannelBufferSize,
		TokenCounter:      model.NewSimpleTokenCounter(),
		TokenTailoringConfig: &model.TokenTailoringConfig{
			ProtocolOverheadTokens: 512,
			ReserveOutputTokens:    2048,
			SafetyMarginRatio:      0.1,
			InputTokensFloor:       1024,
			OutputTokensFloor:      512,
			MaxInputTokensRatio:    0.8,
		},
	}
)

// Option is a function that configures a HuggingFace model.
type Option func(*options)

// WithAPIKey sets the API key for the HuggingFace client.
func WithAPIKey(key string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBaseURL sets the base URL for the HuggingFace API.
func WithBaseURL(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithChannelBufferSize sets the channel buffer size for the HuggingFace client.
func WithChannelBufferSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHTTPClient sets the HTTP client for making requests.
func WithHTTPClient(client *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithChatRequestCallback sets the function to be called before sending a
// chat request. The callback runs synchronously in GenerateContent before
// the response goroutine starts. Start your own goroutine in the callback
// if asynchronous behavior is needed.
func WithChatRequestCallback(fn ChatRequestCallbackFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithChatResponseCallback sets the function to be called after receiving a chat response.
// Used for non-streaming responses.
func WithChatResponseCallback(fn ChatResponseCallbackFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithChatChunkCallback sets the function to be called after receiving a chat chunk.
// Used for streaming responses.
func WithChatChunkCallback(fn ChatChunkCallbackFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithChatStreamCompleteCallback sets the function to be called when
// streaming is completed.
// Called for both successful and failed streaming completions.
func WithChatStreamCompleteCallback(fn ChatStreamCompleteCallbackFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithExtraHeaders sets extra headers to be added to HTTP requests.
func WithExtraHeaders(headers map[string]string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithExtraFields sets extra fields to be added to the HTTP request body.
// These fields will be included in every chat completion request.
func WithExtraFields(extraFields map[string]any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEnableTokenTailoring enables automatic token tailoring based on model context window.
// When enabled, the system will automatically calculate max input tokens using the model's.
// context window minus reserved tokens and protocol overhead.
func WithEnableTokenTailoring(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxInputTokens sets only the input token limit for token tailoring.
// The counter/strategy will be lazily initialized if not provided.
// Defaults to SimpleTokenCounter and MiddleOutStrategy.
func WithMaxInputTokens(limit int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContextWindow sets the model context window size in tokens for this
// model instance.
func WithContextWindow(tokens int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTokenCounter sets the TokenCounter used for token tailoring.
// If not provided and token limit is enabled, a SimpleTokenCounter will be used.
func WithTokenCounter(counter model.TokenCounter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTailoringStrategy sets the TailoringStrategy used for token tailoring.
// If not provided and token limit is enabled, a MiddleOutStrategy will be used.
func WithTailoringStrategy(strategy model.TailoringStrategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTokenTailoringConfig sets custom token tailoring budget parameters.
// This allows advanced users to fine-tune the token allocation strategy.
func WithTokenTailoringConfig(config *model.TokenTailoringConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
