//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/evaluation"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

func newSportsRecapEvaluator(candidateRunner runner.Runner, judgeRunner runner.Runner) (evaluation.AgentEvaluator, error) {
	_ = "STUB: not implemented"
	return *new(evaluation.AgentEvaluator), nil
}

type sportsRecapMetricLocator struct{}

func (sportsRecapMetricLocator) Build(baseDir string, appName string, _ string) string {
	_ = "STUB: not implemented"
	return ""
}

func newJudgeAgent(m model.Model) agent.Agent { _ = "STUB: not implemented"; return *new(agent.Agent) }
