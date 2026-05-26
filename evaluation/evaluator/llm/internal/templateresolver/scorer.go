//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package templateresolver resolves template evaluator runtime components.
package templateresolver

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator/llm/operator/responsescorer"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	// ResponseScorerSingleScoreName identifies the scalar score response scorer.
	ResponseScorerSingleScoreName = "single_score"
	// ResponseScorerRubricScoresName identifies the rubric scores response scorer.
	ResponseScorerRubricScoresName = "rubric_scores"
)

// ResolveResponseScorer returns the response scorer identified by name.
func ResolveResponseScorer(name string) (responsescorer.ResponseScorer, error) {
	_ = "STUB: not implemented"
	return *new(responsescorer.ResponseScorer), nil
}

// StructuredOutput returns the schema associated with the named response scorer.
func StructuredOutput(name string) (*model.StructuredOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
