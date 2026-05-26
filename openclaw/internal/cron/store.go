//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package cron

const (
	storeVersion  = 1
	storeFilePerm = 0o600
	storeDirPerm  = 0o700

	storeTempPattern = defaultJobsFile + ".tmp-*"
)

type storeData struct {
	Version int    `json:"version"`
	Jobs    []*Job `json:"jobs"`
}

func loadJobs(path string) ([]*Job, error) { _ = "STUB: not implemented"; return nil, nil }

func saveJobs(path string, jobs []*Job) error { _ = "STUB: not implemented"; return nil }

func cloneJobs(jobs []*Job) []*Job { _ = "STUB: not implemented"; return nil }
