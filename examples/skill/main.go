//
// Tencent is pleased to support the open source community by making
// trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

// Package main is the main package for the skill example.
package main

import (
	"archive/zip"
	"context"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
	"trpc.group/trpc-go/trpc-agent-go/agent"
	"trpc.group/trpc-go/trpc-agent-go/model"
	"trpc.group/trpc-go/trpc-agent-go/runner"
)

var (
	datasetPath = flag.String(
		"dataset",
		defaultDatasetPath,
		"Path to GAIA dataset",
	)
	dataDir = flag.String(
		"data-dir",
		defaultDataDir,
		"Directory containing data files",
	)
	outputPath = flag.String(
		"output",
		defaultOutputPath,
		"Path to output results",
	)
	maxTasks = flag.Int(
		"tasks",
		0,
		"Maximum number of tasks to run (0 means all)",
	)
	modelName = flag.String(
		"model",
		defaultModelName,
		"Model name to use",
	)
	specificTask = flag.String(
		"task-id",
		"",
		"Run a single task by task ID or 1-based index",
	)
)

const (
	defaultDatasetPath = "../data/gaia_2023_level1_validation.json"
	defaultDataDir     = "../data"
	defaultOutputPath  = "../results/trpc-agent-go.json"
	defaultModelName   = "deepseek-v4-flash"

	maxDocChars = 100000
)

const truncatedSuffix = "\n\n... (Content truncated due to size limit)"

// Save the original working directory.
var originalDir string

var currencySymbolsToStrip = []string{
	"$",
	"€",
	"£",
	"¥",
	"₹",
	"¢",
	"₽",
	"₩",
	"฿",
}

func init() {
	var err error
	originalDir, err = os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	ensureWhisperPython3()
}

func ensureWhisperPython3() { _ = "STUB: not implemented"; return }

// canImportWhisperWithPython3 checks if whisper can be imported using python3.
func canImportWhisperWithPython3() bool { _ = "STUB: not implemented"; return false }

// canImportWhisperWithPython checks if whisper can be imported using python.
func canImportWhisperWithPython() bool { _ = "STUB: not implemented"; return false }

const envGAIADebugModel = "TRPC_AGENT_GAIA_DEBUG_MODEL"

func buildModelDebugCallbacks() *model.Callbacks { _ = "STUB: not implemented"; return nil }

func requestSummary(
	req *model.Request,
) (int, int, *model.Message) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func hashModelRequest(req *model.Request) string { _ = "STUB: not implemented"; return "" }

// GAIATask represents a single task from the dataset
type GAIATask struct {
	TaskID      string `json:"task_id"`
	Question    string `json:"Question"`
	Level       string `json:"Level"`
	FinalAnswer string `json:"Final answer"`
	FileName    string `json:"file_name"`
	FilePath    string `json:"file_path"`
}

// BenchmarkResult stores the result of a single task
type BenchmarkResult struct {
	TaskID          string        `json:"task_id"`
	Question        string        `json:"question"`
	Level           string        `json:"level"`
	PredictedAnswer string        `json:"predicted_answer"`
	GroundTruth     string        `json:"ground_truth"`
	Correct         bool          `json:"correct"`
	Steps           int           `json:"steps"`
	ExecutionTime   time.Duration `json:"execution_time_ms"`
	TokensUsed      int           `json:"tokens_used"`
	ToolCalls       int           `json:"tool_calls"`
	Error           string        `json:"error,omitempty"`
}

// SummaryResult stores aggregated results
type SummaryResult struct {
	Framework       string            `json:"framework"`
	TotalTasks      int               `json:"total_tasks"`
	CorrectCount    int               `json:"correct_count"`
	Accuracy        float64           `json:"accuracy"`
	AvgSteps        float64           `json:"avg_steps"`
	AvgTime         float64           `json:"avg_time_ms"`
	AvgTokens       float64           `json:"avg_tokens"`
	AvgToolCalls    float64           `json:"avg_tool_calls"`
	DetailedResults []BenchmarkResult `json:"detailed_results"`
}

// SystemPrompt is the system instruction used for GAIA evaluation.
const SystemPrompt = `You are an AI assistant designed to answer
questions accurately and concisely.

============================================================
REASONING PROTOCOL
============================================================

STEP 1: UNDERSTAND THE QUESTION
- Identify what is being asked (number, name, list, yes/no, etc.)
- Extract key constraints: time periods, conditions, specific requirements
- Recognize question type: calculation, retrieval, comparison, logic

STEP 2: PLAN YOUR APPROACH
- For CALCULATIONS: Break down into steps, verify formula
- For TIME-BASED queries: Note the exact time period required
  (as of YYYY, in MM/YYYY)
- For LISTS: Ensure you check ALL candidates, don't skip any
- For COMPARISONS: Gather data for all items being compared
- For LOGIC: State the rule or principle being applied

STEP 3: EXECUTE WITH VERIFICATION
- Gather information using appropriate tools
- For numbers: Double-check calculations (consider using multiple methods)
- For historical data: Verify you're using data from the CORRECT time period
- For lists: Maintain a checklist and verify completeness
- NEVER mix current data with historical requests

STEP 4: VALIDATE BEFORE ANSWERING
Checklist:
□ Did I understand the question correctly?
□ Did I use the right time period (if applicable)?
□ Did I check ALL items (for lists)?
□ Did I verify my calculations?
□ Is my answer in the correct format?
□ Did I avoid common errors (mixing dates, skipping items)?

============================================================
CRITICAL RULES
============================================================

FORMAT REQUIREMENTS:
✓ Always end with: FINAL ANSWER: <concise answer>
✓ Numbers: provide just the number (e.g., "42" not "42 units")
✓ Names: provide just the name (e.g., "John Smith")
✓ Lists: use consistent format (e.g., "A, B, C")
✓ NEVER include explanations in the final answer line

COMMON ERRORS TO AVOID:
❌ Using current data when asked about historical periods
❌ Skipping items when checking lists
❌ Calculation errors (verify with multiple approaches)
❌ Confusing similar names/locations
❌ Returning execution plan instead of final answer
❌ Outputting empty or malformed answers

SPECIAL HANDLING:
📊 CALCULATIONS: Show your work, verify results
📅 TIME-BASED: Explicitly state "using data from [time]"
📋 LISTS: Create checklist, verify ALL items checked
🔍 FILE ANALYSIS: Read files completely; don't skip sections.
    Note: Large files (PDF, PPTX) may be truncated; look for truncation
    messages.
🌐 WEB RESEARCH: Cross-verify information from multiple sources

============================================================
AVAILABLE TOOLS
============================================================

- web_search: DuckDuckGo HTML search (general web)
- web_fetch: Fetch one or multiple URLs (HTML/PDF/JSON/XML/text).
  Large PDFs/pages may be truncated.
- wikipedia_wikipedia_search: Wikipedia search
- arxiv_search_search: arXiv search
- read_document: Read local documents (PDF/XLSX/CSV/DOCX/TXT).
  Not for images/audio.
- file_list_file: List files under data-dir (relative paths only)
- file_search_file: Find files by pattern under data-dir
- file_search_content: Search within text files under data-dir
- skill_load / skill_run: Use skills for audio/image when available

Note: Large files (>100KB text) may be truncated to avoid token limits.
You'll see a truncation message if this happens.

============================================================
OUTPUT FORMAT
============================================================

When you have found the answer:

FINAL ANSWER: <your concise, precise answer>

Example formats:
- Number: "FINAL ANSWER: 42"
- Name: "FINAL ANSWER: Albert Einstein"
- List: "FINAL ANSWER: Alice, Bob, Carol"
- Yes/No: "FINAL ANSWER: Yes"

DO NOT include explanations, units, or additional text after "FINAL ANSWER:"`

func main() {
	flag.Parse()

	log.Printf("Starting GAIA Benchmark evaluation for trpc-agent-go")
	log.Printf("Dataset: %s", *datasetPath)
	log.Printf("Data directory: %s", *dataDir)
	log.Printf("Max tasks: %d", *maxTasks)

	// Load dataset
	tasks, err := loadDataset(*datasetPath)
	if err != nil {
		log.Fatalf("Failed to load dataset: %v", err)
	}

	// Filter by specific task if specified
	if *specificTask != "" {
		tasks = filterTasksByID(tasks, *specificTask)
		if len(tasks) == 0 {
			log.Fatalf("No task found matching: %s", *specificTask)
		}
		log.Printf("Running specific task: %s", *specificTask)
	} else if *maxTasks > 0 && len(tasks) > *maxTasks {
		tasks = tasks[:*maxTasks]
	}

	log.Printf("Loaded %d tasks", len(tasks))

	// Create agent
	gaiaAgent := createGAIAAgent()

	// Run benchmark
	results := runBenchmark(gaiaAgent, tasks)

	// Save results
	if err := saveResults(results, *outputPath); err != nil {
		log.Fatalf("Failed to save results: %v", err)
	}

	// Print summary
	printSummary(results)
}

func loadDataset(path string) ([]GAIATask, error) { _ = "STUB: not implemented"; return nil, nil }

// filterTasksByID returns a single task by ID or 1-based index.
func filterTasksByID(tasks []GAIATask, idOrIndex string) []GAIATask {
	_ = "STUB: not implemented"
	return nil
}

// read_document tool.

// ReadFileRequest is the request for read_document.
type ReadFileRequest struct {
	FilePath string `json:"file_path" jsonschema:"description=Path to file"`
}

// ReadFileResponse is the response for read_document.
type ReadFileResponse struct {
	Content  string `json:"content"`
	FileType string `json:"file_type"`
	Error    string `json:"error,omitempty"`
}

// readFileContent reads local documents and extracts text.
//
// It also supports workspace:// and artifact:// file refs, so skill outputs
// can be passed across tools without guessing filesystem roots.
func readFileContent(
	ctx context.Context,
	req ReadFileRequest,
) (ReadFileResponse, error) {
	_ = "STUB: not implemented"
	return *new(ReadFileResponse), nil
}

// Fall back to plain text.

func looksLikeSkillWorkspacePath(p string) bool { _ = "STUB: not implemented"; return false }

func stripInputsPrefix(p string) string { _ = "STUB: not implemented"; return "" }

// readPDFFile reads a PDF file and extracts plain text.
func readPDFFile(filePath string) (ReadFileResponse, error) {
	_ = "STUB: not implemented"
	return *new(ReadFileResponse), nil
}

// readExcelFile reads an Excel file and outputs a text representation.
// If the sheet contains background colors, the output includes color hints.
func readExcelFile(filePath string) (ReadFileResponse, error) {
	_ = "STUB: not implemented"
	return *new(ReadFileResponse), nil
}

const (
	excelARGBPrefix      = "FF"
	excelColorWhite      = "FFFFFF"
	excelColorWhiteShort = "FFFF"
	excelColorBlack      = "000000"
	excelScanExtraCols   = 10
)

const (
	excelColorNote = "Note: This Excel file contains background colors.\n"
	excelColorFmt  = "Format: Cell=Value[#RRGGBB] or Cell=#RRGGBB.\n\n"
)

func maxRowLen(rows [][]string) int { _ = "STUB: not implemented"; return 0 }

func scanExcelSheet(
	f *excelize.File,
	sheet string,
	rows [][]string,
	maxCols int,
) (bool, int) {
	_ = "STUB: not implemented"
	return false, 0
}

func excelCellName(colIdx int, rowIdx int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func excelCellBGColor(
	f *excelize.File,
	sheet string,
	cellName string,
) string {
	_ = "STUB: not implemented"
	return ""
}

func excelIsMeaningfulColor(color string) bool { _ = "STUB: not implemented"; return false }

func writeExcelSheetWithColors(
	b *strings.Builder,
	f *excelize.File,
	sheet string,
	rows [][]string,
	maxUsedCol int,
) {
	_ = "STUB: not implemented"
	return
}

func writeExcelSheetPlain(b *strings.Builder, rows [][]string) { _ = "STUB: not implemented"; return }

// readCSVFile reads a CSV file.
func readCSVFile(filePath string) (ReadFileResponse, error) {
	_ = "STUB: not implemented"
	return *new(ReadFileResponse), nil
}

// Allow variable field counts.

// Fall back to plain text.

// readDocxFile reads a DOCX/PPTX (Office ZIP) and extracts text.
func readDocxFile(filePath string) (ReadFileResponse, error) {
	_ = "STUB: not implemented"
	return *new(ReadFileResponse), nil
}

// extractDocxText extracts text from an Office ZIP payload.
func extractDocxText(data []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

// DOCX
// PPTX slides folder

// For PPTX, add slide markers.

// extractSlideNumber extracts a slide number from a slide XML path.
func extractSlideNumber(filename string) string { _ = "STUB: not implemented"; return "" }

// newZipReader creates a zip reader from an in-memory buffer.
func newZipReader(data []byte) (*zip.Reader, error) { _ = "STUB: not implemented"; return nil, nil }

// extractTextFromXML extracts text from DOCX/PPTX XML documents.
func extractTextFromXML(xml string) string {
	_ = "STUB: not implemented"

	// DOCX: <w:t> tags.
	return ""
}

// PPTX: <a:t> tags.

// readTextFile reads a plain text file.
func readTextFile(filePath string) (ReadFileResponse, error) {
	_ = "STUB: not implemented"
	return *new(ReadFileResponse), nil
}

// DuckDuckGo HTML search implementation for the "web_search" tool.

const (
	ddgFormQueryKey      = "q"
	ddgFormContentType   = "application/x-www-form-urlencoded"
	ddgHTMLSearchURL     = "https://html.duckduckgo.com/html/"
	ddgDefaultMaxResults = 5
	ddgHTTPTimeout       = 30 * time.Second
	httpPrefix           = "http://"
	httpsPrefix          = "https://"
)

const ddgUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
	"AppleWebKit/537.36 (KHTML, like Gecko) " +
	"Chrome/120.0.0.0 Safari/537.36"

const (
	ddgLinkPattern = `class="result__a"[^>]*href="([^"]+)"[^>]*>` +
		`([^<]+)</a>`
	ddgSnippetPattern = `class="result__snippet"[^>]*>([^<]+)</a>`
)

// DDGSearchRequest is the request for web_search.
type DDGSearchRequest struct {
	Query      string `json:"query" jsonschema:"description=Query,required"`
	MaxResults int    `json:"max_results,omitempty" jsonschema:"description=Max"`
}

// DDGSearchResult is a single search result.
type DDGSearchResult struct {
	Title   string `json:"title"`
	URL     string `json:"url"`
	Snippet string `json:"snippet"`
}

// DDGSearchResponse is the response for web_search.
type DDGSearchResponse struct {
	Query   string            `json:"query"`
	Results []DDGSearchResult `json:"results"`
	Summary string            `json:"summary"`
	Error   string            `json:"error,omitempty"`
}

func duckduckgoHTMLSearch(
	ctx context.Context,
	req DDGSearchRequest,
) (DDGSearchResponse, error) {
	_ = "STUB: not implemented"
	return *new(DDGSearchResponse), nil
}

func parseDDGHTML(html string, maxResults int) []DDGSearchResult {
	_ = "STUB: not implemented"
	return nil
}

func cleanHTML(s string) string { _ = "STUB: not implemented"; return "" }

func createGAIAAgent() agent.Agent {
	_ = "STUB: not implemented"
	// Create model
	return *new(agent.Agent)
}

// DuckDuckGo HTML search tool (not the Instant Answer API).

// ArXiv search toolset.

// Wikipedia search toolset.

// Absolute data directory.

// Create the file toolset for directory discovery only (list/search).
// Reading should use read_document or skill_run output_files content.

// read_document tool for local files (PDF/Excel/CSV/DOCX/TXT).

// GPT-5 uses a fixed temperature; keep the default (1.0).

// Code executor for skills.

// Enable skills if configured. This benchmark reuses an older
// output-collection path, so it is pinned to the legacy full
// skill tool surface. New code should not copy this wiring; see
// docs/mkdocs/en/skill.md for the recommended setup (WithSkills +
// an explicit CodeExecutor is enough, no profile needed).

// Keep the CodeExecutor strictly for tool-backed execution
// (skill_run / workspace_exec). Fenced-code auto-execution
// is an independent switch and is opted out here so the
// benchmark is driven only by tool calls.

func runBenchmark(ag agent.Agent, tasks []GAIATask) SummaryResult {
	_ = "STUB: not implemented"
	return *new(SummaryResult)
}

const audioAttachmentHint = `This is an audio file.
- Do NOT use read_document (it returns binary noise).
- Use the whisper skill (skill_load + skill_run) to transcribe it.
- Use the "Skill workspace path" shown above as the input path.
- Write the transcript under out/ (or $OUTPUT_DIR).
- Include the transcript file in output_files so content is inline.
- Answer using the transcript text.
- When listing items from the transcript, copy the exact wording verbatim.
  Keep all descriptors and spelling (do not paraphrase or shorten).
  Example: "freshly squeezed lemon juice" ≠ "fresh squeezed lemon juice".
  Example: "freshly squeezed lemon juice" ≠ "lemon juice".
  Example: "pure vanilla extract" ≠ "vanilla extract".`

const imageAttachmentHint = `This is an image file.
- Do NOT use read_document (it cannot read images).
- Use the ocr skill (skill_load + skill_run) to extract text.
- Use the "Skill workspace path" shown above as the input path.
- Write the OCR result under out/ (or $OUTPUT_DIR).
- Include the OCR output file in output_files so content is inline.
- Answer using the extracted text.
- When listing items from extracted text, copy the exact wording verbatim.
  Keep all descriptors (do not paraphrase or shorten).`

const defaultAttachmentHint = `Use the read_document tool to read this file.`

const skillWorkspacePathPrefix = "inputs/"

func attachmentHintForPath(filePath string) string { _ = "STUB: not implemented"; return "" }

func isAudioExt(ext string) bool { _ = "STUB: not implemented"; return false }

func isImageExt(ext string) bool { _ = "STUB: not implemented"; return false }

func runSingleTask(r runner.Runner, task GAIATask) BenchmarkResult {
	_ = "STUB: not implemented"
	return *new(BenchmarkResult)
}

// Build the user prompt.

// If there is an attachment, add a data-dir-relative path.
//
// Use relative paths in prompts because file tools reject absolute paths
// and ".." traversal. read_document also accepts data-dir-relative paths.

// Create user message

// Run the agent.

// Process events to count steps and tool calls
// Keep only the last assistant message (excluding tool messages).

// Count a step when we get usage.

// Tool call logging.

// Pretty-print JSON args for logs.

// Print tool results.

// Capture the final assistant message (no tool calls).

// Extract the predicted answer from the last assistant message.

// Verify the answer.

// extractFinalAnswer extracts the answer from the last assistant message.
func extractFinalAnswer(content string) string { _ = "STUB: not implemented"; return "" }

// formatAnswer formats an answer string for evaluation output.
func formatAnswer(answer string) string { _ = "STUB: not implemented"; return "" }

// verifyAnswer checks whether the predicted answer matches the ground truth.
func verifyAnswer(predicted, groundTruth string) bool { _ = "STUB: not implemented"; return false }

// normalizeAnswer normalizes answers to make scoring more robust.
func normalizeAnswer(answer string) string { _ = "STUB: not implemented"; return "" }

func calculateSummary(
	framework string,
	results []BenchmarkResult,
) SummaryResult {
	_ = "STUB: not implemented"
	return *new(SummaryResult)
}

func saveResults(summary SummaryResult, path string) error { _ = "STUB: not implemented"; return nil }

// Ensure the path is absolute.

// Create the output directory.

func printSummary(summary SummaryResult) { _ = "STUB: not implemented"; return }

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func floatPtr(f float64) *float64 { _ = "STUB: not implemented"; return nil }

func truncate(s string, maxLen int) string { _ = "STUB: not implemented"; return "" }
