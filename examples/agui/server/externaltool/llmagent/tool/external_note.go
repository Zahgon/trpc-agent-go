//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package tool

import (
	"context"

	agenttool "trpc.group/trpc-go/trpc-agent-go/tool"
)

// ExternalNoteName is the caller-executed note tool name.
const ExternalNoteName = "external_note"

func newExternalNoteTool() agenttool.Tool { _ = "STUB: not implemented"; return *new(agenttool.Tool) }

func externalNoteNotImplemented(_ context.Context, args externalNoteArgs) (externalNoteResult, error) {
	_ = "STUB: not implemented"
	return *new(externalNoteResult), nil
}

type externalNoteArgs struct {
	Topic string `json:"topic" description:"The topic that needs an external note."`
}

type externalNoteResult struct {
	Note string `json:"note" description:"The plain text note returned by the caller."`
}
