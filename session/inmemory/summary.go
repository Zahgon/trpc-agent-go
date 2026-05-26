//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package inmemory

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/session"
)

// CreateSessionSummary generates a summary for the session and stores it on the session object.
// This implementation preserves original events and updates session.Summaries only.
func (s *SessionService) CreateSessionSummary(ctx context.Context, sess *session.Session, filterKey string, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Persist to in-memory store under lock.

// Get app to write summary. Session must exist in storage.

// writeSummaryUnderLock writes a summary for a filterKey under app lock and refreshes TTL.
// When filterKey is "", it represents the full-session summary.
func (s *SessionService) writeSummaryUnderLock(app *appSessions, key session.Key, filterKey string, sum *session.Summary) error {
	_ = "STUB: not implemented"
	return nil
}

// Acquire write lock to protect Summaries access.

// Copy the summary to preserve UpdatedAt calculated by isummary.SummarizeSession.

// GetSessionSummaryText returns previously stored summary from session summaries if present.
// When no options are provided, returns the full-session summary (SummaryFilterKeyAllContents).
// Use session.WithSummaryFilterKey to specify a different filter key.
func (s *SessionService) GetSessionSummaryText(ctx context.Context, sess *session.Session, opts ...session.SummaryOption) (string, bool) {
	_ = "STUB: not implemented"
	// Check session validity.
	return "", false
}

// inmemory only needs in-memory summaries.

// EnqueueSummaryJob enqueues a summary job for asynchronous processing.
func (s *SessionService) EnqueueSummaryJob(ctx context.Context, sess *session.Session, filterKey string, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Fallback to synchronous processing with the same detached context that
// async workers use.
