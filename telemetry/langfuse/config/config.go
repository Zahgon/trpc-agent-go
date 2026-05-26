//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package config provides shared Langfuse connection configuration
// used by both the telemetry exporter and the prompt provider.
package config

// ConnectionConfig holds the credentials and endpoint needed to connect to a
// Langfuse instance. It is shared across telemetry and prompt-provider packages
// so that both can be configured from a single source.
type ConnectionConfig struct {
	PublicKey string
	SecretKey string
	BaseURL   string // e.g. "https://cloud.langfuse.com"
}

// FromEnv creates a ConnectionConfig from standard Langfuse environment
// variables:
//
//   - LANGFUSE_PUBLIC_KEY
//   - LANGFUSE_SECRET_KEY
//   - LANGFUSE_BASE_URL
//   - LANGFUSE_HOST (fallback; converted to a base URL when BASE_URL is unset)
func FromEnv() ConnectionConfig { _ = "STUB: not implemented"; return *new(ConnectionConfig) }

func hostToBaseURL(host string) string { _ = "STUB: not implemented"; return "" }

func getEnv(key, defaultValue string) string { _ = "STUB: not implemented"; return "" }
