//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package messagemerger

import "trpc.group/trpc-go/trpc-agent-go/model"

func mergeConsecutiveMessages(
	messages []model.Message,
	separator string,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func canMergeConsecutiveMessage(msg model.Message) bool { _ = "STUB: not implemented"; return false }

func mergeMessage(
	dst model.Message,
	src model.Message,
	separator string,
) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}

func joinMessageText(first, second, separator string) string { _ = "STUB: not implemented"; return "" }

func cloneMessage(msg model.Message) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}

func mergeMessageContentParts(
	dst model.Message,
	src model.Message,
	separator string,
) []model.ContentPart {
	_ = "STUB: not implemented"
	return nil
}

func orderedMessageContentParts(msg model.Message) []model.ContentPart {
	_ = "STUB: not implemented"
	return nil
}

func shouldInsertMessageSeparator(
	dst model.Message,
	src model.Message,
	separator string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func messageStartsWithText(msg model.Message) bool { _ = "STUB: not implemented"; return false }

func messageEndsWithText(msg model.Message) bool { _ = "STUB: not implemented"; return false }

func textContentPart(text string) model.ContentPart {
	_ = "STUB: not implemented"
	return *new(model.ContentPart)
}
