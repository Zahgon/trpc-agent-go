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

func printBanner(modelName string, location string, failures int) {
	_ = "STUB: not implemented"
	return
}

func printScenarioHeader(name string) { _ = "STUB: not implemented"; return }

func printToolAttempt(attempt int, location string) { _ = "STUB: not implemented"; return }

func printScenarioResult(
	name string,
	attempts int,
	answer string,
	runErr error,
	expectSuccess bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func collectScenarioResult(events <-chan *event.Event) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func printToolCalls(evt *event.Event) { _ = "STUB: not implemented"; return }

func compactJSON(raw []byte) string { _ = "STUB: not implemented"; return "" }
