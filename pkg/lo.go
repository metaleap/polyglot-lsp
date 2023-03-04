package glot

func Map[T any, R any](list []T, convert func(T) R) (ret []R) {
	ret = make([]R, len(list))
	for i, item := range list {
		ret[i] = convert(item)
	}
	return
}
