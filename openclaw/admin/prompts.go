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
	"net/http"
)

const (
	routePromptsJSON        = "/api/prompts"
	routePromptInlineSave   = "/api/prompts/inline"
	routePromptRuntimeSave  = "/api/prompts/runtime"
	routePromptFileSave     = "/api/prompts/file"
	routePromptFileCreate   = "/api/prompts/file/create"
	routePromptFileDelete   = "/api/prompts/file/delete"
	routePersonasJSON       = "/api/personas"
	routePersonaSave        = "/api/personas/save"
	routePersonaDelete      = "/api/personas/delete"
	routePersonaDefaultSave = "/api/personas/default"

	formPromptBundleKey = "bundle_key"
	formPromptPath      = "path"
	formPromptFileName  = "file_name"
	formPromptContent   = "content"

	formPersonaStoreKey = "store_key"
	formPersonaID       = "persona_id"
	formPersonaName     = "persona_name"
	formPersonaPrompt   = "persona_prompt"

	personaKindBuiltIn = "built-in"
	personaKindCustom  = "custom"

	personaStoreSharedTitle   = "Shared Persona Store"
	personaStoreTitleFallback = "Persona Store"
)

type PromptsProvider interface {
	PromptsStatus() (PromptsStatus, error)
	SavePromptInline(bundleKey string, content string) error
	SavePromptRuntime(bundleKey string, content string) error
	SavePromptFile(bundleKey string, path string, content string) error
	CreatePromptFile(bundleKey string, fileName string, content string) error
	DeletePromptFile(bundleKey string, path string) error
}

type PersonasProvider interface {
	PersonasStatus() (PersonasStatus, error)
	SavePersona(
		storeKey string,
		personaID string,
		name string,
		prompt string,
	) error
	DeletePersona(storeKey string, personaID string) error
	SetDefaultPersona(personaID string) error
}

type PromptsStatus struct {
	Enabled  bool                 `json:"enabled"`
	Error    string               `json:"error,omitempty"`
	Sections []PromptSectionState `json:"sections,omitempty"`
	Previews []PromptPreviewState `json:"previews,omitempty"`
	Bundles  []PromptBundleState  `json:"bundles,omitempty"`
}

type PromptSectionState struct {
	Key     string              `json:"key,omitempty"`
	Title   string              `json:"title,omitempty"`
	Summary string              `json:"summary,omitempty"`
	Bundles []PromptBundleState `json:"bundles,omitempty"`
}

type PromptPreviewState struct {
	Key     string `json:"key,omitempty"`
	Title   string `json:"title,omitempty"`
	Summary string `json:"summary,omitempty"`
	Content string `json:"content,omitempty"`
}

type PersonasStatus struct {
	Enabled          bool               `json:"enabled"`
	Error            string             `json:"error,omitempty"`
	DefaultPersonaID string             `json:"default_persona_id,omitempty"`
	DefaultOptions   []PersonaOption    `json:"default_options,omitempty"`
	Stores           []PersonaStoreView `json:"stores,omitempty"`
}

type PromptBundleState struct {
	Key                string            `json:"key,omitempty"`
	Title              string            `json:"title,omitempty"`
	Summary            string            `json:"summary,omitempty"`
	SourceSummary      string            `json:"source_summary,omitempty"`
	ConfiguredLabel    string            `json:"configured_label,omitempty"`
	ConfiguredValue    string            `json:"configured_value,omitempty"`
	EffectiveLabel     string            `json:"effective_label,omitempty"`
	EffectiveValue     string            `json:"effective_value,omitempty"`
	InlineValue        string            `json:"inline_value,omitempty"`
	InlineEditable     bool              `json:"inline_editable"`
	RuntimeValue       string            `json:"runtime_value,omitempty"`
	RuntimeEditable    bool              `json:"runtime_editable"`
	RuntimeOverride    bool              `json:"runtime_override"`
	CreateEnabled      bool              `json:"create_enabled"`
	CreateDir          string            `json:"create_dir,omitempty"`
	LoadError          string            `json:"load_error,omitempty"`
	Files              []PromptFileState `json:"files,omitempty"`
	SupportsFileEdits  bool              `json:"supports_file_edits"`
	SupportsFileCreate bool              `json:"supports_file_create"`
	SupportsFileDelete bool              `json:"supports_file_delete"`
}

type PromptFileState struct {
	Path      string `json:"path,omitempty"`
	Label     string `json:"label,omitempty"`
	Content   string `json:"content,omitempty"`
	Error     string `json:"error,omitempty"`
	Deletable bool   `json:"deletable"`
}

type PersonaOption struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Summary string `json:"summary,omitempty"`
}

type PersonaStoreView struct {
	Key           string        `json:"key,omitempty"`
	Title         string        `json:"title,omitempty"`
	Path          string        `json:"path,omitempty"`
	UsageLabels   []string      `json:"usage_labels,omitempty"`
	CreateEnabled bool          `json:"create_enabled"`
	Personas      []PersonaView `json:"personas,omitempty"`
}

type PersonaView struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Summary   string `json:"summary,omitempty"`
	Prompt    string `json:"prompt,omitempty"`
	BuiltIn   bool   `json:"built_in"`
	Editable  bool   `json:"editable"`
	Deletable bool   `json:"deletable"`
}

func promptSections(status PromptsStatus) []PromptSectionState {
	_ = "STUB: not implemented"
	return nil
}

func promptBlockCount(status PromptsStatus) int { _ = "STUB: not implemented"; return 0 }

func hasPromptValue(raw string) bool { _ = "STUB: not implemented"; return false }

func promptValuesDiffer(left string, right string) bool { _ = "STUB: not implemented"; return false }

const promptSummaryMaxRunes = 120

func promptCollapsedSummary(raw string) string { _ = "STUB: not implemented"; return "" }

func promptSummarySnippet(raw string) string { _ = "STUB: not implemented"; return "" }

func promptLineCount(raw string) int { _ = "STUB: not implemented"; return 0 }

func promptInlineEditorTitle(bundle PromptBundleState) string { _ = "STUB: not implemented"; return "" }

func promptInlineEditorSummary(bundle PromptBundleState) string {
	_ = "STUB: not implemented"
	return ""
}

func promptRuntimeEditorSummary(bundle PromptBundleState) string {
	_ = "STUB: not implemented"
	return ""
}

func personaStoreTitle(store PersonaStoreView) string { _ = "STUB: not implemented"; return "" }

func personaStoreUsageLabels(store PersonaStoreView) []string {
	_ = "STUB: not implemented"
	return nil
}

func personaCustomPersonas(store PersonaStoreView) []PersonaView {
	_ = "STUB: not implemented"
	return nil
}

func personaBuiltInPersonas(store PersonaStoreView) []PersonaView {
	_ = "STUB: not implemented"
	return nil
}

func personaViewsByKind(
	personas []PersonaView,
	builtIn bool,
) []PersonaView {
	_ = "STUB: not implemented"
	return nil
}

func personaStoreBuiltInCount(store PersonaStoreView) int { _ = "STUB: not implemented"; return 0 }

func personaStoreCustomCount(store PersonaStoreView) int { _ = "STUB: not implemented"; return 0 }

func personaDisplayName(view PersonaView) string { _ = "STUB: not implemented"; return "" }

func personaKindLabel(view PersonaView) string { _ = "STUB: not implemented"; return "" }

func personaSummaryText(view PersonaView) string { _ = "STUB: not implemented"; return "" }

func (s *Service) promptsStatus() PromptsStatus {
	_ = "STUB: not implemented"
	return *new(PromptsStatus)
}

func (s *Service) personasStatus() PersonasStatus {
	_ = "STUB: not implemented"
	return *new(PersonasStatus)
}

func (s *Service) handlePromptsJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handlePersonasJSON(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleSavePromptInline(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleSavePromptRuntime(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleSavePromptFile(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleCreatePromptFile(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleDeletePromptFile(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleSavePersona(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleDeletePersona(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) handleSaveDefaultPersona(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = "STUB: not implemented"
	return
}

func (s *Service) requirePromptPOST(
	w http.ResponseWriter,
	r *http.Request,
) (string, string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", "", false
}

func (s *Service) requirePersonaPOST(
	w http.ResponseWriter,
	r *http.Request,
) (string, string, string, string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", "", "", "", false
}

const promptsPageTemplateHTML = `
{{define "promptsPage"}}
    <section class="panels">
      <article class="card">
        <h2>Prompt Control</h2>
        <p class="subtle">
          Edit the prompt blocks you own, inspect the assembled previews,
          and only drop to raw files when you actually need them.
        </p>
        <dl class="meta">
          <dt>Prompt Blocks</dt>
          <dd>{{promptBlockCount .Prompts}}</dd>
          <dt>Final Previews</dt>
          <dd>{{len .Prompts.Previews}}</dd>
          <dt>Personas</dt>
          <dd><a href="/personas">/personas</a></dd>
          <dt>JSON</dt>
          <dd><a href="/api/prompts">/api/prompts</a></dd>
        </dl>
      </article>
    </section>

    {{if .Prompts.Error}}
    <section class="card" style="margin-top: 24px;">
      <div class="notice err">
        {{.Prompts.Error}}
      </div>
    </section>
    {{end}}

    {{if .Prompts.Previews}}
    <section class="card" style="margin-top: 24px;">
      <h2>Final Prompt Preview</h2>
      <p class="subtle">
        Read-only previews of the prompt text produced by the editable
        prompt blocks in this runtime.
      </p>
      {{range .Prompts.Previews}}
      <details class="card prompt-detail" style="margin-top: 18px;" id="preview-{{.Key}}">
        <summary>
          <strong>{{.Title}}</strong>
          {{if .Summary}}
          <p class="subtle prompt-detail-copy">{{.Summary}}</p>
          {{end}}
          <p class="subtle prompt-detail-hint">
            {{promptCollapsedSummary .Content}}
          </p>
        </summary>
        <div class="prompt-detail-body">
          {{if .Content}}
          <div class="memory-preview">{{.Content}}</div>
          {{else}}
          <p class="empty">No prompt content is currently active.</p>
          {{end}}
        </div>
      </details>
      {{end}}
    </section>
    {{end}}

    {{if promptSections .Prompts}}
    {{range promptSections .Prompts}}
    <section class="card" style="margin-top: 24px;">
      <h2>{{.Title}}</h2>
      {{if .Summary}}
      <p class="subtle">{{.Summary}}</p>
      {{end}}
      {{range .Bundles}}
      {{$bundle := .}}
      <article class="card" style="margin-top: 18px;" id="prompt-{{.Key}}">
        <h3>{{.Title}}</h3>
        {{if .Summary}}
        <p class="subtle">{{.Summary}}</p>
        {{end}}
        <dl class="meta">
          <dt>Source Setup</dt>
          <dd>{{if .SourceSummary}}{{.SourceSummary}}{{else}}-{{end}}</dd>
          <dt>Editable Files</dt>
          <dd>{{len .Files}}</dd>
          <dt>Runtime Override</dt>
          <dd>{{if .RuntimeOverride}}active{{else}}off{{end}}</dd>
        </dl>
        {{if .LoadError}}
        <div class="notice err" style="margin-top: 12px;">
          {{.LoadError}}
        </div>
        {{end}}

        <div class="panels" style="grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));">
          <details class="card prompt-detail">
            <summary>
              <strong>
                {{if .EffectiveLabel}}{{.EffectiveLabel}}{{else}}Effective Prompt{{end}}
              </strong>
              <p class="subtle prompt-detail-hint">
                {{promptCollapsedSummary .EffectiveValue}}
              </p>
            </summary>
            <div class="prompt-detail-body">
              {{if .EffectiveValue}}
              <div class="memory-preview">{{.EffectiveValue}}</div>
              {{else}}
              <p class="empty">This block does not add any prompt text.</p>
              {{end}}
            </div>
          </details>
          {{if and (hasPromptValue .ConfiguredValue) (promptValuesDiffer .ConfiguredValue .EffectiveValue)}}
          <details class="card prompt-detail">
            <summary>
              <strong>
                {{if .ConfiguredLabel}}{{.ConfiguredLabel}}{{else}}Configured Prompt{{end}}
              </strong>
              <p class="subtle prompt-detail-hint">
                {{promptCollapsedSummary .ConfiguredValue}}
              </p>
            </summary>
            <div class="prompt-detail-body">
              <div class="memory-preview">{{.ConfiguredValue}}</div>
            </div>
          </details>
          {{end}}
        </div>

        {{if .InlineEditable}}
        <div class="card" style="margin-top: 16px;">
          <h4 style="margin: 0 0 10px;">
            {{promptInlineEditorTitle .}}
          </h4>
          <p class="subtle">
            {{promptInlineEditorSummary .}}
          </p>
          <form method="post" action="/api/prompts/inline">
            <input type="hidden" name="bundle_key" value="{{.Key}}">
            <input type="hidden" name="return_path" value="/prompts">
            <input type="hidden" name="return_to" value="prompt-{{.Key}}">
            <textarea
              name="content"
              style="width: 100%; min-height: 160px; margin-top: 12px;"
            >{{.InlineValue}}</textarea>
            <div class="actions" style="margin-top: 12px;">
              <button type="submit">Save Config Text</button>
            </div>
          </form>
        </div>
        {{end}}

        {{if .RuntimeEditable}}
        <div class="card" style="margin-top: 16px;">
          <h4 style="margin: 0 0 10px;">Runtime Override</h4>
          <p class="subtle">
            {{promptRuntimeEditorSummary .}}
          </p>
          <form method="post" action="/api/prompts/runtime">
            <input type="hidden" name="bundle_key" value="{{.Key}}">
            <input type="hidden" name="return_path" value="/prompts">
            <input type="hidden" name="return_to" value="prompt-{{.Key}}">
            <textarea
              name="content"
              style="width: 100%; min-height: 180px; margin-top: 12px;"
            >{{.RuntimeValue}}</textarea>
            <div class="actions" style="margin-top: 12px;">
              <button type="submit">Save Runtime Override</button>
            </div>
          </form>
        </div>
        {{end}}

        {{if or .CreateEnabled .Files}}
        <details class="card" style="margin-top: 16px;">
          <summary><strong>Advanced Sources</strong></summary>
          <p class="subtle" style="margin-top: 12px;">
            Manage raw files only when you want to split a prompt block
            across multiple markdown files.
          </p>
          {{if .CreateDir}}
          <p class="subtle" style="margin-top: 10px;">
            New files are created in <code>{{.CreateDir}}</code>.
          </p>
          {{end}}
          {{if .CreateEnabled}}
          <form method="post" action="/api/prompts/file/create">
            <input type="hidden" name="bundle_key" value="{{.Key}}">
            <input type="hidden" name="return_path" value="/prompts">
            <input type="hidden" name="return_to" value="prompt-{{.Key}}">
            <div class="panels" style="grid-template-columns: minmax(0, 240px) minmax(0, 1fr); margin-top: 12px;">
              <div>
                <label for="create-{{.Key}}">New file name</label>
                <input
                  id="create-{{.Key}}"
                  type="text"
                  name="file_name"
                  placeholder="local-note.md"
                  style="width: 100%; margin-top: 8px;"
                >
              </div>
              <div>
                <label for="create-body-{{.Key}}">Initial content</label>
                <textarea
                  id="create-body-{{.Key}}"
                  name="content"
                  style="width: 100%; min-height: 120px; margin-top: 8px;"
                ></textarea>
              </div>
            </div>
            <div class="actions" style="margin-top: 12px;">
              <button class="secondary" type="submit">Create File</button>
            </div>
          </form>
          {{end}}

          {{if .Files}}
          {{range .Files}}
          <div class="card" style="margin-top: 12px;">
            <div style="display: flex; gap: 12px; justify-content: space-between; align-items: baseline; flex-wrap: wrap;">
              <strong>{{.Label}}</strong>
              <code>{{.Path}}</code>
            </div>
            {{if .Error}}
            <div class="notice err" style="margin-top: 12px;">
              {{.Error}}
            </div>
            {{end}}
            <form method="post" action="/api/prompts/file" style="margin-top: 12px;">
              <input type="hidden" name="bundle_key" value="{{$bundle.Key}}">
              <input type="hidden" name="path" value="{{.Path}}">
              <input type="hidden" name="return_path" value="/prompts">
              <input type="hidden" name="return_to" value="prompt-{{$bundle.Key}}">
              <textarea
                name="content"
                style="width: 100%; min-height: 180px;"
              >{{.Content}}</textarea>
              <div class="actions" style="margin-top: 12px;">
                <button type="submit">Save File</button>
              </div>
            </form>
            {{if .Deletable}}
            <form method="post" action="/api/prompts/file/delete" style="margin-top: 8px;">
              <input type="hidden" name="bundle_key" value="{{$bundle.Key}}">
              <input type="hidden" name="path" value="{{.Path}}">
              <input type="hidden" name="return_path" value="/prompts">
              <input type="hidden" name="return_to" value="prompt-{{$bundle.Key}}">
              <div class="actions">
                <button
                  class="warn"
                  type="submit"
                  onclick="return confirm('Delete this prompt file?');"
                >
                  Delete File
                </button>
              </div>
            </form>
            {{end}}
          </div>
          {{end}}
          {{else}}
          <p class="empty">
            {{if .CreateEnabled}}
              No markdown files exist in this block yet.
            {{else}}
              This block has no editable files.
            {{end}}
          </p>
          {{end}}
        </details>
        {{end}}
      </article>
      {{end}}
    </section>
    {{end}}
    {{else}}
    <section class="card" style="margin-top: 24px;">
      <p class="empty">Prompt management is not available for this runtime.</p>
    </section>
    {{end}}
{{end}}
`

const personasPageTemplateHTML = `
{{define "personasPage"}}
    <section class="panels">
      <article class="card">
        <h2>Persona Management</h2>
        <p class="subtle">
          Manage the default persona and any file-backed persona
          definitions exposed by this runtime.
        </p>
        <dl class="meta">
          <dt>Default Persona</dt>
          <dd>
            {{if .Personas.DefaultPersonaID}}
              {{.Personas.DefaultPersonaID}}
            {{else}}
              (default)
            {{end}}
          </dd>
          <dt>Persona Stores</dt>
          <dd>{{len .Personas.Stores}}</dd>
          <dt>JSON</dt>
          <dd><a href="/api/personas">/api/personas</a></dd>
        </dl>
      </article>
    </section>

    {{if .Personas.Enabled}}
    <section class="card" style="margin-top: 24px;">
      <h2>Persona Stores</h2>
      <p class="subtle">
        Manage the default persona and any file-backed persona definitions
        exposed by this runtime.
      </p>
      {{if .Personas.Error}}
      <div class="notice err" style="margin-top: 12px;">
        {{.Personas.Error}}
      </div>
      {{end}}
      <article class="card" style="margin-top: 18px;" id="personas-default">
        <h3>Default Persona</h3>
        <form method="post" action="/api/personas/default">
          <input type="hidden" name="return_path" value="/personas">
          <input type="hidden" name="return_to" value="personas-default">
          <label for="default-persona">Persona</label>
          <select id="default-persona" name="persona_id">
            <option value="">(default)</option>
            {{range .Personas.DefaultOptions}}
            <option value="{{.ID}}"{{if eq $.Personas.DefaultPersonaID .ID}} selected{{end}}>
              {{.Name}} ({{.ID}})
            </option>
            {{end}}
          </select>
          <div class="actions" style="margin-top: 12px;">
            <button type="submit">Save Default Persona</button>
          </div>
        </form>
      </article>

      {{range .Personas.Stores}}
      {{$store := .}}
      <article class="card" style="margin-top: 18px;" id="persona-store-{{.Key}}">
        <div style="display: flex; gap: 12px; justify-content: space-between; align-items: baseline; flex-wrap: wrap;">
          <h3 style="margin: 0;">{{personaStoreTitle .}}</h3>
          <div class="skill-badges inline">
            {{if gt (personaStoreCustomCount .) 0}}
            <span class="skill-badge">
              {{personaStoreCustomCount .}} custom
            </span>
            {{end}}
            {{if gt (personaStoreBuiltInCount .) 0}}
            <span class="skill-badge">
              {{personaStoreBuiltInCount .}} built-in
            </span>
            {{end}}
          </div>
        </div>
        {{if .Path}}
        <p class="subtle" style="margin-top: 10px;">
          Store path: <code>{{.Path}}</code>
        </p>
        {{end}}
        {{with personaStoreUsageLabels .}}
        <div class="skill-badges inline" style="margin-top: 10px;">
          {{range .}}
          <span class="skill-badge">{{.}}</span>
          {{end}}
        </div>
        {{end}}

        {{if .CreateEnabled}}
        <details class="card prompt-detail" style="margin-top: 12px;">
          <summary>
            <strong>Create Persona</strong>
            <p class="subtle prompt-detail-hint">
              Add a new persona in this store without scrolling past the
              existing definitions.
            </p>
          </summary>
          <div class="prompt-detail-body">
            <form method="post" action="/api/personas/save">
              <input type="hidden" name="store_key" value="{{$store.Key}}">
              <input type="hidden" name="return_path" value="/personas">
              <input type="hidden" name="return_to" value="persona-store-{{$store.Key}}">
              <label>Name</label>
              <input type="text" name="persona_name">
              <label style="margin-top: 12px;">Prompt</label>
              <textarea
                name="persona_prompt"
                style="width: 100%; min-height: 180px;"
              ></textarea>
              <div class="actions" style="margin-top: 12px;">
                <button class="secondary" type="submit">Create Persona</button>
              </div>
            </form>
          </div>
        </details>
        {{end}}

        {{with personaCustomPersonas .}}
        <div style="margin-top: 16px;">
          <h4 style="margin: 0 0 10px;">Custom Personas</h4>
          {{range .}}
          <details class="card prompt-detail" style="margin-top: 12px;">
            <summary>
              <div style="display: flex; gap: 12px; justify-content: space-between; align-items: baseline; flex-wrap: wrap;">
                <strong>{{personaDisplayName .}}</strong>
                <code>{{.ID}}</code>
              </div>
              <div class="skill-badges inline" style="margin-top: 8px;">
                <span class="skill-badge">{{personaKindLabel .}}</span>
              </div>
              <p class="subtle prompt-detail-hint">
                {{personaSummaryText .}}
              </p>
            </summary>
            <div class="prompt-detail-body">
              {{if .Editable}}
              <form method="post" action="/api/personas/save">
                <input type="hidden" name="store_key" value="{{$store.Key}}">
                <input type="hidden" name="persona_id" value="{{.ID}}">
                <input type="hidden" name="return_path" value="/personas">
                <input type="hidden" name="return_to" value="persona-store-{{$store.Key}}">
                <label>Name</label>
                <input type="text" name="persona_name" value="{{.Name}}">
                <label style="margin-top: 12px;">Prompt</label>
                <textarea
                  name="persona_prompt"
                  style="width: 100%; min-height: 180px;"
                >{{.Prompt}}</textarea>
                <div class="actions" style="margin-top: 12px;">
                  <button type="submit">Save Persona</button>
                </div>
              </form>
              {{end}}
              {{if .Deletable}}
              <form method="post" action="/api/personas/delete" style="margin-top: 8px;">
                <input type="hidden" name="store_key" value="{{$store.Key}}">
                <input type="hidden" name="persona_id" value="{{.ID}}">
                <input type="hidden" name="return_path" value="/personas">
                <input type="hidden" name="return_to" value="persona-store-{{$store.Key}}">
                <div class="actions">
                  <button
                    class="warn"
                    type="submit"
                    onclick="return confirm('Delete this persona override?');"
                  >
                    Delete Persona
                  </button>
                </div>
              </form>
              {{end}}
            </div>
          </details>
          {{end}}
        </div>
        {{end}}

        {{with personaBuiltInPersonas .}}
        <div style="margin-top: 16px;">
          <h4 style="margin: 0 0 10px;">Built-in Personas</h4>
          {{range .}}
          <details class="card prompt-detail" style="margin-top: 12px;">
            <summary>
              <div style="display: flex; gap: 12px; justify-content: space-between; align-items: baseline; flex-wrap: wrap;">
                <strong>{{personaDisplayName .}}</strong>
                <code>{{.ID}}</code>
              </div>
              <div class="skill-badges inline" style="margin-top: 8px;">
                <span class="skill-badge">{{personaKindLabel .}}</span>
              </div>
              <p class="subtle prompt-detail-hint">
                {{personaSummaryText .}}
              </p>
            </summary>
            <div class="prompt-detail-body">
              {{if .Editable}}
              <form method="post" action="/api/personas/save">
                <input type="hidden" name="store_key" value="{{$store.Key}}">
                <input type="hidden" name="persona_id" value="{{.ID}}">
                <input type="hidden" name="return_path" value="/personas">
                <input type="hidden" name="return_to" value="persona-store-{{$store.Key}}">
                <label>Name</label>
                <input type="text" name="persona_name" value="{{.Name}}">
                <label style="margin-top: 12px;">Prompt</label>
                <textarea
                  name="persona_prompt"
                  style="width: 100%; min-height: 180px;"
                >{{.Prompt}}</textarea>
                <div class="actions" style="margin-top: 12px;">
                  <button type="submit">Save Persona</button>
                </div>
              </form>
              {{end}}
              {{if .Deletable}}
              <form method="post" action="/api/personas/delete" style="margin-top: 8px;">
                <input type="hidden" name="store_key" value="{{$store.Key}}">
                <input type="hidden" name="persona_id" value="{{.ID}}">
                <input type="hidden" name="return_path" value="/personas">
                <input type="hidden" name="return_to" value="persona-store-{{$store.Key}}">
                <div class="actions">
                  <button
                    class="warn"
                    type="submit"
                    onclick="return confirm('Delete this persona override?');"
                  >
                    Delete Persona
                  </button>
                </div>
              </form>
              {{end}}
            </div>
          </details>
          {{end}}
        </div>
        {{end}}
      </article>
      {{end}}
    </section>
    {{end}}
{{end}}
`
