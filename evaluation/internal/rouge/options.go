//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package rouge

// options holds internal configuration for ROUGE scoring.
type options struct {
	// rougeTypes holds the requested ROUGE types to compute.
	rougeTypes []string
	// useStemmer enables Porter stemming for tokenization.
	useStemmer bool
	// splitSummaries enables sentence splitting for rougeLsum.
	splitSummaries bool
	// tokenizer overrides the built-in tokenizer when provided.
	tokenizer Tokenizer
}

// newOptions applies functional options to build a scoring configuration.
func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// Option configures ROUGE scoring.
type Option func(*options)

// WithRougeTypes sets the ROUGE types to compute.
func WithRougeTypes(rougeTypes ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStemmer enables or disables Porter stemming in the tokenizer.
func WithStemmer(useStemmer bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithSplitSummaries splits summaries into sentences for rougeLsum.
func WithSplitSummaries(splitSummaries bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTokenizer overrides the built-in tokenizer when provided.
func WithTokenizer(tokenizer Tokenizer) Option { _ = "STUB: not implemented"; return *new(Option) }
