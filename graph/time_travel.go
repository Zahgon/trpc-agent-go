//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

import (
	"context"
	"time"
)

const (
	// CheckpointMetaKeyBaseCheckpointID is stored in CheckpointMetadata.Extra
	// when a checkpoint is created via TimeTravel.EditState.
	CheckpointMetaKeyBaseCheckpointID = "base_checkpoint_id"
	// CheckpointMetaKeyUpdatedKeys is stored in CheckpointMetadata.Extra when a
	// checkpoint is created via TimeTravel.EditState.
	CheckpointMetaKeyUpdatedKeys = "updated_keys"
)

// CheckpointRef is a stable pointer to a checkpoint.
//
// It is intentionally small and "UI friendly":
//   - It can be stored outside the runtime (e.g. in DB / UI state).
//   - It can be converted to:
//   - saver config (for CheckpointSaver APIs)
//   - runtime_state (for GraphAgent / Runner resume)
type CheckpointRef struct {
	LineageID    string
	Namespace    string
	CheckpointID string
}

// Validate returns an error when the ref is incomplete.
func (r CheckpointRef) Validate() error { _ = "STUB: not implemented"; return nil }

// ToSaverConfig converts the ref into a config map for CheckpointSaver.
//
// When CheckpointID is empty, most savers interpret it as
// "latest checkpoint".
func (r CheckpointRef) ToSaverConfig() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToRuntimeState converts the ref into the runtime_state map expected by
// GraphAgent / Runner.
//
// Note: checkpoint_id is always present so callers can use an empty string to
// mean "resume from latest checkpoint" (GraphAgent uses key presence as the
// resume signal).
func (r CheckpointRef) ToRuntimeState() map[string]any { _ = "STUB: not implemented"; return nil }

// CheckpointInfo is a lightweight checkpoint header for history views.
type CheckpointInfo struct {
	Ref              CheckpointRef
	ParentCheckpoint string
	Source           string
	Step             int
	Timestamp        time.Time
}

// StateSnapshot is a checkpoint state snapshot suitable for debugging and
// HITL.
type StateSnapshot struct {
	CheckpointInfo
	State        State
	NextNodes    []string
	NextChannels []string
}

// TimeTravel provides first-class "query / edit / resume" operations built on
// top of the checkpoint system.
//
// It is additive and does not change existing checkpoint/resume semantics
// unless explicitly called by the user.
type TimeTravel struct {
	executor *Executor
	saver    CheckpointSaver
}

// TimeTravel returns a helper bound to this executor.
func (e *Executor) TimeTravel() (*TimeTravel, error) { _ = "STUB: not implemented"; return nil, nil }

// GetState returns the state snapshot at the referenced checkpoint.
//
// If ref.CheckpointID is empty, the latest checkpoint in the namespace is
// used.
func (t *TimeTravel) GetState(
	ctx context.Context,
	ref CheckpointRef,
) (*StateSnapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// History returns checkpoint headers in descending timestamp order.
func (t *TimeTravel) History(
	ctx context.Context,
	lineageID string,
	namespace string,
	limit int,
) ([]CheckpointInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EditStateOption configures EditState.
type EditStateOption func(*editStateOptions)

type editStateOptions struct {
	allowInternalKeys bool
}

// WithAllowInternalKeys allows editing internal runtime keys (keys that start
// with "__" or belong to checkpoint/runtime wiring).
//
// Most users should not enable this.
func WithAllowInternalKeys() EditStateOption {
	_ = "STUB: not implemented"
	return *new(EditStateOption)
}

// EditState creates a new checkpoint derived from base, with patched state.
//
// It writes a new checkpoint with:
//   - Source = "update"
//   - ParentCheckpointID = base checkpoint ID
//
// The new checkpoint is safe to resume from via:
//
//	agent.WithRuntimeState(newRef.ToRuntimeState())
func (t *TimeTravel) EditState(
	ctx context.Context,
	base CheckpointRef,
	patch State,
	opts ...EditStateOption,
) (CheckpointRef, error) {
	_ = "STUB: not implemented"
	return *new(CheckpointRef), nil
}

func (t *TimeTravel) snapshotFromTuple(
	tuple *CheckpointTuple,
) *StateSnapshot {
	_ = "STUB: not implemented"
	return nil
}

func (t *TimeTravel) infoFromTuple(tuple *CheckpointTuple) CheckpointInfo {
	_ = "STUB: not implemented"
	return *new(CheckpointInfo)
}

func (t *TimeTravel) coerceValue(key string, value any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func checkpointRefFromConfig(
	config map[string]any,
	fallbackID string,
) CheckpointRef {
	_ = "STUB: not implemented"
	return *new(CheckpointRef)
}

func isProtectedTimeTravelKey(key string) bool { _ = "STUB: not implemented"; return false }
