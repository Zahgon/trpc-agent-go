//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package json

// defaultNumberTolerance is the default number tolerance.
const defaultNumberTolerance = 1e-6

// options configures JSONCriterion.
type options struct {
	ignore          bool
	ignoreTree      map[string]any
	onlyTree        map[string]any
	matchStrategy   JSONMatchStrategy
	numberTolerance *float64
	valid           bool
	compareName     string
	compare         CompareFunc
}

// newOptions creates a Options with the provided options.
func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// Option is a function that configures JSONCriterion.
type Option func(*options)

// WithIgnore sets the ignore flag.
func WithIgnore(ignore bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithIgnoreTree sets the ignore tree.
func WithIgnoreTree(ignoreTree map[string]any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithOnlyTree sets the only tree.
func WithOnlyTree(onlyTree map[string]any) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMatchStrategy sets the match strategy.
func WithMatchStrategy(matchStrategy JSONMatchStrategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithNumberTolerance sets the number tolerance.
func WithNumberTolerance(tolerance float64) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithValid sets the raw JSON validity flag.
func WithValid(valid bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCompareName sets the name of the registered compare function.
func WithCompareName(compareName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCompare sets the compare function.
func WithCompare(compare CompareFunc) Option { _ = "STUB: not implemented"; return *new(Option) }
