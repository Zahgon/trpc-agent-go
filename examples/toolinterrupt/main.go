//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates interrupting tool execution so an external
// caller can run the tool and send the tool result back to the agent.
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
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

var (
	modelName = flag.String(
		"model",
		defaultModelName,
		"Name of the model to use",
	)
	streaming = flag.Bool(
		"streaming",
		true,
		"Enable streaming mode for responses",
	)
)

const (
	defaultModelName = "deepseek-v4-flash"

	appName   = "tool-interrupt-demo"
	agentName = "tool-interrupt-agent"
	userID    = "demo-user"

	externalToolName  = "external_search"
	externalToolDesc  = "Search an external system for information."
	externalQueryArg  = "query"
	externalQueryDesc = "Search query."

	jsonSchemaTypeObject = "object"
	jsonSchemaTypeString = "string"

	maxTokens   = 1500
	temperature = 0.2

	maxToolLoops = 8
)

const agentInstruction = `You are a helpful assistant.

For every user question:
1) Call external_search with {"query": "<the user question>"}.
2) Wait for the tool result.
3) Answer using ONLY the tool result content.

Do not answer before you receive the tool result.`

func main() {
	flag.Parse()

	fmt.Printf("🚀 Tool Interrupt Demo (External Tool Execution)\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Println(strings.Repeat("=", 60))

	d := &toolInterruptDemo{
		modelName: *modelName,
		streaming: *streaming,
	}

	if err := d.run(); err != nil {
		log.Fatalf("demo failed: %v", err)
	}
}

type toolInterruptDemo struct {
	modelName string
	streaming bool

	runner    runner.Runner
	sessionID string

	externalTools []tool.Tool
}

func (d *toolInterruptDemo) run() error { _ = "STUB: not implemented"; return nil }

func (d *toolInterruptDemo) setup(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func printHelp() { _ = "STUB: not implemented"; return }

func (d *toolInterruptDemo) startChat(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *toolInterruptDemo) processTurn(
	ctx context.Context,
	userInput string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *toolInterruptDemo) runOnce(
	ctx context.Context,
	message model.Message,
) ([]model.ToolCall, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *toolInterruptDemo) printAndCollect(
	eventChan <-chan *event.Event,
) ([]model.ToolCall, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *toolInterruptDemo) executeExternally(
	tc model.ToolCall,
) (model.Message, error) {
	_ = "STUB: not implemented"
	return *new(model.Message), nil
}

type externalSearchInput struct {
	Query string `json:"query"`
}

type externalSearchOutput struct {
	Query   string   `json:"query"`
	Results []string `json:"results"`
}

func externalSearchTool(
	_ context.Context,
	in externalSearchInput,
) (externalSearchOutput, error) {
	_ = "STUB: not implemented"
	return *new(externalSearchOutput), nil
}

func runExternalSearch(args []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

type declarationOnlyTool struct {
	declaration *tool.Declaration
}

func (t *declarationOnlyTool) Declaration() *tool.Declaration {
	_ = "STUB: not implemented"
	return nil
}

func newExternalSearchDeclaration() tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
