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



public class ImplementationParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
}

/// <summary>
/// Represents a location inside a resource, such as a line
/// inside a text file.
/// </summary>
public class Location
{
    public System.String Uri;
    public Range Range;
}

public class ImplementationRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public ImplementationOptions ImplementationOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

public class TypeDefinitionParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
}

public class TypeDefinitionRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public TypeDefinitionOptions TypeDefinitionOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

/// <summary>
/// A workspace folder inside a client.
/// </summary>
public class WorkspaceFolder
{
    /// <summary>
    /// The associated URI for this workspace folder.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// The name of the workspace folder. Used to refer to this
    /// workspace folder in the user interface.
    /// </summary>
    public System.String Name;
}

/// <summary>
/// The parameters of a `workspace/didChangeWorkspaceFolders` notification.
/// </summary>
public class DidChangeWorkspaceFoldersParams
{
    /// <summary>
    /// The actual workspace folder change event.
    /// </summary>
    public WorkspaceFoldersChangeEvent Event;
}

/// <summary>
/// The parameters of a configuration request.
/// </summary>
public class ConfigurationParams
{
    public ConfigurationItem[] Items;
}

/// <summary>
/// Parameters for a <c>DocumentColorRequest</c>.
/// </summary>
public class DocumentColorParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The text document.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
}

/// <summary>
/// Represents a color range from a document.
/// </summary>
public class ColorInformation
{
    /// <summary>
    /// The range in the document where this color appears.
    /// </summary>
    public Range Range;
    /// <summary>
    /// The actual color value for this color range.
    /// </summary>
    public Color Color;
}

public class DocumentColorRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public DocumentColorOptions DocumentColorOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

/// <summary>
/// Parameters for a <c>ColorPresentationRequest</c>.
/// </summary>
public class ColorPresentationParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The text document.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The color to request presentations for.
    /// </summary>
    public Color Color;
    /// <summary>
    /// The range where the color would be inserted. Serves as a context.
    /// </summary>
    public Range Range;
}

public class ColorPresentation
{
    /// <summary>
    /// The label of this color presentation. It will be shown on the color
    /// picker header. By default this is also the text that is inserted when selecting
    /// this color presentation.
    /// </summary>
    public System.String Label;
    /// <summary>
    /// An <c>TextEdit</c> which is applied to a document when selecting
    /// this presentation for the color.  When `falsy` the <c>ColorPresentation.label</c>
    /// is used.
    /// </summary>
    public /*TOpt*/TextEdit? TextEdit;
    /// <summary>
    /// An optional array of additional <c>TextEdit</c> that are applied when
    /// selecting this color presentation. Edits must not overlap with the main <c>ColorPresentation.textEdit</c> nor with themselves.
    /// </summary>
    public /*TOpt*/TextEdit[] AdditionalTextEdits;
}

public class WorkDoneProgressOptions
{
    public /*TOpt*/System.Boolean? WorkDoneProgress;
}

/// <summary>
/// General text document registration options.
/// </summary>
public class TextDocumentRegistrationOptions
{
    /// <summary>
    /// A document selector to identify the scope of the registration. If set to null
    /// the document selector provided on the client side will be used.
    /// </summary>
    public /*TOr1*//*TOpt*/DocumentFilter[] DocumentSelector;
}

/// <summary>
/// Parameters for a <c>FoldingRangeRequest</c>.
/// </summary>
public class FoldingRangeParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The text document.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
}

/// <summary>
/// Represents a folding range. To be valid, start and end line must be bigger than zero and smaller
/// than the number of lines in the document. Clients are free to ignore invalid ranges.
/// </summary>
public class FoldingRange
{
    /// <summary>
    /// The zero-based start line of the range to fold. The folded area starts after the line's last character.
    /// To be valid, the end must be zero or larger and smaller than the number of lines in the document.
    /// </summary>
    public System.UInt64 StartLine;
    /// <summary>
    /// The zero-based character offset from where the folded range starts. If not defined, defaults to the length of the start line.
    /// </summary>
    public /*TOpt*/System.UInt64? StartCharacter;
    /// <summary>
    /// The zero-based end line of the range to fold. The folded area ends with the line's last character.
    /// To be valid, the end must be zero or larger and smaller than the number of lines in the document.
    /// </summary>
    public System.UInt64 EndLine;
    /// <summary>
    /// The zero-based character offset before the folded range ends. If not defined, defaults to the length of the end line.
    /// </summary>
    public /*TOpt*/System.UInt64? EndCharacter;
    /// <summary>
    /// Describes the kind of the folding range such as `comment' or 'region'. The kind
    /// is used to categorize folding ranges and used by commands like 'Fold all comments'.
    /// See <c>FoldingRangeKind</c> for an enumeration of standardized kinds.
    /// </summary>
    public /*TOpt*//*FoldingRangeKind*/string Kind;
    /// <summary>
    /// The text that the client should show when the specified range is
    /// collapsed. If not defined or not supported by the client, a default
    /// will be chosen by the client.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/System.String? CollapsedText;
}

public class FoldingRangeRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public FoldingRangeOptions FoldingRangeOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

public class DeclarationParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
}

public class DeclarationRegistrationOptions
{
    public DeclarationOptions DeclarationOptions; //XTEND
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

/// <summary>
/// A parameter literal used in selection range requests.
/// </summary>
public class SelectionRangeParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The text document.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The positions inside the text document.
    /// </summary>
    public Position[] Positions;
}

/// <summary>
/// A selection range represents a part of a selection hierarchy. A selection range
/// may have a parent selection range that contains it.
/// </summary>
public class SelectionRange
{
    /// <summary>
    /// The <c>Range</c> of this selection range.
    /// </summary>
    public Range Range;
    /// <summary>
    /// The parent selection range containing this range. Therefore `parent.range` must contain `this.range`.
    /// </summary>
    public /*TOpt*/SelectionRange? Parent;
}

public class SelectionRangeRegistrationOptions
{
    public SelectionRangeOptions SelectionRangeOptions; //XTEND
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

public class WorkDoneProgressCreateParams
{
    /// <summary>
    /// The token to be used to report progress.
    /// </summary>
    public ProgressToken Token;
}

public class WorkDoneProgressCancelParams
{
    /// <summary>
    /// The token to be used to report progress.
    /// </summary>
    public ProgressToken Token;
}

/// <summary>
/// The parameter of a `textDocument/prepareCallHierarchy` request.
/// 
/// @since 3.16.0
/// </summary>
public class CallHierarchyPrepareParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
}

/// <summary>
/// Represents programming constructs like functions or constructors in the context
/// of call hierarchy.
/// 
/// @since 3.16.0
/// </summary>
public class CallHierarchyItem
{
    /// <summary>
    /// The name of this item.
    /// </summary>
    public System.String Name;
    /// <summary>
    /// The kind of this item.
    /// </summary>
    public SymbolKind Kind;
    /// <summary>
    /// Tags for this item.
    /// </summary>
    public /*TOpt*/SymbolTag[] Tags;
    /// <summary>
    /// More detail for this item, e.g. the signature of a function.
    /// </summary>
    public /*TOpt*/System.String? Detail;
    /// <summary>
    /// The resource identifier of this item.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// The range enclosing this symbol not including leading/trailing whitespace but everything else, e.g. comments and code.
    /// </summary>
    public Range Range;
    /// <summary>
    /// The range that should be selected and revealed when this symbol is being picked, e.g. the name of a function.
    /// Must be contained by the <c>CallHierarchyItem.range</c>.
    /// </summary>
    public Range SelectionRange;
    /// <summary>
    /// A data entry field that is preserved between a call hierarchy prepare and
    /// incoming calls or outgoing calls requests.
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Data;
}

/// <summary>
/// Call hierarchy options used during static or dynamic registration.
/// 
/// @since 3.16.0
/// </summary>
public class CallHierarchyRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public CallHierarchyOptions CallHierarchyOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

/// <summary>
/// The parameter of a `callHierarchy/incomingCalls` request.
/// 
/// @since 3.16.0
/// </summary>
public class CallHierarchyIncomingCallsParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    public CallHierarchyItem Item;
}

/// <summary>
/// Represents an incoming call, e.g. a caller of a method or constructor.
/// 
/// @since 3.16.0
/// </summary>
public class CallHierarchyIncomingCall
{
    /// <summary>
    /// The item that makes the call.
    /// </summary>
    public CallHierarchyItem From;
    /// <summary>
    /// The ranges at which the calls appear. This is relative to the caller
    /// denoted by <c>CallHierarchyIncomingCall.from</c>.
    /// </summary>
    public Range[] FromRanges;
}

/// <summary>
/// The parameter of a `callHierarchy/outgoingCalls` request.
/// 
/// @since 3.16.0
/// </summary>
public class CallHierarchyOutgoingCallsParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    public CallHierarchyItem Item;
}

/// <summary>
/// Represents an outgoing call, e.g. calling a getter from a method or a method from a constructor etc.
/// 
/// @since 3.16.0
/// </summary>
public class CallHierarchyOutgoingCall
{
    /// <summary>
    /// The item that is called.
    /// </summary>
    public CallHierarchyItem To;
    /// <summary>
    /// The range at which this item is called. This is the range relative to the caller, e.g the item
    /// passed to <c>CallHierarchyItemProvider.provideCallHierarchyOutgoingCalls</c>
    /// and not <c>CallHierarchyOutgoingCall.to</c>.
    /// </summary>
    public Range[] FromRanges;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokensParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The text document.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokens
{
    /// <summary>
    /// An optional result id. If provided and clients support delta updating
    /// the client will include the result id in the next semantic token request.
    /// A server can then instead of computing all semantic tokens again simply
    /// send a delta.
    /// </summary>
    public /*TOpt*/System.String? ResultId;
    /// <summary>
    /// The actual tokens.
    /// </summary>
    public System.UInt64[] Data;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokensPartialResult
{
    public System.UInt64[] Data;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokensRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public SemanticTokensOptions SemanticTokensOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokensDeltaParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The text document.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The result id of a previous response. The result Id can either point to a full response
    /// or a delta response depending on what was received last.
    /// </summary>
    public System.String PreviousResultId;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokensDelta
{
    public /*TOpt*/System.String? ResultId;
    /// <summary>
    /// The semantic token edits to transform a previous result into a new result.
    /// </summary>
    public SemanticTokensEdit[] Edits;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokensDeltaPartialResult
{
    public SemanticTokensEdit[] Edits;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokensRangeParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The text document.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The range the semantic tokens are requested for.
    /// </summary>
    public Range Range;
}

/// <summary>
/// Params to show a resource in the UI.
/// 
/// @since 3.16.0
/// </summary>
public class ShowDocumentParams
{
    /// <summary>
    /// The uri to show.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// Indicates to show the resource in an external program.
    /// To show, for example, `https://code.visualstudio.com/`
    /// in the default WEB browser set `external` to `true`.
    /// </summary>
    public /*TOpt*/System.Boolean? External;
    /// <summary>
    /// An optional property to indicate whether the editor
    /// showing the document should take focus or not.
    /// Clients might ignore this property if an external
    /// program is started.
    /// </summary>
    public /*TOpt*/System.Boolean? TakeFocus;
    /// <summary>
    /// An optional selection range if the document is a text
    /// document. Clients might ignore the property if an
    /// external program is started or the file is not a text
    /// file.
    /// </summary>
    public /*TOpt*/Range? Selection;
}

/// <summary>
/// The result of a showDocument request.
/// 
/// @since 3.16.0
/// </summary>
public class ShowDocumentResult
{
    /// <summary>
    /// A boolean indicating if the show was successful.
    /// </summary>
    public System.Boolean Success;
}

public class LinkedEditingRangeParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
}

/// <summary>
/// The result of a linked editing range request.
/// 
/// @since 3.16.0
/// </summary>
public class LinkedEditingRanges
{
    /// <summary>
    /// A list of ranges that can be edited together. The ranges must have
    /// identical length and contain identical text content. The ranges cannot overlap.
    /// </summary>
    public Range[] Ranges;
    /// <summary>
    /// An optional word pattern (regular expression) that describes valid contents for
    /// the given ranges. If no pattern is provided, the client configuration's word
    /// pattern will be used.
    /// </summary>
    public /*TOpt*/System.String? WordPattern;
}

public class LinkedEditingRangeRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public LinkedEditingRangeOptions LinkedEditingRangeOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

/// <summary>
/// The parameters sent in notifications/requests for user-initiated creation of
/// files.
/// 
/// @since 3.16.0
/// </summary>
public class CreateFilesParams
{
    /// <summary>
    /// An array of all files/folders created in this operation.
    /// </summary>
    public FileCreate[] Files;
}

/// <summary>
/// A workspace edit represents changes to many resources managed in the workspace. The edit
/// should either provide `changes` or `documentChanges`. If documentChanges are present
/// they are preferred over `changes` if the client can handle versioned document edits.
/// 
/// Since version 3.13.0 a workspace edit can contain resource operations as well. If resource
/// operations are present clients need to execute the operations in the order in which they
/// are provided. So a workspace edit for example can consist of the following two changes:
/// (1) a create file a.txt and (2) a text document edit which insert text into file a.txt.
/// 
/// An invalid sequence (e.g. (1) delete file a.txt and (2) insert text into file a.txt) will
/// cause failure of the operation. How the client recovers from the failure is described by
/// the client capability: `workspace.workspaceEdit.failureHandling`
/// </summary>
public class WorkspaceEdit
{
    /// <summary>
    /// Holds changes to existing resources.
    /// </summary>
    public /*TOpt*/System.Collections.Generic.Dictionary<System.String, TextEdit[]> Changes;
    /// <summary>
    /// Depending on the client capability `workspace.workspaceEdit.resourceOperations` document changes
    /// are either an array of `TextDocumentEdit`s to express changes to n different text documents
    /// where each text document edit addresses a specific version of a text document. Or it can contain
    /// above `TextDocumentEdit`s mixed with create, rename and delete file / folder operations.
    /// 
    /// Whether a client supports versioned document edits is expressed via
    /// `workspace.workspaceEdit.documentChanges` client capability.
    /// 
    /// If a client neither supports `documentChanges` nor `workspace.workspaceEdit.resourceOperations` then
    /// only plain `TextEdit`s using the `changes` property are supported.
    /// 
    /// Every object in the array has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/TextDocumentEditOrCreateFileOrRenameFileOrDeleteFile[] DocumentChanges;
    /// <summary>
    /// A map of change annotations that can be referenced in `AnnotatedTextEdit`s or create, rename and
    /// delete file / folder operations.
    /// 
    /// Whether clients honor this property depends on the client capability `workspace.changeAnnotationSupport`.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.Collections.Generic.Dictionary<ChangeAnnotationIdentifier, ChangeAnnotation> ChangeAnnotations;
}

/// <summary>
/// The options to register for file operations.
/// 
/// @since 3.16.0
/// </summary>
public class FileOperationRegistrationOptions
{
    /// <summary>
    /// The actual filters.
    /// </summary>
    public FileOperationFilter[] Filters;
}

/// <summary>
/// The parameters sent in notifications/requests for user-initiated renames of
/// files.
/// 
/// @since 3.16.0
/// </summary>
public class RenameFilesParams
{
    /// <summary>
    /// An array of all files/folders renamed in this operation. When a folder is renamed, only
    /// the folder will be included, and not its children.
    /// </summary>
    public FileRename[] Files;
}

/// <summary>
/// The parameters sent in notifications/requests for user-initiated deletes of
/// files.
/// 
/// @since 3.16.0
/// </summary>
public class DeleteFilesParams
{
    /// <summary>
    /// An array of all files/folders deleted in this operation.
    /// </summary>
    public FileDelete[] Files;
}

public class MonikerParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
}

/// <summary>
/// Moniker definition to match LSIF 0.5 moniker definition.
/// 
/// @since 3.16.0
/// </summary>
public class Moniker
{
    /// <summary>
    /// The scheme of the moniker. For example tsc or .Net
    /// </summary>
    public System.String Scheme;
    /// <summary>
    /// The identifier of the moniker. The value is opaque in LSIF however
    /// schema owners are allowed to define the structure if they want.
    /// </summary>
    public System.String Identifier;
    /// <summary>
    /// The scope in which the moniker is unique
    /// </summary>
    public /*UniquenessLevel*/string Unique;
    /// <summary>
    /// The moniker kind if known.
    /// </summary>
    public /*TOpt*//*MonikerKind*/string Kind;
}

public class MonikerRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public MonikerOptions MonikerOptions; //XTEND
}

/// <summary>
/// The parameter of a `textDocument/prepareTypeHierarchy` request.
/// 
/// @since 3.17.0
/// </summary>
public class TypeHierarchyPrepareParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
}

/// <summary>
/// @since 3.17.0
/// </summary>
public class TypeHierarchyItem
{
    /// <summary>
    /// The name of this item.
    /// </summary>
    public System.String Name;
    /// <summary>
    /// The kind of this item.
    /// </summary>
    public SymbolKind Kind;
    /// <summary>
    /// Tags for this item.
    /// </summary>
    public /*TOpt*/SymbolTag[] Tags;
    /// <summary>
    /// More detail for this item, e.g. the signature of a function.
    /// </summary>
    public /*TOpt*/System.String? Detail;
    /// <summary>
    /// The resource identifier of this item.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// The range enclosing this symbol not including leading/trailing whitespace
    /// but everything else, e.g. comments and code.
    /// </summary>
    public Range Range;
    /// <summary>
    /// The range that should be selected and revealed when this symbol is being
    /// picked, e.g. the name of a function. Must be contained by the
    /// <c>TypeHierarchyItem.range</c>.
    /// </summary>
    public Range SelectionRange;
    /// <summary>
    /// A data entry field that is preserved between a type hierarchy prepare and
    /// supertypes or subtypes requests. It could also be used to identify the
    /// type hierarchy in the server, helping improve the performance on
    /// resolving supertypes and subtypes.
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Data;
}

/// <summary>
/// Type hierarchy options used during static or dynamic registration.
/// 
/// @since 3.17.0
/// </summary>
public class TypeHierarchyRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public TypeHierarchyOptions TypeHierarchyOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

/// <summary>
/// The parameter of a `typeHierarchy/supertypes` request.
/// 
/// @since 3.17.0
/// </summary>
public class TypeHierarchySupertypesParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    public TypeHierarchyItem Item;
}

/// <summary>
/// The parameter of a `typeHierarchy/subtypes` request.
/// 
/// @since 3.17.0
/// </summary>
public class TypeHierarchySubtypesParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    public TypeHierarchyItem Item;
}

/// <summary>
/// A parameter literal used in inline value requests.
/// 
/// @since 3.17.0
/// </summary>
public class InlineValueParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    /// <summary>
    /// The text document.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The document range for which inline values should be computed.
    /// </summary>
    public Range Range;
    /// <summary>
    /// Additional information about the context in which inline values were
    /// requested.
    /// </summary>
    public InlineValueContext Context;
}

/// <summary>
/// Inline value options used during static or dynamic registration.
/// 
/// @since 3.17.0
/// </summary>
public class InlineValueRegistrationOptions
{
    public InlineValueOptions InlineValueOptions; //XTEND
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

/// <summary>
/// A parameter literal used in inlay hint requests.
/// 
/// @since 3.17.0
/// </summary>
public class InlayHintParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    /// <summary>
    /// The text document.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The document range for which inlay hints should be computed.
    /// </summary>
    public Range Range;
}

/// <summary>
/// Inlay hint information.
/// 
/// @since 3.17.0
/// </summary>
public class InlayHint
{
    /// <summary>
    /// The position of this hint.
    /// 
    /// If multiple hints have the same position, they will be shown in the order
    /// they appear in the response.
    /// </summary>
    public Position Position;
    /// <summary>
    /// The label of this hint. A human readable string or an array of
    /// InlayHintLabelPart label parts.
    /// 
    /// *Note* that neither the string nor the label part can be empty.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public StringOrInlayHintLabelParts Label;
    /// <summary>
    /// The kind of this hint. Can be omitted in which case the client
    /// should fall back to a reasonable default.
    /// </summary>
    public /*TOpt*/InlayHintKind Kind;
    /// <summary>
    /// Optional text edits that are performed when accepting this inlay hint.
    /// 
    /// *Note* that edits are expected to change the document so that the inlay
    /// hint (or its nearest variant) is now part of the document and the inlay
    /// hint itself is now obsolete.
    /// </summary>
    public /*TOpt*/TextEdit[] TextEdits;
    /// <summary>
    /// The tooltip text when you hover over this item.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/StringOrMarkupContent Tooltip;
    /// <summary>
    /// Render padding before the hint.
    /// 
    /// Note: Padding should use the editor's background color, not the
    /// background color of the hint itself. That means padding can be used
    /// to visually align/separate an inlay hint.
    /// </summary>
    public /*TOpt*/System.Boolean? PaddingLeft;
    /// <summary>
    /// Render padding after the hint.
    /// 
    /// Note: Padding should use the editor's background color, not the
    /// background color of the hint itself. That means padding can be used
    /// to visually align/separate an inlay hint.
    /// </summary>
    public /*TOpt*/System.Boolean? PaddingRight;
    /// <summary>
    /// A data entry field that is preserved on an inlay hint between
    /// a `textDocument/inlayHint` and a `inlayHint/resolve` request.
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Data;
}

/// <summary>
/// Inlay hint options used during static or dynamic registration.
/// 
/// @since 3.17.0
/// </summary>
public class InlayHintRegistrationOptions
{
    public InlayHintOptions InlayHintOptions; //XTEND
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

/// <summary>
/// Parameters of the document diagnostic request.
/// 
/// @since 3.17.0
/// </summary>
public class DocumentDiagnosticParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The text document.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The additional identifier  provided during registration.
    /// </summary>
    public /*TOpt*/System.String? Identifier;
    /// <summary>
    /// The result id of a previous response if provided.
    /// </summary>
    public /*TOpt*/System.String? PreviousResultId;
}

/// <summary>
/// A partial result for a document diagnostic report.
/// 
/// @since 3.17.0
/// </summary>
public class DocumentDiagnosticReportPartialResult
{
    /// <summary>
    /// Every object in the map has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public System.Collections.Generic.Dictionary<System.String, FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport> RelatedDocuments;
}

/// <summary>
/// Cancellation data returned from a diagnostic request.
/// 
/// @since 3.17.0
/// </summary>
public class DiagnosticServerCancellationData
{
    public System.Boolean RetriggerRequest;
}

/// <summary>
/// Diagnostic registration options.
/// 
/// @since 3.17.0
/// </summary>
public class DiagnosticRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public DiagnosticOptions DiagnosticOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

/// <summary>
/// Parameters of the workspace diagnostic request.
/// 
/// @since 3.17.0
/// </summary>
public class WorkspaceDiagnosticParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The additional identifier provided during registration.
    /// </summary>
    public /*TOpt*/System.String? Identifier;
    /// <summary>
    /// The currently known diagnostic reports with their
    /// previous result ids.
    /// </summary>
    public PreviousResultId[] PreviousResultIds;
}

/// <summary>
/// A workspace diagnostic report.
/// 
/// @since 3.17.0
/// </summary>
public class WorkspaceDiagnosticReport
{
    public WorkspaceDocumentDiagnosticReport[] Items;
}

/// <summary>
/// A partial result for a workspace diagnostic report.
/// 
/// @since 3.17.0
/// </summary>
public class WorkspaceDiagnosticReportPartialResult
{
    public WorkspaceDocumentDiagnosticReport[] Items;
}

/// <summary>
/// The params sent in an open notebook document notification.
/// 
/// @since 3.17.0
/// </summary>
public class DidOpenNotebookDocumentParams
{
    /// <summary>
    /// The notebook document that got opened.
    /// </summary>
    public NotebookDocument NotebookDocument;
    /// <summary>
    /// The text documents that represent the content
    /// of a notebook cell.
    /// </summary>
    public TextDocumentItem[] CellTextDocuments;
}

/// <summary>
/// The params sent in a change notebook document notification.
/// 
/// @since 3.17.0
/// </summary>
public class DidChangeNotebookDocumentParams
{
    /// <summary>
    /// The notebook document that did change. The version number points
    /// to the version after all provided changes have been applied. If
    /// only the text document content of a cell changes the notebook version
    /// doesn't necessarily have to change.
    /// </summary>
    public VersionedNotebookDocumentIdentifier NotebookDocument;
    /// <summary>
    /// The actual changes to the notebook document.
    /// 
    /// The changes describe single state changes to the notebook document.
    /// So if there are two changes c1 (at array index 0) and c2 (at array
    /// index 1) for a notebook in state S then c1 moves the notebook from
    /// S to S' and c2 from S' to S''. So c1 is computed on the state S and
    /// c2 is computed on the state S'.
    /// 
    /// To mirror the content of a notebook using change events use the following approach:
    /// - start with the same initial content
    /// - apply the 'notebookDocument/didChange' notifications in the order you receive them.
    /// - apply the `NotebookChangeEvent`s in a single notification in the order
    ///   you receive them.
    /// </summary>
    public NotebookDocumentChangeEvent Change;
}

/// <summary>
/// The params sent in a save notebook document notification.
/// 
/// @since 3.17.0
/// </summary>
public class DidSaveNotebookDocumentParams
{
    /// <summary>
    /// The notebook document that got saved.
    /// </summary>
    public NotebookDocumentIdentifier NotebookDocument;
}

/// <summary>
/// The params sent in a close notebook document notification.
/// 
/// @since 3.17.0
/// </summary>
public class DidCloseNotebookDocumentParams
{
    /// <summary>
    /// The notebook document that got closed.
    /// </summary>
    public NotebookDocumentIdentifier NotebookDocument;
    /// <summary>
    /// The text documents that represent the content
    /// of a notebook cell that got closed.
    /// </summary>
    public TextDocumentIdentifier[] CellTextDocuments;
}

public class RegistrationParams
{
    public Registration[] Registrations;
}

public class UnregistrationParams
{
    public Unregistration[] Unregisterations;
}

public class InitializeParams
{
    public _InitializeParams _InitializeParams; //XTEND
    public WorkspaceFoldersInitializeParams WorkspaceFoldersInitializeParams; //XTEND
}

/// <summary>
/// The result returned from an initialize request.
/// </summary>
public class InitializeResult
{
    /// <summary>
    /// The capabilities the language server provides.
    /// </summary>
    public ServerCapabilities Capabilities;
    /// <summary>
    /// Information about the server.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/NameStringWithVersionString ServerInfo;
}

/// <summary>
/// The data type of the ResponseError if the
/// initialize request fails.
/// </summary>
public class InitializeError
{
    /// <summary>
    /// Indicates whether the client execute the following retry logic:
    /// (1) show the message provided by the ResponseError to the user
    /// (2) user selects retry or cancel
    /// (3) if user selected retry the initialize method is sent again.
    /// </summary>
    public System.Boolean Retry;
}

public class InitializedParams
{
}

/// <summary>
/// The parameters of a change configuration notification.
/// </summary>
public class DidChangeConfigurationParams
{
    /// <summary>
    /// The actual changed settings
    /// </summary>
    public LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Settings;
}

public class DidChangeConfigurationRegistrationOptions
{
    /// <summary>
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/StringOrStrings Section;
}

/// <summary>
/// The parameters of a notification message.
/// </summary>
public class ShowMessageParams
{
    /// <summary>
    /// The message type. See <c>MessageType</c>
    /// </summary>
    public MessageType Type;
    /// <summary>
    /// The actual message.
    /// </summary>
    public System.String Message;
}

public class ShowMessageRequestParams
{
    /// <summary>
    /// The message type. See <c>MessageType</c>
    /// </summary>
    public MessageType Type;
    /// <summary>
    /// The actual message.
    /// </summary>
    public System.String Message;
    /// <summary>
    /// The message action items to present.
    /// </summary>
    public /*TOpt*/MessageActionItem[] Actions;
}

public class MessageActionItem
{
    /// <summary>
    /// A short title like 'Retry', 'Open Log' etc.
    /// </summary>
    public System.String Title;
}

/// <summary>
/// The log message parameters.
/// </summary>
public class LogMessageParams
{
    /// <summary>
    /// The message type. See <c>MessageType</c>
    /// </summary>
    public MessageType Type;
    /// <summary>
    /// The actual message.
    /// </summary>
    public System.String Message;
}

/// <summary>
/// The parameters sent in an open text document notification
/// </summary>
public class DidOpenTextDocumentParams
{
    /// <summary>
    /// The document that was opened.
    /// </summary>
    public TextDocumentItem TextDocument;
}

/// <summary>
/// The change text document notification's parameters.
/// </summary>
public class DidChangeTextDocumentParams
{
    /// <summary>
    /// The document that did change. The version number points
    /// to the version after all provided content changes have
    /// been applied.
    /// </summary>
    public VersionedTextDocumentIdentifier TextDocument;
    /// <summary>
    /// The actual content changes. The content changes describe single state changes
    /// to the document. So if there are two content changes c1 (at array index 0) and
    /// c2 (at array index 1) for a document in state S then c1 moves the document from
    /// S to S' and c2 from S' to S''. So c1 is computed on the state S and c2 is computed
    /// on the state S'.
    /// 
    /// To mirror the content of a document using change events use the following approach:
    /// - start with the same initial content
    /// - apply the 'textDocument/didChange' notifications in the order you receive them.
    /// - apply the `TextDocumentContentChangeEvent`s in a single notification in the order
    ///   you receive them.
    /// </summary>
    public TextDocumentContentChangeEvent[] ContentChanges;
}

/// <summary>
/// Describe options to be used when registered for text document change events.
/// </summary>
public class TextDocumentChangeRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    /// <summary>
    /// How documents are synced to the server.
    /// </summary>
    public TextDocumentSyncKind SyncKind;
}

/// <summary>
/// The parameters sent in a close text document notification
/// </summary>
public class DidCloseTextDocumentParams
{
    /// <summary>
    /// The document that was closed.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
}

/// <summary>
/// The parameters sent in a save text document notification
/// </summary>
public class DidSaveTextDocumentParams
{
    /// <summary>
    /// The document that was saved.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// Optional the content when saved. Depends on the includeText value
    /// when the save notification was requested.
    /// </summary>
    public /*TOpt*/System.String? Text;
}

/// <summary>
/// Save registration options.
/// </summary>
public class TextDocumentSaveRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public SaveOptions SaveOptions; //XTEND
}

/// <summary>
/// The parameters sent in a will save text document notification.
/// </summary>
public class WillSaveTextDocumentParams
{
    /// <summary>
    /// The document that will be saved.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The 'TextDocumentSaveReason'.
    /// </summary>
    public TextDocumentSaveReason Reason;
}

/// <summary>
/// A text edit applicable to a text document.
/// </summary>
public class TextEdit
{
    /// <summary>
    /// The range of the text document to be manipulated. To insert
    /// text into a document create a range where start === end.
    /// </summary>
    public Range Range;
    /// <summary>
    /// The string to be inserted. For delete operations use an
    /// empty string.
    /// </summary>
    public System.String NewText;
}

/// <summary>
/// The watched files change notification's parameters.
/// </summary>
public class DidChangeWatchedFilesParams
{
    /// <summary>
    /// The actual file events.
    /// </summary>
    public FileEvent[] Changes;
}

/// <summary>
/// Describe options to be used when registered for text document change events.
/// </summary>
public class DidChangeWatchedFilesRegistrationOptions
{
    /// <summary>
    /// The watchers to register.
    /// </summary>
    public FileSystemWatcher[] Watchers;
}

/// <summary>
/// The publish diagnostic notification's parameters.
/// </summary>
public class PublishDiagnosticsParams
{
    /// <summary>
    /// The URI for which diagnostic information is reported.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// Optional the version number of the document the diagnostics are published for.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/System.Int64? Version;
    /// <summary>
    /// An array of diagnostic information items.
    /// </summary>
    public Diagnostic[] Diagnostics;
}

/// <summary>
/// Completion parameters
/// </summary>
public class CompletionParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The completion context. This is only available it the client specifies
    /// to send this using the client capability `textDocument.completion.contextSupport === true`
    /// </summary>
    public /*TOpt*/CompletionContext? Context;
}

/// <summary>
/// A completion item represents a text snippet that is
/// proposed to complete text that is being typed.
/// </summary>
public class CompletionItem
{
    /// <summary>
    /// The label of this completion item.
    /// 
    /// The label property is also by default the text that
    /// is inserted when selecting this completion.
    /// 
    /// If label details are provided the label itself should
    /// be an unqualified name of the completion item.
    /// </summary>
    public System.String Label;
    /// <summary>
    /// Additional details for the label
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/CompletionItemLabelDetails? LabelDetails;
    /// <summary>
    /// The kind of this completion item. Based of the kind
    /// an icon is chosen by the editor.
    /// </summary>
    public /*TOpt*/CompletionItemKind Kind;
    /// <summary>
    /// Tags for this completion item.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/CompletionItemTag[] Tags;
    /// <summary>
    /// A human-readable string with additional information
    /// about this item, like type or symbol information.
    /// </summary>
    public /*TOpt*/System.String? Detail;
    /// <summary>
    /// A human-readable string that represents a doc-comment.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/StringOrMarkupContent Documentation;
    /// <summary>
    /// Indicates if this item is deprecated.
    /// @deprecated Use `tags` instead.
    /// </summary>
    public /*TOpt*/System.Boolean? Deprecated;
    /// <summary>
    /// Select this item when showing.
    /// 
    /// *Note* that only one completion item can be selected and that the
    /// tool / client decides which item that is. The rule is that the *first*
    /// item of those that match best is selected.
    /// </summary>
    public /*TOpt*/System.Boolean? Preselect;
    /// <summary>
    /// A string that should be used when comparing this item
    /// with other items. When `falsy` the <c>CompletionItem.label</c>
    /// is used.
    /// </summary>
    public /*TOpt*/System.String? SortText;
    /// <summary>
    /// A string that should be used when filtering a set of
    /// completion items. When `falsy` the <c>CompletionItem.label</c>
    /// is used.
    /// </summary>
    public /*TOpt*/System.String? FilterText;
    /// <summary>
    /// A string that should be inserted into a document when selecting
    /// this completion. When `falsy` the <c>CompletionItem.label</c>
    /// is used.
    /// 
    /// The `insertText` is subject to interpretation by the client side.
    /// Some tools might not take the string literally. For example
    /// VS Code when code complete is requested in this example
    /// `con<cursor position>` and a completion item with an `insertText` of
    /// `console` is provided it will only insert `sole`. Therefore it is
    /// recommended to use `textEdit` instead since it avoids additional client
    /// side interpretation.
    /// </summary>
    public /*TOpt*/System.String? InsertText;
    /// <summary>
    /// The format of the insert text. The format applies to both the
    /// `insertText` property and the `newText` property of a provided
    /// `textEdit`. If omitted defaults to `InsertTextFormat.PlainText`.
    /// 
    /// Please note that the insertTextFormat doesn't apply to
    /// `additionalTextEdits`.
    /// </summary>
    public /*TOpt*/InsertTextFormat InsertTextFormat;
    /// <summary>
    /// How whitespace and indentation is handled during completion
    /// item insertion. If not provided the clients default value depends on
    /// the `textDocument.completion.insertTextMode` client capability.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/InsertTextMode InsertTextMode;
    /// <summary>
    /// An <c>TextEdit</c> which is applied to a document when selecting
    /// this completion. When an edit is provided the value of
    /// <c>CompletionItem.insertText</c> is ignored.
    /// 
    /// Most editors support two different operations when accepting a completion
    /// item. One is to insert a completion text and the other is to replace an
    /// existing text with a completion text. Since this can usually not be
    /// predetermined by a server it can report both ranges. Clients need to
    /// signal support for `InsertReplaceEdits` via the
    /// `textDocument.completion.insertReplaceSupport` client capability
    /// property.
    /// 
    /// *Note 1:* The text edit's range as well as both ranges from an insert
    /// replace edit must be a [single line] and they must contain the position
    /// at which completion has been requested.
    /// *Note 2:* If an `InsertReplaceEdit` is returned the edit's insert range
    /// must be a prefix of the edit's replace range, that means it must be
    /// contained and starting at the same position.
    /// 
    /// @since 3.16.0 additional type `InsertReplaceEdit`
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/TextEditOrInsertReplaceEdit TextEdit;
    /// <summary>
    /// The edit text used if the completion item is part of a CompletionList and
    /// CompletionList defines an item default for the text edit range.
    /// 
    /// Clients will only honor this property if they opt into completion list
    /// item defaults using the capability `completionList.itemDefaults`.
    /// 
    /// If not provided and a list's default range is provided the label
    /// property is used as a text.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/System.String? TextEditText;
    /// <summary>
    /// An optional array of additional <c>TextEdit</c> that are applied when
    /// selecting this completion. Edits must not overlap (including the same insert position)
    /// with the main <c>CompletionItem.textEdit</c> nor with themselves.
    /// 
    /// Additional text edits should be used to change text unrelated to the current cursor position
    /// (for example adding an import statement at the top of the file if the completion item will
    /// insert an unqualified type).
    /// </summary>
    public /*TOpt*/TextEdit[] AdditionalTextEdits;
    /// <summary>
    /// An optional set of characters that when pressed while this completion is active will accept it first and
    /// then type that character. *Note* that all commit characters should have `length=1` and that superfluous
    /// characters will be ignored.
    /// </summary>
    public /*TOpt*/System.String[] CommitCharacters;
    /// <summary>
    /// An optional <c>Command</c> that is executed *after* inserting this completion. *Note* that
    /// additional modifications to the current document should be described with the
    /// <c>CompletionItem.additionalTextEdits</c>-property.
    /// </summary>
    public /*TOpt*/Command? Command;
    /// <summary>
    /// A data entry field that is preserved on a completion item between a
    /// <c>CompletionRequest</c> and a <c>CompletionResolveRequest</c>.
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Data;
}

/// <summary>
/// Represents a collection of <c>CompletionItem</c> to be presented
/// in the editor.
/// </summary>
public class CompletionList
{
    /// <summary>
    /// This list it not complete. Further typing results in recomputing this list.
    /// 
    /// Recomputed lists have all their items replaced (not appended) in the
    /// incomplete completion sessions.
    /// </summary>
    public System.Boolean IsIncomplete;
    /// <summary>
    /// In many cases the items of an actual completion result share the same
    /// value for properties like `commitCharacters` or the range of a text
    /// edit. A completion list can therefore define item defaults which will
    /// be used if a completion item itself doesn't specify the value.
    /// 
    /// If a completion list specifies a default value and a completion item
    /// also specifies a corresponding value the one from the item is used.
    /// 
    /// Servers are only allowed to return default values if the client
    /// signals support for this via the `completionList.itemDefaults`
    /// capability.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/CommitCharactersStringsWithEditRangeRangeOrInsertRangeWithReplaceRangeWithInsertTextFormatWithInsertTextModeWithDataLSPAny ItemDefaults;
    /// <summary>
    /// The completion items.
    /// </summary>
    public CompletionItem[] Items;
}

/// <summary>
/// Registration options for a <c>CompletionRequest</c>.
/// </summary>
public class CompletionRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public CompletionOptions CompletionOptions; //XTEND
}

/// <summary>
/// Parameters for a <c>HoverRequest</c>.
/// </summary>
public class HoverParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
}

/// <summary>
/// The result of a hover request.
/// </summary>
public class Hover
{
    /// <summary>
    /// The hover's content
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public MarkupContentOrMarkedStringOrMarkedStrings Contents;
    /// <summary>
    /// An optional range inside the text document that is used to
    /// visualize the hover, e.g. by changing the background color.
    /// </summary>
    public /*TOpt*/Range? Range;
}

/// <summary>
/// Registration options for a <c>HoverRequest</c>.
/// </summary>
public class HoverRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public HoverOptions HoverOptions; //XTEND
}

/// <summary>
/// Parameters for a <c>SignatureHelpRequest</c>.
/// </summary>
public class SignatureHelpParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    /// <summary>
    /// The signature help context. This is only available if the client specifies
    /// to send this using the client capability `textDocument.signatureHelp.contextSupport === true`
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/SignatureHelpContext? Context;
}

/// <summary>
/// Signature help represents the signature of something
/// callable. There can be multiple signature but only one
/// active and only one active parameter.
/// </summary>
public class SignatureHelp
{
    /// <summary>
    /// One or more signatures.
    /// </summary>
    public SignatureInformation[] Signatures;
    /// <summary>
    /// The active signature. If omitted or the value lies outside the
    /// range of `signatures` the value defaults to zero or is ignored if
    /// the `SignatureHelp` has no signatures.
    /// 
    /// Whenever possible implementors should make an active decision about
    /// the active signature and shouldn't rely on a default value.
    /// 
    /// In future version of the protocol this property might become
    /// mandatory to better express this.
    /// </summary>
    public /*TOpt*/System.UInt64? ActiveSignature;
    /// <summary>
    /// The active parameter of the active signature. If omitted or the value
    /// lies outside the range of `signatures[activeSignature].parameters`
    /// defaults to 0 if the active signature has parameters. If
    /// the active signature has no parameters it is ignored.
    /// In future version of the protocol this property might become
    /// mandatory to better express the active parameter if the
    /// active signature does have any.
    /// </summary>
    public /*TOpt*/System.UInt64? ActiveParameter;
}

/// <summary>
/// Registration options for a <c>SignatureHelpRequest</c>.
/// </summary>
public class SignatureHelpRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public SignatureHelpOptions SignatureHelpOptions; //XTEND
}

/// <summary>
/// Parameters for a <c>DefinitionRequest</c>.
/// </summary>
public class DefinitionParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
}

/// <summary>
/// Registration options for a <c>DefinitionRequest</c>.
/// </summary>
public class DefinitionRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public DefinitionOptions DefinitionOptions; //XTEND
}

/// <summary>
/// Parameters for a <c>ReferencesRequest</c>.
/// </summary>
public class ReferenceParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    public ReferenceContext Context;
}

/// <summary>
/// Registration options for a <c>ReferencesRequest</c>.
/// </summary>
public class ReferenceRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public ReferenceOptions ReferenceOptions; //XTEND
}

/// <summary>
/// Parameters for a <c>DocumentHighlightRequest</c>.
/// </summary>
public class DocumentHighlightParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
}

/// <summary>
/// A document highlight is a range inside a text document which deserves
/// special attention. Usually a document highlight is visualized by changing
/// the background color of its range.
/// </summary>
public class DocumentHighlight
{
    /// <summary>
    /// The range this highlight applies to.
    /// </summary>
    public Range Range;
    /// <summary>
    /// The highlight kind, default is <c>DocumentHighlightKind.Text</c>.
    /// </summary>
    public /*TOpt*/DocumentHighlightKind Kind;
}

/// <summary>
/// Registration options for a <c>DocumentHighlightRequest</c>.
/// </summary>
public class DocumentHighlightRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public DocumentHighlightOptions DocumentHighlightOptions; //XTEND
}

/// <summary>
/// Parameters for a <c>DocumentSymbolRequest</c>.
/// </summary>
public class DocumentSymbolParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The text document.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
}

/// <summary>
/// Represents information about programming constructs like variables, classes,
/// interfaces etc.
/// </summary>
public class SymbolInformation
{
    public BaseSymbolInformation BaseSymbolInformation; //XTEND
    /// <summary>
    /// Indicates if this symbol is deprecated.
    /// 
    /// @deprecated Use tags instead
    /// </summary>
    public /*TOpt*/System.Boolean? Deprecated;
    /// <summary>
    /// The location of this symbol. The location's range is used by a tool
    /// to reveal the location in the editor. If the symbol is selected in the
    /// tool the range's start information is used to position the cursor. So
    /// the range usually spans more than the actual symbol's name and does
    /// normally include things like visibility modifiers.
    /// 
    /// The range doesn't have to denote a node range in the sense of an abstract
    /// syntax tree. It can therefore not be used to re-construct a hierarchy of
    /// the symbols.
    /// </summary>
    public Location Location;
}

/// <summary>
/// Represents programming constructs like variables, classes, interfaces etc.
/// that appear in a document. Document symbols can be hierarchical and they
/// have two ranges: one that encloses its definition and one that points to
/// its most interesting range, e.g. the range of an identifier.
/// </summary>
public class DocumentSymbol
{
    /// <summary>
    /// The name of this symbol. Will be displayed in the user interface and therefore must not be
    /// an empty string or a string only consisting of white spaces.
    /// </summary>
    public System.String Name;
    /// <summary>
    /// More detail for this symbol, e.g the signature of a function.
    /// </summary>
    public /*TOpt*/System.String? Detail;
    /// <summary>
    /// The kind of this symbol.
    /// </summary>
    public SymbolKind Kind;
    /// <summary>
    /// Tags for this document symbol.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/SymbolTag[] Tags;
    /// <summary>
    /// Indicates if this symbol is deprecated.
    /// 
    /// @deprecated Use tags instead
    /// </summary>
    public /*TOpt*/System.Boolean? Deprecated;
    /// <summary>
    /// The range enclosing this symbol not including leading/trailing whitespace but everything else
    /// like comments. This information is typically used to determine if the clients cursor is
    /// inside the symbol to reveal in the symbol in the UI.
    /// </summary>
    public Range Range;
    /// <summary>
    /// The range that should be selected and revealed when this symbol is being picked, e.g the name of a function.
    /// Must be contained by the `range`.
    /// </summary>
    public Range SelectionRange;
    /// <summary>
    /// Children of this symbol, e.g. properties of a class.
    /// </summary>
    public /*TOpt*/DocumentSymbol[] Children;
}

/// <summary>
/// Registration options for a <c>DocumentSymbolRequest</c>.
/// </summary>
public class DocumentSymbolRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public DocumentSymbolOptions DocumentSymbolOptions; //XTEND
}

/// <summary>
/// The parameters of a <c>CodeActionRequest</c>.
/// </summary>
public class CodeActionParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The document in which the command was invoked.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The range for which the command was invoked.
    /// </summary>
    public Range Range;
    /// <summary>
    /// Context carrying additional information.
    /// </summary>
    public CodeActionContext Context;
}

/// <summary>
/// Represents a reference to a command. Provides a title which
/// will be used to represent a command in the UI and, optionally,
/// an array of arguments which will be passed to the command handler
/// function when invoked.
/// </summary>
public class Command
{
    /// <summary>
    /// Title of the command, like `save`.
    /// </summary>
    public System.String Title;
    /// <summary>
    /// The identifier of the actual command handler.
    /// </summary>
    public System.String Command_;
    /// <summary>
    /// Arguments that the command handler should be
    /// invoked with.
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull[] Arguments;
}

/// <summary>
/// A code action represents a change that can be performed in code, e.g. to fix a problem or
/// to refactor code.
/// 
/// A CodeAction must set either `edit` and/or a `command`. If both are supplied, the `edit` is applied first, then the `command` is executed.
/// </summary>
public class CodeAction
{
    /// <summary>
    /// A short, human-readable, title for this code action.
    /// </summary>
    public System.String Title;
    /// <summary>
    /// The kind of the code action.
    /// 
    /// Used to filter code actions.
    /// </summary>
    public /*TOpt*//*CodeActionKind*/string Kind;
    /// <summary>
    /// The diagnostics that this code action resolves.
    /// </summary>
    public /*TOpt*/Diagnostic[] Diagnostics;
    /// <summary>
    /// Marks this as a preferred action. Preferred actions are used by the `auto fix` command and can be targeted
    /// by keybindings.
    /// 
    /// A quick fix should be marked preferred if it properly addresses the underlying error.
    /// A refactoring should be marked preferred if it is the most reasonable choice of actions to take.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/System.Boolean? IsPreferred;
    /// <summary>
    /// Marks that the code action cannot currently be applied.
    /// 
    /// Clients should follow the following guidelines regarding disabled code actions:
    /// 
    ///   - Disabled code actions are not shown in automatic [lightbulbs](https://code.visualstudio.com/docs/editor/editingevolved#_code-action)
    ///     code action menus.
    /// 
    ///   - Disabled actions are shown as faded out in the code action menu when the user requests a more specific type
    ///     of code action, such as refactorings.
    /// 
    ///   - If the user has a [keybinding](https://code.visualstudio.com/docs/editor/refactoring#_keybindings-for-code-actions)
    ///     that auto applies a code action and only disabled code actions are returned, the client should show the user an
    ///     error message with `reason` in the editor.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/ReasonString Disabled;
    /// <summary>
    /// The workspace edit this code action performs.
    /// </summary>
    public /*TOpt*/WorkspaceEdit? Edit;
    /// <summary>
    /// A command this code action executes. If a code action
    /// provides an edit and a command, first the edit is
    /// executed and then the command.
    /// </summary>
    public /*TOpt*/Command? Command;
    /// <summary>
    /// A data entry field that is preserved on a code action between
    /// a `textDocument/codeAction` and a `codeAction/resolve` request.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Data;
}

/// <summary>
/// Registration options for a <c>CodeActionRequest</c>.
/// </summary>
public class CodeActionRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public CodeActionOptions CodeActionOptions; //XTEND
}

/// <summary>
/// The parameters of a <c>WorkspaceSymbolRequest</c>.
/// </summary>
public class WorkspaceSymbolParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// A query string to filter symbols by. Clients may send an empty
    /// string here to request all symbols.
    /// </summary>
    public System.String Query;
}

/// <summary>
/// A special workspace symbol that supports locations without a range.
/// 
/// See also SymbolInformation.
/// 
/// @since 3.17.0
/// </summary>
public class WorkspaceSymbol
{
    public BaseSymbolInformation BaseSymbolInformation; //XTEND
    /// <summary>
    /// The location of the symbol. Whether a server is allowed to
    /// return a location without a range depends on the client
    /// capability `workspace.symbol.resolveSupport`.
    /// 
    /// See SymbolInformation#location for more details.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public LocationOrUriDocumentUri Location;
    /// <summary>
    /// A data entry field that is preserved on a workspace symbol between a
    /// workspace symbol request and a workspace symbol resolve request.
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Data;
}

/// <summary>
/// Registration options for a <c>WorkspaceSymbolRequest</c>.
/// </summary>
public class WorkspaceSymbolRegistrationOptions
{
    public WorkspaceSymbolOptions WorkspaceSymbolOptions; //XTEND
}

/// <summary>
/// The parameters of a <c>CodeLensRequest</c>.
/// </summary>
public class CodeLensParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The document to request code lens for.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
}

/// <summary>
/// A code lens represents a <c>Command</c> that should be shown along with
/// source text, like the number of references, a way to run tests, etc.
/// 
/// A code lens is _unresolved_ when no command is associated to it. For performance
/// reasons the creation of a code lens and resolving should be done in two stages.
/// </summary>
public class CodeLens
{
    /// <summary>
    /// The range in which this code lens is valid. Should only span a single line.
    /// </summary>
    public Range Range;
    /// <summary>
    /// The command this code lens represents.
    /// </summary>
    public /*TOpt*/Command? Command;
    /// <summary>
    /// A data entry field that is preserved on a code lens item between
    /// a <c>CodeLensRequest</c> and a <c>CodeLensResolveRequest</c>
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Data;
}

/// <summary>
/// Registration options for a <c>CodeLensRequest</c>.
/// </summary>
public class CodeLensRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public CodeLensOptions CodeLensOptions; //XTEND
}

/// <summary>
/// The parameters of a <c>DocumentLinkRequest</c>.
/// </summary>
public class DocumentLinkParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    public PartialResultParams PartialResultParams; //MIXIN
    /// <summary>
    /// The document to provide document links for.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
}

/// <summary>
/// A document link is a range in a text document that links to an internal or external resource, like another
/// text document or a web site.
/// </summary>
public class DocumentLink
{
    /// <summary>
    /// The range this link applies to.
    /// </summary>
    public Range Range;
    /// <summary>
    /// The uri this link points to. If missing a resolve request is sent later.
    /// </summary>
    public /*TOpt*/System.String Target;
    /// <summary>
    /// The tooltip text when you hover over this link.
    /// 
    /// If a tooltip is provided, is will be displayed in a string that includes instructions on how to
    /// trigger the link, such as `{0} (ctrl + click)`. The specific instructions vary depending on OS,
    /// user settings, and localization.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/System.String? Tooltip;
    /// <summary>
    /// A data entry field that is preserved on a document link between a
    /// DocumentLinkRequest and a DocumentLinkResolveRequest.
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Data;
}

/// <summary>
/// Registration options for a <c>DocumentLinkRequest</c>.
/// </summary>
public class DocumentLinkRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public DocumentLinkOptions DocumentLinkOptions; //XTEND
}

/// <summary>
/// The parameters of a <c>DocumentFormattingRequest</c>.
/// </summary>
public class DocumentFormattingParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    /// <summary>
    /// The document to format.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The format options.
    /// </summary>
    public FormattingOptions Options;
}

/// <summary>
/// Registration options for a <c>DocumentFormattingRequest</c>.
/// </summary>
public class DocumentFormattingRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public DocumentFormattingOptions DocumentFormattingOptions; //XTEND
}

/// <summary>
/// The parameters of a <c>DocumentRangeFormattingRequest</c>.
/// </summary>
public class DocumentRangeFormattingParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    /// <summary>
    /// The document to format.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The range to format
    /// </summary>
    public Range Range;
    /// <summary>
    /// The format options
    /// </summary>
    public FormattingOptions Options;
}

/// <summary>
/// Registration options for a <c>DocumentRangeFormattingRequest</c>.
/// </summary>
public class DocumentRangeFormattingRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public DocumentRangeFormattingOptions DocumentRangeFormattingOptions; //XTEND
}

/// <summary>
/// The parameters of a <c>DocumentOnTypeFormattingRequest</c>.
/// </summary>
public class DocumentOnTypeFormattingParams
{
    /// <summary>
    /// The document to format.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The position around which the on type formatting should happen.
    /// This is not necessarily the exact position where the character denoted
    /// by the property `ch` got typed.
    /// </summary>
    public Position Position;
    /// <summary>
    /// The character that has been typed that triggered the formatting
    /// on type request. That is not necessarily the last character that
    /// got inserted into the document since the client could auto insert
    /// characters as well (e.g. like automatic brace completion).
    /// </summary>
    public System.String Ch;
    /// <summary>
    /// The formatting options.
    /// </summary>
    public FormattingOptions Options;
}

/// <summary>
/// Registration options for a <c>DocumentOnTypeFormattingRequest</c>.
/// </summary>
public class DocumentOnTypeFormattingRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public DocumentOnTypeFormattingOptions DocumentOnTypeFormattingOptions; //XTEND
}

/// <summary>
/// The parameters of a <c>RenameRequest</c>.
/// </summary>
public class RenameParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    /// <summary>
    /// The document to rename.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The position at which this request was sent.
    /// </summary>
    public Position Position;
    /// <summary>
    /// The new name of the symbol. If the given name is not valid the
    /// request must return a <c>ResponseError</c> with an
    /// appropriate message set.
    /// </summary>
    public System.String NewName;
}

/// <summary>
/// Registration options for a <c>RenameRequest</c>.
/// </summary>
public class RenameRegistrationOptions
{
    public TextDocumentRegistrationOptions TextDocumentRegistrationOptions; //XTEND
    public RenameOptions RenameOptions; //XTEND
}

public class PrepareRenameParams
{
    public TextDocumentPositionParams TextDocumentPositionParams; //XTEND
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
}

/// <summary>
/// The parameters of a <c>ExecuteCommandRequest</c>.
/// </summary>
public class ExecuteCommandParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    /// <summary>
    /// The identifier of the actual command handler.
    /// </summary>
    public System.String Command;
    /// <summary>
    /// Arguments that the command should be invoked with.
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull[] Arguments;
}

/// <summary>
/// Registration options for a <c>ExecuteCommandRequest</c>.
/// </summary>
public class ExecuteCommandRegistrationOptions
{
    public ExecuteCommandOptions ExecuteCommandOptions; //XTEND
}

/// <summary>
/// The parameters passed via an apply workspace edit request.
/// </summary>
public class ApplyWorkspaceEditParams
{
    /// <summary>
    /// An optional label of the workspace edit. This label is
    /// presented in the user interface for example on an undo
    /// stack to undo the workspace edit.
    /// </summary>
    public /*TOpt*/System.String? Label;
    /// <summary>
    /// The edits to apply.
    /// </summary>
    public WorkspaceEdit Edit;
}

/// <summary>
/// The result returned from the apply workspace edit request.
/// 
/// @since 3.17 renamed from ApplyWorkspaceEditResponse
/// </summary>
public class ApplyWorkspaceEditResult
{
    /// <summary>
    /// Indicates whether the edit was applied or not.
    /// </summary>
    public System.Boolean Applied;
    /// <summary>
    /// An optional textual description for why the edit was not applied.
    /// This may be used by the server for diagnostic logging or to provide
    /// a suitable error for a request that triggered the edit.
    /// </summary>
    public /*TOpt*/System.String? FailureReason;
    /// <summary>
    /// Depending on the client's failure handling strategy `failedChange` might
    /// contain the index of the change that failed. This property is only available
    /// if the client signals a `failureHandlingStrategy` in its client capabilities.
    /// </summary>
    public /*TOpt*/System.UInt64? FailedChange;
}

public class WorkDoneProgressBegin
{
    /// <summary>
    /// The value is always "begin". 
    /// The value is always "begin". 
    /// </summary>
    public System.String Kind;
    /// <summary>
    /// Mandatory title of the progress operation. Used to briefly inform about
    /// the kind of operation being performed.
    /// 
    /// Examples: "Indexing" or "Linking dependencies".
    /// </summary>
    public System.String Title;
    /// <summary>
    /// Controls if a cancel button should show to allow the user to cancel the
    /// long running operation. Clients that don't support cancellation are allowed
    /// to ignore the setting.
    /// </summary>
    public /*TOpt*/System.Boolean? Cancellable;
    /// <summary>
    /// Optional, more detailed associated progress message. Contains
    /// complementary information to the `title`.
    /// 
    /// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
    /// If unset, the previous progress message (if any) is still valid.
    /// </summary>
    public /*TOpt*/System.String? Message;
    /// <summary>
    /// Optional progress percentage to display (value 100 is considered 100%).
    /// If not provided infinite progress is assumed and clients are allowed
    /// to ignore the `percentage` value in subsequent in report notifications.
    /// 
    /// The value should be steadily rising. Clients are free to ignore values
    /// that are not following this rule. The value range is [0, 100].
    /// </summary>
    public /*TOpt*/System.UInt64? Percentage;
}

public class WorkDoneProgressReport
{
    /// <summary>
    /// The value is always "report". 
    /// The value is always "report". 
    /// </summary>
    public System.String Kind;
    /// <summary>
    /// Controls enablement state of a cancel button.
    /// 
    /// Clients that don't support cancellation or don't support controlling the button's
    /// enablement state are allowed to ignore the property.
    /// </summary>
    public /*TOpt*/System.Boolean? Cancellable;
    /// <summary>
    /// Optional, more detailed associated progress message. Contains
    /// complementary information to the `title`.
    /// 
    /// Examples: "3/25 files", "project/src/module2", "node_modules/some_dep".
    /// If unset, the previous progress message (if any) is still valid.
    /// </summary>
    public /*TOpt*/System.String? Message;
    /// <summary>
    /// Optional progress percentage to display (value 100 is considered 100%).
    /// If not provided infinite progress is assumed and clients are allowed
    /// to ignore the `percentage` value in subsequent in report notifications.
    /// 
    /// The value should be steadily rising. Clients are free to ignore values
    /// that are not following this rule. The value range is [0, 100]
    /// </summary>
    public /*TOpt*/System.UInt64? Percentage;
}

public class WorkDoneProgressEnd
{
    /// <summary>
    /// The value is always "end". 
    /// The value is always "end". 
    /// </summary>
    public System.String Kind;
    /// <summary>
    /// Optional, a final message indicating to for example indicate the outcome
    /// of the operation.
    /// </summary>
    public /*TOpt*/System.String? Message;
}

public class SetTraceParams
{
    public /*TraceValues*/string Value;
}

public class LogTraceParams
{
    public System.String Message;
    public /*TOpt*/System.String? Verbose;
}

public class CancelParams
{
    /// <summary>
    /// The request id to cancel.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public IntegerOrString Id;
}

public class ProgressParams
{
    /// <summary>
    /// The progress token provided by the client or server.
    /// </summary>
    public ProgressToken Token;
    /// <summary>
    /// The progress data.
    /// </summary>
    public LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Value;
}

/// <summary>
/// A parameter literal used in requests to pass a text document and a position inside that
/// document.
/// </summary>
public class TextDocumentPositionParams
{
    /// <summary>
    /// The text document.
    /// </summary>
    public TextDocumentIdentifier TextDocument;
    /// <summary>
    /// The position inside the text document.
    /// </summary>
    public Position Position;
}

public class WorkDoneProgressParams
{
    /// <summary>
    /// An optional token that a server can use to report work done progress.
    /// </summary>
    public /*TOpt*/ProgressToken WorkDoneToken;
}

public class PartialResultParams
{
    /// <summary>
    /// An optional token that a server can use to report partial results (e.g. streaming) to
    /// the client.
    /// </summary>
    public /*TOpt*/ProgressToken PartialResultToken;
}

/// <summary>
/// Represents the connection of two locations. Provides additional metadata over normal <c>Location</c>,
/// including an origin range.
/// </summary>
public class LocationLink
{
    /// <summary>
    /// Span of the origin of this link.
    /// 
    /// Used as the underlined span for mouse interaction. Defaults to the word range at
    /// the definition position.
    /// </summary>
    public /*TOpt*/Range? OriginSelectionRange;
    /// <summary>
    /// The target resource identifier of this link.
    /// </summary>
    public System.String TargetUri;
    /// <summary>
    /// The full target range of this link. If the target for example is a symbol then target range is the
    /// range enclosing this symbol not including leading/trailing whitespace but everything else
    /// like comments. This information is typically used to highlight the range in the editor.
    /// </summary>
    public Range TargetRange;
    /// <summary>
    /// The range that should be selected and revealed when this link is being followed, e.g the name of a function.
    /// Must be contained by the `targetRange`. See also `DocumentSymbol#range`
    /// </summary>
    public Range TargetSelectionRange;
}

/// <summary>
/// A range in a text document expressed as (zero-based) start and end positions.
/// 
/// If you want to specify a range that contains a line including the line ending
/// character(s) then use an end position denoting the start of the next line.
/// For example:
/// ```ts
/// {
///     start: { line: 5, character: 23 }
///     end : { line 6, character : 0 }
/// }
/// ```
/// </summary>
public class Range
{
    /// <summary>
    /// The range's start position.
    /// </summary>
    public Position Start;
    /// <summary>
    /// The range's end position.
    /// </summary>
    public Position End;
}

public class ImplementationOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// Static registration options to be returned in the initialize
/// request.
/// </summary>
public class StaticRegistrationOptions
{
    /// <summary>
    /// The id used to register the request. The id can be used to deregister
    /// the request again. See also Registration#id.
    /// </summary>
    public /*TOpt*/System.String? Id;
}

public class TypeDefinitionOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// The workspace folder change event.
/// </summary>
public class WorkspaceFoldersChangeEvent
{
    /// <summary>
    /// The array of added workspace folders
    /// </summary>
    public WorkspaceFolder[] Added;
    /// <summary>
    /// The array of the removed workspace folders
    /// </summary>
    public WorkspaceFolder[] Removed;
}

public class ConfigurationItem
{
    /// <summary>
    /// The scope to get the configuration section for.
    /// </summary>
    public /*TOpt*/System.String ScopeUri;
    /// <summary>
    /// The configuration section asked for.
    /// </summary>
    public /*TOpt*/System.String? Section;
}

/// <summary>
/// A literal to identify a text document in the client.
/// </summary>
public class TextDocumentIdentifier
{
    /// <summary>
    /// The text document's uri.
    /// </summary>
    public System.String Uri;
}

/// <summary>
/// Represents a color in RGBA space.
/// </summary>
public class Color
{
    /// <summary>
    /// The red component of this color in the range [0-1].
    /// </summary>
    public System.Double Red;
    /// <summary>
    /// The green component of this color in the range [0-1].
    /// </summary>
    public System.Double Green;
    /// <summary>
    /// The blue component of this color in the range [0-1].
    /// </summary>
    public System.Double Blue;
    /// <summary>
    /// The alpha component of this color in the range [0-1].
    /// </summary>
    public System.Double Alpha;
}

public class DocumentColorOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

public class FoldingRangeOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

public class DeclarationOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// Position in a text document expressed as zero-based line and character
/// offset. Prior to 3.17 the offsets were always based on a UTF-16 string
/// representation. So a string of the form `a𐐀b` the character offset of the
/// character `a` is 0, the character offset of `𐐀` is 1 and the character
/// offset of b is 3 since `𐐀` is represented using two code units in UTF-16.
/// Since 3.17 clients and servers can agree on a different string encoding
/// representation (e.g. UTF-8). The client announces it's supported encoding
/// via the client capability [`general.positionEncodings`](https://microsoft.github.io/language-server-protocol/specifications/specification-current/#clientCapabilities).
/// The value is an array of position encodings the client supports, with
/// decreasing preference (e.g. the encoding at index `0` is the most preferred
/// one). To stay backwards compatible the only mandatory encoding is UTF-16
/// represented via the string `utf-16`. The server can pick one of the
/// encodings offered by the client and signals that encoding back to the
/// client via the initialize result's property
/// [`capabilities.positionEncoding`](https://microsoft.github.io/language-server-protocol/specifications/specification-current/#serverCapabilities). If the string value
/// `utf-16` is missing from the client's capability `general.positionEncodings`
/// servers can safely assume that the client supports UTF-16. If the server
/// omits the position encoding in its initialize result the encoding defaults
/// to the string value `utf-16`. Implementation considerations: since the
/// conversion from one encoding into another requires the content of the
/// file / line the conversion is best done where the file is read which is
/// usually on the server side.
/// 
/// Positions are line end character agnostic. So you can not specify a position
/// that denotes `\r|\n` or `\n|` where `|` represents the character offset.
/// 
/// @since 3.17.0 - support for negotiated position encoding.
/// </summary>
public class Position
{
    /// <summary>
    /// Line position in a document (zero-based).
    /// 
    /// If a line number is greater than the number of lines in a document, it defaults back to the number of lines in the document.
    /// If a line number is negative, it defaults to 0.
    /// </summary>
    public System.UInt64 Line;
    /// <summary>
    /// Character offset on a line in a document (zero-based).
    /// 
    /// The meaning of this offset is determined by the negotiated
    /// `PositionEncodingKind`.
    /// 
    /// If the character value is greater than the line length it defaults back to the
    /// line length.
    /// </summary>
    public System.UInt64 Character;
}

public class SelectionRangeOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// Call hierarchy options used during static registration.
/// 
/// @since 3.16.0
/// </summary>
public class CallHierarchyOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokensOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
    /// <summary>
    /// The legend used by the server
    /// </summary>
    public SemanticTokensLegend Legend;
    /// <summary>
    /// Server supports providing semantic tokens for a specific range
    /// of a document.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrAnyByString Range;
    /// <summary>
    /// Server supports providing semantic tokens for a full document.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrDeltaBoolean Full;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokensEdit
{
    /// <summary>
    /// The start offset of the edit.
    /// </summary>
    public System.UInt64 Start;
    /// <summary>
    /// The count of elements to remove.
    /// </summary>
    public System.UInt64 DeleteCount;
    /// <summary>
    /// The elements to insert.
    /// </summary>
    public /*TOpt*/System.UInt64[] Data;
}

public class LinkedEditingRangeOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// Represents information on a file/folder create.
/// 
/// @since 3.16.0
/// </summary>
public class FileCreate
{
    /// <summary>
    /// A file:// URI for the location of the file/folder being created.
    /// </summary>
    public System.String Uri;
}

/// <summary>
/// Describes textual changes on a text document. A TextDocumentEdit describes all changes
/// on a document version Si and after they are applied move the document to version Si+1.
/// So the creator of a TextDocumentEdit doesn't need to sort the array of edits or do any
/// kind of ordering. However the edits must be non overlapping.
/// </summary>
public class TextDocumentEdit
{
    /// <summary>
    /// The text document to change.
    /// </summary>
    public OptionalVersionedTextDocumentIdentifier TextDocument;
    /// <summary>
    /// The edits to be applied.
    /// 
    /// @since 3.16.0 - support for AnnotatedTextEdit. This is guarded using a
    /// client capability.
    /// 
    /// Every object in the array has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public TextEditOrAnnotatedTextEdit[] Edits;
}

/// <summary>
/// Create file operation.
/// </summary>
public class CreateFile
{
    public ResourceOperation ResourceOperation; //XTEND
    /// <summary>
    /// A create
    /// The value is always "create". 
    /// The value is always "create". 
    /// </summary>
    public System.String Kind;
    /// <summary>
    /// The resource to create.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// Additional options
    /// </summary>
    public /*TOpt*/CreateFileOptions? Options;
}

/// <summary>
/// Rename file operation
/// </summary>
public class RenameFile
{
    public ResourceOperation ResourceOperation; //XTEND
    /// <summary>
    /// A rename
    /// The value is always "rename". 
    /// The value is always "rename". 
    /// </summary>
    public System.String Kind;
    /// <summary>
    /// The old (existing) location.
    /// </summary>
    public System.String OldUri;
    /// <summary>
    /// The new location.
    /// </summary>
    public System.String NewUri;
    /// <summary>
    /// Rename options.
    /// </summary>
    public /*TOpt*/RenameFileOptions? Options;
}

/// <summary>
/// Delete file operation
/// </summary>
public class DeleteFile
{
    public ResourceOperation ResourceOperation; //XTEND
    /// <summary>
    /// A delete
    /// The value is always "delete". 
    /// The value is always "delete". 
    /// </summary>
    public System.String Kind;
    /// <summary>
    /// The file to delete.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// Delete options.
    /// </summary>
    public /*TOpt*/DeleteFileOptions? Options;
}

/// <summary>
/// Additional information that describes document changes.
/// 
/// @since 3.16.0
/// </summary>
public class ChangeAnnotation
{
    /// <summary>
    /// A human-readable string describing the actual change. The string
    /// is rendered prominent in the user interface.
    /// </summary>
    public System.String Label;
    /// <summary>
    /// A flag which indicates that user confirmation is needed
    /// before applying the change.
    /// </summary>
    public /*TOpt*/System.Boolean? NeedsConfirmation;
    /// <summary>
    /// A human-readable string which is rendered less prominent in
    /// the user interface.
    /// </summary>
    public /*TOpt*/System.String? Description;
}

/// <summary>
/// A filter to describe in which file operation requests or notifications
/// the server is interested in receiving.
/// 
/// @since 3.16.0
/// </summary>
public class FileOperationFilter
{
    /// <summary>
    /// A Uri scheme like `file` or `untitled`.
    /// </summary>
    public /*TOpt*/System.String? Scheme;
    /// <summary>
    /// The actual file operation pattern.
    /// </summary>
    public FileOperationPattern Pattern;
}

/// <summary>
/// Represents information on a file/folder rename.
/// 
/// @since 3.16.0
/// </summary>
public class FileRename
{
    /// <summary>
    /// A file:// URI for the original location of the file/folder being renamed.
    /// </summary>
    public System.String OldUri;
    /// <summary>
    /// A file:// URI for the new location of the file/folder being renamed.
    /// </summary>
    public System.String NewUri;
}

/// <summary>
/// Represents information on a file/folder delete.
/// 
/// @since 3.16.0
/// </summary>
public class FileDelete
{
    /// <summary>
    /// A file:// URI for the location of the file/folder being deleted.
    /// </summary>
    public System.String Uri;
}

public class MonikerOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// Type hierarchy options used during static registration.
/// 
/// @since 3.17.0
/// </summary>
public class TypeHierarchyOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// @since 3.17.0
/// </summary>
public class InlineValueContext
{
    /// <summary>
    /// The stack frame (as a DAP Id) where the execution has stopped.
    /// </summary>
    public System.Int64 FrameId;
    /// <summary>
    /// The document range where execution has stopped.
    /// Typically the end position of the range denotes the line where the inline values are shown.
    /// </summary>
    public Range StoppedLocation;
}

/// <summary>
/// Provide inline value as text.
/// 
/// @since 3.17.0
/// </summary>
public class InlineValueText
{
    /// <summary>
    /// The document range for which the inline value applies.
    /// </summary>
    public Range Range;
    /// <summary>
    /// The text of the inline value.
    /// </summary>
    public System.String Text;
}

/// <summary>
/// Provide inline value through a variable lookup.
/// If only a range is specified, the variable name will be extracted from the underlying document.
/// An optional variable name can be used to override the extracted name.
/// 
/// @since 3.17.0
/// </summary>
public class InlineValueVariableLookup
{
    /// <summary>
    /// The document range for which the inline value applies.
    /// The range is used to extract the variable name from the underlying document.
    /// </summary>
    public Range Range;
    /// <summary>
    /// If specified the name of the variable to look up.
    /// </summary>
    public /*TOpt*/System.String? VariableName;
    /// <summary>
    /// How to perform the lookup.
    /// </summary>
    public System.Boolean CaseSensitiveLookup;
}

/// <summary>
/// Provide an inline value through an expression evaluation.
/// If only a range is specified, the expression will be extracted from the underlying document.
/// An optional expression can be used to override the extracted expression.
/// 
/// @since 3.17.0
/// </summary>
public class InlineValueEvaluatableExpression
{
    /// <summary>
    /// The document range for which the inline value applies.
    /// The range is used to extract the evaluatable expression from the underlying document.
    /// </summary>
    public Range Range;
    /// <summary>
    /// If specified the expression overrides the extracted expression.
    /// </summary>
    public /*TOpt*/System.String? Expression;
}

/// <summary>
/// Inline value options used during static registration.
/// 
/// @since 3.17.0
/// </summary>
public class InlineValueOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// An inlay hint label part allows for interactive and composite labels
/// of inlay hints.
/// 
/// @since 3.17.0
/// </summary>
public class InlayHintLabelPart
{
    /// <summary>
    /// The value of this label part.
    /// </summary>
    public System.String Value;
    /// <summary>
    /// The tooltip text when you hover over this label part. Depending on
    /// the client capability `inlayHint.resolveSupport` clients might resolve
    /// this property late using the resolve request.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/StringOrMarkupContent Tooltip;
    /// <summary>
    /// An optional source code location that represents this
    /// label part.
    /// 
    /// The editor will use this location for the hover and for code navigation
    /// features: This part will become a clickable link that resolves to the
    /// definition of the symbol at the given location (not necessarily the
    /// location itself), it shows the hover that shows at the given location,
    /// and it shows a context menu with further code navigation commands.
    /// 
    /// Depending on the client capability `inlayHint.resolveSupport` clients
    /// might resolve this property late using the resolve request.
    /// </summary>
    public /*TOpt*/Location? Location;
    /// <summary>
    /// An optional command for this label part.
    /// 
    /// Depending on the client capability `inlayHint.resolveSupport` clients
    /// might resolve this property late using the resolve request.
    /// </summary>
    public /*TOpt*/Command? Command;
}

/// <summary>
/// A `MarkupContent` literal represents a string value which content is interpreted base on its
/// kind flag. Currently the protocol supports `plaintext` and `markdown` as markup kinds.
/// 
/// If the kind is `markdown` then the value can contain fenced code blocks like in GitHub issues.
/// See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting
/// 
/// Here is an example how such a string can be constructed using JavaScript / TypeScript:
/// ```ts
/// let markdown: MarkdownContent = {
///  kind: MarkupKind.Markdown,
///  value: [
///    '# Header',
///    'Some text',
///    '```typescript',
///    'someCode();',
///    '```'
///  ].join('\n')
/// };
/// ```
/// 
/// *Please Note* that clients might sanitize the return markdown. A client could decide to
/// remove HTML from the markdown to avoid script execution.
/// </summary>
public class MarkupContent
{
    /// <summary>
    /// The type of the Markup
    /// </summary>
    public /*MarkupKind*/string Kind;
    /// <summary>
    /// The content itself
    /// </summary>
    public System.String Value;
}

/// <summary>
/// Inlay hint options used during static registration.
/// 
/// @since 3.17.0
/// </summary>
public class InlayHintOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
    /// <summary>
    /// The server provides support to resolve additional
    /// information for an inlay hint item.
    /// </summary>
    public /*TOpt*/System.Boolean? ResolveProvider;
}

/// <summary>
/// A full diagnostic report with a set of related documents.
/// 
/// @since 3.17.0
/// </summary>
public class RelatedFullDocumentDiagnosticReport
{
    public FullDocumentDiagnosticReport FullDocumentDiagnosticReport; //XTEND
    /// <summary>
    /// Diagnostics of related documents. This information is useful
    /// in programming languages where code in a file A can generate
    /// diagnostics in a file B which A depends on. An example of
    /// such a language is C/C++ where marco definitions in a file
    /// a.cpp and result in errors in a header file b.hpp.
    /// 
    /// @since 3.17.0
    /// 
    /// Every object in the map has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/System.Collections.Generic.Dictionary<System.String, FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport> RelatedDocuments;
}

/// <summary>
/// An unchanged diagnostic report with a set of related documents.
/// 
/// @since 3.17.0
/// </summary>
public class RelatedUnchangedDocumentDiagnosticReport
{
    public UnchangedDocumentDiagnosticReport UnchangedDocumentDiagnosticReport; //XTEND
    /// <summary>
    /// Diagnostics of related documents. This information is useful
    /// in programming languages where code in a file A can generate
    /// diagnostics in a file B which A depends on. An example of
    /// such a language is C/C++ where marco definitions in a file
    /// a.cpp and result in errors in a header file b.hpp.
    /// 
    /// @since 3.17.0
    /// 
    /// Every object in the map has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/System.Collections.Generic.Dictionary<System.String, FullDocumentDiagnosticReportOrUnchangedDocumentDiagnosticReport> RelatedDocuments;
}

/// <summary>
/// A diagnostic report with a full set of problems.
/// 
/// @since 3.17.0
/// </summary>
public class FullDocumentDiagnosticReport
{
    /// <summary>
    /// A full document diagnostic report.
    /// The value is always "full". 
    /// The value is always "full". 
    /// </summary>
    public System.String Kind;
    /// <summary>
    /// An optional result id. If provided it will
    /// be sent on the next diagnostic request for the
    /// same document.
    /// </summary>
    public /*TOpt*/System.String? ResultId;
    /// <summary>
    /// The actual items.
    /// </summary>
    public Diagnostic[] Items;
}

/// <summary>
/// A diagnostic report indicating that the last returned
/// report is still accurate.
/// 
/// @since 3.17.0
/// </summary>
public class UnchangedDocumentDiagnosticReport
{
    /// <summary>
    /// A document diagnostic report indicating
    /// no changes to the last result. A server can
    /// only return `unchanged` if result ids are
    /// provided.
    /// The value is always "unchanged". 
    /// The value is always "unchanged". 
    /// </summary>
    public System.String Kind;
    /// <summary>
    /// A result id which will be sent on the next
    /// diagnostic request for the same document.
    /// </summary>
    public System.String ResultId;
}

/// <summary>
/// Diagnostic options.
/// 
/// @since 3.17.0
/// </summary>
public class DiagnosticOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
    /// <summary>
    /// An optional identifier under which the diagnostics are
    /// managed by the client.
    /// </summary>
    public /*TOpt*/System.String? Identifier;
    /// <summary>
    /// Whether the language has inter file dependencies meaning that
    /// editing code in one file can result in a different diagnostic
    /// set in another file. Inter file dependencies are common for
    /// most programming languages and typically uncommon for linters.
    /// </summary>
    public System.Boolean InterFileDependencies;
    /// <summary>
    /// The server provides support for workspace diagnostics as well.
    /// </summary>
    public System.Boolean WorkspaceDiagnostics;
}

/// <summary>
/// A previous result id in a workspace pull request.
/// 
/// @since 3.17.0
/// </summary>
public class PreviousResultId
{
    /// <summary>
    /// The URI for which the client knowns a
    /// result id.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// The value of the previous result id.
    /// </summary>
    public System.String Value;
}

/// <summary>
/// A notebook document.
/// 
/// @since 3.17.0
/// </summary>
public class NotebookDocument
{
    /// <summary>
    /// The notebook document's uri.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// The type of the notebook.
    /// </summary>
    public System.String NotebookType;
    /// <summary>
    /// The version number of this document (it will increase after each
    /// change, including undo/redo).
    /// </summary>
    public System.Int64 Version;
    /// <summary>
    /// Additional metadata stored with the notebook
    /// document.
    /// 
    /// Note: should always be an object literal (e.g. LSPObject)
    /// </summary>
    public /*TOpt*/LSPObject Metadata;
    /// <summary>
    /// The cells of a notebook.
    /// </summary>
    public NotebookCell[] Cells;
}

/// <summary>
/// An item to transfer a text document from the client to the
/// server.
/// </summary>
public class TextDocumentItem
{
    /// <summary>
    /// The text document's uri.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// The text document's language identifier.
    /// </summary>
    public System.String LanguageId;
    /// <summary>
    /// The version number of this document (it will increase after each
    /// change, including undo/redo).
    /// </summary>
    public System.Int64 Version;
    /// <summary>
    /// The content of the opened text document.
    /// </summary>
    public System.String Text;
}

/// <summary>
/// A versioned notebook document identifier.
/// 
/// @since 3.17.0
/// </summary>
public class VersionedNotebookDocumentIdentifier
{
    /// <summary>
    /// The version number of this notebook document.
    /// </summary>
    public System.Int64 Version;
    /// <summary>
    /// The notebook document's uri.
    /// </summary>
    public System.String Uri;
}

/// <summary>
/// A change event for a notebook document.
/// 
/// @since 3.17.0
/// </summary>
public class NotebookDocumentChangeEvent
{
    /// <summary>
    /// The changed meta data if any.
    /// 
    /// Note: should always be an object literal (e.g. LSPObject)
    /// </summary>
    public /*TOpt*/LSPObject Metadata;
    /// <summary>
    /// Changes to cells
    /// </summary>
    public /*TOpt*/StructureArrayNotebookCellArrayChangeWithDidOpenTextDocumentItemsWithDidCloseTextDocumentIdentifiersWithDataNotebookCellsWithTextContentDocumentVersionedTextDocumentIdentifierWithChangesTextDocumentContentChangeEventses Cells;
}

/// <summary>
/// A literal to identify a notebook document in the client.
/// 
/// @since 3.17.0
/// </summary>
public class NotebookDocumentIdentifier
{
    /// <summary>
    /// The notebook document's uri.
    /// </summary>
    public System.String Uri;
}

/// <summary>
/// General parameters to register for a notification or to register a provider.
/// </summary>
public class Registration
{
    /// <summary>
    /// The id used to register the request. The id can be used to deregister
    /// the request again.
    /// </summary>
    public System.String Id;
    /// <summary>
    /// The method / capability to register for.
    /// </summary>
    public System.String Method;
    /// <summary>
    /// Options necessary for the registration.
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull RegisterOptions;
}

/// <summary>
/// General parameters to unregister a request or notification.
/// </summary>
public class Unregistration
{
    /// <summary>
    /// The id used to unregister the request or notification. Usually an id
    /// provided during the register request.
    /// </summary>
    public System.String Id;
    /// <summary>
    /// The method to unregister for.
    /// </summary>
    public System.String Method;
}

/// <summary>
/// The initialize parameters
/// </summary>
public class _InitializeParams
{
    public WorkDoneProgressParams WorkDoneProgressParams; //MIXIN
    /// <summary>
    /// The process Id of the parent process that started
    /// the server.
    /// 
    /// Is `null` if the process has not been started by another process.
    /// If the parent process is not alive then the server should exit.
    /// </summary>
    public /*TOr1*//*TOpt*/System.Int64? ProcessId;
    /// <summary>
    /// Information about the client
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/NameStringWithVersionString_ ClientInfo;
    /// <summary>
    /// The locale the client is currently showing the user interface
    /// in. This must not necessarily be the locale of the operating
    /// system.
    /// 
    /// Uses IETF language tags as the value's syntax
    /// (See https://en.wikipedia.org/wiki/IETF_language_tag)
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.String? Locale;
    /// <summary>
    /// The rootPath of the workspace. Is null
    /// if no folder is open.
    /// 
    /// @deprecated in favour of rootUri.
    /// </summary>
    public /*TOpt*//*TOr1*//*TOpt*/System.String? RootPath;
    /// <summary>
    /// The rootUri of the workspace. Is null if no
    /// folder is open. If both `rootPath` and `rootUri` are set
    /// `rootUri` wins.
    /// 
    /// @deprecated in favour of workspaceFolders.
    /// </summary>
    public /*TOr1*//*TOpt*/System.String RootUri;
    /// <summary>
    /// The capabilities provided by the client (editor or tool)
    /// </summary>
    public ClientCapabilities Capabilities;
    /// <summary>
    /// User provided initialization options.
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull InitializationOptions;
    /// <summary>
    /// The initial trace setting. If omitted trace is disabled ('off').
    /// </summary>
    public /*TOpt*//*TraceValues*/string Trace;
}

public class WorkspaceFoldersInitializeParams
{
    /// <summary>
    /// The workspace folders configured in the client when the server starts.
    /// 
    /// This property is only available if the client supports workspace folders.
    /// It can be `null` if the client supports workspace folders but none are
    /// configured.
    /// 
    /// @since 3.6.0
    /// </summary>
    public /*TOpt*//*TOr1*//*TOpt*/WorkspaceFolder[] WorkspaceFolders;
}

/// <summary>
/// Defines the capabilities provided by a language
/// server.
/// </summary>
public class ServerCapabilities
{
    /// <summary>
    /// The position encoding the server picked from the encodings offered
    /// by the client via the client capability `general.positionEncodings`.
    /// 
    /// If the client didn't provide any position encodings the only valid
    /// value that a server can return is 'utf-16'.
    /// 
    /// If omitted it defaults to 'utf-16'.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*//*PositionEncodingKind*/string PositionEncoding;
    /// <summary>
    /// Defines how text documents are synced. Is either a detailed structure
    /// defining each notification or for backwards compatibility the
    /// TextDocumentSyncKind number.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/TextDocumentSyncOptionsOrTextDocumentSyncKind TextDocumentSync;
    /// <summary>
    /// Defines how notebook documents are synced.
    /// 
    /// @since 3.17.0
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/NotebookDocumentSyncOptionsOrNotebookDocumentSyncRegistrationOptions NotebookDocumentSync;
    /// <summary>
    /// The server provides completion support.
    /// </summary>
    public /*TOpt*/CompletionOptions? CompletionProvider;
    /// <summary>
    /// The server provides hover support.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrHoverOptions HoverProvider;
    /// <summary>
    /// The server provides signature help support.
    /// </summary>
    public /*TOpt*/SignatureHelpOptions? SignatureHelpProvider;
    /// <summary>
    /// The server provides Goto Declaration support.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrDeclarationOptionsOrDeclarationRegistrationOptions DeclarationProvider;
    /// <summary>
    /// The server provides goto definition support.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrDefinitionOptions DefinitionProvider;
    /// <summary>
    /// The server provides Goto Type Definition support.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrTypeDefinitionOptionsOrTypeDefinitionRegistrationOptions TypeDefinitionProvider;
    /// <summary>
    /// The server provides Goto Implementation support.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrImplementationOptionsOrImplementationRegistrationOptions ImplementationProvider;
    /// <summary>
    /// The server provides find references support.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrReferenceOptions ReferencesProvider;
    /// <summary>
    /// The server provides document highlight support.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrDocumentHighlightOptions DocumentHighlightProvider;
    /// <summary>
    /// The server provides document symbol support.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrDocumentSymbolOptions DocumentSymbolProvider;
    /// <summary>
    /// The server provides code actions. CodeActionOptions may only be
    /// specified if the client states that it supports
    /// `codeActionLiteralSupport` in its initial `initialize` request.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrCodeActionOptions CodeActionProvider;
    /// <summary>
    /// The server provides code lens.
    /// </summary>
    public /*TOpt*/CodeLensOptions? CodeLensProvider;
    /// <summary>
    /// The server provides document link support.
    /// </summary>
    public /*TOpt*/DocumentLinkOptions? DocumentLinkProvider;
    /// <summary>
    /// The server provides color provider support.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrDocumentColorOptionsOrDocumentColorRegistrationOptions ColorProvider;
    /// <summary>
    /// The server provides workspace symbol support.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrWorkspaceSymbolOptions WorkspaceSymbolProvider;
    /// <summary>
    /// The server provides document formatting.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrDocumentFormattingOptions DocumentFormattingProvider;
    /// <summary>
    /// The server provides document range formatting.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrDocumentRangeFormattingOptions DocumentRangeFormattingProvider;
    /// <summary>
    /// The server provides document formatting on typing.
    /// </summary>
    public /*TOpt*/DocumentOnTypeFormattingOptions? DocumentOnTypeFormattingProvider;
    /// <summary>
    /// The server provides rename support. RenameOptions may only be
    /// specified if the client states that it supports
    /// `prepareSupport` in its initial `initialize` request.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrRenameOptions RenameProvider;
    /// <summary>
    /// The server provides folding provider support.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrFoldingRangeOptionsOrFoldingRangeRegistrationOptions FoldingRangeProvider;
    /// <summary>
    /// The server provides selection range support.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrSelectionRangeOptionsOrSelectionRangeRegistrationOptions SelectionRangeProvider;
    /// <summary>
    /// The server provides execute command support.
    /// </summary>
    public /*TOpt*/ExecuteCommandOptions? ExecuteCommandProvider;
    /// <summary>
    /// The server provides call hierarchy support.
    /// 
    /// @since 3.16.0
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrCallHierarchyOptionsOrCallHierarchyRegistrationOptions CallHierarchyProvider;
    /// <summary>
    /// The server provides linked editing range support.
    /// 
    /// @since 3.16.0
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrLinkedEditingRangeOptionsOrLinkedEditingRangeRegistrationOptions LinkedEditingRangeProvider;
    /// <summary>
    /// The server provides semantic tokens support.
    /// 
    /// @since 3.16.0
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/SemanticTokensOptionsOrSemanticTokensRegistrationOptions SemanticTokensProvider;
    /// <summary>
    /// The server provides moniker support.
    /// 
    /// @since 3.16.0
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrMonikerOptionsOrMonikerRegistrationOptions MonikerProvider;
    /// <summary>
    /// The server provides type hierarchy support.
    /// 
    /// @since 3.17.0
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrTypeHierarchyOptionsOrTypeHierarchyRegistrationOptions TypeHierarchyProvider;
    /// <summary>
    /// The server provides inline values.
    /// 
    /// @since 3.17.0
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrInlineValueOptionsOrInlineValueRegistrationOptions InlineValueProvider;
    /// <summary>
    /// The server provides inlay hints.
    /// 
    /// @since 3.17.0
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrInlayHintOptionsOrInlayHintRegistrationOptions InlayHintProvider;
    /// <summary>
    /// The server has support for pull model diagnostics.
    /// 
    /// @since 3.17.0
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/DiagnosticOptionsOrDiagnosticRegistrationOptions DiagnosticProvider;
    /// <summary>
    /// Workspace specific server capabilities.
    /// </summary>
    public /*TOpt*/WorkspaceFoldersWorkspaceFoldersServerCapabilitiesWithFileOperationsFileOperationOptions Workspace;
    /// <summary>
    /// Experimental server capabilities.
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Experimental;
}

/// <summary>
/// A text document identifier to denote a specific version of a text document.
/// </summary>
public class VersionedTextDocumentIdentifier
{
    public TextDocumentIdentifier TextDocumentIdentifier; //XTEND
    /// <summary>
    /// The version number of this document.
    /// </summary>
    public System.Int64 Version;
}

/// <summary>
/// Save options.
/// </summary>
public class SaveOptions
{
    /// <summary>
    /// The client is supposed to include the content on save.
    /// </summary>
    public /*TOpt*/System.Boolean? IncludeText;
}

/// <summary>
/// An event describing a file change.
/// </summary>
public class FileEvent
{
    /// <summary>
    /// The file's uri.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// The change type.
    /// </summary>
    public FileChangeType Type;
}

public class FileSystemWatcher
{
    /// <summary>
    /// The glob pattern to watch. See <c>GlobPattern</c> for more detail.
    /// 
    /// @since 3.17.0 support for relative patterns.
    /// </summary>
    public GlobPattern GlobPattern;
    /// <summary>
    /// The kind of events of interest. If omitted it defaults
    /// to WatchKind.Create | WatchKind.Change | WatchKind.Delete
    /// which is 7.
    /// </summary>
    public /*TOpt*/WatchKind Kind;
}

/// <summary>
/// Represents a diagnostic, such as a compiler error or warning. Diagnostic objects
/// are only valid in the scope of a resource.
/// </summary>
public class Diagnostic
{
    /// <summary>
    /// The range at which the message applies
    /// </summary>
    public Range Range;
    /// <summary>
    /// The diagnostic's severity. Can be omitted. If omitted it is up to the
    /// client to interpret diagnostics as error, warning, info or hint.
    /// </summary>
    public /*TOpt*/DiagnosticSeverity Severity;
    /// <summary>
    /// The diagnostic's code, which usually appear in the user interface.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/IntegerOrString Code;
    /// <summary>
    /// An optional property to describe the error code.
    /// Requires the code field (above) to be present/not null.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/CodeDescription? CodeDescription;
    /// <summary>
    /// A human-readable string describing the source of this
    /// diagnostic, e.g. 'typescript' or 'super lint'. It usually
    /// appears in the user interface.
    /// </summary>
    public /*TOpt*/System.String? Source;
    /// <summary>
    /// The diagnostic's message. It usually appears in the user interface
    /// </summary>
    public System.String Message;
    /// <summary>
    /// Additional metadata about the diagnostic.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/DiagnosticTag[] Tags;
    /// <summary>
    /// An array of related diagnostic information, e.g. when symbol-names within
    /// a scope collide all definitions can be marked via this property.
    /// </summary>
    public /*TOpt*/DiagnosticRelatedInformation[] RelatedInformation;
    /// <summary>
    /// A data entry field that is preserved between a `textDocument/publishDiagnostics`
    /// notification and `textDocument/codeAction` request.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Data;
}

/// <summary>
/// Contains additional information about the context in which a completion request is triggered.
/// </summary>
public class CompletionContext
{
    /// <summary>
    /// How the completion was triggered.
    /// </summary>
    public CompletionTriggerKind TriggerKind;
    /// <summary>
    /// The trigger character (a single character) that has trigger code complete.
    /// Is undefined if `triggerKind !== CompletionTriggerKind.TriggerCharacter`
    /// </summary>
    public /*TOpt*/System.String? TriggerCharacter;
}

/// <summary>
/// Additional details for a completion item label.
/// 
/// @since 3.17.0
/// </summary>
public class CompletionItemLabelDetails
{
    /// <summary>
    /// An optional string which is rendered less prominently directly after <c>CompletionItem.label</c>,
    /// without any spacing. Should be used for function signatures and type annotations.
    /// </summary>
    public /*TOpt*/System.String? Detail;
    /// <summary>
    /// An optional string which is rendered less prominently after <c>CompletionItem.detail</c>. Should be used
    /// for fully qualified names and file paths.
    /// </summary>
    public /*TOpt*/System.String? Description;
}

/// <summary>
/// A special text edit to provide an insert and a replace operation.
/// 
/// @since 3.16.0
/// </summary>
public class InsertReplaceEdit
{
    /// <summary>
    /// The string to be inserted.
    /// </summary>
    public System.String NewText;
    /// <summary>
    /// The range if the insert is requested
    /// </summary>
    public Range Insert;
    /// <summary>
    /// The range if the replace is requested.
    /// </summary>
    public Range Replace;
}

/// <summary>
/// Completion options.
/// </summary>
public class CompletionOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
    /// <summary>
    /// Most tools trigger completion request automatically without explicitly requesting
    /// it using a keyboard shortcut (e.g. Ctrl+Space). Typically they do so when the user
    /// starts to type an identifier. For example if the user types `c` in a JavaScript file
    /// code complete will automatically pop up present `console` besides others as a
    /// completion item. Characters that make up identifiers don't need to be listed here.
    /// 
    /// If code complete should automatically be trigger on characters not being valid inside
    /// an identifier (for example `.` in JavaScript) list them in `triggerCharacters`.
    /// </summary>
    public /*TOpt*/System.String[] TriggerCharacters;
    /// <summary>
    /// The list of all possible characters that commit a completion. This field can be used
    /// if clients don't support individual commit characters per completion item. See
    /// `ClientCapabilities.textDocument.completion.completionItem.commitCharactersSupport`
    /// 
    /// If a server provides both `allCommitCharacters` and commit characters on an individual
    /// completion item the ones on the completion item win.
    /// 
    /// @since 3.2.0
    /// </summary>
    public /*TOpt*/System.String[] AllCommitCharacters;
    /// <summary>
    /// The server provides support to resolve additional
    /// information for a completion item.
    /// </summary>
    public /*TOpt*/System.Boolean? ResolveProvider;
    /// <summary>
    /// The server supports the following `CompletionItem` specific
    /// capabilities.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/LabelDetailsSupportBoolean CompletionItem;
}

/// <summary>
/// Hover options.
/// </summary>
public class HoverOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// Additional information about the context in which a signature help request was triggered.
/// 
/// @since 3.15.0
/// </summary>
public class SignatureHelpContext
{
    /// <summary>
    /// Action that caused signature help to be triggered.
    /// </summary>
    public SignatureHelpTriggerKind TriggerKind;
    /// <summary>
    /// Character that caused signature help to be triggered.
    /// 
    /// This is undefined when `triggerKind !== SignatureHelpTriggerKind.TriggerCharacter`
    /// </summary>
    public /*TOpt*/System.String? TriggerCharacter;
    /// <summary>
    /// `true` if signature help was already showing when it was triggered.
    /// 
    /// Retriggers occurs when the signature help is already active and can be caused by actions such as
    /// typing a trigger character, a cursor move, or document content changes.
    /// </summary>
    public System.Boolean IsRetrigger;
    /// <summary>
    /// The currently active `SignatureHelp`.
    /// 
    /// The `activeSignatureHelp` has its `SignatureHelp.activeSignature` field updated based on
    /// the user navigating through available signatures.
    /// </summary>
    public /*TOpt*/SignatureHelp? ActiveSignatureHelp;
}

/// <summary>
/// Represents the signature of something callable. A signature
/// can have a label, like a function-name, a doc-comment, and
/// a set of parameters.
/// </summary>
public class SignatureInformation
{
    /// <summary>
    /// The label of this signature. Will be shown in
    /// the UI.
    /// </summary>
    public System.String Label;
    /// <summary>
    /// The human-readable doc-comment of this signature. Will be shown
    /// in the UI but can be omitted.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/StringOrMarkupContent Documentation;
    /// <summary>
    /// The parameters of this signature.
    /// </summary>
    public /*TOpt*/ParameterInformation[] Parameters;
    /// <summary>
    /// The index of the active parameter.
    /// 
    /// If provided, this is used in place of `SignatureHelp.activeParameter`.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.UInt64? ActiveParameter;
}

/// <summary>
/// Server Capabilities for a <c>SignatureHelpRequest</c>.
/// </summary>
public class SignatureHelpOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
    /// <summary>
    /// List of characters that trigger signature help automatically.
    /// </summary>
    public /*TOpt*/System.String[] TriggerCharacters;
    /// <summary>
    /// List of characters that re-trigger signature help.
    /// 
    /// These trigger characters are only active when signature help is already showing. All trigger characters
    /// are also counted as re-trigger characters.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/System.String[] RetriggerCharacters;
}

/// <summary>
/// Server Capabilities for a <c>DefinitionRequest</c>.
/// </summary>
public class DefinitionOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// Value-object that contains additional information when
/// requesting references.
/// </summary>
public class ReferenceContext
{
    /// <summary>
    /// Include the declaration of the current symbol.
    /// </summary>
    public System.Boolean IncludeDeclaration;
}

/// <summary>
/// Reference options.
/// </summary>
public class ReferenceOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// Provider options for a <c>DocumentHighlightRequest</c>.
/// </summary>
public class DocumentHighlightOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// A base for all symbol information.
/// </summary>
public class BaseSymbolInformation
{
    /// <summary>
    /// The name of this symbol.
    /// </summary>
    public System.String Name;
    /// <summary>
    /// The kind of this symbol.
    /// </summary>
    public SymbolKind Kind;
    /// <summary>
    /// Tags for this symbol.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/SymbolTag[] Tags;
    /// <summary>
    /// The name of the symbol containing this symbol. This information is for
    /// user interface purposes (e.g. to render a qualifier in the user interface
    /// if necessary). It can't be used to re-infer a hierarchy for the document
    /// symbols.
    /// </summary>
    public /*TOpt*/System.String? ContainerName;
}

/// <summary>
/// Provider options for a <c>DocumentSymbolRequest</c>.
/// </summary>
public class DocumentSymbolOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
    /// <summary>
    /// A human-readable string that is shown when multiple outlines trees
    /// are shown for the same document.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.String? Label;
}

/// <summary>
/// Contains additional diagnostic information about the context in which
/// a <c>CodeActionProvider.provideCodeActions</c> is run.
/// </summary>
public class CodeActionContext
{
    /// <summary>
    /// An array of diagnostics known on the client side overlapping the range provided to the
    /// `textDocument/codeAction` request. They are provided so that the server knows which
    /// errors are currently presented to the user for the given range. There is no guarantee
    /// that these accurately reflect the error state of the resource. The primary parameter
    /// to compute code actions is the provided range.
    /// </summary>
    public Diagnostic[] Diagnostics;
    /// <summary>
    /// Requested kind of actions to return.
    /// 
    /// Actions not of this kind are filtered out by the client before being shown. So servers
    /// can omit computing them.
    /// </summary>
    public /*TOpt*//*CodeActionKind*/string[] Only;
    /// <summary>
    /// The reason why code actions were requested.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/CodeActionTriggerKind TriggerKind;
}

/// <summary>
/// Provider options for a <c>CodeActionRequest</c>.
/// </summary>
public class CodeActionOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
    /// <summary>
    /// CodeActionKinds that this server may return.
    /// 
    /// The list of kinds may be generic, such as `CodeActionKind.Refactor`, or the server
    /// may list out every specific kind they provide.
    /// </summary>
    public /*TOpt*//*CodeActionKind*/string[] CodeActionKinds;
    /// <summary>
    /// The server provides support to resolve additional
    /// information for a code action.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.Boolean? ResolveProvider;
}

/// <summary>
/// Server capabilities for a <c>WorkspaceSymbolRequest</c>.
/// </summary>
public class WorkspaceSymbolOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
    /// <summary>
    /// The server provides support to resolve additional
    /// information for a workspace symbol.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/System.Boolean? ResolveProvider;
}

/// <summary>
/// Code Lens provider options of a <c>CodeLensRequest</c>.
/// </summary>
public class CodeLensOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
    /// <summary>
    /// Code lens has a resolve provider as well.
    /// </summary>
    public /*TOpt*/System.Boolean? ResolveProvider;
}

/// <summary>
/// Provider options for a <c>DocumentLinkRequest</c>.
/// </summary>
public class DocumentLinkOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
    /// <summary>
    /// Document links have a resolve provider as well.
    /// </summary>
    public /*TOpt*/System.Boolean? ResolveProvider;
}

/// <summary>
/// Value-object describing what options formatting should use.
/// </summary>
public class FormattingOptions
{
    /// <summary>
    /// Size of a tab in spaces.
    /// </summary>
    public System.UInt64 TabSize;
    /// <summary>
    /// Prefer spaces over tabs.
    /// </summary>
    public System.Boolean InsertSpaces;
    /// <summary>
    /// Trim trailing whitespace on a line.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/System.Boolean? TrimTrailingWhitespace;
    /// <summary>
    /// Insert a newline character at the end of the file if one does not exist.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/System.Boolean? InsertFinalNewline;
    /// <summary>
    /// Trim all newlines after the final newline at the end of the file.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/System.Boolean? TrimFinalNewlines;
}

/// <summary>
/// Provider options for a <c>DocumentFormattingRequest</c>.
/// </summary>
public class DocumentFormattingOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// Provider options for a <c>DocumentRangeFormattingRequest</c>.
/// </summary>
public class DocumentRangeFormattingOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
}

/// <summary>
/// Provider options for a <c>DocumentOnTypeFormattingRequest</c>.
/// </summary>
public class DocumentOnTypeFormattingOptions
{
    /// <summary>
    /// A character on which formatting should be triggered, like `{`.
    /// </summary>
    public System.String FirstTriggerCharacter;
    /// <summary>
    /// More trigger characters.
    /// </summary>
    public /*TOpt*/System.String[] MoreTriggerCharacter;
}

/// <summary>
/// Provider options for a <c>RenameRequest</c>.
/// </summary>
public class RenameOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
    /// <summary>
    /// Renames should be checked and tested before being executed.
    /// 
    /// @since version 3.12.0
    /// </summary>
    public /*TOpt*/System.Boolean? PrepareProvider;
}

/// <summary>
/// The server capabilities of a <c>ExecuteCommandRequest</c>.
/// </summary>
public class ExecuteCommandOptions
{
    public WorkDoneProgressOptions WorkDoneProgressOptions; //MIXIN
    /// <summary>
    /// The commands to be executed on the server
    /// </summary>
    public System.String[] Commands;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokensLegend
{
    /// <summary>
    /// The token types a server uses.
    /// </summary>
    public System.String[] TokenTypes;
    /// <summary>
    /// The token modifiers a server uses.
    /// </summary>
    public System.String[] TokenModifiers;
}

/// <summary>
/// A text document identifier to optionally denote a specific version of a text document.
/// </summary>
public class OptionalVersionedTextDocumentIdentifier
{
    public TextDocumentIdentifier TextDocumentIdentifier; //XTEND
    /// <summary>
    /// The version number of this document. If a versioned text document identifier
    /// is sent from the server to the client and the file is not open in the editor
    /// (the server has not received an open notification before) the server can send
    /// `null` to indicate that the version is unknown and the content on disk is the
    /// truth (as specified with document content ownership).
    /// </summary>
    public /*TOr1*//*TOpt*/System.Int64? Version;
}

/// <summary>
/// A special text edit with an additional change annotation.
/// 
/// @since 3.16.0.
/// </summary>
public class AnnotatedTextEdit
{
    public TextEdit TextEdit; //XTEND
    /// <summary>
    /// The actual identifier of the change annotation
    /// </summary>
    public ChangeAnnotationIdentifier AnnotationId;
}

/// <summary>
/// A generic resource operation.
/// </summary>
public class ResourceOperation
{
    /// <summary>
    /// The resource operation kind.
    /// </summary>
    public System.String Kind;
    /// <summary>
    /// An optional annotation identifier describing the operation.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.String? AnnotationId;
}

/// <summary>
/// Options to create a file.
/// </summary>
public class CreateFileOptions
{
    /// <summary>
    /// Overwrite existing file. Overwrite wins over `ignoreIfExists`
    /// </summary>
    public /*TOpt*/System.Boolean? Overwrite;
    /// <summary>
    /// Ignore if exists.
    /// </summary>
    public /*TOpt*/System.Boolean? IgnoreIfExists;
}

/// <summary>
/// Rename file options
/// </summary>
public class RenameFileOptions
{
    /// <summary>
    /// Overwrite target if existing. Overwrite wins over `ignoreIfExists`
    /// </summary>
    public /*TOpt*/System.Boolean? Overwrite;
    /// <summary>
    /// Ignores if target exists.
    /// </summary>
    public /*TOpt*/System.Boolean? IgnoreIfExists;
}

/// <summary>
/// Delete file options
/// </summary>
public class DeleteFileOptions
{
    /// <summary>
    /// Delete the content recursively if a folder is denoted.
    /// </summary>
    public /*TOpt*/System.Boolean? Recursive;
    /// <summary>
    /// Ignore the operation if the file doesn't exist.
    /// </summary>
    public /*TOpt*/System.Boolean? IgnoreIfNotExists;
}

/// <summary>
/// A pattern to describe in which file operation requests or notifications
/// the server is interested in receiving.
/// 
/// @since 3.16.0
/// </summary>
public class FileOperationPattern
{
    /// <summary>
    /// The glob pattern to match. Glob patterns can have the following syntax:
    /// - `*` to match one or more characters in a path segment
    /// - `?` to match on one character in a path segment
    /// - `**` to match any number of path segments, including none
    /// - `{}` to group sub patterns into an OR expression. (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
    /// - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
    /// - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
    /// </summary>
    public System.String Glob;
    /// <summary>
    /// Whether to match files or folders with this pattern.
    /// 
    /// Matches both if undefined.
    /// </summary>
    public /*TOpt*//*FileOperationPatternKind*/string Matches;
    /// <summary>
    /// Additional options used during matching.
    /// </summary>
    public /*TOpt*/FileOperationPatternOptions? Options;
}

/// <summary>
/// A full document diagnostic report for a workspace diagnostic result.
/// 
/// @since 3.17.0
/// </summary>
public class WorkspaceFullDocumentDiagnosticReport
{
    public FullDocumentDiagnosticReport FullDocumentDiagnosticReport; //XTEND
    /// <summary>
    /// The URI for which diagnostic information is reported.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// The version number for which the diagnostics are reported.
    /// If the document is not marked as open `null` can be provided.
    /// </summary>
    public /*TOr1*//*TOpt*/System.Int64? Version;
}

/// <summary>
/// An unchanged document diagnostic report for a workspace diagnostic result.
/// 
/// @since 3.17.0
/// </summary>
public class WorkspaceUnchangedDocumentDiagnosticReport
{
    public UnchangedDocumentDiagnosticReport UnchangedDocumentDiagnosticReport; //XTEND
    /// <summary>
    /// The URI for which diagnostic information is reported.
    /// </summary>
    public System.String Uri;
    /// <summary>
    /// The version number for which the diagnostics are reported.
    /// If the document is not marked as open `null` can be provided.
    /// </summary>
    public /*TOr1*//*TOpt*/System.Int64? Version;
}

/// <summary>
/// A notebook cell.
/// 
/// A cell's document URI must be unique across ALL notebook
/// cells and can therefore be used to uniquely identify a
/// notebook cell or the cell's text document.
/// 
/// @since 3.17.0
/// </summary>
public class NotebookCell
{
    /// <summary>
    /// The cell's kind
    /// </summary>
    public NotebookCellKind Kind;
    /// <summary>
    /// The URI of the cell's text document
    /// content.
    /// </summary>
    public System.String Document;
    /// <summary>
    /// Additional metadata stored with the cell.
    /// 
    /// Note: should always be an object literal (e.g. LSPObject)
    /// </summary>
    public /*TOpt*/LSPObject Metadata;
    /// <summary>
    /// Additional execution summary information
    /// if supported by the client.
    /// </summary>
    public /*TOpt*/ExecutionSummary? ExecutionSummary;
}

/// <summary>
/// A change describing how to move a `NotebookCell`
/// array from state S to S'.
/// 
/// @since 3.17.0
/// </summary>
public class NotebookCellArrayChange
{
    /// <summary>
    /// The start oftest of the cell that changed.
    /// </summary>
    public System.UInt64 Start;
    /// <summary>
    /// The deleted cells
    /// </summary>
    public System.UInt64 DeleteCount;
    /// <summary>
    /// The new cells, if any
    /// </summary>
    public /*TOpt*/NotebookCell[] Cells;
}

/// <summary>
/// Defines the capabilities provided by the client.
/// </summary>
public class ClientCapabilities
{
    /// <summary>
    /// Workspace specific client capabilities.
    /// </summary>
    public /*TOpt*/WorkspaceClientCapabilities? Workspace;
    /// <summary>
    /// Text document specific client capabilities.
    /// </summary>
    public /*TOpt*/TextDocumentClientCapabilities? TextDocument;
    /// <summary>
    /// Capabilities specific to the notebook document support.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/NotebookDocumentClientCapabilities? NotebookDocument;
    /// <summary>
    /// Window specific client capabilities.
    /// </summary>
    public /*TOpt*/WindowClientCapabilities? Window;
    /// <summary>
    /// General client capabilities.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/GeneralClientCapabilities? General;
    /// <summary>
    /// Experimental client capabilities.
    /// </summary>
    public /*TOpt*/LSPObjectOrLSPArrayOrStringOrIntegerOrUintegerOrDecimalOrBooleanOrNull Experimental;
}

public class TextDocumentSyncOptions
{
    /// <summary>
    /// Open and close notifications are sent to the server. If omitted open close notification should not
    /// be sent.
    /// </summary>
    public /*TOpt*/System.Boolean? OpenClose;
    /// <summary>
    /// Change notifications are sent to the server. See TextDocumentSyncKind.None, TextDocumentSyncKind.Full
    /// and TextDocumentSyncKind.Incremental. If omitted it defaults to TextDocumentSyncKind.None.
    /// </summary>
    public /*TOpt*/TextDocumentSyncKind Change;
    /// <summary>
    /// If present will save notifications are sent to the server. If omitted the notification should not be
    /// sent.
    /// </summary>
    public /*TOpt*/System.Boolean? WillSave;
    /// <summary>
    /// If present will save wait until requests are sent to the server. If omitted the request should not be
    /// sent.
    /// </summary>
    public /*TOpt*/System.Boolean? WillSaveWaitUntil;
    /// <summary>
    /// If present save notifications are sent to the server. If omitted the notification should not be
    /// sent.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/BooleanOrSaveOptions Save;
}

/// <summary>
/// Options specific to a notebook plus its cells
/// to be synced to the server.
/// 
/// If a selector provides a notebook document
/// filter but no cell selector all cells of a
/// matching notebook document will be synced.
/// 
/// If a selector provides no notebook document
/// filter but only a cell selector all notebook
/// document that contain at least one matching
/// cell will be synced.
/// 
/// @since 3.17.0
/// </summary>
public class NotebookDocumentSyncOptions
{
    /// <summary>
    /// The notebooks to be synced
    /// </summary>
    public NotebookStringOrNotebookDocumentFilterWithCellsLanguageStrings[] NotebookSelector;
    /// <summary>
    /// Whether save notification should be forwarded to
    /// the server. Will only be honored if mode === `notebook`.
    /// </summary>
    public /*TOpt*/System.Boolean? Save;
}

/// <summary>
/// Registration options specific to a notebook.
/// 
/// @since 3.17.0
/// </summary>
public class NotebookDocumentSyncRegistrationOptions
{
    public NotebookDocumentSyncOptions NotebookDocumentSyncOptions; //XTEND
    public StaticRegistrationOptions StaticRegistrationOptions; //MIXIN
}

public class WorkspaceFoldersServerCapabilities
{
    /// <summary>
    /// The server has support for workspace folders
    /// </summary>
    public /*TOpt*/System.Boolean? Supported;
    /// <summary>
    /// Whether the server wants to receive workspace folder
    /// change notifications.
    /// 
    /// If a string is provided the string is treated as an ID
    /// under which the notification is registered on the client
    /// side. The ID can be used to unregister for these events
    /// using the `client/unregisterCapability` request.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/StringOrBoolean ChangeNotifications;
}

/// <summary>
/// Options for notifications/requests for user operations on files.
/// 
/// @since 3.16.0
/// </summary>
public class FileOperationOptions
{
    /// <summary>
    /// The server is interested in receiving didCreateFiles notifications.
    /// </summary>
    public /*TOpt*/FileOperationRegistrationOptions? DidCreate;
    /// <summary>
    /// The server is interested in receiving willCreateFiles requests.
    /// </summary>
    public /*TOpt*/FileOperationRegistrationOptions? WillCreate;
    /// <summary>
    /// The server is interested in receiving didRenameFiles notifications.
    /// </summary>
    public /*TOpt*/FileOperationRegistrationOptions? DidRename;
    /// <summary>
    /// The server is interested in receiving willRenameFiles requests.
    /// </summary>
    public /*TOpt*/FileOperationRegistrationOptions? WillRename;
    /// <summary>
    /// The server is interested in receiving didDeleteFiles file notifications.
    /// </summary>
    public /*TOpt*/FileOperationRegistrationOptions? DidDelete;
    /// <summary>
    /// The server is interested in receiving willDeleteFiles file requests.
    /// </summary>
    public /*TOpt*/FileOperationRegistrationOptions? WillDelete;
}

/// <summary>
/// Structure to capture a description for an error code.
/// 
/// @since 3.16.0
/// </summary>
public class CodeDescription
{
    /// <summary>
    /// An URI to open with more information about the diagnostic error.
    /// </summary>
    public System.String Href;
}

/// <summary>
/// Represents a related message and source code location for a diagnostic. This should be
/// used to point to code locations that cause or related to a diagnostics, e.g when duplicating
/// a symbol in a scope.
/// </summary>
public class DiagnosticRelatedInformation
{
    /// <summary>
    /// The location of this related diagnostic information.
    /// </summary>
    public Location Location;
    /// <summary>
    /// The message of this related diagnostic information.
    /// </summary>
    public System.String Message;
}

/// <summary>
/// Represents a parameter of a callable-signature. A parameter can
/// have a label and a doc-comment.
/// </summary>
public class ParameterInformation
{
    /// <summary>
    /// The label of this parameter information.
    /// 
    /// Either a string or an inclusive start and exclusive end offsets within its containing
    /// signature label. (see SignatureInformation.label). The offsets are based on a UTF-16
    /// string representation as `Position` and `Range` does.
    /// 
    /// *Note*: a label of type string should be a substring of its containing signature label.
    /// Its intended use case is to highlight the parameter label part in the `SignatureInformation.label`.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public StringOrUintegerWithUinteger Label;
    /// <summary>
    /// The human-readable doc-comment of this parameter. Will be shown
    /// in the UI but can be omitted.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public /*TOpt*/StringOrMarkupContent Documentation;
}

/// <summary>
/// A notebook cell text document filter denotes a cell text
/// document by different properties.
/// 
/// @since 3.17.0
/// </summary>
public class NotebookCellTextDocumentFilter
{
    /// <summary>
    /// A filter that matches against the notebook
    /// containing the notebook cell. If a string
    /// value is provided it matches against the
    /// notebook type. '*' matches every notebook.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public StringOrNotebookDocumentFilter Notebook;
    /// <summary>
    /// A language id like `python`.
    /// 
    /// Will be matched against the language id of the
    /// notebook cell document. '*' matches every language.
    /// </summary>
    public /*TOpt*/System.String? Language;
}

/// <summary>
/// Matching options for the file operation pattern.
/// 
/// @since 3.16.0
/// </summary>
public class FileOperationPatternOptions
{
    /// <summary>
    /// The pattern should be matched ignoring casing.
    /// </summary>
    public /*TOpt*/System.Boolean? IgnoreCase;
}

public class ExecutionSummary
{
    /// <summary>
    /// A strict monotonically increasing value
    /// indicating the execution order of a cell
    /// inside a notebook.
    /// </summary>
    public System.UInt64 ExecutionOrder;
    /// <summary>
    /// Whether the execution was successful or
    /// not if known by the client.
    /// </summary>
    public /*TOpt*/System.Boolean? Success;
}

/// <summary>
/// Workspace specific client capabilities.
/// </summary>
public class WorkspaceClientCapabilities
{
    /// <summary>
    /// The client supports applying batch edits
    /// to the workspace by supporting the request
    /// 'workspace/applyEdit'
    /// </summary>
    public /*TOpt*/System.Boolean? ApplyEdit;
    /// <summary>
    /// Capabilities specific to `WorkspaceEdit`s.
    /// </summary>
    public /*TOpt*/WorkspaceEditClientCapabilities? WorkspaceEdit;
    /// <summary>
    /// Capabilities specific to the `workspace/didChangeConfiguration` notification.
    /// </summary>
    public /*TOpt*/DidChangeConfigurationClientCapabilities? DidChangeConfiguration;
    /// <summary>
    /// Capabilities specific to the `workspace/didChangeWatchedFiles` notification.
    /// </summary>
    public /*TOpt*/DidChangeWatchedFilesClientCapabilities? DidChangeWatchedFiles;
    /// <summary>
    /// Capabilities specific to the `workspace/symbol` request.
    /// </summary>
    public /*TOpt*/WorkspaceSymbolClientCapabilities? Symbol;
    /// <summary>
    /// Capabilities specific to the `workspace/executeCommand` request.
    /// </summary>
    public /*TOpt*/ExecuteCommandClientCapabilities? ExecuteCommand;
    /// <summary>
    /// The client has support for workspace folders.
    /// 
    /// @since 3.6.0
    /// </summary>
    public /*TOpt*/System.Boolean? WorkspaceFolders;
    /// <summary>
    /// The client supports `workspace/configuration` requests.
    /// 
    /// @since 3.6.0
    /// </summary>
    public /*TOpt*/System.Boolean? Configuration;
    /// <summary>
    /// Capabilities specific to the semantic token requests scoped to the
    /// workspace.
    /// 
    /// @since 3.16.0.
    /// </summary>
    public /*TOpt*/SemanticTokensWorkspaceClientCapabilities? SemanticTokens;
    /// <summary>
    /// Capabilities specific to the code lens requests scoped to the
    /// workspace.
    /// 
    /// @since 3.16.0.
    /// </summary>
    public /*TOpt*/CodeLensWorkspaceClientCapabilities? CodeLens;
    /// <summary>
    /// The client has support for file notifications/requests for user operations on files.
    /// 
    /// Since 3.16.0
    /// </summary>
    public /*TOpt*/FileOperationClientCapabilities? FileOperations;
    /// <summary>
    /// Capabilities specific to the inline values requests scoped to the
    /// workspace.
    /// 
    /// @since 3.17.0.
    /// </summary>
    public /*TOpt*/InlineValueWorkspaceClientCapabilities? InlineValue;
    /// <summary>
    /// Capabilities specific to the inlay hint requests scoped to the
    /// workspace.
    /// 
    /// @since 3.17.0.
    /// </summary>
    public /*TOpt*/InlayHintWorkspaceClientCapabilities? InlayHint;
    /// <summary>
    /// Capabilities specific to the diagnostic requests scoped to the
    /// workspace.
    /// 
    /// @since 3.17.0.
    /// </summary>
    public /*TOpt*/DiagnosticWorkspaceClientCapabilities? Diagnostics;
}

/// <summary>
/// Text document specific client capabilities.
/// </summary>
public class TextDocumentClientCapabilities
{
    /// <summary>
    /// Defines which synchronization capabilities the client supports.
    /// </summary>
    public /*TOpt*/TextDocumentSyncClientCapabilities? Synchronization;
    /// <summary>
    /// Capabilities specific to the `textDocument/completion` request.
    /// </summary>
    public /*TOpt*/CompletionClientCapabilities? Completion;
    /// <summary>
    /// Capabilities specific to the `textDocument/hover` request.
    /// </summary>
    public /*TOpt*/HoverClientCapabilities? Hover;
    /// <summary>
    /// Capabilities specific to the `textDocument/signatureHelp` request.
    /// </summary>
    public /*TOpt*/SignatureHelpClientCapabilities? SignatureHelp;
    /// <summary>
    /// Capabilities specific to the `textDocument/declaration` request.
    /// 
    /// @since 3.14.0
    /// </summary>
    public /*TOpt*/DeclarationClientCapabilities? Declaration;
    /// <summary>
    /// Capabilities specific to the `textDocument/definition` request.
    /// </summary>
    public /*TOpt*/DefinitionClientCapabilities? Definition;
    /// <summary>
    /// Capabilities specific to the `textDocument/typeDefinition` request.
    /// 
    /// @since 3.6.0
    /// </summary>
    public /*TOpt*/TypeDefinitionClientCapabilities? TypeDefinition;
    /// <summary>
    /// Capabilities specific to the `textDocument/implementation` request.
    /// 
    /// @since 3.6.0
    /// </summary>
    public /*TOpt*/ImplementationClientCapabilities? Implementation;
    /// <summary>
    /// Capabilities specific to the `textDocument/references` request.
    /// </summary>
    public /*TOpt*/ReferenceClientCapabilities? References;
    /// <summary>
    /// Capabilities specific to the `textDocument/documentHighlight` request.
    /// </summary>
    public /*TOpt*/DocumentHighlightClientCapabilities? DocumentHighlight;
    /// <summary>
    /// Capabilities specific to the `textDocument/documentSymbol` request.
    /// </summary>
    public /*TOpt*/DocumentSymbolClientCapabilities? DocumentSymbol;
    /// <summary>
    /// Capabilities specific to the `textDocument/codeAction` request.
    /// </summary>
    public /*TOpt*/CodeActionClientCapabilities? CodeAction;
    /// <summary>
    /// Capabilities specific to the `textDocument/codeLens` request.
    /// </summary>
    public /*TOpt*/CodeLensClientCapabilities? CodeLens;
    /// <summary>
    /// Capabilities specific to the `textDocument/documentLink` request.
    /// </summary>
    public /*TOpt*/DocumentLinkClientCapabilities? DocumentLink;
    /// <summary>
    /// Capabilities specific to the `textDocument/documentColor` and the
    /// `textDocument/colorPresentation` request.
    /// 
    /// @since 3.6.0
    /// </summary>
    public /*TOpt*/DocumentColorClientCapabilities? ColorProvider;
    /// <summary>
    /// Capabilities specific to the `textDocument/formatting` request.
    /// </summary>
    public /*TOpt*/DocumentFormattingClientCapabilities? Formatting;
    /// <summary>
    /// Capabilities specific to the `textDocument/rangeFormatting` request.
    /// </summary>
    public /*TOpt*/DocumentRangeFormattingClientCapabilities? RangeFormatting;
    /// <summary>
    /// Capabilities specific to the `textDocument/onTypeFormatting` request.
    /// </summary>
    public /*TOpt*/DocumentOnTypeFormattingClientCapabilities? OnTypeFormatting;
    /// <summary>
    /// Capabilities specific to the `textDocument/rename` request.
    /// </summary>
    public /*TOpt*/RenameClientCapabilities? Rename;
    /// <summary>
    /// Capabilities specific to the `textDocument/foldingRange` request.
    /// 
    /// @since 3.10.0
    /// </summary>
    public /*TOpt*/FoldingRangeClientCapabilities? FoldingRange;
    /// <summary>
    /// Capabilities specific to the `textDocument/selectionRange` request.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/SelectionRangeClientCapabilities? SelectionRange;
    /// <summary>
    /// Capabilities specific to the `textDocument/publishDiagnostics` notification.
    /// </summary>
    public /*TOpt*/PublishDiagnosticsClientCapabilities? PublishDiagnostics;
    /// <summary>
    /// Capabilities specific to the various call hierarchy requests.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/CallHierarchyClientCapabilities? CallHierarchy;
    /// <summary>
    /// Capabilities specific to the various semantic token request.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/SemanticTokensClientCapabilities? SemanticTokens;
    /// <summary>
    /// Capabilities specific to the `textDocument/linkedEditingRange` request.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/LinkedEditingRangeClientCapabilities? LinkedEditingRange;
    /// <summary>
    /// Client capabilities specific to the `textDocument/moniker` request.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/MonikerClientCapabilities? Moniker;
    /// <summary>
    /// Capabilities specific to the various type hierarchy requests.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/TypeHierarchyClientCapabilities? TypeHierarchy;
    /// <summary>
    /// Capabilities specific to the `textDocument/inlineValue` request.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/InlineValueClientCapabilities? InlineValue;
    /// <summary>
    /// Capabilities specific to the `textDocument/inlayHint` request.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/InlayHintClientCapabilities? InlayHint;
    /// <summary>
    /// Capabilities specific to the diagnostic pull model.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/DiagnosticClientCapabilities? Diagnostic;
}

/// <summary>
/// Capabilities specific to the notebook document support.
/// 
/// @since 3.17.0
/// </summary>
public class NotebookDocumentClientCapabilities
{
    /// <summary>
    /// Capabilities specific to notebook document synchronization
    /// 
    /// @since 3.17.0
    /// </summary>
    public NotebookDocumentSyncClientCapabilities Synchronization;
}

public class WindowClientCapabilities
{
    /// <summary>
    /// It indicates whether the client supports server initiated
    /// progress using the `window/workDoneProgress/create` request.
    /// 
    /// The capability also controls Whether client supports handling
    /// of progress notifications. If set servers are allowed to report a
    /// `workDoneProgress` property in the request specific server
    /// capabilities.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/System.Boolean? WorkDoneProgress;
    /// <summary>
    /// Capabilities specific to the showMessage request.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/ShowMessageRequestClientCapabilities? ShowMessage;
    /// <summary>
    /// Capabilities specific to the showDocument request.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/ShowDocumentClientCapabilities? ShowDocument;
}

/// <summary>
/// General client capabilities.
/// 
/// @since 3.16.0
/// </summary>
public class GeneralClientCapabilities
{
    /// <summary>
    /// Client capability that signals how the client
    /// handles stale requests (e.g. a request
    /// for which the client will not process the response
    /// anymore since the information is outdated).
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/CancelBooleanWithRetryOnContentModifiedStrings StaleRequestSupport;
    /// <summary>
    /// Client capabilities specific to regular expressions.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/RegularExpressionsClientCapabilities? RegularExpressions;
    /// <summary>
    /// Client capabilities specific to the client's markdown parser.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/MarkdownClientCapabilities? Markdown;
    /// <summary>
    /// The position encodings supported by the client. Client and server
    /// have to agree on the same position encoding to ensure that offsets
    /// (e.g. character position in a line) are interpreted the same on both
    /// sides.
    /// 
    /// To keep the protocol backwards compatible the following applies: if
    /// the value 'utf-16' is missing from the array of position encodings
    /// servers can assume that the client supports UTF-16. UTF-16 is
    /// therefore a mandatory encoding.
    /// 
    /// If omitted it defaults to ['utf-16'].
    /// 
    /// Implementation considerations: since the conversion from one encoding
    /// into another requires the content of the file / line the conversion
    /// is best done where the file is read which is usually on the server
    /// side.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*//*PositionEncodingKind*/string[] PositionEncodings;
}

/// <summary>
/// A relative pattern is a helper to construct glob patterns that are matched
/// relatively to a base URI. The common value for a `baseUri` is a workspace
/// folder root, but it can be another absolute URI as well.
/// 
/// @since 3.17.0
/// </summary>
public class RelativePattern
{
    /// <summary>
    /// A workspace folder or a base URI to which this pattern will be matched
    /// against relatively.
    /// 
    /// This object has "OneOf" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.
    /// </summary>
    public WorkspaceFolderOrURI BaseUri;
    /// <summary>
    /// The actual glob pattern;
    /// </summary>
    public Pattern Pattern;
}

public class WorkspaceEditClientCapabilities
{
    /// <summary>
    /// The client supports versioned document changes in `WorkspaceEdit`s
    /// </summary>
    public /*TOpt*/System.Boolean? DocumentChanges;
    /// <summary>
    /// The resource operations the client supports. Clients should at least
    /// support 'create', 'rename' and 'delete' files and folders.
    /// 
    /// @since 3.13.0
    /// </summary>
    public /*TOpt*//*ResourceOperationKind*/string[] ResourceOperations;
    /// <summary>
    /// The failure handling strategy of a client if applying the workspace edit
    /// fails.
    /// 
    /// @since 3.13.0
    /// </summary>
    public /*TOpt*//*FailureHandlingKind*/string FailureHandling;
    /// <summary>
    /// Whether the client normalizes line endings to the client specific
    /// setting.
    /// If set to `true` the client will normalize line ending characters
    /// in a workspace edit to the client-specified new line
    /// character.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.Boolean? NormalizesLineEndings;
    /// <summary>
    /// Whether the client in general supports change annotations on text edits,
    /// create file, rename file and delete file changes.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/GroupsOnLabelBoolean ChangeAnnotationSupport;
}

public class DidChangeConfigurationClientCapabilities
{
    /// <summary>
    /// Did change configuration notification supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

public class DidChangeWatchedFilesClientCapabilities
{
    /// <summary>
    /// Did change watched files notification supports dynamic registration. Please note
    /// that the current protocol doesn't support static configuration for file changes
    /// from the server side.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// Whether the client has support for <c> RelativePattern relative pattern</c>
    /// or not.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/System.Boolean? RelativePatternSupport;
}

/// <summary>
/// Client capabilities for a <c>WorkspaceSymbolRequest</c>.
/// </summary>
public class WorkspaceSymbolClientCapabilities
{
    /// <summary>
    /// Symbol request supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// Specific capabilities for the `SymbolKind` in the `workspace/symbol` request.
    /// </summary>
    public /*TOpt*/ValueSetSymbolKinds SymbolKind;
    /// <summary>
    /// The client supports tags on `SymbolInformation`.
    /// Clients supporting tags have to handle unknown tags gracefully.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/ValueSetSymbolTags TagSupport;
    /// <summary>
    /// The client support partial workspace symbols. The client will send the
    /// request `workspaceSymbol/resolve` to the server to resolve additional
    /// properties.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/PropertiesStrings ResolveSupport;
}

/// <summary>
/// The client capabilities of a <c>ExecuteCommandRequest</c>.
/// </summary>
public class ExecuteCommandClientCapabilities
{
    /// <summary>
    /// Execute command supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokensWorkspaceClientCapabilities
{
    /// <summary>
    /// Whether the client implementation supports a refresh request sent from
    /// the server to the client.
    /// 
    /// Note that this event is global and will force the client to refresh all
    /// semantic tokens currently shown. It should be used with absolute care
    /// and is useful for situation where a server for example detects a project
    /// wide change that requires such a calculation.
    /// </summary>
    public /*TOpt*/System.Boolean? RefreshSupport;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class CodeLensWorkspaceClientCapabilities
{
    /// <summary>
    /// Whether the client implementation supports a refresh request sent from the
    /// server to the client.
    /// 
    /// Note that this event is global and will force the client to refresh all
    /// code lenses currently shown. It should be used with absolute care and is
    /// useful for situation where a server for example detect a project wide
    /// change that requires such a calculation.
    /// </summary>
    public /*TOpt*/System.Boolean? RefreshSupport;
}

/// <summary>
/// Capabilities relating to events from file operations by the user in the client.
/// 
/// These events do not come from the file system, they come from user operations
/// like renaming a file in the UI.
/// 
/// @since 3.16.0
/// </summary>
public class FileOperationClientCapabilities
{
    /// <summary>
    /// Whether the client supports dynamic registration for file requests/notifications.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// The client has support for sending didCreateFiles notifications.
    /// </summary>
    public /*TOpt*/System.Boolean? DidCreate;
    /// <summary>
    /// The client has support for sending willCreateFiles requests.
    /// </summary>
    public /*TOpt*/System.Boolean? WillCreate;
    /// <summary>
    /// The client has support for sending didRenameFiles notifications.
    /// </summary>
    public /*TOpt*/System.Boolean? DidRename;
    /// <summary>
    /// The client has support for sending willRenameFiles requests.
    /// </summary>
    public /*TOpt*/System.Boolean? WillRename;
    /// <summary>
    /// The client has support for sending didDeleteFiles notifications.
    /// </summary>
    public /*TOpt*/System.Boolean? DidDelete;
    /// <summary>
    /// The client has support for sending willDeleteFiles requests.
    /// </summary>
    public /*TOpt*/System.Boolean? WillDelete;
}

/// <summary>
/// Client workspace capabilities specific to inline values.
/// 
/// @since 3.17.0
/// </summary>
public class InlineValueWorkspaceClientCapabilities
{
    /// <summary>
    /// Whether the client implementation supports a refresh request sent from the
    /// server to the client.
    /// 
    /// Note that this event is global and will force the client to refresh all
    /// inline values currently shown. It should be used with absolute care and is
    /// useful for situation where a server for example detects a project wide
    /// change that requires such a calculation.
    /// </summary>
    public /*TOpt*/System.Boolean? RefreshSupport;
}

/// <summary>
/// Client workspace capabilities specific to inlay hints.
/// 
/// @since 3.17.0
/// </summary>
public class InlayHintWorkspaceClientCapabilities
{
    /// <summary>
    /// Whether the client implementation supports a refresh request sent from
    /// the server to the client.
    /// 
    /// Note that this event is global and will force the client to refresh all
    /// inlay hints currently shown. It should be used with absolute care and
    /// is useful for situation where a server for example detects a project wide
    /// change that requires such a calculation.
    /// </summary>
    public /*TOpt*/System.Boolean? RefreshSupport;
}

/// <summary>
/// Workspace client capabilities specific to diagnostic pull requests.
/// 
/// @since 3.17.0
/// </summary>
public class DiagnosticWorkspaceClientCapabilities
{
    /// <summary>
    /// Whether the client implementation supports a refresh request sent from
    /// the server to the client.
    /// 
    /// Note that this event is global and will force the client to refresh all
    /// pulled diagnostics currently shown. It should be used with absolute care and
    /// is useful for situation where a server for example detects a project wide
    /// change that requires such a calculation.
    /// </summary>
    public /*TOpt*/System.Boolean? RefreshSupport;
}

public class TextDocumentSyncClientCapabilities
{
    /// <summary>
    /// Whether text document synchronization supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// The client supports sending will save notifications.
    /// </summary>
    public /*TOpt*/System.Boolean? WillSave;
    /// <summary>
    /// The client supports sending a will save request and
    /// waits for a response providing text edits which will
    /// be applied to the document before it is saved.
    /// </summary>
    public /*TOpt*/System.Boolean? WillSaveWaitUntil;
    /// <summary>
    /// The client supports did save notifications.
    /// </summary>
    public /*TOpt*/System.Boolean? DidSave;
}

/// <summary>
/// Completion client capabilities
/// </summary>
public class CompletionClientCapabilities
{
    /// <summary>
    /// Whether completion supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// The client supports the following `CompletionItem` specific
    /// capabilities.
    /// </summary>
    public /*TOpt*/SnippetSupportBooleanWithCommitCharactersSupportBooleanWithDocumentationFormatMarkupKindsWithDeprecatedSupportBooleanWithPreselectSupportBooleanWithTagSupportValueSetCompletionItemTagsWithInsertReplaceSupportBooleanWithResolveSupportPropertiesStringsWithInsertTextModeSupportValueSetInsertTextModesWithLabelDetailsSupportBoolean CompletionItem;
    public /*TOpt*/ValueSetCompletionItemKinds CompletionItemKind;
    /// <summary>
    /// Defines how the client handles whitespace and indentation
    /// when accepting a completion item that uses multi line
    /// text in either `insertText` or `textEdit`.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/InsertTextMode InsertTextMode;
    /// <summary>
    /// The client supports to send additional context information for a
    /// `textDocument/completion` request.
    /// </summary>
    public /*TOpt*/System.Boolean? ContextSupport;
    /// <summary>
    /// The client supports the following `CompletionList` specific
    /// capabilities.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/ItemDefaultsStrings CompletionList;
}

public class HoverClientCapabilities
{
    /// <summary>
    /// Whether hover supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// Client supports the following content formats for the content
    /// property. The order describes the preferred format of the client.
    /// </summary>
    public /*TOpt*//*MarkupKind*/string[] ContentFormat;
}

/// <summary>
/// Client Capabilities for a <c>SignatureHelpRequest</c>.
/// </summary>
public class SignatureHelpClientCapabilities
{
    /// <summary>
    /// Whether signature help supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// The client supports the following `SignatureInformation`
    /// specific properties.
    /// </summary>
    public /*TOpt*/DocumentationFormatMarkupKindsWithParameterInformationLabelOffsetSupportBooleanWithActiveParameterSupportBoolean SignatureInformation;
    /// <summary>
    /// The client supports to send additional context information for a
    /// `textDocument/signatureHelp` request. A client that opts into
    /// contextSupport will also support the `retriggerCharacters` on
    /// `SignatureHelpOptions`.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/System.Boolean? ContextSupport;
}

/// <summary>
/// @since 3.14.0
/// </summary>
public class DeclarationClientCapabilities
{
    /// <summary>
    /// Whether declaration supports dynamic registration. If this is set to `true`
    /// the client supports the new `DeclarationRegistrationOptions` return value
    /// for the corresponding server capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// The client supports additional metadata in the form of declaration links.
    /// </summary>
    public /*TOpt*/System.Boolean? LinkSupport;
}

/// <summary>
/// Client Capabilities for a <c>DefinitionRequest</c>.
/// </summary>
public class DefinitionClientCapabilities
{
    /// <summary>
    /// Whether definition supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// The client supports additional metadata in the form of definition links.
    /// 
    /// @since 3.14.0
    /// </summary>
    public /*TOpt*/System.Boolean? LinkSupport;
}

/// <summary>
/// Since 3.6.0
/// </summary>
public class TypeDefinitionClientCapabilities
{
    /// <summary>
    /// Whether implementation supports dynamic registration. If this is set to `true`
    /// the client supports the new `TypeDefinitionRegistrationOptions` return value
    /// for the corresponding server capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// The client supports additional metadata in the form of definition links.
    /// 
    /// Since 3.14.0
    /// </summary>
    public /*TOpt*/System.Boolean? LinkSupport;
}

/// <summary>
/// @since 3.6.0
/// </summary>
public class ImplementationClientCapabilities
{
    /// <summary>
    /// Whether implementation supports dynamic registration. If this is set to `true`
    /// the client supports the new `ImplementationRegistrationOptions` return value
    /// for the corresponding server capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// The client supports additional metadata in the form of definition links.
    /// 
    /// @since 3.14.0
    /// </summary>
    public /*TOpt*/System.Boolean? LinkSupport;
}

/// <summary>
/// Client Capabilities for a <c>ReferencesRequest</c>.
/// </summary>
public class ReferenceClientCapabilities
{
    /// <summary>
    /// Whether references supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// Client Capabilities for a <c>DocumentHighlightRequest</c>.
/// </summary>
public class DocumentHighlightClientCapabilities
{
    /// <summary>
    /// Whether document highlight supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// Client Capabilities for a <c>DocumentSymbolRequest</c>.
/// </summary>
public class DocumentSymbolClientCapabilities
{
    /// <summary>
    /// Whether document symbol supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// Specific capabilities for the `SymbolKind` in the
    /// `textDocument/documentSymbol` request.
    /// </summary>
    public /*TOpt*/ValueSetSymbolKinds SymbolKind;
    /// <summary>
    /// The client supports hierarchical document symbols.
    /// </summary>
    public /*TOpt*/System.Boolean? HierarchicalDocumentSymbolSupport;
    /// <summary>
    /// The client supports tags on `SymbolInformation`. Tags are supported on
    /// `DocumentSymbol` if `hierarchicalDocumentSymbolSupport` is set to true.
    /// Clients supporting tags have to handle unknown tags gracefully.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/ValueSetSymbolTags TagSupport;
    /// <summary>
    /// The client supports an additional label presented in the UI when
    /// registering a document symbol provider.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.Boolean? LabelSupport;
}

/// <summary>
/// The Client Capabilities of a <c>CodeActionRequest</c>.
/// </summary>
public class CodeActionClientCapabilities
{
    /// <summary>
    /// Whether code action supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// The client support code action literals of type `CodeAction` as a valid
    /// response of the `textDocument/codeAction` request. If the property is not
    /// set the request can only return `Command` literals.
    /// 
    /// @since 3.8.0
    /// </summary>
    public /*TOpt*/CodeActionKindValueSetCodeActionKinds CodeActionLiteralSupport;
    /// <summary>
    /// Whether code action supports the `isPreferred` property.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/System.Boolean? IsPreferredSupport;
    /// <summary>
    /// Whether code action supports the `disabled` property.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.Boolean? DisabledSupport;
    /// <summary>
    /// Whether code action supports the `data` property which is
    /// preserved between a `textDocument/codeAction` and a
    /// `codeAction/resolve` request.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.Boolean? DataSupport;
    /// <summary>
    /// Whether the client supports resolving additional code action
    /// properties via a separate `codeAction/resolve` request.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/PropertiesStrings_ ResolveSupport;
    /// <summary>
    /// Whether the client honors the change annotations in
    /// text edits and resource operations returned via the
    /// `CodeAction#edit` property by for example presenting
    /// the workspace edit in the user interface and asking
    /// for confirmation.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.Boolean? HonorsChangeAnnotations;
}

/// <summary>
/// The client capabilities  of a <c>CodeLensRequest</c>.
/// </summary>
public class CodeLensClientCapabilities
{
    /// <summary>
    /// Whether code lens supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// The client capabilities of a <c>DocumentLinkRequest</c>.
/// </summary>
public class DocumentLinkClientCapabilities
{
    /// <summary>
    /// Whether document link supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// Whether the client supports the `tooltip` property on `DocumentLink`.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/System.Boolean? TooltipSupport;
}

public class DocumentColorClientCapabilities
{
    /// <summary>
    /// Whether implementation supports dynamic registration. If this is set to `true`
    /// the client supports the new `DocumentColorRegistrationOptions` return value
    /// for the corresponding server capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// Client capabilities of a <c>DocumentFormattingRequest</c>.
/// </summary>
public class DocumentFormattingClientCapabilities
{
    /// <summary>
    /// Whether formatting supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// Client capabilities of a <c>DocumentRangeFormattingRequest</c>.
/// </summary>
public class DocumentRangeFormattingClientCapabilities
{
    /// <summary>
    /// Whether range formatting supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// Client capabilities of a <c>DocumentOnTypeFormattingRequest</c>.
/// </summary>
public class DocumentOnTypeFormattingClientCapabilities
{
    /// <summary>
    /// Whether on type formatting supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

public class RenameClientCapabilities
{
    /// <summary>
    /// Whether rename supports dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// Client supports testing for validity of rename operations
    /// before execution.
    /// 
    /// @since 3.12.0
    /// </summary>
    public /*TOpt*/System.Boolean? PrepareSupport;
    /// <summary>
    /// Client supports the default behavior result.
    /// 
    /// The value indicates the default behavior used by the
    /// client.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/PrepareSupportDefaultBehavior PrepareSupportDefaultBehavior;
    /// <summary>
    /// Whether the client honors the change annotations in
    /// text edits and resource operations returned via the
    /// rename request's workspace edit by for example presenting
    /// the workspace edit in the user interface and asking
    /// for confirmation.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.Boolean? HonorsChangeAnnotations;
}

public class FoldingRangeClientCapabilities
{
    /// <summary>
    /// Whether implementation supports dynamic registration for folding range
    /// providers. If this is set to `true` the client supports the new
    /// `FoldingRangeRegistrationOptions` return value for the corresponding
    /// server capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// The maximum number of folding ranges that the client prefers to receive
    /// per document. The value serves as a hint, servers are free to follow the
    /// limit.
    /// </summary>
    public /*TOpt*/System.UInt64? RangeLimit;
    /// <summary>
    /// If set, the client signals that it only supports folding complete lines.
    /// If set, client will ignore specified `startCharacter` and `endCharacter`
    /// properties in a FoldingRange.
    /// </summary>
    public /*TOpt*/System.Boolean? LineFoldingOnly;
    /// <summary>
    /// Specific options for the folding range kind.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/ValueSetFoldingRangeKinds FoldingRangeKind;
    /// <summary>
    /// Specific options for the folding range.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/CollapsedTextBoolean FoldingRange;
}

public class SelectionRangeClientCapabilities
{
    /// <summary>
    /// Whether implementation supports dynamic registration for selection range providers. If this is set to `true`
    /// the client supports the new `SelectionRangeRegistrationOptions` return value for the corresponding server
    /// capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// The publish diagnostic client capabilities.
/// </summary>
public class PublishDiagnosticsClientCapabilities
{
    /// <summary>
    /// Whether the clients accepts diagnostics with related information.
    /// </summary>
    public /*TOpt*/System.Boolean? RelatedInformation;
    /// <summary>
    /// Client supports the tag property to provide meta data about a diagnostic.
    /// Clients supporting tags have to handle unknown tags gracefully.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/ValueSetDiagnosticTags TagSupport;
    /// <summary>
    /// Whether the client interprets the version property of the
    /// `textDocument/publishDiagnostics` notification's parameter.
    /// 
    /// @since 3.15.0
    /// </summary>
    public /*TOpt*/System.Boolean? VersionSupport;
    /// <summary>
    /// Client supports a codeDescription property
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.Boolean? CodeDescriptionSupport;
    /// <summary>
    /// Whether code action supports the `data` property which is
    /// preserved between a `textDocument/publishDiagnostics` and
    /// `textDocument/codeAction` request.
    /// 
    /// @since 3.16.0
    /// </summary>
    public /*TOpt*/System.Boolean? DataSupport;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class CallHierarchyClientCapabilities
{
    /// <summary>
    /// Whether implementation supports dynamic registration. If this is set to `true`
    /// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
    /// return value for the corresponding server capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// @since 3.16.0
/// </summary>
public class SemanticTokensClientCapabilities
{
    /// <summary>
    /// Whether implementation supports dynamic registration. If this is set to `true`
    /// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
    /// return value for the corresponding server capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// Which requests the client supports and might send to the server
    /// depending on the server's capability. Please note that clients might not
    /// show semantic tokens or degrade some of the user experience if a range
    /// or full request is advertised by the client but not provided by the
    /// server. If for example the client capability `requests.full` and
    /// `request.range` are both set to true but the server only provides a
    /// range provider the client might not render a minimap correctly or might
    /// even decide to not show any semantic tokens at all.
    /// </summary>
    public RangeBooleanOrAnyByStringWithFullBooleanOrDeltaBoolean Requests;
    /// <summary>
    /// The token types that the client supports.
    /// </summary>
    public System.String[] TokenTypes;
    /// <summary>
    /// The token modifiers that the client supports.
    /// </summary>
    public System.String[] TokenModifiers;
    /// <summary>
    /// The token formats the clients supports.
    /// </summary>
    public /*TokenFormat*/string[] Formats;
    /// <summary>
    /// Whether the client supports tokens that can overlap each other.
    /// </summary>
    public /*TOpt*/System.Boolean? OverlappingTokenSupport;
    /// <summary>
    /// Whether the client supports tokens that can span multiple lines.
    /// </summary>
    public /*TOpt*/System.Boolean? MultilineTokenSupport;
    /// <summary>
    /// Whether the client allows the server to actively cancel a
    /// semantic token request, e.g. supports returning
    /// LSPErrorCodes.ServerCancelled. If a server does the client
    /// needs to retrigger the request.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/System.Boolean? ServerCancelSupport;
    /// <summary>
    /// Whether the client uses semantic tokens to augment existing
    /// syntax tokens. If set to `true` client side created syntax
    /// tokens and semantic tokens are both used for colorization. If
    /// set to `false` the client only uses the returned semantic tokens
    /// for colorization.
    /// 
    /// If the value is `undefined` then the client behavior is not
    /// specified.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/System.Boolean? AugmentsSyntaxTokens;
}

/// <summary>
/// Client capabilities for the linked editing range request.
/// 
/// @since 3.16.0
/// </summary>
public class LinkedEditingRangeClientCapabilities
{
    /// <summary>
    /// Whether implementation supports dynamic registration. If this is set to `true`
    /// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
    /// return value for the corresponding server capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// Client capabilities specific to the moniker request.
/// 
/// @since 3.16.0
/// </summary>
public class MonikerClientCapabilities
{
    /// <summary>
    /// Whether moniker supports dynamic registration. If this is set to `true`
    /// the client supports the new `MonikerRegistrationOptions` return value
    /// for the corresponding server capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// @since 3.17.0
/// </summary>
public class TypeHierarchyClientCapabilities
{
    /// <summary>
    /// Whether implementation supports dynamic registration. If this is set to `true`
    /// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
    /// return value for the corresponding server capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// Client capabilities specific to inline values.
/// 
/// @since 3.17.0
/// </summary>
public class InlineValueClientCapabilities
{
    /// <summary>
    /// Whether implementation supports dynamic registration for inline value providers.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
}

/// <summary>
/// Inlay hint client capabilities.
/// 
/// @since 3.17.0
/// </summary>
public class InlayHintClientCapabilities
{
    /// <summary>
    /// Whether inlay hints support dynamic registration.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// Indicates which properties a client can resolve lazily on an inlay
    /// hint.
    /// </summary>
    public /*TOpt*/PropertiesStrings_ ResolveSupport;
}

/// <summary>
/// Client capabilities specific to diagnostic pull requests.
/// 
/// @since 3.17.0
/// </summary>
public class DiagnosticClientCapabilities
{
    /// <summary>
    /// Whether implementation supports dynamic registration. If this is set to `true`
    /// the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
    /// return value for the corresponding server capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// Whether the clients supports related documents for document diagnostic pulls.
    /// </summary>
    public /*TOpt*/System.Boolean? RelatedDocumentSupport;
}

/// <summary>
/// Notebook specific client capabilities.
/// 
/// @since 3.17.0
/// </summary>
public class NotebookDocumentSyncClientCapabilities
{
    /// <summary>
    /// Whether implementation supports dynamic registration. If this is
    /// set to `true` the client supports the new
    /// `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
    /// return value for the corresponding server capability as well.
    /// </summary>
    public /*TOpt*/System.Boolean? DynamicRegistration;
    /// <summary>
    /// The client supports sending execution summary data per cell.
    /// </summary>
    public /*TOpt*/System.Boolean? ExecutionSummarySupport;
}

/// <summary>
/// Show message request client capabilities
/// </summary>
public class ShowMessageRequestClientCapabilities
{
    /// <summary>
    /// Capabilities specific to the `MessageActionItem` type.
    /// </summary>
    public /*TOpt*/AdditionalPropertiesSupportBoolean MessageActionItem;
}

/// <summary>
/// Client capabilities for the showDocument request.
/// 
/// @since 3.16.0
/// </summary>
public class ShowDocumentClientCapabilities
{
    /// <summary>
    /// The client has support for the showDocument
    /// request.
    /// </summary>
    public System.Boolean Support;
}

/// <summary>
/// Client capabilities specific to regular expressions.
/// 
/// @since 3.16.0
/// </summary>
public class RegularExpressionsClientCapabilities
{
    /// <summary>
    /// The engine's name.
    /// </summary>
    public System.String Engine;
    /// <summary>
    /// The engine's version.
    /// </summary>
    public /*TOpt*/System.String? Version;
}

/// <summary>
/// Client capabilities specific to the used markdown parser.
/// 
/// @since 3.16.0
/// </summary>
public class MarkdownClientCapabilities
{
    /// <summary>
    /// The name of the parser.
    /// </summary>
    public System.String Parser;
    /// <summary>
    /// The version of the parser.
    /// </summary>
    public /*TOpt*/System.String? Version;
    /// <summary>
    /// A list of HTML tags that the client allows / supports in
    /// Markdown.
    /// 
    /// @since 3.17.0
    /// </summary>
    public /*TOpt*/System.String[] AllowedTags;
}
