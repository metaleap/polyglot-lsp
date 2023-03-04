package glot

func Contains[T comparable](list []T, needle T) bool {
	for _, item := range list {
		if needle == item {
			return true
		}
	}
	return false
}

func Copy[T any](list []T) (ret []T) {
	ret = make([]T, len(list))
	copy(ret, list)
	return
}

func Filter[T any](list []T, ok func(T) bool) (ret []T) {
	ret = make([]T, 0, len(list))
	for _, item := range list {
		if ok(item) {
			ret = append(ret, item)
		}
	}
	return
}

func Map[T any, R any](list []T, convert func(T) R) (ret []R) {
	ret = make([]R, len(list))
	for i, item := range list {
		ret[i] = convert(item)
	}
	return
}

func Replace[T comparable](list []T, old T, new T) (ret []T) {
	ret = Copy(list)
	for i, item := range ret {
		if item == old {
			ret[i] = new
		}
	}
	return
}

func Without[T comparable](list []T, sans ...T) []T {
	return Filter(list, func(item T) bool {
		return !Contains(sans, item)
	})
}
