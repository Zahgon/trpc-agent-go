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

	"trpc.group/trpc-go/trpc-agent-go/session"
)

// sessionIndexEntry is the value stored in the per-user session index Hash.
// Structured as JSON to allow future metadata extensions (e.g. lastActiveAt).
type sessionIndexEntry struct {
	CreatedAt time.Time `json:"createdAt"`
}

// addSessionToUserIndex atomically creates session meta (SET NX) and registers
// the session in the per-user index Hash via Lua script.
func (c *Client) addSessionToUserIndex(ctx context.Context, key session.Key, metaJSON []byte, now time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

// removeSessionFromUserIndex removes a session entry from the per-user index Hash
// and deletes all associated data keys via Lua script.
func (c *Client) removeSessionFromUserIndex(ctx context.Context, dataKeys []string, key session.Key) error {
	_ = "STUB: not implemented"
	return nil
}

const userSessionHScanBatchSize = 100

// listSessionIDsFromUserIndex returns all session IDs stored in the per-user
// index Hash via HSCAN iteration.
func (c *Client) listSessionIDsFromUserIndex(ctx context.Context, userKey session.UserKey) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cleanupStaleUserSessionIndexEntries removes orphaned session IDs from the index Hash.
// Called by ListSessions when meta keys have expired but index entries remain.
func (c *Client) cleanupStaleUserSessionIndexEntries(ctx context.Context, userKey session.UserKey, staleIDs []string) {
	_ = "STUB: not implemented"
	return
}
