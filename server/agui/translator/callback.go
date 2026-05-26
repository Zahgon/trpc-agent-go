//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package translator

import (
	"context"

	aguievents "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
	"trpc.group/trpc-go/trpc-agent-go/event"
)

// BeforeTranslateCallback is invoked before translating an internal agent event.
// Returns (customEvent, err).
//   - customEvent: if not nil, the translator will consume this event instead of the original one.
//   - err: if not nil, translation stops and the run returns an error to the client.
type BeforeTranslateCallback func(ctx context.Context, event *event.Event) (*event.Event, error)

// AfterTranslateCallback is invoked after translating into an AG-UI event, but before sending it to the client.
// Returns (customEvent, err).
//   - customEvent: if not nil, the modified event will be emitted to the client.
//   - err: if not nil, emission stops and the run returns an error to the client.
type AfterTranslateCallback func(ctx context.Context, event aguievents.Event) (aguievents.Event, error)

// Callbacks holds translation lifecycle hooks for AG-UI runner.
type Callbacks struct {
	// BeforeTranslate runs before translating internal events to AG-UI events.
	BeforeTranslate []BeforeTranslateCallback
	// AfterTranslate runs after translation, before the AG-UI event is emitted to clients.
	AfterTranslate []AfterTranslateCallback
}

// NewCallbacks creates a new Callbacks instance for translation hooks.
func NewCallbacks() *Callbacks { _ = "STUB: not implemented"; return nil }

// RegisterBeforeTranslate registers a callback executed before translation.
func (c *Callbacks) RegisterBeforeTranslate(cb BeforeTranslateCallback) *Callbacks {
	_ = "STUB: not implemented"
	return nil
}

// RegisterAfterTranslate registers a callback executed after translation.
func (c *Callbacks) RegisterAfterTranslate(cb AfterTranslateCallback) *Callbacks {
	_ = "STUB: not implemented"
	return nil
}

// RunBeforeTranslate runs all before-translate callbacks in order.
// Returns (customEvent, err). If any callback returns a non-nil custom event, translation uses it.
func (c *Callbacks) RunBeforeTranslate(ctx context.Context, event *event.Event) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunAfterTranslate runs all after-translate callbacks in order.
// Returns (customEvent, err). If any callback returns a non-nil custom event, it will be emitted.
func (c *Callbacks) RunAfterTranslate(ctx context.Context, event aguievents.Event) (aguievents.Event, error) {
	_ = "STUB: not implemented"
	return *new(aguievents.Event), nil
}
