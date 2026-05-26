//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package manager provides asynchronous PromptIter run lifecycle management on top of the synchronous engine.
package manager

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/store"
)

// Manager manages asynchronous PromptIter runs.
type Manager interface {
	// Start creates and starts one asynchronous PromptIter run.
	Start(ctx context.Context, request *engine.RunRequest) (*engine.RunResult, error)
	// Get loads one persisted PromptIter run.
	Get(ctx context.Context, runID string) (*engine.RunResult, error)
	// Cancel cancels one running PromptIter run.
	Cancel(ctx context.Context, runID string) error
	// Close stops active runs and releases manager resources.
	Close() error
}

type manager struct {
	appName              string
	engine               engine.Engine
	store                store.Store
	storedResultSlimming engine.RunResultSlimming
	mu                   sync.Mutex
	cancelFuncs          map[string]context.CancelFunc
	closed               bool
}

// New creates a PromptIter run manager for one app.
func New(appName string, engine engine.Engine, opts ...Option) (Manager, error) {
	_ = "STUB: not implemented"
	return *new(Manager), nil
}

// Start creates and starts one asynchronous PromptIter run.
func (m *manager) Start(ctx context.Context, request *engine.RunRequest) (*engine.RunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get loads one persisted PromptIter run.
func (m *manager) Get(ctx context.Context, runID string) (*engine.RunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cancel cancels one running PromptIter run.
func (m *manager) Cancel(ctx context.Context, runID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Close stops active runs and releases manager resources.
func (m *manager) Close() error { _ = "STUB: not implemented"; return nil }

func (m *manager) run(ctx context.Context, runID string, request *engine.RunRequest) {
	_ = "STUB: not implemented"
	return
}

func (m *manager) slimStoredRun(run *engine.RunResult) *engine.RunResult {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) clearCancel(runID string) { _ = "STUB: not implemented"; return }

func validateRunRequest(request *engine.RunRequest) error { _ = "STUB: not implemented"; return nil }

func cloneRunRequest(request *engine.RunRequest) *engine.RunRequest {
	_ = "STUB: not implemented"
	return nil
}

func validateEvalSetInputs(role string, inputs []engine.EvalSetInput) error {
	_ = "STUB: not implemented"
	return nil
}

func isValidLossHintSeverity(severity promptiter.LossSeverity) bool {
	_ = "STUB: not implemented"
	return false
}

func cloneEvalSetInputs(inputs []engine.EvalSetInput) []engine.EvalSetInput {
	_ = "STUB: not implemented"
	return nil
}
