//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package huggingface provides huggingface-compatible model implementations.
// Text-Embeddings-Inference API: https://github.com/huggingface/text-embeddings-inference
//
// Hugging Face Text-Embeddings-Inference (TEI) is a high-performance, production-ready
// inference server for text embedding models. It supports a wide range of transformer-based
// models and provides optimized inference for embedding generation.
//
// The API provides two main endpoints:
// - /embed: Default embedding endpoint with pooling
// - /embed_all: Returns all embeddings without pooling
//
// Usage Example:
//
//	embedder := New(WithBaseURL("http://localhost:8080"))
//	embedding, err := embedder.GetEmbedding(ctx, "Hello world")
//
// Configuration Options:
// - Base URL: API server address
// - Dimensions: Output embedding size
// - Normalize: Whether to normalize embeddings
// - Truncate: Text truncation behavior
// - Embed Route: Choose between /embed and /embed_all endpoints
package huggingface

import (
	"context"
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/embedder"
)

// Verify that Embedder implements the embedder.Embedder interface.
var _ embedder.Embedder = (*Embedder)(nil)

const (
	// DefaultDimensions is the default embedding dimension.
	DefaultDimensions = 1024

	// DefaultBaseURL is the default base URL for the Text-Embeddings-Inference API.
	DefaultBaseURL = "http://localhost:8080"
)

// Embedder implements the embedder.Embedder interface for Hugging Face Text-Embeddings-Inference API.
type Embedder struct {
	baseURL             string
	dimensions          int
	normalize           bool
	promptName          string
	truncate            bool
	truncationDirection TruncateDirection
	embedRoute          EmbedRoute
	client              *http.Client

	serverAddress *string // for telemetry
	serverPort    *int    // for telemetry
}

// TruncateDirection represents the truncation direction for the embedding.
type TruncateDirection string

var (
	TruncateLeft  TruncateDirection = "Left"
	TruncateRight TruncateDirection = "Right"
)

// EmbedRoute represents the route for the embedding request.
type EmbedRoute string

const (
	EmbedDefault EmbedRoute = "/embed"     // default
	EmbedAll     EmbedRoute = "/embed_all" // get all embeddings without pooling
)

// Option represents a functional option for configuring the Embedder.
type Option func(*Embedder)

// WithBaseURL sets the base URL for the Text-Embeddings-Inference API.
// Such as "http://localhost:8080".
func WithBaseURL(baseURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDimensions sets the number of dimensions for the embedding.
func WithDimensions(dimensions int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNormalize sets whether to normalize the embeddings.
func WithNormalize(normalize bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPromptName sets the name of the prompt to use.
func WithPromptName(promptName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTruncate sets whether to truncate the text.
func WithTruncate(truncate bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTruncationDirection sets the truncation direction for the embedding.
func WithTruncationDirection(truncationDirection TruncateDirection) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEmbedRoute sets the route for the embedding request.
func WithEmbedRoute(embed EmbedRoute) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithClient sets the HTTP client to use for requests.
func WithClient(client *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

type requestBody struct {
	Dimensions          int               `json:"dimensions,omitempty"`
	Inputs              string            `json:"inputs,omitempty"`
	Normalize           bool              `json:"normalize,omitempty"`
	PromptName          string            `json:"prompt_name,omitempty"`
	Truncate            bool              `json:"truncate,omitempty"`
	TruncationDirection TruncateDirection `json:"truncation_direction,omitempty"`
}

type embedResponse struct {
	Embeddings [][]float64 `json:"embeddings"`
}

// New creates a new hugging face Embedder instance
func New(opts ...Option) *Embedder { _ = "STUB: not implemented"; return nil }

// GetEmbedding generates an embedding vector for the given text.
func (e *Embedder) GetEmbedding(ctx context.Context, text string) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetEmbeddingWithUsage generates an embedding vector for the given text
// and returns usage information if available.
// TEI don't provide usage information
func (e *Embedder) GetEmbeddingWithUsage(ctx context.Context, text string) ([]float64, map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetDimensions returns the dimensionality of the embeddings produced by this embedder.
// Returns 0 if dimensions are not known or configurable.
func (e *Embedder) GetDimensions() int { _ = "STUB: not implemented"; return 0 }

func (e *Embedder) response(ctx context.Context, text string) (rsp *embedResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Embedder) requestBody(text string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Embedder) parseResponse(resp *http.Response) (rsp *embedResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
