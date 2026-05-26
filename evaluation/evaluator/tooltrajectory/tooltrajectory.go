//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package tooltrajectory provides tool trajectory-based evaluation.
package tooltrajectory

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	ctooltrajectory "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/tooltrajectory"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/status"
)

// toolTrajectoryEvaluator is a tool trajectory evaluator implementation for evaluator.
type toolTrajectoryEvaluator struct {
}

// New creates a new trajectory evaluator.
func New() evaluator.Evaluator { _ = "STUB: not implemented"; return *new(evaluator.Evaluator) }

// Name returns the name of this evaluator.
func (e *toolTrajectoryEvaluator) Name() string { _ = "STUB: not implemented"; return "" }

// Description returns a description of what this evaluator does.
func (e *toolTrajectoryEvaluator) Description() string { _ = "STUB: not implemented"; return "" }

// Evaluate compares tool usage trajectories between actual and expected invocations.
func (e *toolTrajectoryEvaluator) Evaluate(ctx context.Context, actuals, expecteds []*evalset.Invocation,
	evalMetric *metric.EvalMetric) (*evaluator.EvaluateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *toolTrajectoryEvaluator) statusForScore(score float64, evalMetric *metric.EvalMetric) status.EvalStatus {
	_ = "STUB: not implemented"
	return *new(status.EvalStatus)
}

func toolCallsMatch(actual, expected *evalset.Invocation,
	criterion *ctooltrajectory.ToolTrajectoryCriterion) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
