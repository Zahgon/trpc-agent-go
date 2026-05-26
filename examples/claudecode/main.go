//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates using the Claude Code CLI agent with the runner.
package main

import (
	"context"
	"flag"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	claudeBin    = flag.String("claude-bin", "claude", "Claude Code CLI executable path")
	outputFormat = flag.String("output-format", "json", "Transcript output format: json or stream-json")
	logDir       = flag.String("log-dir", "log", "Persist raw stdout/stderr logs under this directory")
)

func main() {
	flag.Parse()
	ag, err := newClaudeAgent(*claudeBin, *outputFormat, *logDir)
	if err != nil {
		log.Fatalf("create agent: %v", err)
	}
	r := runner.NewRunner("claudecode-cli-example", ag)
	defer r.Close()
	ctx := context.Background()
	runInteractive(ctx, r)
}

func runInteractive(ctx context.Context, r runner.Runner) { _ = "STUB: not implemented"; return }

func runOnce(ctx context.Context, r runner.Runner, userID, sessionID, prompt string) error {
	_ = "STUB: not implemented"
	return nil
}

func printToolEvents(evt *event.Event) { _ = "STUB: not implemented"; return }
