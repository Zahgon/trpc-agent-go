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

	coreevaluation "trpc.group/trpc-go/trpc-agent-go/evaluation"
)

func (s *Server) handleRuns(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleCreateRun(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) runEvaluation(ctx context.Context, req *RunEvaluationRequest) (*coreevaluation.EvaluationResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) decodeRunEvaluationRequest(w http.ResponseWriter, r *http.Request) (*RunEvaluationRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateRunEvaluationRequest(req *RunEvaluationRequest) error {
	_ = "STUB: not implemented"
	return nil
}
