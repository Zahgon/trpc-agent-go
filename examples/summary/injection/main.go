//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates session summary injection modes.
//
// It runs a short scripted conversation, forces a summary, then performs
// the same follow-up turn in each injection mode (system vs user) on
// separate sessions to cleanly isolate the comparison.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Model name")
)

func main() {
	flag.Parse()

	d := &injectionDemo{modelName: *modelName}
	if err := d.run(context.Background()); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}
}

type injectionDemo struct {
	modelName string
	reqSeq    int64

	sessionService session.Service
	app            string
	userID         string
}

func (d *injectionDemo) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Phase 1: build conversation history and generate summary.

// Phase 2: system injection mode on the base session.

// Phase 3: user injection mode on a fresh session that shares the same
// conversation history. We replay the base turns and summary so the two
// phases start from identical state.

// Replay the same base turns to build identical history.

// Force summary on the new session too.

func (d *injectionDemo) newRunner(
	llm model.Model,
	mode llmagent.SessionSummaryInjectionMode,
) runner.Runner {
	_ = "STUB: not implemented"
	return *new(runner.Runner)
}

func (d *injectionDemo) runTurn(
	ctx context.Context,
	r runner.Runner,
	sessionID string,
	input string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// runTurnQuiet replays a turn without printing to keep output focused.
func (d *injectionDemo) runTurnQuiet(
	ctx context.Context,
	r runner.Runner,
	sessionID string,
	input string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *injectionDemo) doRun(
	ctx context.Context,
	r runner.Runner,
	sessionID string,
	input string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *injectionDemo) beforeModel(
	_ context.Context,
	args *model.BeforeModelArgs,
) (*model.BeforeModelResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *injectionDemo) fetchSession(ctx context.Context, sessionID string) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *injectionDemo) readSummary(sess *session.Session) string {
	_ = "STUB: not implemented"
	return ""
}

func isSummaryContent(msg model.Message) bool { _ = "STUB: not implemented"; return false }

func preview(s string, max int) string { _ = "STUB: not implemented"; return "" }

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }
