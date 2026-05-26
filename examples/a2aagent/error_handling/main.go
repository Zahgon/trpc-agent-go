//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	appName = "a2a-error-handling"

	defaultHost = "127.0.0.1:18888"

	backendAgentName = "structured-error-agent"
	backendAgentDesc = "Demonstrates structured A2A task errors"

	backendErrorCode = "REMOTE_VALIDATION_FAILED"
	demoUserID       = "demo-user"
	demoPrompt       = "check the remote order status"

	agentCardPath      = "/.well-known/agent-card.json"
	serverReadyTimeout = 5 * time.Second
	probeInterval      = 100 * time.Millisecond
	probeTimeout       = 500 * time.Millisecond
	stopTimeout        = 5 * time.Second
	runTimeout         = 10 * time.Second
)

var host = flag.String(
	"host",
	defaultHost,
	"Host used by the local A2A server",
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

type structuredErrorAgent struct{}

func (a *structuredErrorAgent) Run(
	ctx context.Context,
	invocation *agent.Invocation,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *structuredErrorAgent) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

func (a *structuredErrorAgent) Info() agent.Info {
	_ = "STUB: not implemented"
	return *new(agent.Info)
}

func (a *structuredErrorAgent) SubAgents() []agent.Agent { _ = "STUB: not implemented"; return nil }

func (a *structuredErrorAgent) FindSubAgent(name string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func runScenario(name string, remoteAgent agent.Agent) error { _ = "STUB: not implemented"; return nil }

func responseText(resp *model.Response) string { _ = "STUB: not implemented"; return "" }

func waitForServer(
	serverURL string,
	serverErr <-chan error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func stopServer(server interface {
	Stop(ctx context.Context) error
}) {
	_ = "STUB: not implemented"
	return
}

func ptrValue(value *string) string { _ = "STUB: not implemented"; return "" }

func stringPtr(value string) *string { _ = "STUB: not implemented"; return nil }
