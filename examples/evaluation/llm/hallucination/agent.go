//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	productCatalogToolName = "product_catalog_lookup"
	defaultAgentName       = "hallucination-agent"
	judgeAgentName         = "hallucination-judge"
	productIDAuroraPadX2   = "aurora-pad-x2"
	productIDTerraWatchS   = "terra-watch-s"
	scriptedCandidateModel = "scripted-hallucination"
	productCatalogToolCall = "call_product_catalog_lookup"
	hallucinatedAnswer     = "AuroraPad X2 was released in 2025. It offers 24 hours of battery life. It is designed for classroom teachers."
)

func newQAAgent(modelName string, stream bool) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newJudgeAgent(modelName string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func newForcedHallucinationAgent() agent.Agent { _ = "STUB: not implemented"; return *new(agent.Agent) }

func newProductCatalogTool() tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

type productCatalogArgs struct {
	ProductID string `json:"product_id"`
}

type productCatalogResult struct {
	ProductID        string `json:"product_id"`
	Name             string `json:"name"`
	ReleaseYear      int    `json:"release_year"`
	BatteryLifeHours int    `json:"battery_life_hours"`
	MarketSegment    string `json:"market_segment"`
	Connectivity     string `json:"connectivity"`
}

type scriptedHallucinationAgent struct {
	tools []tool.Tool
}

func (a *scriptedHallucinationAgent) Run(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *scriptedHallucinationAgent) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

func (a *scriptedHallucinationAgent) Info() agent.Info {
	_ = "STUB: not implemented"
	return *new(agent.Info)
}

func (a *scriptedHallucinationAgent) SubAgents() []agent.Agent {
	_ = "STUB: not implemented"
	return nil
}

func (a *scriptedHallucinationAgent) FindSubAgent(string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func lookupProductCatalog(_ context.Context, args productCatalogArgs) (productCatalogResult, error) {
	_ = "STUB: not implemented"
	return *new(productCatalogResult), nil
}

func productCatalogInputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
