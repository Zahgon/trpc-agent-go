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
	"io"
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

type openAPITool struct {
	inputSchema  *tool.Schema
	outputSchema *tool.Schema

	operation *Operation
	config    *config
}

func newOpenAPITool(config *config, operation *Operation) tool.CallableTool {
	_ = "STUB: not implemented"
	return *new(tool.CallableTool)
}

// Call executes the API call.
// parameter replace:  "query", "header", "path" or "cookie"
func (o *openAPITool) Call(ctx context.Context, jsonArgs []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Reject trailing non-whitespace data to match json.Unmarshal strictness.

// Read response

// Parse response based on status code

// If JSON parsing fails, return as string

func (o *openAPITool) prepareRequest(ctx context.Context, args map[string]any) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add headers

// paramValueToString converts JSON scalar tool argument values to strings
// for HTTP query, header, and cookie parameters. Non-scalar values are skipped.
func paramValueToString(value any) (string, bool) { _ = "STUB: not implemented"; return "", false }

func makeRequestURL(endpoint *operationEndpoint, pathParams, queryParams map[string]any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func makeRequestBody(operation *Operation, params map[string]any) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

type marshaler interface {
	Marshal(any) ([]byte, error)
}

type jsonMarshaler struct{}

// Marshal marshals the provided data into JSON.
func (j *jsonMarshaler) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type xmlMarshaler struct{}

// Marshal marshals the provided data into XML.
func (j *xmlMarshaler) Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

var supportedMimeTypes = map[string]marshaler{
	"application/json": &jsonMarshaler{},
	"application/xml":  &xmlMarshaler{},
}

func makeRequestCookies(cookieParams map[string]any) []*http.Cookie {
	_ = "STUB: not implemented"
	return nil
}

func makeRequestHeaders(headerParams map[string]any) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Declaration returns the declaration of the tool.
func (o *openAPITool) Declaration() *tool.Declaration { _ = "STUB: not implemented"; return nil }
