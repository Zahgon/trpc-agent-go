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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

var (
	modelName                  = flag.String("model", "deepseek-v3.2", "Model identifier used by the candidate sports recap agent")
	judgeModelName             = flag.String("judge-model", "gpt-5.2", "Model identifier used by the judge agent")
	workerModelName            = flag.String("worker-model", "gpt-5.2", "Model identifier used by the PromptIter worker agents")
	dataDir                    = flag.String("data-dir", "./data", "Directory containing evaluation set and metric files")
	outputDir                  = flag.String("output-dir", "./output", "Directory where evaluation results will be stored")
	maxRounds                  = flag.Int("max-rounds", 4, "Maximum PromptIter optimization rounds")
	evalCaseParallelism        = flag.Int("eval-case-parallelism", 16, "Maximum number of eval cases processed in parallel")
	backwardCaseParallelism    = flag.Int("backward-case-parallelism", 16, "Maximum number of train eval cases processed in parallel during backward; 0 uses GOMAXPROCS")
	aggregationParallelism     = flag.Int("aggregation-parallelism", 16, "Maximum number of target surfaces aggregated in parallel; 0 uses GOMAXPROCS")
	optimizerParallelism       = flag.Int("optimizer-parallelism", 16, "Maximum number of target surfaces optimized in parallel; 0 uses GOMAXPROCS")
	parallelInferenceEnabled   = flag.Bool("parallel-inference", true, "Enable parallel inference across eval cases")
	parallelEvaluationEnabled  = flag.Bool("parallel-evaluation", true, "Enable parallel evaluation across eval cases")
	parallelBackwardEnabled    = flag.Bool("parallel-backward", false, "Enable parallel backward processing across train eval cases")
	parallelAggregationEnabled = flag.Bool("parallel-aggregation", true, "Enable parallel aggregation across target surfaces")
	parallelOptimizerEnabled   = flag.Bool("parallel-optimization", true, "Enable parallel optimization across target surfaces")
	minScoreGain               = flag.Float64("min-score-gain", 0.01, "Minimum validation score gain required to accept a patch")
	maxRoundsWithoutAcceptance = flag.Int("max-rounds-without-acceptance", 3, "Maximum consecutive rejected rounds before stopping")
	targetScore                = flag.Float64("target-score", 1.01, "Target validation score that stops optimization when reached")
	pollInterval               = flag.Duration("poll-interval", 30*time.Second, "Polling interval used to report PromptIter manager progress")
)

func main() {
	flag.Parse()
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func loadOpenAIModel(modelName string) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}
