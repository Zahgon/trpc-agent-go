//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package knowledge

import (
	"context"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/internal/loader"
)

const loadHeartbeatInterval = 30 * time.Second

type loadReporter struct {
	cfg         *loadConfig
	sourceNames []string
	startTime   time.Time
	buckets     []int
	totalFunc   func() int

	mu            sync.Mutex
	closeOnce     sync.Once
	stats         *loader.Stats
	lastLogged    map[string]int
	heartbeatStop chan struct{}
	heartbeatDone chan struct{}
}

func newLoadReporter(cfg *loadConfig, sourceNames []string, start time.Time, buckets []int, totalFn func() int) *loadReporter {
	_ = "STUB: not implemented"
	return nil
}

func (lr *loadReporter) enabled() bool { _ = "STUB: not implemented"; return false }

func (lr *loadReporter) RecordStat(size int) { _ = "STUB: not implemented"; return }

func (lr *loadReporter) Progress(ctx context.Context, ev LoadProgressEvent) {
	_ = "STUB: not implemented"
	return
}

func (lr *loadReporter) Error(ctx context.Context, ev LoadProgressEvent, err error) {
	_ = "STUB: not implemented"
	return
}

func (lr *loadReporter) Done(ctx context.Context) { _ = "STUB: not implemented"; return }

func (lr *loadReporter) Close() { _ = "STUB: not implemented"; return }

func (lr *loadReporter) runHeartbeat() { _ = "STUB: not implemented"; return }
