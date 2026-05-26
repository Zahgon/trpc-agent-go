//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

func (c *hedgeChat) printBanner() { _ = "STUB: not implemented"; return }

func printCommands() { _ = "STUB: not implemented"; return }

func (c *hedgeChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *hedgeChat) extractContent(choice model.Choice) string {
	_ = "STUB: not implemented"
	return ""
}
