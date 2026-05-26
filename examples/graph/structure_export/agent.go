//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

const (
	nodeStart   = "start"
	nodePrepare = "prepare"
	nodeRoute   = "route"
	nodeTools   = "tools"
	nodeBranchA = "branch_a"
	nodeBranchB = "branch_b"
	nodeJoin    = "join"
	nodeDone    = "done"
)

func buildAgent() (*graphagent.GraphAgent, error) { _ = "STUB: not implemented"; return nil, nil }

func buildGraph() (*graph.Graph, error) { _ = "STUB: not implemented"; return nil, nil }

type searchDocsInput struct {
	Query string `json:"query"`
}

type searchDocsOutput struct {
	Results []string `json:"results"`
}

type summarizeInput struct {
	Notes []string `json:"notes"`
}

type summarizeOutput struct {
	Summary string `json:"summary"`
}
