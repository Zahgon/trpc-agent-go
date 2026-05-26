//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main provides a code check agent for A2A (Agent-to-Agent) communication.
package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	a2aserver "trpc.group/trpc-go/trpc-a2a-go/server"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/log"
	a2a "trpc.group/trpc-go/trpc-agent-go/server/a2a"
)

const (
	agentName = "CodeCheckAgent"
)

func main() {
	// Parse command-line flags.
	host := flag.String("host", "0.0.0.0:8088", "Host to listen on")
	modelName := flag.String("model", "deepseek-v4-flash", "Model to use")
	flag.Parse()

	// Build the code check agent
	codeCheckAgent := buildCodeCheckAgent(*modelName)
	agentCard := buildAgentCard()
	server, err := a2a.New(
		a2a.WithHost(*host),
		a2a.WithAgent(codeCheckAgent, true),
		a2a.WithAgentCard(agentCard),
	)
	if err != nil {
		log.Fatalf("Failed to create a2a server: %v", err)
	}

	// Set up a channel to listen for termination signals.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine.
	go func() {
		log.Infof("Starting server on %s...", *host)
		if err := server.Start(*host); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	sig := <-sigChan
	log.Infof("Received signal %v, shutting down...", sig)

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Stop(ctx); err != nil {
		log.Errorf("Failed to stop server gracefully: %v", err)
	}
}

func buildCodeCheckAgent(modelName string) agent.Agent {
	_ = "STUB: not implemented"
	return *new(agent.Agent)
}

func buildAgentCard() a2aserver.AgentCard {
	_ = "STUB: not implemented"
	return *new(a2aserver.AgentCard)
}

func stringPtr(s string) *string { _ = "STUB: not implemented"; return nil }

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func boolPtr(b bool) *bool { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
