//
// Tencent is pleased to support the open source community
// by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	appName   = "cancelrun-demo"
	userID    = "demo-user"
	sessionID = "demo-session"

	agentName = "slow-writer"

	startMessage = "start"

	maxRunDuration = 15 * time.Second
	chunkDelay     = 120 * time.Millisecond
	eventChanBuf   = 8
)

func main() {
	fmt.Println("Cancel a Run demo")
	fmt.Println("Press Enter to cancel.")
	fmt.Println("Press Ctrl+C to cancel (SIGINT).")
	fmt.Println()

	baseCtx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)
	defer stop()

	ctx, cancel := context.WithTimeout(baseCtx, maxRunDuration)
	defer cancel()

	go cancelOnEnter(cancel)

	r := runner.NewRunner(appName, newSlowWriter(agentName, chunkDelay))
	defer r.Close()

	eventCh, err := r.Run(
		ctx,
		userID,
		sessionID,
		model.NewUserMessage(startMessage),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "run failed: %v\n", err)
		os.Exit(1)
	}

	printEvents(eventCh)
	printExitReason(ctx)
}

func cancelOnEnter(cancel context.CancelFunc) { _ = "STUB: not implemented"; return }

func printEvents(eventCh <-chan *event.Event) { _ = "STUB: not implemented"; return }

func printExitReason(ctx context.Context) { _ = "STUB: not implemented"; return }

type slowWriter struct {
	name  string
	delay time.Duration
}

func newSlowWriter(name string, delay time.Duration) *slowWriter {
	_ = "STUB: not implemented"
	return nil
}

func (a *slowWriter) Run(
	ctx context.Context,
	invocation *agent.Invocation,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *slowWriter) stream(
	ctx context.Context,
	invocation *agent.Invocation,
	out chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func (a *slowWriter) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

func (a *slowWriter) Info() agent.Info { _ = "STUB: not implemented"; return *new(agent.Info) }

func (a *slowWriter) SubAgents() []agent.Agent { _ = "STUB: not implemented"; return nil }

func (a *slowWriter) FindSubAgent(_ string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

const demoChunkFormat = "chunk %d\n"

var demoIntroChunks = []string{
	"Streaming some text...\n",
	"Press Enter (or Ctrl+C) to stop.\n",
	"\n",
}

func demoEvent(
	invocation *agent.Invocation,
	author string,
	content string,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}
