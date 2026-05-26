//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package inmemory provides an in-memory metric manager implementation.
package inmemory

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
)

// manager implements metric.Manager backed by in-memory.
// Each API returns deep-copied objects to avoid accidental mutation.
type manager struct {
	mu      sync.RWMutex
	metrics map[string]map[string][]*metric.EvalMetric // appName -> evalSetID -> []*metric.EvalMetric.
}

// New creates a in-memory metric manager.
func New() metric.Manager { _ = "STUB: not implemented"; return *new(metric.Manager) }

// Close implements metric.Manager.
func (m *manager) Close() error {
	_ = "STUB: not implemented"

	// List lists all metric names identified by the given app name and eval set ID.
	return nil
}

func (m *manager) List(_ context.Context, appName, evalSetID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get gets a metric identified by the given app name, eval set ID and metric name.
func (m *manager) Get(_ context.Context, appName, evalSetID, metricName string) (*metric.EvalMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add adds a metric to EvalSet identified by evalSetID.
func (m *manager) Add(_ context.Context, appName, evalSetID string, metricInput *metric.EvalMetric) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete deletes the metric from EvalSet identified by evalSetID and metricName.
func (m *manager) Delete(_ context.Context, appName, evalSetID, metricName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Update updates the metric identified by evalSetID and metric.MetricName.
func (m *manager) Update(_ context.Context, appName, evalSetID string, metricInput *metric.EvalMetric) error {
	_ = "STUB: not implemented"
	return nil
}
