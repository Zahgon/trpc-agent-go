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
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/client/sse"
	"github.com/ag-ui-protocol/ag-ui/sdks/community/go/pkg/core/events"
)

const (
	defaultEndpoint    = "http://127.0.0.1:8080/agui"
	requestTimeout     = 2 * time.Minute
	connectTimeout     = 30 * time.Second
	readTimeout        = 5 * time.Minute
	streamBufferSize   = 100
	stdinBufferInitial = 64 * 1024
	stdinBufferMax     = 1 << 20
)

func main() {
	endpoint := flag.String("endpoint", defaultEndpoint, "AG-UI SSE endpoint")
	flag.Parse()

	if err := runInteractive(*endpoint); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func runInteractive(endpoint string) error { _ = "STUB: not implemented"; return nil }

func streamConversation(endpoint, prompt string) error { _ = "STUB: not implemented"; return nil }

func newSSEClient(endpoint string) *sse.Client { _ = "STUB: not implemented"; return nil }

func formatEvent(evt events.Event) []string { _ = "STUB: not implemented"; return nil }
