package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "",
			s:        "racecar",
			t:        "carrace",
			expected: true,
		},
		{
			name:     "",
			s:        "jar",
			t:        "jam",
			expected: false,
		},
		{
			name:     "",
			s:        "aaaa",
			t:        "aaaa",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("isAnagram(%v, %v) = %v, expected %v", tt.s, tt.t, result, tt.expected)
			}
		})
	}
}
