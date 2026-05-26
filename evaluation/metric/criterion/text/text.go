//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package text defines text comparison criteria.
package text

import (
	clength "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/length"
)

// CompareFunc defines custom text comparison logic.
type CompareFunc func(actual, expected string) (bool, error)

// TextCriterion governs how two strings should be compared.
type TextCriterion struct {
	// Ignore skips comparison when true.
	Ignore bool `json:"ignore,omitempty"`
	// CaseInsensitive toggles lowercase comparison.
	CaseInsensitive bool `json:"caseInsensitive,omitempty"`
	// MatchStrategy selects the comparison rule.
	MatchStrategy TextMatchStrategy `json:"matchStrategy,omitempty"`
	// Length validates text length.
	Length *clength.LengthCriterion `json:"length,omitempty"`
	// CompareName selects a registered comparison implementation by name.
	CompareName string `json:"compareName,omitempty"`
	// Compare overrides built-in strategies.
	Compare CompareFunc `json:"-"`
}

// TextMatchStrategy enumerates supported text comparison strategies.
type TextMatchStrategy string

const (
	// TextMatchStrategyExact matches strings exactly.
	TextMatchStrategyExact TextMatchStrategy = "exact"
	// TextMatchStrategyContains matches strings that contain the target.
	TextMatchStrategyContains TextMatchStrategy = "contains"
	// TextMatchStrategyRegex matches strings that match the regex.
	TextMatchStrategyRegex TextMatchStrategy = "regex"
	// TextMatchStrategySkip skips built-in string matching.
	TextMatchStrategySkip TextMatchStrategy = "skip"
)

// New creates a new TextCriterion with the provided options.
func New(opt ...Option) *TextCriterion { _ = "STUB: not implemented"; return nil }

// Match compares source and target using the configured strategy.
func (t *TextCriterion) Match(source, target string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Default to exact match.
