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

	"trpc.group/trpc-go/trpc-agent-go/session"
)

// AppendTrackEvent appends a protocol-specific track event to a session.
// Strategy: Track event storage version follows session storage version.
func (s *Service) AppendTrackEvent(
	ctx context.Context,
	sess *session.Session,
	trackEvent *session.TrackEvent,
	opts ...session.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Update in-memory session first

// Snapshot the tracks state for persistence (sess.State["tracks"] is updated by AppendTrackEvent)

// Async persist if enabled

// Sync persist - route based on session version

// enqueueTrackEvent enqueues a track event for async persistence.
func (s *Service) enqueueTrackEvent(ctx context.Context, sess *session.Session, key session.Key, trackEvent *session.TrackEvent, tracksState []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// persistTrackEvent persists track event to the appropriate storage (zset or hashidx).
func (s *Service) persistTrackEvent(ctx context.Context, ver string, key session.Key, trackEvent *session.TrackEvent, tracksState []byte) error {
	_ = "STUB: not implemented"
	// Fast path: use version tag
	return nil
}

// Slow path: no version tag, check storage.
