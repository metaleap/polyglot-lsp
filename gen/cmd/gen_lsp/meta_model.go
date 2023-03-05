package main

import (
	"encoding/json"
	"errors"
	"strconv"

	glot "polyglot-vsx-and-lsp/pkg"
)

type MetaModel struct {
	MetaData struct {
		Version string `json:"version"`
	} `json:"metaData"`
	Enumerations  []MMEnumeration  `json:"enumerations"`
	TypeAliases   []MMTypeAlias    `json:"typeAliases"`
	Structures    []MMStructure    `json:"structures"`
	Notifications []MMNotification `json:"notifications"`
	Requests      []MMRequest      `json:"requests"`
}

type MMBase struct {
	Name          string `json:"name,omitempty"`
	Deprecated    string `json:"deprecated,omitempty"`
	Documentation string `json:"documentation,omitempty"`
	Proposed      bool   `json:"proposed,omitempty"`
	Since         string `json:"since,omitempty"`
}

type MMEnumeration struct {
	MMBase
	SupportsCustomValues bool          `json:"supportsCustomValues,omitempty"`
	Type                 MMType        `json:"type"`
	Values               []MMEnumerant `json:"values"`
}

type MMEnumerant struct {
	MMBase
	Value glot.NumberOrString `json:"value"`
}

type MMStructure struct {
	MMBase
	Extends    []MMType     `json:"extends,omitempty"`
	Mixins     []MMType     `json:"mixins,omitempty"`
	Properties []MMProperty `json:"properties"`
}

type MMStructureLiteral struct {
	MMBase
	Properties []MMProperty `json:"properties"`
}

type MMProperty struct {
	MMBase
	Optional bool   `json:"optional,omitempty"`
	Type     MMType `json:"type"`
}

type MMTypeAlias struct {
	MMBase
	Type MMType `json:"type"`
}

type MMNotification struct {
	MMBase
	MMMessageBase
}

type MMMessageBase struct {
	MessageDirection    MMNotificationMessageDirection `json:"messageDirection"`
	Method              string                         `json:"method,omitempty"`
	Params              MMTypes                        `json:"params,omitempty"`
	RegistrationMethod  string                         `json:"registrationMethod,omitempty"`
	RegistrationOptions *MMType                        `json:"registrationOptions,omitempty"`
}

type MMNotificationMessageDirection string

const (
	MMNotificationMessageDirectionBoth           MMNotificationMessageDirection = "both"
	MMNotificationMessageDirectionClientToServer MMNotificationMessageDirection = "clientToServer"
	MMNotificationMessageDirectionServerToClient MMNotificationMessageDirection = "serverToClient"
)

type MMRequest struct {
	MMBase
	MMMessageBase
	ErrorData     *MMType `json:"errorData,omitempty"`
	PartialResult *MMType `json:"partialResult,omitempty"`
	Result        MMType  `json:"result"`
}

type MMType struct {
	Kind    MMTypeKind   `json:"kind"`              // always
	Name    MMTypeName   `json:"name,omitempty"`    // if EnumerationType, MapKeyType, or `Kind` of base|reference
	Key     *MMType      `json:"key,omitempty"`     // if `Kind` of map
	Value   *MMTypeValue `json:"value,omitempty"`   // if StructureLiteralType or `Kind` of booleanLiteral|integerLiteral|map|stringLiteral
	Element *MMType      `json:"element,omitempty"` // if `Kind` of array
	Items   []MMType     `json:"items,omitempty"`   // if `Kind` of and|or|tuple
}

func (it *MMType) constVal() any {
	switch it.Kind {
	case MMTypeKindBooleanLiteral:
		return it.Value.b
	case MMTypeKindIntegerLiteral:
		return it.Value.i
	case MMTypeKindStringLiteral:
		return it.Value.s
	}
	return nil
}

type MMTypeKind string

const (
	MMTypeKindBase           MMTypeKind = "base"
	MMTypeKindReference      MMTypeKind = "reference"
	MMTypeKindArray          MMTypeKind = "array"
	MMTypeKindMap            MMTypeKind = "map"
	MMTypeKindAnd            MMTypeKind = "and"
	MMTypeKindOr             MMTypeKind = "or"
	MMTypeKindTuple          MMTypeKind = "tuple"
	MMTypeKindLiteral        MMTypeKind = "literal"
	MMTypeKindStringLiteral  MMTypeKind = "stringLiteral"
	MMTypeKindIntegerLiteral MMTypeKind = "integerLiteral"
	MMTypeKindBooleanLiteral MMTypeKind = "booleanLiteral"
)

// MMTypeName can be any, not just the predefined `MMTypeNameFoo` constants.
type MMTypeName string

const (
	MMTypeNameURI         MMTypeName = "URI"
	MMTypeNameDocumentUri MMTypeName = "DocumentUri"
	MMTypeNameInteger     MMTypeName = "integer"
	MMTypeNameUinteger    MMTypeName = "uinteger"
	MMTypeNameNumber      MMTypeName = "number"
	MMTypeNameDecimal     MMTypeName = "decimal"
	MMTypeNameRegExp      MMTypeName = "RegExp"
	MMTypeNameString      MMTypeName = "string"
	MMTypeNameBoolean     MMTypeName = "boolean"
	MMTypeNameNull        MMTypeName = "null"
)

type MMTypeValue struct {
	kind MMTypeValueKind

	m map[string]any
	a []any
	b bool
	s string
	f float64
	i int64
	t *MMType
	l *MMStructureLiteral
}

type MMTypeValueKind int

const (
	_ MMTypeValueKind = iota
	MMTypeValueKindBool
	MMTypeValueKindStr
	MMTypeValueKindFloat
	MMTypeValueKindInt
	MMTypeValueKindArr
	MMTypeValueKindObj
	MMTypeValueKindObjType
	MMTypeValueKindObjStructureLiteral
)

func (it *MMTypeValue) UnmarshalJSON(src []byte) (err error) {
	*it = MMTypeValue{}
	switch s := string(src); {
	case src[0] == '[':
		if err = json.Unmarshal(src, &it.a); err == nil {
			it.kind = MMTypeValueKindArr
		}
	case src[0] == '{':
		if err = json.Unmarshal(src, &it.m); err == nil {
			it.kind = MMTypeValueKindObj
		}
	case src[0] == '"' && src[len(src)-1] == '"':
		if it.s, err = strconv.Unquote(s); err == nil {
			it.kind = MMTypeValueKindStr
		}
	case src[0] >= '0' && src[0] <= '9':
		if it.f, err = strconv.ParseFloat(s, 64); err == nil {
			it.kind = MMTypeValueKindFloat
			if it.i, err = strconv.ParseInt(s, 0, 64); err == nil {
				it.kind = MMTypeValueKindInt
			} else {
				err, it.i = nil, int64(it.f)
			}
		}
	default:
		if it.b, err = strconv.ParseBool(s); err == nil {
			it.kind = MMTypeValueKindBool
		}
	}

	if err == nil && it.m != nil {
		for name, kind_dst := range map[string]glot.Tup[MMTypeValueKind, any]{
			"kind":       {MMTypeValueKindObjType, &it.t},
			"properties": {MMTypeValueKindObjStructureLiteral, &it.l},
		} {
			if _, has := it.m[name]; has {
				if err = json.Unmarshal(src, kind_dst.F2); err != nil {
					err = nil
				} else {
					it.kind = kind_dst.F1
					break
				}
			}
		}
	}
	return
}

func (it *MMTypeValue) MarshalJSON() ([]byte, error) {
	switch {
	case it == nil:
		return []byte("null"), nil
	case it.kind == MMTypeValueKindObjType:
		return json.Marshal(it.t)
	case it.kind == MMTypeValueKindObjStructureLiteral:
		return json.Marshal(it.l)
	case it.kind == MMTypeValueKindObj:
		return json.Marshal(it.m)
	case it.kind == MMTypeValueKindArr:
		return json.Marshal(it.a)
	case it.kind == MMTypeValueKindStr:
		return json.Marshal(it.s)
	case it.kind == MMTypeValueKindBool:
		return json.Marshal(it.b)
	case it.kind == MMTypeValueKindFloat:
		return json.Marshal(it.f)
	case it.kind == MMTypeValueKindInt:
		return json.Marshal(it.i)
	default:
		return nil, errors.New("nothing in here for kind " + strconv.Itoa(int(it.kind)))
	}
}

type MMTypes []MMType

func (it *MMTypes) UnmarshalJSON(src []byte) (err error) {
	var dst []MMType
	switch {
	default:
		err = json.Unmarshal(src, &dst)
	case src[0] == '{':
		dst = make([]MMType, 1)
		err = json.Unmarshal(src, &dst[0])
	}
	*it = dst
	return
}

func (it *MMTypes) MarshalJSON() ([]byte, error) {
	if it == nil {
		return []byte("null"), nil
	}
	slice := *it
	if len(slice) == 1 {
		return json.Marshal((slice)[0])
	}
	return json.Marshal((slice))
}
