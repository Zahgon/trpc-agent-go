//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package chunking

import (
	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
)

// FixedSizeChunking implements a chunking strategy that splits text into fixed-size chunks.
type FixedSizeChunking struct {
	chunkSize int
	overlap   int
}

// Option represents a functional option for configuring FixedSizeChunking.
type Option func(*FixedSizeChunking)

// WithChunkSize sets the maximum size of each chunk in characters.
func WithChunkSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOverlap sets the number of characters to overlap between chunks.
func WithOverlap(overlap int) Option { _ = "STUB: not implemented"; return *new(Option) }

// NewFixedSizeChunking creates a new fixed-size chunking strategy with options.
func NewFixedSizeChunking(opts ...Option) *FixedSizeChunking { _ = "STUB: not implemented"; return nil }

// Apply options.

// Validate parameters.

// Chunk splits the document into fixed-size chunks with optional overlap.
func (f *FixedSizeChunking) Chunk(doc *document.Document) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If content is smaller than chunk size, return as single chunk.

// Use UTF-8 safe splitting to ensure proper character boundaries.

// Apply overlap if specified.

// applyOverlap applies overlap between consecutive chunks while maintaining UTF-8 safety.
func (f *FixedSizeChunking) applyOverlap(chunks []*document.Document) []*document.Document {
	_ = "STUB: not implemented"
	return nil
}

// Get overlap text safely.

// Create new metadata for overlapped chunk.
