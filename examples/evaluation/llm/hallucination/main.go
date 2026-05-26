//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main implements functions for the hallucination evaluation.
package main

import (
	"context"
	"flag"
	"log"

	"trpc.group/trpc-go/trpc-agent-go/evaluation"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
	evalresultlocal "trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult/local"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	evalsetlocal "trpc.group/trpc-go/trpc-agent-go/evaluation/evalset/local"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/registry"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	metriclocal "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/local"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	dataDir            = flag.String("data-dir", "./data", "Directory containing evaluation set and metric files")
	outputDir          = flag.String("output-dir", "./output", "Directory where evaluation results will be stored")
	modelName          = flag.String("model", "gpt-5.4", "Model to use for the candidate agent and judge runner")
	streaming          = flag.Bool("streaming", false, "Enable streaming responses from the agent")
	forceHallucination = flag.Bool("force-hallucination", false, "Use a scripted candidate agent that emits tool traces and returns intentionally wrong facts")
	evalSetID          = flag.String("eval-set", "hallucination-basic", "Evaluation set identifier to execute")
)

const appName = "hallucination-eval-app"

func main() {
	flag.Parse()
	ctx := context.Background()
	actualAgent := newQAAgent(*modelName, *streaming)
	if *forceHallucination {
		actualAgent = newForcedHallucinationAgent()
	}
	actualRunner := runner.NewRunner(appName, actualAgent)
	defer actualRunner.Close()
	judgeRunner := runner.NewRunner(appName+"-judge", newJudgeAgent(*modelName))
	defer judgeRunner.Close()
	evalSetManager := evalsetlocal.New(evalset.WithBaseDir(*dataDir))
	metricManager := metriclocal.New(metric.WithBaseDir(*dataDir))
	evalResultManager := evalresultlocal.New(evalresult.WithBaseDir(*outputDir))
	registry := registry.New()
	agentEvaluator, err := evaluation.New(
		appName,
		actualRunner,
		evaluation.WithEvalSetManager(evalSetManager),
		evaluation.WithMetricManager(metricManager),
		evaluation.WithEvalResultManager(evalResultManager),
		evaluation.WithRegistry(registry),
		evaluation.WithJudgeRunner(judgeRunner),
	)
	if err != nil {
		log.Fatalf("create evaluator: %v", err)
	}
	defer func() { agentEvaluator.Close() }()
	result, err := agentEvaluator.Evaluate(ctx, *evalSetID)
	if err != nil {
		log.Fatalf("evaluate: %v", err)
	}
	printSummary(result, *outputDir)
}

func printSummary(result *evaluation.EvaluationResult, outDir string) {
	_ = "STUB: not implemented"
	return
}
