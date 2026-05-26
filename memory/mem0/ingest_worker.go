//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package mem0

import (
	"context"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

type ingestJob struct {
	Ctx      context.Context
	UserKey  memory.UserKey
	Session  *session.Session
	LatestTs time.Time
	Messages []model.Message
	Options  session.IngestOptions
}

type ingestWorker struct {
	c *client

	asyncMode bool
	version   string

	jobChans []chan *ingestJob
	timeout  time.Duration

	orgID     string
	projectID string

	mu      sync.RWMutex
	wg      sync.WaitGroup
	started bool
}

const (
	ingestEventStatusPending   = "PENDING"
	ingestEventStatusRunning   = "RUNNING"
	ingestEventStatusSucceeded = "SUCCEEDED"
	ingestEventStatusFailed    = "FAILED"
	ingestEventPollInterval    = 2 * time.Second
)

func newIngestWorker(c *client, opts serviceOpts) *ingestWorker {
	_ = "STUB: not implemented"
	return nil
}

func (w *ingestWorker) start() { _ = "STUB: not implemented"; return }

func (w *ingestWorker) Stop() { _ = "STUB: not implemented"; return }

func (w *ingestWorker) tryEnqueue(ctx context.Context, job *ingestJob) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *ingestWorker) process(job *ingestJob) { _ = "STUB: not implemented"; return }

func (w *ingestWorker) ingest(
	ctx context.Context,
	userKey memory.UserKey,
	_ *session.Session,
	messages []model.Message,
	reqOpts session.IngestOptions,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *ingestWorker) awaitQueuedEvents(ctx context.Context, events createMemoryEvents) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *ingestWorker) awaitIngestEvent(ctx context.Context, eventID string) (*eventStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hashUserKey(userKey memory.UserKey) int { _ = "STUB: not implemented"; return 0 }
