//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package registry manages the registration and retrieval of evaluators.
package registry

import (
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evaluator"
)

// Registry defines the interface for evaluators registry.
type Registry interface {
	// Register registers an evaluator to the registry.
	Register(name string, e evaluator.Evaluator) error
	// Get retrieves an evaluator by name.
	Get(name string) (evaluator.Evaluator, error)
	// List returns the names of all registered evaluators.
	List() []string
}

// registry is the default implementation of Registry.
type registry struct {
	mu         sync.RWMutex
	evaluators map[string]evaluator.Evaluator
}

// New creates a evaluator registry
func New() Registry { _ = "STUB: not implemented"; return *new(Registry) }

// Register registers an evaluator to the registry.
// Same name evaluator will be overwritten.
func (r *registry) Register(name string, e evaluator.Evaluator) error {
	_ = "STUB: not implemented"
	return nil
}

// Get gets an evaluator by name.
// Returns os.ErrNotExist if the evaluator is not found.
func (r *registry) Get(name string) (evaluator.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(evaluator.Evaluator), nil
}

// List returns the names of all registered evaluators sorted lexicographically.
func (r *registry) List() []string { _ = "STUB: not implemented"; return nil }
