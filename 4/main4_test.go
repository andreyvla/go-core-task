package main

import (
	"reflect"
	"testing"
)

// TestFindDifference тестирует функцию findDifference.
func TestFindDifference(t *testing.T) {
	testCases := []struct {
		name     string
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			name:     "Example case",
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{
			name:     "No common elements",
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{"d", "e", "f"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "All common elements",
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{"a", "b", "c"},
			expected: []string{},
		},
		{
			name:     "Empty slice1",
			slice1:   []string{},
			slice2:   []string{"a", "b", "c"},
			expected: []string{},
		},
		{
			name:     "Empty slice2",
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Empty slices",
			slice1:   []string{},
			slice2:   []string{},
			expected: []string{},
		},
		{
			name:     "Duplicate in slice1",
			slice1:   []string{"a", "b", "a", "c"},
			slice2:   []string{"b"},
			expected: []string{"a", "a", "c"},
		},
		{
			name:     "Duplicate in slice2",
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{"b", "b"},
			expected: []string{"a", "c"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := findDifference(tc.slice1, tc.slice2)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
