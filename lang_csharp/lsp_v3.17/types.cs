/// <summary>Language Server Protocol (LSP) v3.17 SDK for C#: auto-generated via github.com/metaleap/polyglot-vsx-and-lsp/gen/cmd/gen_lsp</summary>
namespace lsp;


/*TOr*/
public abstract record TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile
{
    private TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile() { }
    public sealed record TextDocumentEdit(TextDocumentEdit? TextDocumentEdit) : TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile;
    public sealed record CreateFile(CreateFile? CreateFile) : TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile;
    public sealed record RenameFile(RenameFile? RenameFile) : TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile;
    public sealed record DeleteFile(DeleteFile? DeleteFile) : TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile;
}

/*TOr*/
public abstract record StringOrInlayHintLabelParts
{
    private StringOrInlayHintLabelParts() { }
    public sealed record String(string? String) : StringOrInlayHintLabelParts;
    public sealed record InlayHintLabelParts(/*TOpt*/InlayHintLabelPart[] InlayHintLabelParts) : StringOrInlayHintLabelParts;
}

/*TOr*/
public abstract record StringOrMarkupContent
{
    private StringOrMarkupContent() { }
    public sealed record String(string? String) : StringOrMarkupContent;
    public sealed record MarkupContent(MarkupContent? MarkupContent) : StringOrMarkupContent;
}

/*TOr*/
public abstract record FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport
{
    private FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport() { }
    public sealed record FullDocumentDiagnosticReport(FullDocumentDiagnosticReport? FullDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport;
    public sealed record UnchangedDocumentDiagnosticReport(UnchangedDocumentDiagnosticReport? UnchangedDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport;
}

/*TStruc*/
public class NameVersion
{
    /// <summary>
    /// The name of the server as defined by the server.
    /// </summary>
    public string Name;
    /// <summary>
    /// The server's version as defined by the server.
    /// </summary>
    public string? Version;
}


/*TOr*/
public abstract record StringOrStrings
{
    private StringOrStrings() { }
    public sealed record String(string? String) : StringOrStrings;
    public sealed record Strings(/*TOpt*/string[] Strings) : StringOrStrings;
}

/*TOr*/
public abstract record StringOrMarkupContent
{
    private StringOrMarkupContent() { }
    public sealed record String(string? String) : StringOrMarkupContent;
    public sealed record MarkupContent(MarkupContent? MarkupContent) : StringOrMarkupContent;
}

/*TOr*/
public abstract record TextEditOrInsertReplaceEdit
{
    private TextEditOrInsertReplaceEdit() { }
    public sealed record TextEdit(TextEdit? TextEdit) : TextEditOrInsertReplaceEdit;
    public sealed record InsertReplaceEdit(InsertReplaceEdit? InsertReplaceEdit) : TextEditOrInsertReplaceEdit;
}

/*TStruc*/
public class InsertReplace
{
    public Range Insert;
    public Range Replace;
}


/*TOr*/
public abstract record RangeOrInsertReplace
{
    private RangeOrInsertReplace() { }
    public sealed record Range(Range? Range) : RangeOrInsertReplace;
    public sealed record InsertReplace(/*TOpt*/__TypeIdentGen__cr06pj99kwku__ InsertReplace) : RangeOrInsertReplace;
}

/*TStruc*/
public class CommitCharactersEditRangeInsertTextFormatInsertTextModeData
{
    /// <summary>
    /// A default commit character set.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/string[] CommitCharacters;
    /// <summary>
    /// A default edit range.
    /// 
    /// @since 3.17.0
    /// 
    /// This object has "OneOf" union-type semantics: only at-most one field in it is ever set, all others will be null/undefined/nil/empty/etc.
    /// </summary>
    public /*TOpt*/__TypeIdentGen__cr06pj99jgwa__ EditRange;
    /// <summary>
    /// A default insert text format.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/InsertTextFormat InsertTextFormat;
    /// <summary>
    /// A default insert text mode.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/InsertTextMode InsertTextMode;
    /// <summary>
    /// A default data value.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/LSPAny Data;
}


/*TOr*/
public abstract record MarkupContentOrMarkedStringOrMarkedStrings
{
    private MarkupContentOrMarkedStringOrMarkedStrings() { }
    public sealed record MarkupContent(MarkupContent? MarkupContent) : MarkupContentOrMarkedStringOrMarkedStrings;
    public sealed record MarkedString(/*TOpt*/MarkedString MarkedString) : MarkupContentOrMarkedStringOrMarkedStrings;
    public sealed record MarkedStrings(/*TOpt*/MarkedString[] MarkedStrings) : MarkupContentOrMarkedStringOrMarkedStrings;
}

/*TStruc*/
public class Reason
{
    /// <summary>
    /// Human readable description of why the code action is currently disabled.
    /// 
    /// This is displayed in the code actions UI.
    /// </summary>
    public string Reason;
}


/*TStruc*/
public class Uri
{
    public string Uri;
}


/*TOr*/
public abstract record LocationOrUri
{
    private LocationOrUri() { }
    public sealed record Location(Location? Location) : LocationOrUri;
    public sealed record Uri(/*TOpt*/__TypeIdentGen__cr06pj9ac2y7__ Uri) : LocationOrUri;
}

/*TOr*/
public abstract record IntegerOrString
{
    private IntegerOrString() { }
    public sealed record Integer(int? Integer) : IntegerOrString;
    public sealed record String(string? String) : IntegerOrString;
}

/*TOr*/
public abstract record BooleanOrAnyByString
{
    private BooleanOrAnyByString() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrAnyByString;
    public sealed record AnyByString(/*TOpt*/System.Collections.Generic.Dictionary<string, object> AnyByString) : BooleanOrAnyByString;
}

/*TStruc*/
public class Delta
{
    /// <summary>
    /// The server supports deltas for full documents.
    /// </summary>
    public bool? Delta;
}


/*TOr*/
public abstract record BooleanOrDelta
{
    private BooleanOrDelta() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrDelta;
    public sealed record Delta(/*TOpt*/__TypeIdentGen__cr06pj9bvire__ Delta) : BooleanOrDelta;
}

/*TOr*/
public abstract record TextEditOrAnnotatedTextEdit
{
    private TextEditOrAnnotatedTextEdit() { }
    public sealed record TextEdit(TextEdit? TextEdit) : TextEditOrAnnotatedTextEdit;
    public sealed record AnnotatedTextEdit(AnnotatedTextEdit? AnnotatedTextEdit) : TextEditOrAnnotatedTextEdit;
}

/*TOr*/
public abstract record StringOrMarkupContent
{
    private StringOrMarkupContent() { }
    public sealed record String(string? String) : StringOrMarkupContent;
    public sealed record MarkupContent(MarkupContent? MarkupContent) : StringOrMarkupContent;
}

/*TOr*/
public abstract record FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport
{
    private FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport() { }
    public sealed record FullDocumentDiagnosticReport(FullDocumentDiagnosticReport? FullDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport;
    public sealed record UnchangedDocumentDiagnosticReport(UnchangedDocumentDiagnosticReport? UnchangedDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport;
}

/*TOr*/
public abstract record FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport
{
    private FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport() { }
    public sealed record FullDocumentDiagnosticReport(FullDocumentDiagnosticReport? FullDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport;
    public sealed record UnchangedDocumentDiagnosticReport(UnchangedDocumentDiagnosticReport? UnchangedDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport;
}

/*TStruc*/
public class ArrayDidOpenDidClose
{
    /// <summary>
    /// The change to the cell array.
    /// </summary>
    public NotebookCellArrayChange Array;
    /// <summary>
    /// Additional opened cell text documents.
    /// </summary>
    public /*TOpt*/TextDocumentItem[] DidOpen;
    /// <summary>
    /// Additional closed cell text documents.
    /// </summary>
    public /*TOpt*/TextDocumentIdentifier[] DidClose;
}


/*TStruc*/
public class DocumentChanges
{
    public VersionedTextDocumentIdentifier Document;
    public TextDocumentContentChangeEvent[] Changes;
}


/*TStruc*/
public class StructureDataTextContent
{
    /// <summary>
    /// Changes to the cell structure to add or
    /// remove cells.
    /// </summary>
    public /*TOpt*/__TypeIdentGen__cr06pj9cve5v__ Structure;
    /// <summary>
    /// Changes to notebook cells properties like its
    /// kind, execution summary or metadata.
    /// </summary>
    public /*TOpt*/NotebookCell[] Data;
    /// <summary>
    /// Changes to the text content of notebook cells.
    /// </summary>
    public /*TOpt*/__TypeIdentGen__cr06pj9cxq98__[] TextContent;
}


/*TStruc*/
public class NameVersion
{
    /// <summary>
    /// The name of the client as defined by the client.
    /// </summary>
    public string Name;
    /// <summary>
    /// The client's version as defined by the client.
    /// </summary>
    public string? Version;
}


/*TOr*/
public abstract record TextDocumentSyncOptionsOrTextDocumentSyncKind
{
    private TextDocumentSyncOptionsOrTextDocumentSyncKind() { }
    public sealed record TextDocumentSyncOptions(TextDocumentSyncOptions? TextDocumentSyncOptions) : TextDocumentSyncOptionsOrTextDocumentSyncKind;
    public sealed record TextDocumentSyncKind(/*TOpt*/TextDocumentSyncKind TextDocumentSyncKind) : TextDocumentSyncOptionsOrTextDocumentSyncKind;
}

/*TOr*/
public abstract record NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions
{
    private NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions() { }
    public sealed record NotebookDocumentSyncOptions(NotebookDocumentSyncOptions? NotebookDocumentSyncOptions) : NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions;
    public sealed record NotebookDocumentSyncRegistrationOptions(NotebookDocumentSyncRegistrationOptions? NotebookDocumentSyncRegistrationOptions) : NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrHoverOptions
{
    private BooleanOrHoverOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrHoverOptions;
    public sealed record HoverOptions(HoverOptions? HoverOptions) : BooleanOrHoverOptions;
}

/*TOr*/
public abstract record BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions
{
    private BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions;
    public sealed record DeclarationOptions(DeclarationOptions? DeclarationOptions) : BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions;
    public sealed record DeclarationRegistrationOptions(DeclarationRegistrationOptions? DeclarationRegistrationOptions) : BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrDefinitionOptions
{
    private BooleanOrDefinitionOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrDefinitionOptions;
    public sealed record DefinitionOptions(DefinitionOptions? DefinitionOptions) : BooleanOrDefinitionOptions;
}

/*TOr*/
public abstract record BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions
{
    private BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions;
    public sealed record TypeDefinitionOptions(TypeDefinitionOptions? TypeDefinitionOptions) : BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions;
    public sealed record TypeDefinitionRegistrationOptions(TypeDefinitionRegistrationOptions? TypeDefinitionRegistrationOptions) : BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrImplementationOptionsOrImplementationRegistrationOptions
{
    private BooleanOrImplementationOptionsOrImplementationRegistrationOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrImplementationOptionsOrImplementationRegistrationOptions;
    public sealed record ImplementationOptions(ImplementationOptions? ImplementationOptions) : BooleanOrImplementationOptionsOrImplementationRegistrationOptions;
    public sealed record ImplementationRegistrationOptions(ImplementationRegistrationOptions? ImplementationRegistrationOptions) : BooleanOrImplementationOptionsOrImplementationRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrReferenceOptions
{
    private BooleanOrReferenceOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrReferenceOptions;
    public sealed record ReferenceOptions(ReferenceOptions? ReferenceOptions) : BooleanOrReferenceOptions;
}

/*TOr*/
public abstract record BooleanOrDocumentHighlightOptions
{
    private BooleanOrDocumentHighlightOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrDocumentHighlightOptions;
    public sealed record DocumentHighlightOptions(DocumentHighlightOptions? DocumentHighlightOptions) : BooleanOrDocumentHighlightOptions;
}

/*TOr*/
public abstract record BooleanOrDocumentSymbolOptions
{
    private BooleanOrDocumentSymbolOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrDocumentSymbolOptions;
    public sealed record DocumentSymbolOptions(DocumentSymbolOptions? DocumentSymbolOptions) : BooleanOrDocumentSymbolOptions;
}

/*TOr*/
public abstract record BooleanOrCodeActionOptions
{
    private BooleanOrCodeActionOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrCodeActionOptions;
    public sealed record CodeActionOptions(CodeActionOptions? CodeActionOptions) : BooleanOrCodeActionOptions;
}

/*TOr*/
public abstract record BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions
{
    private BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions;
    public sealed record DocumentColorOptions(DocumentColorOptions? DocumentColorOptions) : BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions;
    public sealed record DocumentColorRegistrationOptions(DocumentColorRegistrationOptions? DocumentColorRegistrationOptions) : BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrWorkspaceSymbolOptions
{
    private BooleanOrWorkspaceSymbolOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrWorkspaceSymbolOptions;
    public sealed record WorkspaceSymbolOptions(WorkspaceSymbolOptions? WorkspaceSymbolOptions) : BooleanOrWorkspaceSymbolOptions;
}

/*TOr*/
public abstract record BooleanOrDocumentFormattingOptions
{
    private BooleanOrDocumentFormattingOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrDocumentFormattingOptions;
    public sealed record DocumentFormattingOptions(DocumentFormattingOptions? DocumentFormattingOptions) : BooleanOrDocumentFormattingOptions;
}

/*TOr*/
public abstract record BooleanOrDocumentRangeFormattingOptions
{
    private BooleanOrDocumentRangeFormattingOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrDocumentRangeFormattingOptions;
    public sealed record DocumentRangeFormattingOptions(DocumentRangeFormattingOptions? DocumentRangeFormattingOptions) : BooleanOrDocumentRangeFormattingOptions;
}

/*TOr*/
public abstract record BooleanOrRenameOptions
{
    private BooleanOrRenameOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrRenameOptions;
    public sealed record RenameOptions(RenameOptions? RenameOptions) : BooleanOrRenameOptions;
}

/*TOr*/
public abstract record BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions
{
    private BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions;
    public sealed record FoldingRangeOptions(FoldingRangeOptions? FoldingRangeOptions) : BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions;
    public sealed record FoldingRangeRegistrationOptions(FoldingRangeRegistrationOptions? FoldingRangeRegistrationOptions) : BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions
{
    private BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions;
    public sealed record SelectionRangeOptions(SelectionRangeOptions? SelectionRangeOptions) : BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions;
    public sealed record SelectionRangeRegistrationOptions(SelectionRangeRegistrationOptions? SelectionRangeRegistrationOptions) : BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions
{
    private BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions;
    public sealed record CallHierarchyOptions(CallHierarchyOptions? CallHierarchyOptions) : BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions;
    public sealed record CallHierarchyRegistrationOptions(CallHierarchyRegistrationOptions? CallHierarchyRegistrationOptions) : BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions
{
    private BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions;
    public sealed record LinkedEditingRangeOptions(LinkedEditingRangeOptions? LinkedEditingRangeOptions) : BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions;
    public sealed record LinkedEditingRangeRegistrationOptions(LinkedEditingRangeRegistrationOptions? LinkedEditingRangeRegistrationOptions) : BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions;
}

/*TOr*/
public abstract record SemanticTokensOptionsOrSemanticTokensRegistrationOptions
{
    private SemanticTokensOptionsOrSemanticTokensRegistrationOptions() { }
    public sealed record SemanticTokensOptions(SemanticTokensOptions? SemanticTokensOptions) : SemanticTokensOptionsOrSemanticTokensRegistrationOptions;
    public sealed record SemanticTokensRegistrationOptions(SemanticTokensRegistrationOptions? SemanticTokensRegistrationOptions) : SemanticTokensOptionsOrSemanticTokensRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrMonikerOptionsOrMonikerRegistrationOptions
{
    private BooleanOrMonikerOptionsOrMonikerRegistrationOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrMonikerOptionsOrMonikerRegistrationOptions;
    public sealed record MonikerOptions(MonikerOptions? MonikerOptions) : BooleanOrMonikerOptionsOrMonikerRegistrationOptions;
    public sealed record MonikerRegistrationOptions(MonikerRegistrationOptions? MonikerRegistrationOptions) : BooleanOrMonikerOptionsOrMonikerRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions
{
    private BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions;
    public sealed record TypeHierarchyOptions(TypeHierarchyOptions? TypeHierarchyOptions) : BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions;
    public sealed record TypeHierarchyRegistrationOptions(TypeHierarchyRegistrationOptions? TypeHierarchyRegistrationOptions) : BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions
{
    private BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions;
    public sealed record InlineValueOptions(InlineValueOptions? InlineValueOptions) : BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions;
    public sealed record InlineValueRegistrationOptions(InlineValueRegistrationOptions? InlineValueRegistrationOptions) : BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions
{
    private BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions;
    public sealed record InlayHintOptions(InlayHintOptions? InlayHintOptions) : BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions;
    public sealed record InlayHintRegistrationOptions(InlayHintRegistrationOptions? InlayHintRegistrationOptions) : BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions;
}

/*TOr*/
public abstract record DiagnosticOptionsOrDiagnosticRegistrationOptions
{
    private DiagnosticOptionsOrDiagnosticRegistrationOptions() { }
    public sealed record DiagnosticOptions(DiagnosticOptions? DiagnosticOptions) : DiagnosticOptionsOrDiagnosticRegistrationOptions;
    public sealed record DiagnosticRegistrationOptions(DiagnosticRegistrationOptions? DiagnosticRegistrationOptions) : DiagnosticOptionsOrDiagnosticRegistrationOptions;
}

/*TStruc*/
public class WorkspaceFoldersFileOperations
{
    /// <summary>
    /// The server supports workspace folder.
    /// 
    /// @since 3.6.0
    /// </summary>
    public WorkspaceFoldersServerCapabilities? WorkspaceFolders;
    /// <summary>
    /// The server is interested in notifications/requests for operations on files.
    /// 
    /// @since 3.16.0
    /// </summary>
    public FileOperationOptions? FileOperations;
}


/*TOr*/
public abstract record IntegerOrString
{
    private IntegerOrString() { }
    public sealed record Integer(int? Integer) : IntegerOrString;
    public sealed record String(string? String) : IntegerOrString;
}

/*TStruc*/
public class LabelDetailsSupport
{
    /// <summary>
    /// The server has support for completion item label
    /// details (see also `CompletionItemLabelDetails`) when
    /// receiving a completion item in a resolve call.
    /// 
    /// @since 3.17.0
    /// </summary>
    public bool? LabelDetailsSupport;
}


/*TOr*/
public abstract record StringOrMarkupContent
{
    private StringOrMarkupContent() { }
    public sealed record String(string? String) : StringOrMarkupContent;
    public sealed record MarkupContent(MarkupContent? MarkupContent) : StringOrMarkupContent;
}

/*TOr*/
public abstract record BooleanOrSaveOptions
{
    private BooleanOrSaveOptions() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrSaveOptions;
    public sealed record SaveOptions(SaveOptions? SaveOptions) : BooleanOrSaveOptions;
}

/*TOr*/
public abstract record StringOrNotebookDocumentFilter
{
    private StringOrNotebookDocumentFilter() { }
    public sealed record String(string? String) : StringOrNotebookDocumentFilter;
    public sealed record NotebookDocumentFilter(/*TOpt*/NotebookDocumentFilter NotebookDocumentFilter) : StringOrNotebookDocumentFilter;
}

/*TStruc*/
public class Language
{
    public string Language;
}


/*TStruc*/
public class NotebookCells
{
    /// <summary>
    /// The notebook to be synced If a string
    /// value is provided it matches against the
    /// notebook type. '*' matches every notebook.
    /// 
    /// This object has "OneOf" union-type semantics: only at-most one field in it is ever set, all others will be null/undefined/nil/empty/etc.
    /// </summary>
    public /*TOpt*/__TypeIdentGen__cr06pj9hd3hs__ Notebook;
    /// <summary>
    /// The cells of the matching notebook to be synced.
    /// </summary>
    public /*TOpt*/__TypeIdentGen__cr06pj9hf18j__[] Cells;
}


/*TOr*/
public abstract record StringOrBoolean
{
    private StringOrBoolean() { }
    public sealed record String(string? String) : StringOrBoolean;
    public sealed record Boolean(bool? Boolean) : StringOrBoolean;
}

/*TTup*/
public class UintegerWithUinteger
{
    public uint Uinteger0;
    public uint Uinteger1;
}

/*TOr*/
public abstract record StringOrUintegerWithUinteger
{
    private StringOrUintegerWithUinteger() { }
    public sealed record String(string? String) : StringOrUintegerWithUinteger;
    public sealed record UintegerWithUinteger(__TypeIdentGen__cr06pj9hq21m__? UintegerWithUinteger) : StringOrUintegerWithUinteger;
}

/*TOr*/
public abstract record StringOrMarkupContent
{
    private StringOrMarkupContent() { }
    public sealed record String(string? String) : StringOrMarkupContent;
    public sealed record MarkupContent(MarkupContent? MarkupContent) : StringOrMarkupContent;
}

/*TOr*/
public abstract record StringOrNotebookDocumentFilter
{
    private StringOrNotebookDocumentFilter() { }
    public sealed record String(string? String) : StringOrNotebookDocumentFilter;
    public sealed record NotebookDocumentFilter(/*TOpt*/NotebookDocumentFilter NotebookDocumentFilter) : StringOrNotebookDocumentFilter;
}

/*TStruc*/
public class CancelRetryOnContentModified
{
    /// <summary>
    /// The client will actively cancel the request.
    /// </summary>
    public bool Cancel;
    /// <summary>
    /// The list of requests for which the client
    /// will retry the request if it receives a
    /// response with error code `ContentModified`
    /// </summary>
    public string[] RetryOnContentModified;
}


/*TOr*/
public abstract record WorkspaceFolderOrURI
{
    private WorkspaceFolderOrURI() { }
    public sealed record WorkspaceFolder(WorkspaceFolder? WorkspaceFolder) : WorkspaceFolderOrURI;
    public sealed record URI(/*TOpt*/string URI) : WorkspaceFolderOrURI;
}

/*TStruc*/
public class GroupsOnLabel
{
    /// <summary>
    /// Whether the client groups edits with equal labels into tree nodes,
    /// for instance all edits labelled with "Changes in Strings" would
    /// be a tree node.
    /// </summary>
    public bool? GroupsOnLabel;
}


/*TStruc*/
public class ValueSet
{
    /// <summary>
    /// The symbol kind values the client supports. When this
    /// property exists the client also guarantees that it will
    /// handle values outside its set gracefully and falls back
    /// to a default value when unknown.
    /// 
    /// If this property is not present the client only supports
    /// the symbol kinds from `File` to `Array` as defined in
    /// the initial version of the protocol.
    /// </summary>
    public /*TOpt*/SymbolKind[] ValueSet;
}


/*TStruc*/
public class ValueSet
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public SymbolTag[] ValueSet;
}


/*TStruc*/
public class Properties
{
    /// <summary>
    /// The properties that a client can resolve lazily. Usually
    /// `location.range`
    /// </summary>
    public string[] Properties;
}


/*TStruc*/
public class ValueSet
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public CompletionItemTag[] ValueSet;
}


/*TStruc*/
public class Properties
{
    /// <summary>
    /// The properties that a client can resolve lazily.
    /// </summary>
    public string[] Properties;
}


/*TStruc*/
public class ValueSet
{
    public InsertTextMode[] ValueSet;
}


/*TStruc*/
public class SnippetSupportCommitCharactersSupportDocumentationFormatDeprecatedSupportPreselectSupportTagSupportInsertReplaceSupportResolveSupportInsertTextModeSupportLabelDetailsSupport
{
    /// <summary>
    /// Client supports snippets as insert text.
    /// 
    /// A snippet can define tab stops and placeholders with `$1`, `$2`
    /// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
    /// the end of the snippet. Placeholders with equal identifiers are linked,
    /// that is typing in one will update others too.
    /// </summary>
    public bool? SnippetSupport;
    /// <summary>
    /// Client supports commit characters on a completion item.
    /// </summary>
    public bool? CommitCharactersSupport;
    /// <summary>
    /// Client supports the following content formats for the documentation
    /// property. The order describes the preferred format of the client.
    /// </summary>
    public /*TOpt*/MarkupKind[] DocumentationFormat;
    /// <summary>
    /// Client supports the deprecated property on a completion item.
    /// </summary>
    public bool? DeprecatedSupport;
    /// <summary>
    /// Client supports the preselect property on a completion item.
    /// </summary>
    public bool? PreselectSupport;
    /// <summary>
    /// Client supports the tag property on a completion item. Clients supporting
    /// tags have to handle unknown tags gracefully. Clients especially need to
    /// preserve unknown tags when sending a completion item back to the server in
    /// a resolve call.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/__TypeIdentGen__cr06pj9ju0cn__ TagSupport;
    /// <summary>
    /// Client support insert replace edit to control different behavior if a
    /// completion item is inserted in the text or should replace text.
    /// 
    /// @since 3.16.0
    /// </summary>
    public bool? InsertReplaceSupport;
    /// <summary>
    /// Indicates which properties a client can resolve lazily on a completion
    /// item. Before version 3.16.0 only the predefined properties `documentation`
    /// and `details` could be resolved lazily.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/__TypeIdentGen__cr06pj9jvqzs__ ResolveSupport;
    /// <summary>
    /// The client supports the `insertTextMode` property on
    /// a completion item to override the whitespace handling mode
    /// as defined by the client (see `insertTextMode`).
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/__TypeIdentGen__cr06pj9jwpeu__ InsertTextModeSupport;
    /// <summary>
    /// The client has support for completion item label
    /// details (see also `CompletionItemLabelDetails`).
    /// 
    /// @since 3.17.0
    /// </summary>
    public bool? LabelDetailsSupport;
}


/*TStruc*/
public class ValueSet
{
    /// <summary>
    /// The completion item kind values the client supports. When this
    /// property exists the client also guarantees that it will
    /// handle values outside its set gracefully and falls back
    /// to a default value when unknown.
    /// 
    /// If this property is not present the client only supports
    /// the completion items kinds from `Text` to `Reference` as defined in
    /// the initial version of the protocol.
    /// </summary>
    public /*TOpt*/CompletionItemKind[] ValueSet;
}


/*TStruc*/
public class ItemDefaults
{
    /// <summary>
    /// The client supports the following itemDefaults on
    /// a completion list.
    /// 
    /// The value lists the supported property names of the
    /// `CompletionList.itemDefaults` object. If omitted
    /// no properties are supported.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/string[] ItemDefaults;
}


/*TStruc*/
public class LabelOffsetSupport
{
    /// <summary>
    /// The client supports processing label offsets instead of a
    /// simple label string.
    /// 
    /// @since 3.14.0
    /// </summary>
    public bool? LabelOffsetSupport;
}


/*TStruc*/
public class DocumentationFormatParameterInformationActiveParameterSupport
{
    /// <summary>
    /// Client supports the following content formats for the documentation
    /// property. The order describes the preferred format of the client.
    /// </summary>
    public /*TOpt*/MarkupKind[] DocumentationFormat;
    /// <summary>
    /// Client capabilities specific to parameter information.
    /// </summary>
    public /*TOpt*/__TypeIdentGen__cr06pj9k5v2h__ ParameterInformation;
    /// <summary>
    /// The client supports the `activeParameter` property on `SignatureInformation`
    /// literal.
    /// 
    /// @since 3.16.0
    /// </summary>
    public bool? ActiveParameterSupport;
}


/*TStruc*/
public class ValueSet
{
    /// <summary>
    /// The symbol kind values the client supports. When this
    /// property exists the client also guarantees that it will
    /// handle values outside its set gracefully and falls back
    /// to a default value when unknown.
    /// 
    /// If this property is not present the client only supports
    /// the symbol kinds from `File` to `Array` as defined in
    /// the initial version of the protocol.
    /// </summary>
    public /*TOpt*/SymbolKind[] ValueSet;
}


/*TStruc*/
public class ValueSet
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public SymbolTag[] ValueSet;
}


/*TStruc*/
public class ValueSet
{
    /// <summary>
    /// The code action kind values the client supports. When this
    /// property exists the client also guarantees that it will
    /// handle values outside its set gracefully and falls back
    /// to a default value when unknown.
    /// </summary>
    public CodeActionKind[] ValueSet;
}


/*TStruc*/
public class CodeActionKind
{
    /// <summary>
    /// The code action kind is support with the following value
    /// set.
    /// </summary>
    public __TypeIdentGen__cr06pj9kmxbs__ CodeActionKind;
}


/*TStruc*/
public class Properties
{
    /// <summary>
    /// The properties that a client can resolve lazily.
    /// </summary>
    public string[] Properties;
}


/*TStruc*/
public class ValueSet
{
    /// <summary>
    /// The folding range kind values the client supports. When this
    /// property exists the client also guarantees that it will
    /// handle values outside its set gracefully and falls back
    /// to a default value when unknown.
    /// </summary>
    public /*TOpt*/FoldingRangeKind[] ValueSet;
}


/*TStruc*/
public class CollapsedText
{
    /// <summary>
    /// If set, the client signals that it supports setting collapsedText on
    /// folding ranges to display custom labels instead of the default text.
    /// 
    /// @since 3.17.0
    /// </summary>
    public bool? CollapsedText;
}


/*TStruc*/
public class ValueSet
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public DiagnosticTag[] ValueSet;
}


/*TOr*/
public abstract record BooleanOrAnyByString
{
    private BooleanOrAnyByString() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrAnyByString;
    public sealed record AnyByString(/*TOpt*/System.Collections.Generic.Dictionary<string, object> AnyByString) : BooleanOrAnyByString;
}

/*TStruc*/
public class Delta
{
    /// <summary>
    /// The client will send the `textDocument/semanticTokens/full/delta` request if
    /// the server provides a corresponding handler.
    /// </summary>
    public bool? Delta;
}


/*TOr*/
public abstract record BooleanOrDelta
{
    private BooleanOrDelta() { }
    public sealed record Boolean(bool? Boolean) : BooleanOrDelta;
    public sealed record Delta(/*TOpt*/__TypeIdentGen__cr06pj9lga5x__ Delta) : BooleanOrDelta;
}

/*TStruc*/
public class RangeFull
{
    /// <summary>
    /// The client will send the `textDocument/semanticTokens/range` request if
    /// the server provides a corresponding handler.
    /// 
    /// This object has "OneOf" union-type semantics: only at-most one field in it is ever set, all others will be null/undefined/nil/empty/etc.
    /// </summary>
    public /*TOpt*/__TypeIdentGen__cr06pj9lcu9k__ Range;
    /// <summary>
    /// The client will send the `textDocument/semanticTokens/full` request if
    /// the server provides a corresponding handler.
    /// 
    /// This object has "OneOf" union-type semantics: only at-most one field in it is ever set, all others will be null/undefined/nil/empty/etc.
    /// </summary>
    public /*TOpt*/__TypeIdentGen__cr06pj9lf28z__ Full;
}


/*TStruc*/
public class Properties
{
    /// <summary>
    /// The properties that a client can resolve lazily.
    /// </summary>
    public string[] Properties;
}


/*TStruc*/
public class AdditionalPropertiesSupport
{
    /// <summary>
    /// Whether the client supports additional attributes which
    /// are preserved and send back to the server in the
    /// request's response.
    /// </summary>
    public bool? AdditionalPropertiesSupport;
}

