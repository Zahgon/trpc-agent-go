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

// UpdateAppState updates the state by target scope and key.
// Note: AppState key is shared between zset and hashidx (no v2 prefix), so no version routing needed.
func (s *Service) UpdateAppState(ctx context.Context, appName string, state session.StateMap) error {
	_ = "STUB: not implemented"
	return nil
}

// AppState key is the same for both, so either client works

// ListAppStates gets the app states.
// Note: AppState key is shared between zset and hashidx (no v2 prefix).
func (s *Service) ListAppStates(ctx context.Context, appName string) (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

// AppState key is the same for both, so either client works

// DeleteAppState deletes the state by target scope and key.
// Note: AppState key is shared between zset and hashidx (no v2 prefix).
func (s *Service) DeleteAppState(ctx context.Context, appName string, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateUserState updates the state by target scope and key.
// Note: UserState keys are different between zset and hashidx (hashidx uses hashidx: prefix and different hash tag).
func (s *Service) UpdateUserState(ctx context.Context, userKey session.UserKey, state session.StateMap) error {
	_ = "STUB: not implemented"
	return nil
}

// transition mode: write to both hashidx and zset

// Legacy or None mode: write to hashidx only

// ListUserStates lists the state by target scope and key.
func (s *Service) ListUserStates(ctx context.Context, userKey session.UserKey) (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

// Try hashidx first

// Fallback to zset (if zset awareness is enabled: transition or legacy)

// DeleteUserState deletes the state by target scope and key.
func (s *Service) DeleteUserState(ctx context.Context, userKey session.UserKey, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete from hashidx

// Also delete from zset (if zset awareness is enabled: transition or legacy)

// UpdateSessionState updates the session-level state directly without appending an event.
func (s *Service) UpdateSessionState(ctx context.Context, key session.Key, state session.StateMap) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate: disallow app: and user: prefixes

// Check session existence in zset and hashidx

// zset first: if zset exists, route to zset.
