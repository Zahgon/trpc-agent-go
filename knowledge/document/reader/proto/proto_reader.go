//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package proto provides protocol buffer definition file reader implementation.
package proto

import (
	"io"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/internal/codeast"
	codeproto "trpc.group/trpc-go/trpc-agent-go/knowledge/internal/codeast/proto"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/transform"
)

var (
	// supportedExtensions defines the file extensions supported by this reader.
	supportedExtensions = []string{".proto"}
)

// init registers the proto reader with the global registry.
func init() {
	reader.RegisterReader(supportedExtensions, New)
}

// Reader reads protocol buffer definition files and extracts AST-based entities.
type Reader struct {
	chunk        bool
	transformers []transform.Transformer
	parser       *codeproto.Parser
}

// New creates a new proto reader with the given options.
func New(opts ...reader.Option) reader.Reader {
	_ = "STUB: not implemented"
	return *new(reader.Reader)
}

// ReadFromReader reads proto content from an io.Reader and returns a list of documents.
func (r *Reader) ReadFromReader(name string, rd io.Reader) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadFromFile reads a proto file and returns a list of documents.
func (r *Reader) ReadFromFile(filePath string) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadFromURL reads proto content from a URL and returns a list of documents.
func (r *Reader) ReadFromURL(urlStr string) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// processContent processes proto content and extracts AST-based entities.
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

// extractFileNameFromURL extracts a file name from a URL.
func (r *Reader) extractFileNameFromURL(url string) string { _ = "STUB: not implemented"; return "" }

// applyTransformers applies all transformers to the documents.
func (r *Reader) applyTransformers(docs []*document.Document) ([]*document.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Name returns the name of this reader.
func (r *Reader) Name() string { _ = "STUB: not implemented"; return "" }

// SupportedExtensions returns the file extensions this reader supports.
func (r *Reader) SupportedExtensions() []string { _ = "STUB: not implemented"; return nil }
