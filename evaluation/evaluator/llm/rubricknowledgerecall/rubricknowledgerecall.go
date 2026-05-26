//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package rubricknowledgerecall evaluates knowledge recall using LLM judges.
package rubricknowledgerecall

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/invocationsaggregator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/messagesconstructor"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/responsescorer"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/samplesaggregator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

type rubricKnowledgeRecallEvaluator struct {
	llmBaseEvaluator      llm.LLMEvaluator
	messagesConstructor   messagesconstructor.MessagesConstructor
	responsescorer        responsescorer.ResponseScorer
	samplesAggregator     samplesaggregator.SamplesAggregator
	invocationsAggregator invocationsaggregator.InvocationsAggregator
}

// New builds the rubric knowledge recall evaluator.
func New(opt ...Option) evaluator.Evaluator {
	_ = "STUB: not implemented"
	return *new(evaluator.Evaluator)
}

// Name returns the name of the evaluator.
func (e *rubricKnowledgeRecallEvaluator) Name() string { _ = "STUB: not implemented"; return "" }

// Description returns the description of the evaluator.
func (e *rubricKnowledgeRecallEvaluator) Description() string { _ = "STUB: not implemented"; return "" }

// Evaluate evaluates the knowledge recall of the agent.
func (e *rubricKnowledgeRecallEvaluator) Evaluate(ctx context.Context, actuals, expecteds []*evalset.Invocation,
	evalMetric *metric.EvalMetric) (*evaluator.EvaluateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConstructMessages constructs the messages for the evaluator.
func (e *rubricKnowledgeRecallEvaluator) ConstructMessages(ctx context.Context, actuals, expecteds []*evalset.Invocation,
	evalMetric *metric.EvalMetric) ([]model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StructuredOutput delegates structured output schema construction to the prompt builder.
func (e *rubricKnowledgeRecallEvaluator) StructuredOutput(ctx context.Context, actuals,
	expecteds []*evalset.Invocation, evalMetric *metric.EvalMetric) (*model.StructuredOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ScoreBasedOnResponse scores the response of the evaluator.
func (e *rubricKnowledgeRecallEvaluator) ScoreBasedOnResponse(ctx context.Context, response *model.Response,
	evalMetric *metric.EvalMetric) (*evaluator.ScoreResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AggregateSamples aggregates the samples of the evaluator.
func (e *rubricKnowledgeRecallEvaluator) AggregateSamples(ctx context.Context, samples []*evaluator.PerInvocationResult,
	evalMetric *metric.EvalMetric) (*evaluator.PerInvocationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AggregateInvocations aggregates the invocations of the evaluator.
func (e *rubricKnowledgeRecallEvaluator) AggregateInvocations(ctx context.Context, results []*evaluator.PerInvocationResult,
	evalMetric *metric.EvalMetric) (*evaluator.EvaluateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
