//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package conversation

import (
	"trpc.group/trpc-go/trpc-agent-go/plugin"
	"trpc.group/trpc-go/trpc-agent-go/session/summary"
)

const pluginName = "openclaw_conversation"

// Plugin persists request-scoped conversation metadata onto user events.
type Plugin struct{}

// Name implements plugin.Plugin.
func (Plugin) Name() string {
	_ = "STUB: not implemented"

	// Register implements plugin.Plugin.
	return ""
}

func (Plugin) Register(r *plugin.Registry) { _ = "STUB: not implemented"; return }

// PreSummaryHook rewrites summary input using persisted speaker metadata
// when available.
func PreSummaryHook(
	in *summary.PreSummaryHookContext,
) error {
	_ = "STUB: not implemented"
	return nil
}
