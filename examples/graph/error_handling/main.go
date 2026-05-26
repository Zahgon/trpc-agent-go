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
	"fmt"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/graph"
)

const (
	appName = "graph-error-handling"

	stateKeyNote = "note"

	recoverableScenario = "recoverable_local_node"
	fatalScenario       = "fatal_local_node"
	subgraphScenario    = "fatal_child_subgraph"

	recoverableAgentName = "recoverable-graph"
	fatalAgentName       = "fatal-graph"
	parentAgentName      = "parent-error-graph"
	childAgentName       = "child-error-graph"

	softErrorCode       = "LOOKUP_SOFT_TIMEOUT"
	fatalErrorCode      = "WRITE_FATAL"
	childFatalErrorCode = "CHILD_AGENT_FATAL"
)

type codedError struct {
	code        string
	message     string
	recoverable bool
}

func (e codedError) Error() string { _ = "STUB: not implemented"; return "" }

func (e codedError) Code() string { _ = "STUB: not implemented"; return "" }

func (e codedError) Recoverable() bool { _ = "STUB: not implemented"; return false }

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

func buildRecoverableAgent() (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func buildFatalAgent() (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func buildParentSubgraphAgent() (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func buildChildFatalAgent() (agent.Agent, error) {
	_ = "STUB: not implemented"
	return *new(agent.Agent), nil
}

func runScenario(name string, ag agent.Agent) error { _ = "STUB: not implemented"; return nil }

func decodeStringState(
	stateDelta map[string][]byte,
	key string,
) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func readExecutionErrors(
	state graph.State,
) []graph.ExecutionError {
	_ = "STUB: not implemented"
	return nil
}

func ptrValue(value *string) string { _ = "STUB: not implemented"; return "" }
