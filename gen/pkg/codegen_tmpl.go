package glot

import (
	"bytes"
	"strings"
	"text/template"
)

func (it *GenBase) DoDocComments(root *GenDot) string {
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
				rendered := root.gen.tmplExec(nil, root.gen.tmpl("doc_comments_link", "`{{index . 0}}`"), link_strs)
				doc_lines[i] = doc_line[:idx1] + If(rendered == "", link_strs[0], rendered) + doc_line[idx2+idx1+1:]
				doc_line = doc_lines[i]
			}
		}
	}
	return root.gen.tmplExec(nil, root.gen.tmpl("doc_comments", ""), doc_lines)
}

func (it *GenDot) DoType(t GenType) string {
	return it.DoTypeOptional(t, false)
}

func (it *GenDot) DoTypeOptional(t GenType, optional bool) string {
	type GenDotType struct {
		Dot  *GenDot
		Type GenType
	}

	if optional {
		return it.gen.tmplExec(nil, it.gen.tmpl("type_Optional", ""), GenDotType{Dot: it, Type: t})
	}
	switch t := t.(type) {
	case GenTypeBaseType:
		if s := it.Lang.BaseTypeMapping[string(t)]; s != "" {
			return s
		}
		panic("(missing in " + it.gen.filePathLang + ": \"BaseTypeMapping\"{\"" + t.String() + "\":...})")
	case GenTypeReference:
		return t.String()
	default:
		return it.gen.tmplExec(nil, it.gen.tmpl("type_"+t.kind(), ""), GenDotType{Dot: it, Type: t})
	}
}

func (it *GenDot) If(b bool, ifTrue any, ifFalse any) any { return If(b, ifTrue, ifFalse) }
func (it *GenDot) Up0(s string) string                    { return Up0(s) }

func (it *GenDot) IsEnumTypeName(name string) bool {
	return it.gen.tracked.decls.enumerations[name] != nil
}

func (it *GenDot) IsAliasTypeName(name string) bool {
	return it.gen.tracked.decls.typeAliases[name] != nil
}

func (it *GenDot) IsStructTypeName(name string) bool {
	return it.gen.tracked.decls.structures[name] != nil
}

func isTypeKind[T GenType](it *Gen, t GenType) (is bool) {
	if _, is = t.(T); !is {
		if tref, isref := t.(GenTypeReference); isref {
			if alias := it.tracked.decls.typeAliases[string(tref)]; alias != nil {
				return isTypeKind[T](it, alias.Type)
			}
		}
	}
	return
}

func (it *GenDot) IsTypeKindArray(t GenType) bool {
	return isTypeKind[GenTypeArray](it.gen, t)
}

func (it *GenDot) IsTypeKindOr(t GenType) bool {
	return isTypeKind[GenTypeOr](it.gen, t)
}

func (it *GenDot) IsTypeKindMap(t GenType) bool {
	return isTypeKind[GenTypeMap](it.gen, t)
}

func (it *Gen) tmpl(tmplName string, defaultFallback string) (ret *template.Template) {
	file_path := it.dirPathSrc + "/" + tmplName + ".tmpl"
	if ret = tmpls[file_path]; ret == nil {
		file_exists := FileExists(file_path)
		if defaultFallback == "" && !file_exists {
			defaultFallback = it.Dot.Lang.Tmpls[tmplName]
		}
		if defaultFallback != "" && !file_exists {
			ret = template.Must(template.New(tmplName).Parse(defaultFallback))
		} else if file_exists {
			ret = template.Must(template.ParseFiles(file_path))
		} else {
			panic("template '" + tmplName + "' missing: add file '" + it.dirPathSrc + "/" + tmplName + ".tmpl' or add entry Tmpls/" + tmplName + " to " + it.filePathLang)
		}
		tmpls[file_path] = ret
	}
	return
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
	it.tracked.filesGenerated.code = append(it.tracked.filesGenerated.code, file_path)

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
	tmpl := it.tmpl(tmplName, "")
	it.tmplExec(buf, tmpl, nil)
	filePath = it.dirPathDst + "/" + fileName
	FileWrite(filePath, []byte(it.Dot.FileContents))
	return
}
