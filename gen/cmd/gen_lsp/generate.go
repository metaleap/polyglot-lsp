package main

import (
	"strings"

	glot "polyglot-vsx-and-lsp/pkg"
)

func generate(metaModel *MetaModel, ver string, lang string) {
	gen := glot.Gen{LangIdent: lang, Dot: glot.GenDot{
		GenTitle: "Language Server Protocol (LSP)",
		GenIdent: "lsp",
		GenVer:   ver,
		GenRepo:  "github.com/metaleap/polyglot-vsx-and-lsp/gen/cmd/gen_lsp",
	}}
	gen.Generate(metaModel)
}

func (*MetaModel) toGenBase(it *MMBase) glot.GenBase {
	return glot.GenBase{Deprecated: it.Deprecated, Since: it.Since, DocLines: strings.Split(it.Documentation, "\n")}
}

func (it *MetaModel) PerEnum(gen *glot.Gen, do func(*glot.GenEnumeration)) {
	for _, enum := range it.Enumerations {
		if enum.Proposed {
			continue
		}
		do(&glot.GenEnumeration{
			GenBase: it.toGenBase(&enum.MMBase),
			Name:    enum.Name,
			Type:    gen.Type(enum.Type.toGenType()),
			Enumerants: glot.Map(glot.Filter(enum.Values, func(e MMEnumerant) bool { return !e.Proposed }), func(e MMEnumerant) glot.GenEnumerant {
				return glot.GenEnumerant{
					GenBase: it.toGenBase(&e.MMBase),
					Name:    e.Name,
					NameCap: glot.Up0(e.Name),
					Value:   e.Value.String(),
				}
			}),
		})
	}
}

func (it *MMType) toGenType() glot.GenType {
	switch it.Kind {
	case MMTypeKindBase:
		return glot.GenTypeBase(it.Name)
	case MMTypeKindReference:
		return glot.GenTypeReference(it.Name)
	case MMTypeKindArray:
		return glot.GenTypeArr{it.Element.toGenType()}
	case MMTypeKindMap:
		return glot.GenTypeMap{KeyType: it.Key.toGenType(), ValueType: it.Value.t.toGenType()}
	case MMTypeKindAnd:
		return glot.GenTypeAnd(glot.Map(it.Items, func(t MMType) glot.GenType { return t.toGenType() }))
	case MMTypeKindOr:
		return glot.GenTypeOr(glot.Map(it.Items, func(t MMType) glot.GenType { return t.toGenType() }))
	case MMTypeKindTuple:
		return glot.GenTypeTup(glot.Map(it.Items, func(t MMType) glot.GenType { return t.toGenType() }))
	case MMTypeKindStringLiteral:
		return glot.GenTypeLitString(it.Value.s)
	case MMTypeKindIntegerLiteral:
		return glot.GenTypeLitInt(it.Value.i)
	case MMTypeKindBooleanLiteral:
		return glot.GenTypeLitBool(it.Value.b)
	case MMTypeKindLiteral:
		return glot.GenTypeLitStruct(glot.Map(it.Value.l.Properties, func(p MMProperty) glot.GenTypeLitStructProperty {
			return glot.GenTypeLitStructProperty{Name: p.Name, Optional: p.Optional, Type: p.Type.toGenType()}
		}))
	}
	panic(it.Kind)
}
