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
)

const (
	stepKindSystem   = "system"
	stepKindPython   = "python"
	stepKindVenv     = "venv"
	stepKindCommand  = "command"
	stepKindDownload = "download"

	venvFlagSystemSitePackages = "--system-site-packages"
)

const (
	stepStatusApplied  = "applied"
	stepStatusDeferred = "deferred"
	stepStatusFailed   = "failed"

	envGoBin          = "GOBIN"
	defaultToolsDir   = "tools"
	pythonInstallName = "Install Python packages"
)

type Step struct {
	Label           string            `json:"label"`
	Kind            string            `json:"kind"`
	Command         []string          `json:"command"`
	CommandLine     string            `json:"command_line"`
	RequiresRoot    bool              `json:"requires_root,omitempty"`
	Env             map[string]string `json:"env,omitempty"`
	EnsureDirs      []string          `json:"ensure_dirs,omitempty"`
	URL             string            `json:"url,omitempty"`
	TargetPath      string            `json:"target_path,omitempty"`
	Archive         string            `json:"archive,omitempty"`
	Extract         bool              `json:"extract,omitempty"`
	StripComponents int               `json:"strip_components,omitempty"`
}

type Plan struct {
	Platform   Platform  `json:"platform"`
	Toolchain  Toolchain `json:"toolchain"`
	Profiles   []string  `json:"profiles,omitempty"`
	Steps      []Step    `json:"steps,omitempty"`
	Unresolved Missing   `json:"unresolved,omitempty"`
}

type ApplyResult struct {
	Steps []StepResult `json:"steps,omitempty"`
}

type StepResult struct {
	Step     Step   `json:"step"`
	Status   string `json:"status,omitempty"`
	Output   string `json:"output,omitempty"`
	Error    string `json:"error,omitempty"`
	ExitCode int    `json:"exit_code"`
}

type sourceInstallAction struct {
	SourceName string
	Action     InstallAction
}

func BuildPlan(stateDir string, profiles []string) (Plan, error) {
	_ = "STUB: not implemented"
	return *new(Plan), nil
}

func BuildPlanForSources(
	stateDir string,
	names []string,
	sources []Source,
) (Plan, error) {
	_ = "STUB: not implemented"
	return *new(Plan), nil
}

func ApplyPlan(
	ctx context.Context,
	plan Plan,
) (ApplyResult, error) {
	_ = "STUB: not implemented"
	return *new(ApplyResult), nil
}

func collectSystemPackages(
	manager string,
	sources []Source,
	missing Missing,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func collectPythonPackages(pkgs []PythonPackage) []string { _ = "STUB: not implemented"; return nil }

func collectInstallPythonPackages(
	platform Platform,
	sources []Source,
	missing Missing,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func commandInstallSteps(
	toolchain Toolchain,
	platform Platform,
	sources []Source,
	missing Missing,
) ([]Step, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func downloadInstallSteps(
	toolchain Toolchain,
	platform Platform,
	sources []Source,
	missing Missing,
) ([]Step, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func selectInstallActionsForMissing(
	platform Platform,
	sources []Source,
	missing Missing,
	include func(InstallAction) bool,
) []sourceInstallAction {
	_ = "STUB: not implemented"
	return nil
}

func missingBinSet(missing Missing) map[string]struct{} { _ = "STUB: not implemented"; return nil }

func actionNeeded(
	action InstallAction,
	explicitBins map[string]struct{},
	anyBins [][]string,
	satisfiedGroups map[string]struct{},
) bool {
	_ = "STUB: not implemented"
	// Actions without declared bins cannot be filtered precisely, so
	// keep the conservative default and include them in the plan.
	return false
}

func markSatisfiedAnyBinGroups(
	action InstallAction,
	anyBins [][]string,
	satisfiedGroups map[string]struct{},
) {
	_ = "STUB: not implemented"
	return
}

func actionCoversAnyGroupBin(
	action InstallAction,
	group []string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func anyBinGroupKey(group []string) string { _ = "STUB: not implemented"; return "" }

func actionMatchesPlatform(
	action InstallAction,
	goos string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func installActionKey(action InstallAction) string { _ = "STUB: not implemented"; return "" }

func isPythonInstallKind(action InstallAction) bool { _ = "STUB: not implemented"; return false }

func isCommandInstallKind(action InstallAction) bool { _ = "STUB: not implemented"; return false }

func pythonSteps(
	toolchain Toolchain,
	pkgs []string,
) ([]Step, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newSystemStep(manager string, pkgs []string) Step {
	_ = "STUB: not implemented"
	return *new(Step)
}

func unresolvedMissing(
	platform Platform,
	sources []Source,
	missing Missing,
) Missing {
	_ = "STUB: not implemented"
	return *new(Missing)
}

func isActionCoveredByPlanner(
	platform Platform,
	action InstallAction,
) bool {
	_ = "STUB: not implemented"
	return false
}

func packagesForAction(action InstallAction) []string { _ = "STUB: not implemented"; return nil }

func brewPackageName(action InstallAction) string { _ = "STUB: not implemented"; return "" }

func commandInstallStep(
	toolchain Toolchain,
	action InstallAction,
) (Step, error) {
	_ = "STUB: not implemented"
	return *new(Step), nil
}

func goInstallStep(
	toolchain Toolchain,
	action InstallAction,
) (Step, error) {
	_ = "STUB: not implemented"
	return *new(Step), nil
}

func npmInstallStep(
	toolchain Toolchain,
	action InstallAction,
) (Step, error) {
	_ = "STUB: not implemented"
	return *new(Step), nil
}

func actionLabel(action InstallAction, fallback string) string {
	_ = "STUB: not implemented"
	return ""
}

func executePlanStep(
	ctx context.Context,
	toolchain Toolchain,
	step Step,
) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func coversAnyBin(
	bins []string,
	need map[string]struct{},
) bool {
	_ = "STUB: not implemented"
	return false
}

func anyCovered(
	group []string,
	covered map[string]struct{},
) bool {
	_ = "STUB: not implemented"
	return false
}

func mergedPlanEnv(toolchain Toolchain) []string { _ = "STUB: not implemented"; return nil }

func ensureStepWorkingDirs(
	toolchain Toolchain,
	step Step,
) error {
	_ = "STUB: not implemented"
	return nil
}

func hasStepStatus(
	steps []StepResult,
	status string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func setPlanEnv(env []string, key string, value string) []string {
	_ = "STUB: not implemented"
	return nil
}

func managedPythonPath(stateDir string) string { _ = "STUB: not implemented"; return "" }

func shellQuote(parts ...string) string { _ = "STUB: not implemented"; return "" }

func isShellUnsafe(r rune) bool { _ = "STUB: not implemented"; return false }

func profileNames(profiles []Profile) []string { _ = "STUB: not implemented"; return nil }
