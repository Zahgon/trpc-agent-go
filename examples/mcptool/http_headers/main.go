//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates how to use HTTPBeforeRequest to dynamically set HTTP headers for MCP tool calls.
package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool/mcp"
)

// Context keys for passing dynamic data
type contextKey string

const (
	requestIDKey contextKey = "request-id"
	userIDKey    contextKey = "user-id"
	sessionIDKey contextKey = "session-id"
	timestampKey contextKey = "timestamp"
)

func main() {
	fmt.Printf("🚀 MCP HTTP Headers Example\n")
	fmt.Printf("This example shows how to dynamically set HTTP headers for MCP tool calls\n")
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Println(strings.Repeat("=", 50))

	chat := &httpHeadersChat{
		modelName: "deepseek-v4-flash",
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

type httpHeadersChat struct {
	modelName  string
	runner     runner.Runner
	userID     string
	sessionID  string
	mcpToolSet *mcp.ToolSet
}

func (c *httpHeadersChat) run() error { _ = "STUB: not implemented"; return nil }

func (c *httpHeadersChat) setup(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Generate IDs first (needed for setup context)
	return nil
}

// Create OpenAI model

// Create HTTPBeforeRequest function that extracts headers from context

// Extract request ID from context

// Extract user ID from context

// Extract session ID from context

// Extract timestamp from context

// Log the request

// Create SSE MCP toolset with HTTPBeforeRequest

// Create context with values for setup phase
// This ensures that all MCP requests (initialize, tools/list, GET SSE) have headers

// Get tools from the toolset (using context with values)
// This will trigger: initialize, initialized notification, tools/list, GET SSE
// All these requests will have the custom headers from setupCtx

// Create LLM agent with generation config

// Enable streaming

// IMPORTANT: Use WithTools instead of WithToolSets
//
// Two approaches for integrating MCP tools:
//
// Approach A (used here): WithTools(tools)
//   - Manually call toolSet.Tools(ctx) with your own context
//   - Gives you full control over the context for initialize/tools/list
//   - All POST requests (initialize, tools/list, tools/call) have dynamic headers
//   - Recommended when you need per-request authentication or tracing
//
// Approach B (alternative): WithToolSets([]tool.ToolSet{toolSet})
//   - By default, tool discovery uses context.Background(), so
//     initialize/tools/list won't see ctx values.
//   - WithRefreshToolSetsOnRun(true), tool discovery uses the
//     run context (but refreshes tools every run).
//   - tools/call will use the run context either way.
//   - Suitable when you only need static headers (e.g., API keys)
//   - Can combine with WithRequestHeader for static headers
//
// This example uses Approach A to demonstrate full dynamic header control.

// Create runner

func (c *httpHeadersChat) startChat(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *httpHeadersChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// Generate a unique request ID for this message

// Add dynamic values to context
// These will be extracted by the HTTPBeforeRequest function

// Run the agent - context will flow through to HTTPBeforeRequest

func (c *httpHeadersChat) processStreamingResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Detect tool calls

// Detect tool responses

// Process streaming content

// Helper functions for creating pointers to primitive types.
func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
