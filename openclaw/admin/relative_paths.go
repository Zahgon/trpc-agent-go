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
	"bytes"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

const (
	headerContentLength = "Content-Length"
	headerContentType   = "Content-Type"
	headerLocation      = "Location"

	htmlMediaType = "text/html"

	htmlAttrAction     = "action"
	htmlAttrDataPrefix = "data-"
	htmlAttrFormAction = "formaction"
	htmlAttrHref       = "href"
	htmlAttrSrc        = "src"

	htmlAttrActionSuffix = "-action"
	htmlAttrHrefSuffix   = "-href"
	htmlAttrPathSuffix   = "-path"
	htmlAttrSrcSuffix    = "-src"
	htmlAttrURLSuffix    = "-url"

	rootPath           = "/"
	currentPathSegment = "."
	parentPathSegment  = ".."
)

var htmlDataReferenceAttrSuffixes = []string{
	htmlAttrActionSuffix,
	htmlAttrHrefSuffix,
	htmlAttrPathSuffix,
	htmlAttrSrcSuffix,
	htmlAttrURLSuffix,
}

// wrapRelativeLinks keeps admin navigation working when the
// service is exposed behind a reverse-proxy subpath.
func wrapRelativeLinks(base http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func wrapRelativeLinksFunc(
	handler http.HandlerFunc,
) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

type bufferedResponseWriter struct {
	header      http.Header
	body        bytes.Buffer
	status      int
	wroteHeader bool
}

func newBufferedResponseWriter() *bufferedResponseWriter { _ = "STUB: not implemented"; return nil }

func (w *bufferedResponseWriter) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (w *bufferedResponseWriter) WriteHeader(status int) { _ = "STUB: not implemented"; return }

func (w *bufferedResponseWriter) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *bufferedResponseWriter) Flush() { _ = "STUB: not implemented"; return }

func writeRelativeResponse(
	dst http.ResponseWriter,
	req *http.Request,
	src *bufferedResponseWriter,
) {
	_ = "STUB: not implemented"
	return
}

func rewriteHTMLBody(
	requestPath string,
	contentType string,
	body []byte,
) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func isHTMLContentType(contentType string) bool { _ = "STUB: not implemented"; return false }

func rewriteHTMLReferences(node *html.Node, requestPath string) { _ = "STUB: not implemented"; return }

func isHTMLReferenceAttr(key string) bool { _ = "STUB: not implemented"; return false }

func isHTMLDataReferenceAttr(key string) bool { _ = "STUB: not implemented"; return false }

func relativeRequestReference(requestPath string, rawTarget string) string {
	_ = "STUB: not implemented"
	return ""
}

func parseRootRelativeURL(raw string) (*url.URL, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func relativePathReference(requestPath string, targetPath string) string {
	_ = "STUB: not implemented"
	return ""
}

func requestBaseDir(raw string) string { _ = "STUB: not implemented"; return "" }

func requestPath(req *http.Request) string { _ = "STUB: not implemented"; return "" }

func copyHeaders(dst http.Header, src http.Header) { _ = "STUB: not implemented"; return }

func splitURLPath(raw string) []string { _ = "STUB: not implemented"; return nil }

func commonPathPrefixLen(left []string, right []string) int { _ = "STUB: not implemented"; return 0 }
