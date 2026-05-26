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

const waitToolName = "wait_before_answer"

type waitBeforeAnswerArgs struct {
	Reason string `json:"reason,omitempty" description:"Short reason for waiting before answering."`
}

type waitBeforeAnswerResult struct {
	WaitedMS int64  `json:"waited_ms"`
	Message  string `json:"message"`
}

func newWaitTool(quietPeriod time.Duration) tool.Tool {
	_ = "STUB: not implemented"
	return *new(tool.Tool)
}

func sleepWithContext(ctx context.Context, d time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
