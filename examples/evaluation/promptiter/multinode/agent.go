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
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	sportsRecapAgentName  = "promptiter-sports-recap-agent"
	headlineAgentName     = "headline_agent"
	highlightsAgentName   = "highlights_agent"
	statsAngleAgentName   = "stats_angle_agent"
	recapWriterAgentName  = "recap_writer"
	sportsEditorAgentName = "sports_editor"
	nodePrepareGameInput  = "prepare_game_input"
	nodeJoinRecapParts    = "join_recap_parts"
	keyHeadlineInput      = "headline_input"
	keyHighlightsInput    = "highlights_input"
	keyStatsAngleInput    = "stats_angle_input"
	keyHeadline           = "headline"
	keyHighlights         = "highlights"
	keyStatsAngle         = "stats_angle"
	keyWriterInput        = "writer_input"
	keyEditorInput        = "editor_input"
	keyFinalRecap         = "final_recap"
)

func newSportsRecapAgent(m model.Model) (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func newStageAgent(name string, instruction string, m model.Model, cfg model.GenerationConfig) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func buildSportsRecapGraph() (*graph.Graph, error) { _ = "STUB: not implemented"; return nil, nil }

func branchOptions(inputKey string, outputKey string) []graph.Option {
	_ = "STUB: not implemented"
	return nil
}

func prepareGameInput(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func joinRecapParts(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func storeDraft(parent graph.State, result graph.SubgraphResult) graph.State {
	_ = "STUB: not implemented"
	return *new(graph.State)
}

func storeFinal(parent graph.State, result graph.SubgraphResult) graph.State {
	_ = "STUB: not implemented"
	return *new(graph.State)
}

func intPtr(value int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(value float64) *float64 { _ = "STUB: not implemented"; return nil }

const headlineInstruction = `生成标题。`

const highlightsInstruction = `提取比赛高光。`

const statsAngleInstruction = `选择数据角度。`

const recapWriterInstruction = `生成中文战报。`

const sportsEditorInstruction = `润色中文战报。`
