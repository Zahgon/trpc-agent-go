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

	"trpc.group/trpc-go/trpc-agent-go/agent"
	promptiterengine "trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
	promptitermanager "trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/manager"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

func runPromptIter(
	ctx context.Context,
	candidateModel model.Model,
	judgeModel model.Model,
	workerModel model.Model,
) (*promptiterengine.RunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildPromptIterRunRequest() *promptiterengine.RunRequest {
	_ = "STUB: not implemented"
	return nil
}

func waitForRun(
	ctx context.Context,
	manager promptitermanager.Manager,
	runID string,
	interval time.Duration,
) (*promptiterengine.RunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newPromptIterAgent(name string, m model.Model) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func candidateSurfaceIDs() []string { _ = "STUB: not implemented"; return nil }

func isTerminalRunStatus(status promptiterengine.RunStatus) bool {
	_ = "STUB: not implemented"
	return false
}

func terminalRunError(run *promptiterengine.RunResult) error { _ = "STUB: not implemented"; return nil }

func closeRunners(runners ...runner.Runner) { _ = "STUB: not implemented"; return }
