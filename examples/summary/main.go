//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates session summarization with LLM.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

var (
	modelName      = flag.String("model", "deepseek-v4-flash", "Model name to use for LLM summarization and chat")
	streaming      = flag.Bool("streaming", true, "Enable streaming mode for responses")
	flagEvents     = flag.Int("events", 1, "Event count threshold to trigger summarization")
	flagTokens     = flag.Int("tokens", 0, "Token-count threshold to trigger summarization (0=disabled)")
	flagTimeSec    = flag.Int("time-sec", 0, "Time threshold in seconds to trigger summarization (0=disabled)")
	flagMaxWords   = flag.Int("max-words", 0, "Max summary words (0=unlimited)")
	flagSkipRecent = flag.Int("skip-recent", 0, "Number of recent events to skip during summarization (0=skip none)")
	flagAddSum     = flag.Bool("add-summary", true, "Prepend latest branch summary as system message for LLM input")
	flagMaxHist    = flag.Int("max-history", 0, "Max history messages when add-summary=false (0=unlimited)")
)

func main() {
	flag.Parse()

	chat := &summaryChat{
		modelName: *modelName,
	}
	if err := chat.run(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}
}

// summaryChat manages the conversation and summarization demo.
type summaryChat struct {
	modelName      string
	runner         runner.Runner
	sessionService session.Service
	app            string
	userID         string
	sessionID      string
}

func (c *summaryChat) run() error { _ = "STUB: not implemented"; return nil }

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// setup constructs the model, summarizer manager, session service, and runner.
func (c *summaryChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Model used for both chat and summarization.
	return nil
}

// Summarizer with customizable prompt.
// You can customize the summary prompt using WithPrompt().
// Available placeholders:
//   - {conversation_text}: The conversation content to be summarized
//   - {max_summary_words}: The maximum word count for the summary (only included when max-words > 0);
//     when max-words is enabled, include it in either WithPrompt() or WithSystemPrompt()
// You can also add a dedicated system message with WithSystemPrompt(), but it must not contain
// {conversation_text}.

// For example:
// summary.WithPrompt("Summarize this conversation focusing on key decisions: {conversation_text}"),
// summary.WithSystemPrompt("Focus on decisions and keep it within {max_summary_words} words."),

// In-memory session service with summarizer and async config.
// Async summary processing is enabled by default with the following configuration:
// - 2 async workers: handles concurrent summary generation without blocking
// - 100 queue size: buffers summary jobs during high traffic
// You can adjust these values based on your workload:
//   - Low traffic: 1-2 workers, 50-100 queue size
//   - Medium traffic: 2-4 workers, 100-200 queue size
//   - High traffic: 4-8 workers, 200-500 queue size

// 2 async workers for concurrent summary generation
// Queue size 100 to buffer summary jobs during high traffic
// Timeout for each summary job to avoid long-running LLM calls blocking workers.

// Agent and runner (non-streaming for concise output).

// IDs.

// startChat runs the interactive conversation loop.
func (c *summaryChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// processMessage handles one message: run the agent, print the answer, then create and print the summary.
func (c *summaryChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	// Commands
	return nil
}

// Re-fetch session to ensure we read the latest summaries.

// Fallback to service helper if no structured summary was found.

// Normal chat turn (no auto summary printout).

// consumeResponse reads the event stream and returns the final assistant content.
func (c *summaryChat) consumeResponse(evtCh <-chan *event.Event) string {
	_ = "STUB: not implemented"
	return ""
}

// Handle errors.

// Handle content.

// Don't break on Done - wait for all events including finalizeRun.

// extractContent extracts content from the event based on streaming mode.
func (c *summaryChat) extractContent(event *event.Event) string {
	_ = "STUB: not implemented"
	return ""
}

// Helper.
func intPtr(i int) *int {
	_ = "STUB: not implemented"

	// getSummaryFromSession returns a structured summary from the session if present.
	// It returns the first available summary from any branch.
	return nil
}

func getSummaryFromSession(sess *session.Session) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Return the first available summary from any branch.
