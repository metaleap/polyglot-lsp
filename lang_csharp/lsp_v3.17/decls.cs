/// <summary>Language Server Protocol (LSP) v3.17 SDK for C#: auto-generated via github.com/metaleap/polyglot-vsx-and-lsp/gen/cmd/gen_lsp</summary>
namespace lsp;


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
    public /*TOpt*/bool? SnippetSupport;
    /// <summary>
    /// Client supports commit characters on a completion item.
    /// </summary>
    public /*TOpt*/bool? CommitCharactersSupport;
    /// <summary>
    /// Client supports the following content formats for the documentation
    /// property. The order describes the preferred format of the client.
    /// </summary>
    public /*TOpt*//*MarkupKind*/string[] DocumentationFormat;
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
    public /*TOpt*/ValueSetCompletionItemTags TagSupport;
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
    public /*TOpt*/bool? LabelDetailsSupport;
}


/*TOr*/
public abstract record class BooleanOrDeltaBoolean
{
    private BooleanOrDeltaBoolean() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrDeltaBoolean;
    public record class WithDeltaBoolean(/*TOpt*/DeltaBoolean_ DeltaBoolean) : BooleanOrDeltaBoolean;
}

/*TOr*/
public abstract record class InlineValueTextOrInlineValueVariableLookupOrInlineValueEvaluatableExpression
{
    private InlineValueTextOrInlineValueVariableLookupOrInlineValueEvaluatableExpression() { }
    public record class WithInlineValueText(/*TOpt*/InlineValueText? InlineValueText) : InlineValueTextOrInlineValueVariableLookupOrInlineValueEvaluatableExpression;
    public record class WithInlineValueVariableLookup(/*TOpt*/InlineValueVariableLookup? InlineValueVariableLookup) : InlineValueTextOrInlineValueVariableLookupOrInlineValueEvaluatableExpression;
    public record class WithInlineValueEvaluatableExpression(/*TOpt*/InlineValueEvaluatableExpression? InlineValueEvaluatableExpression) : InlineValueTextOrInlineValueVariableLookupOrInlineValueEvaluatableExpression;
}

/*TOr*/
public abstract record class RangeOrRangeWithPlaceholderStringOrDefaultBehaviorBoolean
{
    private RangeOrRangeWithPlaceholderStringOrDefaultBehaviorBoolean() { }
    public record class WithRange(/*TOpt*/Range? Range) : RangeOrRangeWithPlaceholderStringOrDefaultBehaviorBoolean;
    public record class WithRangeWithPlaceholderString(/*TOpt*/RangeWithPlaceholderString RangeWithPlaceholderString) : RangeOrRangeWithPlaceholderStringOrDefaultBehaviorBoolean;
    public record class WithDefaultBehaviorBoolean(/*TOpt*/DefaultBehaviorBoolean DefaultBehaviorBoolean) : RangeOrRangeWithPlaceholderStringOrDefaultBehaviorBoolean;
}

/*TStruc*/
public class UriDocumentUri
{
    public string Uri;
}


/*TOr*/
public abstract record class TextEditOrAnnotatedTextEdit
{
    private TextEditOrAnnotatedTextEdit() { }
    public record class WithTextEdit(/*TOpt*/TextEdit? TextEdit) : TextEditOrAnnotatedTextEdit;
    public record class WithAnnotatedTextEdit(/*TOpt*/AnnotatedTextEdit? AnnotatedTextEdit) : TextEditOrAnnotatedTextEdit;
}

/*TOr*/
public abstract record class BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions
{
    private BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions;
    public record class WithDocumentColorOptions(/*TOpt*/DocumentColorOptions? DocumentColorOptions) : BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions;
    public record class WithDocumentColorRegistrationOptions(/*TOpt*/DocumentColorRegistrationOptions? DocumentColorRegistrationOptions) : BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions;
}

/*TOr*/
public abstract record class StringOrNotebookDocumentFilter
{
    private StringOrNotebookDocumentFilter() { }
    public record class WithString(/*TOpt*/string? String) : StringOrNotebookDocumentFilter;
    public record class WithNotebookDocumentFilter(/*TOpt*/NotebookDocumentFilter NotebookDocumentFilter) : StringOrNotebookDocumentFilter;
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
public class TextString
{
    /// <summary>
    /// The new text of the whole document.
    /// </summary>
    public string Text;
}


/*TOr*/
public abstract record class RangeWithRangeLengthUintegerWithTextStringOrTextString
{
    private RangeWithRangeLengthUintegerWithTextStringOrTextString() { }
    public record class WithRangeWithRangeLengthUintegerWithTextString(/*TOpt*/RangeWithRangeLengthUintegerWithTextString RangeWithRangeLengthUintegerWithTextString) : RangeWithRangeLengthUintegerWithTextStringOrTextString;
    public record class WithTextString(/*TOpt*/TextString TextString) : RangeWithRangeLengthUintegerWithTextStringOrTextString;
}

/*TStruc*/
public class InsertRangeWithReplaceRange
{
    public Range Insert;
    public Range Replace;
}


/*TOr*/
public abstract record class BooleanOrReferenceOptions
{
    private BooleanOrReferenceOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrReferenceOptions;
    public record class WithReferenceOptions(/*TOpt*/ReferenceOptions? ReferenceOptions) : BooleanOrReferenceOptions;
}

/*TTup*/
public class UintegerWithUinteger
{
    public uint Uinteger0;
    public uint Uinteger1;
}


/*TOr*/
public abstract record class BooleanOrDefinitionOptions
{
    private BooleanOrDefinitionOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrDefinitionOptions;
    public record class WithDefinitionOptions(/*TOpt*/DefinitionOptions? DefinitionOptions) : BooleanOrDefinitionOptions;
}

/*TOr*/
public abstract record class BooleanOrDocumentSymbolOptions
{
    private BooleanOrDocumentSymbolOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrDocumentSymbolOptions;
    public record class WithDocumentSymbolOptions(/*TOpt*/DocumentSymbolOptions? DocumentSymbolOptions) : BooleanOrDocumentSymbolOptions;
}

/*TOr*/
public abstract record class BooleanOrSaveOptions
{
    private BooleanOrSaveOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrSaveOptions;
    public record class WithSaveOptions(/*TOpt*/SaveOptions? SaveOptions) : BooleanOrSaveOptions;
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
public class GroupsOnLabelBoolean
{
    /// <summary>
    /// Whether the client groups edits with equal labels into tree nodes,
    /// for instance all edits labelled with "Changes in Strings" would
    /// be a tree node.
    /// </summary>
    public /*TOpt*/bool? GroupsOnLabel;
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
public abstract record class LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull
{
    private LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull() { }
    public record class WithLSPObject(/*TOpt*/LSPObject LSPObject) : LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull;
    public record class WithLSPArray(/*TOpt*/LSPArray LSPArray) : LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull;
    public record class WithString(/*TOpt*/string? String) : LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull;
    public record class WithInteger(/*TOpt*/int? Integer) : LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull;
    public record class WithUinteger(/*TOpt*/uint? Uinteger) : LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull;
    public record class WithDecimal(/*TOpt*/double? Decimal) : LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull;
    public record class WithBoolean(/*TOpt*/bool? Boolean) : LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull;
}

/*TStruc*/
public class RangeWithPlaceholderString
{
    public Range Range;
    public string Placeholder;
}


/*TStruc*/
public class NameStringWithVersionString
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
public abstract record class StringOrBoolean
{
    private StringOrBoolean() { }
    public record class WithString(/*TOpt*/string? String) : StringOrBoolean;
    public record class WithBoolean(/*TOpt*/bool? Boolean) : StringOrBoolean;
}

/*TOr*/
public abstract record class SemanticTokensOptionsOrSemanticTokensRegistrationOptions
{
    private SemanticTokensOptionsOrSemanticTokensRegistrationOptions() { }
    public record class WithSemanticTokensOptions(/*TOpt*/SemanticTokensOptions? SemanticTokensOptions) : SemanticTokensOptionsOrSemanticTokensRegistrationOptions;
    public record class WithSemanticTokensRegistrationOptions(/*TOpt*/SemanticTokensRegistrationOptions? SemanticTokensRegistrationOptions) : SemanticTokensOptionsOrSemanticTokensRegistrationOptions;
}

/*TStruc*/
public class ValueSetInsertTextModes
{
    public InsertTextMode[] ValueSet;
}


/*TStruc*/
public class DefaultBehaviorBoolean
{
    public bool DefaultBehavior;
}


/*TOr*/
public abstract record class FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport
{
    private FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport() { }
    public record class WithFullDocumentDiagnosticReport(/*TOpt*/FullDocumentDiagnosticReport? FullDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport;
    public record class WithUnchangedDocumentDiagnosticReport(/*TOpt*/UnchangedDocumentDiagnosticReport? UnchangedDocumentDiagnosticReport) : FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport;
}

/*TStruc*/
public class DeltaBoolean
{
    /// <summary>
    /// The server supports deltas for full documents.
    /// </summary>
    public /*TOpt*/bool? Delta;
}


/*TOr*/
public abstract record class BooleanOrRenameOptions
{
    private BooleanOrRenameOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrRenameOptions;
    public record class WithRenameOptions(/*TOpt*/RenameOptions? RenameOptions) : BooleanOrRenameOptions;
}

/*TOr*/
public abstract record class BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions
{
    private BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions;
    public record class WithTypeHierarchyOptions(/*TOpt*/TypeHierarchyOptions? TypeHierarchyOptions) : BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions;
    public record class WithTypeHierarchyRegistrationOptions(/*TOpt*/TypeHierarchyRegistrationOptions? TypeHierarchyRegistrationOptions) : BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions;
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
    public /*TOpt*/bool? LabelDetailsSupport;
}


/*TStruc*/
public class CancelBooleanWithRetryOnContentModifiedStrings
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
public abstract record class LocationOrLocations
{
    private LocationOrLocations() { }
    public record class WithLocation(/*TOpt*/Location? Location) : LocationOrLocations;
    public record class WithLocations(/*TOpt*/Location[] Locations) : LocationOrLocations;
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
    public /*TOpt*/uint? RangeLength;
    /// <summary>
    /// The new text for the provided range.
    /// </summary>
    public string Text;
}


/*TStruc*/
public class LanguageStringWithValueString
{
    public string Language;
    public string Value;
}


/*TOr*/
public abstract record class BooleanOrImplementationOptionsOrImplementationRegistrationOptions
{
    private BooleanOrImplementationOptionsOrImplementationRegistrationOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrImplementationOptionsOrImplementationRegistrationOptions;
    public record class WithImplementationOptions(/*TOpt*/ImplementationOptions? ImplementationOptions) : BooleanOrImplementationOptionsOrImplementationRegistrationOptions;
    public record class WithImplementationRegistrationOptions(/*TOpt*/ImplementationRegistrationOptions? ImplementationRegistrationOptions) : BooleanOrImplementationOptionsOrImplementationRegistrationOptions;
}

/*TStruc*/
public class LanguageString
{
    public string Language;
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
    public /*TOpt*/bool? ActiveParameterSupport;
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
public abstract record class StringOrInlayHintLabelParts
{
    private StringOrInlayHintLabelParts() { }
    public record class WithString(/*TOpt*/string? String) : StringOrInlayHintLabelParts;
    public record class WithInlayHintLabelParts(/*TOpt*/InlayHintLabelPart[] InlayHintLabelParts) : StringOrInlayHintLabelParts;
}

/*TOr*/
public abstract record class BooleanOrDocumentFormattingOptions
{
    private BooleanOrDocumentFormattingOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrDocumentFormattingOptions;
    public record class WithDocumentFormattingOptions(/*TOpt*/DocumentFormattingOptions? DocumentFormattingOptions) : BooleanOrDocumentFormattingOptions;
}

/*TOr*/
public abstract record class BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions
{
    private BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions;
    public record class WithInlineValueOptions(/*TOpt*/InlineValueOptions? InlineValueOptions) : BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions;
    public record class WithInlineValueRegistrationOptions(/*TOpt*/InlineValueRegistrationOptions? InlineValueRegistrationOptions) : BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions;
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
public class CollapsedTextBoolean
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
public class AdditionalPropertiesSupportBoolean
{
    /// <summary>
    /// Whether the client supports additional attributes which
    /// are preserved and send back to the server in the
    /// request's response.
    /// </summary>
    public /*TOpt*/bool? AdditionalPropertiesSupport;
}


/*TOr*/
public abstract record class IntegerOrString
{
    private IntegerOrString() { }
    public record class WithInteger(/*TOpt*/int? Integer) : IntegerOrString;
    public record class WithString(/*TOpt*/string? String) : IntegerOrString;
}

/*TOr*/
public abstract record class StringOrMarkupContent
{
    private StringOrMarkupContent() { }
    public record class WithString(/*TOpt*/string? String) : StringOrMarkupContent;
    public record class WithMarkupContent(/*TOpt*/MarkupContent? MarkupContent) : StringOrMarkupContent;
}

/*TStruc*/
public class DocumentVersionedTextDocumentIdentifierWithChangesTextDocumentContentChangeEvents
{
    public VersionedTextDocumentIdentifier Document;
    public TextDocumentContentChangeEvent[] Changes;
}


/*TStruc*/
public class PropertiesStrings
{
    /// <summary>
    /// The properties that a client can resolve lazily.
    /// </summary>
    public string[] Properties;
}


/*TOr*/
public abstract record class BooleanOrAnyByString
{
    private BooleanOrAnyByString() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrAnyByString;
    public record class WithAnyByString(/*TOpt*/System.Collections.Generic.Dictionary<string, object> AnyByString) : BooleanOrAnyByString;
}

/*TOr*/
public abstract record class NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions
{
    private NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions() { }
    public record class WithNotebookDocumentSyncOptions(/*TOpt*/NotebookDocumentSyncOptions? NotebookDocumentSyncOptions) : NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions;
    public record class WithNotebookDocumentSyncRegistrationOptions(/*TOpt*/NotebookDocumentSyncRegistrationOptions? NotebookDocumentSyncRegistrationOptions) : NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions;
}

/*TOr*/
public abstract record class WorkspaceFullDocumentDiagnosticReportOrWorkspaceUnchangedDocumentDiagnosticReport
{
    private WorkspaceFullDocumentDiagnosticReportOrWorkspaceUnchangedDocumentDiagnosticReport() { }
    public record class WithWorkspaceFullDocumentDiagnosticReport(/*TOpt*/WorkspaceFullDocumentDiagnosticReport? WorkspaceFullDocumentDiagnosticReport) : WorkspaceFullDocumentDiagnosticReportOrWorkspaceUnchangedDocumentDiagnosticReport;
    public record class WithWorkspaceUnchangedDocumentDiagnosticReport(/*TOpt*/WorkspaceUnchangedDocumentDiagnosticReport? WorkspaceUnchangedDocumentDiagnosticReport) : WorkspaceFullDocumentDiagnosticReportOrWorkspaceUnchangedDocumentDiagnosticReport;
}

/*TStruc*/
public class NameStringWithVersionString
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


/*TOr*/
public abstract record class StringOrStrings
{
    private StringOrStrings() { }
    public record class WithString(/*TOpt*/string? String) : StringOrStrings;
    public record class WithStrings(/*TOpt*/string[] Strings) : StringOrStrings;
}

/*TOr*/
public abstract record class MarkupContentOrMarkedStringOrMarkedStrings
{
    private MarkupContentOrMarkedStringOrMarkedStrings() { }
    public record class WithMarkupContent(/*TOpt*/MarkupContent? MarkupContent) : MarkupContentOrMarkedStringOrMarkedStrings;
    public record class WithMarkedString(/*TOpt*/MarkedString MarkedString) : MarkupContentOrMarkedStringOrMarkedStrings;
    public record class WithMarkedStrings(/*TOpt*/MarkedString[] MarkedStrings) : MarkupContentOrMarkedStringOrMarkedStrings;
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
public abstract record class RelatedFullDocumentDiagnosticReportOrRelatedUnchangedDocumentDiagnosticReport
{
    private RelatedFullDocumentDiagnosticReportOrRelatedUnchangedDocumentDiagnosticReport() { }
    public record class WithRelatedFullDocumentDiagnosticReport(/*TOpt*/RelatedFullDocumentDiagnosticReport? RelatedFullDocumentDiagnosticReport) : RelatedFullDocumentDiagnosticReportOrRelatedUnchangedDocumentDiagnosticReport;
    public record class WithRelatedUnchangedDocumentDiagnosticReport(/*TOpt*/RelatedUnchangedDocumentDiagnosticReport? RelatedUnchangedDocumentDiagnosticReport) : RelatedFullDocumentDiagnosticReportOrRelatedUnchangedDocumentDiagnosticReport;
}

/*TOr*/
public abstract record class PatternOrRelativePattern
{
    private PatternOrRelativePattern() { }
    public record class WithPattern(/*TOpt*/string? Pattern) : PatternOrRelativePattern;
    public record class WithRelativePattern(/*TOpt*/RelativePattern? RelativePattern) : PatternOrRelativePattern;
}

/*TStruc*/
public class LanguageStringWithSchemeStringWithPatternString
{
    /// <summary>
    /// A language id, like `typescript`.
    /// </summary>
    public /*TOpt*/string? Language;
    /// <summary>
    /// A Uri <c>Uri.scheme</c>, like `file` or `untitled`.
    /// </summary>
    public /*TOpt*/string? Scheme;
    /// <summary>
    /// A glob pattern, like `*.{ts,js}`.
    /// </summary>
    public /*TOpt*/string? Pattern;
}


/*TOr*/
public abstract record class DiagnosticOptionsOrDiagnosticRegistrationOptions
{
    private DiagnosticOptionsOrDiagnosticRegistrationOptions() { }
    public record class WithDiagnosticOptions(/*TOpt*/DiagnosticOptions? DiagnosticOptions) : DiagnosticOptionsOrDiagnosticRegistrationOptions;
    public record class WithDiagnosticRegistrationOptions(/*TOpt*/DiagnosticRegistrationOptions? DiagnosticRegistrationOptions) : DiagnosticOptionsOrDiagnosticRegistrationOptions;
}

/*TOr*/
public abstract record class BooleanOrDeltaBoolean
{
    private BooleanOrDeltaBoolean() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrDeltaBoolean;
    public record class WithDeltaBoolean(/*TOpt*/DeltaBoolean DeltaBoolean) : BooleanOrDeltaBoolean;
}

/*TOr*/
public abstract record class BooleanOrCodeActionOptions
{
    private BooleanOrCodeActionOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrCodeActionOptions;
    public record class WithCodeActionOptions(/*TOpt*/CodeActionOptions? CodeActionOptions) : BooleanOrCodeActionOptions;
}

/*TOr*/
public abstract record class BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions
{
    private BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions;
    public record class WithLinkedEditingRangeOptions(/*TOpt*/LinkedEditingRangeOptions? LinkedEditingRangeOptions) : BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions;
    public record class WithLinkedEditingRangeRegistrationOptions(/*TOpt*/LinkedEditingRangeRegistrationOptions? LinkedEditingRangeRegistrationOptions) : BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions;
}

/*TStruc*/
public class DeltaBoolean
{
    /// <summary>
    /// The client will send the `textDocument/semanticTokens/full/delta` request if
    /// the server provides a corresponding handler.
    /// </summary>
    public /*TOpt*/bool? Delta;
}


/*TOr*/
public abstract record class BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions
{
    private BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions;
    public record class WithSelectionRangeOptions(/*TOpt*/SelectionRangeOptions? SelectionRangeOptions) : BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions;
    public record class WithSelectionRangeRegistrationOptions(/*TOpt*/SelectionRangeRegistrationOptions? SelectionRangeRegistrationOptions) : BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions;
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


/*TStruc*/
public class ValueSetSymbolTags
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public SymbolTag[] ValueSet;
}


/*TStruc*/
public class NotebookTypeStringWithSchemeStringWithPatternString
{
    /// <summary>
    /// The type of the enclosing notebook.
    /// </summary>
    public /*TOpt*/string? NotebookType;
    /// <summary>
    /// A Uri <c>Uri.scheme</c>, like `file` or `untitled`.
    /// </summary>
    public /*TOpt*/string? Scheme;
    /// <summary>
    /// A glob pattern.
    /// </summary>
    public /*TOpt*/string? Pattern;
}


/*TOr*/
public abstract record class TextEditOrInsertReplaceEdit
{
    private TextEditOrInsertReplaceEdit() { }
    public record class WithTextEdit(/*TOpt*/TextEdit? TextEdit) : TextEditOrInsertReplaceEdit;
    public record class WithInsertReplaceEdit(/*TOpt*/InsertReplaceEdit? InsertReplaceEdit) : TextEditOrInsertReplaceEdit;
}

/*TStruc*/
public class ReasonString
{
    /// <summary>
    /// Human readable description of why the code action is currently disabled.
    /// 
    /// This is displayed in the code actions UI.
    /// </summary>
    public string Reason;
}


/*TOr*/
public abstract record class BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions
{
    private BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions;
    public record class WithDeclarationOptions(/*TOpt*/DeclarationOptions? DeclarationOptions) : BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions;
    public record class WithDeclarationRegistrationOptions(/*TOpt*/DeclarationRegistrationOptions? DeclarationRegistrationOptions) : BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions;
}

/*TOr*/
public abstract record class WorkspaceFolderOrURI
{
    private WorkspaceFolderOrURI() { }
    public record class WithWorkspaceFolder(/*TOpt*/WorkspaceFolder? WorkspaceFolder) : WorkspaceFolderOrURI;
    public record class WithURI(/*TOpt*/string URI) : WorkspaceFolderOrURI;
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
    public /*TOpt*/bool? LabelOffsetSupport;
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


/*TOr*/
public abstract record class StringOrLanguageStringWithValueString
{
    private StringOrLanguageStringWithValueString() { }
    public record class WithString(/*TOpt*/string? String) : StringOrLanguageStringWithValueString;
    public record class WithLanguageStringWithValueString(/*TOpt*/LanguageStringWithValueString LanguageStringWithValueString) : StringOrLanguageStringWithValueString;
}

/*TOr*/
public abstract record class TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile
{
    private TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile() { }
    public record class WithTextDocumentEdit(/*TOpt*/TextDocumentEdit? TextDocumentEdit) : TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile;
    public record class WithCreateFile(/*TOpt*/CreateFile? CreateFile) : TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile;
    public record class WithRenameFile(/*TOpt*/RenameFile? RenameFile) : TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile;
    public record class WithDeleteFile(/*TOpt*/DeleteFile? DeleteFile) : TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile;
}

/*TOr*/
public abstract record class RangeOrInsertRangeWithReplaceRange
{
    private RangeOrInsertRangeWithReplaceRange() { }
    public record class WithRange(/*TOpt*/Range? Range) : RangeOrInsertRangeWithReplaceRange;
    public record class WithInsertRangeWithReplaceRange(/*TOpt*/InsertRangeWithReplaceRange InsertRangeWithReplaceRange) : RangeOrInsertRangeWithReplaceRange;
}

/*TOr*/
public abstract record class BooleanOrWorkspaceSymbolOptions
{
    private BooleanOrWorkspaceSymbolOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrWorkspaceSymbolOptions;
    public record class WithWorkspaceSymbolOptions(/*TOpt*/WorkspaceSymbolOptions? WorkspaceSymbolOptions) : BooleanOrWorkspaceSymbolOptions;
}

/*TStruc*/
public class PropertiesStrings
{
    /// <summary>
    /// The properties that a client can resolve lazily. Usually
    /// `location.range`
    /// </summary>
    public string[] Properties;
}


/*TStruc*/
public class ValueSetDiagnosticTags
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public DiagnosticTag[] ValueSet;
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
public abstract record class BooleanOrHoverOptions
{
    private BooleanOrHoverOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrHoverOptions;
    public record class WithHoverOptions(/*TOpt*/HoverOptions? HoverOptions) : BooleanOrHoverOptions;
}

/*TOr*/
public abstract record class BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions
{
    private BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions;
    public record class WithTypeDefinitionOptions(/*TOpt*/TypeDefinitionOptions? TypeDefinitionOptions) : BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions;
    public record class WithTypeDefinitionRegistrationOptions(/*TOpt*/TypeDefinitionRegistrationOptions? TypeDefinitionRegistrationOptions) : BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions;
}

/*TOr*/
public abstract record class BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions
{
    private BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions;
    public record class WithInlayHintOptions(/*TOpt*/InlayHintOptions? InlayHintOptions) : BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions;
    public record class WithInlayHintRegistrationOptions(/*TOpt*/InlayHintRegistrationOptions? InlayHintRegistrationOptions) : BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions;
}

/*TOr*/
public abstract record class BooleanOrMonikerOptionsOrMonikerRegistrationOptions
{
    private BooleanOrMonikerOptionsOrMonikerRegistrationOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrMonikerOptionsOrMonikerRegistrationOptions;
    public record class WithMonikerOptions(/*TOpt*/MonikerOptions? MonikerOptions) : BooleanOrMonikerOptionsOrMonikerRegistrationOptions;
    public record class WithMonikerRegistrationOptions(/*TOpt*/MonikerRegistrationOptions? MonikerRegistrationOptions) : BooleanOrMonikerOptionsOrMonikerRegistrationOptions;
}

/*TOr*/
public abstract record class StringOrUintegerWithUinteger
{
    private StringOrUintegerWithUinteger() { }
    public record class WithString(/*TOpt*/string? String) : StringOrUintegerWithUinteger;
    public record class WithUintegerWithUinteger(/*TOpt*/UintegerWithUinteger? UintegerWithUinteger) : StringOrUintegerWithUinteger;
}

/*TOr*/
public abstract record class TextDocumentFilterOrNotebookCellTextDocumentFilter
{
    private TextDocumentFilterOrNotebookCellTextDocumentFilter() { }
    public record class WithTextDocumentFilter(/*TOpt*/TextDocumentFilter TextDocumentFilter) : TextDocumentFilterOrNotebookCellTextDocumentFilter;
    public record class WithNotebookCellTextDocumentFilter(/*TOpt*/NotebookCellTextDocumentFilter? NotebookCellTextDocumentFilter) : TextDocumentFilterOrNotebookCellTextDocumentFilter;
}

/*TStruc*/
public class CommitCharactersStringsWithEditRangeRangeOrInsertRangeWithReplaceRangeWithInsertTextFormatWithInsertTextModeWithDataLSPAny
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
    public /*TOpt*/LSPAny Data;
}


/*TOr*/
public abstract record class LocationOrUriDocumentUri
{
    private LocationOrUriDocumentUri() { }
    public record class WithLocation(/*TOpt*/Location? Location) : LocationOrUriDocumentUri;
    public record class WithUriDocumentUri(/*TOpt*/UriDocumentUri UriDocumentUri) : LocationOrUriDocumentUri;
}

/*TOr*/
public abstract record class BooleanOrDocumentRangeFormattingOptions
{
    private BooleanOrDocumentRangeFormattingOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrDocumentRangeFormattingOptions;
    public record class WithDocumentRangeFormattingOptions(/*TOpt*/DocumentRangeFormattingOptions? DocumentRangeFormattingOptions) : BooleanOrDocumentRangeFormattingOptions;
}

/*TOr*/
public abstract record class BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions
{
    private BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions;
    public record class WithCallHierarchyOptions(/*TOpt*/CallHierarchyOptions? CallHierarchyOptions) : BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions;
    public record class WithCallHierarchyRegistrationOptions(/*TOpt*/CallHierarchyRegistrationOptions? CallHierarchyRegistrationOptions) : BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions;
}

/*TStruc*/
public class ValueSetCompletionItemTags
{
    /// <summary>
    /// The tags supported by the client.
    /// </summary>
    public CompletionItemTag[] ValueSet;
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
    public /*TOpt*/string[] ItemDefaults;
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
public abstract record class TextDocumentSyncOptionsOrTextDocumentSyncKind
{
    private TextDocumentSyncOptionsOrTextDocumentSyncKind() { }
    public record class WithTextDocumentSyncOptions(/*TOpt*/TextDocumentSyncOptions? TextDocumentSyncOptions) : TextDocumentSyncOptionsOrTextDocumentSyncKind;
    public record class WithTextDocumentSyncKind(/*TOpt*/TextDocumentSyncKind TextDocumentSyncKind) : TextDocumentSyncOptionsOrTextDocumentSyncKind;
}

/*TOr*/
public abstract record class BooleanOrDocumentHighlightOptions
{
    private BooleanOrDocumentHighlightOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrDocumentHighlightOptions;
    public record class WithDocumentHighlightOptions(/*TOpt*/DocumentHighlightOptions? DocumentHighlightOptions) : BooleanOrDocumentHighlightOptions;
}

/*TOr*/
public abstract record class BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions
{
    private BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions() { }
    public record class WithBoolean(/*TOpt*/bool? Boolean) : BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions;
    public record class WithFoldingRangeOptions(/*TOpt*/FoldingRangeOptions? FoldingRangeOptions) : BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions;
    public record class WithFoldingRangeRegistrationOptions(/*TOpt*/FoldingRangeRegistrationOptions? FoldingRangeRegistrationOptions) : BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions;
}

