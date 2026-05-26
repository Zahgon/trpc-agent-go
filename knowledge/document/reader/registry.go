//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package reader defines the interface for document readers.
package reader

import (
	"sync"
)

// Builder is a function that creates a new Reader instance with options.
type Builder func(opts ...Option) Reader

// Registry manages registration of document readers.
type Registry struct {
	mu      sync.RWMutex
	readers map[string]Builder // extension -> builder
}

// globalRegistry is the singleton registry instance.
var globalRegistry = &Registry{
	readers: make(map[string]Builder),
}

// RegisterReader registers a reader builder for specific file extensions.
// Extensions should include the dot prefix (e.g., ".pdf", ".txt").
func RegisterReader(extensions []string, builder Builder) { _ = "STUB: not implemented"; return }

// Normalize extension to lowercase.

// GetReader returns a new reader instance for the given file extension with options.
// The extension should include the dot prefix (e.g., ".pdf").
// Returns nil and false if no reader is registered for the extension.
func GetReader(extension string, opts ...Option) (Reader, bool) {
	_ = "STUB: not implemented"
	return *new(Reader), false
}

// Create a new instance with options

// GetAllReaders returns all registered readers as a map of file type to reader.
// The returned map uses simplified type names (e.g., "text", "pdf") as keys.
// Each call creates new reader instances with the provided options.
func GetAllReaders(opts ...Option) map[string]Reader { _ = "STUB: not implemented"; return nil }

// Skip if we've already processed this type.

// Create a new instance with options

// extensionToType converts a file extension to a simplified type name.
func extensionToType(ext string) string {
	_ = "STUB: not implemented"
	// Remove the dot prefix if present.
	return ""
}

// Map common extensions to type names.

// GetRegisteredExtensions returns all registered file extensions.
func GetRegisteredExtensions() []string { _ = "STUB: not implemented"; return nil }

// ClearRegistry clears all registered readers (mainly for testing).
func ClearRegistry() { _ = "STUB: not implemented"; return }
