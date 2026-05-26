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
	astructure "trpc.group/trpc-go/trpc-agent-go/agent/structure"
)

type structureState struct {
	snapshot        *astructure.Snapshot
	nodeIndex       map[string]astructure.Node
	surfaceIndex    map[string]astructure.Surface
	knownSurfaceIDs map[string]struct{}
}

func newStructureState(snapshot *astructure.Snapshot) (*structureState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildKnownSurfaceIDs(
	surfaces []astructure.Surface,
	nodeIndex map[string]astructure.Node,
) (map[string]struct{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
