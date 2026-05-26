//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package httpclient provides a common HTTP client for Reranker implementations.
package httpclient

import (
	"context"
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/reranker"
)

// Client is a shared HTTP client for Cross-Encoder based rerankers.
// It handles the common logic of sending requests to APIs compatible with Cohere/Infinity.
type Client struct {
	client *http.Client
}

// NewClient creates a new Client.
func NewClient(client *http.Client) *Client { _ = "STUB: not implemented"; return nil }

// RerankRequest represents the request payload for reranking.
type RerankRequest struct {
	Model     string   `json:"model,omitempty"`
	Query     string   `json:"query"`
	Documents []string `json:"documents"`
	TopN      int      `json:"top_n,omitempty"`
}

type rerankResponse struct {
	Results []rerankResult `json:"results"`
}

type rerankResult struct {
	Index          int     `json:"index"`
	RelevanceScore float64 `json:"relevance_score"`
}

// Rerank performs the reranking request.
func (c *Client) Rerank(
	ctx context.Context,
	endpoint string,
	apiKey string,
	reqPayload RerankRequest,
	originalResults []*reranker.Result,
) ([]*reranker.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Map scores back to results

// Sort by score descending
