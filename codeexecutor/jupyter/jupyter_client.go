//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package jupyter

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"trpc.group/trpc-go/trpc-agent-go/codeexecutor"
)

// ConnectionInfo ...
type ConnectionInfo struct {
	Host             string
	Port             int
	Token            string
	KernelName       string
	WaitReadyTimeout time.Duration
}

// Client ...
type Client struct {
	connectionInfo   ConnectionInfo
	baseURL          string
	httpClient       *http.Client
	kernelID         string
	ws               *websocket.Conn
	sessionID        string
	waitReadyTimeout time.Duration
}

// kernelSpec ...
type kernelSpec struct {
	Argv          []string `json:"argv"`
	DisplayName   string   `json:"display_name"`
	Language      string   `json:"language"`
	InterruptMode string   `json:"interrupt_mode"`
}

// kernelInfo ...
type kernelInfo struct {
	Name string     `json:"name"`
	Spec kernelSpec `json:"spec"`
}

// kernelSpecResponse ...
type kernelSpecResponse struct {
	Specs map[string]kernelInfo `json:"kernelspecs"`
}

// executionMessage ...
type executionMessage struct {
	Header struct {
		MsgType string `json:"msg_type"`
		MsgID   string `json:"msg_id"`
	} `json:"header"`
	Content      map[string]any `json:"content"`
	Metadata     map[string]any `json:"metadata"`
	ParentHeader struct {
		MsgID string `json:"msg_id"`
	} `json:"parent_header"`
}

// NewClient creates a new Jupyter client
func NewClient(connectionInfo ConnectionInfo) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CodeBlockDelimiter implements the CodeExecutor interface
func (c *Client) CodeBlockDelimiter() codeexecutor.CodeBlockDelimiter {
	_ = "STUB: not implemented"
	return *new(codeexecutor.CodeBlockDelimiter)
}

// ExecuteCode implements the CodeExecutor interface
func (c *Client) ExecuteCode(ctx context.Context, input codeexecutor.CodeExecutionInput) (codeexecutor.CodeExecutionResult, error) {
	_ = "STUB: not implemented"
	return *new(codeexecutor.CodeExecutionResult), nil
}

// jupyter executor does not support output files yet

// listKernelSpecs lists all available kernel specs
func (c *Client) listKernelSpecs() (kernelSpecResponse, error) {
	_ = "STUB: not implemented"
	return *new(kernelSpecResponse), nil
}

// startKernel starts a new kernel with the given name.
func (c *Client) startKernel(kernelName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Client) waitForReady() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// sendMessage sends a message to the kernel
func (c *Client) sendMessage(content map[string]any, channel string, messageType string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// runCode executes the given code, now only return text output
func (c *Client) runCode(code string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Close closes the client connection.
func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }
