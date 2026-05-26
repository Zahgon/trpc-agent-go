//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main is an AG-UI server example that demonstrates how `skill_run`
// can persist outputs as artifacts and surface stable artifact refs to the
// frontend via `CustomEvent("tool.artifacts")`.
package main

import (
	"context"
	"flag"
	"net/http"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	artifactinmemory "trpc.group/trpc-go/trpc-agent-go/artifact/inmemory"
	localexec "trpc.group/trpc-go/trpc-agent-go/codeexecutor/local"
	"trpc.group/trpc-go/trpc-agent-go/log"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/server/agui"
	sessioninmemory "trpc.group/trpc-go/trpc-agent-go/session/inmemory"
	"trpc.group/trpc-go/trpc-agent-go/skill"
	"trpc.group/trpc-go/trpc-agent-go/tool"
	skilltool "trpc.group/trpc-go/trpc-agent-go/tool/skill"
)

const (
	appName      = "agui-skill-artifacts-demo"
	agentName    = "agui-skill-artifacts-agent"
	demoSkill    = "artifact_demo"
	demoToolCall = "call-skillrun-demo"
)

var (
	address   = flag.String("address", "127.0.0.1:8080", "Listen address")
	path      = flag.String("path", "/agui", "HTTP path")
	skillsDir = flag.String(
		"skills_dir",
		"./server/skill_artifacts/skills",
		"Skill repository root (run from examples/agui)",
	)
)

func main() {
	flag.Parse()

	if _, err := os.Stat(*skillsDir); err != nil {
		log.Fatalf("skills_dir not found: %s (%v)", *skillsDir, err)
	}
	repo, err := skill.NewFSRepository(*skillsDir)
	if err != nil {
		log.Fatalf("create skill repository: %v", err)
	}

	exec := localexec.New()
	runTool := skilltool.NewRunTool(
		repo,
		exec,
		skilltool.WithForceSaveArtifacts(true),
	)

	modelInstance := &scriptedModel{}
	generationConfig := model.GenerationConfig{
		MaxTokens:        intPtr(256),
		Temperature:      floatPtr(0),
		Stream:           false,
		Stop:             nil,
		TopP:             nil,
		PresencePenalty:  nil,
		FrequencyPenalty: nil,
	}

	ag := llmagent.New(
		agentName,
		llmagent.WithDescription("AG-UI demo: skill_run artifacts -> tool.artifacts custom event."),
		llmagent.WithModel(modelInstance),
		llmagent.WithGenerationConfig(generationConfig),
		llmagent.WithTools([]tool.Tool{runTool}),
		llmagent.WithInstruction("You are a demo agent."),
	)

	sessionService := sessioninmemory.NewSessionService()
	artifactService := artifactinmemory.NewService()

	r := runner.NewRunner(
		appName,
		ag,
		runner.WithSessionService(sessionService),
		runner.WithArtifactService(artifactService),
	)
	defer r.Close()

	server, err := agui.New(
		r,
		agui.WithAppName(appName),
		agui.WithSessionService(sessionService),
		agui.WithPath(*path),
		agui.WithMessagesSnapshotEnabled(true),
	)
	if err != nil {
		log.Fatalf("create AG-UI server: %v", err)
	}

	log.Infof("AG-UI: serving agent %q on http://%s%s", ag.Info().Name, *address, *path)
	log.Infof("Try the raw client: cd examples/agui && go run ./client/raw --endpoint http://%s%s", *address, *path)
	log.Infof("You should see a [CUSTOM_EVENT] named 'tool.artifacts' with artifact:// refs.")

	if err := http.ListenAndServe(*address, server.Handler()); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}

// scriptedModel is a deterministic model implementation that always calls
// `skill_run` once, then returns a final assistant message.
//
// This keeps the example runnable without external model credentials.
type scriptedModel struct{}

func (m *scriptedModel) Info() model.Info { _ = "STUB: not implemented"; return *new(model.Info) }

func (m *scriptedModel) GenerateContent(
	ctx context.Context,
	req *model.Request,
) (<-chan *model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// After tool execution, llmagent will call the model again with a tool
// message in the request. Detect that and finish.

// ctx cancellation is handled by llmagent/runner.

func hasToolResult(messages []model.Message) bool { _ = "STUB: not implemented"; return false }

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
