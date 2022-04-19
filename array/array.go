package array

func Filter[T any](src []T, fn func(T) bool) []T {
	var ret []T
	for _, v := range src {
		if match := fn(v); match {
			ret = append(ret, v)
		}
	}
	return ret
}

func Map[T any, M any](src []T, fn func(T) M) []M {
	var ret []M
	for _, v := range src {
		ret = append(ret, fn(v))
	}
	return ret
}

func Remove[T comparable](src []T, tgt T) []T {
	for i, v := range src {
		if v == tgt {
			return src[:i+copy(src[i:], src[i+1:])]
		}
	}
	return src
}
