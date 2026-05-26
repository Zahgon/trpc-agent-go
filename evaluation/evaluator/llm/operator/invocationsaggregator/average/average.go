//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package average aggregates invocation results using arithmetic mean.
package average

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/invocationsaggregator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
)

type averageInvocationsAggregator struct {
}

// New returns an invocations aggregator that averages evaluated scores.
func New() invocationsaggregator.InvocationsAggregator {
	_ = "STUB: not implemented"
	return *new(invocationsaggregator.InvocationsAggregator)
}

// AggregateInvocations summarizes per-invocation results into an overall score while skipping not-evaluated entries.
func (a *averageInvocationsAggregator) AggregateInvocations(ctx context.Context,
	results []*evaluator.PerInvocationResult, evalMetric *metric.EvalMetric) (*evaluator.EvaluateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
