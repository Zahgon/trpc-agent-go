//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates direct host command execution with an LLM
// agent.
package main

import (
	"context"
	"flag"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const appName = "hostexec-demo"

func main() {
	modelName := flag.String("model", "deepseek-v4-flash", "Model name to use")
	baseDir := flag.String("base-dir", ".", "Base directory for commands")
	flag.Parse()

	app, err := newApp(*modelName, *baseDir)
	if err != nil {
		log.Fatalf("setup failed: %v", err)
	}
	defer app.runner.Close()
	defer app.tools.Close()

	if err := app.run(context.Background()); err != nil {
		log.Fatalf("run failed: %v", err)
	}
}

type cliApp struct {
	modelName string
	baseDir   string
	tools     tool.ToolSet
	runner    runner.Runner
	userID    string
	sessionID string
}

func newApp(modelName string, baseDir string) (*cliApp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *cliApp) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *cliApp) runTurn(
	ctx context.Context,
	userText string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func printToolCalls(ev *event.Event) error { _ = "STUB: not implemented"; return nil }

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

const hostExecInstruction = `You are a careful assistant with a direct
host command tool.

Use exec_command for project-local shell work such as listing files,
running builds, running tests, or collecting command output.

Rules:
- Stay inside the configured base directory unless the user explicitly
  asks for another workdir.
- Prefer concise, non-interactive commands.
- For long-running commands, use exec_command with a positive
  yield_time_ms, then continue polling with write_stdin using empty
  chars.
- Use kill_session if a background command should stop.
- Summarize command results clearly after the tool output is available.`
