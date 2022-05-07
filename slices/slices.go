package slices

func Chunk[T any](org []T, chunkSize int) [][]T {
	chunkedLength := func(len, size int) int {
		if size == 1 {
			return len / size
		}
		return len/size + 1
	}(len(org), chunkSize)

	chunked := make([][]T, chunkedLength)
	for i := 0; i < len(chunked); i++ {
		tail := (i + 1) * chunkSize
		if tail > len(org) {
			chunked[i] = org[i*chunkSize:]
		} else {
			chunked[i] = org[i*chunkSize : tail]
		}
	}
	return chunked
}

func Filter[T any](elms []T, fn func(T) bool) []T {
	var ret []T
	for _, v := range elms {
		if match := fn(v); match {
			ret = append(ret, v)
		}
	}
	return ret
}

func Map[T any, R any](elms []T, fn func(T) (R, bool)) []R {
	var ret []R
	for _, v := range elms {
		nv, Ok := fn(v)
		if Ok {
			ret = append(ret, nv)
		}
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
