package glot

import (
	"os"
)

func fsExists(filePath string, dir bool) bool {
	stat, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return stat != nil && stat.IsDir() == dir
}

func DirExists(filePath string) bool {
	return fsExists(filePath, true)
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
