//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package sqldb

import (
	"regexp"
)

// tableNamePattern defines the valid characters for table names and prefixes.
// Must start with a letter or underscore, followed by letters, numbers, or underscores.
var tableNamePattern = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

const (
	// maxTableNameLength is the maximum allowed length for table names.
	// MySQL limit is 64 characters, PostgreSQL is 63.
	// We use 64 as it's more permissive.
	maxTableNameLength = 64
)

// ValidateTableName validates a table name or prefix to prevent SQL injection.
// It checks:
//   - Name is not empty
//   - Name length does not exceed maxTableNameLength
//   - Name matches the allowed pattern (alphanumeric and underscore, starting with letter/underscore)
//
// Returns an error if validation fails.
func ValidateTableName(name string) error { _ = "STUB: not implemented"; return nil }

// ValidateTablePrefix validates a table prefix.
// It applies the same rules as ValidateTableName, but allows empty strings.
func ValidateTablePrefix(prefix string) error {
	_ = "STUB: not implemented"
	// Empty prefix is allowed
	return nil
}

// MustValidateTableName is like ValidateTableName but panics on error.
// This is useful for validating constant table names at package initialization.
func MustValidateTableName(name string) { _ = "STUB: not implemented"; return }

// MustValidateTablePrefix is like ValidateTablePrefix but panics on error.
// This is useful for validating prefixes in option functions.
func MustValidateTablePrefix(prefix string) { _ = "STUB: not implemented"; return }
