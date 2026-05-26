//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package tool defines AG-UI activity metadata for streamed tool results.
package tool

import (
	aguievents "github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
)

const (
	// StreamingToolResultActivityType is the AG-UI activity type for streamed
	// partial tool-result chunks.
	StreamingToolResultActivityType = "tool.result.stream"
)

// IsStreamingToolResultActivityEvent reports whether the event is a synthetic
// activity event derived from partial tool-result chunks.
func IsStreamingToolResultActivityEvent(evt aguievents.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// StreamingToolResultActivityMessageID returns the synthetic activity message ID
// for the tool call.
func StreamingToolResultActivityMessageID(toolCallID string) string {
	_ = "STUB: not implemented"
	return ""
}
