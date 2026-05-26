//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"trpc.group/trpc-go/trpc-agent-go/agent"
)

const (
	// rougeAgentName is the agent name used by the ROUGE evaluation example.
	rougeAgentName = "rouge-agent"
	// rougeAgentInstruction is used to keep the output short and stable for ROUGE matching.
	rougeAgentInstruction = "Answer in exactly one short sentence of plain text. Do not use markdown, lists, or code formatting. Output only the answer."
)

// newRougeAgent creates an LLM-based agent for the ROUGE evaluation example.
func newRougeAgent(modelName string, stream bool) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

// intPtr returns a pointer to v.
func intPtr(v int) *int {
	_ = "STUB: not implemented"

	// floatPtr returns a pointer to v.
	return nil
}

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
