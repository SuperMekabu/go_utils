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

func Includes[T comparable](src []T, tgt T) bool {
	for _, v := range src {
		if v == tgt {
			return true
		}
	}
	return false
}

func RemoveFirst[T comparable](src []T, tgt T) []T {
	for i, v := range src {
		if v == tgt {
			return src[:i+copy(src[i:], src[i+1:])]
		}
	}
	return src
}

func RemoveAll[T comparable](src []T, tgt T) []T {
	tmp := src
	found := false
	for {
		found = false
		for i, v := range tmp {
			if v == tgt {
				tmp = tmp[:i+copy(src[i:], tmp[i+1:])]
				found = true
				break
			}
		}
		if !found {
			break
		}
	}
	return tmp
}

func Every[T any](src []T, fn func(T) bool) bool {
	for _, v := range src {
		if !fn(v) {
			return false
		}
	}
	return true
}

func Some[T any](src []T, fn func(T) bool) bool {
	for _, v := range src {
		if fn(v) {
			return true
		}
	}
	return false
}
