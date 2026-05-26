//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"time"

	"trpc.group/trpc-go/trpc-agent-go/model"
)

const appName = "agui-heartbeat-demo"

type serverConfig struct {
	ModelName         string
	GenerationConfig  model.GenerationConfig
	WaitDuration      time.Duration
	Address           string
	Path              string
	HeartbeatInterval time.Duration
}

func runServer(cfg serverConfig) { _ = "STUB: not implemented"; return }

func durationLabel(d time.Duration) string { _ = "STUB: not implemented"; return "" }
