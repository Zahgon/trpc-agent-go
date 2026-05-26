//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

import (
	"context"
)

// Interrupt interrupts execution at the current node and returns the provided
// prompt value. On resume, it returns the resume value that was provided.
func Interrupt(ctx context.Context, state State, key string, prompt any) (any, error) {
	_ = "STUB: not implemented"
	// Track which interrupts have been used in this invocation.
	// This allows the same resume value to be returned if the node re-executes.
	return *new(any), nil
}

// Check if we've already used a resume value for this key.

// Return the same value that was used before.

// Check if we're resuming.

// Store the used value and return it.

// Clear the resume value to avoid reusing it for other keys.

// Check if we have a resume map with the specific key.

// Store the used value and return it.

// Clear the specific key to avoid reusing it for other keys.

// Not resuming, so interrupt with the prompt.

// ResumeValue extracts a resume value from the state with type safety.
func ResumeValue[T any](ctx context.Context, state State, key string) (T, bool) {
	_ = "STUB: not implemented"

	// Check direct resume channel first.
	return *new(T), false
}

// Clear the resume value to avoid reusing it.

// Check resume map.

// Clear the specific key to avoid reusing it.

// ResumeValueOrDefault extracts a resume value from the state with a default fallback.
func ResumeValueOrDefault[T any](ctx context.Context, state State, key string, defaultValue T) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// HasResumeValue checks if there's a resume value available for the given key.
func HasResumeValue(state State, key string) bool {
	_ = "STUB: not implemented"
	// Check direct resume channel.
	return false
}

// Check resume map.

// ClearResumeValue clears a specific resume value from the state.
func ClearResumeValue(state State, key string) {
	_ = "STUB: not implemented"
	// Clear from resume map.
	return
}

// ClearAllResumeValues clears all resume values from the state.
func ClearAllResumeValues(state State) { _ = "STUB: not implemented"; return }
