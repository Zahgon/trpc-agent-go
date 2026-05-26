//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package workspaceio

import "context"

type workspaceCtxKey struct{}

// WithWorkspace binds w to ctx so callbacks down the chain can resolve
// it via WorkspaceFromContext. Returns ctx unchanged when w is nil so
// callers don't accidentally mask "no executor configured" with a
// non-nil sentinel.
func WithWorkspace(ctx context.Context, w *Workspace) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// WorkspaceFromContext returns the Workspace previously bound via
// WithWorkspace. The boolean return is false when no Workspace was
// installed for the current invocation, allowing callers to gracefully
// fall back when the agent has no code executor configured.
func WorkspaceFromContext(ctx context.Context) (*Workspace, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
