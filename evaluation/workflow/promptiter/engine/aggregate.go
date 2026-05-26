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

// AggregationOptions configures aggregation-stage execution behavior.
type AggregationOptions struct {
	// SurfaceParallelismEnabled enables concurrent aggregation across target surfaces.
	SurfaceParallelismEnabled bool
	// SurfaceParallelism caps concurrent aggregation across target surfaces when SurfaceParallelismEnabled is true. Zero uses GOMAXPROCS.
	SurfaceParallelism int
}

// AggregationResult groups all surfaces after sample gradient merge.
type AggregationResult struct {
	// Surfaces stores merged gradients to feed optimizer per surface.
	Surfaces []promptiter.AggregatedSurfaceGradient
}

func (e *engine) aggregate(
	ctx context.Context,
	structure *structureState,
	backward *BackwardResult,
	targetSurfaceSet targetSurfaceSet,
	options AggregationOptions,
) (*AggregationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
