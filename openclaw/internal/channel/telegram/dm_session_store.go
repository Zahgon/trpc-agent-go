//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package telegram

import (
	"context"
	"sync"
	"time"
)

const (
	dmSessionStoreVersion = 1

	dmSessionStoreFilePrefix = "dm-session-"
	dmSessionStoreFileSuffix = ".json"

	dmSessionStoreDirPerm  = 0o700
	dmSessionStoreFilePerm = 0o600

	dmSessionStoreDayLayout = "20060102"

	dmSessionSaltBytes = 8

	dmSessionTempSuffixBytes = 8
)

type dmSessionResetPolicy struct {
	Idle  time.Duration
	Daily bool
}

type dmSessionStore struct {
	path string
	now  func() time.Time

	mu    sync.Mutex
	state dmSessionStoreState
}

type dmSessionStoreState struct {
	Version int `json:"version"`

	DMs map[string]*dmSessionEntry `json:"dms,omitempty"`
}

type dmSessionEntry struct {
	ActiveSessionID string `json:"active_session_id,omitempty"`

	LastActivityUnix int64 `json:"last_activity_unix,omitempty"`
	LastResetUnix    int64 `json:"last_reset_unix,omitempty"`

	LastDailyResetDay string `json:"last_daily_reset_day,omitempty"`
}

func newDMSessionStore(path string) (*dmSessionStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dmSessionStorePath(stateDir string, bot BotInfo) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *dmSessionStore) load() error { _ = "STUB: not implemented"; return nil }

func (s *dmSessionStore) EnsureActiveSession(
	ctx context.Context,
	userID string,
	legacySessionID string,
	policy dmSessionResetPolicy,
) (sessionID string, rotated bool, err error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func (s *dmSessionStore) Rotate(
	ctx context.Context,
	userID string,
	legacySessionID string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *dmSessionStore) ForgetUser(
	ctx context.Context,
	userID string,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *dmSessionStore) persistLocked(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func randomHex(n int) string { _ = "STUB: not implemented"; return "" }
