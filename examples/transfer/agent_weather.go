//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// createWeatherAgent creates a specialized weather agent.
func (c *transferChat) createWeatherAgent(modelInstance model.Model) agent.Agent {
	_ = "STUB: not implemented"
	// Weather tool.
	return *new(agent.Agent)
}

// getWeather returns weather information for a location.
func (c *transferChat) getWeather(_ context.Context, args weatherArgs) (weatherResult, error) {
	_ = "STUB: not implemented"
	// Simulate weather data based on location.
	return *new(weatherResult), nil
}

// Default response for unknown locations.

// Data structures for weather tool.
type weatherArgs struct {
	Location string `json:"location" jsonschema:"description=The location to get weather for,required"`
}

type weatherResult struct {
	Location       string  `json:"location"`
	Temperature    float64 `json:"temperature"`
	Condition      string  `json:"condition"`
	Humidity       int     `json:"humidity"`
	Recommendation string  `json:"recommendation"`
}
