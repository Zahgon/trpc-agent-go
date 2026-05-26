//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package workspacefacade

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

// DefaultArtifactMaxBytes is the default per-file cap for artifact
// publishing flows when the caller does not provide one.
const DefaultArtifactMaxBytes int64 = 64 * 1024 * 1024

// Reasons returned by ArtifactSaveSkipReason; an empty string means
// the current invocation can persist artifacts.
const (
	SaveReasonNoInvocation = "invocation is missing from context"
	SaveReasonNoService    = "artifact service is not configured"
	SaveReasonNoSession    = "session is missing from invocation"
	SaveReasonNoSessionIDs = "session app/user/session IDs are missing"
)

// ArtifactSaveSkipReasonInv is the single source of truth for "can this
// invocation persist artifacts?". ArtifactSaveSkipReason, WithArtifactContext
// and tool/workspaceexec.SupportsArtifactSave all forward to this helper so
// the predicate stays consistent across the codebase.
func ArtifactSaveSkipReasonInv(inv *agent.Invocation) string { _ = "STUB: not implemented"; return "" }

// ArtifactSaveSkipReason returns a non-empty string explaining why the
// current invocation cannot persist artifacts, or "" when persistence
// is supported.
func ArtifactSaveSkipReason(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// WithArtifactContext copies the invocation's artifact service and
// session info onto ctx so codeexecutor backends can persist files.
// Injection is gated on the same completeness check as
// ArtifactSaveSkipReason — if any prerequisite is missing, ctx is
// returned as-is so backends never receive a half-populated session.
func WithArtifactContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
