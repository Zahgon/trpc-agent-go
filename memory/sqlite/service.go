//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package sqlite provides a SQLite-backed memory service implementation.
package sqlite

import (
	"context"
	"database/sql"

	"trpc.group/trpc-go/trpc-agent-go/memory"
	imemory "trpc.group/trpc-go/trpc-agent-go/memory/internal/memory"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

var _ memory.Service = (*Service)(nil)

// Service is the sqlite memory service.
type Service struct {
	opts      ServiceOpts
	db        *sql.DB
	tableName string

	cachedTools      map[string]tool.Tool
	precomputedTools []tool.Tool
	autoMemoryWorker *imemory.AutoMemoryWorker
}

// NewService creates a new sqlite memory service.
//
// The service owns the passed-in db and will close it in Close().
func NewService(db *sql.DB, options ...ServiceOpt) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddMemory adds or updates a memory for a user (idempotent).
func (s *Service) AddMemory(
	ctx context.Context,
	userKey memory.UserKey,
	memoryStr string,
	topics []string,
	opts ...memory.AddOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) getDeletedAt(
	ctx context.Context,
	userKey memory.UserKey,
	memoryID string,
) (sql.NullInt64, bool, error) {
	_ = "STUB: not implemented"
	return *new(sql.NullInt64), false, nil
}

func (s *Service) enforceMemoryLimit(
	ctx context.Context,
	userKey memory.UserKey,
) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateMemory updates an existing memory for a user.
func (s *Service) UpdateMemory(
	ctx context.Context,
	memoryKey memory.Key,
	memoryStr string,
	topics []string,
	opts ...memory.UpdateOption,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) getEntry(
	ctx context.Context,
	memoryKey memory.Key,
) (*memory.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteMemory deletes a memory for a user.
func (s *Service) DeleteMemory(
	ctx context.Context,
	memoryKey memory.Key,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ClearMemories clears all memories for a user.
func (s *Service) ClearMemories(
	ctx context.Context,
	userKey memory.UserKey,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadMemories reads memories for a user.
func (s *Service) ReadMemories(
	ctx context.Context,
	userKey memory.UserKey,
	limit int,
) ([]*memory.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SearchMemories searches memories for a user.
func (s *Service) SearchMemories(
	ctx context.Context,
	userKey memory.UserKey,
	queryStr string,
	opts ...memory.SearchOption,
) ([]*memory.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tools returns the list of available memory tools.
// In auto memory mode (extractor is set), memory_search is exposed by default,
// memory_load is exposed once enabled, and other enabled tools remain hidden
// unless explicitly exposed.
// Without an extractor, enabled tools are exposed directly.
func (s *Service) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

// EnqueueAutoMemoryJob enqueues an auto memory job.
func (s *Service) EnqueueAutoMemoryJob(
	ctx context.Context,
	sess *session.Session,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Close closes the service and releases resources.
func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }
