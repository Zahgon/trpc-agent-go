//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package llm provides base helpers for LLM-backed evaluators.
package llm

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/invocationsaggregator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/messagesconstructor"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/responsescorer"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/samplesaggregator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// LLMEvaluator defines the LLM-backed evaluator contract.
type LLMEvaluator interface {
	evaluator.Evaluator
	messagesconstructor.MessagesConstructor
	responsescorer.ResponseScorer
	samplesaggregator.SamplesAggregator
	invocationsaggregator.InvocationsAggregator
}

// LLMBaseEvaluator hosts shared orchestration logic for LLM evaluators.
type LLMBaseEvaluator struct {
	LLMEvaluator LLMEvaluator // LLMEvaluator is the concrete LLM evaluator implementation.
}

// New constructs an LLMBaseEvaluator wrapper around the concrete evaluator.
func New(llmEvaluator LLMEvaluator) LLMEvaluator {
	_ = "STUB: not implemented"
	return *new(LLMEvaluator)
}

// Name returns the evaluator name.
func (r *LLMBaseEvaluator) Name() string { _ = "STUB: not implemented"; return "" }

// Description describes the evaluator.
func (r *LLMBaseEvaluator) Description() string { _ = "STUB: not implemented"; return "" }

// Evaluate runs the judge model over paired invocations and aggregates results.
func (r *LLMBaseEvaluator) Evaluate(ctx context.Context, actuals, expecteds []*evalset.Invocation,
	evalMetric *metric.EvalMetric) (*evaluator.EvaluateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AggregateInvocations delegates invocation aggregation to the concrete evaluator.
func (r *LLMBaseEvaluator) AggregateInvocations(ctx context.Context, results []*evaluator.PerInvocationResult,
	evalMetric *metric.EvalMetric) (*evaluator.EvaluateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AggregateSamples delegates sample aggregation to the concrete evaluator.
func (r *LLMBaseEvaluator) AggregateSamples(ctx context.Context, samples []*evaluator.PerInvocationResult,
	evalMetric *metric.EvalMetric) (*evaluator.PerInvocationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ScoreBasedOnResponse delegates response scoring to the concrete evaluator.
func (r *LLMBaseEvaluator) ScoreBasedOnResponse(ctx context.Context, resp *model.Response,
	evalMetric *metric.EvalMetric) (*evaluator.ScoreResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConstructMessages delegates prompt construction to the concrete evaluator.
func (r *LLMBaseEvaluator) ConstructMessages(ctx context.Context, actuals, expecteds []*evalset.Invocation,
	evalMetric *metric.EvalMetric) ([]model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *LLMBaseEvaluator) resolveStructuredOutput(ctx context.Context,
	actuals, expecteds []*evalset.Invocation, evalMetric *metric.EvalMetric) (*model.StructuredOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
