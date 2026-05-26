//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
)

func printSummary(
	result *engine.RunResult,
	dataDir string,
	outputDir string,
	initialInstruction string,
	targetSurfaceID string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func initialValidationScore(result *engine.RunResult) float64 { _ = "STUB: not implemented"; return 0 }

func finalAcceptedValidationScore(result *engine.RunResult) float64 {
	_ = "STUB: not implemented"
	return 0
}

func evaluationResultScore(result *engine.EvaluationResult) float64 {
	_ = "STUB: not implemented"
	return 0
}

func acceptedInstructionText(
	result *engine.RunResult,
	initialInstruction string,
	targetSurfaceID string,
) string {
	_ = "STUB: not implemented"
	return ""
}
