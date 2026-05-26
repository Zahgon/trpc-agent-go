//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package failover provides a model.Model wrapper that falls back across
// candidates before the first non-error chunk is emitted.
package failover

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

// failoverModel wraps multiple candidate models and falls back before the
// first non-error chunk is observed.
type failoverModel struct {
	candidates []model.Model
}

// New creates a failover model wrapper.
func New(opts ...Option) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

// Info returns the primary candidate model info.
func (m *failoverModel) Info() model.Info { _ = "STUB: not implemented"; return *new(model.Info) }

// GenerateContent implements the model.Model interface.
func (m *failoverModel) GenerateContent(
	ctx context.Context,
	request *model.Request,
) (<-chan *model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GenerateContentIter implements the model.IterModel interface.
func (m *failoverModel) GenerateContentIter(
	ctx context.Context,
	request *model.Request,
) (model.Seq[*model.Response], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type attempt struct {
	index     int
	candidate model.Model
	seq       model.Seq[*model.Response]
	failures  []failureRecord
	cancel    context.CancelFunc
}

func (m *failoverModel) prepareAttempt(
	ctx context.Context,
	request *model.Request,
	startIndex int,
	failures []failureRecord,
) (*attempt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *failoverModel) runAttempts(
	ctx context.Context,
	request *model.Request,
	initialAttempt *attempt,
	yield func(*model.Response) bool,
) {
	_ = "STUB: not implemented"
	return
}

func sequenceForCandidate(
	ctx context.Context,
	candidate model.Model,
	request *model.Request,
) (model.Seq[*model.Response], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func cloneRequest(request *model.Request) (*model.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func hasFailoverResponseError(response *model.Response) bool {
	_ = "STUB: not implemented"
	return false
}
