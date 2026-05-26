//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package finalresponse provides deterministic evaluation for agent final responses.
package finalresponse

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	cfinalresponse "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/finalresponse"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/status"
)

// finalResponseEvaluator evaluates final responses using deterministic matching criteria.
type finalResponseEvaluator struct {
}

// New creates a new final response evaluator.
func New() evaluator.Evaluator { _ = "STUB: not implemented"; return *new(evaluator.Evaluator) }

// Name returns the evaluator identifier.
func (e *finalResponseEvaluator) Name() string { _ = "STUB: not implemented"; return "" }

// Description describes the evaluator purpose.
func (e *finalResponseEvaluator) Description() string { _ = "STUB: not implemented"; return "" }

// Evaluate compares final responses between actual and expected invocations.
func (e *finalResponseEvaluator) Evaluate(ctx context.Context, actuals, expecteds []*evalset.Invocation,
	evalMetric *metric.EvalMetric) (*evaluator.EvaluateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// statusForScore maps a numeric score to an evaluation status based on the metric threshold.
func (e *finalResponseEvaluator) statusForScore(score float64, evalMetric *metric.EvalMetric) status.EvalStatus {
	_ = "STUB: not implemented"
	return *new(status.EvalStatus)
}

// finalResponsesMatch performs deterministic matching for the configured final response criterion.
func finalResponsesMatch(ctx context.Context, actual, expected *evalset.Invocation,
	criterion *cfinalresponse.FinalResponseCriterion) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
