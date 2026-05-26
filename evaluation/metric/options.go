//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package metric

// defaultBaseDir is the default base directory for metric files.
const defaultBaseDir = "."

// Options holds the configuration for the metric manager.
type Options struct {
	// BaseDir is the base directory for metric files.
	BaseDir string
	// Locator is the locator for metric files.
	Locator Locator
}

// NewOptions creates a Options with the default values.
func NewOptions(opts ...Option) *Options { _ = "STUB: not implemented"; return nil }

// Option defines a function type for configuring the metric manager.
type Option func(*Options)

// WithBaseDir sets the base directory.
func WithBaseDir(dir string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLocator sets the locator.
func WithLocator(l Locator) Option { _ = "STUB: not implemented"; return *new(Option) }
