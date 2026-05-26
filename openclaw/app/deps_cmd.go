//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package app

import (
	ocdeps "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/deps"
)

const subcmdBootstrap = "bootstrap"

const (
	bootstrapCmdDeps = "deps"
)

type depsCommandOptions struct {
	StateDir           string
	Profiles           string
	Skills             string
	JSON               bool
	Apply              bool
	SkillsRoot         string
	SkillsExtraDirs    string
	SkillsAllowBundled string
}

func runBootstrap(args []string) int { _ = "STUB: not implemented"; return 0 }

func runInspectDeps(args []string) int { _ = "STUB: not implemented"; return 0 }

func runBootstrapDeps(args []string) int { _ = "STUB: not implemented"; return 0 }

func parseDepsCommandOptions(
	top string,
	cmd string,
	args []string,
	allowApply bool,
) (depsCommandOptions, int, error) {
	_ = "STUB: not implemented"
	return *new(depsCommandOptions), 0, nil
}

func resolveDepsSources(
	opts depsCommandOptions,
) (string, []ocdeps.Source, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil, nil
}

func resolveSkillDependencySources(
	stateDir string,
	opts depsCommandOptions,
) ([]ocdeps.Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toolDepsStartupLines(report *ocdeps.Report) []startupLogLine {
	_ = "STUB: not implemented"
	return nil
}

func printDepsReport(
	report ocdeps.Report,
	names []string,
) {
	_ = "STUB: not implemented"
	return
}

func printDepsPlan(plan ocdeps.Plan) { _ = "STUB: not implemented"; return }

func printApplyResult(result ocdeps.ApplyResult) { _ = "STUB: not implemented"; return }

const (
	ocdepsStepStatusApplied  = "applied"
	ocdepsStepStatusDeferred = "deferred"
	ocdepsStepStatusFailed   = "failed"
)

func printApplySteps(
	label string,
	steps []ocdeps.StepResult,
	status string,
) {
	_ = "STUB: not implemented"
	return
}

func printToolchain(toolchain ocdeps.Toolchain) { _ = "STUB: not implemented"; return }

func statusText(found bool, detail string) string { _ = "STUB: not implemented"; return "" }

func anyBinStatusText(status ocdeps.AnyBinStatus) string { _ = "STUB: not implemented"; return "" }

func formatMissing(missing ocdeps.Missing) string { _ = "STUB: not implemented"; return "" }

func printJSON(v any) int { _ = "STUB: not implemented"; return 0 }

func printBootstrapUsage() { _ = "STUB: not implemented"; return }
