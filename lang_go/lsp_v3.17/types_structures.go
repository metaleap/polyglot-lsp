// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-vsx-and-lsp/gen/cmd/gen_lsp
package lsp

type ImplementationParams struct {
}

// Represents a location inside a resource, such as a line
// inside a text file.
type Location struct {
}
type ImplementationRegistrationOptions struct {
}
type TypeDefinitionParams struct {
}
type TypeDefinitionRegistrationOptions struct {
}

// A workspace folder inside a client.
type WorkspaceFolder struct {
}

// The parameters of a `workspace/didChangeWorkspaceFolders` notification.
type DidChangeWorkspaceFoldersParams struct {
}

// The parameters of a configuration request.
type ConfigurationParams struct {
}

// Parameters for a `DocumentColorRequest`.
type DocumentColorParams struct {
}

// Represents a color range from a document.
type ColorInformation struct {
}
type DocumentColorRegistrationOptions struct {
}

// Parameters for a `ColorPresentationRequest`.
type ColorPresentationParams struct {
}
type ColorPresentation struct {
}
type WorkDoneProgressOptions struct {
}

// General text document registration options.
type TextDocumentRegistrationOptions struct {
}

// Parameters for a `FoldingRangeRequest`.
type FoldingRangeParams struct {
}

// Represents a folding range. To be valid, start and end line must be bigger than zero and smaller
// than the number of lines in the document. Clients are free to ignore invalid ranges.
type FoldingRange struct {
}
type FoldingRangeRegistrationOptions struct {
}
type DeclarationParams struct {
}
type DeclarationRegistrationOptions struct {
}

// A parameter literal used in selection range requests.
type SelectionRangeParams struct {
}

// A selection range represents a part of a selection hierarchy. A selection range
// may have a parent selection range that contains it.
type SelectionRange struct {
}
type SelectionRangeRegistrationOptions struct {
}
type WorkDoneProgressCreateParams struct {
}
type WorkDoneProgressCancelParams struct {
}

// The parameter of a `textDocument/prepareCallHierarchy` request.
//
// @since 3.16.0
type CallHierarchyPrepareParams struct {
}

// Represents programming constructs like functions or constructors in the context
// of call hierarchy.
//
// @since 3.16.0
type CallHierarchyItem struct {
}

// Call hierarchy options used during static or dynamic registration.
//
// @since 3.16.0
type CallHierarchyRegistrationOptions struct {
}

// The parameter of a `callHierarchy/incomingCalls` request.
//
// @since 3.16.0
type CallHierarchyIncomingCallsParams struct {
}

// Represents an incoming call, e.g. a caller of a method or constructor.
//
// @since 3.16.0
type CallHierarchyIncomingCall struct {
}

// The parameter of a `callHierarchy/outgoingCalls` request.
//
// @since 3.16.0
type CallHierarchyOutgoingCallsParams struct {
}

// Represents an outgoing call, e.g. calling a getter from a method or a method from a constructor etc.
//
// @since 3.16.0
type CallHierarchyOutgoingCall struct {
}

// @since 3.16.0
type SemanticTokensParams struct {
}

// @since 3.16.0
type SemanticTokens struct {
}

// @since 3.16.0
type SemanticTokensPartialResult struct {
}

// @since 3.16.0
type SemanticTokensRegistrationOptions struct {
}

// @since 3.16.0
type SemanticTokensDeltaParams struct {
}

// @since 3.16.0
type SemanticTokensDelta struct {
}

// @since 3.16.0
type SemanticTokensDeltaPartialResult struct {
}

// @since 3.16.0
type SemanticTokensRangeParams struct {
}

// Params to show a document.
//
// @since 3.16.0
type ShowDocumentParams struct {
}

// The result of a showDocument request.
//
// @since 3.16.0
type ShowDocumentResult struct {
}
type LinkedEditingRangeParams struct {
}

// The result of a linked editing range request.
//
// @since 3.16.0
type LinkedEditingRanges struct {
}
type LinkedEditingRangeRegistrationOptions struct {
}

// The parameters sent in notifications/requests for user-initiated creation of
// files.
//
// @since 3.16.0
type CreateFilesParams struct {
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
}

// The options to register for file operations.
//
// @since 3.16.0
type FileOperationRegistrationOptions struct {
}

// The parameters sent in notifications/requests for user-initiated renames of
// files.
//
// @since 3.16.0
type RenameFilesParams struct {
}

// The parameters sent in notifications/requests for user-initiated deletes of
// files.
//
// @since 3.16.0
type DeleteFilesParams struct {
}
type MonikerParams struct {
}

// Moniker definition to match LSIF 0.5 moniker definition.
//
// @since 3.16.0
type Moniker struct {
}
type MonikerRegistrationOptions struct {
}

// The parameter of a `textDocument/prepareTypeHierarchy` request.
//
// @since 3.17.0
type TypeHierarchyPrepareParams struct {
}

// @since 3.17.0
type TypeHierarchyItem struct {
}

// Type hierarchy options used during static or dynamic registration.
//
// @since 3.17.0
type TypeHierarchyRegistrationOptions struct {
}

// The parameter of a `typeHierarchy/supertypes` request.
//
// @since 3.17.0
type TypeHierarchySupertypesParams struct {
}

// The parameter of a `typeHierarchy/subtypes` request.
//
// @since 3.17.0
type TypeHierarchySubtypesParams struct {
}

// A parameter literal used in inline value requests.
//
// @since 3.17.0
type InlineValueParams struct {
}

// Inline value options used during static or dynamic registration.
//
// @since 3.17.0
type InlineValueRegistrationOptions struct {
}

// A parameter literal used in inlay hint requests.
//
// @since 3.17.0
type InlayHintParams struct {
}

// Inlay hint information.
//
// @since 3.17.0
type InlayHint struct {
}

// Inlay hint options used during static or dynamic registration.
//
// @since 3.17.0
type InlayHintRegistrationOptions struct {
}

// Parameters of the document diagnostic request.
//
// @since 3.17.0
type DocumentDiagnosticParams struct {
}

// A partial result for a document diagnostic report.
//
// @since 3.17.0
type DocumentDiagnosticReportPartialResult struct {
}

// Cancellation data returned from a diagnostic request.
//
// @since 3.17.0
type DiagnosticServerCancellationData struct {
}

// Diagnostic registration options.
//
// @since 3.17.0
type DiagnosticRegistrationOptions struct {
}

// Parameters of the workspace diagnostic request.
//
// @since 3.17.0
type WorkspaceDiagnosticParams struct {
}

// A workspace diagnostic report.
//
// @since 3.17.0
type WorkspaceDiagnosticReport struct {
}

// A partial result for a workspace diagnostic report.
//
// @since 3.17.0
type WorkspaceDiagnosticReportPartialResult struct {
}

// The params sent in an open notebook document notification.
//
// @since 3.17.0
type DidOpenNotebookDocumentParams struct {
}

// The params sent in a change notebook document notification.
//
// @since 3.17.0
type DidChangeNotebookDocumentParams struct {
}

// The params sent in a save notebook document notification.
//
// @since 3.17.0
type DidSaveNotebookDocumentParams struct {
}

// The params sent in a close notebook document notification.
//
// @since 3.17.0
type DidCloseNotebookDocumentParams struct {
}
type RegistrationParams struct {
}
type UnregistrationParams struct {
}
type InitializeParams struct {
}

// The result returned from an initialize request.
type InitializeResult struct {
}

// The data type of the ResponseError if the
// initialize request fails.
type InitializeError struct {
}
type InitializedParams struct {
}

// The parameters of a change configuration notification.
type DidChangeConfigurationParams struct {
}
type DidChangeConfigurationRegistrationOptions struct {
}

// The parameters of a notification message.
type ShowMessageParams struct {
}
type ShowMessageRequestParams struct {
}
type MessageActionItem struct {
}

// The log message parameters.
type LogMessageParams struct {
}

// The parameters sent in an open text document notification
type DidOpenTextDocumentParams struct {
}

// The change text document notification's parameters.
type DidChangeTextDocumentParams struct {
}

// Describe options to be used when registered for text document change events.
type TextDocumentChangeRegistrationOptions struct {
}

// The parameters sent in a close text document notification
type DidCloseTextDocumentParams struct {
}

// The parameters sent in a save text document notification
type DidSaveTextDocumentParams struct {
}

// Save registration options.
type TextDocumentSaveRegistrationOptions struct {
}

// The parameters sent in a will save text document notification.
type WillSaveTextDocumentParams struct {
}

// A text edit applicable to a text document.
type TextEdit struct {
}

// The watched files change notification's parameters.
type DidChangeWatchedFilesParams struct {
}

// Describe options to be used when registered for text document change events.
type DidChangeWatchedFilesRegistrationOptions struct {
}

// The publish diagnostic notification's parameters.
type PublishDiagnosticsParams struct {
}

// Completion parameters
type CompletionParams struct {
}

// A completion item represents a text snippet that is
// proposed to complete text that is being typed.
type CompletionItem struct {
}

// Represents a collection of `CompletionItem` to be presented
// in the editor.
type CompletionList struct {
}

// Registration options for a `CompletionRequest`.
type CompletionRegistrationOptions struct {
}

// Parameters for a `HoverRequest`.
type HoverParams struct {
}

// The result of a hover request.
type Hover struct {
}

// Registration options for a `HoverRequest`.
type HoverRegistrationOptions struct {
}

// Parameters for a `SignatureHelpRequest`.
type SignatureHelpParams struct {
}

// Signature help represents the signature of something
// callable. There can be multiple signature but only one
// active and only one active parameter.
type SignatureHelp struct {
}

// Registration options for a `SignatureHelpRequest`.
type SignatureHelpRegistrationOptions struct {
}

// Parameters for a `DefinitionRequest`.
type DefinitionParams struct {
}

// Registration options for a `DefinitionRequest`.
type DefinitionRegistrationOptions struct {
}

// Parameters for a `ReferencesRequest`.
type ReferenceParams struct {
}

// Registration options for a `ReferencesRequest`.
type ReferenceRegistrationOptions struct {
}

// Parameters for a `DocumentHighlightRequest`.
type DocumentHighlightParams struct {
}

// A document highlight is a range inside a text document which deserves
// special attention. Usually a document highlight is visualized by changing
// the background color of its range.
type DocumentHighlight struct {
}

// Registration options for a `DocumentHighlightRequest`.
type DocumentHighlightRegistrationOptions struct {
}

// Parameters for a `DocumentSymbolRequest`.
type DocumentSymbolParams struct {
}

// Represents information about programming constructs like variables, classes,
// interfaces etc.
type SymbolInformation struct {
}

// Represents programming constructs like variables, classes, interfaces etc.
// that appear in a document. Document symbols can be hierarchical and they
// have two ranges: one that encloses its definition and one that points to
// its most interesting range, e.g. the range of an identifier.
type DocumentSymbol struct {
}

// Registration options for a `DocumentSymbolRequest`.
type DocumentSymbolRegistrationOptions struct {
}

// The parameters of a `CodeActionRequest`.
type CodeActionParams struct {
}

// Represents a reference to a command. Provides a title which
// will be used to represent a command in the UI and, optionally,
// an array of arguments which will be passed to the command handler
// function when invoked.
type Command struct {
}

// A code action represents a change that can be performed in code, e.g. to fix a problem or
// to refactor code.
//
// A CodeAction must set either `edit` and/or a `command`. If both are supplied, the `edit` is applied first, then the `command` is executed.
type CodeAction struct {
}

// Registration options for a `CodeActionRequest`.
type CodeActionRegistrationOptions struct {
}

// The parameters of a `WorkspaceSymbolRequest`.
type WorkspaceSymbolParams struct {
}

// A special workspace symbol that supports locations without a range.
//
// See also SymbolInformation.
//
// @since 3.17.0
type WorkspaceSymbol struct {
}

// Registration options for a `WorkspaceSymbolRequest`.
type WorkspaceSymbolRegistrationOptions struct {
}

// The parameters of a `CodeLensRequest`.
type CodeLensParams struct {
}

// A code lens represents a `Command` that should be shown along with
// source text, like the number of references, a way to run tests, etc.
//
// A code lens is _unresolved_ when no command is associated to it. For performance
// reasons the creation of a code lens and resolving should be done in two stages.
type CodeLens struct {
}

// Registration options for a `CodeLensRequest`.
type CodeLensRegistrationOptions struct {
}

// The parameters of a `DocumentLinkRequest`.
type DocumentLinkParams struct {
}

// A document link is a range in a text document that links to an internal or external resource, like another
// text document or a web site.
type DocumentLink struct {
}

// Registration options for a `DocumentLinkRequest`.
type DocumentLinkRegistrationOptions struct {
}

// The parameters of a `DocumentFormattingRequest`.
type DocumentFormattingParams struct {
}

// Registration options for a `DocumentFormattingRequest`.
type DocumentFormattingRegistrationOptions struct {
}

// The parameters of a `DocumentRangeFormattingRequest`.
type DocumentRangeFormattingParams struct {
}

// Registration options for a `DocumentRangeFormattingRequest`.
type DocumentRangeFormattingRegistrationOptions struct {
}

// The parameters of a `DocumentOnTypeFormattingRequest`.
type DocumentOnTypeFormattingParams struct {
}

// Registration options for a `DocumentOnTypeFormattingRequest`.
type DocumentOnTypeFormattingRegistrationOptions struct {
}

// The parameters of a `RenameRequest`.
type RenameParams struct {
}

// Registration options for a `RenameRequest`.
type RenameRegistrationOptions struct {
}
type PrepareRenameParams struct {
}

// The parameters of a `ExecuteCommandRequest`.
type ExecuteCommandParams struct {
}

// Registration options for a `ExecuteCommandRequest`.
type ExecuteCommandRegistrationOptions struct {
}

// The parameters passed via an apply workspace edit request.
type ApplyWorkspaceEditParams struct {
}

// The result returned from the apply workspace edit request.
//
// @since 3.17 renamed from ApplyWorkspaceEditResponse
type ApplyWorkspaceEditResult struct {
}
type WorkDoneProgressBegin struct {
}
type WorkDoneProgressReport struct {
}
type WorkDoneProgressEnd struct {
}
type SetTraceParams struct {
}
type LogTraceParams struct {
}
type CancelParams struct {
}
type ProgressParams struct {
}

// A parameter literal used in requests to pass a text document and a position inside that
// document.
type TextDocumentPositionParams struct {
}
type WorkDoneProgressParams struct {
}
type PartialResultParams struct {
}

// Represents the connection of two locations. Provides additional metadata over normal `Location`,
// including an origin range.
type LocationLink struct {
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
}
type ImplementationOptions struct {
}

// Static registration options to be returned in the initialize
// request.
type StaticRegistrationOptions struct {
}
type TypeDefinitionOptions struct {
}

// The workspace folder change event.
type WorkspaceFoldersChangeEvent struct {
}
type ConfigurationItem struct {
}

// A literal to identify a text document in the client.
type TextDocumentIdentifier struct {
}

// Represents a color in RGBA space.
type Color struct {
}
type DocumentColorOptions struct {
}
type FoldingRangeOptions struct {
}
type DeclarationOptions struct {
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
}
type SelectionRangeOptions struct {
}

// Call hierarchy options used during static registration.
//
// @since 3.16.0
type CallHierarchyOptions struct {
}

// @since 3.16.0
type SemanticTokensOptions struct {
}

// @since 3.16.0
type SemanticTokensEdit struct {
}
type LinkedEditingRangeOptions struct {
}

// Represents information on a file/folder create.
//
// @since 3.16.0
type FileCreate struct {
}

// Describes textual changes on a text document. A TextDocumentEdit describes all changes
// on a document version Si and after they are applied move the document to version Si+1.
// So the creator of a TextDocumentEdit doesn't need to sort the array of edits or do any
// kind of ordering. However the edits must be non overlapping.
type TextDocumentEdit struct {
}

// Create file operation.
type CreateFile struct {
}

// Rename file operation
type RenameFile struct {
}

// Delete file operation
type DeleteFile struct {
}

// Additional information that describes document changes.
//
// @since 3.16.0
type ChangeAnnotation struct {
}

// A filter to describe in which file operation requests or notifications
// the server is interested in receiving.
//
// @since 3.16.0
type FileOperationFilter struct {
}

// Represents information on a file/folder rename.
//
// @since 3.16.0
type FileRename struct {
}

// Represents information on a file/folder delete.
//
// @since 3.16.0
type FileDelete struct {
}
type MonikerOptions struct {
}

// Type hierarchy options used during static registration.
//
// @since 3.17.0
type TypeHierarchyOptions struct {
}

// @since 3.17.0
type InlineValueContext struct {
}

// Provide inline value as text.
//
// @since 3.17.0
type InlineValueText struct {
}

// Provide inline value through a variable lookup.
// If only a range is specified, the variable name will be extracted from the underlying document.
// An optional variable name can be used to override the extracted name.
//
// @since 3.17.0
type InlineValueVariableLookup struct {
}

// Provide an inline value through an expression evaluation.
// If only a range is specified, the expression will be extracted from the underlying document.
// An optional expression can be used to override the extracted expression.
//
// @since 3.17.0
type InlineValueEvaluatableExpression struct {
}

// Inline value options used during static registration.
//
// @since 3.17.0
type InlineValueOptions struct {
}

// An inlay hint label part allows for interactive and composite labels
// of inlay hints.
//
// @since 3.17.0
type InlayHintLabelPart struct {
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
}

// Inlay hint options used during static registration.
//
// @since 3.17.0
type InlayHintOptions struct {
}

// A full diagnostic report with a set of related documents.
//
// @since 3.17.0
type RelatedFullDocumentDiagnosticReport struct {
}

// An unchanged diagnostic report with a set of related documents.
//
// @since 3.17.0
type RelatedUnchangedDocumentDiagnosticReport struct {
}

// A diagnostic report with a full set of problems.
//
// @since 3.17.0
type FullDocumentDiagnosticReport struct {
}

// A diagnostic report indicating that the last returned
// report is still accurate.
//
// @since 3.17.0
type UnchangedDocumentDiagnosticReport struct {
}

// Diagnostic options.
//
// @since 3.17.0
type DiagnosticOptions struct {
}

// A previous result id in a workspace pull request.
//
// @since 3.17.0
type PreviousResultId struct {
}

// A notebook document.
//
// @since 3.17.0
type NotebookDocument struct {
}

// An item to transfer a text document from the client to the
// server.
type TextDocumentItem struct {
}

// A versioned notebook document identifier.
//
// @since 3.17.0
type VersionedNotebookDocumentIdentifier struct {
}

// A change event for a notebook document.
//
// @since 3.17.0
type NotebookDocumentChangeEvent struct {
}

// A literal to identify a notebook document in the client.
//
// @since 3.17.0
type NotebookDocumentIdentifier struct {
}

// General parameters to register for a notification or to register a provider.
type Registration struct {
}

// General parameters to unregister a request or notification.
type Unregistration struct {
}

// The initialize parameters
type _InitializeParams struct {
}
type WorkspaceFoldersInitializeParams struct {
}

// Defines the capabilities provided by a language
// server.
type ServerCapabilities struct {
}

// A text document identifier to denote a specific version of a text document.
type VersionedTextDocumentIdentifier struct {
}

// Save options.
type SaveOptions struct {
}

// An event describing a file change.
type FileEvent struct {
}
type FileSystemWatcher struct {
}

// Represents a diagnostic, such as a compiler error or warning. Diagnostic objects
// are only valid in the scope of a resource.
type Diagnostic struct {
}

// Contains additional information about the context in which a completion request is triggered.
type CompletionContext struct {
}

// Additional details for a completion item label.
//
// @since 3.17.0
type CompletionItemLabelDetails struct {
}

// A special text edit to provide an insert and a replace operation.
//
// @since 3.16.0
type InsertReplaceEdit struct {
}

// Completion options.
type CompletionOptions struct {
}

// Hover options.
type HoverOptions struct {
}

// Additional information about the context in which a signature help request was triggered.
//
// @since 3.15.0
type SignatureHelpContext struct {
}

// Represents the signature of something callable. A signature
// can have a label, like a function-name, a doc-comment, and
// a set of parameters.
type SignatureInformation struct {
}

// Server Capabilities for a `SignatureHelpRequest`.
type SignatureHelpOptions struct {
}

// Server Capabilities for a `DefinitionRequest`.
type DefinitionOptions struct {
}

// Value-object that contains additional information when
// requesting references.
type ReferenceContext struct {
}

// Reference options.
type ReferenceOptions struct {
}

// Provider options for a `DocumentHighlightRequest`.
type DocumentHighlightOptions struct {
}

// A base for all symbol information.
type BaseSymbolInformation struct {
}

// Provider options for a `DocumentSymbolRequest`.
type DocumentSymbolOptions struct {
}

// Contains additional diagnostic information about the context in which
// a `CodeActionProvider.provideCodeActions` is run.
type CodeActionContext struct {
}

// Provider options for a `CodeActionRequest`.
type CodeActionOptions struct {
}

// Server capabilities for a `WorkspaceSymbolRequest`.
type WorkspaceSymbolOptions struct {
}

// Code Lens provider options of a `CodeLensRequest`.
type CodeLensOptions struct {
}

// Provider options for a `DocumentLinkRequest`.
type DocumentLinkOptions struct {
}

// Value-object describing what options formatting should use.
type FormattingOptions struct {
}

// Provider options for a `DocumentFormattingRequest`.
type DocumentFormattingOptions struct {
}

// Provider options for a `DocumentRangeFormattingRequest`.
type DocumentRangeFormattingOptions struct {
}

// Provider options for a `DocumentOnTypeFormattingRequest`.
type DocumentOnTypeFormattingOptions struct {
}

// Provider options for a `RenameRequest`.
type RenameOptions struct {
}

// The server capabilities of a `ExecuteCommandRequest`.
type ExecuteCommandOptions struct {
}

// @since 3.16.0
type SemanticTokensLegend struct {
}

// A text document identifier to optionally denote a specific version of a text document.
type OptionalVersionedTextDocumentIdentifier struct {
}

// A special text edit with an additional change annotation.
//
// @since 3.16.0.
type AnnotatedTextEdit struct {
}

// A generic resource operation.
type ResourceOperation struct {
}

// Options to create a file.
type CreateFileOptions struct {
}

// Rename file options
type RenameFileOptions struct {
}

// Delete file options
type DeleteFileOptions struct {
}

// A pattern to describe in which file operation requests or notifications
// the server is interested in receiving.
//
// @since 3.16.0
type FileOperationPattern struct {
}

// A full document diagnostic report for a workspace diagnostic result.
//
// @since 3.17.0
type WorkspaceFullDocumentDiagnosticReport struct {
}

// An unchanged document diagnostic report for a workspace diagnostic result.
//
// @since 3.17.0
type WorkspaceUnchangedDocumentDiagnosticReport struct {
}

// A notebook cell.
//
// A cell's document URI must be unique across ALL notebook
// cells and can therefore be used to uniquely identify a
// notebook cell or the cell's text document.
//
// @since 3.17.0
type NotebookCell struct {
}

// A change describing how to move a `NotebookCell`
// array from state S to S'.
//
// @since 3.17.0
type NotebookCellArrayChange struct {
}

// Defines the capabilities provided by the client.
type ClientCapabilities struct {
}
type TextDocumentSyncOptions struct {
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
}

// Registration options specific to a notebook.
//
// @since 3.17.0
type NotebookDocumentSyncRegistrationOptions struct {
}
type WorkspaceFoldersServerCapabilities struct {
}

// Options for notifications/requests for user operations on files.
//
// @since 3.16.0
type FileOperationOptions struct {
}

// Structure to capture a description for an error code.
//
// @since 3.16.0
type CodeDescription struct {
}

// Represents a related message and source code location for a diagnostic. This should be
// used to point to code locations that cause or related to a diagnostics, e.g when duplicating
// a symbol in a scope.
type DiagnosticRelatedInformation struct {
}

// Represents a parameter of a callable-signature. A parameter can
// have a label and a doc-comment.
type ParameterInformation struct {
}

// A notebook cell text document filter denotes a cell text
// document by different properties.
//
// @since 3.17.0
type NotebookCellTextDocumentFilter struct {
}

// Matching options for the file operation pattern.
//
// @since 3.16.0
type FileOperationPatternOptions struct {
}
type ExecutionSummary struct {
}

// Workspace specific client capabilities.
type WorkspaceClientCapabilities struct {
}

// Text document specific client capabilities.
type TextDocumentClientCapabilities struct {
}

// Capabilities specific to the notebook document support.
//
// @since 3.17.0
type NotebookDocumentClientCapabilities struct {
}
type WindowClientCapabilities struct {
}

// General client capabilities.
//
// @since 3.16.0
type GeneralClientCapabilities struct {
}

// A relative pattern is a helper to construct glob patterns that are matched
// relatively to a base URI. The common value for a `baseUri` is a workspace
// folder root, but it can be another absolute URI as well.
//
// @since 3.17.0
type RelativePattern struct {
}
type WorkspaceEditClientCapabilities struct {
}
type DidChangeConfigurationClientCapabilities struct {
}
type DidChangeWatchedFilesClientCapabilities struct {
}

// Client capabilities for a `WorkspaceSymbolRequest`.
type WorkspaceSymbolClientCapabilities struct {
}

// The client capabilities of a `ExecuteCommandRequest`.
type ExecuteCommandClientCapabilities struct {
}

// @since 3.16.0
type SemanticTokensWorkspaceClientCapabilities struct {
}

// @since 3.16.0
type CodeLensWorkspaceClientCapabilities struct {
}

// Capabilities relating to events from file operations by the user in the client.
//
// These events do not come from the file system, they come from user operations
// like renaming a file in the UI.
//
// @since 3.16.0
type FileOperationClientCapabilities struct {
}

// Client workspace capabilities specific to inline values.
//
// @since 3.17.0
type InlineValueWorkspaceClientCapabilities struct {
}

// Client workspace capabilities specific to inlay hints.
//
// @since 3.17.0
type InlayHintWorkspaceClientCapabilities struct {
}

// Workspace client capabilities specific to diagnostic pull requests.
//
// @since 3.17.0
type DiagnosticWorkspaceClientCapabilities struct {
}
type TextDocumentSyncClientCapabilities struct {
}

// Completion client capabilities
type CompletionClientCapabilities struct {
}
type HoverClientCapabilities struct {
}

// Client Capabilities for a `SignatureHelpRequest`.
type SignatureHelpClientCapabilities struct {
}

// @since 3.14.0
type DeclarationClientCapabilities struct {
}

// Client Capabilities for a `DefinitionRequest`.
type DefinitionClientCapabilities struct {
}

// Since 3.6.0
type TypeDefinitionClientCapabilities struct {
}

// @since 3.6.0
type ImplementationClientCapabilities struct {
}

// Client Capabilities for a `ReferencesRequest`.
type ReferenceClientCapabilities struct {
}

// Client Capabilities for a `DocumentHighlightRequest`.
type DocumentHighlightClientCapabilities struct {
}

// Client Capabilities for a `DocumentSymbolRequest`.
type DocumentSymbolClientCapabilities struct {
}

// The Client Capabilities of a `CodeActionRequest`.
type CodeActionClientCapabilities struct {
}

// The client capabilities  of a `CodeLensRequest`.
type CodeLensClientCapabilities struct {
}

// The client capabilities of a `DocumentLinkRequest`.
type DocumentLinkClientCapabilities struct {
}
type DocumentColorClientCapabilities struct {
}

// Client capabilities of a `DocumentFormattingRequest`.
type DocumentFormattingClientCapabilities struct {
}

// Client capabilities of a `DocumentRangeFormattingRequest`.
type DocumentRangeFormattingClientCapabilities struct {
}

// Client capabilities of a `DocumentOnTypeFormattingRequest`.
type DocumentOnTypeFormattingClientCapabilities struct {
}
type RenameClientCapabilities struct {
}
type FoldingRangeClientCapabilities struct {
}
type SelectionRangeClientCapabilities struct {
}

// The publish diagnostic client capabilities.
type PublishDiagnosticsClientCapabilities struct {
}

// @since 3.16.0
type CallHierarchyClientCapabilities struct {
}

// @since 3.16.0
type SemanticTokensClientCapabilities struct {
}

// Client capabilities for the linked editing range request.
//
// @since 3.16.0
type LinkedEditingRangeClientCapabilities struct {
}

// Client capabilities specific to the moniker request.
//
// @since 3.16.0
type MonikerClientCapabilities struct {
}

// @since 3.17.0
type TypeHierarchyClientCapabilities struct {
}

// Client capabilities specific to inline values.
//
// @since 3.17.0
type InlineValueClientCapabilities struct {
}

// Inlay hint client capabilities.
//
// @since 3.17.0
type InlayHintClientCapabilities struct {
}

// Client capabilities specific to diagnostic pull requests.
//
// @since 3.17.0
type DiagnosticClientCapabilities struct {
}

// Notebook specific client capabilities.
//
// @since 3.17.0
type NotebookDocumentSyncClientCapabilities struct {
}

// Show message request client capabilities
type ShowMessageRequestClientCapabilities struct {
}

// Client capabilities for the showDocument request.
//
// @since 3.16.0
type ShowDocumentClientCapabilities struct {
}

// Client capabilities specific to regular expressions.
//
// @since 3.16.0
type RegularExpressionsClientCapabilities struct {
}

// Client capabilities specific to the used markdown parser.
//
// @since 3.16.0
type MarkdownClientCapabilities struct {
}
