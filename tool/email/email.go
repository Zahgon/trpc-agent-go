//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package email provides send email tools for AI agents.
// This tool can send emails to personal email and some Corporate Email.
package email

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	// default name
	defaultName = "email"
)

// MailboxType mailbox
type MailboxType int32

const (
	// MailUnknown unknown mail
	MailUnknown MailboxType = 0
	// MailQQ qq mail
	MailQQ MailboxType = 1
	// Mail163 163 mail
	Mail163 MailboxType = 2
	// MailGmail google mail
	MailGmail MailboxType = 3
)

// MailboxTypeToString convert mailbox type to string
func MailboxTypeToString(mailboxType MailboxType) string {
	_ = "STUB: not implemented"

	// qq mail
	return ""
}

// 163 mail

// google mail

// unknown mail

// Option is a functional option for configuring the file tool set.
type Option func(*emailToolSet)

// emailToolSet implements the ToolSet interface for file operations.
type emailToolSet struct {
	sendEmailEnabled bool
	tools            []tool.Tool
}

// Tools implements the ToolSet interface.
func (e *emailToolSet) Tools(_ context.Context) []tool.Tool {
	_ = "STUB: not implemented"

	// Name implements the ToolSet interface.
	return nil
}

func (e *emailToolSet) Name() string {
	_ = "STUB: not implemented"

	// Close implements the ToolSet interface.
	return ""
}

func (e *emailToolSet) Close() error {
	_ = "STUB: not implemented"
	// No resources to clean up for file tools.
	return nil
}

// NewToolSet creates a new file tool set with the given options.
func NewToolSet(opts ...Option) (tool.ToolSet, error) {
	_ = "STUB: not implemented"
	return *new(tool.ToolSet), nil
}

// Apply user-provided options.

// Create function tools based on enabled features.
