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
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/telemetry/metric/histogram"
	"trpc.group/trpc-go/trpc-agent-go/telemetry/semconv/metrics"
)

var (
	// InvokeAgentMeter is the meter used for recording agent invocation metrics.
	InvokeAgentMeter metric.Meter = MeterProvider.Meter(metrics.MeterNameInvokeAgent)

	// InvokeAgentMetricGenAIRequestCnt records the number of invoke agent requests made.
	InvokeAgentMetricGenAIRequestCnt metric.Int64Counter
	// InvokeAgentMetricGenAIClientTokenUsage records the distribution of input and output token usage.
	InvokeAgentMetricGenAIClientTokenUsage *histogram.DynamicInt64Histogram
	// InvokeAgentMetricGenAIClientTimeToFirstToken records the distribution of time to first token latency in seconds.
	InvokeAgentMetricGenAIClientTimeToFirstToken *histogram.DynamicFloat64Histogram
	// InvokeAgentMetricGenAIClientOperationDuration records the distribution of total agent invocation durations in seconds.
	InvokeAgentMetricGenAIClientOperationDuration *histogram.DynamicFloat64Histogram
)

// invokeAgentAttributes is the attributes for invoke agent metrics.
// It is a subset of chat attributes.
type invokeAgentAttributes struct {
	AgentName string
	AgentID   string
	AppName   string
	UserID    string
	System    string
	Stream    bool
	ErrorType string
	Error     error
}

func (a invokeAgentAttributes) toAttributes() []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// InvokeAgentTracker tracks metrics for a single agent invocation lifecycle.
type InvokeAgentTracker struct {
	ctx                    context.Context
	start                  time.Time
	isFirstToken           bool
	firstTokenTimeDuration time.Duration
	totalCompletionTokens  int
	totalPromptTokens      int

	attributes invokeAgentAttributes
}

// NewInvokeAgentTracker creates a new telemetry tracker for agent invocation.
func NewInvokeAgentTracker(
	ctx context.Context,
	invocation *agent.Invocation,
	stream bool,
	err *error,
) *InvokeAgentTracker {
	_ = "STUB: not implemented"
	return nil
}

// TrackResponse updates telemetry state for each response chunk.
func (t *InvokeAgentTracker) TrackResponse(response *model.Response) {
	_ = "STUB: not implemented"
	return
}

// Track token usage

// SetResponseErrorType updates the response error type seen (for extracting error info).
func (t *InvokeAgentTracker) SetResponseErrorType(errorType string) {
	_ = "STUB: not implemented"
	return
}

// RecordMetrics returns a defer function that records all telemetry metrics.
// Should be called with defer immediately after creating the tracker.
func (t *InvokeAgentTracker) RecordMetrics() func() { _ = "STUB: not implemented"; return nil }

// Increment request counter

// Record request duration

// Record time to first token only when a meaningful payload was observed.

// Record input token usage

// Record output token usage

// FirstTokenTimeDuration returns the time to first token duration.
func (t *InvokeAgentTracker) FirstTokenTimeDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
