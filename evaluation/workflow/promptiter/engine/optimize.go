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
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter"
)

// OptimizerOptions configures optimizer-stage execution behavior.
type OptimizerOptions struct {
	// SurfaceParallelismEnabled enables concurrent optimization across target surfaces.
	SurfaceParallelismEnabled bool
	// SurfaceParallelism caps concurrent optimization across target surfaces when SurfaceParallelismEnabled is true. Zero uses GOMAXPROCS.
	SurfaceParallelism int
}

func (e *engine) optimize(
	ctx context.Context,
	structure *structureState,
	profile *promptiter.Profile,
	aggregation *AggregationResult,
	targetSurfaceSet targetSurfaceSet,
	options OptimizerOptions,
) (*promptiter.PatchSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneAggregatedGradient(gradient promptiter.AggregatedSurfaceGradient) *promptiter.AggregatedSurfaceGradient {
	_ = "STUB: not implemented"
	return nil
}
