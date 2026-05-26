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

func printBanner(modelName string, variant string, streaming bool, prompt string) {
	_ = "STUB: not implemented"
	return
}

func printEvents(eventCh <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

func printToolCalls(evt *event.Event) { _ = "STUB: not implemented"; return }

func printToolResults(evt *event.Event) { _ = "STUB: not implemented"; return }

func printCalculatorExecution(callID string, args calculatorArgs) {
	_ = "STUB: not implemented"
	return
}

func compactJSON(raw []byte) string { _ = "STUB: not implemented"; return "" }

func preview(text string, max int) string { _ = "STUB: not implemented"; return "" }
