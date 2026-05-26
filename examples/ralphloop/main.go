//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	modelName = flag.String(
		"model",
		"deepseek-v4-flash",
		"Name of the model to use",
	)
	streaming = flag.Bool(
		"streaming",
		true,
		"Enable streaming mode for responses",
	)
	variant = flag.String(
		"variant",
		"deepseek",
		"Model variant: openai, deepseek, qwen, hunyuan",
	)
	completionPromise = flag.String(
		"completion-promise",
		"DONE",
		"Stop when assistant outputs <promise>...</promise>",
	)
	maxIterations = flag.Int(
		"max-iterations",
		10,
		"Max RalphLoop iterations per message",
	)
	maxLLMCalls = flag.Int(
		"max-llm-calls",
		0,
		"Max model calls per message (0 = auto)",
	)
)

const (
	appName   = "ralphloop-demo"
	agentName = "ralphloop-assistant"

	commandExit = "/exit"

	promiseTagOpen  = "<promise>"
	promiseTagClose = "</promise>"

	defaultMaxTokens   = 2000
	defaultTemperature = 0.7
	defaultExtraCalls  = 2
)

const agentInstruction = "Be helpful and concise."

func main() {
	flag.Parse()

	fmt.Printf("RalphLoop demo: interactive task loop\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Variant: %s\n", *variant)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Printf(
		"Stop token: %s%s%s\n",
		promiseTagOpen,
		*completionPromise,
		promiseTagClose,
	)
	fmt.Printf("Type %q to exit\n", commandExit)
	fmt.Println(strings.Repeat("=", 50))

	chat := &taskChat{
		modelName: *modelName,
		streaming: *streaming,
		variant:   *variant,

		completionPromise: *completionPromise,
		maxIterations:     *maxIterations,
		maxLLMCalls:       *maxLLMCalls,
	}

	if err := chat.run(context.Background()); err != nil {
		log.Fatalf("RalphLoop demo failed: %v", err)
	}
}

type taskChat struct {
	modelName string
	streaming bool
	variant   string

	completionPromise string
	maxIterations     int
	maxLLMCalls       int

	runner    runner.Runner
	userID    string
	sessionID string
}

func (c *taskChat) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *taskChat) setup() error { _ = "STUB: not implemented"; return nil }

func (c *taskChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *taskChat) processMessage(
	ctx context.Context,
	userMessage string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *taskChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func extractContent(choice model.Choice, streaming bool) string {
	_ = "STUB: not implemented"
	return ""
}

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
