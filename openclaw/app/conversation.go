//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package app

import (
	"trpc.group/trpc-go/trpc-agent-go/openclaw/conversation"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/gateway"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

const includeContentsNone = "none"

func buildConversationRunOptionResolver(
	appName string,
	sessionSvc session.Service,
	historyOpts conversation.HistoryOptions,
) gateway.RunOptionResolver {
	_ = "STUB: not implemented"
	return *new(gateway.RunOptionResolver)
}
