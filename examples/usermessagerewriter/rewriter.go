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
	"trpc.group/trpc-go/trpc-agent-go/model"
)

const (
	rewritePrefix = "rewrite"
	expandPrefix  = "expand"
)

// rewriteUserMessage implements the runner user message rewriter hook.
func (c *rewriterChat) rewriteUserMessage(
	_ context.Context,
	args *agent.UserMessageRewriteArgs,
) ([]model.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// rewriteMessages turns one raw user input into the persisted message sequence.
func (c *rewriterChat) rewriteMessages(userInput string) []model.Message {
	_ = "STUB: not implemented"
	return nil
}
