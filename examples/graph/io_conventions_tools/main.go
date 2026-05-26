//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
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
	"os"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

var (
	modelName = flag.String("model", os.Getenv("MODEL_NAME"), "OpenAI‑compatible model name (e.g., deepseek-v4-flash)")
	baseURL   = flag.String("base-url", os.Getenv("OPENAI_BASE_URL"), "OpenAI‑compatible base URL")
	apiKey    = flag.String("api-key", os.Getenv("OPENAI_API_KEY"), "API key")
	verbose   = flag.Bool("v", false, "Verbose: print model/tool metadata")
)

func main() {
	flag.Parse()
	if *modelName == "" {
		*modelName = "deepseek-v4-flash"
	}
	fmt.Printf("🧩 IO Conventions — Tools Node\nModel: %s\n", *modelName)
	fmt.Println(strings.Repeat("=", 56))
	if *apiKey == "" && os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Println("💡 Hint: provide -api-key/-base-url or set OPENAI_API_KEY/OPENAI_BASE_URL")
	}
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

// Optionally pass ToolCallbacks/ModelCallbacks via GraphAgent if needed.

// tools var kept to signal declared tools; graph nodes hold actual refs.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// buildGraph: parse_input -> llm_decider -> (tools or assistant) -> capture_tool (if tools) -> collect
func buildGraph() (*graph.Graph, map[string]tool.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Define a callable tool that simulates meeting scheduling.

// Allow runtime override from parsed_time if the LLM omitted it

// LLM decider has tool declaration so it may emit tool_calls

// Tools node executes tool calls when present

// After tools, capture the JSON response into state["meeting"]

// Sub‑agent fallback when no tool call

// Collector prints final payload

// Wiring with conditional tools edge: llm_decider -> tools OR assistant

// parseInput writes parsed_time and an intent‑shaping one_shot system message.
func parseInput(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// captureTool finds the latest tool.response in messages and stores structured JSON in state["meeting"].
func captureTool(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// collect prints final payload (meeting if present + last_response + node_responses + parsed_time)
func collect(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// buildSubAgent: simple assistant fallback
func buildSubAgent(name, baseURL, apiKey string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func stream(ch <-chan *event.Event, verbose bool) error { _ = "STUB: not implemented"; return nil }

// parseTime: tiny helper for demo
func parseTime(s string) string { _ = "STUB: not implemented"; return "" }

func printSamples() { _ = "STUB: not implemented"; return }

func printHelp() { _ = "STUB: not implemented"; return }

func runDemo(r runner.Runner, user, session string) { _ = "STUB: not implemented"; return }
