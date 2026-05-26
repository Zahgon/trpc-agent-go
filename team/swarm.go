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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

// SwarmConfig defines optional safety limits for swarm-style handoffs.
//
// All fields are optional. A zero value means "no limit" for that field.
type SwarmConfig struct {
	// MaxHandoffs limits how many transfers can happen in a single run.
	MaxHandoffs int

	// NodeTimeout limits how long a single member agent may run after a
	// transfer. A zero value means no per-node timeout.
	NodeTimeout time.Duration

	// RepetitiveHandoffWindow is the sliding window size used to detect
	// repetitive handoff loops. A zero value disables this check.
	RepetitiveHandoffWindow int

	// RepetitiveHandoffMinUnique is the minimum number of unique agents that
	// must appear in the window. If fewer appear, the transfer is rejected.
	// A zero value disables this check.
	RepetitiveHandoffMinUnique int
}

// SwarmHandoffInputArgs describes one Swarm handoff input rewrite.
type SwarmHandoffInputArgs struct {
	FromAgentName   string
	ToAgentName     string
	RootInput       model.Message
	ParentInput     model.Message
	TransferMessage string
}

// SwarmHandoffInputBuilder builds the target agent input message for a
// Swarm handoff.
type SwarmHandoffInputBuilder func(
	ctx context.Context,
	args SwarmHandoffInputArgs,
) (model.Message, error)

type swarmSessionIDArgs struct {
	ParentSessionID string
	TeamName        string
	EntryAgentName  string
	ToAgentName     string
}

type swarmSessionScope int

const (
	swarmSessionScopeDefault swarmSessionScope = iota
	swarmSessionScopeShared
	swarmSessionScopePerAgent
)

type swarmTurnRouting int

const (
	swarmTurnRoutingDefault swarmTurnRouting = iota
	swarmTurnRoutingEntry
	swarmTurnRoutingTargetTakesOver
)

type swarmHandoffPolicy struct {
	sessionScope swarmSessionScope
	turnRouting  swarmTurnRouting
}

// DefaultSwarmConfig returns conservative defaults that prevent unbounded
// transfer loops while keeping behavior predictable.
func DefaultSwarmConfig() SwarmConfig { _ = "STUB: not implemented"; return *new(SwarmConfig) }

func defaultSwarmSessionID(args swarmSessionIDArgs) string { _ = "STUB: not implemented"; return "" }

func encodeSwarmSessionIDPart(part string) string { _ = "STUB: not implemented"; return "" }

func (p swarmHandoffPolicy) normalizedSessionScope() swarmSessionScope {
	_ = "STUB: not implemented"
	return *new(swarmSessionScope)
}

func (p swarmHandoffPolicy) normalizedTurnRouting() swarmTurnRouting {
	_ = "STUB: not implemented"
	return *new(swarmTurnRouting)
}

func (p swarmHandoffPolicy) usesIsolatedSession() bool { _ = "STUB: not implemented"; return false }

func (p swarmHandoffPolicy) targetTakesOver() bool { _ = "STUB: not implemented"; return false }

func (p swarmHandoffPolicy) needsRootState() bool { _ = "STUB: not implemented"; return false }
