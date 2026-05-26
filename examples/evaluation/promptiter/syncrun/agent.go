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
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	candidateAgentName = "candidate"
)

const defaultCandidateInstruction = "Write one Chinese sentence that summarizes the JSON input. Output only the text."

func newCandidateAgent(m model.Model, instruction string) (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func newBackwarderAgent(m model.Model) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newAggregatorAgent(m model.Model) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newOptimizerAgent(m model.Model) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newPromptIterStageAgent(name string, description string, m model.Model) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newJudgeAgent(m model.Model) agent.Agent { _ = "STUB: not implemented"; return *new(agent.Agent) }

func intPtr(value int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(value float64) *float64 { _ = "STUB: not implemented"; return nil }
