//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package inmemory provides a in-memory manager implementation for evaluation sets.
package inmemory

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
)

// Manager implements the evalset.Manager interface using in-memory manager.
// Each API returns deep-copied objects to avoid accidental mutation.
type manager struct {
	mu        sync.RWMutex
	evalSets  map[string]map[string]*evalset.EvalSet             // appName -> evalSetID -> EvalSet.
	evalCases map[string]map[string]map[string]*evalset.EvalCase // appName -> evalSetID -> evalCaseID -> EvalCase.
}

// New creates a in-memory evaluation set manager.
func New() evalset.Manager { _ = "STUB: not implemented"; return *new(evalset.Manager) }

// Close implements evalset.Manager.
func (m *manager) Close() error {
	_ = "STUB: not implemented"

	// Get gets an EvalSet identified by evalSetID.
	// Returns an error if the EvalSet does not exist.
	return nil
}

func (m *manager) Get(_ context.Context, appName, evalSetID string) (*evalset.EvalSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates an EvalSet.
// Returns an error if the EvalSet already exists.
func (m *manager) Create(_ context.Context, appName, evalSetID string) (*evalset.EvalSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List lists all EvalSet IDs for the given appName.
// Returns an error if the appName does not exist.
func (m *manager) List(_ context.Context, appName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete deletes EvalSet identified by evalSetID.
// Returns an error if the EvalSet does not exist.
func (m *manager) Delete(_ context.Context, appName, evalSetID string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCase gets an EvalCase.
// Returns an error if the EvalCase does not exist.
func (m *manager) GetCase(_ context.Context, appName, evalSetID, evalCaseID string) (*evalset.EvalCase, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddCase adds the given EvalCase to an existing EvalSet identified by evalSetID.
// Returns an error if the EvalSet does not exist or the EvalCase already exists.
func (m *manager) AddCase(_ context.Context, appName, evalSetID string, evalCase *evalset.EvalCase) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateCase updates an existing EvalCase.
// Returns an error if the EvalSet does not exist or the EvalCase does not exist.
func (m *manager) UpdateCase(_ context.Context, appName, evalSetID string, evalCase *evalset.EvalCase) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteCase deletes the given EvalCase.
// Returns an error if the EvalSet does not exist or the EvalCase does not exist.
func (m *manager) DeleteCase(_ context.Context, appName, evalSetID, evalCaseID string) error {
	_ = "STUB: not implemented"
	return nil
}

// ensureAppExist ensures the app exists.
func (m *manager) ensureAppExist(appName string) { _ = "STUB: not implemented"; return }

// loadEvalSet loads the EvalSet.
func (m *manager) loadEvalSet(appName, evalSetID string) (*evalset.EvalSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// loadEvalCase loads the EvalCase.
func (m *manager) loadEvalCase(appName, evalSetID, evalCaseID string) (*evalset.EvalCase, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
