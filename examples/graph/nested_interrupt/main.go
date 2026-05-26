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
	"os"

	_ "github.com/mattn/go-sqlite3"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	graphagent "trpc.group/trpc-go/trpc-agent-go/agent/graphagent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

const (
	modeRun    = "run"
	modeResume = "resume"

	sqliteDriverName = "sqlite3"

	defaultDBPath    = "nested-interrupt.db"
	defaultLineageID = "demo-nested-interrupt"

	minDepth              = 2
	parentAgentName       = "parent"
	nestedAgentNamePrefix = "agent_"

	nodeAsk              = "ask"
	interruptKeyApproval = "approval"
	stateKeyAnswer       = "answer"

	startMessage  = "start"
	resumeMessage = "resume"

	interruptPrompt = "Please type an approval string."
)

type runSummary struct {
	interruptValue any
	answer         string
}

func main() {
	var (
		mode         string
		dbPath       string
		lineageID    string
		checkpointID string
		resumeValue  string
		depth        int
	)

	flag.StringVar(&mode, "mode", modeRun, "run or resume")
	flag.StringVar(&lineageID, "lineage-id", defaultLineageID, "lineage")
	flag.StringVar(&dbPath, "db", defaultDBPath, "sqlite path")
	flag.StringVar(&checkpointID, "checkpoint-id", "", "resume checkpoint")
	flag.StringVar(&resumeValue, "resume-value", "", "resume value")
	flag.IntVar(&depth, "depth", minDepth, "agent nesting depth")
	flag.Parse()

	if err := validateFlags(mode, lineageID, checkpointID, depth); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	ctx := context.Background()
	saver, closeDB, err := openSQLiteSaver(dbPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer closeDB()

	parent, cm, err := buildAgents(saver, depth)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch mode {
	case modeRun:
		err = runMode(ctx, parent, cm, lineageID, depth)
	case modeResume:
		err = resumeMode(ctx, parent, cm, lineageID, checkpointID, resumeValue)
	default:
		err = fmt.Errorf("unknown mode: %s", mode)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func validateFlags(
	mode string,
	lineageID string,
	checkpointID string,
	depth int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func openSQLiteSaver(
	dbPath string,
) (graph.CheckpointSaver, func(), error) {
	_ = "STUB: not implemented"
	return *new(graph.CheckpointSaver), nil, nil
}

func buildAgents(
	saver graph.CheckpointSaver,
	depth int,
) (*graphagent.GraphAgent, *graph.CheckpointManager, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func askNode() graph.NodeFunc { _ = "STUB: not implemented"; return *new(graph.NodeFunc) }

func answerOutputMapper() graph.SubgraphOutputMapper {
	_ = "STUB: not implemented"
	return *new(graph.SubgraphOutputMapper)
}

func runMode(
	ctx context.Context,
	parent *graphagent.GraphAgent,
	cm *graph.CheckpointManager,
	lineageID string,
	depth int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func resumeMode(
	ctx context.Context,
	parent *graphagent.GraphAgent,
	cm *graph.CheckpointManager,
	lineageID string,
	checkpointID string,
	resumeValue string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func runOnce(
	ctx context.Context,
	a agent.Agent,
	inv *agent.Invocation,
) (runSummary, error) {
	_ = "STUB: not implemented"
	return *new(runSummary), nil
}

func readEvents(ch <-chan *event.Event) (runSummary, error) {
	_ = "STUB: not implemented"
	return *new(runSummary), nil
}

func interruptFromEvent(ev *event.Event) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}

func answerFromStateDelta(delta map[string][]byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
