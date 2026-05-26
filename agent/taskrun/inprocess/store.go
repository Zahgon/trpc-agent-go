//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package inprocess

import (
	"context"
	"sync"
	"time"
)

const (
	storeVersion = 1

	storeDirPerm  = 0o700
	storeFilePerm = 0o600

	storeTempSuffix = ".tmp"
)

const (
	errInterruptedByRestart = "interrupted by previous runtime restart"
)

// Store persists task runs.
type Store interface {
	Load(ctx context.Context) ([]Run, error)
	Save(ctx context.Context, runs []Run) error
}

// MemoryStore stores runs in memory. It is useful for tests and stateless
// applications.
type MemoryStore struct {
	mu   sync.Mutex
	runs []Run
}

// NewMemoryStore creates an in-memory store.
func NewMemoryStore() *MemoryStore { _ = "STUB: not implemented"; return nil }

// Load implements Store.
func (s *MemoryStore) Load(ctx context.Context) ([]Run, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Save implements Store.
func (s *MemoryStore) Save(ctx context.Context, runs []Run) error {
	_ = "STUB: not implemented"
	return nil
}

// FileStore persists runs into one JSON file.
type FileStore struct {
	path string
}

// NewFileStore creates a JSON file store.
func NewFileStore(path string) (*FileStore, error) { _ = "STUB: not implemented"; return nil, nil }

// Load implements Store.
func (s *FileStore) Load(ctx context.Context) ([]Run, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Save implements Store.
func (s *FileStore) Save(ctx context.Context, runs []Run) error {
	_ = "STUB: not implemented"
	return nil
}

type storeFile struct {
	Version int   `json:"version"`
	Runs    []Run `json:"runs,omitempty"`
}

func normalizeLoadedRuns(runs map[string]*Run, now time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func ctxErr(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
