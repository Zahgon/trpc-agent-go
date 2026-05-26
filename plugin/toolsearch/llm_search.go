//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package toolsearch

import (
	"context"

	oteltrace "go.opentelemetry.io/otel/trace"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	itelemetry "trpc.group/trpc-go/trpc-agent-go/internal/telemetry"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// llmSearch implements the local searcher interface by asking an LLM to pick
// tool names from a provided candidate set.
//
// It intentionally uses the existing request-building/parsing code paths to keep behavior
// consistent with the ToolSearch callback.
type llmSearch struct {
	model        model.Model
	systemPrompt string
}

const defaultSystemPrompt = "Your goal is to select the most relevant tools for answering the user's query."

func newLlmSearch(model model.Model, systemPrompt string) *llmSearch {
	_ = "STUB: not implemented"
	return nil
}

func (s *llmSearch) Search(ctx context.Context, candidates map[string]tool.Tool, query string, topK int) (context.Context, []string, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

type searchToolResponse struct {
	Tools []string `json:"tools"`
}

func searchTools(ctx context.Context, m model.Model, req *model.Request, tools map[string]tool.Tool) (context.Context, []string, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

func invocationFromContextOrNew(ctx context.Context) *agent.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func generateFinalResponse(ctx context.Context, m model.Model, req *model.Request) (*model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractFirstChoiceContent(final *model.Response) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func trackAndTraceToolSearch(
	ctx context.Context,
	span oteltrace.Span,
	tracker *itelemetry.ChatMetricsTracker,
	invocation *agent.Invocation,
	req *model.Request,
	final *model.Response,
	timingInfo *model.TimingInfo,
	startedSpan bool,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Store usage in context

func parseSearchToolResponse(content string) (searchToolResponse, error) {
	_ = "STUB: not implemented"
	return *new(searchToolResponse), nil
}

// Best-effort: extract a JSON object from surrounding text.

func validateAndDedupeSelectedTools(parsed []string, tools map[string]tool.Tool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findLastUserMessage(messages []model.Message) (model.Message, bool) {
	_ = "STUB: not implemented"
	return *new(model.Message), false
}

func renderToolList(tools []tool.Tool) string { _ = "STUB: not implemented"; return "" }

func toolSelectionSchema(tools []tool.Tool) map[string]any { _ = "STUB: not implemented"; return nil }
