//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package engine implements PromptIter orchestration and runtime flow for a generation round.
package engine

type targetSurfaceSet map[string]struct{}

func compileTargetSurfaceIDs(
	structure *structureState,
	targetSurfaceIDs []string,
) (targetSurfaceSet, error) {
	_ = "STUB: not implemented"
	return *new(targetSurfaceSet), nil
}

func (s targetSurfaceSet) contains(surfaceID string) bool { _ = "STUB: not implemented"; return false }
