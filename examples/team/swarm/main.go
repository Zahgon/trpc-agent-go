//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a Swarm Team.
//
// A Swarm Team starts from an entry Agent. Each Agent can transfer control to
// another Agent using transfer_to_agent. The last Agent reply is the final
// answer.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"

	"trpc.group/trpc-go/trpc-agent-go/examples/team/internal/chat"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	appName  = "team-swarm-example"
	teamName = "team"

	agentOptimist   = "optimist"
	agentSkeptic    = "skeptic"
	agentPragmatist = "pragmatist"
	agentSummarizer = "summarizer"

	entryAgentName = agentOptimist

	defaultModelName = "deepseek-v4-flash"
	defaultVariant   = "openai"

	defaultTimeout = 5 * time.Minute

	defaultMaxTokens   = 2000
	defaultTemperature = 0.7

	sessionPrefix = "demo-"
	demoUserID    = "demo-user"

	dividerWidth = 50
)

var (
	modelName = flag.String(
		"model",
		defaultModelName,
		"Model name",
	)
	variant = flag.String(
		"variant",
		defaultVariant,
		"OpenAI provider variant",
	)
	streaming = flag.Bool(
		"streaming",
		true,
		"Enable streaming",
	)
	timeout = flag.Duration(
		"timeout",
		defaultTimeout,
		"Request timeout",
	)
	crossRequestTransfer = flag.Bool(
		"cross-request-transfer",
		false,
		"Enable cross-request transfer: after a transfer, the next user message "+
			"will start from the target agent (default: false)",
	)
)

func main() {
	flag.Parse()

	runnerInstance, err := buildRunner(
		*modelName,
		*variant,
		*streaming,
	)
	if err != nil {
		log.Fatalf("build runner: %v", err)
	}
	defer runnerInstance.Close()

	sessionID := sessionPrefix + uuid.NewString()

	fmt.Printf("Session: %s\n", sessionID)
	fmt.Printf("Timeout: %s\n", timeout.String())
	fmt.Printf("EntryAgent: %s\n", entryAgentName)
	fmt.Printf("Type %q to exit\n", chat.DefaultExitCommand)
	fmt.Println(strings.Repeat("=", dividerWidth))

	loopCfg := chat.LoopConfig{
		Runner:        runnerInstance,
		UserID:        demoUserID,
		SessionID:     sessionID,
		Timeout:       *timeout,
		ShowInner:     true,
		RootAgentName: teamName,
		ExitCommand:   chat.DefaultExitCommand,
	}

	if err := chat.Run(context.Background(), loopCfg); err != nil {
		log.Fatalf("run: %v", err)
	}
}

func buildRunner(
	modelName string,
	variant string,
	streaming bool,
) (runner.Runner, error) {
	_ = "STUB: not implemented"
	return *new(runner.Runner), nil
}

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
