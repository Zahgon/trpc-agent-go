//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main runs a local llm_judge_template evaluation example.
package main

import (
	"context"
	"flag"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/evaluation"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

type runOptions struct {
	DataDir   string
	OutputDir string
	ModelName string
	EvalSetID string
	Streaming bool
}

type exampleEvaluator interface {
	Evaluate(ctx context.Context, evalSetID string, opt ...evaluation.Option) (*evaluation.EvaluationResult, error)
	Close() error
}

type evaluatorFactory func(appName string, opts runOptions) (exampleEvaluator, error)

var (
	dataDir   = flag.String("data-dir", "./data", "Directory containing evaluation set and metric files")
	outputDir = flag.String("output-dir", "./output", "Directory where evaluation results are stored")
	modelName = flag.String("model", "gpt-5.2", "Model to use for both the agent and the judge")
	streaming = flag.Bool("streaming", false, "Enable streaming responses from the agent")
	evalSetID = flag.String("eval-set", "template-basic", "Evaluation set identifier to execute")
)

const appName = "template-eval-app"

func main() {
	flag.Parse()
	err := runExample(context.Background(), newLocalEvaluator, runOptions{
		DataDir:   *dataDir,
		OutputDir: *outputDir,
		ModelName: *modelName,
		EvalSetID: *evalSetID,
		Streaming: *streaming,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func runExample(ctx context.Context, factory evaluatorFactory, opts runOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func newLocalEvaluator(appName string, opts runOptions) (exampleEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(exampleEvaluator), nil
}

type judgeModelMetricManager struct {
	metric.Manager
	modelName string
}

func (m *judgeModelMetricManager) Get(ctx context.Context, appName, evalSetID, metricName string) (*metric.EvalMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func overrideJudgeModelName(evalMetric *metric.EvalMetric, modelName string) *metric.EvalMetric {
	_ = "STUB: not implemented"
	return nil
}

func printSummary(result *evaluation.EvaluationResult, outDir string) {
	_ = "STUB: not implemented"
	return
}

type closableEvaluator struct {
	evaluation.AgentEvaluator
	runner runner.Runner
}

func (e *closableEvaluator) Close() error { _ = "STUB: not implemented"; return nil }
