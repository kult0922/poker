package util

func All[T comparable](f func(T) bool, ary []T) bool {
	for _, v := range ary {
		if !f(v) {
			return false
		}
	}
	return true
}
