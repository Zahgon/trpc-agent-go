//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package json defines json-based comparison criteria.
package json

import (
	"encoding/json"
)

// CompareFunc defines custom JSON comparison logic.
type CompareFunc func(actual, expected any) (bool, error)

// JSONCriterion compares two JSON objects using exact matching.
type JSONCriterion struct {
	// Ignore skips comparison when true.
	Ignore bool `json:"ignore,omitempty"`
	// IgnoreTree skips nested keys using a structured tree; true leaf ignores the key and its subtree.
	IgnoreTree map[string]any `json:"ignoreTree,omitempty"`
	// OnlyTree compares only selected nested keys; true leaf compares the key and its subtree.
	OnlyTree map[string]any `json:"onlyTree,omitempty"`
	// MatchStrategy selects the comparison rule.
	MatchStrategy JSONMatchStrategy `json:"matchStrategy,omitempty"`
	// NumberTolerance defines the allowed absolute difference between numeric values. 1e-6 is the default.
	NumberTolerance *float64 `json:"numberTolerance,omitempty"`
	// Valid validates raw JSON content when used by callers that receive unparsed content.
	Valid bool `json:"valid,omitempty"`
	// CompareName selects a registered comparison implementation by name.
	CompareName string `json:"compareName,omitempty"`
	// Compare overrides default comparison when provided.
	Compare CompareFunc `json:"-"`
}

// JSONMatchStrategy enumerates supported JSON comparison strategies.
type JSONMatchStrategy string

const (
	// JSONMatchStrategyExact matches json objects exactly.
	JSONMatchStrategyExact JSONMatchStrategy = "exact"
	// JSONMatchStrategySkip skips JSON value matching.
	JSONMatchStrategySkip JSONMatchStrategy = "skip"
)

// New creates a new JSONCriterion with the provided options.
func New(opt ...Option) *JSONCriterion { _ = "STUB: not implemented"; return nil }

// Match compares two JSON values using custom logic or deep equality with numeric tolerance.
func (j *JSONCriterion) Match(actual, expected any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (j *JSONCriterion) normalizeInputs(actual, expected any) (any, any, error) {
	_ = "STUB: not implemented"
	return *new(any), *new(any), nil
}

func validateRawJSON(value any) error { _ = "STUB: not implemented"; return nil }

func parseRawMessage(value any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func parseRawJSON(raw json.RawMessage) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func matchValueOnlyTree(actual, expected any, onlyTree map[string]any, tolerance float64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func compareValueOnlyTree(actual, expected any, onlyTree map[string]any, tolerance float64) error {
	_ = "STUB: not implemented"
	return nil
}

func compareObjectOnlyTree(actual, expected, onlyTree map[string]any, tolerance float64) error {
	_ = "STUB: not implemented"
	return nil
}

func matchValueIgnoreTree(actual, expected any, ignoreTree map[string]any, tolerance float64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func compareValueIgnoreTree(actual, expected any, ignoreTree map[string]any, tolerance float64) error {
	_ = "STUB: not implemented"
	return nil
}

func compareObjectIgnoreTree(actual, expected, ignoreTree map[string]any, tolerance float64) error {
	_ = "STUB: not implemented"
	return nil
}

// isIgnore checks if a key is in the ignore tree.
func isIgnore(ignoreTree map[string]any, key string) bool { _ = "STUB: not implemented"; return false }

func compareValueExact(actual, expected any, tolerance float64) error {
	_ = "STUB: not implemented"
	return nil
}

func compareObjectExact(actual, expected map[string]any, tolerance float64) error {
	_ = "STUB: not implemented"
	return nil
}

func compareArrayExact(actual, expected []any, tolerance float64) error {
	_ = "STUB: not implemented"
	return nil
}

// equalWithTolerance compares two values and applies numeric tolerance when both are numbers.
func equalWithTolerance(actual, expected any, tolerance float64) bool {
	_ = "STUB: not implemented"
	return false
}

// toFloat converts supported numeric types to float64.
func toFloat(v any) (float64, bool) { _ = "STUB: not implemented"; return 0, false }
