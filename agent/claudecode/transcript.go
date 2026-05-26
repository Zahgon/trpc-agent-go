//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package claudecode

import (
	"encoding/json"

	"trpc.group/trpc-go/trpc-agent-go/event"
)

const (
	cliToolTask  = "Task"
	cliToolSkill = "Skill"

	frameworkToolSkillRun = "skill_run"
)

// cliRecord represents one top-level record produced by Claude Code CLI JSON output.
type cliRecord struct {
	Type    string      `json:"type,omitempty"`
	Subtype string      `json:"subtype,omitempty"`
	Result  string      `json:"result,omitempty"`
	Message *cliMessage `json:"message,omitempty"`
}

// cliMessage carries content blocks for an assistant/user message record.
type cliMessage struct {
	Content []*cliContentBlock `json:"content,omitempty"`
}

// cliContentBlock is a single item inside a message content array.
type cliContentBlock struct {
	Type      string          `json:"type,omitempty"`
	ID        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Input     json.RawMessage `json:"input,omitempty"`
	ToolUseID string          `json:"tool_use_id,omitempty"`
	Content   json.RawMessage `json:"content,omitempty"`
}

// toolTextBlock represents one text content block inside a tool_result payload.
type toolTextBlock struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

// taskToolInput is the argument shape for Claude Code Task tool calls.
type taskToolInput struct {
	SubagentType string `json:"subagent_type,omitempty"`
}

// skillToolInput is the argument shape for Claude Code Skill tool calls.
type skillToolInput struct {
	Skill string `json:"skill,omitempty"`
}

// skillRunArgs is the argument shape for framework skill_run events derived from Claude Code Skill tool calls.
type skillRunArgs struct {
	Skill   string `json:"skill"`
	Command string `json:"command"`
}

// parseTranscriptToolEvents parses a Claude Code CLI JSON transcript into tool-call and tool-result events.
//
// It also returns the last transcript "result" field (if present) as the user-visible answer text.
func parseTranscriptToolEvents(stdout []byte, invocationID, author string) ([]*event.Event, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// parseTranscriptRecords decodes the CLI transcript into a list of records.
//
// The CLI supports "json" output (a JSON array) and "stream-json" output (JSONL).
func parseTranscriptRecords(stdout []byte) ([]cliRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// normalizeToolName maps CLI tool names to framework tool names when needed.
func normalizeToolName(name string) string { _ = "STUB: not implemented"; return "" }

// parseTaskSubagentType extracts the transfer target agent name from Task tool arguments.
func parseTaskSubagentType(input json.RawMessage) string { _ = "STUB: not implemented"; return "" }

// newTransferEvent creates an agent.transfer event announcing a sub-agent handoff.
func newTransferEvent(invocationID, author, targetAgent string) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// newToolCallEvent creates a tool-call event for one transcript tool_use block.
func newToolCallEvent(invocationID, author, toolID, toolName string, input json.RawMessage) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// newToolResultEvent creates a tool-result event for one transcript tool_result block.
func newToolResultEvent(invocationID, author, toolID, toolName, result string) *event.Event {
	_ = "STUB: not implemented"
	return nil
}

// normalizeToolInput returns the JSON argument payload for a tool call.
func normalizeToolInput(input json.RawMessage) []byte { _ = "STUB: not implemented"; return nil }

// normalizeToolArguments returns the JSON argument payload for a tool call, applying tool-specific mappings when needed.
func normalizeToolArguments(toolName string, input json.RawMessage) []byte {
	_ = "STUB: not implemented"
	return nil
}

// normalizeSkillRunArguments converts a Claude Code Skill tool call into a framework skill_run argument payload.
func normalizeSkillRunArguments(input json.RawMessage) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// decodeToolResultContent converts tool_result.content into displayable text.
func decodeToolResultContent(raw json.RawMessage) string { _ = "STUB: not implemented"; return "" }

// joinToolTextBlocks concatenates tool_result text blocks using newlines.
func joinToolTextBlocks(blocks []toolTextBlock) string { _ = "STUB: not implemented"; return "" }
