//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main demonstrates a client that connects to the EventEmitter server example
// and displays custom events, progress events, and streaming text events.
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
	defaultEndpoint  = "http://127.0.0.1:8080/agui"
	requestTimeout   = 2 * time.Minute
	connectTimeout   = 30 * time.Second
	readTimeout      = 5 * time.Minute
	streamBufferSize = 100
)

var (
	endpoint = flag.String("endpoint", defaultEndpoint, "AG-UI SSE endpoint")
	prompt   = flag.String("prompt", "process my data", "User prompt to send")
)

func main() {
	flag.Parse()

	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║       EventEmitter Client - Node Custom Events Demo          ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
	fmt.Printf("\n📡 Connecting to: %s\n", *endpoint)
	fmt.Printf("📝 Sending prompt: %q\n\n", *prompt)

	if err := runDemo(*endpoint, *prompt); err != nil {
		fmt.Fprintf(os.Stderr, "\n❌ Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ Demo completed successfully!")
}

func runDemo(endpoint, prompt string) error { _ = "STUB: not implemented"; return nil }

func newSSEClient(endpoint string) *sse.Client { _ = "STUB: not implemented"; return nil }

func displayEvent(evt events.Event) { _ = "STUB: not implemented"; return }

func displayCustomEvent(e *events.CustomEvent) {
	_ = "STUB: not implemented"
	// Parse the custom event based on its name
	return
}

func displayWorkflowEvent(e *events.CustomEvent) { _ = "STUB: not implemented"; return }

func displayProgressEvent(e *events.CustomEvent) { _ = "STUB: not implemented"; return }

// Create progress bar

// New line when complete

func displayTextEvent(e *events.CustomEvent) { _ = "STUB: not implemented"; return }

func displayGenericCustomEvent(e *events.CustomEvent) { _ = "STUB: not implemented"; return }
