//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package processor

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/planner"
)

// PlanningRequestProcessor implements planning request processing logic.
type PlanningRequestProcessor struct {
	// Planner is the planner to use for generating planning instructions.
	Planner planner.Planner
}

// NewPlanningRequestProcessor creates a new planning request processor.
func NewPlanningRequestProcessor(p planner.Planner) *PlanningRequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

// ProcessRequest implements the flow.RequestProcessor interface.
// It generates planning instructions and removes thought markers from requests.
func (p *PlanningRequestProcessor) ProcessRequest(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// Apply thinking configuration for built-in planners.

// For built-in planners, just apply thinking config and return.

// Generate planning instruction.

// Check if planning instruction already exists to avoid duplication.

// hasSystemMessage checks if a system message with the given content already exists.
// It compares only the first few characters of the content for performance reasons,
// as this is usually sufficient to determine content similarity.
func hasSystemMessage(messages []model.Message, content string) bool {
	_ = "STUB: not implemented"
	// Maximum length of content prefix to compare for performance optimization.
	return false
}

// Use content prefix for comparison to avoid performance issues with long content.

// PlanningResponseProcessor implements planning response processing logic.
type PlanningResponseProcessor struct {
	// Planner is the planner to use for processing planning responses.
	Planner planner.Planner
}

// NewPlanningResponseProcessor creates a new planning response processor.
func NewPlanningResponseProcessor(p planner.Planner) *PlanningResponseProcessor {
	_ = "STUB: not implemented"
	return nil
}

// ProcessResponse implements the flow.ResponseProcessor interface.
// It processes planning responses using the configured planner.
func (p *PlanningResponseProcessor) ProcessResponse(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
	rsp *model.Response,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// Process the response using the planner.

// Update the original response with processed content.

// Mark as partial response so it doesn't interfere with full response detection.
