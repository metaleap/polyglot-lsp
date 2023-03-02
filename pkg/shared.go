package glot

import (
	"encoding/json"
	"os"
)

type meta = struct {
	VSX metaEntry
	LSP metaEntry
}

type metaEntry = struct {
	Version string
	File    string
}

var Meta meta

func init() {
	Meta = LoadFromJSONFile[meta]("prereq/meta.json")
}

func ReadFile(filePath string) []byte {
	ret, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return ret
}

func LoadFromJSON[T any](src []byte) T {
	new := new(T)
	if err := json.Unmarshal(src, new); err != nil {
		panic(err)
	}
	return *new
}

func LoadFromJSONFile[T any](filePath string) T {
	src := ReadFile(filePath)
	return LoadFromJSON[T](src)
}
