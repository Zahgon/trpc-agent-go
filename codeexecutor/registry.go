//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package codeexecutor

import (
	"context"
	"sync"
)

// WorkspaceRegistry keeps a process-level mapping of logical IDs to
// created workspaces for reuse within a session.
type WorkspaceRegistry struct {
	mu       sync.Mutex
	byID     map[string]Workspace
	inflight map[string]*workspaceCreateCall
}

type workspaceCreateCall struct {
	done chan struct{}
	ws   Workspace
	err  error
}

// NewWorkspaceRegistry creates a new in-memory registry.
func NewWorkspaceRegistry() *WorkspaceRegistry { _ = "STUB: not implemented"; return nil }

// Acquire creates or returns an existing workspace with the given id.
// Concurrent first-time acquires for the same id coalesce to a single
// CreateWorkspace so init hooks and workspace creation run at most once per id.
func (r *WorkspaceRegistry) Acquire(
	ctx context.Context, m WorkspaceManager, id string,
) (Workspace, error) {
	_ = "STUB: not implemented"
	return *new(Workspace), nil
}

func (r *WorkspaceRegistry) createWorkspace(
	ctx context.Context,
	m WorkspaceManager,
	id string,
	call *workspaceCreateCall,
) {
	_ = "STUB: not implemented"
	return
}

func waitWorkspaceCreate(ctx context.Context, call *workspaceCreateCall) (Workspace, error) {
	_ = "STUB: not implemented"
	return *new(Workspace), nil
}
