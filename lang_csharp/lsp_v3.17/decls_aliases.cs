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



// using Definition = LocationOrLocations;

// using DefinitionLink = LocationLink;

// using LSPArray = LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull[];

// using LSPAny = LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull;

// using Declaration = LocationOrLocations;

// using DeclarationLink = LocationLink;

// using InlineValue = InlineValueTextOrInlineValueVariableLookupOrInlineValueEvaluatableExpression;

// using DocumentDiagnosticReport = RelatedFullDocumentDiagnosticReportOrRelatedUnchangedDocumentDiagnosticReport;

// using PrepareRenameResult = RangeOrRangeWithPlaceholderStringOrDefaultBehaviorBoolean;

// using DocumentSelector = DocumentFilter[];

// using ProgressToken = IntegerOrString;

// using ChangeAnnotationIdentifier = System.String;

// using WorkspaceDocumentDiagnosticReport = WorkspaceFullDocumentDiagnosticReportOrWorkspaceUnchangedDocumentDiagnosticReport;

// using TextDocumentContentChangeEvent = RangeWithRangeLengthUintegerWithTextStringOrTextString;

// using MarkedString = StringOrLanguageStringWithValueString;

// using DocumentFilter = TextDocumentFilterOrNotebookCellTextDocumentFilter;

// using LSPObject = System.Collections.Generic.Dictionary<System.String, LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull>;

// using GlobPattern = PatternOrRelativePattern;

// using TextDocumentFilter = LanguageStringWithSchemeStringWithPatternString;

// using NotebookDocumentFilter = NotebookTypeStringWithSchemeStringWithPatternString;

// using Pattern = System.String;

