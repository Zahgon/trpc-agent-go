//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package document provides a document internal utils.
package document

import (
	"trpc.group/trpc-go/trpc-agent-go/knowledge/document"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/internal/codeast"
)

// CreateDocument creates a new document with the given content and name.
func CreateDocument(content string, name string) *document.Document {
	_ = "STUB: not implemented"
	return nil
}

// CreateDocumentFromPayload creates a document from a code AST payload.
func CreateDocumentFromPayload(payload *codeast.DocumentPayload) *document.Document {
	_ = "STUB: not implemented"
	return nil
}

// GenerateDocumentID generates a unique ID for a document.
// Uses content hash for identification and random bytes for uniqueness.
func GenerateDocumentID(name string, content string) string {
	_ = "STUB: not implemented"
	// Content hash (first 8 bytes = 16 hex chars)
	return ""
}

// Random bytes for uniqueness (8 bytes = 16 hex chars)

// Fallback to timestamp-based uniqueness if crypto/rand fails
