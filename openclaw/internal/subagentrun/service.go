//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package subagentrun

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	coretaskrun "trpc.group/trpc-go/trpc-agent-go/agent/taskrun"
	taskruninprocess "trpc.group/trpc-go/trpc-agent-go/agent/taskrun/inprocess"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/outbound"
	openclawsubagent "trpc.group/trpc-go/trpc-agent-go/openclaw/subagent"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

type Service struct {
	core   *taskruninprocess.Service
	router *outbound.Router

	mu      sync.RWMutex
	baseCtx context.Context
}

func NewService(
	stateDir string,
	r runner.Runner,
	router *outbound.Router,
) (*Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) Start(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *Service) Close() error { _ = "STUB: not implemented"; return nil }

func (s *Service) Spawn(
	ctx context.Context,
	req SpawnRequest,
) (openclawsubagent.Run, error) {
	_ = "STUB: not implemented"
	return *new(openclawsubagent.Run), nil
}

func runOptionsFromContext(ctx context.Context) []agent.RunOption {
	_ = "STUB: not implemented"
	return nil
}

func runContextFromContext(
	ctx context.Context,
) func(context.Context) context.Context {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) ListForUser(
	userID string,
	filter openclawsubagent.ListFilter,
) []openclawsubagent.Run {
	_ = "STUB: not implemented"
	return nil
}

func (s *Service) GetForUser(
	userID string,
	runID string,
) (*openclawsubagent.Run, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) CancelForUser(
	userID string,
	runID string,
) (*openclawsubagent.Run, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (s *Service) WaitForUser(
	ctx context.Context,
	userID string,
	runID string,
) (*openclawsubagent.Run, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Service) OnRunUpdate(ctx context.Context, run coretaskrun.Run) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) notifyCompletion(run coretaskrun.Run) { _ = "STUB: not implemented"; return }

func (s *Service) notificationContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *Service) started() bool { _ = "STUB: not implemented"; return false }

func formatNotification(run coretaskrun.Run) string { _ = "STUB: not implemented"; return "" }

func notificationDetail(run coretaskrun.Run) string { _ = "STUB: not implemented"; return "" }

func (s *Service) runForUser(
	ctx context.Context,
	userID string,
	runID string,
) (*coretaskrun.Run, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateSpawnRequest(req SpawnRequest) error { _ = "STUB: not implemented"; return nil }

func translateCoreError(err error) error { _ = "STUB: not implemented"; return nil }
