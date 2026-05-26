//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package team

import (
	"errors"

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

var (
	errNilTeam      = errors.New("team is nil")
	errNotSwarmTeam = errors.New("team is not a swarm team")
)

// UpdateSwarmMembers replaces the Swarm Team roster at runtime.
//
// This method only applies to Swarm teams (created with NewSwarm).
//
// It rewires every member's SubAgents list (via SetSubAgents) so members can
// transfer to the updated roster.
func (t *Team) UpdateSwarmMembers(members []agent.Agent) error {
	_ = "STUB: not implemented"
	return nil
}

// AddSwarmMember adds one member into a Swarm Team roster at runtime.
//
// This method only applies to Swarm teams (created with NewSwarm).
func (t *Team) AddSwarmMember(member agent.Agent) error { _ = "STUB: not implemented"; return nil }

// RemoveSwarmMember removes a member by name from a Swarm Team roster.
//
// Returns true if the member was removed.
//
// This method only applies to Swarm teams (created with NewSwarm).
func (t *Team) RemoveSwarmMember(name string) bool { _ = "STUB: not implemented"; return false }

func (t *Team) updateSwarmMembersLocked(members []agent.Agent) error {
	_ = "STUB: not implemented"
	return nil
}
