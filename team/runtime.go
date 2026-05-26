//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package team

import (
	"context"
	"errors"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

type swarmRuntime struct {
	mu           sync.Mutex
	teamName     string
	entryName    string
	cfg          SwarmConfig
	handoff      swarmHandoffPolicy
	inputBuilder SwarmHandoffInputBuilder
	handoffs     int
	recent       []string
	sessions     map[string]*session.Session
	branches     map[string]*session.Session
}

func (sr *swarmRuntime) OnTransfer(
	_ context.Context,
	fromAgent string,
	toAgent string,
) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (sr *swarmRuntime) CustomizeTransferInvocation(
	ctx context.Context,
	source *agent.Invocation,
	target *agent.Invocation,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (sr *swarmRuntime) isolateTargetSession(
	ctx context.Context,
	source *agent.Invocation,
	target *agent.Invocation,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (sr *swarmRuntime) sessionForAgentStart(
	ctx context.Context,
	service session.Service,
	root *session.Session,
	agentName string,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sr *swarmRuntime) sessionForTransferTarget(
	ctx context.Context,
	target *agent.Invocation,
	root *session.Session,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sr *swarmRuntime) perAgentSession(
	ctx context.Context,
	service session.Service,
	root *session.Session,
	toAgent string,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sr *swarmRuntime) newIsolatedSession(
	ctx context.Context,
	service session.Service,
	root *session.Session,
	toAgent string,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sr *swarmRuntime) getOrCreateSession(
	ctx context.Context,
	service session.Service,
	root *session.Session,
	sessionID string,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sr *swarmRuntime) isolatedSessionID(
	root *session.Session,
	toAgent string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (sr *swarmRuntime) OnTransferComplete(
	ctx context.Context,
	source *agent.Invocation,
	target *agent.Invocation,
	targetEvent *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func (sr *swarmRuntime) OnTransferTerminalError(
	ctx context.Context,
	source *agent.Invocation,
	target *agent.Invocation,
	targetEvent *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func (sr *swarmRuntime) saveTransferOwner(
	ctx context.Context,
	source *agent.Invocation,
	target *agent.Invocation,
	targetEvent *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func updateRootSessionState(
	ctx context.Context,
	source *agent.Invocation,
	root *session.Session,
	state session.StateMap,
) {
	_ = "STUB: not implemented"
	return
}

func (sr *swarmRuntime) registerInvocationSession(
	invocationID string,
	branch string,
	sess *session.Session,
) {
	_ = "STUB: not implemented"
	return
}

func (sr *swarmRuntime) sessionForEvent(
	evt *event.Event,
) (*session.Session, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (sr *swarmRuntime) sessionForBranchLocked(
	branch string,
) (*session.Session, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func branchMatchesPrefix(branch string, prefix string) bool {
	_ = "STUB: not implemented"
	return false
}

// RouteEvent routes isolated member events to their member session.
func (sr *swarmRuntime) RouteEvent(
	root *agent.Invocation,
	routeEvt *event.Event,
) (*session.Session, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func sourceAgentName(source *agent.Invocation) string { _ = "STUB: not implemented"; return "" }

func sourceMessage(source *agent.Invocation) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}

func rootMessage(source *agent.Invocation) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}

func rootSession(inv *agent.Invocation) *session.Session { _ = "STUB: not implemented"; return nil }

func sameSession(a *session.Session, b *session.Session) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *Team) prepareSwarmStartSession(
	ctx context.Context,
	invocation *agent.Invocation,
	startAgent agent.Agent,
	swarmRun *swarmRuntime,
) (*session.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func appendCurrentTurnUserEvents(
	ctx context.Context,
	invocation *agent.Invocation,
	target *session.Session,
) error {
	_ = "STUB: not implemented"
	return nil
}

func currentTurnUserEvents(invocation *agent.Invocation) []*event.Event {
	_ = "STUB: not implemented"
	return nil
}

func normalizeHandoffInputMessage(msg model.Message) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}

func ensureSwarmRuntime(
	inv *agent.Invocation,
	teamName string,
	entryName string,
	cfg SwarmConfig,
	handoff swarmHandoffPolicy,
	inputBuilder SwarmHandoffInputBuilder,
) *swarmRuntime {
	_ = "STUB: not implemented"
	return nil
}

func cloneRuntimeStateForSwarm(opts *agent.RunOptions) { _ = "STUB: not implemented"; return }

var (
	errRepetitiveHandoff = errors.New("repetitive handoff detected")
)

func uniqueCount(values []string) int { _ = "STUB: not implemented"; return 0 }

func installSwarmTransferController(opts *agent.RunOptions, next agent.TransferController) {
	_ = "STUB: not implemented"
	return
}

func stripSwarmTransferControllers(controller agent.TransferController) agent.TransferController {
	_ = "STUB: not implemented"
	return *new(agent.TransferController)
}

func composeTransferControllers(
	first agent.TransferController,
	second agent.TransferController,
) agent.TransferController {
	_ = "STUB: not implemented"
	return *new(agent.TransferController)
}

type chainedTransferController struct {
	first  agent.TransferController
	second agent.TransferController
}

func (c chainedTransferController) OnTransfer(
	ctx context.Context,
	fromAgent string,
	toAgent string,
) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (c chainedTransferController) CustomizeTransferInvocation(
	ctx context.Context,
	source *agent.Invocation,
	target *agent.Invocation,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c chainedTransferController) OnTransferComplete(
	ctx context.Context,
	source *agent.Invocation,
	target *agent.Invocation,
	targetEvent *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func (c chainedTransferController) OnTransferTerminalError(
	ctx context.Context,
	source *agent.Invocation,
	target *agent.Invocation,
	targetEvent *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func tighterTimeout(a time.Duration, b time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
