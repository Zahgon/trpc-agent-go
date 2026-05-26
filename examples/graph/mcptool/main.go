//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates how to use MCP tools inside a graph‑based
// workflow. The example focuses on calling MCP tools like get_weather
// from a graph, and streaming the intermediate steps.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool/mcp"
)

const (
	defaultModelName = "deepseek-v4-flash"
)

var (
	modelName = flag.String("model", defaultModelName,
		"Name of the model to use")
)

func main() {
	flag.Parse()

	fmt.Printf("🚀 Graph + MCP (STDIO) example\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("This example shows how a graph can call MCP tools")
	fmt.Println("like get_weather via an MCP STDIO server.")
	fmt.Println()
	fmt.Println("💡 Hint:")
	fmt.Println("   1) cd examples/graph/mcptool")
	fmt.Println("   2) go run . -model deepseek-v4-flash")
	fmt.Println("      (the graph will spawn the STDIO MCP server automatically)")
	fmt.Println()

	chat := &mcpGraphChat{
		modelName: *modelName,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("graph+mcp chat failed: %v", err)
	}
}

// mcpGraphChat manages the graph workflow that calls MCP tools.
type mcpGraphChat struct {
	modelName string

	runner runner.Runner

	userID    string
	sessionID string

	mcpToolSet *mcp.ToolSet
}

// run sets up the graph and starts the interactive loop.
func (c *mcpGraphChat) run() error { _ = "STUB: not implemented"; return nil }

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Ensure MCP toolset is closed when finished.

// setup builds the graph, wraps it into a GraphAgent and Runner, and
// initializes the MCP toolset.
func (c *mcpGraphChat) setup(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Hold the toolset for cleanup.

// createMCPGraph creates a simple graph that lets an LLM decide when to call
// MCP tools, executes those tools via a Tools node, and then formats the final
// answer for the user.
func (c *mcpGraphChat) createMCPGraph(ctx context.Context) (*graph.Graph, *mcp.ToolSet, error) {
	_ = "STUB: not implemented"
	// Use the standard messages state schema so that we can rely on
	// graph.StateKeyMessages/graph.StateKeyLastResponse, etc.
	return nil, nil, nil
}

// Create model instance.

// Create MCP toolset: start the local STDIO MCP server defined in
// stdioserver/main.go via "go run".

// LLM node that can decide whether to call MCP tools and also produce the final answer.

// Tools node executes MCP tool calls emitted by the assistant node.

// Final no-op node used only to terminate the graph; the visible answer
// is the last response produced by the assistant node.

// Wiring: assistant -> tools (if tool_calls present) or directly to finish.

// After tools finish, go back to assistant so that it can read tool responses
// and, if needed, decide whether to call tools again or answer the user.

// Entry and finish points.

// startInteractiveMode starts a simple REPL where each line is processed by
// the graph. It shows LLM streaming output, MCP tool calls and tool results.
func (c *mcpGraphChat) startInteractiveMode(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// processRequest runs a single request through the graph and streams the
// intermediate events.
func (c *mcpGraphChat) processRequest(ctx context.Context, content string) error {
	_ = "STUB: not implemented"
	return nil
}

// processStreamingResponse streams events from the graph run and prints:
// - tool calls emitted by the LLM
// - tool responses coming back from MCP tools
// - the assistant's streamed answer
func (c *mcpGraphChat) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle errors.

// Detect and display tool calls emitted by the LLM.

// Detect tool responses (messages with role=tool).

// Stream assistant tokens.
