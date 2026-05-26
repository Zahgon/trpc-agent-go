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
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter"
)

func normalizeProfile(
	structure *structureState,
	profile *promptiter.Profile,
) (*promptiter.Profile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyPatchSet(
	structure *structureState,
	profile *promptiter.Profile,
	patchSet *promptiter.PatchSet,
) (*promptiter.Profile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildOverrideIndex(profile *promptiter.Profile) map[string]promptiter.SurfaceOverride {
	_ = "STUB: not implemented"
	return nil
}

func buildProfileFromOverrideIndex(
	structure *structureState,
	overrideIndex map[string]promptiter.SurfaceOverride,
) *promptiter.Profile {
	_ = "STUB: not implemented"
	return nil
}

func resolveProfileSurface(
	structure *structureState,
	overrideIndex map[string]promptiter.SurfaceOverride,
	surfaceID string,
) (astructure.Surface, error) {
	_ = "STUB: not implemented"
	return *new(astructure.Surface), nil
}
