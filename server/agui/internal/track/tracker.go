//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package track implements the tracker for AG-UI events in the session.
package track

import (
	"context"
	"sync"
	"time"

	aguievents "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/aggregator"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// TrackAGUI is the AG-UI track identifier.
const TrackAGUI session.Track = "agui"

// Tracker is the interface for tracking AG-UI events.
type Tracker interface {
	// AppendEvent appends an AG-UI event to the session track.
	AppendEvent(ctx context.Context, key session.Key, event aguievents.Event) error
	// GetEvents retrieves the AG-UI track events from the session.
	GetEvents(ctx context.Context, key session.Key, opts ...session.Option) (*session.TrackEvents, error)
	// Flush flushes any pending aggregated events for the given session key.
	Flush(ctx context.Context, key session.Key) error
}

// tracker is the implementation of the Tracker interface.
type tracker struct {
	sessionService    session.Service               // sessionService handles session lifecycle.
	trackService      session.TrackService          // trackService persists track events.
	mu                sync.Mutex                    // mu guards the sessionStates map.
	aggregatorFactory aggregator.Factory            // aggregatorFactory builds aggregators for new sessions.
	aggregationOption []aggregator.Option           // aggregationOption applies to newly built aggregators.
	sessionStates     map[session.Key]*sessionState // sessionStates stores the state of each session.
	flushInterval     time.Duration                 // flushInterval is the interval for flushing the session state.
}

// sessionState stores the state of a session.
type sessionState struct {
	mu         sync.Mutex            // mu guards the aggregator and the done channel.
	aggregator aggregator.Aggregator // aggregator aggregates events.
	done       chan struct{}         // done is closed when the session state is removed.
	session    *session.Session      // session caches the ensured session to avoid repeated lookups.
}

// New creates a new tracker.
func New(service session.Service, opt ...Option) (Tracker, error) {
	_ = "STUB: not implemented"
	return *new(Tracker), nil
}

// AppendEvent appends an AG-UI event to the session track.
func (t *tracker) AppendEvent(ctx context.Context, key session.Key, event aguievents.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// GetEvents retrieves the AG-UI track events from the session.
func (t *tracker) GetEvents(ctx context.Context, key session.Key, opts ...session.Option) (*session.TrackEvents, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Flush flushes any pending aggregated events for the given session key.
func (t *tracker) Flush(ctx context.Context, key session.Key) error {
	_ = "STUB: not implemented"
	return nil
}

// persistEvents ensures the session exists and appends track events to storage.
func (t *tracker) persistEvents(ctx context.Context, key session.Key, state *sessionState, events []aguievents.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// ensureSessionExists fetches the session or creates one when absent.
func (t *tracker) ensureSessionExists(ctx context.Context, key session.Key, state *sessionState) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getSessionState returns the cached session state for the key, creating one when missing.
func (t *tracker) getSessionState(ctx context.Context, key session.Key) *sessionState {
	_ = "STUB: not implemented"
	return nil
}

// deleteSessionState removes the cached session state for the session key.
func (t *tracker) deleteSessionState(key session.Key) { _ = "STUB: not implemented"; return }

// flushPeriodically flushes the session state periodically.
func (t *tracker) flushPeriodically(ctx context.Context, key session.Key, state *sessionState) {
	_ = "STUB: not implemented"
	return
}

// flush flushes the session state.
func (t *tracker) flush(ctx context.Context, key session.Key, state *sessionState) error {
	_ = "STUB: not implemented"
	return nil
}
