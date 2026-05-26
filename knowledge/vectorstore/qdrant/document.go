//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// document.go implements document CRUD operations for the VectorStore.
package qdrant

import (
	"context"

	"github.com/qdrant/go-client/qdrant"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/vectorstore"
)

// Add stores a document with its embedding vector.
func (vs *VectorStore) Add(ctx context.Context, doc *document.Document, embedding []float64) error {
	_ = "STUB: not implemented"
	return nil
}

// AddBatch stores multiple documents with their embedding vectors in a single operation.
// This is more efficient than calling Add() multiple times for bulk imports.
func (vs *VectorStore) AddBatch(ctx context.Context, docs []*document.Document, embeddings [][]float64) error {
	_ = "STUB: not implemented"
	return nil
}

// buildPoint creates a Qdrant point from a document and embedding.
func (vs *VectorStore) buildPoint(doc *document.Document, embedding []float64) *qdrant.PointStruct {
	_ = "STUB: not implemented"
	return nil
}

// Get retrieves a document by ID.
func (vs *VectorStore) Get(ctx context.Context, id string) (*document.Document, []float64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// fromRetrievedPoint converts a Qdrant RetrievedPoint to a Document and vector.
func (vs *VectorStore) fromRetrievedPoint(pt *qdrant.RetrievedPoint) (*document.Document, []float64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Update modifies an existing document.
// This performs an upsert: the document is created if it doesn't exist.
func (vs *VectorStore) Update(ctx context.Context, doc *document.Document, embedding []float64) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete removes a document by ID.
func (vs *VectorStore) Delete(ctx context.Context, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteByFilter deletes documents matching criteria.
func (vs *VectorStore) DeleteByFilter(ctx context.Context, opts ...vectorstore.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

// No criteria specified: nothing to delete, succeed silently.

// UpdateByFilter updates documents matching the filter with the specified field values.
// This operation is not yet supported by Qdrant VectorStore.
func (vs *VectorStore) UpdateByFilter(ctx context.Context, opts ...vectorstore.UpdateByFilterOption) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// validateEmbedding checks that the embedding is valid.
func (vs *VectorStore) validateEmbedding(embedding []float64) error {
	_ = "STUB: not implemented"
	return nil
}

// extractDenseVector extracts the dense vector from Qdrant VectorsOutput.
func (vs *VectorStore) extractDenseVector(vectors *qdrant.VectorsOutput) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// Try named vectors first (BM25 mode)

// Fall back to single vector mode
