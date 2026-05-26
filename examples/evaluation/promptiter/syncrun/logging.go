//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"
	"log"
	"sync/atomic"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const maxLoggedPayload = 6000

type loggingRunner struct {
	name    string
	inner   runner.Runner
	logger  *log.Logger
	enabled bool
	nextID  atomic.Uint64
}

type loggedRunnerOutput struct {
	eventCount        int
	finalContent      string
	structuredPayload any
	err               error
}

func newLoggingRunner(
	name string,
	inner runner.Runner,
	logger *log.Logger,
	enabled bool,
) runner.Runner {
	_ = "STUB: not implemented"
	return *new(runner.Runner)
}

func (r *loggingRunner) Run(
	ctx context.Context,
	userID string,
	sessionID string,
	message model.Message,
	runOpts ...agent.RunOption,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *loggingRunner) Close() error { _ = "STUB: not implemented"; return nil }

func (o *loggedRunnerOutput) observe(evt *event.Event) { _ = "STUB: not implemented"; return }

func marshalLogValue(value any) string { _ = "STUB: not implemented"; return "" }

func truncateForLog(value string) string { _ = "STUB: not implemented"; return "" }
