//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates interactive token tailoring using the Runner with interactive command line interface.
package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

var (
	flagProvider             = flag.String("provider", "openai", "Name of the provider to use, openai/anthropic/ollama")
	flagModel                = flag.String("model", "", "Model name (auto-detected if empty)")
	flagEnableTokenTailoring = flag.Bool("enable-token-tailoring", true, "Enable automatic token tailoring based on model context window")
	flagMaxInputTokens       = flag.Int("max-input-tokens", 0, "Max input tokens for token tailoring (0 = auto-calculate from context window)")
	flagCounter              = flag.String("counter", "simple", "Token counter: simple|tiktoken")
	flagStrategy             = flag.String("strategy", "middle", "Tailoring strategy: middle|head|tail")
	flagStreaming            = flag.Bool("streaming", true, "Stream assistant responses")
	flagDebug                = flag.Bool("debug", false, "Enable debug logging")
)

// Interactive demo with /bulk to generate many messages and showcase
// counter/strategy usage without runner/session dependencies.
func main() {
	flag.Parse()
	if *flagDebug {
		log.SetLevel(log.LevelDebug)
	}

	// Auto-detect model name if not specified.
	provider := strings.ToLower(*flagProvider)
	modelName := *flagModel
	if modelName == "" {
		modelName = getDefaultModel(provider)
	}

	// Build model using provider package.
	modelInstance, err := buildModel(provider, modelName)
	if err != nil {
		log.Fatalf("Failed to build model: %v", err)
	}

	fmt.Printf("✂️  Token Tailoring Demo\n")
	fmt.Printf("🔌 provider: %s\n", provider)
	fmt.Printf("🧩 model: %s\n", modelName)
	fmt.Printf("🔧 enable-token-tailoring: %t\n", *flagEnableTokenTailoring)
	if *flagMaxInputTokens > 0 {
		fmt.Printf("🔢 max-input-tokens: %d\n", *flagMaxInputTokens)
	} else {
		fmt.Printf("🔢 max-input-tokens: auto (from context window)\n")
	}
	fmt.Printf("🧮 counter: %s\n", strings.ToLower(*flagCounter))
	fmt.Printf("🎛️ strategy: %s\n", strings.ToLower(*flagStrategy))
	fmt.Printf("📡 streaming: %t\n", *flagStreaming)
	fmt.Println("==================================================")
	fmt.Println("💡 Commands:")
	fmt.Println("  /bulk N     - append N synthetic user messages")
	fmt.Println("  /load FILE  - load messages from JSON file (e.g., /load input.json)")
	fmt.Println("  /history    - show current message count")
	fmt.Println("  /show       - display current messages (head + tail)")
	fmt.Println("  /exit       - quit")
	fmt.Println()

	counter := buildCounter(strings.ToLower(*flagCounter), modelName)
	scanner := bufio.NewScanner(os.Stdin)
	messages := []model.Message{model.NewSystemMessage("You are a helpful assistant.")}
	for {
		fmt.Print("👤 You: ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		handled := handleCommand(&messages, line)
		if handled {
			if line == "/exit" {
				return
			}
			continue
		}
		processTurn(context.Background(), modelInstance, counter, &messages, line)
	}
}

func getDefaultModel(provider string) string { _ = "STUB: not implemented"; return "" }

// buildModel creates a model instance using the provider package with unified configuration.
func buildModel(providerName, modelName string) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

// Build common provider options.

// Set provider-specific callbacks for token statistics display.

// Convert OpenAI messages back to model.Message for token counting.

// Display tailoring statistics.

// Show head and tail messages to visualize what was kept/removed.

// Convert Anthropic messages back to model.Message for token counting.

// Display tailoring statistics.

// Show head and tail messages to visualize what was kept/removed.

// Convert Ollama messages back to model.Message for token counting.

// Display tailoring statistics.

// Show head and tail messages to visualize what was kept/removed.

func buildStrategy(counter model.TokenCounter, strategyName string) model.TailoringStrategy {
	_ = "STUB: not implemented"
	return *new(model.TailoringStrategy)
}

func buildCounter(name string, modelName string) model.TokenCounter {
	_ = "STUB: not implemented"
	return *new(model.TokenCounter)
}

func handleCommand(messages *[]model.Message, line string) bool {
	_ = "STUB: not implemented"
	return false
}

// Replace messages with loaded ones (keep system message if first is not system).

// Prepend system message if not present.

func processTurn(ctx context.Context, m model.Model, counter model.TokenCounter, messages *[]model.Message, userLine string) {
	_ = "STUB: not implemented"
	return
}

// renderResponse prints streaming or non-streaming responses similar to runner example.
func renderResponse(ch <-chan *model.Response, streaming bool) string {
	_ = "STUB: not implemented"
	return ""
}

// Non-streaming mode
