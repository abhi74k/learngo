package ds

import (
	"container/heap"
	"errors"
)

type IntHeap struct {
	data []int
	less func(a, b int) bool
}

func NewMinHeap() *IntHeap {
	h := &IntHeap{
		data: make([]int, 0),
		less: func(a, b int) bool { return a < b },
	}

	heap.Init(h)
	return h
}

func NewMaxHeap() *IntHeap {
	h := &IntHeap{
		data: make([]int, 0),
		less: func(a, b int) bool { return a > b },
	}

	heap.Init(h)
	return h
}

// Sort interface

func (h IntHeap) Len() int {
	return len(h.data)
}

func (h IntHeap) Less(i, j int) bool {
	return h.less(h.data[i], h.data[j])
}

func (h IntHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Heap interface
func (h *IntHeap) Push(x any) {
	h.data = append(h.data, x.(int))
}

func (h *IntHeap) Pop() any {
	n := len(h.data)
	old := h.data[n-1]
	h.data = h.data[0 : n-1]
	return old
}

func (h *IntHeap) PushInt(x int) {
	heap.Push(h, x)
}

func (h *IntHeap) PopInt() (int, error) {
	if len(h.data) == 0 {
		return 0, errors.New("Heap is empty!!")
	}

	return heap.Pop(h).(int), nil
}
