package glot

import (
	"io/fs"
	"os"
	"strconv"
	"strings"
	"time"
)

const SchemasDir = "prereq"

var trashDir = "/tmp/" + strconv.FormatInt(time.Now().UnixNano(), 36)

func init()         { DirCreate(trashDir) }
func FinalCleanUp() { DirRemove(trashDir) }

func Vers(fileNamePrefix string, fileNameSuffix string) (ret []string) {
	if override := os.Getenv("glot_vers"); override != "" {
		return strings.Split(override, ",")
	}
	return Dir(SchemasDir, func(entry fs.DirEntry, path string) (string, bool) {
		name := entry.Name()
		return strings.TrimPrefix(strings.TrimSuffix(name, fileNameSuffix), fileNamePrefix),
			(!entry.IsDir()) && strings.HasPrefix(name, fileNamePrefix) && strings.HasSuffix(name, fileNameSuffix)
	})
}

func Langs() (ret []string) {
	if override := os.Getenv("glot_langs"); override != "" {
		return strings.Split(override, ",")
	}
	return Dir("..", func(entry fs.DirEntry, path string) (string, bool) {
		name := entry.Name()
		return strings.TrimPrefix(name, "lang_"),
			entry.IsDir() && strings.HasPrefix(name, "lang_") && FileExists(path+"/"+name+".json") && DirExists(path+"/_gen")
	})
}
