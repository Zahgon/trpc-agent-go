//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates the error message plugin.
//
// The example wires a minimal agent that always emits a raw error event (the
// same shape llmflow produces for StopError) and runs two Runners:
//
//   - one without the plugin, so Runner applies its default fallback message
//   - one with the plugin, so a custom, user-facing message is used instead
//
// The example does not call any model backend, so no API key is required.
package main

import (
	"fmt"
	"os"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	appName   = "error-message-demo"
	agentName = "stopping-agent"
	userID    = "demo-user"
)

func main() {
	fmt.Println("Error Message Plugin Demo")
	fmt.Println(strings.Repeat("=", 72))

	if err := runWithoutPlugin(); err != nil {
		fmt.Printf("default run failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(strings.Repeat("-", 72))

	if err := runWithPlugin(); err != nil {
		fmt.Printf("plugin run failed: %v\n", err)
		os.Exit(1)
	}
}

func runWithoutPlugin() error { _ = "STUB: not implemented"; return nil }

func runWithPlugin() error { _ = "STUB: not implemented"; return nil }

// Example: produce a friendly message for stop_agent_error, and a generic
// one for any other error type. The structured Response.Error is left
// intact, so debugging and downstream consumers keep the raw reason.

func drainRun(r runner.Runner, sessionID string) error { _ = "STUB: not implemented"; return nil }

// Drain; persistence happens in the runner event loop and we only
// care about what is stored on the session.

func printPersistedErrorEvent(
	svc session.Service,
	sessionID string,
) error {
	_ = "STUB: not implemented"
	return nil
}
