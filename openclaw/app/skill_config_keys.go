//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package app

import (
	"gopkg.in/yaml.v3"
)

const (
	configKeyChannelsPrefix = "channels."

	configKeyToolsPrefix         = "tools."
	configKeyToolProvidersPrefix = "tools.providers."
	configKeyToolSetsPrefix      = "tools.toolsets."

	configKeyPluginsEntriesPrefix = "plugins.entries."
	configKeyPluginsEnabledSuffix = ".enabled"
	configKeyPluginsConfigPrefix  = ".config"

	configKeyExecCommand = "exec_command"
	configKeyWriteStdin  = "write_stdin"
	configKeyKillSession = "kill_session"
	configKeyMessage     = "message"
	configKeyCron        = "cron"
	configKeyLocalExec   = "local_exec"
)

func resolveSkillConfigKeys(opts runOptions) []string { _ = "STUB: not implemented"; return nil }

func addPluginSpecsConfigKeys(
	set map[string]struct{},
	prefix string,
	specs []pluginSpec,
) {
	_ = "STUB: not implemented"
	return
}

func addToolSurfaceKeys(set map[string]struct{}, opts runOptions) {
	_ = "STUB: not implemented"
	return
}

func addPluginConfigNodeKeys(
	set map[string]struct{},
	prefix string,
	node *yaml.Node,
) {
	_ = "STUB: not implemented"
	return
}

func addYAMLConfigKeys(
	set map[string]struct{},
	prefix string,
	node *yaml.Node,
) bool {
	_ = "STUB: not implemented"
	return false
}

func isTruthyScalar(node *yaml.Node) bool { _ = "STUB: not implemented"; return false }

func normalizeConfigSegment(raw string) string { _ = "STUB: not implemented"; return "" }

func addConfigKey(set map[string]struct{}, key string) { _ = "STUB: not implemented"; return }
