//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package xml

type options struct {
	ignore        bool
	valid         bool
	matchStrategy XMLMatchStrategy
	compare       CompareFunc
}

func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// Option configures XMLCriterion.
type Option func(*options)

// WithIgnore sets the ignore flag.
func WithIgnore(ignore bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithValid sets the XML validity flag.
func WithValid(valid bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMatchStrategy sets the XML match strategy.
func WithMatchStrategy(matchStrategy XMLMatchStrategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCompare sets the custom compare function.
func WithCompare(compare CompareFunc) Option { _ = "STUB: not implemented"; return *new(Option) }
