package slicefp_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vulcanlink/web3-models/pkg/slicefp"
)

func TestSliceFilter(t *testing.T) {
	in := []int{1, 5, 3, 9}
	filterFn := func(item int) bool {
		return item > 3
	}

	out := slicefp.SliceFilter(in, filterFn)
	assert.Equal(t, []int{5, 9}, out, "slice is filtered")
}

func TestSliceFilterWG(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	in := []int{1, 5, 3, 9}
	filterFn := func(item int) bool {
		return item > 3
	}

	out := slicefp.SliceFilterWG(&wg, in, filterFn)
	wg.Wait()

	assert.Equal(t, []int{5, 9}, out, "slice is filtered")
}
