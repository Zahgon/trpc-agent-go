//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package app

import (
	"sync"

	"gopkg.in/yaml.v3"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/admin"
)

const (
	adminRuntimeConfigInputText   = "text"
	adminRuntimeConfigInputNumber = "number"
	adminRuntimeConfigInputSelect = "select"

	adminRuntimeConfigApplyRestart = "restart"

	adminRuntimeConfigSourceExplicit  = "explicit"
	adminRuntimeConfigSourceInherited = "inherited"

	adminRuntimeConfigConfiguredExplicit  = "Explicit in config"
	adminRuntimeConfigConfiguredInherited = "Inherited from current runtime"
	adminRuntimeConfigRuntimeSourceLabel  = "Current runtime"

	adminRuntimeConfigValueString = "string"
	adminRuntimeConfigValueInt    = "int"
	adminRuntimeConfigValueBool   = "bool"
)

var runtimeAdminOptionsStore sync.Map

// AdminSourceConfigPathEnvName overrides the writable admin config file path.
const AdminSourceConfigPathEnvName = "TRPC_CLAW_ADMIN_SOURCE_CONFIG_PATH"

type adminRuntimeConfigProvider struct {
	configPath string
	opts       runOptions
}

type adminRuntimeConfigKeyRef struct {
	Preferred string
	Aliases   []string
}

type adminRuntimeConfigSectionSpec struct {
	Key     string
	Title   string
	Summary string
	Fields  []adminRuntimeConfigFieldSpec
}

type adminRuntimeConfigFieldSpec struct {
	Key         string
	Title       string
	Summary     string
	InputType   string
	Placeholder string
	ApplyMode   string
	ValueType   string
	Path        []adminRuntimeConfigKeyRef
	Options     []admin.RuntimeConfigOption
	Runtime     func(runOptions) string
}

type adminRuntimeConfiguredValue struct {
	Value    string
	Explicit bool
}

func buildAdminRuntimeConfigProvider(
	opts runOptions,
) admin.RuntimeConfigProvider {
	_ = "STUB: not implemented"
	return *new(admin.RuntimeConfigProvider)
}

func adminWritableConfigPath(configPath string) string { _ = "STUB: not implemented"; return "" }

func buildAdminOptions(opts runOptions) []admin.Option { _ = "STUB: not implemented"; return nil }

func runtimeAdminOptions(rt *Runtime) []admin.Option { _ = "STUB: not implemented"; return nil }

func setRuntimeAdminOptions(rt *Runtime, opts []admin.Option) { _ = "STUB: not implemented"; return }

func clearRuntimeAdminOptions(rt *Runtime) { _ = "STUB: not implemented"; return }

func (p *adminRuntimeConfigProvider) RuntimeConfigStatus() (
	admin.RuntimeConfigStatus,
	error,
) {
	_ = "STUB: not implemented"
	return *new(admin.RuntimeConfigStatus), nil
}

func (p *adminRuntimeConfigProvider) SaveRuntimeConfigValue(
	key string,
	value string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *adminRuntimeConfigProvider) ResetRuntimeConfigValue(
	key string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func adminRuntimeConfigSectionSpecs() []adminRuntimeConfigSectionSpec {
	_ = "STUB: not implemented"
	return nil
}

func adminRuntimeBoolField(
	key string,
	title string,
	summary string,
	path []adminRuntimeConfigKeyRef,
	runtime func(runOptions) string,
) adminRuntimeConfigFieldSpec {
	_ = "STUB: not implemented"
	return *new(adminRuntimeConfigFieldSpec)
}

func adminRuntimeNumberField(
	key string,
	title string,
	summary string,
	path []adminRuntimeConfigKeyRef,
	runtime func(runOptions) string,
) adminRuntimeConfigFieldSpec {
	_ = "STUB: not implemented"
	return *new(adminRuntimeConfigFieldSpec)
}

func adminRuntimeTextField(
	key string,
	title string,
	summary string,
	placeholder string,
	path []adminRuntimeConfigKeyRef,
	runtime func(runOptions) string,
) adminRuntimeConfigFieldSpec {
	_ = "STUB: not implemented"
	return *new(adminRuntimeConfigFieldSpec)
}

func adminRuntimeSelectField(
	key string,
	title string,
	summary string,
	path []adminRuntimeConfigKeyRef,
	runtime func(runOptions) string,
	options ...string,
) adminRuntimeConfigFieldSpec {
	_ = "STUB: not implemented"
	return *new(adminRuntimeConfigFieldSpec)
}

func adminRuntimeStringOptions(
	values ...string,
) []admin.RuntimeConfigOption {
	_ = "STUB: not implemented"
	return nil
}

func adminRuntimeConfiguredSource(explicit bool) string { _ = "STUB: not implemented"; return "" }

func adminRuntimeConfiguredLabel(explicit bool) string { _ = "STUB: not implemented"; return "" }

func adminRuntimeComparableConfiguredValue(
	spec adminRuntimeConfigFieldSpec,
	value string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func adminRuntimeConfigFieldSpecByKey(
	key string,
) (adminRuntimeConfigFieldSpec, bool) {
	_ = "STUB: not implemented"
	return *new(adminRuntimeConfigFieldSpec), false
}

func adminRuntimeConfigRootFromPath(
	path string,
) (*yaml.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func adminRuntimeConfigDocumentFromPath(
	path string,
) (yaml.Node, *yaml.Node, error) {
	_ = "STUB: not implemented"
	return *new(yaml.Node), nil, nil
}

func adminRuntimeConfiguredFieldValue(
	root *yaml.Node,
	path []adminRuntimeConfigKeyRef,
) adminRuntimeConfiguredValue {
	_ = "STUB: not implemented"
	return *new(adminRuntimeConfiguredValue)
}

func adminRuntimeFieldNode(
	root *yaml.Node,
	path []adminRuntimeConfigKeyRef,
) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

func adminRuntimeFieldParent(
	root *yaml.Node,
	path []adminRuntimeConfigKeyRef,
) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

func adminRuntimeEnsureFieldParent(
	root *yaml.Node,
	path []adminRuntimeConfigKeyRef,
) (*yaml.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func adminRuntimeDeleteField(
	root *yaml.Node,
	path []adminRuntimeConfigKeyRef,
) {
	_ = "STUB: not implemented"
	return
}

func adminRuntimeSetFieldValue(
	parent *yaml.Node,
	spec adminRuntimeConfigFieldSpec,
	raw string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func adminRuntimeOptionValueAllowed(
	value string,
	options []admin.RuntimeConfigOption,
) bool {
	_ = "STUB: not implemented"
	return false
}

func adminRuntimeSetMappingBool(
	parent *yaml.Node,
	key adminRuntimeConfigKeyRef,
	value bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

func adminRuntimeSetMappingInt(
	parent *yaml.Node,
	key adminRuntimeConfigKeyRef,
	value int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func adminRuntimeSetMappingString(
	parent *yaml.Node,
	key adminRuntimeConfigKeyRef,
	value string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func adminRuntimeSetScalarValue(
	parent *yaml.Node,
	key adminRuntimeConfigKeyRef,
	tag string,
	value string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func adminRuntimeEnsureMappingChild(
	parent *yaml.Node,
	key adminRuntimeConfigKeyRef,
) (*yaml.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func adminRuntimeLookupMappingValue(
	parent *yaml.Node,
	key adminRuntimeConfigKeyRef,
) *yaml.Node {
	_ = "STUB: not implemented"
	return nil
}

func adminRuntimeDeleteMappingValue(
	parent *yaml.Node,
	key adminRuntimeConfigKeyRef,
) {
	_ = "STUB: not implemented"
	return
}

func adminRuntimeKey(
	preferred string,
	aliases ...string,
) adminRuntimeConfigKeyRef {
	_ = "STUB: not implemented"
	return *new(adminRuntimeConfigKeyRef)
}

func adminRuntimePreferredKey(key adminRuntimeConfigKeyRef) string {
	_ = "STUB: not implemented"
	return ""
}

func adminRuntimeKeyMatches(
	node *yaml.Node,
	key adminRuntimeConfigKeyRef,
) bool {
	_ = "STUB: not implemented"
	return false
}
