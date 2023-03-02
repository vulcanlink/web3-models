package slicefp

import "sync"

func SliceFilter[T interface{}](in []T, filterFn func(T) bool) []T {
	var out []T
	for _, elem := range in {
		if filterFn(elem) {
			out = append(out, elem)
		}
	}

	return out
}

func SliceFilterWG[T interface{}](wg *sync.WaitGroup, in []T, filterFn func(T) bool) []T {
	defer wg.Done()
	return SliceFilter(in, filterFn)
}

// NOTE: appends are not thread-safe, so you cannot have a Coro implementation similar to SliceMapCoro
