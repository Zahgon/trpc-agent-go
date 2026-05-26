//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package infinity provides a Reranker implementation compatible with Infinity.
package infinity

import (
	"context"
	"errors"
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/reranker"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/reranker/internal/httpclient"
)

var (
	// errEndpointEmpty is returned when the endpoint is empty.
	errEndpointEmpty = errors.New("infinity endpoint cannot be empty")
)

// Reranker implements Reranker using a self-hosted Infinity/TEI instance.
type Reranker struct {
	endpoint   string
	apiKey     string
	modelName  string
	topN       int
	httpClient *httpclient.Client
}

// Option configures Reranker.
type Option func(*Reranker)

// WithAPIKey sets the API key (optional for self-hosted).
func WithAPIKey(key string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithModel sets the model name (optional, depends on server config).
func WithModel(model string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTopN sets the TopN.
func WithTopN(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithEndpoint sets the endpoint URL.
func WithEndpoint(endpoint string) Option { _ = "STUB: not implemented"; return *new(Option) }

// New creates a new Infinity reranker.
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
