//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates multi-turn chat using Runner with multiple
// session backends. It highlights how to switch between sessions while the
// agent keeps per-session context.
//
// Usage:
//
//	go run main.go -session=inmemory
//	go run main.go -session=noop
//	go run main.go -session=sqlite
//	go run main.go -session=redis
//	go run main.go -session=postgres
//	go run main.go -session=pgvector
//	go run main.go -session=mysql
//	go run main.go -session=clickhouse
//
// Environment variables by session type (example usage):
//
//	sqlite:
//		export SQLITE_SESSION_DSN="file:sessions.db?_busy_timeout=5000"
//
//	redis:
//		export REDIS_ADDR="localhost:6379"
//
//	postgres:
//		export PG_HOST="localhost"
//		export PG_PORT="5432"
//		export PG_USER="postgres"
//		export PG_PASSWORD="password"
//		export PG_DATABASE="trpc_agent"
//
//	pgvector:
//		export PGVECTOR_HOST="localhost"
//		export PGVECTOR_PORT="5432"
//		export PGVECTOR_USER="postgres"
//		export PGVECTOR_PASSWORD="password"
//		export PGVECTOR_DATABASE="trpc-agent-go-pgsession"
//		export PGVECTOR_EMBEDDER_MODEL="text-embedding-3-small"
//		export OPENAI_EMBEDDING_API_KEY="$OPENAI_API_KEY"
//		export OPENAI_EMBEDDING_BASE_URL="$OPENAI_BASE_URL"
//
//	mysql:
//		export MYSQL_HOST="localhost"
//		export MYSQL_PORT="3306"
//		export MYSQL_USER="root"
//		export MYSQL_PASSWORD="password"
//		export MYSQL_DATABASE="trpc_agent"
//
//	clickhouse:
//		export CLICKHOUSE_HOST="localhost"
//		export CLICKHOUSE_PORT="9000"
//		export CLICKHOUSE_USER="default"
//		export CLICKHOUSE_PASSWORD=""
//		export CLICKHOUSE_DATABASE="trpc_agent"
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	alog "trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/telemetry/langfuse"
)

const appName = "session-demo"

type ctxKey string

const requestIDKey ctxKey = "requestID"

var (
	modelName = flag.String(
		"model",
		os.Getenv("MODEL_NAME"),
		"Name of the model to use (default: MODEL_NAME env var)",
	)
	sessServiceName = flag.String(
		"session",
		"redis",
		"Name of the session service to use, inmemory / noop / "+
			"sqlite / redis / postgres / pgvector / mysql / tdsql / clickhouse",
	)
	streaming = flag.Bool(
		"streaming",
		true,
		"Enable streaming mode for responses",
	)
	eventLimit = flag.Int(
		"event-limit",
		1000,
		"Maximum number of events to store per session",
	)
	sessionTTL = flag.Duration(
		"session-ttl",
		300*time.Second,
		"Session time-to-live duration",
	)
	searchTopK = flag.Int(
		"search-topk",
		5,
		"Maximum number of recalled events to show when /search is available",
	)
	debugMode = flag.Bool(
		"debug",
		true,
		"Enable debug mode to print session events after each "+
			"turn",
	)
	enableTrace = flag.Bool(
		"enable-trace",
		false,
		"Enable Langfuse tracing for session operations. "+
			"Requires LANGFUSE_SECRET_KEY, LANGFUSE_PUBLIC_KEY, LANGFUSE_HOST env vars.",
	)
)

func main() {
	flag.Parse()

	// Replace the default log functions to inject RequestID from context.
	origInfofContext := alog.InfofContext
	alog.InfofContext = func(ctx context.Context, format string, args ...any) {
		reqID, _ := ctx.Value(requestIDKey).(string)
		if reqID != "" {
			format = fmt.Sprintf("[req:%s] %s", reqID, format)
		}
		origInfofContext(ctx, format, args...)
	}

	if *enableTrace {
		clean, err := langfuse.Start(context.Background())
		if err != nil {
			log.Fatalf("failed to start langfuse tracer: %v", err)
		}
		defer func() {
			if err := clean(context.Background()); err != nil {
				log.Printf("langfuse tracer cleanup: %v", err)
			}
		}()
	}

	fmt.Printf("Session Management Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Printf("Event Limit: %d\n", *eventLimit)
	fmt.Printf("Session TTL: %v\n", *sessionTTL)
	fmt.Printf("Session Backend: %s\n", *sessServiceName)
	fmt.Printf("Debug Mode: %t\n", *debugMode)
	fmt.Println(strings.Repeat("=", 50))

	chat := &multiTurnChat{
		modelName: *modelName,
		streaming: *streaming,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

type multiTurnChat struct {
	modelName      string
	streaming      bool
	runner         runner.Runner
	sessionService session.Service
	searchable     session.SearchableService
	userID         string
	sessionID      string
	debugPersisted bool
}

// run starts the interactive chat session.
func (c *multiTurnChat) run() error { _ = "STUB: not implemented"; return nil }

// Setup the runner.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Start interactive chat.

func (c *multiTurnChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create session service using util.
	return nil
}

// startChat runs the interactive conversation loop.
func (c *multiTurnChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Print session events in debug mode.

// processMessage handles a single message exchange.
func (c *multiTurnChat) processMessage(
	ctx context.Context,
	userMessage string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Inject requestID into context for logging.

// Create a root span if tracing is enabled. The session service
// will attach its child spans (create_session, get_session, append_event)
// to this root span automatically via context propagation.

// Run the agent through the runner.

// Process response.

// processResponse handles streaming and non-streaming responses, with
// tool call visualization.
func (c *multiTurnChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if this is the final event.
// Do not break on tool response events (Done=true but not the
// final assistant response).

// handleEvent processes a single event from the event channel.
func (c *multiTurnChat) handleEvent(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) error {
	_ = "STUB: not implemented"
	// Handle errors.
	return nil
}

// Handle tool calls.

// Handle tool responses.

// Handle content.

// handleToolCalls detects and displays tool calls.
func (c *multiTurnChat) handleToolCalls(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// handleToolResponses detects and displays tool responses.
func (c *multiTurnChat) handleToolResponses(event *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// handleContent processes and displays content.
func (c *multiTurnChat) handleContent(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) {
	_ = "STUB: not implemented"
	return
}

// extractContent extracts content based on streaming mode.
func (c *multiTurnChat) extractContent(choice model.Choice) string {
	_ = "STUB: not implemented"

	// Streaming mode: use delta content.
	return ""
}

// Non-streaming mode: use full message content.

// displayContent prints content to console.
func (c *multiTurnChat) displayContent(
	content string,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) {
	_ = "STUB: not implemented"
	return
}

func (c *multiTurnChat) startNewSession(customID string) { _ = "STUB: not implemented"; return }

func (c *multiTurnChat) listSessions() { _ = "STUB: not implemented"; return }

func (c *multiTurnChat) switchSession(target string) { _ = "STUB: not implemented"; return }

func (c *multiTurnChat) recallSession(
	ctx context.Context,
	query string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func eventDisplay(evt event.Event) (string, string) { _ = "STUB: not implemented"; return "", "" }
