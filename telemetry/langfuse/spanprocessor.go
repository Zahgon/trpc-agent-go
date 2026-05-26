//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package langfuse

import (
	"context"

	"go.opentelemetry.io/otel/baggage"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func newSpanProcessor(e sdktrace.SpanExporter) sdktrace.SpanProcessor {
	_ = "STUB: not implemented"
	return *new(sdktrace.SpanProcessor)
}

// baggageBatchSpanProcessor wraps a BatchSpanProcessor and copies baggage members
// from the span's parent context onto the span as attributes at start time.
//
// This mirrors the behavior of go.opentelemetry.io/contrib/processors/baggagecopy.
type baggageBatchSpanProcessor struct {
	next sdktrace.SpanProcessor
}

var _ sdktrace.SpanProcessor = (*baggageBatchSpanProcessor)(nil)

func (p *baggageBatchSpanProcessor) OnStart(ctx context.Context, span sdktrace.ReadWriteSpan) {
	_ = "STUB: not implemented"
	return
}

// defaultLangfuseTraceAttributeFilter limits which baggage entries get propagated
// onto all spans as attributes for Langfuse querying/aggregation compatibility.
// https://langfuse.com/integrations/native/opentelemetry#propagating-attributes
//
// Propagated attributes:
// - userId: langfuse.user.id or user.id
// - sessionId: langfuse.session.id or session.id
// - metadata: langfuse.trace.metadata.* (top-level metadata keys)
// - version: langfuse.version
// - release: langfuse.release
// - tags: langfuse.trace.tags
func defaultLangfuseTraceAttributeFilter(member baggage.Member) bool {
	_ = "STUB: not implemented"
	return false
}

// Only propagate top-level metadata keys.
// `traceMetadata` is "langfuse.trace.metadata", so allow "langfuse.trace.metadata.<key>".

func (p *baggageBatchSpanProcessor) OnEnd(span sdktrace.ReadOnlySpan) {
	_ = "STUB: not implemented"
	return
}

func (p *baggageBatchSpanProcessor) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *baggageBatchSpanProcessor) ForceFlush(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
