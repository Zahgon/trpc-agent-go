//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package dir provides directory-based knowledge source implementation.
package dir

import (
	"trpc.group/trpc-go/trpc-agent-go/knowledge/chunking"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/extractor"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/ocr"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/source"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/transform"
)

// Option represents a functional option for configuring directory sources.
type Option func(*Source)

// WithName sets the name of the directory source.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadata sets the metadata for the directory source.
func WithMetadata(metadata map[string]any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadataValue adds a single metadata key-value pair.
func WithMetadataValue(key string, value any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFileExtensions sets the file extensions to filter by.
func WithFileExtensions(extensions []string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRecursive sets whether to process subdirectories recursively.
func WithRecursive(recursive bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCustomChunkingStrategy sets a custom chunking strategy for document splitting.
// This overrides the reader's default chunking strategy.
// Note: Most readers have their own optimal chunking strategy.
func WithCustomChunkingStrategy(strategy chunking.Strategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithChunkSize sets the chunk size for the reader's default chunking strategy.
func WithChunkSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithChunkOverlap sets the chunk overlap for the reader's default chunking strategy.
func WithChunkOverlap(overlap int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithOCRExtractor sets an OCR extractor for processing images in documents (e.g., PDFs).
func WithOCRExtractor(extractor ocr.Extractor) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTransformers sets transformers for document processing.
// Transformers are applied before and after chunking.
//
// Example:
//
//	source := dir.New(paths, dir.WithTransformers(
//	    transform.NewCharFilter("\n", "\t"),
//	    transform.NewCharDedup(" "),
//	))
func WithTransformers(transformers ...transform.Transformer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFileReaderType overrides the automatic file type detection for all files in the directory.
// This forces all files to be processed using the specified file type reader.
// Use predefined constants from source package for type safety.
//
// Example:
//
//	source := dir.New([]string{"./data"}, dir.WithFileReaderType(source.FileReaderTypeJSON))
func WithFileReaderType(fileType source.FileReaderType) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithExtractor sets a content extractor for handling complex or unsupported formats.
// When configured, the extractor is used for files whose extension matches
// the extractor's supported formats.
//
// Example:
//
//	source := dir.New(paths, dir.WithExtractor(myVisionExtractor))
func WithExtractor(e extractor.Extractor) Option { _ = "STUB: not implemented"; return *new(Option) }
