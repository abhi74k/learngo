package ds

import (
	"errors"
)

type Ring struct {
	buf      []int
	capacity int
	head     uint64 // data can be read from
	tail     uint64 // data can be written to
}

func (r *Ring) Len() int    { return int(r.tail - r.head) }
func (r *Ring) Full() bool  { return r.Len() == r.capacity }
func (r *Ring) Empty() bool { return r.head == r.tail }

func (r *Ring) physicalidx(index uint64) int { return int(index % uint64(r.capacity)) }
func (r *Ring) get(index uint64) int         { return r.buf[r.physicalidx(index)] }
func (r *Ring) set(index uint64, v int)      { r.buf[r.physicalidx(index)] = v }

func NewRing(capacity int) *Ring {
	return &Ring{
		buf:      make([]int, capacity),
		capacity: capacity,
		head:     0,
		tail:     0,
	}
}

func (r *Ring) PushBack(v int) error {
	if r.Full() {
		return errors.New("Ring full!!!")
	}

	r.set(r.tail, v)
	r.tail++

	return nil
}

func (r *Ring) PopFront() (int, error) {
	if r.Empty() {
		return 0, errors.New("Ring Empty!!!")
	}

	v := r.get(r.head)
	r.head++

	return v, nil
}

func (r *Ring) PeekFront() (int, error) {
	if r.Empty() {
		return 0, errors.New("Ring Empty!!!")
	}

	v := r.get(r.head)
	return v, nil
}

func (r *Ring) PeekBack() (int, error) {
	if r.Empty() {
		return 0, errors.New("Ring Empty!!!")
	}

	v := r.get(r.tail - 1)
	return v, nil
}
