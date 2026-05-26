//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package length defines length-based content criteria.
package length

// LengthCriterion validates that content length is within a configured range.
type LengthCriterion struct {
	// Ignore skips length validation when true.
	Ignore bool `json:"ignore,omitempty"`
	// Min is the inclusive minimum number of Unicode code points.
	Min *int `json:"min,omitempty"`
	// Max is the inclusive maximum number of Unicode code points.
	Max *int `json:"max,omitempty"`
}

// New creates a LengthCriterion with the provided options.
func New(opt ...Option) *LengthCriterion { _ = "STUB: not implemented"; return nil }

// Match validates that content length is within the configured inclusive range.
func (c *LengthCriterion) Match(content string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *LengthCriterion) validate() error { _ = "STUB: not implemented"; return nil }

func (c *LengthCriterion) rangeString() string { _ = "STUB: not implemented"; return "" }
