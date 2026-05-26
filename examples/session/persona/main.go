//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a simple per-session persona example.
//
// Each session stores its own persona in session state. Before every
// runner.Run call, the demo loads that persona and passes it through
// agent.WithGlobalInstruction(...), so the active system prompt is decided
// dynamically for that run.
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"

	util "trpc.group/trpc-go/trpc-agent-go/examples/session"
)

const (
	appName              = "session-persona-demo"
	agentName            = "persona-assistant"
	defaultUserID        = "user"
	personaStateKey      = "assistant_persona"
	defaultModelName     = "deepseek-v4-flash"
	defaultSessionType   = "inmemory"
	defaultEventLimit    = 1000
	defaultBannerWidth   = 72
	defaultPreviewMax    = 96
	escapedNewline       = "\\n"
	actualNewline        = "\n"
	personaSessionPrefix = "persona-session"

	commandExit        = "/exit"
	commandPersona     = "/persona"
	commandShowPersona = "/show-persona"
	commandSessions    = "/sessions"
	commandNew         = "/new"
	commandUse         = "/use"

	defaultSessionTTL = 24 * time.Hour

	defaultPersona = "You are a practical Go mentor for this session. " +
		"Prefer concise answers, explain trade-offs, and keep examples " +
		"compact."
	instructionText = "Answer the latest user request directly. Follow the " +
		"active session persona for tone, expertise, and response style."
	setPersonaUsage = "Usage: /persona <text>"
)

var (
	modelName = flag.String(
		"model",
		os.Getenv("MODEL_NAME"),
		"Name of the model to use (default: MODEL_NAME env var or "+
			"deepseek-v4-flash)",
	)
	sessionType = flag.String(
		"session",
		defaultSessionType,
		"Session backend: inmemory / sqlite / redis / postgres / mysql / "+
			"clickhouse",
	)
	eventLimit = flag.Int(
		"event-limit",
		defaultEventLimit,
		"Maximum number of events to store per session",
	)
	sessionTTL = flag.Duration(
		"session-ttl",
		defaultSessionTTL,
		"Session time-to-live duration",
	)
	streaming = flag.Bool(
		"streaming",
		true,
		"Enable streaming mode",
	)
)

type personaDemo struct {
	modelName      string
	sessionType    string
	eventLimit     int
	sessionTTL     time.Duration
	streaming      bool
	runner         runner.Runner
	sessionService session.Service
	userID         string
	sessionID      string
}

func main() {
	flag.Parse()

	demo := &personaDemo{
		modelName:   getModelName(),
		sessionType: *sessionType,
		eventLimit:  *eventLimit,
		sessionTTL:  *sessionTTL,
		streaming:   *streaming,
	}
	if err := demo.run(); err != nil {
		log.Fatalf("Session persona demo failed: %v", err)
	}
}

func getModelName() string { _ = "STUB: not implemented"; return "" }

func validateSessionType(value string) (util.SessionType, error) {
	_ = "STUB: not implemented"
	return *new(util.SessionType), nil
}

func (d *personaDemo) run() error { _ = "STUB: not implemented"; return nil }

func (d *personaDemo) setup(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *personaDemo) printIntro(ctx context.Context) { _ = "STUB: not implemented"; return }

func (d *personaDemo) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *personaDemo) handleCommand(
	ctx context.Context,
	userInput string,
) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

func hasCommandPrefix(input, command string) bool { _ = "STUB: not implemented"; return false }

func commandArgument(input string) string { _ = "STUB: not implemented"; return "" }

func normalizeInput(value string) string { _ = "STUB: not implemented"; return "" }

func (d *personaDemo) processMessage(
	ctx context.Context,
	userInput string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func buildPersonaInstruction(persona string) string { _ = "STUB: not implemented"; return "" }

func (d *personaDemo) processResponse(
	eventChan <-chan *event.Event,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *personaDemo) extractContent(choice model.Choice) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *personaDemo) setPersona(
	ctx context.Context,
	persona string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *personaDemo) showPersona(ctx context.Context) { _ = "STUB: not implemented"; return }

func (d *personaDemo) currentPersona(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *personaDemo) listSessions(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *personaDemo) switchSession(
	ctx context.Context,
	targetSessionID string,
	announceNew bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *personaDemo) ensureSession(
	ctx context.Context,
	targetSessionID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *personaDemo) loadSession(
	ctx context.Context,
	targetSessionID string,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *personaDemo) updatePersona(
	ctx context.Context,
	targetSessionID string,
	persona string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func personaFromSession(sess *session.Session) string { _ = "STUB: not implemented"; return "" }

func singleLinePersona(persona string) string { _ = "STUB: not implemented"; return "" }

func newSessionID() string { _ = "STUB: not implemented"; return "" }
