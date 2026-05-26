//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates message normalization with the message merger
// plugin against a real OpenAI-compatible backend.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	appName       = "message-merger-demo"
	agentName     = "message-merger-assistant"
	userID        = "demo-user"
	envOpenAIKey  = "OPENAI_API_KEY"
	envOpenAIBase = "OPENAI_BASE_URL"
)

var (
	modelName = flag.String(
		"model",
		"gpt-4o-mini",
		"Name of the model to use",
	)
)

func main() {
	flag.Parse()
	history := demoHistory()
	fmt.Println("Message Merger Plugin Demo")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Auth: OpenAI SDK reads %s and %s from env\n", envOpenAIKey, envOpenAIBase)
	if os.Getenv(envOpenAIKey) == "" {
		fmt.Printf("Hint: %s is not set. Configure a real OpenAI-compatible backend before running this example.\n", envOpenAIKey)
	}
	fmt.Println(strings.Repeat("=", 72))
	printMessages("Caller-supplied messages", history)
	fmt.Println()
	fmt.Println("Running the runner with messagemerger enabled.")
	if err := runScenario(history); err != nil {
		fmt.Printf("Run failed: %v\n", err)
		os.Exit(1)
	}
}

func demoHistory() []model.Message { _ = "STUB: not implemented"; return nil }

func runScenario(history []model.Message) error { _ = "STUB: not implemented"; return nil }
