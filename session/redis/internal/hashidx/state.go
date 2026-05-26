//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package hashidx

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// UpdateSessionState updates the session-level state directly (HashIdx).
func (c *Client) UpdateSessionState(ctx context.Context, key session.Key, state session.StateMap) error {
	_ = "STUB: not implemented"
	return nil
}

// Exists checks if session exists.
func (c *Client) Exists(ctx context.Context, key session.Key) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ExistsPipelined adds a HashIdx session existence check to the pipeline.
// Returns the IntCmd that can be evaluated after pipeline execution.
func (c *Client) ExistsPipelined(ctx context.Context, pipe redis.Pipeliner, key session.Key) *redis.IntCmd {
	_ = "STUB: not implemented"
	return nil
}

// UpdateAppState updates app state.
func (c *Client) UpdateAppState(ctx context.Context, appName string, state session.StateMap, ttl time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteAppState deletes app state key.
func (c *Client) DeleteAppState(ctx context.Context, appName string, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// ListAppStates lists app states.
func (c *Client) ListAppStates(ctx context.Context, appName string) (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

// UpdateUserState updates user state.
func (c *Client) UpdateUserState(ctx context.Context, userKey session.UserKey, state session.StateMap, ttl time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteUserState deletes user state key.
func (c *Client) DeleteUserState(ctx context.Context, userKey session.UserKey, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// ListUserStates lists user states.
func (c *Client) ListUserStates(ctx context.Context, userKey session.UserKey) (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

// RefreshAppStateTTL refreshes the TTL for app state key.
func (c *Client) RefreshAppStateTTL(ctx context.Context, appName string) error {
	_ = "STUB: not implemented"
	return nil
}

// RefreshUserStateTTL refreshes the TTL for user state key.
func (c *Client) RefreshUserStateTTL(ctx context.Context, userKey session.UserKey) error {
	_ = "STUB: not implemented"
	return nil
}
