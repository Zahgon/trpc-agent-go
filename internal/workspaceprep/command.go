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
	"time"
)

// CommandSpec describes a one-shot bootstrap command to execute during
// reconcile. Typical uses are "create a virtualenv once per workspace"
// and "pip install -r requirements.txt when requirements change".
//
// Self-healing notes:
//
//   - When MarkerPath is set, the reconciler treats the marker as the
//     sentinel: if a user removes the marker between reconciles, the
//     command re-runs even when the fingerprint is unchanged.
//   - When ObservedPaths is set, the sentinel is considered satisfied
//     only if all observed paths still exist.
//   - When neither is set, the sentinel is always "present"; the
//     command only re-runs when Fingerprint changes. This matches the
//     behavior documented in the architecture plan.
//
// FingerprintInputs lets callers fold the contents of arbitrary
// workspace-relative files into the fingerprint so that edits to, for
// example, requirements.txt naturally force a re-run.
type CommandSpec struct {
	// Key is the stable Requirement key. When empty a deterministic
	// key is derived from Cmd+Args.
	Key string
	// Cmd is the program to execute, exactly as RunProgramSpec.Cmd.
	Cmd string
	// Args are command-line arguments passed verbatim.
	Args []string
	// Env augments the run environment.
	Env map[string]string
	// Cwd is a workspace-relative working directory.
	Cwd string
	// Timeout bounds a single run.
	Timeout time.Duration
	// MarkerPath is a workspace-relative sentinel file. When set and
	// missing, Apply creates it after a successful run.
	MarkerPath string
	// ObservedPaths are additional workspace-relative paths used as
	// sentinels.
	ObservedPaths []string
	// FingerprintInputs are workspace-relative files whose contents
	// are hashed into Fingerprint. Missing files contribute an empty
	// hash segment rather than causing an error.
	FingerprintInputs []string
	// FingerprintSalt is a caller-supplied version string added to
	// the fingerprint, letting business config force a re-run without
	// changing Cmd/Args.
	FingerprintSalt string
	// Optional marks this requirement as non-blocking.
	Optional bool
}

// NewCommandRequirement builds a Requirement from CommandSpec.
func NewCommandRequirement(spec CommandSpec) (Requirement, error) {
	_ = "STUB: not implemented"
	return *new(Requirement), nil
}

type commandRequirement struct {
	spec CommandSpec
}

func (r *commandRequirement) Key() string    { _ = "STUB: not implemented"; return "" }
func (r *commandRequirement) Kind() Kind     { _ = "STUB: not implemented"; return *new(Kind) }
func (r *commandRequirement) Phase() Phase   { _ = "STUB: not implemented"; return *new(Phase) }
func (r *commandRequirement) Required() bool { _ = "STUB: not implemented"; return false }
func (r *commandRequirement) Target() string { _ = "STUB: not implemented"; return "" }

func (r *commandRequirement) Fingerprint(
	ctx context.Context, rctx ApplyContext,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SentinelExists checks the MarkerPath (if set) and each ObservedPath.
// When neither is configured, the sentinel is considered present so
// that skip decisions are driven purely by Fingerprint.
func (r *commandRequirement) SentinelExists(
	ctx context.Context, rctx ApplyContext,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Apply runs the configured command through eng.Runner() and, on
// success, creates the MarkerPath (when configured) so future
// reconciles can use it as a sentinel.
func (r *commandRequirement) Apply(
	ctx context.Context, rctx ApplyContext,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *commandRequirement) readFile(
	ctx context.Context, rctx ApplyContext, rel string,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fall through to FS() for non-local engines.

func (r *commandRequirement) pathExists(
	ctx context.Context, rctx ApplyContext, rel string,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func cloneEnv(in map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

func trimForError(stderr, stdout string) string { _ = "STUB: not implemented"; return "" }
