//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package openai provides OpenAI-compatible model implementations.
package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	openai "github.com/openai/openai-go"
	openaiopt "github.com/openai/openai-go/option"
	"github.com/openai/openai-go/packages/respjson"
	"github.com/openai/openai-go/packages/ssestream"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	functionToolType string = "function"

	// defaultBatchEndpoint is the default batch endpoint.
	defaultBatchEndpoint = openai.BatchNewParamsEndpointV1ChatCompletions
	//nolint:gosec
	deepSeekAPIKeyName     string = "DEEPSEEK_API_KEY"
	defaultDeepSeekBaseURL string = "https://api.deepseek.com"
	deepSeekAPIHost        string = "api.deepseek.com"

	//nolint:gosec
	qwenAPIKeyName     string = "DASHSCOPE_API_KEY"
	defaultQwenBaseURL string = "https://dashscope.aliyuncs.com/compatible-mode/v1"
)

// Variant represents different model variants with specific behaviors.
type Variant string

const (
	// VariantOpenAI is the default OpenAI variant.
	VariantOpenAI Variant = "openai"
	// VariantHunyuan is the Hunyuan variant with specific file handling.
	VariantHunyuan Variant = "hunyuan"
	// VariantDeepSeek is the DeepSeek variant with specific base_url handling.
	// It backfills empty reasoning_content for assistant history messages by
	// default, including when inferred from the default DeepSeek base URL. Use
	// WithReasoningContentBackfill(false) to disable this behavior.
	VariantDeepSeek Variant = "deepseek"
	// VariantQwen is the Qwen variant with specific base_url handling.
	VariantQwen Variant = "qwen"
)

// thinkingValueConvertor converts ThinkingEnabled bool to the variant-specific value.
type thinkingValueConvertor func(enabled bool) any

// defaultThinkingValueConvertor returns the bool value as-is.
var defaultThinkingValueConvertor = func(enabled bool) any {
	return enabled
}

// deepSeekThinkingValueConvertor converts to the DeepSeek thinking-toggle
// format introduced in v3.2 and reused by v4 (e.g. deepseek-v4-pro /
// deepseek-v4-flash): {"type": "enabled"/"disabled"}.
var deepSeekThinkingValueConvertor = func(enabled bool) any {
	const (
		thinkingTypeEnabled  = "enabled"
		thinkingTypeDisabled = "disabled"
	)
	thinkingType := thinkingTypeDisabled
	if enabled {
		thinkingType = thinkingTypeEnabled
	}
	return map[string]string{"type": thinkingType}
}

// variantConfig holds configuration for different variants.
type variantConfig struct {
	// Default file upload path for this variant.
	fileUploadPath   string
	fileDeletionPath string
	// Default file purpose for this variant.
	filePurpose openai.FilePurpose
	// Default HTTP method for file deletion.
	fileDeletionMethod         string
	fileDeletionBodyConvertor  fileDeletionBodyConvertor
	fileUploadRequestConvertor fileUploadRequestConvertor
	// Whether to skip file type in content parts for this variant.
	skipFileTypeInContent bool
	// Whether user message content must be reduced to text only.
	textOnlyMessageContent bool

	// Default base URL for this variant.
	defaultBaseURL string
	// Default API key name for this variant.
	apiKeyName string
	// Thinking key for this variant.
	thinkingEnabledKey string
	// thinkingValueConvertor converts ThinkingEnabled to variant-specific format.
	thinkingValueConvertor thinkingValueConvertor

	// defaultOptimizeForCache controls the default value for cache optimization
	// when WithOptimizeForCache is not explicitly set.
	defaultOptimizeForCache bool
	// defaultReasoningContentBackfill controls replay-time empty
	// reasoning_content backfill for assistant messages.
	defaultReasoningContentBackfill bool
}
type fileDeletionBodyConvertor func(body []byte, fileID string) []byte

// defaultFileDeletionBodyConvertor is the default file deletion body converter.
var defaultFileDeletionBodyConvertor = func(body []byte, fileID string) []byte {
	return body
}

type fileUploadRequestConvertor func(r *http.Request, file *os.File, fileOpts *FileOptions) (*http.Request, error)

// variantConfigs maps variant names to their configurations.
var variantConfigs = map[Variant]variantConfig{
	VariantOpenAI: {
		fileUploadPath:            "/openapi/v1/files",
		filePurpose:               openai.FilePurposeUserData,
		fileDeletionMethod:        http.MethodDelete,
		skipFileTypeInContent:     false,
		fileDeletionBodyConvertor: defaultFileDeletionBodyConvertor,
		thinkingEnabledKey:        model.ThinkingEnabledKey,
		thinkingValueConvertor:    defaultThinkingValueConvertor,
		defaultOptimizeForCache:   true,
	},
	VariantDeepSeek: {
		fileUploadPath:            "/openapi/v1/files",
		filePurpose:               openai.FilePurposeUserData,
		fileDeletionMethod:        http.MethodDelete,
		skipFileTypeInContent:     false,
		textOnlyMessageContent:    true,
		fileDeletionBodyConvertor: defaultFileDeletionBodyConvertor,
		apiKeyName:                deepSeekAPIKeyName,
		defaultBaseURL:            defaultDeepSeekBaseURL,
		// DeepSeek v3.2+ (incl. v4-pro / v4-flash) uses
		// {"thinking": {"type": "enabled"/"disabled"}} format.
		thinkingEnabledKey:              "thinking",
		thinkingValueConvertor:          deepSeekThinkingValueConvertor,
		defaultReasoningContentBackfill: true,
	},
	VariantHunyuan: {
		fileUploadPath:        "/openapi/v1/files/uploads",
		fileDeletionPath:      "/openapi/v1/files",
		filePurpose:           openai.FilePurpose("file-extract"),
		fileDeletionMethod:    http.MethodPost,
		skipFileTypeInContent: true,
		fileDeletionBodyConvertor: func(body []byte, fileID string) []byte {
			if body != nil {
				return body
			}
			return []byte(`{"file_id":"` + fileID + `"}`)
		},
		fileUploadRequestConvertor: func(r *http.Request, file *os.File, fileOpts *FileOptions) (*http.Request, error) {
			// Create multipart form data.
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			// Add purpose field.
			if err := writer.WriteField("purpose", string(fileOpts.Purpose)); err != nil {
				return nil, fmt.Errorf("failed to write purpose field: %w", err)
			}
			// Add file field.
			fileInfo, err := file.Stat()
			if err != nil {
				return nil, fmt.Errorf("failed to get file info: %w", err)
			}
			part, err := writer.CreateFormFile("file", fileInfo.Name())
			if err != nil {
				return nil, fmt.Errorf("failed to create form file: %w", err)
			}
			// Reset file position and copy file content.
			if _, err := file.Seek(0, 0); err != nil {
				return nil, fmt.Errorf("failed to reset file position: %w", err)
			}
			if _, err := io.Copy(part, file); err != nil {
				return nil, fmt.Errorf("failed to copy file content: %w", err)
			}
			// Close the writer to finalize the multipart data.
			if err := writer.Close(); err != nil {
				return nil, fmt.Errorf("failed to close multipart writer: %w", err)
			}
			// Set the request body and content type.
			r.Body = io.NopCloser(body)
			r.Header.Set("Content-Type", writer.FormDataContentType())
			r.ContentLength = int64(body.Len())
			return r, nil
		},
		thinkingEnabledKey:     model.ThinkingEnabledKey,
		thinkingValueConvertor: defaultThinkingValueConvertor,
	},
	VariantQwen: {
		fileUploadPath:            "/openapi/v1/files",
		filePurpose:               openai.FilePurposeUserData,
		fileDeletionMethod:        http.MethodDelete,
		skipFileTypeInContent:     false,
		fileDeletionBodyConvertor: defaultFileDeletionBodyConvertor,
		apiKeyName:                qwenAPIKeyName,
		defaultBaseURL:            defaultQwenBaseURL,
		// refer:https://help.aliyun.com/zh/model-studio/deep-thinking
		thinkingEnabledKey:     model.EnabledThinkingKey,
		thinkingValueConvertor: defaultThinkingValueConvertor,
	},
}

// Model implements the model.Model interface for OpenAI API.
type Model struct {
	client                     openai.Client
	name                       string
	baseURL                    string
	apiKey                     string
	showToolCallDelta          bool
	channelBufferSize          int
	chatRequestCallback        ChatRequestCallbackFunc
	chatRequestJSONCallback    ChatRequestJSONCallbackFunc
	chatResponseCallback       ChatResponseCallbackFunc
	chatChunkCallback          ChatChunkCallbackFunc
	chatStreamCompleteCallback ChatStreamCompleteCallbackFunc
	chatTelemetry              bool
	extraFields                map[string]any
	variant                    Variant
	variantConfig              variantConfig
	reasoningContentBackfill   bool
	batchCompletionWindow      openai.BatchNewParamsCompletionWindow
	batchMetadata              map[string]string
	batchBaseURL               string
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

	accumulateChunkUsage AccumulateChunkUsage
	optimizeForCache     bool // Optimize message structure for prompt caching
	omitFileContentParts bool
}

// New creates a new OpenAI-like model.
func New(name string, opts ...Option) *Model { _ = "STUB: not implemented"; return nil }

// Set default API key and base URL if not specified.

func inferVariant(baseURL string) Variant { _ = "STUB: not implemented"; return *new(Variant) }

func isDeepSeekBaseURL(raw string) bool { _ = "STUB: not implemented"; return false }

// Info implements the model.Model interface.
func (m *Model) Info() model.Info { _ = "STUB: not implemented"; return *new(model.Info) }

func (m *Model) runChatRequestCallback(
	ctx context.Context,
	chatRequest *openai.ChatCompletionNewParams,
) {
	_ = "STUB: not implemented"
	return
}

func (m *Model) runChatRequestJSONCallback(
	ctx context.Context,
	chatRequest *openai.ChatCompletionNewParams,
) {
	_ = "STUB: not implemented"
	return
}

func (m *Model) runChatResponseCallback(
	ctx context.Context,
	chatRequest *openai.ChatCompletionNewParams,
	chatResponse *openai.ChatCompletion,
) {
	_ = "STUB: not implemented"
	return
}

func (m *Model) runChatChunkCallback(
	ctx context.Context,
	chatRequest *openai.ChatCompletionNewParams,
	chatChunk *openai.ChatCompletionChunk,
) {
	_ = "STUB: not implemented"
	return
}

// prepareChatRequest validates and mutates the request in-place before sending it to the provider.
func (m *Model) prepareChatRequest(
	ctx context.Context,
	request *model.Request,
) (*openai.ChatCompletionNewParams, []openaiopt.RequestOption, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Optimize message structure for cache if enabled.

// Apply token tailoring if configured.

// GenerateContent implements the model.Model interface.
func (m *Model) GenerateContent(
	ctx context.Context,
	request *model.Request,
) (<-chan *model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Execute callback synchronously before starting the goroutine
// to avoid a race where the runner and HTTP handler finish
// (closing the SSE writer) while the callback is still running.

// GenerateContentIter implements the model.IterModel interface.
func (m *Model) GenerateContentIter(
	ctx context.Context,
	request *model.Request,
) (model.Seq[*model.Response], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// optimizeMessagesForCache reorders messages to improve cache hit rates.
// System messages are moved to the front as they are most likely to be cached.
func (m *Model) optimizeMessagesForCache(messages []model.Message) []model.Message {
	_ = "STUB: not implemented"
	return nil
}

// If no reordering needed, return original

// System messages first, then other messages

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

func (m *Model) effectiveOutputReserveTokens(request *model.Request) int {
	_ = "STUB: not implemented"
	return 0
}

func (m *Model) hardInputBudget(contextWindow, outputReserveTokens int) int {
	_ = "STUB: not implemented"
	return 0
}

func (m *Model) estimateToolsTokens(
	ctx context.Context,
	tools map[string]tool.Tool,
) int {
	_ = "STUB: not implemented"
	return 0
}

// buildChatRequest converts our Request to OpenAI request params and options.
func (m *Model) buildChatRequest(request *model.Request) (*openai.ChatCompletionNewParams, []openaiopt.RequestOption) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set response_format for native structured outputs when requested.

// Parallel tool calls can interfere with strict JSON schema
// output.

// MaxTokens is deprecated and not compatible with o-series models.
// Use MaxCompletionTokens instead.

// Use the first stop string for simplicity.

// Add model-level extra fields to the request.

// Add request-level extra fields after model-level fields so they take precedence.

// Add streaming options if needed.

// buildThinkingOption converts our Request to OpenAI request RequestOption.
//
// Note on default behavior: when request.ThinkingEnabled is nil, this function
// does not emit any thinking-toggle field in the outgoing request. The
// upstream provider then applies its server-side default (e.g. DeepSeek v4
// defaults to thinking "enabled"). Callers that want a deterministic on/off
// behavior must set ThinkingEnabled explicitly.
func (m *Model) buildThinkingOption(request *model.Request) []openaiopt.RequestOption {
	_ = "STUB: not implemented"
	return nil
}

// Use variant-specific key and value convertor.

// shouldBackfillReasoningContent reports whether replay should emit
// model.ReasoningContentKey as an empty string for assistant messages. Some
// providers require the key to be present for every assistant message in
// thinking mode even when no reasoning text was returned.
func (m *Model) shouldBackfillReasoningContent(
	msg model.Message,
) bool {
	_ = "STUB: not implemented"
	return false
}

// convertMessages converts our Message format to OpenAI's format.
func (m *Model) convertMessages(messages []model.Message) []openai.ChatCompletionMessageParamUnion {
	_ = "STUB: not implemented"
	return nil
}

// Pass reasoning_content to API if present, or when provider replay
// requires the field for assistant history in thinking mode.

// Default to user message if role is unknown.

// convertSystemMessageContent converts message content to system message content union.
func (m *Model) convertSystemMessageContent(msg model.Message) openai.ChatCompletionSystemMessageParamContentUnion {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionSystemMessageParamContentUnion)
}

// Convert content parts to OpenAI content parts.

// convertUserMessageContent converts a message into an OpenAI user
// message content union.
func (m *Model) convertUserMessageContent(
	msg model.Message,
) (openai.ChatCompletionUserMessageParamContentUnion, map[string]any) {
	_ = "STUB: not implemented"
	// If there are no content parts and Content is not empty, return as string.
	return *new(openai.ChatCompletionUserMessageParamContentUnion), nil
}

func (m *Model) userFileHint(msg model.Message) string { _ = "STUB: not implemented"; return "" }

func onlyFileContentParts(parts []model.ContentPart) bool { _ = "STUB: not implemented"; return false }

func onlyInternalFileContentParts(parts []model.ContentPart) bool {
	_ = "STUB: not implemented"
	return false
}

func fileURLFallbackText(file *model.File) string { _ = "STUB: not implemented"; return "" }

func userTextPart(text string) openai.ChatCompletionContentPartUnionParam {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionContentPartUnionParam)
}

func appendFileURLFallbackText(
	dst *[]openai.ChatCompletionContentPartUnionParam,
	file *model.File,
) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *Model) appendUserContentParts(
	dst *[]openai.ChatCompletionContentPartUnionParam,
	parts []model.ContentPart,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (m *Model) omittedContentHint(parts []model.ContentPart) string {
	_ = "STUB: not implemented"
	return ""
}

func omittedAttachmentHint(
	imageCount int,
	audioCount int,
	fileCount int,
) string {
	_ = "STUB: not implemented"
	return ""
}

func appendFileID(
	extraFields map[string]any,
	part model.ContentPart,
) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func singleUserContentString(
	mainText string,
	fileHint string,
	contentParts []openai.ChatCompletionContentPartUnionParam,
	extraFields map[string]any,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func fileHintForContentParts(parts []model.ContentPart) string {
	_ = "STUB: not implemented"
	return ""
}

// convertAssistantMessageContent converts message content to assistant message content union.
func (m *Model) convertAssistantMessageContent(
	msg model.Message,
) openai.ChatCompletionAssistantMessageParamContentUnion {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionAssistantMessageParamContentUnion)
}

// Convert content parts to OpenAI content parts.

// convertContentPart converts a single content part to OpenAI format.
func (m *Model) convertContentPart(part model.ContentPart) *openai.ChatCompletionContentPartUnionParam {
	_ = "STUB: not implemented"
	return nil
}

// The URL from openai-go can be used either as a URL or as a base64-encoded string.

func imageToURLOrBase64(image *model.Image) string { _ = "STUB: not implemented"; return "" }

func isProviderFileID(fileID string) bool { _ = "STUB: not implemented"; return false }

func isInternalOnlyFile(file *model.File) bool { _ = "STUB: not implemented"; return false }

func safeFileHintName(file *model.File) string { _ = "STUB: not implemented"; return "" }

func fileToParamsOK(
	file *model.File,
) (openai.ChatCompletionContentPartFileFileParam, bool) {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionContentPartFileFileParam), false
}

func fileToParams(file *model.File) openai.ChatCompletionContentPartFileFileParam {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionContentPartFileFileParam)
}

func audioToBase64(audio *model.Audio) string { _ = "STUB: not implemented"; return "" }

func (m *Model) convertToolCalls(toolCalls []model.ToolCall) []openai.ChatCompletionMessageToolCallParam {
	_ = "STUB: not implemented"
	return nil
}

// Pass through ExtraFields transparently (e.g., Gemini 3's thought_signature).

func (m *Model) convertTools(tools map[string]tool.Tool) []openai.ChatCompletionToolParam {
	_ = "STUB: not implemented"
	return nil
}

// Convert the InputSchema to JSON to correctly map to OpenAI's expected format

// Some OpenAI-compatible proxies require object schemas to include
// a `properties` key, even when the tool takes no arguments.

// buildToolDescription builds the description for a tool.
// It appends the output schema to the description.
func buildToolDescription(declaration *tool.Declaration) string {
	_ = "STUB: not implemented"
	return ""
}

// responseEmitter emits a response and returns false to stop streaming.
type responseEmitter func(*model.Response) bool

// handleStreamingResponseWithEmitter handles streaming chat completion responses.
// It returns early when emit returns false.
func (m *Model) handleStreamingResponseWithEmitter(
	ctx context.Context,
	chatRequest openai.ChatCompletionNewParams,
	emit responseEmitter,
	opts ...openaiopt.RequestOption,
) {
	_ = "STUB: not implemented"
	return
}

// Track ID -> Index mapping.

// Track ExtraFields by tool call ID (SDK accumulator doesn't preserve ExtraFields).

// Aggregate reasoning deltas for final message fallback (some providers don't retain it in accumulator).

// Track next available index for tool calls (for providers that don't set correct indices).

// Skip empty chunks.

// Fix tool call indices for providers that return all indices as 0.
// This must be done before updateToolCallIndexMapping and accumulation.

// Collect ExtraFields from chunk tool_calls (SDK accumulator doesn't preserve ExtraFields).

// Track ID -> Index mapping when ID is present (first chunk of each tool call).

// Accumulate chunk for correctness. When a chunk mixes reasoning with
// content or tool-call deltas, strip only the reasoning metadata before
// passing it to the SDK accumulator.

// Suppress chunks that carry no meaningful visible delta (including
// tool_call deltas, which we'll surface only in the final response).
// Note: reasoning content chunks are not suppressed even if they have no other content.

// Call the stream complete callback before the final response is emitted.

// sanitizeChunkForAccumulator returns a defensive copy of the given chunk that
// avoids structures known to cause panics in the upstream OpenAI SDK
// accumulator. In particular, it clears JSON.ToolCalls metadata when it is
// marked present but the typed ToolCalls slice is empty on a finish_reason
// chunk, which would otherwise lead to an out-of-range access in
// chatCompletionResponseState.update.
func sanitizeChunkForAccumulator(chunk openai.ChatCompletionChunk) openai.ChatCompletionChunk {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionChunk)
}

// Only sanitize the specific pattern that is known to be unsafe for the
// accumulator:
//   - finish_reason is set (e.g. "tool_calls" or "stop")
//   - JSON.ToolCalls is marked present
//   - but the typed ToolCalls slice is empty

// Clear the JSON metadata for ToolCalls on the first choice only. This
// preserves finish_reason and usage semantics while preventing the
// accumulator from treating this as a tool-call delta that must have at
// least one element.

// stripReasoningFromChunkForAccumulator returns a defensive copy of the chunk
// with reasoning-only ExtraFields removed from the first choice delta. This
// lets the upstream accumulator keep non-reasoning payloads from mixed chunks
// without mutating the original chunk.
func stripReasoningFromChunkForAccumulator(
	chunk openai.ChatCompletionChunk,
) (openai.ChatCompletionChunk, bool) {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionChunk), false
}

// hasAccumulatorPayloadBeyondReasoning reports whether the stripped chunk still
// contains payload the SDK accumulator must keep, such as content, refusal,
// tool calls, finish reasons, or usage. Pure reasoning-only chunks
// intentionally keep the old behavior and stay out of the accumulator.
func hasAccumulatorPayloadBeyondReasoning(
	chunk openai.ChatCompletionChunk,
) bool {
	_ = "STUB: not implemented"
	return false
}

type toolCallIndexState struct {
	idToIndexMap map[string]int
	indexToID    map[int64]string
	nextIndex    *int
}

func buildIndexToIDMap(
	idToIndexMap map[string]int,
	toolCalls []openai.ChatCompletionChunkChoiceDeltaToolCall,
) map[int64]string {
	_ = "STUB: not implemented"
	return nil
}

func checkIfIndexFixNeeded(
	toolCalls []openai.ChatCompletionChunkChoiceDeltaToolCall,
	idToIndexMap map[string]int,
	indexToID map[int64]string,
) bool {
	_ = "STUB: not implemented"
	return false
}

func updateIDToIndexMapFromToolCalls(
	toolCalls []openai.ChatCompletionChunkChoiceDeltaToolCall,
	idToIndexMap map[string]int,
	nextIndex *int,
) {
	_ = "STUB: not implemented"
	return
}

func createDeepCopyOfChunkForFix(
	chunk openai.ChatCompletionChunk,
	delta openai.ChatCompletionChunkChoiceDelta,
) openai.ChatCompletionChunk {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionChunk)
}

func buildUsedIndicesSet(idToIndexMap map[string]int) map[int64]struct{} {
	_ = "STUB: not implemented"
	return nil
}

func findNextAvailableIndex(usedIndices map[int64]struct{}, startFrom int) int64 {
	_ = "STUB: not implemented"
	return 0
}

func applyToolCallIndexFixes(
	fixedChunk *openai.ChatCompletionChunk,
	state *toolCallIndexState,
	usedIndices map[int64]struct{},
) {
	_ = "STUB: not implemented"
	return
}

// fixToolCallIndices normalizes tool call indices in streaming chunks.
// Some providers incorrectly set ToolCalls[].Index to 0 for every tool call.
// The upstream openai-go accumulator uses ToolCalls[].Index as the slice position.
// When indices are wrong, different tool calls get merged by concatenating Name and Arguments.
// This function uses ToolCalls[].ID as the stable identity and rewrites indices to be consistent.
// This function also handles the case where a single chunk contains multiple tool calls sharing the same index.
// The idToIndexMap stores the canonical index for each tool call ID.
// The nextIndex points to the next available canonical index and is advanced monotonically.
func fixToolCallIndices(
	chunk openai.ChatCompletionChunk,
	idToIndexMap map[string]int,
	nextIndex *int,
) openai.ChatCompletionChunk {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionChunk)
}

// updateToolCallIndexMapping updates the tool call index mapping.
func (m *Model) updateToolCallIndexMapping(chunk openai.ChatCompletionChunk, idToIndexMap map[string]int) {
	_ = "STUB: not implemented"
	return
}

// collectExtraFieldsFromChunk collects ExtraFields from chunk tool_calls.
func (m *Model) collectExtraFieldsFromChunk(
	chunk openai.ChatCompletionChunk,
	extraFieldsMap map[string]map[string]any,
) {
	_ = "STUB: not implemented"
	return
}

// Use ID if available, otherwise use index as key.

func applyOpenAISDKTokenDetailsAccumulationFix(
	acc *openai.ChatCompletionAccumulator,
	chunk openai.ChatCompletionChunk,
) {
	_ = "STUB: not implemented"
	// Temporary workaround for token details accumulation.
	// See https://github.com/trpc-group/trpc-agent-go/issues/1270.
	// Remove this after upgrading openai-go to v3.10.0.
	return
}

// accumulateChunk accumulates non-reasoning deltas into the SDK accumulator and
// always appends reasoning deltas to the reasoning buffer.
func (m *Model) accumulateChunk(
	chunk openai.ChatCompletionChunk,
	acc *openai.ChatCompletionAccumulator,
	reasoningBuf *bytes.Buffer,
) {
	_ = "STUB: not implemented"
	return
}

// Sanitize chunks before feeding them into the upstream accumulator to
// avoid known panics when JSON.ToolCalls is marked present but the
// typed ToolCalls slice is empty, especially on finish_reason chunks.

// Aggregate reasoning delta (if any) for final response fallback.

// sendPartialResponse creates and sends a partial response from a chunk.
func (m *Model) sendPartialResponse(
	ctx context.Context,
	chunk openai.ChatCompletionChunk,
	responseChan chan<- *model.Response,
) error {
	_ = "STUB: not implemented"
	return nil
}

// handleStreamCompleteCallback handles the stream complete callback.
func (m *Model) handleStreamCompleteCallback(
	ctx context.Context,
	chatRequest openai.ChatCompletionNewParams,
	acc openai.ChatCompletionAccumulator,
	streamErr error,
) {
	_ = "STUB: not implemented"
	return
}

func cloneChatCompletionAccumulator(acc openai.ChatCompletionAccumulator) openai.ChatCompletionAccumulator {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionAccumulator)
}

func cloneChatCompletion(chatCompletion openai.ChatCompletion) openai.ChatCompletion {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletion)
}

func cloneChatCompletionChoice(choice openai.ChatCompletionChoice) openai.ChatCompletionChoice {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionChoice)
}

func cloneChatCompletionChoiceLogprobs(
	logprobs openai.ChatCompletionChoiceLogprobs,
) openai.ChatCompletionChoiceLogprobs {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionChoiceLogprobs)
}

func cloneChatCompletionTokenLogprob(
	token openai.ChatCompletionTokenLogprob,
) openai.ChatCompletionTokenLogprob {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionTokenLogprob)
}

func cloneChatCompletionTokenLogprobTopLogprob(
	token openai.ChatCompletionTokenLogprobTopLogprob,
) openai.ChatCompletionTokenLogprobTopLogprob {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionTokenLogprobTopLogprob)
}

func cloneChatCompletionMessage(message openai.ChatCompletionMessage) openai.ChatCompletionMessage {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionMessage)
}

func cloneChatCompletionMessageAnnotation(
	annotation openai.ChatCompletionMessageAnnotation,
) openai.ChatCompletionMessageAnnotation {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionMessageAnnotation)
}

func cloneChatCompletionMessageAnnotationURLCitation(
	urlCitation openai.ChatCompletionMessageAnnotationURLCitation,
) openai.ChatCompletionMessageAnnotationURLCitation {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionMessageAnnotationURLCitation)
}

func cloneChatCompletionAudio(audio openai.ChatCompletionAudio) openai.ChatCompletionAudio {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionAudio)
}

func cloneChatCompletionMessageFunctionCall(
	functionCall openai.ChatCompletionMessageFunctionCall,
) openai.ChatCompletionMessageFunctionCall {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionMessageFunctionCall)
}

func cloneChatCompletionMessageToolCall(
	toolCall openai.ChatCompletionMessageToolCall,
) openai.ChatCompletionMessageToolCall {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionMessageToolCall)
}

func cloneChatCompletionMessageToolCallFunction(
	function openai.ChatCompletionMessageToolCallFunction,
) openai.ChatCompletionMessageToolCallFunction {
	_ = "STUB: not implemented"
	return *new(openai.ChatCompletionMessageToolCallFunction)
}

func cloneCompletionUsage(usage openai.CompletionUsage) openai.CompletionUsage {
	_ = "STUB: not implemented"
	return *new(openai.CompletionUsage)
}

func cloneCompletionUsageCompletionTokensDetails(
	details openai.CompletionUsageCompletionTokensDetails,
) openai.CompletionUsageCompletionTokensDetails {
	_ = "STUB: not implemented"
	return *new(openai.CompletionUsageCompletionTokensDetails)
}

func cloneCompletionUsagePromptTokensDetails(
	details openai.CompletionUsagePromptTokensDetails,
) openai.CompletionUsagePromptTokensDetails {
	_ = "STUB: not implemented"
	return *new(openai.CompletionUsagePromptTokensDetails)
}

func cloneRespJSONFieldMap(fields map[string]respjson.Field) map[string]respjson.Field {
	_ = "STUB: not implemented"
	return nil
}

// shouldSuppressChunk returns true when the chunk contains no meaningful delta
// (no content, no refusal, no non-empty tool calls, and no finish reason).
// This filters out completely empty streaming events that cause noisy logs.
func (m *Model) shouldSuppressChunk(chunk openai.ChatCompletionChunk) bool {
	_ = "STUB: not implemented"
	return false
}

// Check for reasoning content - if present, don't suppress.

// Any meaningful payload disables suppression.

// If this chunk is a tool_calls delta, optionally suppress emission.
// By default we only expose tool calls in the final aggregated response
// to avoid noisy blank chunks. When showToolCallDelta is enabled, treat
// tool_call chunks as meaningful streaming payload.

// shouldSkipEmptyChunk returns true when the chunk contains no meaningful delta.
// This is a defensive check against malformed responses from certain providers
// that may return chunks with valid JSON fields but empty actual content.
//
// The order of checks matters:
// 1. Check reasoning content first - if present, don't skip
// 2. Check content - if valid, don't skip (even if empty string)
// 3. Check refusal - if valid, don't skip
// 4. Check toolcalls - if valid but array is empty, skip (defensive against panic)
// 5. Check usage - if valid, don't skip
// 6. Otherwise, skip
func (m *Model) shouldSkipEmptyChunk(chunk openai.ChatCompletionChunk) bool {
	_ = "STUB: not implemented"
	// Chunks that carry a finish reason are meaningful and should not be
	// skipped, even if they have no content or usage. This ensures that
	// streaming clients can observe termination semantics.
	return false
}

// No choices available, don't skip (let it be processed normally).

// Reasoning content is meaningful even if other fields are empty.

// Extract delta for inspection.

// Content or refusal indicates meaningful output.

// Tool calls are only meaningful when the array is non-empty.

// Otherwise there is no meaningful delta, skip the chunk.

// hasReasoningContent checks if the choices contains reasoning content.
func (m *Model) hasReasoningContent(choices []openai.ChatCompletionChunkChoice) bool {
	_ = "STUB: not implemented"
	return false
}

// extractReasoningContent extracts reasoning content from ExtraFields.
// The extraFields parameter should be a map with values that have a Raw() method.
func extractReasoningContent(extraFields map[string]respjson.Field) string {
	_ = "STUB: not implemented"
	return ""
}

// Ollama and some providers use "reasoning" instead of "reasoning_content".

// convertExtraFields converts SDK's respjson.Field map to a generic map[string]any.
// This preserves all extra fields from the API response (e.g., Gemini 3's thought_signature)
// for transparent passthrough to subsequent requests.
func convertExtraFields(extraFields map[string]respjson.Field) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// createPartialResponse creates a partial response from a chunk.
func (m *Model) createPartialResponse(chunk openai.ChatCompletionChunk) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

// Normalize object for chunks; upstream may emit empty object for toolcall deltas.

// Convert choices for partial responses (content streaming).

// Handle finish reason - FinishReason is a plain string.

func toolCallDeltaIndexPointer(toolCall openai.ChatCompletionChunkChoiceDeltaToolCall) *int {
	_ = "STUB: not implemented"
	return nil
}

// emitStreamingFinalResponse emits the final response with accumulated data.
func (m *Model) emitStreamingFinalResponse(
	ctx context.Context,
	stream *ssestream.Stream[openai.ChatCompletionChunk],
	acc openai.ChatCompletionAccumulator,
	idToIndexMap map[string]int,
	extraFieldsMap map[string]map[string]any,
	aggregatedReasoning string,
	emit responseEmitter,
) {
	_ = "STUB: not implemented"
	return

	// Check accumulated tool calls (batch processing after streaming is complete).
}

// If accumulator is empty but we have aggregated reasoning, create a response with it.

// Send error response.

// processAccumulatedToolCalls processes accumulated tool calls.
func (m *Model) processAccumulatedToolCalls(
	acc openai.ChatCompletionAccumulator,
	idToIndexMap map[string]int,
	extraFieldsMap map[string]map[string]any,
) []model.ToolCall {
	_ = "STUB: not implemented"
	return nil
}

// if openai return function tool call start with index 1 or more
// ChatCompletionAccumulator will return empty tool call for index like 0, skip it.

// Use the original index from ID->Index mapping if available, otherwise use loop index.

// Some providers (e.g., gpt-5-nano) may omit the tool_call ID.
// Synthesize a stable ID from the index to ensure proper pairing.

// Look up ExtraFields by ID first, then by index.

// OpenAI supports function tools for now.

// createFinalResponse creates the final response with accumulated data.
func (m *Model) createFinalResponse(
	acc openai.ChatCompletionAccumulator,
	hasToolCall bool,
	accumulatedToolCalls []model.ToolCall,
	aggregatedReasoning string,
) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

// Extract reasoning content from the accumulated message if available.

// Fallback to aggregated streaming deltas if accumulator didn't retain reasoning.

// If there are tool calls, add them to the final response.
// Usually only the first choice contains tool calls.

// Propagate finish reason from the accumulated choice so that the final
// aggregated response exposes the same termination semantics as the
// underlying provider.

// handleNonStreamingResponseWithEmitter handles non-streaming chat completion responses.
// It returns early when emit returns false.
func (m *Model) handleNonStreamingResponseWithEmitter(
	ctx context.Context,
	chatRequest openai.ChatCompletionNewParams,
	emit responseEmitter,
	opts ...openaiopt.RequestOption,
) {
	_ = "STUB: not implemented"
	return
}

// Some OpenAI-compatible providers return HTTP 200 with an error body
// instead of a proper 4xx/5xx status code. The SDK does not treat these
// as errors because it only inspects the HTTP status. However the error
// payload is still preserved in ChatCompletion.JSON.ExtraFields["error"].
// Detect this case and convert it into an explicit error response so that
// downstream consumers see a clear failure instead of an empty completion
// that can cause silent infinite loops in the agent flow.

// Call response callback on successful completion.

// createResponseFromCompletion converts a provider response into a model.Response.
func (m *Model) createResponseFromCompletion(chatCompletion *openai.ChatCompletion) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

// Convert constant to string

// Convert choices.

// Extract reasoning content from the message if available.

// Synthesize ID for providers that omit it (e.g., gpt-5-nano).

// Handle finish reason - FinishReason is a plain string.

// Convert usage information.

// Set system fingerprint if available.

// FileOptions is the options for file operations.
type FileOptions struct {
	// Path for file operations (default: /openapi/v1/files).
	Path string
	// Purpose for file upload (default: openai.FilePurposeUserData).
	Purpose openai.FilePurpose
	// Method for HTTP request (default: based on operation).
	Method string
	// Body for HTTP request (default: auto-generated based on operation).
	Body []byte
	// BaseURL override for this file request.
	BaseURL string
}

// FileOption is the option for file operations.
type FileOption func(*FileOptions)

// WithPath is the option for setting the file operation path.
func WithPath(path string) FileOption { _ = "STUB: not implemented"; return *new(FileOption) }

// WithPurpose is the option for setting the file upload purpose.
func WithPurpose(purpose openai.FilePurpose) FileOption {
	_ = "STUB: not implemented"
	return *new(FileOption)
}

// WithMethod is the option for setting the HTTP method.
func WithMethod(method string) FileOption { _ = "STUB: not implemented"; return *new(FileOption) }

// WithBody is the option for setting the HTTP request body.
func WithBody(body []byte) FileOption { _ = "STUB: not implemented"; return *new(FileOption) }

// WithFileBaseURL sets a per-request base URL override for file operations.
func WithFileBaseURL(url string) FileOption { _ = "STUB: not implemented"; return *new(FileOption) }

// UploadFile uploads a file to OpenAI and returns the file ID.
// The file can then be referenced in messages using AddFileID().
func (m *Model) UploadFile(ctx context.Context, filePath string, opts ...FileOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Open the file.

// Create middleware to construct multipart form data request.

// Set the correct path.

// Set custom HTTP method if specified.

// Use custom body if specified, otherwise create multipart form data.

// Continue with the modified request.

// Create empty file params since we're handling the file in middleware.

// Upload the file.

// UploadFileData uploads file data to OpenAI and returns the file ID.
// This is useful when you have file data in memory rather than a file path.
func (m *Model) UploadFileData(
	ctx context.Context,
	filename string,
	data []byte,
	opts ...FileOption,
) (string, error) {
	_ = "STUB: not implemented"
	// Apply default options based on variant.
	return "", nil
}

// Create file upload parameters with data reader.

// Set to nil to avoid duplicate multipart form construction by SDK.
// The middleware will handle all request body construction to ensure proper
// filename preservation and field ordering required by Venus platform.

// Create middleware to handle custom options.

// Set the correct path.

// Set custom HTTP method if specified.

// Use custom body if specified.

// Build multipart form to ensure filename suffix is preserved.

// purpose.

// file.

// Upload the file.

// DeleteFile deletes a file from OpenAI.
func (m *Model) DeleteFile(ctx context.Context, fileID string, opts ...FileOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Create middleware to handle custom options.

// Set custom HTTP method if specified.

// Use custom body if specified.

// GetFile retrieves file information from OpenAI.
func (m *Model) GetFile(
	ctx context.Context,
	fileID string,
	opts ...FileOption,
) (*openai.FileObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create middleware to handle custom options.

// Set the correct path.

// Set custom HTTP method if specified.

// Use custom body if specified.

// DownloadFile downloads the content for the given file ID.
func (m *Model) DownloadFile(
	ctx context.Context,
	fileID string,
) ([]byte, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// extractEmbeddedErrorResponse detects an error payload hidden inside an
// otherwise "successful" ChatCompletion. Some OpenAI-compatible providers
// return HTTP 200 with a JSON body like:
//
//	{"error": {"message": "...", "type": "...", "code": "..."}}
//
// The OpenAI SDK only checks HTTP status codes for errors, so it silently
// parses this into an empty ChatCompletion. The error object ends up in
// ChatCompletion.JSON.ExtraFields["error"] because "error" is not a
// recognized field in the ChatCompletion schema.
//
// This function checks for that specific condition and, when found, returns
// a model.Response with the original error information restored.
// Returns nil when no embedded error is detected.
func extractEmbeddedErrorResponse(cc *openai.ChatCompletion) *model.Response {
	_ = "STUB: not implemented"
	return nil
}

// Only inspect ExtraFields when the completion is empty (no choices).
// A completion with valid choices is never treated as an error, even if
// the provider happens to include extra fields named "error".

// Parse the raw error JSON to extract message and type.
// Use Raw() instead of Valid() because the SDK may mark non-schema
// object fields as "invalid" status even though the content is present.

// ExtraFields["error"] exists but is not a recognizable error object;
// leave it for the normal completion path.

// normalizeEmbeddedErrorString extracts a JSON string value.
// Returns "" for absent, null, or non-string values.
func normalizeEmbeddedErrorString(raw json.RawMessage) string { _ = "STUB: not implemented"; return "" }

// normalizeEmbeddedErrorCode converts a JSON code value (string, number, or
// null) into a string suitable for model.ResponseError.Code.
// Returns "" when the value is absent, null, or unparsable.
func normalizeEmbeddedErrorCode(raw json.RawMessage) string { _ = "STUB: not implemented"; return "" }

// Try number (some providers return numeric codes).
