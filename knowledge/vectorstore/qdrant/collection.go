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
)

// ensureCollection checks if the collection exists and creates it if not.
// If the collection exists, it validates the configuration matches our settings.
func (vs *VectorStore) ensureCollection(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Try to create the collection

// createCollection creates a new Qdrant collection with the configured settings.
func (vs *VectorStore) createCollection(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Named vectors configuration: dense + sparse for BM25

// Sparse vector config with IDF modifier for BM25

// Simple single vector configuration (no BM25)

// validateCollectionConfig checks that an existing collection matches our expected configuration.
func (vs *VectorStore) validateCollectionConfig(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate vector configuration

// Validate sparse vector configuration if BM25 is enabled

// validateVectorConfig validates that the collection's vector configuration matches our settings.
func (vs *VectorStore) validateVectorConfig(config *qdrant.VectorsConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Single vector mode (non-BM25)

// Named vectors mode (BM25 enabled)

// Check dimension

// Check distance metric

// validateSparseVectorConfig validates that the collection has sparse vectors configured for BM25.
func (vs *VectorStore) validateSparseVectorConfig(config *qdrant.SparseVectorConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteCollection deletes the entire collection and all its data.
// Use with caution as this operation is irreversible.
func (vs *VectorStore) DeleteCollection(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
