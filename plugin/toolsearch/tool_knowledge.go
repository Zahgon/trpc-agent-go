//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package toolsearch

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/embedder"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/vectorstore"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// ToolKnowledge is a tool knowledge base that uses a vector store to store tools and their embeddings.
type ToolKnowledge struct {
	s vectorstore.VectorStore
	e embedder.Embedder

	mu    sync.RWMutex
	tools map[string]tool.Tool
}

// ToolKnowledgeOption is a function that configures the ToolKnowledge.
type ToolKnowledgeOption func(*ToolKnowledge)

// WithVectorStore sets the vector store for the ToolKnowledge.
func WithVectorStore(s vectorstore.VectorStore) ToolKnowledgeOption {
	_ = "STUB: not implemented"
	return *new(ToolKnowledgeOption)
}

// NewToolKnowledge creates a new ToolKnowledge.
func NewToolKnowledge(e embedder.Embedder, opts ...ToolKnowledgeOption) (*ToolKnowledge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default vector store

func (k *ToolKnowledge) search(ctx context.Context, candidates map[string]tool.Tool, query string, topK int) (context.Context, []string, *model.Usage, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil, nil
}

func (k *ToolKnowledge) upsert(ctx context.Context, ts map[string]tool.Tool) (*model.Usage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toolToText(t tool.Tool) string { _ = "STUB: not implemented"; return "" }

// Add parameter information (mirrors the Python implementation).

// Best-effort inference (useful for partially-filled schemas).
