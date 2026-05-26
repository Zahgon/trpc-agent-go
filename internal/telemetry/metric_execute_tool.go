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
	// ExecuteToolMeter is the meter used for recording tool execution metrics.
	ExecuteToolMeter = MeterProvider.Meter(metrics.MeterNameExecuteTool)

	// ExecuteToolMetricTRPCAgentGoClientRequestCnt records the number of tool execution requests made.
	ExecuteToolMetricTRPCAgentGoClientRequestCnt metric.Int64Counter
	// ExecuteToolMetricGenAIClientOperationDuration records the distribution of tool execution durations in seconds.
	ExecuteToolMetricGenAIClientOperationDuration *histogram.DynamicFloat64Histogram
)

// ExecuteToolAttributes is the attributes for tool execution metrics.
type ExecuteToolAttributes struct {
	RequestModelName string
	ToolName         string
	AppName          string
	AgentName        string
	UserID           string
	SessionID        string
	Error            error
	ErrorType        string
}

func (a ExecuteToolAttributes) toAttributes() []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// ReportExecuteToolMetrics reports the tool execution metrics.
func ReportExecuteToolMetrics(ctx context.Context, attrs ExecuteToolAttributes, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}
