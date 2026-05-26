//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates prompt caching capabilities with Agent, multi-turn conversations, and tools.
// This example shows that prompt cache works even with tool calls.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/model/openai"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session/inmemory"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// generateLongSystemPrompt generates a system prompt with at least 1024 tokens.
// OpenAI requires minimum 1024 tokens for prompt caching to work.
func generateLongSystemPrompt() string { _ = "STUB: not implemented"; return "" }

// CalculatorInput represents input parameters for the calculator tool.
type CalculatorInput struct {
	Operation string  `json:"operation" jsonschema:"description=The operation to perform: add, subtract, multiply, divide, sqrt, power"`
	A         float64 `json:"a" jsonschema:"description=First operand"`
	B         float64 `json:"b,omitempty" jsonschema:"description=Second operand (optional for sqrt)"`
}

// createCalculatorTool creates a calculator tool for mathematical operations.
func createCalculatorTool() tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

// TimeInput represents input for the time tool.
type TimeInput struct {
	Format   string `json:"format,omitempty" jsonschema:"description=Time format: 'full', 'date', 'time', 'unix'"`
	Timezone string `json:"timezone,omitempty" jsonschema:"description=Timezone name like 'UTC', 'America/New_York', 'Asia/Shanghai'"`
}

// createTimeTool creates a tool to get current time.
func createTimeTool() tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

// UsageStats tracks token usage across requests.
type UsageStats struct {
	TotalPromptTokens int
	TotalCachedTokens int
	TotalRequests     int
	RequestsWithCache int
}

func (u *UsageStats) Add(usage *model.Usage) { _ = "STUB: not implemented"; return }

func (u *UsageStats) Print() { _ = "STUB: not implemented"; return }

// OpenAI cached tokens are 50% cheaper

func main() {
	modelName := flag.String("model", "gpt-4o", "Model name to use (e.g. gpt-4o)")
	flag.Parse()

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("⚠️  OPENAI_API_KEY not set")
		fmt.Println("Please set OPENAI_API_KEY environment variable to run this demo")
		return
	}

	fmt.Println("=== Prompt Cache Demo with Agent, Multi-turn & Tools ===")
	fmt.Println()
	fmt.Println("This demo shows that prompt caching works with:")
	fmt.Println("  ✓ Agent-based architecture (llmagent + runner)")
	fmt.Println("  ✓ Multi-turn conversations (session)")
	fmt.Println("  ✓ Tool calls (calculator, time)")
	fmt.Println()
	fmt.Println("OpenAI prompt caching requirements:")
	fmt.Println("  - Minimum 1024 tokens in prompt")
	fmt.Println("  - Same prompt prefix across requests")
	fmt.Println("  - Cache TTL: 5-10 minutes")
	fmt.Println("  - Cache is best-effort (not guaranteed)")
	fmt.Println()

	ctx := context.Background()

	// Cache optimization is enabled by default for OpenAI.
	llm := openai.New(
		*modelName,
		openai.WithAPIKey(apiKey),
	)

	// Create tools
	tools := []tool.Tool{
		createCalculatorTool(),
		createTimeTool(),
	}

	// Create agent with long system prompt for caching
	longSystemPrompt := generateLongSystemPrompt()
	agentInstance := llmagent.New(
		"cache-demo-agent",
		llmagent.WithModel(llm),
		llmagent.WithInstruction(longSystemPrompt),
		llmagent.WithDescription("An AI assistant demonstrating prompt caching with tools"),
		llmagent.WithTools(tools),
		llmagent.WithGenerationConfig(model.GenerationConfig{
			Stream: true,
		}),
	)

	// Create session service for multi-turn conversation
	sessionService := inmemory.NewSessionService()

	// Create runner
	r := runner.NewRunner(
		"prompt-cache-demo",
		agentInstance,
		runner.WithSessionService(sessionService),
	)
	defer r.Close()

	userID := "demo-user"
	sessionID := fmt.Sprintf("cache-session-%d", time.Now().Unix())

	// Track usage statistics
	stats := &UsageStats{}

	// Define test queries - mix of regular questions and tool calls
	testQueries := []struct {
		query       string
		description string
		expectTool  bool
	}{
		{
			query:       "What is the singleton design pattern? Give a brief explanation.",
			description: "Regular question (cache creation)",
			expectTool:  false,
		},
		{
			query:       "Calculate 123 * 456 + 789",
			description: "Tool call - calculator",
			expectTool:  true,
		},
		{
			query:       "What time is it now in UTC?",
			description: "Tool call - time",
			expectTool:  true,
		},
		{
			query:       "What is the factory method pattern? Brief answer.",
			description: "Regular question (expecting cache hit)",
			expectTool:  false,
		},
		{
			query:       "Calculate the square root of 144",
			description: "Tool call - calculator (expecting cache hit)",
			expectTool:  true,
		},
		{
			query:       "Explain the observer pattern briefly.",
			description: "Regular question (expecting cache hit)",
			expectTool:  false,
		},
	}

	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Starting multi-turn conversation with tools...")
	fmt.Println(strings.Repeat("=", 60))

	for i, tq := range testQueries {
		fmt.Printf("\n📝 Turn %d: %s\n", i+1, tq.description)
		fmt.Printf("   Query: %s\n", tq.query)
		if tq.expectTool {
			fmt.Println("   [Expecting tool call]")
		}

		start := time.Now()

		// Run the agent
		message := model.NewUserMessage(tq.query)
		eventChan, err := r.Run(ctx, userID, sessionID, message)
		if err != nil {
			fmt.Printf("   ❌ Error: %v\n", err)
			continue
		}

		// Collect response
		var finalResponse string
		var usage *model.Usage
		var toolCalled bool

		for evt := range eventChan {
			// Handle errors
			if evt.Error != nil {
				fmt.Printf("   ⚠️  Event error: %s\n", evt.Error.Message)
				continue
			}

			// Check for tool calls
			if len(evt.Response.Choices) > 0 && len(evt.Response.Choices[0].Message.ToolCalls) > 0 {
				toolCalled = true
				for _, tc := range evt.Response.Choices[0].Message.ToolCalls {
					fmt.Printf("   🔧 Tool call: %s\n", tc.Function.Name)
				}
			}

			// Collect content
			if len(evt.Response.Choices) > 0 {
				content := evt.Response.Choices[0].Message.Content
				if content != "" {
					finalResponse = content
				}
				// Also check delta for streaming
				delta := evt.Response.Choices[0].Delta.Content
				if delta != "" && finalResponse == "" {
					finalResponse = delta
				}
			}

			// Collect usage (usually in final event)
			if evt.Response.Usage != nil {
				usage = evt.Response.Usage
			}

			// Check for final response
			if evt.IsFinalResponse() {
				break
			}
		}

		elapsed := time.Since(start)

		// Print results
		if finalResponse != "" {
			// Truncate long responses
			displayResp := finalResponse
			if len(displayResp) > 200 {
				displayResp = displayResp[:200] + "..."
			}
			fmt.Printf("   💬 Response: %s\n", displayResp)
		}

		if toolCalled {
			fmt.Println("   ✓ Tool was called")
		}

		if usage != nil {
			stats.Add(usage)
			fmt.Printf("   ⏱️  Time: %v\n", elapsed)
			fmt.Printf("   📊 Prompt tokens: %d, Cached: %d\n",
				usage.PromptTokens,
				usage.PromptTokensDetails.CachedTokens)

			if usage.PromptTokensDetails.CachedTokens > 0 {
				cacheRate := float64(usage.PromptTokensDetails.CachedTokens) / float64(usage.PromptTokens) * 100
				fmt.Printf("   💰 Cache hit rate: %.2f%%\n", cacheRate)
			}
		} else {
			fmt.Printf("   ⏱️  Time: %v\n", elapsed)
			fmt.Println("   📊 Usage info not available in streaming mode")
		}

		// Small delay between requests
		time.Sleep(500 * time.Millisecond)
	}

	// Print final statistics
	stats.Print()

	fmt.Println("\n✅ Demo completed!")
	fmt.Println()
	fmt.Println("Key observations:")
	fmt.Println("• System prompt (>1024 tokens) enables prompt caching")
	fmt.Println("• Cache works across multi-turn conversations with session history")
	fmt.Println("• Tool calls don't prevent caching - system prompt prefix is still cached")
	fmt.Println("• Cache hits reduce cost by up to 50% (cached tokens are half price)")
	fmt.Println()
	fmt.Println("Note: OpenAI caching is best-effort and may not hit on every request.")
	fmt.Println("      Cache hits are more likely in production with sustained traffic.")
	fmt.Println("      Run this demo multiple times to see improved cache hit rates.")
}
