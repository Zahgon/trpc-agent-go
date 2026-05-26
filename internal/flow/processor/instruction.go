//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package processor

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/model"
)

// InstructionRequestProcessor implements instruction processing logic.
type InstructionRequestProcessor struct {
	// Instruction is the instruction to add to requests.
	Instruction string
	// InstructionResolver, if provided, supplies the instruction for each
	// request based on the current invocation. When set, this takes
	// precedence over InstructionGetter and Instruction.
	InstructionResolver func(*agent.Invocation) string
	// InstructionGetter, if provided, dynamically supplies the instruction
	// each time a request is processed. When set, this takes precedence over
	// the static Instruction field.
	InstructionGetter func() string
	// SystemPrompt is the system prompt to add to requests.
	SystemPrompt string
	// SystemPromptResolver, if provided, supplies the system prompt for each
	// request based on the current invocation. When set, this takes
	// precedence over SystemPromptGetter and SystemPrompt.
	SystemPromptResolver func(*agent.Invocation) string
	// SystemPromptGetter, if provided, dynamically supplies the system prompt
	// each time a request is processed. When set, this takes precedence over
	// the static SystemPrompt field.
	SystemPromptGetter func() string
	// OutputSchema is the JSON schema for output validation.
	// When provided, JSON output instructions are automatically injected.
	OutputSchema map[string]any
	// StructuredOutputSchema is the JSON schema generated from structured_output.
	// When provided, it takes precedence over OutputSchema for instruction injection.
	StructuredOutputSchema map[string]any
}

const (
	jsonInstructionsStrictTemplate = `IMPORTANT: Return ONLY a JSON object that
conforms to the schema below.
- Do NOT include the schema itself in your output.
- Do NOT include explanations, comments, or markdown fences.
- Do NOT add keys other than those defined in the schema's properties.
- The response must be a single JSON object instance, not wrapped, and no
  trailing text.

Schema (for reference only, do not include this in your output):
%s
`

	jsonInstructionsToolsTemplate = `IMPORTANT:
- You MAY call tools when needed (function calling).
- While you are calling tools, do NOT provide a user-facing answer.
- When you are ready to provide the final answer, return ONLY a JSON object
  that conforms to the schema below.
- Do NOT include the schema itself in your output.
- Do NOT include explanations, comments, or markdown fences.
- Do NOT add keys other than those defined in the schema's properties.
- The final response must be a single JSON object instance, not wrapped, and
  no trailing text.

Schema (for reference only, do not include this in your output):
%s
`
)

// InstructionRequestProcessorOption is a function that can be used to configure the instruction request processor.
type InstructionRequestProcessorOption func(*InstructionRequestProcessor)

// WithOutputSchema adds the output schema to the instruction request processor.
func WithOutputSchema(outputSchema map[string]any) InstructionRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(InstructionRequestProcessorOption)
}

// WithStructuredOutputSchema adds the structured output schema to the instruction request processor.
// This is used as a fallback when the model provider does not natively enforce JSON Schema.
func WithStructuredOutputSchema(schema map[string]any) InstructionRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(InstructionRequestProcessorOption)
}

// WithInstructionResolver configures a dynamic resolver for instruction
// content based on the current invocation.
func WithInstructionResolver(
	resolver func(*agent.Invocation) string,
) InstructionRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(InstructionRequestProcessorOption)
}

// WithInstructionGetter configures a dynamic getter for instruction content.
// When provided, this getter is called for every request, allowing callers to
// update the instruction at runtime without reconstructing the processor/agent.
func WithInstructionGetter(getter func() string) InstructionRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(InstructionRequestProcessorOption)
}

// WithSystemPromptResolver configures a dynamic resolver for system prompt
// content based on the current invocation.
func WithSystemPromptResolver(
	resolver func(*agent.Invocation) string,
) InstructionRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(InstructionRequestProcessorOption)
}

// WithSystemPromptGetter configures a dynamic getter for system prompt content.
// When provided, this getter is called for every request, allowing callers to
// update the system prompt at runtime without reconstructing the processor/agent.
func WithSystemPromptGetter(getter func() string) InstructionRequestProcessorOption {
	_ = "STUB: not implemented"
	return *new(InstructionRequestProcessorOption)
}

// NewInstructionRequestProcessor creates a new instruction request processor.
func NewInstructionRequestProcessor(
	instruction, systemPrompt string,
	opts ...InstructionRequestProcessorOption,
) *InstructionRequestProcessor {
	_ = "STUB: not implemented"
	return nil
}

// ProcessRequest implements the flow.RequestProcessor interface.
// It adds instruction content and system prompt to the request if provided.
// Prompt placeholders are rendered before any structured-output instructions are appended.
func (p *InstructionRequestProcessor) ProcessRequest(
	ctx context.Context,
	invocation *agent.Invocation,
	req *model.Request,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// Process instruction and system prompt with state injection.

// Update the request messages with processed instructions.

// Send a preprocessing event.

// processInstructionsWithState resolves instruction and system prompt content,
// renders prompt placeholders, and then appends structured-output instructions.
func (p *InstructionRequestProcessor) processInstructionsWithState(
	ctx context.Context,
	invocation *agent.Invocation,
) (string, string) {
	_ = "STUB: not implemented"
	// Prefer invocation-based resolvers, then dynamic getters, then static
	// fields.
	return "", ""
}

// Automatically inject JSON output instructions after prompt rendering so
// literal schema braces are never interpreted as placeholders.
// Precedence: invocation.StructuredOutputSchema > StructuredOutputSchema > OutputSchema.

func (p *InstructionRequestProcessor) resolveStructuredOutputSchema(invocation *agent.Invocation) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// combineInstructions combines existing instruction with new JSON
// instructions.
func (p *InstructionRequestProcessor) combineInstructions(
	existingInstruction, jsonInstructions string,
) string {
	_ = "STUB: not implemented"
	return ""
}

// injectStateIntoContent injects session state into the given content.
func (p *InstructionRequestProcessor) injectStateIntoContent(
	ctx context.Context,
	invocation *agent.Invocation,
	content, contentType string,
) string {
	_ = "STUB: not implemented"
	return ""
}

// updateRequestMessages updates the request messages with processed instructions.
func (p *InstructionRequestProcessor) updateRequestMessages(req *model.Request, processedInstruction, processedSystemPrompt string) {
	_ = "STUB: not implemented"
	return
}

// updateExistingSystemMessage updates an existing system message with new instructions.
func (p *InstructionRequestProcessor) updateExistingSystemMessage(
	req *model.Request, systemMsgIndex int, processedInstruction, processedSystemPrompt string,
) {
	_ = "STUB: not implemented"
	return
}

// createNewSystemMessage creates a new system message with combined instructions.
func (p *InstructionRequestProcessor) createNewSystemMessage(
	req *model.Request, processedInstruction, processedSystemPrompt string,
) {
	_ = "STUB: not implemented"
	return
}

// buildSystemContent builds the content for a new system message.
func (p *InstructionRequestProcessor) buildSystemContent(processedInstruction, processedSystemPrompt string) string {
	_ = "STUB: not implemented"
	return ""
}

// sendPreprocessingEvent sends a preprocessing event if invocation is available.
func (p *InstructionRequestProcessor) sendPreprocessingEvent(
	ctx context.Context,
	invocation *agent.Invocation,
	ch chan<- *event.Event,
) {
	_ = "STUB: not implemented"
	return
}

// findSystemMessageIndex finds the index of the first system message in the messages slice.
// Returns -1 if no system message is found.
func findSystemMessageIndex(messages []model.Message) int { _ = "STUB: not implemented"; return 0 }

func findLastSystemMessageIndex(messages []model.Message) int { _ = "STUB: not implemented"; return 0 }

// containsInstruction checks if the given content already contains the instruction.
func containsInstruction(content, instruction string) bool {
	_ = "STUB: not implemented"
	// strings.Contains handles both exact match and substring cases
	return false
}

func invocationHasTools(invocation *agent.Invocation) bool { _ = "STUB: not implemented"; return false }

func (p *InstructionRequestProcessor) generateStructuredOutputJSONInstructions(
	invocation *agent.Invocation,
	schema map[string]any,
) string {
	_ = "STUB: not implemented"
	return ""
}

// generateJSONInstructions generates JSON output instructions based on a schema.
func (p *InstructionRequestProcessor) generateJSONInstructions(schema map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

func (p *InstructionRequestProcessor) generateJSONInstructionsWithTemplate(
	schema map[string]any,
	template string,
) string {
	_ = "STUB: not implemented"
	// Convert schema to a readable format for the instruction.
	return ""
}

// formatSchemaForInstruction formats the schema for inclusion in instructions.
func (p *InstructionRequestProcessor) formatSchemaForInstruction(schema map[string]any) string {
	_ = "STUB: not implemented"
	// For now, we'll create a simple JSON representation.
	// In a more sophisticated implementation, we could parse the schema more intelligently.
	return ""
}

// Fallback to a simple string representation.
