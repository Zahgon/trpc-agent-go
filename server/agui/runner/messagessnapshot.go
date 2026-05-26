//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package runner

import (
	"context"
	"time"

	aguievents "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/adapter"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// MessagesSnapshotter provides a MessagesSnapshot event stream by replaying persisted AG-UI track events.
type MessagesSnapshotter interface {
	// MessagesSnapshot sends a MessagesSnapshot event stream by replaying persisted AG-UI track events.
	MessagesSnapshot(ctx context.Context, input *adapter.RunAgentInput) (<-chan aguievents.Event, error)
}

// MessagesSnapshot sends a MessagesSnapshot event stream by replaying persisted AG-UI track events.
func (r *runner) MessagesSnapshot(ctx context.Context,
	runAgentInput *adapter.RunAgentInput) (<-chan aguievents.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// messagesSnapshot sends a MessagesSnapshot event stream by replaying persisted AG-UI track events.
func (r *runner) messagesSnapshot(ctx context.Context, input *runInput, events chan<- aguievents.Event) {
	_ = "STUB: not implemented"
	return
}

// Emit a RUN_STARTED event to anchor the synthetic run.

// In order to fetch the history messages as much as possible, still emit the messages even if there is an error.
// Emit a MESSAGES_SNAPSHOT event to send the snapshot payload.

// Emit a RUN_FINISHED event to signal downstream consumers there is no more data.

// getMessagesSnapshotEvent loads AG-UI track events and converts them to an AG-UI MessagesSnapshotEvent.
// In order to fetch the history messages as much as possible, still return the messages even if there is an error.
func (r *runner) getMessagesSnapshotEvent(ctx context.Context,
	sessionKey session.Key) (*aguievents.MessagesSnapshotEvent, *session.TrackEvents, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func trimTrackEventsToHistoryStart(events []session.TrackEvent) ([]session.TrackEvent, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func isUserMessageTrackEvent(trackEvent session.TrackEvent) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *runner) messagesSnapshotFollow(
	ctx context.Context,
	input *runInput,
	events chan<- aguievents.Event,
	initial *session.TrackEvents,
) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) handleMessagesSnapshotFollowTick(
	ctx context.Context,
	input *runInput,
	events chan<- aguievents.Event,
	cursorTime *time.Time,
) bool {
	_ = "STUB: not implemented"
	return false
}

func trackEndsWithTerminalRunEvent(events []session.TrackEvent) bool {
	_ = "STUB: not implemented"
	return false
}

func terminalRunSignal(evt aguievents.Event) (terminal bool, errMessage string) {
	_ = "STUB: not implemented"
	return false, ""
}

func lastTrackTimestamp(trackEvents *session.TrackEvents) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
