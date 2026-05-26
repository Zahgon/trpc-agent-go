//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package approval

import (
	"context"
	"encoding/json"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/approval/review"
	guardtranscript "trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/internal/transcript"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

type transcriptRecord struct {
	index    int
	entry    guardtranscript.Entry
	category guardtranscript.Category
}

func (p *Plugin) buildRequest(ctx context.Context, args *tool.BeforeToolArgs) (*review.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Plugin) buildTranscript(ctx context.Context, invocation *agent.Invocation) []review.TranscriptEntry {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) collectTranscriptEntries(invocation *agent.Invocation) []transcriptRecord {
	_ = "STUB: not implemented"
	return nil
}

func (p *Plugin) countTranscriptTokens(ctx context.Context, entry guardtranscript.Entry) int {
	_ = "STUB: not implemented"
	return 0
}

func declarationDescription(declaration *tool.Declaration) string {
	_ = "STUB: not implemented"
	return ""
}

func messageToTranscriptEntries(msg model.Message) []struct {
	entry    guardtranscript.Entry
	category guardtranscript.Category
} {
	_ = "STUB: not implemented"
	return nil
}

func transcriptContent(msg model.Message) string { _ = "STUB: not implemented"; return "" }

func toolCallSummary(toolCall model.ToolCall) string { _ = "STUB: not implemented"; return "" }

func cloneJSON(input []byte) json.RawMessage {
	_ = "STUB: not implemented"
	return *new(json.RawMessage)
}

func compactJSON(input []byte) string { _ = "STUB: not implemented"; return "" }
