//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main demonstrates how to use the MCP broker tools with LLMAgent.
package main

import (
	"context"
	"flag"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
	"trpc.group/trpc-go/trpc-agent-go/tool/mcpbroker"
)

var (
	modelName = flag.String("model", "gpt-4o-mini", "Name of the model to use")
	variant   = flag.String("variant", "", "Optional provider variant. If empty, infer from OPENAI_BASE_URL")
	streaming = flag.Bool("streaming", true, "Enable streaming mode for responses")
	prompt    = flag.String("prompt", "", "Optional single-turn prompt. If empty, start interactive chat")
)

const (
	appName   = "mcp-broker-demo"
	agentName = "mcp-broker-assistant"
)

func main() {
	flag.Parse()

	chat := &mcpBrokerChat{
		modelName: *modelName,
		variant:   *variant,
		streaming: *streaming,
		prompt:    strings.TrimSpace(*prompt),
	}

	if err := chat.run(); err != nil {
		log.Fatalf("chat failed: %v", err)
	}
}

type mcpBrokerChat struct {
	modelName          string
	variant            string
	streaming          bool
	prompt             string
	runner             runner.Runner
	userID             string
	sessionID          string
	visibleTools       []string
	generatedSkillsDir string
	preferredServer    string
	availableServerIDs []string
	availableSkillIDs  []string
	remoteMCPURL       string
	remoteHTTPDemo     *remoteHTTPDemo
}

func (c *mcpBrokerChat) run() error { _ = "STUB: not implemented"; return nil }

func (c *mcpBrokerChat) setup(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *mcpBrokerChat) buildBrokerOptions(serverPath string) (
	[]mcpbroker.Option,
	[]string,
	string,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil, "", nil
}

func (c *mcpBrokerChat) cleanupGeneratedSkillsDir() { _ = "STUB: not implemented"; return }

func (c *mcpBrokerChat) closeRemoteHTTPDemo() { _ = "STUB: not implemented"; return }

func (c *mcpBrokerChat) printBanner() { _ = "STUB: not implemented"; return }

func (c *mcpBrokerChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *mcpBrokerChat) printTips() { _ = "STUB: not implemented"; return }

func (c *mcpBrokerChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *mcpBrokerChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *mcpBrokerChat) handleToolCalls(
	evt *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *mcpBrokerChat) handleToolResponses(evt *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *mcpBrokerChat) handleContent(
	evt *event.Event,
	toolCallsDetected *bool,
	assistantStarted *bool,
) {
	_ = "STUB: not implemented"
	return
}

func exampleDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

func toolNames(tools []tool.Tool) []string { _ = "STUB: not implemented"; return nil }

func skillNames(repo skill.Repository) []string { _ = "STUB: not implemented"; return nil }

func extractContent(choice model.Choice, streaming bool) string {
	_ = "STUB: not implemented"
	return ""
}

func compactJSON(raw []byte) string { _ = "STUB: not implemented"; return "" }

func preview(text string, max int) string { _ = "STUB: not implemented"; return "" }

func intPtr(v int) *int { _ = "STUB: not implemented"; return nil }
