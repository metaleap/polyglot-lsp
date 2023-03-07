// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-vsx-and-lsp/gen/cmd/gen_lsp
package lsp

type ImplementationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// Represents a location inside a resource, such as a line
// inside a text file.
type Location struct {
	Uri   DocumentURI
	Range Range
}
type ImplementationRegistrationOptions struct {
	TextDocumentRegistrationOptions
	ImplementationOptions
	StaticRegistrationOptions
}
type TypeDefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}
type TypeDefinitionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	TypeDefinitionOptions
	StaticRegistrationOptions
}

// A workspace folder inside a client.
type WorkspaceFolder struct {
	// The associated URI for this workspace folder.
	Uri URI
	// The name of the workspace folder. Used to refer to this
	// workspace folder in the user interface.
	Name string
}

// The parameters of a `workspace/didChangeWorkspaceFolders` notification.
type DidChangeWorkspaceFoldersParams struct {
	// The actual workspace folder change event.
	Event WorkspaceFoldersChangeEvent
}

// The parameters of a configuration request.
type ConfigurationParams struct {
	Items []ConfigurationItem
}

// Parameters for a `DocumentColorRequest`.
type DocumentColorParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The text document.
	TextDocument TextDocumentIdentifier
}

// Represents a color range from a document.
type ColorInformation struct {
	// The range in the document where this color appears.
	Range Range
	// The actual color value for this color range.
	Color Color
}
type DocumentColorRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentColorOptions
	StaticRegistrationOptions
}

// Parameters for a `ColorPresentationRequest`.
type ColorPresentationParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The text document.
	TextDocument TextDocumentIdentifier
	// The color to request presentations for.
	Color Color
	// The range where the color would be inserted. Serves as a context.
	Range Range
}
type ColorPresentation struct {
	// The label of this color presentation. It will be shown on the color
	// picker header. By default this is also the text that is inserted when selecting
	// this color presentation.
	Label string
	// An `TextEdit` which is applied to a document when selecting
	// this presentation for the color.  When `falsy` the `ColorPresentation.label`
	// is used.
	TextEdit/*TOpt*/ *TextEdit
	// An optional array of additional `TextEdit` that are applied when
	// selecting this color presentation. Edits must not overlap with the main `ColorPresentation.textEdit` nor with themselves.
	AdditionalTextEdits/*TOpt*/ []TextEdit
}
type WorkDoneProgressOptions struct {
	WorkDoneProgress /*TOpt*/ *Boolean
}

// General text document registration options.
type TextDocumentRegistrationOptions struct {
	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector /*TOr*/ /*TOpt*/ DocumentSelector
}

// Parameters for a `FoldingRangeRequest`.
type FoldingRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The text document.
	TextDocument TextDocumentIdentifier
}

// Represents a folding range. To be valid, start and end line must be bigger than zero and smaller
// than the number of lines in the document. Clients are free to ignore invalid ranges.
type FoldingRange struct {
	// The zero-based start line of the range to fold. The folded area starts after the line's last character.
	// To be valid, the end must be zero or larger and smaller than the number of lines in the document.
	StartLine uint
	// The zero-based character offset from where the folded range starts. If not defined, defaults to the length of the start line.
	StartCharacter/*TOpt*/ *Uinteger
	// The zero-based end line of the range to fold. The folded area ends with the line's last character.
	// To be valid, the end must be zero or larger and smaller than the number of lines in the document.
	EndLine uint
	// The zero-based character offset before the folded range ends. If not defined, defaults to the length of the end line.
	EndCharacter/*TOpt*/ *Uinteger
	// Describes the kind of the folding range such as `comment' or 'region'. The kind
	// is used to categorize folding ranges and used by commands like 'Fold all comments'.
	// See `FoldingRangeKind` for an enumeration of standardized kinds.
	Kind/*TOpt*/ FoldingRangeKind
	// The text that the client should show when the specified range is
	// collapsed. If not defined or not supported by the client, a default
	// will be chosen by the client.
	//
	// @since 3.17.0
	CollapsedText/*TOpt*/ *String
}
type FoldingRangeRegistrationOptions struct {
	TextDocumentRegistrationOptions
	FoldingRangeOptions
	StaticRegistrationOptions
}
type DeclarationParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}
type DeclarationRegistrationOptions struct {
	DeclarationOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// A parameter literal used in selection range requests.
type SelectionRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The text document.
	TextDocument TextDocumentIdentifier
	// The positions inside the text document.
	Positions []Position
}

// A selection range represents a part of a selection hierarchy. A selection range
// may have a parent selection range that contains it.
type SelectionRange struct {
	// The `Range` of this selection range.
	Range Range
	// The parent selection range containing this range. Therefore `parent.range` must contain `this.range`.
	Parent/*TOpt*/ *SelectionRange
}
type SelectionRangeRegistrationOptions struct {
	SelectionRangeOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}
type WorkDoneProgressCreateParams struct {
	// The token to be used to report progress.
	Token ProgressToken
}
type WorkDoneProgressCancelParams struct {
	// The token to be used to report progress.
	Token ProgressToken
}

// The parameter of a `textDocument/prepareCallHierarchy` request.
//
// @since 3.16.0
type CallHierarchyPrepareParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// Represents programming constructs like functions or constructors in the context
// of call hierarchy.
//
// @since 3.16.0
type CallHierarchyItem struct {
	// The name of this item.
	Name string
	// The kind of this item.
	Kind SymbolKind
	// Tags for this item.
	Tags/*TOpt*/ []SymbolTag
	// More detail for this item, e.g. the signature of a function.
	Detail/*TOpt*/ *String
	// The resource identifier of this item.
	Uri DocumentURI
	// The range enclosing this symbol not including leading/trailing whitespace but everything else, e.g. comments and code.
	Range Range
	// The range that should be selected and revealed when this symbol is being picked, e.g. the name of a function.
	// Must be contained by the `CallHierarchyItem.range`.
	SelectionRange Range
	// A data entry field that is preserved between a call hierarchy prepare and
	// incoming calls or outgoing calls requests.
	Data/*TOpt*/ LSPAny
}

// Call hierarchy options used during static or dynamic registration.
//
// @since 3.16.0
type CallHierarchyRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CallHierarchyOptions
	StaticRegistrationOptions
}

// The parameter of a `callHierarchy/incomingCalls` request.
//
// @since 3.16.0
type CallHierarchyIncomingCallsParams struct {
	WorkDoneProgressParams
	PartialResultParams
	Item CallHierarchyItem
}

// Represents an incoming call, e.g. a caller of a method or constructor.
//
// @since 3.16.0
type CallHierarchyIncomingCall struct {
	// The item that makes the call.
	From CallHierarchyItem
	// The ranges at which the calls appear. This is relative to the caller
	// denoted by `CallHierarchyIncomingCall.from`.
	FromRanges []Range
}

// The parameter of a `callHierarchy/outgoingCalls` request.
//
// @since 3.16.0
type CallHierarchyOutgoingCallsParams struct {
	WorkDoneProgressParams
	PartialResultParams
	Item CallHierarchyItem
}

// Represents an outgoing call, e.g. calling a getter from a method or a method from a constructor etc.
//
// @since 3.16.0
type CallHierarchyOutgoingCall struct {
	// The item that is called.
	To CallHierarchyItem
	// The range at which this item is called. This is the range relative to the caller, e.g the item
	// passed to `CallHierarchyItemProvider.provideCallHierarchyOutgoingCalls`
	// and not `CallHierarchyOutgoingCall.to`.
	FromRanges []Range
}

// @since 3.16.0
type SemanticTokensParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The text document.
	TextDocument TextDocumentIdentifier
}

// @since 3.16.0
type SemanticTokens struct {
	// An optional result id. If provided and clients support delta updating
	// the client will include the result id in the next semantic token request.
	// A server can then instead of computing all semantic tokens again simply
	// send a delta.
	ResultId/*TOpt*/ *String
	// The actual tokens.
	Data []uint
}

// @since 3.16.0
type SemanticTokensPartialResult struct {
	Data []uint
}

// @since 3.16.0
type SemanticTokensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SemanticTokensOptions
	StaticRegistrationOptions
}

// @since 3.16.0
type SemanticTokensDeltaParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The text document.
	TextDocument TextDocumentIdentifier
	// The result id of a previous response. The result Id can either point to a full response
	// or a delta response depending on what was received last.
	PreviousResultId string
}

// @since 3.16.0
type SemanticTokensDelta struct {
	ResultId/*TOpt*/ *String
	// The semantic token edits to transform a previous result into a new result.
	Edits []SemanticTokensEdit
}

// @since 3.16.0
type SemanticTokensDeltaPartialResult struct {
	Edits []SemanticTokensEdit
}

// @since 3.16.0
type SemanticTokensRangeParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The text document.
	TextDocument TextDocumentIdentifier
	// The range the semantic tokens are requested for.
	Range Range
}

// Params to show a document.
//
// @since 3.16.0
type ShowDocumentParams struct {
	// The document uri to show.
	Uri URI
	// Indicates to show the resource in an external program.
	// To show for example `https://code.visualstudio.com/`
	// in the default WEB browser set `external` to `true`.
	External/*TOpt*/ *Boolean
	// An optional property to indicate whether the editor
	// showing the document should take focus or not.
	// Clients might ignore this property if an external
	// program is started.
	TakeFocus/*TOpt*/ *Boolean
	// An optional selection range if the document is a text
	// document. Clients might ignore the property if an
	// external program is started or the file is not a text
	// file.
	Selection/*TOpt*/ *Range
}

// The result of a showDocument request.
//
// @since 3.16.0
type ShowDocumentResult struct {
	// A boolean indicating if the show was successful.
	Success bool
}
type LinkedEditingRangeParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// The result of a linked editing range request.
//
// @since 3.16.0
type LinkedEditingRanges struct {
	// A list of ranges that can be edited together. The ranges must have
	// identical length and contain identical text content. The ranges cannot overlap.
	Ranges []Range
	// An optional word pattern (regular expression) that describes valid contents for
	// the given ranges. If no pattern is provided, the client configuration's word
	// pattern will be used.
	WordPattern/*TOpt*/ *String
}
type LinkedEditingRangeRegistrationOptions struct {
	TextDocumentRegistrationOptions
	LinkedEditingRangeOptions
	StaticRegistrationOptions
}

// The parameters sent in notifications/requests for user-initiated creation of
// files.
//
// @since 3.16.0
type CreateFilesParams struct {
	// An array of all files/folders created in this operation.
	Files []FileCreate
}

// A workspace edit represents changes to many resources managed in the workspace. The edit
// should either provide `changes` or `documentChanges`. If documentChanges are present
// they are preferred over `changes` if the client can handle versioned document edits.
//
// Since version 3.13.0 a workspace edit can contain resource operations as well. If resource
// operations are present clients need to execute the operations in the order in which they
// are provided. So a workspace edit for example can consist of the following two changes:
// (1) a create file a.txt and (2) a text document edit which insert text into file a.txt.
//
// An invalid sequence (e.g. (1) delete file a.txt and (2) insert text into file a.txt) will
// cause failure of the operation. How the client recovers from the failure is described by
// the client capability: `workspace.workspaceEdit.failureHandling`
type WorkspaceEdit struct {
	// Holds changes to existing resources.
	Changes/*TOpt*/ map[DocumentURI][]TextEdit
	// Depending on the client capability `workspace.workspaceEdit.resourceOperations` document changes
	// are either an array of `TextDocumentEdit`s to express changes to n different text documents
	// where each text document edit addresses a specific version of a text document. Or it can contain
	// above `TextDocumentEdit`s mixed with create, rename and delete file / folder operations.
	//
	// Whether a client supports versioned document edits is expressed via
	// `workspace.workspaceEdit.documentChanges` client capability.
	//
	// If a client neither supports `documentChanges` nor `workspace.workspaceEdit.resourceOperations` then
	// only plain `TextEdit`s using the `changes` property are supported.
	//
	// Every object in the array has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	DocumentChanges/*TOpt*/ [] /*TOr*/ struct {
		TextDocumentEdit/*TOpt*/ *TextDocumentEdit
		CreateFile/*TOpt*/ *CreateFile
		RenameFile/*TOpt*/ *RenameFile
		DeleteFile/*TOpt*/ *DeleteFile
	}

	// A map of change annotations that can be referenced in `AnnotatedTextEdit`s or create, rename and
	// delete file / folder operations.
	//
	// Whether clients honor this property depends on the client capability `workspace.changeAnnotationSupport`.
	//
	// @since 3.16.0
	ChangeAnnotations/*TOpt*/ map[ChangeAnnotationIdentifier]ChangeAnnotation
}

// The options to register for file operations.
//
// @since 3.16.0
type FileOperationRegistrationOptions struct {
	// The actual filters.
	Filters []FileOperationFilter
}

// The parameters sent in notifications/requests for user-initiated renames of
// files.
//
// @since 3.16.0
type RenameFilesParams struct {
	// An array of all files/folders renamed in this operation. When a folder is renamed, only
	// the folder will be included, and not its children.
	Files []FileRename
}

// The parameters sent in notifications/requests for user-initiated deletes of
// files.
//
// @since 3.16.0
type DeleteFilesParams struct {
	// An array of all files/folders deleted in this operation.
	Files []FileDelete
}
type MonikerParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// Moniker definition to match LSIF 0.5 moniker definition.
//
// @since 3.16.0
type Moniker struct {
	// The scheme of the moniker. For example tsc or .Net
	Scheme string
	// The identifier of the moniker. The value is opaque in LSIF however
	// schema owners are allowed to define the structure if they want.
	Identifier string
	// The scope in which the moniker is unique
	Unique UniquenessLevel
	// The moniker kind if known.
	Kind/*TOpt*/ MonikerKind
}
type MonikerRegistrationOptions struct {
	TextDocumentRegistrationOptions
	MonikerOptions
}

// The parameter of a `textDocument/prepareTypeHierarchy` request.
//
// @since 3.17.0
type TypeHierarchyPrepareParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// @since 3.17.0
type TypeHierarchyItem struct {
	// The name of this item.
	Name string
	// The kind of this item.
	Kind SymbolKind
	// Tags for this item.
	Tags/*TOpt*/ []SymbolTag
	// More detail for this item, e.g. the signature of a function.
	Detail/*TOpt*/ *String
	// The resource identifier of this item.
	Uri DocumentURI
	// The range enclosing this symbol not including leading/trailing whitespace
	// but everything else, e.g. comments and code.
	Range Range
	// The range that should be selected and revealed when this symbol is being
	// picked, e.g. the name of a function. Must be contained by the
	// `TypeHierarchyItem.range`.
	SelectionRange Range
	// A data entry field that is preserved between a type hierarchy prepare and
	// supertypes or subtypes requests. It could also be used to identify the
	// type hierarchy in the server, helping improve the performance on
	// resolving supertypes and subtypes.
	Data/*TOpt*/ LSPAny
}

// Type hierarchy options used during static or dynamic registration.
//
// @since 3.17.0
type TypeHierarchyRegistrationOptions struct {
	TextDocumentRegistrationOptions
	TypeHierarchyOptions
	StaticRegistrationOptions
}

// The parameter of a `typeHierarchy/supertypes` request.
//
// @since 3.17.0
type TypeHierarchySupertypesParams struct {
	WorkDoneProgressParams
	PartialResultParams
	Item TypeHierarchyItem
}

// The parameter of a `typeHierarchy/subtypes` request.
//
// @since 3.17.0
type TypeHierarchySubtypesParams struct {
	WorkDoneProgressParams
	PartialResultParams
	Item TypeHierarchyItem
}

// A parameter literal used in inline value requests.
//
// @since 3.17.0
type InlineValueParams struct {
	WorkDoneProgressParams
	// The text document.
	TextDocument TextDocumentIdentifier
	// The document range for which inline values should be computed.
	Range Range
	// Additional information about the context in which inline values were
	// requested.
	Context InlineValueContext
}

// Inline value options used during static or dynamic registration.
//
// @since 3.17.0
type InlineValueRegistrationOptions struct {
	InlineValueOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// A parameter literal used in inlay hint requests.
//
// @since 3.17.0
type InlayHintParams struct {
	WorkDoneProgressParams
	// The text document.
	TextDocument TextDocumentIdentifier
	// The document range for which inlay hints should be computed.
	Range Range
}

// Inlay hint information.
//
// @since 3.17.0
type InlayHint struct {
	// The position of this hint.
	Position Position
	// The label of this hint. A human readable string or an array of
	// InlayHintLabelPart label parts.
	//
	// *Note* that neither the string nor the label part can be empty.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Label/*TOr*/ struct {
		String/*TOpt*/ *String
		InlayHintLabelParts/*TOpt*/ []InlayHintLabelPart
	}

	// The kind of this hint. Can be omitted in which case the client
	// should fall back to a reasonable default.
	Kind/*TOpt*/ InlayHintKind
	// Optional text edits that are performed when accepting this inlay hint.
	//
	// *Note* that edits are expected to change the document so that the inlay
	// hint (or its nearest variant) is now part of the document and the inlay
	// hint itself is now obsolete.
	TextEdits/*TOpt*/ []TextEdit
	// The tooltip text when you hover over this item.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Tooltip/*TOpt*/ /*TOr*/ struct {
		String/*TOpt*/ *String
		MarkupContent/*TOpt*/ *MarkupContent
	}

	// Render padding before the hint.
	//
	// Note: Padding should use the editor's background color, not the
	// background color of the hint itself. That means padding can be used
	// to visually align/separate an inlay hint.
	PaddingLeft/*TOpt*/ *Boolean
	// Render padding after the hint.
	//
	// Note: Padding should use the editor's background color, not the
	// background color of the hint itself. That means padding can be used
	// to visually align/separate an inlay hint.
	PaddingRight/*TOpt*/ *Boolean
	// A data entry field that is preserved on an inlay hint between
	// a `textDocument/inlayHint` and a `inlayHint/resolve` request.
	Data/*TOpt*/ LSPAny
}

// Inlay hint options used during static or dynamic registration.
//
// @since 3.17.0
type InlayHintRegistrationOptions struct {
	InlayHintOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// Parameters of the document diagnostic request.
//
// @since 3.17.0
type DocumentDiagnosticParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The text document.
	TextDocument TextDocumentIdentifier
	// The additional identifier  provided during registration.
	Identifier/*TOpt*/ *String
	// The result id of a previous response if provided.
	PreviousResultId/*TOpt*/ *String
}

// A partial result for a document diagnostic report.
//
// @since 3.17.0
type DocumentDiagnosticReportPartialResult struct {
	// Every object in the map has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	RelatedDocuments map[DocumentURI] /*TOr*/ struct {
		FullDocumentDiagnosticReport/*TOpt*/ *FullDocumentDiagnosticReport
		UnchangedDocumentDiagnosticReport/*TOpt*/ *UnchangedDocumentDiagnosticReport
	}
}

// Cancellation data returned from a diagnostic request.
//
// @since 3.17.0
type DiagnosticServerCancellationData struct {
	RetriggerRequest bool
}

// Diagnostic registration options.
//
// @since 3.17.0
type DiagnosticRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DiagnosticOptions
	StaticRegistrationOptions
}

// Parameters of the workspace diagnostic request.
//
// @since 3.17.0
type WorkspaceDiagnosticParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The additional identifier provided during registration.
	Identifier/*TOpt*/ *String
	// The currently known diagnostic reports with their
	// previous result ids.
	PreviousResultIds []PreviousResultId
}

// A workspace diagnostic report.
//
// @since 3.17.0
type WorkspaceDiagnosticReport struct {
	Items []WorkspaceDocumentDiagnosticReport
}

// A partial result for a workspace diagnostic report.
//
// @since 3.17.0
type WorkspaceDiagnosticReportPartialResult struct {
	Items []WorkspaceDocumentDiagnosticReport
}

// The params sent in an open notebook document notification.
//
// @since 3.17.0
type DidOpenNotebookDocumentParams struct {
	// The notebook document that got opened.
	NotebookDocument NotebookDocument
	// The text documents that represent the content
	// of a notebook cell.
	CellTextDocuments []TextDocumentItem
}

// The params sent in a change notebook document notification.
//
// @since 3.17.0
type DidChangeNotebookDocumentParams struct {
	// The notebook document that did change. The version number points
	// to the version after all provided changes have been applied. If
	// only the text document content of a cell changes the notebook version
	// doesn't necessarily have to change.
	NotebookDocument VersionedNotebookDocumentIdentifier
	// The actual changes to the notebook document.
	//
	// The changes describe single state changes to the notebook document.
	// So if there are two changes c1 (at array index 0) and c2 (at array
	// index 1) for a notebook in state S then c1 moves the notebook from
	// S to S' and c2 from S' to S''. So c1 is computed on the state S and
	// c2 is computed on the state S'.
	//
	// To mirror the content of a notebook using change events use the following approach:
	// - start with the same initial content
	// - apply the 'notebookDocument/didChange' notifications in the order you receive them.
	// - apply the `NotebookChangeEvent`s in a single notification in the order
	//   you receive them.
	Change NotebookDocumentChangeEvent
}

// The params sent in a save notebook document notification.
//
// @since 3.17.0
type DidSaveNotebookDocumentParams struct {
	// The notebook document that got saved.
	NotebookDocument NotebookDocumentIdentifier
}

// The params sent in a close notebook document notification.
//
// @since 3.17.0
type DidCloseNotebookDocumentParams struct {
	// The notebook document that got closed.
	NotebookDocument NotebookDocumentIdentifier
	// The text documents that represent the content
	// of a notebook cell that got closed.
	CellTextDocuments []TextDocumentIdentifier
}
type RegistrationParams struct {
	Registrations []Registration
}
type UnregistrationParams struct {
	Unregisterations []Unregistration
}
type InitializeParams struct {
	_InitializeParams
	WorkspaceFoldersInitializeParams
}

// The result returned from an initialize request.
type InitializeResult struct {
	// The capabilities the language server provides.
	Capabilities ServerCapabilities
	// Information about the server.
	//
	// @since 3.15.0
	ServerInfo/*TOpt*/ * /*TStruc*/ struct {
		// The name of the server as defined by the server.
		Name string
		// The server's version as defined by the server.
		Version/*TOpt*/ *String
	}
}

// The data type of the ResponseError if the
// initialize request fails.
type InitializeError struct {
	// Indicates whether the client execute the following retry logic:
	// (1) show the message provided by the ResponseError to the user
	// (2) user selects retry or cancel
	// (3) if user selected retry the initialize method is sent again.
	Retry bool
}
type InitializedParams struct {
}

// The parameters of a change configuration notification.
type DidChangeConfigurationParams struct {
	// The actual changed settings
	Settings LSPAny
}
type DidChangeConfigurationRegistrationOptions struct {
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Section /*TOpt*/ /*TOr*/ struct {
		String/*TOpt*/ *String
		Strings/*TOpt*/ []string
	}
}

// The parameters of a notification message.
type ShowMessageParams struct {
	// The message type. See `MessageType`
	Type MessageType
	// The actual message.
	Message string
}
type ShowMessageRequestParams struct {
	// The message type. See `MessageType`
	Type MessageType
	// The actual message.
	Message string
	// The message action items to present.
	Actions/*TOpt*/ []MessageActionItem
}
type MessageActionItem struct {
	// A short title like 'Retry', 'Open Log' etc.
	Title string
}

// The log message parameters.
type LogMessageParams struct {
	// The message type. See `MessageType`
	Type MessageType
	// The actual message.
	Message string
}

// The parameters sent in an open text document notification
type DidOpenTextDocumentParams struct {
	// The document that was opened.
	TextDocument TextDocumentItem
}

// The change text document notification's parameters.
type DidChangeTextDocumentParams struct {
	// The document that did change. The version number points
	// to the version after all provided content changes have
	// been applied.
	TextDocument VersionedTextDocumentIdentifier
	// The actual content changes. The content changes describe single state changes
	// to the document. So if there are two content changes c1 (at array index 0) and
	// c2 (at array index 1) for a document in state S then c1 moves the document from
	// S to S' and c2 from S' to S''. So c1 is computed on the state S and c2 is computed
	// on the state S'.
	//
	// To mirror the content of a document using change events use the following approach:
	// - start with the same initial content
	// - apply the 'textDocument/didChange' notifications in the order you receive them.
	// - apply the `TextDocumentContentChangeEvent`s in a single notification in the order
	//   you receive them.
	ContentChanges []TextDocumentContentChangeEvent
}

// Describe options to be used when registered for text document change events.
type TextDocumentChangeRegistrationOptions struct {
	TextDocumentRegistrationOptions
	// How documents are synced to the server.
	SyncKind TextDocumentSyncKind
}

// The parameters sent in a close text document notification
type DidCloseTextDocumentParams struct {
	// The document that was closed.
	TextDocument TextDocumentIdentifier
}

// The parameters sent in a save text document notification
type DidSaveTextDocumentParams struct {
	// The document that was saved.
	TextDocument TextDocumentIdentifier
	// Optional the content when saved. Depends on the includeText value
	// when the save notification was requested.
	Text/*TOpt*/ *String
}

// Save registration options.
type TextDocumentSaveRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SaveOptions
}

// The parameters sent in a will save text document notification.
type WillSaveTextDocumentParams struct {
	// The document that will be saved.
	TextDocument TextDocumentIdentifier
	// The 'TextDocumentSaveReason'.
	Reason TextDocumentSaveReason
}

// A text edit applicable to a text document.
type TextEdit struct {
	// The range of the text document to be manipulated. To insert
	// text into a document create a range where start === end.
	Range Range
	// The string to be inserted. For delete operations use an
	// empty string.
	NewText string
}

// The watched files change notification's parameters.
type DidChangeWatchedFilesParams struct {
	// The actual file events.
	Changes []FileEvent
}

// Describe options to be used when registered for text document change events.
type DidChangeWatchedFilesRegistrationOptions struct {
	// The watchers to register.
	Watchers []FileSystemWatcher
}

// The publish diagnostic notification's parameters.
type PublishDiagnosticsParams struct {
	// The URI for which diagnostic information is reported.
	Uri DocumentURI
	// Optional the version number of the document the diagnostics are published for.
	//
	// @since 3.15.0
	Version/*TOpt*/ *Integer
	// An array of diagnostic information items.
	Diagnostics []Diagnostic
}

// Completion parameters
type CompletionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
	// The completion context. This is only available it the client specifies
	// to send this using the client capability `textDocument.completion.contextSupport === true`
	Context/*TOpt*/ *CompletionContext
}

// A completion item represents a text snippet that is
// proposed to complete text that is being typed.
type CompletionItem struct {
	// The label of this completion item.
	//
	// The label property is also by default the text that
	// is inserted when selecting this completion.
	//
	// If label details are provided the label itself should
	// be an unqualified name of the completion item.
	Label string
	// Additional details for the label
	//
	// @since 3.17.0
	LabelDetails/*TOpt*/ *CompletionItemLabelDetails
	// The kind of this completion item. Based of the kind
	// an icon is chosen by the editor.
	Kind/*TOpt*/ CompletionItemKind
	// Tags for this completion item.
	//
	// @since 3.15.0
	Tags/*TOpt*/ []CompletionItemTag
	// A human-readable string with additional information
	// about this item, like type or symbol information.
	Detail/*TOpt*/ *String
	// A human-readable string that represents a doc-comment.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Documentation/*TOpt*/ /*TOr*/ struct {
		String/*TOpt*/ *String
		MarkupContent/*TOpt*/ *MarkupContent
	}

	// Indicates if this item is deprecated.
	// @deprecated Use `tags` instead.
	Deprecated/*TOpt*/ *Boolean
	// Select this item when showing.
	//
	// *Note* that only one completion item can be selected and that the
	// tool / client decides which item that is. The rule is that the *first*
	// item of those that match best is selected.
	Preselect/*TOpt*/ *Boolean
	// A string that should be used when comparing this item
	// with other items. When `falsy` the `CompletionItem.label`
	// is used.
	SortText/*TOpt*/ *String
	// A string that should be used when filtering a set of
	// completion items. When `falsy` the `CompletionItem.label`
	// is used.
	FilterText/*TOpt*/ *String
	// A string that should be inserted into a document when selecting
	// this completion. When `falsy` the `CompletionItem.label`
	// is used.
	//
	// The `insertText` is subject to interpretation by the client side.
	// Some tools might not take the string literally. For example
	// VS Code when code complete is requested in this example
	// `con<cursor position>` and a completion item with an `insertText` of
	// `console` is provided it will only insert `sole`. Therefore it is
	// recommended to use `textEdit` instead since it avoids additional client
	// side interpretation.
	InsertText/*TOpt*/ *String
	// The format of the insert text. The format applies to both the
	// `insertText` property and the `newText` property of a provided
	// `textEdit`. If omitted defaults to `InsertTextFormat.PlainText`.
	//
	// Please note that the insertTextFormat doesn't apply to
	// `additionalTextEdits`.
	InsertTextFormat/*TOpt*/ InsertTextFormat
	// How whitespace and indentation is handled during completion
	// item insertion. If not provided the clients default value depends on
	// the `textDocument.completion.insertTextMode` client capability.
	//
	// @since 3.16.0
	InsertTextMode/*TOpt*/ InsertTextMode
	// An `TextEdit` which is applied to a document when selecting
	// this completion. When an edit is provided the value of
	// `CompletionItem.insertText` is ignored.
	//
	// Most editors support two different operations when accepting a completion
	// item. One is to insert a completion text and the other is to replace an
	// existing text with a completion text. Since this can usually not be
	// predetermined by a server it can report both ranges. Clients need to
	// signal support for `InsertReplaceEdits` via the
	// `textDocument.completion.insertReplaceSupport` client capability
	// property.
	//
	// *Note 1:* The text edit's range as well as both ranges from an insert
	// replace edit must be a [single line] and they must contain the position
	// at which completion has been requested.
	// *Note 2:* If an `InsertReplaceEdit` is returned the edit's insert range
	// must be a prefix of the edit's replace range, that means it must be
	// contained and starting at the same position.
	//
	// @since 3.16.0 additional type `InsertReplaceEdit`
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	TextEdit/*TOpt*/ /*TOr*/ struct {
		TextEdit/*TOpt*/ *TextEdit
		InsertReplaceEdit/*TOpt*/ *InsertReplaceEdit
	}

	// The edit text used if the completion item is part of a CompletionList and
	// CompletionList defines an item default for the text edit range.
	//
	// Clients will only honor this property if they opt into completion list
	// item defaults using the capability `completionList.itemDefaults`.
	//
	// If not provided and a list's default range is provided the label
	// property is used as a text.
	//
	// @since 3.17.0
	TextEditText/*TOpt*/ *String
	// An optional array of additional `TextEdit` that are applied when
	// selecting this completion. Edits must not overlap (including the same insert position)
	// with the main `CompletionItem.textEdit` nor with themselves.
	//
	// Additional text edits should be used to change text unrelated to the current cursor position
	// (for example adding an import statement at the top of the file if the completion item will
	// insert an unqualified type).
	AdditionalTextEdits/*TOpt*/ []TextEdit
	// An optional set of characters that when pressed while this completion is active will accept it first and
	// then type that character. *Note* that all commit characters should have `length=1` and that superfluous
	// characters will be ignored.
	CommitCharacters/*TOpt*/ []string
	// An optional `Command` that is executed *after* inserting this completion. *Note* that
	// additional modifications to the current document should be described with the
	// `CompletionItem.additionalTextEdits`-property.
	Command/*TOpt*/ *Command
	// A data entry field that is preserved on a completion item between a
	// `CompletionRequest` and a `CompletionResolveRequest`.
	Data/*TOpt*/ LSPAny
}

// Represents a collection of `CompletionItem` to be presented
// in the editor.
type CompletionList struct {
	// This list it not complete. Further typing results in recomputing this list.
	//
	// Recomputed lists have all their items replaced (not appended) in the
	// incomplete completion sessions.
	IsIncomplete bool
	// In many cases the items of an actual completion result share the same
	// value for properties like `commitCharacters` or the range of a text
	// edit. A completion list can therefore define item defaults which will
	// be used if a completion item itself doesn't specify the value.
	//
	// If a completion list specifies a default value and a completion item
	// also specifies a corresponding value the one from the item is used.
	//
	// Servers are only allowed to return default values if the client
	// signals support for this via the `completionList.itemDefaults`
	// capability.
	//
	// @since 3.17.0
	ItemDefaults/*TOpt*/ * /*TStruc*/ struct {
		// A default commit character set.
		//
		// @since 3.17.0
		CommitCharacters/*TOpt*/ []string
		// A default edit range.
		//
		// @since 3.17.0
		//
		// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
		EditRange/*TOpt*/ /*TOr*/ struct {
			Range/*TOpt*/ *Range
			InsertReplace/*TOpt*/ * /*TStruc*/ struct {
				Insert  Range
				Replace Range
			}
		}

		// A default insert text format.
		//
		// @since 3.17.0
		InsertTextFormat/*TOpt*/ InsertTextFormat
		// A default insert text mode.
		//
		// @since 3.17.0
		InsertTextMode/*TOpt*/ InsertTextMode
		// A default data value.
		//
		// @since 3.17.0
		Data/*TOpt*/ LSPAny
	}

	// The completion items.
	Items []CompletionItem
}

// Registration options for a `CompletionRequest`.
type CompletionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CompletionOptions
}

// Parameters for a `HoverRequest`.
type HoverParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// The result of a hover request.
type Hover struct {
	// The hover's content
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Contents/*TOr*/ struct {
		MarkupContent/*TOpt*/ *MarkupContent
		MarkedString/*TOpt*/ MarkedString
		MarkedStrings/*TOpt*/ []MarkedString
	}

	// An optional range inside the text document that is used to
	// visualize the hover, e.g. by changing the background color.
	Range/*TOpt*/ *Range
}

// Registration options for a `HoverRequest`.
type HoverRegistrationOptions struct {
	TextDocumentRegistrationOptions
	HoverOptions
}

// Parameters for a `SignatureHelpRequest`.
type SignatureHelpParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	// The signature help context. This is only available if the client specifies
	// to send this using the client capability `textDocument.signatureHelp.contextSupport === true`
	//
	// @since 3.15.0
	Context/*TOpt*/ *SignatureHelpContext
}

// Signature help represents the signature of something
// callable. There can be multiple signature but only one
// active and only one active parameter.
type SignatureHelp struct {
	// One or more signatures.
	Signatures []SignatureInformation
	// The active signature. If omitted or the value lies outside the
	// range of `signatures` the value defaults to zero or is ignored if
	// the `SignatureHelp` has no signatures.
	//
	// Whenever possible implementors should make an active decision about
	// the active signature and shouldn't rely on a default value.
	//
	// In future version of the protocol this property might become
	// mandatory to better express this.
	ActiveSignature/*TOpt*/ *Uinteger
	// The active parameter of the active signature. If omitted or the value
	// lies outside the range of `signatures[activeSignature].parameters`
	// defaults to 0 if the active signature has parameters. If
	// the active signature has no parameters it is ignored.
	// In future version of the protocol this property might become
	// mandatory to better express the active parameter if the
	// active signature does have any.
	ActiveParameter/*TOpt*/ *Uinteger
}

// Registration options for a `SignatureHelpRequest`.
type SignatureHelpRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SignatureHelpOptions
}

// Parameters for a `DefinitionRequest`.
type DefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// Registration options for a `DefinitionRequest`.
type DefinitionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DefinitionOptions
}

// Parameters for a `ReferencesRequest`.
type ReferenceParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
	Context ReferenceContext
}

// Registration options for a `ReferencesRequest`.
type ReferenceRegistrationOptions struct {
	TextDocumentRegistrationOptions
	ReferenceOptions
}

// Parameters for a `DocumentHighlightRequest`.
type DocumentHighlightParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

// A document highlight is a range inside a text document which deserves
// special attention. Usually a document highlight is visualized by changing
// the background color of its range.
type DocumentHighlight struct {
	// The range this highlight applies to.
	Range Range
	// The highlight kind, default is `DocumentHighlightKind.Text`.
	Kind/*TOpt*/ DocumentHighlightKind
}

// Registration options for a `DocumentHighlightRequest`.
type DocumentHighlightRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentHighlightOptions
}

// Parameters for a `DocumentSymbolRequest`.
type DocumentSymbolParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The text document.
	TextDocument TextDocumentIdentifier
}

// Represents information about programming constructs like variables, classes,
// interfaces etc.
type SymbolInformation struct {
	BaseSymbolInformation
	// Indicates if this symbol is deprecated.
	//
	// @deprecated Use tags instead
	Deprecated/*TOpt*/ *Boolean
	// The location of this symbol. The location's range is used by a tool
	// to reveal the location in the editor. If the symbol is selected in the
	// tool the range's start information is used to position the cursor. So
	// the range usually spans more than the actual symbol's name and does
	// normally include things like visibility modifiers.
	//
	// The range doesn't have to denote a node range in the sense of an abstract
	// syntax tree. It can therefore not be used to re-construct a hierarchy of
	// the symbols.
	Location Location
}

// Represents programming constructs like variables, classes, interfaces etc.
// that appear in a document. Document symbols can be hierarchical and they
// have two ranges: one that encloses its definition and one that points to
// its most interesting range, e.g. the range of an identifier.
type DocumentSymbol struct {
	// The name of this symbol. Will be displayed in the user interface and therefore must not be
	// an empty string or a string only consisting of white spaces.
	Name string
	// More detail for this symbol, e.g the signature of a function.
	Detail/*TOpt*/ *String
	// The kind of this symbol.
	Kind SymbolKind
	// Tags for this document symbol.
	//
	// @since 3.16.0
	Tags/*TOpt*/ []SymbolTag
	// Indicates if this symbol is deprecated.
	//
	// @deprecated Use tags instead
	Deprecated/*TOpt*/ *Boolean
	// The range enclosing this symbol not including leading/trailing whitespace but everything else
	// like comments. This information is typically used to determine if the clients cursor is
	// inside the symbol to reveal in the symbol in the UI.
	Range Range
	// The range that should be selected and revealed when this symbol is being picked, e.g the name of a function.
	// Must be contained by the `range`.
	SelectionRange Range
	// Children of this symbol, e.g. properties of a class.
	Children/*TOpt*/ []DocumentSymbol
}

// Registration options for a `DocumentSymbolRequest`.
type DocumentSymbolRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentSymbolOptions
}

// The parameters of a `CodeActionRequest`.
type CodeActionParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The document in which the command was invoked.
	TextDocument TextDocumentIdentifier
	// The range for which the command was invoked.
	Range Range
	// Context carrying additional information.
	Context CodeActionContext
}

// Represents a reference to a command. Provides a title which
// will be used to represent a command in the UI and, optionally,
// an array of arguments which will be passed to the command handler
// function when invoked.
type Command struct {
	// Title of the command, like `save`.
	Title string
	// The identifier of the actual command handler.
	Command string
	// Arguments that the command handler should be
	// invoked with.
	Arguments/*TOpt*/ []LSPAny
}

// A code action represents a change that can be performed in code, e.g. to fix a problem or
// to refactor code.
//
// A CodeAction must set either `edit` and/or a `command`. If both are supplied, the `edit` is applied first, then the `command` is executed.
type CodeAction struct {
	// A short, human-readable, title for this code action.
	Title string
	// The kind of the code action.
	//
	// Used to filter code actions.
	Kind/*TOpt*/ CodeActionKind
	// The diagnostics that this code action resolves.
	Diagnostics/*TOpt*/ []Diagnostic
	// Marks this as a preferred action. Preferred actions are used by the `auto fix` command and can be targeted
	// by keybindings.
	//
	// A quick fix should be marked preferred if it properly addresses the underlying error.
	// A refactoring should be marked preferred if it is the most reasonable choice of actions to take.
	//
	// @since 3.15.0
	IsPreferred/*TOpt*/ *Boolean
	// Marks that the code action cannot currently be applied.
	//
	// Clients should follow the following guidelines regarding disabled code actions:
	//
	//   - Disabled code actions are not shown in automatic [lightbulbs](https://code.visualstudio.com/docs/editor/editingevolved#_code-action)
	//     code action menus.
	//
	//   - Disabled actions are shown as faded out in the code action menu when the user requests a more specific type
	//     of code action, such as refactorings.
	//
	//   - If the user has a [keybinding](https://code.visualstudio.com/docs/editor/refactoring#_keybindings-for-code-actions)
	//     that auto applies a code action and only disabled code actions are returned, the client should show the user an
	//     error message with `reason` in the editor.
	//
	// @since 3.16.0
	Disabled/*TOpt*/ * /*TStruc*/ struct {
		// Human readable description of why the code action is currently disabled.
		//
		// This is displayed in the code actions UI.
		Reason string
	}

	// The workspace edit this code action performs.
	Edit/*TOpt*/ *WorkspaceEdit
	// A command this code action executes. If a code action
	// provides an edit and a command, first the edit is
	// executed and then the command.
	Command/*TOpt*/ *Command
	// A data entry field that is preserved on a code action between
	// a `textDocument/codeAction` and a `codeAction/resolve` request.
	//
	// @since 3.16.0
	Data/*TOpt*/ LSPAny
}

// Registration options for a `CodeActionRequest`.
type CodeActionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeActionOptions
}

// The parameters of a `WorkspaceSymbolRequest`.
type WorkspaceSymbolParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// A query string to filter symbols by. Clients may send an empty
	// string here to request all symbols.
	Query string
}

// A special workspace symbol that supports locations without a range.
//
// See also SymbolInformation.
//
// @since 3.17.0
type WorkspaceSymbol struct {
	BaseSymbolInformation
	// The location of the symbol. Whether a server is allowed to
	// return a location without a range depends on the client
	// capability `workspace.symbol.resolveSupport`.
	//
	// See SymbolInformation#location for more details.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Location/*TOr*/ struct {
		Location/*TOpt*/ *Location
		Uri/*TOpt*/ * /*TStruc*/ struct {
			Uri DocumentURI
		}
	}

	// A data entry field that is preserved on a workspace symbol between a
	// workspace symbol request and a workspace symbol resolve request.
	Data/*TOpt*/ LSPAny
}

// Registration options for a `WorkspaceSymbolRequest`.
type WorkspaceSymbolRegistrationOptions struct {
	WorkspaceSymbolOptions
}

// The parameters of a `CodeLensRequest`.
type CodeLensParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The document to request code lens for.
	TextDocument TextDocumentIdentifier
}

// A code lens represents a `Command` that should be shown along with
// source text, like the number of references, a way to run tests, etc.
//
// A code lens is _unresolved_ when no command is associated to it. For performance
// reasons the creation of a code lens and resolving should be done in two stages.
type CodeLens struct {
	// The range in which this code lens is valid. Should only span a single line.
	Range Range
	// The command this code lens represents.
	Command/*TOpt*/ *Command
	// A data entry field that is preserved on a code lens item between
	// a `CodeLensRequest` and a [CodeLensResolveRequest]
	// (#CodeLensResolveRequest)
	Data/*TOpt*/ LSPAny
}

// Registration options for a `CodeLensRequest`.
type CodeLensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeLensOptions
}

// The parameters of a `DocumentLinkRequest`.
type DocumentLinkParams struct {
	WorkDoneProgressParams
	PartialResultParams
	// The document to provide document links for.
	TextDocument TextDocumentIdentifier
}

// A document link is a range in a text document that links to an internal or external resource, like another
// text document or a web site.
type DocumentLink struct {
	// The range this link applies to.
	Range Range
	// The uri this link points to. If missing a resolve request is sent later.
	Target/*TOpt*/ *String
	// The tooltip text when you hover over this link.
	//
	// If a tooltip is provided, is will be displayed in a string that includes instructions on how to
	// trigger the link, such as `{0} (ctrl + click)`. The specific instructions vary depending on OS,
	// user settings, and localization.
	//
	// @since 3.15.0
	Tooltip/*TOpt*/ *String
	// A data entry field that is preserved on a document link between a
	// DocumentLinkRequest and a DocumentLinkResolveRequest.
	Data/*TOpt*/ LSPAny
}

// Registration options for a `DocumentLinkRequest`.
type DocumentLinkRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentLinkOptions
}

// The parameters of a `DocumentFormattingRequest`.
type DocumentFormattingParams struct {
	WorkDoneProgressParams
	// The document to format.
	TextDocument TextDocumentIdentifier
	// The format options.
	Options FormattingOptions
}

// Registration options for a `DocumentFormattingRequest`.
type DocumentFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentFormattingOptions
}

// The parameters of a `DocumentRangeFormattingRequest`.
type DocumentRangeFormattingParams struct {
	WorkDoneProgressParams
	// The document to format.
	TextDocument TextDocumentIdentifier
	// The range to format
	Range Range
	// The format options
	Options FormattingOptions
}

// Registration options for a `DocumentRangeFormattingRequest`.
type DocumentRangeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentRangeFormattingOptions
}

// The parameters of a `DocumentOnTypeFormattingRequest`.
type DocumentOnTypeFormattingParams struct {
	// The document to format.
	TextDocument TextDocumentIdentifier
	// The position around which the on type formatting should happen.
	// This is not necessarily the exact position where the character denoted
	// by the property `ch` got typed.
	Position Position
	// The character that has been typed that triggered the formatting
	// on type request. That is not necessarily the last character that
	// got inserted into the document since the client could auto insert
	// characters as well (e.g. like automatic brace completion).
	Ch string
	// The formatting options.
	Options FormattingOptions
}

// Registration options for a `DocumentOnTypeFormattingRequest`.
type DocumentOnTypeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentOnTypeFormattingOptions
}

// The parameters of a `RenameRequest`.
type RenameParams struct {
	WorkDoneProgressParams
	// The document to rename.
	TextDocument TextDocumentIdentifier
	// The position at which this request was sent.
	Position Position
	// The new name of the symbol. If the given name is not valid the
	// request must return a `ResponseError` with an
	// appropriate message set.
	NewName string
}

// Registration options for a `RenameRequest`.
type RenameRegistrationOptions struct {
	TextDocumentRegistrationOptions
	RenameOptions
}
type PrepareRenameParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

// The parameters of a `ExecuteCommandRequest`.
type ExecuteCommandParams struct {
	WorkDoneProgressParams
	// The identifier of the actual command handler.
	Command string
	// Arguments that the command should be invoked with.
	Arguments/*TOpt*/ []LSPAny
}

// Registration options for a `ExecuteCommandRequest`.
type ExecuteCommandRegistrationOptions struct {
	ExecuteCommandOptions
}

// The parameters passed via an apply workspace edit request.
type ApplyWorkspaceEditParams struct {
	// An optional label of the workspace edit. This label is
	// presented in the user interface for example on an undo
	// stack to undo the workspace edit.
	Label/*TOpt*/ *String
	// The edits to apply.
	Edit WorkspaceEdit
}

// The result returned from the apply workspace edit request.
//
// @since 3.17 renamed from ApplyWorkspaceEditResponse
type ApplyWorkspaceEditResult struct {
	// Indicates whether the edit was applied or not.
	Applied bool
	// An optional textual description for why the edit was not applied.
	// This may be used by the server for diagnostic logging or to provide
	// a suitable error for a request that triggered the edit.
	FailureReason/*TOpt*/ *String
	// Depending on the client's failure handling strategy `failedChange` might
	// contain the index of the change that failed. This property is only available
	// if the client signals a `failureHandlingStrategy` in its client capabilities.
	FailedChange/*TOpt*/ *Uinteger
}
type WorkDoneProgressBegin struct {
	// The value is always "begin".
	// The value is always "begin".
	Kind string
	// Mandatory title of the progress operation. Used to briefly inform about
	// the kind of operation being performed.
	//
	// Examples: "Indexing" or "Linking dependencies".
	Title string
	// Controls if a cancel button should show to allow the user to cancel the
	// long running operation. Clients that don't support cancellation are allowed
	// to ignore the setting.
	Cancellable/*TOpt*/ *Boolean
	// Optional, more detailed associated progress message. Contains
	// complementary information to the `title`.
	//
	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	// If unset, the previous progress message (if any) is still valid.
	Message/*TOpt*/ *String
	// Optional progress percentage to display (value 100 is considered 100%).
	// If not provided infinite progress is assumed and clients are allowed
	// to ignore the `percentage` value in subsequent in report notifications.
	//
	// The value should be steadily rising. Clients are free to ignore values
	// that are not following this rule. The value range is [0, 100].
	Percentage/*TOpt*/ *Uinteger
}
type WorkDoneProgressReport struct {
	// The value is always "report".
	// The value is always "report".
	Kind string
	// Controls enablement state of a cancel button.
	//
	// Clients that don't support cancellation or don't support controlling the button's
	// enablement state are allowed to ignore the property.
	Cancellable/*TOpt*/ *Boolean
	// Optional, more detailed associated progress message. Contains
	// complementary information to the `title`.
	//
	// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
	// If unset, the previous progress message (if any) is still valid.
	Message/*TOpt*/ *String
	// Optional progress percentage to display (value 100 is considered 100%).
	// If not provided infinite progress is assumed and clients are allowed
	// to ignore the `percentage` value in subsequent in report notifications.
	//
	// The value should be steadily rising. Clients are free to ignore values
	// that are not following this rule. The value range is [0, 100]
	Percentage/*TOpt*/ *Uinteger
}
type WorkDoneProgressEnd struct {
	// The value is always "end".
	// The value is always "end".
	Kind string
	// Optional, a final message indicating to for example indicate the outcome
	// of the operation.
	Message/*TOpt*/ *String
}
type SetTraceParams struct {
	Value TraceValues
}
type LogTraceParams struct {
	Message string
	Verbose/*TOpt*/ *String
}
type CancelParams struct {
	// The request id to cancel.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Id /*TOr*/ struct {
		Integer/*TOpt*/ *Integer
		String/*TOpt*/ *String
	}
}
type ProgressParams struct {
	// The progress token provided by the client or server.
	Token ProgressToken
	// The progress data.
	Value LSPAny
}

// A parameter literal used in requests to pass a text document and a position inside that
// document.
type TextDocumentPositionParams struct {
	// The text document.
	TextDocument TextDocumentIdentifier
	// The position inside the text document.
	Position Position
}
type WorkDoneProgressParams struct {
	// An optional token that a server can use to report work done progress.
	WorkDoneToken /*TOpt*/ ProgressToken
}
type PartialResultParams struct {
	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken /*TOpt*/ ProgressToken
}

// Represents the connection of two locations. Provides additional metadata over normal `Location`,
// including an origin range.
type LocationLink struct {
	// Span of the origin of this link.
	//
	// Used as the underlined span for mouse interaction. Defaults to the word range at
	// the definition position.
	OriginSelectionRange/*TOpt*/ *Range
	// The target resource identifier of this link.
	TargetUri DocumentURI
	// The full target range of this link. If the target for example is a symbol then target range is the
	// range enclosing this symbol not including leading/trailing whitespace but everything else
	// like comments. This information is typically used to highlight the range in the editor.
	TargetRange Range
	// The range that should be selected and revealed when this link is being followed, e.g the name of a function.
	// Must be contained by the `targetRange`. See also `DocumentSymbol#range`
	TargetSelectionRange Range
}

// A range in a text document expressed as (zero-based) start and end positions.
//
// If you want to specify a range that contains a line including the line ending
// character(s) then use an end position denoting the start of the next line.
// For example:
// ```ts
//
//	{
//	    start: { line: 5, character: 23 }
//	    end : { line 6, character : 0 }
//	}
//
// ```
type Range struct {
	// The range's start position.
	Start Position
	// The range's end position.
	End Position
}
type ImplementationOptions struct {
	WorkDoneProgressOptions
}

// Static registration options to be returned in the initialize
// request.
type StaticRegistrationOptions struct {
	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id /*TOpt*/ *String
}
type TypeDefinitionOptions struct {
	WorkDoneProgressOptions
}

// The workspace folder change event.
type WorkspaceFoldersChangeEvent struct {
	// The array of added workspace folders
	Added []WorkspaceFolder
	// The array of the removed workspace folders
	Removed []WorkspaceFolder
}
type ConfigurationItem struct {
	// The scope to get the configuration section for.
	ScopeUri/*TOpt*/ *String
	// The configuration section asked for.
	Section/*TOpt*/ *String
}

// A literal to identify a text document in the client.
type TextDocumentIdentifier struct {
	// The text document's uri.
	Uri DocumentURI
}

// Represents a color in RGBA space.
type Color struct {
	// The red component of this color in the range [0-1].
	Red float64
	// The green component of this color in the range [0-1].
	Green float64
	// The blue component of this color in the range [0-1].
	Blue float64
	// The alpha component of this color in the range [0-1].
	Alpha float64
}
type DocumentColorOptions struct {
	WorkDoneProgressOptions
}
type FoldingRangeOptions struct {
	WorkDoneProgressOptions
}
type DeclarationOptions struct {
	WorkDoneProgressOptions
}

// Position in a text document expressed as zero-based line and character
// offset. Prior to 3.17 the offsets were always based on a UTF-16 string
// representation. So a string of the form `a𐐀b` the character offset of the
// character `a` is 0, the character offset of `𐐀` is 1 and the character
// offset of b is 3 since `𐐀` is represented using two code units in UTF-16.
// Since 3.17 clients and servers can agree on a different string encoding
// representation (e.g. UTF-8). The client announces it's supported encoding
// via the client capability [`general.positionEncodings`](#clientCapabilities).
// The value is an array of position encodings the client supports, with
// decreasing preference (e.g. the encoding at index `0` is the most preferred
// one). To stay backwards compatible the only mandatory encoding is UTF-16
// represented via the string `utf-16`. The server can pick one of the
// encodings offered by the client and signals that encoding back to the
// client via the initialize result's property
// [`capabilities.positionEncoding`](#serverCapabilities). If the string value
// `utf-16` is missing from the client's capability `general.positionEncodings`
// servers can safely assume that the client supports UTF-16. If the server
// omits the position encoding in its initialize result the encoding defaults
// to the string value `utf-16`. Implementation considerations: since the
// conversion from one encoding into another requires the content of the
// file / line the conversion is best done where the file is read which is
// usually on the server side.
//
// Positions are line end character agnostic. So you can not specify a position
// that denotes `\r|\n` or `\n|` where `|` represents the character offset.
//
// @since 3.17.0 - support for negotiated position encoding.
type Position struct {
	// Line position in a document (zero-based).
	//
	// If a line number is greater than the number of lines in a document, it defaults back to the number of lines in the document.
	// If a line number is negative, it defaults to 0.
	Line uint
	// Character offset on a line in a document (zero-based).
	//
	// The meaning of this offset is determined by the negotiated
	// `PositionEncodingKind`.
	//
	// If the character value is greater than the line length it defaults back to the
	// line length.
	Character uint
}
type SelectionRangeOptions struct {
	WorkDoneProgressOptions
}

// Call hierarchy options used during static registration.
//
// @since 3.16.0
type CallHierarchyOptions struct {
	WorkDoneProgressOptions
}

// @since 3.16.0
type SemanticTokensOptions struct {
	WorkDoneProgressOptions
	// The legend used by the server
	Legend SemanticTokensLegend
	// Server supports providing semantic tokens for a specific range
	// of a document.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Range/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		AnyByString/*TOpt*/ map[string]any
	}

	// Server supports providing semantic tokens for a full document.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Full/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		Delta/*TOpt*/ * /*TStruc*/ struct {
			// The server supports deltas for full documents.
			Delta /*TOpt*/ *Boolean
		}
	}
}

// @since 3.16.0
type SemanticTokensEdit struct {
	// The start offset of the edit.
	Start uint
	// The count of elements to remove.
	DeleteCount uint
	// The elements to insert.
	Data/*TOpt*/ []uint
}
type LinkedEditingRangeOptions struct {
	WorkDoneProgressOptions
}

// Represents information on a file/folder create.
//
// @since 3.16.0
type FileCreate struct {
	// A file:// URI for the location of the file/folder being created.
	Uri string
}

// Describes textual changes on a text document. A TextDocumentEdit describes all changes
// on a document version Si and after they are applied move the document to version Si+1.
// So the creator of a TextDocumentEdit doesn't need to sort the array of edits or do any
// kind of ordering. However the edits must be non overlapping.
type TextDocumentEdit struct {
	// The text document to change.
	TextDocument OptionalVersionedTextDocumentIdentifier
	// The edits to be applied.
	//
	// @since 3.16.0 - support for AnnotatedTextEdit. This is guarded using a
	// client capability.
	//
	// Every object in the array has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Edits [] /*TOr*/ struct {
		TextEdit/*TOpt*/ *TextEdit
		AnnotatedTextEdit/*TOpt*/ *AnnotatedTextEdit
	}
}

// Create file operation.
type CreateFile struct {
	ResourceOperation
	// A create
	// The value is always "create".
	// The value is always "create".
	Kind string
	// The resource to create.
	Uri DocumentURI
	// Additional options
	Options/*TOpt*/ *CreateFileOptions
}

// Rename file operation
type RenameFile struct {
	ResourceOperation
	// A rename
	// The value is always "rename".
	// The value is always "rename".
	Kind string
	// The old (existing) location.
	OldUri DocumentURI
	// The new location.
	NewUri DocumentURI
	// Rename options.
	Options/*TOpt*/ *RenameFileOptions
}

// Delete file operation
type DeleteFile struct {
	ResourceOperation
	// A delete
	// The value is always "delete".
	// The value is always "delete".
	Kind string
	// The file to delete.
	Uri DocumentURI
	// Delete options.
	Options/*TOpt*/ *DeleteFileOptions
}

// Additional information that describes document changes.
//
// @since 3.16.0
type ChangeAnnotation struct {
	// A human-readable string describing the actual change. The string
	// is rendered prominent in the user interface.
	Label string
	// A flag which indicates that user confirmation is needed
	// before applying the change.
	NeedsConfirmation/*TOpt*/ *Boolean
	// A human-readable string which is rendered less prominent in
	// the user interface.
	Description/*TOpt*/ *String
}

// A filter to describe in which file operation requests or notifications
// the server is interested in receiving.
//
// @since 3.16.0
type FileOperationFilter struct {
	// A Uri scheme like `file` or `untitled`.
	Scheme/*TOpt*/ *String
	// The actual file operation pattern.
	Pattern FileOperationPattern
}

// Represents information on a file/folder rename.
//
// @since 3.16.0
type FileRename struct {
	// A file:// URI for the original location of the file/folder being renamed.
	OldUri string
	// A file:// URI for the new location of the file/folder being renamed.
	NewUri string
}

// Represents information on a file/folder delete.
//
// @since 3.16.0
type FileDelete struct {
	// A file:// URI for the location of the file/folder being deleted.
	Uri string
}
type MonikerOptions struct {
	WorkDoneProgressOptions
}

// Type hierarchy options used during static registration.
//
// @since 3.17.0
type TypeHierarchyOptions struct {
	WorkDoneProgressOptions
}

// @since 3.17.0
type InlineValueContext struct {
	// The stack frame (as a DAP Id) where the execution has stopped.
	FrameId int
	// The document range where execution has stopped.
	// Typically the end position of the range denotes the line where the inline values are shown.
	StoppedLocation Range
}

// Provide inline value as text.
//
// @since 3.17.0
type InlineValueText struct {
	// The document range for which the inline value applies.
	Range Range
	// The text of the inline value.
	Text string
}

// Provide inline value through a variable lookup.
// If only a range is specified, the variable name will be extracted from the underlying document.
// An optional variable name can be used to override the extracted name.
//
// @since 3.17.0
type InlineValueVariableLookup struct {
	// The document range for which the inline value applies.
	// The range is used to extract the variable name from the underlying document.
	Range Range
	// If specified the name of the variable to look up.
	VariableName/*TOpt*/ *String
	// How to perform the lookup.
	CaseSensitiveLookup bool
}

// Provide an inline value through an expression evaluation.
// If only a range is specified, the expression will be extracted from the underlying document.
// An optional expression can be used to override the extracted expression.
//
// @since 3.17.0
type InlineValueEvaluatableExpression struct {
	// The document range for which the inline value applies.
	// The range is used to extract the evaluatable expression from the underlying document.
	Range Range
	// If specified the expression overrides the extracted expression.
	Expression/*TOpt*/ *String
}

// Inline value options used during static registration.
//
// @since 3.17.0
type InlineValueOptions struct {
	WorkDoneProgressOptions
}

// An inlay hint label part allows for interactive and composite labels
// of inlay hints.
//
// @since 3.17.0
type InlayHintLabelPart struct {
	// The value of this label part.
	Value string
	// The tooltip text when you hover over this label part. Depending on
	// the client capability `inlayHint.resolveSupport` clients might resolve
	// this property late using the resolve request.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Tooltip/*TOpt*/ /*TOr*/ struct {
		String/*TOpt*/ *String
		MarkupContent/*TOpt*/ *MarkupContent
	}

	// An optional source code location that represents this
	// label part.
	//
	// The editor will use this location for the hover and for code navigation
	// features: This part will become a clickable link that resolves to the
	// definition of the symbol at the given location (not necessarily the
	// location itself), it shows the hover that shows at the given location,
	// and it shows a context menu with further code navigation commands.
	//
	// Depending on the client capability `inlayHint.resolveSupport` clients
	// might resolve this property late using the resolve request.
	Location/*TOpt*/ *Location
	// An optional command for this label part.
	//
	// Depending on the client capability `inlayHint.resolveSupport` clients
	// might resolve this property late using the resolve request.
	Command/*TOpt*/ *Command
}

// A `MarkupContent` literal represents a string value which content is interpreted base on its
// kind flag. Currently the protocol supports `plaintext` and `markdown` as markup kinds.
//
// If the kind is `markdown` then the value can contain fenced code blocks like in GitHub issues.
// See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting
//
// Here is an example how such a string can be constructed using JavaScript / TypeScript:
// ```ts
//
//	let markdown: MarkdownContent = {
//	 kind: MarkupKind.Markdown,
//	 value: [
//	   '# Header',
//	   'Some text',
//	   '```typescript',
//	   'someCode();',
//	   '```'
//	 ].join('\n')
//	};
//
// ```
//
// *Please Note* that clients might sanitize the return markdown. A client could decide to
// remove HTML from the markdown to avoid script execution.
type MarkupContent struct {
	// The type of the Markup
	Kind MarkupKind
	// The content itself
	Value string
}

// Inlay hint options used during static registration.
//
// @since 3.17.0
type InlayHintOptions struct {
	WorkDoneProgressOptions
	// The server provides support to resolve additional
	// information for an inlay hint item.
	ResolveProvider/*TOpt*/ *Boolean
}

// A full diagnostic report with a set of related documents.
//
// @since 3.17.0
type RelatedFullDocumentDiagnosticReport struct {
	FullDocumentDiagnosticReport
	// Diagnostics of related documents. This information is useful
	// in programming languages where code in a file A can generate
	// diagnostics in a file B which A depends on. An example of
	// such a language is C/C++ where marco definitions in a file
	// a.cpp and result in errors in a header file b.hpp.
	//
	// @since 3.17.0
	//
	// Every object in the map has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	RelatedDocuments/*TOpt*/ map[DocumentURI] /*TOr*/ struct {
		FullDocumentDiagnosticReport/*TOpt*/ *FullDocumentDiagnosticReport
		UnchangedDocumentDiagnosticReport/*TOpt*/ *UnchangedDocumentDiagnosticReport
	}
}

// An unchanged diagnostic report with a set of related documents.
//
// @since 3.17.0
type RelatedUnchangedDocumentDiagnosticReport struct {
	UnchangedDocumentDiagnosticReport
	// Diagnostics of related documents. This information is useful
	// in programming languages where code in a file A can generate
	// diagnostics in a file B which A depends on. An example of
	// such a language is C/C++ where marco definitions in a file
	// a.cpp and result in errors in a header file b.hpp.
	//
	// @since 3.17.0
	//
	// Every object in the map has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	RelatedDocuments/*TOpt*/ map[DocumentURI] /*TOr*/ struct {
		FullDocumentDiagnosticReport/*TOpt*/ *FullDocumentDiagnosticReport
		UnchangedDocumentDiagnosticReport/*TOpt*/ *UnchangedDocumentDiagnosticReport
	}
}

// A diagnostic report with a full set of problems.
//
// @since 3.17.0
type FullDocumentDiagnosticReport struct {
	// A full document diagnostic report.
	// The value is always "full".
	// The value is always "full".
	Kind string
	// An optional result id. If provided it will
	// be sent on the next diagnostic request for the
	// same document.
	ResultId/*TOpt*/ *String
	// The actual items.
	Items []Diagnostic
}

// A diagnostic report indicating that the last returned
// report is still accurate.
//
// @since 3.17.0
type UnchangedDocumentDiagnosticReport struct {
	// A document diagnostic report indicating
	// no changes to the last result. A server can
	// only return `unchanged` if result ids are
	// provided.
	// The value is always "unchanged".
	// The value is always "unchanged".
	Kind string
	// A result id which will be sent on the next
	// diagnostic request for the same document.
	ResultId string
}

// Diagnostic options.
//
// @since 3.17.0
type DiagnosticOptions struct {
	WorkDoneProgressOptions
	// An optional identifier under which the diagnostics are
	// managed by the client.
	Identifier/*TOpt*/ *String
	// Whether the language has inter file dependencies meaning that
	// editing code in one file can result in a different diagnostic
	// set in another file. Inter file dependencies are common for
	// most programming languages and typically uncommon for linters.
	InterFileDependencies bool
	// The server provides support for workspace diagnostics as well.
	WorkspaceDiagnostics bool
}

// A previous result id in a workspace pull request.
//
// @since 3.17.0
type PreviousResultId struct {
	// The URI for which the client knowns a
	// result id.
	Uri DocumentURI
	// The value of the previous result id.
	Value string
}

// A notebook document.
//
// @since 3.17.0
type NotebookDocument struct {
	// The notebook document's uri.
	Uri URI
	// The type of the notebook.
	NotebookType string
	// The version number of this document (it will increase after each
	// change, including undo/redo).
	Version int
	// Additional metadata stored with the notebook
	// document.
	//
	// Note: should always be an object literal (e.g. LSPObject)
	Metadata/*TOpt*/ LSPObject
	// The cells of a notebook.
	Cells []NotebookCell
}

// An item to transfer a text document from the client to the
// server.
type TextDocumentItem struct {
	// The text document's uri.
	Uri DocumentURI
	// The text document's language identifier.
	LanguageId string
	// The version number of this document (it will increase after each
	// change, including undo/redo).
	Version int
	// The content of the opened text document.
	Text string
}

// A versioned notebook document identifier.
//
// @since 3.17.0
type VersionedNotebookDocumentIdentifier struct {
	// The version number of this notebook document.
	Version int
	// The notebook document's uri.
	Uri URI
}

// A change event for a notebook document.
//
// @since 3.17.0
type NotebookDocumentChangeEvent struct {
	// The changed meta data if any.
	//
	// Note: should always be an object literal (e.g. LSPObject)
	Metadata/*TOpt*/ LSPObject
	// Changes to cells
	Cells/*TOpt*/ * /*TStruc*/ struct {
		// Changes to the cell structure to add or
		// remove cells.
		Structure/*TOpt*/ * /*TStruc*/ struct {
			// The change to the cell array.
			Array NotebookCellArrayChange
			// Additional opened cell text documents.
			DidOpen/*TOpt*/ []TextDocumentItem
			// Additional closed cell text documents.
			DidClose/*TOpt*/ []TextDocumentIdentifier
		}

		// Changes to notebook cells properties like its
		// kind, execution summary or metadata.
		Data/*TOpt*/ []NotebookCell
		// Changes to the text content of notebook cells.
		TextContent/*TOpt*/ [] /*TStruc*/ struct {
			Document VersionedTextDocumentIdentifier
			Changes  []TextDocumentContentChangeEvent
		}
	}
}

// A literal to identify a notebook document in the client.
//
// @since 3.17.0
type NotebookDocumentIdentifier struct {
	// The notebook document's uri.
	Uri URI
}

// General parameters to register for a notification or to register a provider.
type Registration struct {
	// The id used to register the request. The id can be used to deregister
	// the request again.
	Id string
	// The method / capability to register for.
	Method string
	// Options necessary for the registration.
	RegisterOptions/*TOpt*/ LSPAny
}

// General parameters to unregister a request or notification.
type Unregistration struct {
	// The id used to unregister the request or notification. Usually an id
	// provided during the register request.
	Id string
	// The method to unregister for.
	Method string
}

// The initialize parameters
type _InitializeParams struct {
	WorkDoneProgressParams
	// The process Id of the parent process that started
	// the server.
	//
	// Is `null` if the process has not been started by another process.
	// If the parent process is not alive then the server should exit.
	ProcessId/*TOr*/ /*TOpt*/ *Integer

	// Information about the client
	//
	// @since 3.15.0
	ClientInfo/*TOpt*/ * /*TStruc*/ struct {
		// The name of the client as defined by the client.
		Name string
		// The client's version as defined by the client.
		Version/*TOpt*/ *String
	}

	// The locale the client is currently showing the user interface
	// in. This must not necessarily be the locale of the operating
	// system.
	//
	// Uses IETF language tags as the value's syntax
	// (See https://en.wikipedia.org/wiki/IETF_language_tag)
	//
	// @since 3.16.0
	Locale/*TOpt*/ *String
	// The rootPath of the workspace. Is null
	// if no folder is open.
	//
	// @deprecated in favour of rootUri.
	RootPath/*TOpt*/ /*TOr*/ /*TOpt*/ *String

	// The rootUri of the workspace. Is null if no
	// folder is open. If both `rootPath` and `rootUri` are set
	// `rootUri` wins.
	//
	// @deprecated in favour of workspaceFolders.
	RootUri/*TOr*/ /*TOpt*/ *DocumentURI

	// The capabilities provided by the client (editor or tool)
	Capabilities ClientCapabilities
	// User provided initialization options.
	InitializationOptions/*TOpt*/ LSPAny
	// The initial trace setting. If omitted trace is disabled ('off').
	Trace/*TOpt*/ TraceValues
}
type WorkspaceFoldersInitializeParams struct {
	// The workspace folders configured in the client when the server starts.
	//
	// This property is only available if the client supports workspace folders.
	// It can be `null` if the client supports workspace folders but none are
	// configured.
	//
	// @since 3.6.0
	WorkspaceFolders /*TOpt*/ /*TOr*/ /*TOpt*/ []WorkspaceFolder
}

// Defines the capabilities provided by a language
// server.
type ServerCapabilities struct {
	// The position encoding the server picked from the encodings offered
	// by the client via the client capability `general.positionEncodings`.
	//
	// If the client didn't provide any position encodings the only valid
	// value that a server can return is 'utf-16'.
	//
	// If omitted it defaults to 'utf-16'.
	//
	// @since 3.17.0
	PositionEncoding/*TOpt*/ PositionEncodingKind
	// Defines how text documents are synced. Is either a detailed structure
	// defining each notification or for backwards compatibility the
	// TextDocumentSyncKind number.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	TextDocumentSync/*TOpt*/ /*TOr*/ struct {
		TextDocumentSyncOptions/*TOpt*/ *TextDocumentSyncOptions
		TextDocumentSyncKind/*TOpt*/ TextDocumentSyncKind
	}

	// Defines how notebook documents are synced.
	//
	// @since 3.17.0
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	NotebookDocumentSync/*TOpt*/ /*TOr*/ struct {
		NotebookDocumentSyncOptions/*TOpt*/ *NotebookDocumentSyncOptions
		NotebookDocumentSyncRegistrationOptions/*TOpt*/ *NotebookDocumentSyncRegistrationOptions
	}

	// The server provides completion support.
	CompletionProvider/*TOpt*/ *CompletionOptions
	// The server provides hover support.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	HoverProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		HoverOptions/*TOpt*/ *HoverOptions
	}

	// The server provides signature help support.
	SignatureHelpProvider/*TOpt*/ *SignatureHelpOptions
	// The server provides Goto Declaration support.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	DeclarationProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		DeclarationOptions/*TOpt*/ *DeclarationOptions
		DeclarationRegistrationOptions/*TOpt*/ *DeclarationRegistrationOptions
	}

	// The server provides goto definition support.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	DefinitionProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		DefinitionOptions/*TOpt*/ *DefinitionOptions
	}

	// The server provides Goto Type Definition support.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	TypeDefinitionProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		TypeDefinitionOptions/*TOpt*/ *TypeDefinitionOptions
		TypeDefinitionRegistrationOptions/*TOpt*/ *TypeDefinitionRegistrationOptions
	}

	// The server provides Goto Implementation support.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	ImplementationProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		ImplementationOptions/*TOpt*/ *ImplementationOptions
		ImplementationRegistrationOptions/*TOpt*/ *ImplementationRegistrationOptions
	}

	// The server provides find references support.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	ReferencesProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		ReferenceOptions/*TOpt*/ *ReferenceOptions
	}

	// The server provides document highlight support.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	DocumentHighlightProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		DocumentHighlightOptions/*TOpt*/ *DocumentHighlightOptions
	}

	// The server provides document symbol support.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	DocumentSymbolProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		DocumentSymbolOptions/*TOpt*/ *DocumentSymbolOptions
	}

	// The server provides code actions. CodeActionOptions may only be
	// specified if the client states that it supports
	// `codeActionLiteralSupport` in its initial `initialize` request.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	CodeActionProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		CodeActionOptions/*TOpt*/ *CodeActionOptions
	}

	// The server provides code lens.
	CodeLensProvider/*TOpt*/ *CodeLensOptions
	// The server provides document link support.
	DocumentLinkProvider/*TOpt*/ *DocumentLinkOptions
	// The server provides color provider support.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	ColorProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		DocumentColorOptions/*TOpt*/ *DocumentColorOptions
		DocumentColorRegistrationOptions/*TOpt*/ *DocumentColorRegistrationOptions
	}

	// The server provides workspace symbol support.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	WorkspaceSymbolProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		WorkspaceSymbolOptions/*TOpt*/ *WorkspaceSymbolOptions
	}

	// The server provides document formatting.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	DocumentFormattingProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		DocumentFormattingOptions/*TOpt*/ *DocumentFormattingOptions
	}

	// The server provides document range formatting.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	DocumentRangeFormattingProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		DocumentRangeFormattingOptions/*TOpt*/ *DocumentRangeFormattingOptions
	}

	// The server provides document formatting on typing.
	DocumentOnTypeFormattingProvider/*TOpt*/ *DocumentOnTypeFormattingOptions
	// The server provides rename support. RenameOptions may only be
	// specified if the client states that it supports
	// `prepareSupport` in its initial `initialize` request.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	RenameProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		RenameOptions/*TOpt*/ *RenameOptions
	}

	// The server provides folding provider support.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	FoldingRangeProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		FoldingRangeOptions/*TOpt*/ *FoldingRangeOptions
		FoldingRangeRegistrationOptions/*TOpt*/ *FoldingRangeRegistrationOptions
	}

	// The server provides selection range support.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	SelectionRangeProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		SelectionRangeOptions/*TOpt*/ *SelectionRangeOptions
		SelectionRangeRegistrationOptions/*TOpt*/ *SelectionRangeRegistrationOptions
	}

	// The server provides execute command support.
	ExecuteCommandProvider/*TOpt*/ *ExecuteCommandOptions
	// The server provides call hierarchy support.
	//
	// @since 3.16.0
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	CallHierarchyProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		CallHierarchyOptions/*TOpt*/ *CallHierarchyOptions
		CallHierarchyRegistrationOptions/*TOpt*/ *CallHierarchyRegistrationOptions
	}

	// The server provides linked editing range support.
	//
	// @since 3.16.0
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	LinkedEditingRangeProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		LinkedEditingRangeOptions/*TOpt*/ *LinkedEditingRangeOptions
		LinkedEditingRangeRegistrationOptions/*TOpt*/ *LinkedEditingRangeRegistrationOptions
	}

	// The server provides semantic tokens support.
	//
	// @since 3.16.0
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	SemanticTokensProvider/*TOpt*/ /*TOr*/ struct {
		SemanticTokensOptions/*TOpt*/ *SemanticTokensOptions
		SemanticTokensRegistrationOptions/*TOpt*/ *SemanticTokensRegistrationOptions
	}

	// The server provides moniker support.
	//
	// @since 3.16.0
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	MonikerProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		MonikerOptions/*TOpt*/ *MonikerOptions
		MonikerRegistrationOptions/*TOpt*/ *MonikerRegistrationOptions
	}

	// The server provides type hierarchy support.
	//
	// @since 3.17.0
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	TypeHierarchyProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		TypeHierarchyOptions/*TOpt*/ *TypeHierarchyOptions
		TypeHierarchyRegistrationOptions/*TOpt*/ *TypeHierarchyRegistrationOptions
	}

	// The server provides inline values.
	//
	// @since 3.17.0
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	InlineValueProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		InlineValueOptions/*TOpt*/ *InlineValueOptions
		InlineValueRegistrationOptions/*TOpt*/ *InlineValueRegistrationOptions
	}

	// The server provides inlay hints.
	//
	// @since 3.17.0
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	InlayHintProvider/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		InlayHintOptions/*TOpt*/ *InlayHintOptions
		InlayHintRegistrationOptions/*TOpt*/ *InlayHintRegistrationOptions
	}

	// The server has support for pull model diagnostics.
	//
	// @since 3.17.0
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	DiagnosticProvider/*TOpt*/ /*TOr*/ struct {
		DiagnosticOptions/*TOpt*/ *DiagnosticOptions
		DiagnosticRegistrationOptions/*TOpt*/ *DiagnosticRegistrationOptions
	}

	// Workspace specific server capabilities.
	Workspace/*TOpt*/ * /*TStruc*/ struct {
		// The server supports workspace folder.
		//
		// @since 3.6.0
		WorkspaceFolders/*TOpt*/ *WorkspaceFoldersServerCapabilities
		// The server is interested in notifications/requests for operations on files.
		//
		// @since 3.16.0
		FileOperations/*TOpt*/ *FileOperationOptions
	}

	// Experimental server capabilities.
	Experimental/*TOpt*/ LSPAny
}

// A text document identifier to denote a specific version of a text document.
type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	// The version number of this document.
	Version int
}

// Save options.
type SaveOptions struct {
	// The client is supposed to include the content on save.
	IncludeText /*TOpt*/ *Boolean
}

// An event describing a file change.
type FileEvent struct {
	// The file's uri.
	Uri DocumentURI
	// The change type.
	Type FileChangeType
}
type FileSystemWatcher struct {
	// The glob pattern to watch. See `GlobPattern` for more detail.
	//
	// @since 3.17.0 support for relative patterns.
	GlobPattern GlobPattern
	// The kind of events of interest. If omitted it defaults
	// to WatchKind.Create | WatchKind.Change | WatchKind.Delete
	// which is 7.
	Kind/*TOpt*/ WatchKind
}

// Represents a diagnostic, such as a compiler error or warning. Diagnostic objects
// are only valid in the scope of a resource.
type Diagnostic struct {
	// The range at which the message applies
	Range Range
	// The diagnostic's severity. Can be omitted. If omitted it is up to the
	// client to interpret diagnostics as error, warning, info or hint.
	Severity/*TOpt*/ DiagnosticSeverity
	// The diagnostic's code, which usually appear in the user interface.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Code/*TOpt*/ /*TOr*/ struct {
		Integer/*TOpt*/ *Integer
		String/*TOpt*/ *String
	}

	// An optional property to describe the error code.
	// Requires the code field (above) to be present/not null.
	//
	// @since 3.16.0
	CodeDescription/*TOpt*/ *CodeDescription
	// A human-readable string describing the source of this
	// diagnostic, e.g. 'typescript' or 'super lint'. It usually
	// appears in the user interface.
	Source/*TOpt*/ *String
	// The diagnostic's message. It usually appears in the user interface
	Message string
	// Additional metadata about the diagnostic.
	//
	// @since 3.15.0
	Tags/*TOpt*/ []DiagnosticTag
	// An array of related diagnostic information, e.g. when symbol-names within
	// a scope collide all definitions can be marked via this property.
	RelatedInformation/*TOpt*/ []DiagnosticRelatedInformation
	// A data entry field that is preserved between a `textDocument/publishDiagnostics`
	// notification and `textDocument/codeAction` request.
	//
	// @since 3.16.0
	Data/*TOpt*/ LSPAny
}

// Contains additional information about the context in which a completion request is triggered.
type CompletionContext struct {
	// How the completion was triggered.
	TriggerKind CompletionTriggerKind
	// The trigger character (a single character) that has trigger code complete.
	// Is undefined if `triggerKind !== CompletionTriggerKind.TriggerCharacter`
	TriggerCharacter/*TOpt*/ *String
}

// Additional details for a completion item label.
//
// @since 3.17.0
type CompletionItemLabelDetails struct {
	// An optional string which is rendered less prominently directly after `CompletionItem.label`,
	// without any spacing. Should be used for function signatures and type annotations.
	Detail/*TOpt*/ *String
	// An optional string which is rendered less prominently after `CompletionItem.detail`. Should be used
	// for fully qualified names and file paths.
	Description/*TOpt*/ *String
}

// A special text edit to provide an insert and a replace operation.
//
// @since 3.16.0
type InsertReplaceEdit struct {
	// The string to be inserted.
	NewText string
	// The range if the insert is requested
	Insert Range
	// The range if the replace is requested.
	Replace Range
}

// Completion options.
type CompletionOptions struct {
	WorkDoneProgressOptions
	// Most tools trigger completion request automatically without explicitly requesting
	// it using a keyboard shortcut (e.g. Ctrl+Space). Typically they do so when the user
	// starts to type an identifier. For example if the user types `c` in a JavaScript file
	// code complete will automatically pop up present `console` besides others as a
	// completion item. Characters that make up identifiers don't need to be listed here.
	//
	// If code complete should automatically be trigger on characters not being valid inside
	// an identifier (for example `.` in JavaScript) list them in `triggerCharacters`.
	TriggerCharacters/*TOpt*/ []string
	// The list of all possible characters that commit a completion. This field can be used
	// if clients don't support individual commit characters per completion item. See
	// `ClientCapabilities.textDocument.completion.completionItem.commitCharactersSupport`
	//
	// If a server provides both `allCommitCharacters` and commit characters on an individual
	// completion item the ones on the completion item win.
	//
	// @since 3.2.0
	AllCommitCharacters/*TOpt*/ []string
	// The server provides support to resolve additional
	// information for a completion item.
	ResolveProvider/*TOpt*/ *Boolean
	// The server supports the following `CompletionItem` specific
	// capabilities.
	//
	// @since 3.17.0
	CompletionItem/*TOpt*/ * /*TStruc*/ struct {
		// The server has support for completion item label
		// details (see also `CompletionItemLabelDetails`) when
		// receiving a completion item in a resolve call.
		//
		// @since 3.17.0
		LabelDetailsSupport /*TOpt*/ *Boolean
	}
}

// Hover options.
type HoverOptions struct {
	WorkDoneProgressOptions
}

// Additional information about the context in which a signature help request was triggered.
//
// @since 3.15.0
type SignatureHelpContext struct {
	// Action that caused signature help to be triggered.
	TriggerKind SignatureHelpTriggerKind
	// Character that caused signature help to be triggered.
	//
	// This is undefined when `triggerKind !== SignatureHelpTriggerKind.TriggerCharacter`
	TriggerCharacter/*TOpt*/ *String
	// `true` if signature help was already showing when it was triggered.
	//
	// Retriggers occurs when the signature help is already active and can be caused by actions such as
	// typing a trigger character, a cursor move, or document content changes.
	IsRetrigger bool
	// The currently active `SignatureHelp`.
	//
	// The `activeSignatureHelp` has its `SignatureHelp.activeSignature` field updated based on
	// the user navigating through available signatures.
	ActiveSignatureHelp/*TOpt*/ *SignatureHelp
}

// Represents the signature of something callable. A signature
// can have a label, like a function-name, a doc-comment, and
// a set of parameters.
type SignatureInformation struct {
	// The label of this signature. Will be shown in
	// the UI.
	Label string
	// The human-readable doc-comment of this signature. Will be shown
	// in the UI but can be omitted.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Documentation/*TOpt*/ /*TOr*/ struct {
		String/*TOpt*/ *String
		MarkupContent/*TOpt*/ *MarkupContent
	}

	// The parameters of this signature.
	Parameters/*TOpt*/ []ParameterInformation
	// The index of the active parameter.
	//
	// If provided, this is used in place of `SignatureHelp.activeParameter`.
	//
	// @since 3.16.0
	ActiveParameter/*TOpt*/ *Uinteger
}

// Server Capabilities for a `SignatureHelpRequest`.
type SignatureHelpOptions struct {
	WorkDoneProgressOptions
	// List of characters that trigger signature help automatically.
	TriggerCharacters/*TOpt*/ []string
	// List of characters that re-trigger signature help.
	//
	// These trigger characters are only active when signature help is already showing. All trigger characters
	// are also counted as re-trigger characters.
	//
	// @since 3.15.0
	RetriggerCharacters/*TOpt*/ []string
}

// Server Capabilities for a `DefinitionRequest`.
type DefinitionOptions struct {
	WorkDoneProgressOptions
}

// Value-object that contains additional information when
// requesting references.
type ReferenceContext struct {
	// Include the declaration of the current symbol.
	IncludeDeclaration bool
}

// Reference options.
type ReferenceOptions struct {
	WorkDoneProgressOptions
}

// Provider options for a `DocumentHighlightRequest`.
type DocumentHighlightOptions struct {
	WorkDoneProgressOptions
}

// A base for all symbol information.
type BaseSymbolInformation struct {
	// The name of this symbol.
	Name string
	// The kind of this symbol.
	Kind SymbolKind
	// Tags for this symbol.
	//
	// @since 3.16.0
	Tags/*TOpt*/ []SymbolTag
	// The name of the symbol containing this symbol. This information is for
	// user interface purposes (e.g. to render a qualifier in the user interface
	// if necessary). It can't be used to re-infer a hierarchy for the document
	// symbols.
	ContainerName/*TOpt*/ *String
}

// Provider options for a `DocumentSymbolRequest`.
type DocumentSymbolOptions struct {
	WorkDoneProgressOptions
	// A human-readable string that is shown when multiple outlines trees
	// are shown for the same document.
	//
	// @since 3.16.0
	Label/*TOpt*/ *String
}

// Contains additional diagnostic information about the context in which
// a `CodeActionProvider.provideCodeActions` is run.
type CodeActionContext struct {
	// An array of diagnostics known on the client side overlapping the range provided to the
	// `textDocument/codeAction` request. They are provided so that the server knows which
	// errors are currently presented to the user for the given range. There is no guarantee
	// that these accurately reflect the error state of the resource. The primary parameter
	// to compute code actions is the provided range.
	Diagnostics []Diagnostic
	// Requested kind of actions to return.
	//
	// Actions not of this kind are filtered out by the client before being shown. So servers
	// can omit computing them.
	Only/*TOpt*/ []CodeActionKind
	// The reason why code actions were requested.
	//
	// @since 3.17.0
	TriggerKind/*TOpt*/ CodeActionTriggerKind
}

// Provider options for a `CodeActionRequest`.
type CodeActionOptions struct {
	WorkDoneProgressOptions
	// CodeActionKinds that this server may return.
	//
	// The list of kinds may be generic, such as `CodeActionKind.Refactor`, or the server
	// may list out every specific kind they provide.
	CodeActionKinds/*TOpt*/ []CodeActionKind
	// The server provides support to resolve additional
	// information for a code action.
	//
	// @since 3.16.0
	ResolveProvider/*TOpt*/ *Boolean
}

// Server capabilities for a `WorkspaceSymbolRequest`.
type WorkspaceSymbolOptions struct {
	WorkDoneProgressOptions
	// The server provides support to resolve additional
	// information for a workspace symbol.
	//
	// @since 3.17.0
	ResolveProvider/*TOpt*/ *Boolean
}

// Code Lens provider options of a `CodeLensRequest`.
type CodeLensOptions struct {
	WorkDoneProgressOptions
	// Code lens has a resolve provider as well.
	ResolveProvider/*TOpt*/ *Boolean
}

// Provider options for a `DocumentLinkRequest`.
type DocumentLinkOptions struct {
	WorkDoneProgressOptions
	// Document links have a resolve provider as well.
	ResolveProvider/*TOpt*/ *Boolean
}

// Value-object describing what options formatting should use.
type FormattingOptions struct {
	// Size of a tab in spaces.
	TabSize uint
	// Prefer spaces over tabs.
	InsertSpaces bool
	// Trim trailing whitespace on a line.
	//
	// @since 3.15.0
	TrimTrailingWhitespace/*TOpt*/ *Boolean
	// Insert a newline character at the end of the file if one does not exist.
	//
	// @since 3.15.0
	InsertFinalNewline/*TOpt*/ *Boolean
	// Trim all newlines after the final newline at the end of the file.
	//
	// @since 3.15.0
	TrimFinalNewlines/*TOpt*/ *Boolean
}

// Provider options for a `DocumentFormattingRequest`.
type DocumentFormattingOptions struct {
	WorkDoneProgressOptions
}

// Provider options for a `DocumentRangeFormattingRequest`.
type DocumentRangeFormattingOptions struct {
	WorkDoneProgressOptions
}

// Provider options for a `DocumentOnTypeFormattingRequest`.
type DocumentOnTypeFormattingOptions struct {
	// A character on which formatting should be triggered, like `{`.
	FirstTriggerCharacter string
	// More trigger characters.
	MoreTriggerCharacter/*TOpt*/ []string
}

// Provider options for a `RenameRequest`.
type RenameOptions struct {
	WorkDoneProgressOptions
	// Renames should be checked and tested before being executed.
	//
	// @since version 3.12.0
	PrepareProvider/*TOpt*/ *Boolean
}

// The server capabilities of a `ExecuteCommandRequest`.
type ExecuteCommandOptions struct {
	WorkDoneProgressOptions
	// The commands to be executed on the server
	Commands []string
}

// @since 3.16.0
type SemanticTokensLegend struct {
	// The token types a server uses.
	TokenTypes []string
	// The token modifiers a server uses.
	TokenModifiers []string
}

// A text document identifier to optionally denote a specific version of a text document.
type OptionalVersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	// The version number of this document. If a versioned text document identifier
	// is sent from the server to the client and the file is not open in the editor
	// (the server has not received an open notification before) the server can send
	// `null` to indicate that the version is unknown and the content on disk is the
	// truth (as specified with document content ownership).
	Version/*TOr*/ /*TOpt*/ *Integer
}

// A special text edit with an additional change annotation.
//
// @since 3.16.0.
type AnnotatedTextEdit struct {
	TextEdit
	// The actual identifier of the change annotation
	AnnotationId ChangeAnnotationIdentifier
}

// A generic resource operation.
type ResourceOperation struct {
	// The resource operation kind.
	Kind string
	// An optional annotation identifier describing the operation.
	//
	// @since 3.16.0
	AnnotationId/*TOpt*/ *String
}

// Options to create a file.
type CreateFileOptions struct {
	// Overwrite existing file. Overwrite wins over `ignoreIfExists`
	Overwrite/*TOpt*/ *Boolean
	// Ignore if exists.
	IgnoreIfExists/*TOpt*/ *Boolean
}

// Rename file options
type RenameFileOptions struct {
	// Overwrite target if existing. Overwrite wins over `ignoreIfExists`
	Overwrite/*TOpt*/ *Boolean
	// Ignores if target exists.
	IgnoreIfExists/*TOpt*/ *Boolean
}

// Delete file options
type DeleteFileOptions struct {
	// Delete the content recursively if a folder is denoted.
	Recursive/*TOpt*/ *Boolean
	// Ignore the operation if the file doesn't exist.
	IgnoreIfNotExists/*TOpt*/ *Boolean
}

// A pattern to describe in which file operation requests or notifications
// the server is interested in receiving.
//
// @since 3.16.0
type FileOperationPattern struct {
	// The glob pattern to match. Glob patterns can have the following syntax:
	// - `*` to match one or more characters in a path segment
	// - `?` to match on one character in a path segment
	// - `**` to match any number of path segments, including none
	// - `{}` to group sub patterns into an OR expression. (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
	// - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	// - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
	Glob string
	// Whether to match files or folders with this pattern.
	//
	// Matches both if undefined.
	Matches/*TOpt*/ FileOperationPatternKind
	// Additional options used during matching.
	Options/*TOpt*/ *FileOperationPatternOptions
}

// A full document diagnostic report for a workspace diagnostic result.
//
// @since 3.17.0
type WorkspaceFullDocumentDiagnosticReport struct {
	FullDocumentDiagnosticReport
	// The URI for which diagnostic information is reported.
	Uri DocumentURI
	// The version number for which the diagnostics are reported.
	// If the document is not marked as open `null` can be provided.
	Version/*TOr*/ /*TOpt*/ *Integer
}

// An unchanged document diagnostic report for a workspace diagnostic result.
//
// @since 3.17.0
type WorkspaceUnchangedDocumentDiagnosticReport struct {
	UnchangedDocumentDiagnosticReport
	// The URI for which diagnostic information is reported.
	Uri DocumentURI
	// The version number for which the diagnostics are reported.
	// If the document is not marked as open `null` can be provided.
	Version/*TOr*/ /*TOpt*/ *Integer
}

// A notebook cell.
//
// A cell's document URI must be unique across ALL notebook
// cells and can therefore be used to uniquely identify a
// notebook cell or the cell's text document.
//
// @since 3.17.0
type NotebookCell struct {
	// The cell's kind
	Kind NotebookCellKind
	// The URI of the cell's text document
	// content.
	Document DocumentURI
	// Additional metadata stored with the cell.
	//
	// Note: should always be an object literal (e.g. LSPObject)
	Metadata/*TOpt*/ LSPObject
	// Additional execution summary information
	// if supported by the client.
	ExecutionSummary/*TOpt*/ *ExecutionSummary
}

// A change describing how to move a `NotebookCell`
// array from state S to S'.
//
// @since 3.17.0
type NotebookCellArrayChange struct {
	// The start oftest of the cell that changed.
	Start uint
	// The deleted cells
	DeleteCount uint
	// The new cells, if any
	Cells/*TOpt*/ []NotebookCell
}

// Defines the capabilities provided by the client.
type ClientCapabilities struct {
	// Workspace specific client capabilities.
	Workspace/*TOpt*/ *WorkspaceClientCapabilities
	// Text document specific client capabilities.
	TextDocument/*TOpt*/ *TextDocumentClientCapabilities
	// Capabilities specific to the notebook document support.
	//
	// @since 3.17.0
	NotebookDocument/*TOpt*/ *NotebookDocumentClientCapabilities
	// Window specific client capabilities.
	Window/*TOpt*/ *WindowClientCapabilities
	// General client capabilities.
	//
	// @since 3.16.0
	General/*TOpt*/ *GeneralClientCapabilities
	// Experimental client capabilities.
	Experimental/*TOpt*/ LSPAny
}
type TextDocumentSyncOptions struct {
	// Open and close notifications are sent to the server. If omitted open close notification should not
	// be sent.
	OpenClose/*TOpt*/ *Boolean
	// Change notifications are sent to the server. See TextDocumentSyncKind.None, TextDocumentSyncKind.Full
	// and TextDocumentSyncKind.Incremental. If omitted it defaults to TextDocumentSyncKind.None.
	Change/*TOpt*/ TextDocumentSyncKind
	// If present will save notifications are sent to the server. If omitted the notification should not be
	// sent.
	WillSave/*TOpt*/ *Boolean
	// If present will save wait until requests are sent to the server. If omitted the request should not be
	// sent.
	WillSaveWaitUntil/*TOpt*/ *Boolean
	// If present save notifications are sent to the server. If omitted the notification should not be
	// sent.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Save/*TOpt*/ /*TOr*/ struct {
		Boolean/*TOpt*/ *Boolean
		SaveOptions/*TOpt*/ *SaveOptions
	}
}

// Options specific to a notebook plus its cells
// to be synced to the server.
//
// If a selector provides a notebook document
// filter but no cell selector all cells of a
// matching notebook document will be synced.
//
// If a selector provides no notebook document
// filter but only a cell selector all notebook
// document that contain at least one matching
// cell will be synced.
//
// @since 3.17.0
type NotebookDocumentSyncOptions struct {
	// The notebooks to be synced
	NotebookSelector [] /*TStruc*/ struct {
		// The notebook to be synced If a string
		// value is provided it matches against the
		// notebook type. '*' matches every notebook.
		//
		// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
		Notebook/*TOpt*/ /*TOr*/ struct {
			String/*TOpt*/ *String
			NotebookDocumentFilter/*TOpt*/ *NotebookDocumentFilter
		}

		// The cells of the matching notebook to be synced.
		Cells/*TOpt*/ [] /*TStruc*/ struct {
			Language string
		}
	}

	// Whether save notification should be forwarded to
	// the server. Will only be honored if mode === `notebook`.
	Save/*TOpt*/ *Boolean
}

// Registration options specific to a notebook.
//
// @since 3.17.0
type NotebookDocumentSyncRegistrationOptions struct {
	NotebookDocumentSyncOptions
	StaticRegistrationOptions
}
type WorkspaceFoldersServerCapabilities struct {
	// The server has support for workspace folders
	Supported/*TOpt*/ *Boolean
	// Whether the server wants to receive workspace folder
	// change notifications.
	//
	// If a string is provided the string is treated as an ID
	// under which the notification is registered on the client
	// side. The ID can be used to unregister for these events
	// using the `client/unregisterCapability` request.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	ChangeNotifications/*TOpt*/ /*TOr*/ struct {
		String/*TOpt*/ *String
		Boolean/*TOpt*/ *Boolean
	}
}

// Options for notifications/requests for user operations on files.
//
// @since 3.16.0
type FileOperationOptions struct {
	// The server is interested in receiving didCreateFiles notifications.
	DidCreate/*TOpt*/ *FileOperationRegistrationOptions
	// The server is interested in receiving willCreateFiles requests.
	WillCreate/*TOpt*/ *FileOperationRegistrationOptions
	// The server is interested in receiving didRenameFiles notifications.
	DidRename/*TOpt*/ *FileOperationRegistrationOptions
	// The server is interested in receiving willRenameFiles requests.
	WillRename/*TOpt*/ *FileOperationRegistrationOptions
	// The server is interested in receiving didDeleteFiles file notifications.
	DidDelete/*TOpt*/ *FileOperationRegistrationOptions
	// The server is interested in receiving willDeleteFiles file requests.
	WillDelete/*TOpt*/ *FileOperationRegistrationOptions
}

// Structure to capture a description for an error code.
//
// @since 3.16.0
type CodeDescription struct {
	// An URI to open with more information about the diagnostic error.
	Href URI
}

// Represents a related message and source code location for a diagnostic. This should be
// used to point to code locations that cause or related to a diagnostics, e.g when duplicating
// a symbol in a scope.
type DiagnosticRelatedInformation struct {
	// The location of this related diagnostic information.
	Location Location
	// The message of this related diagnostic information.
	Message string
}

// Represents a parameter of a callable-signature. A parameter can
// have a label and a doc-comment.
type ParameterInformation struct {
	// The label of this parameter information.
	//
	// Either a string or an inclusive start and exclusive end offsets within its containing
	// signature label. (see SignatureInformation.label). The offsets are based on a UTF-16
	// string representation as `Position` and `Range` does.
	//
	// *Note*: a label of type string should be a substring of its containing signature label.
	// Its intended use case is to highlight the parameter label part in the `SignatureInformation.label`.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Label/*TOr*/ struct {
		String/*TOpt*/ *String
		UintegerWithUinteger/*TOpt*/ * /*TTup*/ struct {
			Uinteger0 uint
			Uinteger1 uint
		}
	}

	// The human-readable doc-comment of this parameter. Will be shown
	// in the UI but can be omitted.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Documentation/*TOpt*/ /*TOr*/ struct {
		String/*TOpt*/ *String
		MarkupContent/*TOpt*/ *MarkupContent
	}
}

// A notebook cell text document filter denotes a cell text
// document by different properties.
//
// @since 3.17.0
type NotebookCellTextDocumentFilter struct {
	// A filter that matches against the notebook
	// containing the notebook cell. If a string
	// value is provided it matches against the
	// notebook type. '*' matches every notebook.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	Notebook/*TOr*/ struct {
		String/*TOpt*/ *String
		NotebookDocumentFilter/*TOpt*/ *NotebookDocumentFilter
	}

	// A language id like `python`.
	//
	// Will be matched against the language id of the
	// notebook cell document. '*' matches every language.
	Language/*TOpt*/ *String
}

// Matching options for the file operation pattern.
//
// @since 3.16.0
type FileOperationPatternOptions struct {
	// The pattern should be matched ignoring casing.
	IgnoreCase /*TOpt*/ *Boolean
}
type ExecutionSummary struct {
	// A strict monotonically increasing value
	// indicating the execution order of a cell
	// inside a notebook.
	ExecutionOrder uint
	// Whether the execution was successful or
	// not if known by the client.
	Success/*TOpt*/ *Boolean
}

// Workspace specific client capabilities.
type WorkspaceClientCapabilities struct {
	// The client supports applying batch edits
	// to the workspace by supporting the request
	// 'workspace/applyEdit'
	ApplyEdit/*TOpt*/ *Boolean
	// Capabilities specific to `WorkspaceEdit`s.
	WorkspaceEdit/*TOpt*/ *WorkspaceEditClientCapabilities
	// Capabilities specific to the `workspace/didChangeConfiguration` notification.
	DidChangeConfiguration/*TOpt*/ *DidChangeConfigurationClientCapabilities
	// Capabilities specific to the `workspace/didChangeWatchedFiles` notification.
	DidChangeWatchedFiles/*TOpt*/ *DidChangeWatchedFilesClientCapabilities
	// Capabilities specific to the `workspace/symbol` request.
	Symbol/*TOpt*/ *WorkspaceSymbolClientCapabilities
	// Capabilities specific to the `workspace/executeCommand` request.
	ExecuteCommand/*TOpt*/ *ExecuteCommandClientCapabilities
	// The client has support for workspace folders.
	//
	// @since 3.6.0
	WorkspaceFolders/*TOpt*/ *Boolean
	// The client supports `workspace/configuration` requests.
	//
	// @since 3.6.0
	Configuration/*TOpt*/ *Boolean
	// Capabilities specific to the semantic token requests scoped to the
	// workspace.
	//
	// @since 3.16.0.
	SemanticTokens/*TOpt*/ *SemanticTokensWorkspaceClientCapabilities
	// Capabilities specific to the code lens requests scoped to the
	// workspace.
	//
	// @since 3.16.0.
	CodeLens/*TOpt*/ *CodeLensWorkspaceClientCapabilities
	// The client has support for file notifications/requests for user operations on files.
	//
	// Since 3.16.0
	FileOperations/*TOpt*/ *FileOperationClientCapabilities
	// Capabilities specific to the inline values requests scoped to the
	// workspace.
	//
	// @since 3.17.0.
	InlineValue/*TOpt*/ *InlineValueWorkspaceClientCapabilities
	// Capabilities specific to the inlay hint requests scoped to the
	// workspace.
	//
	// @since 3.17.0.
	InlayHint/*TOpt*/ *InlayHintWorkspaceClientCapabilities
	// Capabilities specific to the diagnostic requests scoped to the
	// workspace.
	//
	// @since 3.17.0.
	Diagnostics/*TOpt*/ *DiagnosticWorkspaceClientCapabilities
}

// Text document specific client capabilities.
type TextDocumentClientCapabilities struct {
	// Defines which synchronization capabilities the client supports.
	Synchronization/*TOpt*/ *TextDocumentSyncClientCapabilities
	// Capabilities specific to the `textDocument/completion` request.
	Completion/*TOpt*/ *CompletionClientCapabilities
	// Capabilities specific to the `textDocument/hover` request.
	Hover/*TOpt*/ *HoverClientCapabilities
	// Capabilities specific to the `textDocument/signatureHelp` request.
	SignatureHelp/*TOpt*/ *SignatureHelpClientCapabilities
	// Capabilities specific to the `textDocument/declaration` request.
	//
	// @since 3.14.0
	Declaration/*TOpt*/ *DeclarationClientCapabilities
	// Capabilities specific to the `textDocument/definition` request.
	Definition/*TOpt*/ *DefinitionClientCapabilities
	// Capabilities specific to the `textDocument/typeDefinition` request.
	//
	// @since 3.6.0
	TypeDefinition/*TOpt*/ *TypeDefinitionClientCapabilities
	// Capabilities specific to the `textDocument/implementation` request.
	//
	// @since 3.6.0
	Implementation/*TOpt*/ *ImplementationClientCapabilities
	// Capabilities specific to the `textDocument/references` request.
	References/*TOpt*/ *ReferenceClientCapabilities
	// Capabilities specific to the `textDocument/documentHighlight` request.
	DocumentHighlight/*TOpt*/ *DocumentHighlightClientCapabilities
	// Capabilities specific to the `textDocument/documentSymbol` request.
	DocumentSymbol/*TOpt*/ *DocumentSymbolClientCapabilities
	// Capabilities specific to the `textDocument/codeAction` request.
	CodeAction/*TOpt*/ *CodeActionClientCapabilities
	// Capabilities specific to the `textDocument/codeLens` request.
	CodeLens/*TOpt*/ *CodeLensClientCapabilities
	// Capabilities specific to the `textDocument/documentLink` request.
	DocumentLink/*TOpt*/ *DocumentLinkClientCapabilities
	// Capabilities specific to the `textDocument/documentColor` and the
	// `textDocument/colorPresentation` request.
	//
	// @since 3.6.0
	ColorProvider/*TOpt*/ *DocumentColorClientCapabilities
	// Capabilities specific to the `textDocument/formatting` request.
	Formatting/*TOpt*/ *DocumentFormattingClientCapabilities
	// Capabilities specific to the `textDocument/rangeFormatting` request.
	RangeFormatting/*TOpt*/ *DocumentRangeFormattingClientCapabilities
	// Capabilities specific to the `textDocument/onTypeFormatting` request.
	OnTypeFormatting/*TOpt*/ *DocumentOnTypeFormattingClientCapabilities
	// Capabilities specific to the `textDocument/rename` request.
	Rename/*TOpt*/ *RenameClientCapabilities
	// Capabilities specific to the `textDocument/foldingRange` request.
	//
	// @since 3.10.0
	FoldingRange/*TOpt*/ *FoldingRangeClientCapabilities
	// Capabilities specific to the `textDocument/selectionRange` request.
	//
	// @since 3.15.0
	SelectionRange/*TOpt*/ *SelectionRangeClientCapabilities
	// Capabilities specific to the `textDocument/publishDiagnostics` notification.
	PublishDiagnostics/*TOpt*/ *PublishDiagnosticsClientCapabilities
	// Capabilities specific to the various call hierarchy requests.
	//
	// @since 3.16.0
	CallHierarchy/*TOpt*/ *CallHierarchyClientCapabilities
	// Capabilities specific to the various semantic token request.
	//
	// @since 3.16.0
	SemanticTokens/*TOpt*/ *SemanticTokensClientCapabilities
	// Capabilities specific to the `textDocument/linkedEditingRange` request.
	//
	// @since 3.16.0
	LinkedEditingRange/*TOpt*/ *LinkedEditingRangeClientCapabilities
	// Client capabilities specific to the `textDocument/moniker` request.
	//
	// @since 3.16.0
	Moniker/*TOpt*/ *MonikerClientCapabilities
	// Capabilities specific to the various type hierarchy requests.
	//
	// @since 3.17.0
	TypeHierarchy/*TOpt*/ *TypeHierarchyClientCapabilities
	// Capabilities specific to the `textDocument/inlineValue` request.
	//
	// @since 3.17.0
	InlineValue/*TOpt*/ *InlineValueClientCapabilities
	// Capabilities specific to the `textDocument/inlayHint` request.
	//
	// @since 3.17.0
	InlayHint/*TOpt*/ *InlayHintClientCapabilities
	// Capabilities specific to the diagnostic pull model.
	//
	// @since 3.17.0
	Diagnostic/*TOpt*/ *DiagnosticClientCapabilities
}

// Capabilities specific to the notebook document support.
//
// @since 3.17.0
type NotebookDocumentClientCapabilities struct {
	// Capabilities specific to notebook document synchronization
	//
	// @since 3.17.0
	Synchronization NotebookDocumentSyncClientCapabilities
}
type WindowClientCapabilities struct {
	// It indicates whether the client supports server initiated
	// progress using the `window/workDoneProgress/create` request.
	//
	// The capability also controls Whether client supports handling
	// of progress notifications. If set servers are allowed to report a
	// `workDoneProgress` property in the request specific server
	// capabilities.
	//
	// @since 3.15.0
	WorkDoneProgress/*TOpt*/ *Boolean
	// Capabilities specific to the showMessage request.
	//
	// @since 3.16.0
	ShowMessage/*TOpt*/ *ShowMessageRequestClientCapabilities
	// Capabilities specific to the showDocument request.
	//
	// @since 3.16.0
	ShowDocument/*TOpt*/ *ShowDocumentClientCapabilities
}

// General client capabilities.
//
// @since 3.16.0
type GeneralClientCapabilities struct {
	// Client capability that signals how the client
	// handles stale requests (e.g. a request
	// for which the client will not process the response
	// anymore since the information is outdated).
	//
	// @since 3.17.0
	StaleRequestSupport/*TOpt*/ * /*TStruc*/ struct {
		// The client will actively cancel the request.
		Cancel bool
		// The list of requests for which the client
		// will retry the request if it receives a
		// response with error code `ContentModified`
		RetryOnContentModified []string
	}

	// Client capabilities specific to regular expressions.
	//
	// @since 3.16.0
	RegularExpressions/*TOpt*/ *RegularExpressionsClientCapabilities
	// Client capabilities specific to the client's markdown parser.
	//
	// @since 3.16.0
	Markdown/*TOpt*/ *MarkdownClientCapabilities
	// The position encodings supported by the client. Client and server
	// have to agree on the same position encoding to ensure that offsets
	// (e.g. character position in a line) are interpreted the same on both
	// sides.
	//
	// To keep the protocol backwards compatible the following applies: if
	// the value 'utf-16' is missing from the array of position encodings
	// servers can assume that the client supports UTF-16. UTF-16 is
	// therefore a mandatory encoding.
	//
	// If omitted it defaults to ['utf-16'].
	//
	// Implementation considerations: since the conversion from one encoding
	// into another requires the content of the file / line the conversion
	// is best done where the file is read which is usually on the server
	// side.
	//
	// @since 3.17.0
	PositionEncodings/*TOpt*/ []PositionEncodingKind
}

// A relative pattern is a helper to construct glob patterns that are matched
// relatively to a base URI. The common value for a `baseUri` is a workspace
// folder root, but it can be another absolute URI as well.
//
// @since 3.17.0
type RelativePattern struct {
	// A workspace folder or a base URI to which this pattern will be matched
	// against relatively.
	//
	// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
	BaseUri/*TOr*/ struct {
		WorkspaceFolder/*TOpt*/ *WorkspaceFolder
		URI/*TOpt*/ *URI
	}

	// The actual glob pattern;
	Pattern Pattern
}
type WorkspaceEditClientCapabilities struct {
	// The client supports versioned document changes in `WorkspaceEdit`s
	DocumentChanges/*TOpt*/ *Boolean
	// The resource operations the client supports. Clients should at least
	// support 'create', 'rename' and 'delete' files and folders.
	//
	// @since 3.13.0
	ResourceOperations/*TOpt*/ []ResourceOperationKind
	// The failure handling strategy of a client if applying the workspace edit
	// fails.
	//
	// @since 3.13.0
	FailureHandling/*TOpt*/ FailureHandlingKind
	// Whether the client normalizes line endings to the client specific
	// setting.
	// If set to `true` the client will normalize line ending characters
	// in a workspace edit to the client-specified new line
	// character.
	//
	// @since 3.16.0
	NormalizesLineEndings/*TOpt*/ *Boolean
	// Whether the client in general supports change annotations on text edits,
	// create file, rename file and delete file changes.
	//
	// @since 3.16.0
	ChangeAnnotationSupport/*TOpt*/ * /*TStruc*/ struct {
		// Whether the client groups edits with equal labels into tree nodes,
		// for instance all edits labelled with "Changes in Strings" would
		// be a tree node.
		GroupsOnLabel /*TOpt*/ *Boolean
	}
}
type DidChangeConfigurationClientCapabilities struct {
	// Did change configuration notification supports dynamic registration.
	DynamicRegistration /*TOpt*/ *Boolean
}
type DidChangeWatchedFilesClientCapabilities struct {
	// Did change watched files notification supports dynamic registration. Please note
	// that the current protocol doesn't support static configuration for file changes
	// from the server side.
	DynamicRegistration/*TOpt*/ *Boolean
	// Whether the client has support for ` RelativePattern relative pattern`
	// or not.
	//
	// @since 3.17.0
	RelativePatternSupport/*TOpt*/ *Boolean
}

// Client capabilities for a `WorkspaceSymbolRequest`.
type WorkspaceSymbolClientCapabilities struct {
	// Symbol request supports dynamic registration.
	DynamicRegistration/*TOpt*/ *Boolean
	// Specific capabilities for the `SymbolKind` in the `workspace/symbol` request.
	SymbolKind/*TOpt*/ * /*TStruc*/ struct {
		// The symbol kind values the client supports. When this
		// property exists the client also guarantees that it will
		// handle values outside its set gracefully and falls back
		// to a default value when unknown.
		//
		// If this property is not present the client only supports
		// the symbol kinds from `File` to `Array` as defined in
		// the initial version of the protocol.
		ValueSet /*TOpt*/ []SymbolKind
	}

	// The client supports tags on `SymbolInformation`.
	// Clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.16.0
	TagSupport/*TOpt*/ * /*TStruc*/ struct {
		// The tags supported by the client.
		ValueSet []SymbolTag
	}

	// The client support partial workspace symbols. The client will send the
	// request `workspaceSymbol/resolve` to the server to resolve additional
	// properties.
	//
	// @since 3.17.0
	ResolveSupport/*TOpt*/ * /*TStruc*/ struct {
		// The properties that a client can resolve lazily. Usually
		// `location.range`
		Properties []string
	}
}

// The client capabilities of a `ExecuteCommandRequest`.
type ExecuteCommandClientCapabilities struct {
	// Execute command supports dynamic registration.
	DynamicRegistration /*TOpt*/ *Boolean
}

// @since 3.16.0
type SemanticTokensWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from
	// the server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// semantic tokens currently shown. It should be used with absolute care
	// and is useful for situation where a server for example detects a project
	// wide change that requires such a calculation.
	RefreshSupport /*TOpt*/ *Boolean
}

// @since 3.16.0
type CodeLensWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from the
	// server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// code lenses currently shown. It should be used with absolute care and is
	// useful for situation where a server for example detect a project wide
	// change that requires such a calculation.
	RefreshSupport /*TOpt*/ *Boolean
}

// Capabilities relating to events from file operations by the user in the client.
//
// These events do not come from the file system, they come from user operations
// like renaming a file in the UI.
//
// @since 3.16.0
type FileOperationClientCapabilities struct {
	// Whether the client supports dynamic registration for file requests/notifications.
	DynamicRegistration/*TOpt*/ *Boolean
	// The client has support for sending didCreateFiles notifications.
	DidCreate/*TOpt*/ *Boolean
	// The client has support for sending willCreateFiles requests.
	WillCreate/*TOpt*/ *Boolean
	// The client has support for sending didRenameFiles notifications.
	DidRename/*TOpt*/ *Boolean
	// The client has support for sending willRenameFiles requests.
	WillRename/*TOpt*/ *Boolean
	// The client has support for sending didDeleteFiles notifications.
	DidDelete/*TOpt*/ *Boolean
	// The client has support for sending willDeleteFiles requests.
	WillDelete/*TOpt*/ *Boolean
}

// Client workspace capabilities specific to inline values.
//
// @since 3.17.0
type InlineValueWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from the
	// server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// inline values currently shown. It should be used with absolute care and is
	// useful for situation where a server for example detects a project wide
	// change that requires such a calculation.
	RefreshSupport /*TOpt*/ *Boolean
}

// Client workspace capabilities specific to inlay hints.
//
// @since 3.17.0
type InlayHintWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from
	// the server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// inlay hints currently shown. It should be used with absolute care and
	// is useful for situation where a server for example detects a project wide
	// change that requires such a calculation.
	RefreshSupport /*TOpt*/ *Boolean
}

// Workspace client capabilities specific to diagnostic pull requests.
//
// @since 3.17.0
type DiagnosticWorkspaceClientCapabilities struct {
	// Whether the client implementation supports a refresh request sent from
	// the server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// pulled diagnostics currently shown. It should be used with absolute care and
	// is useful for situation where a server for example detects a project wide
	// change that requires such a calculation.
	RefreshSupport /*TOpt*/ *Boolean
}
type TextDocumentSyncClientCapabilities struct {
	// Whether text document synchronization supports dynamic registration.
	DynamicRegistration/*TOpt*/ *Boolean
	// The client supports sending will save notifications.
	WillSave/*TOpt*/ *Boolean
	// The client supports sending a will save request and
	// waits for a response providing text edits which will
	// be applied to the document before it is saved.
	WillSaveWaitUntil/*TOpt*/ *Boolean
	// The client supports did save notifications.
	DidSave/*TOpt*/ *Boolean
}

// Completion client capabilities
type CompletionClientCapabilities struct {
	// Whether completion supports dynamic registration.
	DynamicRegistration/*TOpt*/ *Boolean
	// The client supports the following `CompletionItem` specific
	// capabilities.
	CompletionItem/*TOpt*/ * /*TStruc*/ struct {
		// Client supports snippets as insert text.
		//
		// A snippet can define tab stops and placeholders with `$1`, `$2`
		// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
		// the end of the snippet. Placeholders with equal identifiers are linked,
		// that is typing in one will update others too.
		SnippetSupport/*TOpt*/ *Boolean
		// Client supports commit characters on a completion item.
		CommitCharactersSupport/*TOpt*/ *Boolean
		// Client supports the following content formats for the documentation
		// property. The order describes the preferred format of the client.
		DocumentationFormat/*TOpt*/ []MarkupKind
		// Client supports the deprecated property on a completion item.
		DeprecatedSupport/*TOpt*/ *Boolean
		// Client supports the preselect property on a completion item.
		PreselectSupport/*TOpt*/ *Boolean
		// Client supports the tag property on a completion item. Clients supporting
		// tags have to handle unknown tags gracefully. Clients especially need to
		// preserve unknown tags when sending a completion item back to the server in
		// a resolve call.
		//
		// @since 3.15.0
		TagSupport/*TOpt*/ * /*TStruc*/ struct {
			// The tags supported by the client.
			ValueSet []CompletionItemTag
		}

		// Client support insert replace edit to control different behavior if a
		// completion item is inserted in the text or should replace text.
		//
		// @since 3.16.0
		InsertReplaceSupport/*TOpt*/ *Boolean
		// Indicates which properties a client can resolve lazily on a completion
		// item. Before version 3.16.0 only the predefined properties `documentation`
		// and `details` could be resolved lazily.
		//
		// @since 3.16.0
		ResolveSupport/*TOpt*/ * /*TStruc*/ struct {
			// The properties that a client can resolve lazily.
			Properties []string
		}

		// The client supports the `insertTextMode` property on
		// a completion item to override the whitespace handling mode
		// as defined by the client (see `insertTextMode`).
		//
		// @since 3.16.0
		InsertTextModeSupport/*TOpt*/ * /*TStruc*/ struct {
			ValueSet []InsertTextMode
		}

		// The client has support for completion item label
		// details (see also `CompletionItemLabelDetails`).
		//
		// @since 3.17.0
		LabelDetailsSupport/*TOpt*/ *Boolean
	}

	CompletionItemKind/*TOpt*/ * /*TStruc*/ struct {
		// The completion item kind values the client supports. When this
		// property exists the client also guarantees that it will
		// handle values outside its set gracefully and falls back
		// to a default value when unknown.
		//
		// If this property is not present the client only supports
		// the completion items kinds from `Text` to `Reference` as defined in
		// the initial version of the protocol.
		ValueSet /*TOpt*/ []CompletionItemKind
	}

	// Defines how the client handles whitespace and indentation
	// when accepting a completion item that uses multi line
	// text in either `insertText` or `textEdit`.
	//
	// @since 3.17.0
	InsertTextMode/*TOpt*/ InsertTextMode
	// The client supports to send additional context information for a
	// `textDocument/completion` request.
	ContextSupport/*TOpt*/ *Boolean
	// The client supports the following `CompletionList` specific
	// capabilities.
	//
	// @since 3.17.0
	CompletionList/*TOpt*/ * /*TStruc*/ struct {
		// The client supports the following itemDefaults on
		// a completion list.
		//
		// The value lists the supported property names of the
		// `CompletionList.itemDefaults` object. If omitted
		// no properties are supported.
		//
		// @since 3.17.0
		ItemDefaults /*TOpt*/ []string
	}
}
type HoverClientCapabilities struct {
	// Whether hover supports dynamic registration.
	DynamicRegistration/*TOpt*/ *Boolean
	// Client supports the following content formats for the content
	// property. The order describes the preferred format of the client.
	ContentFormat/*TOpt*/ []MarkupKind
}

// Client Capabilities for a `SignatureHelpRequest`.
type SignatureHelpClientCapabilities struct {
	// Whether signature help supports dynamic registration.
	DynamicRegistration/*TOpt*/ *Boolean
	// The client supports the following `SignatureInformation`
	// specific properties.
	SignatureInformation/*TOpt*/ * /*TStruc*/ struct {
		// Client supports the following content formats for the documentation
		// property. The order describes the preferred format of the client.
		DocumentationFormat/*TOpt*/ []MarkupKind
		// Client capabilities specific to parameter information.
		ParameterInformation/*TOpt*/ * /*TStruc*/ struct {
			// The client supports processing label offsets instead of a
			// simple label string.
			//
			// @since 3.14.0
			LabelOffsetSupport /*TOpt*/ *Boolean
		}

		// The client supports the `activeParameter` property on `SignatureInformation`
		// literal.
		//
		// @since 3.16.0
		ActiveParameterSupport/*TOpt*/ *Boolean
	}

	// The client supports to send additional context information for a
	// `textDocument/signatureHelp` request. A client that opts into
	// contextSupport will also support the `retriggerCharacters` on
	// `SignatureHelpOptions`.
	//
	// @since 3.15.0
	ContextSupport/*TOpt*/ *Boolean
}

// @since 3.14.0
type DeclarationClientCapabilities struct {
	// Whether declaration supports dynamic registration. If this is set to `true`
	// the client supports the new `DeclarationRegistrationOptions` return value
	// for the corresponding server capability as well.
	DynamicRegistration/*TOpt*/ *Boolean
	// The client supports additional metadata in the form of declaration links.
	LinkSupport/*TOpt*/ *Boolean
}

// Client Capabilities for a `DefinitionRequest`.
type DefinitionClientCapabilities struct {
	// Whether definition supports dynamic registration.
	DynamicRegistration/*TOpt*/ *Boolean
	// The client supports additional metadata in the form of definition links.
	//
	// @since 3.14.0
	LinkSupport/*TOpt*/ *Boolean
}

// Since 3.6.0
type TypeDefinitionClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `TypeDefinitionRegistrationOptions` return value
	// for the corresponding server capability as well.
	DynamicRegistration/*TOpt*/ *Boolean
	// The client supports additional metadata in the form of definition links.
	//
	// Since 3.14.0
	LinkSupport/*TOpt*/ *Boolean
}

// @since 3.6.0
type ImplementationClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `ImplementationRegistrationOptions` return value
	// for the corresponding server capability as well.
	DynamicRegistration/*TOpt*/ *Boolean
	// The client supports additional metadata in the form of definition links.
	//
	// @since 3.14.0
	LinkSupport/*TOpt*/ *Boolean
}

// Client Capabilities for a `ReferencesRequest`.
type ReferenceClientCapabilities struct {
	// Whether references supports dynamic registration.
	DynamicRegistration /*TOpt*/ *Boolean
}

// Client Capabilities for a `DocumentHighlightRequest`.
type DocumentHighlightClientCapabilities struct {
	// Whether document highlight supports dynamic registration.
	DynamicRegistration /*TOpt*/ *Boolean
}

// Client Capabilities for a `DocumentSymbolRequest`.
type DocumentSymbolClientCapabilities struct {
	// Whether document symbol supports dynamic registration.
	DynamicRegistration/*TOpt*/ *Boolean
	// Specific capabilities for the `SymbolKind` in the
	// `textDocument/documentSymbol` request.
	SymbolKind/*TOpt*/ * /*TStruc*/ struct {
		// The symbol kind values the client supports. When this
		// property exists the client also guarantees that it will
		// handle values outside its set gracefully and falls back
		// to a default value when unknown.
		//
		// If this property is not present the client only supports
		// the symbol kinds from `File` to `Array` as defined in
		// the initial version of the protocol.
		ValueSet /*TOpt*/ []SymbolKind
	}

	// The client supports hierarchical document symbols.
	HierarchicalDocumentSymbolSupport/*TOpt*/ *Boolean
	// The client supports tags on `SymbolInformation`. Tags are supported on
	// `DocumentSymbol` if `hierarchicalDocumentSymbolSupport` is set to true.
	// Clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.16.0
	TagSupport/*TOpt*/ * /*TStruc*/ struct {
		// The tags supported by the client.
		ValueSet []SymbolTag
	}

	// The client supports an additional label presented in the UI when
	// registering a document symbol provider.
	//
	// @since 3.16.0
	LabelSupport/*TOpt*/ *Boolean
}

// The Client Capabilities of a `CodeActionRequest`.
type CodeActionClientCapabilities struct {
	// Whether code action supports dynamic registration.
	DynamicRegistration/*TOpt*/ *Boolean
	// The client support code action literals of type `CodeAction` as a valid
	// response of the `textDocument/codeAction` request. If the property is not
	// set the request can only return `Command` literals.
	//
	// @since 3.8.0
	CodeActionLiteralSupport/*TOpt*/ * /*TStruc*/ struct {
		// The code action kind is support with the following value
		// set.
		CodeActionKind /*TStruc*/ struct {
			// The code action kind values the client supports. When this
			// property exists the client also guarantees that it will
			// handle values outside its set gracefully and falls back
			// to a default value when unknown.
			ValueSet []CodeActionKind
		}
	}

	// Whether code action supports the `isPreferred` property.
	//
	// @since 3.15.0
	IsPreferredSupport/*TOpt*/ *Boolean
	// Whether code action supports the `disabled` property.
	//
	// @since 3.16.0
	DisabledSupport/*TOpt*/ *Boolean
	// Whether code action supports the `data` property which is
	// preserved between a `textDocument/codeAction` and a
	// `codeAction/resolve` request.
	//
	// @since 3.16.0
	DataSupport/*TOpt*/ *Boolean
	// Whether the client supports resolving additional code action
	// properties via a separate `codeAction/resolve` request.
	//
	// @since 3.16.0
	ResolveSupport/*TOpt*/ * /*TStruc*/ struct {
		// The properties that a client can resolve lazily.
		Properties []string
	}

	// Whether the client honors the change annotations in
	// text edits and resource operations returned via the
	// `CodeAction#edit` property by for example presenting
	// the workspace edit in the user interface and asking
	// for confirmation.
	//
	// @since 3.16.0
	HonorsChangeAnnotations/*TOpt*/ *Boolean
}

// The client capabilities  of a `CodeLensRequest`.
type CodeLensClientCapabilities struct {
	// Whether code lens supports dynamic registration.
	DynamicRegistration /*TOpt*/ *Boolean
}

// The client capabilities of a `DocumentLinkRequest`.
type DocumentLinkClientCapabilities struct {
	// Whether document link supports dynamic registration.
	DynamicRegistration/*TOpt*/ *Boolean
	// Whether the client supports the `tooltip` property on `DocumentLink`.
	//
	// @since 3.15.0
	TooltipSupport/*TOpt*/ *Boolean
}
type DocumentColorClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `DocumentColorRegistrationOptions` return value
	// for the corresponding server capability as well.
	DynamicRegistration /*TOpt*/ *Boolean
}

// Client capabilities of a `DocumentFormattingRequest`.
type DocumentFormattingClientCapabilities struct {
	// Whether formatting supports dynamic registration.
	DynamicRegistration /*TOpt*/ *Boolean
}

// Client capabilities of a `DocumentRangeFormattingRequest`.
type DocumentRangeFormattingClientCapabilities struct {
	// Whether range formatting supports dynamic registration.
	DynamicRegistration /*TOpt*/ *Boolean
}

// Client capabilities of a `DocumentOnTypeFormattingRequest`.
type DocumentOnTypeFormattingClientCapabilities struct {
	// Whether on type formatting supports dynamic registration.
	DynamicRegistration /*TOpt*/ *Boolean
}
type RenameClientCapabilities struct {
	// Whether rename supports dynamic registration.
	DynamicRegistration/*TOpt*/ *Boolean
	// Client supports testing for validity of rename operations
	// before execution.
	//
	// @since 3.12.0
	PrepareSupport/*TOpt*/ *Boolean
	// Client supports the default behavior result.
	//
	// The value indicates the default behavior used by the
	// client.
	//
	// @since 3.16.0
	PrepareSupportDefaultBehavior/*TOpt*/ PrepareSupportDefaultBehavior
	// Whether the client honors the change annotations in
	// text edits and resource operations returned via the
	// rename request's workspace edit by for example presenting
	// the workspace edit in the user interface and asking
	// for confirmation.
	//
	// @since 3.16.0
	HonorsChangeAnnotations/*TOpt*/ *Boolean
}
type FoldingRangeClientCapabilities struct {
	// Whether implementation supports dynamic registration for folding range
	// providers. If this is set to `true` the client supports the new
	// `FoldingRangeRegistrationOptions` return value for the corresponding
	// server capability as well.
	DynamicRegistration/*TOpt*/ *Boolean
	// The maximum number of folding ranges that the client prefers to receive
	// per document. The value serves as a hint, servers are free to follow the
	// limit.
	RangeLimit/*TOpt*/ *Uinteger
	// If set, the client signals that it only supports folding complete lines.
	// If set, client will ignore specified `startCharacter` and `endCharacter`
	// properties in a FoldingRange.
	LineFoldingOnly/*TOpt*/ *Boolean
	// Specific options for the folding range kind.
	//
	// @since 3.17.0
	FoldingRangeKind/*TOpt*/ * /*TStruc*/ struct {
		// The folding range kind values the client supports. When this
		// property exists the client also guarantees that it will
		// handle values outside its set gracefully and falls back
		// to a default value when unknown.
		ValueSet /*TOpt*/ []FoldingRangeKind
	}

	// Specific options for the folding range.
	//
	// @since 3.17.0
	FoldingRange/*TOpt*/ * /*TStruc*/ struct {
		// If set, the client signals that it supports setting collapsedText on
		// folding ranges to display custom labels instead of the default text.
		//
		// @since 3.17.0
		CollapsedText /*TOpt*/ *Boolean
	}
}
type SelectionRangeClientCapabilities struct {
	// Whether implementation supports dynamic registration for selection range providers. If this is set to `true`
	// the client supports the new `SelectionRangeRegistrationOptions` return value for the corresponding server
	// capability as well.
	DynamicRegistration /*TOpt*/ *Boolean
}

// The publish diagnostic client capabilities.
type PublishDiagnosticsClientCapabilities struct {
	// Whether the clients accepts diagnostics with related information.
	RelatedInformation/*TOpt*/ *Boolean
	// Client supports the tag property to provide meta data about a diagnostic.
	// Clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.15.0
	TagSupport/*TOpt*/ * /*TStruc*/ struct {
		// The tags supported by the client.
		ValueSet []DiagnosticTag
	}

	// Whether the client interprets the version property of the
	// `textDocument/publishDiagnostics` notification's parameter.
	//
	// @since 3.15.0
	VersionSupport/*TOpt*/ *Boolean
	// Client supports a codeDescription property
	//
	// @since 3.16.0
	CodeDescriptionSupport/*TOpt*/ *Boolean
	// Whether code action supports the `data` property which is
	// preserved between a `textDocument/publishDiagnostics` and
	// `textDocument/codeAction` request.
	//
	// @since 3.16.0
	DataSupport/*TOpt*/ *Boolean
}

// @since 3.16.0
type CallHierarchyClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration /*TOpt*/ *Boolean
}

// @since 3.16.0
type SemanticTokensClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration/*TOpt*/ *Boolean
	// Which requests the client supports and might send to the server
	// depending on the server's capability. Please note that clients might not
	// show semantic tokens or degrade some of the user experience if a range
	// or full request is advertised by the client but not provided by the
	// server. If for example the client capability `requests.full` and
	// `request.range` are both set to true but the server only provides a
	// range provider the client might not render a minimap correctly or might
	// even decide to not show any semantic tokens at all.
	Requests/*TStruc*/ struct {
		// The client will send the `textDocument/semanticTokens/range` request if
		// the server provides a corresponding handler.
		//
		// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
		Range/*TOpt*/ /*TOr*/ struct {
			Boolean/*TOpt*/ *Boolean
			AnyByString/*TOpt*/ map[string]any
		}

		// The client will send the `textDocument/semanticTokens/full` request if
		// the server provides a corresponding handler.
		//
		// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
		Full/*TOpt*/ /*TOr*/ struct {
			Boolean/*TOpt*/ *Boolean
			Delta/*TOpt*/ * /*TStruc*/ struct {
				// The client will send the `textDocument/semanticTokens/full/delta` request if
				// the server provides a corresponding handler.
				Delta /*TOpt*/ *Boolean
			}
		}
	}

	// The token types that the client supports.
	TokenTypes []string
	// The token modifiers that the client supports.
	TokenModifiers []string
	// The token formats the clients supports.
	Formats []TokenFormat
	// Whether the client supports tokens that can overlap each other.
	OverlappingTokenSupport/*TOpt*/ *Boolean
	// Whether the client supports tokens that can span multiple lines.
	MultilineTokenSupport/*TOpt*/ *Boolean
	// Whether the client allows the server to actively cancel a
	// semantic token request, e.g. supports returning
	// LSPErrorCodes.ServerCancelled. If a server does the client
	// needs to retrigger the request.
	//
	// @since 3.17.0
	ServerCancelSupport/*TOpt*/ *Boolean
	// Whether the client uses semantic tokens to augment existing
	// syntax tokens. If set to `true` client side created syntax
	// tokens and semantic tokens are both used for colorization. If
	// set to `false` the client only uses the returned semantic tokens
	// for colorization.
	//
	// If the value is `undefined` then the client behavior is not
	// specified.
	//
	// @since 3.17.0
	AugmentsSyntaxTokens/*TOpt*/ *Boolean
}

// Client capabilities for the linked editing range request.
//
// @since 3.16.0
type LinkedEditingRangeClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration /*TOpt*/ *Boolean
}

// Client capabilities specific to the moniker request.
//
// @since 3.16.0
type MonikerClientCapabilities struct {
	// Whether moniker supports dynamic registration. If this is set to `true`
	// the client supports the new `MonikerRegistrationOptions` return value
	// for the corresponding server capability as well.
	DynamicRegistration /*TOpt*/ *Boolean
}

// @since 3.17.0
type TypeHierarchyClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration /*TOpt*/ *Boolean
}

// Client capabilities specific to inline values.
//
// @since 3.17.0
type InlineValueClientCapabilities struct {
	// Whether implementation supports dynamic registration for inline value providers.
	DynamicRegistration /*TOpt*/ *Boolean
}

// Inlay hint client capabilities.
//
// @since 3.17.0
type InlayHintClientCapabilities struct {
	// Whether inlay hints support dynamic registration.
	DynamicRegistration/*TOpt*/ *Boolean
	// Indicates which properties a client can resolve lazily on an inlay
	// hint.
	ResolveSupport/*TOpt*/ * /*TStruc*/ struct {
		// The properties that a client can resolve lazily.
		Properties []string
	}
}

// Client capabilities specific to diagnostic pull requests.
//
// @since 3.17.0
type DiagnosticClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is set to `true`
	// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration/*TOpt*/ *Boolean
	// Whether the clients supports related documents for document diagnostic pulls.
	RelatedDocumentSupport/*TOpt*/ *Boolean
}

// Notebook specific client capabilities.
//
// @since 3.17.0
type NotebookDocumentSyncClientCapabilities struct {
	// Whether implementation supports dynamic registration. If this is
	// set to `true` the client supports the new
	// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
	// return value for the corresponding server capability as well.
	DynamicRegistration/*TOpt*/ *Boolean
	// The client supports sending execution summary data per cell.
	ExecutionSummarySupport/*TOpt*/ *Boolean
}

// Show message request client capabilities
type ShowMessageRequestClientCapabilities struct {
	// Capabilities specific to the `MessageActionItem` type.
	MessageActionItem /*TOpt*/ * /*TStruc*/ struct {
		// Whether the client supports additional attributes which
		// are preserved and send back to the server in the
		// request's response.
		AdditionalPropertiesSupport /*TOpt*/ *Boolean
	}
}

// Client capabilities for the showDocument request.
//
// @since 3.16.0
type ShowDocumentClientCapabilities struct {
	// The client has support for the showDocument
	// request.
	Support bool
}

// Client capabilities specific to regular expressions.
//
// @since 3.16.0
type RegularExpressionsClientCapabilities struct {
	// The engine's name.
	Engine string
	// The engine's version.
	Version/*TOpt*/ *String
}

// Client capabilities specific to the used markdown parser.
//
// @since 3.16.0
type MarkdownClientCapabilities struct {
	// The name of the parser.
	Parser string
	// The version of the parser.
	Version/*TOpt*/ *String
	// A list of HTML tags that the client allows / supports in
	// Markdown.
	//
	// @since 3.17.0
	AllowedTags/*TOpt*/ []string
}
