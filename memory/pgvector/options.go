//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package pgvector

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/embedder"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/memory/extractor"
	imemory "trpc.group/trpc-go/trpc-agent-go/memory/internal/memory"
)

// Default connection settings.
const (
	defaultHost     = "localhost"
	defaultPort     = 5432
	defaultDatabase = "trpc-agent-go-pgmemory"
	defaultSSLMode  = "disable"
)

// Default table and index settings.
const (
	defaultTableName      = "memories"
	defaultIndexDimension = 1536
	defaultMaxResults     = 15
)

// Default HNSW index parameters.
const (
	defaultHNSWM              = 16
	defaultHNSWEfConstruction = 64
)

// Default similarity threshold. Results with cosine similarity below this
// are filtered out even if within the top-K limit. A value of 0 disables
// threshold filtering. The default of 0.35 removes very low relevance
// results that add noise without contributing useful information.
const defaultSimilarityThreshold = 0.30

// Default timeout settings.
const (
	defaultDBInitTimeout = 30 * time.Second
)

// HNSWIndexParams contains parameters for HNSW index.
type HNSWIndexParams struct {
	// M is the maximum number of connections per layer (default: 16, range: 2-100).
	M int
	// EfConstruction is the size of dynamic candidate list for construction.
	// Default: 64, range: 4-1000.
	EfConstruction int
}

var defaultOptions = ServiceOpts{
	tableName:           defaultTableName,
	indexDimension:      defaultIndexDimension,
	maxResults:          defaultMaxResults,
	memoryLimit:         imemory.DefaultMemoryLimit,
	similarityThreshold: defaultSimilarityThreshold,
	toolCreators:        imemory.AllToolCreators,
	enabledTools:        imemory.DefaultEnabledTools,
	asyncMemoryNum:      imemory.DefaultAsyncMemoryNum,
	hnswParams: &HNSWIndexParams{
		M:              defaultHNSWM,
		EfConstruction: defaultHNSWEfConstruction,
	},
}

// ServiceOpts is the options for the pgvector memory service.
type ServiceOpts struct {
	// PostgreSQL connection settings.
	dsn      string
	host     string
	port     int
	user     string
	password string
	database string
	sslMode  string

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

	// Vector index configuration.
	hnswParams *HNSWIndexParams

	// Tool related settings.
	toolCreators      map[string]memory.ToolCreator
	enabledTools      map[string]struct{}
	toolExposed       map[string]struct{}
	toolHidden        map[string]struct{}
	userExplicitlySet map[string]struct{}

	// skipDBInit skips database initialization (table and index creation).
	// Useful when user doesn't have DDL permissions or when tables are managed
	// externally.
	skipDBInit bool

	// schema is the PostgreSQL schema name where tables are created.
	// Default is empty string (uses default schema, typically "public").
	schema string

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

// Clone HNSW params if present.

// ServiceOpt is the option for the pgvector memory service.
type ServiceOpt func(*ServiceOpts)

// WithPGVectorClientDSN sets the PostgreSQL DSN connection string directly.
// (recommended).
// Example: "postgres://user:password@localhost:5432/dbname?sslmode=disable".
//
// Note: WithPGVectorClientDSN has the highest priority.
// If DSN is specified, other connection settings (WithHost, WithPort, etc.).
// will be ignored.
func WithPGVectorClientDSN(dsn string) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithHost sets the PostgreSQL host.
func WithHost(host string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithPort sets the PostgreSQL port.
func WithPort(port int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithUser sets the username for authentication.
func WithUser(user string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithPassword sets the password for authentication.
func WithPassword(password string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithDatabase sets the database name.
func WithDatabase(database string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSSLMode sets the SSL mode for connection.
func WithSSLMode(sslMode string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithPostgresInstance uses a postgres instance from storage.
// Note: Direct connection settings (WithHost, WithPort, etc.) have higher.
// priority than WithPostgresInstance.
// If both are specified, direct connection settings will be used.
func WithPostgresInstance(instanceName string) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithTableName sets the table name for storing memories.
// Default is "memories".
//
// Panics if the table name is invalid.
func WithTableName(tableName string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithIndexDimension sets the vector dimension for the index.
// Default is 1536.
func WithIndexDimension(dimension int) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithMaxResults sets the maximum number of search results.
// Default is 10.
func WithMaxResults(maxResults int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSoftDelete enables or disables soft delete behavior.
// When enabled, delete operations set deleted_at and queries filter deleted rows.
// Default is disabled (hard delete).
func WithSoftDelete(enabled bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithMemoryLimit sets the limit of memories per user.
func WithMemoryLimit(limit int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithCustomTool sets a custom memory tool implementation.
// The tool will be enabled by default.
// If the tool name is invalid or creator is nil, this option will do nothing.
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
// If the tool name is invalid, this option will do nothing.
// User settings via WithToolEnabled take precedence over auto mode
// defaults, regardless of option order.
func WithToolEnabled(toolName string, enabled bool) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithExtraOptions sets the extra options for the postgres memory service.
// These options will be passed to the PostgreSQL client builder.
func WithExtraOptions(extraOptions ...any) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSkipDBInit skips database initialization (table and index creation).
// Useful when.
// - User doesn't have DDL permissions.
// - Tables are managed by migration tools.
// - Running in production environment where schema is pre-created.
func WithSkipDBInit(skip bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSchema sets the PostgreSQL schema name where tables will be created.
// If not set, tables will be created in the default schema (typically "public").
//
// Note: The schema must already exist in the database before using this option.
func WithSchema(schema string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithEmbedder sets the embedder for generating memory embeddings.
// This is required for vector-based memory search.
func WithEmbedder(e embedder.Embedder) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithExtractor sets the memory extractor for auto memory mode.
// When enabled, auto mode defaults are applied to enabledTools.
// but user settings via WithToolEnabled (before or after) take precedence.
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

// WithHNSWIndexParams sets HNSW index parameters.
func WithHNSWIndexParams(params *HNSWIndexParams) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSimilarityThreshold sets the minimum cosine similarity threshold
// for search results. Results below this threshold are filtered out.
// Value should be between 0 and 1. A value of 0 disables filtering.
// Default is 0 (disabled).
func WithSimilarityThreshold(threshold float64) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}
