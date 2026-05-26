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
	"iter"
	"net/http"

	openapi "github.com/getkin/kin-openapi/openapi3"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

func operationName(op *openapi.Operation, path, method string) string {
	_ = "STUB: not implemented"
	return ""
}

func operationDesc(op *openapi.Operation) string { _ = "STUB: not implemented"; return "" }

// Operation represents an operation in the spec.
// parameters are collected from the following sources:
// - path parameters
// - operation parameters
// - request body parameters
// - response parameters
type Operation struct {
	name        string
	description string

	endpoint        *operationEndpoint
	operationParams []*APIParameter
	responseParam   *APIParameter
	originOperation *openapi.Operation
}

func newOperation(name, desc string, endpoint *operationEndpoint, originOperation *openapi.Operation) *Operation {
	_ = "STUB: not implemented"
	return nil
}

func (o *Operation) toolInputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func (o *Operation) toolOutputSchema() *tool.Schema { _ = "STUB: not implemented"; return nil }

func (o *Operation) collectSpecOperationParameters(params openapi.Parameters) {
	_ = "STUB: not implemented"
	return
}

func (o *Operation) collectSpecRequestParameters(request *openapi.RequestBodyRef) {
	_ = "STUB: not implemented"
	return
}

func (o *Operation) collectOperationParameters(params []*APIParameter) {
	_ = "STUB: not implemented"
	return
}

func (o *Operation) collectResponseParameter(responses *openapi.Responses) {
	_ = "STUB: not implemented"
	return
}

type operationEndpoint struct {
	baseURL string
	path    string
	method  string
}

func newOperationEndpoint(baseURL, path, method string) *operationEndpoint {
	_ = "STUB: not implemented"
	return nil
}

// convertOpenAPISchemaToToolSchema converts kin-openapi/openapi3.Schema to tool.Schema
func convertOpenAPISchemaToToolSchema(openapiSchema *openapi.Schema) *tool.Schema {
	_ = "STUB: not implemented"
	return nil
}

// Convert type
// openapi scheme types include: "string", "number", "integer", "boolean", "array", "object"
// tool.Schema type is similar, we take the first type if multiple are defined

// Default type

// Convert description

// Convert properties

// Convert required fields

// Convert items (for array types)

// Convert additional properties

// Convert default value

// Convert enum values

// ParameterLocation defines the location of a parameter.
type ParameterLocation string

const (
	PathParameter     ParameterLocation = "path"
	QueryParameter    ParameterLocation = "query"
	HeaderParameter   ParameterLocation = "header"
	BodyParameter     ParameterLocation = "body"
	CookieParameter   ParameterLocation = "cookie"
	ResponseParameter ParameterLocation = "response"
)

// APIParameter represents an API parameter.
type APIParameter struct {
	OriginalName string            `json:"original_name"`
	Description  string            `json:"description"`
	Location     ParameterLocation `json:"location"`
	Required     bool              `json:"required"`

	level  string
	schema *openapi.Schema
}

var (
	pathItemMethods = [...]string{
		http.MethodConnect,
		http.MethodDelete,
		http.MethodGet,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPost,
		http.MethodPut,
		http.MethodTrace,
	}
)

func methodOperations(p *openapi.PathItem) iter.Seq2[string, *openapi.Operation] {
	_ = "STUB: not implemented"
	return nil
}
