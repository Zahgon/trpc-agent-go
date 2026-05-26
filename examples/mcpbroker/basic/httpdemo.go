//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package main

import (
	"net/http/httptest"
)

const remoteSkillName = "remote-http-mcp"

type remoteHTTPDemo struct {
	server *httptest.Server
	url    string
}

func startRemoteHTTPDemoServer() *remoteHTTPDemo { _ = "STUB: not implemented"; return nil }

func (d *remoteHTTPDemo) Close() { _ = "STUB: not implemented"; return }
