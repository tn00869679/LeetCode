package test

import (
	"LeetCode/SortColors"
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	example1 := []int{2, 0, 2, 1, 1, 0}
	expected1 := []int{0, 0, 1, 1, 2, 2}
	example2 := []int{2, 0, 1}
	expected2 := []int{0, 1, 2}

	SortColors.SortColors(example1)
	if !reflect.DeepEqual(example1, expected1) {
		t.Errorf("SortColors(example1) = %v; want %v", example1, expected1)
	}

	SortColors.SortColors(example2)
	if !reflect.DeepEqual(example2, expected2) {
		t.Errorf("SortColors(example2) = %v; want %v", example2, expected2)
	}
}
