//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package noop provides a session service that keeps no persisted state.
package noop

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

var (
	_ session.Service      = (*Service)(nil)
	_ session.TrackService = (*Service)(nil)
)

// Service implements session.Service without storing sessions or state.
type Service struct{}

// NewService creates a new no-op session service.
func NewService() *Service {
	_ = "STUB: not implemented"

	// CreateSession creates a transient session and does not persist it.
	return nil
}

func (s *Service) CreateSession(
	ctx context.Context,
	key session.Key,
	state session.StateMap,
	opts ...session.Option,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSession always returns nil after validating the key and options.
func (s *Service) GetSession(
	ctx context.Context,
	key session.Key,
	opts ...session.Option,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListSessions always returns an empty list after validating the key and options.
func (s *Service) ListSessions(
	ctx context.Context,
	userKey session.UserKey,
	opts ...session.Option,
) ([]*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteSession validates the key and does not persist any deletion.
func (s *Service) DeleteSession(
	ctx context.Context,
	key session.Key,
	opts ...session.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateAppState validates the app name and drops the state update.
func (s *Service) UpdateAppState(ctx context.Context, appName string, state session.StateMap) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteAppState validates the app name and drops the delete request.
func (s *Service) DeleteAppState(ctx context.Context, appName string, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// ListAppStates validates the app name and returns an empty state map.
func (s *Service) ListAppStates(ctx context.Context, appName string) (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

// UpdateUserState validates the user key and drops the state update.
func (s *Service) UpdateUserState(
	ctx context.Context,
	userKey session.UserKey,
	state session.StateMap,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ListUserStates validates the user key and returns an empty state map.
func (s *Service) ListUserStates(
	ctx context.Context,
	userKey session.UserKey,
) (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

// DeleteUserState validates the user key and drops the delete request.
func (s *Service) DeleteUserState(ctx context.Context, userKey session.UserKey, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateSessionState validates the session key and drops the state update.
func (s *Service) UpdateSessionState(
	ctx context.Context,
	key session.Key,
	state session.StateMap,
) error {
	_ = "STUB: not implemented"
	return nil
}

// AppendEvent updates the transient session object and does not persist the event.
func (s *Service) AppendEvent(
	ctx context.Context,
	sess *session.Session,
	evt *event.Event,
	opts ...session.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

// AppendTrackEvent updates the transient session object and does not persist the event.
func (s *Service) AppendTrackEvent(
	ctx context.Context,
	sess *session.Session,
	trackEvent *session.TrackEvent,
	opts ...session.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateSessionSummary is a no-op.
func (s *Service) CreateSessionSummary(
	ctx context.Context,
	sess *session.Session,
	filterKey string,
	force bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// EnqueueSummaryJob is a no-op.
func (s *Service) EnqueueSummaryJob(
	ctx context.Context,
	sess *session.Session,
	filterKey string,
	force bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// GetSessionSummaryText always reports that no summary exists.
func (s *Service) GetSessionSummaryText(
	ctx context.Context,
	sess *session.Session,
	opts ...session.SummaryOption,
) (string, bool) {
	_ = "STUB: not implemented"

	// Close closes the no-op service.
	return "", false
}

func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }

func applyOptions(opts ...session.Option) *session.Options { _ = "STUB: not implemented"; return nil }
