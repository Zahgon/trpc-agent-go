//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package qdrant

import (
	"context"

	"github.com/qdrant/go-client/qdrant"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/vectorstore"
)

// Search performs similarity search.
func (vs *VectorStore) Search(ctx context.Context, query *vectorstore.SearchQuery) (*vectorstore.SearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// searchByVector performs dense vector similarity search.
func (vs *VectorStore) searchByVector(ctx context.Context, query *vectorstore.SearchQuery) (*vectorstore.SearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Build query based on vector configuration

// Named vector query

// Single vector query

// searchByFilter performs filter-only search without vector similarity.
func (vs *VectorStore) searchByFilter(ctx context.Context, query *vectorstore.SearchQuery) (*vectorstore.SearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// searchByKeyword performs BM25 sparse vector search.
func (vs *VectorStore) searchByKeyword(ctx context.Context, query *vectorstore.SearchQuery) (*vectorstore.SearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use BM25 sparse vector search with Document inference

// searchByHybrid performs hybrid search combining dense vectors and BM25.
// If BM25 is not enabled, it falls back to vector-only search.
func (vs *VectorStore) searchByHybrid(ctx context.Context, query *vectorstore.SearchQuery) (*vectorstore.SearchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prefetch more results for better fusion (capped to avoid excessive memory usage)

// Hybrid search with Prefetch + RRF fusion

// Dense vector search

// BM25 sparse vector search

// buildSearchFilter converts a SearchFilter to a Qdrant Filter.
func (vs *VectorStore) buildSearchFilter(filter *vectorstore.SearchFilter) (*qdrant.Filter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add ID filter if present (uses special HasId condition)

// Build metadata and filter conditions using the converter

// Convert universal filters to Qdrant filter

// Extract conditions from converted filter and add to our conditions

// MustNot needs special handling - wrap in a nested filter
