package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "",
			nums:     []int{3, 4, 5, 6},
			target:   7,
			expected: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			if result[0] != tt.expected[0] || result[1] != tt.expected[1] {
				t.Errorf("twoSum(%v, %v) = %v, expected %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
