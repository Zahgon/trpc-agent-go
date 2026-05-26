//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package langfuse

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
)

// CaseSpec is the framework-native case definition derived from a dataset item.
type CaseSpec struct {
	DatasetItemID string
	TraceName     string
	TraceInput    any
	SessionID     string
	UserID        string
	TraceMetadata map[string]any
	EvalCase      *evalset.EvalCase
}

// CaseBuilder converts one Langfuse dataset item into a framework-native case specification.
type CaseBuilder func(ctx context.Context, item *DatasetItem) (*CaseSpec, error)

func buildCaseSpec(_ context.Context, item *DatasetItem) (*CaseSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func stringifyValue(fieldName string, raw any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
