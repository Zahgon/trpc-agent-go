//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package anthropic provides Anthropic-compatible model implementations.
package anthropic

import (
	"context"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	functionToolType       = "function"
	claudeMythosPreview    = "claude-mythos-preview"
	claudeOpus47           = "claude-opus-4-7"
	claudeOpus46           = "claude-opus-4-6"
	claudeOpus46Alias      = "claude-4.6-opus"
	claudeSonnet46         = "claude-sonnet-4-6"
	claudeSonnet46Alias    = "claude-4.6-sonnet"
	defaultThinkingDisplay = anthropic.ThinkingConfigAdaptiveDisplaySummarized
)

// Model implements the model.Model interface for Anthropic API.
type Model struct {
	client                     anthropic.Client
	name                       string
	baseURL                    string
	apiKey                     string
	channelBufferSize          int
	anthropicRequestOptions    []option.RequestOption
	chatRequestCallback        ChatRequestCallbackFunc
	chatResponseCallback       ChatResponseCallbackFunc
	chatChunkCallback          ChatChunkCallbackFunc
	chatStreamCompleteCallback ChatStreamCompleteCallbackFunc
	enableTokenTailoring       bool                    // Enable automatic token tailoring.
	maxInputTokens             int                     // Max input tokens for token tailoring.
	contextWindow              int                     // Context window for this model instance.
	tokenCounter               model.TokenCounter      // Token counter for token tailoring.
	tailoringStrategy          model.TailoringStrategy // Tailoring strategy for token tailoring.
	// Token tailoring budget parameters (instance-level overrides).
	protocolOverheadTokens int
	reserveOutputTokens    int
	inputTokensFloor       int
	outputTokensFloor      int
	safetyMarginRatio      float64
	maxInputTokensRatio    float64
	// Prompt cache configuration
	cacheSystemPrompt bool
	cacheTools        bool
	cacheMessages     bool
	showToolCallDelta bool
}

// New creates a new Anthropic model adapter.
func New(name string, opts ...Option) *Model { _ = "STUB: not implemented"; return nil }

// Info returns the model information.
func (m *Model) Info() model.Info { _ = "STUB: not implemented"; return *new(model.Info) }

func (m *Model) runChatRequestCallback(
	ctx context.Context,
	chatRequest *anthropic.MessageNewParams,
) {
	_ = "STUB: not implemented"
	return
}

func (m *Model) runChatResponseCallback(
	ctx context.Context,
	chatRequest *anthropic.MessageNewParams,
	chatResponse *anthropic.Message,
) {
	_ = "STUB: not implemented"
	return
}

func (m *Model) runChatChunkCallback(
	ctx context.Context,
	chatRequest *anthropic.MessageNewParams,
	chatChunk *anthropic.MessageStreamEventUnion,
) {
	_ = "STUB: not implemented"
	return
}

func (m *Model) runChatStreamCompleteCallback(
	ctx context.Context,
	chatRequest *anthropic.MessageNewParams,
	chatResponse *anthropic.Message,
	streamErr error,
) {
	_ = "STUB: not implemented"
	return
}

// GenerateContent generates content from the model.
func (m *Model) GenerateContent(
	ctx context.Context,
	request *model.Request,
) (<-chan *model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply token tailoring if configured.

// Execute callback synchronously before starting the goroutine
// to avoid a race where the runner and HTTP handler finish
// (closing the SSE writer) while the callback is still running.

// Send chat request and handle response.

// applyTokenTailoring performs best-effort token tailoring if configured.
// It uses the token tailoring strategy defined in imodel package.
func (m *Model) applyTokenTailoring(ctx context.Context, request *model.Request) {
	_ = "STUB: not implemented"
	// Early return if token tailoring is disabled or no messages to process.
	return
}

// Determine max input tokens using priority: user config > auto calculation > default.

// Auto-calculate based on model context window with custom or default parameters.

// Use custom parameters if any are set.

// Use default parameters.

// Apply token tailoring.

// buildChatRequest builds the chat request for the Anthropic API.
func (m *Model) buildChatRequest(request *model.Request) (*anthropic.MessageNewParams, error) {
	_ = "STUB: not implemented"
	// Convert messages to Anthropic format.
	return nil, nil
}

// Convert tools

// Apply cache control breakpoints if any cache option is enabled.
// Uses multiple independent breakpoints (up to 4 allowed by Anthropic) for optimal caching:
// - System prompt breakpoint: caches stable system instructions
// - Tools breakpoint: caches stable tool definitions
// - Messages breakpoint: caches conversation history at last assistant message

// Build chat request.

// Only apply default MaxTokens when token tailoring is disabled.
// When token tailoring is enabled, respect the value set by applyTokenTailoring
// (or leave it as 0 if token counting failed).

func (m *Model) applyThinkingConfig(
	chatRequest *anthropic.MessageNewParams,
	request *model.Request,
) error {
	_ = "STUB: not implemented"
	return nil
}

func newAdaptiveThinkingConfig() anthropic.ThinkingConfigParamUnion {
	_ = "STUB: not implemented"
	return *new(anthropic.ThinkingConfigParamUnion)
}

func supportsAdaptiveThinking(modelName string) bool { _ = "STUB: not implemented"; return false }

func isClaudeMythosPreview(modelName string) bool { _ = "STUB: not implemented"; return false }

func modelNameMatches(modelName string, targets ...string) bool {
	_ = "STUB: not implemented"
	return false
}

// applyCacheControl applies independent cache control breakpoints.
// Unlike the previous single-breakpoint strategy, this sets multiple breakpoints
// independently (Anthropic supports up to 4). This ensures stable content like
// system prompts and tools always benefit from caching, regardless of whether
// message caching is also enabled.
//
// Breakpoints applied (each independent):
//   - System prompt: always cached when cacheSystemPrompt is true (stable across turns)
//   - Tools: always cached when cacheTools is true (stable across turns)
//   - Last assistant message: cached when cacheMessages is true (opt-in, benefits multi-turn)
func (m *Model) applyCacheControl(
	systemPrompts []anthropic.TextBlockParam,
	tools []anthropic.ToolUnionParam,
	messages []anthropic.MessageParam,
) ([]anthropic.TextBlockParam, []anthropic.ToolUnionParam, []anthropic.MessageParam) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// findLastAssistantMessageIndex finds the index of the last assistant message
// that is not the final message (we want to cache up to but not including the current turn).
func (m *Model) findLastAssistantMessageIndex(messages []anthropic.MessageParam) int {
	_ = "STUB: not implemented"
	// We look for the last assistant message before the final user message
	// In a typical conversation: [user, assistant, user, assistant, user]
	// We want to cache at the second-to-last assistant (index 3)
	return 0
}

// applyCacheControlToMessages adds cache control to a specific message.
// This is used for multi-turn conversation caching.
func (m *Model) applyCacheControlToMessages(messages []anthropic.MessageParam, index int) []anthropic.MessageParam {
	_ = "STUB: not implemented"
	return nil
}

// Create a copy to avoid modifying the original

// Add cache control to the last cacheable content block of the target message.
// We iterate backwards to find the first block that supports cache control,
// in case the last block is an image or other unsupported media type.

// Apply cache control based on content type

// Stop after applying to the first valid block from the end

// applyCacheControlToSystem adds cache control to the last system prompt block.
func (m *Model) applyCacheControlToSystem(systemPrompts []anthropic.TextBlockParam) []anthropic.TextBlockParam {
	_ = "STUB: not implemented"
	return nil
}

// Create a copy to avoid modifying the original slice elements.

// applyCacheControlToTools adds cache control to the last tool definition.
func (m *Model) applyCacheControlToTools(tools []anthropic.ToolUnionParam) []anthropic.ToolUnionParam {
	_ = "STUB: not implemented"
	return nil
}

// Create a copy to avoid modifying the original slice elements.

// handleNonStreamingResponse sends a non-streaming request to the Anthropic API and emits exactly one final response.
func (m *Model) handleNonStreamingResponse(
	ctx context.Context,
	chatRequest anthropic.MessageNewParams,
	responseChan chan<- *model.Response,
) {
	_ = "STUB: not implemented"
	// Issue non-streaming request.
	return
}

// Build final response payload.

// Convert assistant content blocks.

// Set finish reason.

// Set usage.

// Emit final response.

// handleStreamingResponse sends a streaming request to the Anthropic API and emits partial deltas
// followed by a final response.
func (m *Model) handleStreamingResponse(
	ctx context.Context,
	chatRequest anthropic.MessageNewParams,
	responseChan chan<- *model.Response,
) {
	_ = "STUB: not implemented"
	// Issue streaming request.
	return
}

// Accumulator to build final response.

// Accumulate into accumulator.

// Build partial response.

// Emit partial response.

// Propagate stream error.

type streamingMessageAccumulator struct {
	message             anthropic.Message
	inputDeltaStartedAt []bool
	finalized           bool
}

func newStreamingMessageAccumulator() *streamingMessageAccumulator {
	_ = "STUB: not implemented"
	return nil
}

func (a *streamingMessageAccumulator) Message() anthropic.Message {
	_ = "STUB: not implemented"
	return *new(anthropic.Message)
}

func (a *streamingMessageAccumulator) Accumulate(event anthropic.MessageStreamEventUnion) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *streamingMessageAccumulator) Finalize() error { _ = "STUB: not implemented"; return nil }

func (a *streamingMessageAccumulator) lastContentBlock() (*anthropic.ContentBlockUnion, *bool, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (a *streamingMessageAccumulator) finalize() error { _ = "STUB: not implemented"; return nil }

func appendInputJSONDelta(block *anthropic.ContentBlockUnion, started *bool, partial string) {
	_ = "STUB: not implemented"
	return
}

// Treat the first streamed input delta as authoritative to avoid
// concatenating a complete start-event value with a second JSON value.

func finalizeStreamingMessage(message *anthropic.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func refreshContentBlockRawJSON(block *anthropic.ContentBlockUnion) error {
	_ = "STUB: not implemented"
	return nil
}

// buildStreamingPartialResponse builds a partial streaming response for a chunk.
// Returns nil if the chunk should be skipped.
func buildStreamingPartialResponse(acc anthropic.Message,
	chunk anthropic.MessageStreamEventUnion,
	showToolCallDelta bool) (*model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Branch by event type.

func buildStreamingToolCallDelta(
	block anthropic.ToolUseBlock,
	index int,
	partialJSON string,
) model.ToolCall {
	_ = "STUB: not implemented"
	return *new(model.ToolCall)
}

func streamingToolCallTarget(acc anthropic.Message, contentBlockIndex int64) (anthropic.ToolUseBlock, int, bool) {
	_ = "STUB: not implemented"
	return *new(anthropic.ToolUseBlock), 0, false
}

// buildStreamingFinalResponse builds a final streaming response from the accumulator.
func buildStreamingFinalResponse(acc anthropic.Message) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

// Aggregate all blocks into final assistant message.

// Build final response.

// sendErrorResponse sends an error response through the channel.
func (m *Model) sendErrorResponse(ctx context.Context, responseChan chan<- *model.Response, errType string, err error) {
	_ = "STUB: not implemented"
	return
}

// convertContentBlock builds a single assistant message from Anthropic content blocks.
func convertContentBlock(contents []anthropic.ContentBlockUnion) model.Message {
	_ = "STUB: not implemented"
	return *new(model.Message)
}

// convertTools maps our tool declarations to Anthropic tool parameters.
func convertTools(tools map[string]tool.Tool) []anthropic.ToolUnionParam {
	_ = "STUB: not implemented"
	return nil
}

// buildToolDescription builds the description for a tool.
// It appends the output schema to the description.
func buildToolDescription(declaration *tool.Declaration) string {
	_ = "STUB: not implemented"
	return ""
}

// convertMessages builds Anthropic message parameters and system prompts from trpc-agent-go messages.
// Merges consecutive tool results into a single user message and drops empty-content messages.
func convertMessages(messages []model.Message) ([]anthropic.MessageParam, []anthropic.TextBlockParam, error) {
	_ = "STUB: not implemented"
	// Convert messages by role and collect system prompts.
	return nil, nil, nil
}

// Merge consecutive tool result messages into a single user message to support parallel tool invocation.

// Skip empty content messages.

// Forward non-tool result messages.

// Gather contiguous tool results and wrap into a single user message to support parallel tool invocation.

// convertUserMessage converts a user message including supported multimodal parts.
func convertUserMessage(message model.Message) (anthropic.MessageParam, error) {
	_ = "STUB: not implemented"
	return *new(anthropic.MessageParam), nil
}

func convertUserContentPart(part model.ContentPart) (anthropic.ContentBlockParamUnion, bool, error) {
	_ = "STUB: not implemented"
	return *new(anthropic.ContentBlockParamUnion), false, nil
}

func convertImageContentPart(image *model.Image) (anthropic.ContentBlockParamUnion, error) {
	_ = "STUB: not implemented"
	return *new(anthropic.ContentBlockParamUnion), nil
}

func imageURLMediaTypeUnsupported(image *model.Image) bool { _ = "STUB: not implemented"; return false }

func imageURLText(image *model.Image) string { _ = "STUB: not implemented"; return "" }

func convertFileContentPart(file *model.File) (anthropic.ContentBlockParamUnion, error) {
	_ = "STUB: not implemented"
	return *new(anthropic.ContentBlockParamUnion), nil
}

func isTextLikeFileMediaType(mediaType string) bool { _ = "STUB: not implemented"; return false }

func fileDataText(file *model.File, mediaType string) string { _ = "STUB: not implemented"; return "" }

func applyDocumentTitle(block anthropic.ContentBlockParamUnion, name string) anthropic.ContentBlockParamUnion {
	_ = "STUB: not implemented"
	return *new(anthropic.ContentBlockParamUnion)
}

func imageMediaType(format string, data []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func normalizeImageMediaType(value string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func fileMediaType(file *model.File) string { _ = "STUB: not implemented"; return "" }

func fileURLMediaType(file *model.File) string { _ = "STUB: not implemented"; return "" }

func mediaTypeFromURL(rawURL string) string { _ = "STUB: not implemented"; return "" }

func mediaTypeFromName(name string) string { _ = "STUB: not implemented"; return "" }

func normalizeMediaType(value string) string { _ = "STUB: not implemented"; return "" }

// convertAssistantMessageContent converts an assistant message including tool calls into Anthropic format.
func convertAssistantMessageContent(message model.Message) (anthropic.MessageParam, error) {
	_ = "STUB: not implemented"
	// Append text blocks.
	return *new(anthropic.MessageParam), nil
}

// Append tool use blocks.

// decodeToolArguments parses JSON bytes into any, returning an empty object on failure.
func decodeToolArguments(args []byte) any { _ = "STUB: not implemented"; return *new(any) }

// convertToolResult wraps a tool result into a user message with a ToolResult block.
func convertToolResult(message model.Message) anthropic.MessageParam {
	_ = "STUB: not implemented"
	return *new(anthropic.MessageParam)
}

// convertSystemMessageContent converts message content to system message content union.
func convertSystemMessageContent(message model.Message) ([]anthropic.TextBlockParam, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
