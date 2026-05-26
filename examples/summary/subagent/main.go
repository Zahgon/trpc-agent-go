//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates session summarization behaviour when
// the primary agent delegates to a sub-agent via agenttool.
//
// It uses a low token threshold so that summarisation triggers
// quickly, making it easy to observe which events are included in
// the threshold check and when the summary fires.
package main

import (
	"context"
	"fmt"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	appName        = "summary-subagent-demo"
	parentAgent    = "parent-agent"
	childAgentName = "math-specialist"
)

func main() {
	modelName := os.Getenv("MODEL_NAME")
	if modelName == "" {
		modelName = "deepseek-v3.2"
	}
	chat := &demo{modelName: modelName}
	if err := chat.run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

type demo struct {
	modelName      string
	runner         runner.Runner
	sessionService session.Service
	userID         string
	sessionID      string
}

func (d *demo) run() error { _ = "STUB: not implemented"; return nil }

func (d *demo) setup() error { _ = "STUB: not implemented"; return nil }

// Low thresholds to make summary trigger quickly.

// Child agent with a calculator tool.

// Wrap the child agent as a tool for the parent.

// Parent agent.

// Register a BeforeModel callback to show when the summary
// has been injected.

func (d *demo) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *demo) chat(ctx context.Context, msg string) error { _ = "STUB: not implemented"; return nil }

func (d *demo) showSummaries(ctx context.Context) { _ = "STUB: not implemented"; return }

func (d *demo) dumpEvents(ctx context.Context) { _ = "STUB: not implemented"; return }

func extractPreview(e event.Event) string { _ = "STUB: not implemented"; return "" }

func beforeModel(
	_ context.Context, args *model.BeforeModelArgs,
) (*model.BeforeModelResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// calculate performs basic arithmetic.
func calculate(
	_ context.Context, args calcArgs,
) (calcResult, error) {
	_ = "STUB: not implemented"
	return *new(calcResult), nil
}

type calcArgs struct {
	Op string  `json:"operation" jsonschema:"description=add subtract multiply divide"`
	A  float64 `json:"a" jsonschema:"description=First number"`
	B  float64 `json:"b" jsonschema:"description=Second number"`
}

type calcResult struct {
	Op     string  `json:"operation"`
	A      float64 `json:"a"`
	B      float64 `json:"b"`
	Result float64 `json:"result"`
}

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }
