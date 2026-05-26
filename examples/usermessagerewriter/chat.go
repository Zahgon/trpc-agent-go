//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

// rewriterChat manages the demo conversation loop.
type rewriterChat struct {
	modelName      string
	streaming      bool
	runner         runner.Runner
	sessionService session.Service
	userID         string
	sessionID      string
}

func (c *rewriterChat) run() error { _ = "STUB: not implemented"; return nil }

// setup creates the runner and LLM agent used by the demo.
func (c *rewriterChat) setup(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// startChat runs the interactive terminal loop.
func (c *rewriterChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// processMessage runs one conversation turn through the rewriter-enabled runner.
func (c *rewriterChat) processMessage(ctx context.Context, userInput string) error {
	_ = "STUB: not implemented"
	return nil
}

// processResponse prints the assistant response from the event stream.
func (c *rewriterChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// dumpSession prints the persisted session transcript for debugging.
func (c *rewriterChat) dumpSession(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
