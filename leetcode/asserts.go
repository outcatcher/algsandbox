package leetcode

import "testing"

func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()

	if expected != actual {
		t.Fatalf("Expected %v (%[1]T), got %v (%[2]T)", expected, actual)
	}
}
func assertEqualSlices[T comparable](t *testing.T, expected, actual []T) {
	t.Helper()

	for i, exp := range expected {
		if actual[i] != exp {
			t.Fatalf("Expected slice to be %+v but got %+v", expected, actual)
		}
	}
}
