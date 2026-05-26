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
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/knowledge"
)

// progressPrinter renders a multi-line progress display for knowledge loading.
// Each source gets its own progress bar; the display is refreshed in-place
// using ANSI escape codes.
type progressPrinter struct {
	mu      sync.Mutex
	sources []string
	state   map[string]*srcState
	printed int
}

type srcState struct {
	processed int
	total     int
	err       error
}

func newProgressPrinter() *progressPrinter { _ = "STUB: not implemented"; return nil }

func (p *progressPrinter) onProgress(_ context.Context, evt knowledge.LoadProgressEvent) {
	_ = "STUB: not implemented"
	return
}

func (p *progressPrinter) getOrCreate(name string) *srcState { _ = "STUB: not implemented"; return nil }

func (p *progressPrinter) redraw(evt knowledge.LoadProgressEvent) {
	_ = "STUB: not implemented"
	return
}

func (p *progressPrinter) printDoneSummary(evt knowledge.LoadProgressEvent) {
	_ = "STUB: not implemented"
	return
}

func truncate(s string, max int) string { _ = "STUB: not implemented"; return "" }
