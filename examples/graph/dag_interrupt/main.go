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
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
	"trpc.group/trpc-go/trpc-agent-go/graph/checkpoint/inmemory"
)

const (
	engineBSP = "bsp"
	engineDAG = "dag"

	nodeEntry = "entry"
	nodeAsk   = "ask"
	nodeAfter = "after"

	stateKeyAnswer = "answer"

	interruptKey    = "approval"
	interruptPrompt = "please approve: ok?"

	defaultEngine = engineDAG
	defaultResume = "ok"

	runTimeout = 2 * time.Second
)

type interruptMeta struct {
	NodeID       string `json:"nodeID,omitempty"`
	InterruptKey string `json:"interruptKey,omitempty"`
	LineageID    string `json:"lineageId,omitempty"`
	CheckpointID string `json:"checkpointId,omitempty"`
}

func main() {
	var (
		engine = flag.String(
			"engine",
			defaultEngine,
			"Execution engine: bsp|dag",
		)
		resumeValue = flag.String(
			"resume",
			defaultResume,
			"Resume value for the interrupt",
		)
	)
	flag.Parse()

	selectedEngine, err := parseEngine(*engine)
	if err != nil {
		log.Fatalf("parse -engine failed: %v", err)
	}

	g := buildGraph()
	saver := inmemory.NewSaver()

	exec, err := graph.NewExecutor(
		g,
		graph.WithExecutionEngine(selectedEngine),
		graph.WithCheckpointSaver(saver),
		graph.WithMaxConcurrency(2),
	)
	if err != nil {
		log.Fatalf("create executor failed: %v", err)
	}

	lineageID := fmt.Sprintf("dag-interrupt-%d", time.Now().UnixNano())

	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("Engine: %s\n", *engine)
	fmt.Printf("Lineage: %s\n", lineageID)
	fmt.Println("Run #1: expect interrupt")

	ctx1, cancel1 := context.WithTimeout(context.Background(), runTimeout)
	defer cancel1()

	evts1, err := exec.Execute(
		ctx1,
		graph.State{graph.CfgKeyLineageID: lineageID},
		&agent.Invocation{InvocationID: lineageID},
	)
	if err != nil {
		log.Fatalf("execute run #1 failed: %v", err)
	}

	meta, err := waitForInterrupt(evts1, runTimeout)
	if err != nil {
		log.Fatalf("run #1 wait failed: %v", err)
	}
	fmt.Printf(
		"Interrupted: node=%s key=%s checkpoint=%s\n",
		meta.NodeID,
		meta.InterruptKey,
		meta.CheckpointID,
	)

	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("Run #2: resume and expect completion")

	ctx2, cancel2 := context.WithTimeout(context.Background(), runTimeout)
	defer cancel2()

	resumeCmd := (&graph.ResumeCommand{}).WithResumeMap(map[string]any{
		interruptKey: *resumeValue,
	})
	resumeState := graph.State{
		graph.CfgKeyLineageID:    meta.LineageID,
		graph.CfgKeyCheckpointID: meta.CheckpointID,
		graph.StateKeyCommand:    resumeCmd,
	}
	evts2, err := exec.Execute(
		ctx2,
		resumeState,
		&agent.Invocation{InvocationID: lineageID + "-resume"},
	)
	if err != nil {
		log.Fatalf("execute run #2 failed: %v", err)
	}

	finalState, err := waitForCompletion(evts2, runTimeout)
	if err != nil {
		log.Fatalf("run #2 wait failed: %v", err)
	}
	fmt.Printf("Completed: answer=%v\n", finalState[stateKeyAnswer])
}

func parseEngine(raw string) (graph.ExecutionEngine, error) {
	_ = "STUB: not implemented"
	return *new(graph.ExecutionEngine), nil
}

func buildGraph() *graph.Graph { _ = "STUB: not implemented"; return nil }

func waitForInterrupt(
	evts <-chan *event.Event,
	timeout time.Duration,
) (interruptMeta, error) {
	_ = "STUB: not implemented"
	return *new(interruptMeta), nil
}

func waitForCompletion(
	evts <-chan *event.Event,
	timeout time.Duration,
) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func interruptMetaFromEvent(e *event.Event) (interruptMeta, bool) {
	_ = "STUB: not implemented"
	return *new(interruptMeta), false
}

func parseStateDelta(raw map[string][]byte) map[string]any { _ = "STUB: not implemented"; return nil }
