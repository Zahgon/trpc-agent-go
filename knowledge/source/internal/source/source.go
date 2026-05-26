//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package source provides internal source utils.
package source

import (
	"trpc.group/trpc-go/trpc-agent-go/knowledge/chunking"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/ocr"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/transform"

	_ "trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader/csv"
	_ "trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader/golang"
	_ "trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader/json"
	_ "trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader/markdown"
	_ "trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader/proto"
	_ "trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader/text"
)

// ReaderConfig holds configuration for creating readers.
type ReaderConfig struct {
	chunkSize              int
	chunkOverlap           int
	customChunkingStrategy chunking.Strategy
	ocrExtractor           ocr.Extractor
	transformers           []transform.Transformer
}

// ReaderOption is a functional option for configuring readers.
type ReaderOption func(*ReaderConfig)

// WithChunkSize sets the chunk size for readers.
func WithChunkSize(size int) ReaderOption { _ = "STUB: not implemented"; return *new(ReaderOption) }

// WithChunkOverlap sets the chunk overlap for readers.
func WithChunkOverlap(overlap int) ReaderOption {
	_ = "STUB: not implemented"
	return *new(ReaderOption)
}

// WithCustomChunkingStrategy sets a custom chunking strategy for readers.
func WithCustomChunkingStrategy(strategy chunking.Strategy) ReaderOption {
	_ = "STUB: not implemented"
	return *new(ReaderOption)
}

// WithOCRExtractor sets the OCR extractor for PDF reader.
func WithOCRExtractor(extractor ocr.Extractor) ReaderOption {
	_ = "STUB: not implemented"
	return *new(ReaderOption)
}

// WithTransformers sets the transformers for document processing.
func WithTransformers(transformers ...transform.Transformer) ReaderOption {
	_ = "STUB: not implemented"
	return *new(ReaderOption)
}

// GetReaders returns all available readers configured with the given options.
func GetReaders(opts ...ReaderOption) map[string]reader.Reader {
	_ = "STUB: not implemented"
	return nil
}

// Build reader options

// Get readers with options

// buildReaderOptions constructs reader options from config.
func buildReaderOptions(config *ReaderConfig) []reader.Option {
	_ = "STUB: not implemented"
	return nil

	// Pass chunking configurations to readers
}

// GetFileType determines the file type based on the file extension.
func GetFileType(filePath string) string { _ = "STUB: not implemented"; return "" }

// GetFileTypeFromContentType determines the file type based on content type or file extension.
func GetFileTypeFromContentType(contentType, fileName string) string {
	_ = "STUB: not implemented"
	// First try content type.
	return ""
}

// Fall back to file extension.

// Unknown extension, fallback to text reader

// GetReadersWithChunkConfig is deprecated. Use GetReaders with functional options instead.
// Deprecated: Use GetReaders(WithChunkSize(size), WithChunkOverlap(overlap)) instead.
func GetReadersWithChunkConfig(chunkSize, overlap int) map[string]reader.Reader {
	_ = "STUB: not implemented"
	return nil
}

// ResolveFileType returns the file type to use, considering any override.
// If overrideType is non-empty, it returns the override; otherwise returns detectedType.
// This provides a unified way to handle fileReaderType override logic across all sources.
func ResolveFileType(overrideType, detectedType string) string {
	_ = "STUB: not implemented"
	return ""
}
