//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package pgvector

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/session"
)

// CreateSessionSummary creates or updates a session
// summary. It delegates to the configured summarizer
// and persists the result.
func (s *Service) CreateSessionSummary(
	ctx context.Context,
	sess *session.Session,
	filterKey string,
	force bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// EnqueueSummaryJob enqueues a summary job for
// asynchronous processing.
func (s *Service) EnqueueSummaryJob(
	ctx context.Context,
	sess *session.Session,
	filterKey string,
	force bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Fallback to synchronous processing with the same detached context that
// async workers use.

// GetSessionSummaryText gets the summary text for a
// session. Returns the full-session summary by default.
func (s *Service) GetSessionSummaryText(
	ctx context.Context,
	sess *session.Session,
	opts ...session.SummaryOption,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Fallback to full-session summary.
