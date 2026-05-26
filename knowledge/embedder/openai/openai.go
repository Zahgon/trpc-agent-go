//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package openai provides OpenAI embedder implementation.
package openai

import (
	"context"
	"time"

	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/embedder"
)

// Verify that Embedder implements the embedder.Embedder interface.
var _ embedder.Embedder = (*Embedder)(nil)

const (
	// DefaultModel is the default OpenAI embedding model.
	DefaultModel = "text-embedding-3-small"
	// DefaultDimensions is the default embedding dimension for text-embedding-3-small.
	DefaultDimensions = 1536
	// DefaultEncodingFormat is the default encoding format for embeddings.
	DefaultEncodingFormat = "float"
	// DefaultMaxRetries is the default maximum number of retries (same as OpenAI SDK).
	DefaultMaxRetries = 2

	// ModelTextEmbedding3Small represents the text-embedding-3-small model.
	ModelTextEmbedding3Small = "text-embedding-3-small"
	// ModelTextEmbedding3Large represents the text-embedding-3-large model.
	ModelTextEmbedding3Large = "text-embedding-3-large"
	// ModelTextEmbeddingAda002 represents the text-embedding-ada-002 model.
	ModelTextEmbeddingAda002 = "text-embedding-ada-002"

	// EncodingFormatFloat represents the float encoding format.
	EncodingFormatFloat = "float"
	// EncodingFormatBase64 represents the base64 encoding format.
	EncodingFormatBase64 = "base64"

	// textEmbedding3Prefix marks the text-embedding-3 model family
	// (text-embedding-3-small, text-embedding-3-large, ...). The trailing
	// hyphen is required so unrelated ids like text-embedding-30 or
	// text-embedding-3rd-party do not accidentally inherit the legacy
	// default-dimensions forwarding behavior. Members of this family have
	// always received the configured dimensions (with a 1536 default) and
	// we keep that default forwarding to preserve the existing wire
	// behavior for callers that never set WithDimensions.
	textEmbedding3Prefix = "text-embedding-3-"
)

// defaultRetryBackoff is the default backoff durations for retry attempts.
var defaultRetryBackoff = []time.Duration{
	100 * time.Millisecond,
	200 * time.Millisecond,
	400 * time.Millisecond,
	800 * time.Millisecond,
}

// Embedder implements the embedder.Embedder interface for OpenAI API.
type Embedder struct {
	client     openai.Client
	model      string
	dimensions int
	// dimensionsSet indicates whether dimensions was explicitly configured
	// via WithDimensions. When set, the value is forwarded to the API for
	// any model; when unset, see the WithDimensions godoc for which models
	// still receive the historical default.
	dimensionsSet  bool
	encodingFormat string
	user           string
	apiKey         string
	organization   string
	baseURL        string
	requestOptions []option.RequestOption

	// Retry configuration
	maxRetries   int
	retryBackoff []time.Duration
}

// Option represents a functional option for configuring the Embedder.
type Option func(*Embedder)

// WithModel sets the embedding model to use.
func WithModel(model string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDimensions sets the number of dimensions for the embedding.
//
// When set, the value is forwarded as-is to the embeddings endpoint
// regardless of the model id. The caller is responsible for picking a
// value the configured model supports (e.g. text-embedding-3-*, or
// text-embedding-v3/v4 on DashScope-compatible gateways).
//
// When not set, the request includes dimensions only for the
// text-embedding-3-* family (defaulting to DefaultDimensions=1536, which
// preserves the existing wire behavior); for any other model the
// parameter is omitted so the model's server-side default is used.
func WithDimensions(dimensions int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEncodingFormat sets the format for the embeddings.
// Supported formats: "float", "base64".
func WithEncodingFormat(format string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUser sets an optional unique identifier representing your end-user.
func WithUser(user string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAPIKey sets the OpenAI API key.
// If not provided, will use OPENAI_API_KEY environment variable.
func WithAPIKey(apiKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOrganization sets the OpenAI organization ID.
// If not provided, will use OPENAI_ORG_ID environment variable.
func WithOrganization(organization string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBaseURL sets the base URL for OpenAI API.
// Optional, for OpenAI-compatible APIs.
func WithBaseURL(baseURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRequestOptions sets additional options for the OpenAI client requests.
func WithRequestOptions(opts ...option.RequestOption) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMaxRetries sets the maximum number of retries for errors.
// Default is 2 (same as OpenAI SDK default). Negative values are treated as 0.
func WithMaxRetries(maxRetries int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRetryBackoff sets the backoff durations for each retry attempt.
// If the number of retries exceeds the length of backoff slice,
// the last backoff duration will be used for remaining retries.
// Default is [100ms, 200ms, 400ms, 800ms].
func WithRetryBackoff(backoff []time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// New creates a new OpenAI embedder with the given options.
func New(opts ...Option) *Embedder {
	_ = "STUB: not implemented"
	// Create embedder with defaults.
	return nil
}

// Apply functional options.

// Build client options.

// disable openai sdk embedding retries

// Create OpenAI client.

// GetEmbedding implements the embedder.Embedder interface.
// It generates an embedding vector for the given text.
func (e *Embedder) GetEmbedding(ctx context.Context, text string) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract embedding from response.

// GetEmbeddingWithUsage implements the embedder.Embedder interface.
// It generates an embedding vector for the given text and returns usage information.
func (e *Embedder) GetEmbeddingWithUsage(ctx context.Context, text string) ([]float64, map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Extract embedding from response.

// Extract usage information.

// responseWithRetry wraps response with retry logic for errors.
func (e *Embedder) responseWithRetry(ctx context.Context, text string) (*openai.CreateEmbeddingResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No more retries

// Get backoff duration for this attempt and log retry

// getBackoffDuration returns the backoff duration for the given attempt.
// If attempt index exceeds the backoff slice length, returns the last backoff duration.
func (e *Embedder) getBackoffDuration(attempt int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (e *Embedder) response(ctx context.Context, text string) (rsp *openai.CreateEmbeddingResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create embedding request.

// Set optional parameters.

// Forward dimensions when the caller explicitly configured it (any
// model), or implicitly for the text-embedding-3-* family to keep
// the historical default. For other models we omit the parameter so
// the server-side default applies and gateways/models that reject
// the field keep working.

// Combine request options.

// Call OpenAI embeddings API.

// GetDimensions implements the embedder.Embedder interface.
//
// It returns the configured dimensions value (DefaultDimensions when the
// caller never invoked WithDimensions). For non text-embedding-3-* models
// where dimensions was not explicitly configured, the API may return a
// different vector size; in that case prefer calling WithDimensions to
// keep this method consistent with the wire response.
func (e *Embedder) GetDimensions() int { _ = "STUB: not implemented"; return 0 }

// isTextEmbedding3Model reports whether the model belongs to the
// text-embedding-3 family that historically received the configured
// dimensions value (defaulting to 1536) on every request.
func isTextEmbedding3Model(model string) bool { _ = "STUB: not implemented"; return false }
