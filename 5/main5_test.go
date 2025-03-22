package main

import (
	"reflect"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	testCases := []struct {
		name          string
		slice1        []int
		slice2        []int
		expectedBool  bool
		expectedSlice []int
	}{
		{
			name:          "Пересечение есть",
			slice1:        []int{65, 3, 58, 678, 64},
			slice2:        []int{64, 2, 3, 43},
			expectedBool:  true,
			expectedSlice: []int{64, 3},
		},
		{
			name:          "Пересечения нет",
			slice1:        []int{1, 2, 3},
			slice2:        []int{4, 5, 6},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "Пересечение с дубликатами",
			slice1:        []int{1, 2, 3, 3, 2, 1},
			slice2:        []int{3, 2, 1, 1, 2, 3},
			expectedBool:  true,
			expectedSlice: []int{3, 2, 1},
		},
		{
			name:          "Пустые слайсы",
			slice1:        []int{},
			slice2:        []int{},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "Один пустой слайс",
			slice1:        []int{1, 2, 3},
			slice2:        []int{},
			expectedBool:  false,
			expectedSlice: []int{},
		},
		{
			name:          "Другой пустой слайс",
			slice1:        []int{},
			slice2:        []int{1, 2, 3},
			expectedBool:  false,
			expectedSlice: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualBool, actualSlice := findIntersection(tc.slice1, tc.slice2)
			if actualBool != tc.expectedBool {
				t.Errorf("Ожидалось пересечение: %v, имеем: %v", tc.expectedBool, actualBool)
			}
			if !reflect.DeepEqual(actualSlice, tc.expectedSlice) {
				t.Errorf("Ожидался слайс пересечений: %v, имеем: %v", tc.expectedSlice, actualSlice)
			}
		})
	}
}
