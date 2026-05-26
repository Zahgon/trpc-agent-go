//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package telegram

import (
	"context"
	"sync"
	"time"
)

const sentFileRecordTTL = 5 * time.Minute

type sentFileKey struct {
	RequestID string
	ChatID    int64
	ThreadID  int
}

type sentFileTracker struct {
	mu     sync.Mutex
	byPath map[sentFileKey]map[string]time.Time
}

func newSentFileTracker() *sentFileTracker { _ = "STUB: not implemented"; return nil }

func (t *sentFileTracker) Record(
	requestID string,
	chatID int64,
	threadID int,
	paths ...string,
) {
	_ = "STUB: not implemented"
	return
}

func (t *sentFileTracker) Consume(
	requestID string,
	chatID int64,
	threadID int,
) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (t *sentFileTracker) pruneLocked(now time.Time) { _ = "STUB: not implemented"; return }

func currentRequestIDFromContext(ctx context.Context) string { _ = "STUB: not implemented"; return "" }
