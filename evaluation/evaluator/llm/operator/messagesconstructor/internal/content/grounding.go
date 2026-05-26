//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package content

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

type groundingToolCall struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Args any    `json:"args,omitempty"`
}

type groundingToolOutput struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Output any    `json:"output,omitempty"`
}

// ExtractGroundingContext formats invocation artifacts into an ADK-style validation context.
func ExtractGroundingContext(actual *evalset.Invocation) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func formatGroundingMessages(messages []*model.Message) string {
	_ = "STUB: not implemented"
	return ""
}

func formatGroundingTools(tools []*evalset.Tool) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func marshalGroundingSection(payload any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
