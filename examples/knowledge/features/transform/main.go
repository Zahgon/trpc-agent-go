//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates document transformation capabilities using transformers.
//
// Required environment variables:
//   - OPENAI_API_KEY: Your OpenAI API key for embeddings
//   - OPENAI_BASE_URL: (Optional) Custom OpenAI API endpoint
//
// Example usage:
//
//	export OPENAI_API_KEY=sk-xxxx
//	export OPENAI_BASE_URL=https://api.openai.com/v1
//	go run main.go
package main

import (
	"context"
	"fmt"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/transform"

	_ "trpc.group/trpc-go/trpc-agent-go/knowledge/document/reader/text"
)

func main() {
	ctx := context.Background()

	fmt.Println("Transform Demo")
	fmt.Println("==============")

	// Example 1: No transformer (baseline)
	// Shows original content with tabs, multiple spaces, and "xxxxx" patterns
	fmt.Println("\n1. No Transformer (baseline)")
	runDemo(ctx, "No Transform")

	// Example 2: CharFilter - Remove tab characters
	// Removes all \t characters from the content
	fmt.Println("\n2. CharFilter: Remove tabs")
	charFilter := transform.NewCharFilter("\t")
	runDemo(ctx, "CharFilter", charFilter)

	// Example 3: CharDedup - Collapse consecutive spaces and 'x' characters
	// "     " -> " " and "xxxxx" -> "x"
	fmt.Println("\n3. CharDedup: Collapse consecutive spaces and 'x' characters")
	charDedup := transform.NewCharDedup(" ", "x")
	runDemo(ctx, "CharDedup", charDedup)

	// Example 4: Combined transformers
	// First remove tabs, then collapse spaces and 'x' characters
	fmt.Println("\n4. Combined: CharFilter(tabs) + CharDedup(spaces, x)")
	filter := transform.NewCharFilter("\t")
	dedup := transform.NewCharDedup(" ", "x")
	runDemo(ctx, "Combined", filter, dedup)
}

func runDemo(ctx context.Context, demoName string, transformers ...transform.Transformer) {
	_ = "STUB: not implemented"
	// Create source with transformers
	return
}

// Create knowledge base

// Show document info

// Search to get all chunks content

// Show all chunks
