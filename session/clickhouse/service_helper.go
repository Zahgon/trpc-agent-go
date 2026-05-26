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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// getSession retrieves a single session with its events and summaries.
func (s *Service) getSession(
	ctx context.Context,
	key session.Key,
	limit int,
	afterTime time.Time,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	// Query session state using FINAL for deduplication
	return nil, nil
}

// Query app state

// Query user state

// Batch load events for all sessions
// Pass session created_at to filter out events from previous session instances

// Query summaries

// listSessions lists all sessions for a user.
func (s *Service) listSessions(
	ctx context.Context,
	key session.UserKey,
	limit int,
	afterTime time.Time,
	listOnlyMeta bool,
	page *session.ListSessionPage,
) ([]*session.Session, error) {
	_ = "STUB: not implemented"
	// Query app state
	return nil, nil
}

// Query user state

// Query all session states for this user using FINAL

// Build session keys and created_at times for batch loading

// Batch load events for all sessions
// Pass session created_at to filter out events from previous session instances

// Batch load summaries for all sessions

// addEvent adds an event to a session.
func (s *Service) addEvent(ctx context.Context, key session.Key, evt *event.Event) error {
	_ = "STUB: not implemented"

	// Get current session state using FINAL
	return nil
}

// Insert new version of session state (ReplacingMergeTree will deduplicate)

// Insert event if it has response and is not partial
// Events do not have their own expires_at; they are filtered by session's created_at.
// Use UnixMicro to preserve microsecond precision (ClickHouse driver has precision loss issue #1545).

// deleteSessionState soft-deletes a session and its related data.
// It inserts new versions with deleted_at set, which ReplacingMergeTree will use for deduplication.
func (s *Service) deleteSessionState(ctx context.Context, key session.Key) error {
	_ = "STUB: not implemented"

	// Get current session state to preserve fields for soft delete
	return nil
}

// Session not found or already deleted

// Soft delete: INSERT new version with deleted_at set
// ReplacingMergeTree will keep the latest version (with deleted_at)

// Soft delete session events
// Use INSERT INTO ... SELECT ... for batch soft delete

// Soft delete session summaries
// Use INSERT INTO ... SELECT ... for batch soft delete

// getEventsList loads events for multiple sessions in batch.
// sessionCreatedAts contains the created_at time for each session, used to filter out
// events from previous session instances (when a session expires and is recreated with the same ID).
func (s *Service) getEventsList(
	ctx context.Context,
	sessionKeys []session.Key,
	sessionCreatedAts []time.Time,
	limit int,
	afterTime time.Time,
) ([][]event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Build query with multiple conditions
// Each condition includes session key AND event.created_at >= session.created_at

// Map to collect events by session

// Build result in same order as sessionKeys

// getSummary loads summaries for a single session.
func (s *Service) getSummary(
	ctx context.Context,
	key session.Key,
	sessionCreatedAt time.Time,
) (map[string]*session.Summary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getSummariesList loads summaries for multiple sessions of the same user in batch.
// It queries by app_name + user_id, then filters in memory by each session's createdAt.
func (s *Service) getSummariesList(
	ctx context.Context,
	sessionKeys []session.Key,
	sessionCreatedAts []time.Time,
) ([]map[string]*session.Summary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// All sessions belong to the same user (from listSessions context)

// Build sessionCreatedAt lookup map for filtering

// Query all summaries for this user, filter by session createdAt in memory

// Map to collect summaries by session

// Filter by session createdAt to avoid cross-session leakage

// Build result in same order as sessionKeys
