//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package url provides URL-based knowledge source implementation.
package url

import (
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/chunking"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/extractor"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/source"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/transform"
)

// Option represents a functional option for configuring Source.
type Option func(*Source)

// WithName sets a custom name for the URL source.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContentFetchingURL sets the real content fetching URL for the source.
// The real content fetching URL is used to fetch the actual content of the document.
func WithContentFetchingURL(url []string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadata sets additional metadata for the source.
func WithMetadata(metadata map[string]any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadataValue adds a single metadata key-value pair.
func WithMetadataValue(key string, value any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHTTPClient sets a custom HTTP client for URL fetching.
func WithHTTPClient(client *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCustomChunkingStrategy sets a custom chunking strategy for document splitting.
// This overrides the reader's default chunking strategy.
func WithCustomChunkingStrategy(strategy chunking.Strategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithChunkSize sets the chunk size for the reader's default chunking strategy.
func WithChunkSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithChunkOverlap sets the chunk overlap for the reader's default chunking strategy.
func WithChunkOverlap(overlap int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTransformers sets transformers for document processing.
// Transformers are applied before and after chunking.
//
// Example:
//
//	source := url.New(urls, url.WithTransformers(
//	    transform.NewCharFilter("\n", "\t"),
//	    transform.NewCharDedup(" "),
//	))
func WithTransformers(transformers ...transform.Transformer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFileReaderType overrides the automatic file reader type detection based on content-type or URL extension.
// This forces the source to use the specified file type reader.
// Use predefined constants from source package for type safety.
//
// Example:
//
//	source := url.New([]string{"https://example.com/api/data"}, url.WithFileReaderType(source.FileReaderTypeJSON))
func WithFileReaderType(fileType source.FileReaderType) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithExtractor sets a content extractor for handling complex or unsupported formats.
// When configured, the extractor is used for URL content whose inferred file
// extension matches the extractor's supported formats. The source always
// fetches content through the configured HTTP client and passes the response
// body to ExtractFromReader.
func WithExtractor(e extractor.Extractor) Option { _ = "STUB: not implemented"; return *new(Option) }
