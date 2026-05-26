//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package hunyuan

import (
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/model"
	imodel "trpc.group/trpc-go/trpc-agent-go/model/internal/model"
)

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
		tokenCounter: model.NewSimpleTokenCounter(),
	}
)

// Option is a function that configures a Hunyuan model.
type Option func(*options)

// WithSecretId sets the secret ID for Hunyuan API authentication.
func WithSecretId(secretId string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSecretKey sets the secret key for Hunyuan API authentication.
func WithSecretKey(secretKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBaseUrl sets the base URL for the Hunyuan server.
func WithBaseUrl(baseUrl string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHost sets the host for the Hunyuan server.
func WithHost(host string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHttpClient sets the HTTP client to use.
func WithHttpClient(client *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithChannelBufferSize sets the channel buffer size for the Hunyuan client, 256 by default.
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
