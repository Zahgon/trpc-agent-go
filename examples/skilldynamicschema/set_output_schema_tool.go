//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"
	"encoding/json"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

type setOutputSchemaInput struct {
	// Schema is the JSON schema object.
	Schema json.RawMessage `json:"schema,omitempty"`
}

type setOutputSchemaTool struct{}

func (t *setOutputSchemaTool) Declaration() *tool.Declaration {
	_ = "STUB: not implemented"
	return nil
}

func (t *setOutputSchemaTool) Call(ctx context.Context, jsonArgs []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Be tolerant: if a schema is already set for this invocation, treat this
// as a no-op (some models call setter tools redundantly).

// Return a normal tool result (not an error) to avoid noisy framework logs.

var _ tool.Tool = (*setOutputSchemaTool)(nil)
var _ tool.CallableTool = (*setOutputSchemaTool)(nil)
