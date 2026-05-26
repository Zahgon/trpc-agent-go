//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package template evaluates prompt-defined LLM judge metrics.
package template

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/messagesconstructor"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	metricllm "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/llm"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// EvaluatorName is the registered evaluator name for template-based LLM judging.
const EvaluatorName = "llm_judge_template"

type templateEvaluator struct {
	llmBaseEvaluator    llm.LLMEvaluator
	messagesConstructor messagesconstructor.MessagesConstructor
}

// New returns the template evaluator.
func New() evaluator.Evaluator { _ = "STUB: not implemented"; return *new(evaluator.Evaluator) }

// Name returns the evaluator name.
func (e *templateEvaluator) Name() string { _ = "STUB: not implemented"; return "" }

// Description returns the evaluator description.
func (e *templateEvaluator) Description() string { _ = "STUB: not implemented"; return "" }

// Evaluate runs template-based LLM evaluation.
func (e *templateEvaluator) Evaluate(ctx context.Context, actuals, expecteds []*evalset.Invocation,
	evalMetric *metric.EvalMetric) (*evaluator.EvaluateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConstructMessages builds judge prompts from template configuration.
func (e *templateEvaluator) ConstructMessages(ctx context.Context, actuals, expecteds []*evalset.Invocation,
	evalMetric *metric.EvalMetric) ([]model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StructuredOutput delegates structured output schema construction to the prompt builder.
func (e *templateEvaluator) StructuredOutput(ctx context.Context, actuals, expecteds []*evalset.Invocation,
	evalMetric *metric.EvalMetric) (*model.StructuredOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ScoreBasedOnResponse scores the judge response with the configured scorer.
func (e *templateEvaluator) ScoreBasedOnResponse(ctx context.Context, response *model.Response,
	evalMetric *metric.EvalMetric) (*evaluator.ScoreResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AggregateSamples aggregates samples with the configured strategy.
func (e *templateEvaluator) AggregateSamples(ctx context.Context, samples []*evaluator.PerInvocationResult,
	evalMetric *metric.EvalMetric) (*evaluator.PerInvocationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AggregateInvocations aggregates invocation results with the configured strategy.
func (e *templateEvaluator) AggregateInvocations(ctx context.Context, results []*evaluator.PerInvocationResult,
	evalMetric *metric.EvalMetric) (*evaluator.EvaluateResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func responseScorerName(evalMetric *metric.EvalMetric) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func sampleAggregatorName(templateOptions *metricllm.JudgeTemplateOptions) string {
	_ = "STUB: not implemented"
	return ""
}

func invocationAggregatorName(templateOptions *metricllm.JudgeTemplateOptions) string {
	_ = "STUB: not implemented"
	return ""
}

func judgeTemplateOptions(evalMetric *metric.EvalMetric) (*metricllm.JudgeTemplateOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
