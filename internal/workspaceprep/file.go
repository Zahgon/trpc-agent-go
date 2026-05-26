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

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
)

// FileSpec describes a single file-shaped workspace requirement.
//
// Exactly one source strategy is honored, in the following precedence:
//
//  1. Content: inline bytes
//  2. Input: a codeexecutor.InputSpec (artifact://, host://,
//     workspace://, skill://)
//
// The Target is required and must be a workspace-relative path.
// Fingerprint combines the source identity (or content hash) with the
// target path so that moving the same content to a new location is
// treated as a separate requirement.
type FileSpec struct {
	// Key is the stable Requirement key. When empty a deterministic
	// key is derived from Target.
	Key string
	// Target is the workspace-relative destination path.
	Target string
	// Content is inline bytes. When set, Input is ignored.
	Content []byte
	// Mode is the POSIX mode for Content writes. Defaults to 0o644.
	Mode uint32
	// Input is a richer source spec reusing codeexecutor.InputSpec.
	Input *codeexecutor.InputSpec
	// Optional marks this requirement as non-blocking.
	Optional bool
}

// NewFileRequirement builds a Requirement from FileSpec after
// validating the basic invariants (target present, at least one
// source).
func NewFileRequirement(spec FileSpec) (Requirement, error) {
	_ = "STUB: not implemented"
	return *new(Requirement), nil
}

type fileRequirement struct {
	spec FileSpec
}

func (r *fileRequirement) Key() string    { _ = "STUB: not implemented"; return "" }
func (r *fileRequirement) Kind() Kind     { _ = "STUB: not implemented"; return *new(Kind) }
func (r *fileRequirement) Phase() Phase   { _ = "STUB: not implemented"; return *new(Phase) }
func (r *fileRequirement) Required() bool { _ = "STUB: not implemented"; return false }
func (r *fileRequirement) Target() string { _ = "STUB: not implemented"; return "" }

func (r *fileRequirement) Fingerprint(
	ctx context.Context, rctx ApplyContext,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SentinelExists checks whether the target path still exists on the
// local filesystem. For container-backed engines the sentinel is the
// same path inside the workspace; we fall back to FS().Collect when
// the path is not reachable from the host.
func (r *fileRequirement) SentinelExists(
	ctx context.Context, rctx ApplyContext,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Non-local engines can return a permission error here;
// in that case we fall through to FS().Collect.

// Apply writes the file into the workspace. Inline content is routed
// through FS().PutFiles; InputSpec-based sources use the engine's
// StageInputs so that symlink/copy semantics, host:// mounts and
// artifact fetches all reuse existing codeexecutor plumbing.
func (r *fileRequirement) Apply(
	ctx context.Context, rctx ApplyContext,
) error {
	_ = "STUB: not implemented"
	return nil
}

func cleanRel(p string) string { _ = "STUB: not implemented"; return "" }
