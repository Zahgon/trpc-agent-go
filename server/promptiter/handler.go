//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package promptiter

import (
	"context"
	"net/http"
	"time"

	astructure "trpc.group/trpc-go/trpc-agent-go/agent/structure"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter"
	"trpc.group/trpc-go/trpc-agent-go/evaluation/workflow/promptiter/engine"
)

func (s *Server) handleStructure(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleRuns(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleAsyncRuns(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleRunResource(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleCORS(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func (s *Server) redirectTrailingSlashToCanonicalPath(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) respondJSON(w http.ResponseWriter, r *http.Request, statusCode int, payload any) {
	_ = "STUB: not implemented"
	return
}

func statusCodeFromError(err error) int { _ = "STUB: not implemented"; return 0 }

func (s *Server) handleRunByID(w http.ResponseWriter, r *http.Request, runID string) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) runResponse(run *engine.RunResult) *RunResponse {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) handleCancelRun(w http.ResponseWriter, r *http.Request, runID string) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) validateTargetSurfaceIDs(ctx context.Context, targetSurfaceIDs []string) error {
	_ = "STUB: not implemented"
	return nil
}

func isSupportedTargetSurfaceType(surfaceType astructure.SurfaceType) bool {
	_ = "STUB: not implemented"
	return false
}

func newExecutionContext(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func (s *Server) decodeRunRequest(w http.ResponseWriter, r *http.Request) (*RunRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeJSONBody[T any](
	w http.ResponseWriter,
	r *http.Request,
	respond func(w http.ResponseWriter, r *http.Request, statusCode int, payload any),
) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func decodeJSONPayload[T any](body []byte, disallowUnknownFields bool) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func validateRunRequest(req *RunRequest) error { _ = "STUB: not implemented"; return nil }

func validateEngineRunRequest(request *engine.RunRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func validateEvalSetInputs(role string, inputs []engine.EvalSetInput) error {
	_ = "STUB: not implemented"
	return nil
}

func isValidLossHintSeverity(severity promptiter.LossSeverity) bool {
	_ = "STUB: not implemented"
	return false
}
