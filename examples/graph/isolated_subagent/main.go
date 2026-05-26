//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates WithSubgraphIsolatedMessages(true) working correctly
// when the subagent uses tools with the builtin planner (default LLMAgent behavior).
//
// This example shows:
//   - A parent graph that delegates to a child LLMAgent via AddAgentNode
//   - The child LLMAgent has tools and uses the default builtin planner
//   - WithSubgraphIsolatedMessages(true) isolates the child from parent's history
//     while preserving the child's own tool call history within the current invocation
//
// Run with:
//
//	go run . -question "What is 12 + 7?"
//
// Expected behavior: The agent should call the calculator tool once and return the result.
// The child agent correctly sees its own tool calls while being isolated from parent history.
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
)

var (
	modelName = flag.String("model", getEnvOrDefault("MODEL_NAME", "deepseek-v4-flash"), "LLM model name")
	question  = flag.String("question", "", "User question; leave empty for interactive mode")
	isolate   = flag.Bool("isolate", true, "Enable WithSubgraphIsolatedMessages for session isolation")
	maxIter   = flag.Int("max-iter", 3, "Max tool iterations to prevent infinite loop")
	verbose   = flag.Bool("v", false, "Verbose output")
	useReact  = flag.Bool("react", false, "Use ReActPlanner instead of builtin planner")
)

const (
	appName        = "isolated-subagent-demo"
	parentName     = "parent"
	childAgentName = "calculator_agent"
	nodePreprocess = "preprocess"
	nodeAgent      = childAgentName
	nodeCollect    = "collect"
)

func getEnvOrDefault(key, defaultVal string) string { _ = "STUB: not implemented"; return "" }

func main() {
	flag.Parse()

	fmt.Println("=" + strings.Repeat("=", 63))
	fmt.Println("Isolated Subagent Demo - WithSubgraphIsolatedMessages Example")
	fmt.Println("=" + strings.Repeat("=", 63))
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Isolate messages: %v (WithSubgraphIsolatedMessages)\n", *isolate)
	fmt.Printf("Use ReActPlanner: %v\n", *useReact)
	fmt.Printf("Max tool iterations: %d\n", *maxIter)
	fmt.Println()

	prompt := strings.TrimSpace(*question)
	if prompt == "" {
		prompt = readQuestionFromStdin()
	}

	if err := run(prompt); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func readQuestionFromStdin() string { _ = "STUB: not implemented"; return "" }

func run(prompt string) error { _ = "STUB: not implemented"; return nil }

// Build the child LLMAgent with calculator tool

// Build the parent graph that delegates to the child agent

// Create the parent GraphAgent with the child as a sub-agent

// Create runner with in-memory session

func buildChildAgent() agent.Agent {
	_ = "STUB: not implemented"
	// Create model using environment variables for configuration
	return *new(agent.Agent)
}

// Create calculator tool

// Prevent infinite loop

// Optionally use ReActPlanner

// Import and use react planner if needed
// opts = append(opts, llmagent.WithPlanner(react.New()))

func buildParentGraph() (*graph.Graph, error) { _ = "STUB: not implemented"; return nil, nil }

// Preprocess node - just passes through

// Agent node with optional isolation

// WithSubgraphIsolatedMessages(true) isolates the child agent from parent's
// session history while preserving the child's own tool call history within
// the current invocation. This allows proper ReAct loop execution.

// Collect node - gathers the result

// Wire up the graph

func preprocess(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Just pass through - the user input is already in the state

func collect(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Extract the last response

func streamEvents(eventChan <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

// Track tool calls

// Track tool results

// Stream assistant content

// Delta content (streaming)

// Full message content

// Handle errors

// calculator implements the calculator tool
func calculator(_ context.Context, args calculatorArgs) (calculatorResult, error) {
	_ = "STUB: not implemented"
	return *new(calculatorResult), nil
}

type calculatorArgs struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
}

type calculatorResult struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Result    float64 `json:"result"`
}
