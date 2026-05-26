//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package util provides shared utility functions for Redis session implementations.
package util

import (
	"context"

	"github.com/redis/go-redis/v9"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	// ServiceMetaStorageTypeKey is the key in Session.ServiceMeta to store the data version.
	ServiceMetaStorageTypeKey = "storage_type"
	// StorageTypeHashIdx indicates the session is stored in HashIdx format.
	StorageTypeHashIdx = "hashidx"
	// StorageTypeZset indicates the session is stored in Hash format.
	StorageTypeZset = "zset"
)

// ProcessStateCmd processes a HGetAll command result into a StateMap.
func ProcessStateCmd(cmd *redis.MapStringStringCmd) (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

// ProcessEventCmd processes a ZRange (StringSlice) command result into a list of Events.
func ProcessEventCmd(
	ctx context.Context,
	cmd *redis.StringSliceCmd,
) ([]event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MergeState merges app and user state into the session.
func MergeState(appState, userState session.StateMap, sess *session.Session) *session.Session {
	_ = "STUB: not implemented"
	return nil
}

// NormalizeSessionEvents returns the first event list or nil.
func NormalizeSessionEvents(events [][]event.Event) []event.Event {
	_ = "STUB: not implemented"
	return nil
}

// AttachTrackEvents attaches track events to the session.
func AttachTrackEvents(
	sess *session.Session,
	trackEvents []map[session.Track][]session.TrackEvent,
) {
	_ = "STUB: not implemented"
	return
}

// AttachSummaries attaches summaries to the session.
func AttachSummaries(sess *session.Session, summariesCmd *redis.StringCmd) {
	_ = "STUB: not implemented"
	return
}
