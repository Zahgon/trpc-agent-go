//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

// Package s3 provides a reusable S3 client for storage operations.
// It supports AWS S3 and S3-compatible services like MinIO, DigitalOcean Spaces,
// and Cloudflare R2.
package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Client defines the interface for S3 storage operations.
type Client interface {
	PutObject(ctx context.Context, key string, data []byte, contentType string) error
	GetObject(ctx context.Context, key string) ([]byte, string, error)
	ListObjects(ctx context.Context, prefix string) ([]string, error)
	DeleteObjects(ctx context.Context, keys []string) error
	Close() error
}

// s3API defines the subset of AWS S3 API operations used by the client.
// This interface allows mocking the AWS SDK in unit tests.
type s3API interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	DeleteObjects(ctx context.Context, params *s3.DeleteObjectsInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error)
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

// client implements the Client interface using AWS SDK v2.
type client struct {
	s3     s3API
	bucket string
}

// NewClient creates a new S3 client with the given options.
func NewClient(ctx context.Context, opts ...ClientBuilderOpt) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// Custom endpoints need a region; use default fallback

// PutObject uploads an object to S3.
func (c *client) PutObject(ctx context.Context, key string, data []byte, contentType string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetObject downloads an object from S3.
func (c *client) GetObject(ctx context.Context, key string) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// ListObjects lists object keys with the given prefix.
func (c *client) ListObjects(ctx context.Context, prefix string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteObjects deletes multiple objects in batches of 1000.
func (c *client) DeleteObjects(ctx context.Context, keys []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Close implements the Client interface (no-op for S3).
func (c *client) Close() error {
	_ = "STUB: not implemented"

	// wrapError converts AWS SDK errors to sentinel errors.
	return nil
}

func wrapError(err error) error { _ = "STUB: not implemented"; return nil }
