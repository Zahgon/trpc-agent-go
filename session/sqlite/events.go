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
	"context"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// AppendEvent appends an event to a session.
func (s *Service) AppendEvent(
	ctx context.Context,
	sess *session.Session,
	e *event.Event,
	opts ...session.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) appendEventInternal(
	ctx context.Context,
	sess *session.Session,
	e *event.Event,
	key session.Key,
	opts ...session.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) enqueueEventPersist(
	ctx context.Context,
	sess *session.Session,
	key session.Key,
	e *event.Event,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// AppendTrackEvent appends a track event to a session.
func (s *Service) AppendTrackEvent(
	ctx context.Context,
	sess *session.Session,
	trackEvent *session.TrackEvent,
	opts ...session.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) enqueueTrackPersist(
	ctx context.Context,
	sess *session.Session,
	key session.Key,
	e *session.TrackEvent,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) startAsyncPersistWorker() { _ = "STUB: not implemented"; return }
