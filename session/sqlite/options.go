//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package sqlite

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/session/summary"
)

const (
	defaultSessionEventLimit = 1000

	defaultChanBufferSize = 100

	defaultAsyncPersisterNum = 1

	defaultCleanupInterval = 5 * time.Minute

	defaultDBInitTimeout = 30 * time.Second

	defaultAsyncPersistTimeout = 10 * time.Second

	defaultAsyncSummaryNum   = 3
	defaultSummaryQueueSize  = 100
	defaultSummaryJobTimeout = 60 * time.Second
)

// ServiceOpts is the options for the sqlite session service.
type ServiceOpts struct {
	sessionEventLimit int

	sessionTTL         time.Duration
	appStateTTL        time.Duration
	userStateTTL       time.Duration
	enableAsyncPersist bool
	asyncPersisterNum  int
	softDelete         bool
	cleanupInterval    time.Duration

	// summarizer integrates LLM summarization.
	summarizer                summary.SessionSummarizer
	asyncSummaryNum           int
	summaryQueueSize          int
	summaryJobTimeout         time.Duration
	summaryFilterAllowlist    []string
	cascadeFullSessionSummary *bool

	// skipDBInit skips database initialization.
	skipDBInit bool

	// tablePrefix is the prefix for all table names.
	tablePrefix string

	appendEventHooks []session.AppendEventHook
	getSessionHooks  []session.GetSessionHook
}

// ServiceOpt is the option for the sqlite session service.
type ServiceOpt func(*ServiceOpts)

var defaultOptions = ServiceOpts{
	sessionEventLimit: defaultSessionEventLimit,
	asyncPersisterNum: defaultAsyncPersisterNum,
	asyncSummaryNum:   defaultAsyncSummaryNum,
	summaryQueueSize:  defaultSummaryQueueSize,
	summaryJobTimeout: defaultSummaryJobTimeout,
	softDelete:        true,
}

func (opts ServiceOpts) shouldCascadeFullSessionSummary() bool {
	_ = "STUB: not implemented"
	return false
}

// WithSessionEventLimit sets the event limit per session.
func WithSessionEventLimit(limit int) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSessionTTL sets the TTL for session state and event list.
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

// WithEnableAsyncPersist enables async persistence.
func WithEnableAsyncPersist(enable bool) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithAsyncPersisterNum sets the number of async persister workers.
func WithAsyncPersisterNum(num int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSoftDelete enables or disables soft delete.
func WithSoftDelete(enable bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithCleanupInterval sets the cleanup interval for expired data.
func WithCleanupInterval(interval time.Duration) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithSummarizer injects a summarizer for LLM-based summaries.
func WithSummarizer(s summary.SessionSummarizer) ServiceOpt {
	_ = "STUB: not implemented"
	return *new(ServiceOpt)
}

// WithAsyncSummaryNum sets the number of async summary workers.
func WithAsyncSummaryNum(num int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSummaryQueueSize sets the size of the summary job queue.
func WithSummaryQueueSize(size int) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithSummaryJobTimeout sets the timeout for processing one summary job.
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

// WithSkipDBInit skips database initialization (DDL).
func WithSkipDBInit(skip bool) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

// WithTablePrefix sets a prefix for all table names.
//
// Security: Uses internal/session/sqldb.ValidateTablePrefix to prevent SQL
// injection.
func WithTablePrefix(prefix string) ServiceOpt { _ = "STUB: not implemented"; return *new(ServiceOpt) }

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
