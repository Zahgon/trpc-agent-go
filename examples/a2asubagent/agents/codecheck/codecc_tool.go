//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"
)

type readSpecArgs struct {
}

type readSpecResult struct {
	Spec string `json:"spec"`
}

func readSpecFile(ctx context.Context, args readSpecArgs) (readSpecResult, error) {
	_ = "STUB: not implemented"
	return *new(readSpecResult), nil
}
