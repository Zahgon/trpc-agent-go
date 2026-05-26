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

type displayImageInput struct{}

type displayImageOutput struct {
	Result string `json:"result"`
}

func displayImage(ctx context.Context, _ displayImageInput) (displayImageOutput, error) {
	_ = "STUB: not implemented"
	return *new(displayImageOutput), nil
}

var displayImageTool = function.NewFunctionTool(
	displayImage,
	function.WithName("display-image"),
	function.WithDescription("display image"),
)
