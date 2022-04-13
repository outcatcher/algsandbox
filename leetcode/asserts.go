package leetcode

import "testing"

func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()

	if expected != actual {
		t.Fatalf("Expected %v (%[1]T), got %v (%[2]T)", expected, actual)
	}
}
