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
	"time"
)

// InterruptError represents an interrupt in graph execution that can be resumed.
type InterruptError struct {
	// Value is the value that was passed to interrupt().
	Value any
	// Key is the key that was passed to Interrupt().
	Key string
	// NodeID is the ID of the node where the interrupt occurred.
	NodeID string
	// TaskID is the ID of the task that was interrupted.
	TaskID string
	// Step is the step number when the interrupt occurred.
	Step int
	// Timestamp is when the interrupt occurred.
	Timestamp time.Time
	// Path is the execution path to the interrupted node.
	Path []string

	// SkipRerun controls whether the interrupted node is re-executed when
	// resuming from the interrupt checkpoint. The default behavior (false)
	// re-executes the interrupted node to preserve the semantics of
	// graph.Interrupt, which is typically called from inside the node.
	SkipRerun bool

	// NextNodes overrides checkpoint.NextNodes when set. This is useful for
	// interrupts that happen before a step executes, where the frontier is not
	// derivable from channel availability.
	NextNodes []string
}

// Error returns the error message for the interrupt.
func (g *InterruptError) Error() string { _ = "STUB: not implemented"; return "" }

// ResumeCommand represents a command to resume graph execution.
type ResumeCommand struct {
	// Resume contains values to resume execution with.
	Resume any
	// ResumeMap maps task namespaces to resume values.
	ResumeMap map[string]any
}

// NewResumeCommand creates a new resume command.
func NewResumeCommand() *ResumeCommand { _ = "STUB: not implemented"; return nil }

// WithResume sets the resume value.
func (c *ResumeCommand) WithResume(value any) *ResumeCommand { _ = "STUB: not implemented"; return nil }

// WithResumeMap sets the resume map.
func (c *ResumeCommand) WithResumeMap(resumeMap map[string]any) *ResumeCommand {
	_ = "STUB: not implemented"
	return nil
}

// AddResumeValue adds a resume value for a specific task.
func (c *ResumeCommand) AddResumeValue(taskID string, value any) *ResumeCommand {
	_ = "STUB: not implemented"
	return nil
}

// NewInterruptError creates a new InterruptError with the given value.
func NewInterruptError(value any) *InterruptError { _ = "STUB: not implemented"; return nil }

// IsInterruptError checks if an error is a InterruptError.
func IsInterruptError(err error) bool { _ = "STUB: not implemented"; return false }

// GetInterruptError extracts InterruptError from an error.
func GetInterruptError(err error) (*InterruptError, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
