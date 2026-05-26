//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package ollama provides Ollama-compatible model implementations.
package ollama

import (
	"context"
	"net/http"

	"github.com/ollama/ollama/api"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

// Model implements the model.Model interface for Ollama API.
type Model struct {
	client                     *api.Client
	name                       string
	host                       string
	contextWindow              int
	contextWindowConfigured    bool
	contextWindowDiscovered    bool
	httpClient                 *http.Client
	channelBufferSize          int
	chatRequestCallback        ChatRequestCallbackFunc
	chatResponseCallback       ChatResponseCallbackFunc
	chatChunkCallback          ChatChunkCallbackFunc
	chatStreamCompleteCallback ChatStreamCompleteCallbackFunc
	enableTokenTailoring       bool                    // Enable automatic token tailoring.
	maxInputTokens             int                     // Max input tokens for token tailoring.
	tokenCounter               model.TokenCounter      // Token counter for token tailoring.
	tailoringStrategy          model.TailoringStrategy // Tailoring strategy for token tailoring.
	// Token tailoring budget parameters (instance-level overrides).
	protocolOverheadTokens int
	reserveOutputTokens    int
	inputTokensFloor       int
	outputTokensFloor      int
	safetyMarginRatio      float64
	maxInputTokensRatio    float64
	// Additional options for Ollama API. such as temperature/top_p
	options   map[string]any
	keepAlive *api.Duration
}

// New creates a new Ollama model adapter.
func New(name string, opts ...Option) *Model { _ = "STUB: not implemented"; return nil }

// Create Ollama API client.

// Info returns the model information.
func (m *Model) Info() model.Info { _ = "STUB: not implemented"; return *new(model.Info) }

func (m *Model) runChatRequestCallback(
	ctx context.Context,
	chatRequest *api.ChatRequest,
) {
	_ = "STUB: not implemented"
	return
}

func (m *Model) runChatResponseCallback(
	ctx context.Context,
	chatRequest *api.ChatRequest,
	chatResponse *api.ChatResponse,
) {
	_ = "STUB: not implemented"
	return
}

func (m *Model) runChatChunkCallback(
	ctx context.Context,
	chatRequest *api.ChatRequest,
	chatChunk *api.ChatResponse,
) {
	_ = "STUB: not implemented"
	return
}

func (m *Model) runChatStreamCompleteCallback(
	ctx context.Context,
	chatRequest *api.ChatRequest,
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

// buildChatRequest builds the chat request for the Ollama API.
func (m *Model) buildChatRequest(request *model.Request) (*api.ChatRequest, error) {
	_ = "STUB: not implemented"
	// Convert messages to Ollama format.
	return nil, nil
}

// Build chat request.

// Set stream option.

// Set generation parameters.

// Set keep alive if configured.

// handleNonStreamingResponse sends a non-streaming request to the Ollama API.
func (m *Model) handleNonStreamingResponse(
	ctx context.Context,
	chatRequest api.ChatRequest,
	responseID string,
	responseChan chan<- *model.Response,
) {
	_ = "STUB: not implemented"
	// Issue non-streaming request.
	return
}

// Emit final response.

// handleStreamingResponse sends a streaming request to the Ollama API.
func (m *Model) handleStreamingResponse(
	ctx context.Context,
	chatRequest api.ChatRequest,
	responseID string,
	responseChan chan<- *model.Response,
) {
	_ = "STUB: not implemented"
	return
}

// Emit partial response.

// Call the stream complete callback before surfacing the terminal result.

// sendErrorResponse sends an error response through the channel.
func (m *Model) sendErrorResponse(
	ctx context.Context,
	responseChan chan<- *model.Response,
	responseID string,
	errType string,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

func newResponseID() string { _ = "STUB: not implemented"; return "" }

// getContextWindow retrieves the context window size for the model.
// for example, ollama /api/show show model info, and get context_length
//
//	{
//	   "license": "xxx",
//	   "modelfile": "xxx",
//	   "parameters": "xxx",
//	   "template": "xxx",
//	   "details": {
//	   },
//	   "model_info": {
//	       "general.architecture": "llama",
//	       "general.basename": "DeepSeek-R1-Distill-Llama",
//	       "general.file_type": 15,
//	       "general.parameter_count": 8030261312,
//	       "general.quantization_version": 2,
//	       "general.size_label": "8B",
//	       "general.type": "model",
//	       "llama.attention.head_count": 32,
//	       "llama.attention.head_count_kv": 8,
//	       "llama.attention.layer_norm_rms_epsilon": 0.00001,
//	       "llama.block_count": 32,
//	       "llama.context_length": 131072,
//	       "llama.embedding_length": 4096,
//	       "llama.feed_forward_length": 14336,
//	       "llama.rope.dimension_count": 128,
//	       "llama.rope.freq_base": 500000,
//	       "llama.vocab_size": 128256,
//	       "tokenizer.ggml.add_bos_token": true,
//	       "tokenizer.ggml.add_eos_token": false,
//	       "tokenizer.ggml.bos_token_id": 128000,
//	       "tokenizer.ggml.eos_token_id": 128001,
//	       "tokenizer.ggml.merges": null,
//	       "tokenizer.ggml.model": "gpt2",
//	       "tokenizer.ggml.padding_token_id": 128001,
//	       "tokenizer.ggml.pre": "llama-bpe",
//	       "tokenizer.ggml.token_type": null,
//	       "tokenizer.ggml.tokens": null
//	   }
//	}
//
// llama.context_length is the context window size
// ref: https://github.com/ollama/ollama/blob/main/docs/api.md#show-model-information
func (m *Model) getContextWindow() (int, error) { _ = "STUB: not implemented"; return 0, nil }

// convertMessages converts model messages to Ollama messages.
func convertMessages(messages []model.Message) ([]api.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertChatResponse(resp api.ChatResponse, responseID string) (*model.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convertMessage converts a model message to an Ollama message.
// ollama only support system/user/assistant role msg
func convertMessage(msg model.Message) (api.Message, error) {
	_ = "STUB: not implemented"
	return *new(api.Message), nil
}

// ollama role ("system", "user", or "assistant")

// convertTools converts our tool declarations to Ollama tool parameters.
func convertTools(tools map[string]tool.Tool) []api.Tool { _ = "STUB: not implemented"; return nil }

// buildToolDescription builds the description for a tool.
// It appends the output schema to the description.
func buildToolDescription(declaration *tool.Declaration) string {
	_ = "STUB: not implemented"
	return ""
}

func imageToURLOrBase64(image *model.Image) string { _ = "STUB: not implemented"; return "" }

func argsToObject(args []byte) (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }
