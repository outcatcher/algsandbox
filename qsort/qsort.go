package qsort

func quickParallel(src []int) []int {
	if len(src) < 2 {
		// fmt.Println(src)
		return src
	}
	// fmt.Println(src)
	base := src[0]
	var less []int
	var more []int
	for i := 1; i < len(src); i++ {
		v := src[i]
		if v < base {
			less = append(less, v)
		} else {
			more = append(more, v)
		}
	}

	lchan := make(chan []int, 1)
	mchan := make(chan []int, 1)
	go func() {
		lchan <- quickParallel(less)
	}()
	go func() {
		mchan <- quickParallel(more)
	}()

	less, more = <-lchan, <-mchan

	result := make([]int, 0, len(src))
	result = append(less, base)
	result = append(result, more...)
	return result
}

func quick[T int](src []T) []T {
	if len(src) < 2 {
		return src
	}
	base := src[0]
	var less []T
	var more []T
	for i := 1; i < len(src); i++ {
		v := src[i]
		if v < base {
			less = append(less, v)
		} else {
			more = append(more, v)
		}
	}
	less = quick(less)
	result := append(less, base)
	result = append(result, quick(more)...)
	return result
}
