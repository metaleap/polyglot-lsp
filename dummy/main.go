package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"strconv"

	_ "github.com/metaleap/polyglot-lsp/lang_go/lsp_v3.17"
)

func main() {
	stdin, _ := setupJsonIpcPipes(1024 * 1024)
	for stdin.Scan() {
		jsonin := stdin.Bytes()
		m := map[string]any{}
		err := json.Unmarshal(jsonin, &m)
		if err != nil {
			panic(err)
		}

	}
}

func setupJsonIpcPipes(bufferCapacity int) (stdin *bufio.Scanner, stdout *bufio.Writer) {
	stdin = bufio.NewScanner(os.Stdin)
	stdin.Buffer(make([]byte, bufferCapacity), bufferCapacity)
	stdin.Split(func(data []byte, ateof bool) (advance int, token []byte, err error) {
		if i_cl1 := bytes.Index(data, []byte("Content-Length: ")); i_cl1 >= 0 {
			datafromclen := data[i_cl1+16:]
			if i_cl2 := bytes.IndexAny(datafromclen, "\r\n"); i_cl2 > 0 {
				if clen, e := strconv.Atoi(string(datafromclen[:i_cl2])); e != nil {
					err = e
				} else if i_js1 := bytes.Index(datafromclen, []byte("{\"")); i_js1 > i_cl2 {
					if i_js2 := i_js1 + clen; len(datafromclen) >= i_js2 {
						advance = i_cl1 + 16 + i_js2
						token = datafromclen[i_js1:i_js2]
					}
				}
			}
		}
		return
	})
	stdout = bufio.NewWriterSize(os.Stdout, bufferCapacity)
	return
}
