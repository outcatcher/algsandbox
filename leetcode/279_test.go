package leetcode

import (
	"fmt"
	"testing"
)

// Given an integer n, return the least number of perfect square numbers that sum to n.
//
// A perfect square is an integer that is the square of an integer;
// in other words, it is the product of some integer with itself.
// For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.

func numSquares(n int) int {
	solutions := make([]uint16, n+1) // n+1 for easier addressing
	solutions[0] = 0
	for i := 1; i < len(solutions); i++ {
		solutions[i] = uint16(n)
	}

	for target := uint16(1); target <= uint16(n); target++ {
		for left := uint16(1); left*left <= target; left++ {
			local := solutions[target-left*left] + 1
			if local < solutions[target] {
				solutions[target] = local
			}
		}
	}

	return int(solutions[n])
}

func TestMegaOptimized(t *testing.T) {
	i := 12
	fmt.Println(numSquares(i))
}
