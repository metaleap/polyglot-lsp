package glot

import (
	"bytes"
	"io/fs"
	"strings"
)

type Gen struct {
	LangIdent string  // eg. "go" or "cs"
	Main      GenMain // root "dot" object given to templates `decls_*.tmpl` and `file_*.tmpl`

	source       Source
	dirPathLang  string // lang_foo
	dirPathSrc   string // lang_foo/_gen
	dirPathDst   string // lang_foo/{{.GenIdent}}_v{{.GenVer}}
	filePathLang string // lang_foo/lang_foo.json
	tracked      struct {
		types                map[string]GenType
		namedAnonDeclRenders map[string]string
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
	Lang     GenLang // the contents of your lang_foo/lang_foo.json
	GenRepo  string
	GenTitle string
	GenVer   string
	GenIdent string // aka "lsp" or "vsx"
	PkgName  string // defaults to `GenIdent`
	Decl     struct {
		Enums       []*GenEnumeration
		TypeAliases []*GenAlias
		Structures  []*GenStructure

		NamedAnons []string
	}
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
	PostGenCleanUp                    []string // eg ["obj"]
	Dict                              map[string]string
	BaseTypeMapping                   map[string]string
	Tmpls                             map[string]string
	AllowLowerCaseGeneratedTypeIdents bool
}

type Source interface {
	GenTypeAliases(*Gen) []*GenAlias
	GenEnumerations(*Gen) []*GenEnumeration
	GenStructures(*Gen) []*GenStructure
}

func (it *Gen) Generate(source Source) {
	it.Main.gen, it.source = it, source
	it.tracked.types, it.tracked.namedAnonDeclRenders, it.tracked.decls.enumerations, it.tracked.decls.structures, it.tracked.decls.typeAliases = map[string]GenType{}, map[string]string{}, map[string]*GenEnumeration{}, map[string]*GenStructure{}, map[string]*GenAlias{}
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

	it.conv()
	for _, tmpl_name := range Dir(it.dirPathSrc, func(entry fs.DirEntry, _ string) (string, bool) {
		name := entry.Name()
		return strings.TrimSuffix(name, ".tmpl"), strings.HasPrefix(name, "decls_") && strings.HasSuffix(name, ".tmpl")
	}) {
		it.genMainDecls(&buf, tmpl_name)
	}
	it.genNamedAnonDecls(&buf)

	if it.Main.Lang.PostGenTools.Format.ok() && !it.Main.Lang.PostGenTools.Format.PerFile {
		it.Main.Lang.PostGenTools.Format.exec(it, nil)
	}
	for _, check := range it.Main.Lang.PostGenTools.Check {
		if check.ok() && !check.PerFile {
			check.exec(it, nil)
		}
	}
}

func (it *Gen) conv() {
	it.Main.Decl.Enums = it.source.GenEnumerations(it)
	it.Main.Decl.TypeAliases = it.source.GenTypeAliases(it)
	it.Main.Decl.Structures = it.source.GenStructures(it)
	for _, decl := range it.Main.Decl.Enums {
		it.tracked.decls.enumerations[decl.Name] = decl
		it.tracked.decls.enumerations[decl.NameUp] = decl
		it.EnsureTypeTracked(decl.Type)
		ensureDocHintConstVal(decl.Enumerants, func(p GenEnumerant) any { return p.Value })
	}
	for _, decl := range it.Main.Decl.TypeAliases {
		it.tracked.decls.typeAliases[decl.Name] = decl
		it.tracked.decls.typeAliases[decl.NameUp] = decl
		it.EnsureTypeTracked(decl.Type)
		ensureDocHintUnionType(&decl.GenBase, decl.Type, "")
	}
	for _, decl := range it.Main.Decl.Structures {
		it.tracked.decls.structures[decl.Name] = decl
		it.tracked.decls.structures[decl.NameUp] = decl
		it.EnsureTypeTracked(&decl.GenTypeStructure)
		it.ensureTypesTracked(decl.Extends)
		it.ensureTypesTracked(decl.Mixins)
		ensureDocHintConstVal(decl.Properties, func(p GenStructureProperty) any { return p.ConstVal })
		decl.ensureDocHintUnionType()
	}
}

func (it *Gen) genMainDecls(buf *bytes.Buffer, tmplName string) {
	tmpl := it.tmpl(tmplName, "")
	it.tmplExec(buf, tmpl, nil)
	it.toCodeFile(buf, tmplName)
}

func (it *Gen) genNamedAnonDecls(buf *bytes.Buffer) {
	tmpl := it.tmpl("decls", "")
	it.Main.Decl.NamedAnons = MapValues(it.tracked.namedAnonDeclRenders)
	it.tmplExec(buf, tmpl, nil)
	it.toCodeFile(buf, "decls")
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
			base.DocLines = append(base.DocLines, prefix+"\"OneOf\" (union type) semantics: only (at most) one field in it is ever set, all others will be null/undefined/nil/empty/zero-length/etc.")
		} else {
			for i := range t {
				ensureDocHintUnionType(base, t[i], "")
			}
		}
	case GenTypeAnd:
		for i := range t {
			ensureDocHintUnionType(base, t[i], "")
		}
	case GenTypeTuple:
		for i := range t {
			ensureDocHintUnionType(base, t[i], "")
		}
	case *GenTypeStructure:
		t.ensureDocHintUnionType()
	case GenTypeMap:
		ensureDocHintUnionType(base, t.ValueType, "Every object in the map has ")
	case GenTypeArray:
		ensureDocHintUnionType(base, t.ElemType, "Every object in the array has ")
	}
}
