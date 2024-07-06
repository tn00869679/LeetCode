package test

import (
	"LeetCode/MedianofTwoSortedArrays"
	"testing"
)

func TestMedianofTwoSortedArrays(t *testing.T) {
	example1_1 := []int{1, 3}
	example1_2 := []int{2}
	expected1 := 2.0
	example2_1 := []int{1, 2}
	example2_2 := []int{3, 4}
	expected2 := 2.5

	result1 := MedianofTwoSortedArrays.FindMedianSortedArrays(example1_1, example1_2)
	if result1 != expected1 {
		t.Errorf("FindMedianSortedArrays(example1_1, example1_2) = %f; want %f", result1, expected1)
	}

	result2 := MedianofTwoSortedArrays.FindMedianSortedArrays(example2_1, example2_2)
	if result2 != expected2 {
		t.Errorf("FindMedianSortedArrays(example2_1, example2_2) = %f; want %f", result2, expected2)
	}
}
