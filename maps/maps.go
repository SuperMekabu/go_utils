package maps

func Filter[K comparable, V any](elms map[K]V, fn func(K, V) bool) map[K]V {
	ret := make(map[K]V)
	for k, v := range elms {
		if match := fn(k, v); match {
			ret[k] = v
		}
	}
	return ret
}

func Map[K comparable, V any, R any](elms map[K]V, fn func(K, V) (R, bool)) []R {
	var ret []R
	for k, v := range elms {
		nv, Ok := fn(k, v)
		if Ok {
			ret = append(ret, nv)
		}
	}
	return ret
}

func HasKey[K comparable, V any](elms map[K]V, key K) bool {
	for k := range elms {
		if k == key {
			return true
		}
	}
	return false
}

func HasValue[K comparable, V comparable](elms map[K]V, key V) bool {
	for _, v := range elms {
		if v == key {
			return true
		}
	}
	return false
}

func Remove[K comparable, V any](elms map[K]V, key K) map[K]V {
	ret := make(map[K]V, len(elms)-1)
	for k, v := range elms {
		if k == key {
			continue
		}
		ret[k] = v
	}
	return ret
}

func Every[K comparable, V any](elms map[K]V, fn func(K, V) bool) bool {
	for k, v := range elms {
		if !fn(k, v) {
			return false
		}
	}
	return true
}

func Some[K comparable, V any](elms map[K]V, fn func(K, V) bool) bool {
	for k, v := range elms {
		if fn(k, v) {
			return true
		}
	}
	return false
}
