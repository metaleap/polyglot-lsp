/// <summary>Language Server Protocol (LSP) v3.17 SDK for C#: auto-generated via github.com/metaleap/polyglot-lsp/gen/cmd/gen_lsp</summary>
namespace lsp;

using Definition = LocationOrLocations;
using DefinitionLink = LocationLink;
using Declaration = LocationOrLocations;
using DeclarationLink = LocationLink;
using InlineValue = InlineValueTextOrInlineValueVariableLookupOrInlineValueEvaluatableExpression;
using DocumentDiagnosticReport = RelatedFullDocumentDiagnosticReportOrRelatedUnchangedDocumentDiagnosticReport;
using PrepareRenameResult = RangeOrRangeWithPlaceholderStringOrDefaultBehaviorBoolean;
using ProgressToken = IntegerOrString;
using ChangeAnnotationIdentifier = System.String;
using WorkspaceDocumentDiagnosticReport = WorkspaceFullDocumentDiagnosticReportOrWorkspaceUnchangedDocumentDiagnosticReport;
using TextDocumentContentChangeEvent = RangeWithRangeLengthUintegerWithTextStringOrTextString;
using MarkedString = StringOrLanguageStringWithValueString;
using DocumentFilter = TextDocumentFilterOrNotebookCellTextDocumentFilter;
using LSPObject = System.Collections.Generic.Dictionary<System.String, LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull>;
using GlobPattern = PatternOrRelativePattern;
using TextDocumentFilter = LanguageStringWithSchemeStringWithPatternString;
using NotebookDocumentFilter = NotebookTypeStringWithSchemeStringWithPatternString;
using Pattern = System.String;
using LSPAny = LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull;


/*TStruc*/
public class TextString
{
    /// <summary>
    /// The new text of the whole document.
    /// </summary>
    public System.String Text;
}


/*TOr*/
public class TextEditOrAnnotatedTextEdit
{
    public  /*TOpt*/TextEdit? TextEdit;
    public  /*TOpt*/AnnotatedTextEdit? AnnotatedTextEdit;
}

/*TStruc*/
public class NameStringWithVersionString_
{
    /// <summary>
    /// The name of the client as defined by the client.
    /// </summary>
    public System.String Name;
    /// <summary>
    /// The client's version as defined by the client.
    /// </summary>
    public /*TOpt*/System.String? Version;
}


/*TOr*/
public class BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/TypeDefinitionOptions? TypeDefinitionOptions;
    public  /*TOpt*/TypeDefinitionRegistrationOptions? TypeDefinitionRegistrationOptions;
}

/*TStruc*/
public class RangeBooleanOrAnyByStringWithFullBooleanOrDeltaBoolean
{
    /// <summary>
    /// The client will send the `textDocument/semanticTokens/range` request if
    /// the server provides a corresponding handler.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrAnyByString Range;
    /// <summary>
    /// The client will send the `textDocument/semanticTokens/full` request if
    /// the server provides a corresponding handler.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrDeltaBoolean_ Full;
}


/*TOr*/
public class LocationOrLocations
{
    public  /*TOpt*/Location? Location;
    public  /*TOpt*/Location[] Locations;
}

/*TStruc*/
public class LanguageStringWithValueString
{
    public System.String Language;
    public System.String Value;
}


/*TOr*/
public class TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile
{
    public  /*TOpt*/TextDocumentEdit? TextDocumentEdit;
    public  /*TOpt*/CreateFile? CreateFile;
    public  /*TOpt*/RenameFile? RenameFile;
    public  /*TOpt*/DeleteFile? DeleteFile;
}

/*TOr*/
public class BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/DeclarationOptions? DeclarationOptions;
    public  /*TOpt*/DeclarationRegistrationOptions? DeclarationRegistrationOptions;
}

/*TOr*/
public class BooleanOrDefinitionOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/DefinitionOptions? DefinitionOptions;
}

/*TOr*/
public class WorkspaceFullDocumentDiagnosticReportOrWorkspaceUnchangedDocumentDiagnosticReport
{
    public  /*TOpt*/WorkspaceFullDocumentDiagnosticReport? WorkspaceFullDocumentDiagnosticReport;
    public  /*TOpt*/WorkspaceUnchangedDocumentDiagnosticReport? WorkspaceUnchangedDocumentDiagnosticReport;
}

/*TOr*/
public class BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/SelectionRangeOptions? SelectionRangeOptions;
    public  /*TOpt*/SelectionRangeRegistrationOptions? SelectionRangeRegistrationOptions;
}

/*TOr*/
public class BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/TypeHierarchyOptions? TypeHierarchyOptions;
    public  /*TOpt*/TypeHierarchyRegistrationOptions? TypeHierarchyRegistrationOptions;
}

/*TOr*/
public class DiagnosticOptionsOrDiagnosticRegistrationOptions
{
    public  /*TOpt*/DiagnosticOptions? DiagnosticOptions;
    public  /*TOpt*/DiagnosticRegistrationOptions? DiagnosticRegistrationOptions;
}

/*TStruc*/
public class WorkspaceFoldersWorkspaceFoldersServerCapabilitiesWithFileOperationsFileOperationOptions
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


/*TStruc*/
public class DeltaBoolean_
{
    /// <summary>
    /// The client will send the `textDocument/semanticTokens/full/delta` request if
    /// the server provides a corresponding handler.
    /// </summary>
    public /*TOpt*/System.Boolean? Delta;
}


/*TOr*/
public class BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/LinkedEditingRangeOptions? LinkedEditingRangeOptions;
    public  /*TOpt*/LinkedEditingRangeRegistrationOptions? LinkedEditingRangeRegistrationOptions;
}

/*TOr*/
public class BooleanOrSaveOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/SaveOptions? SaveOptions;
}

/*TOr*/
public class LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull
{
    public  /*TOpt*/LSPObject LSPObject;
    public  /*TOpt*/LSPAny[] LSPArray;
    public  /*TOpt*/System.String? String;
    public  /*TOpt*/System.Int64? Integer;
    public  /*TOpt*/System.UInt64? Uinteger;
    public  /*TOpt*/System.Double? Decimal;
    public  /*TOpt*/System.Boolean? Boolean;
}

/*TOr*/
public class StringOrLanguageStringWithValueString
{
    public  /*TOpt*/System.String? String;
    public  /*TOpt*/LanguageStringWithValueString LanguageStringWithValueString;
}

/*TStruc*/
public class LanguageStringWithSchemeStringWithPatternString
{
    /// <summary>
    /// A language id, like `typescript`.
    /// </summary>
    public /*TOpt*/System.String? Language;
    /// <summary>
    /// A Uri <c>Uri.scheme</c>, like `file` or `untitled`.
    /// </summary>
    public /*TOpt*/System.String? Scheme;
    /// <summary>
    /// A glob pattern, like **​/*.{ts,js}. See TextDocumentFilter for examples.
    /// </summary>
    public /*TOpt*/System.String? Pattern;
}


/*TOr*/
public class BooleanOrDeltaBoolean
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/DeltaBoolean DeltaBoolean;
}

/*TOr*/
public class BooleanOrHoverOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/HoverOptions? HoverOptions;
}

/*TOr*/
public class BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/CallHierarchyOptions? CallHierarchyOptions;
    public  /*TOpt*/CallHierarchyRegistrationOptions? CallHierarchyRegistrationOptions;
}

/*TStruc*/
public class ValueSetSymbolTags
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public SymbolTag[] ValueSet;
}


/*TStruc*/
public class LabelOffsetSupportBoolean
{
    /// <summary>
    /// The client supports processing label offsets instead of a
    /// simple label string.
    /// 
    /// @since 3.14.0
    /// </summary>
    public /*TOpt*/System.Boolean? LabelOffsetSupport;
}


/*TStruc*/
public class ValueSetFoldingRangeKinds
{
    /// <summary>
    /// The folding range kind values the client supports. When this
    /// property exists the client also guarantees that it will
    /// handle values outside its set gracefully and falls back
    /// to a default value when unknown.
    /// </summary>
    public /*TOpt*//*FoldingRangeKind*/string[] ValueSet;
}


/*TStruc*/
public class CollapsedTextBoolean
{
    /// <summary>
    /// If set, the client signals that it supports setting collapsedText on
    /// folding ranges to display custom labels instead of the default text.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/System.Boolean? CollapsedText;
}


/*TStruc*/
public class NameStringWithVersionString
{
    /// <summary>
    /// The name of the server as defined by the server.
    /// </summary>
    public System.String Name;
    /// <summary>
    /// The server's version as defined by the server.
    /// </summary>
    public /*TOpt*/System.String? Version;
}


/*TOr*/
public class StringOrStrings
{
    public  /*TOpt*/System.String? String;
    public  /*TOpt*/System.String[] Strings;
}

/*TOr*/
public class TextDocumentSyncOptionsOrTextDocumentSyncKind
{
    public  /*TOpt*/TextDocumentSyncOptions? TextDocumentSyncOptions;
    public  /*TOpt*/TextDocumentSyncKind TextDocumentSyncKind;
}

/*TOr*/
public class NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions
{
    public  /*TOpt*/NotebookDocumentSyncOptions? NotebookDocumentSyncOptions;
    public  /*TOpt*/NotebookDocumentSyncRegistrationOptions? NotebookDocumentSyncRegistrationOptions;
}

/*TStruc*/
public class ValueSetCodeActionKinds
{
    /// <summary>
    /// The code action kind values the client supports. When this
    /// property exists the client also guarantees that it will
    /// handle values outside its set gracefully and falls back
    /// to a default value when unknown.
    /// </summary>
    public /*CodeActionKind*/string[] ValueSet;
}


/*TStruc*/
public class CodeActionKindValueSetCodeActionKinds
{
    /// <summary>
    /// The code action kind is support with the following value
    /// set.
    /// </summary>
    public ValueSetCodeActionKinds CodeActionKind;
}


/*TOr*/
public class SemanticTokensOptionsOrSemanticTokensRegistrationOptions
{
    public  /*TOpt*/SemanticTokensOptions? SemanticTokensOptions;
    public  /*TOpt*/SemanticTokensRegistrationOptions? SemanticTokensRegistrationOptions;
}

/*TStruc*/
public class LanguageString
{
    public System.String Language;
}


/*TStruc*/
public class RangeWithRangeLengthUintegerWithTextString
{
    /// <summary>
    /// The range of the document that changed.
    /// </summary>
    public Range Range;
    /// <summary>
    /// The optional length of the range that got replaced.
    /// 
    /// @deprecated use range instead.
    /// </summary>
    public /*TOpt*/System.UInt64? RangeLength;
    /// <summary>
    /// The new text for the provided range.
    /// </summary>
    public System.String Text;
}


/*TOr*/
public class RangeWithRangeLengthUintegerWithTextStringOrTextString
{
    public  /*TOpt*/RangeWithRangeLengthUintegerWithTextString RangeWithRangeLengthUintegerWithTextString;
    public  /*TOpt*/TextString TextString;
}

/*TOr*/
public class TextDocumentFilterOrNotebookCellTextDocumentFilter
{
    public  /*TOpt*/TextDocumentFilter TextDocumentFilter;
    public  /*TOpt*/NotebookCellTextDocumentFilter? NotebookCellTextDocumentFilter;
}

/*TStruc*/
public class ArrayNotebookCellArrayChangeWithDidOpenTextDocumentItemsWithDidCloseTextDocumentIdentifiers
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
public class BooleanOrReferenceOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/ReferenceOptions? ReferenceOptions;
}

/*TOr*/
public class BooleanOrDocumentFormattingOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/DocumentFormattingOptions? DocumentFormattingOptions;
}

/*TStruc*/
public class GroupsOnLabelBoolean
{
    /// <summary>
    /// Whether the client groups edits with equal labels into tree nodes,
    /// for instance all edits labelled with "Changes in Strings" would
    /// be a tree node.
    /// </summary>
    public /*TOpt*/System.Boolean? GroupsOnLabel;
}


/*TOr*/
public class RelatedFullDocumentDiagnosticReportOrRelatedUnchangedDocumentDiagnosticReport
{
    public  /*TOpt*/RelatedFullDocumentDiagnosticReport? RelatedFullDocumentDiagnosticReport;
    public  /*TOpt*/RelatedUnchangedDocumentDiagnosticReport? RelatedUnchangedDocumentDiagnosticReport;
}

/*TStruc*/
public class RangeWithPlaceholderString
{
    public Range Range;
    public System.String Placeholder;
}


/*TOr*/
public class RangeOrInsertRangeWithReplaceRange
{
    public  /*TOpt*/Range? Range;
    public  /*TOpt*/InsertRangeWithReplaceRange InsertRangeWithReplaceRange;
}

/*TOr*/
public class BooleanOrDocumentSymbolOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/DocumentSymbolOptions? DocumentSymbolOptions;
}

/*TOr*/
public class BooleanOrDocumentRangeFormattingOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/DocumentRangeFormattingOptions? DocumentRangeFormattingOptions;
}

/*TOr*/
public class StringOrUintegerWithUinteger
{
    public  /*TOpt*/System.String? String;
    public  /*TOpt*/UintegerWithUinteger? UintegerWithUinteger;
}

/*TOr*/
public class StringOrNotebookDocumentFilter
{
    public  /*TOpt*/System.String? String;
    public  /*TOpt*/NotebookDocumentFilter NotebookDocumentFilter;
}

/*TStruc*/
public class SnippetSupportBooleanWithCommitCharactersSupportBooleanWithDocumentationFormatMarkupKindsWithDeprecatedSupportBooleanWithPreselectSupportBooleanWithTagSupportValueSetCompletionItemTagsWithInsertReplaceSupportBooleanWithResolveSupportPropertiesStringsWithInsertTextModeSupportValueSetInsertTextModesWithLabelDetailsSupportBoolean
{
    /// <summary>
    /// Client supports snippets as insert text.
    /// 
    /// A snippet can define tab stops and placeholders with `$1`, `$2`
    /// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
    /// the end of the snippet. Placeholders with equal identifiers are linked,
    /// that is typing in one will update others too.
    /// </summary>
    public /*TOpt*/System.Boolean? SnippetSupport;
    /// <summary>
    /// Client supports commit characters on a completion item.
    /// </summary>
    public /*TOpt*/System.Boolean? CommitCharactersSupport;
    /// <summary>
    /// Client supports the following content formats for the documentation
    /// property. The order describes the preferred format of the client.
    /// </summary>
    public /*TOpt*//*MarkupKind*/string[] DocumentationFormat;
    /// <summary>
    /// Client supports the deprecated property on a completion item.
    /// </summary>
    public /*TOpt*/System.Boolean? DeprecatedSupport;
    /// <summary>
    /// Client supports the preselect property on a completion item.
    /// </summary>
    public /*TOpt*/System.Boolean? PreselectSupport;
    /// <summary>
    /// Client supports the tag property on a completion item. Clients supporting
    /// tags have to handle unknown tags gracefully. Clients especially need to
    /// preserve unknown tags when sending a completion item back to the server in
    /// a resolve call.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/ValueSetCompletionItemTags TagSupport;
    /// <summary>
    /// Client support insert replace edit to control different behavior if a
    /// completion item is inserted in the text or should replace text.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.Boolean? InsertReplaceSupport;
    /// <summary>
    /// Indicates which properties a client can resolve lazily on a completion
    /// item. Before version 3.16.0 only the predefined properties `documentation`
    /// and `details` could be resolved lazily.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/PropertiesStrings_ ResolveSupport;
    /// <summary>
    /// The client supports the `insertTextMode` property on
    /// a completion item to override the whitespace handling mode
    /// as defined by the client (see `insertTextMode`).
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/ValueSetInsertTextModes InsertTextModeSupport;
    /// <summary>
    /// The client has support for completion item label
    /// details (see also `CompletionItemLabelDetails`).
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/System.Boolean? LabelDetailsSupport;
}


/*TStruc*/
public class DeltaBoolean
{
    /// <summary>
    /// The server supports deltas for full documents.
    /// </summary>
    public /*TOpt*/System.Boolean? Delta;
}


/*TStruc*/
public class DocumentVersionedTextDocumentIdentifierWithChangesTextDocumentContentChangeEvents
{
    public VersionedTextDocumentIdentifier Document;
    public TextDocumentContentChangeEvent[] Changes;
}


/*TStruc*/
public class StructureArrayNotebookCellArrayChangeWithDidOpenTextDocumentItemsWithDidCloseTextDocumentIdentifiersWithDataNotebookCellsWithTextContentDocumentVersionedTextDocumentIdentifierWithChangesTextDocumentContentChangeEventses
{
    /// <summary>
    /// Changes to the cell structure to add or
    /// remove cells.
    /// </summary>
    public /*TOpt*/ArrayNotebookCellArrayChangeWithDidOpenTextDocumentItemsWithDidCloseTextDocumentIdentifiers Structure;
    /// <summary>
    /// Changes to notebook cells properties like its
    /// kind, execution summary or metadata.
    /// </summary>
    public /*TOpt*/NotebookCell[] Data;
    /// <summary>
    /// Changes to the text content of notebook cells.
    /// </summary>
    public /*TOpt*/DocumentVersionedTextDocumentIdentifierWithChangesTextDocumentContentChangeEvents[] TextContent;
}


/*TOr*/
public class BooleanOrWorkspaceSymbolOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/WorkspaceSymbolOptions? WorkspaceSymbolOptions;
}

/*TOr*/
public class BooleanOrRenameOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/RenameOptions? RenameOptions;
}

/*TStruc*/
public class LabelDetailsSupportBoolean
{
    /// <summary>
    /// The server has support for completion item label
    /// details (see also `CompletionItemLabelDetails`) when
    /// receiving a completion item in a resolve call.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/System.Boolean? LabelDetailsSupport;
}


/*TStruc*/
public class ItemDefaultsStrings
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
    public /*TOpt*/System.String[] ItemDefaults;
}


/*TStruc*/
public class DefaultBehaviorBoolean
{
    public System.Boolean DefaultBehavior;
}


/*TOr*/
public class FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport
{
    public  /*TOpt*/FullDocumentDiagnosticReport? FullDocumentDiagnosticReport;
    public  /*TOpt*/UnchangedDocumentDiagnosticReport? UnchangedDocumentDiagnosticReport;
}

/*TOr*/
public class BooleanOrAnyByString
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/System.Collections.Generic.Dictionary<System.String, object> AnyByString;
}

/*TStruc*/
public class AdditionalPropertiesSupportBoolean
{
    /// <summary>
    /// Whether the client supports additional attributes which
    /// are preserved and send back to the server in the
    /// request's response.
    /// </summary>
    public /*TOpt*/System.Boolean? AdditionalPropertiesSupport;
}


/*TStruc*/
public class NotebookTypeStringWithSchemeStringWithPatternString
{
    /// <summary>
    /// The type of the enclosing notebook.
    /// </summary>
    public /*TOpt*/System.String? NotebookType;
    /// <summary>
    /// A Uri <c>Uri.scheme</c>, like `file` or `untitled`.
    /// </summary>
    public /*TOpt*/System.String? Scheme;
    /// <summary>
    /// A glob pattern.
    /// </summary>
    public /*TOpt*/System.String? Pattern;
}


/*TOr*/
public class StringOrMarkupContent
{
    public  /*TOpt*/System.String? String;
    public  /*TOpt*/MarkupContent? MarkupContent;
}

/*TStruc*/
public class InsertRangeWithReplaceRange
{
    public Range Insert;
    public Range Replace;
}


/*TStruc*/
public class ReasonString
{
    /// <summary>
    /// Human readable description of why the code action is currently disabled.
    /// 
    /// This is displayed in the code actions UI.
    /// </summary>
    public System.String Reason;
}


/*TOr*/
public class LocationOrUriDocumentUri
{
    public  /*TOpt*/Location? Location;
    public  /*TOpt*/UriDocumentUri UriDocumentUri;
}

/*TOr*/
public class MarkupContentOrMarkedStringOrMarkedStrings
{
    public  /*TOpt*/MarkupContent? MarkupContent;
    public  /*TOpt*/MarkedString MarkedString;
    public  /*TOpt*/MarkedString[] MarkedStrings;
}

/*TOr*/
public class BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/InlineValueOptions? InlineValueOptions;
    public  /*TOpt*/InlineValueRegistrationOptions? InlineValueRegistrationOptions;
}

/*TOr*/
public class BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/DocumentColorOptions? DocumentColorOptions;
    public  /*TOpt*/DocumentColorRegistrationOptions? DocumentColorRegistrationOptions;
}

/*TStruc*/
public class NotebookStringOrNotebookDocumentFilterWithCellsLanguageStrings
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
    public /*TOpt*/LanguageString[] Cells;
}


/*TStruc*/
public class PropertiesStrings
{
    /// <summary>
    /// The properties that a client can resolve lazily. Usually
    /// `location.range`
    /// </summary>
    public System.String[] Properties;
}


/*TStruc*/
public class ValueSetCompletionItemKinds
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
public class RangeOrRangeWithPlaceholderStringOrDefaultBehaviorBoolean
{
    public  /*TOpt*/Range? Range;
    public  /*TOpt*/RangeWithPlaceholderString RangeWithPlaceholderString;
    public  /*TOpt*/DefaultBehaviorBoolean DefaultBehaviorBoolean;
}

/*TOr*/
public class StringOrInlayHintLabelParts
{
    public  /*TOpt*/System.String? String;
    public  /*TOpt*/InlayHintLabelPart[] InlayHintLabelParts;
}

/*TStruc*/
public class UriDocumentUri
{
    public System.String Uri;
}


/*TOr*/
public class BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/FoldingRangeOptions? FoldingRangeOptions;
    public  /*TOpt*/FoldingRangeRegistrationOptions? FoldingRangeRegistrationOptions;
}

/*TStruc*/
public class ValueSetCompletionItemTags
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public CompletionItemTag[] ValueSet;
}


/*TOr*/
public class BooleanOrDeltaBoolean_
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/DeltaBoolean_ DeltaBoolean;
}

/*TTup*/
public class UintegerWithUinteger
{
    public System.UInt64 Uinteger0;
    public System.UInt64 Uinteger1;
}


/*TStruc*/
public class ValueSetSymbolKinds
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
public class IntegerOrString
{
    public  /*TOpt*/System.Int64? Integer;
    public  /*TOpt*/System.String? String;
}

/*TOr*/
public class PatternOrRelativePattern
{
    public  /*TOpt*/System.String? Pattern;
    public  /*TOpt*/RelativePattern? RelativePattern;
}

/*TOr*/
public class BooleanOrImplementationOptionsOrImplementationRegistrationOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/ImplementationOptions? ImplementationOptions;
    public  /*TOpt*/ImplementationRegistrationOptions? ImplementationRegistrationOptions;
}

/*TOr*/
public class BooleanOrMonikerOptionsOrMonikerRegistrationOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/MonikerOptions? MonikerOptions;
    public  /*TOpt*/MonikerRegistrationOptions? MonikerRegistrationOptions;
}

/*TOr*/
public class BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/InlayHintOptions? InlayHintOptions;
    public  /*TOpt*/InlayHintRegistrationOptions? InlayHintRegistrationOptions;
}

/*TOr*/
public class StringOrBoolean
{
    public  /*TOpt*/System.String? String;
    public  /*TOpt*/System.Boolean? Boolean;
}

/*TStruc*/
public class DocumentationFormatMarkupKindsWithParameterInformationLabelOffsetSupportBooleanWithActiveParameterSupportBoolean
{
    /// <summary>
    /// Client supports the following content formats for the documentation
    /// property. The order describes the preferred format of the client.
    /// </summary>
    public /*TOpt*//*MarkupKind*/string[] DocumentationFormat;
    /// <summary>
    /// Client capabilities specific to parameter information.
    /// </summary>
    public /*TOpt*/LabelOffsetSupportBoolean ParameterInformation;
    /// <summary>
    /// The client supports the `activeParameter` property on `SignatureInformation`
    /// literal.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.Boolean? ActiveParameterSupport;
}


/*TStruc*/
public class ValueSetDiagnosticTags
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public DiagnosticTag[] ValueSet;
}


/*TOr*/
public class InlineValueTextOrInlineValueVariableLookupOrInlineValueEvaluatableExpression
{
    public  /*TOpt*/InlineValueText? InlineValueText;
    public  /*TOpt*/InlineValueVariableLookup? InlineValueVariableLookup;
    public  /*TOpt*/InlineValueEvaluatableExpression? InlineValueEvaluatableExpression;
}

/*TOr*/
public class TextEditOrInsertReplaceEdit
{
    public  /*TOpt*/TextEdit? TextEdit;
    public  /*TOpt*/InsertReplaceEdit? InsertReplaceEdit;
}

/*TStruc*/
public class CommitCharactersStringsWithEditRangeRangeOrInsertRangeWithReplaceRangeWithInsertTextFormatWithInsertTextModeWithDataLSPAny
{
    /// <summary>
    /// A default commit character set.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/System.String[] CommitCharacters;
    /// <summary>
    /// A default edit range.
    /// 
    /// @since 3.17.0
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/RangeOrInsertRangeWithReplaceRange EditRange;
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
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Data;
}


/*TOr*/
public class WorkspaceFolderOrURI
{
    public  /*TOpt*/WorkspaceFolder? WorkspaceFolder;
    public  /*TOpt*/System.String URI;
}

/*TStruc*/
public class ValueSetInsertTextModes
{
    public InsertTextMode[] ValueSet;
}


/*TOr*/
public class BooleanOrDocumentHighlightOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/DocumentHighlightOptions? DocumentHighlightOptions;
}

/*TOr*/
public class BooleanOrCodeActionOptions
{
    public  /*TOpt*/System.Boolean? Boolean;
    public  /*TOpt*/CodeActionOptions? CodeActionOptions;
}

/*TStruc*/
public class CancelBooleanWithRetryOnContentModifiedStrings
{
    /// <summary>
    /// The client will actively cancel the request.
    /// </summary>
    public System.Boolean Cancel;
    /// <summary>
    /// The list of requests for which the client
    /// will retry the request if it receives a
    /// response with error code `ContentModified`
    /// </summary>
    public System.String[] RetryOnContentModified;
}


/*TStruc*/
public class PropertiesStrings_
{
    /// <summary>
    /// The properties that a client can resolve lazily.
    /// </summary>
    public System.String[] Properties;
}


