//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const createDocumentToolName = "create_document"

type createDocumentArgs struct {
	Title   string `json:"title" description:"A short title for the generated document."`
	Content string `json:"content" description:"The complete generated document content to save."`
}

type createDocumentResult struct {
	DocumentID string `json:"document_id"`
	Title      string `json:"title"`
	Bytes      int    `json:"bytes"`
}

type savedDocument struct {
	ID        string
	Title     string
	Content   string
	CreatedAt time.Time
}

type documentStore struct {
	mu        sync.Mutex
	nextID    int64
	documents map[string]savedDocument
}

func newDocumentStore() *documentStore { _ = "STUB: not implemented"; return nil }

func newCreateDocumentTool(store *documentStore) tool.Tool {
	_ = "STUB: not implemented"
	return *new(tool.Tool)
}

func (s *documentStore) createDocument(ctx context.Context, args createDocumentArgs) (createDocumentResult, error) {
	_ = "STUB: not implemented"
	return *new(createDocumentResult), nil
}
