//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package openapi

import (
	"context"
	"net/url"

	openapi "github.com/getkin/kin-openapi/openapi3"
)

// Loader loads the OpenAPI spec.
type Loader interface {
	// Load loads the OpenAPI spec.
	Load(ctx context.Context) (*openapi.T, error)
}

// LoaderOption is the option for the loader.
type LoaderOption func(*openapi.Loader)

// WithExternalRefs sets whether to allow external $ref references.
func WithExternalRefs(allow bool) LoaderOption {
	_ = "STUB: not implemented"
	return *new(LoaderOption)
}

// WithReadFromURI sets the function to read from URI.
func WithReadFromURI(f openapi.ReadFromURIFunc) LoaderOption {
	_ = "STUB: not implemented"
	return *new(LoaderOption)
}

type dataLoader struct {
	loader *openapi.Loader
	data   []byte
}

// Load loads the OpenAPI spec from io.Reader.
func (d *dataLoader) Load(ctx context.Context) (*openapi.T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewDataLoader creates a new data spec loader.
func NewDataLoader(data []byte, opts ...LoaderOption) (*dataLoader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type fileLoader struct {
	loader *openapi.Loader
	path   string
}

// Load loads the OpenAPI spec from file.
func (f *fileLoader) Load(ctx context.Context) (*openapi.T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewFileLoader creates a new file spec loader.
func NewFileLoader(filePath string, opts ...LoaderOption) (*fileLoader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type urlLoader struct {
	loader   *openapi.Loader
	location *url.URL
}

// Load loads the OpenAPI spec from url.
func (u *urlLoader) Load(ctx context.Context) (*openapi.T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewURILoader creates a new url spec loader.
func NewURILoader(uri string, opts ...LoaderOption) (*urlLoader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newDocProcessor(doc *openapi.T) *docProcessor { _ = "STUB: not implemented"; return nil }

type docProcessor struct {
	doc        *openapi.T
	operations []*Operation
}

// InfoTitle returns the title of the spec.
func (s *docProcessor) InfoTitle() string { _ = "STUB: not implemented"; return "" }

const defaultBaseURL = "/"

// processOperations processes the operations in the spec.
// Currently, only the first server URL is used as the base URL.
// And variables are not supported.
func (s *docProcessor) processOperations() error { _ = "STUB: not implemented"; return nil }

// TODO: security scheme, if any

// TODO: add path-level parameters
