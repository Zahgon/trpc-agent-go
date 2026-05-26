//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package xml defines XML-based content criteria.
package xml

// CompareFunc defines custom XML comparison logic.
type CompareFunc func(actual, expected string) (bool, error)

// XMLMatchStrategy enumerates supported XML matching strategies.
type XMLMatchStrategy string

const (
	// XMLMatchStrategySkip skips XML value matching.
	XMLMatchStrategySkip XMLMatchStrategy = "skip"
)

// XMLCriterion validates XML content.
type XMLCriterion struct {
	// Ignore skips XML validation when true.
	Ignore bool `json:"ignore,omitempty"`
	// Valid validates that the actual content is a well-formed XML document.
	Valid bool `json:"valid,omitempty"`
	// MatchStrategy selects the XML matching rule.
	MatchStrategy XMLMatchStrategy `json:"matchStrategy,omitempty"`
	// Compare overrides default validation when provided.
	Compare CompareFunc `json:"-"`
}

// New creates an XMLCriterion with the provided options.
func New(opt ...Option) *XMLCriterion { _ = "STUB: not implemented"; return nil }

// Match compares or validates XML content using the configured rule.
func (c *XMLCriterion) Match(actual, expected string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func matchValid(content string) (bool, error) { _ = "STUB: not implemented"; return false, nil }
