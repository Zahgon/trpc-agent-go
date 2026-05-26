//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package gemini provides Gemini-compatible model implementations.
package gemini

import (
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

// Accumulator accumulates chunks from a stream
type Accumulator struct {
	Model            string
	FullText         strings.Builder
	ReasoningContent strings.Builder
	FinishReason     strings.Builder
	ToolCalls        []model.ToolCall
	Usage            model.Usage
}

// Accumulate builds up the Message incrementally from a model.Response. The Message then can be used as
// any other Message, except with the caveat that the Message.JSON field which normally can be used to inspect
// the JSON sent over the network may not be populated fully.
func (a *Accumulator) Accumulate(resp *model.Response) { _ = "STUB: not implemented"; return }

// BuildResponse builds up the final a model.Response.
func (a *Accumulator) BuildResponse() *model.Response { _ = "STUB: not implemented"; return nil }
