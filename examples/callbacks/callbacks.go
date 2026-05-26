//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package main

import (
	"context"
	"fmt"

	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

var (
	// Global callback configurations using chain registration.
	// This demonstrates how to create reusable callback configurations.
	_ = model.NewCallbacks().
		RegisterBeforeModel(func(ctx context.Context, args *model.BeforeModelArgs) (*model.BeforeModelResult, error) {
			fmt.Printf("🌐 Global BeforeModel: processing %d messages\n", len(args.Request.Messages))
			return nil, nil
		}).
		RegisterAfterModel(func(ctx context.Context, args *model.AfterModelArgs) (*model.AfterModelResult, error) {
			if args.Error != nil {
				fmt.Printf("🌐 Global AfterModel: error occurred\n")
			} else {
				fmt.Printf("🌐 Global AfterModel: processed successfully\n")
			}
			return nil, nil
		})

	_ = tool.NewCallbacks().
		RegisterBeforeTool(func(ctx context.Context, args *tool.BeforeToolArgs) (*tool.BeforeToolResult, error) {
			fmt.Printf("🌐 Global BeforeTool: executing %s\n", args.ToolName)
			// Return BeforeToolResult.ModifiedArguments when a callback needs
			// the actual tool execution to use updated arguments.
			return nil, nil
		}).
		RegisterAfterTool(func(ctx context.Context, args *tool.AfterToolArgs) (*tool.AfterToolResult, error) {
			if args.Error != nil {
				fmt.Printf("🌐 Global AfterTool: %s failed\n", args.ToolName)
			} else {
				fmt.Printf("🌐 Global AfterTool: %s completed\n", args.ToolName)
			}
			return nil, nil
		})

	_ = agent.NewCallbacks().
		RegisterBeforeAgent(func(ctx context.Context, args *agent.BeforeAgentArgs) (*agent.BeforeAgentResult, error) {
			fmt.Printf("🌐 Global BeforeAgent: starting %s\n", args.Invocation.AgentName)
			return nil, nil
		}).
		RegisterAfterAgent(func(ctx context.Context, args *agent.AfterAgentArgs) (*agent.AfterAgentResult, error) {
			if args.Error != nil {
				fmt.Printf("🌐 Global AfterAgent: execution failed\n")
			} else {
				fmt.Printf("🌐 Global AfterAgent: execution completed\n")
			}
			return nil, nil
		})
)

// createModelCallbacks creates and configures model callbacks.
func (c *multiTurnChatWithCallbacks) createModelCallbacks() *model.Callbacks {
	_ = "STUB: not implemented"
	// Using traditional registration.
	return nil
}

// createBeforeModelCallback creates the before model callback.
func (c *multiTurnChatWithCallbacks) createBeforeModelCallback() model.BeforeModelCallbackStructured {
	_ = "STUB: not implemented"
	return *new(model.BeforeModelCallbackStructured)
}

// You can get the invocation from the context.

// createAfterModelCallback creates the after model callback.
func (c *multiTurnChatWithCallbacks) createAfterModelCallback() model.AfterModelCallbackStructured {
	_ = "STUB: not implemented"
	return *new(model.AfterModelCallbackStructured)
}

// createToolCallbacks creates and configures tool callbacks.
func (c *multiTurnChatWithCallbacks) createToolCallbacks() *tool.Callbacks {
	_ = "STUB: not implemented"
	// Using traditional registration.
	return nil
}

// createBeforeToolCallback creates the before tool callback.
func (c *multiTurnChatWithCallbacks) createBeforeToolCallback() tool.BeforeToolCallbackStructured {
	_ = "STUB: not implemented"
	return *new(tool.BeforeToolCallbackStructured)
}

// Demonstrate argument modification capability.
// Assigning args.Arguments makes the updated value visible to later
// callback logic. Returning ModifiedArguments makes the tool execute
// with the updated arguments.

// Example: normalize the operation while preserving the calculator schema.

// createAfterToolCallback creates the after tool callback.
func (c *multiTurnChatWithCallbacks) createAfterToolCallback() tool.AfterToolCallbackStructured {
	_ = "STUB: not implemented"
	return *new(tool.AfterToolCallbackStructured)
}

// createAgentCallbacks creates and configures agent callbacks.
func (c *multiTurnChatWithCallbacks) createAgentCallbacks() *agent.Callbacks {
	_ = "STUB: not implemented"
	// Using traditional registration.
	return nil
}

// createBeforeAgentCallback creates the before agent callback.
func (c *multiTurnChatWithCallbacks) createBeforeAgentCallback() agent.BeforeAgentCallbackStructured {
	_ = "STUB: not implemented"
	return *new(agent.BeforeAgentCallbackStructured)
}

// createAfterAgentCallback creates the after agent callback.
func (c *multiTurnChatWithCallbacks) createAfterAgentCallback() agent.AfterAgentCallbackStructured {
	_ = "STUB: not implemented"
	return *new(agent.AfterAgentCallbackStructured)
}

// Helper functions for callback logic.

func (c *multiTurnChatWithCallbacks) extractLastUserMessage(req *model.Request) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *multiTurnChatWithCallbacks) shouldReturnCustomResponse(userMsg string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *multiTurnChatWithCallbacks) createCustomResponse() *model.Response {
	_ = "STUB: not implemented"
	return nil
}

func (c *multiTurnChatWithCallbacks) handleModelFinished(resp *model.Response) {
	_ = "STUB: not implemented"
	return
}

func (c *multiTurnChatWithCallbacks) demonstrateOriginalRequestAccess(req *model.Request, resp *model.Response) {
	_ = "STUB: not implemented"
	// Only demonstrate when the response is complete (Done=true) to avoid multiple triggers during streaming.
	return
}

func (c *multiTurnChatWithCallbacks) shouldOverrideResponse(resp *model.Response) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *multiTurnChatWithCallbacks) createOverrideResponse() *model.Response {
	_ = "STUB: not implemented"
	return nil
}

func (c *multiTurnChatWithCallbacks) shouldReturnCustomToolResult(toolName string, jsonArgs []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *multiTurnChatWithCallbacks) createCustomCalculatorResult() calculatorResult {
	_ = "STUB: not implemented"
	return *new(calculatorResult)
}

func (c *multiTurnChatWithCallbacks) shouldFormatTimeResult(toolName string, _ any) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *multiTurnChatWithCallbacks) formatTimeResult(result any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (c *multiTurnChatWithCallbacks) extractResponseContent(invocation *agent.Invocation) string {
	_ = "STUB: not implemented"
	return ""
}
