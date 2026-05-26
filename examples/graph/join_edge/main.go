//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates the join edge (wait-all fan-in) behavior using
// graph.StateGraph.AddJoinEdge.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

const (
	nodeStart = "start"
	nodeA     = "a"
	nodeB     = "b"
	nodeJoin  = "join"

	stateKeyOrder = "order"
)

var (
	sleepA = flag.Duration(
		"sleep-a",
		120*time.Millisecond,
		"Simulated work duration for node a",
	)
	sleepB = flag.Duration(
		"sleep-b",
		40*time.Millisecond,
		"Simulated work duration for node b",
	)
)

func main() {
	flag.Parse()

	g, err := buildGraph(*sleepA, *sleepB)
	if err != nil {
		log.Fatalf("build graph failed: %v", err)
	}
	exec, err := graph.NewExecutor(g)
	if err != nil {
		log.Fatalf("create executor failed: %v", err)
	}

	inv := &agent.Invocation{InvocationID: "join-edge-demo"}
	ch, err := exec.Execute(context.Background(), graph.State{}, inv)
	if err != nil {
		log.Fatalf("execute failed: %v", err)
	}

	order, err := waitForFinalOrder(ch)
	if err != nil {
		log.Fatalf("read final order failed: %v", err)
	}
	fmt.Printf("Final execution order: %v\n", order)
}

func buildGraph(sleepA, sleepB time.Duration) (*graph.Graph, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func waitForFinalOrder(
	ch <-chan *event.Event,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
