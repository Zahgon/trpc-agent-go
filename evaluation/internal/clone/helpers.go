//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package clone provides deep copy helpers for evaluation data structures.
package clone

import (
	"trpc.group/trpc-go/trpc-agent-go/evaluation/epochtime"
)

func cloneEpochTime(src *epochtime.EpochTime) *epochtime.EpochTime {
	_ = "STUB: not implemented"
	return nil
}

func cloneBytes(src []byte) []byte { _ = "STUB: not implemented"; return nil }

func cloneStringSlice(src []string) []string { _ = "STUB: not implemented"; return nil }

func cloneIntPtr(src *int) *int { _ = "STUB: not implemented"; return nil }

func cloneFloat64Ptr(src *float64) *float64 { _ = "STUB: not implemented"; return nil }

func cloneBoolPtr(src *bool) *bool { _ = "STUB: not implemented"; return nil }

func cloneStringPtr(src *string) *string { _ = "STUB: not implemented"; return nil }

func cloneAny(src any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func errNilInput(name string) error { _ = "STUB: not implemented"; return nil }
