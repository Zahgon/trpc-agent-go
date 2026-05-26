//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package admin

import (
	"net/http"
)

const (
	routeConfigPage  = "/config"
	routeConfigJSON  = "/api/config"
	routeConfigSave  = "/api/config/save"
	routeConfigReset = "/api/config/reset"

	formConfigFieldKey = "field_key"
	formConfigValue    = "value"

	configInputText     = "text"
	configInputNumber   = "number"
	configInputSelect   = "select"
	configInputReadOnly = "readonly"

	configApplyHot        = "hot"
	configApplyNextTurn   = "next_turn"
	configApplyRestart    = "restart"
	configSourceExplicit  = "explicit"
	configSourceInherited = "inherited"
)

const pageSummaryConfig = "" +
	"Edit major runtime config fields, compare the saved config with " +
	"the current runtime, and reset fields back to inheritance."

const (
	viewConfig adminView = "config"
)

type RuntimeConfigProvider interface {
	RuntimeConfigStatus() (RuntimeConfigStatus, error)
	SaveRuntimeConfigValue(key string, value string) error
	ResetRuntimeConfigValue(key string) error
}

type RuntimeConfigStatus struct {
	Enabled    bool                   `json:"enabled"`
	Error      string                 `json:"error,omitempty"`
	ConfigPath string                 `json:"config_path,omitempty"`
	Sections   []RuntimeConfigSection `json:"sections,omitempty"`
}

type RuntimeConfigSection struct {
	Key     string               `json:"key,omitempty"`
	Title   string               `json:"title,omitempty"`
	Summary string               `json:"summary,omitempty"`
	Fields  []RuntimeConfigField `json:"fields,omitempty"`
}

type RuntimeConfigField struct {
	Key                   string                `json:"key,omitempty"`
	Title                 string                `json:"title,omitempty"`
	Summary               string                `json:"summary,omitempty"`
	InputType             string                `json:"input_type,omitempty"`
	Placeholder           string                `json:"placeholder,omitempty"`
	ApplyMode             string                `json:"apply_mode,omitempty"`
	EditorValue           string                `json:"editor_value,omitempty"`
	ConfiguredValue       string                `json:"configured_value,omitempty"`
	ConfiguredSourceLabel string                `json:"configured_source_label,omitempty"`
	ConfiguredSource      string                `json:"configured_source,omitempty"`
	RuntimeValue          string                `json:"runtime_value,omitempty"`
	RuntimeSourceLabel    string                `json:"runtime_source_label,omitempty"`
	PendingRestart        bool                  `json:"pending_restart"`
	Resettable            bool                  `json:"resettable"`
	Options               []RuntimeConfigOption `json:"options,omitempty"`
}

type RuntimeConfigOption struct {
	Value string `json:"value,omitempty"`
	Label string `json:"label,omitempty"`
}

func (s *Service) runtimeConfigStatus() RuntimeConfigStatus {
	_ = "STUB: not implemented"
	return *new(RuntimeConfigStatus)
}

func (s *Service) runtimeConfigProvider() RuntimeConfigProvider {
	_ = "STUB: not implemented"
	return *new(RuntimeConfigProvider)
}

func (s *Service) hasRuntimeConfigProvider() bool { _ = "STUB: not implemented"; return false }

func (s *Service) handleConfigPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleConfigJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleSaveRuntimeConfig(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleResetRuntimeConfig(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}
