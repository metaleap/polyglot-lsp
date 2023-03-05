package glot

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
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
	dirPathLang string // lang_foo
	dirPathSrc  string // lang_foo/_gen
	dirPathDst  string // lang_foo/{{.GenIdent}}_v{{.GenVer}}
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

type GenLang struct {
	DisplayName  string // eg. "Go" or "C#"
	PkgFile      string // eg "go.mod", "package.json" etc
	SrcFileExt   string // eg ".go", ".cs" etc
	PostGenTools struct {
		Format *GenLangCmd  // eg "go fmt %in" etc
		Check  []GenLangCmd // compiler, type-checker, parser or other linter. `%in` for absolute path of generated pkg-dir
	}
	PostGenCleanUp []string // eg ["obj"]
	Dict           map[string]string
}

type GenLangCmd struct {
	Disabled bool
	Cmd      string
	Env      map[string]string
	PerFile  bool
}

type GenDots interface {
	PerEnum(*Gen, func(*GenEnumeration))
}

type GenType interface {
	key() string
	kind() string
	String() string
}

type GenTypeBase string

func (it GenTypeBase) String() string { return Up0(string(it)) }
func (it GenTypeBase) kind() string   { return "Base" }
func (it GenTypeBase) key() string    { return "Base(" + string(it) + ")" }

type GenTypeEnumeration string

func (it GenTypeEnumeration) String() string { return "Enum" + string(it) }
func (it GenTypeEnumeration) kind() string   { return "Enumeration" }
func (it GenTypeEnumeration) key() string    { return "Enumeration(" + string(it) + ")" }

type GenTypeMapKey string

func (it GenTypeMapKey) String() string { return "MapKey" + string(it) }
func (it GenTypeMapKey) kind() string   { return "MapKey" }
func (it GenTypeMapKey) key() string    { return "MapKey(" + string(it) + ")" }

type GenTypeReference string

func (it GenTypeReference) String() string { return string(it) }
func (it GenTypeReference) kind() string   { return "Reference" }
func (it GenTypeReference) key() string    { return "Reference(" + string(it) + ")" }

type GenTypeMap struct {
	KeyType   GenType
	ValueType GenType
}

func (it GenTypeMap) String() string {
	return "Map_" + it.KeyType.String() + "_" + it.ValueType.String()
}
func (it GenTypeMap) kind() string { return "Map" }
func (it GenTypeMap) key() string  { return "Map[" + it.KeyType.key() + "]" + it.ValueType.key() }

type GenTypeArr struct {
	ElemType GenType
}

func (it GenTypeArr) String() string { return "Arr_" + it.ElemType.String() }
func (it GenTypeArr) kind() string   { return "Arr" }
func (it GenTypeArr) key() string    { return "[" + it.ElemType.key() + "]" }

type GenTypeAnd []GenType

func (it GenTypeAnd) String() string {
	return "And_" + strings.Join(Map(it, func(gt GenType) string { return gt.String() }), "_")
}
func (it GenTypeAnd) kind() string { return "And" }
func (it GenTypeAnd) key() string {
	return "And(" + strings.Join(Map(it, func(gt GenType) string { return gt.key() }), ",") + ")"
}

type GenTypeOr []GenType

func (it GenTypeOr) String() string {
	return "Or_" + strings.Join(Map(it, func(gt GenType) string { return gt.String() }), "_")
}
func (it GenTypeOr) kind() string { return "Or" }
func (it GenTypeOr) key() string {
	return "Or(" + strings.Join(Map(it, func(gt GenType) string { return gt.key() }), ",") + ")"
}

type GenTypeTup []GenType

func (it GenTypeTup) String() string {
	return "Tup_" + strings.Join(Map(it, func(gt GenType) string { return gt.String() }), "_")
}
func (it GenTypeTup) kind() string { return "Tup" }
func (it GenTypeTup) key() string {
	return "Tup(" + strings.Join(Map(it, func(gt GenType) string { return gt.key() }), ",") + ")"
}

type GenTypeLitBool bool

func (it GenTypeLitBool) String() string { return "Bool_" + strconv.FormatBool(bool(it)) }
func (it GenTypeLitBool) kind() string   { return "LitBool" }
func (it GenTypeLitBool) key() string    { return "LitBool(" + strconv.FormatBool(bool(it)) + ")" }

type GenTypeLitString string

func (it GenTypeLitString) String() string { return "String_" + string(it) }
func (it GenTypeLitString) kind() string   { return "LitString" }
func (it GenTypeLitString) key() string    { return "LitString(" + strconv.Quote(string(it)) + ")" }

type GenTypeLitInt int

func (it GenTypeLitInt) String() string { return "Int_" + strconv.Itoa(int(it)) }
func (it GenTypeLitInt) kind() string   { return "LitInt" }
func (it GenTypeLitInt) key() string    { return "LitInt(" + strconv.FormatInt(int64(it), 36) + ")" }

type GenTypeLitStruct []GenTypeLitStructProperty

type GenTypeLitStructProperty struct {
	Name     string
	Type     GenType
	Optional bool
}

func (it GenTypeLitStruct) String() string { panic(">>>GenTypeLitStruct") }
func (it GenTypeLitStruct) kind() string   { return "LitStruct" }
func (it GenTypeLitStruct) key() string    { panic("GenTypeLitStruct:") }

type GenBase struct {
	DocLines   []string
	Since      string
	Deprecated string
}

func (it *GenBase) base() *GenBase { return it }

func (it *GenBase) DocComments(root *GenDot) string {
	if len(it.DocLines) == 0 {
		return ""
	}
	doc_lines := Copy(it.DocLines)
	for _, s := range doc_lines {
		if idx := strings.Index(s, "{@link"); idx >= 0 {
			if idx2 := strings.Index(s[idx:], "}"); idx2 > 0 {
				idx2 += idx
				doc_str := s[idx+1 : idx2]
				panic(doc_str)
			}
		}
	}
	if it.Since != "" && !Exists(doc_lines, func(s string) bool { return strings.Contains(s, "@since") }) {
		doc_lines = append(doc_lines, "@since "+it.Since)
	}
	if it.Deprecated != "" && !Exists(doc_lines, func(s string) bool { return strings.Contains(s, "@deprecated") }) {
		doc_lines = append(doc_lines, "@deprecated "+it.Deprecated)
	}
	return root.gen.tmplExec(nil, root.gen.tmpl("doc_comments"), doc_lines)
}

type GenEnumeration struct {
	GenBase
	Name       string
	Type       GenType
	Enumerants []GenEnumerant
}
type GenEnumerant struct {
	GenBase
	Name    string
	NameCap string
	Value   string
}

func (it *Gen) tmpl(tmplName string) (ret *template.Template) {
	file_path := it.dirPathSrc + "/" + tmplName + ".tmpl"
	if ret = tmpls[file_path]; ret == nil {
		ret = template.Must(template.ParseFiles(file_path))
		tmpls[file_path] = ret
	}
	return
}

func (it *Gen) Type(t GenType) GenType {
	if key := t.key(); it.types[key] == nil {
		it.types[key] = t
	}
	return t
}

func (it *Gen) Generate(dots GenDots) {
	it.Dot.gen = it
	it.types = map[string]GenType{}
	it.dirPathLang = "../lang_" + it.LangIdent
	it.dirPathSrc = it.dirPathLang + "/_gen"
	it.dirPathDst = it.dirPathLang + "/" + it.Dot.GenIdent + "_v" + it.Dot.GenVer
	defer it.postGenCleanUp()
	DirCreate(it.dirPathDst)

	it.Dot.Lang = LoadFromJSONFile[GenLang](it.dirPathLang + "/lang_" + it.LangIdent + ".json")
	it.Dot.PkgName = If(it.Dot.PkgName == "", it.Dot.GenIdent, it.Dot.PkgName)
	if strings.HasPrefix(it.Dot.Lang.PkgFile, "*") {
		it.Dot.Lang.PkgFile = it.Dot.GenIdent + "_v" + it.Dot.GenVer + it.Dot.Lang.PkgFile[1:]
	}
	println("\n", it.dirPathDst+"...")

	var buf bytes.Buffer
	it.genPkgFile(&buf)
	it.genDots(&buf, map[string]func() []any{
		"types_enums": toAnys(it, dots.PerEnum),
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
		tmpl := it.tmpl(file_name)
		it.Dot.Items = dots()
		it.tmplExec(buf, tmpl, nil)

		it.toCodeFile(buf, file_name)
	}
}

func (it *Gen) genSideTypes(buf *bytes.Buffer) {
	types := make([]any, 0, len(it.types))
	for _, ty := range it.types {
		tmpl := it.tmpl("type_" + ty.kind())
		it.Dot.Items = []any{ty}
		it.tmplExec(buf, tmpl, nil)
		types = append(types, it.Dot.FileContents)
	}
	tmpl := it.tmpl("types")
	it.Dot.Items = types
	it.tmplExec(buf, tmpl, nil)
	it.toCodeFile(buf, "types")
}

func (it *Gen) genPkgFile(buf *bytes.Buffer) {
	tmpl := it.tmpl("file_pkg")
	it.tmplExec(buf, tmpl, nil)
	file_path := it.toOutputFile(buf, "file_pkg", it.Dot.Lang.PkgFile)
	it.filesGenerated.other = append(it.filesGenerated.other, file_path)
}

func (it *Gen) tmplExec(buf *bytes.Buffer, tmpl *template.Template, dot any) (ret string) {
	on_the_fly := (buf == nil)
	if on_the_fly {
		buf = new(bytes.Buffer)
	} else {
		buf.Reset()
	}
	if dot == nil {
		dot = &it.Dot
	}
	if err := tmpl.Execute(buf, dot); err != nil {
		panic(err)
	}
	if ret = buf.String(); !on_the_fly {
		it.Dot.FileContents = ret
	}
	return
}

func (it *Gen) toCodeFile(buf *bytes.Buffer, fileName string) {
	file_path := it.toOutputFile(buf, "file_code", fileName+it.Dot.Lang.SrcFileExt)
	it.filesGenerated.code = append(it.filesGenerated.code, file_path)

	cmd_repl := map[string]string{"{file}": PathAbs(file_path)}
	if it.Dot.Lang.PostGenTools.Format.ok() && it.Dot.Lang.PostGenTools.Format.PerFile {
		it.Dot.Lang.PostGenTools.Format.exec(it, cmd_repl)
	}
	for _, check := range it.Dot.Lang.PostGenTools.Check {
		if check.ok() && check.PerFile {
			check.exec(it, cmd_repl)
		}
	}
}

func (it *Gen) toOutputFile(buf *bytes.Buffer, tmplName string, fileName string) (filePath string) {
	tmpl := it.tmpl(tmplName)
	it.tmplExec(buf, tmpl, nil)
	filePath = it.dirPathDst + "/" + fileName
	FileWrite(filePath, []byte(it.Dot.FileContents))
	return
}

func (it *Gen) postGenCleanUp() {
	for _, name := range it.Dot.Lang.PostGenCleanUp {
		DirRemove(it.dirPathDst + "/" + name)
		FileRemove(it.dirPathDst + "/" + name)
	}
}

func (it *GenLangCmd) ok() bool {
	return it != nil && !it.Disabled
}

func (it *GenLangCmd) exec(gen *Gen, repl map[string]string) {
	if repl == nil {
		repl = map[string]string{}
	}
	for old, new := range map[string]string{
		"{dir}":           PathAbs(gen.dirPathDst),
		"{trash_dir}":     trashDir,
		"{pkg_file_name}": gen.Dot.Lang.PkgFile,
		"{pkg_file_path}": PathAbs(gen.dirPathDst + "/" + gen.Dot.Lang.PkgFile),
		"{user_home_dir}": os.Getenv("HOME"),
	} {
		if repl[old] == "" {
			repl[old] = new
		}
	}

	parts := Without(strings.Split(it.Cmd, " "), "")
	for i := 0; i < len(parts); i++ {
		if parts[i] == "{files}" {
			parts = append(parts[:i], append(gen.filesGenerated.code, parts[i+1:]...)...)
			i--
		} else {
			for old, new := range repl {
				parts[i] = strings.ReplaceAll(parts[i], old, new)
			}
		}
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	for name, value := range it.Env {
		for old, new := range repl {
			value = strings.ReplaceAll(value, old, new)
		}
		it.Env[name] = value // just for the `full_command` print-out further below
		cmd.Env = append(cmd.Env, name+"="+value)
	}
	cmd.Env = append(os.Environ(), cmd.Env...)
	{
		full_command := strings.Join(parts, " ")
		for name, value := range it.Env {
			full_command = name + "=" + value + "\t\t" + full_command
		}
		println(">>>", full_command)
	}
	if output, err := cmd.CombinedOutput(); err != nil {
		panic(err.Error() + ":\n" + string(output))
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
