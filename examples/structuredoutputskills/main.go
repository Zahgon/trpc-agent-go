//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates using structured output together with Agent
// Skills. The agent is free to call tools first, and only the final answer
// must be a single JSON object matching the schema.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	flagModel     = flag.String("model", "deepseek-v4-flash", "model name")
	flagStreaming = flag.Bool("streaming", true, "stream responses")
)

const (
	appName          = "structuredoutput-skills-demo"
	defaultSkillsDir = "skills"
)

const instructionText = `
You can use Agent Skills (tools) and must produce a typed JSON result.

Rules:
- You MAY call tools when needed.
- While calling tools, do not provide a user-facing answer.
- For every user request, do the following:
  1) Call skill_load for the "hello" skill.
  2) Call skill_run to run: bash scripts/hello.sh
  3) Return the final answer as JSON matching the schema.
`

type helloResult struct {
	Skill  string `json:"skill"`
	Output string `json:"output"`
}

func main() {
	flag.Parse()

	fmt.Println("Structured Output + Skills (Typed JSON)")
	fmt.Printf("Model: %s\n", *flagModel)
	fmt.Printf("Streaming: %t\n", *flagStreaming)
	fmt.Println("Type 'exit' to quit.")
	fmt.Println(strings.Repeat("=", 50))

	if err := run(); err != nil {
		fmt.Printf("run failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

// Model (OpenAI-compatible).

// Skills repository (defaults to ./skills in this directory).

// Local workspace executor for skill_run.

// The demo calls skill_run to execute scripts/hello.sh, which
// requires the full skill tool profile; the default profile is
// knowledge_only and does not register skill_run.

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
