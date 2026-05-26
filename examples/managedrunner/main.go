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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	appName      = "managedrunner-demo"
	demoUserID   = "demo-user"
	agentName    = "ticker-agent"
	agentDesc    = "Emits periodic tick events until the context ends."
	messageText  = "start"
	tickFormat   = "tick %d"
	statusFormat = "  status: events=%d last=%s\n"
)

const (
	tickInterval       = 200 * time.Millisecond
	statusPollInterval = 300 * time.Millisecond
)

const (
	requestIDDetached     = "demo-run-detached"
	sessionDetached       = "demo-session-detached"
	requestIDManualCancel = "demo-run-manual-cancel"
	sessionManualCancel   = "demo-session-manual-cancel"
	requestIDMinDeadline  = "demo-run-min-deadline"
	sessionMinDeadline    = "demo-session-min-deadline"
)

const (
	parentCancelAfter   = 500 * time.Millisecond
	maxRunDetached      = 2 * time.Second
	manualCancelAfter   = 1 * time.Second
	maxRunManualCancel  = 10 * time.Second
	parentTimeout       = 1200 * time.Millisecond
	maxRunMinDeadline   = 5 * time.Second
	eventChannelBufSize = 1
)

func main() {
	baseRunner := runner.NewRunner(
		appName,
		newTickerAgent(agentName, tickInterval),
	)
	defer baseRunner.Close()

	managedRunner, ok := baseRunner.(runner.ManagedRunner)
	if !ok {
		fmt.Fprintln(
			os.Stderr,
			"runner does not implement runner.ManagedRunner",
		)
		os.Exit(1)
	}

	if err := demoDetachedCancel(managedRunner); err != nil {
		fmt.Fprintf(os.Stderr, "demo failed: %v\n", err)
		os.Exit(1)
	}
	if err := demoManualCancel(managedRunner); err != nil {
		fmt.Fprintf(os.Stderr, "demo failed: %v\n", err)
		os.Exit(1)
	}
	if err := demoMinDeadline(managedRunner); err != nil {
		fmt.Fprintf(os.Stderr, "demo failed: %v\n", err)
		os.Exit(1)
	}
}

func demoDetachedCancel(managedRunner runner.ManagedRunner) error {
	_ = "STUB: not implemented"
	return nil
}

func demoManualCancel(managedRunner runner.ManagedRunner) error {
	_ = "STUB: not implemented"
	return nil
}

func demoMinDeadline(managedRunner runner.ManagedRunner) error {
	_ = "STUB: not implemented"
	return nil
}

func consumeEvents(
	managedRunner runner.ManagedRunner,
	requestID string,
	start time.Time,
	eventChan <-chan *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func pollStatus(
	managedRunner runner.ManagedRunner,
	requestID string,
	done <-chan struct{},
) {
	_ = "STUB: not implemented"
	return
}

func firstContent(evt *event.Event) string { _ = "STUB: not implemented"; return "" }

type tickerAgent struct {
	name         string
	tickInterval time.Duration
}

func newTickerAgent(name string, tickInterval time.Duration) *tickerAgent {
	_ = "STUB: not implemented"
	return nil
}

func (a *tickerAgent) Run(
	ctx context.Context,
	invocation *agent.Invocation,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *tickerAgent) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

func (a *tickerAgent) Info() agent.Info { _ = "STUB: not implemented"; return *new(agent.Info) }

func (a *tickerAgent) SubAgents() []agent.Agent { _ = "STUB: not implemented"; return nil }

func (a *tickerAgent) FindSubAgent(_ string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func tickEvent(
	invocation *agent.Invocation,
	author string,
	tickCount int,
) *event.Event {
	_ = "STUB: not implemented"
	return nil
}
