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

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/status"
)

const commentaryLengthMetricName = "final_response_length_compliance"

const (
	commentaryPreferredMinLength  = 32
	commentaryPreferredMaxLength  = 58
	commentaryAcceptableMinLength = 18
	commentaryAcceptableMaxLength = 72
)

type commentaryLengthEvaluator struct{}

func newCommentaryLengthEvaluator() evaluator.Evaluator {
	_ = "STUB: not implemented"
	return *new(evaluator.Evaluator)
}

func (e *commentaryLengthEvaluator) Name() string { _ = "STUB: not implemented"; return "" }

func (e *commentaryLengthEvaluator) Description() string { _ = "STUB: not implemented"; return "" }

func (e *commentaryLengthEvaluator) Evaluate(
	_ context.Context,
	actuals, expecteds []*evalset.Invocation,
	evalMetric *metric.EvalMetric,
) (*evaluator.EvaluateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func commentaryFinalResponseText(invocation *evalset.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}

func commentaryLengthScore(actualLength int) (float64, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func commentaryLengthReason(actualLength int, band string, direction string, deltaToPreferred int) string {
	_ = "STUB: not implemented"
	return ""
}

func commentaryLengthDirection(actualLength int) (string, int) {
	_ = "STUB: not implemented"
	return "", 0
}

func commentaryStatusForScore(score float64, evalMetric *metric.EvalMetric) status.EvalStatus {
	_ = "STUB: not implemented"
	return *new(status.EvalStatus)
}
