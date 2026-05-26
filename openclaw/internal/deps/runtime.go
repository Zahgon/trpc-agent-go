//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package deps

import (
	"errors"
)

const (
	defaultToolchainDir = "toolchain"
	defaultPythonEnvDir = "python"

	envPath                = "PATH"
	envVirtualEnv          = "VIRTUAL_ENV"
	envOpenClawToolchain   = "OPENCLAW_TOOLCHAIN_ROOT"
	envOpenClawPython      = "OPENCLAW_TOOLCHAIN_PYTHON"
	envPipDisableVersion   = "PIP_DISABLE_PIP_VERSION_CHECK"
	pipDisableVersionValue = "1"
)

type Platform struct {
	GOOS           string `json:"goos"`
	GOARCH         string `json:"goarch"`
	PackageManager string `json:"package_manager,omitempty"`
}

type PythonRuntime struct {
	Found     bool   `json:"found"`
	Path      string `json:"path,omitempty"`
	Version   string `json:"version,omitempty"`
	Managed   bool   `json:"managed"`
	EnvRoot   string `json:"env_root,omitempty"`
	Bootstrap string `json:"bootstrap,omitempty"`
}

type Toolchain struct {
	StateDir string        `json:"state_dir,omitempty"`
	Root     string        `json:"root,omitempty"`
	BinDir   string        `json:"bin_dir,omitempty"`
	Active   bool          `json:"active"`
	Python   PythonRuntime `json:"python"`
}

type BinStatus struct {
	Name  string `json:"name"`
	Found bool   `json:"found"`
	Path  string `json:"path,omitempty"`
}

type AnyBinStatus struct {
	Names     []string    `json:"names"`
	Satisfied bool        `json:"satisfied"`
	Found     []BinStatus `json:"found,omitempty"`
}

type PythonStatus struct {
	Module  string `json:"module"`
	Package string `json:"package,omitempty"`
	Found   bool   `json:"found"`
}

type SourceReport struct {
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Bins        []BinStatus     `json:"bins,omitempty"`
	AnyBins     []AnyBinStatus  `json:"any_bins,omitempty"`
	Python      []PythonStatus  `json:"python,omitempty"`
	Install     []InstallAction `json:"install,omitempty"`
}

type Missing struct {
	Bins    []string        `json:"bins,omitempty"`
	AnyBins [][]string      `json:"any_bins,omitempty"`
	Python  []PythonPackage `json:"python,omitempty"`
}

type Report struct {
	Platform  Platform       `json:"platform"`
	Toolchain Toolchain      `json:"toolchain"`
	Sources   []SourceReport `json:"sources,omitempty"`
	Missing   Missing        `json:"missing,omitempty"`
}

func Inspect(stateDir string, sources []Source) (Report, error) {
	_ = "STUB: not implemented"
	return *new(Report), nil
}

func InspectStartup(
	stateDir string,
	sources []Source,
) (Report, error) {
	_ = "STUB: not implemented"
	return *new(Report), nil
}

func inspect(
	stateDir string,
	sources []Source,
	includePython bool,
) (Report, error) {
	_ = "STUB: not implemented"
	return *new(Report), nil
}

func InspectProfiles(
	stateDir string,
	profiles []string,
) (Report, error) {
	_ = "STUB: not implemented"
	return *new(Report), nil
}

func DetectToolchain(stateDir string) Toolchain { _ = "STUB: not implemented"; return *new(Toolchain) }

func ManagedToolchainRoot(stateDir string) string { _ = "STUB: not implemented"; return "" }

func ManagedPythonRoot(stateDir string) string { _ = "STUB: not implemented"; return "" }

// ManagedToolPrefix returns the shared install prefix used for managed
// tool binaries. Python, npm, and other managed CLIs intentionally
// share this prefix so one PATH entry can expose all managed tools.
func ManagedToolPrefix(stateDir string) string { _ = "STUB: not implemented"; return "" }

func ManagedBinDir(stateDir string) string { _ = "STUB: not implemented"; return "" }

func ManagedPythonCandidates(stateDir string) []string { _ = "STUB: not implemented"; return nil }

func ToolEnv(stateDir string) map[string]string { _ = "STUB: not implemented"; return nil }

func FindPythonRuntime(stateDir string) PythonRuntime {
	_ = "STUB: not implemented"
	return *new(PythonRuntime)
}

func DetectPackageManager() string { _ = "STUB: not implemented"; return "" }

func CheckPythonPackages(
	python PythonRuntime,
	pkgs []PythonPackage,
) ([]PythonStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inspectSource(
	toolchain Toolchain,
	source Source,
	includePython bool,
) (SourceReport, Missing, error) {
	_ = "STUB: not implemented"
	return *new(SourceReport), *new(Missing), nil
}

func checkBin(toolchain Toolchain, name string) BinStatus {
	_ = "STUB: not implemented"
	return *new(BinStatus)
}

func lookupManagedBin(
	toolchain Toolchain,
	name string,
) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func managedBinCandidates(name string) []string { _ = "STUB: not implemented"; return nil }

func mergeMissing(left, right Missing) Missing { _ = "STUB: not implemented"; return *new(Missing) }

func normalizeMissing(m Missing) Missing { _ = "STUB: not implemented"; return *new(Missing) }

func HasMissing(report Report) bool { _ = "STUB: not implemented"; return false }

func pythonVersion(path string) string { _ = "STUB: not implemented"; return "" }

func systemPythonCandidate() string { _ = "STUB: not implemented"; return "" }

func prependPath(prefix string, current string) string { _ = "STUB: not implemented"; return "" }

func dirExists(path string) bool { _ = "STUB: not implemented"; return false }

func fileExists(path string) bool { _ = "STUB: not implemented"; return false }

var errPythonNotFound = errors.New("python interpreter not found")
