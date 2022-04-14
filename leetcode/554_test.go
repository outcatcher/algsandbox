package leetcode

import (
	"testing"
)

// There is a rectangular brick wall in front of you with n rows of bricks.
// The ith row has some number of bricks each of the same height (i.e., one unit) but they can be of different widths.
// The total width of each row is the same.
//
// Draw a vertical line from the top to the bottom and cross the least bricks. If your line goes through the edge of a
// brick, then the brick is not considered as crossed. You cannot draw a line just along one of the two vertical edges
// of the wall, in which case the line will obviously cross no bricks.
//
// Given the 2D array wall that contains the information about the wall, return the minimum number of crossed bricks
// after drawing such a vertical line.

func calcWidth(wall [][]int) int {
	if len(wall) < 1 {
		panic("this is not a wall")
	}

	w := 0
	// this width is fixed
	for _, b := range wall[0] {
		w += b
	}
	return w
}

func leastBricks(wall [][]int) int {
	width := calcWidth(wall)
	lines := make(map[int]int) // 0 is the worst case here

	for _, row := range wall {
		shift := 0
		for _, brick := range row {
			shift += brick
			if shift == width { // end the row
				continue
			}
			lines[shift-1]++ // there is a gap here, increase gap count
		}
	}

	mostGaps := 0
	for _, v := range lines {
		if v > mostGaps {
			mostGaps = v
		}
	}
	return len(wall) - mostGaps
}

func TestTask(t *testing.T) {
	cases := map[int][][]int{
		3: {{1}, {1}, {1}},
		2: {{1, 2, 2, 1}, {3, 1, 2}, {1, 3, 2}, {2, 4}, {3, 1, 2}, {1, 3, 1, 1}},
	}
	for expected, wall := range cases {
		t.Run("", func(t *testing.T) {
			assertEqual(t, expected, leastBricks(wall))
		})
	}
}
