//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package sessionroute stores internal session-routing controls for runner
// persistence.
package sessionroute

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const (
	stateKey                    = "__trpc_agent_session_route__"
	currentTurnRouteStatePrefix = "__trpc_agent_current_turn_session_route:"
)

// EventRouter routes an event to a non-root session for persistence.
type EventRouter interface {
	RouteEvent(
		root *agent.Invocation,
		routeEvt *event.Event,
	) (*session.Session, bool)
}

type currentTurnRoute struct {
	TargetAgentName string `json:"targetAgentName"`
	SessionID       string `json:"sessionID"`
}

type controller struct {
	mu      sync.Mutex
	routers []EventRouter
}

// CurrentTurnRouteState returns the root-session state update for a current
// turn session route.
func CurrentTurnRouteState(
	ownerAgentName string,
	targetAgentName string,
	root *session.Session,
	target *session.Session,
) (session.StateMap, error) {
	_ = "STUB: not implemented"
	return *new(session.StateMap), nil
}

// ApplyCurrentTurnRouteState applies a current-turn route state update to root.
func ApplyCurrentTurnRouteState(root *session.Session, state session.StateMap) {
	_ = "STUB: not implemented"
	return
}

// HasCurrentTurnRoute reports whether root stores a current-turn route.
func HasCurrentTurnRoute(ownerAgentName string, root *session.Session) bool {
	_ = "STUB: not implemented"
	return false
}

// ResolveCurrentTurnSession returns the session that should receive this user
// turn before the selected agent runs.
func ResolveCurrentTurnSession(
	ctx context.Context,
	service session.Service,
	root *session.Session,
	owner agent.Agent,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AttachEventRouter attaches an internal event persistence router to the
// invocation and its ancestors.
func AttachEventRouter(inv *agent.Invocation, router EventRouter) {
	_ = "STUB: not implemented"
	return
}

// RouteEvent asks attached routers for event persistence decisions.
func RouteEvent(
	inv *agent.Invocation,
	routeEvt *event.Event,
) (*session.Session, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// SnapshotEventIdentity copies the event fields used for route lookup before
// plugins can replace the event.
func SnapshotEventIdentity(src *event.Event) *event.Event { _ = "STUB: not implemented"; return nil }

func attachEventRouter(inv *agent.Invocation, router EventRouter) {
	_ = "STUB: not implemented"
	return
}

func getOrCreateController(inv *agent.Invocation) *controller {
	_ = "STUB: not implemented"
	return nil
}

func controllerFor(inv *agent.Invocation) (*controller, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *controller) eventRouters() []EventRouter { _ = "STUB: not implemented"; return nil }

func currentTurnRouteStateKey(ownerAgentName string) string { _ = "STUB: not implemented"; return "" }

func sameSession(a *session.Session, b *session.Session) bool {
	_ = "STUB: not implemented"
	return false
}
