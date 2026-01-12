package ds

import (
	"container/heap"
	"time"
)

type Task struct {
	ID       string
	deadline time.Time

	index int
}

type TaskHeap []*Task

type Scheduler struct {
	h    TaskHeap
	byID map[string]*Task
}

func (h TaskHeap) Len() int {
	return len(h)
}

func (h TaskHeap) Less(i, j int) bool {
	return h[i].deadline.Before(h[j].deadline)
}

func (h TaskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]

	h[i].index = i
	h[j].index = j
}

func (h *TaskHeap) Push(x any) {
	*h = append(*h, x.(*Task))
}

func (h *TaskHeap) Pop() any {
	old := *h

	n := len(old)
	to_remove := old[n-1]
	*h = old[0 : n-1]
	return to_remove
}

func CreateScheduler() *Scheduler {
	return &Scheduler{
		h:    make(TaskHeap, 0),
		byID: make(map[string]*Task, 0),
	}
}

func (s *Scheduler) AddOrUpdate(ID string, deadline time.Time) {

	if t, ok := s.byID[ID]; ok {

		t.ID = ID
		t.deadline = deadline

		heap.Fix(&s.h, t.index)

		return
	}

	index_to_insert := len(s.h)
	task := &Task{
		ID:       ID,
		deadline: deadline,
		index:    index_to_insert,
	}
	s.byID[ID] = task

	heap.Push(&s.h, task)

}

func (s *Scheduler) Remove(ID string) bool {

	t, ok := s.byID[ID]
	if !ok {
		return false
	}

	heap.Remove(&s.h, t.index)
	delete(s.byID, ID)

	return true
}

func (s *Scheduler) PopDue(deadline time.Time) []*Task {

	var due []*Task = make([]*Task, 0)

	if len(s.h) > 0 {
		t := s.h[0]

		if t.deadline.After(deadline) {
			return due
		}

		dueTask := heap.Pop(&s.h).(*Task)
		delete(s.byID, dueTask.ID)

		due = append(due, dueTask)
	}

	return due
}
