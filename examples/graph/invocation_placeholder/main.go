//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent. All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates {invocation:*} placeholders in a GraphAgent
// (StateGraph + Runner) workflow.
//
// {invocation:*} reads from invocation-scoped state (invocation.SetState),
// which lives only for the current run. This is useful for request metadata
// you do not want to store in the session.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	appName          = "graph-invocation-placeholder-demo"
	defaultModelName = "deepseek-v4-flash"
	defaultUserID    = "user"

	agentName = "invocation-placeholder-agent"
	nodeID    = "assistant"

	invKeyRequestID = "request_id"
	invKeyCase      = "case"

	cmdHelp      = "/help"
	cmdShowState = "/show-state"
	cmdClearCase = "/clear-case"
	cmdCasePref  = "/case "
	exitWord     = "exit"
)

type ctxKeyCase struct{}

type demo struct {
	modelName string

	userID    string
	sessionID string
	caseValue string

	sessionService session.Service
	runner         runner.Runner
}

func main() {
	modelName := flag.String(
		"model",
		defaultModelName,
		"Model name to use",
	)
	flag.Parse()

	fmt.Println("🔖 Graph Invocation Placeholder Demo")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Type '%s' to quit\n", exitWord)
	fmt.Println("Commands:")
	fmt.Printf("  - %s\n", cmdHelp)
	fmt.Printf("  - %s\n", cmdShowState)
	fmt.Printf("  - %s <value>\n", strings.TrimSpace(cmdCasePref))
	fmt.Printf("  - %s\n", cmdClearCase)
	fmt.Println(strings.Repeat("=", 60))

	d := &demo{
		modelName: *modelName,
		userID:    defaultUserID,
	}
	if err := d.run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func (d *demo) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *demo) initialize(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *demo) beforeAgent(
	ctx context.Context,
	args *agent.BeforeAgentArgs,
) (*agent.BeforeAgentResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *demo) loop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *demo) handleCommand(ctx context.Context, line string) bool {
	_ = "STUB: not implemented"
	return false
}

func (d *demo) printHelp() { _ = "STUB: not implemented"; return }

func (d *demo) printSessionState(ctx context.Context) { _ = "STUB: not implemented"; return }

func stream(ch <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }
