//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package query

import "context"

// PassthroughEnhancer is a simple enhancer that returns the original query unchanged.
type PassthroughEnhancer struct {
	// Configuration options can be added here in the future.
}

// Option represents a functional option for configuring PassthroughEnhancer.
type Option func(*PassthroughEnhancer)

// NewPassthroughEnhancer creates a new passthrough query enhancer with options.
func NewPassthroughEnhancer(opts ...Option) *PassthroughEnhancer {
	_ = "STUB: not implemented"
	return nil
}

// Apply options.

// EnhanceQuery implements the Enhancer interface by returning the original query.
func (p *PassthroughEnhancer) EnhanceQuery(ctx context.Context, req *Request) (*Enhanced, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
