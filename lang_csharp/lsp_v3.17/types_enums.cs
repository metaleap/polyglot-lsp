/// <summary>Language Server Protocol (LSP) v3.17 SDK for C#: auto-generated via github.com/metaleap/polyglot-vsx-and-lsp/gen/cmd/gen_lsp</summary>
namespace lsp;


/// <summary>
/// A set of predefined token types. This set is not fixed
/// an clients can specify additional token types via the
/// corresponding client capabilities.
/// 
/// @since 3.16.0
/// </summary>
public static class SemanticTokenTypes
{
    public static readonly string Namespace = "namespace";
    /// <summary>
    /// Represents a generic type. Acts as a fallback for types which can't be mapped to
    /// a specific type like class or enum.
    /// </summary>
    public static readonly string Type = "type";
    public static readonly string Class = "class";
    public static readonly string Enum = "enum";
    public static readonly string Interface = "interface";
    public static readonly string Struct = "struct";
    public static readonly string TypeParameter = "typeParameter";
    public static readonly string Parameter = "parameter";
    public static readonly string Variable = "variable";
    public static readonly string Property = "property";
    public static readonly string EnumMember = "enumMember";
    public static readonly string Event = "event";
    public static readonly string Function = "function";
    public static readonly string Method = "method";
    public static readonly string Macro = "macro";
    public static readonly string Keyword = "keyword";
    public static readonly string Modifier = "modifier";
    public static readonly string Comment = "comment";
    public static readonly string String = "string";
    public static readonly string Number = "number";
    public static readonly string Regexp = "regexp";
    public static readonly string Operator = "operator";
    /// <summary>
    /// @since 3.17.0
    /// </summary>
    public static readonly string Decorator = "decorator";
}

/// <summary>
/// A set of predefined token modifiers. This set is not fixed
/// an clients can specify additional token types via the
/// corresponding client capabilities.
/// 
/// @since 3.16.0
/// </summary>
public static class SemanticTokenModifiers
{
    public static readonly string Declaration = "declaration";
    public static readonly string Definition = "definition";
    public static readonly string Readonly = "readonly";
    public static readonly string Static = "static";
    public static readonly string Deprecated = "deprecated";
    public static readonly string Abstract = "abstract";
    public static readonly string Async = "async";
    public static readonly string Modification = "modification";
    public static readonly string Documentation = "documentation";
    public static readonly string DefaultLibrary = "defaultLibrary";
}

/// <summary>
/// The document diagnostic report kinds.
/// 
/// @since 3.17.0
/// </summary>
public static class DocumentDiagnosticReportKind
{
    /// <summary>
    /// A diagnostic report with a full
    /// set of problems.
    /// </summary>
    public static readonly string Full = "full";
    /// <summary>
    /// A report indicating that the last
    /// returned report is still accurate.
    /// </summary>
    public static readonly string Unchanged = "unchanged";
}

/// <summary>
/// Predefined error codes.
/// </summary>
public enum ErrorCodes
{
    ParseError = -32700,

    InvalidRequest = -32600,

    MethodNotFound = -32601,

    InvalidParams = -32602,

    InternalError = -32603,

    /// <summary>
    /// Error code indicating that a server received a notification or
    /// request before the server has received the `initialize` request.
    /// </summary>
    ServerNotInitialized = -32002,

    UnknownErrorCode = -32001,
}

public enum LSPErrorCodes
{
    /// <summary>
    /// A request failed but it was syntactically correct, e.g the
    /// method name was known and the parameters were valid. The error
    /// message should contain human readable information about why
    /// the request failed.
    /// 
    /// @since 3.17.0
    /// </summary>
    RequestFailed = -32803,

    /// <summary>
    /// The server cancelled the request. This error code should
    /// only be used for requests that explicitly support being
    /// server cancellable.
    /// 
    /// @since 3.17.0
    /// </summary>
    ServerCancelled = -32802,

    /// <summary>
    /// The server detected that the content of a document got
    /// modified outside normal conditions. A server should
    /// NOT send this error code if it detects a content change
    /// in it unprocessed messages. The result even computed
    /// on an older state might still be useful for the client.
    /// 
    /// If a client decides that a result is not of any use anymore
    /// the client should cancel the request.
    /// </summary>
    ContentModified = -32801,

    /// <summary>
    /// The client has canceled a request and a server as detected
    /// the cancel.
    /// </summary>
    RequestCancelled = -32800,
}

/// <summary>
/// A set of predefined range kinds.
/// </summary>
public static class FoldingRangeKind
{
    /// <summary>
    /// Folding range for a comment
    /// </summary>
    public static readonly string Comment = "comment";
    /// <summary>
    /// Folding range for an import or include
    /// </summary>
    public static readonly string Imports = "imports";
    /// <summary>
    /// Folding range for a region (e.g. `#region`)
    /// </summary>
    public static readonly string Region = "region";
}

/// <summary>
/// A symbol kind.
/// </summary>
public enum SymbolKind
{
    File = 1,

    Module = 2,

    Namespace = 3,

    Package = 4,

    Class = 5,

    Method = 6,

    Property = 7,

    Field = 8,

    Constructor = 9,

    Enum = 10,

    Interface = 11,

    Function = 12,

    Variable = 13,

    Constant = 14,

    String = 15,

    Number = 16,

    Boolean = 17,

    Array = 18,

    Object = 19,

    Key = 20,

    Null = 21,

    EnumMember = 22,

    Struct = 23,

    Event = 24,

    Operator = 25,

    TypeParameter = 26,
}

/// <summary>
/// Symbol tags are extra annotations that tweak the rendering of a symbol.
/// 
/// @since 3.16
/// </summary>
public enum SymbolTag
{
    /// <summary>
    /// Render a symbol as obsolete, usually using a strike-out.
    /// </summary>
    Deprecated = 1,
}

/// <summary>
/// Moniker uniqueness level to define scope of the moniker.
/// 
/// @since 3.16.0
/// </summary>
public static class UniquenessLevel
{
    /// <summary>
    /// The moniker is only unique inside a document
    /// </summary>
    public static readonly string Document = "document";
    /// <summary>
    /// The moniker is unique inside a project for which a dump got created
    /// </summary>
    public static readonly string Project = "project";
    /// <summary>
    /// The moniker is unique inside the group to which a project belongs
    /// </summary>
    public static readonly string Group = "group";
    /// <summary>
    /// The moniker is unique inside the moniker scheme.
    /// </summary>
    public static readonly string Scheme = "scheme";
    /// <summary>
    /// The moniker is globally unique
    /// </summary>
    public static readonly string Global = "global";
}

/// <summary>
/// The moniker kind.
/// 
/// @since 3.16.0
/// </summary>
public static class MonikerKind
{
    /// <summary>
    /// The moniker represent a symbol that is imported into a project
    /// </summary>
    public static readonly string Import = "import";
    /// <summary>
    /// The moniker represents a symbol that is exported from a project
    /// </summary>
    public static readonly string Export = "export";
    /// <summary>
    /// The moniker represents a symbol that is local to a project (e.g. a local
    /// variable of a function, a class not visible outside the project, ...)
    /// </summary>
    public static readonly string Local = "local";
}

/// <summary>
/// Inlay hint kinds.
/// 
/// @since 3.17.0
/// </summary>
public enum InlayHintKind
{
    /// <summary>
    /// An inlay hint that for a type annotation.
    /// </summary>
    Type = 1,

    /// <summary>
    /// An inlay hint that is for a parameter.
    /// </summary>
    Parameter = 2,
}

/// <summary>
/// The message type
/// </summary>
public enum MessageType
{
    /// <summary>
    /// An error message.
    /// </summary>
    Error = 1,

    /// <summary>
    /// A warning message.
    /// </summary>
    Warning = 2,

    /// <summary>
    /// An information message.
    /// </summary>
    Info = 3,

    /// <summary>
    /// A log message.
    /// </summary>
    Log = 4,
}

/// <summary>
/// Defines how the host (editor) should sync
/// document changes to the language server.
/// </summary>
public enum TextDocumentSyncKind
{
    /// <summary>
    /// Documents should not be synced at all.
    /// </summary>
    None = 0,

    /// <summary>
    /// Documents are synced by always sending the full content
    /// of the document.
    /// </summary>
    Full = 1,

    /// <summary>
    /// Documents are synced by sending the full content on open.
    /// After that only incremental updates to the document are
    /// send.
    /// </summary>
    Incremental = 2,
}

/// <summary>
/// Represents reasons why a text document is saved.
/// </summary>
public enum TextDocumentSaveReason
{
    /// <summary>
    /// Manually triggered, e.g. by the user pressing save, by starting debugging,
    /// or by an API call.
    /// </summary>
    Manual = 1,

    /// <summary>
    /// Automatic after a delay.
    /// </summary>
    AfterDelay = 2,

    /// <summary>
    /// When the editor lost focus.
    /// </summary>
    FocusOut = 3,
}

/// <summary>
/// The kind of a completion entry.
/// </summary>
public enum CompletionItemKind
{
    Text = 1,

    Method = 2,

    Function = 3,

    Constructor = 4,

    Field = 5,

    Variable = 6,

    Class = 7,

    Interface = 8,

    Module = 9,

    Property = 10,

    Unit = 11,

    Value = 12,

    Enum = 13,

    Keyword = 14,

    Snippet = 15,

    Color = 16,

    File = 17,

    Reference = 18,

    Folder = 19,

    EnumMember = 20,

    Constant = 21,

    Struct = 22,

    Event = 23,

    Operator = 24,

    TypeParameter = 25,
}

/// <summary>
/// Completion item tags are extra annotations that tweak the rendering of a completion
/// item.
/// 
/// @since 3.15.0
/// </summary>
public enum CompletionItemTag
{
    /// <summary>
    /// Render a completion as obsolete, usually using a strike-out.
    /// </summary>
    Deprecated = 1,
}

/// <summary>
/// Defines whether the insert text in a completion item should be interpreted as
/// plain text or a snippet.
/// </summary>
public enum InsertTextFormat
{
    /// <summary>
    /// The primary text to be inserted is treated as a plain string.
    /// </summary>
    PlainText = 1,

    /// <summary>
    /// The primary text to be inserted is treated as a snippet.
    /// 
    /// A snippet can define tab stops and placeholders with `$1`, `$2`
    /// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
    /// the end of the snippet. Placeholders with equal identifiers are linked,
    /// that is typing in one will update others too.
    /// 
    /// See also: https://microsoft.github.io/language-server-protocol/specifications/specification-current/#snippet_syntax
    /// </summary>
    Snippet = 2,
}

/// <summary>
/// How whitespace and indentation is handled during completion
/// item insertion.
/// 
/// @since 3.16.0
/// </summary>
public enum InsertTextMode
{
    /// <summary>
    /// The insertion or replace strings is taken as it is. If the
    /// value is multi line the lines below the cursor will be
    /// inserted using the indentation defined in the string value.
    /// The client will not apply any kind of adjustments to the
    /// string.
    /// </summary>
    AsIs = 1,

    /// <summary>
    /// The editor adjusts leading whitespace of new lines so that
    /// they match the indentation up to the cursor of the line for
    /// which the item is accepted.
    /// 
    /// Consider a line like this: <2tabs><cursor><3tabs>foo. Accepting a
    /// multi line completion item is indented using 2 tabs and all
    /// following lines inserted will be indented using 2 tabs as well.
    /// </summary>
    AdjustIndentation = 2,
}

/// <summary>
/// A document highlight kind.
/// </summary>
public enum DocumentHighlightKind
{
    /// <summary>
    /// A textual occurrence.
    /// </summary>
    Text = 1,

    /// <summary>
    /// Read-access of a symbol, like reading a variable.
    /// </summary>
    Read = 2,

    /// <summary>
    /// Write-access of a symbol, like writing to a variable.
    /// </summary>
    Write = 3,
}

/// <summary>
/// A set of predefined code action kinds
/// </summary>
public static class CodeActionKind
{
    /// <summary>
    /// Empty kind.
    /// </summary>
    public static readonly string Empty = "";
    /// <summary>
    /// Base kind for quickfix actions: 'quickfix'
    /// </summary>
    public static readonly string QuickFix = "quickfix";
    /// <summary>
    /// Base kind for refactoring actions: 'refactor'
    /// </summary>
    public static readonly string Refactor = "refactor";
    /// <summary>
    /// Base kind for refactoring extraction actions: 'refactor.extract'
    /// 
    /// Example extract actions:
    /// 
    /// - Extract method
    /// - Extract function
    /// - Extract variable
    /// - Extract interface from class
    /// - ...
    /// </summary>
    public static readonly string RefactorExtract = "refactor.extract";
    /// <summary>
    /// Base kind for refactoring inline actions: 'refactor.inline'
    /// 
    /// Example inline actions:
    /// 
    /// - Inline function
    /// - Inline variable
    /// - Inline constant
    /// - ...
    /// </summary>
    public static readonly string RefactorInline = "refactor.inline";
    /// <summary>
    /// Base kind for refactoring rewrite actions: 'refactor.rewrite'
    /// 
    /// Example rewrite actions:
    /// 
    /// - Convert JavaScript function to class
    /// - Add or remove parameter
    /// - Encapsulate field
    /// - Make method static
    /// - Move method to base class
    /// - ...
    /// </summary>
    public static readonly string RefactorRewrite = "refactor.rewrite";
    /// <summary>
    /// Base kind for source actions: `source`
    /// 
    /// Source code actions apply to the entire file.
    /// </summary>
    public static readonly string Source = "source";
    /// <summary>
    /// Base kind for an organize imports source action: `source.organizeImports`
    /// </summary>
    public static readonly string SourceOrganizeImports = "source.organizeImports";
    /// <summary>
    /// Base kind for auto-fix source actions: `source.fixAll`.
    /// 
    /// Fix all actions automatically fix errors that have a clear fix that do not require user input.
    /// They should not suppress errors or perform unsafe fixes such as generating new types or classes.
    /// 
    /// @since 3.15.0
    /// </summary>
    public static readonly string SourceFixAll = "source.fixAll";
}

public static class TraceValues
{
    /// <summary>
    /// Turn tracing off.
    /// </summary>
    public static readonly string Off = "off";
    /// <summary>
    /// Trace messages only.
    /// </summary>
    public static readonly string Messages = "messages";
    /// <summary>
    /// Verbose message tracing.
    /// </summary>
    public static readonly string Verbose = "verbose";
}

/// <summary>
/// Describes the content type that a client supports in various
/// result literals like `Hover`, `ParameterInfo` or `CompletionItem`.
/// 
/// Please note that `MarkupKinds` must not start with a `$`. This kinds
/// are reserved for internal usage.
/// </summary>
public static class MarkupKind
{
    /// <summary>
    /// Plain text is supported as a content format
    /// </summary>
    public static readonly string PlainText = "plaintext";
    /// <summary>
    /// Markdown is supported as a content format
    /// </summary>
    public static readonly string Markdown = "markdown";
}

/// <summary>
/// A set of predefined position encoding kinds.
/// 
/// @since 3.17.0
/// </summary>
public static class PositionEncodingKind
{
    /// <summary>
    /// Character offsets count UTF-8 code units.
    /// </summary>
    public static readonly string UTF8 = "utf-8";
    /// <summary>
    /// Character offsets count UTF-16 code units.
    /// 
    /// This is the default and must always be supported
    /// by servers
    /// </summary>
    public static readonly string UTF16 = "utf-16";
    /// <summary>
    /// Character offsets count UTF-32 code units.
    /// 
    /// Implementation note: these are the same as Unicode code points,
    /// so this `PositionEncodingKind` may also be used for an
    /// encoding-agnostic representation of character offsets.
    /// </summary>
    public static readonly string UTF32 = "utf-32";
}

/// <summary>
/// The file event type
/// </summary>
public enum FileChangeType
{
    /// <summary>
    /// The file got created.
    /// </summary>
    Created = 1,

    /// <summary>
    /// The file got changed.
    /// </summary>
    Changed = 2,

    /// <summary>
    /// The file got deleted.
    /// </summary>
    Deleted = 3,
}

public enum WatchKind
{
    /// <summary>
    /// Interested in create events.
    /// </summary>
    Create = 1,

    /// <summary>
    /// Interested in change events
    /// </summary>
    Change = 2,

    /// <summary>
    /// Interested in delete events
    /// </summary>
    Delete = 4,
}

/// <summary>
/// The diagnostic's severity.
/// </summary>
public enum DiagnosticSeverity
{
    /// <summary>
    /// Reports an error.
    /// </summary>
    Error = 1,

    /// <summary>
    /// Reports a warning.
    /// </summary>
    Warning = 2,

    /// <summary>
    /// Reports an information.
    /// </summary>
    Information = 3,

    /// <summary>
    /// Reports a hint.
    /// </summary>
    Hint = 4,
}

/// <summary>
/// The diagnostic tags.
/// 
/// @since 3.15.0
/// </summary>
public enum DiagnosticTag
{
    /// <summary>
    /// Unused or unnecessary code.
    /// 
    /// Clients are allowed to render diagnostics with this tag faded out instead of having
    /// an error squiggle.
    /// </summary>
    Unnecessary = 1,

    /// <summary>
    /// Deprecated or obsolete code.
    /// 
    /// Clients are allowed to rendered diagnostics with this tag strike through.
    /// </summary>
    Deprecated = 2,
}

/// <summary>
/// How a completion was triggered
/// </summary>
public enum CompletionTriggerKind
{
    /// <summary>
    /// Completion was triggered by typing an identifier (24x7 code
    /// complete), manual invocation (e.g Ctrl+Space) or via API.
    /// </summary>
    Invoked = 1,

    /// <summary>
    /// Completion was triggered by a trigger character specified by
    /// the `triggerCharacters` properties of the `CompletionRegistrationOptions`.
    /// </summary>
    TriggerCharacter = 2,

    /// <summary>
    /// Completion was re-triggered as current completion list is incomplete
    /// </summary>
    TriggerForIncompleteCompletions = 3,
}

/// <summary>
/// How a signature help was triggered.
/// 
/// @since 3.15.0
/// </summary>
public enum SignatureHelpTriggerKind
{
    /// <summary>
    /// Signature help was invoked manually by the user or by a command.
    /// </summary>
    Invoked = 1,

    /// <summary>
    /// Signature help was triggered by a trigger character.
    /// </summary>
    TriggerCharacter = 2,

    /// <summary>
    /// Signature help was triggered by the cursor moving or by the document content changing.
    /// </summary>
    ContentChange = 3,
}

/// <summary>
/// The reason why code actions were requested.
/// 
/// @since 3.17.0
/// </summary>
public enum CodeActionTriggerKind
{
    /// <summary>
    /// Code actions were explicitly requested by the user or by an extension.
    /// </summary>
    Invoked = 1,

    /// <summary>
    /// Code actions were requested automatically.
    /// 
    /// This typically happens when current selection in a file changes, but can
    /// also be triggered when file content changes.
    /// </summary>
    Automatic = 2,
}

/// <summary>
/// A pattern kind describing if a glob pattern matches a file a folder or
/// both.
/// 
/// @since 3.16.0
/// </summary>
public static class FileOperationPatternKind
{
    /// <summary>
    /// The pattern matches a file only.
    /// </summary>
    public static readonly string File = "file";
    /// <summary>
    /// The pattern matches a folder only.
    /// </summary>
    public static readonly string Folder = "folder";
}

/// <summary>
/// A notebook cell kind.
/// 
/// @since 3.17.0
/// </summary>
public enum NotebookCellKind
{
    /// <summary>
    /// A markup-cell is formatted source that is used for display.
    /// </summary>
    Markup = 1,

    /// <summary>
    /// A code-cell is source code.
    /// </summary>
    Code = 2,
}

public static class ResourceOperationKind
{
    /// <summary>
    /// Supports creating new files and folders.
    /// </summary>
    public static readonly string Create = "create";
    /// <summary>
    /// Supports renaming existing files and folders.
    /// </summary>
    public static readonly string Rename = "rename";
    /// <summary>
    /// Supports deleting existing files and folders.
    /// </summary>
    public static readonly string Delete = "delete";
}

public static class FailureHandlingKind
{
    /// <summary>
    /// Applying the workspace change is simply aborted if one of the changes provided
    /// fails. All operations executed before the failing operation stay executed.
    /// </summary>
    public static readonly string Abort = "abort";
    /// <summary>
    /// All operations are executed transactional. That means they either all
    /// succeed or no changes at all are applied to the workspace.
    /// </summary>
    public static readonly string Transactional = "transactional";
    /// <summary>
    /// If the workspace edit contains only textual file changes they are executed transactional.
    /// If resource changes (create, rename or delete file) are part of the change the failure
    /// handling strategy is abort.
    /// </summary>
    public static readonly string TextOnlyTransactional = "textOnlyTransactional";
    /// <summary>
    /// The client tries to undo the operations already executed. But there is no
    /// guarantee that this is succeeding.
    /// </summary>
    public static readonly string Undo = "undo";
}

public enum PrepareSupportDefaultBehavior
{
    /// <summary>
    /// The client's default behavior is to select the identifier
    /// according the to language's syntax rule.
    /// </summary>
    Identifier = 1,
}

public static class TokenFormat
{
    public static readonly string Relative = "relative";
}

