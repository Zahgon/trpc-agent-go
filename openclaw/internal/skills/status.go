//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package skills

import (
	"trpc.group/trpc-go/trpc-agent-go/skill"
)

const (
	statusInstallKindPathHint = "path_hint"

	statusInstallIDRestartPath = "path-hint-restart-with-path"
	statusInstallIDMovePath    = "path-hint-existing-path"

	statusInstallLabelRestartPath = "" +
		"Restart openclaw after updating PATH to include " +
		"the binary directory"
	statusInstallLabelMovePath = "" +
		"Move the binary into one of the existing PATH " +
		"directories"
)

type StatusRequirements struct {
	OS      []string `json:"os,omitempty"`
	Bins    []string `json:"bins,omitempty"`
	AnyBins []string `json:"any_bins,omitempty"`
	Env     []string `json:"env,omitempty"`
	Config  []string `json:"config,omitempty"`
}

type StatusInstallOption struct {
	ID    string   `json:"id,omitempty"`
	Kind  string   `json:"kind,omitempty"`
	Label string   `json:"label,omitempty"`
	Bins  []string `json:"bins,omitempty"`
}

type StatusEntry struct {
	Name               string                `json:"name,omitempty"`
	Description        string                `json:"description,omitempty"`
	SkillKey           string                `json:"skill_key,omitempty"`
	ConfigKey          string                `json:"config_key,omitempty"`
	FilePath           string                `json:"file_path,omitempty"`
	BaseDir            string                `json:"base_dir,omitempty"`
	Source             string                `json:"source,omitempty"`
	Reason             string                `json:"reason,omitempty"`
	Emoji              string                `json:"emoji,omitempty"`
	Homepage           string                `json:"homepage,omitempty"`
	PrimaryEnv         string                `json:"primary_env,omitempty"`
	Bundled            bool                  `json:"bundled"`
	Always             bool                  `json:"always"`
	Disabled           bool                  `json:"disabled"`
	Eligible           bool                  `json:"eligible"`
	BlockedByAllowlist bool                  `json:"blocked_by_allowlist"`
	Requirements       StatusRequirements    `json:"requirements,omitempty"`
	Missing            StatusRequirements    `json:"missing,omitempty"`
	Install            []StatusInstallOption `json:"install,omitempty"`
}

type StatusReport struct {
	Skills []StatusEntry `json:"skills,omitempty"`
	Watch  *WatchStatus  `json:"watch,omitempty"`
}

func BuildStatus(roots []string, opts ...Option) (StatusReport, error) {
	_ = "STUB: not implemented"
	return *new(StatusReport), nil
}

func (r *Repository) Status() StatusReport { _ = "STUB: not implemented"; return *new(StatusReport) }

func (r *Repository) statusEntryForSummary(
	summary skill.Summary,
) StatusEntry {
	_ = "STUB: not implemented"
	return *new(StatusEntry)
}

func statusConfigKey(
	r *Repository,
	skillKey string,
	name string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func requiredStatus(meta *openClawMetadata) StatusRequirements {
	_ = "STUB: not implemented"
	return *new(StatusRequirements)
}

func missingStatus(
	meta *openClawMetadata,
	configKeys map[string]struct{},
	cfg SkillConfig,
) StatusRequirements {
	_ = "STUB: not implemented"
	return *new(StatusRequirements)
}

func missingOS(allowlist []string) []string { _ = "STUB: not implemented"; return nil }

func missingBins(bins []string) []string { _ = "STUB: not implemented"; return nil }

func missingAnyBins(bins []string) []string { _ = "STUB: not implemented"; return nil }

func missingEnv(
	names []string,
	primaryEnv string,
	cfg SkillConfig,
) []string {
	_ = "STUB: not implemented"
	return nil
}

func missingConfig(
	keys []string,
	available map[string]struct{},
) []string {
	_ = "STUB: not implemented"
	return nil
}

func normalizeRequirementConfig(keys []string) []string { _ = "STUB: not implemented"; return nil }

func normalizeStatusInstall(
	actions []openClawInstallEntry,
) []StatusInstallOption {
	_ = "STUB: not implemented"
	return nil
}

func appendRuntimePathInstallHints(
	missing StatusRequirements,
	options []StatusInstallOption,
) []StatusInstallOption {
	_ = "STUB: not implemented"
	return nil
}

func statusInstallOptionExists(
	options []StatusInstallOption,
	id string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func missingStatusBins(missing StatusRequirements) []string { _ = "STUB: not implemented"; return nil }

func installLabel(action openClawInstallEntry) string { _ = "STUB: not implemented"; return "" }

func resolveStatusSource(
	baseDir string,
	bundled bool,
	roots []string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func copyStrings(values []string) []string { _ = "STUB: not implemented"; return nil }
