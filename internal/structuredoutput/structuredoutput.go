//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package structuredoutput builds structured output schema data from Go types.
package structuredoutput

import (
	"reflect"
)

// Name returns the provider-facing structured output schema name.
func Name(name string) string { _ = "STUB: not implemented"; return "" }

// FromType returns the schema name, generated JSON schema, and pointer type.
func FromType(examplePtr any, strict bool) (string, map[string]any, reflect.Type) {
	_ = "STUB: not implemented"
	return "", nil, *new(reflect.Type)
}

// TypeOf returns the pointer type used for typed structured output.
func TypeOf(examplePtr any) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
