//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package jsonmap provides helpers for JSON-like map values.
package jsonmap

// Clone returns a best-effort clone of src.
//
// JSON-serializable values are deep-cloned through JSON marshal/unmarshal so
// nested maps and slices are isolated. If src contains values that cannot be
// marshaled as JSON, Clone falls back to a shallow map clone to preserve the
// original values.
func Clone(src map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }
