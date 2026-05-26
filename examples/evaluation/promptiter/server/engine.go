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

	promptiterengine "trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
	promptitermanager "trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/manager"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	appName            = "promptiter-nba-commentary-app"
	candidateAppName   = "promptiter-nba-commentary-candidate"
	judgeAppName       = "promptiter-nba-commentary-judge"
	backwarderAppName  = "promptiter-nba-commentary-backwarder"
	aggregatorAppName  = "promptiter-nba-commentary-aggregator"
	optimizerAppName   = "promptiter-nba-commentary-optimizer"
	sharedMetricFileID = "sports-commentary"
)

type serverConfig struct {
	Addr                      string
	BasePath                  string
	DataDir                   string
	OutputDir                 string
	CandidateModelName        string
	CandidateInstruction      string
	JudgeModelName            string
	WorkerModelName           string
	EvalCaseParallelism       int
	ParallelInferenceEnabled  bool
	ParallelEvaluationEnabled bool
}

type sharedMetricLocator struct {
	metricFileID string
}

type promptIterRuntime struct {
	engine  promptiterengine.Engine
	manager promptitermanager.Manager
	close   func()
}

func buildPromptIterRuntime(ctx context.Context, cfg serverConfig) (*promptIterRuntime, error) {
	_ = "STUB: not implemented"
	return nil, nil
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
