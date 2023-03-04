package lsp




type SemanticTokenTypes ?

const (
	SemanticTokenTypesnamespace SemanticTokenTypes = "namespace"

	SemanticTokenTypestype SemanticTokenTypes = "type"

	SemanticTokenTypesclass SemanticTokenTypes = "class"

	SemanticTokenTypesenum SemanticTokenTypes = "enum"

	SemanticTokenTypesinterface SemanticTokenTypes = "interface"

	SemanticTokenTypesstruct SemanticTokenTypes = "struct"

	SemanticTokenTypestypeParameter SemanticTokenTypes = "typeParameter"

	SemanticTokenTypesparameter SemanticTokenTypes = "parameter"

	SemanticTokenTypesvariable SemanticTokenTypes = "variable"

	SemanticTokenTypesproperty SemanticTokenTypes = "property"

	SemanticTokenTypesenumMember SemanticTokenTypes = "enumMember"

	SemanticTokenTypesevent SemanticTokenTypes = "event"

	SemanticTokenTypesfunction SemanticTokenTypes = "function"

	SemanticTokenTypesmethod SemanticTokenTypes = "method"

	SemanticTokenTypesmacro SemanticTokenTypes = "macro"

	SemanticTokenTypeskeyword SemanticTokenTypes = "keyword"

	SemanticTokenTypesmodifier SemanticTokenTypes = "modifier"

	SemanticTokenTypescomment SemanticTokenTypes = "comment"

	SemanticTokenTypesstring SemanticTokenTypes = "string"

	SemanticTokenTypesnumber SemanticTokenTypes = "number"

	SemanticTokenTypesregexp SemanticTokenTypes = "regexp"

	SemanticTokenTypesoperator SemanticTokenTypes = "operator"

	SemanticTokenTypesdecorator SemanticTokenTypes = "decorator"
)





type SemanticTokenModifiers ?

const (
	SemanticTokenModifiersdeclaration SemanticTokenModifiers = "declaration"

	SemanticTokenModifiersdefinition SemanticTokenModifiers = "definition"

	SemanticTokenModifiersreadonly SemanticTokenModifiers = "readonly"

	SemanticTokenModifiersstatic SemanticTokenModifiers = "static"

	SemanticTokenModifiersdeprecated SemanticTokenModifiers = "deprecated"

	SemanticTokenModifiersabstract SemanticTokenModifiers = "abstract"

	SemanticTokenModifiersasync SemanticTokenModifiers = "async"

	SemanticTokenModifiersmodification SemanticTokenModifiers = "modification"

	SemanticTokenModifiersdocumentation SemanticTokenModifiers = "documentation"

	SemanticTokenModifiersdefaultLibrary SemanticTokenModifiers = "defaultLibrary"
)





type DocumentDiagnosticReportKind ?

const (
	DocumentDiagnosticReportKindFull DocumentDiagnosticReportKind = "full"

	DocumentDiagnosticReportKindUnchanged DocumentDiagnosticReportKind = "unchanged"
)





type ErrorCodes ?

const (
	ErrorCodesParseError ErrorCodes = -32700

	ErrorCodesInvalidRequest ErrorCodes = -32600

	ErrorCodesMethodNotFound ErrorCodes = -32601

	ErrorCodesInvalidParams ErrorCodes = -32602

	ErrorCodesInternalError ErrorCodes = -32603

	ErrorCodesServerNotInitialized ErrorCodes = -32002

	ErrorCodesUnknownErrorCode ErrorCodes = -32001
)





type LSPErrorCodes ?

const (
	LSPErrorCodesRequestFailed LSPErrorCodes = -32803

	LSPErrorCodesServerCancelled LSPErrorCodes = -32802

	LSPErrorCodesContentModified LSPErrorCodes = -32801

	LSPErrorCodesRequestCancelled LSPErrorCodes = -32800
)





type FoldingRangeKind ?

const (
	FoldingRangeKindComment FoldingRangeKind = "comment"

	FoldingRangeKindImports FoldingRangeKind = "imports"

	FoldingRangeKindRegion FoldingRangeKind = "region"
)





type SymbolKind ?

const (
	SymbolKindFile SymbolKind = 1

	SymbolKindModule SymbolKind = 2

	SymbolKindNamespace SymbolKind = 3

	SymbolKindPackage SymbolKind = 4

	SymbolKindClass SymbolKind = 5

	SymbolKindMethod SymbolKind = 6

	SymbolKindProperty SymbolKind = 7

	SymbolKindField SymbolKind = 8

	SymbolKindConstructor SymbolKind = 9

	SymbolKindEnum SymbolKind = 10

	SymbolKindInterface SymbolKind = 11

	SymbolKindFunction SymbolKind = 12

	SymbolKindVariable SymbolKind = 13

	SymbolKindConstant SymbolKind = 14

	SymbolKindString SymbolKind = 15

	SymbolKindNumber SymbolKind = 16

	SymbolKindBoolean SymbolKind = 17

	SymbolKindArray SymbolKind = 18

	SymbolKindObject SymbolKind = 19

	SymbolKindKey SymbolKind = 20

	SymbolKindNull SymbolKind = 21

	SymbolKindEnumMember SymbolKind = 22

	SymbolKindStruct SymbolKind = 23

	SymbolKindEvent SymbolKind = 24

	SymbolKindOperator SymbolKind = 25

	SymbolKindTypeParameter SymbolKind = 26
)





type SymbolTag ?

const (
	SymbolTagDeprecated SymbolTag = 1
)





type UniquenessLevel ?

const (
	UniquenessLeveldocument UniquenessLevel = "document"

	UniquenessLevelproject UniquenessLevel = "project"

	UniquenessLevelgroup UniquenessLevel = "group"

	UniquenessLevelscheme UniquenessLevel = "scheme"

	UniquenessLevelglobal UniquenessLevel = "global"
)





type MonikerKind ?

const (
	MonikerKindimport MonikerKind = "import"

	MonikerKindexport MonikerKind = "export"

	MonikerKindlocal MonikerKind = "local"
)





type InlayHintKind ?

const (
	InlayHintKindType InlayHintKind = 1

	InlayHintKindParameter InlayHintKind = 2
)





type MessageType ?

const (
	MessageTypeError MessageType = 1

	MessageTypeWarning MessageType = 2

	MessageTypeInfo MessageType = 3

	MessageTypeLog MessageType = 4
)





type TextDocumentSyncKind ?

const (
	TextDocumentSyncKindNone TextDocumentSyncKind = 0

	TextDocumentSyncKindFull TextDocumentSyncKind = 1

	TextDocumentSyncKindIncremental TextDocumentSyncKind = 2
)





type TextDocumentSaveReason ?

const (
	TextDocumentSaveReasonManual TextDocumentSaveReason = 1

	TextDocumentSaveReasonAfterDelay TextDocumentSaveReason = 2

	TextDocumentSaveReasonFocusOut TextDocumentSaveReason = 3
)





type CompletionItemKind ?

const (
	CompletionItemKindText CompletionItemKind = 1

	CompletionItemKindMethod CompletionItemKind = 2

	CompletionItemKindFunction CompletionItemKind = 3

	CompletionItemKindConstructor CompletionItemKind = 4

	CompletionItemKindField CompletionItemKind = 5

	CompletionItemKindVariable CompletionItemKind = 6

	CompletionItemKindClass CompletionItemKind = 7

	CompletionItemKindInterface CompletionItemKind = 8

	CompletionItemKindModule CompletionItemKind = 9

	CompletionItemKindProperty CompletionItemKind = 10

	CompletionItemKindUnit CompletionItemKind = 11

	CompletionItemKindValue CompletionItemKind = 12

	CompletionItemKindEnum CompletionItemKind = 13

	CompletionItemKindKeyword CompletionItemKind = 14

	CompletionItemKindSnippet CompletionItemKind = 15

	CompletionItemKindColor CompletionItemKind = 16

	CompletionItemKindFile CompletionItemKind = 17

	CompletionItemKindReference CompletionItemKind = 18

	CompletionItemKindFolder CompletionItemKind = 19

	CompletionItemKindEnumMember CompletionItemKind = 20

	CompletionItemKindConstant CompletionItemKind = 21

	CompletionItemKindStruct CompletionItemKind = 22

	CompletionItemKindEvent CompletionItemKind = 23

	CompletionItemKindOperator CompletionItemKind = 24

	CompletionItemKindTypeParameter CompletionItemKind = 25
)





type CompletionItemTag ?

const (
	CompletionItemTagDeprecated CompletionItemTag = 1
)





type InsertTextFormat ?

const (
	InsertTextFormatPlainText InsertTextFormat = 1

	InsertTextFormatSnippet InsertTextFormat = 2
)





type InsertTextMode ?

const (
	InsertTextModeasIs InsertTextMode = 1

	InsertTextModeadjustIndentation InsertTextMode = 2
)





type DocumentHighlightKind ?

const (
	DocumentHighlightKindText DocumentHighlightKind = 1

	DocumentHighlightKindRead DocumentHighlightKind = 2

	DocumentHighlightKindWrite DocumentHighlightKind = 3
)





type CodeActionKind ?

const (
	CodeActionKindEmpty CodeActionKind = ""

	CodeActionKindQuickFix CodeActionKind = "quickfix"

	CodeActionKindRefactor CodeActionKind = "refactor"

	CodeActionKindRefactorExtract CodeActionKind = "refactor.extract"

	CodeActionKindRefactorInline CodeActionKind = "refactor.inline"

	CodeActionKindRefactorRewrite CodeActionKind = "refactor.rewrite"

	CodeActionKindSource CodeActionKind = "source"

	CodeActionKindSourceOrganizeImports CodeActionKind = "source.organizeImports"

	CodeActionKindSourceFixAll CodeActionKind = "source.fixAll"
)





type TraceValues ?

const (
	TraceValuesOff TraceValues = "off"

	TraceValuesMessages TraceValues = "messages"

	TraceValuesVerbose TraceValues = "verbose"
)





type MarkupKind ?

const (
	MarkupKindPlainText MarkupKind = "plaintext"

	MarkupKindMarkdown MarkupKind = "markdown"
)





type PositionEncodingKind ?

const (
	PositionEncodingKindUTF8 PositionEncodingKind = "utf-8"

	PositionEncodingKindUTF16 PositionEncodingKind = "utf-16"

	PositionEncodingKindUTF32 PositionEncodingKind = "utf-32"
)





type FileChangeType ?

const (
	FileChangeTypeCreated FileChangeType = 1

	FileChangeTypeChanged FileChangeType = 2

	FileChangeTypeDeleted FileChangeType = 3
)





type WatchKind ?

const (
	WatchKindCreate WatchKind = 1

	WatchKindChange WatchKind = 2

	WatchKindDelete WatchKind = 4
)





type DiagnosticSeverity ?

const (
	DiagnosticSeverityError DiagnosticSeverity = 1

	DiagnosticSeverityWarning DiagnosticSeverity = 2

	DiagnosticSeverityInformation DiagnosticSeverity = 3

	DiagnosticSeverityHint DiagnosticSeverity = 4
)





type DiagnosticTag ?

const (
	DiagnosticTagUnnecessary DiagnosticTag = 1

	DiagnosticTagDeprecated DiagnosticTag = 2
)





type CompletionTriggerKind ?

const (
	CompletionTriggerKindInvoked CompletionTriggerKind = 1

	CompletionTriggerKindTriggerCharacter CompletionTriggerKind = 2

	CompletionTriggerKindTriggerForIncompleteCompletions CompletionTriggerKind = 3
)





type SignatureHelpTriggerKind ?

const (
	SignatureHelpTriggerKindInvoked SignatureHelpTriggerKind = 1

	SignatureHelpTriggerKindTriggerCharacter SignatureHelpTriggerKind = 2

	SignatureHelpTriggerKindContentChange SignatureHelpTriggerKind = 3
)





type CodeActionTriggerKind ?

const (
	CodeActionTriggerKindInvoked CodeActionTriggerKind = 1

	CodeActionTriggerKindAutomatic CodeActionTriggerKind = 2
)





type FileOperationPatternKind ?

const (
	FileOperationPatternKindfile FileOperationPatternKind = "file"

	FileOperationPatternKindfolder FileOperationPatternKind = "folder"
)





type NotebookCellKind ?

const (
	NotebookCellKindMarkup NotebookCellKind = 1

	NotebookCellKindCode NotebookCellKind = 2
)





type ResourceOperationKind ?

const (
	ResourceOperationKindCreate ResourceOperationKind = "create"

	ResourceOperationKindRename ResourceOperationKind = "rename"

	ResourceOperationKindDelete ResourceOperationKind = "delete"
)





type FailureHandlingKind ?

const (
	FailureHandlingKindAbort FailureHandlingKind = "abort"

	FailureHandlingKindTransactional FailureHandlingKind = "transactional"

	FailureHandlingKindTextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"

	FailureHandlingKindUndo FailureHandlingKind = "undo"
)





type PrepareSupportDefaultBehavior ?

const (
	PrepareSupportDefaultBehaviorIdentifier PrepareSupportDefaultBehavior = 1
)





type TokenFormat ?

const (
	TokenFormatRelative TokenFormat = "relative"
)




