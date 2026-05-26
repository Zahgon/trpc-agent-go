//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package gateway

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/persona"
)

const personaContextHeader = "Active preset persona for this chat:"

const (
	chatMemoryScopeLabel = "the current chat scope"
	userMemoryScopeLabel = "this user"
)

func (s *Server) injectedContextMessages(
	ctx context.Context,
	userID string,
	sessionID string,
	requestSystemPrompt string,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func requestSystemPromptMessage(prompt string) *model.Message {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) personaContextMessage(
	userID string,
	sessionID string,
) *model.Message {
	_ = "STUB: not implemented"
	return nil
}

func buildPersonaContextText(preset persona.Preset) string { _ = "STUB: not implemented"; return "" }

func (s *Server) memoryFileContextMessages(
	ctx context.Context,
	userID string,
) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) memoryFileContextMessage(
	ctx context.Context,
	appName string,
	userID string,
	scopeLabel string,
	ensure bool,
) *model.Message {
	_ = "STUB: not implemented"
	return nil
}
