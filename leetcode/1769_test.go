package leetcode

import (
	"testing"
)

func minOperations(boxes string) []int {
	results := make([]int, len(boxes))
	left := 0
	left = int(boxes[0]) - '0'
	for i := 1; i < len(boxes); i++ {
		results[i] += results[i-1] + left
		if boxes[i] == '1' {
			left++
		}
	}

	lastIndex := len(boxes) - 1
	right := int(boxes[lastIndex]) - '0'
	prevRight := 0
	for i := len(boxes) - 2; i >= 0; i-- {
		prevRight += right
		results[i] += prevRight
		if boxes[i] == '1' {
			right++
		}
	}
	return results
}

func TestMinOperations(t *testing.T) {
	cases := map[string][]int{
		"110":    {1, 1, 3},
		"001011": {11, 8, 5, 4, 3, 4},
	}
	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			assertEqualSlices(t, expected, minOperations(input))
		})
	}
}
