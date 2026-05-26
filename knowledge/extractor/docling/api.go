//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//

package docling

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"

	"trpc.group/trpc-go/trpc-agent-go/knowledge/extractor"
)

type convertOptions struct {
	ToFormats       []string     `json:"to_formats,omitempty"`
	ImageExportMode ImageRefMode `json:"image_export_mode,omitempty"`
	DoOCR           *bool        `json:"do_ocr,omitempty"`
}

type convertDocumentResponse struct {
	Document exportDocumentResponse `json:"document"`
	Status   string                 `json:"status"`
	Errors   []struct {
		ErrorMessage string `json:"error_message"`
	} `json:"errors"`
}

type exportDocumentResponse struct {
	Filename       string `json:"filename"`
	MdContent      string `json:"md_content"`
	TextContent    string `json:"text_content"`
	HTMLContent    string `json:"html_content"`
	DocTagsContent string `json:"doctags_content"`
}

func buildConvertOptions(extOpts options, eopts *extractor.Options) convertOptions {
	_ = "STUB: not implemented"
	return *new(convertOptions)
}

func doclingOutputFormat(outputFormat string) string { _ = "STUB: not implemented"; return "" }

func decodeConvertResponse(resp *http.Response) (*convertDocumentResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toExtractorResult(convResp *convertDocumentResponse, outputFormat string) (*extractor.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func pickOutputContent(doc exportDocumentResponse, outputFormat string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func writeFileRequest(writer *multipart.Writer, r io.Reader, opts convertOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Extractor) doFileConvert(ctx context.Context, r io.Reader, eopts *extractor.Options) (*extractor.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
