package glot

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func fsExists(path string, dir bool) bool {
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return stat != nil && stat.IsDir() == dir
}

func Dir[T any](dirPath string, maybe func(entry fs.DirEntry, path string) (T, bool)) (ret []T) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if entry, ok := maybe(entry, filepath.Join(dirPath, entry.Name())); ok {
			ret = append(ret, entry)
		}
	}
	return
}

func DirCreate(dirPath string) {
	DirRemove(dirPath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		panic(err)
	}
}

func DirExists(dirPath string) bool {
	return fsExists(dirPath, true)
}

func DirRemove(dirPath string) {
	if err := os.RemoveAll(dirPath); err != nil && !os.IsNotExist(err) {
		panic(err)
	}
}

func FileExists(filePath string) bool {
	return fsExists(filePath, false)
}

func FileRead(filePath string) []byte {
	ret, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return ret
}

func FileWrite(filePath string, data []byte) {
	if err := os.WriteFile(filePath, data, os.ModePerm); err != nil {
		panic(err)
	}
}

func FilesIn(dirPath string, prefix string, suffix string) (ret []string) {
	return Dir(dirPath, func(entry fs.DirEntry, path string) (string, bool) {
		name := entry.Name()
		return path, (!entry.IsDir()) && (prefix == "" || strings.HasPrefix(name, prefix)) && (suffix == "" || strings.HasSuffix(name, suffix))
	})
}
