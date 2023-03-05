// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-vsx-and-lsp/gen/cmd/gen_lsp
package lsp




// A set of predefined token types. This set is not fixed
// an clients can specify additional token types via the
// corresponding client capabilities.
// 
// @since 3.16.0
type SemanticTokenTypes string

// The value is always "namespace".
const SemanticTokenTypesNamespace SemanticTokenTypes = "namespace"

// Represents a generic type. Acts as a fallback for types which can't be mapped to
// a specific type like class or enum.
// 
// The value is always "type".
const SemanticTokenTypesType SemanticTokenTypes = "type"

// The value is always "class".
const SemanticTokenTypesClass SemanticTokenTypes = "class"

// The value is always "enum".
const SemanticTokenTypesEnum SemanticTokenTypes = "enum"

// The value is always "interface".
const SemanticTokenTypesInterface SemanticTokenTypes = "interface"

// The value is always "struct".
const SemanticTokenTypesStruct SemanticTokenTypes = "struct"

// The value is always "typeParameter".
const SemanticTokenTypesTypeParameter SemanticTokenTypes = "typeParameter"

// The value is always "parameter".
const SemanticTokenTypesParameter SemanticTokenTypes = "parameter"

// The value is always "variable".
const SemanticTokenTypesVariable SemanticTokenTypes = "variable"

// The value is always "property".
const SemanticTokenTypesProperty SemanticTokenTypes = "property"

// The value is always "enumMember".
const SemanticTokenTypesEnumMember SemanticTokenTypes = "enumMember"

// The value is always "event".
const SemanticTokenTypesEvent SemanticTokenTypes = "event"

// The value is always "function".
const SemanticTokenTypesFunction SemanticTokenTypes = "function"

// The value is always "method".
const SemanticTokenTypesMethod SemanticTokenTypes = "method"

// The value is always "macro".
const SemanticTokenTypesMacro SemanticTokenTypes = "macro"

// The value is always "keyword".
const SemanticTokenTypesKeyword SemanticTokenTypes = "keyword"

// The value is always "modifier".
const SemanticTokenTypesModifier SemanticTokenTypes = "modifier"

// The value is always "comment".
const SemanticTokenTypesComment SemanticTokenTypes = "comment"

// The value is always "string".
const SemanticTokenTypesString SemanticTokenTypes = "string"

// The value is always "number".
const SemanticTokenTypesNumber SemanticTokenTypes = "number"

// The value is always "regexp".
const SemanticTokenTypesRegexp SemanticTokenTypes = "regexp"

// The value is always "operator".
const SemanticTokenTypesOperator SemanticTokenTypes = "operator"

// @since 3.17.0
// 
// The value is always "decorator".
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
type SemanticTokenModifiers string

// The value is always "declaration".
const SemanticTokenModifiersDeclaration SemanticTokenModifiers = "declaration"

// The value is always "definition".
const SemanticTokenModifiersDefinition SemanticTokenModifiers = "definition"

// The value is always "readonly".
const SemanticTokenModifiersReadonly SemanticTokenModifiers = "readonly"

// The value is always "static".
const SemanticTokenModifiersStatic SemanticTokenModifiers = "static"

// The value is always "deprecated".
const SemanticTokenModifiersDeprecated SemanticTokenModifiers = "deprecated"

// The value is always "abstract".
const SemanticTokenModifiersAbstract SemanticTokenModifiers = "abstract"

// The value is always "async".
const SemanticTokenModifiersAsync SemanticTokenModifiers = "async"

// The value is always "modification".
const SemanticTokenModifiersModification SemanticTokenModifiers = "modification"

// The value is always "documentation".
const SemanticTokenModifiersDocumentation SemanticTokenModifiers = "documentation"

// The value is always "defaultLibrary".
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
type DocumentDiagnosticReportKind string

// A diagnostic report with a full
// set of problems.
// 
// The value is always "full".
const DocumentDiagnosticReportKindFull DocumentDiagnosticReportKind = "full"

// A report indicating that the last
// returned report is still accurate.
// 
// The value is always "unchanged".
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
type ErrorCodes int

// The value is always -32700.
const ErrorCodesParseError ErrorCodes = -32700

// The value is always -32600.
const ErrorCodesInvalidRequest ErrorCodes = -32600

// The value is always -32601.
const ErrorCodesMethodNotFound ErrorCodes = -32601

// The value is always -32602.
const ErrorCodesInvalidParams ErrorCodes = -32602

// The value is always -32603.
const ErrorCodesInternalError ErrorCodes = -32603

// Error code indicating that a server received a notification or
// request before the server has received the `initialize` request.
// 
// The value is always -32002.
const ErrorCodesServerNotInitialized ErrorCodes = -32002

// The value is always -32001.
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


type LSPErrorCodes int

// A request failed but it was syntactically correct, e.g the
// method name was known and the parameters were valid. The error
// message should contain human readable information about why
// the request failed.
// 
// @since 3.17.0
// 
// The value is always -32803.
const LSPErrorCodesRequestFailed LSPErrorCodes = -32803

// The server cancelled the request. This error code should
// only be used for requests that explicitly support being
// server cancellable.
// 
// @since 3.17.0
// 
// The value is always -32802.
const LSPErrorCodesServerCancelled LSPErrorCodes = -32802

// The server detected that the content of a document got
// modified outside normal conditions. A server should
// NOT send this error code if it detects a content change
// in it unprocessed messages. The result even computed
// on an older state might still be useful for the client.
// 
// If a client decides that a result is not of any use anymore
// the client should cancel the request.
// 
// The value is always -32801.
const LSPErrorCodesContentModified LSPErrorCodes = -32801

// The client has canceled a request and a server as detected
// the cancel.
// 
// The value is always -32800.
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
type FoldingRangeKind string

// Folding range for a comment
// 
// The value is always "comment".
const FoldingRangeKindComment FoldingRangeKind = "comment"

// Folding range for an import or include
// 
// The value is always "imports".
const FoldingRangeKindImports FoldingRangeKind = "imports"

// Folding range for a region (e.g. `#region`)
// 
// The value is always "region".
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
type SymbolKind uint

// The value is always 1.
const SymbolKindFile SymbolKind = 1

// The value is always 2.
const SymbolKindModule SymbolKind = 2

// The value is always 3.
const SymbolKindNamespace SymbolKind = 3

// The value is always 4.
const SymbolKindPackage SymbolKind = 4

// The value is always 5.
const SymbolKindClass SymbolKind = 5

// The value is always 6.
const SymbolKindMethod SymbolKind = 6

// The value is always 7.
const SymbolKindProperty SymbolKind = 7

// The value is always 8.
const SymbolKindField SymbolKind = 8

// The value is always 9.
const SymbolKindConstructor SymbolKind = 9

// The value is always 10.
const SymbolKindEnum SymbolKind = 10

// The value is always 11.
const SymbolKindInterface SymbolKind = 11

// The value is always 12.
const SymbolKindFunction SymbolKind = 12

// The value is always 13.
const SymbolKindVariable SymbolKind = 13

// The value is always 14.
const SymbolKindConstant SymbolKind = 14

// The value is always 15.
const SymbolKindString SymbolKind = 15

// The value is always 16.
const SymbolKindNumber SymbolKind = 16

// The value is always 17.
const SymbolKindBoolean SymbolKind = 17

// The value is always 18.
const SymbolKindArray SymbolKind = 18

// The value is always 19.
const SymbolKindObject SymbolKind = 19

// The value is always 20.
const SymbolKindKey SymbolKind = 20

// The value is always 21.
const SymbolKindNull SymbolKind = 21

// The value is always 22.
const SymbolKindEnumMember SymbolKind = 22

// The value is always 23.
const SymbolKindStruct SymbolKind = 23

// The value is always 24.
const SymbolKindEvent SymbolKind = 24

// The value is always 25.
const SymbolKindOperator SymbolKind = 25

// The value is always 26.
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
type SymbolTag uint

// Render a symbol as obsolete, usually using a strike-out.
// 
// The value is always 1.
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
type UniquenessLevel string

// The moniker is only unique inside a document
// 
// The value is always "document".
const UniquenessLevelDocument UniquenessLevel = "document"

// The moniker is unique inside a project for which a dump got created
// 
// The value is always "project".
const UniquenessLevelProject UniquenessLevel = "project"

// The moniker is unique inside the group to which a project belongs
// 
// The value is always "group".
const UniquenessLevelGroup UniquenessLevel = "group"

// The moniker is unique inside the moniker scheme.
// 
// The value is always "scheme".
const UniquenessLevelScheme UniquenessLevel = "scheme"

// The moniker is globally unique
// 
// The value is always "global".
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
type MonikerKind string

// The moniker represent a symbol that is imported into a project
// 
// The value is always "import".
const MonikerKindImport MonikerKind = "import"

// The moniker represents a symbol that is exported from a project
// 
// The value is always "export".
const MonikerKindExport MonikerKind = "export"

// The moniker represents a symbol that is local to a project (e.g. a local
// variable of a function, a class not visible outside the project, ...)
// 
// The value is always "local".
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
type InlayHintKind uint

// An inlay hint that for a type annotation.
// 
// The value is always 1.
const InlayHintKindType InlayHintKind = 1

// An inlay hint that is for a parameter.
// 
// The value is always 2.
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
type MessageType uint

// An error message.
// 
// The value is always 1.
const MessageTypeError MessageType = 1

// A warning message.
// 
// The value is always 2.
const MessageTypeWarning MessageType = 2

// An information message.
// 
// The value is always 3.
const MessageTypeInfo MessageType = 3

// A log message.
// 
// The value is always 4.
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
type TextDocumentSyncKind uint

// Documents should not be synced at all.
// 
// The value is always 0.
const TextDocumentSyncKindNone TextDocumentSyncKind = 0

// Documents are synced by always sending the full content
// of the document.
// 
// The value is always 1.
const TextDocumentSyncKindFull TextDocumentSyncKind = 1

// Documents are synced by sending the full content on open.
// After that only incremental updates to the document are
// send.
// 
// The value is always 2.
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
type TextDocumentSaveReason uint

// Manually triggered, e.g. by the user pressing save, by starting debugging,
// or by an API call.
// 
// The value is always 1.
const TextDocumentSaveReasonManual TextDocumentSaveReason = 1

// Automatic after a delay.
// 
// The value is always 2.
const TextDocumentSaveReasonAfterDelay TextDocumentSaveReason = 2

// When the editor lost focus.
// 
// The value is always 3.
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
type CompletionItemKind uint

// The value is always 1.
const CompletionItemKindText CompletionItemKind = 1

// The value is always 2.
const CompletionItemKindMethod CompletionItemKind = 2

// The value is always 3.
const CompletionItemKindFunction CompletionItemKind = 3

// The value is always 4.
const CompletionItemKindConstructor CompletionItemKind = 4

// The value is always 5.
const CompletionItemKindField CompletionItemKind = 5

// The value is always 6.
const CompletionItemKindVariable CompletionItemKind = 6

// The value is always 7.
const CompletionItemKindClass CompletionItemKind = 7

// The value is always 8.
const CompletionItemKindInterface CompletionItemKind = 8

// The value is always 9.
const CompletionItemKindModule CompletionItemKind = 9

// The value is always 10.
const CompletionItemKindProperty CompletionItemKind = 10

// The value is always 11.
const CompletionItemKindUnit CompletionItemKind = 11

// The value is always 12.
const CompletionItemKindValue CompletionItemKind = 12

// The value is always 13.
const CompletionItemKindEnum CompletionItemKind = 13

// The value is always 14.
const CompletionItemKindKeyword CompletionItemKind = 14

// The value is always 15.
const CompletionItemKindSnippet CompletionItemKind = 15

// The value is always 16.
const CompletionItemKindColor CompletionItemKind = 16

// The value is always 17.
const CompletionItemKindFile CompletionItemKind = 17

// The value is always 18.
const CompletionItemKindReference CompletionItemKind = 18

// The value is always 19.
const CompletionItemKindFolder CompletionItemKind = 19

// The value is always 20.
const CompletionItemKindEnumMember CompletionItemKind = 20

// The value is always 21.
const CompletionItemKindConstant CompletionItemKind = 21

// The value is always 22.
const CompletionItemKindStruct CompletionItemKind = 22

// The value is always 23.
const CompletionItemKindEvent CompletionItemKind = 23

// The value is always 24.
const CompletionItemKindOperator CompletionItemKind = 24

// The value is always 25.
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
type CompletionItemTag uint

// Render a completion as obsolete, usually using a strike-out.
// 
// The value is always 1.
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
type InsertTextFormat uint

// The primary text to be inserted is treated as a plain string.
// 
// The value is always 1.
const InsertTextFormatPlainText InsertTextFormat = 1

// The primary text to be inserted is treated as a snippet.
// 
// A snippet can define tab stops and placeholders with `$1`, `$2`
// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
// the end of the snippet. Placeholders with equal identifiers are linked,
// that is typing in one will update others too.
// 
// See also: https://microsoft.github.io/language-server-protocol/specifications/specification-current/#snippet_syntax
// 
// The value is always 2.
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
type InsertTextMode uint

// The insertion or replace strings is taken as it is. If the
// value is multi line the lines below the cursor will be
// inserted using the indentation defined in the string value.
// The client will not apply any kind of adjustments to the
// string.
// 
// The value is always 1.
const InsertTextModeAsIs InsertTextMode = 1

// The editor adjusts leading whitespace of new lines so that
// they match the indentation up to the cursor of the line for
// which the item is accepted.
// 
// Consider a line like this: <2tabs><cursor><3tabs>foo. Accepting a
// multi line completion item is indented using 2 tabs and all
// following lines inserted will be indented using 2 tabs as well.
// 
// The value is always 2.
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
type DocumentHighlightKind uint

// A textual occurrence.
// 
// The value is always 1.
const DocumentHighlightKindText DocumentHighlightKind = 1

// Read-access of a symbol, like reading a variable.
// 
// The value is always 2.
const DocumentHighlightKindRead DocumentHighlightKind = 2

// Write-access of a symbol, like writing to a variable.
// 
// The value is always 3.
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
type CodeActionKind string

// Empty kind.
// 
// The value is always "".
const CodeActionKindEmpty CodeActionKind = ""

// Base kind for quickfix actions: 'quickfix'
// 
// The value is always "quickfix".
const CodeActionKindQuickFix CodeActionKind = "quickfix"

// Base kind for refactoring actions: 'refactor'
// 
// The value is always "refactor".
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
// 
// The value is always "refactor.extract".
const CodeActionKindRefactorExtract CodeActionKind = "refactor.extract"

// Base kind for refactoring inline actions: 'refactor.inline'
// 
// Example inline actions:
// 
// - Inline function
// - Inline variable
// - Inline constant
// - ...
// 
// The value is always "refactor.inline".
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
// 
// The value is always "refactor.rewrite".
const CodeActionKindRefactorRewrite CodeActionKind = "refactor.rewrite"

// Base kind for source actions: `source`
// 
// Source code actions apply to the entire file.
// 
// The value is always "source".
const CodeActionKindSource CodeActionKind = "source"

// Base kind for an organize imports source action: `source.organizeImports`
// 
// The value is always "source.organizeImports".
const CodeActionKindSourceOrganizeImports CodeActionKind = "source.organizeImports"

// Base kind for auto-fix source actions: `source.fixAll`.
// 
// Fix all actions automatically fix errors that have a clear fix that do not require user input.
// They should not suppress errors or perform unsafe fixes such as generating new types or classes.
// 
// @since 3.15.0
// 
// The value is always "source.fixAll".
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


type TraceValues string

// Turn tracing off.
// 
// The value is always "off".
const TraceValuesOff TraceValues = "off"

// Trace messages only.
// 
// The value is always "messages".
const TraceValuesMessages TraceValues = "messages"

// Verbose message tracing.
// 
// The value is always "verbose".
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
type MarkupKind string

// Plain text is supported as a content format
// 
// The value is always "plaintext".
const MarkupKindPlainText MarkupKind = "plaintext"

// Markdown is supported as a content format
// 
// The value is always "markdown".
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
type PositionEncodingKind string

// Character offsets count UTF-8 code units.
// 
// The value is always "utf-8".
const PositionEncodingKindUTF8 PositionEncodingKind = "utf-8"

// Character offsets count UTF-16 code units.
// 
// This is the default and must always be supported
// by servers
// 
// The value is always "utf-16".
const PositionEncodingKindUTF16 PositionEncodingKind = "utf-16"

// Character offsets count UTF-32 code units.
// 
// Implementation note: these are the same as Unicode code points,
// so this `PositionEncodingKind` may also be used for an
// encoding-agnostic representation of character offsets.
// 
// The value is always "utf-32".
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
type FileChangeType uint

// The file got created.
// 
// The value is always 1.
const FileChangeTypeCreated FileChangeType = 1

// The file got changed.
// 
// The value is always 2.
const FileChangeTypeChanged FileChangeType = 2

// The file got deleted.
// 
// The value is always 3.
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


type WatchKind uint

// Interested in create events.
// 
// The value is always 1.
const WatchKindCreate WatchKind = 1

// Interested in change events
// 
// The value is always 2.
const WatchKindChange WatchKind = 2

// Interested in delete events
// 
// The value is always 4.
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
type DiagnosticSeverity uint

// Reports an error.
// 
// The value is always 1.
const DiagnosticSeverityError DiagnosticSeverity = 1

// Reports a warning.
// 
// The value is always 2.
const DiagnosticSeverityWarning DiagnosticSeverity = 2

// Reports an information.
// 
// The value is always 3.
const DiagnosticSeverityInformation DiagnosticSeverity = 3

// Reports a hint.
// 
// The value is always 4.
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
type DiagnosticTag uint

// Unused or unnecessary code.
// 
// Clients are allowed to render diagnostics with this tag faded out instead of having
// an error squiggle.
// 
// The value is always 1.
const DiagnosticTagUnnecessary DiagnosticTag = 1

// Deprecated or obsolete code.
// 
// Clients are allowed to rendered diagnostics with this tag strike through.
// 
// The value is always 2.
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
type CompletionTriggerKind uint

// Completion was triggered by typing an identifier (24x7 code
// complete), manual invocation (e.g Ctrl+Space) or via API.
// 
// The value is always 1.
const CompletionTriggerKindInvoked CompletionTriggerKind = 1

// Completion was triggered by a trigger character specified by
// the `triggerCharacters` properties of the `CompletionRegistrationOptions`.
// 
// The value is always 2.
const CompletionTriggerKindTriggerCharacter CompletionTriggerKind = 2

// Completion was re-triggered as current completion list is incomplete
// 
// The value is always 3.
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
type SignatureHelpTriggerKind uint

// Signature help was invoked manually by the user or by a command.
// 
// The value is always 1.
const SignatureHelpTriggerKindInvoked SignatureHelpTriggerKind = 1

// Signature help was triggered by a trigger character.
// 
// The value is always 2.
const SignatureHelpTriggerKindTriggerCharacter SignatureHelpTriggerKind = 2

// Signature help was triggered by the cursor moving or by the document content changing.
// 
// The value is always 3.
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
type CodeActionTriggerKind uint

// Code actions were explicitly requested by the user or by an extension.
// 
// The value is always 1.
const CodeActionTriggerKindInvoked CodeActionTriggerKind = 1

// Code actions were requested automatically.
// 
// This typically happens when current selection in a file changes, but can
// also be triggered when file content changes.
// 
// The value is always 2.
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
type FileOperationPatternKind string

// The pattern matches a file only.
// 
// The value is always "file".
const FileOperationPatternKindFile FileOperationPatternKind = "file"

// The pattern matches a folder only.
// 
// The value is always "folder".
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
type NotebookCellKind uint

// A markup-cell is formatted source that is used for display.
// 
// The value is always 1.
const NotebookCellKindMarkup NotebookCellKind = 1

// A code-cell is source code.
// 
// The value is always 2.
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


type ResourceOperationKind string

// Supports creating new files and folders.
// 
// The value is always "create".
const ResourceOperationKindCreate ResourceOperationKind = "create"

// Supports renaming existing files and folders.
// 
// The value is always "rename".
const ResourceOperationKindRename ResourceOperationKind = "rename"

// Supports deleting existing files and folders.
// 
// The value is always "delete".
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


type FailureHandlingKind string

// Applying the workspace change is simply aborted if one of the changes provided
// fails. All operations executed before the failing operation stay executed.
// 
// The value is always "abort".
const FailureHandlingKindAbort FailureHandlingKind = "abort"

// All operations are executed transactional. That means they either all
// succeed or no changes at all are applied to the workspace.
// 
// The value is always "transactional".
const FailureHandlingKindTransactional FailureHandlingKind = "transactional"

// If the workspace edit contains only textual file changes they are executed transactional.
// If resource changes (create, rename or delete file) are part of the change the failure
// handling strategy is abort.
// 
// The value is always "textOnlyTransactional".
const FailureHandlingKindTextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"

// The client tries to undo the operations already executed. But there is no
// guarantee that this is succeeding.
// 
// The value is always "undo".
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


type PrepareSupportDefaultBehavior uint

// The client's default behavior is to select the identifier
// according the to language's syntax rule.
// 
// The value is always 1.
const PrepareSupportDefaultBehaviorIdentifier PrepareSupportDefaultBehavior = 1

// String returns "" or "Identifier", depending on the value of `it`.
func (it PrepareSupportDefaultBehavior) String() string {
	switch it {
		case PrepareSupportDefaultBehaviorIdentifier:
			return "Identifier"
	}
	return ""
}


type TokenFormat string

// The value is always "relative".
const TokenFormatRelative TokenFormat = "relative"

// String returns "" or "Relative", depending on the value of `it`.
func (it TokenFormat) String() string {
	switch it {
		case TokenFormatRelative:
			return "Relative"
	}
	return ""
}

