//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates how to fetch a prompt from Langfuse, render it
// with runtime variables, and dynamically apply it to an llmagent via the
// existing SetInstruction API before each run.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/prompt"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	modelName   = flag.String("model", "deepseek-v4-flash", "Model name to use")
	streaming   = flag.Bool("stream", true, "Enable streaming responses")
	promptName  = flag.String("prompt-name", "movie-critic", "Langfuse prompt name")
	promptLabel = flag.String("prompt-label", "production", "Langfuse prompt label")
	movieTitle  = flag.String("movie", "", "Run once with the given movie title and exit")
	criticLevel = flag.String("critic-level", "expert", "Value for the {{criticlevel}} prompt variable")
	cacheTTL    = flag.Duration("cache-ttl", 30*time.Second, "TTL for the Langfuse prompt source cache")
	followUp    = flag.String("follow-up", "Answer in 2-3 sentences and mention one strength and one weakness.", "User message sent after the instruction is refreshed from Langfuse")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	app := &langfusePromptChat{
		modelName:   *modelName,
		streaming:   *streaming,
		promptName:  *promptName,
		promptLabel: *promptLabel,
		movieTitle:  *movieTitle,
		criticLevel: *criticLevel,
		cacheTTL:    *cacheTTL,
		followUp:    *followUp,
		userID:      "langfuse-demo-user",
	}
	if err := app.run(ctx); err != nil {
		log.Fatalf("example failed: %v", err)
	}
}

type langfusePromptChat struct {
	modelName   string
	streaming   bool
	promptName  string
	promptLabel string
	movieTitle  string
	criticLevel string
	cacheTTL    time.Duration
	followUp    string
	userID      string

	source prompt.Source
	agent  *llmagent.LLMAgent
	runner runner.Runner
}

func (c *langfusePromptChat) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *langfusePromptChat) setup(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *langfusePromptChat) startChat(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *langfusePromptChat) reviewMovie(ctx context.Context, movie string) error {
	_ = "STUB: not implemented"
	return nil
}

// Use the existing llmagent API to update the instruction dynamically.

func (c *langfusePromptChat) renderInstruction(ctx context.Context, movie string) (prompt.Text, string, error) {
	_ = "STUB: not implemented"
	return *new(prompt.Text), "", nil
}

func (c *langfusePromptChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *langfusePromptChat) extractContent(choice model.Choice) string {
	_ = "STUB: not implemented"
	return ""
}

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
