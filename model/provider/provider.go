//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package provider provides a unified interface for constructing model.Model instances from different providers.
package provider

import (
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

func init() {
	Register("openai", openaiProvider)
	Register("anthropic", anthropicProvider)
	Register("gemini", geminiProvider)
	Register("ollama", ollamaProvider)
	Register("hunyuan", hunyuanProvider)
}

// Provider builds a model.Model instance.
type Provider func(opts *Options) (model.Model, error)

var (
	providersMu sync.RWMutex                // providersMu guards providers access.
	providers   = make(map[string]Provider) // providers stores provider name to provider mappings.
)

// Register registers a provider by name.
func Register(name string, provider Provider) { _ = "STUB: not implemented"; return }

// Get returns the provider by name or nil if not found.
func Get(name string) (Provider, bool) { _ = "STUB: not implemented"; return *new(Provider), false }

// Model constructs a model.Model with the given provider name, model name and options.
func Model(providerName, modelName string, opt ...Option) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

// openaiProvider builds an OpenAI-compatible model instance using the resolved options.
func openaiProvider(opts *Options) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

// anthropicProvider builds an Anthropic-compatible model instance using the resolved options.
func anthropicProvider(opts *Options) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

// geminiProvider builds an Gemini-compatible model instance using the resolved options.
func geminiProvider(opts *Options) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

// ollamaProvider builds an Ollama-compatible model instance using the resolved options.
func ollamaProvider(opts *Options) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

// hunyuanProvider builds a Hunyuan-compatible model instance using the resolved options.
func hunyuanProvider(opts *Options) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}
