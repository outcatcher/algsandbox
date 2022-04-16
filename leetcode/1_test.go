package leetcode

import "testing"

// Given an array of integers nums and an integer target, return indices of the two numbers such that
// they add up to target.
//
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
// You can return the answer in any order.

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int) // number: position
	for i, v := range nums {
		if j, ok := hash[target-v]; ok {
			return []int{j, i}
		}
		hash[v] = i
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		output []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			assertEqualSlices(t, c.output, twoSum(c.nums, c.target))
		})
	}
}
