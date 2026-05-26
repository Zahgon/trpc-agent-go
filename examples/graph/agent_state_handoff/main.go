//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates how an Agent node passes data to later nodes by
// writing to graph state.
//
// Key idea:
//   - The Agent node runs a child agent (a sub-agent).
//   - The child agent writes values into its own (child) graph state.
//   - WithSubgraphOutputMapper copies selected values from the child
//     final state back into the parent graph state.
//   - Subsequent parent nodes can read those values from state.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

const (
	emptyString = ""

	parentAgentName = "parent"
	childAgentName  = "child"

	childNodeCompute = "compute"
	parentNodeUse    = "use_child_value"

	keyChildValue     = "child_value"
	keyValueFromChild = "value_from_child"

	childValuePrefix  = "computed: "
	parentValuePrefix = "parent received: "

	flagInputName  = "input"
	flagInputUsage = "Text sent into the parent graph as user input"
	defaultInput   = "hello"

	childAgentDesc  = "Child agent that computes a value into state"
	parentAgentDesc = "Parent agent that calls the child and uses its state"

	exitCodeError = 1

	errNoCompletion          = "no graph completion event received"
	fmtMissingStateKey       = "missing state key: %s"
	fmtUnmarshalStateKey     = "unmarshal state key %s: %w"
	fmtInputLine             = "Input: %s\n"
	fmtValueFromChildLine    = "Value from child (via state): %s\n"
	fmtFinalResponseLine     = "Final response: %s\n"
	fmtErrorLine             = "error: %v\n"
	fmtNoFinalResponseChoice = "(no final response choice)"
)

var input = flag.String(flagInputName, defaultInput, flagInputUsage)

func main() {
	flag.Parse()

	childAgent, err := buildChildAgent()
	if err != nil {
		fmt.Fprintf(os.Stderr, fmtErrorLine, err)
		os.Exit(exitCodeError)
	}

	parentAgent, err := buildParentAgent(childAgent)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmtErrorLine, err)
		os.Exit(exitCodeError)
	}

	completionEvent, err := runOnce(context.Background(), parentAgent, *input)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmtErrorLine, err)
		os.Exit(exitCodeError)
	}

	valueFromChild, err := decodeJSONString(
		completionEvent.StateDelta,
		keyValueFromChild,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmtErrorLine, err)
		os.Exit(exitCodeError)
	}

	fmt.Printf(fmtInputLine, *input)
	fmt.Printf(fmtValueFromChildLine, valueFromChild)
	fmt.Printf(fmtFinalResponseLine, finalResponseText(completionEvent))
}

func buildChildAgent() (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func childComputeNode(ctx context.Context, state graph.State) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func buildParentAgent(childAgent agent.Agent) (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func subgraphOutputMapper(
	parent graph.State,
	result graph.SubgraphResult,
) graph.State {
	_ = "STUB: not implemented"
	return *new(graph.State)
}

func parentUseChildValueNode(
	ctx context.Context,
	state graph.State,
) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func runOnce(
	ctx context.Context,
	a agent.Agent,
	userInput string,
) (*event.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func finalResponseText(ev *event.Event) string { _ = "STUB: not implemented"; return "" }

func decodeJSONString(
	stateDelta map[string][]byte,
	key string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
