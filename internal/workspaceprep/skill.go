//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package workspaceprep

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/internal/skillstage"
	rootskill "trpc.group/trpc-go/trpc-agent-go/skill"
)

// SkillSpec describes a skill working copy that must exist under
// skills/<name> before user commands execute. Source is resolved
// through the provided skill.Repository using the invocation context
// (so per-conversation skill overrides still apply).
type SkillSpec struct {
	// Key is the stable Requirement key. When empty the reconciler
	// uses "skill:<name>".
	Key string
	// Name is the skill name as registered in the Repository.
	Name string
	// Repository resolves skill source paths. It must be non-nil.
	Repository rootskill.Repository
	// ReadOnly requests the legacy read-only staged tree. The
	// default (false) matches the new writable-working-copy
	// contract.
	ReadOnly bool
	// Optional marks this requirement as non-blocking.
	Optional bool
}

// NewSkillRequirement validates SkillSpec and returns a Requirement
// that materializes the skill into skills/<name>. Source resolution
// happens lazily inside Fingerprint/Apply so that context-scoped
// repositories can honor the active invocation.
func NewSkillRequirement(spec SkillSpec) (Requirement, error) {
	_ = "STUB: not implemented"
	return *new(Requirement), nil
}

// spec.Name flows into skills/<name> and into skillstage cleanup.
// Model-driven tool invocations, untrusted skill repositories, or
// a misconfigured caller could otherwise smuggle traversal
// components (absolute paths, "..", backslash-rooted paths) and
// escape the workspace. Normalize and reject anything that does
// not resolve to a single-segment, non-traversing relative name.

// validateSkillName rejects skill names that could escape skills/<name>.
// The check is intentionally strict: we refuse anything that contains
// path separators, parent references, or leading dots. Skill naming in
// the repository layer already follows this convention, so legitimate
// callers are unaffected.
func validateSkillName(name string) error { _ = "STUB: not implemented"; return nil }

// path.Clean must be a no-op for a well-formed single-segment
// name; anything else implies hidden traversal or normalization
// surprises.

type skillRequirement struct {
	spec   SkillSpec
	stager *skillstage.Stager
}

func (r *skillRequirement) Key() string    { _ = "STUB: not implemented"; return "" }
func (r *skillRequirement) Kind() Kind     { _ = "STUB: not implemented"; return *new(Kind) }
func (r *skillRequirement) Phase() Phase   { _ = "STUB: not implemented"; return *new(Phase) }
func (r *skillRequirement) Required() bool { _ = "STUB: not implemented"; return false }
func (r *skillRequirement) Target() string { _ = "STUB: not implemented"; return "" }

// Fingerprint captures the skill source digest plus the staging mode
// so switching between read-only and writable modes forces a
// re-stage.
func (r *skillRequirement) Fingerprint(
	ctx context.Context, rctx ApplyContext,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SentinelExists reports whether the materialized skill tree is still
// structurally intact. We reuse the existing SkillLinksPresent helper
// which validates the out/work/inputs symlinks, and we additionally
// check that SKILL.md exists under skills/<name>.
func (r *skillRequirement) SentinelExists(
	ctx context.Context, rctx ApplyContext,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Apply resolves the skill source path through the repository and
// delegates to skillstage, which already handles idempotent
// restaging, symlink management, and metadata updates.
func (r *skillRequirement) Apply(
	ctx context.Context, rctx ApplyContext,
) error {
	_ = "STUB: not implemented"
	return nil
}
