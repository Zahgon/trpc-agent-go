//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package agent

import (
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	awaitUserReplyInvocationStateKey = "__await_user_reply_route__"
	awaitUserReplySessionStateKey    = "__trpc_agent_await_user_reply_route__"
	awaitUserReplyEventExtensionKey  = "trpc_agent.await_user_reply_route"
	awaitUserReplyRootLookupStateKey = "__await_user_reply_root_lookup__"
)

// AwaitUserReplyRoute describes the target of the next user turn.
//
// Runner consumes this route once when await-user-reply routing is enabled.
type AwaitUserReplyRoute struct {
	// AgentName is the target agent that should resume on the next user turn.
	AgentName string `json:"agent_name"`
	// LookupPath is the stable agent path Runner uses to resolve the target.
	// The first segment is the root agent lookup key. Remaining segments are
	// sub-agent names along the invocation branch.
	LookupPath string `json:"lookup_path,omitempty"`
}

// MarkAwaitingUserReply marks the current invocation so its next terminal
// response persists a one-shot "resume here on the next user reply" route.
func MarkAwaitingUserReply(inv *Invocation) error { _ = "STUB: not implemented"; return nil }

// CurrentAwaitUserReplyRoute returns the route currently staged on an
// invocation by MarkAwaitingUserReply.
func CurrentAwaitUserReplyRoute(
	inv *Invocation,
) (AwaitUserReplyRoute, bool) {
	_ = "STUB: not implemented"
	return *new(AwaitUserReplyRoute), false
}

// PendingAwaitUserReplyRoute reads a persisted next-user-turn route from the
// session state.
func PendingAwaitUserReplyRoute(
	sess *session.Session,
) (AwaitUserReplyRoute, bool, error) {
	_ = "STUB: not implemented"
	return *new(AwaitUserReplyRoute), false, nil
}

// SetAwaitUserReplyRootLookupName stores the stable root lookup key used to
// resume the current invocation branch on the next user turn.
func SetAwaitUserReplyRootLookupName(inv *Invocation, name string) {
	_ = "STUB: not implemented"
	return
}

// ClearAwaitUserReplyRouteState returns the session update that clears the
// pending next-user-turn route.
func ClearAwaitUserReplyRouteState() session.StateMap {
	_ = "STUB: not implemented"
	return *new(session.StateMap)
}

// State encodes the route into session state.
func (r AwaitUserReplyRoute) State() (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

// AttachEvent stores the route in one event's state delta and extensions.
func (r AwaitUserReplyRoute) AttachEvent(evt *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func attachAwaitUserReplyRoute(inv *Invocation, evt *event.Event) {
	_ = "STUB: not implemented"
	return
}

func normalizeAwaitUserReplyRoute(
	route AwaitUserReplyRoute,
) (AwaitUserReplyRoute, error) {
	_ = "STUB: not implemented"
	return *new(AwaitUserReplyRoute), nil
}

func buildAwaitUserReplyLookupPath(inv *Invocation) string { _ = "STUB: not implemented"; return "" }

func normalizeAwaitUserReplyPath(path string) string { _ = "STUB: not implemented"; return "" }

func splitAwaitUserReplyPath(path string) []string { _ = "STUB: not implemented"; return nil }
