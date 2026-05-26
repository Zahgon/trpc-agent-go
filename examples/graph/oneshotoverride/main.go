//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates the OneShot override pattern using the graph
// package with GraphAgent and Runner. It sets one_shot_messages for a single
// round to completely control the model input, then verifies it is cleared.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

const (
	defaultModelName = "deepseek-v4-flash"
	appName          = "oneshot-override"
)

var (
	modelName = flag.String("model", defaultModelName,
		"Name of the model to use")
	inputFlag = flag.String("input", "",
		"User input to place in one_shot_messages. If empty, read stdin")
	sysFlag = flag.String("sys", "You are a domain expert. Use clear steps.",
		"System prompt to place in one_shot_messages")
)

func main() {
	flag.Parse()
	fmt.Printf("🚀 OneShot Override Example\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Println(strings.Repeat("=", 50))

	content := strings.TrimSpace(*inputFlag)
	if content == "" {
		var err error
		content, err = readSingleLine()
		if err != nil {
			log.Fatalf("failed to read input: %v", err)
		}
	}

	if err := runOnce(content, *sysFlag); err != nil {
		log.Fatalf("run failed: %v", err)
	}
}

func readSingleLine() (string, error) { _ = "STUB: not implemented"; return "", nil }

func runOnce(userText string, sysText string) error { _ = "STUB: not implemented"; return nil }

// Create one_shot_messages with system + user for this round.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Pass an empty user message; content is provided via OneShot.

func processEvents(eventChan <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

func truncate(s string, max int) string { _ = "STUB: not implemented"; return "" }

// verifyOneShotCleared confirms that one_shot_messages is cleared after the
// LLM node and prints a short note to the console.
func verifyOneShotCleared(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
