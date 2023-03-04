package main

import (
	glot "polyglot-vsx-and-lsp/pkg"
)

func main() {
	defer glot.FinalCleanUp()
	for _, ver := range glot.Vers("lsp_", ".schema.json") {
		meta_model := glot.LoadFromJSONFile[MetaModel](glot.SchemasDir + "/lsp_" + ver + ".json")
		for _, lang := range glot.Langs() {
			generate(&meta_model, ver, lang)
		}
	}
}
