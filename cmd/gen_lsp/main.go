package main

import (
	glot "polyglot-vsx-and-lsp/pkg"
)

func main() {
	for _, version := range glot.Versions("lsp_", ".schema.json") {
		meta_model := glot.LoadFromJSONFile[MetaModel](glot.SchemasDirName + "/lsp_" + version + ".json")
		generate(&meta_model)
	}
}
