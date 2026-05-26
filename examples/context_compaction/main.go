//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates context compaction with a real model call.
//
// It runs two turns in the same session. The first turn asks the model to call
// a tool that returns a large log. The second turn asks about the previous tool
// result. Use -debug to inspect the exact request projected for each
// model call and verify whether historical tool results were compacted.
//
// Usage:
//
//	# Run from the examples module so this package uses examples/go.mod.
//	cd examples
//
//	# The OpenAI-compatible provider reads credentials from environment
//	# variables. MODEL_NAME is optional when -model is passed explicitly.
//	export OPENAI_API_KEY="..."
//	export MODEL_NAME="gpt-5.2"
//
//	# Debug output is enabled by default. It prints the request after session
//	# history projection and context compaction, immediately before the model
//	# adapter receives it.
//	go run ./context_compaction -model=gpt-5.2
//
//	# Try -skip-recent-events=3 to protect the previous tool chain from Pass 1,
//	# or -force-clean-large-log to always clean the demo tool by name.
package main

import (
	"context"
	"flag"
	"log"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	appName   = "context-compaction-demo"
	agentName = "context-compaction-agent"
	userID    = "demo-user"
)

var (
	modelName = flag.String(
		"model",
		os.Getenv("MODEL_NAME"),
		"Name of the OpenAI-compatible model to use (default: MODEL_NAME env var)",
	)
	streaming = flag.Bool(
		"streaming",
		false,
		"Enable streaming model responses",
	)
	debug = flag.Bool(
		"debug",
		true,
		"Print the projected request before every model call",
	)
	previewBytes = flag.Int(
		"preview-bytes",
		240,
		"Maximum content bytes to print for each request message",
	)
	logLines = flag.Int(
		"log-lines",
		240,
		"Number of synthetic log lines returned by the large_log tool",
	)
	toolResultMaxTokens = flag.Int(
		"tool-result-max-tokens",
		80,
		"Pass 1 token threshold for replacing historical tool results",
	)
	keepRecentRequests = flag.Int(
		"keep-recent-requests",
		0,
		"Number of latest completed requests protected from Pass 1",
	)
	skipRecentEvents = flag.Int(
		"skip-recent-events",
		0,
		"Number of tail events treated as recent by Pass 1",
	)
	oversizedToolResultMaxTokens = flag.Int(
		"oversized-tool-result-max-tokens",
		0,
		"Pass 2 token threshold for head+tail truncation; 0 disables Pass 2",
	)
	forceCleanLargeLog = flag.Bool(
		"force-clean-large-log",
		false,
		"Force the large_log tool result to be cleaned whenever context compaction is enabled",
	)
)

type logRequest struct {
	Lines int `json:"lines"`
}

type logResult struct {
	Summary string `json:"summary"`
	Log     string `json:"log"`
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type demoApp struct {
	runner    runner.Runner
	sessionID string
}

func newDemo() *demoApp { _ = "STUB: not implemented"; return nil }

// The callback sees the final request produced by the agent flow, after
// history projection and context compaction. This is the most direct way
// to verify what will be sent to the provider.

// Context compaction is a prompt-projection feature. It rewrites tool
// result content in the request sent to the model, while the full session
// event remains stored in the session service.

// Pass 1 protects the current request and this many latest completed
// requests. This demo defaults to 0 so the second turn clearly shows the
// previous tool result becoming historical.

// Pass 1 threshold. Historical tool results above this estimated token
// count are replaced with a compact placeholder.

// Pass 2 threshold. When positive, any oversized tool result, including
// one in the current request, is head+tail truncated.

// ForceCleanToolNames is useful for noisy tools whose raw output is
// rarely useful after the tool loop completes, such as shell/log
// dumping tools. This demo exposes it through -force-clean-large-log.

// KeepToolNames has higher priority than ForceCleanToolNames and
// skips context compaction for tools whose exact payload should stay
// visible to the model.

// SkipRecentFunc customizes what counts as "recent" for Pass 1. It
// does not disable Pass 2; an oversized recent tool result can still
// be head+tail truncated when Pass 2 is enabled.

func skipRecentFunc(events []event.Event) int { _ = "STUB: not implemented"; return 0 }

func makeLargeLog(_ context.Context, req logRequest) (logResult, error) {
	_ = "STUB: not implemented"
	return *new(logResult), nil
}

func (d *demoApp) runTurn(ctx context.Context, text string) error {
	_ = "STUB: not implemented"
	return nil
}

func drainEvents(events <-chan *event.Event) { _ = "STUB: not implemented"; return }

func eventError(evt *event.Event) error { _ = "STUB: not implemented"; return nil }

func printEvent(evt *event.Event) { _ = "STUB: not implemented"; return }

func printProjectedRequest(req *model.Request) { _ = "STUB: not implemented"; return }

func formatToolNames(tools map[string]tool.Tool) string { _ = "STUB: not implemented"; return "" }

func preview(content string, maxBytes int) string { _ = "STUB: not implemented"; return "" }
