// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-lsp/gen/cmd/gen_lsp
package lsp

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
