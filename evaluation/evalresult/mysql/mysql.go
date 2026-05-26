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

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/internal/mysqldb"
	storage "trpc.group/trpc-go/trpc-agent-go/storage/mysql"
)

var _ evalresult.Manager = (*manager)(nil)

type manager struct {
	opts   options
	db     storage.Client
	tables mysqldb.Tables
}

// New creates a MySQL-backed eval result manager.
func New(opts ...Option) (evalresult.Manager, error) {
	_ = "STUB: not implemented"
	return *new(evalresult.Manager), nil
}

// Close implements evalresult.Manager.
func (m *manager) Close() error { _ = "STUB: not implemented"; return nil }

// Save upserts an evaluation result into MySQL.
func (m *manager) Save(ctx context.Context, appName string, evalSetResult *evalresult.EvalSetResult) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Get loads an evaluation result from MySQL.
func (m *manager) Get(ctx context.Context, appName, evalSetResultID string) (*evalresult.EvalSetResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List lists evaluation result IDs for the given app from MySQL.
func (m *manager) List(ctx context.Context, appName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
