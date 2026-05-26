//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package sqlite

import (
	"context"
	"database/sql"
	"time"
)

const cleanupTimeout = 5 * time.Minute

func (s *Service) startCleanupRoutine() { _ = "STUB: not implemented"; return }

func (s *Service) stopCleanupRoutine() { _ = "STUB: not implemented"; return }

func (s *Service) cleanupExpiredData(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *Service) cleanupExpiredSessions(ctx context.Context, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) softDeleteExpiredSessions(
	ctx context.Context,
	nowNs int64,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) softDeleteExpiredSessionsTx(
	ctx context.Context,
	tx *sql.Tx,
	nowNs int64,
	whereExpired string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) softDeleteExpiredBySession(
	ctx context.Context,
	tx *sql.Tx,
	table string,
	nowNs int64,
	whereExpired string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) hardDeleteExpiredSessions(
	ctx context.Context,
	nowNs int64,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) hardDeleteExpiredSessionsTx(
	ctx context.Context,
	tx *sql.Tx,
	nowNs int64,
	whereExpired string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) hardDeleteExpiredBySession(
	ctx context.Context,
	tx *sql.Tx,
	table string,
	nowNs int64,
	whereExpired string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) cleanupExpiredAppStates(
	ctx context.Context,
	now time.Time,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) cleanupExpiredUserStates(
	ctx context.Context,
	now time.Time,
) {
	_ = "STUB: not implemented"
	return
}
