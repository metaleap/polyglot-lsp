package main

import (
	"encoding/json"
	glot "polyglot-lsp/pkg"
)

func main() {
	for _, ver := range glot.Vers("lsp_", ".schema.json") {
		meta_model := glot.LoadFromJSONFile[MetaModel](glot.SchemasDir + "/lsp_" + ver + ".json")

		for _, lang := range glot.Langs() {
			generate(&meta_model, ver, lang)
		}

		for _, notif := range meta_model.Notifications {
			src, _ := json.Marshal(notif)
			println("\n\n\n\n", string(src))
		}
	}
}
