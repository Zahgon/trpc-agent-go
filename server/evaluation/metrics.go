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

	"trpc.group/trpc-go/trpc-agent-go/evaluation/metric"
)

func (s *Server) handleMetrics(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleMetricByName(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleListMetrics(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) listMetrics(ctx context.Context, setID string) ([]*metric.EvalMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) handleCreateMetric(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleGetMetric(w http.ResponseWriter, r *http.Request, metricName string) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleUpdateMetric(w http.ResponseWriter, r *http.Request, metricName string) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) handleDeleteMetric(w http.ResponseWriter, r *http.Request, metricName string) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) readRequiredMetricSetID(w http.ResponseWriter, r *http.Request) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (s *Server) decodeCreateMetricRequest(w http.ResponseWriter, r *http.Request) (*CreateMetricRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) decodeUpdateMetricRequest(w http.ResponseWriter, r *http.Request, metricName string) (*UpdateMetricRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateCreateMetricRequest(req *CreateMetricRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func validateUpdateMetricRequest(req *UpdateMetricRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMetricPayload(setID *string, evalMetric *metric.EvalMetric) error {
	_ = "STUB: not implemented"
	return nil
}

func applyMetricNameFromPath(evalMetric *metric.EvalMetric, metricName string) error {
	_ = "STUB: not implemented"
	return nil
}
