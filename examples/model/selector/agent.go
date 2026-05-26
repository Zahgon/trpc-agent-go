//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/model/openai"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	appName                    = "model-selector-demo"
	agentName                  = "calculator-assistant"
	calculatorCalledStateKey   = "example:model_selector:calculator_called"
	calculatorAgentInstruction = `You are a calculator assistant.
For every user request, call calculator before answering.
After the tool result is available, answer in Chinese in one concise sentence.
Do not calculate mentally or invent results.`
)

func run(ctx context.Context, cfg appConfig) error { _ = "STUB: not implemented"; return nil }

func newRunner(baseModel model.Model) runner.Runner {
	_ = "STUB: not implemented"
	return *new(runner.Runner)
}

func openAIModelOptions() []openai.Option { _ = "STUB: not implemented"; return nil }

func selectByToolState(toolCallModel, finalModel model.Model) agent.ModelSelector {
	_ = "STUB: not implemented"
	return *new(agent.ModelSelector)
}
