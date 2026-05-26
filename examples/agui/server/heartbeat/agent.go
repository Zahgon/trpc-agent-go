//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const agentName = "agui-heartbeat-agent"

func newAgent(
	modelName string,
	generationConfig model.GenerationConfig,
	quietPeriod time.Duration,
) *llmagent.LLMAgent {
	_ = "STUB: not implemented"
	return nil
}
