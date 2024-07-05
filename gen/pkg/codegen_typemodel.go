package glot

import (
	"fmt"
	"strings"
)

type GenEnumeration struct {
	GenBase
	Type       GenType
	Enumerants []GenEnumerant
}
type GenEnumerant struct {
	GenBase
	Value NumberOrString
}

type GenAlias struct {
	GenBase
	Type GenType
}

type GenStructure struct {
	GenTypeStructure
	Extends []GenType
	Mixins  []GenType
}
type GenStructureProperty struct {
	GenBase
	Type     GenType
	ConstVal any
	Optional bool
}

func (it *GenStructureProperty) OptionalOrOr() bool {
	return it.Optional || it.Type.kind() == "Or" || it.Type.kind() == "or"
}
func (it *GenStructureProperty) HasConstVal() bool { return it.ConstVal != nil }
func (it *GenStructureProperty) equiv(cmp *GenStructureProperty) bool {
	return it.Name == cmp.Name && sameGenType(it.Type, cmp.Type)
}

type GenBase struct {
	Name     string
	NameUp   string
	DocLines []string

	docHintUnionsEnsured     bool
	docHintStringEnumEnsured bool
}

func (it *GenBase) base() *GenBase { return it }

type GenType interface {
	fmt.Stringer
	NameSuggestion(up bool) string
	key() string
	kind() string
}

func genTypeString(it GenType) string {
	return "<" + it.kind() + ">" + it.key()
}

type GenTypeBaseType string

func (it GenTypeBaseType) NameSuggestion(up bool) string {
	return If(up, Up0, self[string])(string(it))
}
func (it GenTypeBaseType) String() string { return string(it) }
func (it GenTypeBaseType) kind() string   { return "Base" }
func (it GenTypeBaseType) key() string    { return "b#" + string(it) }

type GenTypeEnumeration string

func (it GenTypeEnumeration) NameSuggestion(up bool) string {
	return If(up, Up0, self[string])(string(it)) + "Enum"
}
func (it GenTypeEnumeration) String() string { return genTypeString(it) }
func (it GenTypeEnumeration) kind() string   { return "Enumeration" }
func (it GenTypeEnumeration) key() string    { return "e#" + string(it) }

type GenTypeReference string

func (it GenTypeReference) NameSuggestion(up bool) string {
	return If(up, Up0, self[string])(string(it))
}
func (it GenTypeReference) String() string { return string(it) }
func (it GenTypeReference) kind() string   { return "Reference" }
func (it GenTypeReference) key() string    { return "@" + string(it) }

type GenTypeMapKey string

func (it GenTypeMapKey) NameSuggestion(up bool) string {
	return If(up, Up0, self[string])(string(it)) + "Key"
}
func (it GenTypeMapKey) String() string { return genTypeString(it) }
func (it GenTypeMapKey) kind() string   { return "MapKey" }
func (it GenTypeMapKey) key() string    { return ":" + string(it) }

type GenTypeMap struct {
	KeyType   GenType
	ValueType GenType
}

func (it GenTypeMap) NameSuggestion(up bool) string {
	return it.ValueType.NameSuggestion(up) + "By" + it.KeyType.NameSuggestion(true)
}
func (it GenTypeMap) String() string { return genTypeString(it) }
func (it GenTypeMap) kind() string   { return "Map" }
func (it GenTypeMap) key() string    { return "[" + it.KeyType.key() + ":" + it.ValueType.key() + "]" }

type GenTypeArray struct {
	ElemType GenType
}

func (it GenTypeArray) NameSuggestion(up bool) string {
	elem_ns := it.ElemType.NameSuggestion(up)
	return elem_ns + If(strings.HasSuffix(elem_ns, "s"), "es", "s")
}
func (it GenTypeArray) String() string { return genTypeString(it) }
func (it GenTypeArray) kind() string   { return "Array" }
func (it GenTypeArray) key() string    { return "[" + it.ElemType.key() + "]" }

type GenTypeAnd []GenType

func (it GenTypeAnd) NameSuggestion(up bool) string {
	return strings.Join(MapIdx(it, func(t GenType, i int) string { return t.NameSuggestion(If(i == 0, up, true)) }), "And")
}
func (it GenTypeAnd) String() string { return genTypeString(it) }
func (it GenTypeAnd) kind() string   { return "And" }
func (it GenTypeAnd) key() string {
	return "{" + strings.Join(Map(it, func(t GenType) string { return t.key() }), "&") + "}"
}

type GenTypeOr []GenType

func (it GenTypeOr) NameSuggestion(up bool) string {
	return strings.Join(MapIdx(it, func(t GenType, i int) string { return t.NameSuggestion(If(i == 0, up, true)) }), "Or")
}
func (it GenTypeOr) String() string { return genTypeString(it) }
func (it GenTypeOr) kind() string   { return "Or" }
func (it GenTypeOr) key() string {
	return "{" + strings.Join(Map(it, func(t GenType) string { return t.key() }), "|") + "}"
}
func (it GenTypeOr) NonNull() []GenType {
	return Filter(it, func(t GenType) bool {
		base, is := t.(GenTypeBaseType)
		return (!is) || base != "null"
	})
}

type GenTypeTuple []GenType

func (it GenTypeTuple) NameSuggestion(up bool) string {
	return strings.Join(MapIdx(it, func(t GenType, i int) string { return t.NameSuggestion(If(i == 0, up, true)) }), "With")
}
func (it GenTypeTuple) String() string { return genTypeString(it) }
func (it GenTypeTuple) kind() string   { return "Tuple" }
func (it GenTypeTuple) key() string {
	return "{" + strings.Join(Map(it, func(t GenType) string { return t.key() }), ",") + "}"
}

type GenTypeStructure struct {
	GenBase
	Properties []GenStructureProperty
}

func (it *GenTypeStructure) base() *GenBase { return &it.GenBase }
func (it *GenTypeStructure) ensureDocHintUnionType() {
	for i := range it.Properties {
		ensureDocHintUnionType(&it.Properties[i].GenBase, it.Properties[i].Type, "This object has ")
	}
}
func (it *GenTypeStructure) NameSuggestion(up bool) string {
	if len(it.Properties) == 0 {
		return If(up, "Void", "void")
	}
	up0 := If(up, Up0, self[string])
	return strings.Join(Map(it.Properties, func(p GenStructureProperty) string {
		n1, n2 := up0(p.Name), p.Type.NameSuggestion(up)
		return If(n1 == n2, n1, n1+n2)
	}), "With")
}
func (it *GenTypeStructure) String() string { return genTypeString(it) }
func (it *GenTypeStructure) kind() string   { return "Structure" }
func (it *GenTypeStructure) key() string {
	return it.Name + "{" + strings.Join(Map(it.Properties, func(p GenStructureProperty) string {
		return p.Name + If(p.Optional, "?", "") + " " + p.Type.key() + If(p.ConstVal == nil, "", ValueString(p.ConstVal))
	}), ";") + "}"
}

func sameGenTypes(t1 []GenType, t2 []GenType) (ok bool) {
	if ok = (len(t1) == len(t2)); ok {
		for _, t := range t1 {
			if !Exists(t2, func(t2 GenType) bool { return sameGenType(t, t2) }) {
				return false
			}
		}
	}
	return
}

func sameGenType(t1 GenType, t2 GenType) bool {
	if t1 == nil || t2 == nil {
		return t1 == nil && t2 == nil
	}
	switch t1 := t1.(type) {
	case GenTypeBaseType:
		t2, ok := t2.(GenTypeBaseType)
		return ok && t1 == t2
	case GenTypeEnumeration:
		t2, ok := t2.(GenTypeEnumeration)
		return ok && t1 == t2
	case GenTypeMapKey:
		t2, ok := t2.(GenTypeMapKey)
		return ok && t1 == t2
	case GenTypeReference:
		t2, ok := t2.(GenTypeReference)
		return ok && t1 == t2
	case GenTypeAnd:
		t2, ok := t2.(GenTypeAnd)
		return ok && sameGenTypes(t1, t2)
	case GenTypeOr:
		t2, ok := t2.(GenTypeOr)
		return ok && sameGenTypes(t1, t2)
	case GenTypeTuple:
		t2, ok := t2.(GenTypeTuple)
		return ok && sameGenTypes(t1, t2)
	case GenTypeArray:
		t2, ok := t2.(GenTypeArray)
		return ok && sameGenType(t1.ElemType, t2.ElemType)
	case GenTypeMap:
		t2, ok := t2.(GenTypeMap)
		return ok && sameGenType(t1.KeyType, t2.KeyType) && sameGenType(t1.ValueType, t2.ValueType)
	case *GenTypeStructure:
		t2, ok := t2.(*GenTypeStructure)
		if ok = ok && len(t1.Properties) == len(t2.Properties) && t1.Name == t2.Name; ok {
			for i := range t1.Properties {
				if !Exists(t2.Properties, func(p GenStructureProperty) bool { return t1.Properties[i].equiv(&p) }) {
					return false
				}
			}
		}
		return ok
	default:
		panic(t1)
	}
}

func (it *Gen) ensureTypesTracked(t []GenType) []GenType {
	for i := range t {
		t[i] = it.EnsureTypeTracked(t[i])
	}
	return t
}

func (it *Gen) EnsureTypeTracked(t GenType) GenType {
	if key := t.key(); it.tracked.types[key] == nil {
		it.tracked.types[key] = t // do it again below, this just to avoid infinite-type recursions
		switch ty := t.(type) {
		case GenTypeArray:
			ty.ElemType = it.EnsureTypeTracked(ty.ElemType)
			t = ty
		case GenTypeMap:
			ty.KeyType = it.EnsureTypeTracked(ty.KeyType)
			ty.ValueType = it.EnsureTypeTracked(ty.ValueType)
			t = ty
		case GenTypeAnd:
			ty = it.ensureTypesTracked(ty)
			t = ty
		case GenTypeOr:
			// ditch pattern { {a;b?;c?} | {a?;b;c?} | {a?;b?;c} }
			tstruct, is_struct := ty[0].(*GenTypeStructure)
			if is_struct && AllEq(ty, func(t1 GenType, t2 GenType) bool { return sameGenType(t1, t2) }) {
				for i := range tstruct.Properties {
					tstruct.Properties[i].Optional = true
				}
				t = it.EnsureTypeTracked(tstruct)
			} else {
				ty = it.ensureTypesTracked(ty)
				t = ty
			}
		case GenTypeTuple:
			ty = it.ensureTypesTracked(ty)
			t = ty
		case *GenTypeStructure:
			ensureDocHintConstVal(ty.Properties, func(p GenStructureProperty) any { return p.ConstVal })
			for i, p := range ty.Properties {
				ty.Properties[i].Type = it.EnsureTypeTracked(p.Type)
			}
			t = ty
		}
		it.tracked.types[key] = t
	}
	return t
}

func isGenTypeOf[T GenType](it *Gen, t GenType) (is bool) {
	if _, is = t.(T); !is {
		if tref, isref := t.(GenTypeReference); isref {
			if alias := it.tracked.decls.typeAliases[string(tref)]; alias != nil {
				is = isGenTypeOf[T](it, alias.Type)
			}
		}
	}
	return
}
