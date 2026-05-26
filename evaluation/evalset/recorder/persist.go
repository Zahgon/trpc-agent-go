//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package recorder

import (
	"context"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

type turnToPersist struct {
	appName         string
	evalSetID       string
	evalCaseID      string
	evalMode        evalset.EvalMode
	sessionIn       *evalset.SessionInput
	contextMessages []*model.Message
	invocation      *evalset.Invocation
}

func (t *turnToPersist) lockKey() string { _ = "STUB: not implemented"; return "" }

func (r *Recorder) persistTurn(ctx context.Context, turn *turnToPersist) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Recorder) ensureEvalSet(ctx context.Context, appName, evalSetID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Recorder) appendInvocation(ctx context.Context, turn *turnToPersist) error {
	_ = "STUB: not implemented"
	return nil
}

func sortInvocations(invocations []*evalset.Invocation) { _ = "STUB: not implemented"; return }

func invocationTime(invocation *evalset.Invocation) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func hasInvocation(invocations []*evalset.Invocation, invocationID string) bool {
	_ = "STUB: not implemented"
	return false
}

func conversationByMode(evalCase *evalset.EvalCase, mode evalset.EvalMode) []*evalset.Invocation {
	_ = "STUB: not implemented"
	return nil
}

func appendConversationByMode(evalCase *evalset.EvalCase, mode evalset.EvalMode, invocation *evalset.Invocation) {
	_ = "STUB: not implemented"
	return
}
