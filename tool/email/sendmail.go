//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package email

import (
	"context"

	"trpc.group/trpc-go/trpc-agent-go/tool"
)

const (
	qqMail          = "smtp.qq.com"
	qqPort          = 465
	gmailMail       = "smtp.gmail.com"
	gmailPort       = 587
	netEase163Mail  = "smtp.163.com"
	netEase1163Port = 465
)

// sendMailRequest represents the input for the send mail operation.
type sendMailRequest struct {
	Auth     Auth      `json:"auth" jsonschema:"description=auth of the mail."`
	MailList []*Mail   `json:"mail_list" jsonschema:"description=The list of mail."`
	Extra    ExtraData `json:"extra" jsonschema:"description=extra data of the mail. optional. default is empty."`
}

// Mail represents a mail to be sent.
type Mail struct {
	ToEmail string `json:"to_email" jsonschema:"description=send to email."`
	Subject string `json:"subject" jsonschema:"description=subject of the mail"`
	Content string `json:"content" jsonschema:"description=content of the mail"`
}

// Auth is a struct for email authentication.
type Auth struct {
	Name     string `json:"name" jsonschema:"description=name of the mail."`
	Password string `json:"password" jsonschema:"description=password of the mail."`
}

// ExtraData represents extra data for the mail.
type ExtraData struct {
	SvrAddr string `json:"svr_addr" jsonschema:"description=server address of the mail. optional. default is empty."`
	Port    int    `json:"port" jsonschema:"description=port of the mail. optional. default is empty."`
}

// sendMailResponse represents the output from the send mail operation.
type sendMailResponse struct {
	Message string `json:"message"`
}

// sendMail performs the send mail operation.
func (e *emailToolSet) sendMail(ctx context.Context, req *sendMailRequest) (rsp *sendMailResponse, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// batch send email, not stop if one failed, return err  which join all send error message

//qq mail special error handle
//https://github.com/wneessen/go-mail/issues/463

// getEmailAddr gets the email address and port.
func (e *emailToolSet) getEmailAddr(req *sendMailRequest) (addr string, port int, isSSL bool, err error) {
	_ = "STUB: not implemented"
	return "", 0, false, nil
}

//qq email

//gmail email

//163 email

// not support

// checkMailBoxType checks the mailbox type.
func checkMailBoxType(email string) (MailboxType, error) {
	_ = "STUB: not implemented"
	return *new(MailboxType), nil
}

// to lower

// split by name and domain

// sendMailTool returns a callable tool for send mail.
func (e *emailToolSet) sendMailTool() tool.CallableTool {
	_ = "STUB: not implemented"
	return *new(tool.CallableTool)
}

func qqHandleError(err error) error { _ = "STUB: not implemented"; return nil }

// Check if this is an SMTP RESET error after successful delivery

// https://github.com/wneessen/go-mail/issues/463

// Don't treat this as a delivery failure since mail was sent
