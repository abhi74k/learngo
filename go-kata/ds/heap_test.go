package ds

import (
	"testing"
)

func TestNewMinHeap(t *testing.T) {
	h := NewMinHeap()

	if h == nil {
		t.Fatal("NewMinHeap() returned nil")
	}

	if h.Len() != 0 {
		t.Errorf("Expected empty heap length 0, got %d", h.Len())
	}

	// Test that it's actually a min heap by checking the less function
	if !h.less(1, 2) {
		t.Error("Min heap should have 1 < 2")
	}

	if h.less(2, 1) {
		t.Error("Min heap should not have 2 < 1")
	}
}

func TestNewMaxHeap(t *testing.T) {
	h := NewMaxHeap()

	if h == nil {
		t.Fatal("NewMaxHeap() returned nil")
	}

	if h.Len() != 0 {
		t.Errorf("Expected empty heap length 0, got %d", h.Len())
	}

	// Test that it's actually a max heap by checking the less function
	if h.less(1, 2) {
		t.Error("Max heap should not have 1 < 2")
	}

	if !h.less(2, 1) {
		t.Error("Max heap should have 2 < 1")
	}
}

func TestMinHeapPushPop(t *testing.T) {
	h := NewMinHeap()

	// Test pushing elements
	elements := []int{5, 2, 8, 1, 9, 3}
	for _, elem := range elements {
		h.PushInt(elem)
	}

	if h.Len() != len(elements) {
		t.Errorf("Expected heap length %d, got %d", len(elements), h.Len())
	}

	// Test popping elements - should come out in ascending order for min heap
	expected := []int{1, 2, 3, 5, 8, 9}
	for i, expectedVal := range expected {
		val, err := h.PopInt()
		if err != nil {
			t.Fatalf("Unexpected error popping element %d: %v", i, err)
		}

		if val != expectedVal {
			t.Errorf("Expected %d at position %d, got %d", expectedVal, i, val)
		}
	}

	if h.Len() != 0 {
		t.Errorf("Expected empty heap after popping all elements, length is %d", h.Len())
	}
}

func TestMaxHeapPushPop(t *testing.T) {
	h := NewMaxHeap()

	// Test pushing elements
	elements := []int{5, 2, 8, 1, 9, 3}
	for _, elem := range elements {
		h.PushInt(elem)
	}

	if h.Len() != len(elements) {
		t.Errorf("Expected heap length %d, got %d", len(elements), h.Len())
	}

	// Test popping elements - should come out in descending order for max heap
	expected := []int{9, 8, 5, 3, 2, 1}
	for i, expectedVal := range expected {
		val, err := h.PopInt()
		if err != nil {
			t.Fatalf("Unexpected error popping element %d: %v", i, err)
		}

		if val != expectedVal {
			t.Errorf("Expected %d at position %d, got %d", expectedVal, i, val)
		}
	}

	if h.Len() != 0 {
		t.Errorf("Expected empty heap after popping all elements, length is %d", h.Len())
	}
}

func TestPopFromEmptyHeap(t *testing.T) {
	h := NewMinHeap()

	val, err := h.PopInt()
	if err == nil {
		t.Error("Expected error when popping from empty heap")
	}

	if val != 0 {
		t.Errorf("Expected 0 value when popping from empty heap, got %d", val)
	}

	expectedErrorMsg := "Heap is empty!!"
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMsg, err.Error())
	}
}

func TestSingleElementHeap(t *testing.T) {
	// Test min heap with single element
	minHeap := NewMinHeap()
	minHeap.PushInt(42)

	if minHeap.Len() != 1 {
		t.Errorf("Expected heap length 1, got %d", minHeap.Len())
	}

	val, err := minHeap.PopInt()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if val != 42 {
		t.Errorf("Expected 42, got %d", val)
	}

	if minHeap.Len() != 0 {
		t.Errorf("Expected empty heap after pop, length is %d", minHeap.Len())
	}

	// Test max heap with single element
	maxHeap := NewMaxHeap()
	maxHeap.PushInt(42)

	if maxHeap.Len() != 1 {
		t.Errorf("Expected heap length 1, got %d", maxHeap.Len())
	}

	val, err = maxHeap.PopInt()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if val != 42 {
		t.Errorf("Expected 42, got %d", val)
	}

	if maxHeap.Len() != 0 {
		t.Errorf("Expected empty heap after pop, length is %d", maxHeap.Len())
	}
}

func TestHeapWithDuplicates(t *testing.T) {
	h := NewMinHeap()

	// Push duplicate values
	elements := []int{5, 3, 5, 1, 3, 1, 5}
	for _, elem := range elements {
		h.PushInt(elem)
	}

	// Should pop in sorted order
	expected := []int{1, 1, 3, 3, 5, 5, 5}
	for i, expectedVal := range expected {
		val, err := h.PopInt()
		if err != nil {
			t.Fatalf("Unexpected error popping element %d: %v", i, err)
		}

		if val != expectedVal {
			t.Errorf("Expected %d at position %d, got %d", expectedVal, i, val)
		}
	}
}

func TestHeapWithNegativeNumbers(t *testing.T) {
	h := NewMinHeap()

	elements := []int{-5, 10, -2, 0, 3, -10}
	for _, elem := range elements {
		h.PushInt(elem)
	}

	expected := []int{-10, -5, -2, 0, 3, 10}
	for i, expectedVal := range expected {
		val, err := h.PopInt()
		if err != nil {
			t.Fatalf("Unexpected error popping element %d: %v", i, err)
		}

		if val != expectedVal {
			t.Errorf("Expected %d at position %d, got %d", expectedVal, i, val)
		}
	}
}

func TestHeapMixedOperations(t *testing.T) {
	h := NewMinHeap()

	// Push some elements
	h.PushInt(5)
	h.PushInt(2)
	h.PushInt(8)

	// Pop minimum
	val, err := h.PopInt()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	// Push more elements
	h.PushInt(1)
	h.PushInt(10)

	// Pop should give us 1 (new minimum)
	val, err = h.PopInt()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	// Remaining elements should be 5, 8, 10
	expected := []int{5, 8, 10}
	for i, expectedVal := range expected {
		val, err = h.PopInt()
		if err != nil {
			t.Fatalf("Unexpected error popping element %d: %v", i, err)
		}

		if val != expectedVal {
			t.Errorf("Expected %d at position %d, got %d", expectedVal, i, val)
		}
	}
}

func TestHeapSortInterface(t *testing.T) {
	h := NewMinHeap()

	// Test Len, Less, and Swap methods
	elements := []int{3, 1, 4}
	for _, elem := range elements {
		h.PushInt(elem)
	}

	// Test Len
	if h.Len() != 3 {
		t.Errorf("Expected length 3, got %d", h.Len())
	}

	// The heap should maintain heap property, so we can't directly test Less and Swap
	// without knowing the internal structure, but we can verify the heap works correctly
	// by ensuring elements come out in sorted order
	expected := []int{1, 3, 4}
	for i, expectedVal := range expected {
		val, err := h.PopInt()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if val != expectedVal {
			t.Errorf("Expected %d at position %d, got %d", expectedVal, i, val)
		}
	}
}

func TestLargeHeap(t *testing.T) {
	h := NewMinHeap()

	// Push 1000 elements in reverse order
	for i := 1000; i > 0; i-- {
		h.PushInt(i)
	}

	if h.Len() != 1000 {
		t.Errorf("Expected heap length 1000, got %d", h.Len())
	}

	// Pop all elements - should come out in ascending order
	for i := 1; i <= 1000; i++ {
		val, err := h.PopInt()
		if err != nil {
			t.Fatalf("Unexpected error popping element %d: %v", i, err)
		}

		if val != i {
			t.Errorf("Expected %d, got %d", i, val)
		}
	}

	if h.Len() != 0 {
		t.Errorf("Expected empty heap after popping all elements, length is %d", h.Len())
	}
}
