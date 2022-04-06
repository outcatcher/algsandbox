package qsort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	src := randomSlice(10)
	fmt.Printf("Before: %+v (%[1]T)\n", src)
	res := quickParallel(src)
	fmt.Printf("Q Parallel: %+v (%[1]T)\n", res)
	res2 := quick(src)
	fmt.Printf("Q: %+v (%[1]T)\n", res2)
}

var bRes = randomSlice(100000)

func BenchmarkSortMP(b *testing.B) {
	d := bRes[:]
	quickParallel(d)
}
func BenchmarkSortM(b *testing.B) {
	d := bRes[:]
	quick(d)
}
