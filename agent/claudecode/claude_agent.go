//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package claudecode

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// claudeCodeAgent invokes a locally installed Claude Code CLI and maps its transcript into trpc-agent-go events.
type claudeCodeAgent struct {
	name          string
	description   string
	bin           string
	args          []string
	env           []string
	workDir       string
	commandRunner commandRunner
	rawOutputHook RawOutputHook
}

// New creates a Claude Code CLI agent with the provided options.
func New(opt ...Option) (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

// Info implements agent.Agent.
func (a *claudeCodeAgent) Info() agent.Info { _ = "STUB: not implemented"; return *new(agent.Info) }

// Tools implements agent.Agent.
func (a *claudeCodeAgent) Tools() []tool.Tool {
	_ = "STUB: not implemented"

	// SubAgents implements agent.Agent.
	return nil
}

func (a *claudeCodeAgent) SubAgents() []agent.Agent {
	_ = "STUB: not implemented"

	// FindSubAgent implements agent.Agent.
	return nil
}

func (a *claudeCodeAgent) FindSubAgent(string) agent.Agent {
	_ = "STUB: not implemented"

	// Run implements agent.Agent.
	return *new(agent.Agent)
}

func (a *claudeCodeAgent) Run(ctx context.Context, invocation *agent.Invocation) (<-chan *event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// runInvocation executes the CLI invocation and emits tool events and the final response event.
func (a *claudeCodeAgent) runInvocation(ctx context.Context, invocation *agent.Invocation, out chan<- *event.Event) {
	_ = "STUB: not implemented"
	return
}

// handleRawOutputHook invokes the configured raw output hook.
func (a *claudeCodeAgent) handleRawOutputHook(
	ctx context.Context,
	invocation *agent.Invocation,
	cliSessionID string,
	stdout []byte,
	stderr []byte,
	runErr error,
) error {
	_ = "STUB: not implemented"
	return nil
}

// emitFlowError emits an error response event and stops further invocation processing.
func (a *claudeCodeAgent) emitFlowError(
	ctx context.Context,
	invocation *agent.Invocation,
	out chan<- *event.Event,
	combined []byte,
	flowErr error,
) {
	_ = "STUB: not implemented"
	return
}

// runWithSession executes the CLI with resume-first semantics and returns stdout/stderr.
func (a *claudeCodeAgent) runWithSession(ctx context.Context, sessionID, prompt string) ([]byte, []byte, error) {
	_ = "STUB: not implemented"
	// Copy base args to avoid mutating shared backing arrays across concurrent invocations.
	return nil, nil, nil
}

// emitEvent forwards an event to the caller and logs emission failures.
func (a *claudeCodeAgent) emitEvent(ctx context.Context, invocation *agent.Invocation, out chan<- *event.Event, evt *event.Event) {
	_ = "STUB: not implemented"
	return
}

// cliSessionID returns a deterministic UUID session id suitable for Claude Code CLI flags.
func cliSessionID(sess *session.Session) string { _ = "STUB: not implemented"; return "" }
