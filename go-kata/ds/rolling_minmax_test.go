package ds

import (
	"reflect"
	"testing"
)

func TestRollingMinMax_BasicFunctionality(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		k            int
		expectedMins []int
		expectedMaxs []int
	}{
		{
			name:         "simple increasing sequence",
			nums:         []int{1, 2, 3, 4, 5},
			k:            3,
			expectedMins: []int{1, 2, 3}, // rolling min with window size 3
			expectedMaxs: []int{3, 4, 5}, // rolling max with window size 3
		},
		// {
		// 	name:         "simple decreasing sequence",
		// 	nums:         []int{5, 4, 3, 2, 1},
		// 	k:            3,
		// 	expectedMins: []int{3, 2, 1}, // rolling min with window size 3
		// 	expectedMaxs: []int{5, 4, 3}, // rolling max with window size 3
		// },
		// {
		// 	name:         "mixed sequence",
		// 	nums:         []int{1, 3, 2, 5, 4},
		// 	k:            3,
		// 	expectedMins: []int{1, 2, 2}, // rolling min with window size 3
		// 	expectedMaxs: []int{3, 5, 5}, // rolling max with window size 3
		// },
		// {
		// 	name:         "window size 1",
		// 	nums:         []int{1, 2, 3, 4, 5},
		// 	k:            1,
		// 	expectedMins: []int{1, 2, 3, 4, 5}, // each element is its own min
		// 	expectedMaxs: []int{1, 2, 3, 4, 5}, // each element is its own max
		// },
		// {
		// 	name:         "window size equals array length",
		// 	nums:         []int{3, 1, 4, 1, 5},
		// 	k:            5,
		// 	expectedMins: []int{1}, // only one window, min is 1
		// 	expectedMaxs: []int{5}, // only one window, max is 5
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mins, maxs, err := RollingMinMax(tt.nums, tt.k)
			if err != nil {
				t.Errorf("RollingMinMax(%v, %d) returned error: %v", tt.nums, tt.k, err)
				return
			}
			if !reflect.DeepEqual(mins, tt.expectedMins) {
				t.Errorf("RollingMinMax(%v, %d) mins = %v, want %v", tt.nums, tt.k, mins, tt.expectedMins)
			}
			if !reflect.DeepEqual(maxs, tt.expectedMaxs) {
				t.Errorf("RollingMinMax(%v, %d) maxs = %v, want %v", tt.nums, tt.k, maxs, tt.expectedMaxs)
			}
		})
	}
}

func TestRollingMinMax_EdgeCases(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		k            int
		expectedMins []int
		expectedMaxs []int
	}{
		{
			name:         "empty array",
			nums:         []int{},
			k:            3,
			expectedMins: []int{},
			expectedMaxs: []int{},
		},
		{
			name:         "single element",
			nums:         []int{42},
			k:            1,
			expectedMins: []int{42},
			expectedMaxs: []int{42},
		},
		{
			name:         "single element with k > 1",
			nums:         []int{42},
			k:            3,
			expectedMins: []int{42},
			expectedMaxs: []int{42},
		},
		{
			name:         "duplicate elements",
			nums:         []int{2, 2, 2, 2, 2},
			k:            3,
			expectedMins: []int{2, 2, 2},
			expectedMaxs: []int{2, 2, 2},
		},
		{
			name:         "all same elements",
			nums:         []int{5, 5, 5, 5},
			k:            2,
			expectedMins: []int{5, 5, 5},
			expectedMaxs: []int{5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mins, maxs, err := RollingMinMax(tt.nums, tt.k)
			if err != nil {
				t.Errorf("RollingMinMax(%v, %d) returned error: %v", tt.nums, tt.k, err)
				return
			}
			if !reflect.DeepEqual(mins, tt.expectedMins) {
				t.Errorf("RollingMinMax(%v, %d) mins = %v, want %v", tt.nums, tt.k, mins, tt.expectedMins)
			}
			if !reflect.DeepEqual(maxs, tt.expectedMaxs) {
				t.Errorf("RollingMinMax(%v, %d) maxs = %v, want %v", tt.nums, tt.k, maxs, tt.expectedMaxs)
			}
		})
	}
}

func TestRollingMinMax_NegativeNumbers(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		k            int
		expectedMins []int
		expectedMaxs []int
	}{
		{
			name:         "all negative numbers",
			nums:         []int{-1, -3, -2, -5, -4},
			k:            3,
			expectedMins: []int{-3, -5, -5}, // rolling min of negative numbers
			expectedMaxs: []int{-1, -2, -2}, // rolling max of negative numbers
		},
		{
			name:         "mixed positive and negative",
			nums:         []int{-1, 2, -3, 4, -5},
			k:            3,
			expectedMins: []int{-3, -3, -5}, // rolling min with mixed signs
			expectedMaxs: []int{2, 4, 4},    // rolling max with mixed signs
		},
		{
			name:         "negative to positive transition",
			nums:         []int{-5, -2, 1, 3, -1},
			k:            3,
			expectedMins: []int{-5, -2, -1}, // transition from negative to positive
			expectedMaxs: []int{1, 3, 3},    // transition from negative to positive
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mins, maxs, err := RollingMinMax(tt.nums, tt.k)
			if err != nil {
				t.Errorf("RollingMinMax(%v, %d) returned error: %v", tt.nums, tt.k, err)
				return
			}
			if !reflect.DeepEqual(mins, tt.expectedMins) {
				t.Errorf("RollingMinMax(%v, %d) mins = %v, want %v", tt.nums, tt.k, mins, tt.expectedMins)
			}
			if !reflect.DeepEqual(maxs, tt.expectedMaxs) {
				t.Errorf("RollingMinMax(%v, %d) maxs = %v, want %v", tt.nums, tt.k, maxs, tt.expectedMaxs)
			}
		})
	}
}

func TestRollingMinMax_LargerWindows(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		k            int
		expectedMins []int
		expectedMaxs []int
	}{
		{
			name:         "leetcode example",
			nums:         []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:            3,
			expectedMins: []int{-1, -3, -3, -3, 3, 3}, // classic sliding window minimum
			expectedMaxs: []int{3, 3, 5, 5, 6, 7},     // classic sliding window maximum
		},
		{
			name:         "mountain pattern",
			nums:         []int{1, 2, 3, 4, 3, 2, 1},
			k:            4,
			expectedMins: []int{1, 2, 2, 1}, // valley in the middle
			expectedMaxs: []int{4, 4, 4, 4}, // peak in the middle
		},
		{
			name:         "valley pattern",
			nums:         []int{5, 4, 3, 2, 3, 4, 5},
			k:            4,
			expectedMins: []int{2, 2, 2, 2}, // valley in the middle
			expectedMaxs: []int{5, 4, 4, 5}, // peaks at edges
		},
		{
			name:         "large window",
			nums:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			k:            5,
			expectedMins: []int{1, 2, 3, 4, 5, 6},  // sliding window of size 5
			expectedMaxs: []int{5, 6, 7, 8, 9, 10}, // sliding window of size 5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mins, maxs, err := RollingMinMax(tt.nums, tt.k)
			if err != nil {
				t.Errorf("RollingMinMax(%v, %d) returned error: %v", tt.nums, tt.k, err)
				return
			}
			if !reflect.DeepEqual(mins, tt.expectedMins) {
				t.Errorf("RollingMinMax(%v, %d) mins = %v, want %v", tt.nums, tt.k, mins, tt.expectedMins)
			}
			if !reflect.DeepEqual(maxs, tt.expectedMaxs) {
				t.Errorf("RollingMinMax(%v, %d) maxs = %v, want %v", tt.nums, tt.k, maxs, tt.expectedMaxs)
			}
		})
	}
}

func TestRollingMinMax_SpecialPatterns(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		k            int
		expectedMins []int
		expectedMaxs []int
	}{
		{
			name:         "alternating high low",
			nums:         []int{10, 1, 10, 1, 10, 1},
			k:            2,
			expectedMins: []int{1, 1, 1, 1, 1},      // alternating pattern
			expectedMaxs: []int{10, 10, 10, 10, 10}, // alternating pattern
		},
		{
			name:         "peak at start",
			nums:         []int{10, 5, 3, 2, 1},
			k:            3,
			expectedMins: []int{3, 2, 1},  // minimum decreasing
			expectedMaxs: []int{10, 5, 3}, // maximum at the beginning
		},
		{
			name:         "peak at end",
			nums:         []int{1, 2, 3, 5, 10},
			k:            3,
			expectedMins: []int{1, 2, 3},  // minimum at the beginning
			expectedMaxs: []int{3, 5, 10}, // maximum at the end
		},
		{
			name:         "multiple peaks",
			nums:         []int{8, 2, 9, 1, 7, 3, 6},
			k:            3,
			expectedMins: []int{2, 1, 1, 1, 3}, // multiple local minima
			expectedMaxs: []int{9, 9, 9, 7, 7}, // multiple local maxima
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mins, maxs, err := RollingMinMax(tt.nums, tt.k)
			if err != nil {
				t.Errorf("RollingMinMax(%v, %d) returned error: %v", tt.nums, tt.k, err)
				return
			}
			if !reflect.DeepEqual(mins, tt.expectedMins) {
				t.Errorf("RollingMinMax(%v, %d) mins = %v, want %v", tt.nums, tt.k, mins, tt.expectedMins)
			}
			if !reflect.DeepEqual(maxs, tt.expectedMaxs) {
				t.Errorf("RollingMinMax(%v, %d) maxs = %v, want %v", tt.nums, tt.k, maxs, tt.expectedMaxs)
			}
		})
	}
}

func TestRollingMinMax_LargeK(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		k            int
		expectedMins []int
		expectedMaxs []int
	}{
		{
			name:         "k larger than array",
			nums:         []int{1, 2, 3},
			k:            5,
			expectedMins: []int{1}, // k > len(nums), should return min of entire array
			expectedMaxs: []int{3}, // k > len(nums), should return max of entire array
		},
		{
			name:         "k equals array length",
			nums:         []int{4, 2, 7, 1, 9},
			k:            5,
			expectedMins: []int{1}, // exactly one window
			expectedMaxs: []int{9}, // exactly one window
		},
		{
			name:         "k is array length minus 1",
			nums:         []int{3, 1, 4, 1, 5, 9},
			k:            5,
			expectedMins: []int{1, 1}, // two overlapping windows
			expectedMaxs: []int{5, 9}, // two overlapping windows
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mins, maxs, err := RollingMinMax(tt.nums, tt.k)
			if err != nil {
				t.Errorf("RollingMinMax(%v, %d) returned error: %v", tt.nums, tt.k, err)
				return
			}
			if !reflect.DeepEqual(mins, tt.expectedMins) {
				t.Errorf("RollingMinMax(%v, %d) mins = %v, want %v", tt.nums, tt.k, mins, tt.expectedMins)
			}
			if !reflect.DeepEqual(maxs, tt.expectedMaxs) {
				t.Errorf("RollingMinMax(%v, %d) maxs = %v, want %v", tt.nums, tt.k, maxs, tt.expectedMaxs)
			}
		})
	}
}
