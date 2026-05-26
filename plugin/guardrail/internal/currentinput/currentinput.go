//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package currentinput provides shared helpers for extracting the latest user input and transcript.
package currentinput

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/model"
	guardtranscript "trpc.group/trpc-go/trpc-agent-go/plugin/guardrail/internal/transcript"
)

// Request contains the latest user input plus supporting transcript evidence.
type Request[T any] struct {
	LastUserInput string
	Transcript    []T
}

// Build extracts the latest user input and supporting transcript for content reviewers.
func Build[T any](
	ctx context.Context,
	messages []model.Message,
	tokenCounter model.TokenCounter,
	mapEntry func(guardtranscript.Entry) T,
) *Request[T] {
	_ = "STUB: not implemented"
	return nil
}

func buildTranscript[T any](
	ctx context.Context,
	rawEntries []guardtranscript.Record,
	tokenCounter model.TokenCounter,
	mapEntry func(guardtranscript.Entry) T,
) []T {
	_ = "STUB: not implemented"
	return nil
}

func countTranscriptTokens(
	ctx context.Context,
	tokenCounter model.TokenCounter,
	entry guardtranscript.Entry,
) int {
	_ = "STUB: not implemented"
	return 0
}

func collectTranscriptEntries(messages []model.Message, excludedUserIndex int) []guardtranscript.Record {
	_ = "STUB: not implemented"
	return nil
}

func extractLastUserInput(messages []model.Message) (string, int) {
	_ = "STUB: not implemented"
	return "", 0
}

func extractMessageText(message model.Message) string { _ = "STUB: not implemented"; return "" }
