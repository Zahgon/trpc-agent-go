//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package rubrics extracts and validates LLM judge rubrics shared by evaluator operators.
package rubrics

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// VisibleRubric is a rubric that is visible to the judge prompt.
type VisibleRubric struct {
	ID string
}

// Visible returns rubrics that would be rendered into judge prompts.
func Visible(evalMetric *metric.EvalMetric) []VisibleRubric { _ = "STUB: not implemented"; return nil }

// Count returns the number of rubrics that would be rendered into judge prompts.
func Count(evalMetric *metric.EvalMetric) int { _ = "STUB: not implemented"; return 0 }

// ValidateStructured returns visible rubrics that can safely drive structured output schemas.
func ValidateStructured(evalMetric *metric.EvalMetric) ([]VisibleRubric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ScoresOutput returns a structured output schema for per-rubric scores.
func ScoresOutput(name, description string, visibleRubrics []VisibleRubric) *model.StructuredOutput {
	_ = "STUB: not implemented"
	return nil
}
