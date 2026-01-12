package ds

import (
	"errors"
)

type Deque struct {
	buf      []int
	capacity int
	front    uint64 // Pointer to read from
	back     uint64 // Pointer to write to
}

func NewDeque(capacity int) *Deque {
	return &Deque{
		buf:      make([]int, capacity),
		capacity: capacity,
	}
}

func (q *Deque) Empty() bool { return q.front == q.back }
func (q *Deque) Len() int    { return int(q.back - q.front) }
func (q *Deque) Full() bool  { return q.Len() == q.capacity }

func (q *Deque) idx(index uint64) int    { return int(index % uint64(q.capacity)) }
func (q *Deque) set(index uint64, v int) { q.buf[q.idx(index)] = v }
func (q *Deque) get(index uint64) int    { return q.buf[q.idx(index)] }

func (q *Deque) PushFront(v int) error {
	if q.Full() {
		return errors.New("Dequeue is full!!")
	}

	q.front--
	q.set(q.front, v)

	return nil
}

func (q *Deque) PushBack(v int) error {
	if q.Full() {
		return errors.New("Deque is full!!")
	}

	q.set(q.back, v)
	q.back++

	return nil
}

func (q *Deque) PopFront() (int, error) {
	if q.Empty() {
		return 0, errors.New("Deque is empty!")
	}

	v := q.get(q.front)
	q.front++

	return v, nil
}

func (q *Deque) PopBack() (int, error) {
	if q.Empty() {
		return 0, errors.New("Deque is empty!")
	}

	q.back--
	v := q.get(q.back)

	return v, nil
}

func (q *Deque) PeekFront() (int, error) {
	if q.Empty() {
		return 0, errors.New("Deque is empty!!")
	}

	v := q.get(q.front)
	return v, nil
}

func (q *Deque) PeekBack() (int, error) {
	if q.Empty() {
		return 0, errors.New("Deque is empty!!")
	}

	v := q.get(q.back - 1)
	return v, nil
}
