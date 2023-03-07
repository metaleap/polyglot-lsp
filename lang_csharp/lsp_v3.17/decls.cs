/// <summary>Language Server Protocol (LSP) v3.17 SDK for C#: auto-generated via github.com/metaleap/polyglot-vsx-and-lsp/gen/cmd/gen_lsp</summary>
namespace lsp;


/*TOr*/
public abstract record BooleanOrDocumentHighlightOptions
{
    private BooleanOrDocumentHighlightOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrDocumentHighlightOptions;
    public sealed record DocumentHighlightOptions(/*TOpt*/DocumentHighlightOptions? DocumentHighlightOptions) : BooleanOrDocumentHighlightOptions;
}

/*TOr*/
public abstract record BooleanOrRenameOptions
{
    private BooleanOrRenameOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrRenameOptions;
    public sealed record RenameOptions(/*TOpt*/RenameOptions? RenameOptions) : BooleanOrRenameOptions;
}

/*TOr*/
public abstract record BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions
{
    private BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions;
    public sealed record CallHierarchyOptions(/*TOpt*/CallHierarchyOptions? CallHierarchyOptions) : BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions;
    public sealed record CallHierarchyRegistrationOptions(/*TOpt*/CallHierarchyRegistrationOptions? CallHierarchyRegistrationOptions) : BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions
{
    private BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions;
    public sealed record LinkedEditingRangeOptions(/*TOpt*/LinkedEditingRangeOptions? LinkedEditingRangeOptions) : BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions;
    public sealed record LinkedEditingRangeRegistrationOptions(/*TOpt*/LinkedEditingRangeRegistrationOptions? LinkedEditingRangeRegistrationOptions) : BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions;
}

/*TOr*/
public abstract record StringOrMarkupContent___
{
    private StringOrMarkupContent___() { }
    public sealed record String(/*TOpt*/string? String) : StringOrMarkupContent___;
    public sealed record MarkupContent(/*TOpt*/MarkupContent? MarkupContent) : StringOrMarkupContent___;
}

/*TOr*/
public abstract record StringOrNotebookDocumentFilter_
{
    private StringOrNotebookDocumentFilter_() { }
    public sealed record String(/*TOpt*/string? String) : StringOrNotebookDocumentFilter_;
    public sealed record NotebookDocumentFilter(/*TOpt*/NotebookDocumentFilter NotebookDocumentFilter) : StringOrNotebookDocumentFilter_;
}

/*TOr*/
public abstract record BooleanOrDelta
{
    private BooleanOrDelta() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrDelta;
    public sealed record Delta(/*TOpt*/Delta Delta) : BooleanOrDelta;
}

/*TOr*/
public abstract record BooleanOrDefinitionOptions
{
    private BooleanOrDefinitionOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrDefinitionOptions;
    public sealed record DefinitionOptions(/*TOpt*/DefinitionOptions? DefinitionOptions) : BooleanOrDefinitionOptions;
}

/*TStruc*/
public class Properties_
{
    /// <summary>
    /// The properties that a client can resolve lazily.
    /// </summary>
    public string[] Properties;
}


/*TStruc*/
public class ValueSet______
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public SymbolTag[] ValueSet;
}


/*TOr*/
public abstract record BooleanOrImplementationOptionsOrImplementationRegistrationOptions
{
    private BooleanOrImplementationOptionsOrImplementationRegistrationOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrImplementationOptionsOrImplementationRegistrationOptions;
    public sealed record ImplementationOptions(/*TOpt*/ImplementationOptions? ImplementationOptions) : BooleanOrImplementationOptionsOrImplementationRegistrationOptions;
    public sealed record ImplementationRegistrationOptions(/*TOpt*/ImplementationRegistrationOptions? ImplementationRegistrationOptions) : BooleanOrImplementationOptionsOrImplementationRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrReferenceOptions
{
    private BooleanOrReferenceOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrReferenceOptions;
    public sealed record ReferenceOptions(/*TOpt*/ReferenceOptions? ReferenceOptions) : BooleanOrReferenceOptions;
}

/*TOr*/
public abstract record BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions
{
    private BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions;
    public sealed record FoldingRangeOptions(/*TOpt*/FoldingRangeOptions? FoldingRangeOptions) : BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions;
    public sealed record FoldingRangeRegistrationOptions(/*TOpt*/FoldingRangeRegistrationOptions? FoldingRangeRegistrationOptions) : BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions;
}

/*TStruc*/
public class ValueSet__
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public CompletionItemTag[] ValueSet;
}


/*TStruc*/
public class Properties__
{
    /// <summary>
    /// The properties that a client can resolve lazily.
    /// </summary>
    public string[] Properties;
}


/*TOr*/
public abstract record MarkupContentOrMarkedStringOrMarkedStrings
{
    private MarkupContentOrMarkedStringOrMarkedStrings() { }
    public sealed record MarkupContent(/*TOpt*/MarkupContent? MarkupContent) : MarkupContentOrMarkedStringOrMarkedStrings;
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


/*TOr*/
public abstract record TextEditOrAnnotatedTextEdit
{
    private TextEditOrAnnotatedTextEdit() { }
    public sealed record TextEdit(/*TOpt*/TextEdit? TextEdit) : TextEditOrAnnotatedTextEdit;
    public sealed record AnnotatedTextEdit(/*TOpt*/AnnotatedTextEdit? AnnotatedTextEdit) : TextEditOrAnnotatedTextEdit;
}

/*TStruc*/
public class Language
{
    public string Language;
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
public class Uri
{
    public string Uri;
}


/*TOr*/
public abstract record IntegerOrString
{
    private IntegerOrString() { }
    public sealed record Integer(/*TOpt*/int? Integer) : IntegerOrString;
    public sealed record String(/*TOpt*/string? String) : IntegerOrString;
}

/*TStruc*/
public class Cancel_RetryOnContentModified
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


/*TStruc*/
public class ValueSet_____
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


/*TOr*/
public abstract record BooleanOrWorkspaceSymbolOptions
{
    private BooleanOrWorkspaceSymbolOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrWorkspaceSymbolOptions;
    public sealed record WorkspaceSymbolOptions(/*TOpt*/WorkspaceSymbolOptions? WorkspaceSymbolOptions) : BooleanOrWorkspaceSymbolOptions;
}

/*TOr*/
public abstract record StringOrBoolean
{
    private StringOrBoolean() { }
    public sealed record String(/*TOpt*/string? String) : StringOrBoolean;
    public sealed record Boolean(/*TOpt*/bool? Boolean) : StringOrBoolean;
}

/*TStruc*/
public class DocumentationFormat_ParameterInformation_ActiveParameterSupport
{
    /// <summary>
    /// Client supports the following content formats for the documentation
    /// property. The order describes the preferred format of the client.
    /// </summary>
    public /*TOpt*/MarkupKind[] DocumentationFormat;
    /// <summary>
    /// Client capabilities specific to parameter information.
    /// </summary>
    public /*TOpt*/LabelOffsetSupport ParameterInformation;
    /// <summary>
    /// The client supports the `activeParameter` property on `SignatureInformation`
    /// literal.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/bool? ActiveParameterSupport;
}


/*TOr*/
public abstract record FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport
{
    private FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport() { }
    public sealed record FullDocumentDiagnosticReport(/*TOpt*/FullDocumentDiagnosticReport? FullDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport;
    public sealed record UnchangedDocumentDiagnosticReport(/*TOpt*/UnchangedDocumentDiagnosticReport? UnchangedDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport;
}

/*TStruc*/
public class SnippetSupport_CommitCharactersSupport_DocumentationFormat_DeprecatedSupport_PreselectSupport_TagSupport_InsertReplaceSupport_ResolveSupport_InsertTextModeSupport_LabelDetailsSupport
{
    /// <summary>
    /// Client supports snippets as insert text.
    /// 
    /// A snippet can define tab stops and placeholders with `$1`, `$2`
    /// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
    /// the end of the snippet. Placeholders with equal identifiers are linked,
    /// that is typing in one will update others too.
    /// </summary>
    public /*TOpt*/bool? SnippetSupport;
    /// <summary>
    /// Client supports commit characters on a completion item.
    /// </summary>
    public /*TOpt*/bool? CommitCharactersSupport;
    /// <summary>
    /// Client supports the following content formats for the documentation
    /// property. The order describes the preferred format of the client.
    /// </summary>
    public /*TOpt*/MarkupKind[] DocumentationFormat;
    /// <summary>
    /// Client supports the deprecated property on a completion item.
    /// </summary>
    public /*TOpt*/bool? DeprecatedSupport;
    /// <summary>
    /// Client supports the preselect property on a completion item.
    /// </summary>
    public /*TOpt*/bool? PreselectSupport;
    /// <summary>
    /// Client supports the tag property on a completion item. Clients supporting
    /// tags have to handle unknown tags gracefully. Clients especially need to
    /// preserve unknown tags when sending a completion item back to the server in
    /// a resolve call.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/ValueSet__ TagSupport;
    /// <summary>
    /// Client support insert replace edit to control different behavior if a
    /// completion item is inserted in the text or should replace text.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/bool? InsertReplaceSupport;
    /// <summary>
    /// Indicates which properties a client can resolve lazily on a completion
    /// item. Before version 3.16.0 only the predefined properties `documentation`
    /// and `details` could be resolved lazily.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/Properties_ ResolveSupport;
    /// <summary>
    /// The client supports the `insertTextMode` property on
    /// a completion item to override the whitespace handling mode
    /// as defined by the client (see `insertTextMode`).
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/ValueSet___ InsertTextModeSupport;
    /// <summary>
    /// The client has support for completion item label
    /// details (see also `CompletionItemLabelDetails`).
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/bool? LabelDetailsSupport;
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
    public /*TOpt*/bool? LabelDetailsSupport;
}


/*TOr*/
public abstract record StringOrNotebookDocumentFilter
{
    private StringOrNotebookDocumentFilter() { }
    public sealed record String(/*TOpt*/string? String) : StringOrNotebookDocumentFilter;
    public sealed record NotebookDocumentFilter(/*TOpt*/NotebookDocumentFilter NotebookDocumentFilter) : StringOrNotebookDocumentFilter;
}

/*TOr*/
public abstract record StringOrUintegerWithUinteger
{
    private StringOrUintegerWithUinteger() { }
    public sealed record String(/*TOpt*/string? String) : StringOrUintegerWithUinteger;
    public sealed record UintegerWithUinteger(/*TOpt*/UintegerWithUinteger? UintegerWithUinteger) : StringOrUintegerWithUinteger;
}

/*TOr*/
public abstract record StringOrInlayHintLabelParts
{
    private StringOrInlayHintLabelParts() { }
    public sealed record String(/*TOpt*/string? String) : StringOrInlayHintLabelParts;
    public sealed record InlayHintLabelParts(/*TOpt*/InlayHintLabelPart[] InlayHintLabelParts) : StringOrInlayHintLabelParts;
}

/*TOr*/
public abstract record BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions
{
    private BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions;
    public sealed record InlineValueOptions(/*TOpt*/InlineValueOptions? InlineValueOptions) : BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions;
    public sealed record InlineValueRegistrationOptions(/*TOpt*/InlineValueRegistrationOptions? InlineValueRegistrationOptions) : BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions;
}

/*TOr*/
public abstract record FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport_
{
    private FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport_() { }
    public sealed record FullDocumentDiagnosticReport(/*TOpt*/FullDocumentDiagnosticReport? FullDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport_;
    public sealed record UnchangedDocumentDiagnosticReport(/*TOpt*/UnchangedDocumentDiagnosticReport? UnchangedDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport_;
}

/*TOr*/
public abstract record BooleanOrHoverOptions
{
    private BooleanOrHoverOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrHoverOptions;
    public sealed record HoverOptions(/*TOpt*/HoverOptions? HoverOptions) : BooleanOrHoverOptions;
}

/*TStruc*/
public class ValueSet_
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public SymbolTag[] ValueSet;
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
    public /*TOpt*/bool? LabelOffsetSupport;
}


/*TOr*/
public abstract record TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile
{
    private TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile() { }
    public sealed record TextDocumentEdit(/*TOpt*/TextDocumentEdit? TextDocumentEdit) : TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile;
    public sealed record CreateFile(/*TOpt*/CreateFile? CreateFile) : TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile;
    public sealed record RenameFile(/*TOpt*/RenameFile? RenameFile) : TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile;
    public sealed record DeleteFile(/*TOpt*/DeleteFile? DeleteFile) : TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile;
}

/*TOr*/
public abstract record StringOrMarkupContent__
{
    private StringOrMarkupContent__() { }
    public sealed record String(/*TOpt*/string? String) : StringOrMarkupContent__;
    public sealed record MarkupContent(/*TOpt*/MarkupContent? MarkupContent) : StringOrMarkupContent__;
}

/*TOr*/
public abstract record TextEditOrInsertReplaceEdit
{
    private TextEditOrInsertReplaceEdit() { }
    public sealed record TextEdit(/*TOpt*/TextEdit? TextEdit) : TextEditOrInsertReplaceEdit;
    public sealed record InsertReplaceEdit(/*TOpt*/InsertReplaceEdit? InsertReplaceEdit) : TextEditOrInsertReplaceEdit;
}

/*TOr*/
public abstract record LocationOrUri
{
    private LocationOrUri() { }
    public sealed record Location(/*TOpt*/Location? Location) : LocationOrUri;
    public sealed record Uri(/*TOpt*/Uri Uri) : LocationOrUri;
}

/*TStruc*/
public class Structure_Data_TextContent
{
    /// <summary>
    /// Changes to the cell structure to add or
    /// remove cells.
    /// </summary>
    public /*TOpt*/Array_DidOpen_DidClose Structure;
    /// <summary>
    /// Changes to notebook cells properties like its
    /// kind, execution summary or metadata.
    /// </summary>
    public /*TOpt*/NotebookCell[] Data;
    /// <summary>
    /// Changes to the text content of notebook cells.
    /// </summary>
    public /*TOpt*/Document_Changes[] TextContent;
}


/*TOr*/
public abstract record BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions
{
    private BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions;
    public sealed record DocumentColorOptions(/*TOpt*/DocumentColorOptions? DocumentColorOptions) : BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions;
    public sealed record DocumentColorRegistrationOptions(/*TOpt*/DocumentColorRegistrationOptions? DocumentColorRegistrationOptions) : BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions
{
    private BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions;
    public sealed record InlayHintOptions(/*TOpt*/InlayHintOptions? InlayHintOptions) : BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions;
    public sealed record InlayHintRegistrationOptions(/*TOpt*/InlayHintRegistrationOptions? InlayHintRegistrationOptions) : BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions;
}

/*TStruc*/
public class GroupsOnLabel
{
    /// <summary>
    /// Whether the client groups edits with equal labels into tree nodes,
    /// for instance all edits labelled with "Changes in Strings" would
    /// be a tree node.
    /// </summary>
    public /*TOpt*/bool? GroupsOnLabel;
}


/*TOr*/
public abstract record StringOrStrings
{
    private StringOrStrings() { }
    public sealed record String(/*TOpt*/string? String) : StringOrStrings;
    public sealed record Strings(/*TOpt*/string[] Strings) : StringOrStrings;
}

/*TOr*/
public abstract record StringOrMarkupContent_
{
    private StringOrMarkupContent_() { }
    public sealed record String(/*TOpt*/string? String) : StringOrMarkupContent_;
    public sealed record MarkupContent(/*TOpt*/MarkupContent? MarkupContent) : StringOrMarkupContent_;
}

/*TStruc*/
public class CodeActionKind
{
    /// <summary>
    /// The code action kind is support with the following value
    /// set.
    /// </summary>
    public ValueSet_______ CodeActionKind;
}


/*TStruc*/
public class ValueSet___
{
    public InsertTextMode[] ValueSet;
}


/*TStruc*/
public class ValueSet____
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


/*TOr*/
public abstract record BooleanOrMonikerOptionsOrMonikerRegistrationOptions
{
    private BooleanOrMonikerOptionsOrMonikerRegistrationOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrMonikerOptionsOrMonikerRegistrationOptions;
    public sealed record MonikerOptions(/*TOpt*/MonikerOptions? MonikerOptions) : BooleanOrMonikerOptionsOrMonikerRegistrationOptions;
    public sealed record MonikerRegistrationOptions(/*TOpt*/MonikerRegistrationOptions? MonikerRegistrationOptions) : BooleanOrMonikerOptionsOrMonikerRegistrationOptions;
}

/*TOr*/
public abstract record DiagnosticOptionsOrDiagnosticRegistrationOptions
{
    private DiagnosticOptionsOrDiagnosticRegistrationOptions() { }
    public sealed record DiagnosticOptions(/*TOpt*/DiagnosticOptions? DiagnosticOptions) : DiagnosticOptionsOrDiagnosticRegistrationOptions;
    public sealed record DiagnosticRegistrationOptions(/*TOpt*/DiagnosticRegistrationOptions? DiagnosticRegistrationOptions) : DiagnosticOptionsOrDiagnosticRegistrationOptions;
}

/*TOr*/
public abstract record IntegerOrString_
{
    private IntegerOrString_() { }
    public sealed record Integer(/*TOpt*/int? Integer) : IntegerOrString_;
    public sealed record String(/*TOpt*/string? String) : IntegerOrString_;
}

/*TOr*/
public abstract record TextDocumentSyncOptionsOrTextDocumentSyncKind
{
    private TextDocumentSyncOptionsOrTextDocumentSyncKind() { }
    public sealed record TextDocumentSyncOptions(/*TOpt*/TextDocumentSyncOptions? TextDocumentSyncOptions) : TextDocumentSyncOptionsOrTextDocumentSyncKind;
    public sealed record TextDocumentSyncKind(/*TOpt*/TextDocumentSyncKind TextDocumentSyncKind) : TextDocumentSyncOptionsOrTextDocumentSyncKind;
}

/*TOr*/
public abstract record BooleanOrDocumentSymbolOptions
{
    private BooleanOrDocumentSymbolOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrDocumentSymbolOptions;
    public sealed record DocumentSymbolOptions(/*TOpt*/DocumentSymbolOptions? DocumentSymbolOptions) : BooleanOrDocumentSymbolOptions;
}

/*TStruc*/
public class Name_Version_
{
    /// <summary>
    /// The name of the client as defined by the client.
    /// </summary>
    public string Name;
    /// <summary>
    /// The client's version as defined by the client.
    /// </summary>
    public /*TOpt*/string? Version;
}


/*TOr*/
public abstract record BooleanOrCodeActionOptions
{
    private BooleanOrCodeActionOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrCodeActionOptions;
    public sealed record CodeActionOptions(/*TOpt*/CodeActionOptions? CodeActionOptions) : BooleanOrCodeActionOptions;
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


/*TOr*/
public abstract record BooleanOrDelta_
{
    private BooleanOrDelta_() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrDelta_;
    public sealed record Delta(/*TOpt*/Delta_ Delta) : BooleanOrDelta_;
}

/*TOr*/
public abstract record RangeOrInsert_Replace
{
    private RangeOrInsert_Replace() { }
    public sealed record Range(/*TOpt*/Range? Range) : RangeOrInsert_Replace;
    public sealed record Insert_Replace(/*TOpt*/Insert_Replace Insert_Replace) : RangeOrInsert_Replace;
}

/*TStruc*/
public class CommitCharacters_EditRange_InsertTextFormat_InsertTextMode_Data
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
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/RangeOrInsert_Replace EditRange;
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
public abstract record BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions
{
    private BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions;
    public sealed record DeclarationOptions(/*TOpt*/DeclarationOptions? DeclarationOptions) : BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions;
    public sealed record DeclarationRegistrationOptions(/*TOpt*/DeclarationRegistrationOptions? DeclarationRegistrationOptions) : BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions;
}

/*TStruc*/
public class Notebook_Cells
{
    /// <summary>
    /// The notebook to be synced If a string
    /// value is provided it matches against the
    /// notebook type. '*' matches every notebook.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/StringOrNotebookDocumentFilter Notebook;
    /// <summary>
    /// The cells of the matching notebook to be synced.
    /// </summary>
    public /*TOpt*/Language[] Cells;
}


/*TStruc*/
public class ValueSet_______
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
public class CollapsedText
{
    /// <summary>
    /// If set, the client signals that it supports setting collapsedText on
    /// folding ranges to display custom labels instead of the default text.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/bool? CollapsedText;
}


/*TStruc*/
public class Range_Full
{
    /// <summary>
    /// The client will send the `textDocument/semanticTokens/range` request if
    /// the server provides a corresponding handler.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrAnyByString_ Range;
    /// <summary>
    /// The client will send the `textDocument/semanticTokens/full` request if
    /// the server provides a corresponding handler.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrDelta_ Full;
}


/*TOr*/
public abstract record FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport__
{
    private FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport__() { }
    public sealed record FullDocumentDiagnosticReport(/*TOpt*/FullDocumentDiagnosticReport? FullDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport__;
    public sealed record UnchangedDocumentDiagnosticReport(/*TOpt*/UnchangedDocumentDiagnosticReport? UnchangedDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport__;
}

/*TOr*/
public abstract record NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions
{
    private NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions() { }
    public sealed record NotebookDocumentSyncOptions(/*TOpt*/NotebookDocumentSyncOptions? NotebookDocumentSyncOptions) : NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions;
    public sealed record NotebookDocumentSyncRegistrationOptions(/*TOpt*/NotebookDocumentSyncRegistrationOptions? NotebookDocumentSyncRegistrationOptions) : NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions;
}

/*TOr*/
public abstract record BooleanOrDocumentRangeFormattingOptions
{
    private BooleanOrDocumentRangeFormattingOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrDocumentRangeFormattingOptions;
    public sealed record DocumentRangeFormattingOptions(/*TOpt*/DocumentRangeFormattingOptions? DocumentRangeFormattingOptions) : BooleanOrDocumentRangeFormattingOptions;
}

/*TOr*/
public abstract record BooleanOrSaveOptions
{
    private BooleanOrSaveOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrSaveOptions;
    public sealed record SaveOptions(/*TOpt*/SaveOptions? SaveOptions) : BooleanOrSaveOptions;
}

/*TTup*/
public class UintegerWithUinteger
{
    public uint Uinteger0;
    public uint Uinteger1;
}


/*TStruc*/
public class Delta_
{
    /// <summary>
    /// The client will send the `textDocument/semanticTokens/full/delta` request if
    /// the server provides a corresponding handler.
    /// </summary>
    public /*TOpt*/bool? Delta;
}


/*TStruc*/
public class Name_Version
{
    /// <summary>
    /// The name of the server as defined by the server.
    /// </summary>
    public string Name;
    /// <summary>
    /// The server's version as defined by the server.
    /// </summary>
    public /*TOpt*/string? Version;
}


/*TStruc*/
public class Document_Changes
{
    public VersionedTextDocumentIdentifier Document;
    public TextDocumentContentChangeEvent[] Changes;
}


/*TOr*/
public abstract record BooleanOrAnyByString_
{
    private BooleanOrAnyByString_() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrAnyByString_;
    public sealed record AnyByString(/*TOpt*/System.Collections.Generic.Dictionary<string, object> AnyByString) : BooleanOrAnyByString_;
}

/*TStruc*/
public class Properties___
{
    /// <summary>
    /// The properties that a client can resolve lazily.
    /// </summary>
    public string[] Properties;
}


/*TOr*/
public abstract record BooleanOrAnyByString
{
    private BooleanOrAnyByString() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrAnyByString;
    public sealed record AnyByString(/*TOpt*/System.Collections.Generic.Dictionary<string, object> AnyByString) : BooleanOrAnyByString;
}

/*TOr*/
public abstract record BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions
{
    private BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions;
    public sealed record TypeDefinitionOptions(/*TOpt*/TypeDefinitionOptions? TypeDefinitionOptions) : BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions;
    public sealed record TypeDefinitionRegistrationOptions(/*TOpt*/TypeDefinitionRegistrationOptions? TypeDefinitionRegistrationOptions) : BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions;
}

/*TOr*/
public abstract record SemanticTokensOptionsOrSemanticTokensRegistrationOptions
{
    private SemanticTokensOptionsOrSemanticTokensRegistrationOptions() { }
    public sealed record SemanticTokensOptions(/*TOpt*/SemanticTokensOptions? SemanticTokensOptions) : SemanticTokensOptionsOrSemanticTokensRegistrationOptions;
    public sealed record SemanticTokensRegistrationOptions(/*TOpt*/SemanticTokensRegistrationOptions? SemanticTokensRegistrationOptions) : SemanticTokensOptionsOrSemanticTokensRegistrationOptions;
}

/*TStruc*/
public class WorkspaceFolders_FileOperations
{
    /// <summary>
    /// The server supports workspace folder.
    /// 
    /// @since 3.6.0
    /// </summary>
    public /*TOpt*/WorkspaceFoldersServerCapabilities? WorkspaceFolders;
    /// <summary>
    /// The server is interested in notifications/requests for operations on files.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/FileOperationOptions? FileOperations;
}


/*TOr*/
public abstract record StringOrMarkupContent____
{
    private StringOrMarkupContent____() { }
    public sealed record String(/*TOpt*/string? String) : StringOrMarkupContent____;
    public sealed record MarkupContent(/*TOpt*/MarkupContent? MarkupContent) : StringOrMarkupContent____;
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


/*TOr*/
public abstract record StringOrMarkupContent
{
    private StringOrMarkupContent() { }
    public sealed record String(/*TOpt*/string? String) : StringOrMarkupContent;
    public sealed record MarkupContent(/*TOpt*/MarkupContent? MarkupContent) : StringOrMarkupContent;
}

/*TOr*/
public abstract record BooleanOrDocumentFormattingOptions
{
    private BooleanOrDocumentFormattingOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrDocumentFormattingOptions;
    public sealed record DocumentFormattingOptions(/*TOpt*/DocumentFormattingOptions? DocumentFormattingOptions) : BooleanOrDocumentFormattingOptions;
}

/*TStruc*/
public class ValueSet________
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
public class Insert_Replace
{
    public Range Insert;
    public Range Replace;
}


/*TStruc*/
public class Delta
{
    /// <summary>
    /// The server supports deltas for full documents.
    /// </summary>
    public /*TOpt*/bool? Delta;
}


/*TOr*/
public abstract record BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions
{
    private BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions;
    public sealed record TypeHierarchyOptions(/*TOpt*/TypeHierarchyOptions? TypeHierarchyOptions) : BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions;
    public sealed record TypeHierarchyRegistrationOptions(/*TOpt*/TypeHierarchyRegistrationOptions? TypeHierarchyRegistrationOptions) : BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions;
}

/*TOr*/
public abstract record WorkspaceFolderOrURI
{
    private WorkspaceFolderOrURI() { }
    public sealed record WorkspaceFolder(/*TOpt*/WorkspaceFolder? WorkspaceFolder) : WorkspaceFolderOrURI;
    public sealed record URI(/*TOpt*/string URI) : WorkspaceFolderOrURI;
}

/*TStruc*/
public class ValueSet_________
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public DiagnosticTag[] ValueSet;
}


/*TStruc*/
public class AdditionalPropertiesSupport
{
    /// <summary>
    /// Whether the client supports additional attributes which
    /// are preserved and send back to the server in the
    /// request's response.
    /// </summary>
    public /*TOpt*/bool? AdditionalPropertiesSupport;
}


/*TStruc*/
public class Array_DidOpen_DidClose
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


/*TOr*/
public abstract record BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions
{
    private BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions() { }
    public sealed record Boolean(/*TOpt*/bool? Boolean) : BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions;
    public sealed record SelectionRangeOptions(/*TOpt*/SelectionRangeOptions? SelectionRangeOptions) : BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions;
    public sealed record SelectionRangeRegistrationOptions(/*TOpt*/SelectionRangeRegistrationOptions? SelectionRangeRegistrationOptions) : BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions;
}

