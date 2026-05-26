//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package s3

import (
	"trpc.group/trpc-go/trpc-agent-go/log"
	s3storage "trpc.group/trpc-go/trpc-agent-go/storage/s3"
)

type options struct {
	bucket            string
	client            s3storage.Client
	clientBuilderOpts []s3storage.ClientBuilderOpt
	logger            log.Logger
}

// Option is a function that configures the S3 artifact service.
type Option func(*options)

// WithEndpoint sets a custom endpoint URL for S3-compatible services.
func WithEndpoint(endpoint string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRegion sets the AWS region.
func WithRegion(region string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCredentials sets static AWS credentials.
func WithCredentials(accessKeyID, secretAccessKey string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSessionToken sets the session token for temporary credentials (STS).
func WithSessionToken(token string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPathStyle enables path-style addressing (required for MinIO).
func WithPathStyle(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRetries sets the maximum number of retries (default: 3).
func WithRetries(n int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithClient sets a pre-created S3 client.
// When provided, connection options are ignored.
// The caller retains ownership and must close the client separately.
func WithClient(client s3storage.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLogger sets the logger for operational messages.
func WithLogger(logger log.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }
