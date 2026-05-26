//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package extractor

import (
	"trpc.group/trpc-go/trpc-agent-go/memory"
	memorytool "trpc.group/trpc-go/trpc-agent-go/memory/tool"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// backgroundToolCreators maps tool names to their creator functions.
// These are the tools that can be used by the extractor in background.
var backgroundToolCreators = map[string]func() tool.CallableTool{
	memory.AddToolName:    memorytool.NewAddTool,
	memory.UpdateToolName: memorytool.NewUpdateTool,
	memory.DeleteToolName: memorytool.NewDeleteTool,
	memory.ClearToolName:  memorytool.NewClearTool,
}

// filterTools returns a new tool map containing only tools that are
// enabled by the given set. A nil set keeps all tools enabled, while
// a non-nil empty set disables all tools.
func filterTools(
	all map[string]tool.Tool,
	enabled map[string]struct{},
) map[string]tool.Tool {
	_ = "STUB: not implemented"
	return nil
}

// backgroundTools is the pre-built map of background tools for model request.
// These tools are declaration-only and not callable.
var backgroundTools = func() map[string]tool.Tool {
	tools := make(map[string]tool.Tool, len(backgroundToolCreators))
	for name, creator := range backgroundToolCreators {
		t := creator()
		tools[name] = &declarationOnlyTool{decl: t.Declaration()}
	}
	return tools
}()

// declarationOnlyTool is a tool that only provides declaration, not callable.
type declarationOnlyTool struct {
	decl *tool.Declaration
}

// Declaration returns the tool declaration.
func (t *declarationOnlyTool) Declaration() *tool.Declaration {
	_ = "STUB: not implemented"

	// Argument keys for tool calls.
	return nil
}

const (
	argKeyMemory       = "memory"
	argKeyMemoryID     = "memory_id"
	argKeyTopics       = "topics"
	argKeyMemoryKind   = "memory_kind"
	argKeyEventTime    = "event_time"
	argKeyParticipants = "participants"
	argKeyLocation     = "location"
)

// parseToolCallArgs parses tool call arguments and returns a memory operation.
func parseToolCallArgs(toolName string, args map[string]any) *Operation {
	_ = "STUB: not implemented"
	return nil
}

// parseEpisodicArgs extracts episodic memory fields from tool call arguments.
func parseEpisodicArgs(op *Operation, args map[string]any) { _ = "STUB: not implemented"; return }

// toStringSlice converts an any value to []string.
// Always returns an empty slice instead of nil for consistent downstream handling.
func toStringSlice(v any) []string { _ = "STUB: not implemented"; return nil }
