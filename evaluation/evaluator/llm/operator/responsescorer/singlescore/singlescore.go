//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package singlescore scores JSON judge outputs shaped as {score, reason}.
package singlescore

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/responsescorer"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

type singleScoreResponse struct {
	Score  *float64 `json:"score"`
	Reason *string  `json:"reason"`
}

type singleScoreResponseScorer struct {
}

// New returns a response scorer for single score JSON outputs.
func New() responsescorer.ResponseScorer {
	_ = "STUB: not implemented"
	return *new(responsescorer.ResponseScorer)
}

// ScoreBasedOnResponse parses the structured judge response.
func (s *singleScoreResponseScorer) ScoreBasedOnResponse(ctx context.Context, response *model.Response,
	_ *metric.EvalMetric) (*evaluator.ScoreResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
