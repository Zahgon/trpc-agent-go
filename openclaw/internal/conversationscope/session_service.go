//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package conversationscope

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

type sessionService struct {
	next session.Service
}

// WrapSessionService rewrites persisted session keys using any explicit
// per-request storage user scope carried on the context.
func WrapSessionService(next session.Service) session.Service {
	_ = "STUB: not implemented"
	return *new(session.Service)
}

func (s *sessionService) CreateSession(
	ctx context.Context,
	key session.Key,
	state session.StateMap,
	options ...session.Option,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sessionService) GetSession(
	ctx context.Context,
	key session.Key,
	options ...session.Option,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sessionService) ListSessions(
	ctx context.Context,
	userKey session.UserKey,
	options ...session.Option,
) ([]*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sessionService) DeleteSession(
	ctx context.Context,
	key session.Key,
	options ...session.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sessionService) UpdateAppState(
	ctx context.Context,
	appName string,
	state session.StateMap,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sessionService) DeleteAppState(
	ctx context.Context,
	appName string,
	key string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sessionService) ListAppStates(
	ctx context.Context,
	appName string,
) (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

func (s *sessionService) UpdateUserState(
	ctx context.Context,
	userKey session.UserKey,
	state session.StateMap,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sessionService) ListUserStates(
	ctx context.Context,
	userKey session.UserKey,
) (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

func (s *sessionService) DeleteUserState(
	ctx context.Context,
	userKey session.UserKey,
	key string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sessionService) UpdateSessionState(
	ctx context.Context,
	key session.Key,
	state session.StateMap,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sessionService) AppendEvent(
	ctx context.Context,
	sess *session.Session,
	evt *event.Event,
	options ...session.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sessionService) CreateSessionSummary(
	ctx context.Context,
	sess *session.Session,
	filterKey string,
	force bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sessionService) EnqueueSummaryJob(
	ctx context.Context,
	sess *session.Session,
	filterKey string,
	force bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sessionService) GetSessionSummaryText(
	ctx context.Context,
	sess *session.Session,
	opts ...session.SummaryOption,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (s *sessionService) Close() error { _ = "STUB: not implemented"; return nil }

func rewriteKeyForStorage(
	ctx context.Context,
	key session.Key,
) session.Key {
	_ = "STUB: not implemented"
	return *new(session.Key)
}

func rewriteSessionForStorage(
	ctx context.Context,
	sess *session.Session,
) *session.Session {
	_ = "STUB: not implemented"
	return nil
}

func rewriteSessionForUser(
	sess *session.Session,
	userID string,
) *session.Session {
	_ = "STUB: not implemented"
	return nil
}

// withStorageSessionRuntimeSync keeps the caller's canonical session runtime
// state aligned with backend-visible mutations applied to a storage-scoped
// clone. This preserves canonical UserID semantics while avoiding split-brain
// state during the current invocation.
func (s *sessionService) withStorageSessionRuntimeSync(
	ctx context.Context,
	sess *session.Session,
	apply func(*session.Session) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func syncSessionRuntimeState(
	dst *session.Session,
	src *session.Session,
) {
	_ = "STUB: not implemented"
	return
}
