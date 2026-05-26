//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package status provides functions to summarize the evaluation status.
package status

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/status"
)

// SummarizeMetricsStatus summarizes the metric statuses into a single value.
func SummarizeMetricsStatus(metrics []*evalresult.EvalMetricResult) (status.EvalStatus, error) {
	_ = "STUB: not implemented"
	return *new(status.EvalStatus), nil
}

// Summarize summarizes the evaluation status of a single case.
// The precedence rules are:
// 1. If there is a Failed, the overall status is Failed.
// 2. If there is a Passed, the overall status is Passed.
// 3. Otherwise, the overall status is NotEvaluated.
func Summarize(statuses []status.EvalStatus) (status.EvalStatus, error) {
	_ = "STUB: not implemented"
	return *new(status.EvalStatus), nil
}
