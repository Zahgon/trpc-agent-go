//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package ollama provides Ollama-compatible model implementations.
package ollama

import (
	"context"
	"net/http"
	"time"

	"github.com/ollama/ollama/api"
	"trpc.group/trpc-go/trpc-agent-go/model"
	imodel "trpc.group/trpc-go/trpc-agent-go/model/internal/model"
)

const (
	defaultChannelBufferSize = 256
	functionToolType         = "function"
	// OllamaHost is the environment variable for the Ollama host.
	OllamaHost = "OLLAMA_HOST"
)

var (
	// DefaultHost is the default Ollama host port.
	defaultPort = "11434"
	// DefaultHost is the default Ollama host.
	defaultHost = "http://localhost:11434"
)

// ChatRequestCallbackFunc is the function type for the chat request callback.
type ChatRequestCallbackFunc func(
	ctx context.Context,
	chatRequest *api.ChatRequest,
)

// ChatResponseCallbackFunc is the function type for the chat response callback.
type ChatResponseCallbackFunc func(
	ctx context.Context,
	chatRequest *api.ChatRequest,
	chatResponse *api.ChatResponse,
)

// ChatChunkCallbackFunc is the function type for the chat chunk callback.
type ChatChunkCallbackFunc func(
	ctx context.Context,
	chatRequest *api.ChatRequest,
	chatChunk *api.ChatResponse,
)

// ChatStreamCompleteCallbackFunc is the function type for the chat stream completion callback.
type ChatStreamCompleteCallbackFunc func(
	ctx context.Context,
	chatRequest *api.ChatRequest,
	streamErr error,
)

// options contains configuration options for creating an Ollama model.
type options struct {
	// Host URL for the Ollama server.
	host string
	// HTTP client for the Ollama client.
	httpClient *http.Client
	// Buffer size for response channels (default: 256)
	channelBufferSize int
	// Callback for the chat request.
	chatRequestCallback ChatRequestCallbackFunc
	// Callback for the chat response.
	chatResponseCallback ChatResponseCallbackFunc
	// Callback for the chat chunk.
	chatChunkCallback ChatChunkCallbackFunc
	// Callback for the chat stream completion.
	chatStreamCompleteCallback ChatStreamCompleteCallbackFunc
	// enableTokenTailoring enables automatic token tailoring based on model context window.
	enableTokenTailoring bool
	// tokenCounter count tokens for token tailoring.
	tokenCounter model.TokenCounter
	// tailoringStrategy defines the strategy for token tailoring.
	tailoringStrategy model.TailoringStrategy
	// maxInputTokens is the max input tokens for token tailoring.
	maxInputTokens int
	// contextWindow is the model context window size in tokens.
	contextWindow           int
	contextWindowConfigured bool
	// tokenTailoringConfig allows customization of token tailoring parameters.
	tokenTailoringConfig *model.TokenTailoringConfig
	// Additional options for Ollama API.
	options   map[string]any
	keepAlive *api.Duration
}

var (
	defaultOptions = options{
		channelBufferSize: defaultChannelBufferSize,
		httpClient:        http.DefaultClient,
		tokenTailoringConfig: &model.TokenTailoringConfig{
			ProtocolOverheadTokens: imodel.DefaultProtocolOverheadTokens,
			ReserveOutputTokens:    imodel.DefaultReserveOutputTokens,
			SafetyMarginRatio:      imodel.DefaultSafetyMarginRatio,
			InputTokensFloor:       imodel.DefaultInputTokensFloor,
			MaxInputTokensRatio:    imodel.DefaultMaxInputTokensRatio,
		},
		host:         defaultHost,
		tokenCounter: model.NewSimpleTokenCounter(),
	}
)

// Option is a function that configures an Ollama model.
type Option func(*options)

// WithHost sets the host URL for the Ollama server.
func WithHost(host string) Option { _ = "STUB: not implemented"; return *new(Option) }

// withHttpClient sets the HTTP client to use.
// The site is temporarily not open to the public, as we may implement injection of an internal http client.
func withHttpClient(client *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithChannelBufferSize sets the channel buffer size for the Ollama client, 256 by default.
func WithChannelBufferSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithChatRequestCallback sets the function to be called before sending a
// chat request. The callback runs synchronously in GenerateContent before
// the response goroutine starts. Start your own goroutine in the callback
// if asynchronous behavior is needed.
func WithChatRequestCallback(fn ChatRequestCallbackFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithChatResponseCallback sets the function to be called after receiving a chat response.
func WithChatResponseCallback(fn ChatResponseCallbackFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithChatChunkCallback sets the function to be called after receiving a chat chunk.
func WithChatChunkCallback(fn ChatChunkCallbackFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithChatStreamCompleteCallback sets the function to be called when
// streaming is completed. The callback runs synchronously before the
// terminal streaming result is surfaced to the caller.
func WithChatStreamCompleteCallback(fn ChatStreamCompleteCallbackFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEnableTokenTailoring enables automatic token tailoring based on model context window.
func WithEnableTokenTailoring(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxInputTokens sets only the input token limit for token tailoring.
func WithMaxInputTokens(limit int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContextWindow sets the model context window size in tokens for this
// model instance.
func WithContextWindow(tokens int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTokenCounter sets the TokenCounter used for token tailoring.
func WithTokenCounter(counter model.TokenCounter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTailoringStrategy sets the TailoringStrategy used for token tailoring.
func WithTailoringStrategy(strategy model.TailoringStrategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTokenTailoringConfig sets custom token tailoring budget parameters.
func WithTokenTailoringConfig(config *model.TokenTailoringConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithOptions sets additional options for Ollama API.
func WithOptions(opt map[string]any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithKeepAlive sets the keep alive duration for the Ollama API.
func WithKeepAlive(duration time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }
