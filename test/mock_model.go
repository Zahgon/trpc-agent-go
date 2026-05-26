//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package e2e

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

type Call struct {
	Responses []*model.Response
}

type QueueModel struct {
	Calls []Call
}

func (m *QueueModel) Push(call Call) { _ = "STUB: not implemented"; return }

func (m *QueueModel) GenerateContent(ctx context.Context, request *model.Request) (<-chan *model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *QueueModel) Info() model.Info { _ = "STUB: not implemented"; return *new(model.Info) }
