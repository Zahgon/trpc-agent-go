//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package redis

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/session"
)

// CreateSessionSummary generates a summary for the session (async-ready).
// It performs per-filterKey delta summarization; when filterKey=="", it means full-session summary.
// Strategy: Summary storage version follows session storage version.
func (s *Service) CreateSessionSummary(ctx context.Context, sess *session.Session, filterKey string, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Persist to Redis.

// Fast path: use version tag from session

// Slow path: check which storage has the session

// GetSessionSummaryText returns the latest summary text from the session state if present.
// When no options are provided, returns the full-session summary (SummaryFilterKeyAllContents).
// Use session.WithSummaryFilterKey to specify a different filter key.
// Strategy: Summary storage version follows session storage version.
func (s *Service) GetSessionSummaryText(ctx context.Context, sess *session.Session, opts ...session.SummaryOption) (string, bool) {
	_ = "STUB: not implemented"
	// Check session validity.
	return "", false
}

// Try in-memory summaries first.

// Fast path: use version tag from session to avoid checkSessionExists round-trip.

// Slow path: no version tag, check which storage has the session.

func (s *Service) getSummaryFromHashIdx(ctx context.Context, key session.Key, filterKey string, createdAt time.Time) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (s *Service) getSummaryFromZSet(ctx context.Context, key session.Key, filterKey string, createdAt time.Time) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// EnqueueSummaryJob enqueues a summary job for asynchronous processing.
func (s *Service) EnqueueSummaryJob(ctx context.Context, sess *session.Session, filterKey string, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Fallback to synchronous processing with the same detached context that
// async workers use.
