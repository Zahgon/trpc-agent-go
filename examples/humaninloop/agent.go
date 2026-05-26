//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main provides agent implementations for human-in-the-loop scenarios.
package main

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/agent/llmagent"
)

func reimburse(_ context.Context, _ reimburseInput) (reimburseOutput, error) {
	_ = "STUB: not implemented"
	return *new(reimburseOutput), nil
}

func askForApproval(_ context.Context, i askForApprovalInput) (askForApprovalOutput, error) {
	_ = "STUB: not implemented"
	return *new(askForApprovalOutput), nil
}

type reimburseInput struct {
	Purpose string `json:"purpose"`
	Amount  int    `json:"amount"`
}

type reimburseOutput struct {
	Status string `json:"status"`
}

type askForApprovalInput struct {
	Purpose string `json:"purpose"`
	Amount  int    `json:"amount"`
}

type askForApprovalOutput struct {
	Status   string `json:"status"`
	Amount   int    `json:"amount"`
	TicketID string `json:"ticket_id"`
}

func newLLMAgent(modelName string, streaming bool) *llmagent.LLMAgent {
	_ = "STUB: not implemented"
	return nil
}

// Enable or disable streaming.
