package glot

import (
	"bytes"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var tmpls = map[string]struct {
	Tmpl *template.Template
	Src  string
}{}

func (it *GenBase) DoDocComments(root *GenMain) string {
	if len(it.DocLines) == 0 {
		return ""
	}
	doc_lines := Copy(it.DocLines)
	for i, doc_line := range doc_lines {
		for idx1 := strings.Index(doc_line, "{@link "); idx1 >= 0; idx1 = strings.Index(doc_line, "{@link ") {
			if idx2 := strings.Index(doc_line[idx1:], "}"); idx2 < 0 {
				break
			} else {
				link_strs := []string{doc_line[idx1+len("{@link ") : idx2+idx1]}
				if idx_space := strings.IndexByte(link_strs[0], ' '); idx_space > 0 {
					link_strs = append(link_strs, link_strs[0][idx_space+1:])
					link_strs[0] = link_strs[0][:idx_space]
				}
				rendered := root.gen.tmplExec(nil, root.gen.tmpl("doc_comments_link", "`{{index . 0}}`", false), link_strs)
				doc_lines[i] = doc_line[:idx1] + If(rendered == "", link_strs[0], rendered) + doc_line[idx2+idx1+1:]
				doc_line = doc_lines[i]
			}
		}
	}
	return root.gen.tmplExec(nil, root.gen.tmpl("doc_comments", "", true), doc_lines)
}

func (it *GenMain) DoType(t GenType) string {
	return strings.TrimSpace(it.DoTypeOptional(t, false))
}

func (it *GenMain) DoTypeOptional(t GenType, optional bool) string {
	if optional {
		return it.doType(t, "type_Optional")
	}

	switch t := t.(type) {
	case GenTypeBaseType:
		if s := it.Lang.BaseTypeMapping[string(t)]; s != "" {
			return s
		}
		panic("(missing in " + it.gen.filePathLang + ": \"BaseTypeMapping\"{\"" + t.String() + "\":...})")
	}
	return it.doType(t, "type_"+t.kind())
}

func (it *GenMain) doType(t GenType, tmplName string) (ret string) {
	bag := struct {
		Main         *GenMain
		Type         GenType
		TypeIdentGen string
	}{it, t, "__TypeIdentGen__" + strconv.FormatInt(time.Now().UnixNano(), 36) + "__"}

	tmpl := it.gen.tmpl(tmplName, "", true)
	if ret = it.gen.tmplExec(nil, tmpl, bag); strings.Contains(ret, bag.TypeIdentGen) {
		ident := t.NameSuggestion(!it.Lang.AllowLowerCaseGeneratedTypeIdents)
		ret_new := strings.ReplaceAll(ret, bag.TypeIdentGen, ident)
		for it.gen.tracked.namedAnonDeclRenders[ident] != "" && it.gen.tracked.namedAnonDeclRenders[ident] != ret_new {
			ident = ident + "_"
			ret_new = strings.ReplaceAll(ret, bag.TypeIdentGen, ident)
		}
		ret = ret_new
		it.gen.tracked.namedAnonDeclRenders[ident] = ret
		t = GenTypeReference(ident)
		ret = it.doType(t, "type_Reference")
	}
	if _, is_ref := t.(GenTypeReference); is_ref {
		for rewrite := it.Lang.TypeRefRewrites[ret]; rewrite != ""; rewrite = it.Lang.TypeRefRewrites[ret] {
			ret = rewrite
		}
	}
	return
}

func (it *GenMain) If(b bool, ifTrue any, ifFalse any) any { return If(b, ifTrue, ifFalse) }
func (it *GenMain) Up0(s string) string                    { return Up0(s) }
func (it *GenMain) IsIdentAscii(s string) bool {
	return strings.IndexFunc(s, func(r rune) bool {
		return !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'))
	}) < 0
}

func (it *GenMain) IsEnumTypeName(name string) bool {
	return it.gen.tracked.decls.enumerations[name] != nil
}
func (it *GenMain) IsAliasTypeName(name string) bool {
	return it.gen.tracked.decls.typeAliases[name] != nil
}
func (it *GenMain) IsStructTypeName(name string) bool {
	return it.gen.tracked.decls.structures[name] != nil
}

func (it *GenMain) IsTypeKindOfArray(t GenType) bool { return isGenTypeOf[GenTypeArray](it.gen, t) }
func (it *GenMain) IsTypeKindOfMap(t GenType) bool   { return isGenTypeOf[GenTypeMap](it.gen, t) }
func (it *GenMain) IsTypeKindOfAnd(t GenType) bool   { return isGenTypeOf[GenTypeAnd](it.gen, t) }
func (it *GenMain) IsTypeKindOfOr(t GenType) bool    { return isGenTypeOf[GenTypeOr](it.gen, t) }
func (it *GenMain) IsTypeKindOfTuple(t GenType) bool { return isGenTypeOf[GenTypeTuple](it.gen, t) }
func (it *GenMain) IsTypeKindOfStructure(t GenType) bool {
	return isGenTypeOf[*GenTypeStructure](it.gen, t)
}
func (it *GenMain) IsTypeKindOf(t GenType, s string) (ret bool) {
	var ref string
	switch t := t.(type) {
	case GenTypeBaseType:
		ret = (string(t) == s)
	case GenTypeReference:
		ref, ret = string(t), (string(t) == s)
	case GenTypeEnumeration:
		ref, ret = string(t), (string(t) == s)
	case GenTypeMapKey:
		ref, ret = string(t), (string(t) == s)
	}
	if (!ret) && ref != "" {
		if decl := it.gen.tracked.decls.typeAliases[ref]; decl != nil {
			return it.IsTypeKindOf(decl.Type, s)
		}
		if decl := it.gen.tracked.decls.enumerations[ref]; decl != nil {
			return it.IsTypeKindOf(decl.Type, s)
		}
	}
	return
}

func (it *Gen) tmpl(tmplName string, fallbackSrc string, must bool) (ret *template.Template) {
	ret, _ = it.tmplSrc(tmplName, fallbackSrc, must)
	return
}

func (it *Gen) tmplSrc(tmplName string, fallbackSrc string, must bool) (*template.Template, string) {
	file_path := it.dirPathSrc + "/" + tmplName + ".tmpl"
	tup, got := tmpls[file_path]
	if !got {
		if alt_src := it.Main.Lang.Tmpls[tmplName]; alt_src != "" {
			fallbackSrc = alt_src
		}
		if tup.Src = fallbackSrc; FileExists(file_path) {
			tup.Src = string(FileRead(file_path))
		}
		if tup.Src != "" {
			tup.Tmpl = template.Must(template.New(tmplName).Parse(tup.Src))
		} else if must {
			panic("template '" + tmplName + "' missing: add file '" + it.dirPathSrc + "/" + tmplName + ".tmpl' or add entry Tmpls/" + tmplName + " to " + it.filePathLang)
		}
		tmpls[file_path] = tup
	}
	return tup.Tmpl, tup.Src
}

func (it *Gen) tmplExec(buf *bytes.Buffer, tmpl *template.Template, dot any) (ret string) {
	if dot == nil {
		dot = &it.Main
	}
	on_the_fly := (buf == nil)
	if on_the_fly {
		buf = new(bytes.Buffer)
	} else {
		buf.Reset()
	}
	if err := tmpl.Execute(buf, dot); err != nil {
		panic(err)
	}
	if ret = buf.String(); !on_the_fly {
		it.Main.FileContents = ret
	}
	return
}

func (it *Gen) toCodeFile(buf *bytes.Buffer, fileName string) {
	file_path := it.toOutputFile(buf, "file_code", fileName+it.Main.Lang.SrcFileExt)
	it.tracked.filesGenerated.code = append(it.tracked.filesGenerated.code, file_path)

	cmd_repl := map[string]string{"{file}": PathAbs(file_path)}
	if it.Main.Lang.PostGenTools.Format.ok() && it.Main.Lang.PostGenTools.Format.PerFile {
		it.Main.Lang.PostGenTools.Format.exec(it, cmd_repl, false)
	}
	for _, check := range it.Main.Lang.PostGenTools.Check {
		if check.ok() && check.PerFile {
			check.exec(it, cmd_repl, false)
		}
	}
}

func (it *Gen) toOutputFile(buf *bytes.Buffer, tmplName string, fileName string) (filePath string) {
	tmpl := it.tmpl(tmplName, "", true)
	it.tmplExec(buf, tmpl, nil)
	filePath = it.dirPathDst + "/" + fileName
	FileWrite(filePath, []byte(it.Main.FileContents))
	return
}
