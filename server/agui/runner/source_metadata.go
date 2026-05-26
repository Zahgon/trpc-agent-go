//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package runner

import (
	"context"

	aguievents "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
	agentevent "trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/server/agui/internal/source"
	"trpc.group/trpc-go/trpc-agent-go/session"
)

func (r *runner) attachToolResultInputSourceMetadata(
	ctx context.Context,
	key session.Key,
	event *agentevent.Event,
	toolCallID string,
) {
	_ = "STUB: not implemented"
	return
}

func (r *runner) lookupToolCallSourceMetadata(
	ctx context.Context,
	key session.Key,
	toolCallID string,
) (source.Metadata, bool, error) {
	_ = "STUB: not implemented"
	return *new(source.Metadata), false, nil
}

func toolCallSourceMetadata(
	payload []byte,
	toolCallID string,
) (source.Metadata, bool) {
	_ = "STUB: not implemented"
	return *new(source.Metadata), false
}

func matchesToolCallID(
	event aguievents.Event,
	toolCallID string,
) bool {
	_ = "STUB: not implemented"
	return false
}
