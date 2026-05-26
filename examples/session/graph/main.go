//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a multi-turn graph agent with session persistence.
// It shows that graph completion snapshot keys are visible on runner completion
// events, while only compact business state and response identity are retained
// in session.State.
//
// Usage:
//
//	go run ./graph
//	go run ./graph -debug=false
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	appName         = "session-graph-demo"
	userID          = "user"
	keyNormalized   = "normalized_input"
	keyDraft        = "draft_answer"
	keyAgentReply   = "agent_reply"
	keyAgentReplyID = "agent_reply_id"
	keyBusiness     = "business_result"
)

var (
	modelName = flag.String(
		"model",
		os.Getenv("MODEL_NAME"),
		"Name of the model to use (default: MODEL_NAME env var)",
	)
	streaming = flag.Bool(
		"streaming",
		true,
		"Enable streaming mode for the LLM agent node",
	)
	debugMode = flag.Bool(
		"debug",
		true,
		"Print session events and state after each turn",
	)
)

type graphChat struct {
	runner         runner.Runner
	sessionService session.Service
	sessionID      string
	debug          bool
}

func main() {
	flag.Parse()
	if *modelName == "" {
		*modelName = "deepseek-chat"
	}
	chat := &graphChat{
		sessionID: fmt.Sprintf("graph-session-%d", time.Now().Unix()),
		debug:     *debugMode,
	}
	if err := chat.run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func (c *graphChat) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *graphChat) setup() error { _ = "STUB: not implemented"; return nil }

func (c *graphChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *graphChat) printCommands() { _ = "STUB: not implemented"; return }

func (c *graphChat) processMessage(ctx context.Context, userInput string) error {
	_ = "STUB: not implemented"
	return nil
}

func buildGraphAgent() (*graphagent.GraphAgent, error) { _ = "STUB: not implemented"; return nil, nil }

func buildAssistantAgent() (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func normalizeInput(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func draftAnswer(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func collectAnswer(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func lastRunnerCompletion(events <-chan *event.Event) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func completionText(completion *event.Event) string { _ = "STUB: not implemented"; return "" }

func (c *graphChat) printDebug(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *graphChat) printCurrentState(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func printSessionDebug(sess *session.Session) { _ = "STUB: not implemented"; return }

func printSessionEvents(sess *session.Session) { _ = "STUB: not implemented"; return }

func printSessionState(sess *session.Session) { _ = "STUB: not implemented"; return }

func (c *graphChat) currentSession(ctx context.Context) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *graphChat) startNewSession(customID string) { _ = "STUB: not implemented"; return }

func (c *graphChat) listSessions(ctx context.Context) { _ = "STUB: not implemented"; return }

func graphSnapshotKeys() []string { _ = "STUB: not implemented"; return nil }

func sortedSessionStateKeys(state session.StateMap) []string { _ = "STUB: not implemented"; return nil }

func formatStateValue(raw []byte) string { _ = "STUB: not implemented"; return "" }
