//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package langfuse

import (
	"context"
	"net/http"
	"time"
)

type client struct {
	baseURL    string
	publicKey  string
	secretKey  string
	httpClient *http.Client
}

type datasetRunItem struct {
	ID             string    `json:"id"`
	DatasetRunID   string    `json:"datasetRunId"`
	DatasetRunName string    `json:"datasetRunName"`
	DatasetItemID  string    `json:"datasetItemId"`
	TraceID        string    `json:"traceId"`
	CreatedAt      time.Time `json:"createdAt"`
}

type traceCreateRequest struct {
	ID          string         `json:"id"`
	Timestamp   time.Time      `json:"timestamp"`
	Name        string         `json:"name,omitempty"`
	Input       any            `json:"input,omitempty"`
	Output      any            `json:"output,omitempty"`
	SessionID   string         `json:"sessionId,omitempty"`
	UserID      string         `json:"userId,omitempty"`
	Environment string         `json:"environment,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	Public      *bool          `json:"public,omitempty"`
	Tags        []string       `json:"tags,omitempty"`
}

type datasetRunItemCreateRequest struct {
	RunName        string         `json:"runName"`
	RunDescription string         `json:"runDescription,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
	DatasetItemID  string         `json:"datasetItemId"`
	TraceID        string         `json:"traceId"`
}

type scoreCreateRequest struct {
	Name         string         `json:"name"`
	TraceID      string         `json:"traceId,omitempty"`
	DatasetRunID string         `json:"datasetRunId,omitempty"`
	Value        float64        `json:"value"`
	DataType     string         `json:"dataType"`
	Environment  string         `json:"environment,omitempty"`
	Comment      string         `json:"comment,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`
}

type identifierResponse struct {
	ID string `json:"id"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func newClient(baseURL, publicKey, secretKey string, httpClient *http.Client) *client {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) getDataset(ctx context.Context, datasetName string) (*dataset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) createTrace(ctx context.Context, req traceCreateRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) createDatasetRunItem(ctx context.Context, req datasetRunItemCreateRequest) (*datasetRunItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) createScore(ctx context.Context, req scoreCreateRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) doJSON(
	ctx context.Context,
	method string,
	path string,
	requestBody any,
	responseBody any,
) error {
	_ = "STUB: not implemented"
	return nil
}
