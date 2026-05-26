//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package local provides a local file storage implementation for metrics.
package local

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
)

const (
	defaultTempFileSuffix = ".tmp"
	defaultDirPermission  = 0o755
	defaultFilePermission = 0o644
)

type manager struct {
	mu      sync.RWMutex
	baseDir string
	locator metric.Locator
}

// New creates a filesystem-backed metric manager.
func New(opts ...metric.Option) metric.Manager {
	_ = "STUB: not implemented"
	return *new(metric.Manager)
}

// Close implements metric.Manager.
func (m *manager) Close() error { _ = "STUB: not implemented"; return nil }

func (m *manager) List(_ context.Context, appName, evalSetID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *manager) Get(_ context.Context, appName, evalSetID, metricName string) (*metric.EvalMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add adds a metric to EvalSet identified by evalSetID.
func (m *manager) Add(ctx context.Context, appName, evalSetID string, metricInput *metric.EvalMetric) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete deletes the metric from EvalSet identified by evalSetID and metricName.
func (m *manager) Delete(ctx context.Context, appName, evalSetID, metricName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Update updates the metric identified by evalSetID and metric.MetricName.
func (m *manager) Update(ctx context.Context, appName, evalSetID string, metricInput *metric.EvalMetric) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) metricFilePath(appName, evalSetID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *manager) load(appName, evalSetID string) ([]*metric.EvalMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *manager) store(appName, evalSetID string, metrics []*metric.EvalMetric) error {
	_ = "STUB: not implemented"
	return nil
}
