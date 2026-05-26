//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package langfuse provides Langfuse integration with custom span transformations.
package langfuse

// Option is a function that configures Start options.
type Option func(*config)

// WithSecretKey sets the Langfuse secret key.
func WithSecretKey(secretKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPublicKey sets the Langfuse public key.
func WithPublicKey(publicKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithHost sets the Langfuse host endpoint.
// The provided host should be in "hostname:port" format (no scheme or path).
// For cloud.langfuse.com, use "cloud.langfuse.com:443".
// For local development, use "localhost:3000".
//
// Example:
//
//	WithHost("cloud.langfuse.com:443")      // Production
//	WithHost("localhost:3000")              // Local development
func WithHost(host string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithInsecure configures the exporter to use insecure connections.
// This should only be used for development/testing environments.
// By default, secure connections are used.
func WithInsecure() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithObservationLeafValueMaxBytes configures the max byte length for each leaf
// value in Langfuse observation JSON payloads (and plain string observation values).
//
// If this option is not set, truncation is disabled by default.
// If maxBytes is 0, it truncates everything.
// If maxBytes < 0, truncation is disabled.
func WithObservationLeafValueMaxBytes(maxBytes int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// config holds Langfuse configuration options.
type config struct {
	secretKey                    string
	publicKey                    string
	host                         string
	insecure                     bool
	maxObservationLeafValueBytes *int
}

// newConfigFromEnv creates a Langfuse config from environment variables.
// Supported environment variables:
//
//	LANGFUSE_SECRET_KEY: Langfuse secret key
//	LANGFUSE_PUBLIC_KEY: Langfuse public key
//	LANGFUSE_HOST: Langfuse host in "hostname:port" format (e.g., "cloud.langfuse.com:443")
//	LANGFUSE_INSECURE: Set to "true" for insecure connections (development only)
//	LANGFUSE_OBSERVATION_LEAF_VALUE_MAX_BYTES: Optional; max byte length for each observation JSON leaf value (unset by default)
func newConfigFromEnv() *config { _ = "STUB: not implemented"; return nil }

// getEnv returns the value of the environment variable or the default if not set.
func getEnv(key, defaultValue string) string { _ = "STUB: not implemented"; return "" }

func getEnvIntPtr(key string) *int { _ = "STUB: not implemented"; return nil }
