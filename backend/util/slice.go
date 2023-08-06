package util

func All[T comparable](f func(T) bool, ary []T) bool {
	for _, v := range ary {
		if !f(v) {
			return false
		}
	}
	return true
}

func Map[T comparable](f func(T) T, ary []T) []T {
	for i, _ := range ary {
		ary[i] = f(ary[i])
	}

	return ary
}

func Remove[T comparable](i int, ary []T) []T {
	return ary[:i+copy(ary[i:], ary[i+1:])]
}
