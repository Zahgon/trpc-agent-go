//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package mysql

import (
	"regexp"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/memory/extractor"
	imemory "trpc.group/trpc-go/trpc-agent-go/memory/internal/memory"
)

const (
	// defaultDBInitTimeout is the default timeout for database initialization.
	defaultDBInitTimeout = 30 * time.Second
)

var (
	defaultOptions = ServiceOpts{
		tableName:        "memories",
		memoryLimit:      imemory.DefaultMemoryLimit,
		searchMinScore:   imemory.DefaultSearchMinScore,
		maxSearchResults: imemory.DefaultMaxSearchResults,
		toolCreators:     imemory.AllToolCreators,
		enabledTools:     imemory.DefaultEnabledTools,
		asyncMemoryNum:   imemory.DefaultAsyncMemoryNum,
	}
)

// ServiceOpts is the options for the mysql memory service.
type ServiceOpts struct {
	dsn          string
	instanceName string
	memoryLimit  int
	tableName    string
	softDelete   bool
	// keyword-search settings.
	searchMinScore   float64
	maxSearchResults int

	// Tool related settings.
	toolCreators      map[string]memory.ToolCreator
	enabledTools      map[string]struct{}
	toolExposed       map[string]struct{}
	toolHidden        map[string]struct{}
	userExplicitlySet map[string]struct{}
	extraOptions      []any

	// skipDBInit skips database initialization (table creation).
	// Useful when user doesn't have DDL permissions or when tables are managed externally.
	skipDBInit bool

	// Memory extractor for auto memory mode.
	extractor extractor.MemoryExtractor

	// Async memory worker configuration.
	asyncMemoryNum   int
	memoryQueueSize  int
	memoryJobTimeout time.Duration
}

func (o ServiceOpts) clone() ServiceOpts { _ = "STUB: not implemented"; return *new(ServiceOpts) }

// Initialize userExplicitlySet map (empty for new clone).

// ServiceOpt is the option for the mysql memory service.
type ServiceOpt func(*ServiceOpts)

// WithMySQLClientDSN sets the MySQL DSN connection string directly (recommended).
// Example: "user:password@tcp(localhost:3306)/dbname?parseTime=true&charset=utf8mb4"
//
// This is the preferred way to connect to MySQL as it:
// - Simplifies configuration (all connection params in one string)
// - Supports all MySQL connection parameters
// - Is consistent with session/mysql and storage/mysql
func WithMySQLClientDSN(dsn string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithMySQLInstance uses a MySQL instance from storage.
// The instance must be registered via storage.RegisterMySQLInstance() before use.
//
// Note: WithMySQLClientDSN has higher priority than WithMySQLInstance.
// If both are specified, DSN will be used.
func WithMySQLInstance(instanceName string) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithMemoryLimit sets the limit of memories per user.
func WithMemoryLimit(limit int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithMinSearchScore sets the minimum keyword-search score. Scores below
// this value are filtered out. Default is 0.3.
func WithMinSearchScore(score float64) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithMaxResults sets the maximum number of keyword-search results.
// Default is 10. Use 0 to disable truncation.
func WithMaxResults(maxResults int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithTableName sets the table name for storing memories.
// Default is "memories".
// Table name must start with a letter or underscore and contain only alphanumeric characters and underscores.
// Maximum length is 64 characters.
// Panics if the table name is invalid.
func WithTableName(tableName string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSoftDelete enables or disables soft delete behavior.
// When enabled, delete operations set deleted_at and queries filter deleted rows.
// Default is disabled (hard delete).
func WithSoftDelete(enabled bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithCustomTool sets a custom memory tool implementation.
// The tool will be enabled by default.
// If the tool name is invalid or creator is nil, this option will do nothing.
func WithCustomTool(toolName string, creator memory.ToolCreator) ServiceOpt {
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

// WithExtraOptions sets the extra options for the MySQL memory service.
// These options will be passed to the MySQL client builder.
func WithExtraOptions(extraOptions ...any) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSkipDBInit skips database initialization (table creation).
// Useful when:
// - User doesn't have DDL permissions
// - Tables are managed by migration tools
// - Running in production environment where schema is pre-created
func WithSkipDBInit(skip bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithExtractor sets the memory extractor for auto memory mode.
// When enabled, auto mode defaults are applied to enabledTools,
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

// tableNamePattern is the regex pattern for validating table names.
// Only allows alphanumeric characters and underscores, must start with a letter or underscore.
var tableNamePattern = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

// validateTableName validates the table name to prevent SQL injection.
// Table name must:
// - Start with a letter or underscore.
// - Contain only alphanumeric characters and underscores.
// - Not be empty.
// - Not exceed 64 characters (MySQL limit).
func validateTableName(tableName string) error { _ = "STUB: not implemented"; return nil }
