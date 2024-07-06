package test

import (
	"LeetCode/TwoSum"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}
	expectedStr := TwoSum.GetString(expected)

	result := TwoSum.TwoSum(nums, target)
	resultStr := TwoSum.GetString(result)

	if result[0] != expected[0] {
		t.Errorf("TwoSum(nums, target) = %s; want %s", resultStr, expectedStr)
	}
}
