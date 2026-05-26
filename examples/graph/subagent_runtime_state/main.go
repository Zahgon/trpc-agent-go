//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates using a sub-agent inside a GraphAgent while
// propagating graph state to the sub-agent via Invocation.RunOptions.RuntimeState.
// It shows:
//   - Pre node parses time and loads scene info → writes to state
//   - Sub-agent (LLMAgent) reads the graph state from ctx in model/tool callbacks
//   - Tools use parsed time from runtime state instead of LLM-guessed values
//   - Interactive streaming output from the single graph event channel
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultModelName = "deepseek-v4-flash"
)

var (
	modelName = flag.String("model", defaultModelName, "Model name (OpenAI compatible)")
	sceneID   = flag.String("scene", "scene-1", "Scene ID to fetch context for")
	baseURL   = flag.String("base-url", os.Getenv("OPENAI_BASE_URL"), "OpenAI-compatible base URL (optional)")
	apiKey    = flag.String("api-key", os.Getenv("OPENAI_API_KEY"), "API key (optional; falls back to env)")
)

func main() {
	flag.Parse()
	fmt.Printf("🚀 Sub‑Agent Runtime State (GraphAgent)\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Println(strings.Repeat("=", 50))

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

// Build graph

// GraphAgent with sub‑agent and initial empty state

// Runner + session service

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Interactive loop

// Build runtime state for this run; pre node will also add to/transform it

// buildGraphAndSubAgent constructs:
//   - Graph with a pre node then an agent node "assistant"
//   - LLMAgent sub‑agent with callbacks reading graph state from ctx
func buildGraphAndSubAgent(modelName, baseURL, apiKey string) (*graph.Graph, agent.Agent, error) {
	_ = "STUB: not implemented"
	// 1) Sub‑agent tools
	// schedule_meeting uses parsed_time from runtime state
	return nil, *new(agent.Agent), nil
}

// ignored if runtime provides parsed_time

// override with graph‑parsed time

// 2) Sub‑agent model callbacks – inject scene knowledge (English, tool-friendly)

// Prepend a system message carrying scene knowledge and guidance.
// Always keep it English and tool-friendly to avoid suppressing tool calls.

// Demonstrate we can read parsed_time from runtime state inside sub‑agent callback

// 3) Sub‑agent configuration
// Build model with optional explicit API key/base URL to avoid silent misconfig.

// 4) Graph
// Extend Messages schema with our custom keys so state merges print nicely

// Route from pre -> assistant; without this, the graph would stop after pre.

func mustNewLLMAgent(name string, m model.Model, instruction string, tools []tool.Tool, mc *model.Callbacks, tc *tool.Callbacks) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

// preNode acts as a pre‑processor:
// - loads scene info based on scene_id (dummy in this example)
// - parses a time from current user_input
// - writes them to state so the sub‑agent can read via RuntimeState in callbacks
func preNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Also allow passing from RunOptions at call site

// Scene info in English to prevent non-English model behavior

// Naive time parsing from phrases like "today/tomorrow HH:MM" or explicit 2006-01-02 15:04

// parseTimeInText is a tiny helper for demo purposes.
func parseTimeInText(s string) string { _ = "STUB: not implemented"; return "" }

// Try explicit: 2006-01-02 15:04

// Simple today/tomorrow HH:MM (using keywords: today, tomorrow)

// English: today/tomorrow at HH:MM or HHam/HHpm patterns

// streamEvents prints streaming content and basic execution metadata.
func streamEvents(ch <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

// Print minimal execution metadata from StateDelta (model/tool phases)

// Errors

// Fallback: sub-agent tool.response without graph metadata

// Streaming text

func truncate(s string, n int) string { _ = "STUB: not implemented"; return "" }

// tiny helper to avoid importing reflect in many places here
func reflectTypeOfString() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
