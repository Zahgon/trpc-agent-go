//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates mid-turn summarization in a single run where the
// model performs multiple sequential tool calls before producing the final
// answer.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	defaultToolIterationSlack = 3
	defaultLLMCallSlack       = 6
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Model name to use")
	steps     = flag.Int("steps", 5, "Sequential tool calls required in one turn")
	query     = flag.String("query",
		"Plan and execute the task using the required step tool calls.",
		"User message for the run")
	waitSec             = flag.Int("wait-sec", 8, "Wait seconds for async summary")
	syncSummaryIntraRun = flag.Bool(
		"sync-summary-intra-run",
		false,
		"Enable synchronous summary refresh between LLM iterations in a single run",
	)
)

func main() {
	flag.Parse()
	if *steps <= 0 {
		fmt.Println("steps must be greater than 0")
		os.Exit(1)
	}
	d := &sameTurnDemo{
		modelName:           *modelName,
		steps:               *steps,
		syncSummaryIntraRun: *syncSummaryIntraRun,
	}
	if err := d.run(context.Background(), *query, time.Duration(*waitSec)*time.Second); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}
}

type sameTurnDemo struct {
	modelName           string
	steps               int
	syncSummaryIntraRun bool
	runner              runner.Runner
	sessionService      session.Service
	app                 string
	userID              string
	sessionID           string
	requestSeq          int64
}

func (d *sameTurnDemo) run(ctx context.Context, input string, wait time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *sameTurnDemo) setup() error { _ = "STUB: not implemented"; return nil }

func (d *sameTurnDemo) runSingleTurn(ctx context.Context, input string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *sameTurnDemo) waitSummary(ctx context.Context, wait time.Duration) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *sameTurnDemo) readSummary(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *sameTurnDemo) beforeModel(
	_ context.Context, args *model.BeforeModelArgs,
) (*model.BeforeModelResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isSessionSummaryMessage(msg model.Message) bool { _ = "STUB: not implemented"; return false }

func (d *sameTurnDemo) stepWorker(_ context.Context, req stepArgs) (stepResult, error) {
	_ = "STUB: not implemented"
	return *new(stepResult), nil
}

type stepArgs struct {
	Step int    `json:"step" description:"Current step number, starts at 1"`
	Task string `json:"task" description:"Task description for this step"`
}

type stepResult struct {
	Step    int    `json:"step"`
	Task    string `json:"task"`
	Status  string `json:"status"`
	Detail  string `json:"detail"`
	Checked bool   `json:"checked"`
}

func preview(s string, max int) string { _ = "STUB: not implemented"; return "" }

func sortedSetKeys(m map[string]struct{}) []string { _ = "STUB: not implemented"; return nil }

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
