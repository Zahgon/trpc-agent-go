//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// metadata.go implements metadata operations: Count, GetMetadata, UpdateMetadata.
package qdrant

import (
	"context"

	"github.com/qdrant/go-client/qdrant"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/vectorstore"
)

// Count returns the number of documents.
func (vs *VectorStore) Count(ctx context.Context, opts ...vectorstore.CountOption) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetMetadata retrieves document metadata.
func (vs *VectorStore) GetMetadata(ctx context.Context, opts ...vectorstore.GetMetadataOption) (map[string]vectorstore.DocumentMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Optimize batch size: don't fetch more than needed

// Stop if we've reached the limit or exhausted results

// buildMetadataFilter builds a filter for GetMetadata.
func (vs *VectorStore) buildMetadataFilter(config *vectorstore.GetMetadataConfig) (*qdrant.Filter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateMetadata updates the metadata of an existing document without changing its vector.
// The metadata is merged with existing metadata; to remove a field, set it to nil.
func (vs *VectorStore) UpdateMetadata(ctx context.Context, id string, metadata map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// Nothing to update
