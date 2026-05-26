//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package inprocess

import (
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

type replyAccumulator struct {
	text     string
	builder  strings.Builder
	seenFull bool
	err      error
}

func (a *replyAccumulator) consume(evt *event.Event) { _ = "STUB: not implemented"; return }

func (a *replyAccumulator) consumeFull(rsp *model.Response) { _ = "STUB: not implemented"; return }

func (a *replyAccumulator) consumeDelta(rsp *model.Response) { _ = "STUB: not implemented"; return }

func trimResult(text string) string { _ = "STUB: not implemented"; return "" }

func summarizeText(text string, limit int) string { _ = "STUB: not implemented"; return "" }
