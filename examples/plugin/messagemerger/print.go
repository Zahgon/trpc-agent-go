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

func collectFinalResponse(evCh <-chan *event.Event) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func printMessages(title string, messages []model.Message) { _ = "STUB: not implemented"; return }
