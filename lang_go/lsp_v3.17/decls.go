// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-lsp/gen/cmd/gen_lsp
package lsp

import "net/url"

type DocumentURI = URI
type URI = String

// URL is a convenience `url.Parse(string(it))` call that returns `nil` on error.
// If you do need the `error`, just use the original `url.Parse` directly.
func (it *URI) URL() (ret *url.URL) {
	if it.IsNone() {
		return nil
	}
	ret, _ = url.Parse(string(*it))
	return
}

type String string

// IsNone returns whether `it` is `nil` or the `string` zero value.
func (it *String) IsNone() bool { return it == nil || *it == "" }

// IfNone returns `ifNone` if `it` is `nil` or the `string` zero value; otherwise, `it.Value()` is returned.
func (it *String) IfNone(ifNone string) string {
	if it.IsNone() {
		return ifNone
	}
	return string(*it)
}

// Value returns the `string` zero value if `it` is `nil`; otherwise `*it` is returned.
func (it *String) Value() string { return it.IfNil("") }

// IfNil returns `ifNil` if `it` is `nil`; otherwise, `it.Value()` is returned.
func (it *String) IfNil(ifNil string) string {
	if it == nil {
		return ifNil
	}
	return string(*it)
}

type Integer int

// IsNone returns whether `it` is `nil` or the `int` zero value.
func (it *Integer) IsNone() bool { return it == nil || *it == 0 }

// IfNone returns `ifNone` if `it` is `nil` or the `int` zero value; otherwise, `it.Value()` is returned.
func (it *Integer) IfNone(ifNone int) int {
	if it.IsNone() {
		return ifNone
	}
	return int(*it)
}

// Value returns the `int` zero value if `it` is `nil`; otherwise `*it` is returned.
func (it *Integer) Value() int { return it.IfNil(0) }

// IfNil returns `ifNil` if `it` is `nil`; otherwise, `it.Value()` is returned.
func (it *Integer) IfNil(ifNil int) int {
	if it == nil {
		return ifNil
	}
	return int(*it)
}

type Uinteger uint

// IsNone returns whether `it` is `nil` or the `uint` zero value.
func (it *Uinteger) IsNone() bool { return it == nil || *it == 0 }

// IfNone returns `ifNone` if `it` is `nil` or the `uint` zero value; otherwise, `it.Value()` is returned.
func (it *Uinteger) IfNone(ifNone uint) uint {
	if it.IsNone() {
		return ifNone
	}
	return uint(*it)
}

// Value returns the `uint` zero value if `it` is `nil`; otherwise `*it` is returned.
func (it *Uinteger) Value() uint { return it.IfNil(0) }

// IfNil returns `ifNil` if `it` is `nil`; otherwise, `it.Value()` is returned.
func (it *Uinteger) IfNil(ifNil uint) uint {
	if it == nil {
		return ifNil
	}
	return uint(*it)
}

type Decimal float64

// IsNone returns whether `it` is `nil` or the `float64` zero value.
func (it *Decimal) IsNone() bool {
	return it == nil || *it == 0.0 /*|| *it == math.SmallestNonzeroFloat64*/
}

// IfNone returns `ifNone` if `it` is `nil` or the `float64` zero value; otherwise, `it.Value()` is returned.
func (it *Decimal) IfNone(ifNone float64) float64 {
	if it.IsNone() {
		return ifNone
	}
	return float64(*it)
}

// Value returns the `float64` zero value if `it` is `nil`; otherwise `*it` is returned.
func (it *Decimal) Value() float64 { return it.IfNil(0.0) }

// IfNil returns `ifNil` if `it` is `nil`; otherwise, `it.Value()` is returned.
func (it *Decimal) IfNil(ifNil float64) float64 {
	if it == nil {
		return ifNil
	}
	return float64(*it)
}

type Boolean bool

// IsNone returns whether `it` is `nil` or the `bool` zero value.
func (it *Boolean) IsNone() bool { return it == nil || !(bool)(*it) }

// IfNone returns `ifNone` if `it` is `nil` or the `bool` zero value; otherwise, `it.Value()` is returned.
func (it *Boolean) IfNone(ifNone bool) bool {
	if it.IsNone() {
		return ifNone
	}
	return bool(*it)
}

// Value returns the `bool` zero value if `it` is `nil`; otherwise `*it` is returned.
func (it *Boolean) Value() bool { return it.IfNil(false) }

// IfNil returns `ifNil` if `it` is `nil`; otherwise, `it.Value()` is returned.
func (it *Boolean) IfNil(ifNil bool) bool {
	if it == nil {
		return ifNil
	}
	return bool(*it)
}
