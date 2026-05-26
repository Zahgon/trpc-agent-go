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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// getSession retrieves a single session with its events
// and summaries.
func (s *Service) getSession(
	ctx context.Context,
	key session.Key,
	limit int,
	afterTime time.Time,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	// Query session state.
	// Use NOW() AT TIME ZONE 'localtime' to get the server's local
	// time without timezone, matching the TIMESTAMP column type.
	return nil, nil
}

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
	return nil, nil
}

// addEvent adds an event to a session.
func (s *Service) addEvent(
	ctx context.Context,
	key session.Key,
	evt *event.Event,
) error {
	_ = "STUB: not implemented"
	return nil
}

// addTrackEvent adds a track event to a session.
func (s *Service) addTrackEvent(
	ctx context.Context,
	key session.Key,
	trackEvent *session.TrackEvent,
) error {
	_ = "STUB: not implemented"
	return nil
}

// refreshSessionTTL updates the session's updated_at
// and expires_at timestamps.
func (s *Service) refreshSessionTTL(
	ctx context.Context,
	key session.Key,
) error {
	_ = "STUB: not implemented"
	return nil
}

// deleteSessionState deletes session state plus related
// events, summaries and tracks in a single transaction.
func (s *Service) deleteSessionState(
	ctx context.Context,
	key session.Key,
) error {
	_ = "STUB: not implemented"
	return nil
}

// startAsyncPersistWorker starts goroutine workers for
// asynchronous event and track event persistence.
func (s *Service) startAsyncPersistWorker() { _ = "STUB: not implemented"; return }

// getEventsList batch loads events for multiple sessions.
func (s *Service) getEventsList(
	ctx context.Context,
	sessionKeys []session.Key,
	limit int,
	afterTime time.Time,
) ([][]event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getTrackEvents batch loads track events for multiple
// sessions.
func (s *Service) getTrackEvents(
	ctx context.Context,
	sessionKeys []session.Key,
	sessionStates []*SessionState,
	limit int,
	afterTime time.Time,
) ([]map[session.Track][]session.TrackEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reverse to ascending order.

// getSummariesList batch loads summaries for multiple
// sessions.
func (s *Service) getSummariesList(
	ctx context.Context,
	sessionKeys []session.Key,
) ([]map[string]*session.Summary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
