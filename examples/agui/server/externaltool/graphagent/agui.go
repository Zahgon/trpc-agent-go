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

	"github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/types"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/server/agui"
	aguiadapter "trpc.group/trpc-go/trpc-agent-go/server/agui/adapter"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

func newAGUIServer(r runner.Runner, sessionService session.Service) (*agui.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func resolveRunOptions(_ context.Context, input *aguiadapter.RunAgentInput) ([]agent.RunOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func graphRuntimeState(input *aguiadapter.RunAgentInput, last types.Message) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type graphResumeRef struct {
	lineageID    string
	checkpointID string
}

func resumeRefFromForwardedProps(forwardedProps any) (graphResumeRef, error) {
	_ = "STUB: not implemented"
	return *new(graphResumeRef), nil
}

func forwardedString(props map[string]any, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func toolResumeMapFromMessages(messages []types.Message) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
