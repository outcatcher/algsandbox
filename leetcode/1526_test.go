package leetcode

import (
	"fmt"
	"testing"
)

func minNumberOperations(target []int) int {
	size := len(target)

	count := 0
	for {
		subsets := 0
		lastMoved := false
		for i := 0; i < size; i++ {
			v := target[i]
			if v == 0 {
				if lastMoved {
					subsets++
				}
				lastMoved = false
				continue
			}
			lastMoved = true
			target[i] = v - 1
		}
		if lastMoved { // when the last element is moved, there is at least one subset
			subsets++
		}
		if subsets == 0 { // no changes to the list
			break
		}
		count += subsets
		fmt.Printf("Moved to %v in %d steps\n", target, subsets)
	}
	fmt.Println()
	return count
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
