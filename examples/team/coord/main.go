//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a coordinator Team.
//
// A coordinator Team has one coordinator Agent that consults member Agents
// and produces the final answer.
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
	"trpc.group/trpc-go/trpc-agent-go/team"
)

const (
	appName  = "team-coordinator-example"
	teamName = "team"

	agentRequirementsAnalyst = "requirements_analyst"
	agentSolutionDesigner    = "solution_designer"
	agentQualityReviewer     = "quality_reviewer"

	memberHistoryParent   = "parent"
	memberHistoryIsolated = "isolated"

	memberInnerTextInclude = string(team.InnerTextModeInclude)
	memberInnerTextExclude = string(team.InnerTextModeExclude)

	defaultModelName           = "deepseek-v4-flash"
	defaultVariant             = "openai"
	defaultMemberInnerTextMode = memberInnerTextInclude

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
	showInner = flag.Bool(
		"show-inner",
		true,
		"Show member transcript",
	)
	memberHistory = flag.String(
		"member-history",
		memberHistoryParent,
		"Member history scope: parent or isolated",
	)
	memberInnerText = flag.String(
		"member-inner-text",
		defaultMemberInnerTextMode,
		"Member inner text mode: include or exclude",
	)
	memberSkipSummarization = flag.Bool(
		"member-skip-summarization",
		false,
		"Skip coordinator summary after member tool",
	)
	enableParallelTools = flag.Bool(
		"parallel-tools",
		false,
		"Enable parallel tool execution",
	)
)

func main() {
	flag.Parse()

	runnerInstance, err := buildRunner(
		*modelName,
		*variant,
		*streaming,
		*showInner,
		*memberHistory,
		*memberInnerText,
		*memberSkipSummarization,
		*enableParallelTools,
	)
	if err != nil {
		log.Fatalf("build runner: %v", err)
	}
	defer runnerInstance.Close()

	sessionID := sessionPrefix + uuid.NewString()

	fmt.Printf("Session: %s\n", sessionID)
	fmt.Printf("Timeout: %s\n", timeout.String())
	fmt.Printf("ShowInner: %t\n", *showInner)
	fmt.Printf("MemberHistory: %s\n", *memberHistory)
	fmt.Printf("MemberInnerText: %s\n", *memberInnerText)
	fmt.Printf(
		"MemberSkipSummarization: %t\n",
		*memberSkipSummarization,
	)
	fmt.Printf("ParallelTools: %t\n", *enableParallelTools)
	fmt.Printf("Type %q to exit\n", chat.DefaultExitCommand)
	fmt.Println(strings.Repeat("=", dividerWidth))

	loopCfg := chat.LoopConfig{
		Runner:        runnerInstance,
		UserID:        demoUserID,
		SessionID:     sessionID,
		Timeout:       *timeout,
		ShowInner:     *showInner,
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
	showInner bool,
	memberHistory string,
	memberInnerText string,
	memberSkipSummarization bool,
	parallelTools bool,
) (runner.Runner, error) {
	_ = "STUB: not implemented"
	return *new(runner.Runner), nil
}

func parseMemberInnerTextMode(mode string) (team.InnerTextMode, error) {
	_ = "STUB: not implemented"
	return *new(team.InnerTextMode), nil
}

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
