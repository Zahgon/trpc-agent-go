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
	"context"
	"io"
	"net/url"
	"os"
	"os/exec"
	"sync"
	"time"

	"trpc.group/trpc-go/trpc-agent-go/openclaw/admin"
)

const (
	browserServerManagedPort = "19790"
	browserServerScheme      = "http"

	browserServerDirEnv = "OPENCLAW_BROWSER_SERVER_DIR"

	browserServerAddrEnv  = "OPENCLAW_BROWSER_SERVER_ADDR"
	browserServerTokenEnv = "OPENCLAW_BROWSER_SERVER_TOKEN"

	browserServerDirName    = "browser-server"
	browserServerPackage    = "package.json"
	browserServerBinDirName = "bin"
	browserServerScriptName = "openclaw-browser-server.js"
	browserServerNodeBin    = "node"

	browserServerStateExternal = "external"
	browserServerStateStarting = "starting"
	browserServerStateRunning  = "running"
	browserServerStateStopping = "stopping"
	browserServerStateStopped  = "stopped"
	browserServerStateFailed   = "failed"

	browserServerLogDirName  = "services"
	browserServerLogFileName = "browser-server.log"
	browserServerTailLines   = 40
)

const (
	browserServerProbeTimeout  = 1500 * time.Millisecond
	browserServerStartTimeout  = 20 * time.Second
	browserServerStartInterval = 200 * time.Millisecond
	browserServerStopTimeout   = 5 * time.Second
)

var (
	browserServerNow       = time.Now
	browserServerProbeFunc = probeBrowserServerEndpoint
	browserServerWorkDir   = defaultBrowserServerWorkDir
	browserServerCommand   = defaultBrowserServerCommand
)

type browserServerPlan struct {
	ProviderName string
	ServerURL    string
	Addr         string
	AuthToken    string
}

type browserManagedStatus = admin.BrowserManagedService

type browserServerProcessExit struct {
	Code int
	Err  error
}

type browserServerSup struct {
	mu sync.RWMutex

	plan browserServerPlan
	tail *browserServerTail

	state           string
	managed         bool
	pid             int
	workDir         string
	command         string
	logPath         string
	logRelativePath string
	startedAt       *time.Time
	stoppedAt       *time.Time
	exitCode        *int
	lastError       string
	stopRequested   bool

	cmd    *exec.Cmd
	doneCh chan browserServerProcessExit
}

type browserServerTail struct {
	mu      sync.Mutex
	pending string
	lines   []string
	limit   int
}

func maybeStartBrowserServerSupervisor(
	ctx context.Context,
	specs []pluginSpec,
	debugDir string,
) (*browserServerSup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func detectManagedBrowserServerPlan(
	specs []pluginSpec,
) (browserServerPlan, bool) {
	_ = "STUB: not implemented"
	return *new(browserServerPlan), false
}

func managedBrowserServerAddr(raw string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func parseBrowserServerURL(raw string) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isManagedBrowserServerHost(host string) bool { _ = "STUB: not implemented"; return false }

func newBrowserServerSupervisor(
	plan browserServerPlan,
	debugDir string,
) *browserServerSup {
	_ = "STUB: not implemented"
	return nil
}

func newBrowserServerTail(limit int) *browserServerTail { _ = "STUB: not implemented"; return nil }

func (t *browserServerTail) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (t *browserServerTail) Lines() []string { _ = "STUB: not implemented"; return nil }

func (t *browserServerTail) append(line string) { _ = "STUB: not implemented"; return }

func (s *browserServerSup) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *browserServerSup) startManaged(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *browserServerSup) openLogFile() (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *browserServerSup) waitUntilReady(
	ctx context.Context,
	doneCh <-chan browserServerProcessExit,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *browserServerSup) waitForExit(
	cmd *exec.Cmd,
	logFile *os.File,
	doneCh chan<- browserServerProcessExit,
) {
	_ = "STUB: not implemented"
	return
}

func processExitCode(state *os.ProcessState) int { _ = "STUB: not implemented"; return 0 }

func (s *browserServerSup) probe(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func probeBrowserServerEndpoint(
	ctx context.Context,
	serverURL string,
	authToken string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func defaultBrowserServerWorkDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

func browserServerWorkDirIn(root string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func isBrowserServerWorkDir(dir string) bool { _ = "STUB: not implemented"; return false }

func defaultBrowserServerCommand(
	workDir string,
	env []string,
	stdout io.Writer,
	stderr io.Writer,
) (*exec.Cmd, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (s *browserServerSup) setFailed(err error) { _ = "STUB: not implemented"; return }

func (s *browserServerSup) Close() error { _ = "STUB: not implemented"; return nil }

func (s *browserServerSup) startupLines() []startupLogLine { _ = "STUB: not implemented"; return nil }

func (s *browserServerSup) BrowserManagedStatus() browserManagedStatus {
	_ = "STUB: not implemented"
	return *new(browserManagedStatus)
}
