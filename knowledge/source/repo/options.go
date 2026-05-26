//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package repo provides repository-based knowledge source implementation.
package repo

import (
	"trpc.group/trpc-go/trpc-agent-go/knowledge/transform"
)

// Option represents a functional option for configuring repository sources.
type Option func(*Source)

// WithRepository sets the single repository handled by this Source.
//
// Calling WithRepository more than once keeps the first repository value in
// Source.repository, leaves Source.hasRepository set, and records the duplicate
// configuration in Source.multiRepoError. The deferred error is surfaced by
// Source.ReadDocuments rather than by the option itself.
func WithRepository(repository Repository) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithName sets the source name.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadata sets custom metadata for the source.
func WithMetadata(metadata map[string]any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMetadataValue adds a single metadata key-value pair.
func WithMetadataValue(key string, value any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithFileExtensions limits processing to the given file extensions.
func WithFileExtensions(extensions []string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSkipDirs configures directory names to skip during scanning.
func WithSkipDirs(dirs []string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSkipSuffixes configures file suffixes to skip during scanning.
func WithSkipSuffixes(suffixes []string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTransformers sets document transformers.
func WithTransformers(transformers ...transform.Transformer) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
