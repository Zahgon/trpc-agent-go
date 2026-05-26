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

// =============================================================================
// Summary Operations
// =============================================================================

// Summary storage uses a String key containing a JSON map of all filterKey summaries:
//   Key:   hashidx:sesssum:appName:{userID}:sessionID
//   Type:  String
//   Value: JSON(map[filterKey]*session.Summary)

// CreateSummary creates or updates a summary for the session.
// Uses Lua script to atomically merge filterKey summary only if newer.
// TTL is set atomically inside the Lua script to avoid orphan keys without expiry.
// UpdatedAt is normalized to UTC before serialization so that the Lua-side
// lexicographic comparison of ISO 8601 strings equals chronological order.
func (c *Client) CreateSummary(
	ctx context.Context,
	key session.Key,
	filterKey string,
	sum *session.Summary,
	ttl time.Duration,
) error {
	_ = "STUB: not implemented"
	// Normalize to UTC so Lua string comparison of "updated_at" is correct.
	return nil
}

// GetSummary retrieves all summaries for the session.
func (c *Client) GetSummary(ctx context.Context, key session.Key) (map[string]*session.Summary, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
