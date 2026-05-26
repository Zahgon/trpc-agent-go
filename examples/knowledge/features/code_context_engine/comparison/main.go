//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main compares two agent setups on the same repository questions:
// one agent uses the local code_search tool, and the other uses Augment code
// search through MCP. Each run records the final answer and intermediate tool
// calls/results to a markdown report.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	util "trpc.group/trpc-go/trpc-agent-go/examples/knowledge"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

const (
	defaultOutputDir      = "output/code_context_engine"
	repoURL               = "https://github.com/trpc-group/trpc-agent-go"
	repoName              = "trpc-agent-go"
	repoBranch            = "main"
	augmentServerURL      = "https://api.augmentcode.com/mcp"
	augmentToolName       = "augment_code_search"
	augmentAuthHeader     = "Authorization"
	defaultModelName      = "gpt-5.4"
	localUserID           = "code-context-local-user"
	augmentUserID         = "code-context-augment-user"
	defaultLocalMCPURL    = "http://localhost:3001/mcp"
	defaultLocalMCPToolNm = "code_search"
)

// Run modes selectable via the -mode flag.
const (
	runModeBoth    = "both"
	runModeLocal   = "local"
	runModeAugment = "augment"
)

var (
	runMode          = flag.String("mode", runModeBoth, "Which agent(s) to run: local | augment | both")
	outputDir        = flag.String("output-dir", defaultOutputDir, "Directory where per-case markdown reports are written")
	localMCPURL      = flag.String("local-mcp-url", defaultLocalMCPURL, "URL of our own code-search MCP server used by the local agent")
	localMCPToolName = flag.String("local-mcp-tool", defaultLocalMCPToolNm, "Name of the MCP tool exposed by our local code-search MCP server")
)

type comparisonCase struct {
	Name        string
	Description string
	Prompt      string
}

type toolCallTrace struct {
	CallID    string `json:"call_id"`
	ToolName  string `json:"tool_name"`
	Arguments string `json:"arguments"`
}

type toolResultTrace struct {
	CallID   string `json:"call_id"`
	ToolName string `json:"tool_name"`
	Content  string `json:"content"`
}

type agentRunResult struct {
	FinalAnswer string            `json:"final_answer"`
	ToolCalls   []toolCallTrace   `json:"tool_calls,omitempty"`
	ToolResults []toolResultTrace `json:"tool_results,omitempty"`
}

func main() {
	flag.Parse()
	ctx := context.Background()

	mode := strings.ToLower(strings.TrimSpace(*runMode))
	switch mode {
	case runModeBoth, runModeLocal, runModeAugment:
	default:
		log.Fatalf("invalid -mode %q: expected one of local|augment|both", *runMode)
	}
	runLocal := mode == runModeBoth || mode == runModeLocal
	runAugment := mode == runModeBoth || mode == runModeAugment

	reportDir := strings.TrimSpace(*outputDir)
	if reportDir == "" {
		reportDir = defaultOutputDir
	}
	if err := os.MkdirAll(reportDir, 0o755); err != nil {
		log.Fatalf("failed to create comparison output dir %q: %v", reportDir, err)
	}
	modelName := strings.TrimSpace(util.GetEnvOrDefault("MODEL_NAME", defaultModelName))

	fmt.Println("Code Context Engine Agent Comparison")
	fmt.Println("====================================")
	fmt.Printf("Repository URL: %s\n", repoURL)
	fmt.Printf("Repository Name: %s\n", repoName)
	fmt.Printf("Repository Branch: %s\n", repoBranch)
	fmt.Printf("Agent Model: %s\n", modelName)
	fmt.Printf("Local MCP URL: %s\n", *localMCPURL)
	fmt.Printf("Local MCP Tool: %s\n", *localMCPToolName)
	fmt.Printf("Run Mode: %s\n", mode)
	fmt.Printf("Output Dir: %s\n", reportDir)

	var localAgent *localCodeSearchAgentRunner
	if runLocal {
		la, err := newLocalCodeSearchAgentRunner(localAgentConfig{
			ModelName:    modelName,
			MCPServerURL: *localMCPURL,
			MCPToolName:  *localMCPToolName,
		})
		if err != nil {
			log.Fatalf("create local agent runner failed: %v", err)
		}
		localAgent = la
		defer func() {
			if err := localAgent.Close(); err != nil {
				log.Printf("close local agent runner failed: %v", err)
			}
		}()
	}

	var (
		augmentAgent   *augmentCodeAgentRunner
		augmentInitErr error
	)
	if runAugment {
		augmentAgent, augmentInitErr = newAugmentCodeAgentRunner(modelName)
		if augmentInitErr != nil {
			log.Printf("create augment agent runner failed: %v", augmentInitErr)
		} else {
			defer func() {
				if err := augmentAgent.Close(); err != nil {
					log.Printf("close augment agent runner failed: %v", err)
				}
			}()
		}
	}

	cases := defaultCases()
	fmt.Println("\nStep 2: Per-question agent comparison")
	fmt.Println("------------------------------------")
	for _, c := range cases {
		printCaseHeader(c)

		var (
			localResult   *agentRunResult
			localErr      error
			augmentResult *agentRunResult
			augmentErr    error
		)

		if runLocal {
			localResult, localErr = localAgent.RunCase(ctx, c)
			printAgentSummary("Local agent + code_search", localResult, localErr)
		}

		if runAugment {
			if augmentAgent != nil {
				augmentResult, augmentErr = augmentAgent.RunCase(ctx, c)
				printAgentSummary("Augment agent + MCP", augmentResult, augmentErr)
			} else {
				augmentErr = fmt.Errorf("augment agent runner not initialized: %w", augmentInitErr)
				printAgentSummary("Augment agent + MCP", nil, augmentErr)
			}
		}

		reportPath, err := writeCaseReport(reportDir, mode, c, localResult, localErr, augmentResult, augmentErr)
		if err != nil {
			log.Printf("write comparison report for case %s failed: %v", c.Name, err)
			continue
		}
		fmt.Printf("Report written: %s\n", reportPath)
	}
}

func newAgentModel(modelName string) model.Model {
	_ = "STUB: not implemented"
	return *new(model.Model)
}

func defaultCases() []comparisonCase { _ = "STUB: not implemented"; return nil }

func runAgentInvocation(
	ctx context.Context,
	r runner.Runner,
	userID string,
	sessionID string,
	prompt string,
) (*agentRunResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func printCaseHeader(c comparisonCase) { _ = "STUB: not implemented"; return }

func printAgentSummary(label string, result *agentRunResult, err error) {
	_ = "STUB: not implemented"
	return
}

func writeCaseReport(
	dir string,
	mode string,
	c comparisonCase,
	localResult *agentRunResult,
	localErr error,
	augmentResult *agentRunResult,
	augmentErr error,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func renderAgentReport(result *agentRunResult, err error) string {
	_ = "STUB: not implemented"
	return ""
}

func compactText(text string, limit int) string { _ = "STUB: not implemented"; return "" }

func prettyToolPayload(raw string) string { _ = "STUB: not implemented"; return "" }

func prettyJSON(value any) string { _ = "STUB: not implemented"; return "" }
