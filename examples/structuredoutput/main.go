//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates structured output with LLMAgent using a minimal
// interactive runner-style CLI.
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	streaming = flag.Bool("streaming", true, "Enable streaming mode for responses")
)

func main() {
	flag.Parse()

	fmt.Printf("🚀 Structured Output (JSON Schema)\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Streaming: %t\n", *streaming)
	fmt.Printf("Type 'exit' to end the conversation\n")
	fmt.Println(strings.Repeat("=", 50))

	if err := run(); err != nil {
		log.Fatalf("run failed: %v", err)
	}
}

// placeRecommendation represents a simple structured output for real usage.
type placeRecommendation struct {
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	City       string  `json:"city"`
	Category   string  `json:"category"`
	Rating     float64 `json:"rating"`
	PriceLevel string  `json:"price_level"`
	Notes      string  `json:"notes"`
}

func run() error { _ = "STUB: not implemented"; return nil }

// OpenAI-compatible model.

// Minimal generation config.

// Build agent with structured output using a typed struct; schema auto-generated.

// Runner with in-memory session service.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Minimal event loop: print content as it arrives; show typed payload when available.

// If we got a typed structured output payload, display it succinctly.

// Print content as normal.

// Show raw JSON if not streamed (depends on provider behavior).

func intPtr(i int) *int           { _ = "STUB: not implemented"; return nil }
func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
