//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package runtimeprofile

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/skill"
)

// WorkspaceFromContext returns the active profile workspace policy.
func WorkspaceFromContext(ctx context.Context) (WorkspacePolicy, bool) {
	_ = "STUB: not implemented"
	return *new(WorkspacePolicy), false
}

// CredentialPolicyFromContext returns the active profile credential policy.
func CredentialPolicyFromContext(
	ctx context.Context,
) (CredentialPolicy, bool) {
	_ = "STUB: not implemented"
	return *new(CredentialPolicy), false
}

// ResolveWorkdir applies the profile workspace default and allowed roots.
func ResolveWorkdir(ctx context.Context, requested string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// CheckCredentialRef returns an error when ref is denied by the profile.
func CheckCredentialRef(ctx context.Context, ref string) error {
	_ = "STUB: not implemented"
	return nil
}

// SkillVisibilityFilter filters skills according to runtime profile policy.
func SkillVisibilityFilter(ctx context.Context, summary skill.Summary) bool {
	_ = "STUB: not implemented"
	return false
}

// SkillPathResolver resolves the filesystem path for a skill name.
type SkillPathResolver interface {
	Path(name string) (string, error)
}

// SkillVisibilityFilterForRepository filters skills by include/exclude and
// optional profile roots.
func SkillVisibilityFilterForRepository(
	resolver SkillPathResolver,
) skill.VisibilityFilter {
	_ = "STUB: not implemented"
	return *new(skill.VisibilityFilter)
}

func skillVisibleForProfile(
	ctx context.Context,
	summary skill.Summary,
	resolver SkillPathResolver,
) bool {
	_ = "STUB: not implemented"
	return false
}

func skillInAllowedRoots(
	name string,
	roots []string,
	resolver SkillPathResolver,
) bool {
	_ = "STUB: not implemented"
	return false
}

func isPathAllowed(path string, roots []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func comparablePath(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func cloneWorkspacePolicy(policy WorkspacePolicy) WorkspacePolicy {
	_ = "STUB: not implemented"
	return *new(WorkspacePolicy)
}
