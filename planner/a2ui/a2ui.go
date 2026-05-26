// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package a2ui provides an A2UI-specific planner.
package a2ui

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/planner"
)

// Verify that a2uiPlanner implements the planner interface.
var _ planner.Planner = (*a2uiPlanner)(nil)

// a2uiPlanner implements the planner interface for A2UI.
type a2uiPlanner struct {
	instruction                       string
	clientCapabilitiesSchema          string
	catalogDescriptionSchema          string
	clientToServer                    string
	serverToClient                    string
	serverToClientWithStandardCatalog string
	standardCatalogDefinition         string
}

// New creates a new A2UI planner.
func New(opts ...Option) planner.Planner { _ = "STUB: not implemented"; return *new(planner.Planner) }

// BuildPlanningInstruction injects A2UI protocol constraints.
func (p *a2uiPlanner) BuildPlanningInstruction(ctx context.Context, invocation *agent.Invocation,
	llmRequest *model.Request) string {
	_ = "STUB: not implemented"
	return ""
}

// ProcessPlanningResponse returns nil to indicate that no planning-specific response processing is needed.
func (p *a2uiPlanner) ProcessPlanningResponse(ctx context.Context, invocation *agent.Invocation,
	response *model.Response) *model.Response {
	_ = "STUB: not implemented"
	return nil
}
