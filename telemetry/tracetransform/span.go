//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Copyright The OpenTelemetry Authors
// Copyright (C) 2025 Tencent. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
//

// Package tracetransform provides functions to transform OpenTelemetry traces
// into OTLP traces.
package tracetransform // import "go.opentelemetry.io/otel/exporters/otlp/otlptrace/internal/tracetransform"

import (
	"go.opentelemetry.io/otel/codes"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
)

// Spans transforms a slice of OpenTelemetry spans into a slice of OTLP
// ResourceSpans.
func Spans(sdl []tracesdk.ReadOnlySpan) []*tracepb.ResourceSpans {
	_ = "STUB: not implemented"
	return nil
}

// Either the resource or instrumentation scope were unknown.

// The resource was unknown.

// The resource has been seen before. Check if the instrumentation
// library lookup was unknown because if so we need to add it to the
// ResourceSpans. Otherwise, the instrumentation library has already
// been seen and the append we did above will be included it in the
// ScopeSpans reference.

// Transform the categorized map into a slice

// span transforms a Span into an OTLP span.
func span(sd tracesdk.ReadOnlySpan) *tracepb.Span { _ = "STUB: not implemented"; return nil }

func clampUint32(v int) uint32 { _ = "STUB: not implemented"; return 0 }

// nolint: gosec  // Overflow/Underflow checked.

// safeTimeToUint64 safely converts time.UnixNano() (int64) to uint64
func safeTimeToUint64(nanos int64) uint64 { _ = "STUB: not implemented"; return 0 }

// nolint: gosec  // Negative values checked above.

// safeSpanFlagsToUint32 safely converts SpanFlags to uint32
func safeSpanFlagsToUint32(v tracepb.SpanFlags) uint32 { _ = "STUB: not implemented"; return 0 }

// nolint: gosec  // Negative values checked above.

// status transform a span code and message into an OTLP span status.
func status(status codes.Code, message string) *tracepb.Status {
	_ = "STUB: not implemented"
	return nil
}

// links transforms span Links to OTLP span links.
func links(links []tracesdk.Link) []*tracepb.Span_Link { _ = "STUB: not implemented"; return nil }

// This redefinition is necessary to prevent link.*ID[:] copies
// being reused -- in short we need a new link per iteration.

func buildSpanFlags(sc trace.SpanContext) uint32 { _ = "STUB: not implemented"; return 0 }

// spanEvents transforms span Events to an OTLP span events.
func spanEvents(es []tracesdk.Event) []*tracepb.Span_Event { _ = "STUB: not implemented"; return nil }

// Transform message events

// spanKind transforms a SpanKind to an OTLP span kind.
func spanKind(kind trace.SpanKind) tracepb.Span_SpanKind {
	_ = "STUB: not implemented"
	return *new(tracepb.Span_SpanKind)
}
