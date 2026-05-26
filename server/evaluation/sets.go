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

	"trpc.group/trpc-go/trpc-agent-go/evaluation/evalset"
)

func (s *Server) handleSets(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleSetByID(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) listEvalSets(ctx context.Context) ([]*evalset.EvalSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
