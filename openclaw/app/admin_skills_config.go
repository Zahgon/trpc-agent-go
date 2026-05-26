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
	"gopkg.in/yaml.v3"

	ocskills "trpc.group/trpc-go/trpc-agent-go/openclaw/internal/skills"
)

func (p *adminSkillsProvider) SkillsStatus() (ocskills.StatusReport, error) {
	_ = "STUB: not implemented"
	return *new(ocskills.StatusReport), nil
}

func (p *adminSkillsProvider) SkillsConfigPath() string { _ = "STUB: not implemented"; return "" }

func (p *adminSkillsProvider) SkillsRefreshable() bool { _ = "STUB: not implemented"; return false }

func (p *adminSkillsProvider) RefreshSkills() error { _ = "STUB: not implemented"; return nil }

func (p *adminSkillsProvider) SetSkillEnabled(
	configKey string,
	enabled bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func cloneAdminSkillConfigs(
	src map[string]ocskills.SkillConfig,
) map[string]ocskills.SkillConfig {
	_ = "STUB: not implemented"
	return nil
}

func setSkillEnabledInConfig(
	path string,
	configKey string,
	enabled bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeConfigDocument(data []byte) (yaml.Node, error) {
	_ = "STUB: not implemented"
	return *new(yaml.Node), nil
}

func ensureDocumentMapping(doc *yaml.Node) (*yaml.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ensureMappingChild(
	parent *yaml.Node,
	key string,
) (*yaml.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setMappingBool(
	parent *yaml.Node,
	key string,
	value bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func writeConfigDocument(path string, doc *yaml.Node) error { _ = "STUB: not implemented"; return nil }
