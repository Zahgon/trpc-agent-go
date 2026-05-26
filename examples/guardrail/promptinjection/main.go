//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates the guardrail prompt injection capability with a separate reviewer runner.
package main

import (
	"context"
	"flag"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	appName        = "guardrail-promptinjection-demo"
	mainAgentName  = "promptinjection-assistant"
	reviewerAgent  = "guardrail-promptinjection-reviewer"
	reviewerRunner = "guardrail-promptinjection-reviewer-runner"
	cmdExit        = "/exit"
	cmdHelp        = "/help"
)

var (
	modelName = flag.String("model", "gpt-5.4", "Name of the model to use")
	streaming = flag.Bool("streaming", false, "Enable streaming responses")
)

func main() {
	flag.Parse()
	app := &demoApp{
		modelName: *modelName,
		streaming: *streaming,
	}
	if err := app.run(context.Background()); err != nil {
		log.Fatalf("guardrail prompt injection demo failed: %v", err)
	}
}

type demoApp struct {
	modelName      string
	streaming      bool
	mainRunner     runner.Runner
	reviewerRunner runner.Runner
	userID         string
	sessionID      string
}

func (a *demoApp) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *demoApp) setup() error { _ = "STUB: not implemented"; return nil }

func (a *demoApp) loop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *demoApp) runTurn(ctx context.Context, text string) error {
	_ = "STUB: not implemented"
	return nil
}
