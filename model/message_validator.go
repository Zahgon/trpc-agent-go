//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package model

const nonEmptyContentPlaceholder = " "

type roleGroupKind uint8

const (
	roleGroupUnknown roleGroupKind = iota
	roleGroupUserTool
	roleGroupAssistant
)

// HasPayload reports whether the message has non-empty Content, ContentParts,
// or ReasoningContent.
func HasPayload(msg Message) bool { _ = "STUB: not implemented"; return false }

func roleGroupOf(role Role) roleGroupKind { _ = "STUB: not implemented"; return *new(roleGroupKind) }

// validateAndFixMessageSequence validates and fixes a message sequence to
// ensure it complies with strict chat API requirements.
//
// Requirements enforced for non-system messages:
// - Roles must be one of system/user/assistant/tool.
// - Messages must alternate between user/tool group and assistant group.
// - The last message must be user or tool.
// - Content must not be empty.
//
// The function operates at the round level: a round starts with a user message
// group and includes everything up to (but excluding) the next user message
// group. A round is either fully kept or fully removed.
func validateAndFixMessageSequence(messages []Message) []Message {
	_ = "STUB: not implemented"
	return nil
}

func removeInvalidRoleMessages(messages []Message) []Message { _ = "STUB: not implemented"; return nil }

func ensureNonEmptyContent(messages []Message) []Message { _ = "STUB: not implemented"; return nil }

func splitIntoUserAnchoredRounds(messages []Message) ([]Message, [][]Message) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterValidRounds(rounds [][]Message) [][]Message { _ = "STUB: not implemented"; return nil }

func isRoundValid(round []Message) bool { _ = "STUB: not implemented"; return false }

func ensureLastMessageIsUserOrTool(messages []Message) []Message {
	_ = "STUB: not implemented"
	return nil
}

// Remove trailing system messages. This ensures the last message is a
// strict user/tool message when possible.

// Remove trailing assistant group.
