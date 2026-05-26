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
	"time"

	promptiterengine "trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
	promptitermanager "trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/manager"
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

type asyncRunConfig struct {
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
	PollInterval               time.Duration
}

type sharedMetricLocator struct {
	metricFileID string
}

type promptIterRuntime struct {
	manager promptitermanager.Manager
	close   func()
}

func runAsyncRunExample(ctx context.Context, cfg asyncRunConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func buildPromptIterRuntime(ctx context.Context, cfg asyncRunConfig) (*promptIterRuntime, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildRunRequest(cfg asyncRunConfig, targetSurfaceID string) *promptiterengine.RunRequest {
	_ = "STUB: not implemented"
	return nil
}

func waitForRun(
	ctx context.Context,
	manager promptitermanager.Manager,
	runID string,
	pollInterval time.Duration,
) (*promptiterengine.RunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isTerminalRunStatus(status promptiterengine.RunStatus) bool {
	_ = "STUB: not implemented"
	return false
}

func describeRunProgress(run *promptiterengine.RunResult) string {
	_ = "STUB: not implemented"
	return ""
}

func currentRoundResult(run *promptiterengine.RunResult, roundNumber int) *promptiterengine.RoundResult {
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
