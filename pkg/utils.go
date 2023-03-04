package glot

import (
	"strings"
)

type Tup[T1 any, T2 any] struct {
	F1 T1
	F2 T2
}

func Cap(s string) string {
	if len(s) > 0 {
		s = strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}
