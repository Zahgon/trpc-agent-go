//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package session

import (
	"context"
	"encoding/json"
	"time"
)

const (
	// tracksStateKey stores the track index on the session state.
	tracksStateKey = "tracks"
)

// TrackService provides the interface for appending track events to a session.
type TrackService interface {
	AppendTrackEvent(ctx context.Context, sess *Session, event *TrackEvent, opts ...Option) error
}

// Track represents a logical track that stores track events.
type Track string

// TrackEvent represents a track event.
type TrackEvent struct {
	Track     Track           `json:"track"`
	Payload   json.RawMessage `json:"payload"`
	Timestamp time.Time       `json:"timestamp"`
}

// TrackEvents bundles the track events.
type TrackEvents struct {
	Track  Track        `json:"track"`
	Events []TrackEvent `json:"events"`
}

// TracksFromState returns the tracks stored in the session state.
func TracksFromState(state StateMap) ([]Track, error) { _ = "STUB: not implemented"; return nil, nil }

// ensureTrackExists ensures the track exists in the session state.
func ensureTrackExists(state StateMap, track Track) error { _ = "STUB: not implemented"; return nil }
