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

func newCandidateTravelAgent(modelName string, stream bool) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newReferenceTravelAgent(modelName string, reasoningEffort string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newSimulatorAgent(modelName string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newJudgeAgent(modelName string, reasoningEffort string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newGenerationConfig(maxTokens int, temperature float64, stream bool, reasoningEffort string) model.GenerationConfig {
	_ = "STUB: not implemented"
	return *new(model.GenerationConfig)
}

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }

func stringPtr(v string) *string { _ = "STUB: not implemented"; return nil }
