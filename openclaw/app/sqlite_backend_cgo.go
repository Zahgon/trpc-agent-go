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
	_ "github.com/mattn/go-sqlite3"

	"trpc.group/trpc-go/trpc-agent-go/session"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/registry"
)

func newSQLiteSessionBackend(
	deps registry.SessionDeps,
	spec registry.SessionBackendSpec,
) (session.Service, error) {
	_ = "STUB: not implemented"
	return *new(session.Service), nil
}

func ensureSQLiteDir(path string) error { _ = "STUB: not implemented"; return nil }
