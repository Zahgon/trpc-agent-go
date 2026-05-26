//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package deps

const (
	ProfilePDF             = "pdf"
	ProfileOffice          = "office"
	ProfileAudio           = "audio"
	ProfileVideo           = "video"
	ProfileImage           = "image"
	ProfileOCR             = "ocr"
	ProfileCommonFileTools = "common-file-tools"
)

const (
	InstallKindAPT      = "apt"
	InstallKindBrew     = "brew"
	InstallKindDownload = "download"
	InstallKindGo       = "go"
	InstallKindDNF      = "dnf"
	InstallKindNode     = "node"
	InstallKindNPM      = "npm"
	InstallKindPIP      = "pip"
	InstallKindUV       = "uv"
	InstallKindYUM      = "yum"
)

type PythonPackage struct {
	Module  string `yaml:"module,omitempty" json:"module,omitempty"`
	Package string `yaml:"package,omitempty" json:"package,omitempty"`
	Label   string `yaml:"label,omitempty" json:"label,omitempty"`
}

type Requirement struct {
	Bins    []string        `yaml:"bins,omitempty" json:"bins,omitempty"`
	AnyBins []string        `yaml:"anyBins,omitempty" json:"any_bins,omitempty"`
	Env     []string        `yaml:"env,omitempty" json:"env,omitempty"`
	Config  []string        `yaml:"config,omitempty" json:"config,omitempty"`
	Python  []PythonPackage `yaml:"python,omitempty" json:"python,omitempty"`
}

type InstallAction struct {
	ID              string   `yaml:"id,omitempty" json:"id,omitempty"`
	Kind            string   `yaml:"kind,omitempty" json:"kind,omitempty"`
	Formula         string   `yaml:"formula,omitempty" json:"formula,omitempty"`
	Package         string   `yaml:"package,omitempty" json:"package,omitempty"`
	Packages        []string `yaml:"packages,omitempty" json:"packages,omitempty"`
	Bins            []string `yaml:"bins,omitempty" json:"bins,omitempty"`
	Label           string   `yaml:"label,omitempty" json:"label,omitempty"`
	Tap             string   `yaml:"tap,omitempty" json:"tap,omitempty"`
	Module          string   `yaml:"module,omitempty" json:"module,omitempty"`
	URL             string   `yaml:"url,omitempty" json:"url,omitempty"`
	Archive         string   `yaml:"archive,omitempty" json:"archive,omitempty"`
	TargetDir       string   `yaml:"targetDir,omitempty" json:"target_dir,omitempty"`
	OS              []string `yaml:"os,omitempty" json:"os,omitempty"`
	Extract         bool     `yaml:"extract,omitempty" json:"extract,omitempty"`
	StripComponents int      `yaml:"stripComponents,omitempty" json:"strip_components,omitempty"`
}

type Source struct {
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Requires    Requirement     `json:"requires,omitempty"`
	Install     []InstallAction `json:"install,omitempty"`
}

type Profile struct {
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Expands     []string        `json:"expands,omitempty"`
	Requires    Requirement     `json:"requires,omitempty"`
	Install     []InstallAction `json:"install,omitempty"`
}

var builtinProfiles = map[string]Profile{
	ProfilePDF: {
		Name:        ProfilePDF,
		Description: "PDF readers, text extraction, and Python fallbacks.",
		Requires: Requirement{
			Bins: []string{"pdftotext", "pdfinfo"},
			Python: []PythonPackage{
				{Module: "pypdf", Package: "pypdf"},
				{Module: "PyPDF2", Package: "PyPDF2"},
				{Module: "fitz", Package: "PyMuPDF"},
			},
		},
		Install: []InstallAction{
			systemInstall(
				InstallKindBrew,
				"poppler",
				"Install PDF tools (brew)",
				"pdftotext",
				"pdfinfo",
			),
			systemInstall(
				InstallKindAPT,
				"poppler-utils",
				"Install PDF tools (apt)",
				"pdftotext",
				"pdfinfo",
			),
			systemInstall(
				InstallKindDNF,
				"poppler-utils",
				"Install PDF tools (dnf)",
				"pdftotext",
				"pdfinfo",
			),
			systemInstall(
				InstallKindYUM,
				"poppler-utils",
				"Install PDF tools (yum)",
				"pdftotext",
				"pdfinfo",
			),
		},
	},
	ProfileOffice: {
		Name:        ProfileOffice,
		Description: "Spreadsheet, Word, and slide parsing helpers.",
		Requires: Requirement{
			Python: []PythonPackage{
				{Module: "pandas", Package: "pandas"},
				{Module: "openpyxl", Package: "openpyxl"},
				{Module: "docx", Package: "python-docx"},
				{Module: "pptx", Package: "python-pptx"},
			},
		},
	},
	ProfileAudio: {
		Name:        ProfileAudio,
		Description: "Audio transcoding and inspection tools.",
		Requires: Requirement{
			Bins: []string{"ffmpeg", "ffprobe"},
		},
		Install: []InstallAction{
			systemInstall(
				InstallKindBrew,
				"ffmpeg",
				"Install ffmpeg (brew)",
				"ffmpeg",
				"ffprobe",
			),
			systemInstall(
				InstallKindAPT,
				"ffmpeg",
				"Install ffmpeg (apt)",
				"ffmpeg",
				"ffprobe",
			),
			systemInstall(
				InstallKindDNF,
				"ffmpeg",
				"Install ffmpeg (dnf)",
				"ffmpeg",
				"ffprobe",
			),
			systemInstall(
				InstallKindYUM,
				"ffmpeg",
				"Install ffmpeg (yum)",
				"ffmpeg",
				"ffprobe",
			),
		},
	},
	ProfileVideo: {
		Name:        ProfileVideo,
		Description: "Video frame extraction and transcoding tools.",
		Requires: Requirement{
			Bins: []string{"ffmpeg", "ffprobe"},
		},
		Install: []InstallAction{
			systemInstall(
				InstallKindBrew,
				"ffmpeg",
				"Install ffmpeg (brew)",
				"ffmpeg",
				"ffprobe",
			),
			systemInstall(
				InstallKindAPT,
				"ffmpeg",
				"Install ffmpeg (apt)",
				"ffmpeg",
				"ffprobe",
			),
			systemInstall(
				InstallKindDNF,
				"ffmpeg",
				"Install ffmpeg (dnf)",
				"ffmpeg",
				"ffprobe",
			),
			systemInstall(
				InstallKindYUM,
				"ffmpeg",
				"Install ffmpeg (yum)",
				"ffmpeg",
				"ffprobe",
			),
		},
	},
	ProfileImage: {
		Name:        ProfileImage,
		Description: "Common image conversion and manipulation tools.",
		Requires: Requirement{
			AnyBins: []string{"magick", "convert"},
		},
		Install: []InstallAction{
			systemInstall(
				InstallKindBrew,
				"imagemagick",
				"Install ImageMagick (brew)",
				"magick",
				"convert",
			),
			systemInstall(
				InstallKindAPT,
				"imagemagick",
				"Install ImageMagick (apt)",
				"magick",
				"convert",
			),
			systemInstall(
				InstallKindDNF,
				"ImageMagick",
				"Install ImageMagick (dnf)",
				"magick",
				"convert",
			),
			systemInstall(
				InstallKindYUM,
				"ImageMagick",
				"Install ImageMagick (yum)",
				"magick",
				"convert",
			),
		},
	},
	ProfileOCR: {
		Name:        ProfileOCR,
		Description: "OCR utilities for scanned images and documents.",
		Requires: Requirement{
			Bins: []string{"tesseract"},
		},
		Install: []InstallAction{
			systemInstall(
				InstallKindBrew,
				"tesseract",
				"Install tesseract (brew)",
				"tesseract",
			),
			systemInstall(
				InstallKindAPT,
				"tesseract-ocr",
				"Install tesseract (apt)",
				"tesseract",
			),
			systemInstall(
				InstallKindDNF,
				"tesseract",
				"Install tesseract (dnf)",
				"tesseract",
			),
			systemInstall(
				InstallKindYUM,
				"tesseract",
				"Install tesseract (yum)",
				"tesseract",
			),
		},
	},
	ProfileCommonFileTools: {
		Name:        ProfileCommonFileTools,
		Description: "Recommended default toolchain for common file work.",
		Expands: []string{
			ProfilePDF,
			ProfileOffice,
			ProfileAudio,
			ProfileVideo,
			ProfileImage,
			ProfileOCR,
		},
	},
}

func systemInstall(
	kind string,
	pkg string,
	label string,
	bins ...string,
) InstallAction {
	_ = "STUB: not implemented"
	return *new(InstallAction)
}

func Profiles() []Profile { _ = "STUB: not implemented"; return nil }

func DefaultProfiles() []string { _ = "STUB: not implemented"; return nil }

func ResolveProfiles(names []string) ([]Profile, error) { _ = "STUB: not implemented"; return nil, nil }

func SourcesForProfiles(names []string) ([]Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ProfileNames() []string { _ = "STUB: not implemented"; return nil }

func normalizeProfileNames(names []string) []string { _ = "STUB: not implemented"; return nil }

func appendProfile(
	dst *[]Profile,
	seen map[string]struct{},
	name string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func isAggregateProfile(profile Profile) bool { _ = "STUB: not implemented"; return false }

func normalizeProfile(profile Profile) Profile { _ = "STUB: not implemented"; return *new(Profile) }

func normalizeSource(source Source) Source { _ = "STUB: not implemented"; return *new(Source) }

func normalizeRequirement(req Requirement) Requirement {
	_ = "STUB: not implemented"
	return *new(Requirement)
}

func normalizePythonPackages(
	pkgs []PythonPackage,
) []PythonPackage {
	_ = "STUB: not implemented"
	return nil
}

func normalizeInstallActions(
	actions []InstallAction,
) []InstallAction {
	_ = "STUB: not implemented"
	return nil
}

func normalizeStrings(values []string) []string { _ = "STUB: not implemented"; return nil }

func normalizeOSList(values []string) []string { _ = "STUB: not implemented"; return nil }

func normalizeOSName(raw string) string { _ = "STUB: not implemented"; return "" }

func MergeSources(sources ...Source) []Source { _ = "STUB: not implemented"; return nil }
