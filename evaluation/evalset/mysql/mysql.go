//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package mysql

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/internal/mysqldb"
	storage "trpc.group/trpc-go/trpc-agent-go/storage/mysql"
)

var _ evalset.Manager = (*manager)(nil)

type manager struct {
	opts   options
	db     storage.Client
	tables mysqldb.Tables
}

// New creates a MySQL-backed eval set manager.
func New(opts ...Option) (evalset.Manager, error) {
	_ = "STUB: not implemented"
	return *new(evalset.Manager), nil
}

// Close implements evalset.Manager.
func (m *manager) Close() error { _ = "STUB: not implemented"; return nil }

// ensureEvalSetExists checks whether the specified eval set exists in MySQL.
func (m *manager) ensureEvalSetExists(ctx context.Context, appName, evalSetID string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get retrieves an evaluation set and its cases from MySQL.
func (m *manager) Get(ctx context.Context, appName, evalSetID string) (*evalset.EvalSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new evaluation set in MySQL.
func (m *manager) Create(ctx context.Context, appName, evalSetID string) (*evalset.EvalSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List lists evaluation set IDs for the given app from MySQL.
func (m *manager) List(ctx context.Context, appName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete deletes an evaluation set and its cases from MySQL.
func (m *manager) Delete(ctx context.Context, appName, evalSetID string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCase retrieves an evaluation case from MySQL.
func (m *manager) GetCase(ctx context.Context, appName, evalSetID, evalCaseID string) (*evalset.EvalCase, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddCase adds a new evaluation case to MySQL.
func (m *manager) AddCase(ctx context.Context, appName, evalSetID string, evalCase *evalset.EvalCase) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateCase updates an existing evaluation case in MySQL.
func (m *manager) UpdateCase(ctx context.Context, appName, evalSetID string, evalCase *evalset.EvalCase) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteCase deletes an evaluation case from MySQL.
func (m *manager) DeleteCase(ctx context.Context, appName, evalSetID, evalCaseID string) error {
	_ = "STUB: not implemented"
	return nil
}
