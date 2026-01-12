# Go Concepts Study Guide

This summary covers all the Go concepts, patterns, and techniques used in the go-kata project.

## Table of Contents
1. [Basic Go Syntax & Types](#basic-go-syntax--types)
2. [Data Structures](#data-structures)
3. [Interfaces](#interfaces)
4. [Error Handling](#error-handling)
5. [Memory Management & Pointers](#memory-management--pointers)
6. [Standard Library Usage](#standard-library-usage)
7. [Testing](#testing)
8. [Advanced Patterns](#advanced-patterns)
9. [Algorithms & Problem Solving](#algorithms--problem-solving)

---

## Basic Go Syntax & Types

### Package Declaration & Imports
```go
package ds                    // Package declaration
import (
    "container/heap"         // Standard library import
    "time"                   // Multiple imports in parentheses
    "errors"
)
```

### Variable Declarations
```go
var current_char rune        // var keyword with type
count := 0                   // Short variable declaration
const capacity = 10          // Constants
```

### Basic Types
- `int`, `uint64` - Integer types
- `string` - String type
- `bool` - Boolean type
- `rune` - Unicode character (alias for int32)
- `time.Time` - Time type from standard library

### Control Structures
```go
// For loops
for i := 0; i < len(nums); i++ {
    // Range-based for loop
    for idx, num := range nums {
        // Infinite loop with break
        for {
            if condition {
                break
            }
        }
    }
}

// If statements with initialization
if val, ok := stream.Next(); ok {
    // Use val
}

// Switch statements
switch c {
case '(', '{', '[':
    push(c)
case ')', '}', ']':
    // Handle closing brackets
}
```

---

## Data Structures

### Slices
```go
// Slice creation and manipulation
data := make([]int, 0)           // Empty slice with zero length
buf := make([]int, capacity)     // Slice with specific capacity
slice := nums[0:n-1]             // Slice operations
slice = append(slice, element)   // Appending elements
```

### Maps
```go
// Map creation and usage
byID := make(map[string]*Task, 0)    // Map with initial capacity
counts := make(map[int]int)          // Simple map
seen := make(map[string]struct{})    // Using empty struct as set

// Map operations
if task, ok := byID[ID]; ok {        // Check if key exists
    // Key exists, use task
}
delete(byID, ID)                     // Delete from map
```

### Arrays vs Slices
- Arrays have fixed size: `[5]int`
- Slices are dynamic: `[]int`
- Slices are references to underlying arrays

### Custom Data Structures

#### Heap Implementation
```go
type TaskHeap []*Task

// Implementing heap.Interface
func (h TaskHeap) Len() int           { return len(h) }
func (h TaskHeap) Less(i, j int) bool { return h[i].deadline.Before(h[j].deadline) }
func (h TaskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TaskHeap) Push(x any)        { *h = append(*h, x.(*Task)) }
func (h *TaskHeap) Pop() any          { /* implementation */ }
```

#### Ring Buffer/Circular Buffer
```go
type Ring struct {
    buf      []int
    capacity int
    head     uint64    // Read pointer
    tail     uint64    // Write pointer
}

// Modular arithmetic for circular indexing
func (r *Ring) physicalidx(index uint64) int {
    return int(index % uint64(r.capacity))
}
```

#### Double-Ended Queue (Deque)
```go
type Deque struct {
    buf      []int
    capacity int
    front    uint64
    back     uint64
}

// Operations on both ends
func (q *Deque) PushFront(v int) error  { /* implementation */ }
func (q *Deque) PushBack(v int) error   { /* implementation */ }
func (q *Deque) PopFront() (int, error) { /* implementation */ }
func (q *Deque) PopBack() (int, error)  { /* implementation */ }
```

---

## Interfaces

### Interface Definition
```go
type Stream interface {
    Next() (int, bool)    // Method signature
}
```

### Interface Implementation
- Implicit implementation (no `implements` keyword)
- Any type that has the required methods implements the interface

### Empty Interface
```go
func (h *TaskHeap) Push(x any) {    // any is alias for interface{}
    *h = append(*h, x.(*Task))      // Type assertion
}
```

### Type Assertions
```go
task := x.(*Task)                   // Type assertion with panic on failure
if val, ok := x.(int); ok {         // Safe type assertion
    // Use val
}
```

---

## Error Handling

### Error Type
```go
import "errors"

// Creating errors
return errors.New("Heap is empty!!")
return 0, errors.New("Ring Empty!!!")
```

### Error Handling Pattern
```go
func (h *IntHeap) PopInt() (int, error) {
    if len(h.data) == 0 {
        return 0, errors.New("Heap is empty!!")  // Return zero value + error
    }
    return heap.Pop(h).(int), nil                // Return value + nil error
}

// Usage
if val, err := h.PopInt(); err != nil {
    // Handle error
} else {
    // Use val
}
```

### Multiple Return Values
```go
func Next() (int, bool)              // Value and success indicator
func PopFront() (int, error)         // Value and error
func TwoSum(nums []int, target int) (int, int, bool)  // Multiple values
```

---

## Memory Management & Pointers

### Pointer Usage
```go
type Scheduler struct {
    h    TaskHeap
    byID map[string]*Task    // Pointer to Task
}

// Creating pointers
task := &Task{              // Address-of operator
    ID:       ID,
    deadline: deadline,
}

// Dereferencing
last_merged_interval := &merged_intervals[len(merged_intervals)-1]
last_merged_interval.End = new_end    // Automatic dereferencing
```

### Method Receivers
```go
// Value receiver (copy)
func (h TaskHeap) Len() int {
    return len(h)
}

// Pointer receiver (reference)
func (h *TaskHeap) Push(x any) {
    *h = append(*h, x.(*Task))
}
```

### Memory Allocation
```go
make([]int, capacity)           // Allocate slice
make(map[string]int)           // Allocate map
&Scheduler{...}                // Allocate struct on heap
```

---

## Standard Library Usage

### Container/Heap
```go
import "container/heap"

heap.Init(&h)                  // Initialize heap
heap.Push(&h, item)           // Push to heap
item := heap.Pop(&h)          // Pop from heap
heap.Fix(&h, index)           // Fix heap after modification
heap.Remove(&h, index)        // Remove element at index
```

### Time Package
```go
import "time"

now := time.Now()                    // Current time
deadline := now.Add(time.Hour)      // Add duration
t.deadline.Before(other.deadline)   // Time comparison
t.deadline.After(deadline)          // Time comparison
```

### Strings Package
```go
import "strings"

var buf strings.Builder             // Efficient string building
buf.WriteRune(c)                   // Write rune to builder
buf.WriteString(s)                 // Write string to builder
buf.String()                       // Get final string
buf.Reset()                        // Reset builder

strings.ToLower(s)                 // Convert to lowercase
strings.TrimSpace(s)               // Trim whitespace
strings.Split(s, ":")              // Split string
```

### Sort Package
```go
import "sort"

// Sorting slices with custom comparison
sort.Slice(runes, func(i, j int) bool {
    return runes[i] < runes[j]
})

sort.Slice(intervals, func(i, j int) bool {
    if intervals[i].Start == intervals[j].Start {
        return intervals[i].End < intervals[j].End
    }
    return intervals[i].Start < intervals[j].Start
})
```

### Strconv Package
```go
import "strconv"

money, err := strconv.Atoi(parts[1])    // String to integer
buf.WriteString(strconv.Itoa(count))    // Integer to string
```

### Unicode Package
```go
import "unicode"

if unicode.IsLetter(c) {    // Check if character is letter
    // Process letter
}
```

---

## Testing

### Test Function Structure
```go
func TestFunctionName(t *testing.T) {
    // Arrange
    input := "test input"
    expected := "expected output"
    
    // Act
    result := FunctionToTest(input)
    
    // Assert
    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}
```

### Testing Patterns
```go
// Fatal vs Error
t.Fatal("Stops test execution")
t.Error("Continues test execution")

// Formatted messages
t.Errorf("Expected %d, got %d", expected, actual)
t.Fatalf("Critical error: %v", err)

// Checking conditions
if result == nil {
    t.Fatal("Function returned nil")
}

// Using reflect.DeepEqual for complex comparisons
if !reflect.DeepEqual(got, want) {
    t.Fatalf("got=%v want=%v", got, want)
}
```

---

## Advanced Patterns

### Functional Programming Elements
```go
// Function as variable
less := func(a, b int) bool { return a < b }

// Closures
purge := func(Q *Deque, curr_idx int) {
    // Accesses variables from outer scope
    for !Q.Empty() {
        // Implementation using outer variables
    }
}

// Higher-order functions
monotonic_push := func(Q *Deque, new_idx int, pop_cond func(old int, new int) bool) {
    // Function that takes another function as parameter
}
```

### Method Chaining & Fluent Interface
```go
// Chaining operations
buf.WriteRune(current_char)
if count > 1 {
    buf.WriteString(strconv.Itoa(count))
}
```

### Empty Struct as Set
```go
seen := make(map[string]struct{})    // Using empty struct to save memory
seen[key] = struct{}{}               // Adding to set
if _, ok := seen[key]; ok {          // Checking membership
    // Key exists in set
}
```

### Struct Embedding & Composition
```go
type Task struct {
    ID       string
    deadline time.Time
    index    int        // Additional field for heap operations
}
```

---

## Algorithms & Problem Solving

### Two Pointers Technique
```go
func TwoSum(nums []int, target int) (int, int, bool) {
    i := 0
    j := len(nums) - 1
    
    for i < j {
        sum := nums[i] + nums[j]
        if sum == target {
            return i, j, true
        } else if sum < target {
            i++
        } else {
            j--
        }
    }
    return 0, 0, false
}
```

### Sliding Window Pattern
```go
// Rolling min/max with deque
func RollingMinMax(nums []int, k int) (mins []int, maxs []int, err error) {
    maxQ := NewDeque(k)    // Monotonic deque for maximum
    minQ := NewDeque(k)    // Monotonic deque for minimum
    
    for curr_idx, num := range nums {
        // Maintain window size
        purge(maxQ, curr_idx)
        purge(minQ, curr_idx)
        
        // Maintain monotonic property
        monotonic_push(maxQ, curr_idx, func(old, new int) bool { 
            return old > new 
        })
    }
}
```

### Heap-based Algorithms
```go
// K-way merge using min-heap
func (m *StreamMerger) Next() (int, bool) {
    if len(m.nodes) == 0 {
        return 0, false
    }
    
    node := heap.Pop(&m.nodes).(HeapNode)
    if new_value, ok := node.stream.Next(); ok {
        heap.Push(&m.nodes, HeapNode{val: new_value, stream: node.stream})
    }
    
    return node.val, true
}
```

### String Processing Algorithms
```go
// Run Length Encoding
func RLE(s string) string {
    var buf strings.Builder
    var current_char rune
    count := 0
    
    flush := func() {
        buf.WriteRune(current_char)
        if count > 1 {
            buf.WriteString(strconv.Itoa(count))
        }
    }
    
    // Process each character
    for idx, c := range s {
        if idx == 0 || current_char == c {
            current_char = c
            count++
        } else {
            flush()
            current_char = c
            count = 1
        }
    }
    flush()
    return buf.String()
}
```

### Stack-based Algorithms
```go
// Balanced parentheses using slice as stack
func IsBalanced(s string) bool {
    stack := make([]rune, 0, len(s))
    
    push := func(c rune) {
        stack = append(stack, c)
    }
    
    pop := func() (rune, bool) {
        if len(stack) == 0 {
            return 0, false
        }
        last := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        return last, true
    }
    
    match := map[rune]rune{')': '(', '}': '{', ']': '['}
    
    for _, c := range s {
        switch c {
        case '(', '{', '[':
            push(c)
        case ')', '}', ']':
            if popped, ok := pop(); !ok || popped != match[c] {
                return false
            }
        }
    }
    
    return len(stack) == 0
}
```

---

## Key Go Idioms & Best Practices

### Error Handling
- Always handle errors explicitly
- Return errors as the last return value
- Use meaningful error messages
- Check errors immediately after function calls

### Memory Efficiency
- Use empty struct `struct{}` for sets
- Pre-allocate slices when size is known: `make([]int, 0, capacity)`
- Use pointer receivers for methods that modify the receiver
- Use value receivers for small structs or when you need a copy

### Interface Design
- Keep interfaces small and focused
- Define interfaces where they're used, not where they're implemented
- Use `io.Reader`, `io.Writer` and other standard interfaces when possible

### Naming Conventions
- Use camelCase for unexported names
- Use PascalCase for exported names
- Use short, descriptive names
- Avoid stuttering: `heap.HeapNode` → `heap.Node`

### Concurrency Readiness
- Design with goroutines in mind
- Use channels for communication
- Avoid shared mutable state when possible
- Use sync package for synchronization when needed

This study guide covers the essential Go concepts demonstrated in your go-kata project, from basic syntax to advanced algorithmic patterns. Each concept is illustrated with actual code examples from your codebase.
