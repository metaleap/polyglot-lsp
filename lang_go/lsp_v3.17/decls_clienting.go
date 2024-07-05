// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-lsp/gen/cmd/gen_lsp
package lsp_v3_17

import (
	"io"
)

type Client struct {
	clientServerBase

	// The show message notification is sent from a server to a client to ask
	// the client to display a particular message in the user interface.
	On_window_showMessage func(params *ShowMessageParams) (any, error)

	// The log message notification is sent from the server to the client to ask
	// the client to log a particular message.
	On_window_logMessage func(params *LogMessageParams) (any, error)

	// The telemetry event notification is sent from the server to the client to ask
	// the client to log telemetry data.
	On_telemetry_event func(params *LSPAny) (any, error)

	// Diagnostics notification are sent from the server to the client to signal
	// results of validation runs.
	On_textDocument_publishDiagnostics func(params *PublishDiagnosticsParams) (any, error)

	//
	On___logTrace func(params *LogTraceParams) (any, error)

	//
	On___cancelRequest func(params *CancelParams) (any, error)

	//
	On___progress func(params *ProgressParams) (any, error)

	// The `workspace/workspaceFolders` is sent from the server to the client to fetch the open workspace folders.
	On_workspace_workspaceFolders func(params *Void) (any, error)

	// The 'workspace/configuration' request is sent from the server to the client to fetch a certain
	// configuration setting.
	//
	// This pull model replaces the old push model where the client signaled configuration change via an
	// event. If the server still needs to react to configuration changes (since the server caches the
	// result of `workspace/configuration` requests) the server should register for an empty configuration
	// change event and empty the cache if such an event is received.
	On_workspace_configuration func(params *ConfigurationParams) (any, error)

	// The `window/workDoneProgress/create` request is sent from the server to the client to initiate progress
	// reporting from the server.
	On_window_workDoneProgress_create func(params *WorkDoneProgressCreateParams) (any, error)

	// @since 3.16.0
	On_workspace_semanticTokens_refresh func(params *Void) (any, error)

	// A request to show a document. This request might open an
	// external program depending on the value of the URI to open.
	// For example a request to open `https://code.visualstudio.com/`
	// will very likely open the URI in a WEB browser.
	//
	// @since 3.16.0
	On_window_showDocument func(params *ShowDocumentParams) (any, error)

	// @since 3.17.0
	On_workspace_inlineValue_refresh func(params *Void) (any, error)

	// @since 3.17.0
	On_workspace_inlayHint_refresh func(params *Void) (any, error)

	// The diagnostic refresh request definition.
	//
	// @since 3.17.0
	On_workspace_diagnostic_refresh func(params *Void) (any, error)

	// The `client/registerCapability` request is sent from the server to the client to register a new capability
	// handler on the client side.
	On_client_registerCapability func(params *RegistrationParams) (any, error)

	// The `client/unregisterCapability` request is sent from the server to the client to unregister a previously registered capability
	// handler on the client side.
	On_client_unregisterCapability func(params *UnregistrationParams) (any, error)

	// The show message request is sent from the server to the client to show a message
	// and a set of options actions to the user.
	On_window_showMessageRequest func(params *ShowMessageRequestParams) (any, error)

	// A request to refresh all code actions
	//
	// @since 3.16.0
	On_workspace_codeLens_refresh func(params *Void) (any, error)

	// A request sent from the server to the client to modified certain resources.
	On_workspace_applyEdit func(params *ApplyWorkspaceEditParams) (any, error)
}

// The `workspace/didChangeWorkspaceFolders` notification is sent from the client to the server when the workspace
// folder configuration changes.
func (it *Client) Notify_workspace_didChangeWorkspaceFolders(params *DidChangeWorkspaceFoldersParams) {
	var on_resp func(any)
	go it.send("workspace/didChangeWorkspaceFolders", params, false, on_resp)
}

// The `window/workDoneProgress/cancel` notification is sent from  the client to the server to cancel a progress
// initiated on the server side.
func (it *Client) Notify_window_workDoneProgress_cancel(params *WorkDoneProgressCancelParams) {
	var on_resp func(any)
	go it.send("window/workDoneProgress/cancel", params, false, on_resp)
}

// The did create files notification is sent from the client to the server when
// files were created from within the client.
//
// @since 3.16.0
func (it *Client) Notify_workspace_didCreateFiles(params *CreateFilesParams) {
	var on_resp func(any)
	go it.send("workspace/didCreateFiles", params, false, on_resp)
}

// The did rename files notification is sent from the client to the server when
// files were renamed from within the client.
//
// @since 3.16.0
func (it *Client) Notify_workspace_didRenameFiles(params *RenameFilesParams) {
	var on_resp func(any)
	go it.send("workspace/didRenameFiles", params, false, on_resp)
}

// The will delete files request is sent from the client to the server before files are actually
// deleted as long as the deletion is triggered from within the client.
//
// @since 3.16.0
func (it *Client) Notify_workspace_didDeleteFiles(params *DeleteFilesParams) {
	var on_resp func(any)
	go it.send("workspace/didDeleteFiles", params, false, on_resp)
}

// A notification sent when a notebook opens.
//
// @since 3.17.0
func (it *Client) Notify_notebookDocument_didOpen(params *DidOpenNotebookDocumentParams) {
	var on_resp func(any)
	go it.send("notebookDocument/didOpen", params, false, on_resp)
}

func (it *Client) Notify_notebookDocument_didChange(params *DidChangeNotebookDocumentParams) {
	var on_resp func(any)
	go it.send("notebookDocument/didChange", params, false, on_resp)
}

// A notification sent when a notebook document is saved.
//
// @since 3.17.0
func (it *Client) Notify_notebookDocument_didSave(params *DidSaveNotebookDocumentParams) {
	var on_resp func(any)
	go it.send("notebookDocument/didSave", params, false, on_resp)
}

// A notification sent when a notebook closes.
//
// @since 3.17.0
func (it *Client) Notify_notebookDocument_didClose(params *DidCloseNotebookDocumentParams) {
	var on_resp func(any)
	go it.send("notebookDocument/didClose", params, false, on_resp)
}

// The initialized notification is sent from the client to the
// server after the client is fully initialized and the server
// is allowed to send requests from the server to the client.
func (it *Client) Notify_initialized(params *InitializedParams) {
	var on_resp func(any)
	go it.send("initialized", params, false, on_resp)
}

// The exit event is sent from the client to the server to
// ask the server to exit its process.
func (it *Client) Notify_exit(params *Void) {
	var on_resp func(any)
	go it.send("exit", params, false, on_resp)
}

// The configuration change notification is sent from the client to the server
// when the client's configuration has changed. The notification contains
// the changed configuration as defined by the language client.
func (it *Client) Notify_workspace_didChangeConfiguration(params *DidChangeConfigurationParams) {
	var on_resp func(any)
	go it.send("workspace/didChangeConfiguration", params, false, on_resp)
}

// The document open notification is sent from the client to the server to signal
// newly opened text documents. The document's truth is now managed by the client
// and the server must not try to read the document's truth using the document's
// uri. Open in this sense means it is managed by the client. It doesn't necessarily
// mean that its content is presented in an editor. An open notification must not
// be sent more than once without a corresponding close notification send before.
// This means open and close notification must be balanced and the max open count
// is one.
func (it *Client) Notify_textDocument_didOpen(params *DidOpenTextDocumentParams) {
	var on_resp func(any)
	go it.send("textDocument/didOpen", params, false, on_resp)
}

// The document change notification is sent from the client to the server to signal
// changes to a text document.
func (it *Client) Notify_textDocument_didChange(params *DidChangeTextDocumentParams) {
	var on_resp func(any)
	go it.send("textDocument/didChange", params, false, on_resp)
}

// The document close notification is sent from the client to the server when
// the document got closed in the client. The document's truth now exists where
// the document's uri points to (e.g. if the document's uri is a file uri the
// truth now exists on disk). As with the open notification the close notification
// is about managing the document's content. Receiving a close notification
// doesn't mean that the document was open in an editor before. A close
// notification requires a previous open notification to be sent.
func (it *Client) Notify_textDocument_didClose(params *DidCloseTextDocumentParams) {
	var on_resp func(any)
	go it.send("textDocument/didClose", params, false, on_resp)
}

// The document save notification is sent from the client to the server when
// the document got saved in the client.
func (it *Client) Notify_textDocument_didSave(params *DidSaveTextDocumentParams) {
	var on_resp func(any)
	go it.send("textDocument/didSave", params, false, on_resp)
}

// A document will save notification is sent from the client to the server before
// the document is actually saved.
func (it *Client) Notify_textDocument_willSave(params *WillSaveTextDocumentParams) {
	var on_resp func(any)
	go it.send("textDocument/willSave", params, false, on_resp)
}

// The watched files notification is sent from the client to the server when
// the client detects changes to file watched by the language client.
func (it *Client) Notify_workspace_didChangeWatchedFiles(params *DidChangeWatchedFilesParams) {
	var on_resp func(any)
	go it.send("workspace/didChangeWatchedFiles", params, false, on_resp)
}

func (it *Client) Notify___setTrace(params *SetTraceParams) {
	var on_resp func(any)
	go it.send("$/setTrace", params, false, on_resp)
}

func (it *Client) Notify___cancelRequest(params *CancelParams) {
	var on_resp func(any)
	go it.send("$/cancelRequest", params, false, on_resp)
}

func (it *Client) Notify___progress(params *ProgressParams) {
	var on_resp func(any)
	go it.send("$/progress", params, false, on_resp)
}

// A request to resolve the implementation locations of a symbol at a given text
// document position. The request's parameter is of type {@link TextDocumentPositionParams}
// the response is of type {@link Definition} or a Thenable that resolves to such.
func (it *Client) Request_textDocument_implementation(params *ImplementationParams, onResp func(*DefinitionOrDefinitionLinksOrNull)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/implementation", params, true, on_resp)
}

// A request to resolve the type definition locations of a symbol at a given text
// document position. The request's parameter is of type {@link TextDocumentPositionParams}
// the response is of type {@link Definition} or a Thenable that resolves to such.
func (it *Client) Request_textDocument_typeDefinition(params *TypeDefinitionParams, onResp func(*DefinitionOrDefinitionLinksOrNull)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/typeDefinition", params, true, on_resp)
}

// A request to list all color symbols found in a given text document. The request's
// parameter is of type {@link DocumentColorParams} the
// response is of type {@link ColorInformation ColorInformation[]} or a Thenable
// that resolves to such.
func (it *Client) Request_textDocument_documentColor(params *DocumentColorParams, onResp func([]ColorInformation)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/documentColor", params, true, on_resp)
}

// A request to list all presentation for a color. The request's
// parameter is of type {@link ColorPresentationParams} the
// response is of type {@link ColorInformation ColorInformation[]} or a Thenable
// that resolves to such.
func (it *Client) Request_textDocument_colorPresentation(params *ColorPresentationParams, onResp func([]ColorPresentation)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/colorPresentation", params, true, on_resp)
}

// A request to provide folding ranges in a document. The request's
// parameter is of type {@link FoldingRangeParams}, the
// response is of type {@link FoldingRangeList} or a Thenable
// that resolves to such.
func (it *Client) Request_textDocument_foldingRange(params *FoldingRangeParams, onResp func(* /*TOr*/ []FoldingRange)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/foldingRange", params, true, on_resp)
}

// A request to resolve the type definition locations of a symbol at a given text
// document position. The request's parameter is of type {@link TextDocumentPositionParams}
// the response is of type {@link Declaration} or a typed array of {@link DeclarationLink}
// or a Thenable that resolves to such.
func (it *Client) Request_textDocument_declaration(params *DeclarationParams, onResp func(*DeclarationOrDeclarationLinksOrNull)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/declaration", params, true, on_resp)
}

// A request to provide selection ranges in a document. The request's
// parameter is of type {@link SelectionRangeParams}, the
// response is of type {@link SelectionRange SelectionRange[]} or a Thenable
// that resolves to such.
func (it *Client) Request_textDocument_selectionRange(params *SelectionRangeParams, onResp func(* /*TOr*/ []SelectionRange)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/selectionRange", params, true, on_resp)
}

// A request to result a `CallHierarchyItem` in a document at a given position.
// Can be used as an input to an incoming or outgoing call hierarchy.
//
// @since 3.16.0
func (it *Client) Request_textDocument_prepareCallHierarchy(params *CallHierarchyPrepareParams, onResp func(* /*TOr*/ []CallHierarchyItem)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/prepareCallHierarchy", params, true, on_resp)
}

// A request to resolve the incoming calls for a given `CallHierarchyItem`.
//
// @since 3.16.0
func (it *Client) Request_callHierarchy_incomingCalls(params *CallHierarchyIncomingCallsParams, onResp func(* /*TOr*/ []CallHierarchyIncomingCall)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("callHierarchy/incomingCalls", params, true, on_resp)
}

// A request to resolve the outgoing calls for a given `CallHierarchyItem`.
//
// @since 3.16.0
func (it *Client) Request_callHierarchy_outgoingCalls(params *CallHierarchyOutgoingCallsParams, onResp func(* /*TOr*/ []CallHierarchyOutgoingCall)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("callHierarchy/outgoingCalls", params, true, on_resp)
}

// @since 3.16.0
func (it *Client) Request_textDocument_semanticTokens_full(params *SemanticTokensParams, onResp func(* /*TOr*/ /*TOpt*/ *SemanticTokens)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/semanticTokens/full", params, true, on_resp)
}

// @since 3.16.0
func (it *Client) Request_textDocument_semanticTokens_full_delta(params *SemanticTokensDeltaParams, onResp func(*SemanticTokensOrSemanticTokensDeltaOrNull)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/semanticTokens/full/delta", params, true, on_resp)
}

// @since 3.16.0
func (it *Client) Request_textDocument_semanticTokens_range(params *SemanticTokensRangeParams, onResp func(* /*TOr*/ /*TOpt*/ *SemanticTokens)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/semanticTokens/range", params, true, on_resp)
}

// A request to provide ranges that can be edited together.
//
// @since 3.16.0
func (it *Client) Request_textDocument_linkedEditingRange(params *LinkedEditingRangeParams, onResp func(* /*TOr*/ /*TOpt*/ *LinkedEditingRanges)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/linkedEditingRange", params, true, on_resp)
}

// The will create files request is sent from the client to the server before files are actually
// created as long as the creation is triggered from within the client.
//
// The request can return a `WorkspaceEdit` which will be applied to workspace before the
// files are created. Hence the `WorkspaceEdit` can not manipulate the content of the file
// to be created.
//
// @since 3.16.0
func (it *Client) Request_workspace_willCreateFiles(params *CreateFilesParams, onResp func(* /*TOr*/ /*TOpt*/ *WorkspaceEdit)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("workspace/willCreateFiles", params, true, on_resp)
}

// The will rename files request is sent from the client to the server before files are actually
// renamed as long as the rename is triggered from within the client.
//
// @since 3.16.0
func (it *Client) Request_workspace_willRenameFiles(params *RenameFilesParams, onResp func(* /*TOr*/ /*TOpt*/ *WorkspaceEdit)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("workspace/willRenameFiles", params, true, on_resp)
}

// The did delete files notification is sent from the client to the server when
// files were deleted from within the client.
//
// @since 3.16.0
func (it *Client) Request_workspace_willDeleteFiles(params *DeleteFilesParams, onResp func(* /*TOr*/ /*TOpt*/ *WorkspaceEdit)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("workspace/willDeleteFiles", params, true, on_resp)
}

// A request to get the moniker of a symbol at a given text document position.
// The request parameter is of type {@link TextDocumentPositionParams}.
// The response is of type {@link Moniker Moniker[]} or `null`.
func (it *Client) Request_textDocument_moniker(params *MonikerParams, onResp func(* /*TOr*/ []Moniker)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/moniker", params, true, on_resp)
}

// A request to result a `TypeHierarchyItem` in a document at a given position.
// Can be used as an input to a subtypes or supertypes type hierarchy.
//
// @since 3.17.0
func (it *Client) Request_textDocument_prepareTypeHierarchy(params *TypeHierarchyPrepareParams, onResp func(* /*TOr*/ []TypeHierarchyItem)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/prepareTypeHierarchy", params, true, on_resp)
}

// A request to resolve the supertypes for a given `TypeHierarchyItem`.
//
// @since 3.17.0
func (it *Client) Request_typeHierarchy_supertypes(params *TypeHierarchySupertypesParams, onResp func(* /*TOr*/ []TypeHierarchyItem)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("typeHierarchy/supertypes", params, true, on_resp)
}

// A request to resolve the subtypes for a given `TypeHierarchyItem`.
//
// @since 3.17.0
func (it *Client) Request_typeHierarchy_subtypes(params *TypeHierarchySubtypesParams, onResp func(* /*TOr*/ []TypeHierarchyItem)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("typeHierarchy/subtypes", params, true, on_resp)
}

// A request to provide inline values in a document. The request's parameter is of
// type {@link InlineValueParams}, the response is of type
// {@link InlineValue InlineValue[]} or a Thenable that resolves to such.
//
// @since 3.17.0
func (it *Client) Request_textDocument_inlineValue(params *InlineValueParams, onResp func(* /*TOr*/ []InlineValue)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/inlineValue", params, true, on_resp)
}

// A request to provide inlay hints in a document. The request's parameter is of
// type {@link InlayHintsParams}, the response is of type
// {@link InlayHint InlayHint[]} or a Thenable that resolves to such.
//
// @since 3.17.0
func (it *Client) Request_textDocument_inlayHint(params *InlayHintParams, onResp func(* /*TOr*/ []InlayHint)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/inlayHint", params, true, on_resp)
}

// A request to resolve additional properties for an inlay hint.
// The request's parameter is of type {@link InlayHint}, the response is
// of type {@link InlayHint} or a Thenable that resolves to such.
//
// @since 3.17.0
func (it *Client) Request_inlayHint_resolve(params *InlayHint, onResp func(InlayHint)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("inlayHint/resolve", params, true, on_resp)
}

// The document diagnostic request definition.
//
// @since 3.17.0
func (it *Client) Request_textDocument_diagnostic(params *DocumentDiagnosticParams, onResp func(DocumentDiagnosticReport)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/diagnostic", params, true, on_resp)
}

// The workspace diagnostic request definition.
//
// @since 3.17.0
func (it *Client) Request_workspace_diagnostic(params *WorkspaceDiagnosticParams, onResp func(WorkspaceDiagnosticReport)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("workspace/diagnostic", params, true, on_resp)
}

// The initialize request is sent from the client to the server.
// It is sent once as the request after starting up the server.
// The requests parameter is of type {@link InitializeParams}
// the response if of type {@link InitializeResult} of a Thenable that
// resolves to such.
func (it *Client) Request_initialize(params *InitializeParams, onResp func(InitializeResult)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("initialize", params, true, on_resp)
}

// A shutdown request is sent from the client to the server.
// It is sent once when the client decides to shutdown the
// server. The only notification that is sent after a shutdown request
// is the exit event.
func (it *Client) Request_shutdown(params *Void, onResp func(Void)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("shutdown", params, true, on_resp)
}

// A document will save request is sent from the client to the server before
// the document is actually saved. The request can return an array of TextEdits
// which will be applied to the text document before it is saved. Please note that
// clients might drop results if computing the text edits took too long or if a
// server constantly fails on this request. This is done to keep the save fast and
// reliable.
func (it *Client) Request_textDocument_willSaveWaitUntil(params *WillSaveTextDocumentParams, onResp func(* /*TOr*/ []TextEdit)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/willSaveWaitUntil", params, true, on_resp)
}

// Request to request completion at a given text document position. The request's
// parameter is of type {@link TextDocumentPosition} the response
// is of type {@link CompletionItem CompletionItem[]} or {@link CompletionList}
// or a Thenable that resolves to such.
//
// The request can delay the computation of the {@link CompletionItem.detail `detail`}
// and {@link CompletionItem.documentation `documentation`} properties to the `completionItem/resolve`
// request. However, properties that are needed for the initial sorting and filtering, like `sortText`,
// `filterText`, `insertText`, and `textEdit`, must not be changed during resolve.
func (it *Client) Request_textDocument_completion(params *CompletionParams, onResp func(*CompletionItemsOrCompletionListOrNull)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/completion", params, true, on_resp)
}

// Request to resolve additional information for a given completion item.The request's
// parameter is of type {@link CompletionItem} the response
// is of type {@link CompletionItem} or a Thenable that resolves to such.
func (it *Client) Request_completionItem_resolve(params *CompletionItem, onResp func(CompletionItem)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("completionItem/resolve", params, true, on_resp)
}

// Request to request hover information at a given text document position. The request's
// parameter is of type {@link TextDocumentPosition} the response is of
// type {@link Hover} or a Thenable that resolves to such.
func (it *Client) Request_textDocument_hover(params *HoverParams, onResp func(* /*TOr*/ /*TOpt*/ *Hover)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/hover", params, true, on_resp)
}

func (it *Client) Request_textDocument_signatureHelp(params *SignatureHelpParams, onResp func(* /*TOr*/ /*TOpt*/ *SignatureHelp)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/signatureHelp", params, true, on_resp)
}

// A request to resolve the definition location of a symbol at a given text
// document position. The request's parameter is of type {@link TextDocumentPosition}
// the response is of either type {@link Definition} or a typed array of
// {@link DefinitionLink} or a Thenable that resolves to such.
func (it *Client) Request_textDocument_definition(params *DefinitionParams, onResp func(*DefinitionOrDefinitionLinksOrNull)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/definition", params, true, on_resp)
}

// A request to resolve project-wide references for the symbol denoted
// by the given text document position. The request's parameter is of
// type {@link ReferenceParams} the response is of type
// {@link Location Location[]} or a Thenable that resolves to such.
func (it *Client) Request_textDocument_references(params *ReferenceParams, onResp func(* /*TOr*/ []Location)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/references", params, true, on_resp)
}

// Request to resolve a {@link DocumentHighlight} for a given
// text document position. The request's parameter is of type {@link TextDocumentPosition}
// the request response is an array of type {@link DocumentHighlight}
// or a Thenable that resolves to such.
func (it *Client) Request_textDocument_documentHighlight(params *DocumentHighlightParams, onResp func(* /*TOr*/ []DocumentHighlight)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/documentHighlight", params, true, on_resp)
}

// A request to list all symbols found in a given text document. The request's
// parameter is of type {@link TextDocumentIdentifier} the
// response is of type {@link SymbolInformation SymbolInformation[]} or a Thenable
// that resolves to such.
func (it *Client) Request_textDocument_documentSymbol(params *DocumentSymbolParams, onResp func(*SymbolInformationsOrDocumentSymbolsOrNull)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/documentSymbol", params, true, on_resp)
}

// A request to provide commands for the given text document and range.
func (it *Client) Request_textDocument_codeAction(params *CodeActionParams, onResp func(* /*TOr*/ []*CommandOrCodeAction)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/codeAction", params, true, on_resp)
}

// Request to resolve additional information for a given code action.The request's
// parameter is of type {@link CodeAction} the response
// is of type {@link CodeAction} or a Thenable that resolves to such.
func (it *Client) Request_codeAction_resolve(params *CodeAction, onResp func(CodeAction)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("codeAction/resolve", params, true, on_resp)
}

// A request to list project-wide symbols matching the query string given
// by the {@link WorkspaceSymbolParams}. The response is
// of type {@link SymbolInformation SymbolInformation[]} or a Thenable that
// resolves to such.
//
// @since 3.17.0 - support for WorkspaceSymbol in the returned data. Clients
//
//	need to advertise support for WorkspaceSymbols via the client capability
//	`workspace.symbol.resolveSupport`.
func (it *Client) Request_workspace_symbol(params *WorkspaceSymbolParams, onResp func(*SymbolInformationsOrWorkspaceSymbolsOrNull)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("workspace/symbol", params, true, on_resp)
}

// A request to resolve the range inside the workspace
// symbol's location.
//
// @since 3.17.0
func (it *Client) Request_workspaceSymbol_resolve(params *WorkspaceSymbol, onResp func(WorkspaceSymbol)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("workspaceSymbol/resolve", params, true, on_resp)
}

// A request to provide code lens for the given text document.
func (it *Client) Request_textDocument_codeLens(params *CodeLensParams, onResp func(* /*TOr*/ []CodeLens)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/codeLens", params, true, on_resp)
}

// A request to resolve a command for a given code lens.
func (it *Client) Request_codeLens_resolve(params *CodeLens, onResp func(CodeLens)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("codeLens/resolve", params, true, on_resp)
}

// A request to provide document links
func (it *Client) Request_textDocument_documentLink(params *DocumentLinkParams, onResp func(* /*TOr*/ []DocumentLink)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/documentLink", params, true, on_resp)
}

// Request to resolve additional information for a given document link. The request's
// parameter is of type {@link DocumentLink} the response
// is of type {@link DocumentLink} or a Thenable that resolves to such.
func (it *Client) Request_documentLink_resolve(params *DocumentLink, onResp func(DocumentLink)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("documentLink/resolve", params, true, on_resp)
}

// A request to format a whole document.
func (it *Client) Request_textDocument_formatting(params *DocumentFormattingParams, onResp func(* /*TOr*/ []TextEdit)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/formatting", params, true, on_resp)
}

// A request to format a range in a document.
func (it *Client) Request_textDocument_rangeFormatting(params *DocumentRangeFormattingParams, onResp func(* /*TOr*/ []TextEdit)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/rangeFormatting", params, true, on_resp)
}

// A request to format a document on type.
func (it *Client) Request_textDocument_onTypeFormatting(params *DocumentOnTypeFormattingParams, onResp func(* /*TOr*/ []TextEdit)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/onTypeFormatting", params, true, on_resp)
}

// A request to rename a symbol.
func (it *Client) Request_textDocument_rename(params *RenameParams, onResp func(* /*TOr*/ /*TOpt*/ *WorkspaceEdit)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/rename", params, true, on_resp)
}

// A request to test and perform the setup necessary for a rename.
//
// @since 3.16 - support for default behavior
func (it *Client) Request_textDocument_prepareRename(params *PrepareRenameParams, onResp func(* /*TOr*/ PrepareRenameResult)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("textDocument/prepareRename", params, true, on_resp)
}

// A request send from the client to the server to execute a command. The request might return
// a workspace edit which the client will apply to the workspace.
func (it *Client) Request_workspace_executeCommand(params *ExecuteCommandParams, onResp func(* /*TOr*/ LSPAny)) {
	var on_resp func(any) = clientServerOnResp(&it.clientServerBase, onResp)
	go it.send("workspace/executeCommand", params, true, on_resp)
}

func (it *Client) handleIncoming(raw map[string]any) *jsonRpcError {
	msg_id, msg_method := raw["id"], raw["method"]

	switch msg_method, _ := msg_method.(string); msg_method {
	case "window/showMessage":
		clientServerHandleIncoming(&it.clientServerBase, it.On_window_showMessage, msg_method, msg_id, raw["params"])
	case "window/logMessage":
		clientServerHandleIncoming(&it.clientServerBase, it.On_window_logMessage, msg_method, msg_id, raw["params"])
	case "telemetry/event":
		clientServerHandleIncoming(&it.clientServerBase, it.On_telemetry_event, msg_method, msg_id, raw["params"])
	case "textDocument/publishDiagnostics":
		clientServerHandleIncoming(&it.clientServerBase, it.On_textDocument_publishDiagnostics, msg_method, msg_id, raw["params"])
	case "$/logTrace":
		clientServerHandleIncoming(&it.clientServerBase, it.On___logTrace, msg_method, msg_id, raw["params"])
	case "$/cancelRequest":
		clientServerHandleIncoming(&it.clientServerBase, it.On___cancelRequest, msg_method, msg_id, raw["params"])
	case "$/progress":
		clientServerHandleIncoming(&it.clientServerBase, it.On___progress, msg_method, msg_id, raw["params"])
	case "workspace/workspaceFolders":
		clientServerHandleIncoming(&it.clientServerBase, it.On_workspace_workspaceFolders, msg_method, msg_id, raw["params"])
	case "workspace/configuration":
		clientServerHandleIncoming(&it.clientServerBase, it.On_workspace_configuration, msg_method, msg_id, raw["params"])
	case "window/workDoneProgress/create":
		clientServerHandleIncoming(&it.clientServerBase, it.On_window_workDoneProgress_create, msg_method, msg_id, raw["params"])
	case "workspace/semanticTokens/refresh":
		clientServerHandleIncoming(&it.clientServerBase, it.On_workspace_semanticTokens_refresh, msg_method, msg_id, raw["params"])
	case "window/showDocument":
		clientServerHandleIncoming(&it.clientServerBase, it.On_window_showDocument, msg_method, msg_id, raw["params"])
	case "workspace/inlineValue/refresh":
		clientServerHandleIncoming(&it.clientServerBase, it.On_workspace_inlineValue_refresh, msg_method, msg_id, raw["params"])
	case "workspace/inlayHint/refresh":
		clientServerHandleIncoming(&it.clientServerBase, it.On_workspace_inlayHint_refresh, msg_method, msg_id, raw["params"])
	case "workspace/diagnostic/refresh":
		clientServerHandleIncoming(&it.clientServerBase, it.On_workspace_diagnostic_refresh, msg_method, msg_id, raw["params"])
	case "client/registerCapability":
		clientServerHandleIncoming(&it.clientServerBase, it.On_client_registerCapability, msg_method, msg_id, raw["params"])
	case "client/unregisterCapability":
		clientServerHandleIncoming(&it.clientServerBase, it.On_client_unregisterCapability, msg_method, msg_id, raw["params"])
	case "window/showMessageRequest":
		clientServerHandleIncoming(&it.clientServerBase, it.On_window_showMessageRequest, msg_method, msg_id, raw["params"])
	case "workspace/codeLens/refresh":
		clientServerHandleIncoming(&it.clientServerBase, it.On_workspace_codeLens_refresh, msg_method, msg_id, raw["params"])
	case "workspace/applyEdit":
		clientServerHandleIncoming(&it.clientServerBase, it.On_workspace_applyEdit, msg_method, msg_id, raw["params"])
	default: // msg is an incoming Request or Notification
		return &jsonRpcError{Code: int(ErrorCodesMethodNotFound), Message: "unknown method: " + msg_method}
	}

	return nil
}

// Forever keeps reading and handling LSP JSON-RPC messages incoming over `serverOut`
// until reading from `serverOut` fails, then returns that IO read error.
func (it *Client) Forever(serverOut io.Reader, serverIn io.Writer) error {
	return it.forever(serverOut, serverIn, it.handleIncoming)
}
