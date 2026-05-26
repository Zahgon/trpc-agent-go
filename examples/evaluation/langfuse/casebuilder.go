//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
	langfuseeval "trpc.group/trpc-go/trpc-agent-go/server/evaluation/langfuse"
)

func buildCaseSpec(_ context.Context, item *langfuseeval.DatasetItem) (*langfuseeval.CaseSpec, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func requiredStringField(raw any, objectName string, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func expectedToolsFromMetadata(raw any) ([]*evalset.Tool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
