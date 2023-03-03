package main

import (
	"fmt"

	glot "polyglot-vsx-and-lsp/pkg"
)

func main() {
	for _, version := range glot.Versions("lsp_", ".schema.json") {
		generate(version)
	}
}

func generate(version string) {
	meta_model := glot.LoadFromJSONFile[MetaModel](glot.SchemasDirName + "/lsp_" + version + ".json")
	println(fmt.Sprintf("%+v", meta_model))
}
