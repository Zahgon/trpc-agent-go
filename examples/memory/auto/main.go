//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates automatic memory extraction using the Runner.
// Unlike manual memory tools, auto memory extracts user information from
// conversations automatically in the background without explicit tool calls.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/runner"

	util "trpc.group/trpc-go/trpc-agent-go/examples/memory"
)

var (
	modelName = flag.String(
		"model",
		"deepseek-v4-flash",
		"Model for chat responses",
	)
	extModel = flag.String(
		"ext-model",
		"",
		"Model for memory extraction (defaults to chat model)",
	)
	streaming = flag.Bool(
		"streaming",
		true,
		"Enable streaming mode for responses",
	)
	debug = flag.Bool(
		"debug",
		false,
		"Enable debug mode to print messages sent to model",
	)
	memType = flag.String(
		"memory",
		"inmemory",
		"Memory service type: inmemory, sqlite, sqlitevec, redis, "+
			"postgres, pgvector, mysql, mysqlvec",
	)
)

func main() {
	flag.Parse()

	fmt.Println("🧠 Auto Memory Demo")
	fmt.Printf("Chat Model: %s\n", *modelName)
	extractorModel := *extModel
	if extractorModel == "" {
		extractorModel = *modelName
	}
	fmt.Printf("Extractor Model: %s\n", extractorModel)
	memoryType := util.MemoryType(*memType)
	fmt.Printf("Memory Service: %s\n", memoryType)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
	fmt.Println("💡 Auto memory mode extracts user information automatically.")
	fmt.Println("   No explicit memory tools are needed - the system learns")
	fmt.Println("   about you from natural conversation.")
	fmt.Println()

	chat := &autoMemoryChat{
		modelName:      *modelName,
		extractorModel: extractorModel,
		memoryType:     memoryType,
		appName:        appName,
		streaming:      *streaming,
		debug:          *debug,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

// autoMemoryChat manages the conversation with auto memory capabilities.
type autoMemoryChat struct {
	modelName      string
	extractorModel string
	memoryType     util.MemoryType
	appName        string
	streaming      bool
	debug          bool
	runner         runner.Runner
	memoryService  memory.Service
	userID         string
	sessionID      string
}

// run starts the interactive chat session.
func (c *autoMemoryChat) run() error { _ = "STUB: not implemented"; return nil }

const (
	appName   = "memory-chat"
	agentName = "memory-assistant"
)

// setup creates the runner with LLM agent and auto memory extraction.
func (c *autoMemoryChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create models.
	return nil
}

// Create memory extractor with optional extraction checkers.
// The extractor uses LLM to analyze conversations and extract memories.
// Checkers control when extraction should be triggered.

// Optional: configure extraction checkers.
// By default, extraction happens on every turn.
// Use checkers to control extraction frequency:
//
// Example 1: Extract when messages > 5 OR every 3 minutes (OR logic).
// extractor.WithCheckersAny(
//     extractor.CheckMessageThreshold(5),
//     extractor.CheckTimeInterval(3*time.Minute),
// ),
//
// Example 2: Extract when messages > 10 AND every 5 minutes (AND logic).
// extractor.WithChecker(extractor.CheckMessageThreshold(10)),
// extractor.WithChecker(extractor.CheckTimeInterval(5*time.Minute)),

// Create memory service with auto extraction enabled.
// When extractor is set, write tools (add/update/delete) are hidden, but
// search and clear tools remain available. Load tool is also hidden in auto mode.

// Setup identifiers.

// Create LLM agent with memory tools.
// Only search tool is available since extractor is set.

// Create model callbacks for debug mode.

// Memory preloading injects memories into the system prompt before
// each request.
// Use WithPreloadMemory(N) for adaptive preload: load all memories
// when count <= N, otherwise inject the top-N search results and
// fall back to directly loading up to N memories if search cannot
// provide usable results.
// Use WithPreloadMemory(-1) to load all memories.
// Default is 0 (disabled, use memory_search/memory_load tools instead).

// Create runner with memory service.
// The runner will automatically trigger memory extraction after responses.

// startChat runs the interactive conversation loop.
func (c *autoMemoryChat) startChat(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// showMemories displays the current memories for the user.
func (c *autoMemoryChat) showMemories(ctx context.Context) { _ = "STUB: not implemented"; return }

// processMessage handles a single message exchange.
func (c *autoMemoryChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// processResponse handles both streaming and non-streaming responses.
func (c *autoMemoryChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle tool calls (only load/search in auto mode).

// Handle tool responses.

// Handle content.

// hasToolCalls checks if the event contains tool calls.
func (c *autoMemoryChat) hasToolCalls(evt *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// hasToolResponses checks if the event contains tool responses.
func (c *autoMemoryChat) hasToolResponses(evt *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// handleToolCalls displays tool call information.
func (c *autoMemoryChat) handleToolCalls(evt *event.Event, assistantStarted bool) {
	_ = "STUB: not implemented"
	return
}

// handleToolResponses displays tool response information.
func (c *autoMemoryChat) handleToolResponses(evt *event.Event) { _ = "STUB: not implemented"; return }

// extractContent extracts content from the event based on streaming mode.
func (c *autoMemoryChat) extractContent(evt *event.Event) string {
	_ = "STUB: not implemented"
	return ""
}

// startNewSession creates a new session ID.
func (c *autoMemoryChat) startNewSession() { _ = "STUB: not implemented"; return }

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
