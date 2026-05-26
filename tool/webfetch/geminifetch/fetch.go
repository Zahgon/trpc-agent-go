//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package geminifetch provides a Gemini webfetch tool.
package geminifetch

import (
	"context"

	"google.golang.org/genai"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// Option configures the GeminiFetch tool.
type Option func(*config)

// modelCaller is the interface for the model caller. for testing purposes, we can inject a stub model caller.
type modelCaller interface {
	GenerateContent(ctx context.Context, model string, contents []*genai.Content, config *genai.GenerateContentConfig) (*genai.GenerateContentResponse, error)
}

type config struct {
	apiKey      string
	model       string
	client      *genai.Client
	modelCaller modelCaller
}

// WithAPIKey sets the Google AI API key.
// If not provided, it will use the GEMINI_API_KEY environment variable.
func WithAPIKey(apiKey string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithClient sets a custom Gemini client.
func WithClient(client *genai.Client) Option { _ = "STUB: not implemented"; return *new(Option) }

// fetchRequest is the input for the tool.
type fetchRequest struct {
	Prompt string `json:"prompt" jsonschema:"description=Prompt that includes URLs to fetch and instructions for processing. URLs will be automatically detected and fetched by Gemini."`
}

// fetchResponse is the output.
type fetchResponse struct {
	Content            string              `json:"content"`
	URLContextMetadata *urlContextMetadata `json:"url_context_metadata,omitempty"`
}

type urlContextMetadata struct {
	URLMetadata []urlMetadata `json:"url_metadata"`
}

type urlMetadata struct {
	RetrievedURL       string `json:"retrieved_url"`
	URLRetrievalStatus string `json:"url_retrieval_status"`
}

// NewTool creates the Gemini web-fetch tool.
// This tool uses Gemini's URL Context feature to fetch and process web content.
// modelName: the Gemini model-id to use.
func NewTool(modelName string, opts ...Option) (tool.CallableTool, error) {
	_ = "STUB: not implemented"
	return *new(tool.CallableTool), nil
}

type geminiFetchTool struct {
	cfg *config
}

func newGeminiFetchTool(cfg *config) *geminiFetchTool { _ = "STUB: not implemented"; return nil }

func (t *geminiFetchTool) fetch(ctx context.Context, req fetchRequest) (fetchResponse, error) {
	_ = "STUB: not implemented"
	return *new(fetchResponse), nil
}

// Resolve the model caller so tests can inject a stub without hitting the API.

// Note: Client doesn't have Close method in this version

// Build content parts with the user's prompt
// Gemini will automatically detect URLs in the prompt and fetch them

// Configure with URL context tool
// This enables Gemini to automatically fetch URLs mentioned in the prompt

// Generate content

// Extract content from response

// Extract URL metadata if available
