//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

// sanitizeForCacheKey removes volatile or unsafe keys from a State map so that
// the cache key depends only on deterministic inputs.
func sanitizeForCacheKey(input any) any { _ = "STUB: not implemented"; return *new(any) }

// toCanonicalValue converts an arbitrary value into a form that produces
// deterministic JSON when marshaled: maps are converted to sorted key-value arrays,
// slices are canonicalized element-wise.
func toCanonicalValue(v any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// Unwrap State for convenience.

// Only handle map[string]any and map[string]T consistently; otherwise stringify keys.

// Best-effort: iterate exported fields in name order.

// unexported

// Primitive types are fine as-is.

func canonicalizeMap(m map[string]any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type kv struct {
	K string `json:"k"`
	V any    `json:"v"`
}
