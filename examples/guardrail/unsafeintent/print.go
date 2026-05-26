//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"trpc.group/trpc-go/trpc-agent-go/event"
)

func (a *demoApp) printBanner() { _ = "STUB: not implemented"; return }

func (a *demoApp) printHelp() { _ = "STUB: not implemented"; return }

func (a *demoApp) printEvents(eventCh <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}
