//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package local provides a local file storage manager implementation for evaluation sets.
package local

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
)

const (
	defaultTempFileSuffix = ".tmp"
	defaultDirPermission  = 0o755
	defaultFilePermission = 0o644
)

// manager implements evalset.Manager backed by the local filesystem.
type manager struct {
	mu      sync.RWMutex
	baseDir string
	locator evalset.Locator
}

// New creates a local file evaluation set manager.
func New(opt ...evalset.Option) evalset.Manager {
	_ = "STUB: not implemented"
	return *new(evalset.Manager)
}

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

// List lists all EvalSet ID for the given appName.
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
// If the EvalSet does not exist or the EvalCase already exists, returns an error.
func (m *manager) AddCase(_ context.Context, appName, evalSetID string, evalCase *evalset.EvalCase) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateCase updates an existing EvalCase.
// If the EvalSet does not exist or the EvalCase does not exist, returns an error.
func (m *manager) UpdateCase(_ context.Context, appName, evalSetID string, evalCase *evalset.EvalCase) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteCase deletes the given EvalCase.
// If the EvalSet does not exist or the EvalCase does not exist, returns an error.
func (m *manager) DeleteCase(_ context.Context, appName, evalSetID, evalCaseID string) error {
	_ = "STUB: not implemented"
	return nil
}

// evalSetPath builds the path to the EvalSet file.
func (m *manager) evalSetPath(appName, evalSetID string) string {
	_ = "STUB: not implemented"
	return ""
}

// load loads the EvalSet from the file system.
func (m *manager) load(appName, evalSetID string) (*evalset.EvalSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// store stores the EvalSet to the file system.
func (m *manager) store(appName string, evalSet *evalset.EvalSet) error {
	_ = "STUB: not implemented"
	return nil
}

// remove removes the EvalSet from the file system.
func (m *manager) remove(appName string, evalSetID string) error {
	_ = "STUB: not implemented"
	return nil
}
