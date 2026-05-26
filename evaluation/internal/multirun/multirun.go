//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package multirun provides helpers for summarizing multi-run evaluation results.
package multirun

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/status"
)

// SummarizeMultiRun populates EvalSetResult.Summary based on the current EvalCaseResults.
func SummarizeMultiRun(evalSetResult *evalresult.EvalSetResult, expectedNumRuns int) error {
	_ = "STUB: not implemented"
	return nil
}

func groupCaseResultsByRunID(caseResults []*evalresult.EvalCaseResult, expectedNumRuns int) (map[int][]*evalresult.EvalCaseResult, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func buildEvalSetRunSummaries(runCaseResults map[int][]*evalresult.EvalCaseResult, runIDs []int) ([]*evalresult.EvalSetRunSummary, evalresult.EvalStatusCounts, error) {
	_ = "STUB: not implemented"
	return nil, *new(evalresult.EvalStatusCounts), nil
}

type caseAgg struct {
	runStatusCounts evalresult.EvalStatusCounts
	hasRunError     bool
	runSummaries    []*evalresult.EvalCaseRunSummary
	metricAgg       map[string]*metricAgg
}

func buildEvalCaseSummaries(runCaseResults map[int][]*evalresult.EvalCaseResult, runIDs []int) ([]*evalresult.EvalCaseResultSummary, status.EvalStatus, error) {
	_ = "STUB: not implemented"
	return nil, *new(status.EvalStatus), nil
}

func summarizeOverallFromMetricSummaries(metricSummaries []*evalresult.EvalMetricSummary, hasRunError bool) (status.EvalStatus, error) {
	_ = "STUB: not implemented"
	return *new(status.EvalStatus), nil
}

func addEvalStatus(counts *evalresult.EvalStatusCounts, s status.EvalStatus) error {
	_ = "STUB: not implemented"
	return nil
}

func normalizeCounts(counts evalresult.EvalStatusCounts) *evalresult.EvalStatusCounts {
	_ = "STUB: not implemented"
	return nil
}

type metricAgg struct {
	threshold       float64
	thresholdLoaded bool
	evaluatedCount  int
	scoreSum        float64
	statusCounts    evalresult.EvalStatusCounts
}

func mergeMetricAgg(agg map[string]*metricAgg, metricResults []*evalresult.EvalMetricResult) error {
	_ = "STUB: not implemented"
	return nil
}

func buildMetricSummaries(agg map[string]*metricAgg) []*evalresult.EvalMetricSummary {
	_ = "STUB: not implemented"
	return nil
}

func buildMetricRunSummaries(metricResults []*evalresult.EvalMetricResult) []*evalresult.EvalMetricRunSummary {
	_ = "STUB: not implemented"
	return nil
}
