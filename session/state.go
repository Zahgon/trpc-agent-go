//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package session provides state management functionality.
package session

// State prefix constants for different scope levels
const (
	StateAppPrefix  = "app:"
	StateUserPrefix = "user:"
	StateTempPrefix = "temp:"
)

// State maintains the current value and the pending-commit delta.
type State struct {
	// Value stores the current committed state
	Value StateMap `json:"value"`
	// Delta stores the pending changes that haven't been committed
	Delta StateMap `json:"delta"`
}

// NewState creates a new empty State.
func NewState() *State { _ = "STUB: not implemented"; return nil }

// Set sets the value of a key in the state.
func (s *State) Set(key string, value []byte) { _ = "STUB: not implemented"; return }

// Get gets the value of a key in the state.
// Will return the delta value if it exists, otherwise the value.
func (s *State) Get(key string) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }
