package slices

func Filter[T any](elms []T, fn func(T) bool) []T {
	var ret []T
	for _, v := range elms {
		if match := fn(v); match {
			ret = append(ret, v)
		}
	}
	return ret
}

func Map[T any, R any](elms []T, fn func(T) R) []R {
	var ret []R
	for _, v := range elms {
		ret = append(ret, fn(v))
	}
	return ret
}

func Includes[T comparable](elms []T, tgt T) bool {
	for _, v := range elms {
		if v == tgt {
			return true
		}
	}
	return false
}

func RemoveFirst[T comparable](elms []T, tgt T) []T {
	for i, v := range elms {
		if v == tgt {
			return elms[:i+copy(elms[i:], elms[i+1:])]
		}
	}
	return elms
}

func RemoveAll[T comparable](elms []T, tgt T) []T {
	tmp := elms
	found := false
	for {
		found = false
		for i, v := range tmp {
			if v == tgt {
				tmp = tmp[:i+copy(elms[i:], tmp[i+1:])]
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

func Every[T any](elms []T, fn func(T) bool) bool {
	for _, v := range elms {
		if !fn(v) {
			return false
		}
	}
	return true
}

func Some[T any](elms []T, fn func(T) bool) bool {
	for _, v := range elms {
		if fn(v) {
			return true
		}
	}
	return false
}
