//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package ollama provides Ollama embedder implementation.
package ollama

import (
	"context"
	"net/http"
	"time"

	"github.com/ollama/ollama/api"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/embedder"
)

// Verify that Embedder implements the embedder.Embedder interface.
var _ embedder.Embedder = (*Embedder)(nil)

const (
	// DefaultModel is the default Ollama embedding model.
	DefaultModel = "llama3.2:latest"

	// DefaultDimensions is the default embedding dimension.
	DefaultDimensions = 1536

	// OllamaHost is the environment variable for the Ollama host.
	OllamaHost = "OLLAMA_HOST"
)

// Embedder implements the embedder.Embedder interface for Ollama API.
type Embedder struct {
	useEmbeddings bool
	model         string
	host          string
	httpClient    *http.Client
	options       map[string]any
	keepAlive     time.Duration
	client        *api.Client
	truncate      *bool
	dimensions    int

	serverAddress *string // for telemetry
	serverPort    *int    // for telemetry
}

// Option represents a functional option for configuring the Embedder.
type Option func(*Embedder)

// WithModel sets the embedding model to use.
func WithModel(model string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHost sets the Ollama host.
func WithHost(host string) Option { _ = "STUB: not implemented"; return *new(Option) }

// withHttpClient sets the HTTP client to use.
// The site is temporarily not open to the public, as we may implement injection of an internal http client.
func withHttpClient(client *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTruncate sets the truncate flag.
func WithTruncate(truncate bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUseEmbeddings enables the use of embeddings endpoint(/api/embeddings)
// default is false. means not using the /api/embeddings, but using /api/embed
// ref: https://github.com/ollama/ollama/blob/main/docs/api.md#generate-embedding
func WithUseEmbeddings() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOptions sets the options to use.
// options: https://github.com/ollama/ollama/blob/main/docs/modelfile.mdx#valid-parameters-and-values
func WithOptions(options map[string]any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithKeepAlive sets the keep-alive duration.
func WithKeepAlive(keepAlive time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDimensions sets the number of dimensions for the embedding.
func WithDimensions(dimensions int) Option { _ = "STUB: not implemented"; return *new(Option) }

// embedResponse is a wrapper around api.EmbedResponse that adds an additional field for the embeddings.
// source embeddings is [][]float32, we override it to [][]float64
type embedResponse struct {
	api.EmbedResponse
	Embeddings [][]float64 `json:"embeddings"`
}

// New creates a new Ollama embedder with the given options.
func New(opts ...Option) *Embedder { _ = "STUB: not implemented"; return nil }

// GetEmbedding implements the embedder.Embedder interface.
// It generates an embedding vector for the given text
func (e *Embedder) GetEmbedding(ctx context.Context, text string) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetEmbeddingWithUsage implements the embedder.Embedder interface.
// It generates an embedding vector for the given text and returns usage information.
func (e *Embedder) GetEmbeddingWithUsage(ctx context.Context, text string) ([]float64, map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetDimensions implements the embedder.Embedder interface.
// It returns the number of dimensions in the embedding vectors
func (e *Embedder) GetDimensions() int { _ = "STUB: not implemented"; return 0 }

// response makes a request to the Ollama API and returns the response.
// we prefer to use the /api/embed endpoint, /api/embeddings has been superseded by `/api/embed`
func (e *Embedder) response(ctx context.Context, text string) (rsp *embedResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
