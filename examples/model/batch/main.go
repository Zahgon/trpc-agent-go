//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates how to use the Batch APIs with the OpenAI-like
// model in trpc-agent-go.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	openaisdk "github.com/openai/openai-go"
	"trpc.group/trpc-go/trpc-agent-go/model/openai"
)

// Constants to avoid magic strings and provide sane defaults.
const (
	defaultModelName = "gpt-4o-mini"
	defaultAction    = "list" // Supported: create|get|cancel|list.
	defaultLimit     = int64(5)
	defaultWindow    = "24h" // Batch completion window.
)

func main() {
	// CLI flags for different actions.
	action := flag.String("action", defaultAction, "Action: create|"+
		"get|cancel|list")
	modelName := flag.String("model", defaultModelName, "Model name to use")

	// Create options: user-provided requests.
	requestsInline := flag.String("requests", "", "Inline requests spec. "+
		"Format: 'role: msg || role: msg /// role: msg || role: msg'.")
	requestsFile := flag.String("file", "", "Path to requests spec file. "+
		"Same format as -requests, '///' between requests, '||' between messages.")

	// Flags for get/cancel.
	batchID := flag.String("id", "", "Batch ID for get/cancel")

	// Flags for list.
	after := flag.String("after", "", "Pagination cursor for listing batches")
	limit := flag.Int64("limit", defaultLimit, "Max number of batches to list "+
		"(1-100)")

	flag.Parse()

	fmt.Printf("🚀 Using configuration:\n")
	fmt.Printf("   📝 Model Name: %s\n", *modelName)
	fmt.Printf("   🎛️  Action: %s\n", *action)
	fmt.Printf("   🔑 OpenAI SDK reads OPENAI_API_KEY and OPENAI_BASE_URL from env\n")
	fmt.Println()

	// Initialize model.
	llm := openai.New(*modelName)

	ctx := context.Background()

	var err error
	switch *action {
	case "create":
		err = runCreate(ctx, llm, *requestsInline, *requestsFile)
	case "get":
		err = runGet(ctx, llm, *batchID)
	case "cancel":
		err = runCancel(ctx, llm, *batchID)
	case "list":
		err = runList(ctx, llm, *after, *limit)
	default:
		err = fmt.Errorf("unknown action: %s", *action)
	}

	if err != nil {
		log.Printf("❌ %v", err)
	} else {
		fmt.Println("🎉 Done.")
	}
}

// runCreate builds requests from user spec and creates a batch.
func runCreate(
	ctx context.Context,
	llm *openai.Model,
	inlineSpec string,
	filePath string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// parseRequestsSpec parses a simple textual spec into batch requests.
// Requests are separated by '///'. Messages within a request are separated by
// '||'. Each message line uses 'role: content'. Roles: system|user|assistant.
func parseRequestsSpec(spec string) ([]*openai.BatchRequestInput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// splitBy splits by sep and trims spaces for each piece.
func splitBy(s, sep string) []string { _ = "STUB: not implemented"; return nil }

// splitRoleContent splits 'role: content' into parts.
func splitRoleContent(s string) (role, content string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// runGet retrieves a batch by ID.
func runGet(ctx context.Context, llm *openai.Model, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// If output is available, download and parse it.

// Extract first content from ChatCompletion body.

// Print error details if present.

// runCancel cancels a batch by ID.
func runCancel(ctx context.Context, llm *openai.Model, id string) error {
	_ = "STUB: not implemented"
	return nil
}

// runList lists batches with pagination.
func runList(ctx context.Context, llm *openai.Model, after string, limit int64) error {
	_ = "STUB: not implemented"
	return nil
}

// printBatchListItem prints detailed information for a single batch item in the list.
func printBatchListItem(index int, item openaisdk.Batch) { _ = "STUB: not implemented"; return }

// Show additional details for each batch.

// Add empty line between batches for readability.

// printBatch prints key information of a batch.
func printBatch(prefix string, b *openaisdk.Batch) { _ = "STUB: not implemented"; return }

// Print errors if available.

// ts renders a unix seconds timestamp.
func ts(sec int64) string { _ = "STUB: not implemented"; return "" }
