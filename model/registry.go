//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package model

// RegisterModelContextWindow registers a model's context window size.
// This allows users to add custom models or override existing mappings.
func RegisterModelContextWindow(modelName string, contextWindowSize int) {
	_ = "STUB: not implemented"
	return
}

// RegisterModelContextWindows registers multiple models' context window sizes in batch.
// This is more efficient than calling RegisterModelContextWindow multiple times.
func RegisterModelContextWindows(models map[string]int) { _ = "STUB: not implemented"; return }

// LookupModelContextWindow returns a known context window size for the
// given model name. It returns ok=false when the model is unknown.
func LookupModelContextWindow(modelName string) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}
