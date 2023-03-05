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

	types          map[string]GenType
	filesGenerated struct {
		code  []string
		other []string
	}
	dirPathLang  string // lang_foo
	dirPathSrc   string // lang_foo/_gen
	dirPathDst   string // lang_foo/{{.GenIdent}}_v{{.GenVer}}
	filePathLang string // lang_foo/lang_foo.json
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
}

type GenDots interface {
	PerEnumeration(*Gen, func(*GenEnumeration))
	PerStructure(*Gen, func(*GenStructure))
}

func (it *Gen) Generate(dots GenDots) {
	it.Dot.gen = it
	it.types = map[string]GenType{}
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
	it.genDots(&buf, map[string]func() []any{
		"types_enumerations": toAnys(it, dots.PerEnumeration),
		"types_structures":   toAnys(it, dots.PerStructure),
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

func (it *Gen) genDots(buf *bytes.Buffer, dots_by_file_name map[string]func() []any) {
	for file_name, dots := range dots_by_file_name {
		tmpl := it.tmpl(file_name, "")
		it.Dot.Items = dots()
		it.tmplExec(buf, tmpl, nil)

		it.toCodeFile(buf, file_name)
	}
}

func (it *Gen) genSideTypes(buf *bytes.Buffer) {
	types := make([]any, 0, len(it.types))
	for _, ty := range it.types {
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
	it.filesGenerated.other = append(it.filesGenerated.other, file_path)
}

func (it *Gen) postGenCleanUp() {
	for _, name := range it.Dot.Lang.PostGenCleanUp {
		DirRemove(it.dirPathDst + "/" + name)
		FileRemove(it.dirPathDst + "/" + name)
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
