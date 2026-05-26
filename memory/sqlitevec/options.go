//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package sqlitevec provides a SQLite-backed memory service powered by
// sqlite-vec for vector similarity search.
package sqlitevec

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/embedder"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/memory/extractor"
	imemory "trpc.group/trpc-go/trpc-agent-go/memory/internal/memory"
)

const (
	defaultDBInitTimeout  = 30 * time.Second
	defaultMaxResults     = 10
	defaultIndexDimension = 1536
)

var defaultOptions = ServiceOpts{
	tableName:      defaultTableName,
	indexDimension: defaultIndexDimension,
	maxResults:     defaultMaxResults,

	memoryLimit:      imemory.DefaultMemoryLimit,
	toolCreators:     imemory.AllToolCreators,
	enabledTools:     imemory.DefaultEnabledTools,
	asyncMemoryNum:   imemory.DefaultAsyncMemoryNum,
	memoryQueueSize:  imemory.DefaultMemoryQueueSize,
	memoryJobTimeout: imemory.DefaultMemoryJobTimeout,
}

// ServiceOpts is the options for the sqlite-vec memory service.
type ServiceOpts struct {
	tableName      string
	indexDimension int
	maxResults     int

	memoryLimit int
	softDelete  bool

	// Embedder for generating embeddings for memories and queries.
	embedder embedder.Embedder

	// Tool related settings.
	toolCreators      map[string]memory.ToolCreator
	enabledTools      map[string]struct{}
	toolExposed       map[string]struct{}
	toolHidden        map[string]struct{}
	userExplicitlySet map[string]struct{}

	// skipDBInit skips database initialization (table creation).
	skipDBInit bool

	// Memory extractor for auto memory mode.
	extractor extractor.MemoryExtractor

	// Async memory worker configuration.
	asyncMemoryNum   int
	memoryQueueSize  int
	memoryJobTimeout time.Duration
}

func (o ServiceOpts) clone() ServiceOpts { _ = "STUB: not implemented"; return *new(ServiceOpts) }

// ServiceOpt is the option for the sqlite-vec memory service.
type ServiceOpt func(*ServiceOpts)

// WithTableName sets the table name for storing memories.
// Default is "memories".
//
// Security: Uses internal/session/sqldb.ValidateTableName to prevent SQL
// injection.
func WithTableName(tableName string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithIndexDimension sets the vector dimension for sqlite-vec table.
// If not set, it defaults to the embedder's dimension.
func WithIndexDimension(dim int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithMaxResults sets the max number of results returned by SearchMemories.
func WithMaxResults(maxResults int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithEmbedder sets the embedder for generating embeddings.
func WithEmbedder(e embedder.Embedder) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithMemoryLimit sets the limit of memories per user.
func WithMemoryLimit(limit int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSoftDelete enables or disables soft delete behavior.
// When enabled, delete operations set deleted_at and queries filter deleted
// rows. Default is disabled (hard delete).
func WithSoftDelete(enabled bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSkipDBInit skips database initialization (table creation).
func WithSkipDBInit(skip bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithExtractor sets the memory extractor for auto memory mode.
// When enabled, auto mode defaults are applied to enabledTools, but user
// settings via WithToolEnabled (before or after) take precedence.
func WithExtractor(e extractor.MemoryExtractor) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithAsyncMemoryNum sets the number of async memory workers.
func WithAsyncMemoryNum(num int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithMemoryQueueSize sets the queue size for memory jobs.
func WithMemoryQueueSize(size int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithMemoryJobTimeout sets the timeout for each memory job.
func WithMemoryJobTimeout(timeout time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithCustomTool sets a custom memory tool implementation.
// The tool will be enabled by default.
func WithCustomTool(toolName string, creator memory.ToolCreator) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithToolEnabled enables or disables a memory tool by name.
// User settings take precedence over auto mode defaults regardless of
// option order.
func WithToolEnabled(toolName string, enabled bool) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithAutoMemoryExposedTools exposes enabled tools via Tools() in auto memory
// mode so the agent can call them directly. Invalid tool names are ignored.
func WithAutoMemoryExposedTools(toolNames ...string) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithToolExposed controls whether an enabled memory tool is exposed via
// Tools(). Use WithAutoMemoryExposedTools for the common auto memory case.
func WithToolExposed(toolName string, exposed bool) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}
