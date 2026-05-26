//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights
// reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates an interactive chat using Runner + LLMAgent
// with Agent Skills enabled. It streams content and shows tool calls and
// tool responses during the conversation.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
	"trpc.group/trpc-go/trpc-agent-go/artifact"
	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/model/openai"
	"trpc.group/trpc-go/trpc-agent-go/runner"
	"trpc.group/trpc-go/trpc-agent-go/skill"
)

var (
	flagModel          = flag.String("model", "deepseek-v4-flash", "model name")
	flagStream         = flag.Bool("stream", true, "stream responses")
	flagSkills         = flag.String("skills-root", "", "skills root dir")
	flagSkillsGuidance = flag.Bool(
		"skills-guidance",
		true,
		"include built-in skills tooling/workspace guidance",
	)
	flagSendFileInputs = flag.Bool(
		"send-file-inputs",
		false,
		"send file inputs to the model provider (may be unsupported)",
	)
	flagExec = flag.String(
		"executor", "local",
		"workspace executor: local|container|e2b",
	)
	flagTrustedLocal = flag.Bool(
		"trusted-local",
		false,
		"local executor only: reuse a fixed workspace root (unsafe)",
	)
	flagTrustedRoot = flag.String(
		"trusted-root",
		"",
		"trusted-local workspace root (default: ./skill_workspace)",
	)
	flagInputsHost = flag.String(
		"inputs-host", "",
		"host dir to bind as /opt/trpc-agent/inputs "+
			"(container exec)",
	)
)

const defaultSkillsDir = "skills"
const appName = "skill-run-chat"

const (
	artifactRefPrefix    = "artifact://"
	artifactUploadPrefix = "uploads/"
)

// instructionText guides the assistant behavior in a general way so it
// works with different skills repositories without assuming specifics.
const instructionText = `
Be a concise, helpful assistant that can use Agent Skills.

When a task may need tools, first ask to list skills or suggest one.
Load a skill with skill_load, then run the commands from its docs by
calling workspace_exec inside the loaded skill working copy
(cwd: skills/<name>). Prefer safe defaults; ask clarifying questions
if anything is ambiguous. Summarize results, note saved files, and
propose next steps briefly.

Inside the workspace, treat inputs/ and work/inputs/ as read-only
views of host files unless skill docs say they are writable. Do not
create, move, or modify files under inputs/ or work/inputs/.
User-uploaded file inputs from the conversation are staged under
work/inputs/ before scripts run.

When chaining multiple skills, read previous results directly from
out/ (or $OUTPUT_DIR) and write new files back to out/. If you need
a stable reference to an output file (for example to hand it to
another tool or to surface it back to the user), call
workspace_save_artifact on the workspace path.`

func main() {
	flag.Parse()

	// Resolve skills root (flag, env, default ./skills).
	root := *flagSkills
	if root == "" {
		if s := os.Getenv(skill.EnvSkillsRoot); s != "" {
			root = s
		} else {
			cwd, _ := os.Getwd()
			root = filepath.Join(cwd, defaultSkillsDir)
		}
	}

	// Setup runner and agent.
	chat := &skillChat{
		modelName:  *flagModel,
		stream:     *flagStream,
		skillsRoot: root,
	}
	if err := chat.run(); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}
}

type skillChat struct {
	modelName  string
	stream     bool
	skillsRoot string
	runner     runner.Runner
	userID     string
	sessionID  string
	executor   string
	artSvc     artifact.Service
	model      *openai.Model
	uploaded   []string
}

func (c *skillChat) run() error { _ = "STUB: not implemented"; return nil }

func (c *skillChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Model.
	return nil
}

// Skills repository.

// Choose workspace executor.

// Bind the skills root read-only into the container to
// enable fast in-container copy when staging directories.

// When an inputs-host directory is provided and bound
// into the container, automatically expose it under
// work/inputs (and thus inputs/) inside each workspace
// so skills can read host files via inputs/ paths.

// Optional: bind a host directory for zero-copy inputs.

// Agent with skills enabled. The default knowledge_only skill tool
// profile registers skill_load / skill_list_docs / skill_select_docs;
// the explicit code executor wires workspace_exec (and
// workspace_save_artifact when an artifact service is present) on
// top of it, which is the recommended execution surface.

// Runner + artifact service (in-memory for demo).

// Intro.

func buildAgentOptions(
	mdl model.Model,
	repo skill.Repository,
	exec codeexecutor.CodeExecutor,
	gen model.GenerationConfig,
) []llmagent.Option {
	_ = "STUB: not implemented"
	return nil
}

// Disable fenced-code auto-execution so the model has to go
// through workspace_exec instead of having ```sh blocks
// scraped out of its prose.

func (c *skillChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *skillChat) processMessage(
	ctx context.Context, userMessage string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *skillChat) processModelMessage(
	ctx context.Context, msg model.Message,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *skillChat) processResponse(
	ch <-chan *event.Event,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *skillChat) handleUpload(
	ctx context.Context,
	text string,
	useFileIDs bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *skillChat) handleUploadArtifact(
	ctx context.Context,
	text string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func guessMimeType(name string, data []byte) string { _ = "STUB: not implemented"; return "" }

func (c *skillChat) cleanupUploads(ctx context.Context) { _ = "STUB: not implemented"; return }

// handlePull downloads an artifact from the service and writes it to disk.
// Usage: /pull <name> [version]
func (c *skillChat) handlePull(text string) error { _ = "STUB: not implemented"; return nil }

// handleListArtifacts lists artifact keys for the current session.
func (c *skillChat) handleListArtifacts() error { _ = "STUB: not implemented"; return nil }

func parseInt(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *skillChat) handleEvent(
	ev *event.Event, toolCalls *bool, started *bool, full *string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *skillChat) handleToolCalls(
	ev *event.Event, toolCalls *bool, started *bool,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *skillChat) handleToolResponses(ev *event.Event) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *skillChat) handleContent(
	ev *event.Event, toolCalls *bool, started *bool, full *string,
) {
	_ = "STUB: not implemented"
	return
}

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }
