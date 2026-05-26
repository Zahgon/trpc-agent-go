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

// AppendTrackEvent persists a track event to HashIdx storage.
// Track events are stored in a Hash (data) + ZSet (time index) structure
// with an auto-increment sequence for event IDs.
// tracksState is the serialized tracks state value (from session.State["tracks"])
// that will be set on the session meta atomically.
// This operation is atomic via Lua script.
func (c *Client) AppendTrackEvent(ctx context.Context, key session.Key, trackEvent *session.TrackEvent, tracksState []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Encode tracksState as base64 to match Go's json.Marshal behavior for []byte.
// In sessionMeta JSON, state values (map[string][]byte) are base64-encoded strings.

// GetTrackEvents retrieves track events for a session using Hash+ZSet structure.
// Each track is loaded via a Lua script that reads from ZSet index then HMGETs from Hash.
func (c *Client) GetTrackEvents(
	ctx context.Context,
	key session.Key,
	tracks []session.Track,
	limit int,
	afterTime time.Time,
) (map[session.Track][]session.TrackEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loadTrackEventsViaLua loads track events for a single track via Lua script.
func (c *Client) loadTrackEventsViaLua(
	ctx context.Context,
	key session.Key,
	track session.Track,
	minScore, maxScore string,
	limit int,
) ([]session.TrackEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListTracksForSession returns the list of tracks from session state.
func (c *Client) ListTracksForSession(ctx context.Context, key session.Key) ([]session.Track, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
