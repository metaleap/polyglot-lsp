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
    /// <summary>
    /// The value is always "namespace". 
    /// </summary>
    public static readonly string Namespace = "namespace";
    /// <summary>
    /// Represents a generic type. Acts as a fallback for types which can't be mapped to
    /// a specific type like class or enum.
    /// The value is always "type". 
    /// </summary>
    public static readonly string Type = "type";
    /// <summary>
    /// The value is always "class". 
    /// </summary>
    public static readonly string Class = "class";
    /// <summary>
    /// The value is always "enum". 
    /// </summary>
    public static readonly string Enum = "enum";
    /// <summary>
    /// The value is always "interface". 
    /// </summary>
    public static readonly string Interface = "interface";
    /// <summary>
    /// The value is always "struct". 
    /// </summary>
    public static readonly string Struct = "struct";
    /// <summary>
    /// The value is always "typeParameter". 
    /// </summary>
    public static readonly string TypeParameter = "typeParameter";
    /// <summary>
    /// The value is always "parameter". 
    /// </summary>
    public static readonly string Parameter = "parameter";
    /// <summary>
    /// The value is always "variable". 
    /// </summary>
    public static readonly string Variable = "variable";
    /// <summary>
    /// The value is always "property". 
    /// </summary>
    public static readonly string Property = "property";
    /// <summary>
    /// The value is always "enumMember". 
    /// </summary>
    public static readonly string EnumMember = "enumMember";
    /// <summary>
    /// The value is always "event". 
    /// </summary>
    public static readonly string Event = "event";
    /// <summary>
    /// The value is always "function". 
    /// </summary>
    public static readonly string Function = "function";
    /// <summary>
    /// The value is always "method". 
    /// </summary>
    public static readonly string Method = "method";
    /// <summary>
    /// The value is always "macro". 
    /// </summary>
    public static readonly string Macro = "macro";
    /// <summary>
    /// The value is always "keyword". 
    /// </summary>
    public static readonly string Keyword = "keyword";
    /// <summary>
    /// The value is always "modifier". 
    /// </summary>
    public static readonly string Modifier = "modifier";
    /// <summary>
    /// The value is always "comment". 
    /// </summary>
    public static readonly string Comment = "comment";
    /// <summary>
    /// The value is always "string". 
    /// </summary>
    public static readonly string String = "string";
    /// <summary>
    /// The value is always "number". 
    /// </summary>
    public static readonly string Number = "number";
    /// <summary>
    /// The value is always "regexp". 
    /// </summary>
    public static readonly string Regexp = "regexp";
    /// <summary>
    /// The value is always "operator". 
    /// </summary>
    public static readonly string Operator = "operator";
    /// <summary>
    /// @since 3.17.0
    /// The value is always "decorator". 
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
    /// <summary>
    /// The value is always "declaration". 
    /// </summary>
    public static readonly string Declaration = "declaration";
    /// <summary>
    /// The value is always "definition". 
    /// </summary>
    public static readonly string Definition = "definition";
    /// <summary>
    /// The value is always "readonly". 
    /// </summary>
    public static readonly string Readonly = "readonly";
    /// <summary>
    /// The value is always "static". 
    /// </summary>
    public static readonly string Static = "static";
    /// <summary>
    /// The value is always "deprecated". 
    /// </summary>
    public static readonly string Deprecated = "deprecated";
    /// <summary>
    /// The value is always "abstract". 
    /// </summary>
    public static readonly string Abstract = "abstract";
    /// <summary>
    /// The value is always "async". 
    /// </summary>
    public static readonly string Async = "async";
    /// <summary>
    /// The value is always "modification". 
    /// </summary>
    public static readonly string Modification = "modification";
    /// <summary>
    /// The value is always "documentation". 
    /// </summary>
    public static readonly string Documentation = "documentation";
    /// <summary>
    /// The value is always "defaultLibrary". 
    /// </summary>
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
    /// The value is always "full". 
    /// </summary>
    public static readonly string Full = "full";
    /// <summary>
    /// A report indicating that the last
    /// returned report is still accurate.
    /// The value is always "unchanged". 
    /// </summary>
    public static readonly string Unchanged = "unchanged";
}

/// <summary>
/// Predefined error codes.
/// </summary>
public enum ErrorCodes
{
    /// <summary>
    /// The value is always -32700. 
    /// </summary>
    ParseError = -32700,
    /// <summary>
    /// The value is always -32600. 
    /// </summary>
    InvalidRequest = -32600,
    /// <summary>
    /// The value is always -32601. 
    /// </summary>
    MethodNotFound = -32601,
    /// <summary>
    /// The value is always -32602. 
    /// </summary>
    InvalidParams = -32602,
    /// <summary>
    /// The value is always -32603. 
    /// </summary>
    InternalError = -32603,
    /// <summary>
    /// Error code indicating that a server received a notification or
    /// request before the server has received the `initialize` request.
    /// The value is always -32002. 
    /// </summary>
    ServerNotInitialized = -32002,
    /// <summary>
    /// The value is always -32001. 
    /// </summary>
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
    /// The value is always -32803. 
    /// </summary>
    RequestFailed = -32803,
    /// <summary>
    /// The server cancelled the request. This error code should
    /// only be used for requests that explicitly support being
    /// server cancellable.
    /// 
    /// @since 3.17.0
    /// The value is always -32802. 
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
    /// The value is always -32801. 
    /// </summary>
    ContentModified = -32801,
    /// <summary>
    /// The client has canceled a request and a server as detected
    /// the cancel.
    /// The value is always -32800. 
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
    /// The value is always "comment". 
    /// </summary>
    public static readonly string Comment = "comment";
    /// <summary>
    /// Folding range for an import or include
    /// The value is always "imports". 
    /// </summary>
    public static readonly string Imports = "imports";
    /// <summary>
    /// Folding range for a region (e.g. `#region`)
    /// The value is always "region". 
    /// </summary>
    public static readonly string Region = "region";
}

/// <summary>
/// A symbol kind.
/// </summary>
public enum SymbolKind
{
    /// <summary>
    /// The value is always 1. 
    /// </summary>
    File = 1,
    /// <summary>
    /// The value is always 2. 
    /// </summary>
    Module = 2,
    /// <summary>
    /// The value is always 3. 
    /// </summary>
    Namespace = 3,
    /// <summary>
    /// The value is always 4. 
    /// </summary>
    Package = 4,
    /// <summary>
    /// The value is always 5. 
    /// </summary>
    Class = 5,
    /// <summary>
    /// The value is always 6. 
    /// </summary>
    Method = 6,
    /// <summary>
    /// The value is always 7. 
    /// </summary>
    Property = 7,
    /// <summary>
    /// The value is always 8. 
    /// </summary>
    Field = 8,
    /// <summary>
    /// The value is always 9. 
    /// </summary>
    Constructor = 9,
    /// <summary>
    /// The value is always 10. 
    /// </summary>
    Enum = 10,
    /// <summary>
    /// The value is always 11. 
    /// </summary>
    Interface = 11,
    /// <summary>
    /// The value is always 12. 
    /// </summary>
    Function = 12,
    /// <summary>
    /// The value is always 13. 
    /// </summary>
    Variable = 13,
    /// <summary>
    /// The value is always 14. 
    /// </summary>
    Constant = 14,
    /// <summary>
    /// The value is always 15. 
    /// </summary>
    String = 15,
    /// <summary>
    /// The value is always 16. 
    /// </summary>
    Number = 16,
    /// <summary>
    /// The value is always 17. 
    /// </summary>
    Boolean = 17,
    /// <summary>
    /// The value is always 18. 
    /// </summary>
    Array = 18,
    /// <summary>
    /// The value is always 19. 
    /// </summary>
    Object = 19,
    /// <summary>
    /// The value is always 20. 
    /// </summary>
    Key = 20,
    /// <summary>
    /// The value is always 21. 
    /// </summary>
    Null = 21,
    /// <summary>
    /// The value is always 22. 
    /// </summary>
    EnumMember = 22,
    /// <summary>
    /// The value is always 23. 
    /// </summary>
    Struct = 23,
    /// <summary>
    /// The value is always 24. 
    /// </summary>
    Event = 24,
    /// <summary>
    /// The value is always 25. 
    /// </summary>
    Operator = 25,
    /// <summary>
    /// The value is always 26. 
    /// </summary>
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
    /// The value is always 1. 
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
    /// The value is always "document". 
    /// </summary>
    public static readonly string Document = "document";
    /// <summary>
    /// The moniker is unique inside a project for which a dump got created
    /// The value is always "project". 
    /// </summary>
    public static readonly string Project = "project";
    /// <summary>
    /// The moniker is unique inside the group to which a project belongs
    /// The value is always "group". 
    /// </summary>
    public static readonly string Group = "group";
    /// <summary>
    /// The moniker is unique inside the moniker scheme.
    /// The value is always "scheme". 
    /// </summary>
    public static readonly string Scheme = "scheme";
    /// <summary>
    /// The moniker is globally unique
    /// The value is always "global". 
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
    /// The value is always "import". 
    /// </summary>
    public static readonly string Import = "import";
    /// <summary>
    /// The moniker represents a symbol that is exported from a project
    /// The value is always "export". 
    /// </summary>
    public static readonly string Export = "export";
    /// <summary>
    /// The moniker represents a symbol that is local to a project (e.g. a local
    /// variable of a function, a class not visible outside the project, ...)
    /// The value is always "local". 
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
    /// The value is always 1. 
    /// </summary>
    Type = 1,
    /// <summary>
    /// An inlay hint that is for a parameter.
    /// The value is always 2. 
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
    /// The value is always 1. 
    /// </summary>
    Error = 1,
    /// <summary>
    /// A warning message.
    /// The value is always 2. 
    /// </summary>
    Warning = 2,
    /// <summary>
    /// An information message.
    /// The value is always 3. 
    /// </summary>
    Info = 3,
    /// <summary>
    /// A log message.
    /// The value is always 4. 
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
    /// The value is always 0. 
    /// </summary>
    None = 0,
    /// <summary>
    /// Documents are synced by always sending the full content
    /// of the document.
    /// The value is always 1. 
    /// </summary>
    Full = 1,
    /// <summary>
    /// Documents are synced by sending the full content on open.
    /// After that only incremental updates to the document are
    /// send.
    /// The value is always 2. 
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
    /// The value is always 1. 
    /// </summary>
    Manual = 1,
    /// <summary>
    /// Automatic after a delay.
    /// The value is always 2. 
    /// </summary>
    AfterDelay = 2,
    /// <summary>
    /// When the editor lost focus.
    /// The value is always 3. 
    /// </summary>
    FocusOut = 3,
}

/// <summary>
/// The kind of a completion entry.
/// </summary>
public enum CompletionItemKind
{
    /// <summary>
    /// The value is always 1. 
    /// </summary>
    Text = 1,
    /// <summary>
    /// The value is always 2. 
    /// </summary>
    Method = 2,
    /// <summary>
    /// The value is always 3. 
    /// </summary>
    Function = 3,
    /// <summary>
    /// The value is always 4. 
    /// </summary>
    Constructor = 4,
    /// <summary>
    /// The value is always 5. 
    /// </summary>
    Field = 5,
    /// <summary>
    /// The value is always 6. 
    /// </summary>
    Variable = 6,
    /// <summary>
    /// The value is always 7. 
    /// </summary>
    Class = 7,
    /// <summary>
    /// The value is always 8. 
    /// </summary>
    Interface = 8,
    /// <summary>
    /// The value is always 9. 
    /// </summary>
    Module = 9,
    /// <summary>
    /// The value is always 10. 
    /// </summary>
    Property = 10,
    /// <summary>
    /// The value is always 11. 
    /// </summary>
    Unit = 11,
    /// <summary>
    /// The value is always 12. 
    /// </summary>
    Value = 12,
    /// <summary>
    /// The value is always 13. 
    /// </summary>
    Enum = 13,
    /// <summary>
    /// The value is always 14. 
    /// </summary>
    Keyword = 14,
    /// <summary>
    /// The value is always 15. 
    /// </summary>
    Snippet = 15,
    /// <summary>
    /// The value is always 16. 
    /// </summary>
    Color = 16,
    /// <summary>
    /// The value is always 17. 
    /// </summary>
    File = 17,
    /// <summary>
    /// The value is always 18. 
    /// </summary>
    Reference = 18,
    /// <summary>
    /// The value is always 19. 
    /// </summary>
    Folder = 19,
    /// <summary>
    /// The value is always 20. 
    /// </summary>
    EnumMember = 20,
    /// <summary>
    /// The value is always 21. 
    /// </summary>
    Constant = 21,
    /// <summary>
    /// The value is always 22. 
    /// </summary>
    Struct = 22,
    /// <summary>
    /// The value is always 23. 
    /// </summary>
    Event = 23,
    /// <summary>
    /// The value is always 24. 
    /// </summary>
    Operator = 24,
    /// <summary>
    /// The value is always 25. 
    /// </summary>
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
    /// The value is always 1. 
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
    /// The value is always 1. 
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
    /// The value is always 2. 
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
    /// The value is always 1. 
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
    /// The value is always 2. 
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
    /// The value is always 1. 
    /// </summary>
    Text = 1,
    /// <summary>
    /// Read-access of a symbol, like reading a variable.
    /// The value is always 2. 
    /// </summary>
    Read = 2,
    /// <summary>
    /// Write-access of a symbol, like writing to a variable.
    /// The value is always 3. 
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
    /// The value is always "". 
    /// </summary>
    public static readonly string Empty = "";
    /// <summary>
    /// Base kind for quickfix actions: 'quickfix'
    /// The value is always "quickfix". 
    /// </summary>
    public static readonly string QuickFix = "quickfix";
    /// <summary>
    /// Base kind for refactoring actions: 'refactor'
    /// The value is always "refactor". 
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
    /// The value is always "refactor.extract". 
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
    /// The value is always "refactor.inline". 
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
    /// The value is always "refactor.rewrite". 
    /// </summary>
    public static readonly string RefactorRewrite = "refactor.rewrite";
    /// <summary>
    /// Base kind for source actions: `source`
    /// 
    /// Source code actions apply to the entire file.
    /// The value is always "source". 
    /// </summary>
    public static readonly string Source = "source";
    /// <summary>
    /// Base kind for an organize imports source action: `source.organizeImports`
    /// The value is always "source.organizeImports". 
    /// </summary>
    public static readonly string SourceOrganizeImports = "source.organizeImports";
    /// <summary>
    /// Base kind for auto-fix source actions: `source.fixAll`.
    /// 
    /// Fix all actions automatically fix errors that have a clear fix that do not require user input.
    /// They should not suppress errors or perform unsafe fixes such as generating new types or classes.
    /// 
    /// @since 3.15.0
    /// The value is always "source.fixAll". 
    /// </summary>
    public static readonly string SourceFixAll = "source.fixAll";
}

public static class TraceValues
{
    /// <summary>
    /// Turn tracing off.
    /// The value is always "off". 
    /// </summary>
    public static readonly string Off = "off";
    /// <summary>
    /// Trace messages only.
    /// The value is always "messages". 
    /// </summary>
    public static readonly string Messages = "messages";
    /// <summary>
    /// Verbose message tracing.
    /// The value is always "verbose". 
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
    /// The value is always "plaintext". 
    /// </summary>
    public static readonly string PlainText = "plaintext";
    /// <summary>
    /// Markdown is supported as a content format
    /// The value is always "markdown". 
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
    /// The value is always "utf-8". 
    /// </summary>
    public static readonly string UTF8 = "utf-8";
    /// <summary>
    /// Character offsets count UTF-16 code units.
    /// 
    /// This is the default and must always be supported
    /// by servers
    /// The value is always "utf-16". 
    /// </summary>
    public static readonly string UTF16 = "utf-16";
    /// <summary>
    /// Character offsets count UTF-32 code units.
    /// 
    /// Implementation note: these are the same as Unicode code points,
    /// so this `PositionEncodingKind` may also be used for an
    /// encoding-agnostic representation of character offsets.
    /// The value is always "utf-32". 
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
    /// The value is always 1. 
    /// </summary>
    Created = 1,
    /// <summary>
    /// The file got changed.
    /// The value is always 2. 
    /// </summary>
    Changed = 2,
    /// <summary>
    /// The file got deleted.
    /// The value is always 3. 
    /// </summary>
    Deleted = 3,
}

public enum WatchKind
{
    /// <summary>
    /// Interested in create events.
    /// The value is always 1. 
    /// </summary>
    Create = 1,
    /// <summary>
    /// Interested in change events
    /// The value is always 2. 
    /// </summary>
    Change = 2,
    /// <summary>
    /// Interested in delete events
    /// The value is always 4. 
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
    /// The value is always 1. 
    /// </summary>
    Error = 1,
    /// <summary>
    /// Reports a warning.
    /// The value is always 2. 
    /// </summary>
    Warning = 2,
    /// <summary>
    /// Reports an information.
    /// The value is always 3. 
    /// </summary>
    Information = 3,
    /// <summary>
    /// Reports a hint.
    /// The value is always 4. 
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
    /// The value is always 1. 
    /// </summary>
    Unnecessary = 1,
    /// <summary>
    /// Deprecated or obsolete code.
    /// 
    /// Clients are allowed to rendered diagnostics with this tag strike through.
    /// The value is always 2. 
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
    /// The value is always 1. 
    /// </summary>
    Invoked = 1,
    /// <summary>
    /// Completion was triggered by a trigger character specified by
    /// the `triggerCharacters` properties of the `CompletionRegistrationOptions`.
    /// The value is always 2. 
    /// </summary>
    TriggerCharacter = 2,
    /// <summary>
    /// Completion was re-triggered as current completion list is incomplete
    /// The value is always 3. 
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
    /// The value is always 1. 
    /// </summary>
    Invoked = 1,
    /// <summary>
    /// Signature help was triggered by a trigger character.
    /// The value is always 2. 
    /// </summary>
    TriggerCharacter = 2,
    /// <summary>
    /// Signature help was triggered by the cursor moving or by the document content changing.
    /// The value is always 3. 
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
    /// The value is always 1. 
    /// </summary>
    Invoked = 1,
    /// <summary>
    /// Code actions were requested automatically.
    /// 
    /// This typically happens when current selection in a file changes, but can
    /// also be triggered when file content changes.
    /// The value is always 2. 
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
    /// The value is always "file". 
    /// </summary>
    public static readonly string File = "file";
    /// <summary>
    /// The pattern matches a folder only.
    /// The value is always "folder". 
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
    /// The value is always 1. 
    /// </summary>
    Markup = 1,
    /// <summary>
    /// A code-cell is source code.
    /// The value is always 2. 
    /// </summary>
    Code = 2,
}

public static class ResourceOperationKind
{
    /// <summary>
    /// Supports creating new files and folders.
    /// The value is always "create". 
    /// </summary>
    public static readonly string Create = "create";
    /// <summary>
    /// Supports renaming existing files and folders.
    /// The value is always "rename". 
    /// </summary>
    public static readonly string Rename = "rename";
    /// <summary>
    /// Supports deleting existing files and folders.
    /// The value is always "delete". 
    /// </summary>
    public static readonly string Delete = "delete";
}

public static class FailureHandlingKind
{
    /// <summary>
    /// Applying the workspace change is simply aborted if one of the changes provided
    /// fails. All operations executed before the failing operation stay executed.
    /// The value is always "abort". 
    /// </summary>
    public static readonly string Abort = "abort";
    /// <summary>
    /// All operations are executed transactional. That means they either all
    /// succeed or no changes at all are applied to the workspace.
    /// The value is always "transactional". 
    /// </summary>
    public static readonly string Transactional = "transactional";
    /// <summary>
    /// If the workspace edit contains only textual file changes they are executed transactional.
    /// If resource changes (create, rename or delete file) are part of the change the failure
    /// handling strategy is abort.
    /// The value is always "textOnlyTransactional". 
    /// </summary>
    public static readonly string TextOnlyTransactional = "textOnlyTransactional";
    /// <summary>
    /// The client tries to undo the operations already executed. But there is no
    /// guarantee that this is succeeding.
    /// The value is always "undo". 
    /// </summary>
    public static readonly string Undo = "undo";
}

public enum PrepareSupportDefaultBehavior
{
    /// <summary>
    /// The client's default behavior is to select the identifier
    /// according the to language's syntax rule.
    /// The value is always 1. 
    /// </summary>
    Identifier = 1,
}

public static class TokenFormat
{
    /// <summary>
    /// The value is always "relative". 
    /// </summary>
    public static readonly string Relative = "relative";
}

