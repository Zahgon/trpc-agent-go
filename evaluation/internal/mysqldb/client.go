//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package mysqldb

import (
	storage "trpc.group/trpc-go/trpc-agent-go/storage/mysql"
)

// BuildClient builds a MySQL client with either DSN or a registered instance name.
func BuildClient(dsn, instanceName string, extraOptions []any) (storage.Client, error) {
	_ = "STUB: not implemented"
	return *new(storage.Client), nil
}

// Priority: dsn > instanceName.

// IsDuplicateEntry reports whether the error is a MySQL duplicate entry error.
func IsDuplicateEntry(err error) bool { _ = "STUB: not implemented"; return false }

// IsDuplicateKeyName reports whether the error is a MySQL duplicate index name error.
func IsDuplicateKeyName(err error) bool { _ = "STUB: not implemented"; return false }
