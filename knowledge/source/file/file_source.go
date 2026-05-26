//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package file provides file-based knowledge source implementation.
package file

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/chunking"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/extractor"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/ocr"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/source"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/transform"
)

const (
	defaultFileSourceName = "File Source"
)

// Source represents a knowledge source for file-based content.
type Source struct {
	filePaths              []string
	name                   string
	metadata               map[string]any
	readers                map[string]reader.Reader
	chunkSize              int
	chunkOverlap           int
	customChunkingStrategy chunking.Strategy
	ocrExtractor           ocr.Extractor
	transformers           []transform.Transformer
	fileReaderType         source.FileReaderType
	contentExtractor       extractor.Extractor
}

// New creates a new file knowledge source.
func New(filePaths []string, opts ...Option) *Source { _ = "STUB: not implemented"; return nil }

// Apply options first to capture configuration.

// Build reader options - pass all configurations to internal source layer

// Initialize readers with configuration

// ReadDocuments reads all files and returns documents using appropriate readers.
func (s *Source) ReadDocuments(ctx context.Context) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip if no file paths provided.

// Name returns the name of this source.
func (s *Source) Name() string {
	_ = "STUB: not implemented"

	// Type returns the type of this source.
	return ""
}

func (s *Source) Type() string { _ = "STUB: not implemented"; return "" }

// processFile processes a single file and returns its documents.
func (s *Source) processFile(ctx context.Context, filePath string) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If a content extractor is configured and supports this extension, use it.

// Create metadata for this file.

// Get absolute path for URI
// Not include ip address and port

// Add metadata to all documents.

// extractAndRead uses the content extractor to convert the file, then pipes
// the result through the appropriate reader for chunking and processing.
func (s *Source) extractAndRead(ctx context.Context, filePath, fileName string) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Find the reader matching the extraction output format.

// readWithReader uses the registered reader to process the file directly.
func (s *Source) readWithReader(filePath string) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetReader sets a custom reader for a specific file type.
func (s *Source) SetReader(fileType string, reader reader.Reader) {
	_ = "STUB: not implemented"
	return
}

// SetMetadata sets metadata for this source.
func (s *Source) SetMetadata(key string, value any) { _ = "STUB: not implemented"; return }

// GetMetadata returns the metadata associated with this source.
func (s *Source) GetMetadata() map[string]any { _ = "STUB: not implemented"; return nil }
