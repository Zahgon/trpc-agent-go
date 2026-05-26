//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package clickhouse

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/session"
)

// CreateSessionSummary is the internal implementation that returns the summary.
func (s *Service) CreateSessionSummary(
	ctx context.Context,
	sess *session.Session,
	filterKey string,
	force bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Persist only the updated filterKey summary with atomic set-if-newer to avoid late-write override.

// Note: expires_at is set to NULL - summaries are bound to session
// lifecycle and will be deleted when session is deleted or expires.

// INSERT new version (ReplacingMergeTree will deduplicate based on updated_at).

// EnqueueSummaryJob enqueues a summary job for asynchronous processing.
func (s *Service) EnqueueSummaryJob(ctx context.Context, sess *session.Session, filterKey string, force bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Fallback to synchronous processing with the same detached context that
// async workers use.

// GetSessionSummaryText gets the summary text for a session.
// When no options are provided, returns the full-session summary (SummaryFilterKeyAllContents).
// Use session.WithSummaryFilterKey to specify a different filter key.
func (s *Service) GetSessionSummaryText(
	ctx context.Context,
	sess *session.Session,
	opts ...session.SummaryOption,
) (string, bool) {
	_ = "STUB: not implemented"
	// Check session validity.
	return "", false
}

// Try in-memory summaries first.

// Query database with specified filterKey.

// If requested filterKey not found, try fallback to full-session summary.
