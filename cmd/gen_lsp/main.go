package main

import (
	glot "polyglot-vsx-and-lsp/pkg"
)

func main() {
	for _, version := range glot.Versions("lsp_", ".schema.json") {
		meta_model := glot.LoadFromJSONFile[MetaModel](glot.SchemasDirName + "/lsp_" + version + ".json")
		for _, lang := range glot.Langs() {
			generate(&meta_model, version, lang)
		}
	}
}
