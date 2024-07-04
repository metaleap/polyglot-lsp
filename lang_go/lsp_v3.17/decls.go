// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-lsp/gen/cmd/gen_lsp
package lsp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type DocumentURI = URI
type URI = String

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

type Server struct {
	stdout  *bufio.Writer
	waiters map[string]func()

	// The `workspace/didChangeWorkspaceFolders` notification is sent from the client to the server when the workspace
	// folder configuration changes.
	On_workspace_didChangeWorkspaceFolders func()

	// The `window/workDoneProgress/cancel` notification is sent from  the client to the server to cancel a progress
	// initiated on the server side.
	On_window_workDoneProgress_cancel func()

	// The did create files notification is sent from the client to the server when
	// files were created from within the client.
	//
	// @since 3.16.0
	On_workspace_didCreateFiles func()

	// The did rename files notification is sent from the client to the server when
	// files were renamed from within the client.
	//
	// @since 3.16.0
	On_workspace_didRenameFiles func()

	// The will delete files request is sent from the client to the server before files are actually
	// deleted as long as the deletion is triggered from within the client.
	//
	// @since 3.16.0
	On_workspace_didDeleteFiles func()

	// A notification sent when a notebook opens.
	//
	// @since 3.17.0
	On_notebookDocument_didOpen func()

	//
	On_notebookDocument_didChange func()

	// A notification sent when a notebook document is saved.
	//
	// @since 3.17.0
	On_notebookDocument_didSave func()

	// A notification sent when a notebook closes.
	//
	// @since 3.17.0
	On_notebookDocument_didClose func()

	// The initialized notification is sent from the client to the
	// server after the client is fully initialized and the server
	// is allowed to send requests from the server to the client.
	On_initialized func()

	// The exit event is sent from the client to the server to
	// ask the server to exit its process.
	On_exit func()

	// The configuration change notification is sent from the client to the server
	// when the client's configuration has changed. The notification contains
	// the changed configuration as defined by the language client.
	On_workspace_didChangeConfiguration func()

	// The document open notification is sent from the client to the server to signal
	// newly opened text documents. The document's truth is now managed by the client
	// and the server must not try to read the document's truth using the document's
	// uri. Open in this sense means it is managed by the client. It doesn't necessarily
	// mean that its content is presented in an editor. An open notification must not
	// be sent more than once without a corresponding close notification send before.
	// This means open and close notification must be balanced and the max open count
	// is one.
	On_textDocument_didOpen func()

	// The document change notification is sent from the client to the server to signal
	// changes to a text document.
	On_textDocument_didChange func()

	// The document close notification is sent from the client to the server when
	// the document got closed in the client. The document's truth now exists where
	// the document's uri points to (e.g. if the document's uri is a file uri the
	// truth now exists on disk). As with the open notification the close notification
	// is about managing the document's content. Receiving a close notification
	// doesn't mean that the document was open in an editor before. A close
	// notification requires a previous open notification to be sent.
	On_textDocument_didClose func()

	// The document save notification is sent from the client to the server when
	// the document got saved in the client.
	On_textDocument_didSave func()

	// A document will save notification is sent from the client to the server before
	// the document is actually saved.
	On_textDocument_willSave func()

	// The watched files notification is sent from the client to the server when
	// the client detects changes to file watched by the language client.
	On_workspace_didChangeWatchedFiles func()

	//
	On___setTrace func()

	//
	On___cancelRequest func()

	//
	On___progress func()

	// A request to resolve the implementation locations of a symbol at a given text
	// document position. The request's parameter is of type {@link TextDocumentPositionParams}
	// the response is of type {@link Definition} or a Thenable that resolves to such.
	On_textDocument_implementation func()

	// A request to resolve the type definition locations of a symbol at a given text
	// document position. The request's parameter is of type {@link TextDocumentPositionParams}
	// the response is of type {@link Definition} or a Thenable that resolves to such.
	On_textDocument_typeDefinition func()

	// A request to list all color symbols found in a given text document. The request's
	// parameter is of type {@link DocumentColorParams} the
	// response is of type {@link ColorInformation ColorInformation[]} or a Thenable
	// that resolves to such.
	On_textDocument_documentColor func()

	// A request to list all presentation for a color. The request's
	// parameter is of type {@link ColorPresentationParams} the
	// response is of type {@link ColorInformation ColorInformation[]} or a Thenable
	// that resolves to such.
	On_textDocument_colorPresentation func()

	// A request to provide folding ranges in a document. The request's
	// parameter is of type {@link FoldingRangeParams}, the
	// response is of type {@link FoldingRangeList} or a Thenable
	// that resolves to such.
	On_textDocument_foldingRange func()

	// A request to resolve the type definition locations of a symbol at a given text
	// document position. The request's parameter is of type {@link TextDocumentPositionParams}
	// the response is of type {@link Declaration} or a typed array of {@link DeclarationLink}
	// or a Thenable that resolves to such.
	On_textDocument_declaration func()

	// A request to provide selection ranges in a document. The request's
	// parameter is of type {@link SelectionRangeParams}, the
	// response is of type {@link SelectionRange SelectionRange[]} or a Thenable
	// that resolves to such.
	On_textDocument_selectionRange func()

	// A request to result a `CallHierarchyItem` in a document at a given position.
	// Can be used as an input to an incoming or outgoing call hierarchy.
	//
	// @since 3.16.0
	On_textDocument_prepareCallHierarchy func()

	// A request to resolve the incoming calls for a given `CallHierarchyItem`.
	//
	// @since 3.16.0
	On_callHierarchy_incomingCalls func()

	// A request to resolve the outgoing calls for a given `CallHierarchyItem`.
	//
	// @since 3.16.0
	On_callHierarchy_outgoingCalls func()

	// @since 3.16.0
	On_textDocument_semanticTokens_full func()

	// @since 3.16.0
	On_textDocument_semanticTokens_full_delta func()

	// @since 3.16.0
	On_textDocument_semanticTokens_range func()

	// A request to provide ranges that can be edited together.
	//
	// @since 3.16.0
	On_textDocument_linkedEditingRange func()

	// The will create files request is sent from the client to the server before files are actually
	// created as long as the creation is triggered from within the client.
	//
	// The request can return a `WorkspaceEdit` which will be applied to workspace before the
	// files are created. Hence the `WorkspaceEdit` can not manipulate the content of the file
	// to be created.
	//
	// @since 3.16.0
	On_workspace_willCreateFiles func()

	// The will rename files request is sent from the client to the server before files are actually
	// renamed as long as the rename is triggered from within the client.
	//
	// @since 3.16.0
	On_workspace_willRenameFiles func()

	// The did delete files notification is sent from the client to the server when
	// files were deleted from within the client.
	//
	// @since 3.16.0
	On_workspace_willDeleteFiles func()

	// A request to get the moniker of a symbol at a given text document position.
	// The request parameter is of type {@link TextDocumentPositionParams}.
	// The response is of type {@link Moniker Moniker[]} or `null`.
	On_textDocument_moniker func()

	// A request to result a `TypeHierarchyItem` in a document at a given position.
	// Can be used as an input to a subtypes or supertypes type hierarchy.
	//
	// @since 3.17.0
	On_textDocument_prepareTypeHierarchy func()

	// A request to resolve the supertypes for a given `TypeHierarchyItem`.
	//
	// @since 3.17.0
	On_typeHierarchy_supertypes func()

	// A request to resolve the subtypes for a given `TypeHierarchyItem`.
	//
	// @since 3.17.0
	On_typeHierarchy_subtypes func()

	// A request to provide inline values in a document. The request's parameter is of
	// type {@link InlineValueParams}, the response is of type
	// {@link InlineValue InlineValue[]} or a Thenable that resolves to such.
	//
	// @since 3.17.0
	On_textDocument_inlineValue func()

	// A request to provide inlay hints in a document. The request's parameter is of
	// type {@link InlayHintsParams}, the response is of type
	// {@link InlayHint InlayHint[]} or a Thenable that resolves to such.
	//
	// @since 3.17.0
	On_textDocument_inlayHint func()

	// A request to resolve additional properties for an inlay hint.
	// The request's parameter is of type {@link InlayHint}, the response is
	// of type {@link InlayHint} or a Thenable that resolves to such.
	//
	// @since 3.17.0
	On_inlayHint_resolve func()

	// The document diagnostic request definition.
	//
	// @since 3.17.0
	On_textDocument_diagnostic func()

	// The workspace diagnostic request definition.
	//
	// @since 3.17.0
	On_workspace_diagnostic func()

	// A request to provide inline completions in a document. The request's parameter is of
	// type {@link InlineCompletionParams}, the response is of type
	// {@link InlineCompletion InlineCompletion[]} or a Thenable that resolves to such.
	//
	// @since 3.18.0
	// @proposed
	On_textDocument_inlineCompletion func()

	// The initialize request is sent from the client to the server.
	// It is sent once as the request after starting up the server.
	// The requests parameter is of type {@link InitializeParams}
	// the response if of type {@link InitializeResult} of a Thenable that
	// resolves to such.
	On_initialize func()

	// A shutdown request is sent from the client to the server.
	// It is sent once when the client decides to shutdown the
	// server. The only notification that is sent after a shutdown request
	// is the exit event.
	On_shutdown func()

	// A document will save request is sent from the client to the server before
	// the document is actually saved. The request can return an array of TextEdits
	// which will be applied to the text document before it is saved. Please note that
	// clients might drop results if computing the text edits took too long or if a
	// server constantly fails on this request. This is done to keep the save fast and
	// reliable.
	On_textDocument_willSaveWaitUntil func()

	// Request to request completion at a given text document position. The request's
	// parameter is of type {@link TextDocumentPosition} the response
	// is of type {@link CompletionItem CompletionItem[]} or {@link CompletionList}
	// or a Thenable that resolves to such.
	//
	// The request can delay the computation of the {@link CompletionItem.detail `detail`}
	// and {@link CompletionItem.documentation `documentation`} properties to the `completionItem/resolve`
	// request. However, properties that are needed for the initial sorting and filtering, like `sortText`,
	// `filterText`, `insertText`, and `textEdit`, must not be changed during resolve.
	On_textDocument_completion func()

	// Request to resolve additional information for a given completion item.The request's
	// parameter is of type {@link CompletionItem} the response
	// is of type {@link CompletionItem} or a Thenable that resolves to such.
	On_completionItem_resolve func()

	// Request to request hover information at a given text document position. The request's
	// parameter is of type {@link TextDocumentPosition} the response is of
	// type {@link Hover} or a Thenable that resolves to such.
	On_textDocument_hover func()

	//
	On_textDocument_signatureHelp func()

	// A request to resolve the definition location of a symbol at a given text
	// document position. The request's parameter is of type {@link TextDocumentPosition}
	// the response is of either type {@link Definition} or a typed array of
	// {@link DefinitionLink} or a Thenable that resolves to such.
	On_textDocument_definition func()

	// A request to resolve project-wide references for the symbol denoted
	// by the given text document position. The request's parameter is of
	// type {@link ReferenceParams} the response is of type
	// {@link Location Location[]} or a Thenable that resolves to such.
	On_textDocument_references func()

	// Request to resolve a {@link DocumentHighlight} for a given
	// text document position. The request's parameter is of type {@link TextDocumentPosition}
	// the request response is an array of type {@link DocumentHighlight}
	// or a Thenable that resolves to such.
	On_textDocument_documentHighlight func()

	// A request to list all symbols found in a given text document. The request's
	// parameter is of type {@link TextDocumentIdentifier} the
	// response is of type {@link SymbolInformation SymbolInformation[]} or a Thenable
	// that resolves to such.
	On_textDocument_documentSymbol func()

	// A request to provide commands for the given text document and range.
	On_textDocument_codeAction func()

	// Request to resolve additional information for a given code action.The request's
	// parameter is of type {@link CodeAction} the response
	// is of type {@link CodeAction} or a Thenable that resolves to such.
	On_codeAction_resolve func()

	// A request to list project-wide symbols matching the query string given
	// by the {@link WorkspaceSymbolParams}. The response is
	// of type {@link SymbolInformation SymbolInformation[]} or a Thenable that
	// resolves to such.
	//
	// @since 3.17.0 - support for WorkspaceSymbol in the returned data. Clients
	//  need to advertise support for WorkspaceSymbols via the client capability
	//  `workspace.symbol.resolveSupport`.
	//
	On_workspace_symbol func()

	// A request to resolve the range inside the workspace
	// symbol's location.
	//
	// @since 3.17.0
	On_workspaceSymbol_resolve func()

	// A request to provide code lens for the given text document.
	On_textDocument_codeLens func()

	// A request to resolve a command for a given code lens.
	On_codeLens_resolve func()

	// A request to provide document links
	On_textDocument_documentLink func()

	// Request to resolve additional information for a given document link. The request's
	// parameter is of type {@link DocumentLink} the response
	// is of type {@link DocumentLink} or a Thenable that resolves to such.
	On_documentLink_resolve func()

	// A request to format a whole document.
	On_textDocument_formatting func()

	// A request to format a range in a document.
	On_textDocument_rangeFormatting func()

	// A request to format ranges in a document.
	//
	// @since 3.18.0
	// @proposed
	On_textDocument_rangesFormatting func()

	// A request to format a document on type.
	On_textDocument_onTypeFormatting func()

	// A request to rename a symbol.
	On_textDocument_rename func()

	// A request to test and perform the setup necessary for a rename.
	//
	// @since 3.16 - support for default behavior
	On_textDocument_prepareRename func()

	// A request send from the client to the server to execute a command. The request might return
	// a workspace edit which the client will apply to the workspace.
	On_workspace_executeCommand func()
}

type jsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (it *Server) handleIncoming(jsonRpcMsg []byte) *jsonRpcError {
	raw := map[string]any{}
	if err := json.Unmarshal(jsonRpcMsg, &raw); err != nil {
		return &jsonRpcError{Code: -32700, Message: err.Error()}
	}
	msg_id, msg_method, msg_err_code := fmt.Sprintf("%v", raw["id"]), raw["method"], raw["code"]

	if msg_err_code != nil { // received an error Response
		println(string(jsonRpcMsg)) // goes to stderr
		return nil
	}

	if msg_method != nil { // msg is an incoming Request or Notification
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
		/* workspace/foldingRange/refresh */
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
		/* textDocument/inlineCompletion */
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
		/* workspace/symbol */
		/* workspaceSymbol/resolve */
		/* textDocument/codeLens */
		/* codeLens/resolve */
		/* workspace/codeLens/refresh */
		/* textDocument/documentLink */
		/* documentLink/resolve */
		/* textDocument/formatting */
		/* textDocument/rangeFormatting */
		/* textDocument/rangesFormatting */
		/* textDocument/onTypeFormatting */
		/* textDocument/rename */
		/* textDocument/prepareRename */
		/* workspace/executeCommand */
		/* workspace/applyEdit */
	} else { // msg is an incoming Response
		handler := it.waiters[msg_id]
		delete(it.waiters, msg_id)
		go handler()
	}

	return nil
}

func (it *Server) Serve() error {
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
			err_json, _ := json.Marshal(err)
			_, _ = it.stdout.WriteString("Content-Length: ")
			_, _ = it.stdout.WriteString(strconv.Itoa(len(err_json)))
			_, _ = it.stdout.WriteString("\r\n\r\n")
			_, _ = it.stdout.Write(err_json)
		}
	}
	return stdin.Err()
}
