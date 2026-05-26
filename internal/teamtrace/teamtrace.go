//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package teamtrace provides internal helpers for mounted team node ids.
package teamtrace

import (
	"trpc.group/trpc-go/trpc-agent-go/agent"
)

const memberTraceRootConfigsKey = "__trpc_agent_internal_team_member_trace_root__"

// RootNodeID returns the mounted surface lookup root node id for one team invocation.
func RootNodeID(inv *agent.Invocation, teamName string) string {
	_ = "STUB: not implemented"
	return ""
}

// TraceRootNodeID returns the execution trace root node id for one team invocation.
func TraceRootNodeID(inv *agent.Invocation, teamName string) string {
	_ = "STUB: not implemented"
	return ""
}

// CoordinatorNodeID returns the coordinator node id under one team root.
func CoordinatorNodeID(rootNodeID string) string { _ = "STUB: not implemented"; return "" }

// MemberNodeID returns the member node id under one team root.
func MemberNodeID(rootNodeID string, memberName string) string {
	_ = "STUB: not implemented"
	return ""
}

// WithMemberTraceRoot stores the mounted team root in custom configs.
func WithMemberTraceRoot(cfgs map[string]any, rootNodeID string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// MemberTraceRoot returns the mounted team root from custom configs.
func MemberTraceRoot(cfgs map[string]any) string { _ = "STUB: not implemented"; return "" }

// SetMemberTraceRootForInvocation stores the mounted team root on one invocation.
func SetMemberTraceRootForInvocation(
	inv *agent.Invocation,
	rootNodeID string,
) {
	_ = "STUB: not implemented"
	return
}

// ClearMemberTraceRootForInvocation removes the mounted team root from one invocation.
func ClearMemberTraceRootForInvocation(inv *agent.Invocation) { _ = "STUB: not implemented"; return }

// MemberTraceRootForInvocation returns the mounted team root for one invocation.
func MemberTraceRootForInvocation(inv *agent.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}

func copyConfigs(in map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }
