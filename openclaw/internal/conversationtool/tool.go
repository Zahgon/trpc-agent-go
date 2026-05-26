//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package conversationtool

import (
	"context"
	"errors"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	toolConversationHistory = "conversation_history"

	defaultTurnLimit = 12
	maxTurnLimit     = 50
)

var (
	errToolNotInInvocation = errors.New(
		"conversation_history: current session is unavailable",
	)
)

// Tool inspects the current conversation session.
type Tool struct{}

// NewTool creates a conversation history inspection tool.
func NewTool() *Tool { _ = "STUB: not implemented"; return nil }

func (t *Tool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }

type toolInput struct {
	Limit         *int `json:"limit,omitempty"`
	IncludeSystem bool `json:"include_system,omitempty"`
}

func (t *Tool) Call(ctx context.Context, args []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func conversationLabelOverrides(
	inv *agent.Invocation,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func normalizeLimit(raw *int) int { _ = "STUB: not implemented"; return 0 }
