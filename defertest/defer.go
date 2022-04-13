package defertest

import (
	"sync"
)

type Queue struct {
	sync.Mutex
	arr []int
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Put(elem int) {
	q.Lock()
	q.arr = append(q.arr, elem)
	q.Unlock()
}

func (q *Queue) PutDefer(elem int) {
	q.Lock()
	defer q.Unlock()
	q.arr = append(q.arr, elem)
}

func (q *Queue) Get() int {
	q.Lock()
	if len(q.arr) == 0 {
		q.Unlock()
		return 0
	}
	res := q.arr[0]
	q.arr = q.arr[1:]
	q.Unlock()
	return res
}

func (q *Queue) GetDefer() int {
	q.Lock()
	defer q.Unlock()
	if len(q.arr) == 0 {
		return 0
	}
	res := q.arr[0]
	q.arr = q.arr[1:]
	return res
}
