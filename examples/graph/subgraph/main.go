//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates subgraph usage by composing a parent GraphAgent
// that delegates to a child GraphAgent via a Subgraph node (sugar over Agent node).
// It shows:
// - Building a child graph (LLM + Tools) and wrapping it as a sub-agent
// - Registering sub-agents on the parent GraphAgent
// - Delegating via AddSubgraphNode("assistant") and streaming forwarded events
// - Passing runtime state from parent to subgraph (tool reads from runtime state)
// - Interactive CLI with Runner using default streaming output
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

var (
	modelName = flag.String("model", os.Getenv("MODEL_NAME"), "OpenAI‑compatible model name (e.g., deepseek-v4-flash)")
	baseURL   = flag.String("base-url", os.Getenv("OPENAI_BASE_URL"), "OpenAI‑compatible base URL")
	apiKey    = flag.String("api-key", os.Getenv("OPENAI_API_KEY"), "API key")
	verbose   = flag.Bool("v", false, "Verbose: print model/tool metadata")
	// Subgraph feature toggles
	parentInclude = flag.String("parent-include", "", "Parent include mode: none|filtered|all (empty uses default)")
	subIsolate    = flag.Bool("sub-isolate", false, "Isolate subgraph from parent session messages (include_contents=none)")
	subScope      = flag.String("sub-scope", scopeAssistant, "Subgraph event scope segment (empty disables)")
	subInput      = flag.String("sub-input", inputModeParsed, "Subgraph input mapping: parsed|all")
	subOutput     = flag.String("sub-output", outputModeCustom, "Subgraph output mapping: custom|default")
)

// Deduplicate raw string literals for ids, keys, names, and modes
const (
	defaultModelName = "deepseek-v4-flash"

	appName    = "subgraph-app"
	parentName = "parent"
	// assistantID is the single source of truth for both the subgraph node ID and child agent name.
	assistantID    = "assistant"
	childAgentName = assistantID

	nodeParse     = "parse_input"
	nodeSubgraph  = assistantID
	nodeCollect   = "collect"
	nodeLLMDecide = "llm_decider"
	nodeTools     = "tools"

	toolScheduleName = "schedule_meeting"

	keyParsedTime  = "parsed_time"
	keyMeeting     = "meeting"
	keyFinal       = "final_payload"
	keyChildLast   = "child_last"
	keyChildFinal  = "child_final"
	keyChildFinalK = "child_final_keys"

	helpCmd    = "help"
	samplesCmd = "samples"
	exitCmd    = "exit"
	quitCmd    = "quit"

	includePrefix   = "include "
	includeNone     = "none"
	includeFiltered = "filtered"
	includeAll      = "all"

	scopeAssistant    = assistantID
	inputModeParsed   = "parsed"
	inputModeAll      = "all"
	outputModeCustom  = "custom"
	outputModeDefault = "default"
)

func main() {
	flag.Parse()
	if *modelName == "" {
		*modelName = defaultModelName
	}
	fmt.Printf("🧩 Subgraph Demo (Parent calls Child GraphAgent)\nModel: %s\n", *modelName)
	fmt.Println(strings.Repeat("=", 64))
	if *apiKey == "" && os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Println("💡 Hint: provide -api-key/-base-url or set OPENAI_API_KEY/OPENAI_BASE_URL")
	}
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	_ = "STUB: not implemented"
	// Build child subgraph as a GraphAgent named "assistant".
	return nil
}

// declared tools captured in child; kept to signal tool presence

// Build parent graph that delegates to the child via AddSubgraphNode("assistant").

// Parent GraphAgent with sub-agent registration.

// Runner with in-memory session.

// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)

// Allow overriding parent include_contents at runtime to demonstrate CfgKeyIncludeContents.
// empty = default (all)

// Commands to control parent include_contents: include none|filtered|all

// buildChildSubgraph constructs a child graph (LLM + Tools) and wraps it as a GraphAgent.
// The tool demonstrates reading runtime state injected by the parent (e.g., parsed_time).
func buildChildSubgraph(name string) (*graphagent.GraphAgent, map[string]tool.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Define a callable tool that simulates scheduling and reads parsed time from runtime state if missing.

// Read from runtime state injected by the parent graph.

// Child graph schema and structure (LLM -> Tools (conditional) -> LLM).

// If llm_decider emits tool_calls, go to tools; otherwise finish.

// After tools, route back to LLM to get a natural assistant message summarizing tool results.

// buildParentGraph constructs the parent graph.
// Nodes: parse_input -> assistant(subgraph) -> collect
type parentConfig struct {
	SubIsolate bool
	SubScope   string
	SubInput   string // parsed|all
	SubOutput  string // custom|default
}

func buildParentGraph(cfg parentConfig) (*graph.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Construct options per flags to demonstrate all subgraph features.

// no input mapper (pass full state via default path)

// no custom output mapper; fallback to default last_response + node_responses

// parseInput writes parsed_time and a one-shot system message to shape the next LLM prompt.
func parseInput(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// collect extracts tool JSON result into state["meeting"], then prepares a final payload.
func collect(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	// Extract latest tool.response JSON (search backward until user).
	return *new(any), nil
}

// last_response may be the parent or child; we also recorded child_last via OutputMapper.

// child_final holds the entire subgraph final state snapshot (serializable keys only).

// stream consumes events and prints streaming model/tool details and final outputs.
func stream(ch <-chan *event.Event, verbose bool) error { _ = "STUB: not implemented"; return nil }

// Print streaming deltas.

// Verbose: print tool/model metadata and filter key.

// On completion, show final payload if available.

// Helpers

var (
	reTomorrow = regexp.MustCompile("(?i)tomorrow") // simplistic demo parser
	reToday    = regexp.MustCompile("(?i)today|now")
)

func parseTime(s string) string { _ = "STUB: not implemented"; return "" }

// Fallback: return input verbatim; real parsers can be plugged here.

func printHelp() { _ = "STUB: not implemented"; return }

func printSamples() { _ = "STUB: not implemented"; return }
