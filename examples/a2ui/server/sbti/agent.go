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
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

const (
	graphAgentName         = "sbti-a2ui-graph"
	directorAgentName      = "sbti_director"
	rendererAgentName      = "sbti_a2ui_renderer"
	directorStateOutputKey = "sbti_director_state"
	platformOutputTextKey  = "{{input.output_text}}"
)

func newAgent() (agent.Agent, error) { _ = "STUB: not implemented"; return *new(agent.Agent), nil }

func buildGraph() (*graph.Graph, error) { _ = "STUB: not implemented"; return nil, nil }

func buildDirectorAgent(schema map[string]any) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func localDirectorInstructionText() string { _ = "STUB: not implemented"; return "" }

func buildRendererAgent() agent.Agent { _ = "STUB: not implemented"; return *new(agent.Agent) }

func localRendererInstructionText() string { _ = "STUB: not implemented"; return "" }

func directorOutputSchemaMap() (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }
