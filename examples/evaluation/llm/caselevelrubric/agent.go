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

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

// newTravelAgent builds an agent with weather, news, time, and ticket tools.
func newTravelAgent(modelName string, stream bool) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

type weatherArgs struct {
	City string `json:"city"`
}

type weatherResult struct {
	City      string  `json:"city"`
	Condition string  `json:"condition"`
	TempC     float64 `json:"tempC"`
}

func getWeather(_ context.Context, args weatherArgs) (weatherResult, error) {
	_ = "STUB: not implemented"
	return *new(weatherResult), nil
}

type newsArgs struct {
	City string `json:"city"`
}

type newsResult struct {
	City     string `json:"city"`
	Headline string `json:"headline"`
}

func getNews(_ context.Context, args newsArgs) (newsResult, error) {
	_ = "STUB: not implemented"
	return *new(newsResult), nil
}

type timeArgs struct {
}

type timeResult struct {
	Timestamp string `json:"timestamp"`
}

func getTime(_ context.Context, _ timeArgs) (timeResult, error) {
	_ = "STUB: not implemented"
	return *new(timeResult), nil
}

type ticketArgs struct {
	City string `json:"city"`
	Time string `json:"time"`
}

type ticketResult struct {
	City      string `json:"city"`
	Time      string `json:"time"`
	SeatsLeft int    `json:"seatsLeft"`
}

func getTicket(_ context.Context, args ticketArgs) (ticketResult, error) {
	_ = "STUB: not implemented"
	return *new(ticketResult), nil
}

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(v float64) *float64 { _ = "STUB: not implemented"; return nil }
