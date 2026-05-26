//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates safe-boundary user steering in a single run.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	modelName = flag.String(
		"model",
		defaultModelName(),
		"Name of the model to use",
	)
	toolDelay = flag.Duration(
		"tool-delay",
		2*time.Second,
		"How long the tool should wait before returning",
	)
	steerAfter = flag.Duration(
		"steer-after",
		1*time.Second,
		"When to enqueue the extra user message",
	)
	question = flag.String(
		"question",
		defaultQuestion,
		"Initial user message",
	)
	steerText = flag.String(
		"steer",
		defaultSteerText,
		"Extra user message to insert into the same run",
	)
)

const (
	fallbackModelName = "gpt-4.1-mini"
	openAIAPIKeyEnv   = "OPENAI_API_KEY"

	appName   = "steer-demo"
	agentName = "steer-agent"
	userID    = "demo-user"

	toolName        = "load_launch_brief"
	toolDescription = "Load the launch brief for a project."

	defaultQuestion  = "Draft a short launch announcement for Project Atlas."
	defaultSteerText = "Update the draft: make the tone warmer " +
		"and explicitly mention the May 20 launch date."
)

const agentInstruction = `You are a launch announcement assistant.

You must call load_launch_brief exactly once before answering.
After the tool result arrives, if there are newer user messages in the
conversation, you must incorporate the newest user requirements while keeping
the factual details from the tool result.
Keep the final answer under 120 words.`

type launchBriefRequest struct {
	Project string `json:"project"`
}

type launchBrief struct {
	Project    string   `json:"project"`
	LaunchDate string   `json:"launch_date"`
	Audience   string   `json:"audience"`
	Highlights []string `json:"highlights"`
}

type steerDemo struct {
	modelName  string
	toolDelay  time.Duration
	steerAfter time.Duration
	question   string
	steerText  string
}

func main() {
	flag.Parse()

	demo := &steerDemo{
		modelName:  *modelName,
		toolDelay:  *toolDelay,
		steerAfter: *steerAfter,
		question:   *question,
		steerText:  *steerText,
	}

	if err := demo.run(context.Background()); err != nil {
		log.Fatalf("demo failed: %v", err)
	}
}

func (d *steerDemo) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (d *steerDemo) enqueueSteer(
	ctx context.Context,
	r runner.Runner,
	requestID string,
) {
	_ = "STUB: not implemented"
	return
}

func (d *steerDemo) loadLaunchBrief(
	ctx context.Context,
	req launchBriefRequest,
) (launchBrief, error) {
	_ = "STUB: not implemented"
	return *new(launchBrief), nil
}

func (d *steerDemo) printRun(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func defaultModelName() string { _ = "STUB: not implemented"; return "" }

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
