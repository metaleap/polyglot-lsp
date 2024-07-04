// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-lsp/gen/cmd/gen_lsp
package lsp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Server struct {
	stdout  *bufio.Writer
	waiters map[string]func()

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

	// The initialize request is sent from the client to the server.
	// It is sent once as the request after starting up the server.
	// The requests parameter is of type {@link InitializeParams}
	// the response if of type {@link InitializeResult} of a Thenable that
	// resolves to such.
	On_initialize func(params *InitializeParams) (any, error)

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
		go handler()
	case "workspace/didChangeWorkspaceFolders":
		serverHandleIncoming(it, it.On_workspace_didChangeWorkspaceFolders, msg_id, raw["params"])
	case "window/workDoneProgress/cancel":
		serverHandleIncoming(it, it.On_window_workDoneProgress_cancel, msg_id, raw["params"])
	case "workspace/didCreateFiles":
		serverHandleIncoming(it, it.On_workspace_didCreateFiles, msg_id, raw["params"])
	case "workspace/didRenameFiles":
		serverHandleIncoming(it, it.On_workspace_didRenameFiles, msg_id, raw["params"])
	case "workspace/didDeleteFiles":
		serverHandleIncoming(it, it.On_workspace_didDeleteFiles, msg_id, raw["params"])
	case "notebookDocument/didOpen":
		serverHandleIncoming(it, it.On_notebookDocument_didOpen, msg_id, raw["params"])
	case "notebookDocument/didChange":
		serverHandleIncoming(it, it.On_notebookDocument_didChange, msg_id, raw["params"])
	case "notebookDocument/didSave":
		serverHandleIncoming(it, it.On_notebookDocument_didSave, msg_id, raw["params"])
	case "notebookDocument/didClose":
		serverHandleIncoming(it, it.On_notebookDocument_didClose, msg_id, raw["params"])
	case "initialized":
		serverHandleIncoming(it, it.On_initialized, msg_id, raw["params"])
	case "exit":
		serverHandleIncoming(it, it.On_exit, msg_id, raw["params"])
	case "workspace/didChangeConfiguration":
		serverHandleIncoming(it, it.On_workspace_didChangeConfiguration, msg_id, raw["params"])
	case "textDocument/didOpen":
		serverHandleIncoming(it, it.On_textDocument_didOpen, msg_id, raw["params"])
	case "textDocument/didChange":
		serverHandleIncoming(it, it.On_textDocument_didChange, msg_id, raw["params"])
	case "textDocument/didClose":
		serverHandleIncoming(it, it.On_textDocument_didClose, msg_id, raw["params"])
	case "textDocument/didSave":
		serverHandleIncoming(it, it.On_textDocument_didSave, msg_id, raw["params"])
	case "textDocument/willSave":
		serverHandleIncoming(it, it.On_textDocument_willSave, msg_id, raw["params"])
	case "workspace/didChangeWatchedFiles":
		serverHandleIncoming(it, it.On_workspace_didChangeWatchedFiles, msg_id, raw["params"])
	case "$/setTrace":
		serverHandleIncoming(it, it.On___setTrace, msg_id, raw["params"])
	case "$/cancelRequest":
		serverHandleIncoming(it, it.On___cancelRequest, msg_id, raw["params"])
	case "$/progress":
		serverHandleIncoming(it, it.On___progress, msg_id, raw["params"])
	case "textDocument/implementation":
		serverHandleIncoming(it, it.On_textDocument_implementation, msg_id, raw["params"])
	case "textDocument/typeDefinition":
		serverHandleIncoming(it, it.On_textDocument_typeDefinition, msg_id, raw["params"])
	case "textDocument/documentColor":
		serverHandleIncoming(it, it.On_textDocument_documentColor, msg_id, raw["params"])
	case "textDocument/colorPresentation":
		serverHandleIncoming(it, it.On_textDocument_colorPresentation, msg_id, raw["params"])
	case "textDocument/foldingRange":
		serverHandleIncoming(it, it.On_textDocument_foldingRange, msg_id, raw["params"])
	case "textDocument/declaration":
		serverHandleIncoming(it, it.On_textDocument_declaration, msg_id, raw["params"])
	case "textDocument/selectionRange":
		serverHandleIncoming(it, it.On_textDocument_selectionRange, msg_id, raw["params"])
	case "textDocument/prepareCallHierarchy":
		serverHandleIncoming(it, it.On_textDocument_prepareCallHierarchy, msg_id, raw["params"])
	case "callHierarchy/incomingCalls":
		serverHandleIncoming(it, it.On_callHierarchy_incomingCalls, msg_id, raw["params"])
	case "callHierarchy/outgoingCalls":
		serverHandleIncoming(it, it.On_callHierarchy_outgoingCalls, msg_id, raw["params"])
	case "textDocument/semanticTokens/full":
		serverHandleIncoming(it, it.On_textDocument_semanticTokens_full, msg_id, raw["params"])
	case "textDocument/semanticTokens/full/delta":
		serverHandleIncoming(it, it.On_textDocument_semanticTokens_full_delta, msg_id, raw["params"])
	case "textDocument/semanticTokens/range":
		serverHandleIncoming(it, it.On_textDocument_semanticTokens_range, msg_id, raw["params"])
	case "textDocument/linkedEditingRange":
		serverHandleIncoming(it, it.On_textDocument_linkedEditingRange, msg_id, raw["params"])
	case "workspace/willCreateFiles":
		serverHandleIncoming(it, it.On_workspace_willCreateFiles, msg_id, raw["params"])
	case "workspace/willRenameFiles":
		serverHandleIncoming(it, it.On_workspace_willRenameFiles, msg_id, raw["params"])
	case "workspace/willDeleteFiles":
		serverHandleIncoming(it, it.On_workspace_willDeleteFiles, msg_id, raw["params"])
	case "textDocument/moniker":
		serverHandleIncoming(it, it.On_textDocument_moniker, msg_id, raw["params"])
	case "textDocument/prepareTypeHierarchy":
		serverHandleIncoming(it, it.On_textDocument_prepareTypeHierarchy, msg_id, raw["params"])
	case "typeHierarchy/supertypes":
		serverHandleIncoming(it, it.On_typeHierarchy_supertypes, msg_id, raw["params"])
	case "typeHierarchy/subtypes":
		serverHandleIncoming(it, it.On_typeHierarchy_subtypes, msg_id, raw["params"])
	case "textDocument/inlineValue":
		serverHandleIncoming(it, it.On_textDocument_inlineValue, msg_id, raw["params"])
	case "textDocument/inlayHint":
		serverHandleIncoming(it, it.On_textDocument_inlayHint, msg_id, raw["params"])
	case "inlayHint/resolve":
		serverHandleIncoming(it, it.On_inlayHint_resolve, msg_id, raw["params"])
	case "textDocument/diagnostic":
		serverHandleIncoming(it, it.On_textDocument_diagnostic, msg_id, raw["params"])
	case "workspace/diagnostic":
		serverHandleIncoming(it, it.On_workspace_diagnostic, msg_id, raw["params"])
	case "initialize":
		serverHandleIncoming(it, it.On_initialize, msg_id, raw["params"])
	case "shutdown":
		serverHandleIncoming(it, it.On_shutdown, msg_id, raw["params"])
	case "textDocument/willSaveWaitUntil":
		serverHandleIncoming(it, it.On_textDocument_willSaveWaitUntil, msg_id, raw["params"])
	case "textDocument/completion":
		serverHandleIncoming(it, it.On_textDocument_completion, msg_id, raw["params"])
	case "completionItem/resolve":
		serverHandleIncoming(it, it.On_completionItem_resolve, msg_id, raw["params"])
	case "textDocument/hover":
		serverHandleIncoming(it, it.On_textDocument_hover, msg_id, raw["params"])
	case "textDocument/signatureHelp":
		serverHandleIncoming(it, it.On_textDocument_signatureHelp, msg_id, raw["params"])
	case "textDocument/definition":
		serverHandleIncoming(it, it.On_textDocument_definition, msg_id, raw["params"])
	case "textDocument/references":
		serverHandleIncoming(it, it.On_textDocument_references, msg_id, raw["params"])
	case "textDocument/documentHighlight":
		serverHandleIncoming(it, it.On_textDocument_documentHighlight, msg_id, raw["params"])
	case "textDocument/documentSymbol":
		serverHandleIncoming(it, it.On_textDocument_documentSymbol, msg_id, raw["params"])
	case "textDocument/codeAction":
		serverHandleIncoming(it, it.On_textDocument_codeAction, msg_id, raw["params"])
	case "codeAction/resolve":
		serverHandleIncoming(it, it.On_codeAction_resolve, msg_id, raw["params"])
	case "workspaceSymbol/resolve":
		serverHandleIncoming(it, it.On_workspaceSymbol_resolve, msg_id, raw["params"])
	case "textDocument/codeLens":
		serverHandleIncoming(it, it.On_textDocument_codeLens, msg_id, raw["params"])
	case "codeLens/resolve":
		serverHandleIncoming(it, it.On_codeLens_resolve, msg_id, raw["params"])
	case "textDocument/documentLink":
		serverHandleIncoming(it, it.On_textDocument_documentLink, msg_id, raw["params"])
	case "documentLink/resolve":
		serverHandleIncoming(it, it.On_documentLink_resolve, msg_id, raw["params"])
	case "textDocument/formatting":
		serverHandleIncoming(it, it.On_textDocument_formatting, msg_id, raw["params"])
	case "textDocument/rangeFormatting":
		serverHandleIncoming(it, it.On_textDocument_rangeFormatting, msg_id, raw["params"])
	case "textDocument/onTypeFormatting":
		serverHandleIncoming(it, it.On_textDocument_onTypeFormatting, msg_id, raw["params"])
	case "textDocument/rename":
		serverHandleIncoming(it, it.On_textDocument_rename, msg_id, raw["params"])
	case "textDocument/prepareRename":
		serverHandleIncoming(it, it.On_textDocument_prepareRename, msg_id, raw["params"])
	case "workspace/executeCommand":
		serverHandleIncoming(it, it.On_workspace_executeCommand, msg_id, raw["params"])
	default: // msg is an incoming Request or Notification
		/* workspace/didChangeWorkspaceFolders */
		/* window/workDoneProgress/cancel */
		/* workspace/didCreateFiles */
		/* workspace/didRenameFiles */
		/* workspace/didDeleteFiles */
		/* notebookDocument/didOpen */
		/* notebookDocument/didChange */
		/* notebookDocument/didSave */
		/* notebookDocument/didClose */
		/* initialized */
		/* exit */
		/* workspace/didChangeConfiguration */
		/* window/showMessage */
		/* window/logMessage */
		/* telemetry/event */
		/* textDocument/didOpen */
		/* textDocument/didChange */
		/* textDocument/didClose */
		/* textDocument/didSave */
		/* textDocument/willSave */
		/* workspace/didChangeWatchedFiles */
		/* textDocument/publishDiagnostics */
		/* $/setTrace */
		/* $/logTrace */
		/* $/cancelRequest */
		/* $/progress */
		/* textDocument/implementation */
		/* textDocument/typeDefinition */
		/* workspace/workspaceFolders */
		/* workspace/configuration */
		/* textDocument/documentColor */
		/* textDocument/colorPresentation */
		/* textDocument/foldingRange */
		/* textDocument/declaration */
		/* textDocument/selectionRange */
		/* window/workDoneProgress/create */
		/* textDocument/prepareCallHierarchy */
		/* callHierarchy/incomingCalls */
		/* callHierarchy/outgoingCalls */
		/* textDocument/semanticTokens/full */
		/* textDocument/semanticTokens/full/delta */
		/* textDocument/semanticTokens/range */
		/* workspace/semanticTokens/refresh */
		/* window/showDocument */
		/* textDocument/linkedEditingRange */
		/* workspace/willCreateFiles */
		/* workspace/willRenameFiles */
		/* workspace/willDeleteFiles */
		/* textDocument/moniker */
		/* textDocument/prepareTypeHierarchy */
		/* typeHierarchy/supertypes */
		/* typeHierarchy/subtypes */
		/* textDocument/inlineValue */
		/* workspace/inlineValue/refresh */
		/* textDocument/inlayHint */
		/* inlayHint/resolve */
		/* workspace/inlayHint/refresh */
		/* textDocument/diagnostic */
		/* workspace/diagnostic */
		/* workspace/diagnostic/refresh */
		/* client/registerCapability */
		/* client/unregisterCapability */
		/* initialize */
		/* shutdown */
		/* window/showMessageRequest */
		/* textDocument/willSaveWaitUntil */
		/* textDocument/completion */
		/* completionItem/resolve */
		/* textDocument/hover */
		/* textDocument/signatureHelp */
		/* textDocument/definition */
		/* textDocument/references */
		/* textDocument/documentHighlight */
		/* textDocument/documentSymbol */
		/* textDocument/codeAction */
		/* codeAction/resolve */
		/* workspaceSymbol/resolve */
		/* textDocument/codeLens */
		/* codeLens/resolve */
		/* workspace/codeLens/refresh */
		/* textDocument/documentLink */
		/* documentLink/resolve */
		/* textDocument/formatting */
		/* textDocument/rangeFormatting */
		/* textDocument/onTypeFormatting */
		/* textDocument/rename */
		/* textDocument/prepareRename */
		/* workspace/executeCommand */
		/* workspace/applyEdit */
	}

	return nil
}

func serverHandleIncoming[T any](it *Server, handler func(*T) (any, error), msgIdMaybe any, msgParams any) {
	var req_id string
	if msgIdMaybe != nil {
		req_id = fmt.Sprintf("%v", msgIdMaybe)
	}
	var params T
	if msgParams != nil {
		json_bytes, _ := json.Marshal(msgParams)
		if err := json.Unmarshal(json_bytes, &params); err != nil {
			it.sendErrMsg(&jsonRpcError{Code: -32700, Message: err.Error()})
			return
		}
	}
}

func (it *Server) sendErrMsg(err any) {
	if err == nil {
		return
	}
	var json_rpc_err_msg *jsonRpcError
	if json_rpc_err_msg, _ = err.(*jsonRpcError); json_rpc_err_msg == nil {
		json_rpc_err_msg = &jsonRpcError{Code: -32603, Message: fmt.Sprintf("%v", err)}
	}
	err_json, _ := json.Marshal(json_rpc_err_msg)
	_, _ = it.stdout.WriteString("Content-Length: ")
	_, _ = it.stdout.WriteString(strconv.Itoa(len(err_json)))
	_, _ = it.stdout.WriteString("\r\n\r\n")
	_, _ = it.stdout.Write(err_json)
}

// ServeForever keeps reading and handling LSP JSON-RPC messages incoming over `os.Stdin`
// until reading from `os.Stdin` fails, then returns that IO read error.
func (it *Server) ServeForever() error {
	const buf_cap = 1024 * 1024

	it.stdout = bufio.NewWriterSize(os.Stdout, buf_cap)
	it.waiters = map[string]func(){}

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

	for stdin.Scan() {
		err := it.handleIncoming(stdin.Bytes())
		if err != nil {
			it.sendErrMsg(err)
		}
	}
	return stdin.Err()
}