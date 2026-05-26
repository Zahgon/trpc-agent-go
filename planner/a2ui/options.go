// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package a2ui

const defaultInstruction = `A2UI server-to-client output MUST be JSONL-compatible.
Each outbound text content message must contain exactly one complete JSON object and must be emitted as one complete line.
Do not split one message across multiple chunks.
Each message must be independently valid JSON.
Only these message keys are allowed: beginRendering, surfaceUpdate, dataModelUpdate, deleteSurface.
Do not add markdown fences, code fences, or any extra explanatory text.`

// Option configures the A2UI planner.
type Option func(*options)

// options stores A2UI planner extension points.
type options struct {
	instruction                       string
	serverToClientWithStandardCatalog string
	clientToServer                    string
	clientCapabilitiesSchema          string
	serverToClient                    string
	standardCatalogDefinition         string
	catalogDescriptionSchema          string
}

// newOptions creates a new options instance.
func newOptions(opt ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithInstruction sets the planner instruction appended to the system prompt.
func WithInstruction(instruction string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithServerToClientWithStandardCatalogSchema sets the standard catalog server-to-client schema.
func WithServerToClientWithStandardCatalogSchema(schema string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithClientToServerSchema sets the client-to-server schema.
func WithClientToServerSchema(schema string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithClientCapabilitiesSchema sets the client capabilities schema.
func WithClientCapabilitiesSchema(schema string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithServerToClientSchema sets the server-to-client schema.
func WithServerToClientSchema(schema string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithStandardCatalogDefinition sets the standard catalog definition for planning.
func WithStandardCatalogDefinition(definition string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithCatalogDescriptionSchema sets the catalog description schema.
func WithCatalogDescriptionSchema(schema string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
