package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func randomSlice(size int) []int {
	res := make([]int, 0, size)
	for len(res) < size {
		res = append(res, rand.Intn(20))
	}
	return res
}

type sortFunc func([]int)

func checkSort(t *testing.T, src []int, f sortFunc) {
	t.Helper()

	before := make([]int, len(src))
	cnt := copy(before, src)
	if cnt != len(src) {
		t.Fatalf("error copying source list")
	}
	t.Logf("Before sort: %+v (%[1]T)\n", before)
	f(before)
	t.Logf("After sort: %+v (%[1]T)\n", before)

	prev := before[0]
	for i := 1; i < len(before); i++ {
		if v := before[i]; v < prev {
			t.Fatalf("Array %+v is not sorted: %d is misplaced", before, v)
		}
		prev = before[i]
	}
}

const (
	benchSize  = 10000
	iterations = 3
)

func benchSort(b *testing.B, name string, f sortFunc, maxSize int) {
	b.Helper()
	step := maxSize / iterations
	for i := step; i <= maxSize; i += maxSize / iterations {
		data := randomSlice(i)
		b.Run(fmt.Sprintf("%s (%d)", name, i), func(b *testing.B) {
			f(data)
		})
	}
}
