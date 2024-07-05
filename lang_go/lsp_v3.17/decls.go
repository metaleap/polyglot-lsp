// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-lsp/gen/cmd/gen_lsp
package lsp_v3_17

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strconv"
	"sync"
	"time"
)

/*TOr*/
type WorkspaceFullDocumentDiagnosticReportOrWorkspaceUnchangedDocumentDiagnosticReport struct {
	WorkspaceFullDocumentDiagnosticReport/*TOpt*/ *WorkspaceFullDocumentDiagnosticReport
	WorkspaceUnchangedDocumentDiagnosticReport/*TOpt*/ *WorkspaceUnchangedDocumentDiagnosticReport
}

func (it *WorkspaceFullDocumentDiagnosticReportOrWorkspaceUnchangedDocumentDiagnosticReport) MarshalJSON() ([]byte, error) {
	if !isNone(it.WorkspaceFullDocumentDiagnosticReport) {
		return json.Marshal(it.WorkspaceFullDocumentDiagnosticReport)
	}
	if !isNone(it.WorkspaceUnchangedDocumentDiagnosticReport) {
		return json.Marshal(it.WorkspaceUnchangedDocumentDiagnosticReport)
	}

	return []byte("null"), nil
}
func (it *WorkspaceFullDocumentDiagnosticReportOrWorkspaceUnchangedDocumentDiagnosticReport) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *WorkspaceFullDocumentDiagnosticReport
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.WorkspaceFullDocumentDiagnosticReport = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *WorkspaceUnchangedDocumentDiagnosticReport
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.WorkspaceUnchangedDocumentDiagnosticReport = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type StringOrInlayHintLabelParts struct {
	String/*TOpt*/ *String
	InlayHintLabelParts/*TOpt*/ []InlayHintLabelPart
}

func (it *StringOrInlayHintLabelParts) MarshalJSON() ([]byte, error) {
	if !isNone(it.String) {
		return json.Marshal(it.String)
	}
	if !isNone(it.InlayHintLabelParts) {
		return json.Marshal(it.InlayHintLabelParts)
	}

	return []byte("null"), nil
}
func (it *StringOrInlayHintLabelParts) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *String
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.String = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ []InlayHintLabelPart
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.InlayHintLabelParts = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	TypeHierarchyOptions/*TOpt*/ *TypeHierarchyOptions
	TypeHierarchyRegistrationOptions/*TOpt*/ *TypeHierarchyRegistrationOptions
}

func (it *BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.TypeHierarchyOptions) {
		return json.Marshal(it.TypeHierarchyOptions)
	}
	if !isNone(it.TypeHierarchyRegistrationOptions) {
		return json.Marshal(it.TypeHierarchyRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *TypeHierarchyOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.TypeHierarchyOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *TypeHierarchyRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.TypeHierarchyRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type WorkspaceFolderOrURI struct {
	WorkspaceFolder/*TOpt*/ *WorkspaceFolder
	URI/*TOpt*/ *URI
}

func (it *WorkspaceFolderOrURI) MarshalJSON() ([]byte, error) {
	if !isNone(it.WorkspaceFolder) {
		return json.Marshal(it.WorkspaceFolder)
	}
	if !isNone(it.URI) {
		return json.Marshal(it.URI)
	}

	return []byte("null"), nil
}
func (it *WorkspaceFolderOrURI) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *WorkspaceFolder
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.WorkspaceFolder = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *URI
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.URI = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type TextEditOrInsertReplaceEdit struct {
	TextEdit/*TOpt*/ *TextEdit
	InsertReplaceEdit/*TOpt*/ *InsertReplaceEdit
}

func (it *TextEditOrInsertReplaceEdit) MarshalJSON() ([]byte, error) {
	if !isNone(it.TextEdit) {
		return json.Marshal(it.TextEdit)
	}
	if !isNone(it.InsertReplaceEdit) {
		return json.Marshal(it.InsertReplaceEdit)
	}

	return []byte("null"), nil
}
func (it *TextEditOrInsertReplaceEdit) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *TextEdit
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.TextEdit = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *InsertReplaceEdit
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.InsertReplaceEdit = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions struct {
	NotebookDocumentSyncOptions/*TOpt*/ *NotebookDocumentSyncOptions
	NotebookDocumentSyncRegistrationOptions/*TOpt*/ *NotebookDocumentSyncRegistrationOptions
}

func (it *NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.NotebookDocumentSyncOptions) {
		return json.Marshal(it.NotebookDocumentSyncOptions)
	}
	if !isNone(it.NotebookDocumentSyncRegistrationOptions) {
		return json.Marshal(it.NotebookDocumentSyncRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *NotebookDocumentSyncOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.NotebookDocumentSyncOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *NotebookDocumentSyncRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.NotebookDocumentSyncRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrCodeActionOptions struct {
	Boolean/*TOpt*/ *Boolean
	CodeActionOptions/*TOpt*/ *CodeActionOptions
}

func (it *BooleanOrCodeActionOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.CodeActionOptions) {
		return json.Marshal(it.CodeActionOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrCodeActionOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *CodeActionOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.CodeActionOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type PatternOrRelativePattern struct {
	Pattern/*TOpt*/ *String
	RelativePattern/*TOpt*/ *RelativePattern
}

func (it *PatternOrRelativePattern) MarshalJSON() ([]byte, error) {
	if !isNone(it.Pattern) {
		return json.Marshal(it.Pattern)
	}
	if !isNone(it.RelativePattern) {
		return json.Marshal(it.RelativePattern)
	}

	return []byte("null"), nil
}
func (it *PatternOrRelativePattern) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *String
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Pattern = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *RelativePattern
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.RelativePattern = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type LocationOrUriDocumentUri struct {
	Location/*TOpt*/ *Location
	UriDocumentUri/*TOpt*/ * /*TStruc*/ struct {
		Uri DocumentURI `json:"uri,omitempty"`
	}
}

func (it *LocationOrUriDocumentUri) MarshalJSON() ([]byte, error) {
	if !isNone(it.Location) {
		return json.Marshal(it.Location)
	}
	if !isNone(it.UriDocumentUri) {
		return json.Marshal(it.UriDocumentUri)
	}

	return []byte("null"), nil
}
func (it *LocationOrUriDocumentUri) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Location
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Location = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ * /*TStruc*/ struct {
			Uri DocumentURI `json:"uri,omitempty"`
		}
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.UriDocumentUri = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrImplementationOptionsOrImplementationRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	ImplementationOptions/*TOpt*/ *ImplementationOptions
	ImplementationRegistrationOptions/*TOpt*/ *ImplementationRegistrationOptions
}

func (it *BooleanOrImplementationOptionsOrImplementationRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.ImplementationOptions) {
		return json.Marshal(it.ImplementationOptions)
	}
	if !isNone(it.ImplementationRegistrationOptions) {
		return json.Marshal(it.ImplementationRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrImplementationOptionsOrImplementationRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *ImplementationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.ImplementationOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *ImplementationRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.ImplementationRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type StringOrLanguageStringWithValueString struct {
	String/*TOpt*/ *String
	LanguageStringWithValueString/*TOpt*/ * /*TStruc*/ struct {
		Language string `json:"language,omitempty"`

		Value string `json:"value,omitempty"`
	}
}

func (it *StringOrLanguageStringWithValueString) MarshalJSON() ([]byte, error) {
	if !isNone(it.String) {
		return json.Marshal(it.String)
	}
	if !isNone(it.LanguageStringWithValueString) {
		return json.Marshal(it.LanguageStringWithValueString)
	}

	return []byte("null"), nil
}
func (it *StringOrLanguageStringWithValueString) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *String
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.String = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ * /*TStruc*/ struct {
			Language string `json:"language,omitempty"`

			Value string `json:"value,omitempty"`
		}
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.LanguageStringWithValueString = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type StringOrMarkupContent struct {
	String/*TOpt*/ *String
	MarkupContent/*TOpt*/ *MarkupContent
}

func (it *StringOrMarkupContent) MarshalJSON() ([]byte, error) {
	if !isNone(it.String) {
		return json.Marshal(it.String)
	}
	if !isNone(it.MarkupContent) {
		return json.Marshal(it.MarkupContent)
	}

	return []byte("null"), nil
}
func (it *StringOrMarkupContent) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *String
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.String = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *MarkupContent
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.MarkupContent = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrDefinitionOptions struct {
	Boolean/*TOpt*/ *Boolean
	DefinitionOptions/*TOpt*/ *DefinitionOptions
}

func (it *BooleanOrDefinitionOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.DefinitionOptions) {
		return json.Marshal(it.DefinitionOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrDefinitionOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *DefinitionOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DefinitionOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type SemanticTokensOptionsOrSemanticTokensRegistrationOptions struct {
	SemanticTokensOptions/*TOpt*/ *SemanticTokensOptions
	SemanticTokensRegistrationOptions/*TOpt*/ *SemanticTokensRegistrationOptions
}

func (it *SemanticTokensOptionsOrSemanticTokensRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.SemanticTokensOptions) {
		return json.Marshal(it.SemanticTokensOptions)
	}
	if !isNone(it.SemanticTokensRegistrationOptions) {
		return json.Marshal(it.SemanticTokensRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *SemanticTokensOptionsOrSemanticTokensRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *SemanticTokensOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.SemanticTokensOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *SemanticTokensRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.SemanticTokensRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type InlineValueTextOrInlineValueVariableLookupOrInlineValueEvaluatableExpression struct {
	InlineValueText/*TOpt*/ *InlineValueText
	InlineValueVariableLookup/*TOpt*/ *InlineValueVariableLookup
	InlineValueEvaluatableExpression/*TOpt*/ *InlineValueEvaluatableExpression
}

func (it *InlineValueTextOrInlineValueVariableLookupOrInlineValueEvaluatableExpression) MarshalJSON() ([]byte, error) {
	if !isNone(it.InlineValueText) {
		return json.Marshal(it.InlineValueText)
	}
	if !isNone(it.InlineValueVariableLookup) {
		return json.Marshal(it.InlineValueVariableLookup)
	}
	if !isNone(it.InlineValueEvaluatableExpression) {
		return json.Marshal(it.InlineValueEvaluatableExpression)
	}

	return []byte("null"), nil
}
func (it *InlineValueTextOrInlineValueVariableLookupOrInlineValueEvaluatableExpression) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *InlineValueText
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.InlineValueText = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *InlineValueVariableLookup
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.InlineValueVariableLookup = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *InlineValueEvaluatableExpression
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.InlineValueEvaluatableExpression = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type DefinitionOrDefinitionLinksOrNull struct {
	Definition/*TOpt*/ Definition
	DefinitionLinks/*TOpt*/ []DefinitionLink
}

func (it *DefinitionOrDefinitionLinksOrNull) MarshalJSON() ([]byte, error) {
	if !isNone(it.Definition) {
		return json.Marshal(it.Definition)
	}
	if !isNone(it.DefinitionLinks) {
		return json.Marshal(it.DefinitionLinks)
	}

	return []byte("null"), nil
}
func (it *DefinitionOrDefinitionLinksOrNull) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ Definition
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Definition = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ []DefinitionLink
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DefinitionLinks = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrAnyByString struct {
	Boolean/*TOpt*/ *Boolean
	AnyByString/*TOpt*/ map[string]any
}

func (it *BooleanOrAnyByString) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.AnyByString) {
		return json.Marshal(it.AnyByString)
	}

	return []byte("null"), nil
}
func (it *BooleanOrAnyByString) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ map[string]any
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.AnyByString = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type RangeOrRangeWithPlaceholderStringOrDefaultBehaviorBoolean struct {
	Range/*TOpt*/ *Range
	RangeWithPlaceholderString/*TOpt*/ * /*TStruc*/ struct {
		Range Range `json:"range,omitempty"`

		Placeholder string `json:"placeholder,omitempty"`
	}
	DefaultBehaviorBoolean/*TOpt*/ * /*TStruc*/ struct {
		DefaultBehavior bool `json:"defaultBehavior,omitempty"`
	}
}

func (it *RangeOrRangeWithPlaceholderStringOrDefaultBehaviorBoolean) MarshalJSON() ([]byte, error) {
	if !isNone(it.Range) {
		return json.Marshal(it.Range)
	}
	if !isNone(it.RangeWithPlaceholderString) {
		return json.Marshal(it.RangeWithPlaceholderString)
	}
	if !isNone(it.DefaultBehaviorBoolean) {
		return json.Marshal(it.DefaultBehaviorBoolean)
	}

	return []byte("null"), nil
}
func (it *RangeOrRangeWithPlaceholderStringOrDefaultBehaviorBoolean) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Range
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Range = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ * /*TStruc*/ struct {
			Range Range `json:"range,omitempty"`

			Placeholder string `json:"placeholder,omitempty"`
		}
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.RangeWithPlaceholderString = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ * /*TStruc*/ struct {
			DefaultBehavior bool `json:"defaultBehavior,omitempty"`
		}
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DefaultBehaviorBoolean = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type SemanticTokensOrSemanticTokensDeltaOrNull struct {
	SemanticTokens/*TOpt*/ *SemanticTokens
	SemanticTokensDelta/*TOpt*/ *SemanticTokensDelta
}

func (it *SemanticTokensOrSemanticTokensDeltaOrNull) MarshalJSON() ([]byte, error) {
	if !isNone(it.SemanticTokens) {
		return json.Marshal(it.SemanticTokens)
	}
	if !isNone(it.SemanticTokensDelta) {
		return json.Marshal(it.SemanticTokensDelta)
	}

	return []byte("null"), nil
}
func (it *SemanticTokensOrSemanticTokensDeltaOrNull) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *SemanticTokens
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.SemanticTokens = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *SemanticTokensDelta
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.SemanticTokensDelta = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type SymbolInformationsOrWorkspaceSymbolsOrNull struct {
	SymbolInformations/*TOpt*/ []SymbolInformation
	WorkspaceSymbols/*TOpt*/ []WorkspaceSymbol
}

func (it *SymbolInformationsOrWorkspaceSymbolsOrNull) MarshalJSON() ([]byte, error) {
	if !isNone(it.SymbolInformations) {
		return json.Marshal(it.SymbolInformations)
	}
	if !isNone(it.WorkspaceSymbols) {
		return json.Marshal(it.WorkspaceSymbols)
	}

	return []byte("null"), nil
}
func (it *SymbolInformationsOrWorkspaceSymbolsOrNull) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ []SymbolInformation
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.SymbolInformations = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ []WorkspaceSymbol
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.WorkspaceSymbols = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type RangeOrInsertRangeWithReplaceRange struct {
	Range/*TOpt*/ *Range
	InsertRangeWithReplaceRange/*TOpt*/ * /*TStruc*/ struct {
		Insert Range `json:"insert,omitempty"`

		Replace Range `json:"replace,omitempty"`
	}
}

func (it *RangeOrInsertRangeWithReplaceRange) MarshalJSON() ([]byte, error) {
	if !isNone(it.Range) {
		return json.Marshal(it.Range)
	}
	if !isNone(it.InsertRangeWithReplaceRange) {
		return json.Marshal(it.InsertRangeWithReplaceRange)
	}

	return []byte("null"), nil
}
func (it *RangeOrInsertRangeWithReplaceRange) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Range
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Range = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ * /*TStruc*/ struct {
			Insert Range `json:"insert,omitempty"`

			Replace Range `json:"replace,omitempty"`
		}
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.InsertRangeWithReplaceRange = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrRenameOptions struct {
	Boolean/*TOpt*/ *Boolean
	RenameOptions/*TOpt*/ *RenameOptions
}

func (it *BooleanOrRenameOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.RenameOptions) {
		return json.Marshal(it.RenameOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrRenameOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *RenameOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.RenameOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile struct {
	TextDocumentEdit/*TOpt*/ *TextDocumentEdit
	CreateFile/*TOpt*/ *CreateFile
	RenameFile/*TOpt*/ *RenameFile
	DeleteFile/*TOpt*/ *DeleteFile
}

func (it *TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile) MarshalJSON() ([]byte, error) {
	if !isNone(it.TextDocumentEdit) {
		return json.Marshal(it.TextDocumentEdit)
	}
	if !isNone(it.CreateFile) {
		return json.Marshal(it.CreateFile)
	}
	if !isNone(it.RenameFile) {
		return json.Marshal(it.RenameFile)
	}
	if !isNone(it.DeleteFile) {
		return json.Marshal(it.DeleteFile)
	}

	return []byte("null"), nil
}
func (it *TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *TextDocumentEdit
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.TextDocumentEdit = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *CreateFile
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.CreateFile = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *RenameFile
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.RenameFile = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *DeleteFile
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DeleteFile = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type RangeWithRangeLengthUintegerWithTextStringOrTextString struct {
	RangeWithRangeLengthUintegerWithTextString/*TOpt*/ * /*TStruc*/ struct {

		// The range of the document that changed.
		Range Range `json:"range,omitempty"`

		// The optional length of the range that got replaced.
		//
		// @deprecated use range instead.
		RangeLength/*TOpt*/ *Uinteger `json:"rangeLength,omitempty"`

		// The new text for the provided range.
		Text string `json:"text,omitempty"`
	}
	TextString/*TOpt*/ * /*TStruc*/ struct {

		// The new text of the whole document.
		Text string `json:"text,omitempty"`
	}
}

func (it *RangeWithRangeLengthUintegerWithTextStringOrTextString) MarshalJSON() ([]byte, error) {
	if !isNone(it.RangeWithRangeLengthUintegerWithTextString) {
		return json.Marshal(it.RangeWithRangeLengthUintegerWithTextString)
	}
	if !isNone(it.TextString) {
		return json.Marshal(it.TextString)
	}

	return []byte("null"), nil
}
func (it *RangeWithRangeLengthUintegerWithTextStringOrTextString) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ * /*TStruc*/ struct {

			// The range of the document that changed.
			Range Range `json:"range,omitempty"`

			// The optional length of the range that got replaced.
			//
			// @deprecated use range instead.
			RangeLength/*TOpt*/ *Uinteger `json:"rangeLength,omitempty"`

			// The new text for the provided range.
			Text string `json:"text,omitempty"`
		}
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.RangeWithRangeLengthUintegerWithTextString = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ * /*TStruc*/ struct {

			// The new text of the whole document.
			Text string `json:"text,omitempty"`
		}
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.TextString = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type MarkupContentOrMarkedStringOrMarkedStrings struct {
	MarkupContent/*TOpt*/ *MarkupContent
	MarkedString/*TOpt*/ MarkedString
	MarkedStrings/*TOpt*/ []MarkedString
}

func (it *MarkupContentOrMarkedStringOrMarkedStrings) MarshalJSON() ([]byte, error) {
	if !isNone(it.MarkupContent) {
		return json.Marshal(it.MarkupContent)
	}
	if !isNone(it.MarkedString) {
		return json.Marshal(it.MarkedString)
	}
	if !isNone(it.MarkedStrings) {
		return json.Marshal(it.MarkedStrings)
	}

	return []byte("null"), nil
}
func (it *MarkupContentOrMarkedStringOrMarkedStrings) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *MarkupContent
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.MarkupContent = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ MarkedString
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.MarkedString = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ []MarkedString
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.MarkedStrings = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrDeltaBoolean struct {
	Boolean/*TOpt*/ *Boolean
	DeltaBoolean/*TOpt*/ * /*TStruc*/ struct {

		// The server supports deltas for full documents.
		Delta /*TOpt*/ *Boolean `json:"delta,omitempty"`
	}
}

func (it *BooleanOrDeltaBoolean) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.DeltaBoolean) {
		return json.Marshal(it.DeltaBoolean)
	}

	return []byte("null"), nil
}
func (it *BooleanOrDeltaBoolean) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ * /*TStruc*/ struct {

			// The server supports deltas for full documents.
			Delta /*TOpt*/ *Boolean `json:"delta,omitempty"`
		}
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DeltaBoolean = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type DiagnosticOptionsOrDiagnosticRegistrationOptions struct {
	DiagnosticOptions/*TOpt*/ *DiagnosticOptions
	DiagnosticRegistrationOptions/*TOpt*/ *DiagnosticRegistrationOptions
}

func (it *DiagnosticOptionsOrDiagnosticRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.DiagnosticOptions) {
		return json.Marshal(it.DiagnosticOptions)
	}
	if !isNone(it.DiagnosticRegistrationOptions) {
		return json.Marshal(it.DiagnosticRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *DiagnosticOptionsOrDiagnosticRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *DiagnosticOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DiagnosticOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *DiagnosticRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DiagnosticRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrSaveOptions struct {
	Boolean/*TOpt*/ *Boolean
	SaveOptions/*TOpt*/ *SaveOptions
}

func (it *BooleanOrSaveOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.SaveOptions) {
		return json.Marshal(it.SaveOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrSaveOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *SaveOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.SaveOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type CommandOrCodeAction struct {
	Command/*TOpt*/ *Command
	CodeAction/*TOpt*/ *CodeAction
}

func (it *CommandOrCodeAction) MarshalJSON() ([]byte, error) {
	if !isNone(it.Command) {
		return json.Marshal(it.Command)
	}
	if !isNone(it.CodeAction) {
		return json.Marshal(it.CodeAction)
	}

	return []byte("null"), nil
}
func (it *CommandOrCodeAction) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Command
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Command = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *CodeAction
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.CodeAction = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type StringOrStrings struct {
	String/*TOpt*/ *String
	Strings/*TOpt*/ []string
}

func (it *StringOrStrings) MarshalJSON() ([]byte, error) {
	if !isNone(it.String) {
		return json.Marshal(it.String)
	}
	if !isNone(it.Strings) {
		return json.Marshal(it.Strings)
	}

	return []byte("null"), nil
}
func (it *StringOrStrings) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *String
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.String = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ []string
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Strings = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrDocumentHighlightOptions struct {
	Boolean/*TOpt*/ *Boolean
	DocumentHighlightOptions/*TOpt*/ *DocumentHighlightOptions
}

func (it *BooleanOrDocumentHighlightOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.DocumentHighlightOptions) {
		return json.Marshal(it.DocumentHighlightOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrDocumentHighlightOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *DocumentHighlightOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DocumentHighlightOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	CallHierarchyOptions/*TOpt*/ *CallHierarchyOptions
	CallHierarchyRegistrationOptions/*TOpt*/ *CallHierarchyRegistrationOptions
}

func (it *BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.CallHierarchyOptions) {
		return json.Marshal(it.CallHierarchyOptions)
	}
	if !isNone(it.CallHierarchyRegistrationOptions) {
		return json.Marshal(it.CallHierarchyRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *CallHierarchyOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.CallHierarchyOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *CallHierarchyRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.CallHierarchyRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type StringOrUintegerWithUinteger struct {
	String/*TOpt*/ *String
	UintegerWithUinteger/*TOpt*/ * /*TTup*/ []uint
}

func (it *StringOrUintegerWithUinteger) MarshalJSON() ([]byte, error) {
	if !isNone(it.String) {
		return json.Marshal(it.String)
	}
	if !isNone(it.UintegerWithUinteger) {
		return json.Marshal(it.UintegerWithUinteger)
	}

	return []byte("null"), nil
}
func (it *StringOrUintegerWithUinteger) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *String
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.String = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ * /*TTup*/ []uint
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.UintegerWithUinteger = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type RelatedFullDocumentDiagnosticReportOrRelatedUnchangedDocumentDiagnosticReport struct {
	RelatedFullDocumentDiagnosticReport/*TOpt*/ *RelatedFullDocumentDiagnosticReport
	RelatedUnchangedDocumentDiagnosticReport/*TOpt*/ *RelatedUnchangedDocumentDiagnosticReport
}

func (it *RelatedFullDocumentDiagnosticReportOrRelatedUnchangedDocumentDiagnosticReport) MarshalJSON() ([]byte, error) {
	if !isNone(it.RelatedFullDocumentDiagnosticReport) {
		return json.Marshal(it.RelatedFullDocumentDiagnosticReport)
	}
	if !isNone(it.RelatedUnchangedDocumentDiagnosticReport) {
		return json.Marshal(it.RelatedUnchangedDocumentDiagnosticReport)
	}

	return []byte("null"), nil
}
func (it *RelatedFullDocumentDiagnosticReportOrRelatedUnchangedDocumentDiagnosticReport) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *RelatedFullDocumentDiagnosticReport
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.RelatedFullDocumentDiagnosticReport = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *RelatedUnchangedDocumentDiagnosticReport
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.RelatedUnchangedDocumentDiagnosticReport = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type TextDocumentSyncOptionsOrTextDocumentSyncKind struct {
	TextDocumentSyncOptions/*TOpt*/ *TextDocumentSyncOptions
	TextDocumentSyncKind/*TOpt*/ TextDocumentSyncKind
}

func (it *TextDocumentSyncOptionsOrTextDocumentSyncKind) MarshalJSON() ([]byte, error) {
	if !isNone(it.TextDocumentSyncOptions) {
		return json.Marshal(it.TextDocumentSyncOptions)
	}
	if !isNone(it.TextDocumentSyncKind) {
		return json.Marshal(it.TextDocumentSyncKind)
	}

	return []byte("null"), nil
}
func (it *TextDocumentSyncOptionsOrTextDocumentSyncKind) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *TextDocumentSyncOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.TextDocumentSyncOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ TextDocumentSyncKind
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.TextDocumentSyncKind = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type StringOrNotebookDocumentFilter struct {
	String/*TOpt*/ *String
	NotebookDocumentFilter/*TOpt*/ *NotebookDocumentFilter
}

func (it *StringOrNotebookDocumentFilter) MarshalJSON() ([]byte, error) {
	if !isNone(it.String) {
		return json.Marshal(it.String)
	}
	if !isNone(it.NotebookDocumentFilter) {
		return json.Marshal(it.NotebookDocumentFilter)
	}

	return []byte("null"), nil
}
func (it *StringOrNotebookDocumentFilter) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *String
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.String = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *NotebookDocumentFilter
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.NotebookDocumentFilter = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrDeltaBoolean_ struct {
	Boolean/*TOpt*/ *Boolean
	DeltaBoolean/*TOpt*/ * /*TStruc*/ struct {

		// The client will send the `textDocument/semanticTokens/full/delta` request if
		// the server provides a corresponding handler.
		Delta /*TOpt*/ *Boolean `json:"delta,omitempty"`
	}
}

func (it *BooleanOrDeltaBoolean_) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.DeltaBoolean) {
		return json.Marshal(it.DeltaBoolean)
	}

	return []byte("null"), nil
}
func (it *BooleanOrDeltaBoolean_) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ * /*TStruc*/ struct {

			// The client will send the `textDocument/semanticTokens/full/delta` request if
			// the server provides a corresponding handler.
			Delta /*TOpt*/ *Boolean `json:"delta,omitempty"`
		}
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DeltaBoolean = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type LocationOrLocations struct {
	Location/*TOpt*/ *Location
	Locations/*TOpt*/ []Location
}

func (it *LocationOrLocations) MarshalJSON() ([]byte, error) {
	if !isNone(it.Location) {
		return json.Marshal(it.Location)
	}
	if !isNone(it.Locations) {
		return json.Marshal(it.Locations)
	}

	return []byte("null"), nil
}
func (it *LocationOrLocations) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Location
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Location = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ []Location
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Locations = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrDocumentRangeFormattingOptions struct {
	Boolean/*TOpt*/ *Boolean
	DocumentRangeFormattingOptions/*TOpt*/ *DocumentRangeFormattingOptions
}

func (it *BooleanOrDocumentRangeFormattingOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.DocumentRangeFormattingOptions) {
		return json.Marshal(it.DocumentRangeFormattingOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrDocumentRangeFormattingOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *DocumentRangeFormattingOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DocumentRangeFormattingOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	InlineValueOptions/*TOpt*/ *InlineValueOptions
	InlineValueRegistrationOptions/*TOpt*/ *InlineValueRegistrationOptions
}

func (it *BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.InlineValueOptions) {
		return json.Marshal(it.InlineValueOptions)
	}
	if !isNone(it.InlineValueRegistrationOptions) {
		return json.Marshal(it.InlineValueRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *InlineValueOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.InlineValueOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *InlineValueRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.InlineValueRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrWorkspaceSymbolOptions struct {
	Boolean/*TOpt*/ *Boolean
	WorkspaceSymbolOptions/*TOpt*/ *WorkspaceSymbolOptions
}

func (it *BooleanOrWorkspaceSymbolOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.WorkspaceSymbolOptions) {
		return json.Marshal(it.WorkspaceSymbolOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrWorkspaceSymbolOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *WorkspaceSymbolOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.WorkspaceSymbolOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type IntegerOrString struct {
	Integer/*TOpt*/ *Integer
	String/*TOpt*/ *String
}

func (it *IntegerOrString) MarshalJSON() ([]byte, error) {
	if !isNone(it.Integer) {
		return json.Marshal(it.Integer)
	}
	if !isNone(it.String) {
		return json.Marshal(it.String)
	}

	return []byte("null"), nil
}
func (it *IntegerOrString) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Integer
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Integer = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *String
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.String = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type TextDocumentFilterOrNotebookCellTextDocumentFilter struct {
	TextDocumentFilter/*TOpt*/ *TextDocumentFilter
	NotebookCellTextDocumentFilter/*TOpt*/ *NotebookCellTextDocumentFilter
}

func (it *TextDocumentFilterOrNotebookCellTextDocumentFilter) MarshalJSON() ([]byte, error) {
	if !isNone(it.TextDocumentFilter) {
		return json.Marshal(it.TextDocumentFilter)
	}
	if !isNone(it.NotebookCellTextDocumentFilter) {
		return json.Marshal(it.NotebookCellTextDocumentFilter)
	}

	return []byte("null"), nil
}
func (it *TextDocumentFilterOrNotebookCellTextDocumentFilter) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *TextDocumentFilter
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.TextDocumentFilter = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *NotebookCellTextDocumentFilter
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.NotebookCellTextDocumentFilter = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type CompletionItemsOrCompletionListOrNull struct {
	CompletionItems/*TOpt*/ []CompletionItem
	CompletionList/*TOpt*/ *CompletionList
}

func (it *CompletionItemsOrCompletionListOrNull) MarshalJSON() ([]byte, error) {
	if !isNone(it.CompletionItems) {
		return json.Marshal(it.CompletionItems)
	}
	if !isNone(it.CompletionList) {
		return json.Marshal(it.CompletionList)
	}

	return []byte("null"), nil
}
func (it *CompletionItemsOrCompletionListOrNull) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ []CompletionItem
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.CompletionItems = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *CompletionList
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.CompletionList = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type TextEditOrAnnotatedTextEdit struct {
	TextEdit/*TOpt*/ *TextEdit
	AnnotatedTextEdit/*TOpt*/ *AnnotatedTextEdit
}

func (it *TextEditOrAnnotatedTextEdit) MarshalJSON() ([]byte, error) {
	if !isNone(it.TextEdit) {
		return json.Marshal(it.TextEdit)
	}
	if !isNone(it.AnnotatedTextEdit) {
		return json.Marshal(it.AnnotatedTextEdit)
	}

	return []byte("null"), nil
}
func (it *TextEditOrAnnotatedTextEdit) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *TextEdit
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.TextEdit = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *AnnotatedTextEdit
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.AnnotatedTextEdit = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	InlayHintOptions/*TOpt*/ *InlayHintOptions
	InlayHintRegistrationOptions/*TOpt*/ *InlayHintRegistrationOptions
}

func (it *BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.InlayHintOptions) {
		return json.Marshal(it.InlayHintOptions)
	}
	if !isNone(it.InlayHintRegistrationOptions) {
		return json.Marshal(it.InlayHintRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *InlayHintOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.InlayHintOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *InlayHintRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.InlayHintRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrDocumentFormattingOptions struct {
	Boolean/*TOpt*/ *Boolean
	DocumentFormattingOptions/*TOpt*/ *DocumentFormattingOptions
}

func (it *BooleanOrDocumentFormattingOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.DocumentFormattingOptions) {
		return json.Marshal(it.DocumentFormattingOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrDocumentFormattingOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *DocumentFormattingOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DocumentFormattingOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	FoldingRangeOptions/*TOpt*/ *FoldingRangeOptions
	FoldingRangeRegistrationOptions/*TOpt*/ *FoldingRangeRegistrationOptions
}

func (it *BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.FoldingRangeOptions) {
		return json.Marshal(it.FoldingRangeOptions)
	}
	if !isNone(it.FoldingRangeRegistrationOptions) {
		return json.Marshal(it.FoldingRangeRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *FoldingRangeOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.FoldingRangeOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *FoldingRangeRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.FoldingRangeRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type DeclarationOrDeclarationLinksOrNull struct {
	Declaration/*TOpt*/ Declaration
	DeclarationLinks/*TOpt*/ []DeclarationLink
}

func (it *DeclarationOrDeclarationLinksOrNull) MarshalJSON() ([]byte, error) {
	if !isNone(it.Declaration) {
		return json.Marshal(it.Declaration)
	}
	if !isNone(it.DeclarationLinks) {
		return json.Marshal(it.DeclarationLinks)
	}

	return []byte("null"), nil
}
func (it *DeclarationOrDeclarationLinksOrNull) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ Declaration
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Declaration = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ []DeclarationLink
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DeclarationLinks = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport struct {
	FullDocumentDiagnosticReport/*TOpt*/ *FullDocumentDiagnosticReport
	UnchangedDocumentDiagnosticReport/*TOpt*/ *UnchangedDocumentDiagnosticReport
}

func (it *FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport) MarshalJSON() ([]byte, error) {
	if !isNone(it.FullDocumentDiagnosticReport) {
		return json.Marshal(it.FullDocumentDiagnosticReport)
	}
	if !isNone(it.UnchangedDocumentDiagnosticReport) {
		return json.Marshal(it.UnchangedDocumentDiagnosticReport)
	}

	return []byte("null"), nil
}
func (it *FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *FullDocumentDiagnosticReport
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.FullDocumentDiagnosticReport = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *UnchangedDocumentDiagnosticReport
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.UnchangedDocumentDiagnosticReport = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrHoverOptions struct {
	Boolean/*TOpt*/ *Boolean
	HoverOptions/*TOpt*/ *HoverOptions
}

func (it *BooleanOrHoverOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.HoverOptions) {
		return json.Marshal(it.HoverOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrHoverOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *HoverOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.HoverOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	DeclarationOptions/*TOpt*/ *DeclarationOptions
	DeclarationRegistrationOptions/*TOpt*/ *DeclarationRegistrationOptions
}

func (it *BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.DeclarationOptions) {
		return json.Marshal(it.DeclarationOptions)
	}
	if !isNone(it.DeclarationRegistrationOptions) {
		return json.Marshal(it.DeclarationRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *DeclarationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DeclarationOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *DeclarationRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DeclarationRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrReferenceOptions struct {
	Boolean/*TOpt*/ *Boolean
	ReferenceOptions/*TOpt*/ *ReferenceOptions
}

func (it *BooleanOrReferenceOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.ReferenceOptions) {
		return json.Marshal(it.ReferenceOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrReferenceOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *ReferenceOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.ReferenceOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	DocumentColorOptions/*TOpt*/ *DocumentColorOptions
	DocumentColorRegistrationOptions/*TOpt*/ *DocumentColorRegistrationOptions
}

func (it *BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.DocumentColorOptions) {
		return json.Marshal(it.DocumentColorOptions)
	}
	if !isNone(it.DocumentColorRegistrationOptions) {
		return json.Marshal(it.DocumentColorRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *DocumentColorOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DocumentColorOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *DocumentColorRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DocumentColorRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type StringOrBoolean struct {
	String/*TOpt*/ *String
	Boolean/*TOpt*/ *Boolean
}

func (it *StringOrBoolean) MarshalJSON() ([]byte, error) {
	if !isNone(it.String) {
		return json.Marshal(it.String)
	}
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}

	return []byte("null"), nil
}
func (it *StringOrBoolean) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *String
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.String = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type SymbolInformationsOrDocumentSymbolsOrNull struct {
	SymbolInformations/*TOpt*/ []SymbolInformation
	DocumentSymbols/*TOpt*/ []DocumentSymbol
}

func (it *SymbolInformationsOrDocumentSymbolsOrNull) MarshalJSON() ([]byte, error) {
	if !isNone(it.SymbolInformations) {
		return json.Marshal(it.SymbolInformations)
	}
	if !isNone(it.DocumentSymbols) {
		return json.Marshal(it.DocumentSymbols)
	}

	return []byte("null"), nil
}
func (it *SymbolInformationsOrDocumentSymbolsOrNull) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ []SymbolInformation
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.SymbolInformations = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ []DocumentSymbol
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DocumentSymbols = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	TypeDefinitionOptions/*TOpt*/ *TypeDefinitionOptions
	TypeDefinitionRegistrationOptions/*TOpt*/ *TypeDefinitionRegistrationOptions
}

func (it *BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.TypeDefinitionOptions) {
		return json.Marshal(it.TypeDefinitionOptions)
	}
	if !isNone(it.TypeDefinitionRegistrationOptions) {
		return json.Marshal(it.TypeDefinitionRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *TypeDefinitionOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.TypeDefinitionOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *TypeDefinitionRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.TypeDefinitionRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrDocumentSymbolOptions struct {
	Boolean/*TOpt*/ *Boolean
	DocumentSymbolOptions/*TOpt*/ *DocumentSymbolOptions
}

func (it *BooleanOrDocumentSymbolOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.DocumentSymbolOptions) {
		return json.Marshal(it.DocumentSymbolOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrDocumentSymbolOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *DocumentSymbolOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.DocumentSymbolOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	SelectionRangeOptions/*TOpt*/ *SelectionRangeOptions
	SelectionRangeRegistrationOptions/*TOpt*/ *SelectionRangeRegistrationOptions
}

func (it *BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.SelectionRangeOptions) {
		return json.Marshal(it.SelectionRangeOptions)
	}
	if !isNone(it.SelectionRangeRegistrationOptions) {
		return json.Marshal(it.SelectionRangeRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *SelectionRangeOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.SelectionRangeOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *SelectionRangeRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.SelectionRangeRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	LinkedEditingRangeOptions/*TOpt*/ *LinkedEditingRangeOptions
	LinkedEditingRangeRegistrationOptions/*TOpt*/ *LinkedEditingRangeRegistrationOptions
}

func (it *BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.LinkedEditingRangeOptions) {
		return json.Marshal(it.LinkedEditingRangeOptions)
	}
	if !isNone(it.LinkedEditingRangeRegistrationOptions) {
		return json.Marshal(it.LinkedEditingRangeRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *LinkedEditingRangeOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.LinkedEditingRangeOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *LinkedEditingRangeRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.LinkedEditingRangeRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

/*TOr*/
type BooleanOrMonikerOptionsOrMonikerRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	MonikerOptions/*TOpt*/ *MonikerOptions
	MonikerRegistrationOptions/*TOpt*/ *MonikerRegistrationOptions
}

func (it *BooleanOrMonikerOptionsOrMonikerRegistrationOptions) MarshalJSON() ([]byte, error) {
	if !isNone(it.Boolean) {
		return json.Marshal(it.Boolean)
	}
	if !isNone(it.MonikerOptions) {
		return json.Marshal(it.MonikerOptions)
	}
	if !isNone(it.MonikerRegistrationOptions) {
		return json.Marshal(it.MonikerRegistrationOptions)
	}

	return []byte("null"), nil
}
func (it *BooleanOrMonikerOptionsOrMonikerRegistrationOptions) UnmarshalJSON(jsonBytes []byte) error {
	if bytes.Equal(jsonBytes, []byte("null")) {
		return nil
	}
	{
		var tmp /*TOpt*/ *Boolean
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.Boolean = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *MonikerOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.MonikerOptions = tmp
			return nil
		}
	}
	{
		var tmp /*TOpt*/ *MonikerRegistrationOptions
		err := json.Unmarshal(jsonBytes, &tmp)
		if err == nil && !isNone(tmp) {
			it.MonikerRegistrationOptions = tmp
			return nil
		}
	}

	return nil
}

type DocumentURI = URI
type URI = String
type Void = struct{}

func isNone(it any) bool {
	if it == nil {
		return true
	}
	if has_is_none, _ := it.(interface{ IsNone() bool }); has_is_none != nil {
		return has_is_none.IsNone()
	}
	rv := reflect.ValueOf(it)
	if maybe, is := it.(reflect.Value); is {
		if rv = maybe; rv.IsValid() {
			if has_is_none, _ := maybe.Interface().(interface{ IsNone() bool }); has_is_none != nil {
				return has_is_none.IsNone()
			}
		}
	}
	return (!rv.IsValid()) || rv.IsZero() || ((rv.Kind() == reflect.Pointer) && isNone(rv.Elem()))
}

// URL is a convenience `url.Parse(string(it))` call that returns `nil` on error.
// If you do need the `error`, just use the original `url.Parse` directly.
func (it *URI) URL() (ret *url.URL) {
	if it.IsNone() {
		return nil
	}
	ret, _ = url.Parse(string(*it))
	return
}

type String string

// IsNone returns whether `it` is `nil` or the `string` zero value.
func (it *String) IsNone() bool { return it == nil || *it == "" }

// IfNone returns `ifNone` if `it` is `nil` or the `string` zero value; otherwise, `it.Value()` is returned.
func (it *String) IfNone(ifNone string) string {
	if it.IsNone() {
		return ifNone
	}
	return string(*it)
}

// Value returns the `string` zero value if `it` is `nil`; otherwise `*it` is returned.
func (it *String) Value() string { return it.IfNil("") }

// IfNil returns `ifNil` if `it` is `nil`; otherwise, `it.Value()` is returned.
func (it *String) IfNil(ifNil string) string {
	if it == nil {
		return ifNil
	}
	return string(*it)
}

type Integer int

// IsNone returns whether `it` is `nil` or the `int` zero value.
func (it *Integer) IsNone() bool { return it == nil || *it == 0 }

// IfNone returns `ifNone` if `it` is `nil` or the `int` zero value; otherwise, `it.Value()` is returned.
func (it *Integer) IfNone(ifNone int) int {
	if it.IsNone() {
		return ifNone
	}
	return int(*it)
}

// Value returns the `int` zero value if `it` is `nil`; otherwise `*it` is returned.
func (it *Integer) Value() int { return it.IfNil(0) }

// IfNil returns `ifNil` if `it` is `nil`; otherwise, `it.Value()` is returned.
func (it *Integer) IfNil(ifNil int) int {
	if it == nil {
		return ifNil
	}
	return int(*it)
}

type Uinteger uint

// IsNone returns whether `it` is `nil` or the `uint` zero value.
func (it *Uinteger) IsNone() bool { return it == nil || *it == 0 }

// IfNone returns `ifNone` if `it` is `nil` or the `uint` zero value; otherwise, `it.Value()` is returned.
func (it *Uinteger) IfNone(ifNone uint) uint {
	if it.IsNone() {
		return ifNone
	}
	return uint(*it)
}

// Value returns the `uint` zero value if `it` is `nil`; otherwise `*it` is returned.
func (it *Uinteger) Value() uint { return it.IfNil(0) }

// IfNil returns `ifNil` if `it` is `nil`; otherwise, `it.Value()` is returned.
func (it *Uinteger) IfNil(ifNil uint) uint {
	if it == nil {
		return ifNil
	}
	return uint(*it)
}

type Decimal float64

// IsNone returns whether `it` is `nil` or the `float64` zero value.
func (it *Decimal) IsNone() bool {
	return it == nil || *it == 0.0 /*|| *it == math.SmallestNonzeroFloat64*/
}

// IfNone returns `ifNone` if `it` is `nil` or the `float64` zero value; otherwise, `it.Value()` is returned.
func (it *Decimal) IfNone(ifNone float64) float64 {
	if it.IsNone() {
		return ifNone
	}
	return float64(*it)
}

// Value returns the `float64` zero value if `it` is `nil`; otherwise `*it` is returned.
func (it *Decimal) Value() float64 { return it.IfNil(0.0) }

// IfNil returns `ifNil` if `it` is `nil`; otherwise, `it.Value()` is returned.
func (it *Decimal) IfNil(ifNil float64) float64 {
	if it == nil {
		return ifNil
	}
	return float64(*it)
}

type Boolean bool

// IsNone returns whether `it` is `nil` or the `bool` zero value.
func (it *Boolean) IsNone() bool { return it == nil || !(bool)(*it) }

// IfNone returns `ifNone` if `it` is `nil` or the `bool` zero value; otherwise, `it.Value()` is returned.
func (it *Boolean) IfNone(ifNone bool) bool {
	if it.IsNone() {
		return ifNone
	}
	return bool(*it)
}

// Value returns the `bool` zero value if `it` is `nil`; otherwise `*it` is returned.
func (it *Boolean) Value() bool { return it.IfNil(false) }

// IfNil returns `ifNil` if `it` is `nil`; otherwise, `it.Value()` is returned.
func (it *Boolean) IfNil(ifNil bool) bool {
	if it == nil {
		return ifNil
	}
	return bool(*it)
}

type jsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ptr[T any](value T) *T { return &value }

func iIf[T any](chk bool, ifTrue T, ifFalse T) T {
	if chk {
		return ifTrue
	}
	return ifFalse
}

type clientServerBase struct {
	stdioLock sync.Mutex // to sync writes to stdout
	stdout    io.Writer
	waiters   map[any]func(any)

	LogPrefixSendRecvJsons string

	// Initialized is for informational purposes only, to the importer who shall not set or mutate them.
	// Its fields are set automatically at the appropriate initialization lifecycle instant.
	Initialized struct {
		Client *InitializeParams
		Server *InitializeResult
	}
}

func (it *clientServerBase) sendErrMsg(err any) {
	if err == nil {
		return
	}
	json_rpc_err_msg, is_json_rpc_err_msg := err.(*jsonRpcError)
	if json_rpc_err_msg == nil {
		if is_json_rpc_err_msg {
			return
		}
		json_rpc_err_msg = &jsonRpcError{Code: int(ErrorCodesInternalError), Message: fmt.Sprintf("%v", err)}
	}
	it.sendMsg(json_rpc_err_msg)
}

func (it *clientServerBase) sendMsg(jsonable any) {
	json_bytes, _ := json.Marshal(jsonable)
	it.stdioLock.Lock()
	defer it.stdioLock.Unlock()
	if it.LogPrefixSendRecvJsons != "" {
		println(it.LogPrefixSendRecvJsons + ".SEND>>" + string(json_bytes) + ">>")
	}
	_, _ = it.stdout.Write([]byte("Content-Length: "))
	_, _ = it.stdout.Write([]byte(strconv.Itoa(len(json_bytes))))
	_, _ = it.stdout.Write([]byte("\r\n\r\n"))
	_, _ = it.stdout.Write(json_bytes)
}

func (it *clientServerBase) send(methodName string, params any, isReq bool, onResp func(any)) {
	req_id := strconv.FormatInt(time.Now().UnixNano(), 36)
	req := map[string]any{"method": methodName, "params": params}
	if onResp != nil {
		it.waiters[req_id] = onResp
	}
	if isReq {
		req["id"] = req_id
	}
	it.sendMsg(req)
}

// forever keeps reading and handling LSP JSON-RPC messages incoming over
// `in` until reading from `in` fails, then returns that IO read error.
func (it *clientServerBase) forever(in io.Reader, out io.Writer, handleIncoming func(map[string]any) *jsonRpcError) error {
	const buf_cap = 1024 * 1024

	it.stdout = out
	it.waiters = map[any]func(any){}

	stdin := bufio.NewScanner(in)
	stdin.Split(func(data []byte, ateof bool) (advance int, token []byte, err error) {
		if i_cl1 := bytes.Index(data, []byte("Content-Length: ")); i_cl1 >= 0 {
			datafromclen := data[i_cl1+16:]
			if i_cl2 := bytes.IndexAny(datafromclen, "\r\n"); i_cl2 > 0 {
				if clen, e := strconv.Atoi(string(datafromclen[:i_cl2])); e != nil {
					err = e
				} else if i_js1 := bytes.Index(datafromclen, []byte("{\"")); i_js1 > i_cl2 {
					if i_js2 := i_js1 + clen; len(datafromclen) >= i_js2 {
						advance = i_cl1 + 16 + i_js2
						token = datafromclen[i_js1:i_js2]
					}
				}
			}
		}
		return
	})

	for stdin.Scan() {
		raw := map[string]any{}
		json_bytes := stdin.Bytes()
		if it.LogPrefixSendRecvJsons != "" {
			it.stdioLock.Lock()
			println(it.LogPrefixSendRecvJsons + ".RECV<<" + string(json_bytes) + "<<")
			it.stdioLock.Unlock()
		}
		if err := json.Unmarshal(json_bytes, &raw); err != nil {
			it.sendErrMsg(&jsonRpcError{Code: int(ErrorCodesParseError), Message: err.Error()})
			continue
		}

		if raw["code"] != nil { // received an error message
			it.stdioLock.Lock()
			println(string(json_bytes)) // goes to stderr
			it.stdioLock.Unlock()
			continue
		}
		if msg_id := raw["id"]; raw["method"] == nil { // received a Response message
			handler := it.waiters[msg_id]
			delete(it.waiters, msg_id)
			go handler(raw["result"])
		} else {
			it.sendErrMsg(handleIncoming(raw))
		}
	}
	return stdin.Err()
}

func clientServerOnResp[T any](it *clientServerBase, onResp func(T)) func(any) {
	if onResp == nil {
		return nil
	}
	return func(resultAsMap any) {
		var result, none T
		if resultAsMap != nil {
			json_bytes, _ := json.Marshal(resultAsMap)
			if err := json.Unmarshal(json_bytes, &result); err != nil {
				it.sendErrMsg(err)
				return
			}
		}
		onResp(iIf(resultAsMap == nil, none, result))
	}
}

func clientServerHandleIncoming[T any](it *clientServerBase, handler func(*T) (any, error), msgMethodName string, msgId any, msgParams any) {
	if handler == nil {
		if msgId != nil {
			it.sendErrMsg(errors.New("unimplemented: " + msgMethodName))
		}
		return
	}
	var params T
	if msgParams != nil {
		json_bytes, _ := json.Marshal(msgParams)
		if err := json.Unmarshal(json_bytes, &params); err != nil {
			it.sendErrMsg(&jsonRpcError{Code: int(ErrorCodesInvalidParams), Message: err.Error()})
			return
		}
	}
	go func(params *T) {
		if msgParams == nil {
			params = nil
		}
		result, err := handler(params)
		resp := map[string]any{
			"jsonrpc": "2.0",
			"result":  result,
			"id":      msgId,
		}
		if err != nil {
			if msgId != nil {
				resp["error"] = &jsonRpcError{Code: int(ErrorCodesInternalError), Message: fmt.Sprintf("%v", err)}
			} else {
				it.sendErrMsg(err)
				return
			}
		}
		if msgId != nil {
			it.sendMsg(resp)
		}
	}(&params)
}
