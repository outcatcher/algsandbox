package merges

func simpleMerge(src1, src2 []int) []int {
	res := make([]int, len(src1)+len(src2))
	copy(res, src1)
	copy(res[len(src1):], src2)
	return res
}

func iterMerge(src1, src2 []int) []int {
	res := make([]int, len(src1)+len(src2))
	for i := 0; i < len(src1); i++ {
		res[i] = src1[i]
	}
	for i := 0; i < len(src2); i++ {
		res[i+len(src1)] = src2[i]
	}
	return res
}
