//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package n8n

import (
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/agent"
)

// Option configures the N8nAgent.
type Option func(*N8nAgent)

// WithWebhookURL sets the webhook URL of the n8n service.
func WithWebhookURL(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithName sets the name of the agent.
func WithName(name string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDescription sets the agent description.
func WithDescription(description string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAuthType sets the authentication type for n8n webhook requests.
func WithAuthType(authType AuthType) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAuthConfig sets the authentication configuration for n8n webhook requests.
func WithAuthConfig(config *AuthConfig) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCustomRequestConverter sets a custom request converter.
func WithCustomRequestConverter(converter RequestConverter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCustomResponseConverter sets a custom response converter.
func WithCustomResponseConverter(converter ResponseConverter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithEnableStreaming explicitly controls whether to use streaming protocol.
// If not set (nil), the agent defaults to non-streaming.
// This option can be overridden per-run via RunOptions.Stream.
func WithEnableStreaming(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStreamingChannelBufSize sets the buffer size of the streaming event channel.
func WithStreamingChannelBufSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStreamingRespHandler sets a handler function to process streaming responses.
func WithStreamingRespHandler(handler StreamingRespHandler) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHTTPClient sets a custom HTTP client for the agent.
func WithHTTPClient(client *http.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithGetHTTPClientFunc sets a custom function to create an HTTP client for each invocation.
func WithGetHTTPClientFunc(fn func(*agent.Invocation) (*http.Client, error)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTransferStateKey appends keys from session state to transfer to the n8n request inputs.
func WithTransferStateKey(key ...string) Option { _ = "STUB: not implemented"; return *new(Option) }
