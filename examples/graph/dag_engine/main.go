//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

const (
	engineBSP  = "bsp"
	engineDAG  = "dag"
	engineBoth = "both"
)

const (
	nodeSplit   = "split"
	nodeSlowA   = "slow_a"
	nodeFastB   = "fast_b"
	nodeMidC    = "mid_c"
	nodeFastNxt = "fast_b_next"
)

const (
	slowDuration = 800 * time.Millisecond
	fastDuration = 200 * time.Millisecond
	midDuration  = 400 * time.Millisecond
	nextDuration = 120 * time.Millisecond
)

func main() {
	var engine string
	flag.StringVar(
		&engine,
		"engine",
		engineBSP,
		"Execution engine: bsp|dag|both",
	)
	flag.Parse()

	engine = strings.ToLower(strings.TrimSpace(engine))
	if engine != engineBSP && engine != engineDAG && engine != engineBoth {
		panic(fmt.Errorf("unknown engine %q", engine))
	}

	switch engine {
	case engineBoth:
		run(engineBSP)
		fmt.Println(strings.Repeat("-", 60))
		run(engineDAG)
	default:
		run(engine)
	}
}

func run(engine string) { _ = "STUB: not implemented"; return }

func buildGraph(l *logger) *graph.Graph { _ = "STUB: not implemented"; return nil }

func drainEvents(evts <-chan *event.Event) error { _ = "STUB: not implemented"; return nil }

type logger struct {
	start time.Time
	mu    sync.Mutex
}

func newLogger(start time.Time) *logger { _ = "STUB: not implemented"; return nil }

func (l *logger) Printf(format string, args ...any) { _ = "STUB: not implemented"; return }
