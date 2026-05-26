//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package llm

import "trpc.group/trpc-go/trpc-agent-go/model"

var (
	// DefaultNumSamples sets the default judge sample count.
	DefaultNumSamples = 1
	// defaultMaxTokens sets the default max tokens for judge generation.
	defaultMaxTokens = 2000
	// defaultTemperature sets the default temperature for judge generation.
	defaultTemperature = 0.8
	// defaultStream sets the default streaming behavior for judge generation.
	defaultStream = false
	// DefaultGeneration sets the default generation configuration for judge.
	DefaultGeneration = model.GenerationConfig{
		MaxTokens:   &defaultMaxTokens,
		Temperature: &defaultTemperature,
		Stream:      defaultStream,
	}
)

// options captures judge model configuration overrides.
type options struct {
	rubrics     []*Rubric               // rubrics is the list of rubrics to use.
	variant     string                  // variant selects OpenAI-compatible model behavior.
	baseURL     string                  // baseURL is a custom base URL for the judge model.
	apiKey      string                  // apiKey is the credential for the judge model provider.
	extraFields map[string]any          // extraFields holds provider-specific extras.
	numSamples  int                     // numSamples is the number of samples to request.
	generation  *model.GenerationConfig // generation configures the judge model generation behavior.
}

// newOptions applies Option overrides on top of sensible defaults.
func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// Option configures judge model settings.
type Option func(*options)

// WithRubrics sets the list of rubrics to use.
func WithRubrics(rubrics []*Rubric) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithVariant sets the OpenAI-compatible variant for the judge model.
func WithVariant(variant string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBaseURL sets a custom base URL for the judge model endpoint.
func WithBaseURL(baseURL string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAPIKey sets the API key used when invoking the judge model provider.
func WithAPIKey(apiKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExtraFields supplies provider-specific parameters for the judge model.
func WithExtraFields(extraFields map[string]any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithNumSamples overrides how many judge samples to collect.
func WithNumSamples(numSamples int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGeneration sets the generation configuration for the judge model.
func WithGeneration(generation *model.GenerationConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
