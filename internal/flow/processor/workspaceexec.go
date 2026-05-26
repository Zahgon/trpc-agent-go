//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package processor

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/skill"
)

const (
	workspaceExecGuidanceHeader       = "Workspace shell guidance:"
	legacyWorkspaceExecGuidanceHeader = "Executor workspace guidance:"
)

type workspaceExecRequestProcessorOptions struct {
	sessionTools     bool
	hasSkillsRepo    bool
	repoResolver     func(*agent.Invocation) skill.Repository
	enabledResolver  func(*agent.Invocation) bool
	sessionsResolver func(*agent.Invocation) bool
}

// WorkspaceExecRequestProcessorOption configures
// WorkspaceExecRequestProcessor.
type WorkspaceExecRequestProcessorOption func(*workspaceExecRequestProcessorOptions)

// WithWorkspaceExecSessionsEnabled tells the processor that the
// workspace_exec companion session tools are registered.
func WithWorkspaceExecSessionsEnabled() WorkspaceExecRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(WorkspaceExecRequestProcessorOption)
}

// WithWorkspaceExecEnabledResolver sets an invocation-aware workspace_exec
// capability resolver.
func WithWorkspaceExecEnabledResolver(
	resolver func(*agent.Invocation) bool,
) WorkspaceExecRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(WorkspaceExecRequestProcessorOption)
}

// WithWorkspaceExecSessionsResolver sets an invocation-aware resolver for
// workspace_exec session helper capability.
func WithWorkspaceExecSessionsResolver(
	resolver func(*agent.Invocation) bool,
) WorkspaceExecRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(WorkspaceExecRequestProcessorOption)
}

// WithWorkspaceExecSkillsRepo indicates that skills are configured, so the
// workspace guidance can mention existing paths under skills/.
func WithWorkspaceExecSkillsRepo() WorkspaceExecRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(WorkspaceExecRequestProcessorOption)
}

// WithWorkspaceExecSkillsRepositoryResolver sets an invocation-aware skills repository resolver.
func WithWorkspaceExecSkillsRepositoryResolver(
	resolver func(*agent.Invocation) skill.Repository,
) WorkspaceExecRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(WorkspaceExecRequestProcessorOption)
}

// WorkspaceExecRequestProcessor injects system guidance for executor-side
// workspace_exec tools independently of skills repo wiring.
type WorkspaceExecRequestProcessor struct {
	sessionTools     bool
	staticSkillsRepo bool
	repoResolver     func(*agent.Invocation) skill.Repository
	enabledResolver  func(*agent.Invocation) bool
	sessionsResolver func(*agent.Invocation) bool
}

// NewWorkspaceExecRequestProcessor creates a new
// WorkspaceExecRequestProcessor.
func NewWorkspaceExecRequestProcessor(
	opts ...WorkspaceExecRequestProcessorOption,
) *WorkspaceExecRequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

// ProcessRequest implements flow.RequestProcessor.
func (p *WorkspaceExecRequestProcessor) ProcessRequest(
	ctx context.Context,
	inv *agent.Invocation,
	req *model.Request,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func hasWorkspaceExecGuidance(content string) bool { _ = "STUB: not implemented"; return false }

func (p *WorkspaceExecRequestProcessor) guidanceText(
	inv *agent.Invocation,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *WorkspaceExecRequestProcessor) enabledForInvocation(
	inv *agent.Invocation,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *WorkspaceExecRequestProcessor) sessionToolsForInvocation(
	inv *agent.Invocation,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (p *WorkspaceExecRequestProcessor) hasSkillsRepo(
	inv *agent.Invocation,
) bool {
	_ = "STUB: not implemented"
	return false
}
