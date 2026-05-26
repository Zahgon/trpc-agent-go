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

	"trpc.group/trpc-go/trpc-agent-go/model"
)

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

func newUserPrompt(location string) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}
