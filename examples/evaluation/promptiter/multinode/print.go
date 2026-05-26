//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	promptiterengine "trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
)

func printProgress(runID string, run *promptiterengine.RunResult) {
	_ = "STUB: not implemented"
	return
}

func reportBaseline(runID string, run *promptiterengine.RunResult, reported bool) bool {
	_ = "STUB: not implemented"
	return false
}

func reportRoundMilestones(
	runID string,
	run *promptiterengine.RunResult,
	reportedTrainRounds map[int]struct{},
	reportedValidationRounds map[int]struct{},
	reportedCompletedRounds map[int]struct{},
) {
	_ = "STUB: not implemented"
	return
}

func describeRunProgress(run *promptiterengine.RunResult) string {
	_ = "STUB: not implemented"
	return ""
}

func currentRoundResult(run *promptiterengine.RunResult, roundNumber int) *promptiterengine.RoundResult {
	_ = "STUB: not implemented"
	return nil
}

func printRunSummary(result *promptiterengine.RunResult, candidateSurfaceIDs []string) {
	_ = "STUB: not implemented"
	return
}

func finalAcceptedValidationScore(result *promptiterengine.RunResult) float64 {
	_ = "STUB: not implemented"
	return 0
}
