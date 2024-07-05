package main

import (
	"strings"

	glot "polyglot-lsp/pkg"
)

func generate(metaModel *MetaModel, ver string, lang string) {
	gen := glot.Gen{LangIdent: lang, Main: glot.GenMain{
		GenTitle: "Language Server Protocol (LSP)",
		GenIdent: "lsp",
		PkgName:  "lsp_v" + strings.ReplaceAll(ver, ".", "_"),
		GenVer:   ver,
		GenRepo:  "github.com/metaleap/polyglot-lsp/gen/cmd/gen_lsp",
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

func (it *MetaModel) GenExtras(gen *glot.Gen) (ret []any) {
	repl := strings.NewReplacer("/", "_", "$", "_")

	prepForTmpl := func(msgBase *MMMessageBase, base *MMBase, result *MMType) bool {
		msgBase.MethodNameSafe = repl.Replace(msgBase.Method)
		msgBase.IsClientToServer = (msgBase.MessageDirection == MMMessageDirectionClientToServer) || (msgBase.MessageDirection == MMMessageDirectionBoth)
		msgBase.IsServerToClient = (msgBase.MessageDirection == MMMessageDirectionServerToClient) || (msgBase.MessageDirection == MMMessageDirectionBoth)
		msgBase.DocLines = strings.Split(base.Documentation, "\n")
		if len(msgBase.Params) == 1 {
			msgBase.UnaryParamsTypeName = strings.TrimSpace(string(msgBase.Params[0].Name))
		}
		if len(msgBase.Params) <= 1 && msgBase.UnaryParamsTypeName == "" {
			msgBase.UnaryParamsTypeName = "Void"
		}
		msgBase.IsReq, msgBase.IsInit = (result != nil), (msgBase.Method == "initialize")
		if result != nil {
			msgBase.ResultType = result.toGenType(it, gen)
		}

		if idx := strings.IndexByte(base.Since, ' '); idx > 0 {
			base.Since = base.Since[:idx]
		}
		return !(base.Since > it.MetaData.Version)
	}
	for _, it := range it.Notifications {
		if prepForTmpl(&it.MMMessageBase, &it.MMBase, nil) {
			ret = append(ret, it)
		}
	}
	for _, it := range it.Requests {
		if prepForTmpl(&it.MMMessageBase, &it.MMBase, &it.Result) {
			ret = append(ret, it)
		}
	}
	return
}

func (it *MetaModel) GenEnumerations(gen *glot.Gen) []*glot.GenEnumeration {
	return glot.Map(glot.Filter(it.Enumerations, func(en MMEnumeration) bool {
		return !en.Proposed
	}), func(en MMEnumeration) *glot.GenEnumeration {
		return &glot.GenEnumeration{
			GenBase: it.toGenBase(&en.MMBase),
			Type:    en.Type.toGenType(it, gen),
			Enumerants: glot.Map(glot.Filter(en.Values, func(e MMEnumerant) bool { return !e.Proposed }), func(e MMEnumerant) glot.GenEnumerant {
				return glot.GenEnumerant{
					GenBase: it.toGenBase(&e.MMBase),
					Value:   e.Value,
				}
			}),
		}
	})
}

func (it *MetaModel) GenTypeAliases(gen *glot.Gen) []*glot.GenAlias {
	return glot.Map(glot.Filter(it.TypeAliases, func(ta MMTypeAlias) bool {
		return !ta.Proposed
	}), func(ta MMTypeAlias) *glot.GenAlias {
		return &glot.GenAlias{
			GenBase: it.toGenBase(&ta.MMBase),
			Type:    ta.Type.toGenType(it, gen),
		}
	})
}

func (it *MetaModel) GenStructures(gen *glot.Gen) []*glot.GenStructure {
	return glot.Map(glot.Filter(it.Structures, func(st MMStructure) bool {
		return !st.Proposed
	}), func(st MMStructure) *glot.GenStructure {
		return &glot.GenStructure{
			GenTypeStructure: glot.GenTypeStructure{
				GenBase: it.toGenBase(&st.MMBase),
				Properties: glot.Map(glot.Filter(st.Properties, func(p MMProperty) bool { return !p.Proposed }), func(p MMProperty) glot.GenStructureProperty {
					return glot.GenStructureProperty{
						GenBase:  it.toGenBase(&p.MMBase),
						Type:     p.Type.toGenType(it, gen),
						Optional: p.Optional,
						ConstVal: p.Type.constVal(),
					}
				}),
			},
			Mixins:  glot.Map(st.Mixins, func(t MMType) glot.GenType { return t.toGenType(it, gen) }),
			Extends: glot.Map(st.Extends, func(t MMType) glot.GenType { return t.toGenType(it, gen) }),
		}
	})
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
