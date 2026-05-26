//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package engine implements PromptIter orchestration and runtime flow for a generation round.
package engine

import (
	atrace "trpc.group/trpc-go/trpc-agent-go/agent/trace"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter"
)

type lossHintMetricKey struct {
	evalSetID  string
	evalCaseID string
	metricName string
}

func (e *engine) loss(result *EvaluationResult) ([]promptiter.CaseLoss, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeLossHints(
	losses []promptiter.CaseLoss,
	result *EvaluationResult,
	inputs []EvalSetInput,
) ([]promptiter.CaseLoss, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func indexLossHints(inputs []EvalSetInput) map[caseResultKey][]LossHint {
	_ = "STUB: not implemented"
	return nil
}

func indexLossHintTargets(
	losses []promptiter.CaseLoss,
) (map[caseResultKey]int, map[lossHintMetricKey][]string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findMetricResult(metrics []MetricResult, metricName string) (MetricResult, bool) {
	_ = "STUB: not implemented"
	return *new(MetricResult), false
}

func traceTerminalStepIDs(trace *atrace.Trace) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sortStepIDs(stepIDs map[string]struct{}, stepOrder map[string]int) []string {
	_ = "STUB: not implemented"
	return nil
}

func sortCaseLosses(losses []promptiter.CaseLoss) { _ = "STUB: not implemented"; return }

func terminalLossLess(left promptiter.TerminalLoss, right promptiter.TerminalLoss) bool {
	_ = "STUB: not implemented"
	return false
}
