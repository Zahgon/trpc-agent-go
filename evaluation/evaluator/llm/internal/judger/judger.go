//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package judger runs judge requests for LLM-based evaluators.
package judger

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	criterionllm "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/llm"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

// Judge runs the configured judge and returns its final response.
func Judge(ctx context.Context, messages []model.Message, evalMetric *metric.EvalMetric, opt ...Option) (*model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func judgeWithModel(ctx context.Context, judgeModel *criterionllm.JudgeModelOptions,
	messages []model.Message, opts *options) (*model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func judgeWithRunner(ctx context.Context, judgeRunner runner.Runner, messages []model.Message, opts *options) (*model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildRunnerOptions(opts *options) ([]agent.RunOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func materializeStructuredOutputContent(resp *model.Response, payload any, opts *options) error {
	_ = "STUB: not implemented"
	return nil
}
