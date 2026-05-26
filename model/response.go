//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package model

import "time"

// Error type constants for ResponseError.Type field.
const (
	ErrorTypeStreamError = "stream_error"
	ErrorTypeAPIError    = "api_error"
	ErrorTypeFlowError   = "flow_error"
	ErrorTypeRunError    = "run_error"
	// ErrorTypeCancelled indicates an error resulting from the caller actively canceling the operation or from a timeout.
	ErrorTypeCancelled = "cancelled"
)

// Object type constants for Response.Object field.
const (
	ObjectTypeError = "error"
	// ObjectTypeToolResponse is the object type for tool response events.
	ObjectTypeToolResponse = "tool.response"
	// ObjectTypePreprocessingBasic is the object type for basic preprocessing events.
	ObjectTypePreprocessingBasic = "preprocessing.basic"
	// ObjectTypePreprocessingContent is the object type for content preprocessing events.
	ObjectTypePreprocessingContent = "preprocessing.content"
	// ObjectTypePreprocessingIdentity is the object type for identity preprocessing events.
	ObjectTypePreprocessingIdentity = "preprocessing.identity"
	// ObjectTypePreprocessingInstruction is the object type for instruction preprocessing events.
	ObjectTypePreprocessingInstruction = "preprocessing.instruction"
	// ObjectTypePreprocessingPlanning is the object type for planning preprocessing events.
	ObjectTypePreprocessingPlanning = "preprocessing.planning"
	// ObjectTypePostprocessingPlanning is the object type for planning postprocessing events.
	ObjectTypePostprocessingPlanning = "postprocessing.planning"
	// ObjectTypePostprocessingCodeExecution is the object type for code execution postprocessing events.
	ObjectTypePostprocessingCodeExecution = "postprocessing.code_execution"
	// ObjectTypeTransfer is the object type for agent transfer events.
	ObjectTypeTransfer = "agent.transfer"
	// ObjectTypeRunnerCompletion is the object type for runner completion events.
	ObjectTypeRunnerCompletion = "runner.completion"
	// ObjectTypeStateUpdate is the object type for state update events.
	ObjectTypeStateUpdate = "state.update"

	// ObjectTypeChatCompletionChunk is the object type for chat completion chunk events.
	ObjectTypeChatCompletionChunk = "chat.completion.chunk"
	// ObjectTypeChatCompletion is the object type for chat completion events.
	ObjectTypeChatCompletion = "chat.completion"
)

// Choice represents a single completion choice.
type Choice struct {
	// Index is the index of the choice.
	Index int `json:"index"`

	// Message is the message content.
	Message Message `json:"message,omitempty"`

	// Delta is the delta message content.
	Delta Message `json:"delta,omitempty"`

	// FinishReason is the reason the choice was finished.
	// "stop", "length", "content_filter", etc.
	FinishReason *string `json:"finish_reason,omitempty"`
}

// TimingInfo represents timing information for token generation.
// It accumulates durations across multiple LLM calls within a flow.
type TimingInfo struct {

	// FirstTokenDuration is the accumulated duration from request start to the first meaningful token.
	// A "meaningful token" is defined as the first chunk containing reasoning content, regular content, or tool calls.
	// Empty chunks are skipped. For multiple LLM calls (e.g., tool calls), this accumulates the duration of each call.
	FirstTokenDuration time.Duration `json:"time_to_first_token,omitempty"`

	// ReasoningDuration is the accumulated duration of reasoning phases (streaming mode only).
	// Measured from the first reasoning chunk to the last reasoning chunk in each LLM call.
	// For non-streaming requests, this field will remain 0 as reasoning duration cannot be measured precisely.
	ReasoningDuration time.Duration `json:"reasoning_duration,omitempty"`
}

// Usage represents token usage information.
type Usage struct {
	// PromptTokens is the number of tokens in the prompt.
	PromptTokens int `json:"prompt_tokens"`

	// CompletionTokens is the number of tokens in the completion.
	CompletionTokens int `json:"completion_tokens"`

	// TotalTokens is the total number of tokens in the response.
	TotalTokens int `json:"total_tokens"`

	// PromptTokensDetails is the details of the prompt tokens.
	PromptTokensDetails PromptTokensDetails `json:"prompt_tokens_details"`

	// CompletionTokensDetails is the details of the completion tokens.
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`

	// TimingInfo contains detailed timing information for token generation.
	TimingInfo *TimingInfo `json:"timing_info,omitempty"`
}

// PromptTokensDetails is the details of the prompt tokens.
type PromptTokensDetails struct {
	// CachedTokens is the number of cached tokens in the prompt.
	CachedTokens int `json:"cached_tokens"`
	// CacheCreationTokens is the number of tokens used to create the cache (Anthropic).
	CacheCreationTokens int `json:"cache_creation_tokens,omitempty"`
	// CacheReadTokens is the number of tokens read from cache (Anthropic).
	CacheReadTokens int `json:"cache_read_tokens,omitempty"`
}

// CompletionTokensDetails is the details of the completion tokens.
// It intentionally exposes reasoning tokens first; other OpenAI SDK completion
// detail fields such as audio tokens and prediction token counters are deferred
// until the model package has provider-agnostic semantics for them.
type CompletionTokensDetails struct {
	// ReasoningTokens is the number of tokens generated for reasoning.
	ReasoningTokens int `json:"reasoning_tokens,omitempty"`
}

// Response is the response from the model.
//
// Error Handling Note:
// The Error field in this struct represents API-level errors that occur
// after successful communication with the model service. This is different
// from function-level errors returned by GenerateContent(), which indicate
// system-level failures that prevent communication entirely.
//
// Examples of Response.Error:
// - API rate limit exceeded
// - Content filtered by safety systems
// - Model-specific processing errors
// - Streaming connection errors
//
// Examples of function-level errors:
// - Invalid request parameters
// - Network connectivity issues
// - Authentication failures
type Response struct {
	// ID is the unique identifier for this response.
	ID string `json:"id"`

	// Object describes the type of object returned (e.g., "chat.completion").
	Object string `json:"object"`

	// Created is the Unix timestamp when the response was created.
	Created int64 `json:"created"`

	// Model is the model used to generate the response.
	Model string `json:"model"`

	// Choices contains the completion choices.
	Choices []Choice `json:"choices"`

	// Usage contains token usage information (may be nil for streaming responses).
	Usage *Usage `json:"usage,omitempty"`

	// SystemFingerprint is a unique identifier for the backend configuration.
	SystemFingerprint *string `json:"system_fingerprint,omitempty"`

	// Error contains API-level error information if the request failed.
	// This is nil for successful responses.
	// Note: This is different from function-level errors returned by GenerateContent().
	Error *ResponseError `json:"error,omitempty"`

	// Timestamp when this response chunk was received (for streaming).
	Timestamp time.Time `json:"timestamp"`

	// Done indicates if the llm flow should stop.
	Done bool `json:"done"`

	// IsPartial indicates if this is a partial response.
	IsPartial bool `json:"is_partial"`
}

// Clone creates a deep copy of the response.
func (rsp *Response) Clone() *Response { _ = "STUB: not implemented"; return nil }

// Deep copy TimingInfo if present

// Deep copy Error if present.

// Deep copy SystemFingerprint if present.

// IsValidContent checks if the response has valid content for message generation.
func (rsp *Response) IsValidContent() bool { _ = "STUB: not implemented"; return false }

// Check if event has choices with meaningful payload.

// IsUserMessage checks if the response is a user message.
func (rsp *Response) IsUserMessage() bool { _ = "STUB: not implemented"; return false }

// IsToolResultResponse  checks if the response is a tool call result response.
func (rsp *Response) IsToolResultResponse() bool { _ = "STUB: not implemented"; return false }

// IsToolCallResponse checks if the response is related to tool calls.
func (rsp *Response) IsToolCallResponse() bool { _ = "STUB: not implemented"; return false }

// GetToolCallIDs gets the IDs of tool calls from the response.
func (rsp *Response) GetToolCallIDs() []string { _ = "STUB: not implemented"; return nil }

// GetToolResultIDs gets the IDs of tool results from the response.
func (rsp *Response) GetToolResultIDs() []string { _ = "STUB: not implemented"; return nil }

// IsFinalResponse checks if the Response is a final response.
func (rsp *Response) IsFinalResponse() bool { _ = "STUB: not implemented"; return false }

// Consider response final if it's marked as done and has content or error.

// ResponseError represents an error response from the API.
type ResponseError struct {
	// Message is the error message.
	Message string `json:"message"`

	// Type is the type of error.
	Type string `json:"type"`

	// Param is the parameter that caused the error.
	Param *string `json:"param,omitempty"`

	// Code is the error code.
	Code *string `json:"code,omitempty"`
}

// Error implements the error interface.
func (e *ResponseError) Error() string { _ = "STUB: not implemented"; return "" }
