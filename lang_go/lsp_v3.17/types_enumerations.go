// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-vsx-and-lsp/gen/cmd/gen_lsp
package lsp

// A set of predefined token types. This set is not fixed
// an clients can specify additional token types via the
// corresponding client capabilities.
//
// @since 3.16.0
type SemanticTokenTypes String

const SemanticTokenTypesNamespace SemanticTokenTypes = "namespace"

// Represents a generic type. Acts as a fallback for types which can't be mapped to
// a specific type like class or enum.
const SemanticTokenTypesType SemanticTokenTypes = "type"

const SemanticTokenTypesClass SemanticTokenTypes = "class"

const SemanticTokenTypesEnum SemanticTokenTypes = "enum"

const SemanticTokenTypesInterface SemanticTokenTypes = "interface"

const SemanticTokenTypesStruct SemanticTokenTypes = "struct"

const SemanticTokenTypesTypeParameter SemanticTokenTypes = "typeParameter"

const SemanticTokenTypesParameter SemanticTokenTypes = "parameter"

const SemanticTokenTypesVariable SemanticTokenTypes = "variable"

const SemanticTokenTypesProperty SemanticTokenTypes = "property"

const SemanticTokenTypesEnumMember SemanticTokenTypes = "enumMember"

const SemanticTokenTypesEvent SemanticTokenTypes = "event"

const SemanticTokenTypesFunction SemanticTokenTypes = "function"

const SemanticTokenTypesMethod SemanticTokenTypes = "method"

const SemanticTokenTypesMacro SemanticTokenTypes = "macro"

const SemanticTokenTypesKeyword SemanticTokenTypes = "keyword"

const SemanticTokenTypesModifier SemanticTokenTypes = "modifier"

const SemanticTokenTypesComment SemanticTokenTypes = "comment"

const SemanticTokenTypesString SemanticTokenTypes = "string"

const SemanticTokenTypesNumber SemanticTokenTypes = "number"

const SemanticTokenTypesRegexp SemanticTokenTypes = "regexp"

const SemanticTokenTypesOperator SemanticTokenTypes = "operator"

// @since 3.17.0
const SemanticTokenTypesDecorator SemanticTokenTypes = "decorator"

// String returns "" or "Namespace" or "Type" or "Class" or "Enum" or "Interface" or "Struct" or "TypeParameter" or "Parameter" or "Variable" or "Property" or "EnumMember" or "Event" or "Function" or "Method" or "Macro" or "Keyword" or "Modifier" or "Comment" or "String" or "Number" or "Regexp" or "Operator" or "Decorator", depending on the value of `it`.
func (it SemanticTokenTypes) String() string {
	switch it {
	case SemanticTokenTypesNamespace:
		return "Namespace"
	case SemanticTokenTypesType:
		return "Type"
	case SemanticTokenTypesClass:
		return "Class"
	case SemanticTokenTypesEnum:
		return "Enum"
	case SemanticTokenTypesInterface:
		return "Interface"
	case SemanticTokenTypesStruct:
		return "Struct"
	case SemanticTokenTypesTypeParameter:
		return "TypeParameter"
	case SemanticTokenTypesParameter:
		return "Parameter"
	case SemanticTokenTypesVariable:
		return "Variable"
	case SemanticTokenTypesProperty:
		return "Property"
	case SemanticTokenTypesEnumMember:
		return "EnumMember"
	case SemanticTokenTypesEvent:
		return "Event"
	case SemanticTokenTypesFunction:
		return "Function"
	case SemanticTokenTypesMethod:
		return "Method"
	case SemanticTokenTypesMacro:
		return "Macro"
	case SemanticTokenTypesKeyword:
		return "Keyword"
	case SemanticTokenTypesModifier:
		return "Modifier"
	case SemanticTokenTypesComment:
		return "Comment"
	case SemanticTokenTypesString:
		return "String"
	case SemanticTokenTypesNumber:
		return "Number"
	case SemanticTokenTypesRegexp:
		return "Regexp"
	case SemanticTokenTypesOperator:
		return "Operator"
	case SemanticTokenTypesDecorator:
		return "Decorator"
	}
	return ""
}

// A set of predefined token modifiers. This set is not fixed
// an clients can specify additional token types via the
// corresponding client capabilities.
//
// @since 3.16.0
type SemanticTokenModifiers String

const SemanticTokenModifiersDeclaration SemanticTokenModifiers = "declaration"

const SemanticTokenModifiersDefinition SemanticTokenModifiers = "definition"

const SemanticTokenModifiersReadonly SemanticTokenModifiers = "readonly"

const SemanticTokenModifiersStatic SemanticTokenModifiers = "static"

const SemanticTokenModifiersDeprecated SemanticTokenModifiers = "deprecated"

const SemanticTokenModifiersAbstract SemanticTokenModifiers = "abstract"

const SemanticTokenModifiersAsync SemanticTokenModifiers = "async"

const SemanticTokenModifiersModification SemanticTokenModifiers = "modification"

const SemanticTokenModifiersDocumentation SemanticTokenModifiers = "documentation"

const SemanticTokenModifiersDefaultLibrary SemanticTokenModifiers = "defaultLibrary"

// String returns "" or "Declaration" or "Definition" or "Readonly" or "Static" or "Deprecated" or "Abstract" or "Async" or "Modification" or "Documentation" or "DefaultLibrary", depending on the value of `it`.
func (it SemanticTokenModifiers) String() string {
	switch it {
	case SemanticTokenModifiersDeclaration:
		return "Declaration"
	case SemanticTokenModifiersDefinition:
		return "Definition"
	case SemanticTokenModifiersReadonly:
		return "Readonly"
	case SemanticTokenModifiersStatic:
		return "Static"
	case SemanticTokenModifiersDeprecated:
		return "Deprecated"
	case SemanticTokenModifiersAbstract:
		return "Abstract"
	case SemanticTokenModifiersAsync:
		return "Async"
	case SemanticTokenModifiersModification:
		return "Modification"
	case SemanticTokenModifiersDocumentation:
		return "Documentation"
	case SemanticTokenModifiersDefaultLibrary:
		return "DefaultLibrary"
	}
	return ""
}

// The document diagnostic report kinds.
//
// @since 3.17.0
type DocumentDiagnosticReportKind String

// A diagnostic report with a full
// set of problems.
const DocumentDiagnosticReportKindFull DocumentDiagnosticReportKind = "full"

// A report indicating that the last
// returned report is still accurate.
const DocumentDiagnosticReportKindUnchanged DocumentDiagnosticReportKind = "unchanged"

// String returns "" or "Full" or "Unchanged", depending on the value of `it`.
func (it DocumentDiagnosticReportKind) String() string {
	switch it {
	case DocumentDiagnosticReportKindFull:
		return "Full"
	case DocumentDiagnosticReportKindUnchanged:
		return "Unchanged"
	}
	return ""
}

// Predefined error codes.
type ErrorCodes Integer

const ErrorCodesParseError ErrorCodes = -32700

const ErrorCodesInvalidRequest ErrorCodes = -32600

const ErrorCodesMethodNotFound ErrorCodes = -32601

const ErrorCodesInvalidParams ErrorCodes = -32602

const ErrorCodesInternalError ErrorCodes = -32603

// Error code indicating that a server received a notification or
// request before the server has received the `initialize` request.
const ErrorCodesServerNotInitialized ErrorCodes = -32002

const ErrorCodesUnknownErrorCode ErrorCodes = -32001

// String returns "" or "ParseError" or "InvalidRequest" or "MethodNotFound" or "InvalidParams" or "InternalError" or "ServerNotInitialized" or "UnknownErrorCode", depending on the value of `it`.
func (it ErrorCodes) String() string {
	switch it {
	case ErrorCodesParseError:
		return "ParseError"
	case ErrorCodesInvalidRequest:
		return "InvalidRequest"
	case ErrorCodesMethodNotFound:
		return "MethodNotFound"
	case ErrorCodesInvalidParams:
		return "InvalidParams"
	case ErrorCodesInternalError:
		return "InternalError"
	case ErrorCodesServerNotInitialized:
		return "ServerNotInitialized"
	case ErrorCodesUnknownErrorCode:
		return "UnknownErrorCode"
	}
	return ""
}

type LSPErrorCodes Integer

// A request failed but it was syntactically correct, e.g the
// method name was known and the parameters were valid. The error
// message should contain human readable information about why
// the request failed.
//
// @since 3.17.0
const LSPErrorCodesRequestFailed LSPErrorCodes = -32803

// The server cancelled the request. This error code should
// only be used for requests that explicitly support being
// server cancellable.
//
// @since 3.17.0
const LSPErrorCodesServerCancelled LSPErrorCodes = -32802

// The server detected that the content of a document got
// modified outside normal conditions. A server should
// NOT send this error code if it detects a content change
// in it unprocessed messages. The result even computed
// on an older state might still be useful for the client.
//
// If a client decides that a result is not of any use anymore
// the client should cancel the request.
const LSPErrorCodesContentModified LSPErrorCodes = -32801

// The client has canceled a request and a server as detected
// the cancel.
const LSPErrorCodesRequestCancelled LSPErrorCodes = -32800

// String returns "" or "RequestFailed" or "ServerCancelled" or "ContentModified" or "RequestCancelled", depending on the value of `it`.
func (it LSPErrorCodes) String() string {
	switch it {
	case LSPErrorCodesRequestFailed:
		return "RequestFailed"
	case LSPErrorCodesServerCancelled:
		return "ServerCancelled"
	case LSPErrorCodesContentModified:
		return "ContentModified"
	case LSPErrorCodesRequestCancelled:
		return "RequestCancelled"
	}
	return ""
}

// A set of predefined range kinds.
type FoldingRangeKind String

// Folding range for a comment
const FoldingRangeKindComment FoldingRangeKind = "comment"

// Folding range for an import or include
const FoldingRangeKindImports FoldingRangeKind = "imports"

// Folding range for a region (e.g. `#region`)
const FoldingRangeKindRegion FoldingRangeKind = "region"

// String returns "" or "Comment" or "Imports" or "Region", depending on the value of `it`.
func (it FoldingRangeKind) String() string {
	switch it {
	case FoldingRangeKindComment:
		return "Comment"
	case FoldingRangeKindImports:
		return "Imports"
	case FoldingRangeKindRegion:
		return "Region"
	}
	return ""
}

// A symbol kind.
type SymbolKind Uinteger

const SymbolKindFile SymbolKind = 1

const SymbolKindModule SymbolKind = 2

const SymbolKindNamespace SymbolKind = 3

const SymbolKindPackage SymbolKind = 4

const SymbolKindClass SymbolKind = 5

const SymbolKindMethod SymbolKind = 6

const SymbolKindProperty SymbolKind = 7

const SymbolKindField SymbolKind = 8

const SymbolKindConstructor SymbolKind = 9

const SymbolKindEnum SymbolKind = 10

const SymbolKindInterface SymbolKind = 11

const SymbolKindFunction SymbolKind = 12

const SymbolKindVariable SymbolKind = 13

const SymbolKindConstant SymbolKind = 14

const SymbolKindString SymbolKind = 15

const SymbolKindNumber SymbolKind = 16

const SymbolKindBoolean SymbolKind = 17

const SymbolKindArray SymbolKind = 18

const SymbolKindObject SymbolKind = 19

const SymbolKindKey SymbolKind = 20

const SymbolKindNull SymbolKind = 21

const SymbolKindEnumMember SymbolKind = 22

const SymbolKindStruct SymbolKind = 23

const SymbolKindEvent SymbolKind = 24

const SymbolKindOperator SymbolKind = 25

const SymbolKindTypeParameter SymbolKind = 26

// String returns "" or "File" or "Module" or "Namespace" or "Package" or "Class" or "Method" or "Property" or "Field" or "Constructor" or "Enum" or "Interface" or "Function" or "Variable" or "Constant" or "String" or "Number" or "Boolean" or "Array" or "Object" or "Key" or "Null" or "EnumMember" or "Struct" or "Event" or "Operator" or "TypeParameter", depending on the value of `it`.
func (it SymbolKind) String() string {
	switch it {
	case SymbolKindFile:
		return "File"
	case SymbolKindModule:
		return "Module"
	case SymbolKindNamespace:
		return "Namespace"
	case SymbolKindPackage:
		return "Package"
	case SymbolKindClass:
		return "Class"
	case SymbolKindMethod:
		return "Method"
	case SymbolKindProperty:
		return "Property"
	case SymbolKindField:
		return "Field"
	case SymbolKindConstructor:
		return "Constructor"
	case SymbolKindEnum:
		return "Enum"
	case SymbolKindInterface:
		return "Interface"
	case SymbolKindFunction:
		return "Function"
	case SymbolKindVariable:
		return "Variable"
	case SymbolKindConstant:
		return "Constant"
	case SymbolKindString:
		return "String"
	case SymbolKindNumber:
		return "Number"
	case SymbolKindBoolean:
		return "Boolean"
	case SymbolKindArray:
		return "Array"
	case SymbolKindObject:
		return "Object"
	case SymbolKindKey:
		return "Key"
	case SymbolKindNull:
		return "Null"
	case SymbolKindEnumMember:
		return "EnumMember"
	case SymbolKindStruct:
		return "Struct"
	case SymbolKindEvent:
		return "Event"
	case SymbolKindOperator:
		return "Operator"
	case SymbolKindTypeParameter:
		return "TypeParameter"
	}
	return ""
}

// Symbol tags are extra annotations that tweak the rendering of a symbol.
//
// @since 3.16
type SymbolTag Uinteger

// Render a symbol as obsolete, usually using a strike-out.
const SymbolTagDeprecated SymbolTag = 1

// String returns "" or "Deprecated", depending on the value of `it`.
func (it SymbolTag) String() string {
	switch it {
	case SymbolTagDeprecated:
		return "Deprecated"
	}
	return ""
}

// Moniker uniqueness level to define scope of the moniker.
//
// @since 3.16.0
type UniquenessLevel String

// The moniker is only unique inside a document
const UniquenessLevelDocument UniquenessLevel = "document"

// The moniker is unique inside a project for which a dump got created
const UniquenessLevelProject UniquenessLevel = "project"

// The moniker is unique inside the group to which a project belongs
const UniquenessLevelGroup UniquenessLevel = "group"

// The moniker is unique inside the moniker scheme.
const UniquenessLevelScheme UniquenessLevel = "scheme"

// The moniker is globally unique
const UniquenessLevelGlobal UniquenessLevel = "global"

// String returns "" or "Document" or "Project" or "Group" or "Scheme" or "Global", depending on the value of `it`.
func (it UniquenessLevel) String() string {
	switch it {
	case UniquenessLevelDocument:
		return "Document"
	case UniquenessLevelProject:
		return "Project"
	case UniquenessLevelGroup:
		return "Group"
	case UniquenessLevelScheme:
		return "Scheme"
	case UniquenessLevelGlobal:
		return "Global"
	}
	return ""
}

// The moniker kind.
//
// @since 3.16.0
type MonikerKind String

// The moniker represent a symbol that is imported into a project
const MonikerKindImport MonikerKind = "import"

// The moniker represents a symbol that is exported from a project
const MonikerKindExport MonikerKind = "export"

// The moniker represents a symbol that is local to a project (e.g. a local
// variable of a function, a class not visible outside the project, ...)
const MonikerKindLocal MonikerKind = "local"

// String returns "" or "Import" or "Export" or "Local", depending on the value of `it`.
func (it MonikerKind) String() string {
	switch it {
	case MonikerKindImport:
		return "Import"
	case MonikerKindExport:
		return "Export"
	case MonikerKindLocal:
		return "Local"
	}
	return ""
}

// Inlay hint kinds.
//
// @since 3.17.0
type InlayHintKind Uinteger

// An inlay hint that for a type annotation.
const InlayHintKindType InlayHintKind = 1

// An inlay hint that is for a parameter.
const InlayHintKindParameter InlayHintKind = 2

// String returns "" or "Type" or "Parameter", depending on the value of `it`.
func (it InlayHintKind) String() string {
	switch it {
	case InlayHintKindType:
		return "Type"
	case InlayHintKindParameter:
		return "Parameter"
	}
	return ""
}

// The message type
type MessageType Uinteger

// An error message.
const MessageTypeError MessageType = 1

// A warning message.
const MessageTypeWarning MessageType = 2

// An information message.
const MessageTypeInfo MessageType = 3

// A log message.
const MessageTypeLog MessageType = 4

// String returns "" or "Error" or "Warning" or "Info" or "Log", depending on the value of `it`.
func (it MessageType) String() string {
	switch it {
	case MessageTypeError:
		return "Error"
	case MessageTypeWarning:
		return "Warning"
	case MessageTypeInfo:
		return "Info"
	case MessageTypeLog:
		return "Log"
	}
	return ""
}

// Defines how the host (editor) should sync
// document changes to the language server.
type TextDocumentSyncKind Uinteger

// Documents should not be synced at all.
const TextDocumentSyncKindNone TextDocumentSyncKind = 0

// Documents are synced by always sending the full content
// of the document.
const TextDocumentSyncKindFull TextDocumentSyncKind = 1

// Documents are synced by sending the full content on open.
// After that only incremental updates to the document are
// send.
const TextDocumentSyncKindIncremental TextDocumentSyncKind = 2

// String returns "" or "None" or "Full" or "Incremental", depending on the value of `it`.
func (it TextDocumentSyncKind) String() string {
	switch it {
	case TextDocumentSyncKindNone:
		return "None"
	case TextDocumentSyncKindFull:
		return "Full"
	case TextDocumentSyncKindIncremental:
		return "Incremental"
	}
	return ""
}

// Represents reasons why a text document is saved.
type TextDocumentSaveReason Uinteger

// Manually triggered, e.g. by the user pressing save, by starting debugging,
// or by an API call.
const TextDocumentSaveReasonManual TextDocumentSaveReason = 1

// Automatic after a delay.
const TextDocumentSaveReasonAfterDelay TextDocumentSaveReason = 2

// When the editor lost focus.
const TextDocumentSaveReasonFocusOut TextDocumentSaveReason = 3

// String returns "" or "Manual" or "AfterDelay" or "FocusOut", depending on the value of `it`.
func (it TextDocumentSaveReason) String() string {
	switch it {
	case TextDocumentSaveReasonManual:
		return "Manual"
	case TextDocumentSaveReasonAfterDelay:
		return "AfterDelay"
	case TextDocumentSaveReasonFocusOut:
		return "FocusOut"
	}
	return ""
}

// The kind of a completion entry.
type CompletionItemKind Uinteger

const CompletionItemKindText CompletionItemKind = 1

const CompletionItemKindMethod CompletionItemKind = 2

const CompletionItemKindFunction CompletionItemKind = 3

const CompletionItemKindConstructor CompletionItemKind = 4

const CompletionItemKindField CompletionItemKind = 5

const CompletionItemKindVariable CompletionItemKind = 6

const CompletionItemKindClass CompletionItemKind = 7

const CompletionItemKindInterface CompletionItemKind = 8

const CompletionItemKindModule CompletionItemKind = 9

const CompletionItemKindProperty CompletionItemKind = 10

const CompletionItemKindUnit CompletionItemKind = 11

const CompletionItemKindValue CompletionItemKind = 12

const CompletionItemKindEnum CompletionItemKind = 13

const CompletionItemKindKeyword CompletionItemKind = 14

const CompletionItemKindSnippet CompletionItemKind = 15

const CompletionItemKindColor CompletionItemKind = 16

const CompletionItemKindFile CompletionItemKind = 17

const CompletionItemKindReference CompletionItemKind = 18

const CompletionItemKindFolder CompletionItemKind = 19

const CompletionItemKindEnumMember CompletionItemKind = 20

const CompletionItemKindConstant CompletionItemKind = 21

const CompletionItemKindStruct CompletionItemKind = 22

const CompletionItemKindEvent CompletionItemKind = 23

const CompletionItemKindOperator CompletionItemKind = 24

const CompletionItemKindTypeParameter CompletionItemKind = 25

// String returns "" or "Text" or "Method" or "Function" or "Constructor" or "Field" or "Variable" or "Class" or "Interface" or "Module" or "Property" or "Unit" or "Value" or "Enum" or "Keyword" or "Snippet" or "Color" or "File" or "Reference" or "Folder" or "EnumMember" or "Constant" or "Struct" or "Event" or "Operator" or "TypeParameter", depending on the value of `it`.
func (it CompletionItemKind) String() string {
	switch it {
	case CompletionItemKindText:
		return "Text"
	case CompletionItemKindMethod:
		return "Method"
	case CompletionItemKindFunction:
		return "Function"
	case CompletionItemKindConstructor:
		return "Constructor"
	case CompletionItemKindField:
		return "Field"
	case CompletionItemKindVariable:
		return "Variable"
	case CompletionItemKindClass:
		return "Class"
	case CompletionItemKindInterface:
		return "Interface"
	case CompletionItemKindModule:
		return "Module"
	case CompletionItemKindProperty:
		return "Property"
	case CompletionItemKindUnit:
		return "Unit"
	case CompletionItemKindValue:
		return "Value"
	case CompletionItemKindEnum:
		return "Enum"
	case CompletionItemKindKeyword:
		return "Keyword"
	case CompletionItemKindSnippet:
		return "Snippet"
	case CompletionItemKindColor:
		return "Color"
	case CompletionItemKindFile:
		return "File"
	case CompletionItemKindReference:
		return "Reference"
	case CompletionItemKindFolder:
		return "Folder"
	case CompletionItemKindEnumMember:
		return "EnumMember"
	case CompletionItemKindConstant:
		return "Constant"
	case CompletionItemKindStruct:
		return "Struct"
	case CompletionItemKindEvent:
		return "Event"
	case CompletionItemKindOperator:
		return "Operator"
	case CompletionItemKindTypeParameter:
		return "TypeParameter"
	}
	return ""
}

// Completion item tags are extra annotations that tweak the rendering of a completion
// item.
//
// @since 3.15.0
type CompletionItemTag Uinteger

// Render a completion as obsolete, usually using a strike-out.
const CompletionItemTagDeprecated CompletionItemTag = 1

// String returns "" or "Deprecated", depending on the value of `it`.
func (it CompletionItemTag) String() string {
	switch it {
	case CompletionItemTagDeprecated:
		return "Deprecated"
	}
	return ""
}

// Defines whether the insert text in a completion item should be interpreted as
// plain text or a snippet.
type InsertTextFormat Uinteger

// The primary text to be inserted is treated as a plain string.
const InsertTextFormatPlainText InsertTextFormat = 1

// The primary text to be inserted is treated as a snippet.
//
// A snippet can define tab stops and placeholders with `$1`, `$2`
// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
// the end of the snippet. Placeholders with equal identifiers are linked,
// that is typing in one will update others too.
//
// See also: https://microsoft.github.io/language-server-protocol/specifications/specification-current/#snippet_syntax
const InsertTextFormatSnippet InsertTextFormat = 2

// String returns "" or "PlainText" or "Snippet", depending on the value of `it`.
func (it InsertTextFormat) String() string {
	switch it {
	case InsertTextFormatPlainText:
		return "PlainText"
	case InsertTextFormatSnippet:
		return "Snippet"
	}
	return ""
}

// How whitespace and indentation is handled during completion
// item insertion.
//
// @since 3.16.0
type InsertTextMode Uinteger

// The insertion or replace strings is taken as it is. If the
// value is multi line the lines below the cursor will be
// inserted using the indentation defined in the string value.
// The client will not apply any kind of adjustments to the
// string.
const InsertTextModeAsIs InsertTextMode = 1

// The editor adjusts leading whitespace of new lines so that
// they match the indentation up to the cursor of the line for
// which the item is accepted.
//
// Consider a line like this: <2tabs><cursor><3tabs>foo. Accepting a
// multi line completion item is indented using 2 tabs and all
// following lines inserted will be indented using 2 tabs as well.
const InsertTextModeAdjustIndentation InsertTextMode = 2

// String returns "" or "AsIs" or "AdjustIndentation", depending on the value of `it`.
func (it InsertTextMode) String() string {
	switch it {
	case InsertTextModeAsIs:
		return "AsIs"
	case InsertTextModeAdjustIndentation:
		return "AdjustIndentation"
	}
	return ""
}

// A document highlight kind.
type DocumentHighlightKind Uinteger

// A textual occurrence.
const DocumentHighlightKindText DocumentHighlightKind = 1

// Read-access of a symbol, like reading a variable.
const DocumentHighlightKindRead DocumentHighlightKind = 2

// Write-access of a symbol, like writing to a variable.
const DocumentHighlightKindWrite DocumentHighlightKind = 3

// String returns "" or "Text" or "Read" or "Write", depending on the value of `it`.
func (it DocumentHighlightKind) String() string {
	switch it {
	case DocumentHighlightKindText:
		return "Text"
	case DocumentHighlightKindRead:
		return "Read"
	case DocumentHighlightKindWrite:
		return "Write"
	}
	return ""
}

// A set of predefined code action kinds
type CodeActionKind String

// Empty kind.
const CodeActionKindEmpty CodeActionKind = ""

// Base kind for quickfix actions: 'quickfix'
const CodeActionKindQuickFix CodeActionKind = "quickfix"

// Base kind for refactoring actions: 'refactor'
const CodeActionKindRefactor CodeActionKind = "refactor"

// Base kind for refactoring extraction actions: 'refactor.extract'
//
// Example extract actions:
//
// - Extract method
// - Extract function
// - Extract variable
// - Extract interface from class
// - ...
const CodeActionKindRefactorExtract CodeActionKind = "refactor.extract"

// Base kind for refactoring inline actions: 'refactor.inline'
//
// Example inline actions:
//
// - Inline function
// - Inline variable
// - Inline constant
// - ...
const CodeActionKindRefactorInline CodeActionKind = "refactor.inline"

// Base kind for refactoring rewrite actions: 'refactor.rewrite'
//
// Example rewrite actions:
//
// - Convert JavaScript function to class
// - Add or remove parameter
// - Encapsulate field
// - Make method static
// - Move method to base class
// - ...
const CodeActionKindRefactorRewrite CodeActionKind = "refactor.rewrite"

// Base kind for source actions: `source`
//
// Source code actions apply to the entire file.
const CodeActionKindSource CodeActionKind = "source"

// Base kind for an organize imports source action: `source.organizeImports`
const CodeActionKindSourceOrganizeImports CodeActionKind = "source.organizeImports"

// Base kind for auto-fix source actions: `source.fixAll`.
//
// Fix all actions automatically fix errors that have a clear fix that do not require user input.
// They should not suppress errors or perform unsafe fixes such as generating new types or classes.
//
// @since 3.15.0
const CodeActionKindSourceFixAll CodeActionKind = "source.fixAll"

// String returns "" or "Empty" or "QuickFix" or "Refactor" or "RefactorExtract" or "RefactorInline" or "RefactorRewrite" or "Source" or "SourceOrganizeImports" or "SourceFixAll", depending on the value of `it`.
func (it CodeActionKind) String() string {
	switch it {
	case CodeActionKindEmpty:
		return "Empty"
	case CodeActionKindQuickFix:
		return "QuickFix"
	case CodeActionKindRefactor:
		return "Refactor"
	case CodeActionKindRefactorExtract:
		return "RefactorExtract"
	case CodeActionKindRefactorInline:
		return "RefactorInline"
	case CodeActionKindRefactorRewrite:
		return "RefactorRewrite"
	case CodeActionKindSource:
		return "Source"
	case CodeActionKindSourceOrganizeImports:
		return "SourceOrganizeImports"
	case CodeActionKindSourceFixAll:
		return "SourceFixAll"
	}
	return ""
}

type TraceValues String

// Turn tracing off.
const TraceValuesOff TraceValues = "off"

// Trace messages only.
const TraceValuesMessages TraceValues = "messages"

// Verbose message tracing.
const TraceValuesVerbose TraceValues = "verbose"

// String returns "" or "Off" or "Messages" or "Verbose", depending on the value of `it`.
func (it TraceValues) String() string {
	switch it {
	case TraceValuesOff:
		return "Off"
	case TraceValuesMessages:
		return "Messages"
	case TraceValuesVerbose:
		return "Verbose"
	}
	return ""
}

// Describes the content type that a client supports in various
// result literals like `Hover`, `ParameterInfo` or `CompletionItem`.
//
// Please note that `MarkupKinds` must not start with a `$`. This kinds
// are reserved for internal usage.
type MarkupKind String

// Plain text is supported as a content format
const MarkupKindPlainText MarkupKind = "plaintext"

// Markdown is supported as a content format
const MarkupKindMarkdown MarkupKind = "markdown"

// String returns "" or "PlainText" or "Markdown", depending on the value of `it`.
func (it MarkupKind) String() string {
	switch it {
	case MarkupKindPlainText:
		return "PlainText"
	case MarkupKindMarkdown:
		return "Markdown"
	}
	return ""
}

// A set of predefined position encoding kinds.
//
// @since 3.17.0
type PositionEncodingKind String

// Character offsets count UTF-8 code units.
const PositionEncodingKindUTF8 PositionEncodingKind = "utf-8"

// Character offsets count UTF-16 code units.
//
// This is the default and must always be supported
// by servers
const PositionEncodingKindUTF16 PositionEncodingKind = "utf-16"

// Character offsets count UTF-32 code units.
//
// Implementation note: these are the same as Unicode code points,
// so this `PositionEncodingKind` may also be used for an
// encoding-agnostic representation of character offsets.
const PositionEncodingKindUTF32 PositionEncodingKind = "utf-32"

// String returns "" or "UTF8" or "UTF16" or "UTF32", depending on the value of `it`.
func (it PositionEncodingKind) String() string {
	switch it {
	case PositionEncodingKindUTF8:
		return "UTF8"
	case PositionEncodingKindUTF16:
		return "UTF16"
	case PositionEncodingKindUTF32:
		return "UTF32"
	}
	return ""
}

// The file event type
type FileChangeType Uinteger

// The file got created.
const FileChangeTypeCreated FileChangeType = 1

// The file got changed.
const FileChangeTypeChanged FileChangeType = 2

// The file got deleted.
const FileChangeTypeDeleted FileChangeType = 3

// String returns "" or "Created" or "Changed" or "Deleted", depending on the value of `it`.
func (it FileChangeType) String() string {
	switch it {
	case FileChangeTypeCreated:
		return "Created"
	case FileChangeTypeChanged:
		return "Changed"
	case FileChangeTypeDeleted:
		return "Deleted"
	}
	return ""
}

type WatchKind Uinteger

// Interested in create events.
const WatchKindCreate WatchKind = 1

// Interested in change events
const WatchKindChange WatchKind = 2

// Interested in delete events
const WatchKindDelete WatchKind = 4

// String returns "" or "Create" or "Change" or "Delete", depending on the value of `it`.
func (it WatchKind) String() string {
	switch it {
	case WatchKindCreate:
		return "Create"
	case WatchKindChange:
		return "Change"
	case WatchKindDelete:
		return "Delete"
	}
	return ""
}

// The diagnostic's severity.
type DiagnosticSeverity Uinteger

// Reports an error.
const DiagnosticSeverityError DiagnosticSeverity = 1

// Reports a warning.
const DiagnosticSeverityWarning DiagnosticSeverity = 2

// Reports an information.
const DiagnosticSeverityInformation DiagnosticSeverity = 3

// Reports a hint.
const DiagnosticSeverityHint DiagnosticSeverity = 4

// String returns "" or "Error" or "Warning" or "Information" or "Hint", depending on the value of `it`.
func (it DiagnosticSeverity) String() string {
	switch it {
	case DiagnosticSeverityError:
		return "Error"
	case DiagnosticSeverityWarning:
		return "Warning"
	case DiagnosticSeverityInformation:
		return "Information"
	case DiagnosticSeverityHint:
		return "Hint"
	}
	return ""
}

// The diagnostic tags.
//
// @since 3.15.0
type DiagnosticTag Uinteger

// Unused or unnecessary code.
//
// Clients are allowed to render diagnostics with this tag faded out instead of having
// an error squiggle.
const DiagnosticTagUnnecessary DiagnosticTag = 1

// Deprecated or obsolete code.
//
// Clients are allowed to rendered diagnostics with this tag strike through.
const DiagnosticTagDeprecated DiagnosticTag = 2

// String returns "" or "Unnecessary" or "Deprecated", depending on the value of `it`.
func (it DiagnosticTag) String() string {
	switch it {
	case DiagnosticTagUnnecessary:
		return "Unnecessary"
	case DiagnosticTagDeprecated:
		return "Deprecated"
	}
	return ""
}

// How a completion was triggered
type CompletionTriggerKind Uinteger

// Completion was triggered by typing an identifier (24x7 code
// complete), manual invocation (e.g Ctrl+Space) or via API.
const CompletionTriggerKindInvoked CompletionTriggerKind = 1

// Completion was triggered by a trigger character specified by
// the `triggerCharacters` properties of the `CompletionRegistrationOptions`.
const CompletionTriggerKindTriggerCharacter CompletionTriggerKind = 2

// Completion was re-triggered as current completion list is incomplete
const CompletionTriggerKindTriggerForIncompleteCompletions CompletionTriggerKind = 3

// String returns "" or "Invoked" or "TriggerCharacter" or "TriggerForIncompleteCompletions", depending on the value of `it`.
func (it CompletionTriggerKind) String() string {
	switch it {
	case CompletionTriggerKindInvoked:
		return "Invoked"
	case CompletionTriggerKindTriggerCharacter:
		return "TriggerCharacter"
	case CompletionTriggerKindTriggerForIncompleteCompletions:
		return "TriggerForIncompleteCompletions"
	}
	return ""
}

// How a signature help was triggered.
//
// @since 3.15.0
type SignatureHelpTriggerKind Uinteger

// Signature help was invoked manually by the user or by a command.
const SignatureHelpTriggerKindInvoked SignatureHelpTriggerKind = 1

// Signature help was triggered by a trigger character.
const SignatureHelpTriggerKindTriggerCharacter SignatureHelpTriggerKind = 2

// Signature help was triggered by the cursor moving or by the document content changing.
const SignatureHelpTriggerKindContentChange SignatureHelpTriggerKind = 3

// String returns "" or "Invoked" or "TriggerCharacter" or "ContentChange", depending on the value of `it`.
func (it SignatureHelpTriggerKind) String() string {
	switch it {
	case SignatureHelpTriggerKindInvoked:
		return "Invoked"
	case SignatureHelpTriggerKindTriggerCharacter:
		return "TriggerCharacter"
	case SignatureHelpTriggerKindContentChange:
		return "ContentChange"
	}
	return ""
}

// The reason why code actions were requested.
//
// @since 3.17.0
type CodeActionTriggerKind Uinteger

// Code actions were explicitly requested by the user or by an extension.
const CodeActionTriggerKindInvoked CodeActionTriggerKind = 1

// Code actions were requested automatically.
//
// This typically happens when current selection in a file changes, but can
// also be triggered when file content changes.
const CodeActionTriggerKindAutomatic CodeActionTriggerKind = 2

// String returns "" or "Invoked" or "Automatic", depending on the value of `it`.
func (it CodeActionTriggerKind) String() string {
	switch it {
	case CodeActionTriggerKindInvoked:
		return "Invoked"
	case CodeActionTriggerKindAutomatic:
		return "Automatic"
	}
	return ""
}

// A pattern kind describing if a glob pattern matches a file a folder or
// both.
//
// @since 3.16.0
type FileOperationPatternKind String

// The pattern matches a file only.
const FileOperationPatternKindFile FileOperationPatternKind = "file"

// The pattern matches a folder only.
const FileOperationPatternKindFolder FileOperationPatternKind = "folder"

// String returns "" or "File" or "Folder", depending on the value of `it`.
func (it FileOperationPatternKind) String() string {
	switch it {
	case FileOperationPatternKindFile:
		return "File"
	case FileOperationPatternKindFolder:
		return "Folder"
	}
	return ""
}

// A notebook cell kind.
//
// @since 3.17.0
type NotebookCellKind Uinteger

// A markup-cell is formatted source that is used for display.
const NotebookCellKindMarkup NotebookCellKind = 1

// A code-cell is source code.
const NotebookCellKindCode NotebookCellKind = 2

// String returns "" or "Markup" or "Code", depending on the value of `it`.
func (it NotebookCellKind) String() string {
	switch it {
	case NotebookCellKindMarkup:
		return "Markup"
	case NotebookCellKindCode:
		return "Code"
	}
	return ""
}

type ResourceOperationKind String

// Supports creating new files and folders.
const ResourceOperationKindCreate ResourceOperationKind = "create"

// Supports renaming existing files and folders.
const ResourceOperationKindRename ResourceOperationKind = "rename"

// Supports deleting existing files and folders.
const ResourceOperationKindDelete ResourceOperationKind = "delete"

// String returns "" or "Create" or "Rename" or "Delete", depending on the value of `it`.
func (it ResourceOperationKind) String() string {
	switch it {
	case ResourceOperationKindCreate:
		return "Create"
	case ResourceOperationKindRename:
		return "Rename"
	case ResourceOperationKindDelete:
		return "Delete"
	}
	return ""
}

type FailureHandlingKind String

// Applying the workspace change is simply aborted if one of the changes provided
// fails. All operations executed before the failing operation stay executed.
const FailureHandlingKindAbort FailureHandlingKind = "abort"

// All operations are executed transactional. That means they either all
// succeed or no changes at all are applied to the workspace.
const FailureHandlingKindTransactional FailureHandlingKind = "transactional"

// If the workspace edit contains only textual file changes they are executed transactional.
// If resource changes (create, rename or delete file) are part of the change the failure
// handling strategy is abort.
const FailureHandlingKindTextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"

// The client tries to undo the operations already executed. But there is no
// guarantee that this is succeeding.
const FailureHandlingKindUndo FailureHandlingKind = "undo"

// String returns "" or "Abort" or "Transactional" or "TextOnlyTransactional" or "Undo", depending on the value of `it`.
func (it FailureHandlingKind) String() string {
	switch it {
	case FailureHandlingKindAbort:
		return "Abort"
	case FailureHandlingKindTransactional:
		return "Transactional"
	case FailureHandlingKindTextOnlyTransactional:
		return "TextOnlyTransactional"
	case FailureHandlingKindUndo:
		return "Undo"
	}
	return ""
}

type PrepareSupportDefaultBehavior Uinteger

// The client's default behavior is to select the identifier
// according the to language's syntax rule.
const PrepareSupportDefaultBehaviorIdentifier PrepareSupportDefaultBehavior = 1

// String returns "" or "Identifier", depending on the value of `it`.
func (it PrepareSupportDefaultBehavior) String() string {
	switch it {
	case PrepareSupportDefaultBehaviorIdentifier:
		return "Identifier"
	}
	return ""
}

type TokenFormat String

const TokenFormatRelative TokenFormat = "relative"

// String returns "" or "Relative", depending on the value of `it`.
func (it TokenFormat) String() string {
	switch it {
	case TokenFormatRelative:
		return "Relative"
	}
	return ""
}
