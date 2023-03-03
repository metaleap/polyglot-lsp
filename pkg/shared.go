package glot

import (
	"os"
	"strings"
)

const SchemasDirName = "prereq"

func init() {
}

type Tup[T1 any, T2 any] struct {
	F1 T1
	F2 T2
}

func Versions(fileNamePrefix string, fileNameSuffix string) (ret []string) {
	entries, err := os.ReadDir(SchemasDirName)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if name := entry.Name(); (!entry.IsDir()) && strings.HasPrefix(name, fileNamePrefix) && strings.HasSuffix(name, fileNameSuffix) {
			ret = append(ret, strings.TrimPrefix(strings.TrimSuffix(name, fileNameSuffix), fileNamePrefix))
		}
	}
	return
}

func ReadFile(filePath string) []byte {
	ret, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return ret
}
