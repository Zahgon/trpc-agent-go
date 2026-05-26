//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

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
	dataDir   = flag.String("data-dir", "./data", "Directory containing evaluation set and metric files")
	outputDir = flag.String("output-dir", "./output", "Directory where evaluation results will be stored")
	evalSetID = flag.String("eval-set", "claudecode-basic", "Evaluation set identifier to execute")
	numRuns   = flag.Int("runs", 1, "Number of times to repeat the evaluation loop per case")

	claudeBin    = flag.String("claude-bin", "claude", "Claude Code CLI executable path")
	outputFormat = flag.String("output-format", "json", "Transcript output format: json or stream-json")
	workDir      = flag.String("work-dir", ".", "Claude Code project directory (contains .mcp.json and .claude)")
	logDir       = flag.String("log-dir", "", "Optional directory to persist raw Claude CLI stdout/stderr logs")
)

const appName = "claudecode-eval-app"

func main() {
	flag.Parse()
	ctx := context.Background()

	ag, err := newClaudeCodeEvalAgent(*claudeBin, *outputFormat, *workDir, *logDir)
	if err != nil {
		log.Fatalf("create agent: %v", err)
	}

	r := runner.NewRunner(appName, ag)
	defer r.Close()

	evalSetManager := evalsetlocal.New(evalset.WithBaseDir(*dataDir))
	metricManager := metriclocal.New(metric.WithBaseDir(*dataDir))
	evalResultManager := evalresultlocal.New(evalresult.WithBaseDir(*outputDir))
	reg := registry.New()

	agentEvaluator, err := evaluation.New(
		appName,
		r,
		evaluation.WithEvalSetManager(evalSetManager),
		evaluation.WithMetricManager(metricManager),
		evaluation.WithEvalResultManager(evalResultManager),
		evaluation.WithRegistry(reg),
		evaluation.WithNumRuns(*numRuns),
	)
	if err != nil {
		log.Fatalf("create evaluator: %v", err)
	}
	defer func() { _ = agentEvaluator.Close() }()

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
