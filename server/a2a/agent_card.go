//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package a2a

import (
	"net/http"

	a2aprotocolserver "trpc.group/trpc-go/trpc-a2a-go/server"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const agentCardContentType = "application/json; charset=utf-8"

// AgentCardOption configures optional parameters for NewAgentCard.
type AgentCardOption func(*agentCardOptions)

type agentCardOptions struct {
	tools []tool.Tool
}

// WithCardTools sets the tools to be advertised as skills in the AgentCard.
// Each tool's Declaration (Name, Description) will be converted to an AgentSkill
// with the "tool" tag. A default agent-level skill is always prepended.
//
// If no tools are provided, only the default agent-level skill is created.
func WithCardTools(tools ...tool.Tool) AgentCardOption {
	_ = "STUB: not implemented"
	return *new(AgentCardOption)
}

// NewAgentCard builds a basic AgentCard from explicit metadata.
// Optional AgentCardOption values can be provided to customize the card,
// e.g. WithCardTools to advertise agent tools as skills.
func NewAgentCard(
	name string,
	description string,
	host string,
	streaming bool,
	opts ...AgentCardOption,
) (a2aprotocolserver.AgentCard, error) {
	_ = "STUB: not implemented"
	return *new(a2aprotocolserver.AgentCard), nil
}

// buildSkillsFromCardTools converts tool declarations to AgentSkills.
// It always prepends a default agent-level skill, then appends one skill
// per tool that has a non-nil Declaration.
func buildSkillsFromCardTools(
	tools []tool.Tool,
	agentName string,
	agentDesc string,
) []a2aprotocolserver.AgentSkill {
	_ = "STUB: not implemented"
	return nil
}

// NewAgentCardHandler returns a handler that serves AgentCard snapshots
// provided by getter. The getter can read from any caller-managed state.
func NewAgentCardHandler(getter func() a2aprotocolserver.AgentCard) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func writeAgentCard(
	w http.ResponseWriter,
	r *http.Request,
	getter func() a2aprotocolserver.AgentCard,
) {
	_ = "STUB: not implemented"
	return
}
