//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package template assembles judge prompts from template configuration.
package template

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/messagesconstructor"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	metricllm "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/llm"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/prompt"
)

type templateMessagesConstructor struct {
}

// New returns a messages constructor for template prompts.
func New() messagesconstructor.MessagesConstructor {
	_ = "STUB: not implemented"
	return *new(messagesconstructor.MessagesConstructor)
}

// ConstructMessages renders the configured judge template into a user message.
func (c *templateMessagesConstructor) ConstructMessages(ctx context.Context, actuals, expecteds []*evalset.Invocation,
	evalMetric *metric.EvalMetric) ([]model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StructuredOutput returns the structured output schema for the configured response scorer.
func (c *templateMessagesConstructor) StructuredOutput(ctx context.Context, actuals, expecteds []*evalset.Invocation,
	evalMetric *metric.EvalMetric) (*model.StructuredOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func judgeTemplateOptions(evalMetric *metric.EvalMetric) (*metricllm.JudgeTemplateOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveTemplateValues(actuals, expecteds []*evalset.Invocation,
	bindings []*metricllm.TemplateVariableBinding) (prompt.Vars, error) {
	_ = "STUB: not implemented"
	return *new(prompt.Vars), nil
}

func resolveBindingValue(actuals, expecteds []*evalset.Invocation,
	source *metricllm.TemplateVariableSource) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func resolveActualValue(actuals []*evalset.Invocation, field metricllm.TemplateVariableField) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func resolveExpectedValue(expecteds []*evalset.Invocation, field metricllm.TemplateVariableField) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
