//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package main runs an OpenClaw runtime and consumes it via A2A.
package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/openclaw/app"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	defaultQuestion = "What's the weather in Shanghai today?"
	defaultFollowUp = "What about tomorrow?"

	defaultSkillsRoot  = "./skills"
	defaultA2ABase     = "/a2a"
	defaultA2AName     = "openclaw-sandbox"
	defaultA2ADesc     = "Sandbox agent for bundled skills."
	defaultInstruction = "For live weather or forecast questions, " +
		"call skill_load for the weather skill before skill_run. " +
		"Use the loaded weather skill instead of guessing current " +
		"conditions."
	weatherSkillName  = "weather"
	skillsLoadSession = "session"
	openAIMode        = "openai"

	exampleRunnerName = "openclaw-a2a-example"
	exampleAppName    = "openclaw-a2a-example"
	exampleUserID     = "example-user"
	exampleSessionID  = "example-session"

	startupTimeout = 15 * time.Second
	requestTimeout = 150 * time.Second
	shutdownWait   = 5 * time.Second
	pollInterval   = 100 * time.Millisecond
)

var (
	addrFlag = flag.String(
		"addr",
		"",
		"HTTP listen address for OpenClaw (default random loopback port)",
	)
	skillsRootFlag = flag.String(
		"skills-root",
		defaultSkillsRoot,
		"Path to the OpenClaw skills root",
	)
	stateDirFlag = flag.String(
		"state-dir",
		"",
		"State dir for the embedded OpenClaw runtime",
	)
	modelFlag = flag.String(
		"model",
		"",
		"Optional OpenAI model override",
	)
	baseURLFlag = flag.String(
		"openai-base-url",
		"",
		"Optional OpenAI base URL override",
	)
	questionFlag = flag.String(
		"question",
		defaultQuestion,
		"First user question sent through A2A",
	)
	followUpFlag = flag.String(
		"follow-up",
		defaultFollowUp,
		"Optional follow-up question using the same session",
	)
	streamingFlag = flag.Bool(
		"streaming",
		true,
		"Enable streaming on the OpenClaw A2A surface",
	)
	advertiseToolsFlag = flag.Bool(
		"advertise-tools",
		false,
		"Advertise individual tools in the agent card",
	)
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

func newRuntime(
	configPath string,
	a2aURL string,
	addr string,
	stateDir string,
) (*app.Runtime, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newHTTPServer(rt *app.Runtime) *http.Server { _ = "STUB: not implemented"; return nil }

func listenLoopback(rawAddr string) (net.Listener, string, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), "", nil
}

func resolveStateDir(raw string) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func writeConfigStub() (string, func(), error) { _ = "STUB: not implemented"; return "", nil, nil }

func waitForReady(
	ctx context.Context,
	addr string,
	cardPath string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func prompts() []string { _ = "STUB: not implemented"; return nil }

func ask(procRunner runner.Runner, prompt string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func collectAnswer(eventCh <-chan *event.Event) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func shutdownServer(httpSrv *http.Server) { _ = "STUB: not implemented"; return }

func receiveServeErr(serveErr <-chan error) error { _ = "STUB: not implemented"; return nil }
