//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package mem0

import (
	"trpc.group/trpc-go/trpc-agent-go/memory"
	memorytool "trpc.group/trpc-go/trpc-agent-go/memory/tool"
	"trpc.group/trpc-go/trpc-agent-go/tool"
)

func buildReadOnlyTools(s *Service) []tool.Tool { _ = "STUB: not implemented"; return nil }

func newSearchTool(s *Service) tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

func newLoadTool(s *Service) tool.Tool { _ = "STUB: not implemented"; return *new(tool.Tool) }

func buildToolSearchOptions(req *memorytool.SearchMemoryRequest) (memory.SearchOptions, error) {
	_ = "STUB: not implemented"
	return *new(memory.SearchOptions), nil
}

func entryToResult(e *memory.Entry) memorytool.Result {
	_ = "STUB: not implemented"
	return *new(memorytool.Result)
}
