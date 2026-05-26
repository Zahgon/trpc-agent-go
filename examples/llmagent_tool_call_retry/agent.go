//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	defaultModelName = "deepseek-v4-flash"
	defaultLocation  = "Shenzhen"
	toolName         = "get_weather"
)

const llmInstruction = `You are a careful weather assistant.

Rules:
1. You must call the get_weather tool exactly once before answering.
2. Use the user's requested location as the tool argument.
3. Do not answer before you receive the tool result.`

func buildAgent(
	modelName string,
	baseURL string,
	apiKey string,
	service *flakyWeatherService,
	retryPolicy *tool.RetryPolicy,
) *llmagent.LLMAgent {
	_ = "STUB: not implemented"
	return nil
}

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
