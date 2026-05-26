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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/session"
)

// UpdateUserState updates user state.
func (s *Service) UpdateUserState(
	ctx context.Context,
	userKey session.UserKey,
	state session.StateMap,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) upsertUserState(
	ctx context.Context,
	appName string,
	userID string,
	key string,
	value []byte,
	now time.Time,
	expiresAt *int64,
) error {
	_ = "STUB: not implemented"
	return nil
}

// ListUserStates lists user states.
func (s *Service) ListUserStates(
	ctx context.Context,
	userKey session.UserKey,
) (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

// DeleteUserState deletes a user state key.
func (s *Service) DeleteUserState(
	ctx context.Context,
	userKey session.UserKey,
	key string,
) error {
	_ = "STUB: not implemented"
	return nil
}
