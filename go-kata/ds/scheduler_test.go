package ds

import (
	"testing"
	"time"
)

func TestCreateScheduler(t *testing.T) {
	s := CreateScheduler()

	if s == nil {
		t.Fatal("CreateScheduler() returned nil")
	}

	if s.h == nil {
		t.Error("Scheduler heap should be initialized")
	}

	if s.byID == nil {
		t.Error("Scheduler byID map should be initialized")
	}

	if len(s.h) != 0 {
		t.Errorf("Expected empty heap, got length %d", len(s.h))
	}

	if len(s.byID) != 0 {
		t.Errorf("Expected empty byID map, got length %d", len(s.byID))
	}
}

func TestAddOrUpdateNewTask(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()
	deadline := now.Add(time.Hour)

	s.AddOrUpdate("task1", deadline)

	if len(s.h) != 1 {
		t.Errorf("Expected heap length 1, got %d", len(s.h))
	}

	if len(s.byID) != 1 {
		t.Errorf("Expected byID map length 1, got %d", len(s.byID))
	}

	task, exists := s.byID["task1"]
	if !exists {
		t.Error("Task should exist in byID map")
	}

	if task.ID != "task1" {
		t.Errorf("Expected task ID 'task1', got '%s'", task.ID)
	}

	if !task.deadline.Equal(deadline) {
		t.Errorf("Expected deadline %v, got %v", deadline, task.deadline)
	}

	if task.index != 0 {
		t.Errorf("Expected task index 0, got %d", task.index)
	}
}

func TestAddOrUpdateExistingTask(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()
	originalDeadline := now.Add(time.Hour)
	newDeadline := now.Add(2 * time.Hour)

	// Add initial task
	s.AddOrUpdate("task1", originalDeadline)

	// Update existing task
	s.AddOrUpdate("task1", newDeadline)

	// Should still have only one task
	if len(s.h) != 1 {
		t.Errorf("Expected heap length 1, got %d", len(s.h))
	}

	if len(s.byID) != 1 {
		t.Errorf("Expected byID map length 1, got %d", len(s.byID))
	}

	task, exists := s.byID["task1"]
	if !exists {
		t.Error("Task should exist in byID map")
	}

	if !task.deadline.Equal(newDeadline) {
		t.Errorf("Expected updated deadline %v, got %v", newDeadline, task.deadline)
	}
}

func TestAddMultipleTasks(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()

	tasks := []struct {
		id       string
		deadline time.Time
	}{
		{"task1", now.Add(3 * time.Hour)},
		{"task2", now.Add(1 * time.Hour)}, // This should be at the top of heap
		{"task3", now.Add(2 * time.Hour)},
	}

	for _, task := range tasks {
		s.AddOrUpdate(task.id, task.deadline)
	}

	if len(s.h) != 3 {
		t.Errorf("Expected heap length 3, got %d", len(s.h))
	}

	if len(s.byID) != 3 {
		t.Errorf("Expected byID map length 3, got %d", len(s.byID))
	}

	// The heap should maintain min-heap property (earliest deadline at top)
	if s.h[0].ID != "task2" {
		t.Errorf("Expected task2 at heap top, got %s", s.h[0].ID)
	}
}

func TestRemoveExistingTask(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()

	s.AddOrUpdate("task1", now.Add(time.Hour))
	s.AddOrUpdate("task2", now.Add(2*time.Hour))

	removed := s.Remove("task1")

	if !removed {
		t.Error("Remove should return true for existing task")
	}

	if len(s.h) != 1 {
		t.Errorf("Expected heap length 1 after removal, got %d", len(s.h))
	}

	if len(s.byID) != 1 {
		t.Errorf("Expected byID map length 1 after removal, got %d", len(s.byID))
	}

	_, exists := s.byID["task1"]
	if exists {
		t.Error("Removed task should not exist in byID map")
	}

	// Remaining task should be task2
	if s.h[0].ID != "task2" {
		t.Errorf("Expected task2 to remain, got %s", s.h[0].ID)
	}
}

func TestRemoveNonExistentTask(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()

	s.AddOrUpdate("task1", now.Add(time.Hour))

	removed := s.Remove("nonexistent")

	if removed {
		t.Error("Remove should return false for non-existent task")
	}

	if len(s.h) != 1 {
		t.Errorf("Expected heap length 1, got %d", len(s.h))
	}

	if len(s.byID) != 1 {
		t.Errorf("Expected byID map length 1, got %d", len(s.byID))
	}
}

func TestRemoveFromEmptyScheduler(t *testing.T) {
	s := CreateScheduler()

	removed := s.Remove("task1")

	if removed {
		t.Error("Remove should return false for empty scheduler")
	}
}

func TestPopDueWithDueTask(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()
	pastDeadline := now.Add(-time.Hour) // Task is overdue

	s.AddOrUpdate("task1", pastDeadline)

	due := s.PopDue(now)

	if len(due) != 1 {
		t.Errorf("Expected 1 due task, got %d", len(due))
	}

	if due[0].ID != "task1" {
		t.Errorf("Expected task1 to be due, got %s", due[0].ID)
	}

	if !due[0].deadline.Equal(pastDeadline) {
		t.Errorf("Expected deadline %v, got %v", pastDeadline, due[0].deadline)
	}

	// Task should be removed from scheduler
	if len(s.h) != 0 {
		t.Errorf("Expected empty heap after popping due task, got length %d", len(s.h))
	}

	if len(s.byID) != 0 {
		t.Errorf("Expected empty byID map after popping due task, got length %d", len(s.byID))
	}
}

func TestPopDueWithNoDueTasks(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()
	futureDeadline := now.Add(time.Hour) // Task is not due yet

	s.AddOrUpdate("task1", futureDeadline)

	due := s.PopDue(now)

	if len(due) != 0 {
		t.Errorf("Expected 0 due tasks, got %d", len(due))
	}

	// Task should remain in scheduler
	if len(s.h) != 1 {
		t.Errorf("Expected heap length 1, got %d", len(s.h))
	}

	if len(s.byID) != 1 {
		t.Errorf("Expected byID map length 1, got %d", len(s.byID))
	}
}

func TestPopDueFromEmptyScheduler(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()

	due := s.PopDue(now)

	if len(due) != 0 {
		t.Errorf("Expected 0 due tasks from empty scheduler, got %d", len(due))
	}
}

func TestPopDueOnlyPopsEarliestTask(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()

	// Add multiple tasks, some due, some not
	s.AddOrUpdate("task1", now.Add(-2*time.Hour)) // Due (earliest)
	s.AddOrUpdate("task2", now.Add(-1*time.Hour)) // Due (later)
	s.AddOrUpdate("task3", now.Add(time.Hour))    // Not due

	due := s.PopDue(now)

	// PopDue should only return one task (the earliest due task)
	if len(due) != 1 {
		t.Errorf("Expected 1 due task, got %d", len(due))
	}

	if due[0].ID != "task1" {
		t.Errorf("Expected task1 (earliest) to be popped, got %s", due[0].ID)
	}

	// Scheduler should still have 2 tasks
	if len(s.h) != 2 {
		t.Errorf("Expected heap length 2 after popping one task, got %d", len(s.h))
	}

	if len(s.byID) != 2 {
		t.Errorf("Expected byID map length 2 after popping one task, got %d", len(s.byID))
	}
}

func TestComplexScenario(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()

	// Add multiple tasks
	s.AddOrUpdate("urgent", now.Add(-time.Hour)) // Overdue
	s.AddOrUpdate("medium", now.Add(time.Hour))  // Future
	s.AddOrUpdate("low", now.Add(2*time.Hour))   // Further future

	// Update an existing task
	s.AddOrUpdate("medium", now.Add(-30*time.Minute)) // Make it overdue too

	// Remove a task
	removed := s.Remove("low")
	if !removed {
		t.Error("Should have removed 'low' task")
	}

	// Now we should have 2 tasks, both overdue
	if len(s.h) != 2 {
		t.Errorf("Expected heap length 2, got %d", len(s.h))
	}

	// Pop due tasks - should get the most urgent first
	due1 := s.PopDue(now)
	if len(due1) != 1 {
		t.Errorf("Expected 1 due task, got %d", len(due1))
	}

	// The most urgent should be the one with earlier deadline
	expectedFirst := "urgent" // -1 hour is earlier than -30 minutes
	if due1[0].ID != expectedFirst {
		t.Errorf("Expected %s to be popped first, got %s", expectedFirst, due1[0].ID)
	}

	// Pop again - should get the second overdue task
	due2 := s.PopDue(now)
	if len(due2) != 1 {
		t.Errorf("Expected 1 due task, got %d", len(due2))
	}

	if due2[0].ID != "medium" {
		t.Errorf("Expected 'medium' to be popped second, got %s", due2[0].ID)
	}

	// Scheduler should now be empty
	if len(s.h) != 0 {
		t.Errorf("Expected empty heap, got length %d", len(s.h))
	}

	if len(s.byID) != 0 {
		t.Errorf("Expected empty byID map, got length %d", len(s.byID))
	}
}

func TestTaskHeapInterface(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()

	// Add tasks with different deadlines
	tasks := []struct {
		id       string
		deadline time.Time
	}{
		{"task3", now.Add(3 * time.Hour)},
		{"task1", now.Add(1 * time.Hour)},
		{"task2", now.Add(2 * time.Hour)},
	}

	for _, task := range tasks {
		s.AddOrUpdate(task.id, task.deadline)
	}

	// Test that heap maintains min-heap property
	heap := s.h

	// Test Len
	if heap.Len() != 3 {
		t.Errorf("Expected heap length 3, got %d", heap.Len())
	}

	// Test Less - earlier deadline should be "less"
	if !heap.Less(0, 1) && !heap.Less(0, 2) {
		t.Error("Root should have earliest deadline")
	}

	// The root should be task1 (earliest deadline)
	if heap[0].ID != "task1" {
		t.Errorf("Expected task1 at root, got %s", heap[0].ID)
	}
}

func TestConcurrentDeadlines(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()
	sameDeadline := now.Add(time.Hour)

	// Add multiple tasks with the same deadline
	s.AddOrUpdate("task1", sameDeadline)
	s.AddOrUpdate("task2", sameDeadline)
	s.AddOrUpdate("task3", sameDeadline)

	if len(s.h) != 3 {
		t.Errorf("Expected heap length 3, got %d", len(s.h))
	}

	// All tasks should be considered not due if deadline is in future
	due := s.PopDue(now)
	if len(due) != 0 {
		t.Errorf("Expected 0 due tasks, got %d", len(due))
	}

	// All tasks should be considered due if deadline is in past
	due = s.PopDue(sameDeadline.Add(time.Minute))
	if len(due) != 1 {
		t.Errorf("Expected 1 due task, got %d", len(due))
	}
}

func TestUpdateTaskToEarlierDeadline(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()

	// Add tasks
	s.AddOrUpdate("task1", now.Add(3*time.Hour))
	s.AddOrUpdate("task2", now.Add(2*time.Hour))
	s.AddOrUpdate("task3", now.Add(1*time.Hour))

	// task3 should be at the root
	if s.h[0].ID != "task3" {
		t.Errorf("Expected task3 at root, got %s", s.h[0].ID)
	}

	// Update task1 to have the earliest deadline
	s.AddOrUpdate("task1", now.Add(30*time.Minute))

	// Now task1 should be at the root
	if s.h[0].ID != "task1" {
		t.Errorf("Expected task1 at root after update, got %s", s.h[0].ID)
	}
}

func TestUpdateTaskToLaterDeadline(t *testing.T) {
	s := CreateScheduler()
	now := time.Now()

	// Add tasks
	s.AddOrUpdate("task1", now.Add(1*time.Hour)) // This will be at root
	s.AddOrUpdate("task2", now.Add(2*time.Hour))
	s.AddOrUpdate("task3", now.Add(3*time.Hour))

	// task1 should be at the root
	if s.h[0].ID != "task1" {
		t.Errorf("Expected task1 at root, got %s", s.h[0].ID)
	}

	// Update task1 to have the latest deadline
	s.AddOrUpdate("task1", now.Add(4*time.Hour))

	// Now task2 should be at the root
	if s.h[0].ID != "task2" {
		t.Errorf("Expected task2 at root after update, got %s", s.h[0].ID)
	}
}
