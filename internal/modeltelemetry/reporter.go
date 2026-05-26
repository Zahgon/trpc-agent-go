//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package modeltelemetry provides opt-in telemetry helpers for direct model usage.
package modeltelemetry

import (
	"context"

	oteltrace "go.opentelemetry.io/otel/trace"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	itelemetry "trpc.group/trpc-go/trpc-agent-go/internal/telemetry"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// Reporter records chat trace and metrics for one direct model call.
type Reporter struct {
	ctx          context.Context
	invocation   *agent.Invocation
	request      *model.Request
	span         oteltrace.Span
	startedSpan  bool
	tracker      *itelemetry.ChatMetricsTracker
	recordMetric func()
	ended        bool
	err          error
}

// StartChat starts opt-in chat telemetry for direct model usage.
func StartChat(
	ctx context.Context,
	llm model.Model,
	request *model.Request,
	enabled bool,
) *Reporter {
	_ = "STUB: not implemented"
	return nil
}

func invocationView(ctx context.Context, llm model.Model) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

// TrackResponse records telemetry state for one model response.
func (r *Reporter) TrackResponse(response *model.Response) { _ = "STUB: not implemented"; return }

// End finishes chat metrics and trace span recording.
func (r *Reporter) End() { _ = "STUB: not implemented"; return }
