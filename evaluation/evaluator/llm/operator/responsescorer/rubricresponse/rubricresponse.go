//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package rubricresponse scores rubric-graded judge outputs.
package rubricresponse

import (
	"context"
	"regexp"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/responsescorer"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const passedVerdict = "yes"

// rubricBlockRegex extracts rubric ID, property, evidence, reason, and verdict blocks.
var rubricBlockRegex = regexp.MustCompile(
	`(?ms)ID:\s*(.*?)\s*` + // 1: rubric id
		`Rubric:\s*(.*?)\s*` + // 2: rubric text
		`Evidence:\s*(.*?)\s*` + // 3: evidence text
		`Reason:\s*(.*?)\s*` + // 4: reason text
		`Verdict:\s*(.*?)\s*$`, // 5: verdict yes/no
)

type rubricResponseScorer struct {
}

// New returns a response scorer for rubric responses.
func New() responsescorer.ResponseScorer {
	_ = "STUB: not implemented"
	return *new(responsescorer.ResponseScorer)
}

// ScoreBasedOnResponse scores rubric responses.
func (e *rubricResponseScorer) ScoreBasedOnResponse(ctx context.Context, response *model.Response,
	evalMetric *metric.EvalMetric) (*evaluator.ScoreResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func configuredRubricCount(evalMetric *metric.EvalMetric) int { _ = "STUB: not implemented"; return 0 }
