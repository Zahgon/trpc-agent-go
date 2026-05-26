//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package chat

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	DefaultExitCommand = "/exit"

	defaultPrompt = "You: "
)

type LoopConfig struct {
	Runner        runner.Runner
	UserID        string
	SessionID     string
	Timeout       time.Duration
	ShowInner     bool
	RootAgentName string
	ExitCommand   string
}

func Run(ctx context.Context, cfg LoopConfig) error { _ = "STUB: not implemented"; return nil }

func printEvents(
	eventChannel <-chan *event.Event,
	showInner bool,
	rootAgentName string,
) {
	_ = "STUB: not implemented"
	return
}

func printToolCalls(ev *event.Event, showArgs bool) { _ = "STUB: not implemented"; return }

func recordToolIDs(toolNameByID map[string]string, ev *event.Event) {
	_ = "STUB: not implemented"
	return
}

func printToolResults(
	toolNameByID map[string]string,
	printed map[string]bool,
	ev *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

func firstDelta(ev *event.Event) string { _ = "STUB: not implemented"; return "" }

func firstContent(ev *event.Event) string { _ = "STUB: not implemented"; return "" }
