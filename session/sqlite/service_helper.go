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
	"database/sql"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

func (s *Service) getSession(
	ctx context.Context,
	key session.Key,
	limit int,
	afterTime time.Time,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

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

func (s *Service) addEvent(
	ctx context.Context,
	key session.Key,
	evt *event.Event,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) addTrackEvent(
	ctx context.Context,
	key session.Key,
	trackEvent *session.TrackEvent,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) deleteSessionState(
	ctx context.Context,
	key session.Key,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) softDeleteSessionTx(
	ctx context.Context,
	tx *sql.Tx,
	key session.Key,
	nowNs int64,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) hardDeleteSessionTx(
	ctx context.Context,
	tx *sql.Tx,
	key session.Key,
) error {
	_ = "STUB: not implemented"
	return nil
}
