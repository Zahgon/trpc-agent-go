//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package sqldb

// BuildTableName constructs a full table name with optional prefix.
// If prefix is empty, returns the base table name.
// If prefix is provided, automatically adds an underscore separator if not present.
//
// Examples:
//   - BuildTableName("", "session_states") -> "session_states"
//   - BuildTableName("test", "session_states") -> "test_session_states"
//   - BuildTableName("test_", "session_states") -> "test_session_states"
func BuildTableName(prefix, base string) string { _ = "STUB: not implemented"; return "" }

// Automatically add underscore if not present

// BuildIndexName constructs an index name based on table name and suffix.
// The format is: idx_{tableName}_{suffix}
//
// Examples:
//   - BuildIndexName("", "session_states", "unique_active")
//     -> "idx_session_states_unique_active"
//   - BuildIndexName("test", "session_states", "lookup")
//     -> "idx_test_session_states_lookup"
func BuildIndexName(prefix, tableName, suffix string) string { _ = "STUB: not implemented"; return "" }

// BuildAllTableNames builds all table names with the given prefix.
// Returns a map of base table name to full table name.
func BuildAllTableNames(prefix string) map[string]string { _ = "STUB: not implemented"; return nil }

// BuildTableNameWithSchema constructs a full table name with optional schema and prefix.
// This is primarily used by PostgreSQL which supports schema namespaces.
// MySQL typically doesn't use schemas in the same way (databases serve a similar purpose).
//
// Examples:
//   - BuildTableNameWithSchema("", "", "session_states") -> "session_states"
//   - BuildTableNameWithSchema("", "test", "session_states") -> "test_session_states"
//   - BuildTableNameWithSchema("myschema", "", "session_states") -> "myschema.session_states"
//   - BuildTableNameWithSchema("myschema", "test", "session_states") -> "myschema.test_session_states"
func BuildTableNameWithSchema(schema, prefix, base string) string {
	_ = "STUB: not implemented"
	return ""
}

// BuildIndexNameWithSchema constructs an index name based on schema, table name and suffix.
// For PostgreSQL with schema support, the schema part is replaced with underscore to create a valid index name.
// The format is: idx_{schema}_{tableName}_{suffix} (if schema is provided)
//
//	or idx_{tableName}_{suffix} (if schema is empty)
//
// Examples:
//   - BuildIndexNameWithSchema("", "", "session_states", "unique_active")
//     -> "idx_session_states_unique_active"
//   - BuildIndexNameWithSchema("", "test", "session_states", "lookup")
//     -> "idx_test_session_states_lookup"
//   - BuildIndexNameWithSchema("myschema", "", "session_states", "unique_active")
//     -> "idx_myschema_session_states_unique_active"
//   - BuildIndexNameWithSchema("myschema", "test", "session_states", "lookup")
//     -> "idx_myschema_test_session_states_lookup"
func BuildIndexNameWithSchema(schema, prefix, tableName, suffix string) string {
	_ = "STUB: not implemented"
	return ""
}
