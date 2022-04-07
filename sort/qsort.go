package sort

func quickParallel(src []int) {
	if len(src) < 2 {
		return
	}
	base := src[0]
	less := make([]int, 0, len(src))
	more := make([]int, 0, len(src))
	for i := 1; i < len(src); i++ {
		v := src[i]
		if v < base {
			less = append(less, v)
		} else {
			more = append(more, v)
		}
	}

	lchan := make(chan []int)
	mchan := make(chan []int)
	go func() {
		less := less
		quickParallel(less)
		lchan <- less
	}()
	go func() {
		more := more
		quickParallel(more)
		mchan <- more
	}()

	less, more = <-lchan, <-mchan

	copy(src, less)
	src[len(less)] = base
	copy(src[len(less)+1:], more)
}

func quickParallelMagic(src []int) {
	if len(src) < 2 {
		return
	}
	base := src[0]
	less := make([]int, 0, len(src))
	more := make([]int, 0, len(src))
	for i := 1; i < len(src); i++ {
		v := src[i]
		if v < base {
			less = append(less, v)
		} else {
			more = append(more, v)
		}
	}

	lchan := make(chan []int)
	mchan := make(chan []int)
	go func() {
		quickParallel(less)
		lchan <- less
	}()
	go func() {
		quickParallel(more)
		mchan <- more
	}()

	_, _ = <-lchan, <-mchan

	copy(src, less)
	src[len(less)] = base
	copy(src[len(less)+1:], more)
}

func quick(src []int) {
	if len(src) < 2 {
		return
	}
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

	quick(less)
	quick(more)

	copy(src, less)
	src[len(less)] = base
	copy(src[len(less)+1:], more)
}
