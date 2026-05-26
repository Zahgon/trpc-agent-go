//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates the guardrail tool approval capability with the hostexec tool set.
package main

import (
	"context"
	"flag"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	appName         = "guardrail-approval-demo"
	mainAgentName   = "hostexec-assistant"
	reviewerAgent   = "guardrail-approval-reviewer"
	reviewerRunner  = "guardrail-approval-reviewer-runner"
	cmdExit         = "/exit"
	cmdHelp         = "/help"
	toolExecCommand = "hostexec_exec_command"
	toolWriteStdin  = "hostexec_write_stdin"
	toolKillSession = "hostexec_kill_session"
)

var (
	modelName = flag.String("model", "gpt-5.4", "Name of the model to use")
	streaming = flag.Bool("streaming", false, "Enable streaming responses")
	baseDir   = flag.String("base-dir", ".", "Base directory for host commands")
)

func main() {
	flag.Parse()
	app := &demoApp{
		modelName: *modelName,
		streaming: *streaming,
		baseDir:   *baseDir,
	}
	if err := app.run(context.Background()); err != nil {
		log.Fatalf("guardrail approval demo failed: %v", err)
	}
}

type demoApp struct {
	modelName      string
	streaming      bool
	baseDir        string
	toolSet        tool.ToolSet
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
