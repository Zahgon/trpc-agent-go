//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package transcript provides shared transcript shaping helpers for guardrail reviewers.
package transcript

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

// Default transcript shaping limits and markers used by guardrail reviewers.
const (
	DefaultMessageTranscriptBudget = 10000
	DefaultToolTranscriptBudget    = 10000
	DefaultMessageEntryCap         = 2000
	DefaultToolEntryCap            = 1000
	DefaultRecentNonUserEntryLimit = 40
	DefaultOmissionNote            = "[Earlier context omitted.]"
	DefaultTruncatedSuffix         = " [truncated]"
)

// Category identifies the transcript budget bucket for a record.
type Category int

// Transcript record categories used during transcript shaping.
const (
	CategoryMessage Category = iota
	CategoryTool
)

// Entry is a normalized transcript entry passed to reviewers.
type Entry struct {
	Role    model.Role
	Content string
}

// Record preserves the original order and category of a transcript entry.
type Record struct {
	Index    int
	Entry    Entry
	Category Category
}

// Options configures transcript shaping budgets and truncation behavior.
type Options struct {
	MessageTranscriptBudget int
	ToolTranscriptBudget    int
	MessageEntryCap         int
	ToolEntryCap            int
	RecentNonUserEntryLimit int
	OmissionNote            string
	TruncatedSuffix         string
}

// CountTokensFunc counts tokens for a normalized transcript entry.
type CountTokensFunc func(ctx context.Context, entry Entry) int

type preparedRecord struct {
	index     int
	entry     Entry
	category  Category
	tokens    int
	truncated bool
}

// DefaultOptions returns the default transcript shaping configuration.
func DefaultOptions() Options { _ = "STUB: not implemented"; return *new(Options) }

// Build shapes raw transcript records into reviewer-facing transcript entries.
func Build(ctx context.Context, raw []Record, countTokens CountTokensFunc, options Options) []Entry {
	_ = "STUB: not implemented"
	return nil
}

// TruncateContent truncates content to the rune limit and reports whether truncation happened.
func TruncateContent(content string, maxRunes int, suffix string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func normalizeOptions(options Options) Options { _ = "STUB: not implemented"; return *new(Options) }

func prepareRecords(
	ctx context.Context,
	raw []Record,
	countTokens CountTokensFunc,
	opts Options,
) []preparedRecord {
	_ = "STUB: not implemented"
	return nil
}

func selectEntries(records []preparedRecord, opts Options) ([]Entry, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func reversePreparedRecords(records []preparedRecord) { _ = "STUB: not implemented"; return }
