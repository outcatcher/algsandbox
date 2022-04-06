package qsort

import "math/rand"

func randomSlice(size int) []int {
	res := make([]int, 0, size)
	for len(res) < size {
		res = append(res, rand.Intn(20))
	}
	return res
}
