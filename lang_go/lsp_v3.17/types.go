// Language Server Protocol (LSP) v3.17 SDK for Go: auto-generated via github.com/metaleap/polyglot-vsx-and-lsp/gen/cmd/gen_lsp
package lsp

import "net/url"

type DocumentURI string
type URI string

func (it DocumentURI) URL() (*url.URL, error) { return url.Parse(string(it)) }
func (it URI) URL() (*url.URL, error)         { return url.Parse(string(it)) }
