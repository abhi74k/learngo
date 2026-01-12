package ds

import (
	"testing"
)

func TestNewDeque(t *testing.T) {
	capacity := 5
	d := NewDeque(capacity)

	if d == nil {
		t.Fatal("NewDeque returned nil")
	}

	if d.capacity != capacity {
		t.Errorf("Expected capacity %d, got %d", capacity, d.capacity)
	}

	if d.front != 0 {
		t.Errorf("Expected front to be 0, got %d", d.front)
	}

	if d.back != 0 {
		t.Errorf("Expected back to be 0, got %d", d.back)
	}

	if !d.Empty() {
		t.Error("New deque should be empty")
	}

	if d.Full() {
		t.Error("New deque should not be full")
	}

	if d.Len() != 0 {
		t.Errorf("Expected length 0, got %d", d.Len())
	}
}

func TestPushBack(t *testing.T) {
	d := NewDeque(3)

	// Test pushing elements to back
	err := d.PushBack(10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if d.Len() != 1 {
		t.Errorf("Expected length 1, got %d", d.Len())
	}

	if d.Empty() {
		t.Error("Deque should not be empty after push")
	}

	err = d.PushBack(20)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if d.Len() != 2 {
		t.Errorf("Expected length 2, got %d", d.Len())
	}

	err = d.PushBack(30)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if d.Len() != 3 {
		t.Errorf("Expected length 3, got %d", d.Len())
	}

	if !d.Full() {
		t.Error("Deque should be full")
	}
}

func TestPushFront(t *testing.T) {
	d := NewDeque(3)

	// Test pushing elements to front
	err := d.PushFront(10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if d.Len() != 1 {
		t.Errorf("Expected length 1, got %d", d.Len())
	}

	err = d.PushFront(20)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if d.Len() != 2 {
		t.Errorf("Expected length 2, got %d", d.Len())
	}

	err = d.PushFront(30)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if d.Len() != 3 {
		t.Errorf("Expected length 3, got %d", d.Len())
	}

	if !d.Full() {
		t.Error("Deque should be full")
	}
}

func TestPushWhenFull(t *testing.T) {
	d := NewDeque(2)

	// Fill the deque
	d.PushBack(1)
	d.PushBack(2)

	// Try to push back when full
	err := d.PushBack(3)
	if err == nil {
		t.Error("Expected error when pushing back to full deque")
	}

	expectedError := "Deque is full!!"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}

	// Try to push front when full
	err = d.PushFront(3)
	if err == nil {
		t.Error("Expected error when pushing front to full deque")
	}

	expectedErrorFront := "Dequeue is full!!"
	if err.Error() != expectedErrorFront {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorFront, err.Error())
	}
}

func TestPopBack(t *testing.T) {
	d := NewDeque(3)

	// Push some elements
	d.PushBack(10)
	d.PushBack(20)
	d.PushBack(30)

	// Pop elements from back in LIFO order
	val, err := d.PopBack()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}
	if d.Len() != 2 {
		t.Errorf("Expected length 2, got %d", d.Len())
	}

	val, err = d.PopBack()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}
	if d.Len() != 1 {
		t.Errorf("Expected length 1, got %d", d.Len())
	}

	val, err = d.PopBack()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}
	if d.Len() != 0 {
		t.Errorf("Expected length 0, got %d", d.Len())
	}

	if !d.Empty() {
		t.Error("Deque should be empty after popping all elements")
	}
}

func TestPopFrontFIFO(t *testing.T) {
	d := NewDeque(3)

	// Push some elements
	d.PushBack(10)
	d.PushBack(20)
	d.PushBack(30)

	// Pop elements from front in FIFO order
	val, err := d.PopFront()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}
	if d.Len() != 2 {
		t.Errorf("Expected length 2, got %d", d.Len())
	}

	val, err = d.PopFront()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}
	if d.Len() != 1 {
		t.Errorf("Expected length 1, got %d", d.Len())
	}

	val, err = d.PopFront()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}
	if d.Len() != 0 {
		t.Errorf("Expected length 0, got %d", d.Len())
	}

	if !d.Empty() {
		t.Error("Deque should be empty after popping all elements")
	}
}

func TestPopWhenEmpty(t *testing.T) {
	d := NewDeque(3)

	// Try to pop front from empty deque
	val, err := d.PopFront()
	if err == nil {
		t.Error("Expected error when popping front from empty deque")
	}

	expectedError := "Deque is empty!"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}

	if val != 0 {
		t.Errorf("Expected default value 0, got %d", val)
	}

	// Try to pop back from empty deque
	val, err = d.PopBack()
	if err == nil {
		t.Error("Expected error when popping back from empty deque")
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}

	if val != 0 {
		t.Errorf("Expected default value 0, got %d", val)
	}
}

func TestPeekFrontBasic(t *testing.T) {
	d := NewDeque(3)

	// Push some elements
	d.PushBack(100)
	d.PushBack(200)

	// Peek should return first element without removing it
	val, err := d.PeekFront()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 100 {
		t.Errorf("Expected 100, got %d", val)
	}

	// Length should remain the same
	if d.Len() != 2 {
		t.Errorf("Expected length 2, got %d", d.Len())
	}

	// Peek again should return same value
	val, err = d.PeekFront()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 100 {
		t.Errorf("Expected 100, got %d", val)
	}
}

func TestPeekBackBasic(t *testing.T) {
	d := NewDeque(3)

	// Push some elements
	d.PushBack(100)
	d.PushBack(200)
	d.PushBack(300)

	// Peek should return last element without removing it
	val, err := d.PeekBack()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 300 {
		t.Errorf("Expected 300, got %d", val)
	}

	// Length should remain the same
	if d.Len() != 3 {
		t.Errorf("Expected length 3, got %d", d.Len())
	}

	// Peek again should return same value
	val, err = d.PeekBack()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 300 {
		t.Errorf("Expected 300, got %d", val)
	}
}

func TestPeekWhenEmpty(t *testing.T) {
	d := NewDeque(3)

	// Try to peek front from empty deque
	val, err := d.PeekFront()
	if err == nil {
		t.Error("Expected error when peeking front of empty deque")
	}

	expectedError := "Deque is empty!!"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}

	if val != 0 {
		t.Errorf("Expected default value 0, got %d", val)
	}

	// Try to peek back from empty deque
	val, err = d.PeekBack()
	if err == nil {
		t.Error("Expected error when peeking back of empty deque")
	}

	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}

	if val != 0 {
		t.Errorf("Expected default value 0, got %d", val)
	}
}

func TestDequeStates(t *testing.T) {
	d := NewDeque(2)

	// Test empty state
	if !d.Empty() {
		t.Error("New deque should be empty")
	}
	if d.Full() {
		t.Error("New deque should not be full")
	}
	if d.Len() != 0 {
		t.Errorf("Empty deque length should be 0, got %d", d.Len())
	}

	// Test partial state
	d.PushBack(1)
	if d.Empty() {
		t.Error("Deque with one element should not be empty")
	}
	if d.Full() {
		t.Error("Deque with one element should not be full")
	}
	if d.Len() != 1 {
		t.Errorf("Deque with one element should have length 1, got %d", d.Len())
	}

	// Test full state
	d.PushBack(2)
	if d.Empty() {
		t.Error("Full deque should not be empty")
	}
	if !d.Full() {
		t.Error("Deque should be full")
	}
	if d.Len() != 2 {
		t.Errorf("Full deque should have length 2, got %d", d.Len())
	}
}

func TestMixedFrontBackOperations(t *testing.T) {
	d := NewDeque(4)

	// Push to back, then front
	d.PushBack(10)
	d.PushFront(5)
	d.PushBack(20)
	d.PushFront(1)

	// Should be: [1, 5, 10, 20]
	if d.Len() != 4 {
		t.Errorf("Expected length 4, got %d", d.Len())
	}

	// Peek both ends
	front, _ := d.PeekFront()
	back, _ := d.PeekBack()
	if front != 1 || back != 20 {
		t.Errorf("Expected front=1, back=20, got front=%d, back=%d", front, back)
	}

	// Pop from front: should get 1
	val, _ := d.PopFront()
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	// Pop from back: should get 20
	val, _ = d.PopBack()
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}

	// Remaining should be [5, 10]
	if d.Len() != 2 {
		t.Errorf("Expected length 2, got %d", d.Len())
	}

	front, _ = d.PeekFront()
	back, _ = d.PeekBack()
	if front != 5 || back != 10 {
		t.Errorf("Expected front=5, back=10, got front=%d, back=%d", front, back)
	}
}

func TestPushFrontOrder(t *testing.T) {
	d := NewDeque(3)

	// Push elements to front
	d.PushFront(10)
	d.PushFront(20)
	d.PushFront(30)

	// Should be: [30, 20, 10]
	// Pop from front should give: 30, 20, 10
	val, _ := d.PopFront()
	if val != 30 {
		t.Errorf("Expected 30, got %d", val)
	}

	val, _ = d.PopFront()
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}

	val, _ = d.PopFront()
	if val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}
}

func TestCircularBehaviorAdvanced(t *testing.T) {
	d := NewDeque(3)

	// Fill the deque from back
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	// Pop one from front
	val, _ := d.PopFront()
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	// Push to back (should wrap around)
	err := d.PushBack(4)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Pop one from front
	val, _ = d.PopFront()
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	// Push to front (should wrap around)
	err = d.PushFront(0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify order: [0, 3, 4]
	val, _ = d.PopFront()
	if val != 0 {
		t.Errorf("Expected 0, got %d", val)
	}

	val, _ = d.PopFront()
	if val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}

	val, _ = d.PopFront()
	if val != 4 {
		t.Errorf("Expected 4, got %d", val)
	}
}

func TestSingleElementDeque(t *testing.T) {
	d := NewDeque(1)

	// Test with single capacity
	err := d.PushBack(42)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !d.Full() {
		t.Error("Single element deque should be full")
	}

	// Both peek operations should return the same value
	front, _ := d.PeekFront()
	back, _ := d.PeekBack()
	if front != 42 || back != 42 {
		t.Errorf("Expected both peeks to return 42, got front=%d, back=%d", front, back)
	}

	// Pop from front should return the element
	val, _ := d.PopFront()
	if val != 42 {
		t.Errorf("Expected 42, got %d", val)
	}

	if !d.Empty() {
		t.Error("Deque should be empty after popping single element")
	}

	// Test push front with single capacity
	err = d.PushFront(99)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Pop from back should return the element
	val, _ = d.PopBack()
	if val != 99 {
		t.Errorf("Expected 99, got %d", val)
	}
}

func TestMultipleWraparoundsDeque(t *testing.T) {
	d := NewDeque(2)

	// Test multiple cycles of operations
	for cycle := 0; cycle < 3; cycle++ {
		// Fill from back
		d.PushBack(cycle*10 + 1)
		d.PushBack(cycle*10 + 2)

		if !d.Full() {
			t.Errorf("Deque should be full in cycle %d", cycle)
		}

		// Empty from front
		val1, _ := d.PopFront()
		val2, _ := d.PopFront()

		if val1 != cycle*10+1 || val2 != cycle*10+2 {
			t.Errorf("Cycle %d: expected %d, %d got %d, %d",
				cycle, cycle*10+1, cycle*10+2, val1, val2)
		}

		if !d.Empty() {
			t.Errorf("Deque should be empty after cycle %d", cycle)
		}
	}
}

func TestAlternatingOperations(t *testing.T) {
	d := NewDeque(4)

	// Alternate between front and back operations
	d.PushBack(1)
	d.PushFront(0)
	d.PushBack(2)
	d.PushFront(-1)

	// Should be: [-1, 0, 1, 2]
	if d.Len() != 4 {
		t.Errorf("Expected length 4, got %d", d.Len())
	}

	// Pop alternating
	val, _ := d.PopBack() // Should get 2
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	val, _ = d.PopFront() // Should get -1
	if val != -1 {
		t.Errorf("Expected -1, got %d", val)
	}

	val, _ = d.PopBack() // Should get 1
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	val, _ = d.PopFront() // Should get 0
	if val != 0 {
		t.Errorf("Expected 0, got %d", val)
	}

	if !d.Empty() {
		t.Error("Deque should be empty")
	}
}
