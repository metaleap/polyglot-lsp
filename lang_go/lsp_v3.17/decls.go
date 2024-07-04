// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-lsp/gen/cmd/gen_lsp
package lsp

import (
	"net/url"
)

/*TOr*/
type RangeOrRangeWithPlaceholderStringOrDefaultBehaviorBoolean struct {
	Range/*TOpt*/ *Range
	RangeWithPlaceholderString/*TOpt*/ * /*TStruc*/ struct {
		Range Range `json:"range"`

		Placeholder string `json:"placeholder"`
	}
	DefaultBehaviorBoolean/*TOpt*/ * /*TStruc*/ struct {
		DefaultBehavior bool `json:"defaultBehavior"`
	}
}

/*TOr*/
type StringOrLanguageStringWithValueString struct {
	String/*TOpt*/ *String
	LanguageStringWithValueString/*TOpt*/ * /*TStruc*/ struct {
		Language string `json:"language"`

		Value string `json:"value"`
	}
}

/*TOr*/
type StringOrStrings struct {
	String/*TOpt*/ *String
	Strings/*TOpt*/ []string
}

/*TOr*/
type BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	InlineValueOptions/*TOpt*/ *InlineValueOptions
	InlineValueRegistrationOptions/*TOpt*/ *InlineValueRegistrationOptions
}

/*TOr*/
type BooleanOrDocumentFormattingOptions struct {
	Boolean/*TOpt*/ *Boolean
	DocumentFormattingOptions/*TOpt*/ *DocumentFormattingOptions
}

/*TOr*/
type InlineValueTextOrInlineValueVariableLookupOrInlineValueEvaluatableExpression struct {
	InlineValueText/*TOpt*/ *InlineValueText
	InlineValueVariableLookup/*TOpt*/ *InlineValueVariableLookup
	InlineValueEvaluatableExpression/*TOpt*/ *InlineValueEvaluatableExpression
}

/*TOr*/
type WorkspaceFullDocumentDiagnosticReportOrWorkspaceUnchangedDocumentDiagnosticReport struct {
	WorkspaceFullDocumentDiagnosticReport/*TOpt*/ *WorkspaceFullDocumentDiagnosticReport
	WorkspaceUnchangedDocumentDiagnosticReport/*TOpt*/ *WorkspaceUnchangedDocumentDiagnosticReport
}

/*TOr*/
type TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile struct {
	TextDocumentEdit/*TOpt*/ *TextDocumentEdit
	CreateFile/*TOpt*/ *CreateFile
	RenameFile/*TOpt*/ *RenameFile
	DeleteFile/*TOpt*/ *DeleteFile
}

/*TOr*/
type MarkupContentOrMarkedStringOrMarkedStrings struct {
	MarkupContent/*TOpt*/ *MarkupContent
	MarkedString/*TOpt*/ MarkedString
	MarkedStrings/*TOpt*/ []MarkedString
}

/*TOr*/
type BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	TypeDefinitionOptions/*TOpt*/ *TypeDefinitionOptions
	TypeDefinitionRegistrationOptions/*TOpt*/ *TypeDefinitionRegistrationOptions
}

/*TOr*/
type BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	DocumentColorOptions/*TOpt*/ *DocumentColorOptions
	DocumentColorRegistrationOptions/*TOpt*/ *DocumentColorRegistrationOptions
}

/*TOr*/
type BooleanOrWorkspaceSymbolOptions struct {
	Boolean/*TOpt*/ *Boolean
	WorkspaceSymbolOptions/*TOpt*/ *WorkspaceSymbolOptions
}

/*TOr*/
type BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	SelectionRangeOptions/*TOpt*/ *SelectionRangeOptions
	SelectionRangeRegistrationOptions/*TOpt*/ *SelectionRangeRegistrationOptions
}

/*TOr*/
type StringOrMarkupContent struct {
	String/*TOpt*/ *String
	MarkupContent/*TOpt*/ *MarkupContent
}

/*TOr*/
type SemanticTokensOptionsOrSemanticTokensRegistrationOptions struct {
	SemanticTokensOptions/*TOpt*/ *SemanticTokensOptions
	SemanticTokensRegistrationOptions/*TOpt*/ *SemanticTokensRegistrationOptions
}

/*TOr*/
type BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	InlayHintOptions/*TOpt*/ *InlayHintOptions
	InlayHintRegistrationOptions/*TOpt*/ *InlayHintRegistrationOptions
}

/*TOr*/
type StringOrNotebookDocumentFilter struct {
	String/*TOpt*/ *String
	NotebookDocumentFilter/*TOpt*/ *NotebookDocumentFilter
}

/*TOr*/
type FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport struct {
	FullDocumentDiagnosticReport/*TOpt*/ *FullDocumentDiagnosticReport
	UnchangedDocumentDiagnosticReport/*TOpt*/ *UnchangedDocumentDiagnosticReport
}

/*TOr*/
type TextEditOrInsertReplaceEdit struct {
	TextEdit/*TOpt*/ *TextEdit
	InsertReplaceEdit/*TOpt*/ *InsertReplaceEdit
}

/*TOr*/
type BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	DeclarationOptions/*TOpt*/ *DeclarationOptions
	DeclarationRegistrationOptions/*TOpt*/ *DeclarationRegistrationOptions
}

/*TOr*/
type BooleanOrDefinitionOptions struct {
	Boolean/*TOpt*/ *Boolean
	DefinitionOptions/*TOpt*/ *DefinitionOptions
}

/*TOr*/
type BooleanOrSaveOptions struct {
	Boolean/*TOpt*/ *Boolean
	SaveOptions/*TOpt*/ *SaveOptions
}

/*TOr*/
type StringOrBoolean struct {
	String/*TOpt*/ *String
	Boolean/*TOpt*/ *Boolean
}

/*TOr*/
type StringOrInlayHintLabelParts struct {
	String/*TOpt*/ *String
	InlayHintLabelParts/*TOpt*/ []InlayHintLabelPart
}

/*TOr*/
type BooleanOrDocumentSymbolOptions struct {
	Boolean/*TOpt*/ *Boolean
	DocumentSymbolOptions/*TOpt*/ *DocumentSymbolOptions
}

/*TOr*/
type BooleanOrAnyByString struct {
	Boolean/*TOpt*/ *Boolean
	AnyByString/*TOpt*/ map[string]any
}

/*TOr*/
type WorkspaceFolderOrURI struct {
	WorkspaceFolder/*TOpt*/ *WorkspaceFolder
	URI/*TOpt*/ *URI
}

/*TOr*/
type LocationOrLocations struct {
	Location/*TOpt*/ *Location
	Locations/*TOpt*/ []Location
}

/*TOr*/
type RangeWithRangeLengthUintegerWithTextStringOrTextString struct {
	RangeWithRangeLengthUintegerWithTextString/*TOpt*/ * /*TStruc*/ struct {

		// The range of the document that changed.
		Range Range `json:"range"`

		// The optional length of the range that got replaced.
		//
		// @deprecated use range instead.
		RangeLength/*TOpt*/ *Uinteger `json:"rangeLength"`

		// The new text for the provided range.
		Text string `json:"text"`
	}
	TextString/*TOpt*/ * /*TStruc*/ struct {

		// The new text of the whole document.
		Text string `json:"text"`
	}
}

/*TOr*/
type BooleanOrDocumentHighlightOptions struct {
	Boolean/*TOpt*/ *Boolean
	DocumentHighlightOptions/*TOpt*/ *DocumentHighlightOptions
}

/*TOr*/
type TextEditOrAnnotatedTextEdit struct {
	TextEdit/*TOpt*/ *TextEdit
	AnnotatedTextEdit/*TOpt*/ *AnnotatedTextEdit
}

/*TOr*/
type BooleanOrCodeActionOptions struct {
	Boolean/*TOpt*/ *Boolean
	CodeActionOptions/*TOpt*/ *CodeActionOptions
}

/*TOr*/
type RelatedFullDocumentDiagnosticReportOrRelatedUnchangedDocumentDiagnosticReport struct {
	RelatedFullDocumentDiagnosticReport/*TOpt*/ *RelatedFullDocumentDiagnosticReport
	RelatedUnchangedDocumentDiagnosticReport/*TOpt*/ *RelatedUnchangedDocumentDiagnosticReport
}

/*TOr*/
type TextDocumentFilterOrNotebookCellTextDocumentFilter struct {
	TextDocumentFilter/*TOpt*/ *TextDocumentFilter
	NotebookCellTextDocumentFilter/*TOpt*/ *NotebookCellTextDocumentFilter
}

/*TOr*/
type BooleanOrImplementationOptionsOrImplementationRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	ImplementationOptions/*TOpt*/ *ImplementationOptions
	ImplementationRegistrationOptions/*TOpt*/ *ImplementationRegistrationOptions
}

/*TOr*/
type BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	CallHierarchyOptions/*TOpt*/ *CallHierarchyOptions
	CallHierarchyRegistrationOptions/*TOpt*/ *CallHierarchyRegistrationOptions
}

/*TOr*/
type DiagnosticOptionsOrDiagnosticRegistrationOptions struct {
	DiagnosticOptions/*TOpt*/ *DiagnosticOptions
	DiagnosticRegistrationOptions/*TOpt*/ *DiagnosticRegistrationOptions
}

/*TOr*/
type BooleanOrRenameOptions struct {
	Boolean/*TOpt*/ *Boolean
	RenameOptions/*TOpt*/ *RenameOptions
}

/*TOr*/
type BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	TypeHierarchyOptions/*TOpt*/ *TypeHierarchyOptions
	TypeHierarchyRegistrationOptions/*TOpt*/ *TypeHierarchyRegistrationOptions
}

/*TOr*/
type IntegerOrString struct {
	Integer/*TOpt*/ *Integer
	String/*TOpt*/ *String
}

/*TOr*/
type NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions struct {
	NotebookDocumentSyncOptions/*TOpt*/ *NotebookDocumentSyncOptions
	NotebookDocumentSyncRegistrationOptions/*TOpt*/ *NotebookDocumentSyncRegistrationOptions
}

/*TOr*/
type BooleanOrDocumentRangeFormattingOptions struct {
	Boolean/*TOpt*/ *Boolean
	DocumentRangeFormattingOptions/*TOpt*/ *DocumentRangeFormattingOptions
}

/*TOr*/
type StringOrUintegerWithUinteger struct {
	String/*TOpt*/ *String
	UintegerWithUinteger/*TOpt*/ * /*TTup*/ []uint
}

/*TOr*/
type LocationOrUriDocumentUri struct {
	Location/*TOpt*/ *Location
	UriDocumentUri/*TOpt*/ * /*TStruc*/ struct {
		Uri DocumentURI `json:"uri"`
	}
}

/*TOr*/
type BooleanOrDeltaBoolean struct {
	Boolean/*TOpt*/ *Boolean
	DeltaBoolean/*TOpt*/ * /*TStruc*/ struct {

		// The server supports deltas for full documents.
		Delta /*TOpt*/ *Boolean `json:"delta"`
	}
}

/*TOr*/
type BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	FoldingRangeOptions/*TOpt*/ *FoldingRangeOptions
	FoldingRangeRegistrationOptions/*TOpt*/ *FoldingRangeRegistrationOptions
}

/*TOr*/
type BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	LinkedEditingRangeOptions/*TOpt*/ *LinkedEditingRangeOptions
	LinkedEditingRangeRegistrationOptions/*TOpt*/ *LinkedEditingRangeRegistrationOptions
}

/*TOr*/
type BooleanOrMonikerOptionsOrMonikerRegistrationOptions struct {
	Boolean/*TOpt*/ *Boolean
	MonikerOptions/*TOpt*/ *MonikerOptions
	MonikerRegistrationOptions/*TOpt*/ *MonikerRegistrationOptions
}

/*TOr*/
type RangeOrInsertRangeWithReplaceRange struct {
	Range/*TOpt*/ *Range
	InsertRangeWithReplaceRange/*TOpt*/ * /*TStruc*/ struct {
		Insert Range `json:"insert"`

		Replace Range `json:"replace"`
	}
}

/*TOr*/
type TextDocumentSyncOptionsOrTextDocumentSyncKind struct {
	TextDocumentSyncOptions/*TOpt*/ *TextDocumentSyncOptions
	TextDocumentSyncKind/*TOpt*/ TextDocumentSyncKind
}

/*TOr*/
type BooleanOrHoverOptions struct {
	Boolean/*TOpt*/ *Boolean
	HoverOptions/*TOpt*/ *HoverOptions
}

/*TOr*/
type BooleanOrDeltaBoolean_ struct {
	Boolean/*TOpt*/ *Boolean
	DeltaBoolean/*TOpt*/ * /*TStruc*/ struct {

		// The client will send the `textDocument/semanticTokens/full/delta` request if
		// the server provides a corresponding handler.
		Delta /*TOpt*/ *Boolean `json:"delta"`
	}
}

/*TOr*/
type PatternOrRelativePattern struct {
	Pattern/*TOpt*/ *String
	RelativePattern/*TOpt*/ *RelativePattern
}

/*TOr*/
type BooleanOrReferenceOptions struct {
	Boolean/*TOpt*/ *Boolean
	ReferenceOptions/*TOpt*/ *ReferenceOptions
}

type DocumentURI = URI
type URI = String
type Void = struct{}

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
