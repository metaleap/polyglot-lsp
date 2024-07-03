package main

import (
	glot "polyglot-lsp/pkg"
)

func main() {
	for _, ver := range glot.Vers("lsp_", ".schema.json") {
		meta_model := glot.LoadFromJSONFile[MetaModel](glot.SchemasDir + "/lsp_" + ver + ".json")

		for _, lang := range glot.Langs() {
			generate(&meta_model, ver, lang)
		}
	}
}
