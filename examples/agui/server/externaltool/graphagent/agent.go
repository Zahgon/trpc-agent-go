//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

func newGraphAgent(g *graph.Graph) (*graphagent.GraphAgent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newGenerationConfig() model.GenerationConfig {
	_ = "STUB: not implemented"
	return *new(model.GenerationConfig)
}

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
