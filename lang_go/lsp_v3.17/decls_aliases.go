// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-lsp/gen/cmd/gen_lsp
package lsp

// The definition of a symbol represented as one or many `Location`.
// For most programming languages there is only one location at which a symbol is
// defined.
//
// Servers should prefer returning `DefinitionLink` over `Definition` if supported
// by the client.
//
// "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
type Definition /*TOr*/ struct {
	Location/*TOpt*/ *Location
	Locations/*TOpt*/ []Location
}

// Information about where a symbol is defined.
//
// Provides additional metadata over normal `Location` definitions, including the range of
// the defining symbol
type DefinitionLink = LocationLink

// LSP arrays.
// @since 3.17.0
type LSPArray = []LSPAny

// The LSP any type.
// Please note that strictly speaking a property with the value `undefined`
// can't be converted into JSON preserving the property name. However for
// convenience it is allowed and assumed that all these properties are
// optional as well.
// @since 3.17.0
//
// "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
type LSPAny = any

// The declaration of a symbol representation as one or many `Location`.
//
// "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
type Declaration /*TOr*/ struct {
	Location/*TOpt*/ *Location
	Locations/*TOpt*/ []Location
}

// Information about where a symbol is declared.
//
// Provides additional metadata over normal `Location` declarations, including the range of
// the declaring symbol.
//
// Servers should prefer returning `DeclarationLink` over `Declaration` if supported
// by the client.
type DeclarationLink = LocationLink

// Inline value information can be provided by different means:
// - directly as a text value (class InlineValueText).
// - as a name to use for a variable lookup (class InlineValueVariableLookup)
// - as an evaluatable expression (class InlineValueEvaluatableExpression)
// The InlineValue types combines all inline value types into one type.
//
// @since 3.17.0
//
// "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
type InlineValue /*TOr*/ struct {
	InlineValueText/*TOpt*/ *InlineValueText
	InlineValueVariableLookup/*TOpt*/ *InlineValueVariableLookup
	InlineValueEvaluatableExpression/*TOpt*/ *InlineValueEvaluatableExpression
}

// The result of a document diagnostic pull request. A report can
// either be a full report containing all diagnostics for the
// requested document or an unchanged report indicating that nothing
// has changed in terms of diagnostics in comparison to the last
// pull request.
//
// @since 3.17.0
//
// "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
type DocumentDiagnosticReport /*TOr*/ struct {
	RelatedFullDocumentDiagnosticReport/*TOpt*/ *RelatedFullDocumentDiagnosticReport
	RelatedUnchangedDocumentDiagnosticReport/*TOpt*/ *RelatedUnchangedDocumentDiagnosticReport
}

// "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
type PrepareRenameResult /*TOr*/ struct {
	Range/*TOpt*/ *Range
	RangeWithPlaceholderString/*TOpt*/ * /*TStruc*/ struct {
		Range Range `json:"range"`

		Placeholder string `json:"placeholder"`
	}
	DefaultBehaviorBoolean/*TOpt*/ * /*TStruc*/ struct {
		DefaultBehavior bool `json:"defaultBehavior"`
	}
}

// A document selector is the combination of one or many document filters.
//
// @sample `let sel:DocumentSelector = [{ language: 'typescript' }, { language: 'json', pattern: '**∕tsconfig.json' }]`;
//
// The use of a string as a document filter is deprecated @since 3.16.0.
type DocumentSelector = []DocumentFilter

// "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
type ProgressToken /*TOr*/ struct {
	Integer/*TOpt*/ *Integer
	String/*TOpt*/ *String
}

// An identifier to refer to a change annotation stored with a workspace edit.
type ChangeAnnotationIdentifier = String

// A workspace diagnostic document report.
//
// @since 3.17.0
//
// "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
type WorkspaceDocumentDiagnosticReport /*TOr*/ struct {
	WorkspaceFullDocumentDiagnosticReport/*TOpt*/ *WorkspaceFullDocumentDiagnosticReport
	WorkspaceUnchangedDocumentDiagnosticReport/*TOpt*/ *WorkspaceUnchangedDocumentDiagnosticReport
}

// An event describing a change to a text document. If only a text is provided
// it is considered to be the full content of the document.
//
// "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
type TextDocumentContentChangeEvent /*TOr*/ struct {
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

// MarkedString can be used to render human readable text. It is either a markdown string
// or a code-block that provides a language and a code snippet. The language identifier
// is semantically equal to the optional language identifier in fenced code blocks in GitHub
// issues. See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting
//
// The pair of a language and a value is an equivalent to markdown:
// ```${language}
// ${value}
// ```
//
// Note that markdown strings will be sanitized - that means html will be escaped.
// @deprecated use MarkupContent instead.
//
// "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
type MarkedString /*TOr*/ struct {
	String/*TOpt*/ *String
	LanguageStringWithValueString/*TOpt*/ * /*TStruc*/ struct {
		Language string `json:"language"`

		Value string `json:"value"`
	}
}

// A document filter describes a top level text document or
// a notebook cell document.
//
// @since 3.17.0 - proposed support for NotebookCellTextDocumentFilter.
//
// "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
type DocumentFilter /*TOr*/ struct {
	TextDocumentFilter/*TOpt*/ *TextDocumentFilter
	NotebookCellTextDocumentFilter/*TOpt*/ *NotebookCellTextDocumentFilter
}

// LSP object definition.
// @since 3.17.0
type LSPObject = map[string]LSPAny

// The glob pattern. Either a string pattern or a relative pattern.
//
// @since 3.17.0
//
// "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
type GlobPattern /*TOr*/ struct {
	Pattern/*TOpt*/ *String
	RelativePattern/*TOpt*/ *RelativePattern
}

// A document filter denotes a document by different properties like
// the `TextDocument.languageId`, the `Uri.scheme` of
// its resource, or a glob-pattern that is applied to the `TextDocument.fileName`.
//
// Glob patterns can have the following syntax:
// - `*` to match one or more characters in a path segment
// - `?` to match on one character in a path segment
// - `**` to match any number of path segments, including none
// - `{}` to group sub patterns into an OR expression. (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
// - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
// - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
//
// @sample A language filter that applies to typescript files on disk: `{ language: 'typescript', scheme: 'file' }`
// @sample A language filter that applies to all package.json paths: `{ language: 'json', pattern: '**package.json' }`
//
// @since 3.17.0
type TextDocumentFilter /*TStruc*/ struct {

	// A language id, like `typescript`.
	Language/*TOpt*/ *String `json:"language"`

	// A Uri `Uri.scheme`, like `file` or `untitled`.
	Scheme/*TOpt*/ *String `json:"scheme"`

	// A glob pattern, like **​/*.{ts,js}. See TextDocumentFilter for examples.
	Pattern/*TOpt*/ *String `json:"pattern"`
}

// A notebook document filter denotes a notebook document by
// different properties. The properties will be match
// against the notebook's URI (same as with documents)
//
// @since 3.17.0
type NotebookDocumentFilter /*TStruc*/ struct {

	// The type of the enclosing notebook.
	NotebookType/*TOpt*/ *String `json:"notebookType"`

	// A Uri `Uri.scheme`, like `file` or `untitled`.
	Scheme/*TOpt*/ *String `json:"scheme"`

	// A glob pattern.
	Pattern/*TOpt*/ *String `json:"pattern"`
}

// The glob pattern to watch relative to the base path. Glob patterns can have the following syntax:
// - `*` to match one or more characters in a path segment
// - `?` to match on one character in a path segment
// - `**` to match any number of path segments, including none
// - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
// - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
// - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
//
// @since 3.17.0
type Pattern = String
