//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package jsonschema

type options struct {
	strict bool
}

// Option configures a Generator.
type Option func(*options)

// WithStrict enables strict structured-output-compatible schema generation.
func WithStrict() Option { _ = "STUB: not implemented"; return *new(Option) }

func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }
