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
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
	"trpc.group/trpc-go/trpc-agent-go/internal/skillstage"
)

// defaultReconciler is the process-local, single-node implementation
// of Reconciler. It uses a keyed mutex on ws.Path to serialize
// reconciles for the same workspace, reads/writes WorkspaceMetadata
// through the shared skillstage helpers, and enforces a fixed phase
// order (PhaseFile -> PhaseSkill -> PhaseCommand).
type defaultReconciler struct {
	locker *keyedLocker
	stager *skillstage.Stager
}

// NewReconciler returns the default Reconciler used by workspace_exec
// and other workspace-aware tools.
func NewReconciler() Reconciler { _ = "STUB: not implemented"; return *new(Reconciler) }

// Reconcile implements Reconciler.
func (r *defaultReconciler) Reconcile(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	reqs []Requirement,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *defaultReconciler) saveReconcileMetadata(
	ctx context.Context,
	eng codeexecutor.Engine,
	ws codeexecutor.Workspace,
	base codeexecutor.WorkspaceMetadata,
	md codeexecutor.WorkspaceMetadata,
	changedKeys []string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func mergeReconcileMetadata(
	latest codeexecutor.WorkspaceMetadata,
	base codeexecutor.WorkspaceMetadata,
	updated codeexecutor.WorkspaceMetadata,
	changedKeys []string,
) codeexecutor.WorkspaceMetadata {
	_ = "STUB: not implemented"
	return *new(codeexecutor.WorkspaceMetadata)
}

func mergeDirectMetadataChanges(
	merged *codeexecutor.WorkspaceMetadata,
	base codeexecutor.WorkspaceMetadata,
	updated codeexecutor.WorkspaceMetadata,
) {
	_ = "STUB: not implemented"
	return
}

func cloneReconcileMetadata(
	md codeexecutor.WorkspaceMetadata,
) codeexecutor.WorkspaceMetadata {
	_ = "STUB: not implemented"
	return *new(codeexecutor.WorkspaceMetadata)
}

// runOne applies a single requirement. It returns whether work was
// done (so the caller knows to persist metadata), a non-empty warning
// string that callers should surface, and an error on hard failure.
func (r *defaultReconciler) runOne(
	ctx context.Context,
	rctx ApplyContext,
	req Requirement,
) (bool, string, error) {
	_ = "STUB: not implemented"
	return false, "", nil
}

// sortRequirements orders requirements by Phase and, within a phase,
// by their original position (Go's sort.SliceStable preserves
// insertion order for equal keys). Callers should pass the slice in
// the order Providers were registered so that behavior is
// deterministic.
func sortRequirements(reqs []Requirement) { _ = "STUB: not implemented"; return }

// dedupeRequirements removes duplicate requirements by Key while
// preserving the first occurrence. This lets multiple Providers
// contribute overlapping requirements without forcing them to
// coordinate; the reconciler simply honors the first one it saw.
func dedupeRequirements(in []Requirement) []Requirement { _ = "STUB: not implemented"; return nil }

// keyedLocker is a small process-local keyed mutex used to serialize
// reconciles for the same workspace. A sync.Map-backed implementation
// would be acceptable too; the simple map+mutex version is chosen for
// clarity because contention is rare (same-session parallel tool calls
// reconciling at the same time).
type keyedLocker struct {
	mu    sync.Mutex
	locks map[string]*keyedLock
}

type keyedLock struct {
	mu   sync.Mutex
	refs int
}

func newKeyedLocker() *keyedLocker { _ = "STUB: not implemented"; return nil }

// lock acquires the mutex for the given key and returns an unlock
// function. The lock is reference-counted so parallel callers for
// different keys never contend on the outer mutex for longer than
// needed.
func (k *keyedLocker) lock(key string) func() {
	_ = "STUB: not implemented"

	// Fall back to a shared lock for empty keys so callers still
	// get serialization even when ws.Path is unexpectedly empty.
	return nil
}
