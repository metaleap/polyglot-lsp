package glot

import (
	"bytes"
	"strings"
	"text/template"
)

var tmpls = map[string]*template.Template{}

type Gen struct {
	LangIdent string // eg. "go" or "cs"
	Dot       GenDot

	dirPathLang  string // lang_foo
	dirPathSrc   string // lang_foo/_gen
	dirPathDst   string // lang_foo/{{.GenIdent}}_v{{.GenVer}}
	filePathLang string // lang_foo/lang_foo.json
	tracked      struct {
		types     map[string]GenType
		sideTypes map[string]GenType
		decls     struct {
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

type GenDot struct {
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

type GenDots interface {
	PerTypeAlias(*Gen, func(*GenAlias))
	PerEnumeration(*Gen, func(*GenEnumeration))
	PerStructure(*Gen, func(*GenStructure))
}

func (it *Gen) Generate(dots GenDots) {
	it.Dot.gen = it
	it.tracked.types, it.tracked.sideTypes, it.tracked.decls.enumerations, it.tracked.decls.structures, it.tracked.decls.typeAliases = map[string]GenType{}, map[string]GenType{}, map[string]*GenEnumeration{}, map[string]*GenStructure{}, map[string]*GenAlias{}
	it.dirPathLang = "../lang_" + it.LangIdent
	it.dirPathSrc = it.dirPathLang + "/_gen"
	it.dirPathDst = it.dirPathLang + "/" + it.Dot.GenIdent + "_v" + it.Dot.GenVer
	it.filePathLang = it.dirPathLang + "/lang_" + it.LangIdent + ".json"
	defer it.postGenCleanUp()
	DirCreate(it.dirPathDst)

	it.Dot.Lang = LoadFromJSONFile[GenLang](it.filePathLang)
	it.Dot.PkgName = If(it.Dot.PkgName == "", it.Dot.GenIdent, it.Dot.PkgName)
	if strings.HasPrefix(it.Dot.Lang.PkgFile, "*") {
		it.Dot.Lang.PkgFile = it.Dot.GenIdent + "_v" + it.Dot.GenVer + it.Dot.Lang.PkgFile[1:]
	}
	println("\n", it.dirPathDst+"...")

	var buf bytes.Buffer
	it.genPkgFile(&buf)
	it.genDots(&buf, []Tup[string, func() []any]{ // no map here because the order matters:
		{"types_enumerations", toAnys(it, dots.PerEnumeration)},
		{"types_structures", toAnys(it, dots.PerStructure)},
		{"types_aliases", toAnys(it, dots.PerTypeAlias)},
	})
	it.genSideTypes(&buf)
	if it.Dot.Lang.PostGenTools.Format.ok() && !it.Dot.Lang.PostGenTools.Format.PerFile {
		it.Dot.Lang.PostGenTools.Format.exec(it, nil)
	}
	for _, check := range it.Dot.Lang.PostGenTools.Check {
		if check.ok() && !check.PerFile {
			check.exec(it, nil)
		}
	}
}

func (it *Gen) genDots(buf *bytes.Buffer, dots_by_file_name []Tup[string, func() []any]) {
	for _, tup := range dots_by_file_name {
		file_name, dots := tup.F1, tup.F2
		tmpl := it.tmpl(file_name, "")
		it.Dot.Items = dots()
		for _, item := range it.Dot.Items {
			switch decl := item.(type) {
			case *GenEnumeration:
				it.tracked.decls.enumerations[decl.Name] = decl
				it.tracked.decls.enumerations[decl.NameUp] = decl
				ensureDocHintConstVal(decl.Enumerants, func(p GenEnumerant) any { return p.Value })
			case *GenStructure:
				it.tracked.decls.structures[decl.Name] = decl
				it.tracked.decls.structures[decl.NameUp] = decl
				ensureDocHintConstVal(decl.Properties, func(p GenStructureProperty) any { return p.ConstVal })
				decl.ensureDocHintUnionType()
			case *GenAlias:
				it.tracked.decls.typeAliases[decl.Name] = decl
				it.tracked.decls.typeAliases[decl.NameUp] = decl
				ensureDocHintUnionType(&decl.GenBase, decl.Type, "")
			}
		}
		it.tmplExec(buf, tmpl, nil)

		it.toCodeFile(buf, file_name)
	}
}

func (it *Gen) genSideTypes(buf *bytes.Buffer) {
	types := make([]any, 0, len(it.tracked.sideTypes))
	for _, ty := range it.tracked.sideTypes {
		if _, is_ref := ty.(GenTypeReference); is_ref {
			continue
		}
		tmpl := it.tmpl("type_"+ty.kind(), "")
		it.Dot.Items = []any{ty}
		it.tmplExec(buf, tmpl, nil)
		types = append(types, it.Dot.FileContents)
	}
	tmpl := it.tmpl("types", "")
	it.Dot.Items = types
	it.tmplExec(buf, tmpl, nil)
	it.toCodeFile(buf, "types")
}

func (it *Gen) genPkgFile(buf *bytes.Buffer) {
	tmpl := it.tmpl("file_pkg", "")
	it.tmplExec(buf, tmpl, nil)
	file_path := it.toOutputFile(buf, "file_pkg", it.Dot.Lang.PkgFile)
	it.tracked.filesGenerated.other = append(it.tracked.filesGenerated.other, file_path)
}

func (it *Gen) postGenCleanUp() {
	for _, name := range it.Dot.Lang.PostGenCleanUp {
		DirRemove(it.dirPathDst + "/" + name)
		FileRemove(it.dirPathDst + "/" + name)
	}
}

func ensureDocHintConstVal[T any](items []T, constVal func(T) any) {
	for i, item := range items {
		if base, const_val := any(&items[i]).(interface{ base() *GenBase }).base(), constVal(item); const_val != nil {
			base.DocLines = append(base.DocLines, "The value is always "+ValueString(const_val))
		}
	}
}

func ensureDocHintUnionType(base *GenBase, t GenType, prefix string) {
	switch t := t.(type) {
	case GenTypeOr:
		if len(t.NonNull()) > 1 && base != nil && !base.docHintUnionsEnsured {
			base.docHintUnionsEnsured = true
			base.DocLines = append(base.DocLines, prefix+"\"OneOf\" union-type semantics: only at-most one field is ever set, all others will be null/undefined/nil/empty/etc.")
		}
	case *GenTypeStructure:
		t.ensureDocHintUnionType()
	case GenTypeMap:
		ensureDocHintUnionType(base, t.ValueType, "Every map value has ")
	case GenTypeArray:
		ensureDocHintUnionType(base, t.ElemType, "Every array element has ")
	}
}

func toAnys[T any](gen *Gen, each func(*Gen, func(*T))) func() []any {
	return func() (ret []any) {
		each(gen, func(dot *T) {
			ret = append(ret, dot)
		})
		return
	}
}
