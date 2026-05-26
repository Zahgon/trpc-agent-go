//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package weknora

import (
	"time"

	"github.com/Tencent/WeKnora/client"

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

// Option configures the WeKnoraAgent
type Option func(*WeKnoraAgent)

// WithBaseUrl sets the base URL of the WeKnora service
func WithBaseUrl(baseUrl string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithToken sets the authentication token for WeKnora service
func WithToken(token string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithName sets the name of agent
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDescription sets the agent description
func WithDescription(description string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAgentID sets the custom agent ID for WeKnora
func WithAgentID(agentID string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithKnowledgeBaseIDs sets the knowledge base IDs for WeKnora
func WithKnowledgeBaseIDs(ids []string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithWebSearchEnabled sets whether to enable web search
func WithWebSearchEnabled(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTimeout sets the timeout for WeKnora requests
func WithTimeout(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGetWeKnoraClientFunc sets a custom function to create WeKnora client for each invocation.
func WithGetWeKnoraClientFunc(fn func(*agent.Invocation) (*client.Client, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
