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

	"trpc.group/trpc-go/trpc-agent-go/tool/function"
)

type generateImageInput struct {
	Prompt string `json:"prompt"`
}

type generateImageOutput struct {
	Result string `json:"result"`
}

func generateImage(ctx context.Context, input generateImageInput) (generateImageOutput, error) {
	_ = "STUB: not implemented"
	return *new(generateImageOutput), nil
}

type image struct {
	mimeType string
	content  []byte
	url      string
}

func mockLLMGenerateImage(ctx context.Context, prompt string) ([]image, error) {
	_ = "STUB: not implemented"
	// Get the current directory
	return nil, nil
}

// Resource directory path

// Read directory contents

// Process each image file

// Check if it's an image file

// Read file content

// Determine MIME type based on file extension

// default to PNG

// Create image struct

var generateImageTool = function.NewFunctionTool(
	generateImage,
	function.WithName("text-to-image"),
	function.WithDescription("generate image by input text"),
)
