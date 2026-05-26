//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package evaluation

import (
	"context"
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalresult"
)

func (s *Server) handleResults(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleResultByID(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) listEvalResults(ctx context.Context, filterSetID string) ([]*evalresult.EvalSetResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readSetIDFilter(r *http.Request) string { _ = "STUB: not implemented"; return "" }
