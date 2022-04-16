package leetcode

import (
	"fmt"
	"testing"
)

func minNumberOperations(target []int) int {
	operations := 0
	prev := 0
	for i := 0; i < len(target); i++ {
		this := target[i]
		if add := this - prev; add > 0 {
			operations += add
		}
		prev = this
	}
	return operations
}

func TestMinNumber(t *testing.T) {
	cases := map[int][]int{
		3: {1, 2, 3, 2, 1},
		4: {3, 1, 1, 2},
		7: {3, 1, 5, 4, 2},
	}

	for expected, v := range cases {
		t.Run(fmt.Sprintf("minimal-%d", expected), func(t *testing.T) {
			count := minNumberOperations(v)
			assertEqual(t, expected, count)
		})
	}
}
