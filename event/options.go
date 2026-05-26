//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package event provides the event system for agent communication.
package event

import (
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// Option is a function that can be used to configure the Event.
type Option func(*Event)

// WithBranch sets the branch for the event.
func WithBranch(branch string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithResponse sets the response for the event.
func WithResponse(response *model.Response) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithObject sets the object for the event.
func WithObject(o string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStateDelta sets state delta for the event.
func WithStateDelta(stateDelta map[string][]byte) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithStructuredOutputPayload sets a typed structured output payload on the event.
// This data is not serialized and is intended for immediate consumption.
func WithStructuredOutputPayload(payload any) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSkipSummarization sets the SkipSummarization action on the event.
func WithSkipSummarization() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTag sets the tag for the event.
func WithTag(tag string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithExtension stores one serialized extension on the event.
func WithExtension(key string, value any) Option { _ = "STUB: not implemented"; return *new(Option) }

// SetExtension stores one serialized extension on the event.
func SetExtension(e *Event, key string, value any) error { _ = "STUB: not implemented"; return nil }

// GetExtension decodes one typed extension from the event.
func GetExtension[T any](e *Event, key string) (T, bool, error) {
	_ = "STUB: not implemented"
	return *new(T), false, nil
}
