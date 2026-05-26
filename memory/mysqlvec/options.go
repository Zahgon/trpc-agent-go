//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package mysqlvec

import (
	"regexp"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/embedder"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/memory/extractor"
	imemory "trpc.group/trpc-go/trpc-agent-go/memory/internal/memory"
)

// Default settings.
const (
	defaultTableName      = "memories"
	defaultIndexDimension = 1536
	defaultMaxResults     = 15
	defaultDBInitTimeout  = 30 * time.Second
)

// Default similarity threshold. Results with cosine similarity below this
// are filtered out. The default of 0.30 removes very low relevance results.
const defaultSimilarityThreshold = 0.30

var defaultOptions = ServiceOpts{
	tableName:           defaultTableName,
	indexDimension:      defaultIndexDimension,
	maxResults:          defaultMaxResults,
	memoryLimit:         imemory.DefaultMemoryLimit,
	similarityThreshold: defaultSimilarityThreshold,
	toolCreators:        imemory.AllToolCreators,
	enabledTools:        imemory.DefaultEnabledTools,
	asyncMemoryNum:      imemory.DefaultAsyncMemoryNum,
}

// ServiceOpts is the options for the mysqlvec memory service.
type ServiceOpts struct {
	// MySQL connection settings.
	dsn          string
	instanceName string
	extraOptions []any

	tableName      string
	indexDimension int
	maxResults     int
	memoryLimit    int
	softDelete     bool

	// similarityThreshold filters out search results with cosine similarity
	// below this value (range 0-1). A value of 0 disables filtering.
	similarityThreshold float64

	// Tool related settings.
	toolCreators      map[string]memory.ToolCreator
	enabledTools      map[string]struct{}
	toolExposed       map[string]struct{}
	toolHidden        map[string]struct{}
	userExplicitlySet map[string]struct{}

	// skipDBInit skips database initialization (table and index creation).
	skipDBInit bool

	// Embedder for generating memory embeddings.
	embedder embedder.Embedder

	// Memory extractor for auto memory mode.
	extractor extractor.MemoryExtractor

	// Async memory worker configuration.
	asyncMemoryNum   int
	memoryQueueSize  int
	memoryJobTimeout time.Duration
}

func (o ServiceOpts) clone() ServiceOpts { _ = "STUB: not implemented"; return *new(ServiceOpts) }

// Initialize userExplicitlySet map (empty for new clone).

// ServiceOpt is the option for the mysqlvec memory service.
type ServiceOpt func(*ServiceOpts)

// WithMySQLClientDSN sets the MySQL DSN connection string directly (recommended).
// Example: "user:password@tcp(localhost:3306)/dbname?parseTime=true&charset=utf8mb4"
//
// Note: WithMySQLClientDSN has the highest priority.
// If DSN is specified, WithMySQLInstance will be ignored.
func WithMySQLClientDSN(dsn string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithMySQLInstance uses a MySQL instance from storage.
// The instance must be registered via storage.RegisterMySQLInstance() before use.
//
// Note: WithMySQLClientDSN has higher priority than WithMySQLInstance.
func WithMySQLInstance(instanceName string) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithTableName sets the table name for storing memories.
// Default is "memories".
//
// Panics if the table name is invalid.
func WithTableName(tableName string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithIndexDimension sets the vector dimension for the embedding column.
// Default is 1536.
func WithIndexDimension(dimension int) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithMaxResults sets the maximum number of search results.
// Default is 15.
func WithMaxResults(maxResults int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSoftDelete enables or disables soft delete behavior.
// When enabled, delete operations set deleted_at and queries filter deleted rows.
// Default is disabled (hard delete).
func WithSoftDelete(enabled bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithMemoryLimit sets the limit of memories per user.
func WithMemoryLimit(limit int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithCustomTool sets a custom memory tool implementation.
// The tool will be enabled by default.
func WithCustomTool(toolName string, creator memory.ToolCreator) ServiceOpt {
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

// WithToolEnabled sets which tool is enabled.
// User settings via WithToolEnabled take precedence over auto mode
// defaults, regardless of option order.
func WithToolEnabled(toolName string, enabled bool) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithExtraOptions sets extra options passed to the MySQL client builder.
func WithExtraOptions(extraOptions ...any) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSkipDBInit skips database initialization (table creation).
func WithSkipDBInit(skip bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithEmbedder sets the embedder for generating memory embeddings.
// This is required for vector-based memory search.
func WithEmbedder(e embedder.Embedder) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithExtractor sets the memory extractor for auto memory mode.
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

// WithSimilarityThreshold sets the minimum cosine similarity threshold
// for search results. Results below this threshold are filtered out.
// Value should be between 0 and 1. A value of 0 disables filtering.
func WithSimilarityThreshold(threshold float64) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// tableNamePattern is the regex pattern for validating table names.
var tableNamePattern = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

// validateTableName validates the table name to prevent SQL injection.
func validateTableName(tableName string) error { _ = "STUB: not implemented"; return nil }
