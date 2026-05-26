//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package model

import (
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// Role represents the role of a message author.
type Role string

// Role constants for message authors.
const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleTool      Role = "tool"
)

// Thinking parameter keys used in API requests.
const (
	// ThinkingEnabledKey is the key used for enabling thinking mode in API requests.
	ThinkingEnabledKey = "thinking_enabled"
	// ThinkingTokensKey is the key used for thinking tokens configuration in API requests.
	ThinkingTokensKey = "thinking_tokens"
	// ReasoningContentKey is the key used for reasoning content in API responses.
	ReasoningContentKey = "reasoning_content"
	// ReasoningContentKeyAlt is the alternative key used by some providers (e.g. Ollama).
	ReasoningContentKeyAlt = "reasoning"
	// EnabledThinkingKey is the key used for enabling thinking mode in API requests e.g. Qwen model.
	EnabledThinkingKey = "enabled_thinking"
)

// String returns the string representation of the role.
func (r Role) String() string {
	_ = "STUB: not implemented"

	// IsValid checks if the role is one of the defined constants.
	return ""
}

func (r Role) IsValid() bool { _ = "STUB: not implemented"; return false }

// Message represents a single message in a conversation.
type Message struct {
	// Role is the role of the message author.
	Role Role `json:"role"`
	// Content is the message content.
	Content string `json:"content,omitempty"`
	// ContentParts is the content parts for multimodal messages.
	ContentParts []ContentPart `json:"content_parts,omitempty"`
	// ToolID is the ID of the tool used by tool response.
	ToolID string `json:"tool_id,omitempty"`
	// ToolName is the name of the tool used by tool response.
	ToolName string `json:"tool_name,omitempty"`
	// ToolCalls is the optional tool calls for the message.
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
	// ReasoningContent is hunyuan or deepseek think content
	// - https://api-docs.deepseek.com/api/create-chat-completion#responses
	ReasoningContent string `json:"reasoning_content,omitempty"`
	// ReasoningSignature is a token that verifies the reasoning text was generated
	// by the model. When passing a reasoning block back to the API in a multi-turn
	// conversation, include the text and its signature unmodified.
	// Currently used by AWS Bedrock (Claude) models.
	ReasoningSignature string `json:"reasoning_signature,omitempty"`
}

// AddFilePath adds a file path to the message.
func (m *Message) AddFilePath(filepath string) error { _ = "STUB: not implemented"; return nil }

// AddFileData adds a file data to the message.
// The argument of data is the raw file data without base64 encoding.
func (m *Message) AddFileData(name string, data []byte, mimetype string) {
	_ = "STUB: not implemented"
	return
}

// AddFileURL adds a URL-based file to the message.
func (m *Message) AddFileURL(name, url, mimetype string) { _ = "STUB: not implemented"; return }

// AddFileID adds a file ID to the message.
// The file id can be obtained from the response of the upload file API.
func (m *Message) AddFileID(fileID string) { _ = "STUB: not implemented"; return }

// AddFileIDWithName adds a file ID and filename to the message.
//
// The filename is not always sent to the model provider when FileID is used,
// but it can be useful for downstream tooling (for example, staging user file
// inputs into a skill workspace).
func (m *Message) AddFileIDWithName(fileID, name string) { _ = "STUB: not implemented"; return }

// AddImageURL adds an image URL to the message.
// The argument of detail is the detail level: "low", "high", "auto".
// If detail is empty, it will be set to "auto".
func (m *Message) AddImageURL(url, detail string) { _ = "STUB: not implemented"; return }

// AddImageFilePath adds an image file path to the message.
// The argument detail specifies the detail level: "low", "high", or "auto".
// If detail is empty, it will be set to "auto".
// Supported formats:
//
//   - PNG (.png)
//   - JPEG (.jpeg, .jpg)
//   - WEBP (.webp)
//   - Non-animated GIF (.gif)
//
// Reference: https://platform.openai.com/docs/guides/images-vision.
func (m *Message) AddImageFilePath(path string, detail string) error {
	_ = "STUB: not implemented"
	return nil
}

// Infer format from the file extension.

// AddImageData adds an image data to the message.
// The argument of data is the raw image data without base64 encoding.
// The argument of detail is the detail level: "low", "high", "auto".
// If detail is empty, it will be set to "auto".
func (m *Message) AddImageData(data []byte, detail, format string) {
	_ = "STUB: not implemented"
	return
}

// AddAudioFilePath adds an audio file path to the message.
func (m *Message) AddAudioFilePath(path string) error { _ = "STUB: not implemented"; return nil }

// Infer format from the file extension.

// AddAudioData adds an audio data to the message.
// The argument of data is the raw audio data without base64 encoding.
// The argument of format is the format of the audio data.
// Currently supports "wav" and "mp3".
func (m *Message) AddAudioData(data []byte, format string) { _ = "STUB: not implemented"; return }

// ContentType represents the type of content.
type ContentType string

// ContentType constants for content types.
const (
	ContentTypeText  ContentType = "text"
	ContentTypeImage ContentType = "image"
	ContentTypeAudio ContentType = "audio"
	ContentTypeFile  ContentType = "file"
)

// ContentPart represents a single content part in a multimodal message.
type ContentPart struct {
	// Type is the type of content: "text", "image", "audio", "file"
	Type ContentType `json:"type"`
	// Text is the text content.
	Text *string `json:"text,omitempty"`
	// Image is the image data.
	Image *Image `json:"image,omitempty"`
	// Audio is the audio data.
	Audio *Audio `json:"audio,omitempty"`
	// File is the file data.
	File *File `json:"file,omitempty"`
}

// File represents file content for file input models.
type File struct {
	// Name is the name of the file, used when passing the file to the model as a string.
	Name string `json:"filename"`
	// URL is the URL of the file.
	URL string `json:"url,omitempty"`
	// Data is the raw file data, used when passing the file to the model as a string.
	Data []byte `json:"data"`
	// FileID is the ID of an uploaded file to use as input.
	FileID string `json:"file_id"`
	// MimeType is the format of the file data.
	MimeType string `json:"format,omitempty"`
}

// FileURLText returns a textual representation for providers that cannot accept URL-based files.
func FileURLText(file *File) string { _ = "STUB: not implemented"; return "" }

// Image represents an image data for vision models.
type Image struct {
	// URL is the URL of the image.
	URL string `json:"url"`
	// Data is the raw image data.
	Data []byte `json:"data"`
	// Detail is the detail level: "low", "high", "auto".
	Detail string `json:"detail,omitempty"`
	// Format is the image format.
	// Data-backed images usually use a subtype such as "png".
	// URL-backed images may use a full MIME type such as "image/png".
	Format string `json:"format,omitempty"`
}

// Audio represents audio input for audio models.
type Audio struct {
	// Data is the raw audio data.
	Data []byte `json:"data"`
	// Format is the format of the encoded audio data. Currently supports "wav" and "mp3".
	Format string `json:"format"`
}

// NewSystemMessage creates a new system message.
func NewSystemMessage(content string) Message { _ = "STUB: not implemented"; return *new(Message) }

// NewUserMessage creates a new user message.
func NewUserMessage(content string) Message { _ = "STUB: not implemented"; return *new(Message) }

// NewToolMessage creates a new tool message.
func NewToolMessage(toolID, toolName, content string) Message {
	_ = "STUB: not implemented"
	return *new(Message)
}

// NewAssistantMessage creates a new assistant message.
func NewAssistantMessage(content string) Message { _ = "STUB: not implemented"; return *new(Message) }

// GenerationConfig contains configuration for text generation.
type GenerationConfig struct {
	// MaxTokens is the maximum number of tokens to generate.
	MaxTokens *int `json:"max_tokens,omitempty"`

	// Temperature controls randomness (0.0 to 2.0).
	Temperature *float64 `json:"temperature,omitempty"`

	// TopP controls nucleus sampling (0.0 to 1.0).
	TopP *float64 `json:"top_p,omitempty"`

	// Stream indicates whether to stream the response.
	Stream bool `json:"stream"`

	// Stop sequences where the API will stop generating further tokens.
	Stop []string `json:"stop,omitempty"`

	// PresencePenalty penalizes new tokens based on their existing frequency.
	PresencePenalty *float64 `json:"presence_penalty,omitempty"`

	// FrequencyPenalty penalizes new tokens based on their frequency in the text so far.
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`

	// ReasoningEffort limits the reasoning effort for reasoning models.
	// The accepted values depend on the provider:
	//   - OpenAI o-series: "low", "medium", "high".
	//   - Anthropic adaptive thinking: "low", "medium", "high", "max", and
	//     "xhigh" on models that support it.
	//   - DeepSeek v4 (deepseek-v4-pro / deepseek-v4-flash): "high", "max".
	//     For backward compatibility, the DeepSeek service maps "low" and
	//     "medium" to "high", and "xhigh" to "max", so older configurations
	//     still work without errors.
	// See:
	//   - https://platform.openai.com/docs/api-reference/chat/create
	//   - https://platform.claude.com/docs/en/build-with-claude/adaptive-thinking
	//   - https://api-docs.deepseek.com/api/create-chat-completion
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`

	// ThinkingEnabled toggles thinking/reasoning mode for providers that
	// expose it (e.g. DeepSeek v4 via the "thinking" object, Claude / Gemini
	// via the OpenAI-compatible API).
	//
	// When left nil, trpc-agent-go does NOT emit any thinking-toggle field
	// in the outgoing request, so the provider's server-side default takes
	// effect (DeepSeek v4 defaults to "enabled"). Set to *true / *false to
	// force a specific behavior. Anthropic adaptive-thinking models map *true
	// to thinking.type=adaptive, while older Anthropic models map *true to
	// thinking.type=enabled with ThinkingTokens.
	ThinkingEnabled *bool `json:"thinking_enabled,omitempty"`

	// ThinkingTokens controls the fixed thinking token budget for providers that support it.
	// Anthropic adaptive-thinking models ignore this field and use ReasoningEffort instead.
	ThinkingTokens *int `json:"thinking_tokens,omitempty"`
}

// GenerationConfigPatch selectively overrides fields in GenerationConfig.
//
// It is designed for per-request overrides where callers only want to change
// a subset of parameters (e.g. temperature or max_tokens) without rebuilding
// agents or graphs.
//
// Semantics:
//   - Pointer fields: nil means "do not override".
//   - Stop: nil means "do not override"; an empty slice clears Stop.
//   - Stream: nil means "do not override".
type GenerationConfigPatch struct {
	MaxTokens        *int     `json:"max_tokens,omitempty"`
	Temperature      *float64 `json:"temperature,omitempty"`
	TopP             *float64 `json:"top_p,omitempty"`
	Stream           *bool    `json:"stream,omitempty"`
	Stop             []string `json:"stop,omitempty"`
	PresencePenalty  *float64 `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`
	ReasoningEffort  *string  `json:"reasoning_effort,omitempty"`
	ThinkingEnabled  *bool    `json:"thinking_enabled,omitempty"`
	ThinkingTokens   *int     `json:"thinking_tokens,omitempty"`
}

// ApplyGenerationConfigPatch applies patch to base and returns the merged
// configuration.
func ApplyGenerationConfigPatch(
	base GenerationConfig,
	patch GenerationConfigPatch,
) GenerationConfig {
	_ = "STUB: not implemented"
	return *new(GenerationConfig)
}

// Request is the request to the model.
type Request struct {
	// Messages is the conversation history.
	Messages []Message `json:"messages"`

	// GenerationConfig contains the generation parameters.
	GenerationConfig `json:"generation_config,omitempty"`

	// StructuredOutput defines how the model should produce structured output.
	// When set, the underlying model adapter may use native structured output
	// capabilities (e.g. OpenAI response_format with json_schema) to enforce
	// JSON formatting. This field is optional and provider-agnostic.
	StructuredOutput *StructuredOutput `json:"structured_output,omitempty"`

	// ExtraFields stores provider-specific top-level request body fields.
	// Model adapters merge these with model-level extra fields when supported;
	// request-level values take precedence.
	ExtraFields map[string]any `json:"-"`

	Tools map[string]tool.Tool `json:"-"` // Tools are not serialized, handled separately
}

// RequestOption configures a Request.
type RequestOption func(*Request)

// NewRequest creates a model request from messages and applies options.
func NewRequest(messages []Message, opts ...RequestOption) *Request {
	_ = "STUB: not implemented"
	return nil
}

// WithStructuredOutputJSON sets JSON schema structured output for a request.
// The schema is constructed automatically from the provided example type.
//
// This configures provider-native structured output where supported. Direct
// model callers still receive normal model.Response values and should unmarshal
// the final JSON content themselves.
func WithStructuredOutputJSON(examplePtr any, strict bool, description string) RequestOption {
	_ = "STUB: not implemented"
	return *new(RequestOption)
}

// ToolCall represents a call to a tool (function) in the model response.
type ToolCall struct {
	// Type of the tool. Currently, only `function` is supported.
	Type string `json:"type"`
	// Function definition for the tool
	Function FunctionDefinitionParam `json:"function,omitempty"`
	// The ID of the tool call returned by the model.
	ID string `json:"id,omitempty"`

	// Index is the index of the tool call in the message for streaming responses.
	Index *int `json:"index,omitempty"`

	// ExtraFields stores additional provider-specific fields for transparent passthrough.
	// For example, Gemini 3's thought_signature for multi-turn function calling.
	ExtraFields map[string]any `json:"extra_fields,omitempty"`
}

// FunctionDefinitionParam represents the parameters for a function definition in tool calls.
type FunctionDefinitionParam struct {
	// The name of the function to be called. Must be a-z, A-Z, 0-9, or contain
	// underscores and dashes, with a maximum length of 64.
	Name string `json:"name"`
	// Whether to enable strict schema adherence when generating the function call. If
	// set to true, the model will follow the exact schema defined in the `parameters`
	// field. Only a subset of JSON Schema is supported when `strict` is `true`. Learn
	// more about Structured Outputs in the
	// [function calling guide](docs/guides/function-calling).
	Strict bool `json:"strict,omitempty"`
	// A description of what the function does, used by the model to choose when and
	// how to call the function.
	Description string `json:"description,omitempty"`

	// Optional arguments to pass to the function, json-encoded.
	Arguments []byte `json:"arguments,omitempty"`
}

// MarshalJSON customizes JSON marshaling for FunctionDefinitionParam.
// This prevents double-encoding of the Arguments field by treating it as a string.
func (f FunctionDefinitionParam) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalJSON customizes JSON unmarshaling for FunctionDefinitionParam.
// This ensures the Arguments field is properly decoded from JSON string to []byte.
func (f *FunctionDefinitionParam) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// toMIME maps file extensions to their corresponding MIME types.
var toMIME = map[string]string{
	".txt":  "text/plain",
	".md":   "text/markdown",
	".html": "text/html",
	".json": "application/json",
	".doc":  "application/msword",
	".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	".pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	".pdf":  "application/pdf",
	".c":    "text/x-c",
	".cpp":  "text/x-c++",
	".cs":   "text/x-csharp",
	".java": "text/x-java",
	".js":   "text/javascript",
	".ts":   "application/typescript",
	".py":   "text/x-python",
	".rb":   "text/x-ruby",
	".css":  "text/css",
	".sh":   "application/x-sh",
	".php":  "text/x-php",
	".tex":  "text/x-tex",
}

// inferMimeType infers the MIME type from the file extension of the given path.
// Returns the MIME type string, or an error if the extension is unknown.
func inferMimeType(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// StructuredOutputType defines the type of structured output.
type StructuredOutputType string

const (
	// StructuredOutputJSONSchema enables structured JSON output.
	StructuredOutputJSONSchema StructuredOutputType = "json_schema"
)

// JSONSchemaConfig defines the configuration for JSON schema structured output.
type JSONSchemaConfig struct {
	// Name is the name of the structured output format.
	Name string `json:"name,omitempty"`
	// Schema is the JSON schema definition.
	Schema map[string]any `json:"schema"`
	// Strict controls whether to enforce strict schema adherence.
	Strict bool `json:"strict,omitempty"`
	// Description provides context for the model about the structured output.
	Description string `json:"description,omitempty"`
}

// StructuredOutput defines how the model should produce structured output.
type StructuredOutput struct {
	// Type specifies the structured output type.
	Type StructuredOutputType `json:"type"`
	// JSONSchema is used when Type is StructuredOutputJSONSchema.
	JSONSchema *JSONSchemaConfig `json:"json_schema,omitempty"`
}

func structuredOutputJSON(examplePtr any, strict bool, description string) *StructuredOutput {
	_ = "STUB: not implemented"
	return nil
}
