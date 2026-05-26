//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package elasticsearch provides Elasticsearch-based vector storage implementation.
package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/vectorstore"
)

const (
	// scriptParamQueryVector is the name of the script parameter for the query vector.
	scriptParamQueryVector = "query_vector"
)

// buildVectorSearchQuery builds a vector similarity search query.
func (vs *VectorStore) buildVectorSearchQuery(query *vectorstore.SearchQuery) (*types.SearchRequestBody, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Marshal query vector to a valid JSON array for script params.

// Build script source dynamically to support custom embedding field.

// Create script for cosine similarity using esdsl.

// Create match_all query using esdsl.

// Create script_score query using esdsl.

// Build the complete search request using official SearchRequestBody.

// Add filters if specified.

func (vs *VectorStore) buildFilterSearchQuery(query *vectorstore.SearchQuery) (*types.SearchRequestBody, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Build the complete search request using official SearchRequestBody.

// buildKeywordSearchQuery builds a keyword-based search query.
func (vs *VectorStore) buildKeywordSearchQuery(query *vectorstore.SearchQuery) (*types.SearchRequestBody, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create multi_match query using esdsl.

// Build the complete search request using official SearchRequestBody.

// Add filters if specified.

// buildHybridSearchQuery builds a hybrid search query combining vector and keyword search.
func (vs *VectorStore) buildHybridSearchQuery(query *vectorstore.SearchQuery) (*types.SearchRequestBody, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Marshal query vector to a valid JSON array for script params.

// Build script with custom embedding field.

// Create match_all query for script_score.

// Create script_score query.

// Combine queries using bool query.

// Build the complete search request using official SearchRequestBody.

// Add filters if specified.

// buildFilterQuery builds a filter query for search results.
func (vs *VectorStore) buildFilterQuery(filter *vectorstore.SearchFilter) (types.QueryVariant, error) {
	_ = "STUB: not implemented"
	return *new(types.QueryVariant), nil
}

// Filter by document IDs.

// Filter by metadata.

func (vs *VectorStore) getMaxResult(maxResults int) int { _ = "STUB: not implemented"; return 0 }
