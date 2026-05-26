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
	"log"

	promptiterengine "trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	appName             = "promptiter-nba-commentary-app"
	candidateAppName    = "promptiter-nba-commentary-candidate"
	judgeAppName        = "promptiter-nba-commentary-judge"
	backwarderAppName   = "promptiter-nba-commentary-backwarder"
	aggregatorAppName   = "promptiter-nba-commentary-aggregator"
	optimizerAppName    = "promptiter-nba-commentary-optimizer"
	trainEvalSetID      = "nba-commentary-train"
	validationEvalSetID = "nba-commentary-validation"
	sharedMetricFileID  = "sports-commentary"
)

type syncRunConfig struct {
	DataDir                    string
	OutputDir                  string
	CandidateModelName         string
	CandidateInstruction       string
	JudgeModelName             string
	WorkerModelName            string
	MaxRounds                  int
	MinScoreGain               float64
	MaxRoundsWithoutAcceptance int
	TargetScore                float64
	EvalCaseParallelism        int
	ParallelInferenceEnabled   bool
	ParallelEvaluationEnabled  bool
	DebugIO                    bool
	Logger                     *log.Logger
}

type sharedMetricLocator struct {
	metricFileID string
}

type promptIterRuntime struct {
	engine promptiterengine.Engine
	close  func()
}

func runSyncRunExample(ctx context.Context, cfg syncRunConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func runSyncRun(
	ctx context.Context,
	cfg syncRunConfig,
) (*promptiterengine.RunResult, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func buildPromptIterRuntime(ctx context.Context, cfg syncRunConfig) (*promptIterRuntime, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildRunRequest(cfg syncRunConfig, targetSurfaceID string) *promptiterengine.RunRequest {
	_ = "STUB: not implemented"
	return nil
}

// Build maps every eval set to the shared metric file used by the example.
func (l *sharedMetricLocator) Build(baseDir, appName, _ string) string {
	_ = "STUB: not implemented"
	return ""
}

func loadOpenAIModel(modelName string) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}
