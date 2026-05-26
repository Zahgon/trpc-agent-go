//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package text

import clength "trpc.group/trpc-go/trpc-agent-go/evaluation/metric/criterion/length"

type options struct {
	ignore          bool
	caseInsensitive bool
	matchStrategy   TextMatchStrategy
	length          *clength.LengthCriterion
	compareName     string
	compare         CompareFunc
}

func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// Option configures TextCriterion.
type Option func(*options)

// WithIgnore sets the ignore flag.
func WithIgnore(ignore bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCaseInsensitive sets case-insensitive comparison.
func WithCaseInsensitive(caseInsensitive bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithMatchStrategy sets the match strategy.
func WithMatchStrategy(matchStrategy TextMatchStrategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithLengthCriterion sets the length criterion.
func WithLengthCriterion(criterion *clength.LengthCriterion) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCompareName sets the name of the registered compare function.
func WithCompareName(compareName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCompare sets the custom compare function.
func WithCompare(compare CompareFunc) Option { _ = "STUB: not implemented"; return *new(Option) }
