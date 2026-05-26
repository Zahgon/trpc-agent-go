//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package inmemory provides an in-memory storage evaluation result manager implementation.
package inmemory

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
)

// manager implements evalresult.Manager backed by in-memory.
// Each API returns deep-copied objects to avoid accidental mutation.
type manager struct {
	mu             sync.RWMutex
	evalSetResults map[string]map[string]*evalresult.EvalSetResult // appName -> evalSetResultID -> EvalSetResult.
}

// New creates a in-memory evaluation result manager.
func New() evalresult.Manager { _ = "STUB: not implemented"; return *new(evalresult.Manager) }

// Close implements evalresult.Manager.
func (m *manager) Close() error {
	_ = "STUB: not implemented"

	// Save stores a evaluation result keyed by EvalSetResultID.
	// If the eval set result id is empty, it will be generated.
	// Returns an error if the app name is empty or the eval set result is nil or the eval set id is empty.
	return nil
}

func (m *manager) Save(_ context.Context, appName string, evalSetResult *evalresult.EvalSetResult) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Get retrieves evaluation result by evalSetResultID.
func (m *manager) Get(_ context.Context, appName, evalSetResultID string) (*evalresult.EvalSetResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List returns all stored evaluation results.
func (m *manager) List(_ context.Context, appName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
