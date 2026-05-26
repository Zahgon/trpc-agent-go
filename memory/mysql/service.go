//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package mysql provides the mysql memory service.
package mysql

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/memory"
	imemory "trpc.group/trpc-go/trpc-agent-go/memory/internal/memory"
	"trpc.group/trpc-go/trpc-agent-go/session"
	storage "trpc.group/trpc-go/trpc-agent-go/storage/mysql"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

var _ memory.Service = (*Service)(nil)

// Service is the mysql memory service.
// Storage structure:
//
//	Table: memories
//	Columns: app_name, user_id, memory_id, memory_data (JSON), created_at, updated_at.
//	Primary Key: (app_name, user_id, memory_id).
//	Index: (app_name, user_id).
type Service struct {
	opts      ServiceOpts
	db        storage.Client
	tableName string

	cachedTools      map[string]tool.Tool
	precomputedTools []tool.Tool
	autoMemoryWorker *imemory.AutoMemoryWorker
}

// NewService creates a new mysql memory service.
func NewService(options ...ServiceOpt) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Apply user options.
}

// Apply auto mode defaults after all options are applied.
// User settings via WithToolEnabled take precedence regardless of option order.

// Priority: dsn > instanceName.

// Initialize database if needed

// Pre-compute tools list to avoid lock contention in Tools() method.

// Initialize auto memory worker if extractor is configured.

// AddMemory adds or updates a memory for a user (idempotent).
func (s *Service) AddMemory(ctx context.Context, userKey memory.UserKey,
	memoryStr string, topics []string,
	opts ...memory.AddOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Enforce memory limit.

// UpdateMemory updates an existing memory for a user.
func (s *Service) UpdateMemory(ctx context.Context, memoryKey memory.Key,
	memoryStr string, topics []string,
	opts ...memory.UpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteMemory deletes a memory for a user (soft delete).
func (s *Service) DeleteMemory(ctx context.Context, memoryKey memory.Key) error {
	_ = "STUB: not implemented"
	return nil
}

// ClearMemories clears all memories for a user (soft delete).
func (s *Service) ClearMemories(ctx context.Context, userKey memory.UserKey) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadMemories reads memories for a user.
func (s *Service) ReadMemories(ctx context.Context, userKey memory.UserKey, limit int) ([]*memory.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SearchMemories searches memories for a user.
// Results are ranked by a normalized BM25-style keyword score and only
// the top entries above the minimum threshold are returned.
func (s *Service) SearchMemories(ctx context.Context, userKey memory.UserKey,
	query string, opts ...memory.SearchOption) ([]*memory.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tools returns the list of available memory tools.
// In auto memory mode (extractor is set), memory_search is exposed by default,
// memory_load is exposed once enabled, and other enabled tools remain hidden
// unless explicitly exposed.
// Without an extractor, enabled tools are exposed directly.
// The tools list is pre-computed at service creation time.
func (s *Service) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

// EnqueueAutoMemoryJob enqueues an auto memory extraction job for async
// processing. The session contains the full transcript and state for
// incremental extraction.
func (s *Service) EnqueueAutoMemoryJob(ctx context.Context, sess *session.Session) error {
	_ = "STUB: not implemented"
	return nil
}

// Close closes the database connection and stops async workers.
func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }
