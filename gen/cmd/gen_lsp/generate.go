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
	it.Documentation = strings.TrimSpace(it.Documentation)

	doc_lines := glot.If[[]string](it.Documentation == "", nil, strings.Split(it.Documentation, "\n"))
	if it.Since != "" && !glot.Exists(doc_lines, func(s string) bool { return strings.Contains(s, "@since") }) {
		doc_lines = append(doc_lines, "", "@since "+it.Since)
	}
	if it.Deprecated != "" && !glot.Exists(doc_lines, func(s string) bool { return strings.Contains(s, "@deprecated") }) {
		doc_lines = append(doc_lines, "", "@deprecated "+it.Deprecated)
	}

	return glot.GenBase{Name: it.Name, NameUp: glot.Up0(it.Name), DocLines: doc_lines}
}

func (it *MetaModel) PerTypeAlias(gen *glot.Gen, do func(*glot.GenAlias)) {
	for _, type_alias := range it.TypeAliases {
		if type_alias.Proposed {
			continue
		}
		do(&glot.GenAlias{
			GenBase: it.toGenBase(&type_alias.MMBase),
			Type:    type_alias.Type.toGenType(it, gen),
		})
	}
}

func (it *MetaModel) PerEnumeration(gen *glot.Gen, do func(*glot.GenEnumeration)) {
	for _, enumeration := range it.Enumerations {
		if enumeration.Proposed {
			continue
		}
		do(&glot.GenEnumeration{
			GenBase: it.toGenBase(&enumeration.MMBase),
			Type:    enumeration.Type.toGenType(it, gen),
			Enumerants: glot.Map(glot.Filter(enumeration.Values, func(e MMEnumerant) bool { return !e.Proposed }), func(e MMEnumerant) glot.GenEnumerant {
				return glot.GenEnumerant{
					GenBase: it.toGenBase(&e.MMBase),
					Value:   e.Value,
				}
			}),
		})
	}
}

func (it *MetaModel) PerStructure(gen *glot.Gen, do func(*glot.GenStructure)) {
	for _, structure := range it.Structures {
		if structure.Proposed {
			continue
		}
		do(&glot.GenStructure{
			GenTypeStructure: glot.GenTypeStructure{
				GenBase: it.toGenBase(&structure.MMBase),
				Properties: glot.Map(glot.Filter(structure.Properties, func(p MMProperty) bool { return !p.Proposed }), func(p MMProperty) glot.GenStructureProperty {
					return glot.GenStructureProperty{
						GenBase:  it.toGenBase(&p.MMBase),
						Type:     p.Type.toGenType(it, gen),
						Optional: p.Optional,
						ConstVal: p.Type.constVal(),
					}
				}),
			},
			Mixins:  glot.Map(structure.Mixins, func(t MMType) glot.GenType { return t.toGenType(it, gen) }),
			Extends: glot.Map(structure.Extends, func(t MMType) glot.GenType { return t.toGenType(it, gen) }),
		})
	}
}

func (it *MMType) toGenType(mm *MetaModel, gen *glot.Gen) (ret glot.GenType) {
	switch it.Kind {
	case MMTypeKindBase:
		ret = glot.GenTypeBaseType(it.Name)
	case MMTypeKindReference:
		ret = glot.GenTypeReference(it.Name)
	case MMTypeKindArray:
		ret = glot.GenTypeArray{it.Element.toGenType(mm, gen)}
	case MMTypeKindMap:
		ret = glot.GenTypeMap{KeyType: it.Key.toGenType(mm, gen), ValueType: it.Value.t.toGenType(mm, gen)}
	case MMTypeKindAnd:
		ret = glot.GenTypeAnd(glot.Map(it.Items, func(t MMType) glot.GenType { return t.toGenType(mm, gen) }))
	case MMTypeKindOr:
		ret = glot.GenTypeOr(glot.Map(it.Items, func(t MMType) glot.GenType { return t.toGenType(mm, gen) }))
	case MMTypeKindTuple:
		ret = glot.GenTypeTuple(glot.Map(it.Items, func(t MMType) glot.GenType { return t.toGenType(mm, gen) }))
	case MMTypeKindStringLiteral:
		ret = (&MMType{Kind: MMTypeKindBase, Name: MMTypeNameString}).toGenType(mm, gen)
	case MMTypeKindIntegerLiteral:
		ret = (&MMType{Kind: MMTypeKindBase, Name: MMTypeNameInteger}).toGenType(mm, gen)
	case MMTypeKindBooleanLiteral:
		ret = (&MMType{Kind: MMTypeKindBase, Name: MMTypeNameBoolean}).toGenType(mm, gen)
	case MMTypeKindLiteral:
		props := glot.Filter(it.Value.l.Properties, func(p MMProperty) bool { return !p.Proposed })
		if len(props) == 0 {
			ret = glot.GenTypeMap{
				KeyType:   glot.GenTypeBaseType(MMTypeNameString),
				ValueType: glot.GenTypeBaseType("any"),
			}
		} else {
			ret = &glot.GenTypeStructure{
				GenBase: mm.toGenBase(&it.Value.l.MMBase),
				Properties: glot.Map(props, func(p MMProperty) glot.GenStructureProperty {
					return glot.GenStructureProperty{
						GenBase:  mm.toGenBase(&p.MMBase),
						Type:     p.Type.toGenType(mm, gen),
						Optional: p.Optional,
						ConstVal: p.Type.constVal(),
					}
				})}
		}
	default:
		panic(it.Kind)
	}
	return gen.EnsureTypeTracked(ret)
}
