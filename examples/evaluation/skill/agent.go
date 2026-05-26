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
	"trpc.group/trpc-go/trpc-agent-go/skill"
)

func newSkillAgent(modelName string, stream bool, repo skill.Repository) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

// This evaluation asserts skill_run behavior, so opt into the
// full skill tool profile; the default knowledge_only profile
// does not register skill_run.

func intPtr(v int) *int           { _ = "STUB: not implemented"; return nil }
func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
