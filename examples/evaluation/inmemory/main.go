//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"
	"flag"

	"trpc.group/trpc-go/trpc-agent-go/evaluation"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
	evalresultinmemory "trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult/inmemory"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	evalsetinmemory "trpc.group/trpc-go/trpc-agent-go/evaluation/evalset/inmemory"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/registry"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	metricinmemory "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/inmemory"
	"trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	modelName = flag.String("model", "deepseek-v4-flash", "Model to use for evaluation runs")
	streaming = flag.Bool("streaming", false, "Enable streaming responses from the agent")
	numRuns   = flag.Int("runs", 1, "Number of times to repeat the evaluation loop per case")
)

const (
	appName   = "math-eval-app"
	evalSetID = "math-basic"
)

func main() {
	flag.Parse()
	ctx := context.Background()
	// New runner.
	run := runner.NewRunner(appName, newCalculatorAgent(*modelName, *streaming))

	// Ensure runner resources are cleaned up (trpc-agent-go >= v0.5.0)
	defer run.Close()

	// New manager and registry for evaluation.
	evalSetManager := evalsetinmemory.New()
	metricManager := metricinmemory.New()
	evalResultManager := evalresultinmemory.New()
	registry := registry.New()
	// Prepare evalset and metric.
	if err := prepareEvalSet(ctx, evalSetManager); err != nil {
		log.Fatalf("prepare eval set: %v", err)
	}
	if err := prepareMetric(ctx, metricManager); err != nil {
		log.Fatalf("prepare metric: %v", err)
	}
	// New agent evaluator.
	agentEvaluator, err := evaluation.New(
		appName,
		run,
		evaluation.WithEvalSetManager(evalSetManager),
		evaluation.WithMetricManager(metricManager),
		evaluation.WithEvalResultManager(evalResultManager),
		evaluation.WithRegistry(registry),
		evaluation.WithNumRuns(*numRuns),
	)
	if err != nil {
		log.Fatalf("create evaluator: %v", err)
	}
	defer func() { agentEvaluator.Close() }()
	// Run evaluate.
	result, err := agentEvaluator.Evaluate(ctx, evalSetID)
	if err != nil {
		log.Fatalf("evaluate: %v", err)
	}
	printSummary(ctx, result, evalResultManager)
}

func printSummary(ctx context.Context, result *evaluation.EvaluationResult, evalResultManager evalresult.Manager) {
	_ = "STUB: not implemented"
	return
}

func prepareEvalSet(ctx context.Context, evalSetManager evalset.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func prepareMetric(ctx context.Context, metricManager metric.Manager) error {
	_ = "STUB: not implemented"
	return nil
}
