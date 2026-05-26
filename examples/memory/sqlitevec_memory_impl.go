//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

//go:build cgo && sqliteveccgo

package util

import (
	"trpc.group/trpc-go/trpc-agent-go/memory"
)

func newSQLiteVecMemoryService(
	cfg MemoryServiceConfig,
) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}
