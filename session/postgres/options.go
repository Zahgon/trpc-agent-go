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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/session/summary"
)

const (
	defaultSessionEventLimit     = 1000
	defaultChanBufferSize        = 100
	defaultAsyncPersisterNum     = 10
	defaultCleanupIntervalSecond = 5 * time.Minute // 5 min
	defaultAsyncPersistTimeout   = 5 * time.Second

	defaultAsyncSummaryNum   = 3
	defaultSummaryQueueSize  = 100
	defaultSummaryJobTimeout = 60 * time.Second

	defaultHost     = "localhost"
	defaultPort     = 5432
	defaultDatabase = "trpc-agent-go-pgsession"
	defaultSSLMode  = "disable"
)

// ServiceOpts is the options for the postgres session service.
type ServiceOpts struct {
	sessionEventLimit int

	// PostgreSQL connection settings
	dsn      string
	host     string
	port     int
	user     string
	password string
	database string
	sslMode  string

	instanceName string
	extraOptions []any

	sessionTTL         time.Duration // TTL for session state and event list
	appStateTTL        time.Duration // TTL for app state
	userStateTTL       time.Duration // TTL for user state
	enableAsyncPersist bool
	asyncPersisterNum  int           // number of worker goroutines for async persistence
	softDelete         bool          // enable soft delete (default: true)
	cleanupInterval    time.Duration // interval for automatic cleanup of expired data
	// summarizer integrates LLM summarization.
	summarizer summary.SessionSummarizer
	// asyncSummaryNum is the number of worker goroutines for async summary.
	asyncSummaryNum int
	// summaryQueueSize is the size of summary job queue.
	summaryQueueSize int
	// summaryJobTimeout is the timeout for processing a single summary job.
	summaryJobTimeout time.Duration
	// summaryFilterAllowlist restricts which non-empty filterKeys may trigger
	// branch summaries.
	summaryFilterAllowlist []string
	// cascadeFullSessionSummary controls whether allowed branch summaries also
	// refresh the full-session summary. Nil preserves the legacy default of
	// enabling full-session cascade for zero-value options.
	cascadeFullSessionSummary *bool
	// skipDBInit skips database initialization (table and index creation).
	// Useful when user doesn't have DDL permissions or when tables are managed externally.
	skipDBInit bool
	// tablePrefix is the prefix for all table names.
	// Default is empty string (no prefix).
	tablePrefix string
	// schema is the PostgreSQL schema name where tables are created.
	// Default is empty string (uses default schema, typically "public").
	schema string
	// hooks for session operations.
	appendEventHooks []session.AppendEventHook
	getSessionHooks  []session.GetSessionHook
}

// ServiceOpt is the option for the postgres session service.
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

var (
	defaultOptions = ServiceOpts{
		sessionEventLimit:  defaultSessionEventLimit,
		sessionTTL:         0,
		appStateTTL:        0,
		userStateTTL:       0,
		asyncPersisterNum:  defaultAsyncPersisterNum,
		enableAsyncPersist: false,
		asyncSummaryNum:    defaultAsyncSummaryNum,
		summaryQueueSize:   defaultSummaryQueueSize,
		summaryJobTimeout:  defaultSummaryJobTimeout,
		softDelete:         true, // Enable soft delete by default
		cleanupInterval:    0,
	}
)

func (opts ServiceOpts) shouldCascadeFullSessionSummary() bool {
	_ = "STUB: not implemented"
	return false
}

// WithSessionEventLimit sets the limit of events in a session.
func WithSessionEventLimit(limit int) ServiceOpt {
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

// WithExtraOptions sets the extra options for the postgres session service.
// this option mainly used for the customized postgres client builder, it will be passed to the builder.
func WithExtraOptions(extraOptions ...any) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSessionTTL sets the TTL for session state and event list.
// If not set, session will not expire, set 0 will not expire.
func WithSessionTTL(ttl time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithAppStateTTL sets the TTL for app state.
// If not set, app state will not expire.
func WithAppStateTTL(ttl time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithUserStateTTL sets the TTL for user state.
// If not set, user state will not expire.
func WithUserStateTTL(ttl time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithEnableAsyncPersist enables async persistence for session state and event list.
// if not set, default is false.
func WithEnableAsyncPersist(enable bool) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithAsyncPersisterNum sets the number of workers for async persistence.
func WithAsyncPersisterNum(num int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSummarizer injects a summarizer for LLM-based summaries.
func WithSummarizer(s summary.SessionSummarizer) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithAsyncSummaryNum sets the number of workers for async summary processing.
func WithAsyncSummaryNum(num int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSummaryQueueSize sets the size of the summary job queue.
func WithSummaryQueueSize(size int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSummaryJobTimeout sets the timeout for processing a single summary job.
// If not set, a sensible default will be applied.
func WithSummaryJobTimeout(timeout time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSummaryFilterAllowlist restricts which non-empty filterKeys may trigger
// branch summaries. Keys use the same exact format as event filter keys.
func WithSummaryFilterAllowlist(filterKeys ...string) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithCascadeFullSessionSummary controls whether an allowed branch summary also
// refreshes the full-session summary keyed by SummaryFilterKeyAllContents.
func WithCascadeFullSessionSummary(enable bool) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSoftDelete enables or disables soft delete.
// When enabled (default), DELETE operations set deleted_at timestamp instead of removing records.
// When disabled, DELETE operations permanently remove records from database.
func WithSoftDelete(enable bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithCleanupInterval sets the interval for automatic cleanup of expired data.
// If set to 0, automatic cleanup will be determined based on TTL configuration.
// Default cleanup interval is 5 minutes if any TTL is configured.
func WithCleanupInterval(interval time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSkipDBInit skips database initialization (table and index creation).
// Useful when:
// - User doesn't have DDL permissions
// - Tables are managed by migration tools
// - Running in production environment where schema is pre-created
func WithSkipDBInit(skip bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithTablePrefix sets a prefix for all table names.
// For example, with prefix "trpc", tables will be named:
// - trpc_session_states
// - trpc_session_events
// - etc.
//
// Note: An underscore will be automatically added if not present.
// "trpc" and "trpc_" both result in "trpc_" prefix.
//
// Security: Uses internal/session/sqldb.ValidateTablePrefix to prevent SQL injection.
func WithTablePrefix(prefix string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// Use internal/session/sqldb validation

// Automatically add underscore if not present

// WithSchema sets the PostgreSQL schema name where tables will be created.
// If not set, tables will be created in the default schema (typically "public").
// For example, with schema "my_schema", tables will be qualified as:
// - my_schema.session_states
// - my_schema.session_events
// - etc.
//
// Note: The schema must already exist in the database before using this option.
// Security: Uses internal/session/sqldb.ValidateTableName to prevent SQL injection.
func WithSchema(schema string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// Use internal/session/sqldb validation

// WithAppendEventHook adds AppendEvent hooks.
func WithAppendEventHook(hooks ...session.AppendEventHook) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithGetSessionHook adds GetSession hooks.
func WithGetSessionHook(hooks ...session.GetSessionHook) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}
