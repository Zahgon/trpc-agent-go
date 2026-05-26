//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package openai provides OpenAI-compatible model implementations.
package openai

import (
	"context"

	openai "github.com/openai/openai-go"
	openaiopt "github.com/openai/openai-go/option"
	"trpc.group/trpc-go/trpc-agent-go/model"
	imodel "trpc.group/trpc-go/trpc-agent-go/model/internal/model"
)

const (
	// defaultChannelBufferSize is the default channel buffer size.
	defaultChannelBufferSize = 256
	// defaultBatchCompletionWindow is the default batch completion window.
	defaultBatchCompletionWindow = "24h"
)

// ChatRequestCallbackFunc is the function type for the chat request callback.
type ChatRequestCallbackFunc func(
	ctx context.Context,
	chatRequest *openai.ChatCompletionNewParams,
)

// ChatRequestJSONCallbackFunc is the function type for the chat request
// JSON callback.
type ChatRequestJSONCallbackFunc func(
	ctx context.Context,
	chatRequestJSON []byte,
	marshalErr error,
)

// ChatResponseCallbackFunc is the function type for the chat response callback.
type ChatResponseCallbackFunc func(
	ctx context.Context,
	chatRequest *openai.ChatCompletionNewParams,
	chatResponse *openai.ChatCompletion,
)

// ChatChunkCallbackFunc is the function type for the chat chunk callback.
type ChatChunkCallbackFunc func(
	ctx context.Context,
	chatRequest *openai.ChatCompletionNewParams,
	chatChunk *openai.ChatCompletionChunk,
)

// ChatStreamCompleteCallbackFunc is the function type for the chat stream completion callback.
// This callback is invoked when streaming is completely finished (success or error).
type ChatStreamCompleteCallbackFunc func(
	ctx context.Context,
	chatRequest *openai.ChatCompletionNewParams,
	accumulator *openai.ChatCompletionAccumulator, // nil if streamErr is not nil
	streamErr error, // nil if streaming completed successfully
)

// options contains configuration options for creating a Model.
type options struct {
	// API key for the OpenAI client.
	APIKey string
	// Base URL for the OpenAI client. It is optional for OpenAI-compatible APIs.
	BaseURL string
	// Buffer size for response channels (default: 256)
	ChannelBufferSize int
	// Options for the HTTP client.
	HTTPClientOptions []HTTPClientOption
	// Callback for the chat request.
	ChatRequestCallback ChatRequestCallbackFunc
	// Callback for the marshaled chat request JSON.
	ChatRequestJSONCallback ChatRequestJSONCallbackFunc
	// Callback for the chat response.
	ChatResponseCallback ChatResponseCallbackFunc
	// Callback for the chat chunk.
	ChatChunkCallback ChatChunkCallbackFunc
	// Callback for the chat stream completion.
	ChatStreamCompleteCallback ChatStreamCompleteCallbackFunc
	// ChatTelemetry enables opt-in telemetry for direct model usage.
	ChatTelemetry bool
	// Options for the OpenAI client.
	OpenAIOptions []openaiopt.RequestOption
	// Extra fields to be added to the HTTP request body.
	ExtraFields map[string]any
	// Variant for model-specific behavior.
	Variant Variant
	// variantSet tracks whether WithVariant was explicitly provided.
	variantSet bool
	// Batch completion window for batch processing.
	BatchCompletionWindow openai.BatchNewParamsCompletionWindow
	// Batch metadata for batch processing.
	BatchMetadata map[string]string
	// BatchBaseURL overrides the base URL for batch requests (batches/files).
	BatchBaseURL string
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
	// ShowToolCallDelta controls whether to expose tool call
	// deltas in streaming responses. When true, raw tool_call
	// chunks from the provider will be forwarded via
	// Response.Choices[].Delta.ToolCalls instead of being
	// suppressed until the final aggregated response.
	ShowToolCallDelta bool
	// ReasoningContentBackfill controls whether assistant messages should
	// replay an empty reasoning_content field when the message has no
	// reasoning text.
	ReasoningContentBackfill    bool
	reasoningContentBackfillSet bool
	accumulateChunkUsage        AccumulateChunkUsage
	// OptimizeForCache controls whether to optimize message structure for prompt caching.
	// When enabled, system messages will be moved to the front to improve cache hit rates.
	// OpenAI's prompt caching is automatic and doesn't require explicit cache control,
	// but message ordering affects cache effectiveness.
	OptimizeForCache    bool
	optimizeForCacheSet bool

	// OmitFileContentParts controls whether file content parts are removed
	// from requests sent to the model provider.
	//
	// This can be useful when a provider rejects file inputs in chat messages,
	// while still keeping the file parts in-memory for downstream tools.
	OmitFileContentParts bool
}

var (
	defaultOptions = options{
		Variant:               VariantOpenAI, // The default variant is VariantOpenAI.
		ChannelBufferSize:     defaultChannelBufferSize,
		BatchCompletionWindow: defaultBatchCompletionWindow,
		TokenCounter:          model.NewSimpleTokenCounter(),
		TokenTailoringConfig: &model.TokenTailoringConfig{
			ProtocolOverheadTokens: imodel.DefaultProtocolOverheadTokens,
			ReserveOutputTokens:    imodel.DefaultReserveOutputTokens,
			SafetyMarginRatio:      imodel.DefaultSafetyMarginRatio,
			InputTokensFloor:       imodel.DefaultInputTokensFloor,
			MaxInputTokensRatio:    imodel.DefaultMaxInputTokensRatio,
		},
		OptimizeForCache: false,
	}
)

// Option is a function that configures an OpenAI model.
type Option func(*options)

// WithAPIKey sets the API key for the OpenAI client.
func WithAPIKey(key string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBaseURL sets the base URL for the OpenAI client.
func WithBaseURL(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithChannelBufferSize sets the channel buffer size for the OpenAI client.
func WithChannelBufferSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithChatRequestCallback sets the function to be called before sending a
// chat request. The callback runs synchronously in GenerateContent before
// the response goroutine starts. Start your own goroutine in the callback
// if asynchronous behavior is needed.
func WithChatRequestCallback(fn ChatRequestCallbackFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithChatRequestJSONCallback sets the function to be called with the
// marshaled chat request JSON before sending the request.
func WithChatRequestJSONCallback(fn ChatRequestJSONCallbackFunc) Option {
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
// streaming is completed. The callback runs synchronously before the
// terminal streaming result is surfaced to the caller.
// Called for both successful and failed streaming completions.
func WithChatStreamCompleteCallback(fn ChatStreamCompleteCallbackFunc) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithChatTelemetry enables chat trace and metric reporting for direct
// model/openai usage.
//
// This option is intended for users who call Model.GenerateContent or
// Model.GenerateContentIter directly. The recommended runner + agent path
// already reports chat telemetry from llmflow. Reusing a model with this
// option enabled inside runner + agent may report duplicate chat telemetry.
func WithChatTelemetry(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithReasoningContentBackfill enables replay-time reasoning_content backfill
// for assistant messages that have no reasoning text.
func WithReasoningContentBackfill(enabled bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHTTPClientOptions sets the HTTP client options for the OpenAI client.
func WithHTTPClientOptions(httpOpts ...HTTPClientOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithOpenAIOptions sets the OpenAI options for the OpenAI client.
// E.g. use its middleware option:
//
//	import (
//		openai "github.com/openai/openai-go"
//		openaiopt "github.com/openai/openai-go/option"
//	)
//
//	WithOpenAIOptions(openaiopt.WithMiddleware(
//		func(req *http.Request, next openaiopt.MiddlewareNext) (*http.Response, error) {
//			// do something
//			return next(req)
//		}
//	)))
func WithOpenAIOptions(openaiOpts ...openaiopt.RequestOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHeaders appends static HTTP headers to all OpenAI requests.
func WithHeaders(headers map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExtraFields sets extra fields to be added to the HTTP request body.
// These fields will be included in every chat completion request.
// E.g.:
//
//	WithExtraFields(map[string]any{
//		"custom_metadata": map[string]string{
//			"session_id": "abc",
//		},
//	})
//
// and "session_id" : "abc" will be added to the HTTP request json body.
func WithExtraFields(extraFields map[string]any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithVariant sets the model variant for specific behavior.
// The default variant is VariantOpenAI.
// Optional variants are:
// - VariantHunyuan: Hunyuan variant with specific file handling.
func WithVariant(variant Variant) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOmitFileContentParts controls whether file content parts are removed
// from requests sent to the model provider.
func WithOmitFileContentParts(omit bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBatchCompletionWindow sets the batch completion window.
func WithBatchCompletionWindow(window openai.BatchNewParamsCompletionWindow) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithBatchMetadata sets the batch metadata.
func WithBatchMetadata(metadata map[string]string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithBatchBaseURL sets a base URL override for batch requests (batches/files).
// When set, batch operations will use this base URL via per-request override.
func WithBatchBaseURL(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEnableTokenTailoring enables automatic token tailoring based on model context window.
// When enabled, the system will automatically calculate max input tokens using the model's
// context window minus reserved tokens and protocol overhead.
func WithEnableTokenTailoring(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMaxInputTokens sets only the input token limit for token tailoring.
// The counter/strategy will be lazily initialized if not provided.
// Defaults to SimpleTokenCounter and MiddleOutStrategy.
func WithMaxInputTokens(limit int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContextWindow sets the model context window size in tokens for this
// model instance.
func WithContextWindow(tokens int) Option { _ = "STUB: not implemented"; return *new(Option) }

// AccumulateChunkUsage is the function type for accumulating chunk usage.
type AccumulateChunkUsage func(u model.Usage, delta model.Usage) model.Usage

// WithAccumulateChunkTokenUsage sets the function to be called to accumulate chunk token usage.
func WithAccumulateChunkTokenUsage(a AccumulateChunkUsage) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// inverseOpenAISDKAddChunkUsage calculates the inverse of OPENAISKDAddChunkUsage, related to the current openai sdk version
func inverseOpenAISDKAddChunkUsage(u model.Usage, delta model.Usage) model.Usage {
	_ = "STUB: not implemented"
	return *new(model.Usage)
}

// completionUsageToModelUsage converts openai.CompletionUsage to model.Usage.
func completionUsageToModelUsage(usage openai.CompletionUsage) model.Usage {
	_ = "STUB: not implemented"
	return *new(model.Usage)
}

// modelUsageToCompletionUsage converts model.Usage to openai.CompletionUsage.
func modelUsageToCompletionUsage(usage model.Usage) openai.CompletionUsage {
	_ = "STUB: not implemented"
	return *new(openai.CompletionUsage)
}

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
//
// Example:
//
//	openai.WithTokenTailoringConfig(&model.TokenTailoringConfig{
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

// WithShowToolCallDelta controls whether to expose tool call
// deltas in streaming responses. When enabled, the model will
// forward provider tool_call chunks via Delta.ToolCalls so
// callers can reconstruct arguments incrementally.
func WithShowToolCallDelta(show bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOptimizeForCache controls whether to optimize message structure for prompt caching.
// When enabled, system messages will be moved to the front of the message list
// to improve cache hit rates with OpenAI's automatic prompt caching.
//
// OpenAI's prompt caching works automatically and doesn't require explicit cache control,
// but consistent message ordering significantly improves cache effectiveness.
//
// Enable this for:
// - Multi-turn conversations with consistent system prompts
// - Scenarios where system context is reused across requests
//
// Disable this for:
// - Cases where message order must be strictly preserved
// - Testing scenarios requiring deterministic behavior
// - When you want explicit control over message ordering
//
// This option is enabled by default for VariantOpenAI.
//
// Example:
//
//	model := openai.New("gpt-4o",
//	    openai.WithOptimizeForCache(false),  // Disable cache optimization
//	)
func WithOptimizeForCache(optimize bool) Option { _ = "STUB: not implemented"; return *new(Option) }
