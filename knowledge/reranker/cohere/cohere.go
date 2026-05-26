//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package cohere provides a Reranker implementation using Cohere's Rerank API.
package cohere

import (
	"context"
	"errors"
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/reranker"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/reranker/internal/httpclient"
)

var (
	// errEndpointEmpty is returned when the endpoint is empty.
	errEndpointEmpty = errors.New("cohere endpoint cannot be empty")
)

const (
	defaultCohereEndpoint = "https://api.cohere.ai/v1/rerank"
	defaultCohereModel    = "rerank-english-v3.0"
)

// Reranker implements Reranker using Cohere's API.
type Reranker struct {
	apiKey     string
	modelName  string
	endpoint   string
	topN       int
	httpClient *httpclient.Client
}

// Option configures Reranker.
type Option func(*Reranker)

// WithModel sets the model name.
func WithModel(model string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTopN sets the TopN.
func WithTopN(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAPIKey sets the API key.
func WithAPIKey(key string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEndpoint sets the endpoint URL.
func WithEndpoint(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

// New creates a new Cohere reranker.
func New(opts ...Option) (*Reranker, error) { _ = "STUB: not implemented"; return nil, nil }

// Rerank implements the Reranker interface.
func (r *Reranker) Rerank(
	ctx context.Context,
	query *reranker.Query,
	results []*reranker.Result,
) ([]*reranker.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply TopN locally as a safeguard, though API handles it too
