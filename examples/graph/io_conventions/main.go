//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates the I/O conventions of a graph.
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
	fmt.Printf("🔧 IO Conventions Example\nModel: %s\n", *modelName)
	fmt.Println(strings.Repeat("=", 56))
	if *apiKey == "" && os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Println("💡 Hint: provide -api-key/-base-url or set OPENAI_API_KEY/OPENAI_BASE_URL")
	}

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	_ = "STUB: not implemented"
	// Build the graph
	return nil
}

// Sub‑agent: simple English assistant that can read runtime state

// GraphAgent

// Runner + memory session

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Interactive

// Run with runtime state (could carry caller metadata)

// buildGraph defines: parse_input -> llm_summary -> subagent -> collect
func buildGraph() (*graph.Graph, error) { _ = "STUB: not implemented"; return nil, nil }

// Custom keys to illustrate I/O flow

// LLM node: summarizes intent. It consumes one_shot_messages (from parse_input)

// Collect LLM's textual output into a custom key for downstream use

// Sub‑agent node

// Final collector shows how to read last_response/node_responses/custom keys

// Wiring

// parseInput extracts a simple time expression and writes both one_shot_messages and parsed_time.
func parseInput(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Provide one_shot_messages to guide the LLM node. Make it tool-friendly and constrained.

// captureLLM copies last_response into a custom key
func captureLLM(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// collect builds a final payload by reading outputs from previous nodes
func collect(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Pretty‑print once for demo

// buildSubAgent creates a simple assistant that reads parsed_time from runtime state
func buildSubAgent(name, baseURL, apiKey string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func stream(ch <-chan *event.Event, verbose bool) error { _ = "STUB: not implemented"; return nil }

// Model/Tool metadata (from graph LLM node)

// Streaming text

// parseTime: tiny helper for demo
func parseTime(s string) string { _ = "STUB: not implemented"; return "" }

func printSamples() { _ = "STUB: not implemented"; return }

func printHelp() { _ = "STUB: not implemented"; return }

func runDemo(r runner.Runner, user, session string) { _ = "STUB: not implemented"; return }
