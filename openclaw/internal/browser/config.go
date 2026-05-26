//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package browser

import (
	"time"

	mcptool "trpc.group/trpc-go/trpc-agent-go/tool/mcp"
)

const (
	ToolName = "browser"

	defaultProfileName = "openclaw"

	transportStdio      = "stdio"
	transportSSE        = "sse"
	transportStreamable = "streamable"
	transportStreamHTTP = "streamable_http"
)

type reconnectConfig struct {
	Enabled     bool `yaml:"enabled,omitempty"`
	MaxAttempts int  `yaml:"max_attempts,omitempty"`
}

// NodeConfig describes one remote browser server node.
type NodeConfig struct {
	ID        string `yaml:"id,omitempty"`
	Name      string `yaml:"name,omitempty"`
	ServerURL string `yaml:"server_url,omitempty"`
	AuthToken string `yaml:"auth_token,omitempty"`
}

// ProfileConfig describes one browser execution profile.
type ProfileConfig struct {
	Name             string            `yaml:"name,omitempty"`
	Description      string            `yaml:"description,omitempty"`
	AuthToken        string            `yaml:"auth_token,omitempty"`
	BrowserServerURL string            `yaml:"browser_server_url,omitempty"`
	Transport        string            `yaml:"transport,omitempty"`
	ServerURL        string            `yaml:"server_url,omitempty"`
	Headers          map[string]string `yaml:"headers,omitempty"`
	Command          string            `yaml:"command,omitempty"`
	Args             []string          `yaml:"args,omitempty"`
	Timeout          time.Duration     `yaml:"timeout,omitempty"`
	Reconnect        *reconnectConfig  `yaml:"reconnect,omitempty"`
}

// Config describes the native browser tool configuration.
type Config struct {
	DefaultProfile   string          `yaml:"default_profile,omitempty"`
	EvaluateEnabled  *bool           `yaml:"evaluate_enabled,omitempty"`
	ServerURL        string          `yaml:"server_url,omitempty"`
	AuthToken        string          `yaml:"auth_token,omitempty"`
	SandboxServerURL string          `yaml:"sandbox_server_url,omitempty"`
	SandboxAuthToken string          `yaml:"sandbox_auth_token,omitempty"`
	AllowedDomains   []string        `yaml:"allowed_domains,omitempty"`
	BlockedDomains   []string        `yaml:"blocked_domains,omitempty"`
	AllowLoopback    *bool           `yaml:"allow_loopback,omitempty"`
	AllowPrivateNet  *bool           `yaml:"allow_private_networks,omitempty"`
	AllowFileURLs    *bool           `yaml:"allow_file_urls,omitempty"`
	Nodes            []NodeConfig    `yaml:"nodes,omitempty"`
	Profiles         []ProfileConfig `yaml:"profiles,omitempty"`
}

type resolvedConfig struct {
	DefaultProfile  string
	EvaluateEnabled bool
	Navigation      navigationPolicy
	HostServer      *serverTargetConfig
	SandboxServer   *serverTargetConfig
	NodeTargets     map[string]serverTargetConfig
	Profiles        []resolvedProfile
}

type resolvedProfile struct {
	Name             string
	Description      string
	BrowserServerURL string
	AuthToken        string
	Connection       mcptool.ConnectionConfig
	Reconnect        *reconnectConfig
}

type serverTargetConfig struct {
	ID        string
	ServerURL string
	AuthToken string
}

func resolveConfig(cfg Config) (resolvedConfig, error) {
	_ = "STUB: not implemented"
	return *new(resolvedConfig), nil
}

func resolveNavigationPolicy(cfg Config) navigationPolicy {
	_ = "STUB: not implemented"
	return *new(navigationPolicy)
}

func resolveProfile(
	cfg ProfileConfig,
	index int,
	allowEmptyConnection bool,
) (resolvedProfile, error) {
	_ = "STUB: not implemented"
	return *new(resolvedProfile), nil
}

func resolveServerTarget(
	id string,
	serverURL string,
	authToken string,
) *serverTargetConfig {
	_ = "STUB: not implemented"
	return nil
}

func resolveNodeTargets(
	nodes []NodeConfig,
) map[string]serverTargetConfig {
	_ = "STUB: not implemented"
	return nil
}

func validateConnection(cfg mcptool.ConnectionConfig) error { _ = "STUB: not implemented"; return nil }
