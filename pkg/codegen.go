package glot

import (
	"bytes"
	"text/template"
)

var tmpls = map[string]*template.Template{}

type GenConfig struct {
	PkgFile           string // eg "go.mod", "package.json" etc
	CommentLinePrefix string // eg "// ", "# " etc
	Named             map[string]string
}

type Gen struct {
	Lang    string
	Version string
	Dot     struct {
		GenConfig
		GenType      string // aka "lsp" or "vsx"
		Items        []any
		FileContents string
	}

	dirPathLang string // lang_foo
	dirPathSrc  string // lang_foo/_gen
	dirPathDst  string // lang_foo/GenType_Version
}

type GenDots interface {
	PerEnum(func(*GenEnumeration))
}

type GenEnumeration struct {
	DocLines   string
	TypeName   string
	BaseType   string
	Enumerants []GenEnumerant
}

type GenEnumerant struct {
	Name  string
	Value string
}

func (it *Gen) tmpl(filePath string) (ret *template.Template) {
	if ret = tmpls[filePath]; ret == nil {
		ret = template.Must(template.ParseFiles(filePath))
		tmpls[filePath] = ret
	}
	return
}

func (it *Gen) Generate(dots GenDots) {
	it.dirPathLang = "lang_" + it.Lang
	it.Dot.GenConfig = LoadFromJSONFile[GenConfig](it.dirPathLang + "/" + it.dirPathLang + ".json")
	it.dirPathSrc = "lang_" + it.Lang + "/_gen"
	it.dirPathDst = it.dirPathLang + "/" + it.Dot.GenType + "_" + it.Version
	println(it.dirPathDst + "...")
	DirCreate(it.dirPathDst)

	var buf bytes.Buffer
	for file_name, dots := range map[string]func() []any{
		"types_enums": toAnys(dots.PerEnum),
	} {
		tmpl := it.tmpl(it.dirPathSrc + "/" + file_name + ".tmpl")
		it.Dot.Items = dots()
		buf.Reset()
		tmpl.Execute(&buf, &it.Dot)
		it.Dot.FileContents = buf.String()

		tmpl = it.tmpl(it.dirPathSrc + "/" + "file_code.tmpl")
		buf.Reset()
		tmpl.Execute(&buf, &it.Dot)
		it.Dot.FileContents = buf.String()
		if it.Lang == "go" {
			// src, err := format.Source([]byte(it.Dot.FileContents))
			// if err == nil {
			// 	it.Dot.FileContents = string(src)
			// } else {
			// 	panic(err)
			// }
		}
		FileWrite(it.dirPathDst+"/"+file_name+".go", []byte(it.Dot.FileContents))
	}
}

func toAnys[T any](each func(func(*T))) func() []any {
	return func() (ret []any) {
		each(func(dot *T) {
			ret = append(ret, dot)
		})
		return
	}
}
