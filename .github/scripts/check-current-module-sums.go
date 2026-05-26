//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main provides a PR-time module zip sum checker.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/mod/module"
)

type moduleEntry struct {
	GoModPath  string
	Dir        string
	ModulePath string
}

type mismatchEntry struct {
	GoModPath  string
	ModulePath string
	VCS        string
	Working    string
}

type unsupportedEntry struct {
	GoModPath  string
	ModulePath string
	Reason     string
}

func main() {
	version := flag.String("version", "v0.0.0", "Synthetic version used for module zip prefix.")
	revision := flag.String("revision", "HEAD", "Git revision used for VCS archive mode.")
	flag.Parse()

	repoRoot, err := gitRepoRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "::error::Unable to determine git repository root: %v\n", err)
		os.Exit(1)
	}

	entries, initialUnsupported, err := discoverModules(repoRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "::error::Unable to discover go modules: %v\n", err)
		os.Exit(1)
	}
	if len(entries) == 0 {
		fmt.Fprintf(os.Stderr, "::error::No go.mod files found.\n")
		os.Exit(1)
	}

	tmpDir, err := os.MkdirTemp("", "trpc-agent-go-sumcheck-*")
	if err != nil {
		fmt.Fprintf(os.Stderr, "::error::Unable to create temp dir: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	hasErrors := len(initialUnsupported) > 0
	skipped := []string{}
	mismatches := []mismatchEntry{}
	unsupported := append([]unsupportedEntry(nil), initialUnsupported...)

	for _, entry := range entries {
		if isDoNotUseGoMod(entry.GoModPath) {
			skipped = append(skipped, rel(repoRoot, entry.Dir))
			continue
		}

		subdir := rel(repoRoot, entry.Dir)
		if subdir == "." {
			subdir = ""
		}
		subdir = filepath.ToSlash(strings.TrimPrefix(subdir, "./"))

		fmt.Printf("::group::%s\n", fmt.Sprintf("%s@%s (%s)", entry.ModulePath, *version, humanModuleDir(subdir)))

		m := module.Version{Path: entry.ModulePath, Version: *version}
		if err := module.Check(m.Path, m.Version); err != nil {
			fmt.Fprintf(os.Stderr, "::error file=%s::Invalid synthetic module version for %s: %v\n", rel(repoRoot, entry.GoModPath), entry.ModulePath, err)
			unsupported = append(unsupported, unsupportedEntry{
				GoModPath:  rel(repoRoot, entry.GoModPath),
				ModulePath: entry.ModulePath,
				Reason:     fmt.Sprintf("invalid synthetic module version: %v", err),
			})
			hasErrors = true
			fmt.Println("::endgroup::")
			continue
		}

		vcsZipPath, err := createModuleZipFromVCS(tmpDir, m, repoRoot, *revision, subdir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "::error file=%s::Failed to create VCS module zip for %s: %v\n", rel(repoRoot, entry.GoModPath), entry.ModulePath, err)
			unsupported = append(unsupported, unsupportedEntry{
				GoModPath:  rel(repoRoot, entry.GoModPath),
				ModulePath: entry.ModulePath,
				Reason:     fmt.Sprintf("failed to create VCS zip: %v", err),
			})
			hasErrors = true
			fmt.Println("::endgroup::")
			continue
		}
		vcsSum, files, err := hashZipAndListFiles(vcsZipPath)
		_ = os.Remove(vcsZipPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "::error file=%s::Failed to compute VCS zip sum for %s: %v\n", rel(repoRoot, entry.GoModPath), entry.ModulePath, err)
			unsupported = append(unsupported, unsupportedEntry{
				GoModPath:  rel(repoRoot, entry.GoModPath),
				ModulePath: entry.ModulePath,
				Reason:     fmt.Sprintf("failed to compute VCS zip sum: %v", err),
			})
			hasErrors = true
			fmt.Println("::endgroup::")
			continue
		}

		zipPrefix := fmt.Sprintf("%s@%s/", m.Path, m.Version)
		wtSum, err := hashFilesFromWorkingTree(files, zipPrefix, repoRoot, entry.Dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "::error file=%s::Failed to compute working tree sum for %s: %v\n", rel(repoRoot, entry.GoModPath), entry.ModulePath, err)
			unsupported = append(unsupported, unsupportedEntry{
				GoModPath:  rel(repoRoot, entry.GoModPath),
				ModulePath: entry.ModulePath,
				Reason:     fmt.Sprintf("failed to compute working tree sum: %v", err),
			})
			hasErrors = true
			fmt.Println("::endgroup::")
			continue
		}

		fmt.Printf("vcs zip: %s\n", vcsSum)
		fmt.Printf("wt  zip: %s\n", wtSum)

		if vcsSum != wtSum {
			fmt.Fprintf(os.Stderr, "::error file=%s::Module zip sum mismatch for %s (VCS != working tree).\n", rel(repoRoot, entry.GoModPath), entry.ModulePath)
			mismatches = append(mismatches, mismatchEntry{
				GoModPath:  rel(repoRoot, entry.GoModPath),
				ModulePath: entry.ModulePath,
				VCS:        vcsSum,
				Working:    wtSum,
			})
			hasErrors = true
		}

		fmt.Println("::endgroup::")
	}

	if len(skipped) > 0 {
		sort.Strings(skipped)
		fmt.Println("::group::Skipped modules")
		for _, s := range skipped {
			fmt.Printf("- %s\n", s)
		}
		fmt.Println("::endgroup::")
	}

	if len(mismatches) > 0 {
		sort.Slice(mismatches, func(i, j int) bool {
			return mismatches[i].ModulePath < mismatches[j].ModulePath
		})
		fmt.Println("::group::Mismatched modules")
		for _, m := range mismatches {
			fmt.Printf("- %s (%s)\n", m.ModulePath, m.GoModPath)
			fmt.Printf("  vcs: %s\n", m.VCS)
			fmt.Printf("  wt : %s\n", m.Working)
		}
		fmt.Println("::endgroup::")
	}

	if len(unsupported) > 0 {
		sort.Slice(unsupported, func(i, j int) bool {
			return unsupported[i].ModulePath < unsupported[j].ModulePath
		})
		fmt.Println("::group::Unsupported modules")
		for _, u := range unsupported {
			fmt.Printf("- %s (%s): %s\n", u.ModulePath, u.GoModPath, u.Reason)
		}
		fmt.Println("::endgroup::")
	}

	if hasErrors {
		fmt.Fprintf(os.Stderr, "::error::Some modules have inconsistent sums between VCS and working tree.\n")
		os.Exit(1)
	}

	fmt.Println("All modules match between VCS and working tree.")
}

func discoverModules(repoRoot string) ([]moduleEntry, []unsupportedEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func readModulePath(goModPath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func isDoNotUseGoMod(goModPath string) bool { _ = "STUB: not implemented"; return false }

func createModuleZipFromVCS(tmpDir string, m module.Version, repoRoot, revision, subdir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func hashZipAndListFiles(zipPath string) (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func hashFilesFromWorkingTree(files []string, zipPrefix, repoRoot, moduleDir string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func humanModuleDir(subdir string) string { _ = "STUB: not implemented"; return "" }

func gitRepoRoot() (string, error) { _ = "STUB: not implemented"; return "", nil }

func git(repoRoot string, args ...string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func splitLines(s string) []string { _ = "STUB: not implemented"; return nil }

func rel(base, target string) string { _ = "STUB: not implemented"; return "" }
