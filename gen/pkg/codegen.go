package glot

import (
	"bytes"
	"strings"
)

type Gen struct {
	LangIdent string  // eg. "go" or "cs"
	Main      GenMain // root "dot" object given to templates `types_*.tmpl` and `file_*.tmpl`

	dirPathLang  string // lang_foo
	dirPathSrc   string // lang_foo/_gen
	dirPathDst   string // lang_foo/{{.GenIdent}}_v{{.GenVer}}
	filePathLang string // lang_foo/lang_foo.json
	tracked      struct {
		types                map[string]GenType
		namedAnonTypeRenders []string
		decls                struct {
			enumerations map[string]*GenEnumeration
			structures   map[string]*GenStructure
			typeAliases  map[string]*GenAlias
		}
		filesGenerated struct {
			code  []string
			other []string
		}
	}
}

type GenMain struct {
	Lang         GenLang // the contents of your lang_foo/lang_foo.json
	GenRepo      string
	GenTitle     string
	GenVer       string
	GenIdent     string // aka "lsp" or "vsx"
	PkgName      string // defaults to `GenIdent`
	Items        []any
	FileContents string

	gen *Gen
}

type GenLang struct { // the contents of your lang_foo/lang_foo.json
	DisplayName  string // eg. "Go" or "C#"
	PkgFile      string // eg "go.mod", "package.json" etc
	SrcFileExt   string // eg ".go", ".cs" etc
	PostGenTools struct {
		Format *GenLangCmd  // eg "go fmt %in" etc
		Check  []GenLangCmd // compiler, type-checker, parser or other linter. `%in` for absolute path of generated pkg-dir
	}
	PostGenCleanUp  []string // eg ["obj"]
	Dict            map[string]string
	BaseTypeMapping map[string]string
	Tmpls           map[string]string
}

type Source interface {
	GenTypeAliases(*Gen) []*GenAlias
	GenEnumerations(*Gen) []*GenEnumeration
	GenStructures(*Gen) []*GenStructure
}

func (it *Gen) Generate(source Source) {
	it.Main.gen = it
	it.tracked.types, it.tracked.decls.enumerations, it.tracked.decls.structures, it.tracked.decls.typeAliases = map[string]GenType{}, map[string]*GenEnumeration{}, map[string]*GenStructure{}, map[string]*GenAlias{}
	it.dirPathLang = "../lang_" + it.LangIdent
	it.dirPathSrc = it.dirPathLang + "/_gen"
	it.dirPathDst = it.dirPathLang + "/" + it.Main.GenIdent + "_v" + it.Main.GenVer
	it.filePathLang = it.dirPathLang + "/lang_" + it.LangIdent + ".json"
	defer it.postGenCleanUp()
	DirCreate(it.dirPathDst)

	it.Main.Lang = LoadFromJSONFile[GenLang](it.filePathLang)
	it.Main.PkgName = If(it.Main.PkgName == "", it.Main.GenIdent, it.Main.PkgName)
	if strings.HasPrefix(it.Main.Lang.PkgFile, "*") {
		it.Main.Lang.PkgFile = it.Main.GenIdent + "_v" + it.Main.GenVer + it.Main.Lang.PkgFile[1:]
	}
	println("\n", it.dirPathDst+"...")

	var buf bytes.Buffer
	it.genPkgFile(&buf)
	for _, decls_by_file_name := range []Tup[string, []any]{ // no map here because the ordering of these matters:
		{"types_enumerations", toAnys(source.GenEnumerations(it))},
		{"types_aliases", toAnys(source.GenTypeAliases(it))},
		{"types_structures", toAnys(source.GenStructures(it))},
	} {
		it.genMainDecls(&buf, decls_by_file_name.F1, decls_by_file_name.F2)
	}
	it.genNamedAnonTypes(&buf)

	if it.Main.Lang.PostGenTools.Format.ok() && !it.Main.Lang.PostGenTools.Format.PerFile {
		it.Main.Lang.PostGenTools.Format.exec(it, nil)
	}
	for _, check := range it.Main.Lang.PostGenTools.Check {
		if check.ok() && !check.PerFile {
			check.exec(it, nil)
		}
	}
}

func (it *Gen) genMainDecls(buf *bytes.Buffer, fileName string, decls []any) {
	it.Main.Items = decls
	tmpl := it.tmpl(fileName, "")
	for _, item := range it.Main.Items {
		switch decl := item.(type) {
		case *GenEnumeration:
			it.tracked.decls.enumerations[decl.Name] = decl
			it.tracked.decls.enumerations[decl.NameUp] = decl
			it.EnsureTypeTracked(decl.Type)
			ensureDocHintConstVal(decl.Enumerants, func(p GenEnumerant) any { return p.Value })
		case *GenAlias:
			it.tracked.decls.typeAliases[decl.Name] = decl
			it.tracked.decls.typeAliases[decl.NameUp] = decl
			it.EnsureTypeTracked(decl.Type)
			ensureDocHintUnionType(&decl.GenBase, decl.Type, "")
		case *GenStructure:
			it.tracked.decls.structures[decl.Name] = decl
			it.tracked.decls.structures[decl.NameUp] = decl
			it.EnsureTypeTracked(&decl.GenTypeStructure)
			it.ensureTypesTracked(decl.Extends)
			it.ensureTypesTracked(decl.Mixins)
			ensureDocHintConstVal(decl.Properties, func(p GenStructureProperty) any { return p.ConstVal })
			decl.ensureDocHintUnionType()
		default:
			panic(decl)
		}
	}
	it.tmplExec(buf, tmpl, nil)
	it.toCodeFile(buf, fileName)
}

func (it *Gen) genNamedAnonTypes(buf *bytes.Buffer) {
	tmpl := it.tmpl("types", "")
	it.Main.Items = toAnys(it.tracked.namedAnonTypeRenders)
	it.tmplExec(buf, tmpl, nil)
	it.toCodeFile(buf, "types")
}

func (it *Gen) genPkgFile(buf *bytes.Buffer) {
	tmpl := it.tmpl("file_pkg", "")
	it.tmplExec(buf, tmpl, nil)
	file_path := it.toOutputFile(buf, "file_pkg", it.Main.Lang.PkgFile)
	it.tracked.filesGenerated.other = append(it.tracked.filesGenerated.other, file_path)
}

func (it *Gen) postGenCleanUp() {
	for _, name := range it.Main.Lang.PostGenCleanUp {
		DirRemove(it.dirPathDst + "/" + name)
		FileRemove(it.dirPathDst + "/" + name)
	}
}

func ensureDocHintConstVal[T any](items []T, constVal func(T) any) {
	for i, item := range items {
		if base, const_val := any(&items[i]).(interface{ base() *GenBase }).base(), constVal(item); const_val != nil {
			base.DocLines = append(base.DocLines, "The value is always "+ValueString(const_val)+". ")
		}
	}
}

func ensureDocHintUnionType(base *GenBase, t GenType, prefix string) {
	switch t := t.(type) {
	case GenTypeOr:
		if len(t.NonNull()) > 1 && base != nil && !base.docHintUnionsEnsured {
			if base.docHintUnionsEnsured = true; len(base.DocLines) > 0 {
				base.DocLines = append(base.DocLines, "")
			}
			base.DocLines = append(base.DocLines, prefix+"\"OneOf\" union-type semantics: only at-most one field in it is ever set, all others will be null/undefined/nil/empty/etc.")
		}
	case *GenTypeStructure:
		t.ensureDocHintUnionType()
	case GenTypeMap:
		ensureDocHintUnionType(base, t.ValueType, "Every object in the map has ")
	case GenTypeArray:
		ensureDocHintUnionType(base, t.ElemType, "Every object in the array has ")
	}
}

func toAnys[T any](items []T) []any { return Map(items, func(it T) any { return it }) }
