//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package skills

import (
	"context"
	"sync"

	"trpc.group/trpc-go/trpc-agent-go/skill"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/internal/deps"
)

const (
	skillsPathEnvName = "PATH"

	skillsBinPathFixHint = "" +
		"move the binary into an existing PATH dir or restart " +
		"openclaw with PATH including its directory"
	skillsAnyBinPathFixHint = "" +
		"install one of these binaries into an existing PATH " +
		"dir or restart openclaw with PATH including its " +
		"directory"

	missingBinsReasonFormat = "" +
		"missing bins: %s; searched PATH dirs: %s; fix: %s"
	missingAnyBinsReasonFormat = "" +
		"missing anyBins (need one): %s; searched PATH dirs: %s; " +
		"fix: %s"

	emptySkillsSearchDirs = "(empty)"
)

type SkillConfig struct {
	Enabled *bool
	APIKey  string
	Env     map[string]string
}

type Repository struct {
	mu sync.RWMutex

	base  skill.Repository
	roots []string

	eligible map[string]struct{}
	reasons  map[string]string

	baseDirs map[string]string

	metas    map[string]*openClawMetadata
	skillKey map[string]string

	debug bool

	configKeys map[string]struct{}

	allowBundled map[string]struct{}
	bundledRoot  string

	skillConfigs map[string]SkillConfig
}

type Option func(*Repository)

func WithDebug(debug bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithConfigKeys(keys []string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithBundledSkillsRoot(root string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithAllowBundled(allow []string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSkillConfigs(cfg map[string]SkillConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func NewRepository(roots []string, opts ...Option) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) Summaries() []skill.Summary { _ = "STUB: not implemented"; return nil }

func (r *Repository) Get(name string) (*skill.Skill, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) Path(name string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *Repository) SkillRunEnv(
	_ context.Context,
	skillName string,
) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) DependencySources(
	names []string,
) ([]deps.Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Repository) Refresh() error { _ = "STUB: not implemented"; return nil }

func (r *Repository) SetSkillEnabled(
	configKey string,
	enabled bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Repository) indexLocked() { _ = "STUB: not implemented"; return }

func evaluateSkill(
	skillName string,
	meta openClawMetadata,
	hasOpenClawMeta bool,
	configKeys map[string]struct{},
	cfg SkillConfig,
	isBundled bool,
	allowBundled map[string]struct{},
) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func evaluateOpenClawRequirements(
	meta openClawMetadata,
	configKeys map[string]struct{},
	cfg SkillConfig,
) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func evaluateOpenClawOS(allowlist []string) (bool, string) {
	_ = "STUB: not implemented"
	return false, ""
}

func normalizeOpenClawOS(raw string) string { _ = "STUB: not implemented"; return "" }

func evaluateRequiredBins(bins []string) string { _ = "STUB: not implemented"; return "" }

func evaluateRequiredAnyBins(bins []string) string { _ = "STUB: not implemented"; return "" }

func formatSkillsSearchDirs() string { _ = "STUB: not implemented"; return "" }

func evaluateRequiredEnv(
	names []string,
	primaryEnv string,
	cfg SkillConfig,
) string {
	_ = "STUB: not implemented"
	return ""
}

func evaluateRequiredConfig(
	keys []string,
	available map[string]struct{},
) string {
	_ = "STUB: not implemented"
	return ""
}

func normalizeConfigKeys(keys []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }

func normalizeConfigKey(raw string) string { _ = "STUB: not implemented"; return "" }

func hasConfigKey(keys map[string]struct{}, want string) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Repository) resolveSkillConfig(
	skillKey string,
	skillName string,
) (SkillConfig, bool) {
	_ = "STUB: not implemented"
	return *new(SkillConfig), false
}

func (r *Repository) isBundledSkill(baseDir string) bool { _ = "STUB: not implemented"; return false }

func normalizeAllowlist(allow []string) map[string]struct{} { _ = "STUB: not implemented"; return nil }

func isAllowedByAllowlist(
	allow map[string]struct{},
	skillKey string,
	skillName string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func normalizeSkillConfigs(
	cfg map[string]SkillConfig,
) map[string]SkillConfig {
	_ = "STUB: not implemented"
	return nil
}

func copySkillEnv(env map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

func containsString(values []string, want string) bool { _ = "STUB: not implemented"; return false }

func containsSource(sources []deps.Source, want string) bool {
	_ = "STUB: not implemented"
	return false
}

func normalizeSkillNames(names []string) []string { _ = "STUB: not implemented"; return nil }

const (
	envLDPreload           = "LD_PRELOAD"
	envLDLibraryPath       = "LD_LIBRARY_PATH"
	envDYLDInsertLibraries = "DYLD_INSERT_LIBRARIES"
	envDYLDLibraryPath     = "DYLD_LIBRARY_PATH"
	envDYLDForceFlatNS     = "DYLD_FORCE_FLAT_NAMESPACE"
	envOpenSSLConf         = "OPENSSL_CONF"
)

func isBlockedSkillEnvKey(key string) bool { _ = "STUB: not implemented"; return false }

var _ skill.Repository = (*Repository)(nil)
var _ skill.RefreshableRepository = (*Repository)(nil)
