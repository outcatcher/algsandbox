package leetcode

import (
	"testing"
)

func minOperations(boxes string) []int {
	results := make([]int, len(boxes))
	left := 0
	prevLeft := 0
	for i := 0; i < len(boxes); i++ {
		prevLeft += left
		results[i] = prevLeft
		left += int(boxes[i]) - '0'
	}

	right := int(boxes[len(boxes)-1]) - '0'
	prevRight := 0
	for i := len(boxes) - 2; i >= 0; i-- {
		prevRight += right
		results[i] += prevRight
		right += int(boxes[i]) - '0'
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
