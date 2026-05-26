//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"context"
	"flag"
	"fmt"
	"net/http/httptest"
	"os"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

type ctxKey string

const userTokenKey ctxKey = "user-token"

var (
	mode  = flag.String("mode", "named", "Broker target mode: named or adhoc")
	token = flag.String("token", "demo-user-token", "User token injected through context for the successful path")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "auth hook example failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

func callBrokerTool(ctx context.Context, tools []tool.Tool, name string, args map[string]any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func printJSON(title string, value any) { _ = "STUB: not implemented"; return }

type protectedDemoServer struct {
	server *httptest.Server
	URL    string
}

func startProtectedDemoServer(requiredAuth string) *protectedDemoServer {
	_ = "STUB: not implemented"
	return nil
}

func (s *protectedDemoServer) Close() { _ = "STUB: not implemented"; return }
