//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates Runner plugins with a small interactive chat.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	defaultModelName = "deepseek-v4-flash"
	defaultVariant   = "openai"

	appName   = "plugin-demo"
	agentName = "chat-assistant"

	cmdExit = "/exit"
	cmdHelp = "/help"

	separatorWidth = 50

	globalInstruction = "Follow security policies. Be helpful and concise."
	agentInstruction  = "Use tools for exact math when needed."
)

var (
	modelName = flag.String(
		"model",
		defaultModelName,
		"Name of the model to use",
	)
	variant = flag.String(
		"variant",
		defaultVariant,
		"OpenAI provider variant",
	)
	streaming = flag.Bool(
		"streaming",
		false,
		"Enable streaming responses",
	)
	debug = flag.Bool(
		"debug",
		false,
		"Print plugin debug lines",
	)
)

func main() {
	flag.Parse()

	fmt.Println("🔌 Runner plugin demo")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Printf("Debug: %t\n", *debug)
	fmt.Printf("Type %q to exit, %q for help\n", cmdExit, cmdHelp)
	fmt.Println(strings.Repeat("=", separatorWidth))

	app := &chatApp{
		modelName: *modelName,
		variant:   *variant,
		streaming: *streaming,
		debug:     *debug,
	}
	if err := app.run(context.Background()); err != nil {
		fmt.Printf("❌ error: %v\n", err)
		os.Exit(1)
	}
}

type chatApp struct {
	modelName string
	variant   string
	streaming bool
	debug     bool

	runner    runner.Runner
	userID    string
	sessionID string
}

func (a *chatApp) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *chatApp) setup() error { _ = "STUB: not implemented"; return nil }

func (a *chatApp) loop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *chatApp) printHelp() { _ = "STUB: not implemented"; return }

func (a *chatApp) runOnce(ctx context.Context, userText string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *chatApp) printEvents(evCh <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *chatApp) printToolCalls(evt *event.Event) { _ = "STUB: not implemented"; return }

func (a *chatApp) printToolResults(evt *event.Event) { _ = "STUB: not implemented"; return }
