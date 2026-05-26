//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates how to use the OpenAI-like model with environment variables.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/model/openai"
)

func main() {
	// Read configuration from command line flags.
	modelName := flag.String("model", "gpt-4o-mini", "Name of the model to use")
	flag.Parse()

	fmt.Printf("🚀 Using configuration:\n")
	fmt.Printf("   📝 Model Name: %s\n", *modelName)
	fmt.Printf("   🔑 OpenAI SDK will automatically read OPENAI_API_KEY and OPENAI_BASE_URL from environment\n")
	fmt.Println()

	// Create a new OpenAI-like model instance.
	// The OpenAI SDK will automatically read OPENAI_API_KEY and OPENAI_BASE_URL from environment variables.
	llm := openai.New(*modelName)

	ctx := context.Background()

	fmt.Println("🔄 === Non-streaming Example ===")
	if err := nonStreamingExample(ctx, llm); err != nil {
		log.Printf("❌ Non-streaming example failed: %v", err)
	}

	fmt.Println("\n🌊 === Streaming Example ===")
	if err := streamingExample(ctx, llm); err != nil {
		log.Printf("❌ Streaming example failed: %v", err)
	}

	fmt.Println("\n⚡ === Advanced Example with Parameters ===")
	if err := advancedExample(ctx, llm); err != nil {
		log.Printf("❌ Advanced example failed: %v", err)
	}

	fmt.Println("\n🧪 === Parameter Testing Example ===")
	if err := parameterTestingExample(ctx, llm); err != nil {
		log.Printf("❌ Parameter testing example failed: %v", err)
	}

	fmt.Println("🎉 === Demo Complete ===")
}

// nonStreamingExample demonstrates non-streaming usage.
func nonStreamingExample(ctx context.Context, llm *openai.Model) error {
	_ = "STUB: not implemented"
	return nil
}

// streamingExample demonstrates streaming usage.
func streamingExample(ctx context.Context, llm *openai.Model) error {
	_ = "STUB: not implemented"
	return nil
}

// advancedExample demonstrates advanced parameters and conversation.
func advancedExample(ctx context.Context, llm *openai.Model) error {
	_ = "STUB: not implemented"
	return nil
}

// Display response metadata.

// parameterTestingExample demonstrates various parameter combinations.
func parameterTestingExample(ctx context.Context, llm *openai.Model) error {
	_ = "STUB: not implemented"
	return nil
}

// testRequest sends a request and displays the response.
func testRequest(ctx context.Context, llm *openai.Model, request *model.Request, description string) error {
	_ = "STUB: not implemented"
	return nil
}

// Helper functions for creating pointers to primitive types.
func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }
