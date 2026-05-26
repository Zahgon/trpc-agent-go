//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package qdrant provides a Qdrant-based implementation of the VectorStore interface.
package qdrant

import (
	"context"
	"sync"

	"github.com/qdrant/go-client/qdrant"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/searchfilter"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/vectorstore"
	qdrantstorage "trpc.group/trpc-go/trpc-agent-go/storage/qdrant"
)

var _ vectorstore.VectorStore = (*VectorStore)(nil)

// VectorStore implements vectorstore.VectorStore using Qdrant.
type VectorStore struct {
	client          qdrantstorage.Client
	ownsClient      bool
	opts            options
	filterConverter searchfilter.Converter[*qdrant.Filter]
	retryCfg        retryConfig
	closeOnce       sync.Once
}

// New creates a new Qdrant VectorStore.
func New(ctx context.Context, opts ...Option) (*VectorStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate configuration

// Close closes the connection if it was created by the VectorStore.
// Safe to call multiple times; only the first call closes the client.
func (vs *VectorStore) Close() error { _ = "STUB: not implemented"; return nil }
