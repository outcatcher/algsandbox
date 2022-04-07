package merges

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

func validateMerge(t *testing.T, src1, src2, res []int) {
	t.Helper()
	t.Logf("Sources: %+v, %+v", src1, src2)
	t.Logf("Result: %+v", res)

	l1, l2, lr := len(src1), len(src2), len(res)
	if l1+l2 != lr {
		t.Fatalf("merged size %d not matching source size %d, %d", lr, l1, l2)
	}

	for i := 0; i < l1; i++ {
		if res[i] != src1[i] {
			t.Fatalf("not same after merge")
		}
	}
	for i := 0; i < l2; i++ {
		if res[i+l1] != src2[i] {
			t.Fatalf("not same after merge (2)")
		}
	}
}

func TestMerge(t *testing.T) {
	s1, s2 := randomSlice(5), randomSlice(5)
	t.Run("copy", func(t *testing.T) {
		r := simpleMerge(s1, s2)
		validateMerge(t, s1, s2, r)
	})
	t.Run("iter", func(t *testing.T) {
		r := iterMerge(s1, s2)
		validateMerge(t, s1, s2, r)
	})
}

type mergeFunc func([]int, []int) []int

func BenchmarkMerge(b *testing.B) {
	max := 5000000
	step := max / 2

	cases := map[string]mergeFunc{
		"copy":          simpleMerge,
		"iter":          iterMerge,
		"append":        appendMerge,
		"append_stupid": appendStupidMerge,
	}
	for name, f := range cases {
		for size := step; size <= max; size += step {
			s1, s2 := randomSlice(size), randomSlice(size)
			b.Run(fmt.Sprintf("%s(%d)", name, size), func(b *testing.B) {
				f(s1, s2)
			})
		}
	}
}
