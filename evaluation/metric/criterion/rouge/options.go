//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package rouge

type options struct {
	ignore         bool
	rougeType      string
	measure        RougeMeasure
	threshold      Score
	useStemmer     bool
	splitSummaries bool
	tokenizerName  string
	tokenizer      Tokenizer
}

func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// Option configures RougeCriterion.
type Option func(*options)

// WithIgnore sets the ignore flag.
func WithIgnore(ignore bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRougeType sets the ROUGE variant.
func WithRougeType(rougeType string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithMeasure sets the primary ROUGE measure.
func WithMeasure(measure RougeMeasure) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithThreshold sets the minimum score thresholds.
func WithThreshold(threshold Score) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUseStemmer enables Porter stemming for the built-in tokenizer.
func WithUseStemmer(useStemmer bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSplitSummaries enables sentence splitting for rougeLsum.
func WithSplitSummaries(splitSummaries bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTokenizerName sets the name of the registered tokenizer.
func WithTokenizerName(tokenizerName string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTokenizer sets the custom tokenizer.
func WithTokenizer(tokenizer Tokenizer) Option { _ = "STUB: not implemented"; return *new(Option) }
