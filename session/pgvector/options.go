//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package pgvector provides a PostgreSQL session service
// with built-in pgvector-based semantic search.
package pgvector

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/knowledge/embedder"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/session/summary"
)

// Default settings.
const (
	defaultSessionEventLimit     = 1000
	defaultChanBufferSize        = 100
	defaultAsyncPersisterNum     = 10
	defaultCleanupIntervalSecond = 5 * time.Minute
	defaultEmbedTimeout          = 30 * time.Second

	defaultAsyncSummaryNum   = 3
	defaultSummaryQueueSize  = 100
	defaultSummaryJobTimeout = 60 * time.Second

	defaultHost     = "localhost"
	defaultPort     = 5432
	defaultDatabase = "trpc-agent-go-pgsession"
	defaultSSLMode  = "disable"

	defaultIndexDimension = 1536
	defaultMaxResults     = 5
	defaultHNSWM          = 16
	defaultHNSWEf         = 200
	defaultHybridRRFK     = 60
	defaultCandidateRatio = 3
)

// IndexTextBuilder customizes the searchable text stored
// for an event before embedding.
type IndexTextBuilder func(
	sess *session.Session,
	evt *event.Event,
	baseText string,
	role model.Role,
) string

// ServiceOpts holds all configuration for the pgvector
// session service.
type ServiceOpts struct {
	sessionEventLimit int

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

	sessionTTL         time.Duration
	appStateTTL        time.Duration
	userStateTTL       time.Duration
	enableAsyncPersist bool
	asyncPersisterNum  int
	softDelete         bool
	cleanupInterval    time.Duration

	// Summarizer integrates LLM summarization.
	summarizer                summary.SessionSummarizer
	asyncSummaryNum           int
	summaryQueueSize          int
	summaryJobTimeout         time.Duration
	summaryFilterAllowlist    []string
	cascadeFullSessionSummary *bool

	skipDBInit  bool
	tablePrefix string
	schema      string

	// Hooks for session operations.
	appendEventHooks []session.AppendEventHook
	getSessionHooks  []session.GetSessionHook

	// Vector index settings.
	indexDimension int
	maxResults     int
	hnswM          int
	hnswEf         int
	hybridRRFK     int
	candidateRatio int

	// Embedder generates event embeddings.
	embedder     embedder.Embedder
	embedTimeout time.Duration
	// syncIndexing forces embedding generation to happen
	// in the caller/worker path instead of a detached
	// goroutine.
	syncIndexing bool
	// indexTextBuilder customizes the stored searchable
	// text before embedding.
	indexTextBuilder IndexTextBuilder
}

// ServiceOpt is a functional option for the pgvector
// session service.
type ServiceOpt func(*ServiceOpts)

var defaultOptions = ServiceOpts{
	sessionEventLimit:  defaultSessionEventLimit,
	asyncPersisterNum:  defaultAsyncPersisterNum,
	enableAsyncPersist: false,
	asyncSummaryNum:    defaultAsyncSummaryNum,
	summaryQueueSize:   defaultSummaryQueueSize,
	summaryJobTimeout:  defaultSummaryJobTimeout,
	softDelete:         true,
	indexDimension:     defaultIndexDimension,
	maxResults:         defaultMaxResults,
	hnswM:              defaultHNSWM,
	hnswEf:             defaultHNSWEf,
	hybridRRFK:         defaultHybridRRFK,
	candidateRatio:     defaultCandidateRatio,
	embedTimeout:       defaultEmbedTimeout,
}

func (opts ServiceOpts) shouldCascadeFullSessionSummary() bool {
	_ = "STUB: not implemented"
	return false
}

// WithPostgresClientDSN sets the PostgreSQL DSN connection
// string directly (recommended).
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

// WithPostgresInstance uses a named postgres instance
// from storage.
func WithPostgresInstance(name string) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithExtraOptions sets extra options for the postgres
// client builder.
func WithExtraOptions(extra ...any) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSessionEventLimit sets the limit of events in a
// session.
func WithSessionEventLimit(limit int) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSessionTTL sets the TTL for session state and
// event list.
func WithSessionTTL(ttl time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithAppStateTTL sets the TTL for app state.
func WithAppStateTTL(ttl time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithUserStateTTL sets the TTL for user state.
func WithUserStateTTL(ttl time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithEnableAsyncPersist enables async persistence for
// session events.
func WithEnableAsyncPersist(enable bool) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithAsyncPersisterNum sets the number of workers for
// async persistence.
func WithAsyncPersisterNum(num int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSummarizer injects a summarizer for LLM-based
// summaries.
func WithSummarizer(s summary.SessionSummarizer) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithAsyncSummaryNum sets the number of workers for
// async summary processing.
func WithAsyncSummaryNum(num int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSummaryQueueSize sets the size of the summary job
// queue.
func WithSummaryQueueSize(size int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSummaryJobTimeout sets the timeout for processing
// a single summary job.
func WithSummaryJobTimeout(
	timeout time.Duration,
) ServiceOpt {
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
func WithSoftDelete(enable bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithCleanupInterval sets the interval for automatic
// cleanup of expired data.
func WithCleanupInterval(
	interval time.Duration,
) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSkipDBInit skips database initialization.
func WithSkipDBInit(skip bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithTablePrefix sets a prefix for all table names.
func WithTablePrefix(prefix string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSchema sets the PostgreSQL schema name.
func WithSchema(schema string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithAppendEventHook adds AppendEvent hooks.
func WithAppendEventHook(
	hooks ...session.AppendEventHook,
) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithGetSessionHook adds GetSession hooks.
func WithGetSessionHook(
	hooks ...session.GetSessionHook,
) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithEmbedder sets the embedder for generating event
// embeddings. Required for pgvector service
// initialization and vector search support.
func WithEmbedder(e embedder.Embedder) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithEmbedTimeout sets the timeout for embedding API calls.
// Default is 30 seconds. Increase this if you experience
// timeout errors with slow embedding APIs.
func WithEmbedTimeout(timeout time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSyncIndexing controls whether event embeddings are
// generated synchronously after persistence.
func WithSyncIndexing(sync bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithIndexTextBuilder customizes the text used for
// event embeddings.
func WithIndexTextBuilder(builder IndexTextBuilder) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithIndexDimension sets the embedding vector dimension
// (default: 1536). It must match the configured embedder
// dimension when the embedder reports one.
func WithIndexDimension(dim int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithMaxResults sets the default max results for
// SearchEvents (default: 5).
func WithMaxResults(n int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithHNSWM sets the HNSW index M parameter
// (default: 16).
func WithHNSWM(m int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithHNSWEfConstruction sets the HNSW index
// ef_construction parameter (default: 200).
func WithHNSWEfConstruction(ef int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithHybridRRFK sets the RRF constant used when
// SearchModeHybrid is enabled (default: 60).
func WithHybridRRFK(k int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithHybridCandidateRatio sets how many candidates each
// hybrid branch fetches before fusion (default: 3x).
func WithHybridCandidateRatio(ratio int) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}
