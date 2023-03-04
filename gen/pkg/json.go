package glot

import (
	"encoding/json"
	"strconv"
)

type NumberOrString json.Number

func (it *NumberOrString) UnmarshalJSON(src []byte) (err error) {
	if src[0] == '"' {
		var s string
		s, err = strconv.Unquote(string(src))
		*it = NumberOrString(s)
		return
	}

	var n json.Number
	err = json.Unmarshal(src, &n)
	*it = NumberOrString(n)
	return
}

func (it *NumberOrString) MarshalJSON() ([]byte, error) {
	return []byte(it.String()), nil
}

func (it *NumberOrString) String() string {
	if it == nil {
		return "null"
	}
	n := json.Number(*it)
	if i64, ei := n.Int64(); ei == nil {
		return strconv.FormatInt(i64, 10)
	} else if f64, ef := n.Float64(); ef == nil {
		return strconv.FormatFloat(f64, 'g', -1, 64)
	}
	return strconv.Quote(string(*it))
}

func LoadFromJSON[T any](src []byte) T {
	new := new(T)
	if err := json.Unmarshal(src, new); err != nil {
		panic(err)
	}
	return *new
}

func LoadFromJSONFile[T any](filePath string) T {
	src := FileRead(filePath)
	return LoadFromJSON[T](src)
}
