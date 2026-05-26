//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package outbound

import (
	"context"
	"sync"
)

type sentTextRecorderContextKey struct{}

// SentTextRecorder tracks successful text sends within one agent run.
type SentTextRecorder struct {
	mu      sync.Mutex
	sent    map[sentTextKey]struct{}
	targets map[sentTargetKey]struct{}
}

type sentTextKey struct {
	Channel string
	Target  string
	Text    string
}

type sentTargetKey struct {
	Channel string
	Target  string
}

// NewSentTextRecorder creates an empty per-run delivery recorder.
func NewSentTextRecorder() *SentTextRecorder { _ = "STUB: not implemented"; return nil }

// WithSentTextRecorder attaches a per-run recorder to a context.
func WithSentTextRecorder(
	ctx context.Context,
	recorder *SentTextRecorder,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Record stores a successful text delivery.
func (r *SentTextRecorder) Record(target DeliveryTarget, text string) {
	_ = "STUB: not implemented"
	return
}

// Contains reports whether the exact text target was already delivered.
func (r *SentTextRecorder) Contains(
	target DeliveryTarget,
	text string,
) bool {
	_ = "STUB: not implemented"
	return false
}

// ContainsTarget reports whether any text was delivered to the target.
func (r *SentTextRecorder) ContainsTarget(target DeliveryTarget) bool {
	_ = "STUB: not implemented"
	return false
}

func sentTextRecorderFromContext(
	ctx context.Context,
) (*SentTextRecorder, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func sentTextKeyFor(
	target DeliveryTarget,
	text string,
) (sentTextKey, bool) {
	_ = "STUB: not implemented"
	return *new(sentTextKey), false
}

func sentTargetKeyFor(target DeliveryTarget) (sentTargetKey, bool) {
	_ = "STUB: not implemented"
	return *new(sentTargetKey), false
}
