//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"
	"sync"
)

const llmInstruction = `You are a careful weather assistant.

Rules:
1. You must call the get_weather tool exactly once before answering.
2. Use the user's requested location as the tool argument.
3. After the tool result arrives, answer in one concise sentence.
4. Do not answer before you receive the tool result.`

type weatherArgs struct {
	Location string `json:"location"`
}

type flakyWeatherService struct {
	mu                sync.Mutex
	failuresRemaining int
	attempts          int
}

func (s *flakyWeatherService) Attempts() int { _ = "STUB: not implemented"; return 0 }

func (s *flakyWeatherService) getWeather(
	ctx context.Context,
	args weatherArgs,
) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
