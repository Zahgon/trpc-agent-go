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
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// createResearchAgent creates a specialized research agent.
func (c *transferChat) createResearchAgent(modelInstance model.Model) agent.Agent {
	_ = "STUB: not implemented"
	// Search tool.
	return *new(agent.Agent)
}

// search performs information search.
func (c *transferChat) search(_ context.Context, args searchArgs) (searchResult, error) {
	_ = "STUB: not implemented"
	// Simulate search results based on query.
	return *new(searchResult), nil
}

// Data structures for search tool.
type searchArgs struct {
	Query string `json:"query" jsonschema:"description=The search query,required"`
}

type searchResult struct {
	Query   string   `json:"query"`
	Results []string `json:"results"`
	Count   int      `json:"count"`
}
