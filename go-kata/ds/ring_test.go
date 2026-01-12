package ds

import (
	"testing"
)

func TestNewRing(t *testing.T) {
	capacity := 5
	r := NewRing(capacity)

	if r == nil {
		t.Fatal("NewRing returned nil")
	}

	if r.capacity != capacity {
		t.Errorf("Expected capacity %d, got %d", capacity, r.capacity)
	}

	if r.head != 0 {
		t.Errorf("Expected head to be 0, got %d", r.head)
	}

	if r.tail != 0 {
		t.Errorf("Expected tail to be 0, got %d", r.tail)
	}

	if !r.Empty() {
		t.Error("New ring should be empty")
	}

	if r.Full() {
		t.Error("New ring should not be full")
	}

	if r.Len() != 0 {
		t.Errorf("Expected length 0, got %d", r.Len())
	}
}

func TestPushBackAndLen(t *testing.T) {
	r := NewRing(3)

	// Test pushing elements
	err := r.PushBack(10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if r.Len() != 1 {
		t.Errorf("Expected length 1, got %d", r.Len())
	}

	if r.Empty() {
		t.Error("Ring should not be empty after push")
	}

	err = r.PushBack(20)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if r.Len() != 2 {
		t.Errorf("Expected length 2, got %d", r.Len())
	}

	err = r.PushBack(30)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if r.Len() != 3 {
		t.Errorf("Expected length 3, got %d", r.Len())
	}

	if !r.Full() {
		t.Error("Ring should be full")
	}
}

func TestPushBackWhenFull(t *testing.T) {
	r := NewRing(2)

	// Fill the ring
	r.PushBack(1)
	r.PushBack(2)

	// Try to push when full
	err := r.PushBack(3)
	if err == nil {
		t.Error("Expected error when pushing to full ring")
	}

	expectedError := "Ring full!!!"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}
}

func TestPopFront(t *testing.T) {
	r := NewRing(3)

	// Push some elements
	r.PushBack(10)
	r.PushBack(20)
	r.PushBack(30)

	// Pop elements in FIFO order
	val, err := r.PopFront()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}
	if r.Len() != 2 {
		t.Errorf("Expected length 2, got %d", r.Len())
	}

	val, err = r.PopFront()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}
	if r.Len() != 1 {
		t.Errorf("Expected length 1, got %d", r.Len())
	}

	val, err = r.PopFront()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}
	if r.Len() != 0 {
		t.Errorf("Expected length 0, got %d", r.Len())
	}

	if !r.Empty() {
		t.Error("Ring should be empty after popping all elements")
	}
}

func TestPopFrontWhenEmpty(t *testing.T) {
	r := NewRing(3)

	// Try to pop from empty ring
	val, err := r.PopFront()
	if err == nil {
		t.Error("Expected error when popping from empty ring")
	}

	expectedError := "Ring Empty!!!"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}

	if val != 0 {
		t.Errorf("Expected default value 0, got %d", val)
	}
}

func TestPeekFront(t *testing.T) {
	r := NewRing(3)

	// Push some elements
	r.PushBack(100)
	r.PushBack(200)

	// Peek should return first element without removing it
	val, err := r.PeekFront()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 100 {
		t.Errorf("Expected 100, got %d", val)
	}

	// Length should remain the same
	if r.Len() != 2 {
		t.Errorf("Expected length 2, got %d", r.Len())
	}

	// Peek again should return same value
	val, err = r.PeekFront()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 100 {
		t.Errorf("Expected 100, got %d", val)
	}
}

func TestPeekFrontWhenEmpty(t *testing.T) {
	r := NewRing(3)

	// Try to peek from empty ring
	val, err := r.PeekFront()
	if err == nil {
		t.Error("Expected error when peeking empty ring")
	}

	expectedError := "Ring Empty!!!"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}

	if val != 0 {
		t.Errorf("Expected default value 0, got %d", val)
	}
}

func TestPeekBack(t *testing.T) {
	r := NewRing(3)

	// Push some elements
	r.PushBack(100)
	r.PushBack(200)
	r.PushBack(300)

	// Peek should return last element without removing it
	val, err := r.PeekBack()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 300 {
		t.Errorf("Expected 300, got %d", val)
	}

	// Length should remain the same
	if r.Len() != 3 {
		t.Errorf("Expected length 3, got %d", r.Len())
	}

	// Peek again should return same value
	val, err = r.PeekBack()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 300 {
		t.Errorf("Expected 300, got %d", val)
	}
}

func TestPeekBackWhenEmpty(t *testing.T) {
	r := NewRing(3)

	// Try to peek from empty ring
	val, err := r.PeekBack()
	if err == nil {
		t.Error("Expected error when peeking empty ring")
	}

	expectedError := "Ring Empty!!!"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}

	if val != 0 {
		t.Errorf("Expected default value 0, got %d", val)
	}
}

func TestRingStates(t *testing.T) {
	r := NewRing(2)

	// Test empty state
	if !r.Empty() {
		t.Error("New ring should be empty")
	}
	if r.Full() {
		t.Error("New ring should not be full")
	}
	if r.Len() != 0 {
		t.Errorf("Empty ring length should be 0, got %d", r.Len())
	}

	// Test partial state
	r.PushBack(1)
	if r.Empty() {
		t.Error("Ring with one element should not be empty")
	}
	if r.Full() {
		t.Error("Ring with one element should not be full")
	}
	if r.Len() != 1 {
		t.Errorf("Ring with one element should have length 1, got %d", r.Len())
	}

	// Test full state
	r.PushBack(2)
	if r.Empty() {
		t.Error("Full ring should not be empty")
	}
	if !r.Full() {
		t.Error("Ring should be full")
	}
	if r.Len() != 2 {
		t.Errorf("Full ring should have length 2, got %d", r.Len())
	}
}

func TestCircularBehavior(t *testing.T) {
	r := NewRing(3)

	// Fill the ring
	r.PushBack(1)
	r.PushBack(2)
	r.PushBack(3)

	// Pop one element
	val, _ := r.PopFront()
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	// Push another element (should wrap around)
	err := r.PushBack(4)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify order is maintained
	val, _ = r.PopFront()
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	val, _ = r.PopFront()
	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}

	val, _ = r.PopFront()
	if val != 4 {
		t.Errorf("Expected 4, got %d", val)
	}
}

func TestMultipleWraparounds(t *testing.T) {
	r := NewRing(2)

	// Test multiple cycles of fill/empty
	for cycle := 0; cycle < 3; cycle++ {
		// Fill
		r.PushBack(cycle*10 + 1)
		r.PushBack(cycle*10 + 2)

		if !r.Full() {
			t.Errorf("Ring should be full in cycle %d", cycle)
		}

		// Empty
		val1, _ := r.PopFront()
		val2, _ := r.PopFront()

		if val1 != cycle*10+1 || val2 != cycle*10+2 {
			t.Errorf("Cycle %d: expected %d, %d got %d, %d",
				cycle, cycle*10+1, cycle*10+2, val1, val2)
		}

		if !r.Empty() {
			t.Errorf("Ring should be empty after cycle %d", cycle)
		}
	}
}

func TestMixedOperations(t *testing.T) {
	r := NewRing(4)

	// Push some elements
	r.PushBack(10)
	r.PushBack(20)

	// Peek front and back
	front, _ := r.PeekFront()
	back, _ := r.PeekBack()
	if front != 10 || back != 20 {
		t.Errorf("Expected front=10, back=20, got front=%d, back=%d", front, back)
	}

	// Pop one
	val, _ := r.PopFront()
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	// Push more
	r.PushBack(30)
	r.PushBack(40)

	// Peek again
	front, _ = r.PeekFront()
	back, _ = r.PeekBack()
	if front != 20 || back != 40 {
		t.Errorf("Expected front=20, back=40, got front=%d, back=%d", front, back)
	}

	// Verify length
	if r.Len() != 3 {
		t.Errorf("Expected length 3, got %d", r.Len())
	}
}

func TestSingleElementRing(t *testing.T) {
	r := NewRing(1)

	// Test with single capacity
	err := r.PushBack(42)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !r.Full() {
		t.Error("Single element ring should be full")
	}

	// Both peek operations should return the same value
	front, _ := r.PeekFront()
	back, _ := r.PeekBack()
	if front != 42 || back != 42 {
		t.Errorf("Expected both peeks to return 42, got front=%d, back=%d", front, back)
	}

	// Pop should return the element
	val, _ := r.PopFront()
	if val != 42 {
		t.Errorf("Expected 42, got %d", val)
	}

	if !r.Empty() {
		t.Error("Ring should be empty after popping single element")
	}
}
