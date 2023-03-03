package glot

import (
	"os"
	"strings"
)

const SchemasDirName = "prereq"

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

func Langs() (ret []string) {
	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if name := entry.Name(); entry.IsDir() && strings.HasPrefix(name, "lang_") &&
			FileExists(name+"/"+name+".json") && DirExists(name+"/gen") {
			ret = append(ret, name[len("lang_"):])
		}
	}
	return
}
