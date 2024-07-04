// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-lsp/gen/cmd/gen_lsp
package lsp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Server struct {
	sync.Mutex // sync writes to stdout
	stdout     *bufio.Writer
	waiters    map[string]func(any)

	Lang struct {
		CompletionTriggerChars []string
		SignatureTriggerChars  []string
		Commands               []string
	}

	// Initialized is for informational purposes only, to the importer who shall not set or mutate them.
	// Its fields are set automatically at the appropriate initialization lifecycle instant.
	Initialized struct {
		Client *InitializeParams
		Server *InitializeResult
	}

	// The `workspace/didChangeWorkspaceFolders` notification is sent from the client to the server when the workspace
	// folder configuration changes.
	On_workspace_didChangeWorkspaceFolders func(params *DidChangeWorkspaceFoldersParams) (any, error)

	// The `window/workDoneProgress/cancel` notification is sent from  the client to the server to cancel a progress
	// initiated on the server side.
	On_window_workDoneProgress_cancel func(params *WorkDoneProgressCancelParams) (any, error)

	// The did create files notification is sent from the client to the server when
	// files were created from within the client.
	//
	// @since 3.16.0
	On_workspace_didCreateFiles func(params *CreateFilesParams) (any, error)

	// The did rename files notification is sent from the client to the server when
	// files were renamed from within the client.
	//
	// @since 3.16.0
	On_workspace_didRenameFiles func(params *RenameFilesParams) (any, error)

	// The will delete files request is sent from the client to the server before files are actually
	// deleted as long as the deletion is triggered from within the client.
	//
	// @since 3.16.0
	On_workspace_didDeleteFiles func(params *DeleteFilesParams) (any, error)

	// A notification sent when a notebook opens.
	//
	// @since 3.17.0
	On_notebookDocument_didOpen func(params *DidOpenNotebookDocumentParams) (any, error)

	//
	On_notebookDocument_didChange func(params *DidChangeNotebookDocumentParams) (any, error)

	// A notification sent when a notebook document is saved.
	//
	// @since 3.17.0
	On_notebookDocument_didSave func(params *DidSaveNotebookDocumentParams) (any, error)

	// A notification sent when a notebook closes.
	//
	// @since 3.17.0
	On_notebookDocument_didClose func(params *DidCloseNotebookDocumentParams) (any, error)

	// The initialized notification is sent from the client to the
	// server after the client is fully initialized and the server
	// is allowed to send requests from the server to the client.
	On_initialized func(params *InitializedParams) (any, error)

	// The exit event is sent from the client to the server to
	// ask the server to exit its process.
	On_exit func(params *Void) (any, error)

	// The configuration change notification is sent from the client to the server
	// when the client's configuration has changed. The notification contains
	// the changed configuration as defined by the language client.
	On_workspace_didChangeConfiguration func(params *DidChangeConfigurationParams) (any, error)

	// The document open notification is sent from the client to the server to signal
	// newly opened text documents. The document's truth is now managed by the client
	// and the server must not try to read the document's truth using the document's
	// uri. Open in this sense means it is managed by the client. It doesn't necessarily
	// mean that its content is presented in an editor. An open notification must not
	// be sent more than once without a corresponding close notification send before.
	// This means open and close notification must be balanced and the max open count
	// is one.
	On_textDocument_didOpen func(params *DidOpenTextDocumentParams) (any, error)

	// The document change notification is sent from the client to the server to signal
	// changes to a text document.
	On_textDocument_didChange func(params *DidChangeTextDocumentParams) (any, error)

	// The document close notification is sent from the client to the server when
	// the document got closed in the client. The document's truth now exists where
	// the document's uri points to (e.g. if the document's uri is a file uri the
	// truth now exists on disk). As with the open notification the close notification
	// is about managing the document's content. Receiving a close notification
	// doesn't mean that the document was open in an editor before. A close
	// notification requires a previous open notification to be sent.
	On_textDocument_didClose func(params *DidCloseTextDocumentParams) (any, error)

	// The document save notification is sent from the client to the server when
	// the document got saved in the client.
	On_textDocument_didSave func(params *DidSaveTextDocumentParams) (any, error)

	// A document will save notification is sent from the client to the server before
	// the document is actually saved.
	On_textDocument_willSave func(params *WillSaveTextDocumentParams) (any, error)

	// The watched files notification is sent from the client to the server when
	// the client detects changes to file watched by the language client.
	On_workspace_didChangeWatchedFiles func(params *DidChangeWatchedFilesParams) (any, error)

	//
	On___setTrace func(params *SetTraceParams) (any, error)

	//
	On___cancelRequest func(params *CancelParams) (any, error)

	//
	On___progress func(params *ProgressParams) (any, error)

	// A request to resolve the implementation locations of a symbol at a given text
	// document position. The request's parameter is of type {@link TextDocumentPositionParams}
	// the response is of type {@link Definition} or a Thenable that resolves to such.
	On_textDocument_implementation func(params *ImplementationParams) (any, error)

	// A request to resolve the type definition locations of a symbol at a given text
	// document position. The request's parameter is of type {@link TextDocumentPositionParams}
	// the response is of type {@link Definition} or a Thenable that resolves to such.
	On_textDocument_typeDefinition func(params *TypeDefinitionParams) (any, error)

	// A request to list all color symbols found in a given text document. The request's
	// parameter is of type {@link DocumentColorParams} the
	// response is of type {@link ColorInformation ColorInformation[]} or a Thenable
	// that resolves to such.
	On_textDocument_documentColor func(params *DocumentColorParams) (any, error)

	// A request to list all presentation for a color. The request's
	// parameter is of type {@link ColorPresentationParams} the
	// response is of type {@link ColorInformation ColorInformation[]} or a Thenable
	// that resolves to such.
	On_textDocument_colorPresentation func(params *ColorPresentationParams) (any, error)

	// A request to provide folding ranges in a document. The request's
	// parameter is of type {@link FoldingRangeParams}, the
	// response is of type {@link FoldingRangeList} or a Thenable
	// that resolves to such.
	On_textDocument_foldingRange func(params *FoldingRangeParams) (any, error)

	// A request to resolve the type definition locations of a symbol at a given text
	// document position. The request's parameter is of type {@link TextDocumentPositionParams}
	// the response is of type {@link Declaration} or a typed array of {@link DeclarationLink}
	// or a Thenable that resolves to such.
	On_textDocument_declaration func(params *DeclarationParams) (any, error)

	// A request to provide selection ranges in a document. The request's
	// parameter is of type {@link SelectionRangeParams}, the
	// response is of type {@link SelectionRange SelectionRange[]} or a Thenable
	// that resolves to such.
	On_textDocument_selectionRange func(params *SelectionRangeParams) (any, error)

	// A request to result a `CallHierarchyItem` in a document at a given position.
	// Can be used as an input to an incoming or outgoing call hierarchy.
	//
	// @since 3.16.0
	On_textDocument_prepareCallHierarchy func(params *CallHierarchyPrepareParams) (any, error)

	// A request to resolve the incoming calls for a given `CallHierarchyItem`.
	//
	// @since 3.16.0
	On_callHierarchy_incomingCalls func(params *CallHierarchyIncomingCallsParams) (any, error)

	// A request to resolve the outgoing calls for a given `CallHierarchyItem`.
	//
	// @since 3.16.0
	On_callHierarchy_outgoingCalls func(params *CallHierarchyOutgoingCallsParams) (any, error)

	// @since 3.16.0
	On_textDocument_semanticTokens_full func(params *SemanticTokensParams) (any, error)

	// @since 3.16.0
	On_textDocument_semanticTokens_full_delta func(params *SemanticTokensDeltaParams) (any, error)

	// @since 3.16.0
	On_textDocument_semanticTokens_range func(params *SemanticTokensRangeParams) (any, error)

	// A request to provide ranges that can be edited together.
	//
	// @since 3.16.0
	On_textDocument_linkedEditingRange func(params *LinkedEditingRangeParams) (any, error)

	// The will create files request is sent from the client to the server before files are actually
	// created as long as the creation is triggered from within the client.
	//
	// The request can return a `WorkspaceEdit` which will be applied to workspace before the
	// files are created. Hence the `WorkspaceEdit` can not manipulate the content of the file
	// to be created.
	//
	// @since 3.16.0
	On_workspace_willCreateFiles func(params *CreateFilesParams) (any, error)

	// The will rename files request is sent from the client to the server before files are actually
	// renamed as long as the rename is triggered from within the client.
	//
	// @since 3.16.0
	On_workspace_willRenameFiles func(params *RenameFilesParams) (any, error)

	// The did delete files notification is sent from the client to the server when
	// files were deleted from within the client.
	//
	// @since 3.16.0
	On_workspace_willDeleteFiles func(params *DeleteFilesParams) (any, error)

	// A request to get the moniker of a symbol at a given text document position.
	// The request parameter is of type {@link TextDocumentPositionParams}.
	// The response is of type {@link Moniker Moniker[]} or `null`.
	On_textDocument_moniker func(params *MonikerParams) (any, error)

	// A request to result a `TypeHierarchyItem` in a document at a given position.
	// Can be used as an input to a subtypes or supertypes type hierarchy.
	//
	// @since 3.17.0
	On_textDocument_prepareTypeHierarchy func(params *TypeHierarchyPrepareParams) (any, error)

	// A request to resolve the supertypes for a given `TypeHierarchyItem`.
	//
	// @since 3.17.0
	On_typeHierarchy_supertypes func(params *TypeHierarchySupertypesParams) (any, error)

	// A request to resolve the subtypes for a given `TypeHierarchyItem`.
	//
	// @since 3.17.0
	On_typeHierarchy_subtypes func(params *TypeHierarchySubtypesParams) (any, error)

	// A request to provide inline values in a document. The request's parameter is of
	// type {@link InlineValueParams}, the response is of type
	// {@link InlineValue InlineValue[]} or a Thenable that resolves to such.
	//
	// @since 3.17.0
	On_textDocument_inlineValue func(params *InlineValueParams) (any, error)

	// A request to provide inlay hints in a document. The request's parameter is of
	// type {@link InlayHintsParams}, the response is of type
	// {@link InlayHint InlayHint[]} or a Thenable that resolves to such.
	//
	// @since 3.17.0
	On_textDocument_inlayHint func(params *InlayHintParams) (any, error)

	// A request to resolve additional properties for an inlay hint.
	// The request's parameter is of type {@link InlayHint}, the response is
	// of type {@link InlayHint} or a Thenable that resolves to such.
	//
	// @since 3.17.0
	On_inlayHint_resolve func(params *InlayHint) (any, error)

	// The document diagnostic request definition.
	//
	// @since 3.17.0
	On_textDocument_diagnostic func(params *DocumentDiagnosticParams) (any, error)

	// The workspace diagnostic request definition.
	//
	// @since 3.17.0
	On_workspace_diagnostic func(params *WorkspaceDiagnosticParams) (any, error)

	// A shutdown request is sent from the client to the server.
	// It is sent once when the client decides to shutdown the
	// server. The only notification that is sent after a shutdown request
	// is the exit event.
	On_shutdown func(params *Void) (any, error)

	// A document will save request is sent from the client to the server before
	// the document is actually saved. The request can return an array of TextEdits
	// which will be applied to the text document before it is saved. Please note that
	// clients might drop results if computing the text edits took too long or if a
	// server constantly fails on this request. This is done to keep the save fast and
	// reliable.
	On_textDocument_willSaveWaitUntil func(params *WillSaveTextDocumentParams) (any, error)

	// Request to request completion at a given text document position. The request's
	// parameter is of type {@link TextDocumentPosition} the response
	// is of type {@link CompletionItem CompletionItem[]} or {@link CompletionList}
	// or a Thenable that resolves to such.
	//
	// The request can delay the computation of the {@link CompletionItem.detail `detail`}
	// and {@link CompletionItem.documentation `documentation`} properties to the `completionItem/resolve`
	// request. However, properties that are needed for the initial sorting and filtering, like `sortText`,
	// `filterText`, `insertText`, and `textEdit`, must not be changed during resolve.
	On_textDocument_completion func(params *CompletionParams) (any, error)

	// Request to resolve additional information for a given completion item.The request's
	// parameter is of type {@link CompletionItem} the response
	// is of type {@link CompletionItem} or a Thenable that resolves to such.
	On_completionItem_resolve func(params *CompletionItem) (any, error)

	// Request to request hover information at a given text document position. The request's
	// parameter is of type {@link TextDocumentPosition} the response is of
	// type {@link Hover} or a Thenable that resolves to such.
	On_textDocument_hover func(params *HoverParams) (any, error)

	//
	On_textDocument_signatureHelp func(params *SignatureHelpParams) (any, error)

	// A request to resolve the definition location of a symbol at a given text
	// document position. The request's parameter is of type {@link TextDocumentPosition}
	// the response is of either type {@link Definition} or a typed array of
	// {@link DefinitionLink} or a Thenable that resolves to such.
	On_textDocument_definition func(params *DefinitionParams) (any, error)

	// A request to resolve project-wide references for the symbol denoted
	// by the given text document position. The request's parameter is of
	// type {@link ReferenceParams} the response is of type
	// {@link Location Location[]} or a Thenable that resolves to such.
	On_textDocument_references func(params *ReferenceParams) (any, error)

	// Request to resolve a {@link DocumentHighlight} for a given
	// text document position. The request's parameter is of type {@link TextDocumentPosition}
	// the request response is an array of type {@link DocumentHighlight}
	// or a Thenable that resolves to such.
	On_textDocument_documentHighlight func(params *DocumentHighlightParams) (any, error)

	// A request to list all symbols found in a given text document. The request's
	// parameter is of type {@link TextDocumentIdentifier} the
	// response is of type {@link SymbolInformation SymbolInformation[]} or a Thenable
	// that resolves to such.
	On_textDocument_documentSymbol func(params *DocumentSymbolParams) (any, error)

	// A request to provide commands for the given text document and range.
	On_textDocument_codeAction func(params *CodeActionParams) (any, error)

	// Request to resolve additional information for a given code action.The request's
	// parameter is of type {@link CodeAction} the response
	// is of type {@link CodeAction} or a Thenable that resolves to such.
	On_codeAction_resolve func(params *CodeAction) (any, error)

	// A request to list project-wide symbols matching the query string given
	// by the {@link WorkspaceSymbolParams}. The response is
	// of type {@link SymbolInformation SymbolInformation[]} or a Thenable that
	// resolves to such.
	//
	// @since 3.17.0 - support for WorkspaceSymbol in the returned data. Clients
	//  need to advertise support for WorkspaceSymbols via the client capability
	//  `workspace.symbol.resolveSupport`.
	//
	On_workspace_symbol func(params *WorkspaceSymbolParams) (any, error)

	// A request to resolve the range inside the workspace
	// symbol's location.
	//
	// @since 3.17.0
	On_workspaceSymbol_resolve func(params *WorkspaceSymbol) (any, error)

	// A request to provide code lens for the given text document.
	On_textDocument_codeLens func(params *CodeLensParams) (any, error)

	// A request to resolve a command for a given code lens.
	On_codeLens_resolve func(params *CodeLens) (any, error)

	// A request to provide document links
	On_textDocument_documentLink func(params *DocumentLinkParams) (any, error)

	// Request to resolve additional information for a given document link. The request's
	// parameter is of type {@link DocumentLink} the response
	// is of type {@link DocumentLink} or a Thenable that resolves to such.
	On_documentLink_resolve func(params *DocumentLink) (any, error)

	// A request to format a whole document.
	On_textDocument_formatting func(params *DocumentFormattingParams) (any, error)

	// A request to format a range in a document.
	On_textDocument_rangeFormatting func(params *DocumentRangeFormattingParams) (any, error)

	// A request to format a document on type.
	On_textDocument_onTypeFormatting func(params *DocumentOnTypeFormattingParams) (any, error)

	// A request to rename a symbol.
	On_textDocument_rename func(params *RenameParams) (any, error)

	// A request to test and perform the setup necessary for a rename.
	//
	// @since 3.16 - support for default behavior
	On_textDocument_prepareRename func(params *PrepareRenameParams) (any, error)

	// A request send from the client to the server to execute a command. The request might return
	// a workspace edit which the client will apply to the workspace.
	On_workspace_executeCommand func(params *ExecuteCommandParams) (any, error)
}

// The show message notification is sent from a server to a client to ask
// the client to display a particular message in the user interface.
func (it *Server) Notify_window_showMessage(params *ShowMessageParams) {
	var on_resp func(any)
	go it.send("window/showMessage", params, false, on_resp)
}

// The log message notification is sent from the server to the client to ask
// the client to log a particular message.
func (it *Server) Notify_window_logMessage(params *LogMessageParams) {
	var on_resp func(any)
	go it.send("window/logMessage", params, false, on_resp)
}

// The telemetry event notification is sent from the server to the client to ask
// the client to log telemetry data.
func (it *Server) Notify_telemetry_event(params *LSPAny) {
	var on_resp func(any)
	go it.send("telemetry/event", params, false, on_resp)
}

// Diagnostics notification are sent from the server to the client to signal
// results of validation runs.
func (it *Server) Notify_textDocument_publishDiagnostics(params *PublishDiagnosticsParams) {
	var on_resp func(any)
	go it.send("textDocument/publishDiagnostics", params, false, on_resp)
}

func (it *Server) Notify___logTrace(params *LogTraceParams) {
	var on_resp func(any)
	go it.send("$/logTrace", params, false, on_resp)
}

func (it *Server) Notify___cancelRequest(params *CancelParams) {
	var on_resp func(any)
	go it.send("$/cancelRequest", params, false, on_resp)
}

func (it *Server) Notify___progress(params *ProgressParams) {
	var on_resp func(any)
	go it.send("$/progress", params, false, on_resp)
}

// The `workspace/workspaceFolders` is sent from the server to the client to fetch the open workspace folders.
func (it *Server) Request_workspace_workspaceFolders(params *Void, onResp func(* /*TOr*/ /*TOpt*/ []WorkspaceFolder)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("workspace/workspaceFolders", params, true, on_resp)
}

// The 'workspace/configuration' request is sent from the server to the client to fetch a certain
// configuration setting.
//
// This pull model replaces the old push model where the client signaled configuration change via an
// event. If the server still needs to react to configuration changes (since the server caches the
// result of `workspace/configuration` requests) the server should register for an empty configuration
// change event and empty the cache if such an event is received.
func (it *Server) Request_workspace_configuration(params *ConfigurationParams, onResp func(*[]LSPAny)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("workspace/configuration", params, true, on_resp)
}

// The `window/workDoneProgress/create` request is sent from the server to the client to initiate progress
// reporting from the server.
func (it *Server) Request_window_workDoneProgress_create(params *WorkDoneProgressCreateParams, onResp func(*Void)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("window/workDoneProgress/create", params, true, on_resp)
}

// @since 3.16.0
func (it *Server) Request_workspace_semanticTokens_refresh(params *Void, onResp func(*Void)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("workspace/semanticTokens/refresh", params, true, on_resp)
}

// A request to show a document. This request might open an
// external program depending on the value of the URI to open.
// For example a request to open `https://code.visualstudio.com/`
// will very likely open the URI in a WEB browser.
//
// @since 3.16.0
func (it *Server) Request_window_showDocument(params *ShowDocumentParams, onResp func(*ShowDocumentResult)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("window/showDocument", params, true, on_resp)
}

// @since 3.17.0
func (it *Server) Request_workspace_inlineValue_refresh(params *Void, onResp func(*Void)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("workspace/inlineValue/refresh", params, true, on_resp)
}

// @since 3.17.0
func (it *Server) Request_workspace_inlayHint_refresh(params *Void, onResp func(*Void)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("workspace/inlayHint/refresh", params, true, on_resp)
}

// The diagnostic refresh request definition.
//
// @since 3.17.0
func (it *Server) Request_workspace_diagnostic_refresh(params *Void, onResp func(*Void)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("workspace/diagnostic/refresh", params, true, on_resp)
}

// The `client/registerCapability` request is sent from the server to the client to register a new capability
// handler on the client side.
func (it *Server) Request_client_registerCapability(params *RegistrationParams, onResp func(*Void)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("client/registerCapability", params, true, on_resp)
}

// The `client/unregisterCapability` request is sent from the server to the client to unregister a previously registered capability
// handler on the client side.
func (it *Server) Request_client_unregisterCapability(params *UnregistrationParams, onResp func(*Void)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("client/unregisterCapability", params, true, on_resp)
}

// The show message request is sent from the server to the client to show a message
// and a set of options actions to the user.
func (it *Server) Request_window_showMessageRequest(params *ShowMessageRequestParams, onResp func(* /*TOr*/ /*TOpt*/ *MessageActionItem)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("window/showMessageRequest", params, true, on_resp)
}

// A request to refresh all code actions
//
// @since 3.16.0
func (it *Server) Request_workspace_codeLens_refresh(params *Void, onResp func(*Void)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("workspace/codeLens/refresh", params, true, on_resp)
}

// A request sent from the server to the client to modified certain resources.
func (it *Server) Request_workspace_applyEdit(params *ApplyWorkspaceEditParams, onResp func(*ApplyWorkspaceEditResult)) {
	var on_resp func(any) = serverOnResp(it, onResp)
	go it.send("workspace/applyEdit", params, true, on_resp)
}

func serverOnResp[T any](it *Server, onResp func(*T)) func(any) {
	if onResp == nil {
		return nil
	}
	return func(resultAsMap any) {
		var result T
		if resultAsMap != nil {
			json_bytes, _ := json.Marshal(resultAsMap)
			if err := json.Unmarshal(json_bytes, &result); err != nil {
				it.sendErrMsg(err)
				return
			}
		}
		onResp(iIf(resultAsMap == nil, nil, &result))
	}
}

func (it *Server) handleIncoming(jsonRpcMsg []byte) *jsonRpcError {
	raw := map[string]any{}
	if err := json.Unmarshal(jsonRpcMsg, &raw); err != nil {
		return &jsonRpcError{Code: -32700, Message: err.Error()}
	}
	msg_id, msg_method, msg_err_code := raw["id"], raw["method"], raw["code"]

	if msg_err_code != nil { // received an error Response
		println(string(jsonRpcMsg)) // goes to stderr
		return nil
	}

	switch msg_method, _ := msg_method.(string); msg_method {
	case "": // msg is an incoming Response
		msg_id := fmt.Sprintf("%v", msg_id)
		handler := it.waiters[msg_id]
		delete(it.waiters, msg_id)
		go handler(raw["result"])
	case "workspace/didChangeWorkspaceFolders":
		serverHandleIncoming(it, it.On_workspace_didChangeWorkspaceFolders, msg_method, msg_id, raw["params"])
	case "window/workDoneProgress/cancel":
		serverHandleIncoming(it, it.On_window_workDoneProgress_cancel, msg_method, msg_id, raw["params"])
	case "workspace/didCreateFiles":
		serverHandleIncoming(it, it.On_workspace_didCreateFiles, msg_method, msg_id, raw["params"])
	case "workspace/didRenameFiles":
		serverHandleIncoming(it, it.On_workspace_didRenameFiles, msg_method, msg_id, raw["params"])
	case "workspace/didDeleteFiles":
		serverHandleIncoming(it, it.On_workspace_didDeleteFiles, msg_method, msg_id, raw["params"])
	case "notebookDocument/didOpen":
		serverHandleIncoming(it, it.On_notebookDocument_didOpen, msg_method, msg_id, raw["params"])
	case "notebookDocument/didChange":
		serverHandleIncoming(it, it.On_notebookDocument_didChange, msg_method, msg_id, raw["params"])
	case "notebookDocument/didSave":
		serverHandleIncoming(it, it.On_notebookDocument_didSave, msg_method, msg_id, raw["params"])
	case "notebookDocument/didClose":
		serverHandleIncoming(it, it.On_notebookDocument_didClose, msg_method, msg_id, raw["params"])
	case "initialized":
		serverHandleIncoming(it, it.On_initialized, msg_method, msg_id, raw["params"])
	case "exit":
		serverHandleIncoming(it, it.On_exit, msg_method, msg_id, raw["params"])
	case "workspace/didChangeConfiguration":
		serverHandleIncoming(it, it.On_workspace_didChangeConfiguration, msg_method, msg_id, raw["params"])
	case "textDocument/didOpen":
		serverHandleIncoming(it, it.On_textDocument_didOpen, msg_method, msg_id, raw["params"])
	case "textDocument/didChange":
		serverHandleIncoming(it, it.On_textDocument_didChange, msg_method, msg_id, raw["params"])
	case "textDocument/didClose":
		serverHandleIncoming(it, it.On_textDocument_didClose, msg_method, msg_id, raw["params"])
	case "textDocument/didSave":
		serverHandleIncoming(it, it.On_textDocument_didSave, msg_method, msg_id, raw["params"])
	case "textDocument/willSave":
		serverHandleIncoming(it, it.On_textDocument_willSave, msg_method, msg_id, raw["params"])
	case "workspace/didChangeWatchedFiles":
		serverHandleIncoming(it, it.On_workspace_didChangeWatchedFiles, msg_method, msg_id, raw["params"])
	case "$/setTrace":
		serverHandleIncoming(it, it.On___setTrace, msg_method, msg_id, raw["params"])
	case "$/cancelRequest":
		serverHandleIncoming(it, it.On___cancelRequest, msg_method, msg_id, raw["params"])
	case "$/progress":
		serverHandleIncoming(it, it.On___progress, msg_method, msg_id, raw["params"])
	case "textDocument/implementation":
		serverHandleIncoming(it, it.On_textDocument_implementation, msg_method, msg_id, raw["params"])
	case "textDocument/typeDefinition":
		serverHandleIncoming(it, it.On_textDocument_typeDefinition, msg_method, msg_id, raw["params"])
	case "textDocument/documentColor":
		serverHandleIncoming(it, it.On_textDocument_documentColor, msg_method, msg_id, raw["params"])
	case "textDocument/colorPresentation":
		serverHandleIncoming(it, it.On_textDocument_colorPresentation, msg_method, msg_id, raw["params"])
	case "textDocument/foldingRange":
		serverHandleIncoming(it, it.On_textDocument_foldingRange, msg_method, msg_id, raw["params"])
	case "textDocument/declaration":
		serverHandleIncoming(it, it.On_textDocument_declaration, msg_method, msg_id, raw["params"])
	case "textDocument/selectionRange":
		serverHandleIncoming(it, it.On_textDocument_selectionRange, msg_method, msg_id, raw["params"])
	case "textDocument/prepareCallHierarchy":
		serverHandleIncoming(it, it.On_textDocument_prepareCallHierarchy, msg_method, msg_id, raw["params"])
	case "callHierarchy/incomingCalls":
		serverHandleIncoming(it, it.On_callHierarchy_incomingCalls, msg_method, msg_id, raw["params"])
	case "callHierarchy/outgoingCalls":
		serverHandleIncoming(it, it.On_callHierarchy_outgoingCalls, msg_method, msg_id, raw["params"])
	case "textDocument/semanticTokens/full":
		serverHandleIncoming(it, it.On_textDocument_semanticTokens_full, msg_method, msg_id, raw["params"])
	case "textDocument/semanticTokens/full/delta":
		serverHandleIncoming(it, it.On_textDocument_semanticTokens_full_delta, msg_method, msg_id, raw["params"])
	case "textDocument/semanticTokens/range":
		serverHandleIncoming(it, it.On_textDocument_semanticTokens_range, msg_method, msg_id, raw["params"])
	case "textDocument/linkedEditingRange":
		serverHandleIncoming(it, it.On_textDocument_linkedEditingRange, msg_method, msg_id, raw["params"])
	case "workspace/willCreateFiles":
		serverHandleIncoming(it, it.On_workspace_willCreateFiles, msg_method, msg_id, raw["params"])
	case "workspace/willRenameFiles":
		serverHandleIncoming(it, it.On_workspace_willRenameFiles, msg_method, msg_id, raw["params"])
	case "workspace/willDeleteFiles":
		serverHandleIncoming(it, it.On_workspace_willDeleteFiles, msg_method, msg_id, raw["params"])
	case "textDocument/moniker":
		serverHandleIncoming(it, it.On_textDocument_moniker, msg_method, msg_id, raw["params"])
	case "textDocument/prepareTypeHierarchy":
		serverHandleIncoming(it, it.On_textDocument_prepareTypeHierarchy, msg_method, msg_id, raw["params"])
	case "typeHierarchy/supertypes":
		serverHandleIncoming(it, it.On_typeHierarchy_supertypes, msg_method, msg_id, raw["params"])
	case "typeHierarchy/subtypes":
		serverHandleIncoming(it, it.On_typeHierarchy_subtypes, msg_method, msg_id, raw["params"])
	case "textDocument/inlineValue":
		serverHandleIncoming(it, it.On_textDocument_inlineValue, msg_method, msg_id, raw["params"])
	case "textDocument/inlayHint":
		serverHandleIncoming(it, it.On_textDocument_inlayHint, msg_method, msg_id, raw["params"])
	case "inlayHint/resolve":
		serverHandleIncoming(it, it.On_inlayHint_resolve, msg_method, msg_id, raw["params"])
	case "textDocument/diagnostic":
		serverHandleIncoming(it, it.On_textDocument_diagnostic, msg_method, msg_id, raw["params"])
	case "workspace/diagnostic":
		serverHandleIncoming(it, it.On_workspace_diagnostic, msg_method, msg_id, raw["params"])
	case "shutdown":
		serverHandleIncoming(it, it.On_shutdown, msg_method, msg_id, raw["params"])
	case "textDocument/willSaveWaitUntil":
		serverHandleIncoming(it, it.On_textDocument_willSaveWaitUntil, msg_method, msg_id, raw["params"])
	case "textDocument/completion":
		serverHandleIncoming(it, it.On_textDocument_completion, msg_method, msg_id, raw["params"])
	case "completionItem/resolve":
		serverHandleIncoming(it, it.On_completionItem_resolve, msg_method, msg_id, raw["params"])
	case "textDocument/hover":
		serverHandleIncoming(it, it.On_textDocument_hover, msg_method, msg_id, raw["params"])
	case "textDocument/signatureHelp":
		serverHandleIncoming(it, it.On_textDocument_signatureHelp, msg_method, msg_id, raw["params"])
	case "textDocument/definition":
		serverHandleIncoming(it, it.On_textDocument_definition, msg_method, msg_id, raw["params"])
	case "textDocument/references":
		serverHandleIncoming(it, it.On_textDocument_references, msg_method, msg_id, raw["params"])
	case "textDocument/documentHighlight":
		serverHandleIncoming(it, it.On_textDocument_documentHighlight, msg_method, msg_id, raw["params"])
	case "textDocument/documentSymbol":
		serverHandleIncoming(it, it.On_textDocument_documentSymbol, msg_method, msg_id, raw["params"])
	case "textDocument/codeAction":
		serverHandleIncoming(it, it.On_textDocument_codeAction, msg_method, msg_id, raw["params"])
	case "codeAction/resolve":
		serverHandleIncoming(it, it.On_codeAction_resolve, msg_method, msg_id, raw["params"])
	case "workspace/symbol":
		serverHandleIncoming(it, it.On_workspace_symbol, msg_method, msg_id, raw["params"])
	case "workspaceSymbol/resolve":
		serverHandleIncoming(it, it.On_workspaceSymbol_resolve, msg_method, msg_id, raw["params"])
	case "textDocument/codeLens":
		serverHandleIncoming(it, it.On_textDocument_codeLens, msg_method, msg_id, raw["params"])
	case "codeLens/resolve":
		serverHandleIncoming(it, it.On_codeLens_resolve, msg_method, msg_id, raw["params"])
	case "textDocument/documentLink":
		serverHandleIncoming(it, it.On_textDocument_documentLink, msg_method, msg_id, raw["params"])
	case "documentLink/resolve":
		serverHandleIncoming(it, it.On_documentLink_resolve, msg_method, msg_id, raw["params"])
	case "textDocument/formatting":
		serverHandleIncoming(it, it.On_textDocument_formatting, msg_method, msg_id, raw["params"])
	case "textDocument/rangeFormatting":
		serverHandleIncoming(it, it.On_textDocument_rangeFormatting, msg_method, msg_id, raw["params"])
	case "textDocument/onTypeFormatting":
		serverHandleIncoming(it, it.On_textDocument_onTypeFormatting, msg_method, msg_id, raw["params"])
	case "textDocument/rename":
		serverHandleIncoming(it, it.On_textDocument_rename, msg_method, msg_id, raw["params"])
	case "textDocument/prepareRename":
		serverHandleIncoming(it, it.On_textDocument_prepareRename, msg_method, msg_id, raw["params"])
	case "workspace/executeCommand":
		serverHandleIncoming(it, it.On_workspace_executeCommand, msg_method, msg_id, raw["params"])
	case "initialize":
		serverHandleIncoming(it, func(params *InitializeParams) (any, error) {
			init := &it.Initialized
			init.Client = params
			init.Server = &InitializeResult{
				ServerInfo: &struct {
					Name    string  "json:\"name\""
					Version *String "json:\"version\""
				}{Name: os.Args[0]},
			}
			caps := &init.Server.Capabilities
			caps.TextDocumentSync.TextDocumentSyncOptions = &TextDocumentSyncOptions{
				OpenClose: ptr(Boolean(it.On_textDocument_didClose != nil || it.On_textDocument_didOpen != nil)),
				Change:    iIf(it.On_textDocument_didChange == nil, TextDocumentSyncKindNone, TextDocumentSyncKindFull),
			}
			if it.On_textDocument_completion != nil {
				caps.CompletionProvider = &CompletionOptions{TriggerCharacters: it.Lang.CompletionTriggerChars}
			}
			caps.HoverProvider.Boolean = ptr(Boolean(it.On_textDocument_hover != nil))
			if it.On_textDocument_signatureHelp != nil {
				caps.SignatureHelpProvider = &SignatureHelpOptions{TriggerCharacters: it.Lang.SignatureTriggerChars}
			}
			caps.DeclarationProvider.Boolean = ptr(Boolean(it.On_textDocument_declaration != nil))
			caps.DefinitionProvider.Boolean = ptr(Boolean(it.On_textDocument_definition != nil))
			caps.TypeDefinitionProvider.Boolean = ptr(Boolean(it.On_textDocument_typeDefinition != nil))
			caps.ImplementationProvider.Boolean = ptr(Boolean(it.On_textDocument_implementation != nil))
			caps.ReferencesProvider.Boolean = ptr(Boolean(it.On_textDocument_references != nil))
			caps.DocumentHighlightProvider.Boolean = ptr(Boolean(it.On_textDocument_documentHighlight != nil))
			caps.DocumentSymbolProvider.Boolean = ptr(Boolean(it.On_textDocument_documentSymbol != nil))
			caps.CodeActionProvider.Boolean = ptr(Boolean(it.On_textDocument_codeAction != nil))
			if it.On_textDocument_codeLens != nil {
				caps.CodeLensProvider = &CodeLensOptions{}
			}
			caps.DocumentFormattingProvider.Boolean = ptr(Boolean(it.On_textDocument_formatting != nil))
			caps.DocumentRangeFormattingProvider.Boolean = ptr(Boolean(it.On_textDocument_rangeFormatting != nil))
			caps.RenameProvider.RenameOptions = iIf(it.On_textDocument_rename == nil, nil, &RenameOptions{
				PrepareProvider: ptr(Boolean(it.On_textDocument_prepareRename != nil)),
			})
			if it.On_workspace_executeCommand != nil {
				caps.ExecuteCommandProvider = &ExecuteCommandOptions{Commands: it.Lang.Commands}
			}
			caps.SelectionRangeProvider.Boolean = ptr(Boolean(it.On_textDocument_selectionRange != nil))
			caps.CallHierarchyProvider.Boolean = ptr(Boolean(it.On_textDocument_prepareCallHierarchy != nil && it.On_callHierarchy_incomingCalls != nil && it.On_callHierarchy_outgoingCalls != nil))
			caps.TypeHierarchyProvider.Boolean = ptr(Boolean(it.On_textDocument_prepareTypeHierarchy != nil && it.On_typeHierarchy_subtypes != nil && it.On_typeHierarchy_supertypes != nil))
			caps.InlineValueProvider.Boolean = ptr(Boolean(it.On_textDocument_inlineValue != nil))
			caps.InlayHintProvider.Boolean = ptr(Boolean(it.On_textDocument_inlayHint != nil))
			caps.WorkspaceSymbolProvider.Boolean = ptr(Boolean(it.On_workspace_symbol != nil))
			caps.Workspace = &struct {
				WorkspaceFolders *WorkspaceFoldersServerCapabilities "json:\"workspaceFolders\""
				FileOperations   *FileOperationOptions               "json:\"fileOperations\""
			}{
				WorkspaceFolders: &WorkspaceFoldersServerCapabilities{
					Supported: ptr(Boolean(it.On_workspace_didChangeWorkspaceFolders != nil)),
					ChangeNotifications: struct {
						String  *String
						Boolean *Boolean
					}{
						Boolean: ptr(Boolean(it.On_workspace_didChangeWorkspaceFolders != nil)),
					},
				},
			}

			return init.Server, nil
		}, msg_method, msg_id, raw["params"])
	default: // msg is an incoming Request or Notification
		return &jsonRpcError{Code: -32601, Message: "unknown method: " + msg_method}
	}

	return nil
}

func ptr[T any](value T) *T { return &value }

func iIf[T any](chk bool, ifTrue T, ifFalse T) T {
	if chk {
		return ifTrue
	}
	return ifFalse
}

func serverHandleIncoming[T any](it *Server, handler func(*T) (any, error), msgMethodName string, msgIdMaybe any, msgParams any) {
	var req_id string
	if msgIdMaybe != nil {
		req_id = fmt.Sprintf("%v", msgIdMaybe)
	}
	if handler == nil {
		if req_id != "" {
			it.sendErrMsg(errors.New("unimplemented: " + msgMethodName))
		}
		return
	}
	var params T
	if msgParams != nil {
		json_bytes, _ := json.Marshal(msgParams)
		if err := json.Unmarshal(json_bytes, &params); err != nil {
			it.sendErrMsg(&jsonRpcError{Code: -32602, Message: err.Error()})
			return
		}
	}
	go func(params *T) {
		if msgParams == nil {
			params = nil
		}
		result, err := handler(params)
		resp := map[string]any{
			"result": result,
			"id":     req_id,
		}
		if err != nil {
			if msgIdMaybe != nil {
				resp["error"] = &jsonRpcError{Code: -32603, Message: fmt.Sprintf("%v", err)}
			} else {
				it.sendErrMsg(err)
				return
			}
		}
		if msgIdMaybe != nil {
			it.sendMsg(resp)
		}
	}(&params)
}

func (it *Server) sendErrMsg(err any) {
	if err == nil {
		return
	}
	var json_rpc_err_msg *jsonRpcError
	if json_rpc_err_msg, _ = err.(*jsonRpcError); json_rpc_err_msg == nil {
		json_rpc_err_msg = &jsonRpcError{Code: -32603, Message: fmt.Sprintf("%v", err)}
	}
	it.sendMsg(json_rpc_err_msg)
}

func (it *Server) sendMsg(jsonable any) {
	err_json, _ := json.Marshal(jsonable)
	it.Lock()
	defer it.Unlock()
	_, _ = it.stdout.WriteString("Content-Length: ")
	_, _ = it.stdout.WriteString(strconv.Itoa(len(err_json)))
	_, _ = it.stdout.WriteString("\r\n\r\n")
	_, _ = it.stdout.Write(err_json)
}

func (it *Server) send(methodName string, params any, isReq bool, onResp func(any)) {
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

// ServeForever keeps reading and handling LSP JSON-RPC messages incoming over `os.Stdin`
// until reading from `os.Stdin` fails, then returns that IO read error.
func (it *Server) ServeForever() error {
	const buf_cap = 1024 * 1024

	it.stdout = bufio.NewWriterSize(os.Stdout, buf_cap)
	it.waiters = map[string]func(any){}

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer(make([]byte, buf_cap), buf_cap)
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

	{ // users shouldn't have to set up no-op handlers for these routine teardown lifecycle messages:
		old_shutdown, old_exit, old_initialized := it.On_shutdown, it.On_exit, it.On_initialized
		it.On_shutdown = func(params *Void) (any, error) {
			if old_shutdown != nil {
				return old_shutdown(params)
			}
			return nil, nil
		}
		it.On_exit = func(params *Void) (any, error) {
			if old_exit != nil {
				return old_exit(params)
			}
			os.Exit(0)
			return nil, nil
		}
		it.On_initialized = func(params *InitializedParams) (any, error) {
			if it.On_workspace_didChangeWatchedFiles != nil {
				it.Request_client_registerCapability(&RegistrationParams{
					Registrations: []Registration{
						{Method: "workspace/didChangeWatchedFiles", Id: "workspace/didChangeWatchedFiles",
							RegisterOptions: DidChangeWatchedFilesRegistrationOptions{Watchers: []FileSystemWatcher{
								{Kind: WatchKindChange | WatchKindCreate | WatchKindDelete,
									GlobPattern: GlobPattern{Pattern: ptr(String("**/*"))}}}}},
					},
				}, func(*Void) {})
			}
			if old_initialized != nil {
				return old_initialized(params)
			}
			return nil, nil
		}
	}

	for stdin.Scan() {
		err := it.handleIncoming(stdin.Bytes())
		if err != nil {
			it.sendErrMsg(err)
		}
	}
	return stdin.Err()
}
