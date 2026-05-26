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

// RetryCondition determines whether an error is retryable.
type RetryCondition interface {
	Match(err error) bool
}

// RetryConditionFunc is an adapter to allow the use of
// ordinary functions as RetryCondition.
type RetryConditionFunc func(error) bool

// Match calls f(err).
func (f RetryConditionFunc) Match(err error) bool {
	_ = "STUB: not implemented"

	// RetryPolicy defines per-node or default retry configuration.
	// Attempts are counted inclusive of the first try. For example,
	// MaxAttempts=3 means 1 initial try + up to 2 retries.
	return false
}

type RetryPolicy struct {
	MaxAttempts     int
	InitialInterval time.Duration
	BackoffFactor   float64
	MaxInterval     time.Duration
	Jitter          bool
	RetryOn         []RetryCondition

	// Optional total time budget across retries; 0 to disable.
	MaxElapsedTime time.Duration
	// Optional per-attempt timeout override; 0 to use executor's node timeout.
	PerAttemptTimeout time.Duration
}

// NextDelay returns the backoff delay before the given attempt number.
// attempt starts at 1 for the first try; delay applies before the next retry,
// so callers typically pass the current attempt count.
func (p RetryPolicy) NextDelay(attempt int) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Compute exponential backoff based on attempt index (attempt-1 increments)

// Default to no exponential growth if misconfigured

// Clamp to MaxInterval if set

// Full jitter in [d, 2d) style is common; here use [0, d) additive jitter.
// Use crypto/rand to avoid gosec G404 complaint.

// ShouldRetry reports whether the given error matches any of the policy's conditions.
func (p RetryPolicy) ShouldRetry(err error) bool { _ = "STUB: not implemented"; return false }

// RetryOnErrors creates a condition that matches when errors.Is(err, any target).
func RetryOnErrors(targets ...error) RetryCondition {
	_ = "STUB: not implemented"
	return *new(RetryCondition)
}

// RetryOnPredicate creates a condition that defers matching to the provided function.
func RetryOnPredicate(match func(error) bool) RetryCondition {
	_ = "STUB: not implemented"
	return *new(RetryCondition)
}

// DefaultTransientCondition matches common transient errors worthy of retry:
// - context.DeadlineExceeded
// - net.Error with Timeout() or Temporary()
func DefaultTransientCondition() RetryCondition {
	_ = "STUB: not implemented"
	return *new(RetryCondition)
}

// Temporary() is deprecated but widely implemented
// so still consider it when available.

// WithSimpleRetry is a convenience constructor for a basic retry policy.
// Example defaults: attempts=3, initial=500ms, factor=2.0, max=8s, jitter=true,
// retrying on DefaultTransientCondition.
func WithSimpleRetry(attempts int) RetryPolicy { _ = "STUB: not implemented"; return *new(RetryPolicy) }
