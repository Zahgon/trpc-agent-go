//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package postgres

import (
	"regexp"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/memory/extractor"
	imemory "trpc.group/trpc-go/trpc-agent-go/memory/internal/memory"
)

const (
	defaultHost     = "localhost"
	defaultPort     = 5432
	defaultDatabase = "trpc-agent-go-pgmemory"
	defaultSSLMode  = "disable"
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

// ServiceOpts is the options for the postgres memory service.
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

	tableName   string
	memoryLimit int
	softDelete  bool
	// keyword-search settings.
	searchMinScore   float64
	maxSearchResults int

	// Tool related settings.
	toolCreators      map[string]memory.ToolCreator
	enabledTools      map[string]struct{}
	toolExposed       map[string]struct{}
	toolHidden        map[string]struct{}
	userExplicitlySet map[string]struct{}

	// skipDBInit skips database initialization (table and index creation).
	// Useful when user doesn't have DDL permissions or when tables are managed externally.
	skipDBInit bool

	// schema is the PostgreSQL schema name where tables are created.
	// Default is empty string (uses default schema, typically "public").
	schema string

	// Memory extractor for auto memory mode.
	extractor extractor.MemoryExtractor

	// Async memory worker configuration.
	asyncMemoryNum   int
	memoryQueueSize  int
	memoryJobTimeout time.Duration
}

func (o ServiceOpts) clone() ServiceOpts { _ = "STUB: not implemented"; return *new(ServiceOpts) }

// Initialize userExplicitlySet map (empty for new clone).

// ServiceOpt is the option for the postgres memory service.
type ServiceOpt func(*ServiceOpts)

// WithPostgresClientDSN sets the PostgreSQL DSN connection string directly (recommended).
// Example: "postgres://user:password@localhost:5432/dbname?sslmode=disable"
//
// Note: WithPostgresClientDSN has the highest priority.
// If DSN is specified, other connection settings (WithHost, WithPort, etc.) will be ignored.
func WithPostgresClientDSN(dsn string) ServiceOpt {
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
// Note: Direct connection settings (WithHost, WithPort, etc.) have higher priority than WithPostgresInstance.
// If both are specified, direct connection settings will be used.
func WithPostgresInstance(instanceName string) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithTableName sets the table name for storing memories.
// Default is "memories".
// Table name must start with a letter or underscore and contain only alphanumeric characters and underscores.
// Maximum length is 63 characters (PostgreSQL limit).
// Panics if the table name is invalid.
func WithTableName(tableName string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSoftDelete enables or disables soft delete behavior.
// When enabled, delete operations set deleted_at and queries filter deleted rows.
// Default is disabled (hard delete).
func WithSoftDelete(enabled bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

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

// WithExtraOptions sets the extra options for the postgres memory service.
// These options will be passed to the PostgreSQL client builder.
func WithExtraOptions(extraOptions ...any) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSkipDBInit skips database initialization (table and index creation).
// Useful when:
// - User doesn't have DDL permissions
// - Tables are managed by migration tools
// - Running in production environment where schema is pre-created
func WithSkipDBInit(skip bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSchema sets the PostgreSQL schema name where tables will be created.
// If not set, tables will be created in the default schema (typically "public").
// For example, with schema "my_schema", tables will be qualified as:
// - my_schema.memories
//
// Note: The schema must already exist in the database before using this option.
// Security: Uses internal/session/sqldb.ValidateTableName to prevent SQL injection.
func WithSchema(schema string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// Use internal/session/sqldb validation.

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
// - Not exceed 63 characters (PostgreSQL limit).
func validateTableName(tableName string) error { _ = "STUB: not implemented"; return nil }
