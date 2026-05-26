//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package tooltrajectory defines tool trajectory comparison criteria.
package tooltrajectory

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	criterionjson "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/json"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/text"
)

// CompareFunc defines custom tool trajectory comparison logic.
type CompareFunc func(actual, expected *evalset.Invocation) (bool, error)

// New creates a ToolTrajectoryCriterion with the provided options.
func New(opt ...Option) *ToolTrajectoryCriterion { _ = "STUB: not implemented"; return nil }

// ToolTrajectoryCriterion provides comparison rules for tool call and response sequences.
type ToolTrajectoryCriterion struct {
	// DefaultStrategy applies when no tool-specific strategy is provided.
	DefaultStrategy *ToolTrajectoryStrategy `json:"defaultStrategy,omitempty"`
	// ToolStrategy holds per-tool strategies keyed by tool name.
	ToolStrategy map[string]*ToolTrajectoryStrategy `json:"toolStrategy,omitempty"`
	// OrderSensitive requires tools to match in sequence when true; when false, matching is order-agnostic.
	OrderSensitive bool `json:"orderSensitive,omitempty"`
	// SubsetMatching allows expected tool list to be a subset of actual list.
	SubsetMatching bool `json:"subsetMatching,omitempty"`
	// CompareName selects a registered comparison implementation by name.
	CompareName string `json:"compareName,omitempty"`
	// Compare allows custom comparison override.
	Compare CompareFunc `json:"-"`
}

// ToolTrajectoryStrategy defines comparison strategies for a single tool.
type ToolTrajectoryStrategy struct {
	Name      *text.TextCriterion          `json:"name,omitempty"`      // Name compares tool names.
	Arguments *criterionjson.JSONCriterion `json:"arguments,omitempty"` // Arguments compares tool call arguments.
	Result    *criterionjson.JSONCriterion `json:"result,omitempty"`    // Result compares tool call results.
}

// Match compares actual and expected invocations according to tool trajectory rules.
func (t *ToolTrajectoryCriterion) Match(actual, expected *evalset.Invocation) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// validateToolCounts validates the tool counts of actual and expected invocations.
func (t *ToolTrajectoryCriterion) validateToolCounts(actual, expected *evalset.Invocation) error {
	_ = "STUB: not implemented"
	return nil
}

// orderedMatch matches actual and expected tool calls in order.
func (t *ToolTrajectoryCriterion) orderedMatch(actual, expected []*evalset.Tool) error {
	_ = "STUB: not implemented"
	return nil
}

// unorderedMatch matches actual and expected tool calls in no order.
func (t *ToolTrajectoryCriterion) unorderedMatch(actual, expected []*evalset.Tool) error {
	_ = "STUB: not implemented"
	return nil
}

// matchTool matches a single tool call and response.
func (t *ToolTrajectoryCriterion) matchTool(actualTool, expectedTool *evalset.Tool) error {
	_ = "STUB: not implemented"
	return nil
}

// getStrategy picks the comparison strategy for a specific tool pair.
func (t *ToolTrajectoryCriterion) getStrategy(actualTool, expectedTool *evalset.Tool) *ToolTrajectoryStrategy {
	_ = "STUB: not implemented"
	return nil
}

// Match compares a single tool call against the expected tool call using the configured strategies.
func (t *ToolTrajectoryStrategy) Match(actual, expected *evalset.Tool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
