//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package rubricscores scores JSON judge outputs shaped as {rubricScores: [...]}.
package rubricscores

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/responsescorer"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

type rubricScoreItem struct {
	ID     *string  `json:"id"`
	Score  *float64 `json:"score"`
	Reason *string  `json:"reason"`
}

type rubricScoresResponse struct {
	RubricScores []rubricScoreItem `json:"rubricScores"`
}

type rubricScoresResponseScorer struct {
}

// New returns a response scorer for rubric scores JSON outputs.
func New() responsescorer.ResponseScorer {
	_ = "STUB: not implemented"
	return *new(responsescorer.ResponseScorer)
}

// ScoreBasedOnResponse parses the structured judge response.
func (s *rubricScoresResponseScorer) ScoreBasedOnResponse(ctx context.Context, response *model.Response,
	evalMetric *metric.EvalMetric) (*evaluator.ScoreResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func expectedRubricIDs(evalMetric *metric.EvalMetric) (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateRubricScoreID(id string, seen, expected map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func missingRubricScoreID(seen, expected map[string]struct{}) string {
	_ = "STUB: not implemented"
	return ""
}
