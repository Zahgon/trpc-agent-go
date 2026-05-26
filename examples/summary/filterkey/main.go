//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates session summarization with custom filterKey support.
// This example shows how to use AppendEventHook to set custom filterKeys for
// categorizing conversations, enabling separate summaries per category.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

var (
	modelName   = flag.String("model", "deepseek-v4-flash", "Model name for LLM summarization")
	streaming   = flag.Bool("streaming", true, "Enable streaming mode for responses")
	maxWords    = flag.Int("max-words", 0, "Max summary words (0=unlimited)")
	debug       = flag.Bool("debug", false, "Enable debug mode to print request messages")
	allowlist   = flag.String("allowlist", "", "Comma-separated short filterKeys to summarize (empty=allow all branch filterKeys)")
	cascadeFull = flag.Bool("cascade-full", true, "Refresh the full-session summary when a branch summary runs")
)

const defaultFilterKey = "default"

func main() {
	flag.Parse()

	chat := &filterKeyChat{
		modelName:                 *modelName,
		currentFilterKey:          defaultFilterKey,
		allowedBranchFilterKeys:   parseFilterKeyList(*allowlist),
		cascadeFullSessionSummary: *cascadeFull,
	}
	if err := chat.run(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}
}

// filterKeyChat manages the conversation and filterKey summarization demo.
type filterKeyChat struct {
	modelName                 string
	runner                    runner.Runner
	sessionService            session.Service
	app                       string
	userID                    string
	sessionID                 string
	allowedBranchFilterKeys   []string
	cascadeFullSessionSummary bool

	// currentFilterKey is the active filterKey for new messages.
	currentFilterKey string
	filterKeyMu      sync.RWMutex
}

func (c *filterKeyChat) run() error { _ = "STUB: not implemented"; return nil }

// Ensure runner resources are cleaned up.

// setup constructs the model, summarizer manager, session service, and runner.
func (c *filterKeyChat) setup(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// Model used for both chat and summarization.

// Summarizer with custom filterKey support.

// In-memory session service with summarizer and AppendEventHook.
// The hook sets filterKey based on the current user-defined key.

// Create tools for the agent.

// Agent and runner with tools.

// Add debug callback if enabled.

// IDs.

// setEventFilterKey sets the filterKey based on the current user-defined key.
func (c *filterKeyChat) setEventFilterKey(evt *event.Event) { _ = "STUB: not implemented"; return }

// Use app-prefixed keys so they match the invocation's filter prefix.

// setCurrentFilterKey updates the current filterKey.
func (c *filterKeyChat) setCurrentFilterKey(key string) { _ = "STUB: not implemented"; return }

// getCurrentFilterKey returns the current filterKey.
func (c *filterKeyChat) getCurrentFilterKey() string { _ = "STUB: not implemented"; return "" }

func (c *filterKeyChat) prefixedFilterKeys(keys []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// startChat runs the interactive conversation loop.
func (c *filterKeyChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *filterKeyChat) printHelp() { _ = "STUB: not implemented"; return }

func parseFilterKeyList(raw string) []string { _ = "STUB: not implemented"; return nil }

func formatFilterKeyList(keys []string) string { _ = "STUB: not implemented"; return "" }

// processMessage handles one message: run the agent, print the answer.
func (c *filterKeyChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	// Handle commands.
	return nil
}

// Normal chat turn.

// handleKeyCommand switches the current filterKey.
func (c *filterKeyChat) handleKeyCommand(userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// handleShowCommand shows the summary for a specific filterKey.
func (c *filterKeyChat) handleShowCommand(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// handleChatTurn handles normal chat messages.
func (c *filterKeyChat) handleChatTurn(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

// displaySummary displays a summary for the given filter key.
func (c *filterKeyChat) displaySummary(sess *session.Session, filterKey, displayName string) {
	_ = "STUB: not implemented"
	return
}

// handleListSummaries prints all filterKeys and their summaries in the session.
func (c *filterKeyChat) handleListSummaries(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Extract display name from filterKey (e.g., "app/math" -> "math").

// consumeResponse reads the event stream and displays the assistant response.
func (c *filterKeyChat) consumeResponse(evtCh <-chan *event.Event, streaming bool) string {
	_ = "STUB: not implemented"
	return ""
}

// Handle errors.

// Tool call events.

// Tool response events.

// Handle content.

// Final response.

// handleToolCalls logs tool calls when the LLM requests them.
func (c *filterKeyChat) handleToolCalls(evt *event.Event, streaming bool, seen map[string]struct{}) bool {
	_ = "STUB: not implemented"
	return false
}

// handleToolResponses logs tool outputs.
func (c *filterKeyChat) handleToolResponses(evt *event.Event, seen map[string]struct{}) bool {
	_ = "STUB: not implemented"
	return false
}

// extractContent extracts content from the event based on streaming mode.
func (c *filterKeyChat) extractContent(evt *event.Event, streaming bool) string {
	_ = "STUB: not implemented"
	return ""
}
