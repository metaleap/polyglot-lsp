package main

import (
	"fmt"
	glot "polyglot-vsx-and-lsp/pkg"
)

type MetaModel struct {
	Enumerations []struct {
		Name string `json:"name"`
	} `json:"enumerations"`
}

var metaModel MetaModel

func main() {
	meta := glot.Meta.LSP
	metaModel := glot.LoadFromJSONFile[MetaModel](meta.File)
	for _, enum := range metaModel.Enumerations {
		println(fmt.Sprintf("%+v", enum))
	}
}
