package main

import "testing"

func TestHasDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "empty slice",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: false,
		},
		{
			name:     "two same elements",
			nums:     []int{1, 1},
			expected: true,
		},
		{
			name:     "no duplicates",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "with duplicates",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "duplicates at beginning",
			nums:     []int{1, 1, 2, 3},
			expected: true,
		},
		{
			name:     "negative numbers with duplicates",
			nums:     []int{-1, -2, -1, 3},
			expected: true,
		},
		{
			name:     "negative numbers no duplicates",
			nums:     []int{-1, -2, -3},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("hasDuplicate(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}
