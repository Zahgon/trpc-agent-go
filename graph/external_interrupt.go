//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package graph

import (
	"context"
	"sync"
	"time"
)

type graphInterruptKey struct{}

type graphInterruptState struct {
	mu      sync.RWMutex
	timeout *time.Duration

	done chan struct{}
	once sync.Once
}

type graphInterruptOptions struct {
	timeout *time.Duration
}

// GraphInterruptOption configures behavior for WithGraphInterrupt().
type GraphInterruptOption func(*graphInterruptOptions)

// WithGraphInterruptTimeout specifies the max waiting time before forcing an
// interrupt. After the timeout the executor will cancel in-flight work and
// interrupt as soon as it can.
func WithGraphInterruptTimeout(
	timeout time.Duration,
) GraphInterruptOption {
	_ = "STUB: not implemented"
	return *new(GraphInterruptOption)
}

// WithGraphInterrupt creates a context that can be interrupted externally.
//
// When the returned context is used to execute a graph, calling the returned
// interrupt function requests the run to pause and save an interrupt
// checkpoint.
//
// By default the executor waits for the current step's tasks to finish and
// interrupts before starting the next step. When WithGraphInterruptTimeout is
// provided, the executor will cancel in-flight tasks after the timeout.
func WithGraphInterrupt(
	parent context.Context,
) (
	ctx context.Context,
	interrupt func(opts ...GraphInterruptOption),
) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func graphInterruptFromContext(
	ctx context.Context,
) *graphInterruptState {
	_ = "STUB: not implemented"
	return nil
}

func (s *graphInterruptState) requested() bool { _ = "STUB: not implemented"; return false }

func (s *graphInterruptState) timeoutOrNil() *time.Duration { _ = "STUB: not implemented"; return nil }

func (s *graphInterruptState) doneCh() <-chan struct{} { _ = "STUB: not implemented"; return nil }
