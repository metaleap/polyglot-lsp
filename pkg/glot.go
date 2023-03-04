package glot

import (
	"io/fs"
	"strings"
)

const SchemasDirName = "prereq"

func Versions(fileNamePrefix string, fileNameSuffix string) (ret []string) {
	return Dir(SchemasDirName, func(entry fs.DirEntry, path string) (string, bool) {
		name := entry.Name()
		return strings.TrimPrefix(strings.TrimSuffix(name, fileNameSuffix), fileNamePrefix),
			(!entry.IsDir()) && strings.HasPrefix(name, fileNamePrefix) && strings.HasSuffix(name, fileNameSuffix)
	})
}

func Langs() (ret []string) {
	return Dir(".", func(entry fs.DirEntry, path string) (string, bool) {
		name := entry.Name()
		return strings.TrimPrefix(name, "lang_"),
			entry.IsDir() && strings.HasPrefix(name, "lang_") && FileExists(name+"/"+name+".json") && DirExists(name+"/_gen")
	})
}
