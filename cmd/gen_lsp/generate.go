package main

import (
	glot "polyglot-vsx-and-lsp/pkg"
)

func generate(metaModel *MetaModel, version string, lang string) {
	gen := glot.Gen{Lang: lang, Version: version}
	gen.Dot.GenType = "lsp"
	gen.Generate(metaModel)
}

func (it *MetaModel) PerEnum(do func(*glot.GenEnumeration)) {
	for _, enum := range it.Enumerations {
		do(&glot.GenEnumeration{
			TypeName: enum.Name,
			BaseType: enum.Type.String(),
			Enumerants: glot.Map(enum.Values, func(e MMEnumerant) glot.GenEnumerant {
				return glot.GenEnumerant{
					Name:  e.Name,
					Value: e.Value.String(),
				}
			}),
		})
	}

	return
}
