// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-vsx-and-lsp/gen/cmd/gen_lsp
package lsp

import "net/url"

type DocumentURI string

func (it DocumentURI) URL() (*url.URL, error) { return url.Parse(string(it)) }

type URI string

func (it URI) URL() (*url.URL, error) { return url.Parse(string(it)) }

type String string

func (it *String) IsNone() bool { return it == nil || *it == "" }
func (it *String) IfNone(ifNone string) string {
	if it.IsNone() {
		return ifNone
	}
	return string(*it)
}
func (it *String) IfNil(ifNil string) string {
	if it == nil {
		return ifNil
	}
	return string(*it)
}

type Integer int

func (it *Integer) IsNone() bool { return it == nil || *it == 0 }
func (it *Integer) IfNone(ifNone int) int {
	if it.IsNone() {
		return ifNone
	}
	return int(*it)
}
func (it *Integer) IfNil(ifNil int) int {
	if it == nil {
		return ifNil
	}
	return int(*it)
}

type Uinteger uint

func (it *Uinteger) IsNone() bool { return it == nil || *it == 0 }
func (it *Uinteger) IfNone(ifNone uint) uint {
	if it.IsNone() {
		return ifNone
	}
	return uint(*it)
}
func (it *Uinteger) IfNil(ifNil uint) uint {
	if it == nil {
		return ifNil
	}
	return uint(*it)
}

type Decimal float64

func (it *Decimal) IsNone() bool {
	return it == nil || *it == 0.0 /*|| *it == math.SmallestNonzeroFloat64*/
}
func (it *Decimal) IfNone(ifNone float64) float64 {
	if it.IsNone() {
		return ifNone
	}
	return float64(*it)
}
func (it *Decimal) IfNil(ifNil float64) float64 {
	if it == nil {
		return ifNil
	}
	return float64(*it)
}

type Boolean bool

func (it *Boolean) IsNone() bool { return it == nil || !(bool)(*it) }
func (it *Boolean) IfNone(ifNone bool) bool {
	if it.IsNone() {
		return ifNone
	}
	return bool(*it)
}
func (it *Boolean) IfNil(ifNil bool) bool {
	if it == nil {
		return ifNil
	}
	return bool(*it)
}
