//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
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

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/tool"

	agenttool "trpc.group/trpc-go/trpc-agent-go/tool/agent"
)

// Team is an agent that coordinates multiple member agents.
//
// Team implements agent.Agent so it can be used anywhere an agent is
// expected.
type Team struct {
	name        string
	description string

	mode        Mode
	coordinator agent.Agent
	entryName   string

	mu sync.RWMutex

	members      []agent.Agent
	memberByName map[string]agent.Agent

	memberToolSet     tool.ToolSet
	swarm             SwarmConfig
	swarmHandoff      swarmHandoffPolicy
	swarmHandoffInput SwarmHandoffInputBuilder
}

// Mode controls how a Team runs.
type Mode int

const (
	// ModeCoordinator uses a coordinator agent to call members as tools and
	// then respond to the user.
	ModeCoordinator Mode = iota

	// ModeSwarm starts from an entry member and lets members transfer control
	// to each other via transfer_to_agent.
	ModeSwarm
)

const (
	// SwarmActiveAgentKeyPrefix is the session state key prefix for storing the active agent in a Swarm team.
	// When a Swarm team member transfers to another member, the target agent name is stored under:
	// SwarmActiveAgentKeyPrefix + teamName.
	// The next user message will start from this agent instead of the entry member.
	SwarmActiveAgentKeyPrefix = "swarm_active_agent:"
	// SwarmTeamNameKey is the session state key for storing the Swarm team name.
	// This is used to identify if a session belongs to a Swarm team.
	SwarmTeamNameKey = "swarm_team_name"
)

const swarmTraceNodeIDKey = "__swarm_trace_node_id__"

var (
	errEmptyTeamName  = errors.New("team name is empty")
	errNilCoordinator = errors.New("coordinator is nil")
)

// New creates a coordinator team.
//
// The coordinator must support dynamic tool sets (LLMAgent does) so Team can
// expose members as AgentTools.
//
// The created Team uses coordinator.Info().Name as its own name.
func New(
	coordinator agent.Agent,
	members []agent.Agent,
	opts ...Option,
) (*Team, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSwarm creates a swarm team.
//
// entryName must be the name of a member in members.
func NewSwarm(
	name string,
	entryName string,
	members []agent.Agent,
	opts ...Option,
) (*Team, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run implements agent.Agent.
func (t *Team) Run(
	ctx context.Context,
	invocation *agent.Invocation,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Team) runCoordinator(
	ctx context.Context,
	invocation *agent.Invocation,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func wrapCoordinatorInvocationState(
	invocation *agent.Invocation,
	src <-chan *event.Event,
) <-chan *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func (t *Team) runSwarm(
	ctx context.Context,
	invocation *agent.Invocation,
) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try to get the active agent from session state (for cross-request transfer).

// If no active agent (either cross-request transfer disabled or no active agent stored),
// fall back to entry member.

func wrapSwarmInvocationState(
	invocation *agent.Invocation,
	src <-chan *event.Event,
) <-chan *event.Event {
	_ = "STUB: not implemented"
	return nil
}

func (t *Team) markSwarmRootSession(
	invocation *agent.Invocation,
	traceRootNodeID string,
) {
	_ = "STUB: not implemented"
	return
}

// getActiveAgent retrieves the active agent from session state for cross-request transfer.
// Returns nil if no active agent is stored or if the stored agent doesn't exist.
func (t *Team) getActiveAgent(invocation *agent.Invocation) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

// Get the active agent name from session state.

// Look up the agent in memberByName.

// Active agent doesn't exist, return nil to fall back to entry member.

// Tools implements agent.Agent.
func (t *Team) Tools() []tool.Tool { _ = "STUB: not implemented"; return nil }

// Info implements agent.Agent.
func (t *Team) Info() agent.Info { _ = "STUB: not implemented"; return *new(agent.Info) }

// SubAgents implements agent.Agent.
func (t *Team) SubAgents() []agent.Agent { _ = "STUB: not implemented"; return nil }

// FindSubAgent implements agent.Agent.
func (t *Team) FindSubAgent(name string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func buildMemberIndex(
	coordinatorName string,
	members []agent.Agent,
) (map[string]agent.Agent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newMemberToolSet(
	cfg memberToolOptions,
	members []agent.Agent,
) tool.ToolSet {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet)
}

func agentToolHistoryScope(scope HistoryScope) agenttool.HistoryScope {
	_ = "STUB: not implemented"
	return *new(agenttool.HistoryScope)
}

type staticToolSet struct {
	name  string
	tools []tool.Tool
}

func (s *staticToolSet) Tools(context.Context) []tool.Tool { _ = "STUB: not implemented"; return nil }

func (s *staticToolSet) Close() error { _ = "STUB: not implemented"; return nil }

func (s *staticToolSet) Name() string { _ = "STUB: not implemented"; return "" }

func wireSwarmRoster(members []agent.Agent) error { _ = "STUB: not implemented"; return nil }

func swarmActiveAgentKey(teamName string) string { _ = "STUB: not implemented"; return "" }
