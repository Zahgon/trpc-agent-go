//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package surface provides internal helpers for PromptIter surface semantics.
package surface

import (
	astructure "trpc.group/trpc-go/trpc-agent-go/agent/structure"
)

// IsSupportedType reports whether the surface type is supported by PromptIter.
func IsSupportedType(surfaceType astructure.SurfaceType) bool {
	_ = "STUB: not implemented"
	return false
}

// ValidateValue validates that one surface value matches the target surface type.
func ValidateValue(surfaceType astructure.SurfaceType, value astructure.SurfaceValue) error {
	_ = "STUB: not implemented"
	return nil
}

// BuildIndex validates surfaces and indexes them by surface ID.
func BuildIndex(surfaces []astructure.Surface) (map[string]astructure.Surface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SanitizeValue validates one surface value and removes empty noise fields.
func SanitizeValue(
	surfaceType astructure.SurfaceType,
	value astructure.SurfaceValue,
) (astructure.SurfaceValue, error) {
	_ = "STUB: not implemented"
	return *new(astructure.SurfaceValue), nil
}

// CloneValue deep-copies one supported PromptIter surface value.
func CloneValue(value astructure.SurfaceValue) astructure.SurfaceValue {
	_ = "STUB: not implemented"
	return *new(astructure.SurfaceValue)
}

func cloneText(value *string) *string { _ = "STUB: not implemented"; return nil }

func cloneExamples(examples []astructure.FewShotExample) []astructure.FewShotExample {
	_ = "STUB: not implemented"
	return nil
}

func cloneMessages(messages []astructure.FewShotMessage) []astructure.FewShotMessage {
	_ = "STUB: not implemented"
	return nil
}

func cloneModel(modelValue *astructure.ModelRef) *astructure.ModelRef {
	_ = "STUB: not implemented"
	return nil
}

func isEmptyModel(modelValue *astructure.ModelRef) bool { _ = "STUB: not implemented"; return false }
