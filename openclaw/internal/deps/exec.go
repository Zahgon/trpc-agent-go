//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package deps

import (
	"context"
	"os/exec"
)

type executableSpec struct {
	path string
}

func planStepCommand(
	toolchain Toolchain,
	step Step,
) (*exec.Cmd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pythonExecCommand(
	pythonPath string,
	args ...string,
) (*exec.Cmd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func combinedOutputContext(
	ctx context.Context,
	cmd *exec.Cmd,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pythonCommandSpec(command string) (executableSpec, error) {
	_ = "STUB: not implemented"
	return *new(executableSpec), nil
}

func systemCommandSpec(manager string) (executableSpec, error) {
	_ = "STUB: not implemented"
	return *new(executableSpec), nil
}

func resolveExecutable(
	command string,
	fallback string,
	searchPath string,
) (executableSpec, error) {
	_ = "STUB: not implemented"
	return *new(executableSpec), nil
}

func lookPath(
	name string,
	searchPath string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func validateExecutablePath(path string) (executableSpec, error) {
	_ = "STUB: not implemented"
	return *new(executableSpec), nil
}

func newExecCommand(
	spec executableSpec,
	args ...string,
) *exec.Cmd {
	_ = "STUB: not implemented"
	return nil
}

func mergeStepEnv(
	base []string,
	overrides map[string]string,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func envValue(
	env []string,
	key string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func commandDir(command string) string { _ = "STUB: not implemented"; return "" }

func commandBase(command string) string { _ = "STUB: not implemented"; return "" }
