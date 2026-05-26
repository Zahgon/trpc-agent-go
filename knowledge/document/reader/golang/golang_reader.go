//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package golang provides Go source file reader implementation.
package golang

import (
	"io"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/internal/codeast"
	codegolang "trpc.group/trpc-go/trpc-agent-go/knowledge/internal/codeast/golang"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/transform"
)

var (
	// supportedExtensions defines the file extensions supported by this reader.
	supportedExtensions = []string{".go"}
)

// init registers the Go reader with the global registry.
func init() {
	reader.RegisterReader(supportedExtensions, New)
}

// Reader reads Go files and extracts AST-based entities.
type Reader struct {
	chunk        bool
	transformers []transform.Transformer
	parser       *codegolang.Parser
}

// New creates a new Go reader with the given options.
func New(opts ...reader.Option) reader.Reader {
	_ = "STUB: not implemented"
	return *new(reader.Reader)
}

// ReadFromReader reads Go content from an io.Reader and returns a list of documents.
func (r *Reader) ReadFromReader(name string, rd io.Reader) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadFromFile reads a Go file and returns a list of AST entity documents.
func (r *Reader) ReadFromFile(filePath string) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadFromURL reads Go content from a URL and returns a list of documents.
func (r *Reader) ReadFromURL(urlStr string) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadFromDirectory reads a Go module or directory and returns AST entity documents.
// It performs package-aware parsing across the directory instead of processing files independently.
func (r *Reader) ReadFromDirectory(dirPath string) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) processContent(content, name string, baseMetadata map[string]any) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) nodesToDocuments(result *codeast.Result, baseMetadata map[string]any) []*document.Document {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reader) createFileDocument(content, name string, baseMetadata map[string]any) *document.Document {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reader) createFileDocumentFromInfo(content, name string, baseMetadata map[string]any, fileInfo *codeast.FileInfo) *document.Document {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reader) applyTransformers(docs []*document.Document) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) extractFileNameFromURL(urlStr string) string { _ = "STUB: not implemented"; return "" }

// Name returns the name of this reader.
func (r *Reader) Name() string {
	_ = "STUB: not implemented"

	// SupportedExtensions returns the file extensions this reader supports.
	return ""
}

func (r *Reader) SupportedExtensions() []string { _ = "STUB: not implemented"; return nil }

// resolveScope returns the AST scope ("code" or "example") for a file-level
// document. When baseMetadata provides a repository root under
// source.MetaRepoPath, detection is anchored at that root.
func resolveScope(filePath string, baseMetadata map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

func repoRootFromMetadata(baseMetadata map[string]any) string { _ = "STUB: not implemented"; return "" }
