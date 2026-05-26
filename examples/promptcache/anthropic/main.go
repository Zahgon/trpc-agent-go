//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates Anthropic prompt caching with three independent controls.
// The example is structured in three phases to clearly show each cache option's effect:
//
//	Phase 1: System prompt caching only - stable system instructions
//	Phase 2: System + Tools caching - adding tool definitions caching
//	Phase 3: System + Tools + Messages caching - multi-turn dynamic breakpoint
//
// Usage:
//
//	export ANTHROPIC_API_KEY=sk-ant-xxx
//	go run main.go
package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/model/anthropic"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/session/inmemory"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// generateLongSystemPrompt generates a system prompt with at least 1024 tokens.
// Anthropic requires minimum 1024 tokens for prompt caching to work.
func generateLongSystemPrompt() string { _ = "STUB: not implemented"; return "" }

// CalculatorInput represents input parameters for the calculator tool.
type CalculatorInput struct {
	Operation string  `json:"operation" jsonschema:"description=The operation to perform: add, subtract, multiply, divide, sqrt, power"`
	A         float64 `json:"a" jsonschema:"description=First operand"`
	B         float64 `json:"b,omitempty" jsonschema:"description=Second operand (optional for sqrt)"`
}

func createCalculatorTool() tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

// TimeInput represents input for the time tool.
type TimeInput struct {
	Format   string `json:"format,omitempty" jsonschema:"description=Time format: 'full', 'date', 'time', 'unix'"`
	Timezone string `json:"timezone,omitempty" jsonschema:"description=Timezone name like 'UTC', 'America/New_York', 'Asia/Shanghai'"`
}

func createTimeTool() tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

// turnUsage records the cache metrics for a single request turn.
type turnUsage struct {
	Turn                int
	Phase               string
	Query               string
	InputTokens         int // new tokens (not cached)
	CacheReadTokens     int // tokens served from cache
	CacheCreationTokens int // tokens written to cache
	Elapsed             time.Duration
}

func (t *turnUsage) totalInput() int { _ = "STUB: not implemented"; return 0 }

func (t *turnUsage) cacheHitRate() float64 { _ = "STUB: not implemented"; return 0 }

// printTurnResult prints the result for a single turn.
func printTurnResult(tu *turnUsage, response string) { _ = "STUB: not implemented"; return }

// printPhaseStats prints aggregate stats for a phase.
func printPhaseStats(usages []*turnUsage) { _ = "STUB: not implemented"; return }

// runTurn executes one turn and collects usage metrics.
func runTurn(ctx context.Context, r runner.Runner, userID, sessionID, query string) (string, *model.Usage, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// extractUsage extracts cache-related metrics from model.Usage.
func extractUsage(usage *model.Usage) (inputTokens, cacheRead, cacheCreation int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// CacheCreationTokens is reported via PromptTokens when cache is first created.
// Anthropic SDK reports it separately; we capture it from the raw usage if available.
// In the current model.Usage mapping, cache_creation_input_tokens is not directly exposed
// as a separate field, but the relationship is:
//   total billed input = input_tokens + cache_read + cache_creation
// For display purposes, if no cache read and inputTokens is large, that likely includes creation.

func main() {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("ANTHROPIC_AUTH_TOKEN")
	}
	if apiKey == "" {
		fmt.Println("Please set ANTHROPIC_API_KEY environment variable")
		return
	}

	fmt.Println("=== Anthropic Prompt Cache - Three Phase Verification ===")
	fmt.Println()
	fmt.Println("Three independent cache options:")
	fmt.Println("  WithCacheSystemPrompt(true) - cache stable system instructions")
	fmt.Println("  WithCacheTools(true)         - cache tool definitions")
	fmt.Println("  WithCacheMessages(true)      - cache conversation history (dynamic breakpoint)")
	fmt.Println()
	fmt.Println("Anthropic cache key facts:")
	fmt.Println("  - Minimum 1024 tokens required for a cache breakpoint")
	fmt.Println("  - Cache read: 90% cheaper than normal input")
	fmt.Println("  - Cache creation: 25% extra cost (one-time)")
	fmt.Println("  - Cache TTL: ~5 minutes")
	fmt.Println()

	ctx := context.Background()
	tools := []tool.Tool{createCalculatorTool(), createTimeTool()}
	longSystemPrompt := generateLongSystemPrompt()
	maxTokens := 1024

	// ================================================================
	// Phase 1: System prompt caching only
	// ================================================================
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Phase 1: WithCacheSystemPrompt(true) ONLY")
	fmt.Println("  - System prompt (~1200 tokens) gets a cache breakpoint")
	fmt.Println("  - Tools and messages are NOT cached")
	fmt.Println("  - Expected: Turn 1 creates cache, Turn 2 reads from cache")
	fmt.Println(strings.Repeat("=", 60))

	llm1 := anthropic.New("claude-4-5-sonnet-20250929",
		anthropic.WithAPIKey(apiKey),
		anthropic.WithHeaders(map[string]string{"Venus-Sticky-Routing": "token"}),
		anthropic.WithCacheSystemPrompt(true),
	)
	agent1 := llmagent.New("phase1-agent",
		llmagent.WithModel(llm1),
		llmagent.WithInstruction(longSystemPrompt),
		llmagent.WithDescription("Phase 1: system prompt cache only"),
		llmagent.WithTools(tools),
		llmagent.WithGenerationConfig(model.GenerationConfig{
			MaxTokens: &maxTokens,
		}),
	)
	session1 := inmemory.NewSessionService()
	r1 := runner.NewRunner("phase1", agent1, runner.WithSessionService(session1))
	defer r1.Close()

	sid1 := fmt.Sprintf("phase1-%d", time.Now().Unix())
	phase1Queries := []struct {
		query string
		desc  string
	}{
		{"What is the singleton design pattern? One sentence answer.", "Turn 1: cache creation for system prompt"},
		{"What is the factory method pattern? One sentence answer.", "Turn 2: system prompt should hit cache"},
	}

	var phase1Usages []*turnUsage
	for i, q := range phase1Queries {
		fmt.Printf("\n[Turn %d] %s\n", i+1, q.desc)
		fmt.Printf("   Query: %s\n", q.query)

		start := time.Now()
		resp, usage, err := runTurn(ctx, r1, "user1", sid1, q.query)
		elapsed := time.Since(start)
		if err != nil {
			fmt.Printf("   Error: %v\n", err)
			continue
		}

		tu := &turnUsage{Turn: i + 1, Phase: "Phase1", Query: q.query, Elapsed: elapsed}
		if usage != nil {
			tu.InputTokens = usage.PromptTokens
			tu.CacheReadTokens = usage.PromptTokensDetails.CachedTokens
		}
		phase1Usages = append(phase1Usages, tu)
		printTurnResult(tu, resp)
		time.Sleep(1 * time.Second)
	}
	printPhaseStats(phase1Usages)

	// ================================================================
	// Phase 2: System + Tools caching
	// ================================================================
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Phase 2: WithCacheSystemPrompt(true) + WithCacheTools(true)")
	fmt.Println("  - Both system prompt and tool definitions get cache breakpoints")
	fmt.Println("  - Messages are NOT cached")
	fmt.Println("  - Expected: more tokens cached (system + tools prefix)")
	fmt.Println(strings.Repeat("=", 60))

	llm2 := anthropic.New("claude-4-5-sonnet-20250929",
		anthropic.WithAPIKey(apiKey),
		anthropic.WithHeaders(map[string]string{"Venus-Sticky-Routing": "token"}),
		anthropic.WithCacheSystemPrompt(true),
		anthropic.WithCacheTools(true),
	)
	agent2 := llmagent.New("phase2-agent",
		llmagent.WithModel(llm2),
		llmagent.WithInstruction(longSystemPrompt),
		llmagent.WithDescription("Phase 2: system + tools cache"),
		llmagent.WithTools(tools),
		llmagent.WithGenerationConfig(model.GenerationConfig{
			MaxTokens: &maxTokens,
		}),
	)
	session2 := inmemory.NewSessionService()
	r2 := runner.NewRunner("phase2", agent2, runner.WithSessionService(session2))
	defer r2.Close()

	sid2 := fmt.Sprintf("phase2-%d", time.Now().Unix())
	phase2Queries := []struct {
		query string
		desc  string
	}{
		{"Calculate 123 * 456", "Turn 1: cache creation (system + tools)"},
		{"Calculate 789 + 321", "Turn 2: system + tools should hit cache"},
	}

	var phase2Usages []*turnUsage
	for i, q := range phase2Queries {
		fmt.Printf("\n[Turn %d] %s\n", i+1, q.desc)
		fmt.Printf("   Query: %s\n", q.query)

		start := time.Now()
		resp, usage, err := runTurn(ctx, r2, "user1", sid2, q.query)
		elapsed := time.Since(start)
		if err != nil {
			fmt.Printf("   Error: %v\n", err)
			continue
		}

		tu := &turnUsage{Turn: i + 1, Phase: "Phase2", Query: q.query, Elapsed: elapsed}
		if usage != nil {
			tu.InputTokens = usage.PromptTokens
			tu.CacheReadTokens = usage.PromptTokensDetails.CachedTokens
		}
		phase2Usages = append(phase2Usages, tu)
		printTurnResult(tu, resp)
		time.Sleep(1 * time.Second)
	}
	printPhaseStats(phase2Usages)

	// ================================================================
	// Phase 3: All three - System + Tools + Messages (dynamic breakpoint)
	// ================================================================
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Phase 3: All caching enabled (System + Tools + Messages)")
	fmt.Println("  - WithCacheMessages(true) adds a dynamic cache breakpoint")
	fmt.Println("    at the LAST assistant message in conversation history.")
	fmt.Println()
	fmt.Println("  How dynamic message caching works:")
	fmt.Println("    Turn 1: [user] -> no assistant msg yet, only system+tools cached")
	fmt.Println("    Turn 2: [user, asst, user] -> breakpoint at asst (index 1)")
	fmt.Println("    Turn 3: [user, asst, user, asst, user] -> breakpoint moves to asst (index 3)")
	fmt.Println("    Each turn, the breakpoint moves forward to cover more history.")
	fmt.Println()
	fmt.Println("  This means:")
	fmt.Println("    - Turn 1: cache miss (first request, creating cache)")
	fmt.Println("    - Turn 2: cache HIT on system+tools, new cache created at asst msg")
	fmt.Println("    - Turn 3: cache HIT on system+tools+msgs up to prev asst msg")
	fmt.Println("    - Turn N: increasingly more tokens served from cache")
	fmt.Println(strings.Repeat("=", 60))

	llm3 := anthropic.New("claude-4-5-sonnet-20250929",
		anthropic.WithAPIKey(apiKey),
		anthropic.WithHeaders(map[string]string{"Venus-Sticky-Routing": "token"}),
		anthropic.WithCacheSystemPrompt(true),
		anthropic.WithCacheTools(true),
		anthropic.WithCacheMessages(true),
	)
	agent3 := llmagent.New("phase3-agent",
		llmagent.WithModel(llm3),
		llmagent.WithInstruction(longSystemPrompt),
		llmagent.WithDescription("Phase 3: full cache with dynamic message breakpoint"),
		llmagent.WithTools(tools),
		llmagent.WithGenerationConfig(model.GenerationConfig{
			MaxTokens: &maxTokens,
		}),
	)
	session3 := inmemory.NewSessionService()
	r3 := runner.NewRunner("phase3", agent3, runner.WithSessionService(session3))
	defer r3.Close()

	sid3 := fmt.Sprintf("phase3-%d", time.Now().Unix())

	// 6 turns to clearly show the dynamic breakpoint moving forward
	phase3Queries := []struct {
		query string
		desc  string
	}{
		{
			"What is the observer pattern? One sentence.",
			"Turn 1: no history yet, cache creation (system+tools)",
		},
		{
			"Please use the calculator tool to add 12345 and 67890.",
			"Turn 2: system+tools from cache, breakpoint set at Turn 1's assistant msg",
		},
		{
			"Please use the get_current_time tool to tell me the current time in UTC.",
			"Turn 3: system+tools+Turn1-2 history from cache, breakpoint moves to Turn 2's assistant msg",
		},
		{
			"What is the strategy pattern? One sentence.",
			"Turn 4: even more history cached, breakpoint moves to Turn 3's assistant msg",
		},
		{
			"Please use the calculator tool to multiply 999 by 111.",
			"Turn 5: most of the conversation served from cache",
		},
		{
			"Summarize all the design patterns you've mentioned so far in this conversation.",
			"Turn 6: references prior turns - large context mostly from cache",
		},
	}

	var phase3Usages []*turnUsage
	for i, q := range phase3Queries {
		fmt.Printf("\n[Turn %d] %s\n", i+1, q.desc)
		fmt.Printf("   Query: %s\n", q.query)

		start := time.Now()
		resp, usage, err := runTurn(ctx, r3, "user1", sid3, q.query)
		elapsed := time.Since(start)
		if err != nil {
			fmt.Printf("   Error: %v\n", err)
			continue
		}

		tu := &turnUsage{Turn: i + 1, Phase: "Phase3", Query: q.query, Elapsed: elapsed}
		if usage != nil {
			tu.InputTokens = usage.PromptTokens
			tu.CacheReadTokens = usage.PromptTokensDetails.CachedTokens
		}
		phase3Usages = append(phase3Usages, tu)
		printTurnResult(tu, resp)
		time.Sleep(1 * time.Second)
	}
	printPhaseStats(phase3Usages)

	// ================================================================
	// Final Summary: Compare all three phases
	// ================================================================
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("FINAL COMPARISON")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("%-10s %-8s %-12s %-12s %-10s\n", "Phase", "Turns", "CacheRead", "NewTokens", "HitRate")
	fmt.Println(strings.Repeat("-", 55))

	allPhases := []struct {
		name   string
		usages []*turnUsage
	}{
		{"Phase1", phase1Usages},
		{"Phase2", phase2Usages},
		{"Phase3", phase3Usages},
	}

	for _, p := range allPhases {
		var totalNew, totalCacheRead int
		for _, u := range p.usages {
			totalNew += u.InputTokens
			totalCacheRead += u.CacheReadTokens
		}
		total := totalNew + totalCacheRead
		rate := 0.0
		if total > 0 {
			rate = float64(totalCacheRead) / float64(total) * 100
		}
		fmt.Printf("%-10s %-8d %-12d %-12d %.1f%%\n",
			p.name, len(p.usages), totalCacheRead, totalNew, rate)
	}

	fmt.Println()
	fmt.Println("Key takeaways:")
	fmt.Println("  Phase 1: Only system prompt cached - baseline savings")
	fmt.Println("  Phase 2: System + tools cached - more tokens covered")
	fmt.Println("  Phase 3: Dynamic message caching - cache grows with conversation")
	fmt.Println("           Breakpoint moves to latest assistant message each turn,")
	fmt.Println("           so cache hit rate increases as conversation gets longer.")
	fmt.Println()
	fmt.Println("When to use each option:")
	fmt.Println("  WithCacheSystemPrompt(true) - system prompt is stable (recommended)")
	fmt.Println("  WithCacheTools(true)         - tools don't change often (recommended)")
	fmt.Println("  WithCacheMessages(true)      - multi-turn conversations with 3+ turns")
	fmt.Println("                                 (25% creation cost per turn, 90% savings on reads)")
}
