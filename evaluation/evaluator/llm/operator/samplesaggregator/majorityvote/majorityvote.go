//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package majorityvote picks a representative sample based on vote counts.
package majorityvote

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/samplesaggregator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
)

type majorityVoteSamplesAggregator struct {
}

// New returns a samples aggregator that selects a representative by majority vote.
func New() samplesaggregator.SamplesAggregator {
	_ = "STUB: not implemented"
	return *new(samplesaggregator.SamplesAggregator)
}

// AggregateSamples resolves multiple judge samples to one invocation result, preferring the majority status.
func (s *majorityVoteSamplesAggregator) AggregateSamples(ctx context.Context, samples []*evaluator.PerInvocationResult,
	evalMetric *metric.EvalMetric) (*evaluator.PerInvocationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// On tie or majority fail, return a representative failing sample to preserve conservative scoring.
