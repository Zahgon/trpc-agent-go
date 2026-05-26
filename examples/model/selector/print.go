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

func printBanner(cfg appConfig) { _ = "STUB: not implemented"; return }

func printEvents(eventChan <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

func printToolCalls(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

func printToolResponses(evt *event.Event) bool { _ = "STUB: not implemented"; return false }

func extractContent(choice model.Choice) string { _ = "STUB: not implemented"; return "" }
