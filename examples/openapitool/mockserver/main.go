//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main gives an example of mock server.
package main

import (
	"bytes"
	"context"
	"log"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	ctx := context.Background()

	// Load OpenAPI specification
	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
	doc, err := loader.LoadFromFile("../petstore3.yaml")
	if err != nil {
		log.Fatalf("Failed to load OpenAPI spec: %v", err)
	}

	// Validate the specification
	if err := doc.Validate(ctx); err != nil {
		log.Fatalf("OpenAPI spec validation failed: %v", err)
	}

	// Create HTTP server
	handler := &MockServerHandler{doc: doc}
	handler.setupRoutes()

	log.Println("Starting OpenAPI Mock Server on :8080")
	log.Println("Available endpoints:")
	for _, path := range doc.Paths.InMatchingOrder() {
		pathItem := doc.Paths.Find(path)
		if pathItem != nil {
			if pathItem.Get != nil {
				log.Printf("  GET    %s", path)
			}
			if pathItem.Post != nil {
				log.Printf("  POST   %s", path)
			}
			if pathItem.Put != nil {
				log.Printf("  PUT    %s", path)
			}
			if pathItem.Delete != nil {
				log.Printf("  DELETE %s", path)
			}
		}
	}

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

type MockServerHandler struct {
	doc *openapi3.T
	mux *http.ServeMux
}

// responseCapturer captures the response body for logging
type responseCapturer struct {
	http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (rc *responseCapturer) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (rc *responseCapturer) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

func (h *MockServerHandler) setupRoutes() { _ = "STUB: not implemented"; return }

// Setup routes for all paths in the OpenAPI spec

// Add a health check endpoint

func (h *MockServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	// Handle CORS
	return
}

// Create response capturer to log response body

// Handle the request

// Log response body if it exists

func (h *MockServerHandler) createHandler(path string, operation *openapi3.Operation) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Read and log request body

// Restore the body so it can be read again by other handlers

// Set content type based on Accept header or default to JSON

// Generate mock response based on the operation's responses

// Try to get a 200 response first, then fallback to other success codes

// If no success response found, use the first available response

// Generate mock data based on response schema

// Set appropriate status code

// Simple XML response (for demonstration)

// JSON response

func (h *MockServerHandler) generateMockData(response *openapi3.Response, path, operationID string) any {
	_ = "STUB: not implemented"
	// Simple mock data generation based on operation ID and path
	return *new(any)
}

// Generic response for other operations

func (h *MockServerHandler) sendError(w http.ResponseWriter, message string, statusCode int) {
	_ = "STUB: not implemented"
	return
}
