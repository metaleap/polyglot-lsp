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
	Value string
}

type GenTypeAlias struct {
	GenBase
	Type GenType
}

type GenStructure struct {
	GenBase
	Extends    []GenType
	Mixins     []GenType
	Properties []GenStructureProperty
}
type GenStructureProperty struct {
	GenBase
	Type     GenType
	Optional bool
}

type GenBase struct {
	Name       string
	NameUp     string
	DocLines   []string
	Since      string
	Deprecated string
}

func (it *GenBase) base() *GenBase { return it }

type GenType interface {
	fmt.Stringer
	NameSuggestion(bool) string
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

type GenTypeMapKey string

func (it GenTypeMapKey) NameSuggestion(up bool) string {
	return If(up, Up0, self[string])(string(it)) + "Key"
}
func (it GenTypeMapKey) String() string { return genTypeString(it) }
func (it GenTypeMapKey) kind() string   { return "MapKey" }
func (it GenTypeMapKey) key() string    { return ":" + string(it) }

type GenTypeReference string

func (it GenTypeReference) NameSuggestion(up bool) string {
	return If(up, Up0, self[string])(string(it))
}
func (it GenTypeReference) String() string { return Up0(string(it)) }
func (it GenTypeReference) kind() string   { return "Reference" }
func (it GenTypeReference) key() string    { return "@" + string(it) }

type GenTypeMap struct {
	KeyType   GenType
	ValueType GenType
}

func (it GenTypeMap) NameSuggestion(up bool) string {
	return it.ValueType.NameSuggestion(up) + "By" + it.KeyType.NameSuggestion(up)
}
func (it GenTypeMap) String() string { return genTypeString(it) }
func (it GenTypeMap) kind() string   { return "Map" }
func (it GenTypeMap) key() string    { return "[" + it.KeyType.key() + ":" + it.ValueType.key() + "]" }

type GenTypeArray struct {
	ElemType GenType
}

func (it GenTypeArray) NameSuggestion(up bool) string { return it.ElemType.NameSuggestion(up) + "s" }
func (it GenTypeArray) String() string                { return genTypeString(it) }
func (it GenTypeArray) kind() string                  { return "Array" }
func (it GenTypeArray) key() string                   { return "[" + it.ElemType.key() + "]" }

type GenTypeAnd []GenType

func (it GenTypeAnd) NameSuggestion(up bool) string {
	return strings.Join(Map(it, func(t GenType) string { return t.NameSuggestion(up) }), "And")
}
func (it GenTypeAnd) String() string { return genTypeString(it) }
func (it GenTypeAnd) kind() string   { return "And" }
func (it GenTypeAnd) key() string {
	return "{" + strings.Join(Map(it, func(t GenType) string { return t.key() }), "&") + "}"
}

type GenTypeOr []GenType

func (it GenTypeOr) NameSuggestion(up bool) string {
	return strings.Join(Map(it, func(t GenType) string { return t.NameSuggestion(up) }), "Or")
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
	return strings.Join(Map(it, func(t GenType) string { return t.NameSuggestion(up) }), "With")
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

func (it GenTypeStructure) NameSuggestion(up bool) string {
	return strings.Join(Map(it.Properties, func(p GenStructureProperty) string {
		return If(up, Up0, self[string])(p.Name)
	}), "")
}
func (it GenTypeStructure) String() string { return genTypeString(it) }
func (it GenTypeStructure) kind() string   { return "Structure" }
func (it GenTypeStructure) key() string {
	return "{" + strings.Join(Map(it.Properties, func(p GenStructureProperty) string {
		return p.Name + If(p.Optional, "?", "") + " " + p.Type.key()
	}), ";") + "}"
}

func (it *Gen) Type(t GenType) GenType {
	return t
}
