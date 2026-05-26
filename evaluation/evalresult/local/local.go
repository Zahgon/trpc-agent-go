//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package local provides a local file storage evaluation result manager implementation.
package local

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
)

const (
	defaultTempFileSuffix = ".tmp"
	defaultDirPermission  = 0o755
	defaultFilePermission = 0o644
)

// manager implements evalresult.Manager backed by the local filesystem.
type manager struct {
	mu      sync.RWMutex
	baseDir string
	locator evalresult.Locator
}

// New creates a new local file evaluation result manager.
func New(opt ...evalresult.Option) evalresult.Manager {
	_ = "STUB: not implemented"
	return *new(evalresult.Manager)
}

// Close implements evalresult.Manager.
func (m *manager) Close() error {
	_ = "STUB: not implemented"

	// Save stores an evaluation result.
	// Returns an error if the eval set result is nil or the eval set id is empty.
	return nil
}

func (m *manager) Save(_ context.Context, appName string, evalSetResult *evalresult.EvalSetResult) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Get retrieves an evaluation result by evalSetResultID.
func (m *manager) Get(_ context.Context, appName, evalSetResultID string) (*evalresult.EvalSetResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List returns all available evaluation results.
func (m *manager) List(_ context.Context, appName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// evalSetResultPath builds the path to the EvalSetResult file.
func (m *manager) evalSetResultPath(appName, evalSetResultID string) string {
	_ = "STUB: not implemented"
	return ""
}

// load loads the EvalSetResult from the file system.
func (m *manager) load(appName, evalSetResultID string) (*evalresult.EvalSetResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Keep backward compatibility with legacy string-wrapped results.

// store stores the EvalSetResult to the file system.
func (m *manager) store(appName string, evalSetResult *evalresult.EvalSetResult) error {
	_ = "STUB: not implemented"
	return nil
}
