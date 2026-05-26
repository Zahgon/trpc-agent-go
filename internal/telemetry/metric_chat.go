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
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/session"
	"trpc.group/trpc-go/trpc-agent-go/telemetry/metric/histogram"
	"trpc.group/trpc-go/trpc-agent-go/telemetry/semconv/metrics"
)

var (

	// ChatMeter is the meter used for recording chat-related metrics.
	ChatMeter metric.Meter = MeterProvider.Meter(metrics.MeterNameChat)

	// ChatMetricTRPCAgentGoClientRequestCnt records the number of chat requests made.
	ChatMetricTRPCAgentGoClientRequestCnt metric.Int64Counter
	// ChatMetricGenAIClientTokenUsage records the distribution of token usage (both input and output tokens).
	ChatMetricGenAIClientTokenUsage *histogram.DynamicInt64Histogram
	// ChatMetricGenAIClientOperationDuration records the distribution of total chat operation durations in seconds.
	ChatMetricGenAIClientOperationDuration *histogram.DynamicFloat64Histogram
	// ChatMetricGenAIServerTimeToFirstToken records the distribution of time to first token latency in seconds.
	// This measures the time from request start until the first meaningful response payload is received.
	ChatMetricGenAIServerTimeToFirstToken *histogram.DynamicFloat64Histogram
	// ChatMetricTRPCAgentGoClientTimeToFirstToken records the distribution of time to first token latency in seconds.
	// Note: This metric is reported alongside ChatMetricGenAIServerTimeToFirstToken with the same value.
	ChatMetricTRPCAgentGoClientTimeToFirstToken *histogram.DynamicFloat64Histogram
	// ChatMetricTRPCAgentGoClientTimePerOutputToken records the distribution of average time per output token in seconds.
	// This metric measures the decode phase performance by calculating (total_duration - time_to_first_token) / (output_tokens - first_token_count).
	ChatMetricTRPCAgentGoClientTimePerOutputToken *histogram.DynamicFloat64Histogram
	// ChatMetricTRPCAgentGoClientOutputTokenPerTime records the distribution of output token per time for client.
	// 1 / ChatMetricTRPCAgentGoClientTimePerOutputToken.
	ChatMetricTRPCAgentGoClientOutputTokenPerTime *histogram.DynamicFloat64Histogram
)

// chatAttributes is the attributes for chat metrics.
type chatAttributes struct {
	RequestModelName  string
	ResponseModelName string
	Stream            bool
	AgentName         string
	TaskType          string

	AppName   string
	UserID    string
	SessionID string

	ErrorType string
	Error     error
}

func (a chatAttributes) toAttributes() []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

// ChatMetricsTracker tracks metrics for a single chat request lifecycle.
type ChatMetricsTracker struct {
	ctx                            context.Context
	start                          time.Time
	isFirstToken                   bool
	firstTokenTimeDuration         time.Duration
	firstCompleteToken             int
	totalCompletionTokens          int
	totalPromptTokens              int
	totalPromptCachedTokens        int
	totalPromptCacheReadTokens     int
	totalPromptCacheCreationTokens int
	lastEvent                      *event.Event

	// Timing tracking for streaming reasoning phases
	firstReasoningTime time.Time
	lastReasoningTime  time.Time

	// TimingInfo is response timing info that will be recorded in session and attached to events
	timingInfo *model.TimingInfo

	// Configuration
	invocation       *agent.Invocation
	sourceInvocation *agent.Invocation
	llmRequest       *model.Request
	taskType         *string
	err              *error // pointer to capture final error
}

// NewChatMetricsTracker creates a new telemetry tracker.
// The timingInfo parameter should be obtained from invocation state to ensure
// only the first LLM call records timing information.
func NewChatMetricsTracker(
	ctx context.Context,
	invocation *agent.Invocation,
	llmRequest *model.Request,
	timingInfo *model.TimingInfo,
	taskType *string,
	err *error,
) *ChatMetricsTracker {
	_ = "STUB: not implemented"
	return nil
}

func metricsSessionView(sess *session.Session) *session.Session {
	_ = "STUB: not implemented"
	return nil
}

func metricsInvocationView(invocation *agent.Invocation) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

// TrackResponse updates telemetry state and timing info for each response chunk.
// This method tracks both token usage metrics and timing information (FirstTokenDuration and ReasoningDuration).
// Call this for each response received from the LLM.
func (t *ChatMetricsTracker) TrackResponse(response *model.Response) {
	_ = "STUB: not implemented"
	return
}

// Track first token timing (for both metrics and timing info)

// Record TTFT only when the first meaningful response payload arrives.

// Update FirstTokenDuration in TimingInfo only if not already recorded (first LLM call only).

// Track token usage

// Track reasoning duration (streaming mode only, first LLM call only)
// Measures from first reasoning chunk to last reasoning chunk

// Track reasoning phase start and continuation

// Reasoning phase ended (first non-reasoning chunk received), record duration

// SetLastEvent updates the last event seen (for extracting response model name and error).
func (t *ChatMetricsTracker) SetLastEvent(evt *event.Event) {
	_ = "STUB: not implemented"

	// FirstTokenTimeDuration returns the time to first token duration.
	return
}

func (t *ChatMetricsTracker) FirstTokenTimeDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// GetTimingInfo returns the current TimingInfo for attaching to responses.
func (t *ChatMetricsTracker) GetTimingInfo() *model.TimingInfo {
	_ = "STUB: not implemented"
	return nil

	// SetInvocationState refreshes invocation-scoped tracking state for later chunks.
}

func (t *ChatMetricsTracker) SetInvocationState(
	invocation *agent.Invocation,
	timingInfo *model.TimingInfo,
) {
	_ = "STUB: not implemented"
	return
}

func firstNonEmptyString(values ...string) string { _ = "STUB: not implemented"; return "" }

func sameMetricsSessionView(previous *session.Session, current *session.Session) bool {
	_ = "STUB: not implemented"
	return false
}

func mergeMetricsSessionView(
	previous *session.Session,
	current *session.Session,
) *session.Session {
	_ = "STUB: not implemented"
	return nil
}

func mergeInvocationForMetrics(
	previous *agent.Invocation,
	current *agent.Invocation,
) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func chatMetricsEnabled() bool { _ = "STUB: not implemented"; return false }

// RecordMetrics returns a defer function that records all telemetry metrics.
// Should be called with defer immediately after creating the tracker.
func (t *ChatMetricsTracker) RecordMetrics() func() { _ = "STUB: not implemented"; return nil }

// Increment chat request counter

// Record chat request duration

// Record time to first token only when a meaningful payload was observed.

// Record input token usage

// Record cached prompt token usage (subset of input tokens)

// Record tokens read from prompt cache (Anthropic)

// Record tokens used to create prompt cache (Anthropic)

// Record output token usage

// Calculate and record derived metrics

// buildAttributes constructs chatAttributes from tracked state.
func (t *ChatMetricsTracker) buildAttributes() chatAttributes {
	_ = "STUB: not implemented"
	return *

	// Extract error
	new(chatAttributes)
}

// Extract request attributes

// Extract invocation attributes (with nil safety)

// Extract response attributes from last event

// recordDerivedMetrics calculates and records time-per-token and token-per-time metrics.
func (t *ChatMetricsTracker) recordDerivedMetrics(otelAttrs []attribute.KeyValue, requestDuration time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Record time per output token

// Record output token per time

// Record time per output token

// Record output token per time
