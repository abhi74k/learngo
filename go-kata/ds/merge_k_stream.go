package ds

import "container/heap"

type Stream interface {
	Next() (int, bool)
}

type HeapNode struct {
	val    int
	stream Stream
}

type StreamHeap []HeapNode

type StreamMerger struct {
	nodes StreamHeap
}

func (h StreamHeap) Len() int {
	return len(h)
}

func (h StreamHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h StreamHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StreamHeap) Push(x any) {
	*h = append(*h, x.(HeapNode))
}

func (h *StreamHeap) Pop() any {

	n := h.Len()
	old := *h
	to_remove := old[n-1]
	*h = old[0 : n-1]

	return to_remove
}

func CreateMerger(streams []Stream) *StreamMerger {
	merger := StreamMerger{}

	heap.Init(&merger.nodes)

	for _, stream := range streams {

		if val, ok := stream.Next(); ok {
			heap.Push(&merger.nodes, HeapNode{val: val, stream: stream})
		}
	}

	return &merger
}

func (m *StreamMerger) Next() (int, bool) {

	if len(m.nodes) == 0 {
		return 0, false
	}

	node := heap.Pop(&m.nodes).(HeapNode)
	v := node.val
	s := node.stream

	if new_value, ok := s.Next(); ok {
		heap.Push(&m.nodes, HeapNode{val: new_value, stream: s})
	}

	return v, true
}
