package ds

import (
	"testing"
)

// MockStream implements the Stream interface for testing
type MockStream struct {
	values []int
	index  int
}

func NewMockStream(values []int) *MockStream {
	return &MockStream{
		values: values,
		index:  0,
	}
}

func (m *MockStream) Next() (int, bool) {
	if m.index >= len(m.values) {
		return 0, false
	}
	val := m.values[m.index]
	m.index++
	return val, true
}

// EmptyStream implements Stream interface but returns no values
type EmptyStream struct{}

func (e *EmptyStream) Next() (int, bool) {
	return 0, false
}

func TestCreateMerger(t *testing.T) {
	t.Run("CreateMerger with multiple streams", func(t *testing.T) {
		streams := []Stream{
			NewMockStream([]int{1, 4, 7}),
			NewMockStream([]int{2, 5, 8}),
			NewMockStream([]int{3, 6, 9}),
		}

		merger := CreateMerger(streams)
		if merger == nil {
			t.Fatal("CreateMerger returned nil")
		}

		if len(merger.nodes) != 3 {
			t.Errorf("Expected 3 nodes in heap, got %d", len(merger.nodes))
		}
	})

	t.Run("CreateMerger with empty streams", func(t *testing.T) {
		streams := []Stream{
			&EmptyStream{},
			&EmptyStream{},
		}

		merger := CreateMerger(streams)
		if merger == nil {
			t.Fatal("CreateMerger returned nil")
		}

		if len(merger.nodes) != 0 {
			t.Errorf("Expected 0 nodes in heap for empty streams, got %d", len(merger.nodes))
		}
	})

	t.Run("CreateMerger with mixed empty and non-empty streams", func(t *testing.T) {
		streams := []Stream{
			NewMockStream([]int{1, 3, 5}),
			&EmptyStream{},
			NewMockStream([]int{2, 4, 6}),
		}

		merger := CreateMerger(streams)
		if merger == nil {
			t.Fatal("CreateMerger returned nil")
		}

		if len(merger.nodes) != 2 {
			t.Errorf("Expected 2 nodes in heap (empty stream should be ignored), got %d", len(merger.nodes))
		}
	})

	t.Run("CreateMerger with no streams", func(t *testing.T) {
		streams := []Stream{}

		merger := CreateMerger(streams)
		if merger == nil {
			t.Fatal("CreateMerger returned nil")
		}

		if len(merger.nodes) != 0 {
			t.Errorf("Expected 0 nodes in heap for no streams, got %d", len(merger.nodes))
		}
	})
}

func TestStreamMergerNext(t *testing.T) {
	t.Run("Merge sorted streams in order", func(t *testing.T) {
		streams := []Stream{
			NewMockStream([]int{1, 4, 7}),
			NewMockStream([]int{2, 5, 8}),
			NewMockStream([]int{3, 6, 9}),
		}

		merger := CreateMerger(streams)
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		for i, expectedVal := range expected {
			val, ok := merger.Next()
			if !ok {
				t.Fatalf("Expected value at position %d, but Next() returned false", i)
			}
			if val != expectedVal {
				t.Errorf("At position %d: expected %d, got %d", i, expectedVal, val)
			}
		}

		// Should return false when all streams are exhausted
		_, ok := merger.Next()
		if ok {
			t.Error("Expected Next() to return false when all streams are exhausted")
		}
	})

	t.Run("Merge streams with duplicates", func(t *testing.T) {
		streams := []Stream{
			NewMockStream([]int{1, 3, 5, 5}),
			NewMockStream([]int{2, 3, 4, 5}),
		}

		merger := CreateMerger(streams)
		expected := []int{1, 2, 3, 3, 4, 5, 5, 5}

		for i, expectedVal := range expected {
			val, ok := merger.Next()
			if !ok {
				t.Fatalf("Expected value at position %d, but Next() returned false", i)
			}
			if val != expectedVal {
				t.Errorf("At position %d: expected %d, got %d", i, expectedVal, val)
			}
		}
	})

	t.Run("Merge single stream", func(t *testing.T) {
		streams := []Stream{
			NewMockStream([]int{10, 20, 30}),
		}

		merger := CreateMerger(streams)
		expected := []int{10, 20, 30}

		for i, expectedVal := range expected {
			val, ok := merger.Next()
			if !ok {
				t.Fatalf("Expected value at position %d, but Next() returned false", i)
			}
			if val != expectedVal {
				t.Errorf("At position %d: expected %d, got %d", i, expectedVal, val)
			}
		}
	})

	t.Run("Merge streams with different lengths", func(t *testing.T) {
		streams := []Stream{
			NewMockStream([]int{1}),
			NewMockStream([]int{2, 4, 6, 8, 10}),
			NewMockStream([]int{3, 5}),
		}

		merger := CreateMerger(streams)
		expected := []int{1, 2, 3, 4, 5, 6, 8, 10}

		for i, expectedVal := range expected {
			val, ok := merger.Next()
			if !ok {
				t.Fatalf("Expected value at position %d, but Next() returned false", i)
			}
			if val != expectedVal {
				t.Errorf("At position %d: expected %d, got %d", i, expectedVal, val)
			}
		}
	})

	t.Run("Next on empty merger", func(t *testing.T) {
		streams := []Stream{}
		merger := CreateMerger(streams)

		val, ok := merger.Next()
		if ok {
			t.Errorf("Expected Next() to return false for empty merger, got true with value %d", val)
		}
	})

	t.Run("Next on merger with only empty streams", func(t *testing.T) {
		streams := []Stream{
			&EmptyStream{},
			&EmptyStream{},
		}
		merger := CreateMerger(streams)

		val, ok := merger.Next()
		if ok {
			t.Errorf("Expected Next() to return false for merger with only empty streams, got true with value %d", val)
		}
	})

	t.Run("Merge streams with negative numbers", func(t *testing.T) {
		streams := []Stream{
			NewMockStream([]int{-5, -1, 3}),
			NewMockStream([]int{-3, 0, 2}),
		}

		merger := CreateMerger(streams)
		expected := []int{-5, -3, -1, 0, 2, 3}

		for i, expectedVal := range expected {
			val, ok := merger.Next()
			if !ok {
				t.Fatalf("Expected value at position %d, but Next() returned false", i)
			}
			if val != expectedVal {
				t.Errorf("At position %d: expected %d, got %d", i, expectedVal, val)
			}
		}
	})
}

func TestStreamMergerEdgeCases(t *testing.T) {
	t.Run("Large number of streams", func(t *testing.T) {
		streams := make([]Stream, 100)
		for i := 0; i < 100; i++ {
			streams[i] = NewMockStream([]int{i * 100, i*100 + 50})
		}

		merger := CreateMerger(streams)
		if merger == nil {
			t.Fatal("CreateMerger returned nil")
		}

		// Just verify we can get some values without error
		count := 0
		for {
			_, ok := merger.Next()
			if !ok {
				break
			}
			count++
			if count > 200 { // Should have exactly 200 values (100 streams * 2 values each)
				t.Error("Got more values than expected")
				break
			}
		}

		if count != 200 {
			t.Errorf("Expected 200 values, got %d", count)
		}
	})

	t.Run("Stream with single value", func(t *testing.T) {
		streams := []Stream{
			NewMockStream([]int{42}),
		}

		merger := CreateMerger(streams)
		val, ok := merger.Next()
		if !ok {
			t.Fatal("Expected to get a value from single-value stream")
		}
		if val != 42 {
			t.Errorf("Expected 42, got %d", val)
		}

		_, ok = merger.Next()
		if ok {
			t.Error("Expected Next() to return false after single value is consumed")
		}
	})

	t.Run("Streams with zero values", func(t *testing.T) {
		streams := []Stream{
			NewMockStream([]int{0, 0, 0}),
			NewMockStream([]int{0, 0}),
		}

		merger := CreateMerger(streams)
		expected := []int{0, 0, 0, 0, 0}

		for i, expectedVal := range expected {
			val, ok := merger.Next()
			if !ok {
				t.Fatalf("Expected value at position %d, but Next() returned false", i)
			}
			if val != expectedVal {
				t.Errorf("At position %d: expected %d, got %d", i, expectedVal, val)
			}
		}
	})
}

func TestStreamInterface(t *testing.T) {
	t.Run("MockStream implements Stream interface", func(t *testing.T) {
		var _ Stream = &MockStream{}
		var _ Stream = &EmptyStream{}
	})

	t.Run("MockStream Next behavior", func(t *testing.T) {
		stream := NewMockStream([]int{1, 2, 3})

		val, ok := stream.Next()
		if !ok || val != 1 {
			t.Errorf("Expected (1, true), got (%d, %t)", val, ok)
		}

		val, ok = stream.Next()
		if !ok || val != 2 {
			t.Errorf("Expected (2, true), got (%d, %t)", val, ok)
		}

		val, ok = stream.Next()
		if !ok || val != 3 {
			t.Errorf("Expected (3, true), got (%d, %t)", val, ok)
		}

		val, ok = stream.Next()
		if ok {
			t.Errorf("Expected (0, false), got (%d, %t)", val, ok)
		}
	})

	t.Run("EmptyStream Next behavior", func(t *testing.T) {
		stream := &EmptyStream{}

		val, ok := stream.Next()
		if ok {
			t.Errorf("Expected (0, false), got (%d, %t)", val, ok)
		}
	})
}

// Benchmark tests
func BenchmarkStreamMerger(b *testing.B) {
	streams := []Stream{
		NewMockStream([]int{1, 4, 7, 10, 13}),
		NewMockStream([]int{2, 5, 8, 11, 14}),
		NewMockStream([]int{3, 6, 9, 12, 15}),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Reset streams for each iteration
		streams = []Stream{
			NewMockStream([]int{1, 4, 7, 10, 13}),
			NewMockStream([]int{2, 5, 8, 11, 14}),
			NewMockStream([]int{3, 6, 9, 12, 15}),
		}

		merger := CreateMerger(streams)
		for {
			_, ok := merger.Next()
			if !ok {
				break
			}
		}
	}
}
