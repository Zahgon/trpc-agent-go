//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package recorder

import (
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

type accumulator struct {
	mu                    sync.Mutex
	finalized             bool
	capturedRunInputs     bool
	hasUserContent        bool
	userContent           model.Message
	hasFinalResponse      bool
	finalResponse         model.Message
	hasRunError           bool
	runError              model.ResponseError
	sessionInputState     map[string]any
	contextMessages       []model.Message
	tools                 []*evalset.Tool
	toolIDIdx             map[string]int
	intermediateResponses []model.Message
}

func newAccumulator() *accumulator { _ = "STUB: not implemented"; return nil }

func (a *accumulator) captureRunInputs(runtimeState map[string]any, contextMessages []model.Message) {
	_ = "STUB: not implemented"
	return
}

func (a *accumulator) setUserContent(msg model.Message) { _ = "STUB: not implemented"; return }

func (a *accumulator) setFinalResponse(msg model.Message) { _ = "STUB: not implemented"; return }

func (a *accumulator) addIntermediateResponse(msg model.Message) { _ = "STUB: not implemented"; return }

func (a *accumulator) setRunError(err model.ResponseError) { _ = "STUB: not implemented"; return }

func (a *accumulator) addToolCall(tc model.ToolCall) { _ = "STUB: not implemented"; return }

func (a *accumulator) addToolResult(toolID, toolName, content string) {
	_ = "STUB: not implemented"
	return
}

func (a *accumulator) isFinalized() bool { _ = "STUB: not implemented"; return false }

func parseToolCallArguments(arguments []byte) any { _ = "STUB: not implemented"; return *new(any) }

func parseToolResultContent(content string) any { _ = "STUB: not implemented"; return *new(any) }

type turnSnapshot struct {
	finalized             bool
	hasUserContent        bool
	userContent           model.Message
	hasFinalResponse      bool
	finalResponse         model.Message
	hasRunError           bool
	runError              model.ResponseError
	sessionInputState     map[string]any
	contextMessages       []model.Message
	tools                 []*evalset.Tool
	intermediateResponses []model.Message
}

func (a *accumulator) finalizeAndSnapshot() turnSnapshot {
	_ = "STUB: not implemented"
	return *new(turnSnapshot)
}

func cloneStateMap(state map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

func cloneValue[T any](name string, value T) T { _ = "STUB: not implemented"; return *new(T) }

func normalizeStateValue(value any) any { _ = "STUB: not implemented"; return *new(any) }
