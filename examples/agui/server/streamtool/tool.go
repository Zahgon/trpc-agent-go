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
	"time"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	countProgressToolName   = "count_progress"
	defaultCountSteps       = 5
	maxCountSteps           = 20
	defaultCountDelayMS     = 200
	minCountDelayMS         = 50
	maxCountDelayMS         = 2000
	defaultCountStepLatency = 200 * time.Millisecond
)

type countProgressArgs struct {
	Steps   int `json:"steps,omitempty" description:"How many steps to count before finishing."`
	DelayMS int `json:"delay_ms,omitempty" description:"Delay in milliseconds between streamed updates."`
}

type countProgressResult struct {
	Completed int `json:"completed"`
	Total     int `json:"total"`
}

func newCountProgressTool() tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

func normalizeCountProgressArgs(args countProgressArgs) countProgressArgs {
	_ = "STUB: not implemented"
	return *new(countProgressArgs)
}

func runCountProgress(ctx context.Context, args countProgressArgs, writer *tool.StreamWriter) {
	_ = "STUB: not implemented"
	return
}

func sendCountProgressUpdate(
	ctx context.Context,
	writer *tool.StreamWriter,
	current int,
	total int,
	delay time.Duration,
) error {
	_ = "STUB: not implemented"
	return nil
}

func sleepWithContext(ctx context.Context, d time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
