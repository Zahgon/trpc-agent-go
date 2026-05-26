//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// processResponse handles both streaming and non-streaming responses.
func (c *appendEventChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// handleEvent processes a single event from the event channel.
func (c *appendEventChat) handleEvent(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) error {
	_ = "STUB: not implemented"
	// Handle errors.
	return nil
}

// Handle tool calls.

// Handle tool responses.

// Handle content.

// handleToolCalls detects and displays tool calls.
func (c *appendEventChat) handleToolCalls(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

// handleToolResponses detects and displays tool responses.
func (c *appendEventChat) handleToolResponses(event *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

// handleContent processes and displays content.
func (c *appendEventChat) handleContent(
	event *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) {
	_ = "STUB: not implemented"
	return
}

// extractContent extracts content based on streaming mode.
func (c *appendEventChat) extractContent(choice model.Choice) string {
	_ = "STUB: not implemented"
	return ""
}

// displayContent prints content to console.
func (c *appendEventChat) displayContent(
	content string,
	toolCallsDetected *bool,
	assistantStarted *bool,
	fullContent *string,
) {
	_ = "STUB: not implemented"
	return
}

func (c *appendEventChat) startNewSession() { _ = "STUB: not implemented"; return }

func (c *appendEventChat) rememberSession(id string) { _ = "STUB: not implemented"; return }

func (c *appendEventChat) listSessions() { _ = "STUB: not implemented"; return }

func (c *appendEventChat) switchSession(target string) { _ = "STUB: not implemented"; return }

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
