package leetcode

import (
	"fmt"
	"testing"
)

// You are given an integer n, the number of teams in a tournament that has strange rules:
//
//    If the current number of teams is even, each team gets paired with another team. A total of n / 2 matches are
//    played, and n / 2 teams advance to the next round.
//    If the current number of teams is odd, one team randomly advances in the tournament, and the rest gets paired.
//    A total of (n - 1) / 2 matches are played, and (n - 1) / 2 + 1 teams advance to the next round.
//
// Return the number of matches played in the tournament until a winner is decided.
//
//

// that was so easy that I haven't noticed that
func numberOfMatches(n int) int {
	return n - 1
}

func TestMatches(t *testing.T) {
	cases := map[int]int{
		14: 13,
		7:  6,
		3:  2,
	}
	for input, expected := range cases {
		t.Run(fmt.Sprintf("For %d", input), func(t *testing.T) {
			assertEqual(t, expected, numberOfMatches(input))
		})
	}
}
