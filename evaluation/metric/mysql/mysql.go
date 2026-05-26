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

	"trpc.group/trpc-go/trpc-agent-go/evaluation/internal/mysqldb"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
	storage "trpc.group/trpc-go/trpc-agent-go/storage/mysql"
)

var _ metric.Manager = (*manager)(nil)

type manager struct {
	opts   options
	db     storage.Client
	tables mysqldb.Tables
}

// New creates a MySQL-backed metric manager.
func New(opts ...Option) (metric.Manager, error) {
	_ = "STUB: not implemented"
	return *new(metric.Manager), nil
}

// Close implements metric.Manager.
func (m *manager) Close() error { _ = "STUB: not implemented"; return nil }

// List lists metric names for the specified evaluation set from MySQL.
func (m *manager) List(ctx context.Context, appName, evalSetID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get retrieves a metric definition from MySQL.
func (m *manager) Get(ctx context.Context, appName, evalSetID, metricName string) (*metric.EvalMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add inserts a new metric definition into MySQL.
func (m *manager) Add(ctx context.Context, appName, evalSetID string, metricInput *metric.EvalMetric) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete removes a metric definition from MySQL.
func (m *manager) Delete(ctx context.Context, appName, evalSetID, metricName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Update updates an existing metric definition in MySQL.
func (m *manager) Update(ctx context.Context, appName, evalSetID string, metricInput *metric.EvalMetric) error {
	_ = "STUB: not implemented"
	return nil
}
