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

func newReviewerAgent(modelInstance model.Model) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newMainAgent(modelInstance model.Model, streaming bool) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
