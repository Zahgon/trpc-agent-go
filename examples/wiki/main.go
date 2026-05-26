//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates the Wikipedia search tool usage.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"trpc.group/trpc-go/trpc-agent-go/event"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	modelName  = flag.String("model", "deepseek-v4-flash", "Name of the model to use")
	language   = flag.String("lang", "en", "Wikipedia language (en, zh, es, etc.)")
	maxResults = flag.Int("maxresults", 3, "Maximum number of search results")
)

func main() {
	// Parse command line flags
	flag.Parse()

	fmt.Printf("🔍 Wikipedia Search Tool Demo\n")
	fmt.Printf("Model: %s\n", *modelName)
	fmt.Printf("Language: %s\n", *language)
	fmt.Println(strings.Repeat("=", 50))

	// Create and run the chat
	chat := &wikiChat{
		modelName:  *modelName,
		language:   *language,
		maxResults: *maxResults,
	}

	if err := chat.run(); err != nil {
		log.Fatalf("Chat failed: %v", err)
	}
}

type wikiChat struct {
	modelName  string
	language   string
	runner     runner.Runner
	userID     string
	sessionID  string
	maxResults int
}

func (c *wikiChat) run() error { _ = "STUB: not implemented"; return nil }

func (c *wikiChat) setup(_ context.Context) error {
	_ = "STUB: not implemented"
	// Create model instance
	return nil
}

// Create Wikipedia tool set

// Create generation config

// Create LLM agent

// Create runner

// Setup identifiers

// constructs the instruction for the LLM agent(felxibly change to fit your needs)
func (c *wikiChat) buildInstruction() string { _ = "STUB: not implemented"; return "" }

func (c *wikiChat) startChat(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Handle special commands

// Process the user message

func (c *wikiChat) printExamples() { _ = "STUB: not implemented"; return }

func (c *wikiChat) printToolInfo() { _ = "STUB: not implemented"; return }

func (c *wikiChat) processMessage(ctx context.Context, userMessage string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *wikiChat) processResponse(eventChan <-chan *event.Event) error {
	_ = "STUB: not implemented"
	return nil
}

// Detect tool calls

// 解析并显示搜索参数

// 显示语言参数
// 默认语言

// 构建Wikipedia搜索URL

// Handle tool results (when the tool returns data)

// 检查是否是工具返回的结果

// Handle content

// parseToolArguments parses tool call arguments JSON
func (c *wikiChat) parseToolArguments(args string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// displayToolResults displays formatted tool results
func (c *wikiChat) displayToolResults(content string) { _ = "STUB: not implemented"; return }

// 显示搜索结果摘要

// 显示找到的文章链接

// 截断描述以避免输出过长
// if len(desc) > 100 {
// 	desc = desc[:100] + "..."
// }

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }
