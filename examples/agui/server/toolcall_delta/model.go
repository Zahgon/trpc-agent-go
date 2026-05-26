//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/model/openai"
)

func newLLMModel(name string, baseURL string, apiKey string) model.Model {
	_ = "STUB: not implemented"
	return *new(model.Model)
}

func newGenerationConfig(stream bool) model.GenerationConfig {
	_ = "STUB: not implemented"
	return *new(model.GenerationConfig)
}

func buildOpenAIOptions(baseURL string, apiKey string) []openai.Option {
	_ = "STUB: not implemented"
	return nil
}

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
