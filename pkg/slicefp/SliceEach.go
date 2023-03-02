package slicefp

import "sync"

func SliceEach[T interface{}](in []T, eachFn func(T)) {
	for _, elem := range in {
		eachFn(elem)
	}
}

func SliceEachWG[T interface{}](wg *sync.WaitGroup, in []T, eachFn func(T)) {
	defer wg.Done()
	SliceEach(in, eachFn)
}

func SliceEachCoro[T interface{}](in []T, eachFn func(T)) {
	var wg sync.WaitGroup
	wg.Add(len(in))

	wrappedEachFn := func(i int, elem T) {
		defer wg.Done()
		eachFn(elem)
	}

	for i, elem := range in {
		go wrappedEachFn(i, elem)
	}

	wg.Wait()
}

func SliceEachCoroWG[T interface{}](wg *sync.WaitGroup, in []T, eachFn func(T)) {
	defer wg.Done()
	SliceEachCoro(in, eachFn)
}

func SliceEachCoroIdx[T interface{}](in []T, eachFn func(T, int)) {
	var wg sync.WaitGroup
	wg.Add(len(in))

	wrappedEachFn := func(i int, elem T) {
		defer wg.Done()
		eachFn(elem, i)
	}

	for i, elem := range in {
		go wrappedEachFn(i, elem)
	}

	wg.Wait()
}

func SliceEachCoroIdxWG[T interface{}](wg *sync.WaitGroup, in []T, eachFn func(T, int)) {
	defer wg.Done()
	SliceEachCoroIdx(in, eachFn)
}
