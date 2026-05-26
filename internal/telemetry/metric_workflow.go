//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package telemetry

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"trpc.group/trpc-go/trpc-agent-go/telemetry/metric/histogram"
	"trpc.group/trpc-go/trpc-agent-go/telemetry/semconv/metrics"
)

var (
	// WorkflowMeter is the meter used for recording workflow execution metrics.
	WorkflowMeter metric.Meter = MeterProvider.Meter(metrics.MeterNameWorkflow)

	// WorkflowMetricGenAIClientOperationDuration records graph workflow/node execution durations in seconds.
	WorkflowMetricGenAIClientOperationDuration *histogram.DynamicFloat64Histogram
)

// WorkflowAttributes is the attributes for workflow execution metrics.
type WorkflowAttributes struct {
	System       string
	AppName      string
	UserID       string
	AgentID      string
	AgentName    string
	WorkflowID   string
	WorkflowName string
	WorkflowType string
	Error        error
	ErrorType    string
}

func (a WorkflowAttributes) toAttributes() []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// ReportWorkflowMetrics reports the workflow execution metrics.
func ReportWorkflowMetrics(ctx context.Context, attrs WorkflowAttributes, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}
