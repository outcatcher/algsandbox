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

func appendMerge(src1, src2 []int) []int {
	return append(src1, src2...)
}

func appendStupidMerge(src1, src2 []int) []int {
	var res []int
	for s1 := range src1 {
		res = append(res, s1)
	}
	for s2 := range src2 {
		res = append(res, s2)
	}
	return res
}
