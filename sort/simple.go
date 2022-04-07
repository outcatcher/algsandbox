package sort

func stupid(src []int) {
	i := 0
	for i < len(src)-1 {
		if src[i] > src[i+1] {
			src[i], src[i+1] = src[i+1], src[i]
			i = 0
		} else {
			i++
		}
	}
}

func bubble(src []int) {
	end := len(src) - 1
	for end > 0 {
		for i := 0; i < end; i++ {
			if src[i] > src[i+1] {
				src[i], src[i+1] = src[i+1], src[i]
			}
		}
		end--
	}
}

func shaker(src []int) {
	end := len(src) - 1
	start := 0
	for end > start {
		for i := start; i < end; i++ {
			if src[i] > src[i+1] {
				src[i], src[i+1] = src[i+1], src[i]
			}
		}
		end--
		for i := end; i > start; i-- {
			if src[i] < src[i-1] {
				src[i], src[i-1] = src[i-1], src[i]
			}
		}
		start++
	}
}

func minPosition(src []int) int {
	p := 0
	for i := 1; i < len(src); i++ {
		if src[i] < src[p] {
			p = i
		}
	}
	return p
}

func selection(src []int) {
	for start := 0; start < len(src); start++ {
		pos := minPosition(src[start:]) + start
		src[start], src[pos] = src[pos], src[start]
	}
}

func insertion(src []int) {
	result := make([]int, len(src))
	i := 0
	for start := 0; start < len(src); start++ {
		target := src[start:]
		result[i] = target[minPosition(target)]
		i++
	}
}
