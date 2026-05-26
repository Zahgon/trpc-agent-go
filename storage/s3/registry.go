//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package s3

import "sync"

var (
	registryMu sync.RWMutex
	s3Registry = make(map[string][]ClientBuilderOpt)
)

// RegisterS3Instance registers a named S3 instance with its configuration options.
// If an instance with the same name already exists, it will be overwritten.
//
// Example:
//
//	s3.RegisterS3Instance("production",
//	    s3.WithBucket("my-bucket"),
//	    s3.WithRegion("us-west-2"),
//	    s3.WithCredentials(accessKey, secretKey),
//	)
//
//	// Later, retrieve and use the instance
//	opts, ok := s3.GetS3Instance("production")
//	if ok {
//	    client, err := s3.NewClient(ctx, opts...)
//	}
func RegisterS3Instance(name string, opts ...ClientBuilderOpt) { _ = "STUB: not implemented"; return }

// GetS3Instance retrieves the configuration options for a named S3 instance.
// Returns a copy of the options and true if found, or nil and false if not found.
func GetS3Instance(name string) ([]ClientBuilderOpt, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Copy to prevent external modifications

// UnregisterS3Instance removes a named S3 instance from the registry.
func UnregisterS3Instance(name string) { _ = "STUB: not implemented"; return }

// ListS3Instances returns a list of all registered instance names.
func ListS3Instances() []string { _ = "STUB: not implemented"; return nil }
