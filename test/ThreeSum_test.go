package test

import (
	"LeetCode/ThreeSum"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	example1 := []int{-1, 0, 1, 2, -1, -4}
	expected1 := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	example2 := []int{0, 1, 1}
	expected2 := [][]int{}
	example3 := []int{0, 0, 0}
	expected3 := [][]int{{0, 0, 0}}

	result1 := ThreeSum.ThreeSum(example1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("ThreeSum(example1) = %v; want %v", result1, expected1)
	}

	result2 := ThreeSum.ThreeSum(example2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("ThreeSum(example2) = %v; want %v", result2, expected2)
	}

	result3 := ThreeSum.ThreeSum(example3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("ThreeSum(example3) = %v; want %v", result3, expected3)
	}
}
