//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package evalset

// defaultBaseDir is the default base directory for eval set files.
const defaultBaseDir = "."

// Options configure the local evaluation set manager.
type Options struct {
	BaseDir string  // BaseDir is the base directory for eval set files.
	Locator Locator // Locator is the locator for eval set files.
}

// NewOptions constructs Options with the default values.
func NewOptions(opts ...Option) *Options { _ = "STUB: not implemented"; return nil }

// Option is a functional option for configuring the eval set manager.
type Option func(*Options)

// WithBaseDir sets the root directory for storing eval set JSON files.
func WithBaseDir(dir string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithLocator sets the locator.
func WithLocator(p Locator) Option { _ = "STUB: not implemented"; return *new(Option) }
