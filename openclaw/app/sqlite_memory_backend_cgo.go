//go:build cgo

//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package app

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"trpc.group/trpc-go/trpc-agent-go/memory"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/registry"
)

func newSQLiteMemoryBackend(
	deps registry.MemoryDeps,
	spec registry.MemoryBackendSpec,
) (memory.Service, error) {
	_ = "STUB: not implemented"
	return *new(memory.Service), nil
}

func openSQLiteMemoryDB(
	path string,
	dsn string,
	missingPathErr string,
) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
