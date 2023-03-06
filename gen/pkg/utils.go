package glot

import (
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type Tup[T1 any, T2 any] struct {
	F1 T1
	F2 T2
}

func self[T any](it T) T { return it }

func CmpEq(a any, b any, equateEmpty bool, ignoreUnexported bool) bool {
	var opts []cmp.Option
	if equateEmpty {
		opts = append(opts, cmpopts.EquateEmpty())
	}
	if ignoreUnexported {
		opts = append(opts, cmpopts.IgnoreUnexported())
	}
	return cmp.Equal(a, b, opts...)
}

func If[T any](b bool, ifTrue T, ifFalse T) T {
	if b {
		return ifTrue
	}
	return ifFalse
}

func Up0(s string) string {
	if len(s) > 0 && !(s[0] >= 'A' && s[0] <= 'Z') {
		s = strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}

func ValueString(v any) string {
	if v == nil {
		return ""
	}
	switch it := v.(type) {
	case bool:
		return strconv.FormatBool(it)
	case string:
		return strconv.Quote(it)
	case uint:
		return strconv.FormatUint(uint64(it), 64)
	case uint8:
		return strconv.FormatUint(uint64(it), 8)
	case uint16:
		return strconv.FormatUint(uint64(it), 16)
	case uint32:
		return strconv.FormatUint(uint64(it), 32)
	case uint64:
		return strconv.FormatUint(it, 64)
	case int:
		return strconv.FormatInt(int64(it), 64)
	case int8:
		return strconv.FormatInt(int64(it), 8)
	case int16:
		return strconv.FormatInt(int64(it), 16)
	case int32:
		return strconv.FormatInt(int64(it), 32)
	case int64:
		return strconv.FormatInt(it, 64)
	case float32:
		return strconv.FormatFloat(float64(it), 'g', -1, 32)
	case float64:
		return strconv.FormatFloat(it, 'g', -1, 32)
	case *NumberOrString:
		return it.String()
	case NumberOrString:
		return it.String()
	}
	panic(v)
}
