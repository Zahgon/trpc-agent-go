//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package provider

import (
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/model/anthropic"
	"trpc.group/trpc-go/trpc-agent-go/model/gemini"
	"trpc.group/trpc-go/trpc-agent-go/model/hunyuan"
	"trpc.group/trpc-go/trpc-agent-go/model/ollama"
	"trpc.group/trpc-go/trpc-agent-go/model/openai"
)

// Option configures how a model instance should be constructed.
type Option func(*Options)

// Options contains resolved settings used when constructing provider-backed models.
type Options struct {
	ProviderName         string                      // ProviderName is the provider identifier passed to Model.
	ModelName            string                      // ModelName is the concrete model identifier.
	Variant              string                      // Variant selects OpenAI-compatible model behavior.
	APIKey               string                      // APIKey holds the credential used for downstream SDK initialization.
	BaseURL              string                      // BaseURL overrides the default endpoint when specified.
	HTTPClientName       string                      // HTTPClientName is the logical name applied to the HTTP client.
	HTTPClientTransport  http.RoundTripper           // HTTPClientTransport allows customizing the HTTP transport.
	Callbacks            *Callbacks                  // Callbacks captures provider-specific callback hooks.
	ChannelBufferSize    *int                        // ChannelBufferSize is the response channel buffer size.
	Headers              map[string]string           // Headers are appended to outbound provider requests.
	ExtraFields          map[string]any              // ExtraFields are serialized into provider-specific request payloads.
	EnableTokenTailoring *bool                       // EnableTokenTailoring toggles automatic token tailoring.
	MaxInputTokens       *int                        // MaxInputTokens sets the maximum input tokens for token tailoring.
	ContextWindow        *int                        // ContextWindow sets the model context window size in tokens.
	TokenCounter         model.TokenCounter          // TokenCounter provides a custom token counting implementation.
	TailoringStrategy    model.TailoringStrategy     // TailoringStrategy defines the strategy for token tailoring.
	TokenTailoringConfig *model.TokenTailoringConfig // TokenTailoringConfig customizes token tailoring budget parameters for all providers.
	OpenAIOption         []openai.Option             // OpenAIOption stores additional OpenAI options.
	AnthropicOption      []anthropic.Option          // AnthropicOption stores additional Anthropic options.
	GeminiOption         []gemini.Option             // GeminiOption stores additional Gemini options.
	OllamaOption         []ollama.Option             // OllamaOption stores additional Ollama options.
	HunyuanOption        []hunyuan.Option            // HunyuanOption stores additional Hunyuan options.
}

// Callbacks collects provider specific callback hooks.
type Callbacks struct {
	// OpenAIChatRequest runs before dispatching a chat request to OpenAI providers.
	OpenAIChatRequest openai.ChatRequestCallbackFunc
	// OpenAIChatResponse runs after receiving a full chat response from OpenAI providers.
	OpenAIChatResponse openai.ChatResponseCallbackFunc
	// OpenAIChatChunk runs for each streaming chunk from OpenAI providers.
	OpenAIChatChunk openai.ChatChunkCallbackFunc
	// OpenAIStreamComplete runs after an OpenAI streaming session completes.
	OpenAIStreamComplete openai.ChatStreamCompleteCallbackFunc
	// AnthropicChatRequest runs before dispatching a chat request to Anthropic providers.
	AnthropicChatRequest anthropic.ChatRequestCallbackFunc
	// AnthropicChatResponse runs after receiving a full chat response from Anthropic providers.
	AnthropicChatResponse anthropic.ChatResponseCallbackFunc
	// AnthropicChatChunk runs for each streaming chunk from Anthropic providers.
	AnthropicChatChunk anthropic.ChatChunkCallbackFunc
	// AnthropicStreamComplete runs after an Anthropic streaming session completes.
	AnthropicStreamComplete anthropic.ChatStreamCompleteCallbackFunc
	// GeminiChatRequest runs before dispatching a chat request to Gemini providers.
	GeminiChatRequest gemini.ChatRequestCallbackFunc
	// GeminiChatResponse runs after receiving a full chat response from Gemini providers.
	GeminiChatResponse gemini.ChatResponseCallbackFunc
	// GeminiChatChunk runs for each streaming chunk from Gemini providers.
	GeminiChatChunk gemini.ChatChunkCallbackFunc
	// GeminiStreamComplete runs after an Gemini streaming session completes.
	GeminiStreamComplete gemini.ChatStreamCompleteCallbackFunc
	// OllamaChatRequest runs before dispatching a chat request to Ollama providers.
	OllamaChatRequest ollama.ChatRequestCallbackFunc
	// OllamaChatResponse runs after receiving a full chat response from Ollama providers.
	OllamaChatResponse ollama.ChatResponseCallbackFunc
	// OllamaChatChunk runs for each streaming chunk from Ollama providers.
	OllamaChatChunk ollama.ChatChunkCallbackFunc
	// OllamaStreamComplete runs after an Ollama streaming session completes.
	OllamaStreamComplete ollama.ChatStreamCompleteCallbackFunc
	// HunyuanChatRequest runs before dispatching a chat request to Hunyuan providers.
	HunyuanChatRequest hunyuan.ChatRequestCallbackFunc
	// HunyuanChatResponse runs after receiving a full chat response from Hunyuan providers.
	HunyuanChatResponse hunyuan.ChatResponseCallbackFunc
	// HunyuanChatChunk runs for each streaming chunk from Hunyuan providers.
	HunyuanChatChunk hunyuan.ChatChunkCallbackFunc
	// HunyuanStreamComplete runs after a Hunyuan streaming session completes.
	HunyuanStreamComplete hunyuan.ChatStreamCompleteCallbackFunc
}

// WithAPIKey records the API key for the provider.
func WithAPIKey(key string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBaseURL records the base URL for the provider.
func WithBaseURL(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithVariant records the OpenAI-compatible variant.
func WithVariant(variant string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHTTPClientName records the logical HTTP client name.
func WithHTTPClientName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHTTPClientTransport configures the HTTP transport for the provider.
func WithHTTPClientTransport(transport http.RoundTripper) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHeaders appends static HTTP headers for supported providers.
func WithHeaders(headers map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCallbacks registers provider specific callbacks.
func WithCallbacks(cb Callbacks) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithChannelBufferSize overrides the response channel buffer size for supported providers.
func WithChannelBufferSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExtraFields stores provider-specific extra fields for request payload customization.
func WithExtraFields(fields map[string]any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableTokenTailoring toggles automatic token tailoring for supported providers.
func WithEnableTokenTailoring(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxInputTokens sets the maximum input tokens when token tailoring is enabled.
func WithMaxInputTokens(limit int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContextWindow sets the model context window size in tokens for the
// constructed model instance.
func WithContextWindow(tokens int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTokenCounter supplies a custom token counter for token tailoring.
func WithTokenCounter(counter model.TokenCounter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTailoringStrategy supplies a custom token tailoring strategy.
func WithTailoringStrategy(strategy model.TailoringStrategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTokenTailoringConfig sets custom token tailoring budget parameters for all providers.
// This allows advanced users to fine-tune the token allocation strategy.
//
// Example:
//
//	provider.WithTokenTailoringConfig(&model.TokenTailoringConfig{
//	    ProtocolOverheadTokens: 1024,
//	    ReserveOutputTokens:    4096,
//	    SafetyMarginRatio:      0.15,
//	})
//
// Note: It is recommended to use the default values unless you have specific
// requirements.
func WithTokenTailoringConfig(config *model.TokenTailoringConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithOpenAIOption appends raw OpenAI options.
func WithOpenAIOption(opt ...openai.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAnthropicOption appends raw Anthropic options.
func WithAnthropicOption(opt ...anthropic.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithGeminiOption appends raw Gemini options.
func WithGeminiOption(opt ...gemini.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOllamaOption appends raw Ollama options.
func WithOllamaOption(opt ...ollama.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHunyuanOption appends raw Hunyuan options.
func WithHunyuanOption(opt ...hunyuan.Option) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
