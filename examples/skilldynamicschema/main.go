//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a Skill-driven pattern for dynamically switching
// structured output JSON schema at runtime:
//
//   - The model chooses a Skill based on user input.
//   - The model loads the Skill content, extracts a JSON Schema, and calls a
//     user-defined set_output_schema tool to update invocation.StructuredOutput.
//   - Subsequent model calls in the same invocation are constrained by that schema.
//   - The final JSON is extracted into event.StructuredOutput (untyped map/slice/etc).
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

var (
	flagModel      = flag.String("model", "deepseek-v4-flash", "model name (OpenAI-compatible)")
	flagStreaming  = flag.Bool("streaming", false, "stream responses")
	flagTraceTools = flag.Bool("trace_tools", true, "print tool call/response trace")
)

const (
	appName          = "skill-dynamic-schema-demo"
	defaultSkillsDir = "skills"

	// OutputKey is only used to make sure OutputResponseProcessor is installed
	// (so we can read event.StructuredOutput without changing the framework).
	outputKey = "__so_tmp__"
)

const instructionText = `
You are a tool-using agent.

You have two skills available: plan_route and recommend_poi.

For every user request:
- Choose exactly one skill to run.
  - If the user explicitly mentions "plan_route" or "recommend_poi", use that one.
  - Otherwise, use plan_route for route/ETA/distance requests, and recommend_poi for POI/city/recommendation requests.
- Perform the following steps using tool calls only (no assistant content):
  1) Call skill_load for the chosen skill.
  2) From the loaded skill content, find the JSON schema under the section "Output JSON Schema".
  3) Call set_output_schema with {"schema": <that JSON schema object>}.
     - Do NOT call set_output_schema without "schema".
     - If set_output_schema returns {"ok": false, ...}, extract the schema again and retry.
     - Do not proceed to step 4 until set_output_schema returns {"ok": true, ...}.
  4) Call skill_run with {"skill":"<chosen skill>","command":"cat result.json"}.
- Finally, return ONLY the JSON object from step 4 stdout. No extra text.

Do not output analysis.
`

func main() {
	flag.Parse()

	fmt.Println("Skill-Driven Dynamic Structured Output (JSON Schema)")
	fmt.Printf("Model: %s\n", *flagModel)
	fmt.Printf("Streaming: %t\n", *flagStreaming)
	fmt.Printf("Trace tools: %t\n", *flagTraceTools)
	fmt.Println("Type 'exit' to quit.")
	fmt.Println(strings.Repeat("=", 60))

	if err := run(); err != nil {
		fmt.Printf("run failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

// This demo drives skill_run, which lives in the full skill
// tool profile. The default knowledge_only profile omits it.

// Key point: install OutputResponseProcessor without a static schema.

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }

func temperatureForModel(name string) *float64 { _ = "STUB: not implemented"; return nil }

// gpt-5 only supports the default temperature; omit it to avoid 400s.

func printToolCalls(
	toolCalls []model.ToolCall,
	toolNameByID map[string]string,
	seen map[string]struct{},
) {
	_ = "STUB: not implemented"
	return
}

func printToolResult(
	msg model.Message,
	toolNameByID map[string]string,
	seen map[string]struct{},
) {
	_ = "STUB: not implemented"
	return
}

func formatToolResult(content string) string { _ = "STUB: not implemented"; return "" }

func formatInlineJSON(b []byte) string { _ = "STUB: not implemented"; return "" }

func toolIcon(name string) string { _ = "STUB: not implemented"; return "" }
