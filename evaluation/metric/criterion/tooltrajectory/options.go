//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package tooltrajectory

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/json"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/text"
)

// defaultToolTrajectoryStrategy is used when no user strategy is supplied.
var (
	defaultJsonCriterion          = json.New()
	defaultToolTrajectoryStrategy = &ToolTrajectoryStrategy{
		Name:      &text.TextCriterion{MatchStrategy: text.TextMatchStrategyExact},
		Arguments: defaultJsonCriterion,
		Result:    defaultJsonCriterion,
	}
)

// options configures ToolTrajectoryCriterion.
type options struct {
	// defaultStrategy sets the fallback strategy when no tool-specific strategy is defined.
	defaultStrategy *ToolTrajectoryStrategy
	// toolStrategy configures per-tool strategies keyed by tool name.
	toolStrategy map[string]*ToolTrajectoryStrategy
	// orderSensitive enforces ordered matching when true; when false, tools can match out of order.
	orderSensitive bool
	// subsetMatching allows expected tool list to be a subset of actual list.
	subsetMatching bool
	// compareName selects a registered comparison implementation by name.
	compareName string
	// compare allows overriding comparison logic entirely.
	compare CompareFunc
}

// newOptions applies provided options for ToolTrajectoryCriterion.
func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// Option is a function that configures ToolTrajectoryCriterion.
type Option func(*options)

// WithDefault sets the default tool trajectory strategy.
func WithDefault(defaultStrategy *ToolTrajectoryStrategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTool sets the per-tool strategies keyed by tool name.
func WithTool(tool map[string]*ToolTrajectoryStrategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithOrderSensitive controls whether tool matching must follow sequence order.
func WithOrderSensitive(orderSensitive bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSubsetMatching allows expected tool list to be a subset of actual list.
func WithSubsetMatching(subsetMatching bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCompareName sets the name of the registered compare function.
func WithCompareName(compareName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCompare sets the tool trajectory comparison logic.
func WithCompare(compare CompareFunc) Option { _ = "STUB: not implemented"; return *new(Option) }
