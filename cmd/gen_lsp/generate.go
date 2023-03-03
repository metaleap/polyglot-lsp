package main

import (
	"fmt"

	glot "polyglot-vsx-and-lsp/pkg"
)

func generate(metaModel *MetaModel) {
	println(fmt.Sprintf("%+v", metaModel.MetaData.Version))
	for _, lang := range glot.Langs() {
		println(lang)
	}
}
