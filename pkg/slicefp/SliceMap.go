package slicefp

import "sync"

func SliceMap[T interface{}, U interface{}](in []T, mapFn func(T) U) []U {
	out := make([]U, len(in))
	for i, elem := range in {
		out[i] = mapFn(elem)
	}

	return out
}

func SliceMapWG[T interface{}, U interface{}](wg *sync.WaitGroup, in []T, mapFn func(T) U) []U {
	defer wg.Done()
	return SliceMap(in, mapFn)
}

func SliceMapCoro[T interface{}, U interface{}](in []T, mapFn func(T) U) []U {
	out := make([]U, len(in))
	var wg sync.WaitGroup
	wg.Add(len(in))

	wrappedMapFn := func(i int, elem T) {
		defer wg.Done()
		out[i] = mapFn(elem)
	}

	for i, elem := range in {
		go wrappedMapFn(i, elem)
	}

	wg.Wait()
	return out
}

func SliceMapCoroWG[T interface{}, U interface{}](wg *sync.WaitGroup, in []T, mapFn func(T) U) []U {
	defer wg.Done()
	return SliceMapCoro(in, mapFn)
}

func SliceMapCoroIdx[T interface{}, U interface{}](in []T, mapFn func(T, int) U) []U {
	out := make([]U, len(in))
	var wg sync.WaitGroup
	wg.Add(len(in))

	wrappedMapFn := func(i int, elem T) {
		defer wg.Done()
		out[i] = mapFn(elem, i)
	}

	for i, elem := range in {
		go wrappedMapFn(i, elem)
	}

	wg.Wait()
	return out
}

func SliceMapCoroIdxWG[T interface{}, U interface{}](wg *sync.WaitGroup, in []T, mapFn func(T, int) U) []U {
	defer wg.Done()
	return SliceMapCoroIdx(in, mapFn)
}
