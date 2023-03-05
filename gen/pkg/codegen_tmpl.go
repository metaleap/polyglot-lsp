package glot

import (
	"bytes"
	"text/template"
)

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
	tmpl := it.tmpl(tmplName, "")
	it.tmplExec(buf, tmpl, nil)
	filePath = it.dirPathDst + "/" + fileName
	FileWrite(filePath, []byte(it.Dot.FileContents))
	return
}
