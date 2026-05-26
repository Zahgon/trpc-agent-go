//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"
	"flag"
	"log"

	"trpc.group/trpc-go/trpc-a2a-go/client"
	"trpc.group/trpc-go/trpc-a2a-go/protocol"
)

const (
	defaultAgentURL = "http://localhost:8888/agents/math"
	defaultMessage  = "hello"
)

func main() {
	agentURL := flag.String("url", defaultAgentURL, "A2A agent base URL")
	message := flag.String("msg", defaultMessage, "User message")
	flag.Parse()

	c, err := client.NewA2AClient(*agentURL)
	if err != nil {
		log.Fatalf("create client: %v", err)
	}

	msg := protocol.NewMessage(
		protocol.MessageRoleUser,
		[]protocol.Part{
			protocol.NewTextPart(*message),
		},
	)

	rsp, err := c.SendMessage(
		context.Background(),
		protocol.SendMessageParams{Message: msg},
	)
	if err != nil {
		log.Fatalf("send message: %v", err)
	}
	printMessageResult(rsp)
}

func printMessageResult(rsp *protocol.MessageResult) { _ = "STUB: not implemented"; return }

func printTask(task *protocol.Task) { _ = "STUB: not implemented"; return }

func printTextParts(parts []protocol.Part) { _ = "STUB: not implemented"; return }

func asTextPart(part protocol.Part) *protocol.TextPart { _ = "STUB: not implemented"; return nil }
