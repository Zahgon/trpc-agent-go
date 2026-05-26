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
	"errors"
	"sync"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	tracepb "go.opentelemetry.io/proto/otlp/trace/v1"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

var _ trace.SpanExporter = (*exporter)(nil)

type exporter struct {
	client otlptrace.Client

	mu      sync.RWMutex
	started bool

	startOnce sync.Once
	stopOnce  sync.Once
}

func newExporter(ctx context.Context, opts ...otlptracehttp.Option) (*exporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *exporter) ExportSpans(ctx context.Context, ss []trace.ReadOnlySpan) error {
	_ = "STUB: not implemented"
	return nil
}

func transform(ss []*tracepb.ResourceSpans) []*tracepb.ResourceSpans {
	_ = "STUB: not implemented"
	return nil
}

// transformSpan applies langfuse-specific transformations to a span
func transformSpan(span *tracepb.Span) { _ = "STUB: not implemented"; return }

// Find the operation name

func transformInvokeAgent(span *tracepb.Span) { _ = "STUB: not implemented"; return }

// Skip token usage attributes for InvokeAgent observations.
//
// Reason: the top-level span represents the whole trace, and Langfuse aggregates token
// usage across all spans in the trace. We recently started emitting token usage for
// InvokeAgent spans (to support Galileo); previously only Chat spans had token usage.
// Keeping token attributes on InvokeAgent would make Langfuse double count tokens
// compared to the old behavior (Chat-only token accounting).

// transformCallLLM transforms LLM call spans for Langfuse
// llmSpanCollected holds the intermediate values collected from an LLM span's attributes.
type llmSpanCollected struct {
	sessionID          *commonpb.AnyValue
	llmRequest         *string
	llmResponse        *string
	inputMessagesOTel  *string
	outputMessagesOTel *string
	toolDefinitions    *string
	usage              usageDetails
	attrs              []*commonpb.KeyValue // non-LLM attributes to keep
}

// transformCallLLM transforms LLM call spans for Langfuse.
func transformCallLLM(span *tracepb.Span) { _ = "STUB: not implemented"; return }

// observation.input

// observation.output

// observation.model_parameters (generation_config from llm request)

// observation.usage_details

// collectLLMSpanAttributes iterates over the raw span attributes once, collecting
// the pieces needed by subsequent build steps and filtering out OTEL-specific keys.
func collectLLMSpanAttributes(attrs []*commonpb.KeyValue) llmSpanCollected {
	_ = "STUB: not implemented"
	return *new(llmSpanCollected)
}

// buildLLMObservationInput constructs the Langfuse observation.input value from
// collected LLM span data, wrapping tools+messages when both are present.
func buildLLMObservationInput(c llmSpanCollected) string { _ = "STUB: not implemented"; return "" }

func buildLLMObservationOutput(c llmSpanCollected) string { _ = "STUB: not implemented"; return "" }

func otelObservationInput(otelMessages *string) *string { _ = "STUB: not implemented"; return nil }

func otelObservationOutput(otelMessages *string) *string { _ = "STUB: not implemented"; return nil }

// wrapWithToolsIfPresent returns messagesJSON as-is, or wraps it with tool
// definitions into {"tools":..., "messages":...} when toolDefs is non-empty.
func wrapWithToolsIfPresent(messagesJSON string, toolDefs *string) string {
	_ = "STUB: not implemented"
	return ""
}

// extractModelParameters extracts "generation_config" from the LLM request JSON
// and returns it as an observation.model_parameters attribute, or nil.
func extractModelParameters(llmRequest *string) *commonpb.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// stringKV is a helper to build a string-valued KeyValue proto attribute.
func stringKV(key, value string) *commonpb.KeyValue { _ = "STUB: not implemented"; return nil }

// getStringPtr returns a pointer to the string value of v, or nil if v is nil.
func getStringPtr(v *commonpb.AnyValue) *string { _ = "STUB: not implemented"; return nil }

// stringValueOrNA returns the string value of v, or "N/A" if v is nil.
func stringValueOrNA(v *commonpb.AnyValue) string { _ = "STUB: not implemented"; return "" }

func stringPtrValueOrNA(v *string) string { _ = "STUB: not implemented"; return "" }

func truncateObservationInputMessages(raw string) string { _ = "STUB: not implemented"; return "" }

func truncateObservationOutputChoices(raw string) string { _ = "STUB: not implemented"; return "" }

func isOTelMessagesPayload(raw string) bool { _ = "STUB: not implemented"; return false }

func truncateObservationLLMInput(raw string) string { _ = "STUB: not implemented"; return "" }

func truncateObservationLLMResponse(raw string) string { _ = "STUB: not implemented"; return "" }

func truncateObservationJSONLeafValues(raw string) string { _ = "STUB: not implemented"; return "" }

func truncateJSONLeafValue(v any, maxLeafBytes int) any {
	_ = "STUB: not implemented"
	return *new(any)
}

type truncateMessagesPlan struct {
	textLimit   int
	binaryLimit int
}

type observationTelemetryMessage struct {
	Role             model.Role          `json:"role"`
	Content          string              `json:"content,omitempty"`
	ContentParts     []model.ContentPart `json:"content_parts,omitempty"`
	ToolCallID       string              `json:"tool_call_id,omitempty"`
	Name             string              `json:"name,omitempty"`
	ToolCalls        []model.ToolCall    `json:"tool_calls,omitempty"`
	ReasoningContent string              `json:"reasoning_content,omitempty"`
}

type observationTelemetryChoice struct {
	Index        int                         `json:"index"`
	Message      observationTelemetryMessage `json:"message,omitempty"`
	Delta        observationTelemetryMessage `json:"delta,omitempty"`
	FinishReason *string                     `json:"finish_reason,omitempty"`
}

func sanitizeMessagesForObservation(messages []model.Message, plan truncateMessagesPlan) {
	_ = "STUB: not implemented"
	return
}

func sanitizeTelemetryMessagesForObservation(messages []observationTelemetryMessage, plan truncateMessagesPlan) {
	_ = "STUB: not implemented"
	return
}

func sanitizeSingleTelemetryMessageForObservation(msg *observationTelemetryMessage, plan truncateMessagesPlan) {
	_ = "STUB: not implemented"
	return
}

func sanitizeSingleMessageForObservation(msg *model.Message, plan truncateMessagesPlan) {
	_ = "STUB: not implemented"
	return
}

func truncateBytesHeadTail(b []byte, maxBytes int) []byte { _ = "STUB: not implemented"; return nil }

func buildObservationInputPrompt(messagesJSON, toolDefsJSON string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func extractMessagesJSONFromRequestJSON(requestJSON string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// transformExecuteTool transforms tool execution spans for Langfuse
func transformExecuteTool(span *tracepb.Span) { _ = "STUB: not implemented"; return }

// Add observation type

// Process existing attributes

// Skip this attribute (delete it)

// Skip this attribute (delete it)

// Keep other attributes

// use post set session id

// Replace span attributes

// transformWorkflow transforms workflow spans for Langfuse.
func transformWorkflow(span *tracepb.Span) { _ = "STUB: not implemented"; return }

// Add observation type

// Process existing attributes

// Skip this attribute (delete it)

// Skip this attribute (delete it)

// Keep other attributes

// use post set session id

// Replace span attributes

func (e *exporter) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

var errAlreadyStarted = errors.New("already started")

func (e *exporter) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// MarshalLog is the marshaling function used by the logging system to represent this exporter.
func (e *exporter) MarshalLog() any { _ = "STUB: not implemented"; return *new(any) }
